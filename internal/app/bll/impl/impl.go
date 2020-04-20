package impl

import (
	"go.uber.org/dig"
	"knife-panel/internal/app/bll"
	"knife-panel/internal/app/bll/impl/internal"
)

// Inject 注入bll实现
// 使用方式：
//   container := dig.New()
//   Inject(container)
//   container.Invoke(func(foo IFileBrowser) {
//   })
func Inject(container *dig.Container) error {
	_ = container.Provide(internal.NewTrans)
	_ = container.Provide(func(b *internal.Trans) bll.ITrans { return b })
	_ = container.Provide(internal.NewLogin)
	_ = container.Provide(func(b *internal.Login) bll.ILogin { return b })
	_ = container.Provide(internal.NewMenu)
	_ = container.Provide(func(b *internal.Menu) bll.IMenu { return b })
	_ = container.Provide(internal.NewRole)
	_ = container.Provide(func(b *internal.Role) bll.IRole { return b })
	_ = container.Provide(internal.NewUser)
	_ = container.Provide(func(b *internal.User) bll.IUser { return b })
	_ = container.Provide(internal.NewFileBrowser)
	_ = container.Provide(func(b *internal.FileBrowser) bll.IFileBrowser { return b })
	_ = container.Provide(internal.NewSystemMonitor)
	_ = container.Provide(func(b *internal.SystemMonitor) bll.ISystemMonitor { return b })
	_ = container.Provide(internal.NewPackage)
	_ = container.Provide(func(b *internal.Package) bll.IPackage { return b })
	return nil
}
