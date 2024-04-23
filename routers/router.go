package routers

import (
	"blog_system/controllers"
	"fmt"

	"github.com/astaxie/beego"
)

func init() {
	fmt.Println("init router")
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/new", &controllers.NewController{})
	beego.Router("/view/:id", &controllers.ViewController{})
	beego.Router("/edit/:id", &controllers.EditController{})
	beego.Router("/delete/:id", &controllers.DeleteController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
}
