package repository

import (
	"pasteBin/internal/database/models"
	"pasteBin/pkg/exception"

	. "pasteBin/internal/models/auth"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(createUser *RegisterModel) (*models.User, *exception.AppError) {
	user := &models.User{
		Username: createUser.UserName,
		Email:    createUser.Email,
		Password: &createUser.Password,
	}
	result := r.db.Create(user)
	if result.Error != nil {
		return nil, exception.NewInternalServerError("Failed to create user", result.Error)
	}
	return user, nil
}
func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := r.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)
if result.RowsAffected == 0 {
		return nil, exception.NewNotFoundError("User not found")
	}

	if result.Error != nil {
		return nil, exception.NewAppError("Failed to get user by email", 500, result.Error)
	}
		
	return &user, nil
}

