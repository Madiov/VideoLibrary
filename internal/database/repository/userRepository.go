package repository

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique;->" binding:"min=2,max=30"`
	PassWord string `gorm:"<-:update"`
	Email    string `gorm:"<-:update" validate:"email"`
}
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
func (u *UserRepository) ReadUserByUsername(username string) (*User, error) {
	var user User
	res := u.db.Where("username = ?", username).First(&user)
	if res.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
