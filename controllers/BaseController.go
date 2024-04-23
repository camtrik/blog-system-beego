package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
	IsLoggedIn bool
	Username   string
}

func (c *BaseController) Prepare() {
	userID := c.GetSession("uid")
	if userID != nil {
		c.IsLoggedIn = true
        c.Username = c.GetSession("uname").(string)
	}

	c.Data["IsLoggedIn"] = c.IsLoggedIn
	c.Data["Username"] = c.Username
}