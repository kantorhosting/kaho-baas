# Kaho BaaS - Open-Source Backend as a Service

Kaho BaaS is an open-source **Backend as a Service (BaaS)** built with Go. Our goal is to provide authentication, databases, file storage, realtime, functions and site hosting behind one API, so you can focus on building your application while Kaho handles the backend. Kaho is designed to be **easy to deploy with Docker** and **stateless**, so it scales by simply running more containers.

> **Early development:** Kaho BaaS is not ready for production use yet. APIs may change without notice.

## ✨ Status

**Available now:**

- ✅ **Email and password accounts** – register and log in users per project
- ✅ **Docker setup** – development with hot reload, and a minimal production image
- ✅ **API documentation** – Swagger UI at `/docs`

**Planned:** stateless sessions, complete authentication (OAuth, MFA, teams), databases, file storage, realtime, functions and Netlify-like site hosting. See the [Roadmap](ROADMAP.md) for details and priorities.

---

## 🚀 Getting Started

### Prerequisites
- Go (latest version recommended)
- Docker (optional, for containerized deployment)
- PostgreSQL (or any supported database backend)

### Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/kantorhosting/kaho-baas.git
   cd kaho-baas
   ```

2. **Install Dependencies**
   ```bash
   go mod tidy
   ```

3. **Run the Server**
   ```bash
   go run cmd/api/main.go
   ```

4. **Access the API**
   The server runs on `http://localhost:8080` by default. You can access API endpoints using cURL, Postman, or a frontend client.

---

## 📖 Documentation

Full documentation is available at **[Kaho BaaS Docs](#)** (coming soon).

---

## 🤝 Contributing

We welcome contributions! Please read the [Contributing Guide](CONTRIBUTING.md) before opening a Pull Request. All contributors must sign the [Contributor License Agreement](CLA.md), and everyone taking part in the project is expected to follow our [Code of Conduct](CODE_OF_CONDUCT.md).

Found a security issue? **Do not open a public issue.** See our [Security Policy](SECURITY.md).

---

## 🛠 Configuration

Copy `.env.example` to `.env` and adjust the values:
```env
PORT=8080
APP_ENV=local
DB_HOST=localhost
DB_PORT=5432
DB_DATABASE=kaho
DB_USERNAME=kaho
DB_PASSWORD=change-me
DB_SCHEMA=public
DB_EXPOSE_PORT=5432
```

When running with Docker Compose, `DB_HOST` and `DB_PORT` are set automatically to point at the `db` service.

---

## 🐳 Docker

**Development** (hot reload with [air](https://github.com/air-verse/air), source code mounted into the container, Postgres exposed on `DB_EXPOSE_PORT`):
```bash
make docker-dev
```

**Production** (static binary on a minimal Alpine image, running as a non-root user, Postgres only reachable inside the Docker network):
```bash
make docker-run
```

Stop the containers:
```bash
make docker-down
```

---

## 📜 License

Kaho BaaS is released under the **GNU Affero General Public License v3.0 (AGPL-3.0)**. See [LICENSE](LICENSE) for details.

**What does this mean for me?**

- **Using Kaho as a backend for your app** (calling its API from your web, mobile or server application) does **not** require you to open source your application.
- **Self-hosting an unmodified Kaho** for your own use or for your customers is allowed.
- **Modifying Kaho and offering it to others over a network** (for example as a hosted service) requires you to make the source code of your modified version available to its users under the AGPL-3.0.

For a commercial license without the AGPL-3.0 obligations, contact Kantor Hosting at [hello@kantorhosting.com](mailto:hello@kantorhosting.com).

The AGPL-3.0 covers the source code only. It does not grant permission to use the "Kaho" name for your own products or services.

---

## ⭐ Support the Project

If you find Kaho BaaS useful, consider giving us a ⭐ on GitHub! 😊
