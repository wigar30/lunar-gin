package model

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,alphanum"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type JwtClaims struct {
	ID   string `json:"id"`
	Role int    `json:"role"`
	jwt.StandardClaims
}

type AuthUseCaseInterface interface {
	Login(c *gin.Context, req *LoginRequest) (*LoginResponse, error)
}
