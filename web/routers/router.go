package routers

import (
	"web/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/form_house", &controllers.MainController{}, "get:SetHouseInfo")
	beego.Router("/house_search", &controllers.MainController{}, "get:GetHouseInfo")
}
