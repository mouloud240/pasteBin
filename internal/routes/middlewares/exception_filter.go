package middlewares

import (
	"errors"
	"net/http"
	"pasteBin/pkg/exception"

	"github.com/gin-gonic/gin"
)
func ExceptionFilterMiddleware(c *gin.Context)  {
	c.Next()
if len(c.Errors)>0{
	  err := c.Errors.Last().Err

              var appErr *exception.AppError
              if errors.As(err, &appErr) {
                  c.JSON(appErr.Status, gin.H{
                      "status":  appErr.Status,
                      "message": appErr.Message,
											"error":appErr.Err,
                  })
                  return
              }

              // Fallback
              c.JSON(http.StatusInternalServerError, gin.H{
                  "status":  http.StatusInternalServerError,
                  "message": "Internal server error",
              })
}
	
	
}
