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


// CreatePasteHanlder godoc
// @Summary Create a new paste
// @Description Create a new paste with optional password protection, expiration, and view limit
// @Tags Pastes
// @Accept json
// @Produce json
// @Param body body models.CreatePaste true "Paste data"
// @Success 201 {object} map[string]interface{} "Paste created successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /pastes [post]
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
// GetPastesHanlder godoc
// @Summary Get list of pastes
// @Description Get paginated list of all public pastes
// @Tags Pastes
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} map[string]interface{} "List of pastes"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /pastes [get]
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
// GetPasteByIdHandler godoc
// @Summary Get a specific paste
// @Description Get paste by ID, requires password if paste is protected
// @Tags Pastes
// @Accept json
// @Produce json
// @Param pasteId path string true "Paste ID"
// @Param password query string false "Password for protected paste"
// @Success 200 {object} map[string]interface{} "Paste data"
// @Failure 400 {object} map[string]interface{} "Bad request or invalid password"
// @Router /pastes/{pasteId} [get]
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
// DeletePasteHandler godoc
// @Summary Delete a paste
// @Description Delete a paste (only owner can delete)
// @Tags Pastes
// @Accept json
// @Produce json
// @Param pasteId path string true "Paste ID"
// @Success 204 "Paste deleted successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Security CookieAuth
// @Router /pastes/{pasteId} [delete]
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

