package controllers

import (
	"blog_system/models"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type EditController struct {
	beego.Controller
}

func (c *EditController) Get() {
	var err error
	id, _ := strconv.Atoi(c.Ctx.Input.Params()[":id"])
	c.Data["Post"], err = models.GetBlog(id)
	if err != nil {
		beego.Error(err)
	}
	c.Layout = "layout.tpl"
	c.TplName = "edit.tpl"
}

func (c *EditController) Post() {
	inputs := c.Input()
	id, _ := strconv.Atoi(c.Ctx.Input.Params()[":id"])
	var blog = models.Blog{
		Id:      id,
		Title:   inputs.Get("title"),
		Content: inputs.Get("content"),
		Created: time.Now(),
	}
	models.SaveBlog(&blog)
	c.Redirect("/", 302)
}
