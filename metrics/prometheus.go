package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartMetricsServer(router *gin.Engine) {
	registerRequestMetrics(prometheus.DefaultRegisterer)

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
