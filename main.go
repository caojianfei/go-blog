package main

import (
	"blog/web/admin/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./web/views", ".html"))

	mvc.Configure(app.Party("/admin"), amdinMvc)

	err := app.Run(iris.Addr(":8080"))
	if err != nil {
		panic(err.Error())
	}

}


func amdinMvc(app *mvc.Application) {
	//app.Party("/user/login").Handle(new(controllers.UserController))

	app.Party("/user").Handle(new(controllers.UserController))
	//app.Handle(new(controllers.UserController))
}