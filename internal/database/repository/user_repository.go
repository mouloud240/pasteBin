package repository

import (
	"context"
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

func (r *UserRepository) CreateUser(ctx context.Context, createUser *RegisterModel) (*models.User, *exception.AppError) {
	user := &models.User{
		Username: createUser.UserName,
		Email:    createUser.Email,
		Password: &createUser.Password,
	}
	err := gorm.G[models.User](r.db).Create(ctx,user)
	if err!= nil {
		return nil, exception.NewInternalServerError("Failed to create user", err.Error)
	}
	return user, nil
}
func (r *UserRepository) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	result := r.db.WithContext(ctx).First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &user, nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	result := r.db.WithContext(ctx).Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	result := r.db.WithContext(ctx).Where("email = ?", email).First(&user)
if result.RowsAffected == 0 {
		return nil, exception.NewNotFoundError("User not found")
	}

	if result.Error != nil {
		return nil, exception.NewAppError("Failed to get user by email", 500, result.Error)
	}
		
	return &user, nil
}

