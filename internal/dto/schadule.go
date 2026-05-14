package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type CreateScheduleReq struct {
	MovieID  string          `json:"movie_id" validate:"required"`
	StudioID string          `json:"studio_id" validate:"required"`
	ShowTime string          `json:"show_time" validate:"required"`
	EndTime  string          `json:"end_time" validate:"required"`
	Price    decimal.Decimal `json:"price" validate:"required"`
}

type UpdateScheduleReq struct {
	ID       string          `json:"id" validate:"required"`
	MovieID  string          `json:"movie_id" validate:"required"`
	StudioID string          `json:"studio_id" validate:"required"`
	ShowTime string          `json:"show_time" validate:"required"`
	EndTime  string          `json:"end_time" validate:"required"`
	Price    decimal.Decimal `json:"price" validate:"required"`
}

type ScheduleResp struct {
	ID            string          `json:"id"`
	MovieID       string          `json:"movie_id" `
	MovieName     string          `json:"movie_name"`
	MovieDuration int             `json:"movie_duration"`
	StudioID      string          `json:"studio_id" `
	StudioName    string          `json:"studio_name"`
	ShowTime      time.Time       `json:"show_time" `
	EndTime       time.Time       `json:"end_time" `
	Price         decimal.Decimal `json:"price" `
}
