package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "main.html", nil)
}
