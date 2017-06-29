package controllers

import (
	"log"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type HomeController struct {
	beego.Controller
}
type BillController struct {
	beego.Controller
}
type DownloadController struct {
	beego.Controller
}
type FormController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "scoller.html"
}
func (c *MainController) Post() {
	name := c.GetString("name")
	log.Println(name)
}
func (c *HomeController) Get() {
	c.TplName = "amaui.html"
}
func (c *FormController) Get() {
	c.TplName = "add_form.html"
}
func (c *FormController) Post() {
	name := c.GetString("name")
	log.Println(name)
	if name != "" {
		c.TplName = "amaui.html"
	} else {
		c.Ctx.WriteString("数据有误")
	}
	// c.Ctx.WriteString("成功提交数据")
	// c.Redirect("/home", 302)
	// return
}
func (this *DownloadController) Get() {
	param := this.Ctx.Input.Param(":name")
	log.Println(param)
	if param == "" {
		this.TplName = "amaui.html"
	}
	this.Ctx.Output.Download("download/" + param)
}
func (this *DownloadController) Page() {
	params := this.Ctx.Input.Params()
	log.Println(params["0"])
	this.TplName = "download.html"
}
func (this *BillController) Show() {
	this.TplName = "bill.html"
}
