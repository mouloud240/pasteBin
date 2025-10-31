package routes

import (
	"pasteBin/internal/handlers"

	"github.com/gin-gonic/gin"
)
func BindUser(r *gin.Engine, userHandlers *handlers.UserHandler, middlwareFunc func (c *gin.Context) )  {
	
	userGroup:=r.Group("/user").Use(middlwareFunc)
	userGroup.GET("/me",userHandlers.CurrentUserHandler)

}
