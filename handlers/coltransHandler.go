package handlers

import (
	"cryptoapp/ciphers/coltrans"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ColTransHandler(context *gin.Context) {
	if context.Request.Method == "GET" {
		context.HTML(http.StatusOK, "coltrans.html", nil)
		return
	}

	message := context.PostForm("message")
	key := context.PostForm("key")
	pad := context.PostForm("pad")[0]
	operation := context.PostForm("operation")

	res, err := coltrans.Operation(pad, message, key, operation)

	if err != nil {
		context.HTML(http.StatusNotAcceptable, "coltrans.html", err)
		return
	}
	context.HTML(http.StatusOK, "coltrans.html", res)

}
