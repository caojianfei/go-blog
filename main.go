package main

import (
	"blog/web/admin/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./web/views", ".html").Reload(true))

	// 静态资源服务
	app.StaticWeb("/static", "./web/static")

	// 管理端路由
	mvc.Configure(app.Party("/admin"), adminRouter)

	err := app.Run(iris.Addr(":8080"))
	if err != nil {
		panic(err.Error())
	}
}

func adminRouter(app *mvc.Application) {
	app.Party("/user").Handle(new(controllers.UserController))
}