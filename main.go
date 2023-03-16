package main

import (
	"cryptoapp/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("templates/*.html")

	// main handler
	server.GET("/", handlers.MainHandler)

	// affine routes
	server.GET("/affine", handlers.AffineHandler)
	server.POST("/affine", handlers.AffineHandler)

	// atbash routes
	server.GET("/atbash", handlers.AtbashHandler)
	server.POST("/atbash", handlers.AtbashHandler)

	// autokey routes
	server.GET("/autokey", handlers.AutokeyHandler)
	server.POST("/autokey", handlers.AutokeyHandler)

	// baconian routes
	server.GET("/baconian", handlers.BaconianHandler)
	server.POST("/baconian", handlers.BaconianHandler)

	// beaufort routes
	server.GET("/beaufort", handlers.BeaufortHandler)
	server.POST("/beaufort", handlers.BeaufortHandler)

	// caesar routes
	server.GET("/caesar", handlers.CaesarHandler)
	server.POST("/caesar", handlers.CaesarHandler)

	// columnar transposition routes
	server.GET("/coltrans", handlers.ColTransHandler)
	server.POST("/coltrans", handlers.ColTransHandler)

	// four square routes
	server.GET("/foursquare", handlers.FourSquareHandler)
	server.POST("/foursquare", handlers.FourSquareHandler)

	// playfair routes
	server.GET("/playfair", handlers.PlayfairHandler)
	server.POST("/playfair", handlers.PlayfairHandler)

	// porta routes
	server.GET("/porta", handlers.PortaHandler)
	server.POST("/porta", handlers.PortaHandler)

	// rot13 routes
	server.GET("/rot13", handlers.ROT13Handler)
	server.POST("/rot13", handlers.ROT13Handler)

	// simple substitution routes
	server.GET("/simplesubstitution", handlers.SimpleSubstitutionHandler)
	server.POST("/simplesubstitution", handlers.SimpleSubstitutionHandler)

	// vigenere routes
	server.GET("/vigenere", handlers.VigenereHandler)
	server.POST("/vigenere", handlers.VigenereHandler)

	server.Run(":8080")
}
