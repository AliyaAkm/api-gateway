package router

import (
	"net/http"
	"testing"

	domain "gateway/internal/domain/gateway"
	healthhandler "gateway/internal/transport/http/handlers"

	"github.com/gin-gonic/gin"
)

type stubProxyFactory struct{}

func (stubProxyFactory) Build(route domain.Route) (http.Handler, error) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}), nil
}

func TestNewRegistersRoutesWithoutConflicts(t *testing.T) {
	gin.SetMode(gin.TestMode)

	routes := []domain.Route{
		{
			Name:           "auth",
			GatewayPrefix:  "/api/v1/auth",
			UpstreamPrefix: "/auth",
			Upstream:       domain.Upstream{Name: "auth", BaseURL: "http://localhost:8080"},
		},
		{
			Name:           "courses",
			GatewayPrefix:  "/api/v1/courses",
			UpstreamPrefix: "/courses",
			Upstream:       domain.Upstream{Name: "course", BaseURL: "http://localhost:8083"},
		},
	}

	health := healthhandler.New("gateway", nil)

	engine, err := New(routes, stubProxyFactory{}, health, []gin.HandlerFunc{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if engine == nil {
		t.Fatal("expected non-nil gin engine")
	}
}
