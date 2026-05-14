package controller

import (
	"errors"
	"mkpticket/internal/dto"
	"mkpticket/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ScheduleController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	List(c *gin.Context)
	Delete(c *gin.Context)
}

type scheduleControllerImpl struct {
	ScheduleService service.ScheduleService
}

func NewScheduleControllerImpl(scheduleService service.ScheduleService) ScheduleController {
	return &scheduleControllerImpl{
		ScheduleService: scheduleService,
	}
}

func (s *scheduleControllerImpl) Create(c *gin.Context) {
	var body dto.CreateScheduleReq

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse(err.Error()))
		return
	}

	if err := s.ScheduleService.CreateSchedule(c.Request.Context(), &body); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse("failed register user"))
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResponse(nil, "success create schedule"))
}

func (s *scheduleControllerImpl) Update(c *gin.Context) {
	var body dto.UpdateScheduleReq

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse(err.Error()))
		return
	}

	id := c.Param("id")
	body.ID = id

	if err := s.ScheduleService.UpdateSchedule(c.Request.Context(), &body); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse("failed update schedule"))
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResponse(nil, "success update schedule"))
}

func (s *scheduleControllerImpl) List(c *gin.Context) {
	schedule, err := s.ScheduleService.ListSchedule(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse("failed get schedule"))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse(schedule, "success get schedule"))
}

func (s *scheduleControllerImpl) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := s.ScheduleService.DeleteSchedule(c.Request.Context(), id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse("failed delete schedule"))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse(nil, "success delete schedule"))
}
