# Updating the Anthropos Platform

This guide covers keeping your local development environment **up-to-date** with the latest code, dependencies, and database schemas.

> **Prerequisites**: This guide assumes you have completed the [Platform Setup Guide](setup_guide.md) and have a working local environment.

## When to Update

Run through this guide when:
- Starting work after being away (daily/weekly sync)
- A teammate mentions breaking changes
- You see dependency errors or schema mismatches
- Before starting work on a new feature

---

## Update Execution Guidelines (UPDATE STEP)

When executing each step:

1. **Check Current State**: Verify what needs updating before making changes.
2. **Pull Before Build**: Always fetch latest code before rebuilding.
3. **Handle Conflicts**: If git conflicts occur, resolve before proceeding.
4. **Verify After Update**: Confirm services still work after updates.

---

## 1. Stop Running Services

Before updating, stop any running services to avoid conflicts.

### Check What's Running

```bash
cd anthropos-dev/platform
make ps
```

### Stop Services

```bash
# Stop Docker services
make down

# Stop frontend (Ctrl+C in terminal, or)
pkill -f "pnpm dev:web" 2>/dev/null || true
```

---

## 2. Update Repository Code

Pull latest changes from all repositories using the Makefile.

### Navigate to Platform Directory

```bash
cd anthropos-dev/platform
```

### Update All Repositories

The `make pull` command updates all repos defined in `repos.yml`. It automatically:
- Stashes dirty changes before pulling
- Checks out main/master branch
- Pulls with rebase

```bash
make pull
```

### Check Repository Status

```bash
make status
```

This shows branch, dirty status, and commits behind for all repos.

### Handle Git Conflicts

If you see merge conflicts:

```bash
cd [repo-with-conflict]
git status  # See conflicting files

# Option 1: Stash your changes and pull
git stash
git pull origin main
git stash pop  # Reapply your changes

# Option 2: Reset to remote (WARNING: loses local changes)
git fetch origin
git reset --hard origin/main
```

---

## 3. Update Dependencies

After pulling code, install any new or updated dependencies.

### Frontend (Next.js Monorepo)

```bash
cd anthropos-dev/next-web-app
pnpm install
```

*Verification*: No errors during install, `node_modules` updated.

### Studio-Desk

```bash
cd anthropos-dev/studio-desk
npm install
```

### Studio-Room (Python)

```bash
cd anthropos-dev/studio-room
pip3 install -r requirements.txt --upgrade
```

### Go Services (if building locally)

Go modules update automatically on build, but you can force an update:

```bash
cd anthropos-dev/backend
go mod download
go mod tidy
```

---

## 4. Run Database Migrations

If schema changes were made, apply new migrations.

### Apply All Migrations

```bash
cd anthropos-dev/platform
make migrate
```

This automatically runs Atlas migrations for all repos with `migrations: true` in `repos.yml` (app, cms, jobsimulation, skiller, skillpath, chronos).

### Apply Single Service Migration

```bash
make migrate S=cms
```

*Expected*: Each command shows applied migrations or "No migration files to apply".

### If Migrations Fail

**Connection refused**: PostgreSQL not started. Run `make up` first.

**Dirty migration state**: Check status, then consider `make reset-db` (WARNING: data loss).

---

## 5. Rebuild and Start Services

The `make up` command automatically rebuilds images from local code (`--build` flag) and starts services.

```bash
cd anthropos-dev/platform
make up
```

### Verify Updated Environment

```bash
# Check all containers are healthy
make ps

# Test backend
curl -s http://localhost:8082/health || echo "Backend not responding"

# Test GraphQL
curl -s http://localhost:5050/health || echo "GraphQL not responding"
```

### Start and Test Frontend

```bash
cd anthropos-dev/next-web-app
pnpm dev:web
```

Open http://localhost:3000 and verify the application loads.

---

## 6. Quick Update Scenarios

### Scenario: Daily Sync (Quick)

Minimal update when starting work:

```bash
cd anthropos-dev/platform
make pull              # Pull all repos
make up                # Rebuild and start
cd ../next-web-app && pnpm install && pnpm dev:web
```

### Scenario: Weekly Sync (Full)

Comprehensive update:

```bash
cd anthropos-dev/platform
make down              # Stop everything
make pull              # Pull all repos
make up                # Rebuild and start
make migrate           # Apply any new migrations

# Update frontend deps and start
cd ../next-web-app && pnpm install && pnpm dev:web
```

### Scenario: After Major Release

When significant changes are announced:

```bash
cd anthropos-dev/platform
make down              # Stop everything
make pull              # Pull all repos
make reset-db          # Fresh database + migrations
make up                # Rebuild and start

# Clean and reinstall frontend deps
cd ../next-web-app && rm -rf node_modules && pnpm install && pnpm dev:web
```

---

## 8. Troubleshooting

### "Package not found" After Pull

Dependencies changed but weren't installed.

```bash
# Frontend
cd next-web-app && pnpm install

# Studio-Desk
cd studio-desk && npm install
```

### Migration Conflicts

Schema out of sync with code.

```bash
# Check migration status
cd [service]
atlas migrate status --env local

# If stuck, full reset (WARNING: data loss)
cd anthropos-dev/platform
make reset-db
```

### Docker Build Errors After Update

Image cache may be stale.

```bash
# Clean Docker cache
docker system prune -f

# Rebuild
make up
```

### Git Conflicts During Pull

Local changes conflict with remote.

```bash
# See what's conflicting
git status

# Option 1: Keep your changes
git stash
git pull origin main
git stash pop

# Option 2: Discard local changes
git checkout -- .
git pull origin main
```

### Frontend Compilation Errors After Update

TypeScript types may be out of sync.

```bash
cd next-web-app
pnpm clean  # Clean build artifacts
pnpm install  # Reinstall deps
pnpm dev:web  # Try again
```

---

## 9. Maintenance Guidelines

This guide follows the same iterative improvement pattern as other ops guides:

- **Document Issues**: When you discover new update scenarios, add them.
- **Keep Scripts Updated**: If repo structure changes, update the commands.
- **Track Patterns**: Common issues should become troubleshooting entries.

### Related Guides

- [Platform Setup](setup_guide.md) - Initial environment build
- [Platform Run](run_guide.md) - Starting the platform
