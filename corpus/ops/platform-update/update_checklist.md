# Anthropos Platform Update Checklist

**Date**: _______________
**Update Type**: [ ] Daily Sync / [ ] Weekly Sync / [ ] Major Release

> Copy this checklist to `anthropos-dev/update_progress.md` to track your progress.

---

## 1. Pre-Update

- [ ] Running services stopped (Docker, frontend)
- [ ] No uncommitted local changes (or stashed)

## 2. Repository Updates

- [ ] `platform` pulled
- [ ] `backend` pulled
- [ ] `cms` pulled
- [ ] `jobsimulation` pulled
- [ ] `skiller` pulled
- [ ] `next-web-app` pulled
- [ ] `studio-desk` pulled
- [ ] `studio-room` pulled
- [ ] `cms/studio` pulled (synced with studio-room)

## 3. Dependency Updates

- [ ] Frontend: `pnpm install` in next-web-app
- [ ] Studio-Desk: `npm install`
- [ ] Studio-Room: `pip3 install -r requirements.txt`

## 4. Database Migrations

- [ ] PostgreSQL running
- [ ] Backend migrations applied
- [ ] CMS migrations applied
- [ ] Jobsimulation migrations applied
- [ ] Skiller migrations applied

## 5. Docker Rebuild (if needed)

- [ ] Services rebuilt (`docker compose build`)
- [ ] Images up to date

## 6. Verification

- [ ] Docker services running and healthy
- [ ] Backend API responding (http://localhost:8082)
- [ ] GraphQL gateway responding (http://localhost:5050)
- [ ] Frontend loads (http://localhost:3000)

---

## Quick Commands

```bash
# Stop services
docker compose -p ant-rosetta down

# Pull all repos
for repo in platform backend cms jobsimulation skiller next-web-app studio-desk studio-room cms/studio; do
  (cd "$repo" 2>/dev/null && git pull origin main) || true
done

# Update deps
(cd next-web-app && pnpm install)

# Run migrations
docker compose -p ant-rosetta up -d postgresql && sleep 5
(cd backend && atlas migrate apply --env local)
(cd cms && atlas migrate apply --env local)
(cd jobsimulation && atlas migrate apply --env local)
(cd skiller && atlas migrate apply --env local)

# Start services
docker compose -p ant-rosetta --profile graphql up -d
```

---

## Issues Encountered

| Step | Issue | Resolution |
|------|-------|------------|
| | | |
| | | |
| | | |
