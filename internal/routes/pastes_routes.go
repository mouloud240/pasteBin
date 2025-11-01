package routes

import (
	. "pasteBin/internal/handlers"
	"pasteBin/internal/routes/middlewares"

	"github.com/gin-gonic/gin"
)
func BindPaste(r *gin.Engine,p *PastesHandlers, authMiddleware gin.HandlerFunc)  {

  privateRoutes:=r.Group("/pastes").Use(authMiddleware)
	pastesRoutes:=r.Group("/pastes")
	pastesPublicPrivateRoutes:=r.Group("/pastes").Use(middlewares.PublicPrivate(),authMiddleware)
	pastesPublicPrivateRoutes.POST("/",p.CreatePasteHanlder)
	pastesRoutes.GET("/",p.GetPastesHanlder)
	pastesPublicPrivateRoutes.GET("/:pasteId",p.GetPasteByIdHandler)
	privateRoutes.DELETE("/:pasteId",p.DeletePasteHandler)
}
