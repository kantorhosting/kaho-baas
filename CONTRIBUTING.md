# Contributing to Kaho BaaS

Thanks for your interest in contributing! This guide explains how to set up the project, the conventions we follow, and how to get your change merged.

By participating in this project you agree to follow our [Code of Conduct](CODE_OF_CONDUCT.md).

## Before you start

- **Security issues:** do not open a public issue. Follow the [Security Policy](SECURITY.md) instead.
- **What to work on:** the [Roadmap](ROADMAP.md) lists our priorities and planned milestones.
- **Bugs and features:** search [existing issues](https://github.com/kantorhosting/kaho-baas/issues) first. If none exists, open one using the issue templates.
- **Larger changes:** open an issue to discuss the approach before writing code, so your work is not wasted if the direction does not fit the project.

## Contributor License Agreement (CLA)

Before we can merge your first Pull Request, you must sign the [Contributor License Agreement](CLA.md). A bot will comment on your Pull Request with instructions; signing only takes a single comment and is needed once.

Kaho BaaS is licensed under the [AGPL-3.0](LICENSE), and Kantor Hosting also offers it as a hosted service and under commercial licenses. The CLA gives Kantor Hosting the rights needed to do so, while you keep the copyright of your contribution.

## Development setup

### Prerequisites

- [Go](https://go.dev/dl/): the version used in the `Dockerfile` (`FROM golang:<version>`) is recommended, `go.mod` declares the minimum
- [Docker](https://docs.docker.com/get-docker/) with Docker Compose

### Run the project

1. Fork the repository and clone your fork:
   ```bash
   git clone https://github.com/<your-username>/kaho-baas.git
   cd kaho-baas
   ```
2. Create your environment file:
   ```bash
   cp .env.example .env
   ```
3. Start the app with hot reload and a Postgres database:
   ```bash
   make docker-dev
   ```
   Or run Postgres yourself and start the app locally with `make run` / `make watch`.

### Run the tests

```bash
make test
```

Some tests start a Postgres container with [testcontainers](https://golang.testcontainers.org/), so Docker must be running.

Before opening a Pull Request, also make sure these pass:

```bash
gofmt -l .        # must print nothing
go vet ./...
go test -race ./...
```

The same checks run automatically in CI on every Pull Request, together with a Docker build and a check that the Swagger docs are up to date. A Pull Request can only be merged when CI passes.

### API documentation

If you add or change an endpoint, update its Swagger annotations and regenerate the docs:

```bash
make swagger
```

This runs [swag](https://github.com/swaggo/swag) at the version pinned in `go.mod`, so you do not need to install it yourself.

Commit the generated files in `docs/` together with your change. CI fails if the committed docs differ from the generated ones.

## Project structure

```
cmd/api/            Application entry point
internal/apps/      Feature modules (e.g. account): handlers, services, repositories, models, routes
internal/database/  Database connection and migrations
internal/pkg/       Shared packages (utils, constants, session manager, ...)
internal/server/    HTTP server setup and route registration
docs/               Generated Swagger documentation
```

Each module follows the same layering: `handlers` → `services` → `repositories`. Business logic belongs in services, database access in repositories.

## Branches and commits

Create a branch from `main` named after the type of change:

| Prefix      | Use for                               |
|-------------|---------------------------------------|
| `feat/`     | New features                          |
| `fix/`      | Bug fixes                             |
| `docs/`     | Documentation only                    |
| `refactor/` | Code changes that do not change behavior |
| `test/`     | Adding or fixing tests                |
| `chore/`    | Tooling, build, dependencies, CI      |

Commit messages follow [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>(<optional scope>): <short summary>

feat(account): add logout endpoint
fix(account): query user by email
docs: add contributing guide
```

Keep commits focused. Make sure your Git `user.name` and `user.email` are set correctly, as they are recorded in the project history.

## Pull Requests

1. Keep each Pull Request focused on a single change. Smaller Pull Requests are reviewed faster.
2. Fill in the Pull Request template: the related issue, what you did, and proof that it works (test output or screenshots).
3. Add or update tests for your change.
4. Rebase on the latest `main` and resolve conflicts before requesting a review.
5. At least one maintainer approval is required before merging.

Maintainers may ask for changes. Please respond to review comments or let us know if you are no longer able to continue, so someone else can pick up the work.

## Questions

Join our [Discord](https://discord.gg/FsWEpr7YzD) or open a discussion in an issue.
