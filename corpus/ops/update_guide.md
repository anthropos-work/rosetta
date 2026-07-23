# Updating the Anthropos Platform

This guide covers keeping your local development environment **up-to-date** with the latest code, dependencies, and database schemas.

> **Prerequisites**: This guide assumes you have completed the [Platform Setup Guide](setup_guide.md) and have a working local environment.

## When to Update

Run through this guide when:
- Starting work after being away (daily/weekly sync)
- A teammate mentions breaking changes
- You see dependency errors or schema mismatches
- Before starting work on a new feature

> **⚠️ Consolidation re-sync (v2.7 "july jitter", M246) — skillpath is decommissioned into `app`.**
> The skiller→app merge (v2.1) was one step of a program that consolidates every runtime engine into
> `app`. As of current `origin/main`, **`skillpath` is fully decommissioned**: it is gone from
> `repos.yml`, from `docker-compose.yml` (no skillpath service/container), and from the GraphQL
> federation — the router now composes **3 subgraphs** (app/backend, cms, jobsimulation), not 4. The
> skill-path **runtime session** table moved schemas: writes now go to **`public.skill_path_sessions`**,
> not `skillpath.skill_path_sessions` (the old schema table survives as an empty legacy husk).
> A `/stack-update` that crosses this consolidation is expected; the rext seeder tooling was re-pointed
> to `public.skill_path_sessions` in the same release so seeding a freshly-updated stack does not break.
> (The full corpus reconciliation of skillpath→app across all service docs is M247.) Proven green by a
> cold `/demo-up` on the consolidated platform at M246.

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
cd stack-dev/platform
make ps
```

### Stop Services

```bash
# Stop Docker services
make down

# Stop frontend (Ctrl+C in terminal, or)
pkill -f "pnpm dev:web" 2>/dev/null || true
```

> **Caveat**: `make down` is plain `docker compose down` with no profile or `--remove-orphans`. If the previous run was on an older `docker-compose.yml` that included services since removed (e.g. chronos, intelligence), those orphaned containers will keep the network alive and `make down` will fail with `Network anthropos_app-network  Resource is still in use`. In that case run:
>
> ```bash
> docker compose --profile all down --remove-orphans
> ```
>
> This stops every service the current compose file knows about, plus any orphaned containers from older revisions. Use it after pulling platform changes that remove or rename services. See troubleshooting below for context.

---

## 2. Update Repository Code

Pull latest changes from all repositories using the Makefile.

### Navigate to Platform Directory

```bash
cd stack-dev/platform
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
cd stack-dev/next-web-app
pnpm install
```

*Verification*: No errors during install, `node_modules` updated.

> **Node version**: As of 2026-05, `next-web-app/package.json` requires `node >=24.0.0`. If you're on an older Node you'll see `WARN  Unsupported engine` and pnpm will refuse to wipe `node_modules` from a non-TTY shell with `ERR_PNPM_ABORTED_REMOVE_MODULES_DIR_NO_TTY`. Fix:
>
> ```bash
> nvm install 24 && nvm use 24
> pnpm install
> ```
>
> Use `nvm alias default 24` to make it stick across shells.

### Studio-Desk

```bash
cd stack-dev/studio-desk
npm install
```

### Studio-Room (Python)

```bash
cd stack-dev/studio-room
pip3 install -r requirements.txt --upgrade
```

### Go Services (if building locally)

Go modules update automatically on build, but you can force an update:

```bash
cd stack-dev/backend
go mod download
go mod tidy
```

---

## 4. Run Database Migrations

If schema changes were made, apply new migrations.

### Apply All Migrations

```bash
cd stack-dev/platform
make migrate
```

This automatically runs Atlas migrations for all repos with `migrations: true` in `repos.yml` (currently: app, cms, jobsimulation — **skillpath is decommissioned into `app`** as of v2.7; see the consolidation re-sync note above).

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
cd stack-dev/platform
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
cd stack-dev/next-web-app
pnpm dev:web
```

Open http://localhost:3000 and verify the application loads.

---

## 6. Quick Update Scenarios

### Scenario: Daily Sync (Quick)

Minimal update when starting work:

```bash
cd stack-dev/platform
make pull              # Pull all repos
make up                # Rebuild and start
cd ../next-web-app && pnpm install && pnpm dev:web
```

### Scenario: Weekly Sync (Full)

Comprehensive update:

```bash
cd stack-dev/platform
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
cd stack-dev/platform
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
cd stack-dev/platform
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

### `make down` Fails: "Network anthropos_app-network  Resource is still in use"

Containers from a previous `docker-compose.yml` revision are still running but no longer listed in the current compose file (typical after pulling commits that remove a service). `docker compose down` doesn't see them, so the network can't be torn down.

**Fix**:
```bash
docker compose --profile all down --remove-orphans
```

This catches every profile member plus any container whose service no longer exists in the current compose file. Encountered today (2026-05-11) after `chronos` and `intelligence` were removed from `repos.yml`; see ops report `op_20260511_173232_update_orphaned_chronos_intelligence.md`.

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
