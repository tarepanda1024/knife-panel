package ctl

import (
	"github.com/gin-gonic/gin"
	"knife-panel/internal/app/bll"
	"knife-panel/internal/app/ginplus"
	"strconv"
)

func NewPackage(packageBll bll.IPackage) *Package {
	return &Package{
		PackageBll: packageBll,
	}
}

type Package struct {
	PackageBll bll.IPackage
}

func (a *Package) List(c *gin.Context) {
	installed, err := strconv.ParseBool(c.Query("installed"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	result, err := a.PackageBll.List(ginplus.NewContext(c), installed)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}

	ginplus.ResSuccess(c, result)
}
