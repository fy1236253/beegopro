package controllers

import (
	"io"
	"log"
	"os"
	"time"

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
	p := c.Ctx.Request.Form
	log.Println(p)
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	timer := tm.Format("2006-01-02 15:04:05")
	f, _ := os.OpenFile("logs/user.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	data := "参赛单位:" + p["unit"][0] + "\n" + "参赛作品名称:" + p["pro_name"][0] + "\n" + "联系人电话:" + p["tel"][0] + "\n" + "联系人姓名:" + p["user"][0] + "\n" + "邮箱:" + p["email"][0] + "\n" + "通讯地址:" + p["address"][0] + "\n"

	io.WriteString(f, "\n【"+timer+"】\n==================================\n"+data)
	defer f.Close()
	c.TplName = "pay.html"
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
	if params["0"] == "" {
		this.TplName = "download.html"
	} else {
		this.Ctx.Output.Download("download/" + params["0"])
	}
}
func (this *BillController) Show() {
	p := this.Ctx.Request.Form
	log.Println(p)
	if this.Ctx.Request.Method == "POST" {
		timestamp := time.Now().Unix()
		tm := time.Unix(timestamp, 0)
		timer := tm.Format("2006-01-02 15:04:05")
		f, _ := os.OpenFile("logs/ticket.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		data := "参赛单位:" + p["unit"][0] + "\n" + "纳税人识别码:" + p["code"][0] + "\n" + "单位地址:" + p["address"][0] + "\n" + "单位电话:" + p["tel"][0] + "\n" + "开户银行:" + p["bank_name"][0] + "\n" + "开户银行账号:" + p["bank"][0] + "\n"

		io.WriteString(f, "\n【"+timer+"】\n==================================\n"+data)
		defer f.Close()
		this.Ctx.WriteString("申请提交成功")
	} else {
		this.TplName = "bill.html"
	}
}
