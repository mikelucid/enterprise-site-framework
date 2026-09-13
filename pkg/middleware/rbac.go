package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikelucid/enterprise-site-framework/pkg/auth"
)

func RBAC(enforcer *auth.RBAC, permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, _ := c.Get("roles")
		roleList, ok := roles.([]string)
		if !ok || !enforcer.HasPermission(roleList, permission) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}
		c.Next()
	}
}
