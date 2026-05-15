package router

import (
	"fmt"
	"gateway/internal/service/jwt"
	"net/http"

	domain "gateway/internal/domain/gateway"
	"gateway/internal/transport/http/middleware"

	healthhandler "gateway/internal/transport/http/handlers"

	"github.com/gin-gonic/gin"
)

const (
	RoleStudent = "student"
	RoleTeacher = "teacher"
	RoleManager = "manager"
	RoleAdmin   = "admin"
)

type ProxyFactory interface {
	Build(route domain.Route) (http.Handler, error)
}

func New(
	routes []domain.Route,
	proxyFactory ProxyFactory,
	health *healthhandler.HealthHandler,
	globalMiddlewares []gin.HandlerFunc,
	jwtMgr *jwt.Manager,
) (*gin.Engine, error) {
	engine := gin.New()
	engine.Use(globalMiddlewares...)

	engine.GET("/health", health.Handle)
	engine.GET("/api/v1/health", health.Handle)

	for _, route := range routes {
		proxyHandler, err := proxyFactory.Build(route)
		if err != nil {
			return nil, fmt.Errorf("build route %s: %w", route.Name, err)
		}

		routeHandlers := make([]gin.HandlerFunc, 0, 2)

		if route.Protected {
			routeHandlers = append(routeHandlers, middleware.Authenticate(jwtMgr))

			if len(route.AllowedRoles) > 0 {
				routeHandlers = append(routeHandlers, middleware.RequireRole(route.AllowedRoles...))
			}

			if len(route.WriteRoles) > 0 {
				routeHandlers = append(routeHandlers, middleware.RequireRoleForWriteMethodsExcept(route.WriteRoleExemptSuffixes, route.WriteRoles...))
			}
		}

		routeHandlers = append(routeHandlers, gin.WrapH(proxyHandler))

		engine.Any(route.GatewayPrefix, routeHandlers...)
		engine.Any(route.GatewayPrefix+"/*proxyPath", routeHandlers...)
	}

	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "route not found",
		})
	})

	return engine, nil
}
