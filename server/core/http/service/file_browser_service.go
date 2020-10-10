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
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"knife-panel/server/core/http/model"
	"os"
	"path"
	"sort"
)

func NewFileBrowser() *FileBrowser {
	return &FileBrowser{}
}

type FileBrowser struct {
}

// List 查询数据
func (a *FileBrowser) List(basePath string) (model.FileItems, error) {
	fileList, e := ioutil.ReadDir(basePath)
	if e != nil {
		return nil, e
	}
	fileItems := model.FileItems{}
	for _, v := range fileList {
		if len(v.Name()) == 0 {
			continue
		}
		fileItem := model.FileItem{
			AbsPath:    path.Join(basePath, v.Name()),
			Name:       v.Name(),
			Dir:        v.IsDir(),
			Size:       v.Size(),
			ModifyTime: v.ModTime().Unix(),
		}
		fileItems = append(fileItems, &fileItem)
	}
	sort.Sort(fileItems)
	return fileItems, nil
}

// Download 查询指定数据
func (a *FileBrowser) Download(ctx *fiber.Ctx) error {
	file := ctx.Params("absPath")
	return ctx.Download(file)
}

// Upload 创建数据
func (a *FileBrowser) Upload(ctx *fiber.Ctx) (bool, error) {
	file, _ := ctx.FormFile("file")
	dst := ctx.FormValue("dst")
	return true, ctx.SaveFile(file, dst)
}

// Delete 删除数据
func (a *FileBrowser) Delete(absPath string) (bool, error) {
	if info, err := os.Stat(absPath); err == nil {
		if info.IsDir() {
			return true, os.RemoveAll(absPath)
		} else {
			return true, os.Remove(absPath)
		}
	} else {
		return false, err
	}
}
