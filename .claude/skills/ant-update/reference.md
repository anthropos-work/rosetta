# Anthropos Update Skill - Technical Reference

## Architecture

This skill implements a **stop-update-verify** loop:
1. Stop running services to avoid conflicts
2. Update code, dependencies, and schemas
3. Rebuild affected components
4. Verify everything works together

## File References

### Documentation Sources
- `corpus/ops/platform-update/update_guide.md` - Master update guide
- `corpus/ops/platform-update/update_checklist.md` - Progress tracking checklist
- `corpus/architecture/` - Platform architecture context

### Working Directory
- `anthropos-dev/` - Git-ignored scratchpad for all platform activities

## Repository Update Order

```
1. platform (Docker configs, shared .env)
2. backend (main API)
3. cms (content management)
4. jobsimulation
5. skiller
6. next-web-app (frontend)
7. studio-desk
8. studio-room
9. cms/studio (symlinked studio-room)
```

## Tool Usage Strategy

### Read-Only Tools (Verification)
- `Bash` for status checks (`git status`, `docker ps`, `npm list`)
- `Read` for documentation and configs

### Write Tools (Execution & Documentation)
- `Bash` for git operations, package installation (after confirmation)
- `Edit` for updating documentation (auto-improvement)

### Interactive Tools
- `AskUserQuestion` for all confirmations before:
  - Stopping services
  - Pulling changes
  - Running migrations
  - Rebuilding images
  - Data-destructive operations

## UPDATE STEP Guidelines Implementation

### Guideline 1: Check Current State

```bash
# Check running services
docker ps --filter "name=ant-rosetta" --format "table {{.Names}}\t{{.Status}}"

# Check git status for each repo
for repo in platform backend cms jobsimulation skiller next-web-app; do
  echo "=== $repo ==="
  (cd "anthropos-dev/$repo" && git status -s) 2>/dev/null || echo "  Not found"
done

# Check for uncommitted changes
(cd anthropos-dev/backend && git diff --stat)
```

### Guideline 2: Pull Before Build

```bash
# Standard pull
git pull origin main

# If behind with local commits
git fetch origin
git rebase origin/main

# Check what changed
git log --oneline HEAD..origin/main
```

### Guideline 3: Handle Conflicts

```bash
# Option 1: Stash local changes
git stash
git pull origin main
git stash pop

# Option 2: Discard local changes (after confirmation)
git fetch origin
git reset --hard origin/main

# Option 3: Merge conflicts manually
git status  # See conflicting files
# Edit files to resolve
git add .
git commit -m "Resolve merge conflicts"
```

### Guideline 4: Verify After Update

```bash
# Start services
docker compose -p ant-rosetta --profile graphql up -d

# Wait for startup
sleep 10

# Health checks
docker ps --filter "name=ant-rosetta" --format "table {{.Names}}\t{{.Status}}"
curl -s http://localhost:8082/health || echo "Backend not responding"
curl -s http://localhost:5050/health || echo "GraphQL not responding"

# Start frontend and verify
cd anthropos-dev/next-web-app && pnpm dev:web &
sleep 10
curl -s http://localhost:3000 > /dev/null && echo "Frontend OK" || echo "Frontend not responding"
```

## Update Scenario Details

### Daily Sync

```bash
cd anthropos-dev

# Stop services
cd platform && docker compose -p ant-rosetta down && cd ..

# Pull main repos
(cd platform && git pull origin main)
(cd backend && git pull origin main)
(cd next-web-app && git pull origin main)

# Update frontend deps
(cd next-web-app && pnpm install)

# Restart
cd platform && docker compose -p ant-rosetta --profile graphql up -d
```

**When to use**: Starting daily work, quick sync

### Weekly Sync

```bash
cd anthropos-dev

# Full stop
cd platform && docker compose -p ant-rosetta down && cd ..

# Pull all repos
for repo in platform backend cms jobsimulation skiller next-web-app studio-desk studio-room cms/studio; do
  echo "Updating $repo..."
  (cd "$repo" 2>/dev/null && git pull origin main) || echo "  Skipped: not found"
done

# Update all dependencies
(cd next-web-app && pnpm install)
(cd studio-desk && npm install)
(cd studio-room && pip3 install -r requirements.txt --upgrade)

# Apply migrations
cd platform && docker compose -p ant-rosetta up -d postgresql && sleep 5 && cd ..
(cd backend && atlas migrate apply --env local)
(cd cms && atlas migrate apply --env local)
(cd jobsimulation && atlas migrate apply --env local)
(cd skiller && atlas migrate apply --env local)

# Rebuild and start
cd platform
docker compose -p ant-rosetta --profile graphql up -d --build
```

**When to use**: Weekly maintenance, after being away

### Major Release

```bash
cd anthropos-dev

# Full stop with volume removal
cd platform && docker compose -p ant-rosetta down -v && cd ..

# Fresh pull (discard local changes)
for repo in platform backend cms jobsimulation skiller next-web-app studio-desk studio-room; do
  (cd "$repo" 2>/dev/null && git fetch origin && git reset --hard origin/main) || true
done

# Clean reinstall
(cd next-web-app && rm -rf node_modules && pnpm install)
(cd studio-desk && rm -rf node_modules && npm install)

# Fresh database and migrations
cd platform && docker compose -p ant-rosetta up -d postgresql redis && sleep 10 && cd ..
(cd backend && atlas migrate apply --env local)
(cd cms && atlas migrate apply --env local)
(cd jobsimulation && atlas migrate apply --env local)
(cd skiller && atlas migrate apply --env local)

# Fresh build
cd platform && docker compose -p ant-rosetta --profile graphql up -d --build
```

**When to use**: Breaking changes announced, major version bumps

## Migration Patterns

### Check Migration Status

```bash
cd anthropos-dev/backend
atlas migrate status --env local
```

### Apply Migrations

```bash
# Ensure PostgreSQL is running
docker compose -p ant-rosetta up -d postgresql
sleep 5

# Apply for each service
(cd backend && atlas migrate apply --env local)
(cd cms && atlas migrate apply --env local)
(cd jobsimulation && atlas migrate apply --env local)
(cd skiller && atlas migrate apply --env local)
```

### Handle Migration Failures

```yaml
Scenario: Migration partially applied
Diagnosis:
  1. Check status: atlas migrate status --env local
  2. Check schema: docker exec ant-rosetta-postgresql-1 psql -U postgres -c "\dn"
Recovery:
  - Clean state: docker compose -p ant-rosetta down -v
  - Re-apply: Start fresh with all migrations
```

## Dependency Update Patterns

### Frontend (pnpm)

```bash
cd anthropos-dev/next-web-app

# Install (respects lockfile)
pnpm install

# Update all packages
pnpm update

# Check for outdated
pnpm outdated
```

### Studio-Desk (npm)

```bash
cd anthropos-dev/studio-desk

# Install
npm install

# Update
npm update

# Check for outdated
npm outdated
```

### Studio-Room (pip)

```bash
cd anthropos-dev/studio-room

# Install/upgrade
pip3 install -r requirements.txt --upgrade

# Check installed versions
pip3 list
```

### Go Services

```bash
cd anthropos-dev/backend

# Download dependencies
go mod download

# Tidy up
go mod tidy

# Verify
go mod verify
```

## Docker Rebuild Patterns

### Specific Service

```bash
docker compose -p ant-rosetta build backend
docker compose -p ant-rosetta up -d backend
```

### All Services

```bash
docker compose -p ant-rosetta build
docker compose -p ant-rosetta up -d
```

### Force Fresh Build

```bash
# Clear build cache
docker system prune -f

# Build without cache
docker compose -p ant-rosetta build --no-cache
```

## Error Recovery Patterns

### Pattern 1: Git Conflict

```yaml
Symptoms: "CONFLICT (content)" during pull
Recovery:
  1. Check status: git status
  2. Option A - Stash: git stash && git pull && git stash pop
  3. Option B - Discard: git checkout -- . && git pull
  4. Option C - Manual: Edit files, git add ., git commit
```

### Pattern 2: Package Not Found

```yaml
Symptoms: "Module not found" or "Package not found" at runtime
Recovery:
  1. Identify: Which package?
  2. Reinstall: pnpm install (frontend) or npm install (studio-desk)
  3. If persists: rm -rf node_modules && pnpm install
```

### Pattern 3: Migration Error

```yaml
Symptoms: "migration failed" or "schema mismatch"
Recovery:
  1. Check status: atlas migrate status --env local
  2. If dirty: May need fresh database
  3. Fresh start: docker compose -p ant-rosetta down -v
  4. Re-apply all migrations
```

### Pattern 4: Docker Build Error

```yaml
Symptoms: Build fails during docker compose build
Recovery:
  1. Check logs: docker compose -p ant-rosetta build [service] 2>&1 | tail -100
  2. SSH issues: eval "$(ssh-agent -s)" && ssh-add
  3. Cache issues: docker system prune -f
  4. Retry: docker compose -p ant-rosetta build --no-cache
```

### Pattern 5: Service Won't Start After Update

```yaml
Symptoms: Container exits or unhealthy after rebuild
Recovery:
  1. Check logs: docker compose -p ant-rosetta logs [service]
  2. Missing env: Verify platform/.env has all required keys
  3. Schema mismatch: Apply pending migrations
  4. Fresh start: docker compose -p ant-rosetta down -v && restart
```

## Progress Tracking

### Two-Level Tracking System

**Level 1: Local Checklist** (`anthropos-dev/update_progress.md`)
- Copied from update_checklist.md at start
- Updated as each phase completes
- Contains issues encountered

**Level 2: TodoWrite Tool**
- High-level phase tracking
- Blockers requiring user input

Example TodoWrite tracking:
```markdown
- [x] Phase 1: Services stopped
- [x] Phase 2: platform pulled
- [x] Phase 2: backend pulled
- [ ] Phase 2: cms pulled
- [ ] Phase 3: Dependencies updated
- [ ] Phase 4: Migrations applied
```

## Success Validation

At completion, verify:

1. **Git Clean**: `git status` shows clean in all repos
2. **Deps Installed**: No missing package errors
3. **Migrations Applied**: `atlas migrate status` shows no pending
4. **Services Healthy**: `docker ps` shows all healthy
5. **Frontend Works**: http://localhost:3000 loads

## Skill Invocation

```bash
# Interactive - asks for scenario
/ant-update

# Specific scenario
/ant-update daily        # Quick daily sync
/ant-update weekly       # Full weekly update
/ant-update full         # Major release update

# Specific phase
/ant-update pull         # Just pull repos
/ant-update deps         # Just update dependencies
/ant-update migrate      # Just run migrations
/ant-update rebuild      # Just rebuild Docker images
```

## Integration with Other Skills

### Workflow Order

1. `/ant-update` - Sync code and dependencies
2. `/ant-run` - Start the platform
3. `/ant-integrate` - Document any issues found

### Handoff to ant-run

After successful update:
1. Services are stopped
2. Code is current
3. Dependencies are installed
4. Migrations are applied
5. User should run `/ant-run` to start

### Handoff to ant-integrate

If update reveals issues:
1. Document in update_progress.md
2. Note patterns that failed
3. Suggest `/ant-integrate D` to improve docs
