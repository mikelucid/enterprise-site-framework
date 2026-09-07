# enterprise-site-framework

Hybrid Golang platform that combines:
- **Enterprise multi-site management** (provisioning, deployment, operations, finance, HR/legal, lobbying)
- **NOVA creator economy** (real-person likeness, collaborative characters, royalty splits, licensing, consent)

## Framework Decision: Gin vs Fiber vs Echo

**Selected: Gin**

This project standardizes on **Gin** to align with the architecture requirements and enable straightforward integration with **gRPC-gateway**, middleware, and OpenAPI tooling.

## Target Architecture (Golang)

- **HTTP:** Gin
- **RPC:** gRPC + Protobuf
- **Data:** PostgreSQL + Redis
- **Async:** NATS
- **Infra:** Kubernetes + Helm + Terraform (AWS/GCP/Azure)
- **Security:** JWT, RBAC, mTLS
- **Observability:** OpenTelemetry, Prometheus, Jaeger

## Planned Project Structure (Clean Architecture)

```text
cmd/
  gateway/
  sitectl/
  creatorctl/
internal/
  shared/
    config logger errors auth db messaging validators contracts
  services/
    site-manager
    infra-provisioner
    ai-engine
    budget-manager
    ads-manager
    payment-collector
    accounting
    financial-reports
    tax-manager
    hr-team-manager
    legal-compliance
    lobbying-manager
    legislation-tracker
    creator-service
    real-person-model
    character-management
    consent-agreement
    royalty-engine
    anti-impersonation
    usage-analytics
api/
  proto/
  openapi/
deploy/
  helm/
  k8s/
infra/
  terraform/
migrations/
```

## Milestones and Estimated Completion Dates (1 FTE, 40 hrs/week)

Start date baseline: **2026-09-07**

| Milestone | Scope | Window | Estimated Completion |
|---|---|---|---|
| M1: Foundation & DevOps | Scaffolding, shared libs, migrations, Docker, CI/CD skeleton | Weeks 1–4 | **2026-10-04** |
| M2: Enterprise Core Services | Site manager, infra provisioner, AI engine, budget, ads services | Weeks 5–12 | **2026-11-29** |
| M3: Financial Services | Payments, accounting, reports, tax, invoices/receipts | Weeks 13–18 | **2027-01-10** |
| M4: HR & Legal | Hiring/payroll/onboarding, legal compliance, document automation | Weeks 19–24 | **2027-02-21** |
| M5: Lobbying & Gov Relations | Lobbying tracking, legislation APIs, compliance reporting | Weeks 25–28 | **2027-03-21** |
| M6: Creator Economy & Likeness | Real-person onboarding/KYC, consent, royalty split engine, dashboards | Weeks 29–40 | **2027-06-13** |
| M7: Integration & Gateway | Gin gateway, auth, throttling, service mesh, integration tests | Weeks 41–44 | **2027-07-11** |
| M8: Infra & Deployment | Helm/K8s/Terraform production deployment stack | Weeks 45–50 | **2027-08-22** |
| M9: Testing, Docs & Launch | >75% critical coverage, E2E, API docs, runbooks, launch readiness | Weeks 51–56 | **2027-10-03** |

## Success Targets

- 15+ services with gRPC + REST exposure
- Clean migrations on fresh PostgreSQL
- Kubernetes deployment healthy across core services
- E2E validation for payment, hiring, creator onboarding, royalty distribution
- Security enforcement: mTLS + JWT + RBAC
- Full CI/CD checks green before merge
