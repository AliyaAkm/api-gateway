package gateway

import (
	domain "gateway/internal/domain/gateway"
)

type BuilderInput struct {
	Auth         domain.Upstream
	Course       domain.Upstream
	Lesson       domain.Upstream
	Enrollment   domain.Upstream
	Progress     domain.Upstream
	Notification domain.Upstream
}

type Service struct {
	routes []domain.Route
}

func NewService(input BuilderInput) *Service {
	return &Service{
		routes: []domain.Route{
			{
				Name:           "auth",
				GatewayPrefix:  "/api/v1/auth",
				UpstreamPrefix: "/auth",
				Protected:      false,
				Upstream:       input.Auth,
			},
			{
				Name:           "users",
				GatewayPrefix:  "/api/v1/users",
				UpstreamPrefix: "/users",
				Protected:      true,
				Upstream:       input.Auth,
			},
			{
				Name:           "roles",
				GatewayPrefix:  "/api/v1/roles",
				UpstreamPrefix: "/roles",
				Protected:      true,
				Upstream:       input.Auth,
			},
			{
				Name:           "user_roles",
				GatewayPrefix:  "/api/v1/user_roles",
				UpstreamPrefix: "/user_roles",
				Protected:      true,
				Upstream:       input.Auth,
			},
			{
				Name:           "courses",
				GatewayPrefix:  "/api/v1/courses",
				UpstreamPrefix: "/courses",
				Protected:      false,
				Upstream:       input.Course,
			},
			{
				Name:           "lessons",
				GatewayPrefix:  "/api/v1/lessons",
				UpstreamPrefix: "/lessons",
				Protected:      false,
				Upstream:       input.Lesson,
			},
			{
				Name:           "enrollments",
				GatewayPrefix:  "/api/v1/enrollments",
				UpstreamPrefix: "/enrollments",
				Protected:      true,
				Upstream:       input.Enrollment,
			},
			{
				Name:           "progress",
				GatewayPrefix:  "/api/v1/progress",
				UpstreamPrefix: "/progress",
				Protected:      true,
				Upstream:       input.Progress,
			},
			{
				Name:           "notifications",
				GatewayPrefix:  "/api/v1/notifications",
				UpstreamPrefix: "/notifications",
				Protected:      true,
				Upstream:       input.Notification,
			},
		},
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
