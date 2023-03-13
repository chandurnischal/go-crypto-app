package handlers

import (
	"cryptoapp/ciphers/caesar"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CaesarHandler(context *gin.Context) {
	if context.Request.Method == "GET" {
		context.HTML(http.StatusOK, "caesar.html", nil)
		return
	}
	message := context.PostForm("message")
	off := context.PostForm("offset")
	operation := context.PostForm("operation")
	offset, err := strconv.Atoi(off)

	if err != nil {
		context.HTML(http.StatusNotAcceptable, "caesar.html", err)
		return
	}
	res, err := caesar.Operation(offset, message, operation)

	if err != nil {
		context.HTML(http.StatusNotAcceptable, "caesar.html", err)
		return
	}

	context.HTML(http.StatusOK, "caesar.html", res)

}
