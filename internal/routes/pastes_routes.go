package routes

import (
.	"pasteBin/internal/handlers" 

	"github.com/gin-gonic/gin"
)
func BindPaste(r *gin.Engine){
	
	p :=PastesHandlers{}
	pastesRoutes:=r.Group("/pastes")
	pastesRoutes.POST("/",p.CreatePasteHanlder)
	pastesRoutes.GET("/",p.GetPastesHanlder)
	pastesRoutes.GET("/:pasteId",p.GetPasteByIdHandler)
	pastesRoutes.DELETE("/:pasteId",p.DeletePasteHandler)
}
