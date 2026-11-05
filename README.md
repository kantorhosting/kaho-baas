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

We welcome contributions! Please read the [Contributing Guide](CONTRIBUTING.md) before opening a Pull Request. All contributors must sign the [Contributor License Agreement](CLA.md), and everyone taking part in the project is expected to follow our [Code of Conduct](CODE_OF_CONDUCT.md).

Found a security issue? **Do not open a public issue.** See our [Security Policy](SECURITY.md).

---

## 🛠 Configuration

Environment variables for configuring Kaho BaaS:
```env
PORT=8080
APP_ENV=local
BLUEPRINT_DB_HOST=psql_bp
BLUEPRINT_DB_PORT=5432
BLUEPRINT_DB_DATABASE=blueprint
BLUEPRINT_DB_USERNAME=melkey
BLUEPRINT_DB_PASSWORD=password1234
BLUEPRINT_DB_SCHEMA=public
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
