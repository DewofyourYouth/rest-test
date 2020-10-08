package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "it works!",
		})
	})

	r.GET("/hello/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "Hello %s!", name)
	})

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
