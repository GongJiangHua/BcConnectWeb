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
	l.ParseForm(&user)
	fmt.Println(user.Phone)
}
