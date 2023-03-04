package handlers

import (
	"cryptoapp/ciphers/autokey"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AutokeyHandler(context *gin.Context) {
	if context.Request.Method == "GET" {
		context.HTML(http.StatusOK, "autokey.html", nil)
		return
	}
	message := context.PostForm("message")
	key := context.PostForm("key")
	operation := context.PostForm("operation")
	res, err := autokey.Operation(message, key, operation)

	if err != nil {
		context.HTML(http.StatusNotFound, "autokey.html", err)
		return
	}

	context.HTML(http.StatusOK, "autokey.html", res)

}
