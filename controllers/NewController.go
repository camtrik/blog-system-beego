package controllers

import (
	"blog_system/models"
	"time"

	"github.com/astaxie/beego"
)

type NewController struct {
	beego.Controller
}

func (c *NewController) Get() {
	c.Layout = "layout.tpl"
	c.TplName = "new.tpl"
}

func (c *NewController) Post() {
	inputs := c.Input()
	blog := models.Blog{
		// 系统会显示自动插入Id的值
		Title:   inputs.Get("title"), // 从表单中获得字段
		Content: inputs.Get("content"),
		Created: time.Now(),
	}

	if err := models.SaveBlog(&blog); err != nil {
		beego.Error(err)
	}

	c.Redirect("/", 302)
}
