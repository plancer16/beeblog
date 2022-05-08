package controllers

import (
	"github.com/beego/beego"
	"github.com/plancer16/beeblog/models"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	this.Data["blogs"] = models.GetAll()
	this.Layout = "layout.tpl"
	this.TplName = "index.tpl"
}