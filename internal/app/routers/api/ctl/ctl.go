package ctl

import (
	"go.uber.org/dig"
)

// Inject 注入ctl
// 使用方式：
//   container := dig.New()
//   Inject(container)
//   container.Invoke(func(foo *ctl.FileBrowser) {
//   })
func Inject(container *dig.Container) error {
	_ = container.Provide(NewLogin)
	_ = container.Provide(NewMenu)
	_ = container.Provide(NewRole)
	_ = container.Provide(NewUser)
	_ = container.Provide(NewFileBrowser)
	_ = container.Provide(NewSystemMonitor)
	_ = container.Provide(NewPackage)
	return nil
}
