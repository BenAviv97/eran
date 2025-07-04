
> **Mission** — A scalable, secure, creator-friendly live-streaming platform that supports up to 1 M concurrent viewers across thousands of live channels with sub-second latency.

---

## 1  Architecture at a Glance

| Layer                  | Key Components                                                    |
| ---------------------- | ----------------------------------------------------------------- |
| **Client**             | Web (Next.js + HLS.js/WebRTC) & React-Native mobile apps          |
| **Creator Studio**     | WebRTC device manager, scene composer, real-time health widgets   |
| **Realtime**           | WebSocket chat, presence, tipping (NATS JetStream + Redis)        |
| **Ingest → Transcode** | RTMPS/SRT edge POPs → FFmpeg/NVIDIA ladder → S3-compatible origin |
| **Delivery**           | Multi-CDN (Fastly + CloudFront) with TLS 1.3 & QUIC               |
| **Control Plane**      | Go microservices (gRPC) on Kubernetes, Istio mTLS                 |
| **Data Stores**        | PostgreSQL 15, ScyllaDB 5, Redis Cluster, ClickHouse 24           |
| **Observability**      | Prometheus + Loki, OpenTelemetry tracing, Grafana dashboards      |
| **Security**           | Vault secrets, OPA Gatekeeper, Falco IDS, CrowdSec WAF            |
| **FinOps**             | Kubecost cost monitoring & budget alerts                          |

---

## 2  Tech Stack

| Tier     | Technology                                     |
| -------- | ---------------------------------------------- |
| Frontend | React (Next.js), TypeScript, Tailwind, Zustand |
| Mobile   | React-Native 0.74, AVPlayer/ExoPlayer bridges  |
| Backend  | Go 1.22, gRPC, Echo, Buf build                 |
| Realtime | Redis 7 cluster, NATS JetStream                |
| Media    | FFmpeg, NVIDIA A10G/A16, SVT-AV1 (roadmap)     |
| Infra    | Docker, Kubernetes 1.30, Helm, Terraform       |
| Edge     | Cloudflare Workers / Lambda\@Edge              |
| CI/CD    | GitHub Actions, Helmfile, KEDA autoscaling     |

---

## 3  Repository Layout

```
/home/$USER/livestreampro
├── README.md
├── frontend/              # Next.js web client
├── mobile/                # React-Native app
├── backend/               # Go microservices
│   ├── cmd/
│   ├── internal/
│   └── proto/
├── infra/
│   ├── docker/compose.local.yaml
│   ├── k8s/overlays/
│   └── terraform/
├── docs/
│   ├── architecture.png
│   ├── moderation-policy.md
│   ├── dmca-policy.md
│   └── sre/runbooks/
│       ├── general.md
│       └── stream-outage.md
└── .github/workflows/     # CI/CD pipelines
```

---

## 4  Quick Start (Local Dev)

### Prerequisites

| Tool           | Version |
| -------------- | ------- |
| Docker Desktop | ≥ 4.30  |
| Node.js        | 20 LTS  |
| Go             | 1.22    |
| Make           | 4.4     |

### One-Command Bootstrap

```bash
git clone https://github.com/your-org/livestreampro.git ~/livestreampro
cd ~/livestreampro
make dev                # builds images & starts full stack
# Web UI:      http://localhost:3000
# API Gateway: http://localhost:8080
# Grafana:     http://localhost:3001 (admin/admin)
```
```bash
cd mobile
npm install
npm run android   # or: npm run ios
```

---

## 5  Configuration

Create the following files (never commit secrets):

```ini
# backend/.env
POSTGRES_URL=postgres://stream:stream@db:5432/streams?sslmode=disable
REDIS_URL=redis://redis:6379/0
NATS_URL=nats://nats:4222
JWT_SECRET=change-me
VAULT_ADDR=http://vault:8200

# frontend/.env.local
NEXT_PUBLIC_API_BASE=http://localhost:8080
NEXT_PUBLIC_HLS_URL=http://localhost:8080/hls
```

---

## 6  Testing & Quality Gates

| Stage         | Tools                      | CI Trigger        |
| ------------- | -------------------------- | ----------------- |
| Unit          | `go test`, Jest            | push / PR         |
| Integration   | Testcontainers, Playwright | PR matrix         |
| Security      | Trivy (fail on HIGH+), Snyk | PR blocker        |
| Accessibility | Lighthouse CI              | FE PRs            |
| Chaos         | LitmusChaos                | nightly (staging) |

Trivy scans all built Docker images and the job fails if any vulnerabilities with severity **HIGH** or greater are detected.

Run locally:

```bash
make test-backend      # Go unit tests
make test-frontend     # Jest + React Testing Library
make chaos-smoke       # quick chaos suite against dev stack
```

---

## 7  Deployment (Staging → Prod)

```bash
# 1. Build & push images
make docker-push REGISTRY=ghcr.io/your-org

# 2. Provision EKS / core infra
cd infra/terraform/prod && terraform apply

# 3. Deploy services
cd infra/k8s/overlays/prod && helmfile sync

# 4. Smoke & chaos tests
make chaos-smoke
```

Autoscaling is handled by **KEDA** (CPU, GPU, NATS lag, WebRTC sessions).

---

## 8  Security Hardening Checklist

* TLS 1.3, HSTS, OCSP stapling
* Istio mTLS & SPIFFE identities
* Vault-managed dynamic DB creds
* OPA Gatekeeper: deny `:latest`, enforce non-root UID, read-only FS
* Trivy scan gates (CVSS ≥ 7 → fail)
* Falco runtime alerts → PagerDuty
* CrowdSec WAF + AWS Shield Advanced
* Immutable logs → Loki → SIEM
* GDPR/CCPA DSAR API, COPPA kid-mode switch

---

## 9  Observability & FinOps

* **Prometheus** — service metrics, GPU & encoder stats
* **OpenTelemetry** — distributed traces across gRPC and WebRTC
* **Loki** — structured logs; Grafana dashboards & alerts
* **Kubecost** — live GPU / egress spend, budget alerts at 80 % threshold

---

## 10  Scaling Notes

| Component   | Strategy                                    |
| ----------- | ------------------------------------------- |
| Ingest POPs | Horizontal scale on RTMP session count      |
| Transcode   | GPU node pool + HPA at 70 % NVENC           |
| Chat        | Redis slot hashing, NATS sharding           |
| CDN         | Anycast multi-CDN with RUM-based steering   |
| DB          | Patroni HA Postgres, Scylla for high-TPS KV |

---

## 11  Contributing

1. **Fork** → `git checkout -b feat/my-feature`
2. Use Conventional Commits (`feat:`, `fix:`…).
3. Run `make pre-commit` before pushing.
4. Open PR → CI must pass, at least one maintainer review.

Full guidelines in [docs/contributing.md](docs/contributing.md).

---

## 12  License

Apache License 2.0 — see `/LICENSE`.

---

## 13  Useful Make Targets

| Command             | Description                            |
| ------------------- | -------------------------------------- |
| `make dev`          | Full local stack (Docker Compose)      |
| `make proto`        | Compile gRPC stubs (Go + TS)           |
| `make lint`         | golangci-lint, ESLint, Prettier        |
| `make docker-push`  | Build & push images                    |
| `make chaos-smoke`  | Fast chaos tests                       |
| `make budget-check` | Kubecost forecast vs budget            |
| `make lighthouse`   | Lighthouse accessibility & perf checks |

