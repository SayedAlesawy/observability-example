package main

import (
	example "observability-example/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", example.ExampleHandler)

	router.Run()
}
