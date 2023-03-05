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

	server.GET("/baconian", handlers.BaconianHandler)
	server.POST("/baconian", handlers.BaconianHandler)

	server.GET("/beaufort", handlers.BeaufortHandler)
	server.POST("/beaufort", handlers.BeaufortHandler)

	server.GET("/caesar", handlers.CasarHandler)
	server.POST("/caesar", handlers.CasarHandler)

	server.Run(":8080")
}
