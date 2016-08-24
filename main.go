package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Index",
	})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("html/*")
	router.GET("/",Index)
	router.Run(":8080")
	fmt.Println("server start port 8080.")
}
