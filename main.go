package main

import (
	"encoding/gob"
	_ "xinmall/routers"
	"github.com/astaxie/beego"
	"xinmall/common"
	"xinmall/models"
	"github.com/astaxie/beego/plugins/cors"

	_ "github.com/astaxie/beego/session/redis"
)

func main() {
	//add func map
	beego.AddFuncMap("timestapToData",common.TimestampToDate)
	models.DB.LogMode(true)
	beego.AddFuncMap("formatImage", common.FormatImage)
	beego.AddFuncMap("mul", common.Mul)
	beego.AddFuncMap("formatAttribute", common.FormatAttribute)

	
	beego.InsertFilter("*",beego.BeforeRouter,cors.Allow(&cors.Options{
		AllowOrigins: []string{"127.0.0.1"},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS"},
		AllowHeaders: []string{
			"Origin",
			"Authorization",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Content-Type"},
		ExposeHeaders: []string{
			"Content-Length",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Content-Type"},
		AllowCredentials: true, //cookie
	}))
	
	gob.Register(models.Administrator{})
	
	defer models.DB.Close()
	//
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"

	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"

	
	beego.Run()
}

