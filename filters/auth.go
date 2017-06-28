package filters

import (
	"github.com/gin-gonic/gin"
)

func AuthMIddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("X-Auth-Token")
		appid := c.Request.Header.Get("X-Auth-Appid")

		if token == "" {
			c.Abort()
			return
		}

		c.Next()
	}
}
