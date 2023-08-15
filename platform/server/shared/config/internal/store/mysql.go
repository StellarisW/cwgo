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

package store

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mysqlDb *gorm.DB
)

func InitMysqlDB(dsn string) error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	mysqlDb = db

	return nil
}

func GetMysqlDB(dsn ...string) (*gorm.DB, error) {
	if dsn != nil {
		err := InitMysqlDB(dsn[0])
		if err != nil {
			return nil, err
		}
	}
	if mysqlDb != nil {
		return mysqlDb, nil
	}
	return nil, errors.New("mysql db not initialized")
}
