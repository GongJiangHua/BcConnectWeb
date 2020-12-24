package routers

import (
	"BitConnectWeb/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//默认网址是注册页面
    beego.Router("/", &controllers.MainController{})
    //注册页面
    beego.Router("/register",&controllers.RegControllers{})
    //跳转首页面
    beego.Router("/index",&controllers.LoginControllers{})
}
