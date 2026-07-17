---
milestone: M227
slug: the-notes
version: v2.4 "casting call"
milestone_shape: section
status: archived
created: 2026-07-17
last_updated: 2026-07-17
depends_on: M222, M223, M224, M225, M226
delivers: the believability corrections + retuned gate thresholds into corpus/services/hiring.md (+ the seeding/coverage/playthrough specs)
---

# M227 — The notes

> **Believability-hardening backfill** — 4 demo-data realism defects surfaced by LIVE feedback AFTER M226 shipped
> the working hiring demo (Meridian Talent, proven live on `billion`). The demo works, but doesn't fully *read* as
> real. **All fixes in the rext seed/content tooling — 0 platform-repo edits.** v2.4 stays OPEN until the release
> reads right; the user runs `/developer-kit:close-release` on explicit sign-off. The billion re-prove is **M228**.

## Goal
The hiring demo reads as **believable** — the 4 fixes land + the compare gate retunes to the realistic per-position
count, proven GREEN **locally** (cold reset-to-seed on an offset-port demo, asserted from this Mac; never billion).

## Shape — section
4 enumerable seed/content fixes + a gate retune + a local re-prove. Reuses the M223 funnel + M225 coverage/playthrough
machinery (never forked).

## Sections (the In: list)

### 1. Hiring-only content
The hiring org (Meridian Talent) surfaces ONLY `SIMULATION_TYPE_HIRING` sims; the training/assessment sims are
scoped OUT of its surfaces. **Root cause (traced):** the recruiter's AI-Simulations Results list
(`InsightsByJobSimulations`, `app/internal/organization/intelligence.go:1472`) reads `public.local_jobsimulation_sessions`
grouped by sim and shows every distinct sim that has a mirror row for an org member. The mirror rows for the hiring
org are written by TWO seeders: `HiringFunnelSeeder` (HIRING-typed only ✓) **and `PersonaSeeder`** — which writes the
candidate heroes Cara/Cody a workforce-style verified-skill chain against **generic (non-hiring) sims** (`refs.sims`),
whose mirror rows surface those non-hiring sims in the recruiter list. **Fix:** the generic workforce activity seeders
skip hiring orgs so a hiring org's ENTIRE simulation footprint is the HiringFunnel (hiring-only). The candidate
heroes' usable profile comes from the HiringFunnel `/home` assignment + scored positions, not the (unreachable,
admin-gated `/profile`) verified-skill chain — so skipping it is complete, not a partial.

### 2. External candidate emails
Candidates are outside applicants → private/external domains (gmail.com / outlook.com / proton.me / icloud.com / …),
NOT `@meridian-talent.com`; only employees (recruiters/admins) keep the org domain. **Root cause:** `users.go` keys
the email domain on the STORY (`storyEmailDomainFor(st)`), role-blind. **Fix:** key the domain on ROLE — `candidate`
→ a deterministic external domain; else the org domain. Keep the Clerkenstein roster email == the seeded user email
== the login (consistency), incl. the candidate heroes Cara/Cody (whose `login:` in the preset must match).

### 3. 1 sim per candidate + gate retune
Each candidate does exactly **1** of the 5 hiring positions (the role applied for), not all 5 → **~8 candidates per
position** (from ~40 assessed, distributed evenly across the 5). **Fix:** `HiringFunnelSeeder.seedHiringOrgFunnel`
assigns each assessed candidate ONE position by a deterministic even split; keep the mirror pair correct + closure
green (the M219 mirror-table trap). Then **retune the compare gate `≥40 → the realistic floor` EVERYWHERE it is
asserted**: M224 `GATE-DECISION D1` (decisions.md), the M226 (+ new M228) exit_gate condition (2), the M225 coverage
manifest (`reservedHiringSimRefs`/floor) + hiring playthrough, `corpus/services/hiring.md`, and the render-probe
`RENDER_GATE_FLOOR` default. Confirm the seeded per-position count first, then set the floor with a small safety
margin below it.

### 4. Gender-consistent avatars
Every user's avatar photo matches the name's inferred gender, across ALL orgs (not just hiring). **Root cause:**
`photoAvatarDataURI(uid)` picks 1 of 12 bundled faces by `hash(uid)` — gender-blind. The 12 faces are F={01,03,05,07,10}
M={02,04,06,08,09,11,12}. **Fix:** infer gender from the (first) name and pick the face from the gender-matching
subset — for generated fill members AND the fixed heroes. Deterministic, $0, offline-safe (no re-generation of the
LLM cache).

### 5. Local re-prove
Bring up a FRESH LOCAL demo (cached images → fast; offset ports; NEVER billion) consuming the new M227 rext tag,
cold reset-to-seed, and confirm on the corrected data (asserted from this Mac against the LOCAL offset ports):
- the recruiter comparison renders **≥ the retuned floor per each of the 5 positions** (each candidate on 1 sim);
- hiring-only sims in the org;
- candidates show external emails;
- avatars match names;
- the hiring coverage sweep + the hiring playthrough GREEN.

## Out
The billion re-prove (M228).

## Depends on
M222–M226 (the shipped hiring demo).

## KB dependencies
`corpus/services/hiring.md` · `corpus/ops/demo/stories-spec.md` · `corpus/ops/seeding-spec.md` ·
`corpus/ops/demo/coverage-protocol.md` · `corpus/ops/demo/playthroughs.md`

## Delivers → knowledge/corpus
The believability corrections + retuned gate thresholds into `corpus/services/hiring.md` (+ the specs).

## Demo-patch?
Pure tooling (seed + content-scope). No platform-render gap. 0 platform-repo edits.

## Hard constraints (carried)
- **ZERO platform-repo edits** — only rosetta `corpus/`/`knowledge/` + the rext tooling.
- **LOCAL demo only, never billion** (billion is M228). One driver.
- rext tooling → the authoring clone `.agentspace/rosetta-extensions` (commit + tag `casting-call-m227-*`, sync
  consumption, bump `.agentspace/rext.tag`); rosetta docs/plan → `m227/the-notes`.
- **M219 mirror-table trap** (fix #3): the score source is `public.local_jobsimulation_sessions` (the mirror pair
  the HiringFunnelSeeder writes) — changing to 1-sim/candidate must keep the mirror pair correct + closure green.
- Do NOT touch the pre-existing unrelated working-tree changes (`.gitignore` + `_email-assets/…`).
