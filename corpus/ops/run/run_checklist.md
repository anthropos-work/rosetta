# Anthropos Platform Run Checklist

**Date Started**: _______________
**Platform**: [ ] macOS / [ ] Linux

> Copy this checklist to `anthropos-dev/run_progress.md` to track your progress.

---

## 1. Docker Environment

- [ ] Docker daemon is running
- [ ] No conflicting Anthropos containers from other environments

## 2. Infrastructure Services

- [ ] PostgreSQL container started (`ant-rosetta-postgresql-1`)
- [ ] PostgreSQL accepting connections (health check passed)
- [ ] Redis container started (`ant-rosetta-redis-1`)
- [ ] Redis responding to ping (health check passed)

## 3. Backend Services

Choose one:

### Option A: Full Backend Stack (Recommended)
- [ ] Started with `--profile graphql`
- [ ] Sentinel service running
- [ ] Backend service running
- [ ] CMS service running
- [ ] Skiller service running
- [ ] Skillpath service running
- [ ] Storage service running
- [ ] Chronos service running
- [ ] Jobsimulation service running
- [ ] Intelligence service running
- [ ] GraphQL gateway running (http://localhost:5050)

### Option B: Minimal Backend
- [ ] Sentinel service running
- [ ] Backend service running
- [ ] CMS service running

## 4. Frontend (Next.js Web App)

- [ ] Dependencies installed (`node_modules` exists)
- [ ] Development server started (`pnpm dev:web`)
- [ ] Accessible at http://localhost:3000
- [ ] Login page or dashboard loads

## 5. Studio Services (Optional)

### Studio-Desk
- [ ] Dependencies installed
- [ ] `.env` file configured
- [ ] Development server started (`npm run dev`)
- [ ] Accessible at http://localhost:3100

### Studio-Room
- [ ] Python dependencies installed
- [ ] Test generation successful

---

## Service URLs Verified

| Service | URL | Status |
|---------|-----|--------|
| Frontend | http://localhost:3000 | [ ] OK |
| GraphQL | http://localhost:5050 | [ ] OK |
| Studio-Desk | http://localhost:3100 | [ ] OK / [ ] N/A |
| Backend API | http://localhost:8082 | [ ] OK |

---

## Notes / Issues Encountered

| Step | Issue | Resolution |
|------|-------|------------|
| | | |
| | | |
| | | |

---

## Quick Commands Reference

```bash
# Start everything from scratch
cd anthropos-dev/platform
docker compose -p ant-rosetta up -d postgresql redis
docker compose -p ant-rosetta --profile graphql up -d
cd ../next-web-app && pnpm dev:web

# Stop everything
cd anthropos-dev/platform
docker compose -p ant-rosetta down

# Check status
docker ps --filter "name=ant-rosetta" --format "table {{.Names}}\t{{.Status}}"

# View logs
docker compose -p ant-rosetta logs -f [service-name]
```
