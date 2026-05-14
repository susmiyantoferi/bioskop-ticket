package controller

import (
	"mkpticket/internal/dto"
	"mkpticket/internal/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type authControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthControllerImpl(authService service.AuthService) AuthController {
	return &authControllerImpl{
		AuthService: authService,
	}
}

func (a *authControllerImpl) Register(c *gin.Context) {
	var body dto.RegisterReq

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse(err.Error()))
		return
	}

	if err := a.AuthService.Register(c.Request.Context(), &body); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse("failed register user"))
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResponse(nil, "success register user"))
}

func (a *authControllerImpl) Login(c *gin.Context) {
	var body dto.LoginReq

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse(err.Error()))
		return
	}

	response, err := a.AuthService.Login(c.Request.Context(), &body)
	if err != nil {
		if strings.Contains(err.Error(), "CRED_ERR") {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse("email or password wrong"))
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse("failed login"))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse(response, "success login"))
}
