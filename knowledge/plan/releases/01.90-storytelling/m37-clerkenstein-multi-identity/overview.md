---
milestone: M37
slug: clerkenstein-multi-identity
version: v1.9 "storytelling"
milestone_shape: section
status: archived
created: 2026-06-22
last_updated: 2026-06-23
complexity: medium-large
delivers: rosetta-extensions/clerkenstein (a users/orgs registry replacing the single DefaultDemoUser + an active-user selection mechanism + an Alignment DNA for the multi-identity surface) + corpus pointer (rosetta_demo.md / clerk-integration.md)
depends_on: M35
parallel_with: M36
spec_ref: .agentspace/seeding_gaps.md §4f (Clerkenstein single→multi identity, G20); builds on the wip/clerkenstein-browser-login branch
---

# M37 — Clerkenstein multi-identity

## Goal
A demo stack can **switch the active browser identity** among the seeded heroes/orgs — the seat-switch the
cockpit's "login as" requires. Today Clerkenstein resolves every session to **one** `DefaultDemoUser`
(`clerk-frontend/resources.go:21`, `SingleSessionMode: true`); M37 makes it multi-identity.

## Why section
The boundary is clear: JWT minting is **already** fully per-identity (universal HS256 key + the claim shape
`{AuthID, Eid, Org*, OrgRole}`), so the work is a **registry + active-user selection**, not crypto. A WIP
branch (`wip/clerkenstein-browser-login`) already exists to build on. Fixed checklist.

## Scope
**In:**
- **A users/orgs registry** in `clerk-frontend` — replace the single `user DemoUser` field with a map keyed
  by the seeded heroes/orgs (sourced from the stories.yaml identity list M35 defines).
- **An active-user selection mechanism** — how the browser session resolves to a chosen identity. **O11:**
  spike both **token-injection** (mint + inject the cookie/storage clerk-js reads) **and** a **parameterized
  FAPI handshake**, pick the simpler robust one early.
- **An Alignment DNA** for the new multi-identity surface (author via `/align-dna`; score via `/align-run`) —
  it is a **new measured surface**, and the existing **100%/100%** on all 4 Clerkenstein surfaces must stay
  green (no regression of the single-identity behaviour).
- **Build on `wip/clerkenstein-browser-login`** — fold its work in rather than starting fresh.

**Out:** the cockpit panel (M38) — M37 delivers the capability; M38 the UI that drives it.

## Repo split
- **`rosetta-extensions`** `clerkenstein` (`clerk-frontend/` + `alignment/` DNA + goldens).
- **`rosetta`** corpus: a pointer doc-update in `rosetta_demo.md` / `clerk-integration.md` (the cockpit's
  identity layer).

## Parallelism
**M37 ∥ M36** — `clerkenstein` is a different ext section with no shared files; M37 needs only M35's
hero-identity list, so it can run alongside the dashboard work (M36).

## Risk
A new alignment-measured surface; the seat-switch mechanism is unproven (**O11**). **Mitigation:** an early
token-injection-vs-handshake spike; hold the alignment gate green throughout (`/align-run` each iter).

## Done-when
A demo stack can present as any seeded hero (correct org claims + role); the multi-identity Alignment DNA
scores 100%/100% alongside the existing 4 surfaces; `wip/clerkenstein-browser-login` is reconciled/retired.
