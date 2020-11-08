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
	"knife-panel/server/core/http/model"
	"knife-panel/server/core/http/service"
	"knife-panel/server/knife"
	"net/http"

	"strconv"
)

func RegisterDBApi(app *fiber.App) {
	dbService := service.NewDBService()

	dbApis := app.Group("/xhr/v1/dbs")

	dbApis.Post("/create", func(ctx *fiber.Ctx) error {
		dbProp := &model.DBProp{}
		if err := ctx.BodyParser(dbProp); err != nil {
			return knife.Fail(ctx, http.StatusBadRequest, "Illegal params")
		}
		if err := dbService.Create(dbProp); err != nil {
			return knife.Fail(ctx, http.StatusBadRequest, "Illegal params")
		}
		return knife.Success(ctx, nil)
	})

	dbApis.Get("/list", func(ctx *fiber.Ctx) error {
		result := dbService.List()
		return knife.Success(ctx, result)
	})

	dbApis.Get("/connect", func(ctx *fiber.Ctx) error {
		if id, err := strconv.ParseUint(ctx.Query("id"), 10, 32); err != nil {
			return knife.Fail(ctx, http.StatusBadRequest, "Illegal params")
		} else {
			if err := dbService.Connect(uint(id)); err != nil {
				return knife.Fail(ctx, http.StatusBadRequest, "Illegal params")
			}
		}
		return knife.Success(ctx, nil)
	})

	dbApis.Get("/listTables", func(ctx *fiber.Ctx) error {
		if id, err := strconv.ParseUint(ctx.Query("id"), 10, 32); err != nil {
			return knife.Fail(ctx, http.StatusBadRequest, "Illegal params")
		} else {
			if result, err := dbService.ListTables(uint(id)); err != nil {
				return knife.Fail(ctx, http.StatusBadRequest, err.Error())
			} else {
				return knife.Success(ctx, result)
			}
		}
	})

	dbApis.Get("/listColumns", func(ctx *fiber.Ctx) error {
		tblName := ctx.Query("tblName")
		if id, err := strconv.ParseUint(ctx.Query("id"), 10, 32); err != nil {
			return knife.Fail(ctx, http.StatusBadRequest, "Illegal params")
		} else {
			if result, err := dbService.ListColumns(uint(id), tblName); err != nil {
				return knife.Fail(ctx, http.StatusBadRequest, err.Error())
			} else {
				return knife.Success(ctx, result)
			}
		}
	})
}
