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
	//Global Middlewares
	r.Use(middlewares.RequestIdMiddleware)
	r.Use(middlewares.ExceptionFilterMiddleware)

sessionManager:=sessions.NewSessionManager("auth")
userRepo:=repository.NewUserRepository(db)
middlwareFunc:=middlewares.NewAuthMiddleware(repository.NewUserRepository(db),sessionManager)
pastesHandlers:=handlers.NewPastesHandlers(repository.NewPastesRepository(db))
authHandlers:=handlers.NewAuthHandlers(userRepo,sessionManager)
//Bind Routes
//Pastes Routes
BindPaste(r,pastesHandlers,middlwareFunc.AuthMiddleware)	

//Auth Routes
BindAuth(r,authHandlers)
//User Routes
BindUser(r,handlers.NewUserHandler(userRepo),middlwareFunc.AuthMiddleware)
	
}
