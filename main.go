package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getMethod(c *gin.Context) {
	fmt.Println("goApi.go 'GetMethod' called")
	message := "Get Method Called"
	c.JSON(http.StatusOK, message)
}

func postMethod(c *gin.Context) {
	fmt.Println("goApi.go 'PostMethod' called")
	message := "Post Method called"
	c.JSON(http.StatusAccepted, message)
}

func main() {
	router := gin.Default()
	router.GET("/", getMethod)
	router.POST("/", postMethod)
	listinPort := "4000"
	err := router.Run(":" + listinPort)
	if err != nil {
		return
	}
}
