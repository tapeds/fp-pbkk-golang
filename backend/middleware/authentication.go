package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tapeds/fp-pbkk-golang/dto"
	"github.com/tapeds/fp-pbkk-golang/service"
	"github.com/tapeds/fp-pbkk-golang/utils"
)

type AuthConfig struct {
	RequiredRole *string
}
type Option func(*AuthConfig)

func WithRole(role string) Option {
	return func(config *AuthConfig) {
		config.RequiredRole = &role
	}
}

func Authenticate(jwtService service.JWTService, opts ...Option) gin.HandlerFunc {
	config := &AuthConfig{
		RequiredRole: nil,
	}

	for _, opt := range opts {
		opt(config)
	}

	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_TOKEN_NOT_FOUND, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if !strings.Contains(authHeader, "Bearer ") {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_TOKEN_NOT_VALID, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_TOKEN_NOT_VALID, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if !token.Valid {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_DENIED_ACCESS, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId, err := jwtService.GetUserIDByToken(tokenString)
		if err != nil {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userRole, err := jwtService.GetUserRoleByToken(tokenString)
		if err != nil {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if config.RequiredRole != nil && *config.RequiredRole != userRole {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, "Insufficient privileges", nil)
			ctx.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}

		ctx.Set("token", tokenString)
		ctx.Set("user_id", userId)
		ctx.Set("role", userRole)
		ctx.Next()
	}
}
