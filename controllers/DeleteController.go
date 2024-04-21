package controllers

import (
	"blog_system/models"
	"strconv"

	"github.com/astaxie/beego"
)

type DeleteController struct {
	beego.Controller
}

func (c *DeleteController) Get() {
	id, _ := strconv.Atoi(c.Ctx.Input.Params()[":id"])
	blog, err := models.GetBlog(id)
	if err != nil {
		beego.Error(err)
	}

	c.Data["Post"] = blog
	models.DelBlog(&blog)
	c.Ctx.Redirect(302, "/")
}
