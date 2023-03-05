package handlers

import (
	"cryptoapp/ciphers/polybius"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PolybiusHandler(context *gin.Context) {
	if context.Request.Method == "GET" {
		context.HTML(http.StatusOK, "polybius.html", nil)
		return
	}
	message := context.PostForm("message")
	key := context.PostForm("key")
	chars := context.PostForm("chars")[:5]
	operation := context.PostForm("operation")

	res, err := polybius.Operation(message, key, chars, operation)

	if err != nil {
		context.HTML(http.StatusNotAcceptable, "polybius.html", err)
		return
	}

	context.HTML(http.StatusOK, "polybius.html", res)
}
