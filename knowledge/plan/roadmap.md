# Roadmap

Active development plan for **Project Rosetta** (the Anthropos documentation corpus + environment-
builder skills). This file holds the **active major** only; the retired **v1.x** history (M0 … M46, all
SHIPPED) lives in [`roadmap-legacy.md`](roadmap-legacy.md). Future versions + the unscheduled backlog live in
[`roadmap-vision.md`](roadmap-vision.md). The live source of truth for *current/next* is [`state.md`](state.md).

> **Designed 2026-06-28** via `/developer-kit:design-roadmap`. **v2.0 "opening night"** opens a **NEW MAJOR** —
> **Playthroughs** is a new pillar (functional-flow *testing*: proving the platform's core user journeys actually
> work end-to-end), distinct from the v1.x demo/seeding lineage. v2+ adopts the **`Mxyy`** milestone-numbering
> scheme (M201, M202, M203, M204 — major 2, milestone 01/02/03/04). The v1.x flat counter (M0 … M46) is closed and
> archived in `roadmap-legacy.md`.

## Version plan

| Version | Codename | Theme | Milestones | Status |
|---------|----------|-------|------------|--------|
| **v2.0** | **opening night** | The platform's core user journeys, **proven to actually work** — a new **Playthroughs** pillar: a manifest-driven, deterministic e2e suite that *pretends to be the human* and proves the platform does its job | M201 ✅ ∥ M202 → { M203 ∥ M204 } → ship | ⏸ **PAUSED** after M201 (branch `release/02.00-opening-night`) — a **v1.10 backfill** is interposed (re-sync + re-ground prod), then v2.0 resumes |

> The complete v1.x version-plan table (v1.0 "body double" … v1.10 "method acting", all ✅ SHIPPED) is preserved
> in [`roadmap-legacy.md`](roadmap-legacy.md) § Version plan.

The Playthroughs capability is governed by the consolidated **capability spec**
[`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md) (v0.3 — all in-scope decisions made +
review-hardened). v2.0's milestones build the contract that spec defines. Hard constraints carry over from the
v1.x lineage: **no modification to any platform repo** (the platform stays read-only — a surface that can't be
driven without a platform edit *escalates*, it does not edit), and all stack-operating tooling lives in
**`rosetta-extensions`** (built + tested in the `.agentspace/rosetta-extensions/` authoring copy, tagged, then
consumed per-stack at a pinned tag). Playthroughs reuse the M42 e2e foundation + the seeding machinery — they are
the **functional** sibling of M42's **presence**-only coverage sweep.

---

## In Development — v2.0 "opening night"

> **Theme:** *the platform's core user journeys, proven to actually work.* A **Playthrough** is an automated
> actor that **is the user** — it logs in as a seeded hero, sets out with a goal, plays through a real journey
> across the platform start-to-finish the way a person would, then proves the platform delivered the outcome.
> The capability is the **canonical, living set of these journeys**: the platform's user-facing functionality,
> continuously **proven to actually work** — cleanly decoupled from *"the pixels are identical"* (a Playthrough
> breaks **only when a capability breaks**). It is the **functional** sibling of v1.x's M42 coverage sweep
> (which proves *presence* — every reachable page **shows** real content); Playthroughs prove the hero can **do**
> the things that world is for.
>
> **Designed 2026-06-28** via `/developer-kit:design-roadmap`, from the consolidated capability spec
> [`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md) (v0.3). **A new MAJOR** — Playthroughs is
> a new pillar distinct from the demo/seeding lineage; v2+ uses **`Mxyy`** milestone numbering. **Tooling + docs
> only — zero platform-repo edits** (the read-only platform line carries over; an un-drivable surface escalates
> via the `unimplementable-without-platform-edit` state, it never edits the platform).

### Execution graph

```
v2.0 "opening night"
  M201 ──┐                          (manifest corpus — prose, user-guided)
  M202 ──┼──→ M203 ─┐
                M204 ─┴──→ ship
```

**M201 (the manifest corpus) and M202 (the foundation) open in parallel.** M201 is the **user-guided manifest
curation** — prose-only (the goal-aligned Product → Story → Use Case corpus), so it carries **no code dependency**
and can be authored before / alongside M202. M202 is the **Playthroughs foundation** (the section, the manifest
model + the §5.3 **validator**, the page-object layer, the dedicated seed + reset lifecycle, the runner + 4-state
reporting, one trivial proof Playthrough) — it builds the validator + dedicated seed to **match** the M201 corpus.
Then the two **vantage-coverage** milestones — **M203** (employee) and **M204** (manager) — run **in parallel**,
both `iterative`, implementing Playthroughs against the M201-declared use cases on the M202 foundation; the release
ships when both gates fire.

**M201 ∥ M202 (manifest ∥ foundation).** No hard ordering: M201 produces the **prose contract** (the use-case
manifest); M202 produces the **machinery** (validator + dedicated seed) that validates + seeds against it. They
reconcile when M202's validator runs over the M201 corpus. Where an M201 use case names a **precondition the demo
seed lacks**, that feeds the **M202 dedicated-seed expansion** (M201 records the need; M202 builds the seed).

**Parallelism note (M203 ∥ M204).** The two coverage milestones share an **additive merge surface**: the
per-surface **landmark registry** + the **locator index** (the §5.6 page-object layer every Playthrough imports).
Each vantage adds its own surfaces/anchors to that shared layer — an **additive** merge, not a conflicting one.
Both are `iterative` (the use-cases are *declarable* in the M201 corpus, but getting them green against the real
antd UI + the AI-sim assertion boundary is exploratory, like M42e/M42m), so they advance independently toward
their own exit gates and reconcile the registry additively at merge.

### Milestones

**M201 — Manifest corpus** · `iterative` · **USER-GUIDED** · complexity **large** · depends on: **none** (the
manifest is prose — authorable before/parallel to the M202 foundation).
**Status:** ✅ **`done` — closed-on-gate 2026-06-29.** 9 products · 26 stories · 28 use-cases authored,
**adversarially re-grounded** (11-agent wf `wvpnpvozh` → 15/27 runnable), **user-signed-off**. Records:
[`releases/02.00-opening-night/m201-manifest-corpus/`](releases/02.00-opening-night/m201-manifest-corpus/)
(deliverable: `manifest-draft.yaml`). The close discovered the **stale-clone drift** (next-web 115+ commits behind
prod) → **v2.0 PAUSED for a v1.10 backfill** (re-sync + re-ground + re-validate, user-driven) before resuming.
**Goal:** top-down, **user-directed** curation of the **full goal-aligned Product → Story → Use Case manifest
corpus** — the build + regression contract every coverage milestone (M203/M204) implements against. The flow per
pass: **outline** (products → stories → use cases) → **validate** (against the real platform surface + the spec's
manifest model) → **write the prose-intent manifest YAML** (spec §5.3, **one file per product**).
**Explicitly NOT bounded by the current minimal/partially-aligned demo stories seed** — it captures what the goal
says must be proven; where a use case needs preconditions the demo lacks, that **feeds the M202 dedicated-seed
expansion** (noted, not resolved here).
**Shape:** `iterative`, **driven by the user** — worked conversationally + via `/developer-kit:work-mstone-iters`,
not autonomously.
**Exit gate:** **the manifest corpus is comprehensively outlined, validated, and written as prose-intent YAML (one
file per product)** — covering the platform's products × their must-work user journeys, each use case carrying
**goal + actor + flow + intermediate/final expectations**, structurally valid (the spec §5.3 validator passes,
**ids unique + both-way**) — **and the USER signs off the corpus as the complete-enough v2.0 coverage contract.**
**iteration_protocol_ref:** the capability spec
[`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md) (esp. §2 model, §4 use-case shape, §5.3
manifest format).
**Delivers →** the prose-intent manifest YAML corpus (one file per product); **lands in the rext `playthroughs`
section once M202 exists**, until then drafted under the milestone dir / `spec-drafts/playthroughs/manifest-draft/`.
**Candidate starting outline (the user directs — NOT fixed):** (a) the demo-covered products — **Skill Paths, AI
Simulations, Profile & Skills, Workforce Intelligence, Hiring, Academy**; (b) goal-aligned areas the demo barely
covers (flag *to confirm with the user*) — **Auth & Onboarding, Billing & Entitlements/tier-gates, Org Admin &
Settings, cross-product journeys** (candidate→employee).

**M202 — Playthroughs Foundation** · `section` · complexity **medium** · depends on: **none** (reuses the M42
harness + the seeding machinery; the M201 manifest corpus is its build+regression contract, authorable in parallel).
**Goal:** stand up the **`playthroughs` rext section** on the **shared M42 e2e foundation**, proven by **one
trivial end-to-end Playthrough**.
**Scope — In:**
- the **manifest model + a light validator** — both-way id integrity (use-case ↔ Playthrough, traceable both
  directions) + precondition-coverage (every declared `seed`/`preconditions` resolves to a named seeded world,
  no silent "ideally"), **datadna-gated** (the Playthrough seed is covered by the same `datadna` conformance gate
  as the seeding machinery);
- the **per-surface locator/landmark page-object layer** (the §5.6 shared registry every Playthrough imports —
  a UI/antd/copy shift is absorbed by editing the per-surface registry, not N tests) — **1 surface to start**;
- the **dedicated, decoupled seed** preset (test data ≠ demo data; the demo seed is the *starting point* but
  kept separate) — **spans entitlement tiers + multi-org-private**;
- the **reset-to-seed lifecycle + serial-default runner** (`workers: 1`, `fullyParallel: false`; reset via the
  real `--reset` path honoring its contract — **additive re-seed is FORBIDDEN as a reset**);
- the **4-state reporting map** — `passing` / `failing` / `unimplemented` / `unimplementable-without-platform-edit`
  (the last being the P3 zero-edit escape valve — it escalates, never edits);
- **one trivial proof Playthrough** — **login → /profile → assert hero identity** (the foundation's smoke test).
**Out:** real product coverage (M203+); the AI-sim / integration mirror tier; cross-vantage.
**Delivers →** a corpus runbook that **graduates the playthroughs spec** (e.g.
[`corpus/ops/demo/playthroughs.md`](../../corpus/ops/demo/playthroughs.md)) — becomes the `iteration_protocol_ref`
for M203/M204.
**KB deps (read as contract):** the playthroughs spec-draft
[`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md), the **M201 manifest corpus**
([`releases/02.00-opening-night/m201-manifest-corpus/`](releases/02.00-opening-night/m201-manifest-corpus/) — the
prose contract the validator + seed implement against),
[`corpus/ops/demo/coverage-protocol.md`](../../corpus/ops/demo/coverage-protocol.md),
[`corpus/ops/seeding-spec.md`](../../corpus/ops/seeding-spec.md),
[`corpus/ops/idempotency.md`](../../corpus/ops/idempotency.md).
**Reuse paths (cite in spec-notes):** `stack-demo/rosetta-extensions/stack-verify/e2e/lib/{cockpit-login,
section-assert,empty-states,coverage-manifest}.ts`, `stack-demo/rosetta-extensions/stack-seeding/`.

**M203 — Employee-vantage coverage** · `iterative` · complexity **large** · depends on: **M202** ·
parallel-with: **M204** (caveat: the shared **landmark-registry + locator index** is an *additive* merge surface;
both are iterative).
**Goal:** **Maya's** core **employee** journeys play green (declared in the M201 manifest corpus) —
- **Skill Paths** (browse → enroll → complete → verify-skill),
- **AI Simulations** (chat/code launch → complete → score-in-range, **NON-voice**),
- **Profile** (verified-skill chart + the claimed-vs-verified gap + work/education timeline).
**Exit gate:** **every declared employee-vantage use case has a passing Playthrough on a COLD reset-to-seed demo
stack (the 3 employee stories), with 0 false-fails over 5 consecutive reset runs.**
**iteration_protocol_ref:** the playthroughs spec / the M202-delivered runbook
([`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md) until M202 graduates it to
`corpus/ops/demo/playthroughs.md`).
**Why iterative:** the use-cases are *declarable* (in the M201 corpus), but getting them green against the real
antd UI (the landmark layer) + the AI-sim assertion boundary is **exploratory**, like M42e.
**Re-scope trigger:** a surface that can't be driven without a platform edit (the
`unimplementable-without-platform-edit` state) → **escalate, don't edit**.

**M204 — Manager-vantage coverage** · `iterative` · complexity **large** · depends on: **M202** ·
parallel-with: **M203**.
**Goal:** **Dan's** core **manager** journeys play green (declared in the M201 manifest corpus) —
- **Workforce funnel** + member roster,
- **member drill-down** (the activity-dashboard),
- **succession / at-risk** (the Growth tab) signals.
**Exit gate:** **same shape as M203, manager-vantage** — every declared manager-vantage use case has a passing
Playthrough on a COLD reset-to-seed demo stack, with 0 false-fails over 5 consecutive reset runs.
**iteration_protocol_ref:** same as M203 (the spec / the M202-delivered runbook).
**Why iterative:** same as M203 — declarable use-cases, exploratory path against the real manager UI + the
assertion boundary.
**Re-scope trigger:** same — `unimplementable-without-platform-edit` → escalate, don't edit.

### Top risks

- **manifest completeness → no auto-gate, user owns "enough".** The M201 manifest is a **build reference** with
  **no introspectable schema for "what user-facing capabilities exist"** (spec §5.9) — an *added* platform
  capability with no use case cannot be auto-detected. The corpus's completeness is a **user judgement** (the M201
  exit gate's sign-off), not a machine check. *Mitigation:* M201 is **user-guided + iterative** (the user directs
  each top-down pass + signs off the complete-enough contract); the cadence-review stance (§5.9) carries forward.
- **antd-a11y → the landmark layer is load-bearing.** zero-platform-edit means we **cannot** add `data-testid`;
  locators bind to the **accessibility tree** (`getByRole`/`getByLabel`/`getByText`) over the **real antd UI**,
  with a Rosetta-side **landmark registry** for ambiguous surfaces. If antd's a11y is thin on a surface, that
  surface's landmark anchors carry the test — the registry's quality is the risk. *Mitigation:* the per-surface
  page-object layer (re-pin is O(surfaces), not O(tests)); start with **1 surface** in M202 to prove the pattern.
- **determinism-under-mutation → M202's reset must be solid.** P6 ("same inputs → same result") holds **only if**
  the world is reset to the known seed between runs, and an *additive* re-seed silently leaves stale state (the
  M42e "green-but-wrong" trap). The whole determinism headline rests on M202's **reset-to-seed lifecycle** being
  correct — it is a **foundation** risk, surfaced and owned in M202 before M203/M204 lean on it.
- **hero-login → demo-N only.** Hero-driven Playthroughs run on **demo-N** (or a Clerkenstein-injected dev-N) —
  a plain dev-N is real Clerk + one identity + `dev-min`, so the hero suite is **not** the same on dev today. The
  target environment is the demo stack; the dev-stack hero-flow generalization is an explicit **later** item
  (spec §5.4), not v2.0 scope.
- **AI-sim mirror tier is future.** The signature voice/recording AI-simulation journey needs a **mirror engine**
  (Clerkenstein mocks **only** Clerk — no LiveKit/Chime/Stripe/Brevo mirror). v2.0 covers the **NON-voice**
  chat/code/document sims (playable as-is, asserted at the launch/completion boundary); voice + recording +
  payments + email are parked as `later — needs a mirror engine` → **M206** ([`roadmap-vision.md`](roadmap-vision.md)).

---

## Shipped releases

The complete shipped history — **v1.0 "body double"** (2026-06-03, tag `v1.0`) through **v1.10 "method acting"**
(2026-06-27, tag `v1.10`), 11 versions / milestones M0 … M46 — is preserved in
[`roadmap-legacy.md`](roadmap-legacy.md) (the retired v1.x major). Records are archived under
[`releases/archive/`](releases/archive/). No v2.0 release has shipped yet.

## Notes

- **Milestone numbering — v2+ uses `Mxyy`** (`M` + major digit + two-digit milestone): **M201, M202, M203, M204**
  for v2.0. This is the major-version scheme `context.md` reserved for *"a future *major* v2+"*; the v1.x flat
  sequential counter (M0 … M46, with the `a`/`b`/`c`/`e`/`m` suffix conventions) is **closed** and lives in
  [`roadmap-legacy.md`](roadmap-legacy.md) § Notes.
- **Milestone shapes** mix within v2.0: **M201 is `iterative` + USER-GUIDED** (the manifest corpus — a top-down,
  user-directed prose curation toward a sign-off gate); **M202 is `section`** (a fixed In-scope checklist — the
  foundation is decomposable up front); **M203 + M204 are `iterative`** (a measurable exit gate, exploratory path
  — getting declared use-cases green against the real antd UI + the AI-sim assertion boundary, like the M42e/M42m
  precedent).
- Date format throughout: ISO `YYYY-MM-DD`.
- The Playthroughs capability **graduated from spec-draft to active development** at v2.0 design (2026-06-28); the
  governing spec is [`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md), graduated to a corpus
  runbook (`corpus/ops/demo/playthroughs.md`) by M202.

_Last updated: 2026-06-28 (**v2.0 "opening night" DESIGNED + PROMOTED** via `/developer-kit:design-roadmap` — a
NEW MAJOR opening the **Playthroughs** pillar; **4 milestones M201 ∥ M202 → { M203 ∥ M204 }** [M201 manifest
corpus inserted as the user-guided prose contract]; branch `release/02.00-opening-night` cut from `main`. v1.x
history rotated to `roadmap-legacy.md`. Designed from the consolidated capability spec
`spec-drafts/playthroughs/spec.md` v0.3. Tooling + docs only — zero platform-repo edits.)_
