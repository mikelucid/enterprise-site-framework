package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logger() gin.HandlerFunc {
	l, _ := zap.NewProduction()
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		l.Info("http_request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", time.Since(start)),
			zap.String("request_id", c.Writer.Header().Get(HeaderRequestID)),
		)
	}
}
