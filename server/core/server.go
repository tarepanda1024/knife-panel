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
package core

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"knife-panel/server/global"
	"knife-panel/server/initialize"
)

func RunServer() {
	address := fmt.Sprintf("127.0.0.1:%d", global.AppConfig.System.Addr)
	app := fiber.New()
	initialize.InitRouter(app)
	global.Log.Info("listening at :", address)
	if err := app.Listen(address); err != nil {
		global.Log.Error(err)
	}
}
