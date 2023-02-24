package main

import (
	"cryptoapp/pkg/autokey"
	"fmt"
)

func main() {
	auto, _ := autokey.New("german")

	fmt.Println(auto.Decrypt("leiavrwymjrnxydrtzyfnsxhtqi"))
	// server := gin.Default()
	// server.LoadHTMLGlob("templates/*.html")
	// server.GET("/", handlers.MainHandler)
	// server.GET("/affine", handlers.AffineHandler)
	// server.POST("/affine", handlers.AffineHandler)
	// server.GET("/atbash", handlers.AtbashHandler)
	// server.POST("/atbash", handlers.AtbashHandler)
	// server.Run(":8080")
}
