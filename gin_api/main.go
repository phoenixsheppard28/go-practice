package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func handle(c *gin.Context) {
	c.Writer.Write([]byte("Hello world"))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load environment variables")
	}
	router := gin.Default()
	router.GET("/handle", handle)

	router.Run(os.Getenv("PORT"))
}
