package http

import (
	"gateway/internal/config"
	domain "gateway/internal/domain/gateway"
	"gateway/internal/infrastructure/proxy"
	"gateway/internal/service/jwt"
	healthhandler "gateway/internal/transport/http/handlers"
	"gateway/internal/transport/http/middleware"
	"gateway/internal/transport/http/router"
	gatewayuc "gateway/internal/usecase/gateway"
	"github.com/gin-gonic/gin"
)

func NewHandler(cfg *config.Config) (*gin.Engine, error) {
	gatewayService := gatewayuc.NewService(gatewayuc.BuilderInput{
		Auth:         domain.Upstream{Name: "auth", BaseURL: cfg.Auth.BaseURL()},
		Curriculum:   domain.Upstream{Name: "curriculum", BaseURL: cfg.Curriculum.BaseURL()},
		Lesson:       domain.Upstream{Name: "lesson", BaseURL: cfg.Lesson.BaseURL()},
		Enrollment:   domain.Upstream{Name: "enrollment", BaseURL: cfg.Enrollment.BaseURL()},
		Progress:     domain.Upstream{Name: "progress", BaseURL: cfg.Progress.BaseURL()},
		Notification: domain.Upstream{Name: "notification", BaseURL: cfg.Notification.BaseURL()},
	})

	proxyFactory := proxy.NewFactory()
	health := healthhandler.New(cfg.ServiceName, gatewayService.Services())
	jwtMgr := jwt.New(
		[]byte(cfg.JWT.Secret),
		cfg.JWT.Issuer,
		cfg.JWT.Audience,
		cfg.JWT.AccessTTL,
	)

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
		jwtMgr,
	)
}
