package handlers

import (
	"cryptoapp/ciphers/simplesubstitution"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SimpleSubstitutionHandler(context *gin.Context) {
	if context.Request.Method == "GET" {
		context.HTML(http.StatusOK, "simplesubstitution.html", nil)
		return
	}
	message := context.PostForm("message")
	key := context.PostForm("key")
	operation := context.PostForm("operation")

	res, err := simplesubstitution.Operation(message, key, operation)

	if err != nil {
		context.HTML(http.StatusMethodNotAllowed, "simplesubstitution.html", err)
		return
	}

	context.HTML(http.StatusOK, "simplesubstitution.html", res)
}
