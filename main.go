package main

import (
	_ "blog.api/models"
	_ "blog.api/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	/*
	var FilterToken = func(ctx *context.Context) {
		logs.Info("当前路由器路径为 ", ctx.Request.RequestURI)
		if ctx.Request.RequestURI != "/login" && ctx.Input.Header("Authorization") == "" {
			logs.Error("without token, unauthorized !!")
			ctx.ResponseWriter.WriteHeader(401)
			ctx.ResponseWriter.Write([]byte("no permission"))
		}
		if ctx.Request.RequestURI != "/login" && ctx.Input.Header("Authorization") != "" {
			token := ctx.Input.Header("Authorization")
			token = strings.Split(token, "")[1]
			logs.Info("curernttoken: ", token)
			// validate token
			// invoke ValidateToken in utils/token
			// invalid or expired todo res 401
		}
	}

	//请求的路由token验证
	beego.InsertFilter("*", beego.BeforeRouter, FilterToken)
	*/

	//允许跨域
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.Run()
}
