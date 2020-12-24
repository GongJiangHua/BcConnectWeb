package controllers

import (
	"BitConnectWeb/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginControllers struct {
	beego.Controller
}

func (l *LoginControllers) Post() {
	var user models.User
	err := l.ParseForm(&user)
	if err != nil {
		l.Ctx.WriteString("用户")
	}
	fmt.Println(user.Phone)
}
