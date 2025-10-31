package models
type RegisterModel struct{

	UserName string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}
