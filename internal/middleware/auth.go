package middleware

import (
	"fmt"
	"mkpticket/infrastructure/config"
	"mkpticket/internal/dto"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authentication(cfg *config.JWTConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse("unauthorization"))
			ctx.Abort()
			return
		}

		tokenPart := strings.Split(authHeader, " ")
		if len(tokenPart) != 2 || tokenPart[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse("invalid authorization format"))
			ctx.Abort()
			return
		}

		tokenStr := tokenPart[1]
		jwtSecret := []byte(cfg.Secret)
		claim := &dto.TokenClaims{}

		token, err := jwt.ParseWithClaims(tokenStr, claim, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse("invalid token"))
			ctx.Abort()
			return
		}

		ctx.Set("user", claim)
		ctx.Next()
	}

}

func RoleAccessMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userClaims, exist := ctx.Get("user")
		if !exist {
			ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse("user not found"))
			ctx.Abort()
			return
		}

		user := userClaims.(*dto.TokenClaims)
		role := user.Role

		for _, v := range allowedRoles {
			if role == v {
				ctx.Next()
				return
			}
		}

		ctx.JSON(http.StatusForbidden, dto.ErrorResponse("insufficient role"))
		ctx.Abort()
	}
}
