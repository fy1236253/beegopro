package controllers

import (
	"github.com/astaxie/beego"
)

type WebController struct {
	beego.Controller
}

func (this *WebController) Show() {
	this.TplName = "dezhen.html"
}
