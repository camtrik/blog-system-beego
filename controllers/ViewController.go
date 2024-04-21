package controllers

import (
	"blog_system/models"
	"strconv"

	"github.com/astaxie/beego"
)

type ViewController struct {
	beego.Controller
}

func (c *ViewController) Get() {
	sess := c.StartSession()
	var visitCount int
	if count := sess.Get("visitCount"); count != nil {
		visitCount = count.(int)
	}
	visitCount++
	sess.Set("visitCount", visitCount)
	c.Data["visitCount"] = visitCount
	beego.Info("visitCount:", visitCount) 

	var err error
	id, _ := strconv.Atoi(c.Ctx.Input.Params()[":id"])
	c.Data["Post"], err = models.GetBlog(id)
	if err != nil {
		beego.Error(err)
		return
	}
	c.Layout = "layout.tpl"
	c.TplName = "view.tpl"
}
