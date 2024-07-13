package main

import (
	"flag"
	example "observability-example/controllers"
	"observability-example/metrics"

	"github.com/gin-gonic/gin"
)

func main() {
	var successChance = flag.Int("success-chance", 90, "Ratio of requests that will succeed with non 4XX or 5XX code")
	var fastChance = flag.Int("fast-chance", 95, "Ratio of requests that will be faster than -fast-duration threshold")
	var fastDuration = flag.Int("fast-duration", 300, "Maximum duration of fast requests")
	flag.Parse()

	router := gin.Default()

	metrics.StartMetricsServer(router)
	router.Use(metrics.MetricsMiddleware())

	router.GET("/ping/:id/status", example.ExampleHandler(*successChance, *fastChance, *fastDuration))
	router.GET("/ping/:id/info", example.ExampleHandler(*successChance, *fastChance, *fastDuration))

	router.Run()
}
