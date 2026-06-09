# M20 — Spec Notes

_Technical notes accumulated during build — mechanisms, file paths (with line cites), gotchas, and the concrete shape of each change. Populated by `/developer-kit:build-milestone`. The verified code locations from the design-time research are in the milestone `overview.md` and `.agentspace/demo-up-issue.md`._

## Pre-flight audits — Section 1 (set-dress chaining + cold-start doc)

**KB-fidelity (Phase 0b): GREEN.** Report: `kb-fidelity-audit.md`. Sha at audit: `3ddb277` (pre-flight). Topic→doc→code triples:
- Set-dress chaining → `snapshot-spec.md` §"Dev as a full-fidelity peer (M13)" + `safety.md` §2.5 → `dev-stack/dev-setdress.sh` (reuse), `demo-stack/up-injected.sh` (chain site). ALIGNED.
- Capture-source policy → `snapshot-spec.md` §"The capture-source policy" + `safety.md` §1.4 → `stack-snapshot/source/source.go`, `cmd/stacksnap`. ALIGNED (precedence list, `Available()`, bounded-session SQL, no-offline-file-reader all match byte-for-byte).
- Re-run safety → `idempotency.md` (M17) → replay TRUNCATE-then-reload + idempotent seed COPY. ALIGNED.
- Cold-start fresh-box workflow → BLIND-AREA → the milestone deliverable `corpus/ops/snapshot-cold-start.md` (overview `Delivers →`). Not a blocker; authored as first work.
- demo preset → `seeding-spec.md` §"The shipped presets" → `stack-seeding/presets/small-200.seed.yaml`. ALIGNED.

## Key code locations (verified design-time + this audit)
- `demo-stack/up-injected.sh:195` — `migrate-demo.sh` call; the set-dress chain goes AFTER it, BEFORE the M18 verify (`:211`).
- `dev-stack/dev-setdress.sh` — the proven M13 pass: `snapshot_step()` (provision-plan check-env → cache-first `stacksnap replay taxonomy,directus`) + `seed_step()` (preset seed). The reuse target.
- `stack-snapshot/source/source.go` — `DefaultPrecedence`, `BoundedSession.SetupSQL()`, `Resolve()`. The cold-start DSN-export path documents this.
- `stack-snapshot/cmd/stacksnap/main.go:152+` — `capture` flagset; `--dsn` required, `--source` optional, NO `--dump`-file flag.
