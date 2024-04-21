package main

import (
	models "blog_system/models"
	_ "blog_system/routers"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

func formatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

var GlobalSessions *session.Manager

func init() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = beego.AppConfig.String("sessionprovider")
	beego.BConfig.WebConfig.Session.SessionName = beego.AppConfig.String("sessionname")
	maxtime, _ := beego.AppConfig.Int64("sessiongcmaxlifetime")
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = maxtime

	models.Init()
	beego.AddFuncMap("formatTime", formatTime)

	// Initialize session manager
	sessionConfig := &session.ManagerConfig{
		CookieName:      beego.BConfig.WebConfig.Session.SessionName,
		EnableSetCookie: true,
		Gclifetime:      beego.BConfig.WebConfig.Session.SessionGCMaxLifetime,
		Maxlifetime:     beego.BConfig.WebConfig.Session.SessionGCMaxLifetime,
	}

	GlobalSessions, _ = session.NewManager("memory", sessionConfig)
	go GlobalSessions.GC()
}
func main() {

	beego.Run()
}
