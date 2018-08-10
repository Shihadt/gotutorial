package routers

import (
	"github.com/user/examples/web_tutorial/my_project/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
