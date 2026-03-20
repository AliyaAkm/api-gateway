package gateway

import "testing"

func TestRouteMatches(t *testing.T) {
	route := Route{GatewayPrefix: "/api/v1/courses"}

	cases := []struct {
		path string
		want bool
	}{
		{path: "/api/v1/courses", want: true},
		{path: "/api/v1/courses/42", want: true},
		{path: "/api/v1/course", want: false},
		{path: "/api/v1/courseships", want: false},
	}

	for _, tc := range cases {
		if got := route.Matches(tc.path); got != tc.want {
			t.Fatalf("path %q: want %v, got %v", tc.path, tc.want, got)
		}
	}
}

func TestRouteRewritePath(t *testing.T) {
	route := Route{
		GatewayPrefix:  "/api/v1/lessons",
		UpstreamPrefix: "/lessons",
	}

	cases := []struct {
		path string
		want string
	}{
		{path: "/api/v1/lessons", want: "/lessons"},
		{path: "/api/v1/lessons/15", want: "/lessons/15"},
		{path: "/api/v1/lessons/15/materials", want: "/lessons/15/materials"},
	}

	for _, tc := range cases {
		if got := route.RewritePath(tc.path); got != tc.want {
			t.Fatalf("path %q: want %q, got %q", tc.path, tc.want, got)
		}
	}
}
