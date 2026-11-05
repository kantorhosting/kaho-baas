# Kaho BaaS Roadmap

This roadmap describes where Kaho BaaS is heading and in which order we plan to build it. It is a living document: priorities can change based on feedback from users and contributors.

Milestones have no fixed dates. Each milestone is tracked as a [GitHub Milestone](https://github.com/kantorhosting/kaho-baas/milestones) with linked issues.

## Vision

Kaho BaaS is an open-source Backend as a Service: authentication, databases, file storage, realtime, functions and site hosting behind one API. It should be **easy to deploy anywhere with Docker**, and **stateless**, so it scales by simply running more containers. Later, Kaho should also let developers deploy their own stateless apps and frontends as easily as on Netlify and similar platforms.

The same container image powers both self-hosted installations and Kantor Hosting's managed Kaho Cloud.

## Principles

Every feature must follow these rules, so a Kaho deployment stays simple and horizontally scalable.

**Stateless application.** Any replica can serve any request, and any container can be stopped at any time without losing data or logging users out. Concretely:

- **One image, configured only through environment variables.** No config files baked into the container ([twelve-factor](https://12factor.net/config)). Invalid or missing configuration fails fast at startup with a clear error.
- **No local disk.** The app never writes to its container filesystem except temporary files in `/tmp`. Files go to S3-compatible object storage.
- **Postgres is the only required dependency.** Sessions, background jobs, rate limits, schedules and realtime fan-out (`LISTEN/NOTIFY`) are stored in or coordinated through Postgres. Optional services (for example Redis) may be added later for very large installations, but are never required.
- **Migrations are a separate step.** Schema changes run through an explicit `migrate` command, not on every application start, so multiple replicas can start at the same time safely.
- **Operable by default.** Graceful shutdown, separate liveness and readiness endpoints, structured JSON logs and metrics.

**Easy deployment.** A working installation is `docker compose up`. Scaling is `docker compose up --scale app=3` behind a load balancer, or more replicas on Kubernetes.

## Milestones

### v0.1 Foundation

Make the current code base production-ready and stateless-friendly.

- [ ] Fix known issues: `/health` exposes database details, `GET /v1/account/` lists all users without authentication, failed logins return `400` instead of `401`, the server returns `nil` when the database is unreachable, and an invalid `PORT` is silently ignored
- [ ] `config` package: load and validate all environment variables at startup
- [ ] Use a single database connection pool
- [ ] Versioned SQL migrations and a `migrate` command, guarded by a Postgres advisory lock; remove `AutoMigrate` on startup
- [ ] `/healthz` (liveness) and `/readyz` (readiness, checks the database) endpoints
- [ ] Structured JSON logging with a request ID on every request
- [ ] Configurable CORS (allowed origins, `X-Kaho-Project` header)
- [ ] `projects` table: requests reference a real project instead of an arbitrary `X-Kaho-Project` value
- [ ] Server API keys per project
- [ ] Release workflow that publishes multi-architecture images (`amd64`, `arm64`) to GitHub Container Registry with semantic version tags
- [ ] Quick start: `docker compose up` with the published image, no local build needed

**Done when:** a new user can start Kaho from the published image in a few minutes, and two replicas can start at the same time against the same database.

### v0.2 Stateless sessions and core authentication

- [ ] `sessions` table: random, opaque session tokens stored hashed in Postgres, sent as a cookie or an `Authorization` header
- [ ] List, revoke and log out sessions (current session and all sessions)
- [ ] Rate limiting for authentication endpoints, shared between replicas through Postgres
- [ ] Background job queue in Postgres (`SELECT ... FOR UPDATE SKIP LOCKED`)
- [ ] Outgoing email through SMTP, sent by the job queue
- [ ] Email verification and password reset
- [ ] Account endpoints: get and update the current user, change password

**Done when:** an end-to-end test runs two replicas behind a reverse proxy, stops one of them, and the user stays logged in.

### v0.3 Complete authentication

- [ ] OAuth2 login providers (starting with Google and GitHub)
- [ ] Multi-factor authentication with TOTP
- [ ] Teams, roles and a permission model reused by all later features
- [ ] Scoped API keys (limit a key to specific services and actions)
- [ ] Start of the admin Console (web UI to manage projects and users), developed as a separate track

### v0.4 Database

- [ ] Collections and documents per project, stored as JSONB in Postgres
- [ ] Per-collection and per-document permissions
- [ ] Query API: filters, sorting, pagination
- [ ] Indexes managed through the API
- [ ] Client SDKs generated from the OpenAPI specification, starting with JavaScript

### v0.5 Storage

- [ ] Buckets with permissions, maximum file size and allowed file types
- [ ] S3-compatible storage backend (AWS S3, MinIO, Cloudflare R2 and others); no files on local disk
- [ ] Upload, download and delete files; pre-signed URLs for direct transfers
- [ ] MinIO added to the Docker Compose setup for self-hosting

### v0.6 Realtime

- [ ] WebSocket endpoint to subscribe to changes in databases, storage and accounts
- [ ] Fan-out between replicas with Postgres `LISTEN/NOTIFY`, so a client connected to any replica receives every event
- [ ] Permission checks per channel and per event

### v0.7 Functions

- [ ] Separate, optional **executor** service that runs user code in isolated containers; the core API stays stateless and never runs containers itself
- [ ] Function code stored in object storage
- [ ] Runtimes: Node.js, Python and Go to start
- [ ] Triggers: HTTP, events (for example "user created") and schedules (scheduler coordinated through Postgres, so only one replica runs each schedule)
- [ ] Environment variables and secrets per function

### v0.8 Sites

Deploy stateless apps and frontends on Kaho, similar to Netlify.

- [ ] Deploy static sites and single-page applications by uploading a build or by building from a Git repository (builds run in the executor)
- [ ] Build output stored in object storage and served by Kaho based on the request host
- [ ] Atomic deploys and instant rollback to any previous deployment
- [ ] Preview deployments
- [ ] Build-time environment variables
- [ ] Custom domains with automatic TLS through a reverse proxy (for example Caddy or Traefik); certificates are never stored on the app container's disk

### v1.0 Stable

- [ ] Stable, versioned API with a deprecation policy
- [ ] Stable SDKs
- [ ] Documentation site
- [ ] Backup, restore and upgrade guides
- [ ] Helm chart for Kubernetes
- [ ] Security review
- [ ] Kaho Cloud runs on the same open-source image

## Not planned for now

These ideas may be revisited later:

- Redis as an optional backend for very large installations
- GraphQL API
- Multi-region deployments

## Contributing to the roadmap

- Want to work on an item? Look for an existing issue or open one to discuss the approach first. See the [Contributing Guide](CONTRIBUTING.md).
- Have an idea or think priorities should change? Open a [feature request](https://github.com/kantorhosting/kaho-baas/issues/new/choose) or talk to us on [Discord](https://discord.gg/FsWEpr7YzD).
