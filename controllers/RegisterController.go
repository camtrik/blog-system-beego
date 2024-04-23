package controllers

import (
	"blog_system/models"
	"fmt"
	"regexp"

	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "register.tpl"
}

func (c *RegisterController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	email := c.GetString("email")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		beego.Error("Password hashing: ", err)
	}

	userVal := CheckUsername(username)
	passwordVal := CheckPassword(password)
	if !userVal || !passwordVal {
		c.Layout = "layout.tpl"
		c.TplName = "register.tpl"
		if !passwordVal {
			c.Data["PasswordErr"] = "Password must consist of 6-12 characters, and contain at least two of the following: letters, numbers, and symbols"
		}
		if !userVal {
			c.Data["UsernameErr"] = "Username must consist of 4-12 characters, letters and numbers only"
		}
		c.Data["Username"] = username
		c.Data["Email"] = email
		c.Render()
		return
	}

	if !passwordVal {
		c.TplName = "register.tpl"

		c.Render()
		return
	}

	user := models.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	}

	if err := models.CreateUser(&user); err != nil {
		c.Data["Error"] = err.Error()
		c.TplName = "register.tpl"
		c.Render()
	} else {
		c.Data["Success"] = true
		c.TplName = "register.tpl"
		c.Render()
		fmt.Println("register sucess: ", username)
	}
}

func CheckUsername(username string) bool {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,12}$", username); !ok {
		return false
	}
	return true
}

func CheckPassword(password string) bool {
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasLetter := regexp.MustCompile(`[a-zA-Z]`).MatchString(password)
	hasSymbol := regexp.MustCompile(`[\W_]`).MatchString(password) // symbols

	return hasNumber && hasLetter || hasSymbol && hasNumber || hasSymbol && hasLetter
}
