package models

import "time"


type CreatePaste struct {
	Content     string `json:"content" binding:"required"`
	Password *string `json:"password" binding:"omitempty,min=8"`
	MaxViews *int  `json:"maxViews" binding:"omitempty,min=1"`
	Expires_at  *time.Time  `json:"expires_at"  binding:"omitempty ,datetime,gt"`
}
