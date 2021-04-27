package main

import (
	"github.com/fullstacker-go/practice_gin/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*.html")
	r.GET("/", handler.HomeHandler)
	r.GET("/users", handler.GetAll)
	r.GET("/user/:id", handler.GetAUser)
	r.POST("/user", handler.CreateUser)

	r.Run(":3000")
}
