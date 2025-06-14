package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoleMiddleware(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		rolesInterface, exists := c.Get("roles")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Role tidak ditemukan!"})
			return
		}

		userRoles, ok := rolesInterface.([]string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Format role tidak valid!"})
			return
		}

		roleMap := make(map[string]bool)
		for _, role := range userRoles {
			roleMap[role] = true
		}

		for _, required := range requiredRoles {
			if _, ok := roleMap[required]; !ok {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied!"})
				return
			}
		}
		c.Next()
	}
}
