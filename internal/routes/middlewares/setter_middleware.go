package middlewares

import "github.com/gin-gonic/gin"
/*
 @Param key string - The key under which the value will be stored in the Gin context.
 @Param value any - The value to be stored in the Gin context.
 @Return gin.HandlerFunc - A Gin middleware function that sets the specified key-value pair in the context.
 @Description This middleware sets a specified key-value pair in the Gin context for each incoming request, this works similart to nestjs set Metadata funcction with decorators.
*/
func SetMiddleware(key string, value any)gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(key, value)
		c.Next()
	}
}
