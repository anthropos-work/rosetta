---
milestone: M219
slug: readiness-renders
version: v2.3 "cue to cue"
milestone_shape: section
status: planned
created: 2026-07-13
last_updated: 2026-07-13
complexity: medium
depends_on: M217
parallel_with: M218, M220
delivers: the AI-readiness story VISIBLE (Dana's filled page, Ben's started workflow, Aria's completed state) + the ai-readiness playthrough manifest (their journeys are not e2e-proven today)
issues: "the SEEDING is a verified no-op (already default-on and proven); the story does not RENDER — Dana's page needs M217's re-pinned patch AND a new demo-patch for the CycleID==nil → buildLiveResponse default that bypasses the frozen-snapshot seed"
---

# M219 — Readiness renders

## Goal
The AI-readiness story is **visible**, not merely seeded.

## The seeding is a VERIFIED NO-OP — do not budget seeder work

The user's ask was *"make sure a well-seeded stack contains at least 1 organization with AI readiness enabled (if
it's already there, ok, no need to do it)"*. **It is already there.** Confirmed independently by two agents, in code,
not docs:

| Fact | Evidence |
|------|----------|
| **Northwind Aviation** (`narrative: ai-readiness`, 200 members) with heroes **Aria COMPLETED / Ben STARTED / Dana manager** — in the **DEFAULT** preset | `stack-seeding/presets/stories.seed.yaml:118-153` |
| All 3 seeders (`OrgSettingsSeeder`, `AIReadinessConfigSeeder`, `AIReadinessFunnelSeeder`) registered **unconditionally**; they self-gate on the narrative discriminator | `cmd/stackseed/main.go:410,411,431` in `buildRegistry`, called at `:470`; `seeders/org_settings.go:18,69` |
| "AI-readiness **enabled**" = a `public.organization_settings` row (`setting='ai_readiness', is_enabled=true`) — **written** | `seeders/org_settings.go:72-73` |
| The 2nd gate (PostHog `flag_ai_readiness`) is satisfied because the demo bakes **no** `NEXT_PUBLIC_POSTHOG_KEY` | `corpus/services/ai-readiness.md:36-40` |
| **STARTED vs COMPLETED are both produced** (stage 1 = Step-1 only; stage 3 = all 3 steps; manager stage 0) | `seeders/ai_readiness_funnel.go:177-196` |
| **Live proof:** the last run on `billion` wrote `org rows=3`, `ai-readiness-config rows=6`, `org-settings rows=1` | `coldrun2.log:311-338` |

**The gap is RENDERING, not seeding.**

## Why section
The deliverables are enumerable: one known platform-shaped read-path gap (→ a new demo-patch), three render proofs,
one e2e manifest, one stale comment.

## Scope

### In
1. **Dana (manager) sees a FILLED AI-readiness page.** Two blockers, both known:
   - M217's **re-pinned `app-aireadiness-snapshot-loadmembers` patch** (dead on sha-drift today).
   - **The default GET takes `buildLiveResponse` when `CycleID == nil`**
     (`app/internal/workforce/ai_readiness.go:285,301`) — so the **frozen-snapshot seed is bypassed** unless the
     frontend passes `?cycle=`. This is **platform-shaped** ⇒ it routes to a **NEW sha-pinned demo-patch** (the
     sanctioned hatch, per **D-DESIGN-2**), **never a platform edit**.
2. **Ben's from-scratch STARTED workflow is visible on his dashboard** (the seeded stage-1 funnel state renders).
3. **Aria's COMPLETED state renders.**
4. Fix the **stale ACTIVE-vs-CLOSED comment** at `stories.seed.yaml:112-117` — the code writes `status='closed'`
   (`ai_readiness_config.go:98,143`).

### In — inherited from M218 (Fate-3, added at the M218 close, 2026-07-14)

M218 could not land these **without invalidating its own exit gate**: its gate is *a p95 over 5 consecutive cold
reset-to-seed cycles*, graded on a specific binary, and **iter-05 D13 established that a demo-runtime change
restarts the 5-cycle count**. Shipping a runtime fix *after* the battery would mean shipping something other
than what was measured. They land here, where a fresh battery is affordable.

5. **`FIX-M219-bapi-org-eid` (F-11) — the BAPI fabricates the org eid.** Clerkenstein's
   `organizationWithEid` returns `"org_eid_" + orgID` in `organization.public_metadata.eid` for any org that
   isn't the hardcoded demo org, instead of the **real** org UUID the demo roster carries (real Clerk gets it
   via `UpdateClerkOrganizationWithExternalId`). It is the **ORG-level twin of the user-identity stub that cost
   M218 ~6 s per authenticated render** — same defect, same blind spot, one field over.
   - **It currently ships as a deliberately RED alignment gene** (`MembershipOrgIdentity/real-org-eid`,
     `standard` weight) so the score stops lying — the Go surface reports **97.2% overall / 100% critical**
     (gate ≥95/=100 ⇒ still MET), and the divergence is named on **every single run** until this lands. See
     **M218 D16**. **Do not "fix" the score by omitting the gene** — a silently-omitted field is exactly how
     M218's headline bug survived four releases.
   - **DoD:** `resources.go` returns the roster's real `org_eid`; the gene goes green; **and a fresh 5-cycle
     cold battery is run** (the change is on the demo's runtime path).
6. **`TEST-M219-expressrun-dep-gate` — `expressrun` is UNMEASURABLE, and scores nothing.** Without
   `@clerk/express` `node_modules` it exits **rc=2 with no score** — and nothing treats that as a failure.
   Consequence, recorded honestly: **M218 iter-04's "all 5 surfaces 100%" claim is not reproducible** — only
   **4 of 5** surfaces were measured at the M218 harden pass. **Pre-existing** (identical at baseline
   `f296e5e`), not an M218 regression. **DoD:** a missing dependency **fails loud**; it must never present as
   *absence of a score* (the D17 *absence-read-as-success* class).
7. **`TEST-M219-freshness-gate-skips` — the demo-patch live-clone freshness gate SKIPS when
   `stack-demo/next-web-app` is absent**, so a box without that clone gets **no anchor-drift protection** at
   all. Itself an instance of *absence read as success*. **DoD:** the skip is explicit and loud, or the gate
   finds another source of truth.

### Out
- Any seeder work (the seed is proven — see above).
- The manager page's **speed** — that is M218's business, and per **D-DESIGN-1** the grid's data-load is reported,
  not gated. **This milestone proves it RENDERS; M218 proves the LOGIN is fast.**

## Delivers → knowledge/corpus
- **An `ai-readiness` playthrough manifest** — **BLIND AREA.** The e2e suites cover profile / workforce /
  skill-paths / ai-simulations / assignment-monitoring only. **Aria's and Ben's journeys are not e2e-proven at
  all.** Plus its section in `corpus/ops/demo/playthroughs.md`.
- Updates to **`corpus/services/ai-readiness.md`** — the doc is otherwise **exemplary** (`:120-192` documents the
  seeder contract properly). **Build against it; do not re-derive it.** Add the read-path/`?cycle=` finding.

## KB dependencies
- `corpus/services/ai-readiness.md` (the seeder contract + the product surfaces — the contract to build against)
- `corpus/ops/demo/stories-spec.md` (the 7-table verified-skill fan-out) · `corpus/ops/demo/seed-manifest-spec.md`
- `corpus/ops/demo/playthroughs.md` (the manifest model the new playthrough must conform to)
- `corpus/ops/demo/demopatch-spec.md` ← **authored by M217** (the contract for the new patch this milestone writes)
