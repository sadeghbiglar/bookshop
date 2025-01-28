package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sadeghbiglar/bookshop/config"
)

func main() {
	app := gin.Default()
	app.Use(gin.Recovery())
	app.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message":    "server is running fine!",
			"statusCode": 200,
		})
	})
	fmt.Println("Go app started successfully !")
	port := config.GetEnvProperty("port")
	fmt.Println(port)
	app.Run(port)
}
