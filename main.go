package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AffineAPI struct {
	Multiplier int    `json:"multiplier"`
	Offset     int    `json:"offset"`
	Message    string `json:"message"`
	Encrypt    bool   `json:"encrypt"`
}

func affineHandler(context *gin.Context) {
	var aff AffineAPI
	err := context.BindJSON(&aff)
	if err != nil {
		context.HTML(http.StatusNotFound, "result.html", gin.H{"message": "invalid input"})
	}
	if aff.Encrypt {
		context.HTML(http.StatusOK, "result.html", gin.H{"message": "encrypt"})
		return
	}
	context.HTML(http.StatusOK, "result.html", gin.H{"message": "decrypt"})
}

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("templates/*.html")
	server.GET("/affine", affineHandler)
	server.Run(":8080")
}
