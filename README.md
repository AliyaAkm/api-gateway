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
- `/api/v1/courses/*` -> course-service
- `/api/v1/lessons/*` -> lesson-service
- `/api/v1/enrollments/*` -> enrollment-service
- `/api/v1/progress/*` -> progress-service
- `/api/v1/notifications/*` -> notification-service
- `/health` and `/api/v1/health` -> gateway health check

## Auth-service integration

The gateway forwards:

- `/api/v1/auth/*` -> `http://localhost:8080/auth/*`

This matches your `auth-service`, because it listens on `HTTP_ADDR=:8080` and registers auth endpoints under `/auth`.

## Run

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

## Notes

- The gateway uses `gin` for routing and middleware.
- Protected routes require the `Authorization` header.
- The proxy adds:
  - `X-Request-ID`
  - `X-Forwarded-Host`
  - `X-Forwarded-Proto`
  - `X-Forwarded-Prefix`
  - `X-Gateway-Route`
