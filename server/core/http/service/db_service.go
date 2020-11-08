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
package service

import (
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"knife-panel/server/core/http/model"
	"knife-panel/server/global"
)

type DBService struct {
	dbInstances map[uint]*gorm.DB
}

func NewDBService() *DBService {
	return &DBService{
		dbInstances: map[uint]*gorm.DB{},
	}
}

func (this *DBService) Create(dbProp *model.DBProp) error {
	global.DB.Create(dbProp)
	return nil
}

func (this *DBService) List() []model.DBProp {
	dbs := make([]model.DBProp, 0)
	global.DB.Find(&dbs)
	return dbs
}

func (this *DBService) getDBInstance(id uint) (*gorm.DB, error) {
	if this.dbInstances[id] != nil {
		return this.dbInstances[id], nil
	}
	dbProp := model.DBProp{}
	global.DB.First(&dbProp, id)
	switch dbProp.Type {
	case "sqlite":
		instance, err := gorm.Open(sqlite.Open(dbProp.Path), &gorm.Config{})
		if err == nil {
			this.dbInstances[dbProp.ID] = instance
			return this.dbInstances[id], nil
		}
		return nil, err

	default:
		return nil, errors.New("unSupport db type")
	}
}

func (this *DBService) Connect(id uint) error {
	if _, err := this.getDBInstance(id); err != nil {
		return err
	} else {
		return nil
	}
}

func (this *DBService) ListTables(id uint) ([]string, error) {
	if db, err := this.getDBInstance(id); err != nil {
		return nil, err
	} else {
		tables := make([]string, 0)
		dbType := db.Dialector.Name()
		switch dbType {
		case "sqlite":
			sqlite3Tables := make([]model.Sqlite3Table, 0)
			db.Raw("select * from sqlite_master where type = ?", "table").Scan(&sqlite3Tables)
			for _, table := range sqlite3Tables {
				tables = append(tables, table.Name)
			}
			return tables, nil
		}
		return nil, nil
	}
}

func (this *DBService) ListColumns(id uint, tblName string) ([]model.TableColumn, error) {
	if db, err := this.getDBInstance(id); err != nil {
		return nil, err
	} else {
		tableColumns := make([]model.TableColumn, 0)
		dbType := db.Dialector.Name()
		switch dbType {
		case "sqlite":
			sqlite3Columns := make([]model.Sqlite3Column, 0)
			//TODO:可能存在安全问题
			db.Debug().Raw("PRAGMA table_info(" + tblName + ")").Scan(&sqlite3Columns)
			for _, columnItem := range sqlite3Columns {
				tableColumns = append(tableColumns, model.TableColumn{
					Name:         columnItem.Name,
					Type:         columnItem.Type,
					NotNull:      columnItem.NotNull,
					DefaultValue: columnItem.DefaultValue,
					PK:           columnItem.PK,
				})
			}
			return tableColumns, nil
		}
		return nil, nil
	}
}
