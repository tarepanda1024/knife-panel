// Copyright 2020 liuxiaodong Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package initialize

import (
	_ "github.com/mattn/go-sqlite3"
	"knife-panel/server/core/http/model"
	"knife-panel/server/global"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() {
	conf := global.AppConfig
	var conn string
	var dialect gorm.Dialector
	if "sqlite3" == conf.System.DbType {
		conn = conf.Sqlite.Path
		dialect = sqlite.Open(conn)
	}

	if dialect == nil {
		panic("unSupport db type")
	}

	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&model.DBProp{})
	if err != nil {
		panic("migrate failed.")
	}
	global.DB = db
}
