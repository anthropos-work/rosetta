---
title: "KB Fidelity Audit — M25 field-bake"
date: 2026-06-13
scope: milestone:M25
invoked-by: build-milestone
---

## Verdict
YELLOW

No blind areas; no stale load-bearing claim. The one M25-critical flag (`--local-content`) is ALIGNED in
parser + body + threading. Two doc-surface arg-hint drifts fixed inline (KB-1, KB-2). The live runs are
unblocked — and the field-bake itself is the ultimate fidelity test (it proves the documented behaviors by
observation).

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| local-Directus provisioning + exit-0/4 split | `corpus/ops/directus-local.md`, `snapshot-spec.md` | `stack-snapshot/cmd/stacksnap/main.go`, `autoprovision.go` | PAIRED · ALIGNED |
| asset-plane-stays-prod | `directus-local.md`, `snapshot-spec.md` | `stack-injection/gen_injected_override.py`, `stack-core/gen_override.py` | PAIRED · ALIGNED |
| demo default-on / dev opt-in / N=0 guard | `directus-local.md`, `rosetta_demo.md` | `dev-stack/dev-setdress.sh` | PAIRED · ALIGNED |
| structure+rows captured together (M21 fold) | `snapshot-spec.md`, `snapshot-cold-start.md` | `stack-snapshot/directus/directus.go`, `capture/capture.go`, `manifest/manifest.go` | PAIRED · ALIGNED |
| cold-start `--dsn` + AssertPublicOnly + no-MCP-source | `snapshot-cold-start.md` | `stack-snapshot/firewall/firewall.go`, `source/source.go`, `cmd/stacksnap/adapters.go` | PAIRED · ALIGNED |
| new Directus verify probes | `verification.md` | `stack-verify/lib/services.sh`, `lib/readiness.sh` | PAIRED · ALIGNED |
| teardown reclaims directus + registry | `rosetta_demo.md`, `safety.md` | `demo-stack/rosetta-demo` (down), `dev-stack/dev-stack` (down) | PAIRED · ALIGNED |
| skill flags (`--local-content`, env-var toggles) | `dev-up`/`demo-up`/`stack-snapshot` SKILL.md | the CLI parsers | PAIRED · 2 drifts fixed |

## Fidelity Findings

1. **KB-1 — `demo-up` arg-hint false promise.** `argument-hint` advertised `--full` (nonexistent anywhere),
   `--no-ui`, `--no-setdress` as `rosetta-demo` CLI flags. The real interface is env vars on `up-injected.sh`
   (`DEMO_NO_UI` / `DEMO_NO_SETDRESS` / `DEMO_NO_LOCAL_CONTENT`); the SKILL **body** already uses the env-var
   form correctly. Verdict STALE (arg-hint only). **Fixed inline** — arg-hint now states the env-var reality.
2. **KB-2 — `dev-up` arg-hint incomplete.** `--local-content` is in the parser (`dev-stack:88`) and the body,
   but was missing from `argument-hint`. Verdict STALE (incomplete). **Fixed inline** — added `[--local-content]`.
3. **`stack-snapshot` exit-code framing** (agent flagged STALE): the description/body say "replay exits 0 on a
   `--local-content` stack; exit 4 on a stack without `--local-content` (live-from-prod)". The nuance that
   exit 4 also fires on a genuinely unprovisioned schema regardless of mode is true, but for the *normal* path
   a non-local-content stack simply has no provisioned per-stack Directus → directus replay → exit 4. The
   framing is accurate for the documented paths (M24 already converged this). Verdict ALIGNED-enough — and the
   field-bake will *observe* the actual exit codes, settling it empirically. No edit.

## Completeness Gaps
None critical. The five M25 done-bar behaviors all have a documented contract and a code anchor.

## Applied Fixes
- `.claude/skills/demo-up/SKILL.md` — arg-hint rewritten to the env-var reality (drop fictional `--full`).
- `.claude/skills/dev-up/SKILL.md` — arg-hint gains `[--local-content]`.
- `spec-notes.md` — Pre-flight audits block with the topic→doc→code triples.

## Open Items (require user decision)
None.

## Gate Result
YELLOW — proceed. The two arg-hint drifts are fixed; the `stack-snapshot` exit-code framing is accurate for
the documented paths and will be empirically confirmed by the live runs.
