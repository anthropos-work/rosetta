---
title: "KB Fidelity Audit — M246 re-sync & re-point"
date: 2026-07-23
scope: milestone:M246
invoked-by: build-milestone
---

## Verdict
YELLOW

Proceed with tracking. No blind areas; every code anchor the milestone cites is real and
accurate; every KB-dependency doc exists. The known-stale skillpath-is-live corpus claims are
**Fate-2 owned by M247** (the corpus re-ground) — tracked, not M246 blockers, because M246's own
implementation re-points AWAY from `skillpath` to `public` driven by the platform-consolidation
research + the empirical `/demo-up` prove, NOT by any corpus prose it reads as truth.

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Seeder re-point `skillpath.*→public.*` | `corpus/services/skiller.md` (redirect pattern); NOT corpus-anchored as public yet (M247) | rext `stack-seeding/{cmd/stackseed/main.go, seeders/hero_activity.go, skillpath_sessions.go, content_nonsim.go}` + `dna/data-dna.json` | PAIRED |
| Demo clone pins / `DEMO_ADVANCE_CLONES=pinned` | `corpus/ops/update_guide.md`, `corpus/ops/rosetta_demo.md` | rext `demo-stack/ensure-clones.sh:172-355` (mechanism); `stack-demo/clones.pin.json` (to author) | PAIRED |
| Injection subgraph set (3 subgraphs) | `corpus/ops/rosetta_demo.md` (demo lifecycle) | rext `stack-injection/gen_injected_override.py:15-17` | PAIRED |
| Cold `/demo-up` bring-up + demopatch escape | `corpus/ops/rosetta_demo.md`, `corpus/ops/demo/demopatch-spec.md` | rext `demo-stack/up-injected.sh`, `demopatch` | PAIRED |

## Fidelity Findings

1. **Seeder write-site anchors — ALIGNED.** overview.md cites `cmd/stackseed/main.go:97` (reset
   TRUNCATE entry `"skillpath.skill_path_sessions"`) and `seeders/hero_activity.go:180`
   (`CopyRowsIdempotent(ctx, "skillpath", "skill_path_sessions", …)`). Both line anchors verified
   exact. Full write-site set enumerated in spec-notes triples below.

2. **`DEMO_ADVANCE_CLONES=pinned` path — ALIGNED, already wired (M237).** `ensure-clones.sh:206`
   already implements the `pinned` advance mode reading `$DEMO/clones.pin.json`; `:315` implements
   the pin-drift freshness check. overview.md's "wire the advance path" is therefore *already done* —
   M246's real delta for this section is **authoring `stack-demo/clones.pin.json`** (currently
   absent) with pins to current `origin/main`. Scope clarification, not a blocker.

3. **Injection comment `gen_injected_override.py:16` — STALE (in-scope fix).** The comment (lines
   15-16) reads "the injected/subgraph services are backend/app, cms, jobsimulation, skillpath" — 4
   subgraphs incl. skillpath. Consolidated reality = **3 subgraphs** (skillpath decommissioned into
   app). This is M246's declared in-scope fix. Fix owner: update comment.

4. **Injection CODE still lists skillpath as a live service — STALE, scope-watch (NOT the comment).**
   Beyond the comment: `gen_injected_override.py:17` `INJECTED = {…, "skillpath": "skillpath"}`,
   `:458` service enumeration, and `exposure_claim_guard.py:124` (`"skillpath": {"ports":[…8095…]}`)
   still treat skillpath as an injectable/exposed subgraph. M246's declared scope fixes **only the
   comment** (line 16). Whether the bring-up needs the *code* entries handled (a decommissioned
   skillpath has no compose service to inject → likely a benign no-op, but possibly a hard error) is a
   **go/no-go surface to watch** during the cold `/demo-up` prove — transcribe the observed behavior
   into the M247 drift ledger.

5. **Corpus asserts skillpath live Tier-1 (~30 files) — STALE, Fate-2 COVERED BY M247.** The corpus
   docs (`rosetta_demo.md`, `update_guide.md`, `skiller.md`, service taxonomy, etc.) still describe
   skillpath as a live standalone Tier-1 service / 4th subgraph. M247 (corpus re-ground) explicitly
   owns this reconciliation (overview.md §Out: "The corpus doc reconciliation (M247)"). Per the
   three-fate rule this is **confirmed covered by M247** — no new M246 deferral, no M246 edit. M246
   does not read these claims as truth.

## Completeness Gaps

None load-bearing for M246. The `public.skill_path_sessions` existence claim (the seeder's re-point
target) is not yet corpus-anchored — but that is the exact fact M246 exists to **prove** empirically
via the cold `/demo-up` bring-up, and M247 will anchor it in the corpus afterward. Not a gap that
blocks M246.

## Applied Fixes

None inline. Finding 3 (the comment) is M246's own in-scope build work (not an audit-time doc fix).
Findings 4-5 are tracked below / owned by M247.

## Open Items (require user decision)

None. Finding 4 (injection code skillpath entries) is a scope-watch to be resolved empirically by the
bring-up + transcribed to the M247 ledger — not a pre-build user decision.

## Gate Result

**YELLOW — proceed with tracking.** Findings recorded as KB-1..KB-5 in `decisions.md`. Finding 3 is
in-scope build work; Finding 4 is a go/no-go scope-watch for the bring-up; Finding 5 is Fate-2 owned
by M247.
