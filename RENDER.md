# Render Setup

This service is intended to run as a public Render web service.
Only configured upstreams are exposed as routes.

## What this service can use

- `AUTH_SERVICE_URL`
- `CURRICULUM_SERVICE_URL`
- `LESSON_SERVICE_URL`
- `ENROLLMENT_SERVICE_URL`
- `PROGRESS_SERVICE_URL`
- `NOTIFICATION_SERVICE_URL`

Each upstream value can be provided in either format:

- full URL, for example `https://auth-service.onrender.com`
- internal Render host and port, for example `auth-service:10000`

The gateway normalizes bare `host:port` values to `http://host:port`.

For the first deployment you can configure only `CURRICULUM_SERVICE_URL`.
This repository now also includes a ready path for `AUTH_SERVICE_URL`.
The remaining upstreams can be added later without redeploying any private API contract.

## Deploy order

1. Deploy all private backend services first.
2. Make sure each private service is healthy inside Render.
3. Deploy the public gateway.
4. Render will inject the private services' `hostport` values into the gateway if you use `render.yaml`.

## Health check

- `/api/v1/health`
