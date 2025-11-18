package handlers

import (
	"pasteBin/internal/database/repository"
	"pasteBin/pkg/sessions"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{
	userRepo *repository.UserRepository
}

func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{userRepo:userRepo}
}


// User Handler Doc
// @Summary Gets the current logged in user info
// @Accept json
// @Produce json  
// @Success 200  {object} 	models.UserResponseDto
// @Failure 401 {object} map[string]interface{} "unauthorized"
// @Router /user/me [get]
func (h *UserHandler) CurrentUserHandler(c *gin.Context) {
  currentUser,exists:= c.Get("currentUser")
	if !exists{
		c.JSON(401,gin.H{"status":401,"error":"unauthorized"})
		return;
	}
	parsedUser:=currentUser.(*sessions.SessionPayload)
	ctx := c.Request.Context()
	user,err:=h.userRepo.GetUserByID(ctx, parsedUser.UserID)
	if err!=nil{
		c.JSON(500,gin.H{"status":500,"error":"something went wrong"})
		return;
	}
	c.JSON(200,gin.H{"status":200,"user":user})
}


