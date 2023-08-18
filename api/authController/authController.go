package authController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"testproject/internal/services/authService/contracts"
	"testproject/internal/services/authService/dto"
)

type IAuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	MiddleAuth(ctx *gin.Context)
}

type AuthController struct {
	service contracts.IAuthService
}

func NewAuthController(service contracts.IAuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (c *AuthController) MiddleAuth(ctx *gin.Context) {
	tokenStr := ctx.GetHeader("Authorization")
	token, err := jwt.ParseWithClaims(tokenStr, &dto.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	if claims, ok := token.Claims.(dto.Claims); ok && token.Valid {
		fmt.Printf("%v %v", claims.UserID, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
	}
	ctx.Next()
}

func (c *AuthController) Login(ctx *gin.Context) {
	var user dto.LoginRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := c.service.Login(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}
func (c *AuthController) Register(ctx *gin.Context) {
	var user dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := c.service.Register(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}
