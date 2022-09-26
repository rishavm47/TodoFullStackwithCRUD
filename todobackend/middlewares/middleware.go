package middlewares

import "github.com/gin-gonic/gin"

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ctx := c.Request.Context()
		// fields := make(map[string]interface{})
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		if c.Request.Method == "OPTIONS" {
			// fields["Method"] = "OPTIONS"
			// logger.LogError(&ctx, "Cors Middleware Failed", &fields)
			c.AbortWithStatus(204)
			return
		}
		// fields["CorsMiddleare"] = "Passed"
		// logger.LogInfo(&ctx, "Cors Middleware Passes", &fields)
		c.Next()
	}
}
