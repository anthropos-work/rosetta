# M235 — Spec notes

_(technical details / thresholds accumulate here during build)_

## Pre-flight audits — iter-01 (bootstrap tok)

- **KB-fidelity (Phase 0b):** **YELLOW** — report `kb-fidelity-audit.md`. No blind areas; every M231-M234
  topic PAIRED + code anchors resolve in the rext clone. One finding **KB-1** (stale code-comment): the
  `stack-seeding/contentsession/fixture/content-sessions.yaml` header still describes the SUPERSEDED
  synthesize-first "provably PII-free" posture; the authoritative corpus doc (`session-clone-spec.md`) correctly
  documents copy-real+scrub. Routed to the first fixture-extension tik (it edits that file anyway). Proceed.

## Topic → doc → code triples (audit fast-start)

| Topic | Knowledge doc | Code (rext clone) |
|---|---|---|
| Result routes / prove-by-render (M231) | `corpus/ops/demo/content-stories-routes.md` | `stack-seeding/contentsession/` |
| Session-clone seeder (M232) | `corpus/ops/demo/session-clone-spec.md` | `seeders/content_stories{,_write,_modality}.go`, `cmd/content-capture`, `scrub/` |
| Content manifest + honesty gate (M233) | `corpus/ops/demo/content-stories-spec.md` §1–§6 | `seeders/content_manifest.go`, `cmd/stackseed` `--content-export` |
| Cockpit tab + content-player seats (M234) | `corpus/ops/demo/content-stories-spec.md` §7 | `demo-stack/cockpit.py`, `seeders/roster.go`, `storyPopulationNames` |
| Playthroughs (function) | `corpus/ops/demo/playthroughs.md` | `playthroughs/{manifest,e2e,seed,report}`, `seed/pt-world.seed.yaml` |
| Coverage sweep (presence) | `corpus/ops/demo/coverage-protocol.md` | `stack-verify/e2e/` |
| Academy demo-fill (M230 c/f) | `corpus/services/ant-academy.md`, `frontend-tier.md` | rext tag `playbill-m230-academy-fs-published`, `app/cmd/academy-seed` |
