package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func mainHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "main.html", nil)
}

func affineHandler(context *gin.Context) {
	if context.Request.Method == "GET" {
		context.HTML(http.StatusOK, "affine.html", nil)
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "hello world"})
}

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("templates/*.html")
	server.GET("/", mainHandler)
	server.GET("/affine", affineHandler)
	server.POST("/affine", affineHandler)
	server.Run(":8080")
}
