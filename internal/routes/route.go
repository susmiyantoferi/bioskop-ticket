package routes

import (
	"mkpticket/infrastructure/config"
	"mkpticket/internal/controller"
	"mkpticket/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRoutes(authentication controller.AuthController, schedule controller.ScheduleController, cfg *config.JWTConfig) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authentication.Register)
			auth.POST("/login", authentication.Login)

		}

		s := v1.Group("/schedules")
		s.Use(middleware.Authentication(cfg))
		{
			s.GET("", middleware.RoleAccessMiddleware("ADMIN", "CUSTOMER"), schedule.List)
		}

		admin := s.Group("")
		admin.Use(middleware.RoleAccessMiddleware("ADMIN"))
		{
			admin.POST("", schedule.Create)
			admin.PATCH("/:id", schedule.Update)
			admin.DELETE("/:id", schedule.Delete)
		}
	}

	return r

}
