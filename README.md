# Go User Activity Logging (Local Dev + ELK Stack)

A lightweight Go service for collecting and storing **user activity logs** using:

- **Go + Fiber** (API)
- **PostgreSQL** (storage)
- **Elasticsearch + Logstash + Kibana** (observability)
- **Docker Compose** (local development)
- **Air** (hot‑reload for Go)

This setup is ideal for experimenting with log pipelines, Go service patterns, and local ELK workflows.

---

## 🚀 Features

- REST API service in Go (hot‑reload enabled)
- Logs processed through Logstash into Elasticsearch
- Kibana dashboard for querying logs
- PostgreSQL table for structured activity logging
- Fully containerized local environment

---

## 📦 Requirements

- Docker + Docker Compose
- Go (optional if using only containers)

---

## 🔧 Local Development

### Start the full environment

```bash
docker-compose up --build
```

### Follow logs for specific services

```bash
docker-compose logs -f go-service
docker-compose logs -f logstash
docker-compose logs -f elasticsearch
```

---

## ▶️ API Service (Go)

The Go server auto‑reloads thanks to **Air**.

If you want to run only the Go service:

```bash
docker-compose up go-service
```

---

## 🗄 Database (init.sql)

The PostgreSQL instance initializes with:

```sql
CREATE TABLE user_activity (
    id SERIAL PRIMARY KEY,
    user_id TEXT NOT NULL,
    action TEXT NOT NULL,
    metadata JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO user_activity (user_id, action, metadata)
VALUES 
  ('user123', 'login', '{"ip":"192.168.1.10"}'),
  ('user456', 'view_page', '{"page":"home"}'),
  ('user789', 'purchase', '{"item":"book","price":19.99}');
```

---

## 📊 ELK Stack

- **Elasticsearch**
- **Logstash** (pipeline from Go logs → Elasticsearch)
- **Kibana** (port 5601)

Open Kibana:

http://localhost:5601

---

## 🐳 Production Build Example

```
docker build -f Dockerfile -t gouserlog:prod .
docker run -p 8080:8080 gouserlog:prod
```

or with Docker Compose:

```
docker-compose -f docker-compose.prod.yml up --build
```

---

## 📝 Notes

- Hot‑reload is enabled in the **api** container via `air`.
- This is a public‑friendly repo structure for collaboration.

---

## 🤝 Contributions

PRs & issues welcome!

---

## 📄 License

MIT
