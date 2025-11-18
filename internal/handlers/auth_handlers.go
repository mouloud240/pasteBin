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

func (h *AuthHandlers) RegisterHandler(c *gin.Context) {
	var  body *models.RegisterModel
	if err:=c.BindJSON(&body); err!=nil{
   c.JSON(400,gin.H{"status":400, "message":"Bad request", "error":err.Error()})
	 return;
	}
	user,err:=h.userRepo.CreateUser(body)
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

func (h *AuthHandlers) LoginHandler(c *gin.Context) {
	var body *models.BasicLoginModel
	if err:=c.BindJSON(&body) ;err!=nil{
		c.JSON(400,gin.H{
			"status":400,
			"message":"Bad request",
			"error":err.Error(),
		})
	}
	user,err:=h.userRepo.GetUserByEmail(body.Email)
	
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
log.Print("Hashed Pass",*user.Password)
		log.Print("User provided",body.Password)

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
func (h *AuthHandlers) LogoutHandler(c *gin.Context) {
	h.sm.Destroy(c.Request,c.Writer)
	c.JSON(204,nil)
}
