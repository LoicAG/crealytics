package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "google.golang.org/api/compute/v1"
)

func main() {
    r := gin.Default()

    r.GET("/healthcheck", HealthCheck)
    r.POST("/v1/instances/create", CreateInstance)

    r.Run(":8080")
}

func HealthCheck(c *gin.Context) {
    c.String(200, "")
}

func CreateInstance(c *gin.Context) {
    //TODO
}
