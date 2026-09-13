# Enterprise Site Framework

## 1. Project overview
NOVA Enterprise Multi-Site Management Framework using Gin, gRPC, GORM, and GWAV/OpenAI-ready AI integration.
This repository also includes an agent runtime scaffold (`agent:run`) so automation agents can live inside the framework and manage build/ops workflows.

## 2. Architecture diagram
- API Gateway (`cmd/api-gateway`) fronts service modules.
- Domain modules follow clean architecture under `internal/*`.
- Shared cross-cutting libraries live in `pkg/*`.
- Bootstrap container (`bootstrap/app.go`) provides DI and provider registration.

## 3. Technology stack
- Go 1.25
- Gin HTTP framework
- gRPC/protobuf APIs
- GORM + PostgreSQL
- NATS messaging
- Zap logging
- OpenTelemetry tracing

## 4. Quick start
```bash
cp .env.example .env
make setup
make test
make build
make docker-up
```

## 5. Configuration
YAML configuration files are in `config/`:
- `app.yaml`
- `database.yaml`
- `ai.yaml`
- `queue.yaml`
- `auth.yaml`

Environment variables override YAML values.

## 6. Project structure
See the repository tree for a clean-architecture Go layout with `cmd/*` service entrypoints, domain-driven `internal/*` modules, shared `pkg/*` libraries, and `app`/`bootstrap` orchestration layers.

## 7. Extending
- Add services in `internal/<service>` with domain/usecase/adapter/port layers.
- Register providers in `bootstrap/providers.go`.
- Add AI model adapters in `internal/ai/adapter` and expose via facades.
- Extend the internal automation agent in `internal/agent` and run it with `go run ./cmd/agent-runner agent:run`.

## 8. Testing
```bash
make test
```
Shared libraries include focused unit tests in their package directories.

## 9. Deployment
- Docker: `Dockerfile`, `docker-compose.yml`
- Kubernetes manifests: `deploy/kubernetes`
- Helm chart seed: `deploy/helm/enterprise-framework`
- Terraform module seeds: `deploy/terraform/{aws,gcp,azure}`

## 10. Contributing
1. Create a branch
2. Make focused changes
3. Run tests/lint/build locally
4. Open PR with clear scope and verification details
