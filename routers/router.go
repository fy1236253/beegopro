package routers

import (
	"beegopro/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.DownloadController{})
	beego.AutoRouter(&controllers.WebController{})
	beego.AutoRouter(&controllers.BillController{})
	beego.Router("/home", &controllers.HomeController{})
	beego.Router("/form", &controllers.FormController{})
}
