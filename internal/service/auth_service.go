package service

import (
	"context"
	"errors"
	"fmt"
	"mkpticket/infrastructure/config"
	"mkpticket/internal/dto"
	"mkpticket/internal/entity"
	"mkpticket/internal/repository"
	"mkpticket/utils/hashing"
	"mkpticket/utils/token"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(c context.Context, req *dto.RegisterReq) error
	Login(c context.Context, req *dto.LoginReq) (*dto.TokenResponse, error)
}

type authServiceImpl struct {
	DB       *gorm.DB
	AuthRepo repository.AuthRepository
	Log      *logrus.Logger
	Validate *validator.Validate
	JWT      *config.JWTConfig
}

func NewAuthServiceImpl(db *gorm.DB, authRepo repository.AuthRepository, log *logrus.Logger, validate *validator.Validate, jwt *config.JWTConfig) AuthService {
	return &authServiceImpl{
		DB:       db,
		AuthRepo: authRepo,
		Log:      log,
		Validate: validate,
		JWT:      jwt,
	}
}

func (a *authServiceImpl) Register(c context.Context, req *dto.RegisterReq) error {
	db := a.DB.WithContext(c)

	if err := a.Validate.Struct(req); err != nil {
		a.Log.WithError(err).Error("Validation failed")
		return err
	}

	pass, err := hashing.HashPassword(req.Password)
	if err != nil {
		a.Log.WithError(err).Error("failed hash password")
		return fmt.Errorf("hashing: %w", err)
	}

	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: pass,
		Role:     entity.Customer,
	}

	if err := a.AuthRepo.Create(db, &user); err != nil {
		a.Log.WithError(err).Error("auth service: failed register user")
		return err
	}

	return nil
}

func (a *authServiceImpl) Login(c context.Context, req *dto.LoginReq) (*dto.TokenResponse, error) {
	db := a.DB.WithContext(c)

	if err := a.Validate.Struct(req); err != nil {
		a.Log.WithError(err).Error("Validation failed")
		return nil, err
	}

	user, err := a.AuthRepo.FindByEmail(db, req.Email)
	if err != nil {
		a.Log.WithError(err).Error("failed get user by emai")
		return nil, err
	}

	if !hashing.CompareHashPassword(user.Password, req.Password) {
		a.Log.WithError(err).Error("failed compare password")
		return nil, errors.New("CRED_ERR")
	}

	tokenExp := a.JWT.Expire

	accessToken, err := token.GenerateToken(user.ID.String(), user.Name, user.Email, string(user.Role), time.Duration(tokenExp), a.JWT)
	if err != nil {
		a.Log.WithError(err).Error("failed generate access token")
		return nil, err
	}

	refreshToken, err := token.GenerateToken(user.ID.String(), user.Name, user.Email, string(user.Role), time.Duration(tokenExp*7), a.JWT)
	if err != nil {
		a.Log.WithError(err).Error("failed generate refresh token")
		return nil, err
	}

	return &dto.TokenResponse{
		Email:        user.Email,
		Token:        accessToken,
		TokenRefresh: refreshToken,
		ExipresIn:    tokenExp * 3600,
	}, nil

}
