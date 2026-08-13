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
docker exec -it ant-rosetta-postgresql-1 psql -U postgres -c "SET search_path TO cms;" -c "\dt"
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
cd stack-dev/next-web-app && pnpm dev:web
```

### Build web app

```bash
cd stack-dev/next-web-app && pnpm build:web
```

### Clean and reinstall deps

```bash
cd stack-dev/next-web-app && rm -rf node_modules && pnpm install
```

---

## Migrations

### Apply all migrations

> **⚠️ `app` is the ONLY repo with migrations to run**, and the directory is `app` — **not `backend`**.
> `backend` is the *deployed service name*; the repo `make init` clones is `app`, so `cd backend` fails
> with *no such file or directory*. `repos.yml` states the rule in the platform's own words: the folded
> repos *"own no local schema, no compose service and no clone entry here."* The `cms` and
> `jobsimulation` legs this recipe used to carry named two directories `make init` no longer creates
> **and** two schemas the platform no longer creates — running them against a current stack is the
> failure gate clause 4 exists to prevent. Use `make migrate` from `platform/` unless you need atlas
> directly.

```bash
cd stack-dev
(cd app && atlas migrate apply --env local)
```

### Check migration status

```bash
cd stack-dev/app && atlas migrate status --env local
```

---

## Git

### Pull all repos

```bash
cd stack-dev
for repo in platform backend cms jobsimulation next-web-app studio-desk studio-room; do
  (cd "$repo" 2>/dev/null && git pull origin main) || true
done
```

---

## Ports Reference

| Service | Port |
|---------|------|
| Frontend | 3000 |
| GraphQL | **8082** (`/graphql/query`, served by `backend`; the `:5050` Cosmo router was deleted from compose at platform `2adcf71`) |
| Backend | 8082 |
| Studio-Desk | 3100 |
| PostgreSQL | 5432 |
| Redis | 6379 |

> **No local Directus by default.** The platform compose has no directus service — content is read **live from
> prod** (`content.anthropos.work`). A local Directus (port 8055, offset on additional stacks) exists only when
> the v1.5 "prop room" tooling stands one up: demo-default / dev-opt-in (`--local-content`). See
> [`directus-local.md`](./directus-local.md).
