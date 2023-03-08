package handlers

import (
	"cryptoapp/ciphers/vigenere"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VigenereHandler(context *gin.Context) {
	if context.Request.Method == "GET" {
		context.HTML(http.StatusOK, "vigenere.html", nil)
		return
	}
	message := context.PostForm("message")
	key := context.PostForm("key")
	operation := context.PostForm("operation")

	res, err := vigenere.Operation(message, key, operation)

	if err != nil {
		context.HTML(http.StatusMethodNotAllowed, "vigenere.html", err)
		return
	}
	context.HTML(http.StatusOK, "vigenere.html", res)
}
