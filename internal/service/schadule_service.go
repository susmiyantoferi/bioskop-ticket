package service

import (
	"context"
	"errors"
	"mkpticket/internal/dto"
	"mkpticket/internal/entity"
	"mkpticket/internal/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ScheduleService interface {
	CreateSchedule(c context.Context, req *dto.CreateScheduleReq) error
	UpdateSchedule(c context.Context, req *dto.UpdateScheduleReq) error
	ListSchedule(c context.Context) ([]*dto.ScheduleResp, error)
	DeleteSchedule(c context.Context, id string) error
}

type scheduleServiceImpl struct {
	DB           *gorm.DB
	ScheduleRepo repository.ScheduleRepository
	Log          *logrus.Logger
	Validate     *validator.Validate
}

func NewScheduleServiceImpl(db *gorm.DB, scheduleRepo repository.ScheduleRepository, log *logrus.Logger, validate *validator.Validate) ScheduleService {
	return &scheduleServiceImpl{
		DB:           db,
		ScheduleRepo: scheduleRepo,
		Log:          log,
		Validate:     validate,
	}
}

func timeFormat(showTime, endTime string) (*time.Time, *time.Time, error) {
	now := time.Now()
	date := now.Format("2006-01-02")

	showTimeStr := date + " " + showTime
	endTimeStr := date + " " + endTime

	layout := "2006-01-02 15:04"

	loc, _ := time.LoadLocation("Asia/Jakarta")

	show, err := time.ParseInLocation(layout, showTimeStr, loc)
	if err != nil {
		return nil, nil, err
	}

	end, err := time.ParseInLocation(layout, endTimeStr, loc)
	if err != nil {
		return nil, nil, err
	}

	return &show, &end, nil
}

func (s *scheduleServiceImpl) CreateSchedule(c context.Context, req *dto.CreateScheduleReq) error {
	db := s.DB.WithContext(c)

	if err := s.Validate.Struct(req); err != nil {
		s.Log.WithError(err).Error("Validation failed")
		return err
	}

	//cekk movie id
	movieRepo := &repository.RepositoryImpl[entity.Movie]{}
	movie, err := movieRepo.GetByID(db, req.MovieID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.Log.WithError(err).Error("get movie by id not found")
			return gorm.ErrRecordNotFound
		}
		s.Log.WithError(err).Error("failed get by id")
		return err

	}

	//cekk studio id
	studioRepo := &repository.RepositoryImpl[entity.Studio]{}
	studio, err := studioRepo.GetByID(db, req.StudioID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.Log.WithError(err).Error("get studio by id not found")
			return gorm.ErrRecordNotFound
		}
		s.Log.WithError(err).Error("failed get by id")
		return err

	}

	showTime, endTime, err := timeFormat(req.ShowTime, req.EndTime)
	if err != nil {
		return err
	}

	schedule := entity.Schedule{
		MovieID:  movie.ID,
		StudioID: studio.ID,
		ShowTime: *showTime,
		EndTime:  *endTime,
		Price:    req.Price,
	}

	if err := s.ScheduleRepo.Create(db, &schedule); err != nil {
		s.Log.WithError(err).Error("schedule: failed create")
		return err
	}

	return nil
}

func (s *scheduleServiceImpl) UpdateSchedule(c context.Context, req *dto.UpdateScheduleReq) error {
	db := s.DB.WithContext(c)

	if err := s.Validate.Struct(req); err != nil {
		s.Log.WithError(err).Error("Validation failed")
		return err
	}

	//cekk movie id exist
	movieRepo := &repository.RepositoryImpl[entity.Movie]{}
	movie, err := movieRepo.GetByID(db, req.MovieID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.Log.WithError(err).Error("get movie by id not found")
			return gorm.ErrRecordNotFound
		}
		s.Log.WithError(err).Error("failed get by id")
		return err

	}

	//cekk studio id exist
	studioRepo := &repository.RepositoryImpl[entity.Studio]{}
	studio, err := studioRepo.GetByID(db, req.StudioID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.Log.WithError(err).Error("get studio by id not found")
			return gorm.ErrRecordNotFound
		}
		s.Log.WithError(err).Error("failed get by id")
		return err

	}

	sch, err := s.ScheduleRepo.GetByID(db, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.Log.WithError(err).Error("get studio by id not found")
			return gorm.ErrRecordNotFound
		}
		s.Log.WithError(err).Error("failed get by id")
		return err

	}

	schedule, err := s.ScheduleRepo.GetByID(db, sch.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.Log.WithError(err).Error("get by id not found")
			return gorm.ErrRecordNotFound
		}
		s.Log.WithError(err).Error("failed get by id")
		return err
	}

	showTime, endTime, err := timeFormat(req.ShowTime, req.EndTime)
	if err != nil {
		return err
	}

	update := entity.Schedule{
		MovieID:  movie.ID,
		StudioID: studio.ID,
		ShowTime: *showTime,
		EndTime:  *endTime,
		Price:    req.Price,
	}

	if err := s.ScheduleRepo.Update(db, &update, schedule.ID); err != nil {
		s.Log.WithError(err).Error("failed to update schedule")
		return err
	}

	return nil
}

func (s *scheduleServiceImpl) ListSchedule(c context.Context) ([]*dto.ScheduleResp, error) {
	list, err := s.ScheduleRepo.ListSchedule(s.DB.WithContext(c))
	if err != nil {
		s.Log.WithError(err).Error("failed get all schedule")
		return nil, err
	}

	var responses []*dto.ScheduleResp
	for _, v := range list {
		response := dto.ScheduleResp{
			ID:            v.ID.String(),
			MovieID:       v.MovieID.String(),
			MovieName:     v.Movie.Name,
			MovieDuration: v.Movie.DurationMinutes,
			StudioID:      v.StudioID.String(),
			StudioName:    v.Studio.Name,
			ShowTime:      v.ShowTime,
			EndTime:       v.EndTime,
			Price:         v.Price,
		}

		responses = append(responses, &response)
	}

	return responses, nil
}

func (s *scheduleServiceImpl) DeleteSchedule(c context.Context, id string) error {
	db := s.DB.WithContext(c)

	schedule, err := s.ScheduleRepo.GetByID(db, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.Log.WithError(err).Error("get by id not found")
			return gorm.ErrRecordNotFound
		}
		s.Log.WithError(err).Error("failed get by id")
		return err
	}

	if err := s.ScheduleRepo.Delete(db, schedule.ID); err != nil {
		s.Log.WithError(err).Error("failed delete by id")
		return err
	}

	return nil
}
