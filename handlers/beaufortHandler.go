package handlers

import (
	"cryptoapp/ciphers/beaufort"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BeaufortHandler(context *gin.Context) {
	if context.Request.Method == "GET" {
		context.HTML(http.StatusOK, "beaufort.html", nil)
		return
	}

	message := context.PostForm("message")
	key := context.PostForm("key")

	res, err := beaufort.Operation(message, key)

	if err != nil {
		context.HTML(http.StatusNotAcceptable, "beaufort.html", err)
		return
	}
	context.HTML(http.StatusOK, "beaufort.html", res)

}
