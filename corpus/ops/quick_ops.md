# Quick Operations Reference

Common commands for working with the Anthropos platform.

---

## Database

### Access PostgreSQL CLI

```bash
docker exec -it ant-rosetta-postgresql-1 psql -U postgres
```

### List all schemas

```bash
docker exec ant-rosetta-postgresql-1 psql -U postgres -c "\dn"
```

### Run a SQL query

```bash
docker exec ant-rosetta-postgresql-1 psql -U postgres -c "SELECT * FROM users LIMIT 5;"
```

### Connect to a specific schema

```bash
docker exec -it ant-rosetta-postgresql-1 psql -U postgres -c "SET search_path TO skiller;" -c "\dt"
```

---

## Docker

### Check running containers

```bash
docker ps --filter "name=ant-rosetta" --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"
```

### View logs for a service

```bash
docker compose -p ant-rosetta logs -f backend
```

### Restart a single service

```bash
docker compose -p ant-rosetta restart backend
```

### Rebuild and restart a service

```bash
docker compose -p ant-rosetta up -d --build backend
```

### Stop everything

```bash
docker compose -p ant-rosetta down
```

### Stop and wipe data

```bash
docker compose -p ant-rosetta down -v
```

---

## Redis

### Access Redis CLI

```bash
docker exec -it ant-rosetta-redis-1 redis-cli
```

### Ping Redis

```bash
docker exec ant-rosetta-redis-1 redis-cli ping
```

### List all keys

```bash
docker exec ant-rosetta-redis-1 redis-cli KEYS "*"
```

---

## Frontend

### Start web app (dev)

```bash
cd anthropos-dev/next-web-app && pnpm dev:web
```

### Build web app

```bash
cd anthropos-dev/next-web-app && pnpm build:web
```

### Clean and reinstall deps

```bash
cd anthropos-dev/next-web-app && rm -rf node_modules && pnpm install
```

---

## Migrations

### Apply all migrations

```bash
cd anthropos-dev
(cd backend && atlas migrate apply --env local)
(cd cms && atlas migrate apply --env local)
(cd jobsimulation && atlas migrate apply --env local)
(cd skiller && atlas migrate apply --env local)
```

### Check migration status

```bash
cd anthropos-dev/backend && atlas migrate status --env local
```

---

## Git

### Pull all repos

```bash
cd anthropos-dev
for repo in platform backend cms jobsimulation skiller next-web-app studio-desk studio-room; do
  (cd "$repo" 2>/dev/null && git pull origin main) || true
done
```

---

## Ports Reference

| Service | Port |
|---------|------|
| Frontend | 3000 |
| GraphQL | 5050 |
| Backend | 8082 |
| Studio-Desk | 3100 |
| PostgreSQL | 5432 |
| Redis | 6379 |
| Directus | 8055 |
