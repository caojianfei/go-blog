package controllers

import (
	"github.com/kataras/iris"
)

type UserController struct {}

func (ac *UserController) GetLogin(ctx iris.Context) {
	//fmt.Println(ctx.Request().Method)
	err := ctx.View("admin/user/login.html")
	if err != nil {
		panic(err.Error())
	}
}