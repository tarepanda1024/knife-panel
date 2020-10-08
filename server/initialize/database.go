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
	"knife-panel/server/global"
	"os"
	"time"
	"xorm.io/xorm"
)

func InitDatabase() {
	conf := global.AppConfig
	var conn string
	if "sqlite3" == conf.System.DbType {
		conn = conf.Sqlite.Path
	}
	engine, err := xorm.NewEngine(conf.System.DbType, conn)

	if err != nil {
		global.Log.Fatalf("dao connect failed: %#v\n", err.Error())
		os.Exit(-1)
	}
	if engine.Ping() != nil {
		global.Log.Fatalf("dao connect error: %#v\n", err.Error())
		os.Exit(-1)
	}
	//if err := engine.Sync2(new(model.Book)); err != nil {
	//	global.Log.Fatalf("sync error: %#v\n", err.Error())
	//	os.Exit(-1)
	//}

	engine.ShowSQL(false)
	engine.SetMaxIdleConns(5)
	engine.SetMaxOpenConns(30)
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	global.Log.Info("[NewDBEngine] connect to dao success")

	timer := time.NewTicker(time.Minute * 30)
	go func(engine *xorm.Engine) {
		for _ = range timer.C {
			err = engine.Ping()
			if err != nil {
				global.Log.Fatalf("dao connect error: %#v\n", err.Error())
			}
		}
	}(engine)
	global.DB = engine
}
