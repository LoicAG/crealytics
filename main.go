package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/healthcheck", HealthCheck)
    r.POST("/v1/instances/create", CreateInstance)

    r.Run(":8080")
}

func HealthCheck(c *gin.Context) {
    //TODO
}

func CreateInstance(c *gin.Context) {
    //TODO
}
