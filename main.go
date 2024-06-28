package main

import (
	example "observability-example/controllers"
	"observability-example/metrics"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	metrics.StartMetricsServer()
	router.Use(metrics.MetricsMiddleware())

	router.GET("/ping/:id/status", example.ExampleHandler)
	router.GET("/ping/:id/info", example.ExampleHandler)

	router.Run()
}
