package handlers

import (
	. "pasteBin/internal/database/repository"
	. "pasteBin/internal/models"
	"pasteBin/pkg/sessions"
	"pasteBin/pkg/sessions/extractors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PastesHandlers  struct{
	pastesRepo *PastesRepository
}
func NewPastesHandlers(pastesRepo *PastesRepository) *PastesHandlers {
	return &PastesHandlers{pastesRepo:pastesRepo}
}


func (h *PastesHandlers) CreatePasteHanlder(c *gin.Context){
	var body CreatePaste;
	
	
paste_repo:=h.pastesRepo
  if err:=c.BindJSON(&body);err!=nil{
		c.JSON(400,gin.H{"error":err.Error()})
		return;
	}
	isAuthenticated,exists:=c.Get("isAuthenticated")
	var userId *uint;
	if exists && isAuthenticated.(bool){
		currentUser,exists:=c.Get("currentUser")
		if exists{
			parsedUser:=currentUser.(*sessions.SessionPayload)
			userId=&parsedUser.UserID
		}

	}
	ctx := c.Request.Context()
  createdPaste,err:=paste_repo.CreatePaste(ctx, body,userId)

	if (err!=nil){
		c.JSON(500,gin.H{"message":"something went wrong","error":err.Error()})
	}

	
		
c.JSON(201,createdPaste)
	
}
func (h *PastesHandlers) GetPastesHanlder(c *gin.Context){
	paste_repo:=h.pastesRepo
	pageRaw:=c.Query("page")
	limitRaw:=c.Query("limit")


	 page,err:=strconv.Atoi(pageRaw)
	 if err!=nil{
		page=1
	 }
	 limit,err:=strconv.Atoi(limitRaw)
	 if err!=nil{
		limit=10
	 }
	
	 ctx := c.Request.Context()
	 pastes,err:=paste_repo.GetPastes(ctx, &page,&limit);
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
ctx := c.Request.Context()
paste,err:=paste_repo.GetPaste(ctx, pasteId,password)
if err!=nil{
	c.JSON(400,gin.H{"status":400,"error":err.Error()})
	return
}
c.JSON(200,gin.H{"status":200,"paste":paste})
}
func (h *PastesHandlers) DeletePasteHandler(c *gin.Context){

paste_repo:=h.pastesRepo
pasteId:=	c.Param("pasteId")
user,exists:=extractors.ExtractUserSessionPayload(c)

if !exists{
	c.JSON(401,gin.H{"status":401,"error":"unauthorized"})
	return;
}

if pasteId==""{
	c.JSON(400,gin.H{"status":400,"error":"paste Id must be provided"})
	return;
}
ctx := c.Request.Context()
if err:= paste_repo.DeletePaste(ctx, pasteId,user.UserID);err!=nil{
	c.JSON(500,gin.H{"status":500,"error":err.Error()})
	return;
}
c.JSON(204,gin.H{})
}

