package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mikelucid/enterprise-site-framework/pkg/ai"
)

func AIRecommendations(engine *ai.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		recs := engine.Recommend(c.Request.URL.Path)
		c.Writer.Header().Set("X-AI-Recommendations", strings.Join(recs, ","))
	}
}
