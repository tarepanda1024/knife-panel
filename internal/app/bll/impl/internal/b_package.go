package internal

import (
	"context"
	"knife-panel/internal/app/schema"
)

// NewPackage 创建demo
func NewPackage() *Package {
	return &Package{}
}

// Package 示例程序
type Package struct {
}

// List 查询数据
func (a *Package) List(ctx context.Context, installed bool) (*schema.PackageItem, error) {
	return nil, nil
}
