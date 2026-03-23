package gateway

import (
	domain "gateway/internal/domain/gateway"
)

type BuilderInput struct {
	Auth         domain.Upstream
	Curriculum   domain.Upstream
	Lesson       domain.Upstream
	Enrollment   domain.Upstream
	Progress     domain.Upstream
	Notification domain.Upstream
}

type Service struct {
	routes []domain.Route
}

func NewService(input BuilderInput) *Service {
	routes := make([]domain.Route, 0, 9)

	appendRoute := func(route domain.Route) {
		if route.Upstream.BaseURL == "" {
			return
		}
		routes = append(routes, route)
	}

	appendRoute(domain.Route{
		Name:           "auth",
		GatewayPrefix:  "/api/v1/auth",
		UpstreamPrefix: "/auth",
		Protected:      false,
		Upstream:       input.Auth,
	})
	appendRoute(domain.Route{
		Name:           "users",
		GatewayPrefix:  "/api/v1/users",
		UpstreamPrefix: "/users",
		Protected:      true,
		Upstream:       input.Auth,
	})
	appendRoute(domain.Route{
		Name:           "roles",
		GatewayPrefix:  "/api/v1/roles",
		UpstreamPrefix: "/roles",
		Protected:      true,
		Upstream:       input.Auth,
	})
	appendRoute(domain.Route{
		Name:           "user_roles",
		GatewayPrefix:  "/api/v1/user_roles",
		UpstreamPrefix: "/user_roles",
		Protected:      true,
		Upstream:       input.Auth,
	})
	appendRoute(domain.Route{
		Name:           "course",
		GatewayPrefix:  "/api/v1/course",
		UpstreamPrefix: "/course",
		Protected:      false,
		Upstream:       input.Curriculum,
	})
	appendRoute(domain.Route{
		Name:           "lessons",
		GatewayPrefix:  "/api/v1/lessons",
		UpstreamPrefix: "/lessons",
		Protected:      false,
		Upstream:       input.Lesson,
	})
	appendRoute(domain.Route{
		Name:           "enrollments",
		GatewayPrefix:  "/api/v1/enrollments",
		UpstreamPrefix: "/enrollments",
		Protected:      true,
		Upstream:       input.Enrollment,
	})
	appendRoute(domain.Route{
		Name:           "progress",
		GatewayPrefix:  "/api/v1/progress",
		UpstreamPrefix: "/progress",
		Protected:      true,
		Upstream:       input.Progress,
	})
	appendRoute(domain.Route{
		Name:           "notifications",
		GatewayPrefix:  "/api/v1/notifications",
		UpstreamPrefix: "/notifications",
		Protected:      true,
		Upstream:       input.Notification,
	})

	return &Service{
		routes: routes,
	}
}

func (s *Service) Routes() []domain.Route {
	routes := make([]domain.Route, len(s.routes))
	copy(routes, s.routes)
	return routes
}

func (s *Service) Services() []domain.ServiceInfo {
	items := make([]domain.ServiceInfo, 0, len(s.routes))
	for _, route := range s.routes {
		items = append(items, route.Snapshot())
	}
	return items
}
