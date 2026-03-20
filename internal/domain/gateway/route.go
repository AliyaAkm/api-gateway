package gateway

import "strings"

type Upstream struct {
	Name    string
	BaseURL string
}

type Route struct {
	Name           string
	GatewayPrefix  string
	UpstreamPrefix string
	Protected      bool
	Upstream       Upstream
}

type ServiceInfo struct {
	Name          string `json:"name"`
	GatewayPrefix string `json:"gateway_prefix"`
	UpstreamURL   string `json:"upstream_url"`
	Protected     bool   `json:"protected"`
}

func (r Route) Matches(path string) bool {
	prefix := normalizePath(r.GatewayPrefix)
	return path == prefix || strings.HasPrefix(path, prefix+"/")
}

func (r Route) RewritePath(path string) string {
	gatewayPrefix := normalizePath(r.GatewayPrefix)
	upstreamPrefix := normalizePath(r.UpstreamPrefix)

	if path == gatewayPrefix {
		return upstreamPrefix
	}

	suffix := strings.TrimPrefix(path, gatewayPrefix)
	if suffix == "" {
		return upstreamPrefix
	}

	return joinPath(upstreamPrefix, suffix)
}

func (r Route) Snapshot() ServiceInfo {
	return ServiceInfo{
		Name:          r.Name,
		GatewayPrefix: normalizePath(r.GatewayPrefix),
		UpstreamURL:   r.Upstream.BaseURL,
		Protected:     r.Protected,
	}
}

func normalizePath(path string) string {
	if path == "" || path == "/" {
		return "/"
	}

	trimmed := strings.TrimSuffix(path, "/")
	if trimmed == "" {
		return "/"
	}

	return trimmed
}

func joinPath(base, suffix string) string {
	switch {
	case base == "" || base == "/":
		if strings.HasPrefix(suffix, "/") {
			return suffix
		}
		return "/" + suffix
	case strings.HasSuffix(base, "/") && strings.HasPrefix(suffix, "/"):
		return base + strings.TrimPrefix(suffix, "/")
	case !strings.HasSuffix(base, "/") && !strings.HasPrefix(suffix, "/"):
		return base + "/" + suffix
	default:
		return base + suffix
	}
}
