package handlers

import (
	"cryptoapp/ciphers/atbash"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AtbashHandler(context *gin.Context) {
	if context.Request.Method == "GET" {
		context.HTML(http.StatusOK, "atbash.html", nil)
		return
	}
	message := context.PostForm("message")
	res, err := atbash.Operation(message)

	if err != nil {
		context.HTML(http.StatusNotAcceptable, "atbash.html", err)
		return
	}
	context.HTML(http.StatusOK, "atbash.html", res)
}
