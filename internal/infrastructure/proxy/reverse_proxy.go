package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	domain "gateway/internal/domain/gateway"
	"gateway/internal/shared/httpx"
)

type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) Build(route domain.Route) (http.Handler, error) {
	target, err := url.Parse(route.Upstream.BaseURL)
	if err != nil {
		return nil, err
	}

	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			originalHost := req.Host
			originalScheme := "http"
			if req.TLS != nil {
				originalScheme = "https"
			}

			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.URL.Path = singleJoiningSlash(target.Path, route.RewritePath(req.URL.Path))
			req.Host = target.Host

			req.Header.Set("X-Forwarded-Host", originalHost)
			req.Header.Set("X-Forwarded-Proto", originalScheme)
			req.Header.Set("X-Forwarded-Prefix", route.GatewayPrefix)
			req.Header.Set("X-Gateway-Route", route.Name)
		},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			log.Printf("proxy error for route %s: %v", route.Name, err)
			httpx.WriteJSON(w, http.StatusBadGateway, map[string]any{
				"error":   "upstream service unavailable",
				"service": route.Name,
			})
		},
	}

	return proxy, nil
}

func singleJoiningSlash(a, b string) string {
	aslash := len(a) > 0 && a[len(a)-1] == '/'
	bslash := len(b) > 0 && b[0] == '/'

	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	default:
		return a + b
	}
}
