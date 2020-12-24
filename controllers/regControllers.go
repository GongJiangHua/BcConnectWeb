package controllers

import (
	"BitConnectWeb/models"
	"github.com/astaxie/beego"
)

type RegControllers struct {
	beego.Controller
}

func (r *RegControllers) Post()  {
	var user models.User
	err := r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("注册用户失败，请重新尝试！！")
		return
	}
	//fmt.Println(user)
	//将扒取到的User数据存入数据库
	_,err = user.SaveUser()
	if err != nil {
		r.Ctx.WriteString("用户信息注册失败，请重试！")
		return
	}
	//fmt.Println("数据库影响了：",rows)
	r.TplName = "login.html"
}
