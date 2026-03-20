package http

import (
	"gateway/internal/config"
	domain "gateway/internal/domain/gateway"
	"gateway/internal/infrastructure/proxy"
	healthhandler "gateway/internal/transport/http/handlers"
	"gateway/internal/transport/http/middleware"
	"gateway/internal/transport/http/router"
	gatewayuc "gateway/internal/usecase/gateway"

	"github.com/gin-gonic/gin"
)

func NewHandler(cfg *config.Config) (*gin.Engine, error) {
	gatewayService := gatewayuc.NewService(gatewayuc.BuilderInput{
		Auth:         domain.Upstream{Name: "auth", BaseURL: cfg.Auth.URL},
		Course:       domain.Upstream{Name: "course", BaseURL: cfg.Course.URL},
		Lesson:       domain.Upstream{Name: "lesson", BaseURL: cfg.Lesson.URL},
		Enrollment:   domain.Upstream{Name: "enrollment", BaseURL: cfg.Enrollment.URL},
		Progress:     domain.Upstream{Name: "progress", BaseURL: cfg.Progress.URL},
		Notification: domain.Upstream{Name: "notification", BaseURL: cfg.Notification.URL},
	})

	proxyFactory := proxy.NewFactory()
	health := healthhandler.New(cfg.ServiceName, gatewayService.Services())

	return router.New(
		gatewayService.Routes(),
		proxyFactory,
		health,
		[]gin.HandlerFunc{
			middleware.CORS(cfg.CORS.AllowedOrigins, cfg.CORS.AllowedMethods, cfg.CORS.AllowedHeaders),
			middleware.RequestID(),
			middleware.Logger(),
			middleware.Recover(),
		},
	)
}
