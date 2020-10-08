// Copyright 2019 liuxiaodong Authors
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
package tty

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"knife-panel/server/thirdpart/terminal/backend/localcommand"
	"knife-panel/server/thirdpart/terminal/server"
	"knife-panel/server/utils"
	"os"
)

func RegisterRouter(app *fiber.App) {
	tty := app.Group("/tty")
	{
		tty.Get("/", websocket.New(func(c *websocket.Conn) {

			appOptions := &server.Options{
				PermitWrite: true,
				Width:       600,
				Height:      800,
			}
			if err := utils.ApplyDefaultValues(appOptions); err != nil {
				exit(err, 1)
			}
			backendOptions := &localcommand.Options{}
			if err := utils.ApplyDefaultValues(backendOptions); err != nil {
				exit(err, 1)
			}

			factory, err := localcommand.NewFactory("/usr/bin/bash", nil, backendOptions)
			if err != nil {
				exit(err, 3)
			}
			gCtx, _ := context.WithCancel(context.Background())
			srv, err := server.New(factory, appOptions)
			if err != nil {
				exit(err, 3)
			}
			if err := srv.ProcessWSConn(gCtx, c); err != nil {
				_ = c.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			}
		}))
	}
}
func exit(err error, code int) {
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}
