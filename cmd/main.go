package main

import (
	"mkpticket/infrastructure/config"
	"mkpticket/infrastructure/datastore"
	"mkpticket/infrastructure/logger"
	"mkpticket/internal/controller"
	"mkpticket/internal/repository"
	"mkpticket/internal/routes"
	"mkpticket/internal/service"

	"github.com/go-playground/validator/v10"
)

func main() {
	config, err := config.NewViper()
	if err != nil {
		panic(err)
	}

	log := logger.NewLogrus(&config.Logger)
	db := datastore.NewDatabase(&config.Postgres)
	validate := validator.New()

	//auth 
	authRepo := repository.NewAuthRepositoryImpl()
	authService := service.NewAuthServiceImpl(db, authRepo, log, validate, &config.Jwt)
	authController := controller.NewAuthControllerImpl(authService)

	//schedule
	scheduleRepo := repository.NewScheduleRepositoryImpl()
	scheduleService := service.NewScheduleServiceImpl(db, scheduleRepo, log, validate)
	scheduleController := controller.NewScheduleControllerImpl(scheduleService)


	router := routes.NewRoutes(authController, scheduleController, &config.Jwt)
	router.Run(config.App.Port)
}
