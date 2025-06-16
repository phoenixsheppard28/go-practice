package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Input struct {
	Num1     *int   `json:"num1" binding:"required"`
	Num2     *int   `json:"num2" binding:"required"`
	Operator string `json:"operator" binding:"required"`
}

func handle(c *gin.Context) {
	var inp Input
	if err := c.ShouldBindBodyWithJSON(&inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	num1 := *inp.Num1
	num2 := *inp.Num2
	oper := inp.Operator

	var res int
	switch oper {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2

	case "/":
		if num2 == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "divide by zero",
			})
		}
		res = num1 / num2
	case "*":
		res = num1 * num2

	case "%":
		if num2 == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "mod by zero",
			})
		}
		res = num1 % num2

	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid operator, only support + - / * %",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": res,
	})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load environment variables")
	}
	router := gin.Default()
	router.POST("/handle", handle)

	router.Run(os.Getenv("PORT"))
}
