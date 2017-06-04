package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "google.golang.org/api/compute/v1"
)

func main() {
    r := gin.Default()

    computeService, err := compute.New(&http.Client{}) //TODO configure the http.Client

    if err != nil {
        panic(err)
    }

    r.GET("/healthcheck", HealthCheck)
    r.POST("/v1/instances/create", CreateInstance(computeService))

    r.Run(":8080")
}

func HealthCheck(c *gin.Context) {
    c.String(200, "")
}

func CreateInstance(computeService *compute.Service) gin.HandlerFunc {
    fn := func(c *gin.Context) {
        //TODO call computeService.Whatever
    }

    return gin.HandlerFunc(fn)
}
