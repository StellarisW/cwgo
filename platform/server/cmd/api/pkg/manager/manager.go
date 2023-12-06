/*
 *
 *  * Copyright 2022 CloudWeGo Authors
 *  *
 *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  * you may not use this file except in compliance with the License.
 *  * You may obtain a copy of the License at
 *  *
 *  *     http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 *
 */

package manager

import (
	"context"
	"fmt"
	"github.com/cloudwego/cwgo/platform/server/cmd/api/pkg/dispatcher"
	"github.com/cloudwego/cwgo/platform/server/shared/config/app"
	"github.com/cloudwego/cwgo/platform/server/shared/consts"
	"github.com/cloudwego/cwgo/platform/server/shared/dao"
	"github.com/cloudwego/cwgo/platform/server/shared/kitex_gen/agent"
	"github.com/cloudwego/cwgo/platform/server/shared/kitex_gen/agent/agentservice"
	"github.com/cloudwego/cwgo/platform/server/shared/kitex_gen/model"
	"github.com/cloudwego/cwgo/platform/server/shared/logger"
	"github.com/cloudwego/cwgo/platform/server/shared/registry"
	"github.com/cloudwego/cwgo/platform/server/shared/service"
	"github.com/cloudwego/cwgo/platform/server/shared/task"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"go.uber.org/zap"
	"sync"
	"time"
)

type Manager struct {
	sync.Mutex
	updateTaskInterval    time.Duration
	currentUpdateTaskTime time.Time
	lastUpdateTaskTime    time.Time

	agents []*service.Service

	syncAgentServiceInterval time.Duration
	syncRepositoryInterval   time.Duration
	syncIdlInterval          time.Duration

	daoManager *dao.Manager
	dispatcher dispatcher.IDispatcher
	registry   registry.IRegistry
	resolver   discovery.Resolver
}

func NewManager(appConf app.Config, daoManager *dao.Manager, dispatcher dispatcher.IDispatcher, registry registry.IRegistry, resolver discovery.Resolver) *Manager {
	manager := &Manager{
		Mutex:                 sync.Mutex{},
		updateTaskInterval:    3 * time.Second,
		currentUpdateTaskTime: time.Time{},
		lastUpdateTaskTime:    time.Now(),

		agents: make([]*service.Service, 0),

		syncAgentServiceInterval: appConf.GetSyncAgentServiceInterval(),
		syncRepositoryInterval:   appConf.GetSyncRepositoryInterval(),
		syncIdlInterval:          appConf.GetSyncIdlInterval(),

		daoManager: daoManager,
		dispatcher: dispatcher,
		registry:   registry,
		resolver:   resolver,
	}

	page := 1
	for {
		idlModels, total, err := daoManager.Idl.GetIDLList(context.Background(), int32(page), 1000, consts.OrderNumDec, "update_time")
		if err != nil {
			panic(fmt.Sprintf("get idl list failed, err: %v", err))
		}
		for _, idlModel := range idlModels {
			err = manager.AddTask(
				task.NewTask(
					model.Type_sync_idl_data,
					manager.syncRepositoryInterval.String(),
					&model.Data{
						SyncIdlData: &model.SyncIdlData{
							IdlId: idlModel.Id,
						},
					},
				))
			if err != nil {
				panic(err)
			}
		}
		if int64(page)*1000 >= total {
			break
		}
		page++
	}

	go manager.StartUpdate()

	return manager
}

func (m *Manager) GetAgentClient() (agentservice.Client, error) {
	c, err := agentservice.NewClient(
		consts.ServiceNameAgent,
		client.WithResolver(m.resolver),
	)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (m *Manager) GetAgentClientByServiceId(serviceId string) (agentservice.Client, error) {
	c, err := agentservice.NewClient(
		consts.ServiceNameAgent,
		client.WithResolver(m.resolver),
		client.WithTag("service_id", serviceId),
	)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (m *Manager) AddTask(t *model.Task) error {
	err := m.dispatcher.AddTask(t)
	if err != nil {
		return fmt.Errorf("add task to dispatcher failed, err: %v", err)
	}

	m.Lock()
	m.currentUpdateTaskTime = time.Now()
	m.Unlock()

	return nil
}

func (m *Manager) UpdateAgentTasks() {
	var wg sync.WaitGroup
	for _, svr := range m.agents {
		wg.Add(1)
		go func(serviceId string) {
			defer wg.Done()

			c, err := m.GetAgentClientByServiceId(serviceId)
			if err != nil {
				logger.Logger.Error("get agent client failed", zap.Error(err))
			}

			tasks := m.dispatcher.GetTaskByServiceId(serviceId)

			tasksModels := make([]*model.Task, len(tasks))
			for i, t := range tasks {
				tasksModels[i] = &model.Task{
					Id:           t.Id,
					Type:         t.Type,
					ScheduleTime: t.ScheduleTime,
					Data:         t.Data,
				}
			}

			rpcRes, err := c.UpdateTasks(context.Background(), &agent.UpdateTasksReq{Tasks: tasksModels})
			if err != nil {
				logger.Logger.Error("update tasks to rpc client failed", zap.Error(err))
			} else if rpcRes.Code != 0 {
				logger.Logger.Error("update tasks failed", zap.String("err", rpcRes.Msg))
			}
		}(svr.Id)
	}
	wg.Wait()
}

func (m *Manager) StartUpdate() {
	go func() {
		for {
			time.Sleep(m.syncAgentServiceInterval)
			m.SyncService()
		}
	}()

	go func() {
		for {
			time.Sleep(100 * time.Millisecond)

			m.Lock()
			if m.lastUpdateTaskTime != m.currentUpdateTaskTime {
				if m.currentUpdateTaskTime.Add(m.updateTaskInterval).After(time.Now()) {
					m.UpdateAgentTasks()
					m.lastUpdateTaskTime = m.currentUpdateTaskTime
				}
			}
			m.Unlock()
		}
	}()
}

func (m *Manager) SyncService() {
	services, err := m.registry.GetAllService()
	if err != nil {
		return
	}

	seta := make(map[string]struct{})
	setb := make(map[string]struct{})
	var addServiceIds, delServicesIds []string

	for _, svr := range m.agents {
		seta[svr.Id] = struct{}{}
	}
	for _, svr := range services {
		setb[svr.Id] = struct{}{}
	}

	for serviceId := range seta {
		if _, ok := setb[serviceId]; !ok {
			delServicesIds = append(delServicesIds, serviceId)
		}
	}

	for serviceId := range setb {
		if _, ok := seta[serviceId]; !ok {
			addServiceIds = append(addServiceIds, serviceId)
		}
	}

	for _, serviceId := range addServiceIds {
		err = m.dispatcher.AddService(serviceId)
		if err != nil {
			continue
		}
	}

	for _, serviceId := range delServicesIds {
		err = m.dispatcher.DelService(serviceId)
		if err != nil {
			continue
		}
	}

	m.agents = services

	if len(addServiceIds) != 0 || len(delServicesIds) != 0 {
		// service changed, update cron
		m.UpdateAgentTasks()
	}
}

func (m *Manager) GetDispatcher() dispatcher.IDispatcher {
	return m.dispatcher
}

func (m *Manager) GetRegistry() registry.IRegistry {
	return m.registry
}
