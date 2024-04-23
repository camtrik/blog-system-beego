package controllers

import (
	"blog_system/models"

	"github.com/astaxie/beego"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {

	var err error
	c.Data["blogs"], err = models.GetAll()
	if err != nil {
		beego.Error(err)
	}
	c.TplName = "index.tpl"
	c.Layout = "layout.tpl"
}
