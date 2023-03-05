package handlers

import (
	"cryptoapp/ciphers/foursquare"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FourSquareHandler(context *gin.Context) {
	if context.Request.Method == "GET" {
		context.HTML(http.StatusOK, "foursquare.html", nil)
		return
	}

	message := context.PostForm("message")
	key1 := context.PostForm("key 1")
	key2 := context.PostForm("key 2")
	operation := context.PostForm("operation")

	res, err := foursquare.Operation(message, key1, key2, operation)

	if err != nil {
		context.HTML(http.StatusNotAcceptable, "foursquare.html", err)
		return
	}

	context.HTML(http.StatusOK, "foursquare.html", res)

}
