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
package api

import (
	"github.com/gofiber/fiber/v2"
	"knife-panel/server/core/http/service"
	"knife-panel/server/knife"
)

func RegisterDockerApi(app *fiber.App) {
	dockerService := service.NewDockerService()
	dockerManage := app.Group("/xhr/v1/dockers")

	dockerManage.Get("/listContainers", func(ctx *fiber.Ctx) error {
		if result, err := dockerService.ListContainers(); err == nil {
			return knife.Success(ctx, result)
		} else {
			return err
		}
	})

	dockerManage.Get("/listImages", func(ctx *fiber.Ctx) error {
		if result, err := dockerService.ListImages(); err == nil {
			return knife.Success(ctx, result)
		} else {
			return err
		}
	})

	dockerManage.Get("/listNets", func(ctx *fiber.Ctx) error {
		if result, err := dockerService.ListNets(); err == nil {
			return knife.Success(ctx, result)
		} else {
			return err
		}
	})

	dockerManage.Get("/listVolumes", func(ctx *fiber.Ctx) error {
		if result, err := dockerService.ListVolumes(); err == nil {
			return knife.Success(ctx, result)
		} else {
			return err
		}
	})
}
