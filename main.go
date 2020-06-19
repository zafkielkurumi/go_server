package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("start")
	app := gin.Default()
	app.GET("/hello", func(c *gin.Context) {
		fmt.Printf("firstname, lastname")
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // 是 c.Request.URL.Query().Get("lastname") 的简写
		fmt.Printf(firstname, lastname)
		c.JSON(200, gin.H{"message": "hello"})
	})
	app.Run(":5000")
}
