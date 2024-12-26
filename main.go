package main

import (
    "github.com/gin-gonic/gin"
	
)

func main() {
    // Create a new Gin router
    router := gin.Default()

    // Define a GET endpoint
    router.GET("/hello", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, World!",
        })
    })

    // Start the server on port 8082
    router.Run(":8082")
}
