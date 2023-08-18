package contracts

import (
	"testproject/internal/services/authService/dto"
)

type IAuthService interface {
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
	Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error)
}
