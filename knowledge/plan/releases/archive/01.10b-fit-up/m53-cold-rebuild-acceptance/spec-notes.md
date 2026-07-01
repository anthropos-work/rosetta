# M53 — spec notes

_Technical notes accumulate here during build (file:line surfaces, rext tag, schema findings)._

## Pre-flight audits — §1 academy F6 (first section)
- **Phase 0b verdict: GREEN.** Report: `m53-cold-rebuild-acceptance/kb-fidelity-audit.md` (2026-07-01).
  All 9 acceptance-bar contract docs ALIGNED with rext authoring HEAD `36d7430`. Two low-severity caveats
  recorded as KB-1 (AB2 replay-from-filled-cache) + KB-2 (AB5 78.4%/199). Academy F6 = the one planned
  new-code surface (menu-link + authenticated session net-new; content ships with clone).

## Topic → doc → code triples (fast-start for future audits)
| Topic | Doc | Code |
|---|---|---|
| bring-up phases / #7 abort | `rosetta_demo.md`, demo-up SKILL/GUIDE | `demo-stack/up-injected.sh` |
| demo-down --purge (M49 #6) | `rosetta_demo.md` | `demo-stack/rosetta-demo:162-171` |
| cold-start capture | `snapshot-cold-start.md` | `stack-snapshot/`, `pg/pg.go:114-134` |
| idempotency | `idempotency.md` | replay/seed/casbin |
| auto-verify | `verification.md` | `stack-verify/live/autoverify.sh`,`verify.sh` |
| coverage | `coverage-protocol.md` | `stack-verify/e2e/run-coverage.sh` |
| AI-readiness | `ai-readiness.md`,`seeding-spec.md` | `internal/workforce/ai_readiness.go`, patch, seeders |
| manifest download | `cockpit-spec.md`,`seed-manifest-spec.md` | `demo-stack/cockpit.py`,`stack-seeding/manifest/` |
| academy | `ant-academy.md`,`frontend-tier.md` | `demo-stack/ant-academy.sh`, `stack-demo/ant-academy/code/` |

## Key file surfaces (build reference)
- **Academy launcher (F6 target):** `demo-stack/ant-academy.sh` — runs anonymous via `BENCHMARK_VISUAL_BYPASS=1`
  + `REQUIRE_ORGANIZATION_MEMBERSHIP=0`; port `3077 + N*10000`; writes gitignored `code/.env.local`. Launched
  default-on at `up-injected.sh:648`.
- **Cockpit deep-link catalog:** `stack-seeding/seeders/cockpit.go:59-77` (`DeepLinkCatalog`, next-web-only — no
  academy entry). Cockpit render: `demo-stack/cockpit.py` (per-hero `[Log in as]`→jump_to via handshake).
- **AI-readiness perf-patch:** `demo-stack/patches/app-aireadiness-snapshot-loadmembers/` applied at
  `up-injected.sh:447-460` (M51, non-fatal).
- **Snapshot cache (AB2):** `.agentspace/snapshots/{taxonomy,directus,sim-embeddings}/<hash>/` — 1.4 GB, filled.
- **Coverage runner:** `stack-verify/e2e/run-coverage.sh <N> {employee|manager}`; gate = `gateMet` on
  frontier-exhausted crawl.
- **rext authoring HEAD:** `36d7430` (past `fit-up-m52`); tags `fit-up-m47..m52` present; `v1.10.1` to roll in §2.
- **rext.tag currently:** `fit-up-m51` → bump to `v1.10.1` in §6.
