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

func RegisterFileBrowser(app *fiber.App) {
	fileManageService := service.NewFileBrowser()

	fileManage := app.Group("/xhr/v1/files")

	fileManage.Get("/list", func(ctx *fiber.Ctx) error {
		basePath := ctx.Query("basePath")
		if result, err := fileManageService.List(basePath); err == nil {
			return knife.Success(ctx, result)
		} else {
			return err
		}
	})

	fileManage.Post("/upload", func(ctx *fiber.Ctx) error {
		if result, err := fileManageService.Upload(ctx); err == nil {
			return knife.Success(ctx, result)
		} else {
			return err
		}
	})

	fileManage.Get("/download", func(ctx *fiber.Ctx) error {
		return fileManageService.Download(ctx)
	})

	fileManage.Get("/delete", func(ctx *fiber.Ctx) error {
		basePath := ctx.Query("basePath")
		if result, err := fileManageService.Delete(basePath); err == nil {
			return knife.Success(ctx, result)
		} else {
			return err
		}
	})

}
