package token

import (
	"errors"
	"fmt"
	"mkpticket/infrastructure/config"
	"mkpticket/internal/dto"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId string, name, email, role string, exp time.Duration, cfg *config.JWTConfig) (string, error) {
	jwtSecret := []byte(cfg.Secret)
	jwtExp := time.Now().Add(exp * time.Hour)

	tokenCLaim := &dto.TokenClaims{
		UserID: userId,
		Name:   name,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(jwtExp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenCLaim)
	tokenStr, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenStr, nil
}

func ClaimTokenRefresh(tokenUser string, cfg *config.JWTConfig) (*dto.TokenClaims, error) {
	jwtSecret := []byte(cfg.Secret)
	claim := &dto.TokenClaims{}

	token, err := jwt.ParseWithClaims(tokenUser, claim, func(t *jwt.Token) (any, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")

	}

	return claim, nil
}
