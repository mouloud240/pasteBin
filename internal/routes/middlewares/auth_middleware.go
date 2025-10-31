package middlewares

import (
	"log"
	"pasteBin/internal/database/repository"
	"pasteBin/pkg/sessions"

	"github.com/gin-gonic/gin"
)

type AuthMiddlewareStruct struct {
	repo *repository.UserRepository
	sm   *sessions.SessionManager
}

func NewAuthMiddleware(repo *repository.UserRepository, sm *sessions.SessionManager) *AuthMiddlewareStruct{
	return &AuthMiddlewareStruct{
		repo:repo,
		sm:sm,
	}
}
func (m *AuthMiddlewareStruct) AuthMiddleware(c *gin.Context ){

	payload, err := m.sm.Get(c.Request)
	if err != nil || payload == nil {
		c.AbortWithStatusJSON(401, gin.H{
			"status":  401,
			"message": "Unauthorized",
			"error":   "Invalid or missing session",
		})
		return
	}

	log.Print("Authenticated user ID: ", payload.UserID)
	user,err:=m.repo.GetUserByID(payload.UserID)

	if user==nil{
		
		c.AbortWithStatusJSON(401, gin.H{
			"status":  401,
			"message": "Unauthorized",
			"error":   "User not found",
		})
		return
	}
	if  err!=nil{
		c.AbortWithStatusJSON(401, gin.H{
			"status":  401,
			"message": "Unauthorized",
			"error":   "User not found",
		})
		return
	}
	c.Set("currentUser", payload)
}
