package gateway

import (
	domain "gateway/internal/domain/gateway"
	"gateway/internal/transport/http/router"
)

type BuilderInput struct {
	Auth       domain.Upstream
	Curriculum domain.Upstream
	Payment    domain.Upstream
}

type Service struct {
	routes []domain.Route
}

func NewService(input BuilderInput) *Service {
	routes := make([]domain.Route, 0, 24)

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
		Name:           "profiles",
		GatewayPrefix:  "/api/v1/profiles",
		UpstreamPrefix: "/profiles",
		Protected:      true,
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
		Name:           "progress",
		GatewayPrefix:  "/api/v1/progress",
		UpstreamPrefix: "/progress",
		Protected:      true,
		Upstream:       input.Curriculum,
	})
	appendRoute(domain.Route{
		Name:           "achievements",
		GatewayPrefix:  "/api/v1/achievements",
		UpstreamPrefix: "/achievements",
		Protected:      true,
		Upstream:       input.Curriculum,
	})
	appendRoute(domain.Route{
		Name:           "order",
		GatewayPrefix:  "/api/v1/order",
		UpstreamPrefix: "/order",
		Protected:      true,
		Upstream:       input.Payment,
	})
	appendRoute(domain.Route{
		Name:           "price",
		GatewayPrefix:  "/api/v1/price",
		UpstreamPrefix: "/price",
		Protected:      true,
		Upstream:       input.Payment,
	})
	appendRoute(domain.Route{
		Name:           "payment_method",
		GatewayPrefix:  "/api/v1/payment_method",
		UpstreamPrefix: "/payment_method",
		Protected:      true,
		Upstream:       input.Payment,
	})
	appendRoute(domain.Route{
		Name:           "payment",
		GatewayPrefix:  "/api/v1/payment",
		UpstreamPrefix: "/payment",
		Protected:      true,
		Upstream:       input.Payment,
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
		Name:                    "lesson",
		GatewayPrefix:           "/api/v1/lesson",
		UpstreamPrefix:          "/lesson",
		Protected:               true,
		WriteRoles:              []string{router.RoleTeacher, router.RoleAdmin},
		WriteRoleExemptSuffixes: []string{"/complete"},
		Upstream:                input.Curriculum,
	})
	appendRoute(domain.Route{
		Name:           "practice",
		GatewayPrefix:  "/api/v1/practice",
		UpstreamPrefix: "/practice",
		Protected:      true,
		Upstream:       input.Curriculum,
	})
	appendRoute(domain.Route{
		Name:           "quiz",
		GatewayPrefix:  "/api/v1/quiz",
		UpstreamPrefix: "/quiz",
		Protected:      true,
		Upstream:       input.Curriculum,
	})
	appendRoute(domain.Route{
		Name:           "review",
		GatewayPrefix:  "/api/v1/review",
		UpstreamPrefix: "/review",
		Protected:      true,
		Upstream:       input.Curriculum,
	})
	appendRoute(domain.Route{
		Name:           "point",
		GatewayPrefix:  "/api/v1/point",
		UpstreamPrefix: "/point",
		Protected:      true,
		Upstream:       input.Curriculum,
	})
	appendRoute(domain.Route{
		Name:           "leaderboard",
		GatewayPrefix:  "/api/v1/leaderboard",
		UpstreamPrefix: "/leaderboard",
		Protected:      true,
		Upstream:       input.Curriculum,
	})
	appendRoute(domain.Route{
		Name:           "streak",
		GatewayPrefix:  "/api/v1/streak",
		UpstreamPrefix: "/streak",
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
