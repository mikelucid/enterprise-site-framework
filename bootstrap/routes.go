package bootstrap

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikelucid/enterprise-site-framework/pkg/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(
		middleware.RequestID(),
		middleware.Logger(),
		middleware.Recovery(),
		middleware.CORS(),
	)
	r.GET("/healthz", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })
	return r
}
