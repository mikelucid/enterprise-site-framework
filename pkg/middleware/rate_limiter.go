package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type bucket struct {
	tokens float64
	last   time.Time
}

func RateLimiter(rate, burst float64) gin.HandlerFunc {
	var mu sync.Mutex
	buckets := map[string]*bucket{}
	return func(c *gin.Context) {
		key := c.ClientIP()
		mu.Lock()
		b := buckets[key]
		if b == nil {
			b = &bucket{tokens: burst, last: time.Now()}
			buckets[key] = b
		}
		now := time.Now()
		elapsed := now.Sub(b.last).Seconds()
		b.tokens = minf(burst, b.tokens+elapsed*rate)
		b.last = now
		if b.tokens < 1 {
			mu.Unlock()
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
			return
		}
		b.tokens -= 1
		mu.Unlock()
		c.Next()
	}
}

func minf(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
