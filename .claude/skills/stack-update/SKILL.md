---
name: stack-update
description: Sync a stack's code, dependencies, and database schemas with the latest — pull platform repos, install deps, apply migrations, rebuild. Works on the main dev stack or a named dev-N. Use when asked to update / sync a stack.
argument-hint: [dev-N] [scenario: 'daily' | 'weekly' | 'full']
---

# Stack Update — sync a stack's code, deps, and schemas

Updates a stack by following `corpus/ops/update_guide.md` with systematic verification — pull platform repos,
install dependencies, apply migrations, rebuild + restart, verify health. Operates on the **main dev stack**
(N=0) by default, or a named **`dev-N`** (its platform dir + its pinned tooling tag). (Formerly
`/update-platform`, now stack-targeted.) Source of truth:
[`corpus/ops/update_guide.md`](../../../corpus/ops/update_guide.md).

> **Demo stacks aren't updated in place** — a `demo-N` is disposable and clones each platform repo at a pinned
> **release tag** for reproducibility. To "update" a demo, tear it down (`/demo-down`) and bring it up again
> (`/demo-up`) at the desired refs. This skill targets the **dev** side, whose repos track `main`.

## Your mission

1. **Read the guide** — `corpus/ops/update_guide.md` is the source of truth.
2. **Apply UPDATE STEP principles** to every step:

   | Principle | Action |
   |-----------|--------|
   | Check Current State | Verify what needs updating before changing anything |
   | Pull Before Build | Always fetch latest code before rebuilding |
   | Handle Conflicts | If git conflicts occur, resolve before proceeding |
   | Verify After Update | Confirm services still work after updates |

3. **Track progress** via TodoWrite: services stopped (`make down`) → repos updated (`make pull`) → deps
   installed (frontend `pnpm install`) → migrations applied (`make migrate`) → rebuilt + started (`make up`)
   → verified healthy (`make ps` + health checks).

## Confirmation policy

**Proceed without confirmation:** checking git status / service state (`make status`, `make ps`), health
checks. **Ask first:** stopping services (`make down`), pulling (`make pull`), migrations (`make migrate`),
rebuilds (`make up`), destructive ops (`make reset-db`).

## Target resolution

- **No target / `dev-0` / the main dev stack** — operate in `stack-dev/platform` against the main `anthropos`
  project; `make pull`/`migrate`/`up` as usual.
- **`dev-N` (N ≥ 1)** — operate in that stack's platform dir with its offset project (`-p dev-N`). A `dev-N`
  built by `/dev-up` consumes its tooling at a pinned `rosetta-extensions` tag; **updating that tooling is a
  tag bump, not an in-place edit** (see the note below).

## Updating a stack's tooling (NOT editing scripts in place)

Each gitignored `stack-*/` dir spans one full local stack — its platform service repos **plus** its own clone
of `rosetta-extensions` consumed at a pinned tag (`stack-<role>/rosetta-extensions @ <tag>`). To update that
tooling, **bump the stack's pinned `rosetta-extensions` tag** — re-clone or checkout at the new tag — rather
than hand-editing scripts inside the stack dir. New tooling work happens first in the AUTHORING copy at
`.agentspace/rosetta-extensions/` (build + TEST, then commit + TAG); stacks only ever consume tooling at a
pinned tag.

## Error handling

1. Don't skip errors — resolve first.
2. Check logs: `cd stack-dev/platform && make logs S=[service]`.
3. Write an ops-report (`stack-dev/ops-reports/op_YYYYMMDD_HHMMSS_update_<topic>.md`: Issue / Context /
   Resolution / Suggested doc update) — these feed `/update-knowledge`.

## Critical rules

- Work in the stack's `stack-*/` workspace only; use `make` for all platform operations.
- Stop services before pulling code; handle git conflicts before continuing; verify health after.
- Follow the guide — don't improvise.

## Related skills

| Skill | Use when |
|-------|----------|
| `/dev-up` | Start a stack after updating |
| `/stack-list` | List live stacks |
| `/update-knowledge` | Process ops-reports into the corpus |

## Additional resources

- For technical reference (update scenarios, troubleshooting), see [reference.md](reference.md).
