# API Gateway

API Gateway for the graduation project:

**Проектирование и разработка интерактивного образовательного портала для дополнительного обучения**  
**Design and development of an interactive learning platform for additional education**

The project follows clean architecture principles:

- `cmd/gateway` - application entry point
- `internal/config` - reading and validating `env`
- `internal/domain/gateway` - routing domain entities
- `internal/usecase/gateway` - public gateway route definitions
- `internal/infrastructure/proxy` - reverse proxy for upstream services
- `internal/transport/http` - `gin` router, middleware, and health endpoint

## Gateway routes

- `/api/v1/auth/*` -> auth-service
- `/api/v1/users/*` -> auth-service
- `/api/v1/course/*` -> curriculum-service
- `/api/v1/module/*` -> curriculum-service
- `/api/v1/lesson/*` -> curriculum-service
- `/api/v1/practice/*` -> curriculum-service
- `/api/v1/quiz/*` -> curriculum-service
- `/api/v1/review/*` -> curriculum-service
- `/api/v1/point/*` -> curriculum-service
- `/api/v1/leaderboard/*` -> curriculum-service
- `/api/v1/progress/*` -> curriculum-service
- `/api/v1/achievements/*` -> curriculum-service
- `/api/v1/streak/*` -> curriculum-service
- `/api/v1/dictionary/*` -> curriculum-service
- `/api/v1/order/*` -> payment-service
- `/api/v1/price/*` -> payment-service
- `/api/v1/payment_method/*` -> payment-service
- `/api/v1/payment/*` -> payment-service
- `/health` and `/api/v1/health` -> gateway health check

## Run locally

1. Create `.env` from `.env.example`
2. Run:

```bash
go run ./cmd/gateway
```

## Config style

The gateway keeps your preferred config loading style:

```go
_ = godotenv.Load(".env")

cfg, err := config.ReadEnv()
```

## Render deployment notes

- The gateway supports Render's `PORT` environment variable automatically.
- On Render, `PORT` has priority over `HTTP_ADDR`.
- `HTTP_ADDR` is still used for local development.
- Only configured upstream services are registered as routes.
- Upstream URLs must be reachable from Render.
- `localhost` upstream URLs will not work after deployment to Render.
- Use either Render internal service URLs or public service URLs for upstream services.
- Render internal `host:port` values are also supported and are normalized to `http://host:port`.

## Notes

- The gateway uses `gin` for routing and middleware.
- Protected routes require the `Authorization` header.
- The proxy adds:
  - `X-Request-ID`
  - `X-Forwarded-Host`
  - `X-Forwarded-Proto`
  - `X-Forwarded-Prefix`
  - `X-Gateway-Route`
