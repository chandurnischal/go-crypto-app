package main

import (
	"cryptoapp/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	server.LoadHTMLGlob("templates/*.html")
	server.GET("/", handlers.MainHandler)
	server.GET("/affine", handlers.AffineHandler)
	server.POST("/affine", handlers.AffineHandler)
	server.GET("/atbash", handlers.AtbashHandler)
	server.POST("/atbash", handlers.AtbashHandler)
	server.GET("/autokey", handlers.AutokeyHandler)
	server.POST("/autokey", handlers.AutokeyHandler)
	server.Run(":8080")
}
