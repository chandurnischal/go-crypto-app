package handlers

import (
	"cryptoapp/pkg/affine"
	"cryptoapp/pkg/atbash"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "main.html", nil)
}

func AffineHandler(context *gin.Context) {
	if context.Request.Method == "GET" {
		context.HTML(http.StatusOK, "affine.html", nil)
		return
	}
	m := context.PostForm("multiplier")
	o := context.PostForm("offset")
	message := context.PostForm("message")
	action := context.PostForm("operation")
	res, err := affine.Operation(m, o, message, action)

	if err != nil {
		context.HTML(http.StatusNotAcceptable, "affine.html", res)
		return
	}
	context.HTML(http.StatusOK, "affine.html", res)
}

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
