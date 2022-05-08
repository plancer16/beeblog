package main

import (
	"github.com/plancer16/beeblog/controllers"
	"github.com/beego/beego"
	"github.com/plancer16/beeblog/models"
)

func main() {
	models.GetLink()
	//显示博客首页
	beego.Router("/", &controllers.IndexController{})
	//查看博客详细信息
	beego.Router("/view/:id([0-9]+)", &controllers.ViewController{})
	//新建博客博文
	beego.Router("/new", &controllers.NewController{})
	//删除博文
	beego.Router("/delete/:id([0-9]+)", &controllers.DeleteController{})
	//编辑博文
	beego.Router("/edit/:id([0-9]+)", &controllers.EditController{})

	beego.Run("127.0.0.1:8080")
}