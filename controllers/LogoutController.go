package controllers

import "github.com/astaxie/beego"

type LogoutController struct {
	beego.Controller
}

func (c *LogoutController) Get() {
	c.DestroySession()
	c.Redirect("/", 302)
}