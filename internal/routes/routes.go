package routes

import (
	"pasteBin/internal/database/repository"
	"pasteBin/internal/handlers"
	"pasteBin/internal/routes/middlewares"
	"pasteBin/pkg/sessions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func SetupRoutes(r *gin.Engine,db *gorm.DB){

	//Init dependencies to inject into handlers (Will probaly use a DI framework later (even implement a simple one))

sessionManager:=sessions.NewSessionManager("auth")
middlwareFunc:=middlewares.NewAuthMiddleware(repository.NewUserRepository(db),sessionManager)
pastesHandlers:=handlers.NewPastesHandlers(repository.NewPastesRepository(db))
authHandlers:=handlers.NewAuthHandlers(repository.NewUserRepository(db),sessionManager)
//Bind Routes
//Pastes Routes
BindPaste(r,pastesHandlers)	

//Auth Routes
BindAuth(r,authHandlers)
//User Routes
BindUser(r,handlers.NewUserHandler(repository.NewUserRepository(db)),middlwareFunc.AuthMiddleware)
	
}
