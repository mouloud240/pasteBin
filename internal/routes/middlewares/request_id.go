package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)
func RequestIdMiddleware(c *gin.Context) {
	
	requestId:=uuid.New().String()
	c.Set("RequestId",requestId)
	c.Header("X-Request-Id",requestId)
	c.Next()
}
