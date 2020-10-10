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
package model

type FileItem struct {
	AbsPath    string `json:"absPath"`
	Name       string `json:"name"`
	Dir        bool   `json:"dir"`
	Size       int64  `json:"size"`
	ModifyTime int64  `json:"modifyTime"`
}

type FileItems []*FileItem

func (p FileItems) Len() int { return len(p) }

// 根据元素的年龄降序排序 （此处按照自己的业务逻辑写）
func (p FileItems) Less(i, j int) bool {
	return p[i].Dir
}

// 交换数据
func (p FileItems) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
