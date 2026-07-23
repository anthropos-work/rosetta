# M246 — Decisions

_(Implementation decisions with rationale, D-numbered, recorded during build.)_

## KB-fidelity findings (Phase 0b, 2026-07-23) — verdict YELLOW

- **KB-1** — Seeder write-site anchors ALIGNED. `main.go:97` + `hero_activity.go:180` verified exact; full write-site inventory in spec-notes. In-scope build work.
- **KB-2** — `DEMO_ADVANCE_CLONES=pinned` path ALREADY WIRED (M237, `ensure-clones.sh:206`). M246's delta = author `stack-demo/clones.pin.json`, not wire the path. Scope clarification.
- **KB-3** — Injection comment `gen_injected_override.py:16` STALE (4→3 subgraphs). M246 in-scope fix (comment only).
- **KB-4** — Injection CODE (`INJECTED` dict `:17`, enum `:458`, `exposure_claim_guard.py:124`) still lists skillpath as a live service. Out of the comment-only declared scope → **go/no-go scope-watch** for the bring-up; observed behavior → M247 drift ledger.
- **KB-5** — Corpus asserts skillpath live Tier-1 (~30 files). **Fate-2 confirmed covered by M247** (corpus re-ground) — no M246 deferral, no M246 edit. M246 does not read these as truth.

## D-1 — Seeder re-point: hard-cut to `public`, no dual-schema tolerance (2026-07-23)
**Decision.** Re-point every `skill_path_sessions` write from schema `skillpath` → `public` in one
shot (live seeder code + tests + DNA + comments); do NOT add a runtime dual-schema fallback.
**Rationale.** The open question in overview.md asked dual-schema-tolerant vs hard-cut. Hard-cut wins:
(a) the demo builds from a SINGLE pinned clone set (this milestone bumps it to current `origin/main`,
where skillpath is decommissioned) — there is no mixed-topology demo to tolerate; (b) a dual-schema
seeder would need a live schema probe per write (complexity + a new failure mode) to serve a
transition window the demo never occupies; (c) `rext` is consumed per-stack at a pinned tag, so a
stack still on the old platform pin simply consumes the OLD rext tag — the version skew is handled by
the tag pin, not by seeder branching. Commit `97585f5` (rext).
**Scope kept tight.** Surface NAME `skillpath-sessions`, Go symbols, and the mirror
`public.local_skill_path_sessions` untouched (the mirror was already `public`).
**Empirical dependency.** Correctness of the target (`public.skill_path_sessions` existing) is proven
by the cold `/demo-up` — the seeder COPY to public succeeds only if the consolidated migrations create
it there. That is the go/no-go, not an assumption baked in here.
