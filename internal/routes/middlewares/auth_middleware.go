package middlewares

import (
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
	public,	exists := c.Get("isPublic")
	if exists && public.(bool) {
		c.Next()
		return
	}

	public_private,	exists := c.Get("isPublicPrivate")

	payload, err := m.sm.Get(c.Request)
	if err != nil || payload == nil {
		if public_private!=nil && public_private.(bool){
			c.Set("isAuthenticated", false)
			c.Next()
			return
		}

		c.AbortWithStatusJSON(401, gin.H{
			"status":  401,
			"message": "Unauthorized",
			"error":   "Invalid or missing session",
		})
		return
	}

	ctx := c.Request.Context()
	user,err:=m.repo.GetUserByID(ctx, payload.UserID)

	if user==nil{
		if public_private!=nil && public_private.(bool){
			c.Set("isAuthenticated", false)
			c.Next()
			return
		}
		
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
	c.Set("isAuthenticated", true)
	c.Set("currentUser", payload)
	c.Next()
}
