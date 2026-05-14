package dto

import "github.com/golang-jwt/jwt/v5"

type TokenClaims struct {
	UserID string   `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type TokenResponse struct {
	Email         string `json:"email"`
	Token        string `json:"access_token"`
	TokenRefresh string `json:"refresh_token"`
	ExipresIn    int    `json:"expires_in"`
}
