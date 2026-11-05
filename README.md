# Kaho BaaS - Open-Source Backend as a Service

Kaho BaaS is a high-performance, open-source **Backend as a Service (BaaS)** built with Go. It provides authentication, database management, real-time sync, and serverless functions—so you can focus on building great applications while Kaho handles the backend. Scalable, flexible, and cloud-ready! 🚀

## ✨ Features

✅ **Authentication & Authorization** – Secure user authentication with JWT, OAuth, and more.  
✅ **Database Management** – Scalable and easy-to-use database solutions.  
✅ **Real-Time & Offline Sync** – Keep your data in sync across devices seamlessly.  
✅ **Serverless Functions** – Deploy custom backend logic without managing servers.  
✅ **REST & GraphQL API** – Access your data effortlessly with modern API support.  
✅ **Self-Hosted & Cloud Ready** – Deploy anywhere, from local setups to cloud platforms.  

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

We welcome contributions! To contribute:
1. Fork this repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Commit your changes (`git commit -m "Add new feature"`).
4. Push to the branch (`git push origin feature-branch`).
5. Open a Pull Request.

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

Kaho BaaS is released under the **MIT License**. See [LICENSE](LICENSE) for details.

---

## ⭐ Support the Project

If you find Kaho BaaS useful, consider giving us a ⭐ on GitHub! 😊
