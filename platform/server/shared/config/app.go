/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"time"

	"github.com/cloudwego/cwgo/platform/server/shared/consts"
)

type AppConfig struct {
	RunMode                  string `mapstructure:"runMode"`
	Timezone                 string `mapstructure:"timezone"`
	ProxyUrl                 string `mapstructure:"proxyUrl"`
	SyncAgentServiceInterval string `mapstructure:"syncAgentServiceInterval"`
	SyncIdlInterval          string `mapstructure:"syncIdlInterval"`
}

const (
	defaultSyncAgentServiceInterval = 10 * time.Second
	defaultSyncIdlInterval          = 3 * time.Minute
)

func (conf *AppConfig) Init() {
	conf.RunMode = consts.ServerRunModeCluster
}

func (conf *AppConfig) GetSyncAgentServiceInterval() time.Duration {
	duration, err := time.ParseDuration(conf.SyncAgentServiceInterval)
	if err == nil {
		return duration
	}

	return defaultSyncAgentServiceInterval
}

func (conf *AppConfig) GetSyncIdlInterval() time.Duration {
	duration, err := time.ParseDuration(conf.SyncIdlInterval)
	if err == nil {
		return duration
	}

	return defaultSyncIdlInterval
}
