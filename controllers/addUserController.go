package controllers

import (
	"blog_system/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/go-playground/validator/v10"
)

type AddUserController struct {
	beego.Controller
}

func (c *AddUserController) Get() {
	c.Layout = "layout.tpl"
	c.TplName = "adduser.tpl"
}

func (c *AddUserController) Post() {
	var user models.User
	if err := c.ParseForm(&user); err != nil {
		beego.Error("Fail to parse form: ", err)
		c.Data["FormError"] = "Failed to parse the form. Please check your input."
		c.TplName = "adduser.tpl"
		c.Render()
		return
	}
	fmt.Printf("Parsed User: %+v\n", user)
	if err := user.Validate(); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMap := make(map[string]string)
		for _, err := range validationErrors {
			switch err.Tag() {
			case "required":
				errorMap[err.Field()] = fmt.Sprintf("%s is required", err.Field())
			case "max":
				if err.Field() == "Username" {
					errorMap[err.Field()] = "Username must be no more than 8 characters"
				}
			case "email":
				errorMap[err.Field()] = "Invalid email format"
			case "numeric":
				errorMap[err.Field()] = "Must be a number"
			}
		}
		c.Data["Errors"] = errorMap
		c.TplName = "adduser.tpl"
		c.Render()
		return
	}

	if err := models.UserInsert(&user); err != nil {
		beego.Error("Failed to insert user: ", err)
		c.Data["FormError"] = "Failed to insert user. Please try again."
	} else {
		c.Data["Success"] = "User added successfully! redirect to home in 3 seconds"
	}

	c.TplName = "adduser.tpl"
	c.Render()
}
