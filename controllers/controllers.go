package controllers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ExampleHandler(c *gin.Context) {
	// simulate request taking time
	time.Sleep(getRandomLatencyMs())

	// simulate a random status code
	c.JSON(getRandomStatusCode(), gin.H{
		"response": "pong",
	})
}

// getRandomLatencyMs generates a random request latency between 100ms and 1000ms
func getRandomLatencyMs() time.Duration {
	return time.Duration(generateRandomInt(100, 1000) * int(time.Millisecond))
}

// getRandomStatusCode generates a random status code that's 200 50% of the time
func getRandomStatusCode() int {
	codes := []int{
		http.StatusBadRequest,
		http.StatusOK,
		http.StatusGatewayTimeout,
	}

	chance200 := generateRandomInt(1, 100)
	if chance200 <= 50 {
		return http.StatusOK
	}

	return codes[generateRandomInt(0, len(codes)-1)]
}

// generateRandomInt generates a random integer between min and max
// note that in go 1.20, there is no need to call rand.seed,
// but in older versions it should be called
func generateRandomInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
