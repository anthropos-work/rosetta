---
name: ant-update
description: Sync Anthropos code, dependencies, and database schemas with latest changes
argument-hint: [scenario: 'daily' | 'weekly' | 'full']
---

# Anthropos Platform Update

Execute platform updates by following `corpus/ops/platform-update/update_guide.md` while applying UPDATE STEP guidelines and verifying compatibility.

## Your Mission

1. **Follow the guide**: Use `corpus/ops/platform-update/update_guide.md` as your source of truth
2. **Apply UPDATE STEP guidelines**: Check state, pull before build, handle conflicts, verify after
3. **Track progress locally**: Copy checklist to `anthropos-dev/update_progress.md` and update as you go
4. **Auto-improve docs**: Update update_guide.md when you discover issues or better approaches

## UPDATE STEP Guidelines (Apply to Every Step)

### 1. Check Current State
Verify what needs updating before making changes:
```bash
# Check running services
docker ps --filter "name=ant-rosetta"

# Check git status
cd anthropos-dev/backend && git status
```

### 2. Pull Before Build
Always fetch latest code before rebuilding:
```bash
git pull origin main
```

### 3. Handle Conflicts
If git conflicts occur, resolve before proceeding:
```bash
git stash
git pull origin main
git stash pop
```

### 4. Verify After Update
Confirm services still work after updates:
```bash
docker ps --filter "name=ant-rosetta" --format "{{.Status}}"
curl -s http://localhost:8082/health
```

## Initial Setup

1. Copy checklist: `cp corpus/ops/platform-update/update_checklist.md anthropos-dev/update_progress.md`
2. Read `corpus/ops/platform-update/update_guide.md` to understand the process
3. Navigate to `anthropos-dev/` workspace

## Update Scenarios

### Daily Sync (Quick)
Minimal update for starting work:
- Pull main repos (platform, backend, next-web-app)
- Update frontend deps
- Restart services

### Weekly Sync (Full)
Comprehensive update:
- Pull all repos
- Update all dependencies
- Apply migrations
- Rebuild Docker images

### Major Release
After significant changes:
- Stop everything with volume removal
- Fresh pull all repos
- Clean reinstall dependencies
- Apply migrations to fresh database
- Fresh Docker build

## Request Confirmation

**ALWAYS ask user before**:
- Stopping running services
- Pulling repository changes
- Running database migrations
- Rebuilding Docker images
- Data-destructive operations (down -v)

Use AskUserQuestion tool.

## Update Sequence

### Phase 1: Stop Running Services
```bash
cd anthropos-dev/platform
docker compose -p ant-rosetta down

# Stop frontend (Ctrl+C or)
pkill -f "pnpm dev:web" 2>/dev/null || true
```

### Phase 2: Update Repository Code
```bash
cd anthropos-dev

# Platform and backend
(cd platform && git pull origin main)
(cd backend && git pull origin main)
(cd cms && git pull origin main)

# Frontend
(cd next-web-app && git pull origin main)

# Studio (if used)
(cd studio-desk && git pull origin main)
(cd studio-room && git pull origin main)
```

### Phase 3: Update Dependencies
```bash
# Frontend
(cd next-web-app && pnpm install)

# Studio-Desk
(cd studio-desk && npm install)

# Studio-Room
(cd studio-room && pip3 install -r requirements.txt --upgrade)
```

### Phase 4: Database Migrations
```bash
# Start PostgreSQL first
docker compose -p ant-rosetta up -d postgresql
sleep 5

# Apply migrations
(cd backend && atlas migrate apply --env local)
(cd cms && atlas migrate apply --env local)
(cd jobsimulation && atlas migrate apply --env local)
(cd skiller && atlas migrate apply --env local)
```

### Phase 5: Rebuild Docker Images
```bash
cd anthropos-dev/platform

# Rebuild specific service
docker compose -p ant-rosetta build backend

# Rebuild all
docker compose -p ant-rosetta build

# Force fresh build (no cache)
docker compose -p ant-rosetta build --no-cache
```

### Phase 6: Verify Updated Environment
```bash
# Start services
docker compose -p ant-rosetta --profile graphql up -d

# Health checks
docker ps --filter "name=ant-rosetta" --format "table {{.Names}}\t{{.Status}}"
curl -s http://localhost:8082/health
curl -s http://localhost:5050/health
```

## Error Handling

1. **Do NOT skip errors** - resolve them first
2. Document error message verbatim
3. Research solution
4. Test fix
5. Add to update_guide.md troubleshooting section
6. Continue

## Progress Tracking

Use TodoWrite for high-level tracking:
```markdown
- [x] Copied checklist to anthropos-dev/
- [x] Services stopped
- [ ] Repositories updated
- [ ] Dependencies installed
- [ ] Migrations applied
- [ ] Docker rebuilt
- [ ] Services verified
```

Use local checklist (`anthropos-dev/update_progress.md`) for detailed tracking.

## Critical Rules

- Work in `anthropos-dev/` scratchpad only
- Use `-p ant-rosetta` for Docker isolation
- Ask before every destructive operation
- Stop services before pulling code
- Handle git conflicts before continuing
- Verify health after updates
- Follow the guide - don't improvise unless needed
- Update docs immediately when improvements found

## Troubleshooting Quick Reference

### "Package not found" After Pull
```bash
cd next-web-app && pnpm install
```

### Migration Conflicts
```bash
atlas migrate status --env local
# May need to reset: docker compose -p ant-rosetta down -v
```

### Docker Build Errors
```bash
docker system prune -f
docker compose -p ant-rosetta build --no-cache
```

### Git Conflicts
```bash
git stash
git pull origin main
git stash pop  # or git checkout -- . to discard
```

## Success Criteria

Update complete when:
1. All repos pulled without conflicts
2. Dependencies installed without errors
3. Migrations applied successfully
4. Docker services running and healthy
5. Frontend loads without errors
6. Documentation improvements committed (if any)

**Follow `corpus/ops/platform-update/update_guide.md` as your primary reference. Apply these guidelines to update the platform reliably.**
