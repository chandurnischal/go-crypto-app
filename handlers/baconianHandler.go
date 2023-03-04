package handlers

import (
	"cryptoapp/ciphers/baconian"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BaconianHandler(context *gin.Context) {
	if context.Request.Method == "GET" {
		context.HTML(http.StatusOK, "baconian.html", nil)
		return
	}
	m := context.PostForm("message")
	o := context.PostForm("operation")
	res, err := baconian.Operation(m, o)

	if err != nil {
		context.HTML(http.StatusNotAcceptable, "baconian.html", err)
		return
	}
	context.HTML(http.StatusOK, "baconian.html", res)
}
