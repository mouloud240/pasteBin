package routes

import (
	"pasteBin/internal/handlers"

	"github.com/gin-gonic/gin"
)
func BindAuth(r *gin.Engine, a *handlers.AuthHandlers) {
	authRoutes := r.Group("/auth")
authRoutes.POST("/register", a.RegisterHandler) 
authRoutes.POST("/login", a.LoginHandler)
	authRoutes.POST("/logout", a.LogoutHandler)
}
