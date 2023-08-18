package authService

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"testproject/internal/database"
	"testproject/internal/database/repository"
	"testproject/internal/services/authService/dto"
	"testproject/internal/utils"
	"time"
)

type AuthService struct {
	uow     *database.UnitOfWork
	userIds map[string]uint
}

func NewAuthService(uow *database.UnitOfWork) *AuthService {
	return &AuthService{
		uow:     uow,
		userIds: make(map[string]uint),
	}
}
func (a *AuthService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	dbUser, err := a.uow.UserRepository.ReadUserByUsername(req.UserName)
	if err != nil {
		return nil, err
	}
	IsValid, err := utils.ComparePasswords(req.PassWord, dbUser.PassWord)
	if err != nil || !IsValid {
		return nil, err
	}
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &dto.Claims{
		Username: dbUser.UserName,
		UserID:   dbUser.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("Super-Secret-Stuff"))
	if err != nil {
		return nil, err
	}
	result := dto.LoginResponse{
		Token:   tokenString,
		Expires: expirationTime,
	}
	return &result, nil
}
func (a *AuthService) Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	dbUser, err := a.uow.UserRepository.ReadUserByUsername(req.UserName)
	if err != nil {
		return nil, err
	}
	if dbUser != nil {
		return nil, errors.New("user Already Exists")
	}
	hashedPassword, err := utils.HashPassword(req.PassWord)
	if err != nil {
		return nil, err
	}
	user := repository.User{
		UserName: req.UserName,
		PassWord: hashedPassword,
		Email:    req.Email,
	}
	err = a.uow.BaseRepository.Create(&user)
	if err != nil {
		return nil, err
	}
	result := dto.RegisterResponse{UserName: user.UserName}
	return &result, nil
}
