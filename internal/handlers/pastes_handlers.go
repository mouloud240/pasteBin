package handlers

import (
	. "pasteBin/internal/database/repository"
	. "pasteBin/internal/models"

	"github.com/gin-gonic/gin"
)

type PastesHandlers  struct{
	pastesRepo *PastesRepository
}
func NewPastesHandlers(pastesRepo *PastesRepository) *PastesHandlers {
	return &PastesHandlers{pastesRepo:pastesRepo}
}
//TODO: Figure out how to handle injection of the repository properly


func (h *PastesHandlers) CreatePasteHanlder(c *gin.Context){
	var body CreatePaste;
	
paste_repo:=h.pastesRepo
  if err:=c.BindJSON(&body);err!=nil{
		c.JSON(400,gin.H{"error":err.Error()})
		return;
	}
  createdPaste,err:=paste_repo.CreatePaste(body)
	if (err!=nil){
		c.JSON(500,gin.H{"message":"something went wrong","error":err.Error()})
	}

	
		
c.JSON(201,createdPaste)
	
}
func (h *PastesHandlers) GetPastesHanlder(c *gin.Context){
	paste_repo:=h.pastesRepo
	 pastes,err:=paste_repo.GetPastes();
	 if err!=nil{
		 c.JSON(500,gin.H{"message":"something went wrong","error":err.Error()})
	 }
	 c.JSON(200,gin.H{"status":200,"pastes":pastes})

}
func (h *PastesHandlers) GetPasteByIdHandler(c *gin.Context){

paste_repo:=h.pastesRepo
pasteId:=	c.Param("pasteId")
password:=c.Query("password")

if pasteId==""{
	c.JSON(400,gin.H{"status":400,"error":"paste Id must be provided"})
	return;
}
paste,err:=paste_repo.GetPaste(pasteId,password)
if err!=nil{
	c.JSON(400,gin.H{"status":400,"error":err.Error()})
	return
}
c.JSON(200,gin.H{"status":200,"paste":paste})
}
func (h *PastesHandlers) DeletePasteHandler(c *gin.Context){

paste_repo:=h.pastesRepo
pasteId:=	c.Param("pasteId")

if pasteId==""{
	c.JSON(400,gin.H{"status":400,"error":"paste Id must be provided"})
	return;
}
if err:= paste_repo.DeletePaste(pasteId);err!=nil{
	c.JSON(500,gin.H{"status":500,"error":err.Error()})
	return;
}
c.JSON(204,gin.H{})
}

