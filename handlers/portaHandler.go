package handlers

import (
	"cryptoapp/ciphers/porta"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PortaHandler(context *gin.Context) {
	if context.Request.Method == "GET" {
		context.HTML(http.StatusOK, "porta.html", nil)
		return
	}
	message := context.PostForm("message")
	key := context.PostForm("key")
	res, err := porta.Operation(message, key)

	if err != nil {
		context.HTML(http.StatusNotAcceptable, "porta.html", err)
		return
	}

	context.HTML(http.StatusOK, "porta.html", res)
}
