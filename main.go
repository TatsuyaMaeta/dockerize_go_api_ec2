package main

import (
	"go_docker_ec2/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "nya~~",
		})
	})

	r.GET("/products", func(c *gin.Context) {
		var product models.Product

		models.DB.First(&product, 1)
		c.JSON(200, product)
	})
	r.Run()
}
