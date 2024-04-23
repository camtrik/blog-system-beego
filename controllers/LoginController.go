package controllers

import (
	"blog_system/models"
	"fmt"

	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.tpl"
}

func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")

	user, err := models.GetUser(username)
	if err != nil {
		c.Data["Error"] = "User does not exist"
		c.TplName = "login.tpl"
	} else if VerifyPassword(user, password) {
		// login success, set session
		c.SetSession("uid", user.ID)
		c.SetSession("uname", user.Username)
		c.Redirect("/", 302)
		// set session
		fmt.Println("login success: ", username)
		return
	} else {
		c.Data["Error"] = "Password is incorrect"
		c.TplName = "login.tpl"

	}
	c.Render()

}

func VerifyPassword(user models.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
