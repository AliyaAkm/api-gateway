package gateway

import (
	domain "gateway/internal/domain/gateway"
	"gateway/internal/transport/http/router"
	"testing"
)

func TestPracticeReviewRoutes(t *testing.T) {
	service := NewService(BuilderInput{
		Curriculum: domain.Upstream{Name: "curriculum", BaseURL: "http://curriculum-service:8083"},
	})

	routes := service.Routes()

	practiceRoute := findRoute(t, routes, "practice")
	if !practiceRoute.Matches("/api/v1/practice/70c54ffc-fe87-470f-ae8e-2557001f5197/submissions") {
		t.Fatal("practice route should match manual submission create endpoint")
	}
	if len(practiceRoute.WriteRoleExemptSuffixes) != 2 {
		t.Fatalf("practice route exemptions: want 2, got %d", len(practiceRoute.WriteRoleExemptSuffixes))
	}

	studentRoute := findRoute(t, routes, "practice_submissions")
	if !studentRoute.Matches("/api/v1/practice-submissions/my") {
		t.Fatal("practice_submissions route should match student submission list")
	}
	if practiceRoute.Matches("/api/v1/practice-submissions/my") {
		t.Fatal("practice route must not match practice-submissions route")
	}

	teacherRoute := findRoute(t, routes, "teacher_practice_submissions")
	if !teacherRoute.Matches("/api/v1/teacher/practice-submissions/123/review") {
		t.Fatal("teacher_practice_submissions route should match teacher review endpoint")
	}
	if len(teacherRoute.AllowedRoles) != 2 ||
		teacherRoute.AllowedRoles[0] != router.RoleTeacher ||
		teacherRoute.AllowedRoles[1] != router.RoleAdmin {
		t.Fatalf("teacher route roles: want teacher/admin, got %#v", teacherRoute.AllowedRoles)
	}
}

func findRoute(t *testing.T, routes []domain.Route, name string) domain.Route {
	t.Helper()
	for _, route := range routes {
		if route.Name == name {
			return route
		}
	}
	t.Fatalf("route %q not found", name)
	return domain.Route{}
}
