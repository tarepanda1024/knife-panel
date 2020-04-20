package api

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"knife-panel/internal/app/middleware"
	"knife-panel/internal/app/routers/api/ctl"
	"knife-panel/pkg/auth"
	"os"
)

// RegisterRouter 注册/api路由
func RegisterRouter(app *gin.Engine, container *dig.Container) error {

	return container.Invoke(func(
		a auth.Auther,
		e *casbin.SyncedEnforcer,
		cFileBrowser *ctl.FileBrowser,
		cPackage *ctl.Package,
		cLogin *ctl.Login,
		cMenu *ctl.Menu,
		cRole *ctl.Role,
		cUser *ctl.User,
		monitor *ctl.SystemMonitor,
	) error {

		g := app.Group("/api")

		// 用户身份授权
		g.Use(middleware.UserAuthMiddleware(a,
			middleware.AllowPathPrefixSkipper("/api/v1/pub/login"),
		))

		// casbin权限校验中间件
		g.Use(middleware.CasbinMiddleware(e,
			middleware.AllowPathPrefixSkipper("/api/v1/pub"),
		))

		// 请求频率限制中间件
		g.Use(middleware.RateLimiterMiddleware())

		v1 := g.Group("/v1")
		{
			pub := v1.Group("/pub")
			{
				// 注册/api/v1/pub/login
				gLogin := pub.Group("login")
				{
					gLogin.GET("captchaid", cLogin.GetCaptcha)
					gLogin.GET("captcha", cLogin.ResCaptcha)
					gLogin.POST("", cLogin.Login)
					gLogin.POST("exit", cLogin.Logout)
				}

				// 注册/api/v1/pub/refresh-token
				pub.POST("/refresh-token", cLogin.RefreshToken)

				// 注册/api/v1/pub/current
				gCurrent := pub.Group("current")
				{
					gCurrent.PUT("password", cLogin.UpdatePassword)
					gCurrent.GET("user", cLogin.GetUserInfo)
					gCurrent.GET("menutree", cLogin.QueryUserMenuTree)
				}

			}

			// 注册/api/v1/demos
			fileBrowser := v1.Group("file-browser")
			{
				fileBrowser.GET("", cFileBrowser.List)
				fileBrowser.GET(":id", cFileBrowser.Download)
				fileBrowser.POST("", cFileBrowser.Upload)
				fileBrowser.DELETE(":id", cFileBrowser.Delete)
			}

			packageManage := v1.Group("/package-manager")
			{
				packageManage.GET("", cPackage.List)
			}

			// 注册/api/v1/menus
			gMenu := v1.Group("menus")
			{
				gMenu.GET("", cMenu.Query)
				gMenu.GET(":id", cMenu.Get)
				gMenu.POST("", cMenu.Create)
				gMenu.PUT(":id", cMenu.Update)
				gMenu.DELETE(":id", cMenu.Delete)
			}
			v1.GET("/menus.tree", cMenu.QueryTree)

			// 注册/api/v1/roles
			gRole := v1.Group("roles")
			{
				gRole.GET("", cRole.Query)
				gRole.GET(":id", cRole.Get)
				gRole.POST("", cRole.Create)
				gRole.PUT(":id", cRole.Update)
				gRole.DELETE(":id", cRole.Delete)
			}
			v1.GET("/roles.select", cRole.QuerySelect)

			// 注册/api/v1/users
			gUser := v1.Group("users")
			{
				gUser.GET("", cUser.Query)
				gUser.GET(":id", cUser.Get)
				gUser.POST("", cUser.Create)
				gUser.PUT(":id", cUser.Update)
				gUser.DELETE(":id", cUser.Delete)
				gUser.PATCH(":id/enable", cUser.Enable)
				gUser.PATCH(":id/disable", cUser.Disable)
			}

			systemMonitor := v1.Group("system-monitor")
			{
				systemMonitor.GET("", monitor.GetSystemInfo)
			}

		}

		return nil
	})

}
func exit(err error, code int) {
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}
