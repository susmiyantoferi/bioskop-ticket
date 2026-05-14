package repository

import (
	"mkpticket/internal/entity"

	"gorm.io/gorm"
)

type ScheduleRepository interface {
	Repository[entity.Schedule]
	ListSchedule(db *gorm.DB) ([]*entity.Schedule, error)
}

type scheduleRepositoryImpl struct {
	RepositoryImpl[entity.Schedule]
}

func NewScheduleRepositoryImpl() ScheduleRepository {
	return &scheduleRepositoryImpl{}
}

func (s *scheduleRepositoryImpl) ListSchedule(db *gorm.DB) ([]*entity.Schedule, error) {
	var schedule []*entity.Schedule
	return schedule, db.Preload("Movie").Preload("Studio").Find(&schedule).Error
}
