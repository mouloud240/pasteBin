package handlers

import (
	"log"
	. "pasteBin/internal/database/models"
	"pasteBin/internal/database/repository"
	models "pasteBin/internal/models/auth"
	"pasteBin/pkg/hash"
	"pasteBin/pkg/sessions"

	"github.com/gin-gonic/gin"
)

func setupSessionPayload(user User) *sessions.SessionPayload {
	return &sessions.SessionPayload{
		UserID:  user.ID,
		Email:   user.Email,
		IsAdmin: user.IsAdmin,
	}
}

type AuthHandlers struct {
	userRepo *repository.UserRepository
	sm *sessions.SessionManager
}
func NewAuthHandlers(userRepo *repository.UserRepository, sm *sessions.SessionManager) *AuthHandlers {
	return &AuthHandlers{userRepo: userRepo,sm: sm}
}

// RegisterHandler godoc
// @Summary Register a new user
// @Description Register a new user with username, email and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body models.RegisterModel true "Registration credentials"
// @Success 201 {object} map[string]interface{} "User created successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/register [post]
func (h *AuthHandlers) RegisterHandler(c *gin.Context) {
	var  body *models.RegisterModel
	if err:=c.BindJSON(&body); err!=nil{
   c.JSON(400,gin.H{"status":400, "message":"Bad request", "error":err.Error()})
	 return;
	}
	ctx:=c.Request.Context()
	user,err:=h.userRepo.CreateUser(ctx, body)
	if err!=nil{
		c.Error(err)
				return
	}
	payload:=setupSessionPayload(*user)
	h.sm.Set(c.Request,c.Writer,payload)

	c.JSON(201,gin.H{
		"status":201,
		"data":gin.H{"user":user},
	})

	
}

// LoginHandler godoc
// @Summary Login user
// @Description Authenticate user with email and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body models.BasicLoginModel true "Login credentials"
// @Success 200 {object} map[string]interface{} "Login successful"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 401 {object} map[string]interface{} "Invalid credentials"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/login [post]
func (h *AuthHandlers) LoginHandler(c *gin.Context) {
	var body *models.BasicLoginModel
	if err:=c.BindJSON(&body) ;err!=nil{
		c.JSON(400,gin.H{
			"status":400,
			"message":"Bad request",
			"error":err.Error(),
		})
	}
	ctx := c.Request.Context()
	user,err:=h.userRepo.GetUserByEmail(ctx, body.Email)
	
	if err!=nil{
		c.Error(err)
		return;
	}

	if user.Password==nil{
		log.Print("User password is nil for email: ", body.Email)
		c.JSON(401,gin.H{
			"status":401,
			"message":"Unauthorized",
			"error":"Invalid email or password",
		})
		return

	}
	match,err:=hash.Compare(body.Password,*user.Password)
	if err!=nil{
				c.JSON(500,gin.H{
			"status":500,
			"message":"Internal server error",
			"error":err.Error(),
		})
		return
	}
	if *match==false{

			c.JSON(401,gin.H{
			"status":401,
			"message":"Unauthorized",
			"error":"Invalid email or password",
		})
		return
	}
	payload:=	 setupSessionPayload(*user)
	h.sm.Set(c.Request,c.Writer,payload)

	c.JSON(200,gin.H{
		"status":200,
		"data":gin.H{"user":user},
	})
}
// LogoutHandler godoc
// @Summary Logout user
// @Description Destroy user session
// @Tags Auth
// @Accept json
// @Produce json
// @Success 204 "Logout successful"
// @Security CookieAuth
// @Router /auth/logout [delete]
func (h *AuthHandlers) LogoutHandler(c *gin.Context) {
	h.sm.Destroy(c.Request,c.Writer)
	c.JSON(204,nil)
}
