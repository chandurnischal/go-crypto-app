package handlers

import (
	"cryptoapp/ciphers/playfair"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PlayfairHandler(context *gin.Context) {
	if context.Request.Method == "GET" {
		context.HTML(http.StatusOK, "playfair.html", nil)
		return
	}

	message := context.PostForm("message")
	key := context.PostForm("key")
	operation := context.PostForm("operation")

	res, err := playfair.Operation(message, key, operation)

	if err != nil {
		context.HTML(http.StatusNotAcceptable, "playfair.html", err)
		return
	}

	context.HTML(http.StatusOK, "playfair.html", res)

}
