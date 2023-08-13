package main

import (
	"ginpackage/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static")
	router.GET("/", handler.IndexPage)
	router.POST("/", handler.IndexPage)
	router.GET("/signup", handler.Signup)
	router.POST("/signuppost", handler.SignupPost)
	router.GET("/login", handler.Login)
	router.POST("/loginpost", handler.LoginPost)
	router.GET("/home", handler.HomeMethod)
	router.POST("/logout", handler.Logout)
	router.Run()

}
