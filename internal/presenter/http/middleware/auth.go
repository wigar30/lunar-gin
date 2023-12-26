package middleware

import (
	"lunar/internal/model"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthMiddleware struct {
	userRepo model.UserRepositoryInterface

	config *model.EnvConfigs
}

func NewAuthMiddleware(userRepo model.UserRepositoryInterface, config *model.EnvConfigs) *AuthMiddleware {
	return &AuthMiddleware{
		userRepo: userRepo,
		config:   config,
	}
}

func (am *AuthMiddleware) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")
		if bearerToken == "" {
			model.OnError(c, &model.ErrorResponse{
				Code:    401,
				Message: "Unauthorized",
			})
			return
		}

		token := strings.TrimPrefix(bearerToken, "Bearer ")

		jwtClaims, err := jwt.ParseWithClaims(token, &model.JwtClaims{}, func(jwtToken *jwt.Token) (interface{}, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, &model.ErrorResponse{
					Code:    401,
					Message: "Invalid Token",
				}
			}
			return []byte(am.config.JwtSecret), nil
		})
		if err != nil {
			model.OnError(c, &model.ErrorResponse{
				Code:    401,
				Message: "Invalid Token",
			})
			return
		}

		claims, ok := jwtClaims.Claims.(*model.JwtClaims)
		claimID, err := strconv.ParseInt(claims.ID, 10, 64)
		if err != nil {
			model.OnError(c, &model.ErrorResponse{
				Code:    401,
				Message: "Invalid Token",
			})
			return
		}

		user, err := am.userRepo.GetUserByID(claimID, false)
		if err != nil {
			model.OnError(c, &model.ErrorResponse{
				Code:    401,
				Message: "Invalid Token",
			})
			return
		}

		if user.StatusID != "1" {
			model.OnError(c, &model.ErrorResponse{
				Code:    401,
				Message: "User not verified",
			})
			return
		}

		if !ok || !jwtClaims.Valid {
			model.OnError(c, &model.ErrorResponse{
				Code:    401,
				Message: "Invalid Token",
			})
			return
		}

		c.Set("claims", claims)

		c.Next()
	}
}
