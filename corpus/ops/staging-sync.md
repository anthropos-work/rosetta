# Staging Sync — Daily Force-Reset to `origin/main`

The sync routine that keeps every Anthropos personal staging (Ithaca, Calypso, your `wip-*`) mirroring `origin/main` HEAD across every repo. **Force-main, no feature branches, ever.** This doc covers cadence, mechanics, skip-worktree handling, and recovery.

Companion to:

- [`staging-bringup.md`](./staging-bringup.md) — fresh-VM onboarding (where this routine gets installed).
- [`staging-clerk.md`](./staging-clerk.md) — dev Clerk app shared across all stagings.

---

## The rule: staging is always `main`

> Every Anthropos staging clone must always be on `origin/main` HEAD, across every repo, and the docker stack must always be built from that. **No feature branches on staging, ever.**

The rule was hardened by an incident: on 2026-05-14 Ithaca's `next-web-app` had been sitting on a feature branch for 8 days, 164 commits behind `origin/main`. Talk-to-Data and Workforce Intelligence (both shipped on `main` during that window) were silently missing from staging the whole time. Stefano discovered it by trying to demo a feature that wasn't there.

The previous sync logic — "preserve feature branches, fetch-only when not on main" — was the bug: it treated feature branches as a respected user choice instead of as drift.

The current sync **force-resets every repo to `origin/main` every run, no exceptions**, and the hourly drift script flags any repo on a feature branch as **WARN** so you get advance warning before the 06:00 UTC reset.

---

## What if I want to test a feature branch with prod-shape data?

**Not on a staging host.** You have three other options:

- **Best:** spawn an `--unmanaged` agentspace workspace against a local clone on your laptop — same data backend, but on your machine, no shared infrastructure risk.
- **OK:** spin up a new tailnet `wip-<your-initials>` host with its own dump restore. It will *not* be added to the staging-sync routine; you control its branch.
- **Not OK:** check out a feature branch on `ithaca:/home/devops/<repo>` or `calypso:/home/devops/<repo>`. The 06:00 UTC sync will force-overwrite it and your work will live in the safety patches dir until you manually `git apply` it back.

---

## Architecture (two scripts, two timers)

| Cadence              | Script                                                | What it does                                                                                                                                                                                                            |
| -------------------- | ----------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Hourly               | `~/.local/bin/anthropos-staging-drift.sh`             | Parallel `git fetch --prune` across 16 repos (~10s wall). Writes `drift.json` + `drift.summary` to `~/.local/state/anthropos-staging-sync/`. Surfaces drift + flags feature-branched repos as **WARN**.                 |
| Daily<br>06:00 UTC   | `~/.local/bin/anthropos-staging-sync.sh`              | Full sync: per-repo safety patch → `git checkout main; git reset --hard origin/main` → re-apply skip-worktree → `docker compose build` + `up -d` for each changed service → Playwright smoke test. ~10-60 min wall. |

Both run as systemd **user** units. `loginctl enable-linger $USER` is required so they fire even when no SSH session is attached.

### Why 06:00 UTC?

Heaviest service: full `next-web-app` rebuild takes ~8 minutes. 06:00 UTC = 07:00 / 08:00 CET/CEST = before Stefano starts his day. Builds + smoke complete well before standup.

### Why hourly drift?

Cheap (~10s wall, network-only). Surfaces "your staging is N commits behind" in the morning, and "someone left a feature branch" before the force-reset destroys it. Banner / motd reads from `drift.summary`.

---

## Force-main mechanics per repo

Phase 1 of `anthropos-staging-sync.sh` is the force-main loop. Per repo:

1. `git fetch --prune`. If this fails, the repo is left untouched and the error is logged.
2. **Capture skip-worktree marks** before touching the worktree (the [skip-worktree pattern](#skip-worktree-handling) hides long-lived staging-only patches from upstream PRs).
3. **Save a safety patch** to `~/.local/state/anthropos-staging-sync/safety/<repo>-<TS>.{patch,head}` — the working-tree diff (with skip-worktree files temporarily unmasked so their changes are captured too) and the previous HEAD sha. Recoverable via `git reset --hard <head>; git apply <patch>`. 30-day retention.
4. If current branch ≠ default, **unmark skip-worktree** (so the next checkout can move tracked files) and `git checkout -B <default> origin/<default>`.
5. `git reset --hard origin/<default>`. This discards any uncommitted local work — that's **intentional**; the safety patch from step 3 is the recovery surface.
6. **Re-apply skip-worktree marks** for every captured file that still exists upstream. Files that vanished upstream (e.g., a removed Dockerfile or moved config) are logged but not re-marked.
7. Default branch is auto-detected via `git symbolic-ref refs/remotes/origin/HEAD` — works for repos that use `master` instead of `main`.

Untracked staging-local files (`vendor-colony/`, `apps/web/.env.production`, `clerk-fetch-fix.js`, etc.) listed in `.git/info/exclude` survive `reset --hard` naturally — git doesn't touch ignored or unknown files.

---

## Safety patches

Every sync run, before touching any repo, writes:

```
~/.local/state/anthropos-staging-sync/safety/
├── app-20260514T080621Z.patch       # diff of working tree pre-reset
├── app-20260514T080621Z.head        # previous HEAD sha
├── next-web-app-20260514T080621Z.patch
├── next-web-app-20260514T080621Z.head
└── …
```

Files pruned at 30 days. Recovery sequence after an accidentally-clobbered WIP:

```bash
cd ~/<repo>
LATEST=$(ls -t ~/.local/state/anthropos-staging-sync/safety/<repo>-*.patch | head -1)
LATEST_HEAD=${LATEST%.patch}.head
git checkout -B recovered $(cat $LATEST_HEAD)
git apply $LATEST
git status                            # inspect what you got back
```

In practice this almost never fires — the rule is "no WIP on staging clones". The patches exist to cover the "oh shit I had WIP" case once the rule is misremembered, not as a sanctioned workflow.

---

## Repo scope

The 16 repos the routine covers (same on every staging host):

**Service repos (rebuild on change):** `app`, `next-web-app`, `cms`, `skiller`, `skillpath`, `jobsimulation`, `storage`, `sentinel`, `roadrunner`, `messenger`, `customerio-sync`, `studio-desk`, `graphql-wundergraph`.

**Plain repos (no docker rebuild):** `rosetta`, `anthropos-knowledge-base`, `ant-singularity`.

**Excluded by design:**

- `platform/` — runtime config, not a service. Sync-managed manually because the staging-local `docker-compose.yml` and `.env` patches are too host-specific to force-reset safely (the skip-worktree handling does protect them, but the maintainer's call was to leave `platform` off the force-reset path; revisit if staging-local platform divergence becomes painful).
- `ant-academy/` — historically optional, not part of the prod stack.
- `colony/` — vendored snapshot (no `.git`), not a git checkout.
- `skill-path-builder/` — not deployed on staging.

---

## Skip-worktree handling

The `skip-worktree` pattern lets the docker stack read staging-only patches from disk while keeping them invisible to git (so agent commits stay clean). Service clones (`app`, `cms`, `skiller`, `skillpath`, `jobsimulation`, `storage`, `sentinel`, `messenger`, `next-web-app`, `platform`) carry these — see [`staging-bringup.md` Quirk #19](./staging-bringup.md#bringup-quirks-consolidated-as-a-procedural-narrative).

### Apply once per staging clone (idempotent)

```bash
cd ~/<repo>

# Tracked-and-modified files → skip-worktree (git pretends they're unchanged)
git diff --name-only | while read f; do
  git update-index --skip-worktree -- "$f"
done

# Untracked dirs/files → local-only ignore (.git/info/exclude is per-repo, never committed)
echo 'vendor-colony/' >> .git/info/exclude
echo 'apps/web/.env.production' >> .git/info/exclude
```

After: `git status` shows only what the agent actually changed; `git add .` stages just that; `git diff main` is empty until the agent edits something real.

### Typical mark inventory per repo

(Varies by 1-2 entries across hosts due to `.agentspace/.gitkeep` differences.)

| Repo                       | Skip-worktree files                                                                                                  |
| -------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `app`                      | `Dockerfile.dev`, `go.mod`, `go.sum`, `internal/cors/cors.go`, `internal/web/backend/graphql/graph/handler.go`       |
| `cms`                      | `Dockerfile.dev`, `go.mod`                                                                                           |
| `skiller`, `skillpath`, `jobsimulation`, `storage`, `sentinel`, `messenger` | `Dockerfile.dev`, `go.mod` (+ `go.sum` on some)                                          |
| `next-web-app`             | `Dockerfile.dev`                                                                                                     |
| `platform`                 | `Makefile`, `docker-compose.yml`                                                                                     |

### Force-reset interaction (the dance)

The sync's Phase 1 unmarks `skip-worktree` transiently so it can `checkout -B main; reset --hard origin/main`, then re-applies marks for every file that still exists upstream. Files that vanished upstream (e.g. a Dockerfile renamed or a removed module) are logged.

### Caveat: `git pull` refuses on skip-worktree'd files

If you ever do `git pull` directly (don't — the sync handles this), it will refuse when upstream modifies a skip-worktree'd file: `your local changes would be overwritten`. Recover:

```bash
git update-index --no-skip-worktree <file>
git stash
git pull --rebase
git stash pop
git update-index --skip-worktree <file>
```

Rare in practice — patched files don't change often upstream.

---

## Docker rebuild

After Phase 1, only services whose source repo SHA actually moved get rebuilt. Mapping:

| Repo                | Docker compose service |
| ------------------- | ---------------------- |
| `app`               | `backend`              |
| `next-web-app`      | `next-web-app`         |
| `cms`               | `cms`                  |
| `skiller`           | `skiller`              |
| `skillpath`         | `skillpath`            |
| `jobsimulation`     | `jobsimulation`        |
| `storage`           | `storage`              |
| `sentinel`          | `sentinel`             |
| `roadrunner`        | `roadrunner`           |
| `messenger`         | `messenger`            |
| `customerio-sync`   | `customerio-sync`      |
| `studio-desk`       | `studio-desk`          |
| `graphql-wundergraph` | `graphql`            |

Builds run **serially** (1-2 builds in parallel exhaust RAM on a 16 GB box). Build failures are logged but don't abort the rest of the run — they show up in `errors[]` of `last.json`.

**Never rebuilt:** `postgresql`, `redis`, `gotenberg` (vendored images, no local source).

### Atlas migrations are NOT run by sync

The daily sync pulls new source and rebuilds containers, but it does **not** run `atlas migrate apply`. The Go services boot fine against an out-of-date schema; the breakage only surfaces the first time code paths reach a missing table or column (e.g., `ask_conversations does not exist` for Talk to Data, `skill_translations does not exist` for the skiller subgraph). On 2026-05-14 both Ithaca and Calypso had 6–11 pending migrations sitting unapplied since the initial dump restore, undetected for weeks because the smoke test exercises Clerk + `/home` only.

**This is an operator responsibility, not a sync-routine job.** Reasons sync doesn't run Atlas itself:

- `atlas` is not preinstalled on a fresh staging box; the sync routine can't fail soft on a missing binary across the fleet without losing more than it gains.
- Some pending migrations (out-of-order files, baseline mismatches) need interactive `migrate set` / `--exec-order` decisions — wrong choice truncates `atlas_schema_revisions` and is a non-trivial recovery.
- Migrations occasionally fail with destructive intent (rename + drop in two passes), and silently applying them at 06:00 UTC against prod-shape data is a worse failure mode than "Talk to Data 500s until the operator runs Atlas."

**What to do periodically** (≥weekly, or whenever a new feature lands on `main`):

```bash
# Install atlas once (idempotent)
command -v atlas >/dev/null || curl -sSf https://atlasgo.sh | sh

# Check + apply per service. Schemas per service in
# staging-bringup.md § 4.5.
for svc_schema in "app:public" "skiller:skiller" "jobsimulation:jobsimulation" "cms:cms" "skillpath:skillpath"; do
  svc="${svc_schema%%:*}"; schema="${svc_schema##*:}"
  echo "=== $svc → $schema ==="
  (cd ~/$svc && atlas migrate status --env local \
    --url "postgresql://postgres@localhost:5432/postgres?sslmode=disable&search_path=$schema") || true
done

# Then apply where needed:
cd ~/<service>
atlas migrate apply --env local \
  --url "postgresql://postgres@localhost:5432/postgres?sslmode=disable&search_path=<schema>"

# Restart any service whose schema moved
cd ~/platform && docker compose restart <service>
```

Full procedure incl. the out-of-order / baseline gotcha lives in [`staging-bringup.md § 4.5`](./staging-bringup.md#45-apply-pending-atlas-migrations).

**Open question:** whether to fold this into the daily sync once Atlas is preinstalled fleet-wide. Stefano's call as of 2026-05-14 is "keep it manual" — the explicit-pause-before-DDL property is more valuable than the operational tax of remembering. Revisit if the fleet grows beyond 3-4 hosts.

---

## Smoke test

Real-user login flow via Playwright (`~/.local/bin/anthropos-staging-smoke.js`):

- Email/password from [`staging-clerk.md` § Shared test login](./staging-clerk.md#shared-test-login) (overridable via `SMOKE_EMAIL` / `SMOKE_PASSWORD` env).
- 90-120s timeout on `/home` rendering (first-load is slow in dev — known issue, not a regression).
- Screenshot to `/tmp/post_update_home.png`.
- Pass if dashboard greeting appears OR the body has meaningful content after redirect.

Fail does **not** roll back the sync — it just gets logged as `smoke:fail` in `last.json` AND writes `~/.local/state/anthropos-staging-sync/last.alert` with the timestamp + log path. The alert file is *cleared* on the next successful smoke. Read it from a banner or login motd to surface staging health.

---

## State files

All under `~/.local/state/anthropos-staging-sync/`:

| File                  | Cadence            | Purpose                                                                                            |
| --------------------- | ------------------ | -------------------------------------------------------------------------------------------------- |
| `<timestamp>.log`     | per run            | Full transcript of a sync run. 30-day retention (auto-pruned).                                     |
| `last.json`           | per sync run       | Final state: per-repo branch_before/before/after SHAs, action description, errors, smoke result.   |
| `last.summary`        | per sync run       | One-liner (timestamp, changed count, error count, smoke result, host).                             |
| `last.alert`          | only on smoke FAIL | Cleared on next smoke PASS. Banner / motd source.                                                  |
| `safety/`             | per-repo per-run   | Pre-reset working-tree diffs (`.patch`) + previous HEAD (`.head`). 30-day retention.               |
| `drift.json`          | per drift check    | Per-repo behind-count + branch + SHAs + flag (`ok`, `behind`, `warn:on-feature-branch`).           |
| `drift.summary`       | per drift check    | One-liner of total drift + WARN count.                                                             |
| `drift.log`           | append per check   | Drift check transcript (auto-truncated at 1MB).                                                    |
| `drift.lock`          | runtime            | flock to prevent overlapping drift runs.                                                           |

---

## Manual run

Sync everything right now:

```bash
~/.local/bin/anthropos-staging-sync.sh
# or via systemd (status to journalctl):
systemctl --user start anthropos-staging-sync.service
journalctl --user -u anthropos-staging-sync.service -f
```

Just a drift check:

```bash
~/.local/bin/anthropos-staging-drift.sh
cat ~/.local/state/anthropos-staging-sync/drift.summary
```

Read the last sync result:

```bash
jq . ~/.local/state/anthropos-staging-sync/last.json | less
```

---

## Disable / re-enable

```bash
# Stop both timers (sync stays installed but won't fire)
systemctl --user disable --now anthropos-staging-sync.timer anthropos-staging-drift.timer

# Re-enable
systemctl --user enable --now anthropos-staging-sync.timer anthropos-staging-drift.timer
```

The older [`anthropos-pull.sh`](#related) (SSH-login background pull) is **unaffected** — it keeps running on every login. The two coexist; this one does the heavy lifting, the old one keeps clean main branches fresh on every login.

---

## Bootstrap a new staging from an existing one

To install the sync routine on `wip-<initials>` after bringup:

```bash
# From your laptop, with both hosts on the tailnet
scp -r ithaca:~/.local/bin/anthropos-staging-{sync,drift,smoke}.{sh,js} \
       wip-<initials>:~/.local/bin/
scp -r ithaca:~/.config/systemd/user/anthropos-staging-{sync,drift}.{service,timer} \
       wip-<initials>:~/.config/systemd/user/

ssh wip-<initials> '
  sudo loginctl enable-linger $USER
  systemctl --user daemon-reload
  systemctl --user enable --now anthropos-staging-sync.timer anthropos-staging-drift.timer
  ~/.local/bin/anthropos-staging-drift.sh
'
```

Then ssh in and verify `cat ~/.local/state/anthropos-staging-sync/drift.summary` shows green.

---

## Related

- `~/.local/bin/anthropos-pull.sh` — the older, lighter SSH-login background pull. Runs on every SSH login, fetches all repos, ff-only-pulls clean main checkouts. Coexists with this routine.
- [`staging-bringup.md`](./staging-bringup.md) — fresh-VM onboarding.
- [`staging-clerk.md`](./staging-clerk.md) — dev Clerk app + shared test login.
- [`staging_from_dump.md`](./staging_from_dump.md) — original verbose engineer-rebind reference.
- Ant-singularity catalog: [`auto-anthropos-staging-dev-loop.md`](https://github.com/stefano-anthropos/ant-singularity/blob/main/knowledge/singularity-catalog/auto-anthropos-staging-dev-loop.md) — org-level workflow framing.
