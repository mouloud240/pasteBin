package routes

import "github.com/gin-gonic/gin"
func BindUser(r *gin.Engine)  {
	
	r.Group("/user")
}
