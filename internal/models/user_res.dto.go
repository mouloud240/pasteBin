package models

import "time"
type UserResponseDto struct {
	Id string `json:"id"`
	Username     string  `json:"username"`
	Email        string  `json:"email"`
	Password *string  `json:"-"`
	IsAdmin 		bool `json:"isAdmin"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
