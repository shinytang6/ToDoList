package routers

import (
	"quickstart/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/task/", &controllers.TaskController{},"get:GetTask;post:AddTask")
      
}
