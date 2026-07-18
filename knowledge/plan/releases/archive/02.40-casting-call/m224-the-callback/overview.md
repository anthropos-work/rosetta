---
milestone: M224
slug: the-callback
version: v2.4 "casting call"
milestone_shape: iterative
status: archived
created: 2026-07-15
depends_on: M222, M223
exit_gate: "On a COLD reset-to-seed, the manager hero's comparison surface (/enterprise/activity-dashboard → AI-Simulations → [simId]) renders ≥40 comparable candidate rows per EACH of the 5 shared sims, with a realistic non-degenerate score distribution, 0 junk skills/roles/names (closure green), 0 prod-eject escapes, over ≥3 consecutive cold runs. Latency is REPORTED, not gated, here (gated at M226, per the v2.3 D-DESIGN-1 lineage)."
iteration_protocol_ref: corpus/ops/demo/coverage-protocol.md (+ corpus/ops/verification.md — the live-browser measure → attribute → fix → re-measure loop, here driving a hiring-comparison RENDER gate)
delivers: click→[Log in as] the recruiter hero lands on the comparison surface with ~45 comparable non-junk rows per each of 5 positions; the org reads as hiring; a candidate hero renders a usable assessed /profile. Plus the render-path + hiring-vantage section into corpus/services/hiring.md + cockpit-spec.md
---

# M224 — The callback

## Goal
Click **[Log in as]** the recruiter hero → land on `/enterprise/activity-dashboard` and see, for **each** of the 5
positions, **~45 rankable, comparable, non-junk** candidate rows; the org reads as hiring; and a candidate hero
renders a usable assessed `/profile`.

## Exit gate (measurable)
On a **cold reset-to-seed**, the manager hero's comparison surface renders **≥40 comparable candidate rows per each
of the 5 sims** with a **realistic non-degenerate score distribution**, **0 junk skills/roles/names** (closure
green), **0 prod-eject escapes**, over **≥3 consecutive cold runs**. **Latency is REPORTED, not gated, here** —
gated at M226 (per the v2.3 D-DESIGN-1 access-vs-render lineage).

## Why iterative (not section)
The **render-risk heart.** Whether ~45 comparable, non-junk candidate rows actually paint on the manager's
comparison surface is **measurement-driven** — the exact read-path is a blind area (BA-1), and the M219 lesson is
that **a render gate silently bypasses the seed** (`CycleID==nil→buildLiveResponse`; the PostHog `flag-gate`; junk
skills from a dry pool). You **seed → render → attribute the gap → fix (data-only seeding OR a sha-pinned demo-patch
OR Clerkenstein wiring) → re-render.** A fixed `In:` list would be speculative.

## Iteration protocol
`corpus/ops/demo/coverage-protocol.md` (the live-browser measure → attribute → fix → re-measure loop) +
`corpus/ops/verification.md`. Reuse the M37 cockpit seat-switch for hero login + the M42 e2e foundation (never
forked).

## Scope

### In
1. **The hero trio (cockpit seats, login-only — the v2.3.2 change: one [Log in as] CTA per hero, no academy
   button):**
   - **1 recruiter/manager** hero (`vantage: manager` → `jump_to` the comparison surface).
   - **2 candidate exemplars** on the *same* position, showing two funnel states: **one assigned AND assessed** on a
     hiring simulation, **one only assigned** (not yet taken) → `/profile` | `/profile/skills`. The candidate heroes
     render usable assessed `/profile` surfaces.
2. **DeepLinkCatalog entries** for the recruiter surfaces (`/enterprise/activity-dashboard`, per-`[simId]`
   comparison) — none exist today — with proper `[Log in as]` labels + a `NeedsID` entry if the jump targets a
   specific `[simId]`.
3. **Extend Clerkenstein FAPI/BAPI to emit org `public_metadata.isHiring=true`** (D-DESIGN-1: a **genuine** hiring
   org — the org and candidates read as hiring end-to-end). **Tooling-owned, NOT a platform patch.** Today the
   roster emits org `public_metadata.eid` only.
4. **Drive the render loop to green** against the exit gate.

### Out
- The coverage-sweep gate + playthrough (M225) · the live billion proof (M226).

## Depends on
**M222, M223.**

## KB dependencies
- `corpus/ops/demo/coverage-protocol.md` · `corpus/ops/demo/cockpit-spec.md` · `corpus/services/clerkenstein.md` ·
  `corpus/ops/demo/stories-spec.md` · `corpus/services/hiring.md` (M222) · `corpus/ops/demo/demopatch-spec.md`

## Delivers → knowledge/corpus
The render-path + hiring-vantage section into `corpus/services/hiring.md` + `corpus/ops/demo/cockpit-spec.md`.

## Demo-patch? — LIKELY (the release's D-DESIGN-2 order-of-preference call)
1. **`next-web-hiring-flag-gate`** — if a comparison surface hard-gates on `isHiringOrg` the way M219's member
   surface gated on an `undefined` PostHog flag. **Prefer the Clerkenstein `isHiring` wiring FIRST** (tooling, no
   patch); only if a surface *still* won't mount does it route to a sha-pinned demo-patch (the M219 precedent).
2. **A perf demo-patch** for the whole-org-hydration of the 45×5 compare table (the AI-readiness `loadMembers`
   180 s→19 ms precedent) — but **latency is not this milestone's gate**, so this may defer to M226.

**A platform-repo edit is never in bounds; an un-patchable surface escalates** (the 4-state
`unimplementable-without-platform-edit`).

## Alignment guard (BLOCKING for any `clerkenstein/clerk-frontend/` change)
Extending Clerkenstein FAPI/BAPI to emit `public_metadata.isHiring` touches the fake FAPI/BAPI. **Any change to
`clerkenstein/clerk-frontend/` MUST carry an explicit `/align-run` step** (CI is inert). The identity genes are
`critical` (100%, no partial credit); a new `isHiring` field must not perturb the existing decoded-identity genes.

## Risks carried
- **R1 (blocks-milestone)** — a platform HIRING render path bypasses the seed → "45 seeded candidates" shows as 0
  rows. Mitigation: prefer Clerkenstein wiring; a genuine render gate → a sha-pinned demo-patch; un-patchable →
  escalate.
- **R3 (degrades-quality)** — junk-or-identical assessments. Mitigation: closure green + a realistic distribution
  (seeded at M223, swept at M225).
- **R4 (latency)** — the 45×5 compare-table hydration. Measured + REPORTED here; a perf demo-patch may pin here but
  is **gated** at M226.
