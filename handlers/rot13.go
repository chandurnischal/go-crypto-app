package handlers

import (
	"cryptoapp/ciphers/rot13"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ROT13Handler(context *gin.Context) {
	if context.Request.Method == "GET" {
		context.HTML(http.StatusOK, "rot13.html", nil)
		return
	}
	message := context.PostForm("message")
	operation := context.PostForm("operation")

	res, err := rot13.Operation(message, operation)

	if err != nil {
		context.HTML(http.StatusNotAcceptable, "rot13.html", err)
		return
	}
	context.HTML(http.StatusOK, "rot13.html", res)

}
