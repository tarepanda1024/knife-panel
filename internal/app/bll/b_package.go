package bll

import (
	"context"
	"knife-panel/internal/app/schema"
)

// IFileBrowser demo业务逻辑接口
type IPackage interface {
	// 查询数据
	List(ctx context.Context, installed bool) (*schema.PackageItem, error)
}
