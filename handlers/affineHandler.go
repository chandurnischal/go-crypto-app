package handlers

import (
	"cryptoapp/pkg/affine"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
