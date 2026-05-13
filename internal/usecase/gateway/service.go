package gateway

import (
	domain "gateway/internal/domain/gateway"
	"gateway/internal/transport/http/router"
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
		AllowedRoles:   []string{router.RoleAdmin, router.RoleManager},
		Upstream:       input.Auth,
	})
	appendRoute(domain.Route{
		Name:           "roles",
		GatewayPrefix:  "/api/v1/roles",
		UpstreamPrefix: "/roles",
		Protected:      true,
		AllowedRoles:   []string{router.RoleAdmin},
		Upstream:       input.Auth,
	})
	appendRoute(domain.Route{
		Name:           "user_roles",
		GatewayPrefix:  "/api/v1/user_roles",
		UpstreamPrefix: "/user_roles",
		Protected:      true,
		AllowedRoles:   []string{router.RoleAdmin},
		Upstream:       input.Auth,
	})
	appendRoute(domain.Route{
		Name:           "course",
		GatewayPrefix:  "/api/v1/course",
		UpstreamPrefix: "/course",
		Protected:      true,
		Upstream:       input.Curriculum,
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
	appendRoute(domain.Route{
		Name:           "course_dictionary_status",
		GatewayPrefix:  "/api/v1/dictionary/status",
		UpstreamPrefix: "/dictionary/status",
		Protected:      true,
		Upstream:       input.Curriculum,
	})

	appendRoute(domain.Route{
		Name:           "course_dictionary_level",
		GatewayPrefix:  "/api/v1/dictionary/level",
		UpstreamPrefix: "/dictionary/level",
		Protected:      true,
		Upstream:       input.Curriculum,
	})

	appendRoute(domain.Route{
		Name:           "course_dictionary_duration_category",
		GatewayPrefix:  "/api/v1/dictionary/duration_category",
		UpstreamPrefix: "/dictionary/duration_category",
		Protected:      true,
		Upstream:       input.Curriculum,
	})

	appendRoute(domain.Route{
		Name:           "course_dictionary_topic",
		GatewayPrefix:  "/api/v1/dictionary/topic",
		UpstreamPrefix: "/dictionary/topic",
		Protected:      true,
		Upstream:       input.Curriculum,
	})

	appendRoute(domain.Route{
		Name:           "course_dictionary_tag",
		GatewayPrefix:  "/api/v1/dictionary/tag",
		UpstreamPrefix: "/dictionary/tag",
		Protected:      true,
		Upstream:       input.Curriculum,
	})

	appendRoute(domain.Route{
		Name:           "course_dictionary_locale",
		GatewayPrefix:  "/api/v1/dictionary/locale",
		UpstreamPrefix: "/dictionary/locale",
		Protected:      true,
		Upstream:       input.Curriculum,
	})
	appendRoute(domain.Route{
		Name:           "module",
		GatewayPrefix:  "/api/v1/module",
		UpstreamPrefix: "/module",
		Protected:      true,
		Upstream:       input.Curriculum,
	})
	appendRoute(domain.Route{
		Name:           "lesson",
		GatewayPrefix:  "/api/v1/lesson",
		UpstreamPrefix: "/lesson",
		Protected:      true,
		Upstream:       input.Curriculum,
	})
	appendRoute(domain.Route{
		Name:           "streak",
		GatewayPrefix:  "/api/v1/streak",
		UpstreamPrefix: "streak",
		Protected:      true,
		Upstream:       input.Curriculum,
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
