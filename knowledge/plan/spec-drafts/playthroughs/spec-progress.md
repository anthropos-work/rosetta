# Playthroughs — Spec Progress (open points tracker)

> **Status:** Draft · spec-draft · 2026-06-28 (tracks [`spec.md`](spec.md) `v0.3`)
> Tracker + decision log for the spec, [`spec.md`](spec.md). We work decisions **one at a time** and record each
> here. Rows **10–20** were decided 2026-06-28 by the in-depth review (see the decision log).

**Legend:** 🔴 not decided · 🟡 discussing / proposed · ✅ decided · ⏭️ deferred (→ [`next-release.md`](next-release.md))

| # | Topic (plain English) | Status | Decision |
|---|------------------------|:------:|----------|
| A | What is the capability? | ✅ | A manifest-driven suite of **e2e functional-flow tests** that **pretend to be the human** and prove the platform completes real user journeys. Defined in [`spec.md`](spec.md) §1. |
| B | The model — how the manifest is structured | ✅ | Four levels: **Product → Story → Use Case → Playthrough**, mirroring seeding's stories model. A **Use Case** = goal + flow + intermediate & final expectations. [`spec.md`](spec.md) §2, §4. |
| C | One use case = one test | ✅ | **1:1** — every declared use case maps to exactly one Playthrough, traceable both ways by a stable id (P4). [`spec.md`](spec.md) §3. |
| D | The manifest doubles as build **and** regression reference | ✅ | Yes — manifest-first (P5): a use case can be declared *before* its Playthrough exists (a build-reference gap), and the suite is the regression reference once green. Coverage = passing ÷ declared. [`spec.md`](spec.md) §1.2, §5.5. |
| E | Resilient-to-UI-churn is the cardinal principle | ✅ | **Functional truth, not pixel truth** (P2): assert on goal/outcome/state, never exact copy/DOM/CSS/layout. Slight UI or copy change must NOT break a test. [`spec.md`](spec.md) §3. |
| F | Tooling | ✅ | **Playwright** — semantic locators (role/label/accessibility-tree), auto-waiting, tracing; extends the existing M42 `stack-verify/e2e` harness. [`spec.md`](spec.md) §5.1. |
| G | Document the principles so building + extending stay aligned | ✅ | Done — [`spec.md`](spec.md) §3 (P1–P8) is the alignment contract; P8 makes the doc itself the mechanism. |
| 1 | **The brand name** for the tests / the feature | ✅ | **Decided: "Playthroughs."** A complete **play through** the product's content, following the story — the gaming sense (a playthrough completes the content/quests following the story). Maps exactly to "follow the stories, complete all their use cases," and a playthrough is about completing the journey, **not** the cosmetics. Considered: Run-throughs / Encore / Odysseys. |
| 2 | Pretend-to-be-human boundary — backdoors? | ✅ | The **action under test** uses no API/DB backdoor (P1). Backdoors allowed for **setup/teardown only** (seed/reset). [`spec.md`](spec.md) §3 P1. |
| 3 | **How a test locates elements** (given zero-platform-edit) | ✅ | **Pure-semantic by default** (role / accessible-name / label / a11y-tree) **+ a thin Rosetta-side landmark registry** for ambiguous surfaces (anchors we *find*, not hooks we *add*). No platform edits. [`spec.md`](spec.md) §5.2. |
| 4 | The manifest — how expressive + how laid out | ✅ | **Prose-intent**: `flow`/`expectations` as plain-language *what*, not executable *how* (code implements). **One file per product**; a **light validator** (unique ids, every use case → a Playthrough id or `TODO`). [`spec.md`](spec.md) §5.3. |
| 5 | **The seed/world** the suite runs against | ✅ | A **dedicated Playthrough seed, decoupled from the demo seed** (may fork it as a starting point). The seed *machinery* runs on dev-N and demo-N; but **hero-driven login is demo-N-only** (the Clerkenstein seat-switch — refined by **#20**). Rationale: demo data changes faster than tests can chase + they serve different purposes. [`spec.md`](spec.md) §5.4. |
| 6 | **Where Playthroughs live** in rext | ✅ | A **dedicated `playthroughs` section** (own manifest, dedicated-seed wiring, lifecycle) on a **shared e2e foundation** — Playwright + the **existing cockpit-login helper** (not a generic Clerkenstein-login) + a per-surface locator/landmark page-object layer shared with `stack-verify` (#20, §5.6). [`spec.md`](spec.md) §5.6. |
| 7 | First **coverage targets** (which products/stories first) | ⏭️ | **After** this spec — the build backlog, not part of defining the capability. The user's directive was define capability/principles/tech now, **not** list/build tests. → [`next-release.md`](next-release.md). |
| 8 | Relationship to the M42 coverage sweep | ✅ | **Complementary, not redundant.** Coverage = *presence* (pages show real content); Playthroughs = *function* (the hero can complete the use case). [`spec.md`](spec.md) §6. |
| 9 | What this capability is **not** | ✅ | Not visual-regression (the opposite of P2), not perf/load, not unit/integration, not API-contract, not security/a11y-auditing. [`spec.md`](spec.md) §7. |
| 10 | **Test-data lifecycle** — reset + concurrency | ✅ | **Per-suite reset-to-seed** via the dedicated seed's machinery; **additive re-seed forbidden** as a reset; P6 is *conditional on* the reset running. **Concurrency: serial by default** (`workers:1`, `fullyParallel:false`) on the shared `organization_id` DB; throughput-reclaim = stack-per-worker or per-worker org/hero partitions. [`spec.md`](spec.md) §5.7; cross-ref `idempotency.md` + `seeding-spec.md`. |
| 11 | **Integration-dependent / AI-sim flows** posture | ✅ | Posture **per third party**: chat/code/document **playable as-is**; voice (LiveKit) / recording (Chime) / Stripe / Brevo **need a mirror first** (parked `later`); a live-AI/opaque-media flow asserts at the **launch/completion boundary**, not in-widget. Scoring is **deterministic rubric** (NOT AI); creds **ARE provisioned** dev+demo (only `DIRECTUS_TOKEN` stripped). [`spec.md`](spec.md) §5.8. |
| 12 | **Negative / error-path outcome** model | ✅ | `outcome: success \| blocked \| error` discriminator (default `success`); a **correct refusal is functional truth** (P2). Model/vocabulary only — enumeration deferred. [`spec.md`](spec.md) §4.2. |
| 13 | **Engines, surfaces & coexistence** (skill-paths migration) | ✅ | A goal-level UC is **engine-agnostic** (legacy→new must not break a green Playthrough); per-UC `engine: legacy \| new-academy` + entitlement-tier precondition; **re-point, don't delete** (green legacy = "nothing lost at sunset" net). [`spec.md`](spec.md) §4.4; cross-ref [`path-migration/spec.md`](../path-migration/spec.md). |
| 14 | **Entitlement tiers + private-path isolation** | ✅ | Actor's **entitlement tier** is a declared seed/precondition; the seed spans tiers + **multi-org private content** (uses the negative-expectation primitive of #12). Model only. [`spec.md`](spec.md) §4.3, §5.4. |
| 15 | **i18n test-locale** pin | ✅ | Pin a **canonical test locale (English)** as part of the known seeded state; §5.2 locators resolve against it; copy-immunity (P2) is *within-locale*, a locale change is a separate pinned axis. Language-switch/fallback **parked**. [`spec.md`](spec.md) §5.4. |
| 16 | **Intermediate-expectation binding** | ✅ | `intermediate[i]` ↔ the i-th asserted checkpoint by **stable label**, reported individually; a failed intermediate **fails AND aborts**; §5.5 map optionally surfaces per-checkpoint pass/fail. [`spec.md`](spec.md) §4.1, §5.5. |
| 17 | **Manifest currency & ownership** | ✅ | Owner+trigger = `/stack-update`/release cadence; a new surface enters as a UC (`playthrough: TODO`); **deliberately NO automated drift gate** (no introspectable "what capabilities exist" schema — the datadna contrast) = **accepted risk**. [`spec.md`](spec.md) §5.9. |
| 18 | **Validator tightening** | ✅ | Both-way id integrity (inherits P4: UC→id-or-`TODO` **and** every tagged PT→a live UC + every non-`TODO` id exists); **precondition coverage** (every `seed`/`preconditions`/`entitlement` resolves to what the seed provides, validate-time); same `datadna` gate covers the Playthrough seed. [`spec.md`](spec.md) §5.3. |
| 19 | **Zero-edit hard-wall escape valve** | ✅ | 4th map state **`unimplementable-without-platform-edit`** (distinct from `unimplemented`) — escalate, do **not** edit the platform; mirrors `coverage-protocol.md`'s re-scope trigger. [`spec.md`](spec.md) §5.5. |
| 20 | **dev-vs-demo login substrate** | ✅ | Named-hero login **is** the demo-only Clerkenstein seat-switch (M37 roster→fake-FAPI→cockpit handshake); a plain dev-N = real Clerk + one identity + `dev-min`. Hero-driven Playthroughs are **demo-N** (or a Clerkenstein-injected dev-N); dev-N roster+fake-FAPI wiring is an **open build item**. [`spec.md`](spec.md) §5.4, §6, §8. |

---

## Decision log

### Points A–G — the capability, model & principles (decided 2026-06-28, from the founding brief)

The founding direction fixed the spine: a **manifest-driven** suite where the manifest declares **Products →
Stories → Use Cases** (an interconnected flow of product use, each use case = goal + serving flow + intermediate
& final expectations), and **each use case is covered by exactly one** automated, deterministic e2e test
(**1:1**). The manifest is **both** a building reference (declare the functional surface) **and** a regression
reference (keep it green as the platform evolves). The cardinal principle is **implementation-agnostic
resilience**: tests capture *functionality at work*, not micro-UI — a slight UI/copy change must not break them —
which points at **Playwright** (semantic / accessibility-tree locating) as the enabling tool. And the principles
themselves must be **documented so building and extending stay aligned** — which is what [`spec.md`](spec.md) §3
(P1–P8) is for. All captured in [`spec.md`](spec.md).

### Point 1 — brand name (decided 2026-06-28)

**"Playthroughs."** In gaming, a *playthrough* is completing a game's content by following its story — which is
exactly what these tests do: follow a story and complete all its use cases. It reads naturally as both the noun
for one test ("a Playthrough of this use case") and the suite ("the Playthroughs are green"), and it carries the
right connotation — *completing the journey*, not matching the cosmetics (reinforcing P2). Considered and set
aside: **Run-throughs** (theatre; the original proposal), **Encore** (regression-flavoured but undersold the
first-build role), **Odysseys** (evocative but less precise). The whole spec uses it; the spec-draft dir is
`playthroughs/`.

### Point 3 — how a test locates elements (decided 2026-06-28)

Zero-platform-edit rules out `data-testid` hooks, so locating is **pure-semantic by default** — ARIA role,
accessible name, label, the accessibility tree (what the *user* perceives), never markup. Because real
enterprise surfaces (unlabelled data grids, duplicate names, custom widgets) aren't always cleanly accessible,
we add a **thin Rosetta-side landmark registry**: a small, shared map of stable anchors (a region heading, a
unique visible label, a parent role to scope within) for the ambiguous cases — defined once, reused, fixed in
one place if the platform shifts. Crucially these are anchors we **find**, never hooks we **add** — no platform
edit, no platform dependency. Strictly-pure (no fallback) was set aside as impractical for messy surfaces;
adding test-ids to the platform is ruled out by zero-platform-edit.

### Point 4 — the manifest (decided 2026-06-28)

The manifest is **human-readable intent**, not code-in-YAML. A use case's `flow` and `expectations` are
plain-language steps and outcome statements — *what* the user does and what must be true — and the **Playthrough
code** does the actual clicking/reading. This keeps the manifest the readable build+regression contract and
honors P2 (the manifest never describes selectors or mechanics). **Layout: one file per product** (use cases
nested under stories), with a **light structural validator** (unique ids; every use case maps to a Playthrough
id or an explicit `TODO` marking a build-reference gap). The exact field set refines during the first build.
Set aside: a structured/executable-steps manifest (couples to UI mechanics, fights P2/P3) and a single
monolithic file (unwieldy at scale).

### Point 5 — the seed/world the suite runs against (decided 2026-06-28)

A **dedicated Playthrough seed, decoupled from the demo seed.** Playthroughs are for **both dev and demo** — the
same suite runs on a dev-N or a demo-N stack. They reuse seeding's machinery and the stories model, and the demo
seed can be the **starting point** (fork it), but test data is then maintained **separately** from demo data.
Why: the demo evolves for showcase reasons faster than the tests can follow, and Playthroughs need stable,
purpose-built preconditions — coupling the suite to a moving demo target would make it brittle. Different data
for different purposes. (Reusing the demo world as-is, and a shared-base + per-use-case-fixtures hybrid, were set
aside for exactly this reason — the decoupling is the point.)

### Point 6 — where Playthroughs live in the tooling (decided 2026-06-28)

A **dedicated `playthroughs` rext section** rather than folding into `stack-verify`. Coverage (presence) and
Playthroughs (function) are different capabilities — and having decoupled their *data* (Point 5), we decouple
their *code/ownership* too: Playthroughs get their own manifest, dedicated-seed wiring, and lifecycle/reporting.
The low-level e2e plumbing — Playwright, Clerkenstein-login, the locator/landmark helpers, stack binding — is
**shared** via a common foundation, so nothing is rebuilt. Folding into `stack-verify` was set aside: it would
mix two purposes in one section and muddy as the suite grows.

### Point 7 — first coverage targets (deferred 2026-06-28)

Per the founding directive, **defining the capability comes first; listing and building tests does not.** Which
products/stories/use-cases to declare first is the build backlog — parked in [`next-release.md`](next-release.md)
until the capability spec is agreed (and ideally until `/developer-kit:design-roadmap` turns it into a versioned
plan).

### Points 10–20 — the in-depth-review fold-in (decided 2026-06-28, per in-depth review)

A multi-lens in-depth review of `v0.2` found the design sound but surfaced scope holes (no test-data lifecycle, the
AI-sim tier never confronted, no negative-path vocabulary) and a few stale-status slips. `v0.3` folds in **every**
recommendation. The substantive decisions:

- **#10 Test-data lifecycle** — *(review's recommendation — user can override)* **per-suite reset-to-seed** via the
  dedicated seed's machinery, **forbidding** an additive `ON CONFLICT DO NOTHING` re-seed as a reset (the M42e
  "green-but-wrong" trap), with P6's determinism *conditional on* the reset running. Bundled with the
  **isolation** decision (reset and isolation are one problem): **serial by default** (`workers:1`,
  `fullyParallel:false`) on the single shared `organization_id` Postgres — Playwright's parallel default would bite
  immediately — with sanctioned throughput-reclaim via stack-per-worker (the registry) or per-worker org/hero
  partitions. Cross-refs `idempotency.md` + `seeding-spec.md`'s `--reset`/`--force`/N=0 contract. ([`spec.md`](spec.md) §5.7.)
- **#11 Integration-dependent flows** — *(review's recommendation — user can override)* a posture **per third
  party** (playable-as-is {chat/code/document} | needs-a-mirror {voice/LiveKit, recording/Chime, Stripe test-mode,
  Brevo sink} | deferred), and an **assertion-boundary-per-flow-class** rule: a live-AI / opaque-media flow asserts
  at the **launch/completion boundary** (launched + interactive + outcome-artifact materialized), not by driving
  the in-widget interaction. Mirror-less third parties parked `later — needs a mirror engine`. **Corrections
  honored:** sim **scoring is deterministic rubric-based** (EU-AI-Act, NOT an AI scorer — `ai_architecture.md`
  §Evaluation System) so the *risk* is the live voice/recording/conversation legs + per-session computed result,
  **not** "an AI grader varies"; and **creds ARE provisioned** on dev-N *and* demo-N (`/stack-secrets` — only
  `DIRECTUS_TOKEN` stripped), so a hero is **not** cred-blocked from a sim. ([`spec.md`](spec.md) §5.8.)
- **#12 Negative/error-path** — an `outcome: success | blocked | error` discriminator (default `success`); a P2
  sentence that asserting a **correct refusal** (a Sentinel deny, an entitlement-gate redirect, a prevented
  duplicate) is functional truth. Model/vocabulary only — enumeration of specific negative UCs stays in the
  backlog. ([`spec.md`](spec.md) §4.2.)
- **#13 Engines, surfaces & coexistence** — a goal-level skill-path UC is **engine-agnostic** (a legacy→new move
  must not break a green Playthrough); a per-UC `engine: legacy | new-academy` field + entitlement-tier
  precondition; the **re-point-don't-delete** migration policy where a green legacy Playthrough is the "nothing
  lost at sunset" net, retired only after the new-engine equivalent is green. Cross-refs
  [`path-migration/spec.md`](../path-migration/spec.md). ([`spec.md`](spec.md) §4.4.)
- **#14 Entitlement tiers + private-path isolation** — the actor's **entitlement tier** is a first-class declared
  seed/precondition; the §5.4 seed spans tiers + **multiple orgs with private content** (the seedable
  *"org-X's private path is invisible to org-Y"* class, built on the #12 negative-expectation primitive). Model
  only. ([`spec.md`](spec.md) §4.3, §5.4.)
- **#15 i18n test-locale** — *(review's recommendation — user can override)* pin a **canonical test locale
  (English)** as part of the known seeded state; §5.2 locators resolve against it; copy-immunity (P2) is
  *within-locale*, while a locale change is a separate pinned axis. Language-switch/fallback parked as its own
  flow class. ([`spec.md`](spec.md) §5.4.)
- **#16 Intermediate-expectation binding** — `intermediate[i]` ↔ the i-th asserted checkpoint by **stable label**,
  reported individually; a failed intermediate **fails AND aborts** the remaining flow; the §5.5 map optionally
  surfaces per-checkpoint pass/fail so a failure localizes the broken step. ([`spec.md`](spec.md) §4.1, §5.5.)
- **#17 Manifest currency & ownership** — owner + trigger tied to the `/stack-update` / release cadence; a new
  surface enters as a use case (`playthrough: TODO`); and a recorded **accepted-risk** line that there is
  **deliberately NO automated drift gate** (no introspectable "what capabilities exist" schema — the contrast with
  seeding's `datadna`). ([`spec.md`](spec.md) §5.9.)
- **#18 Validator tightening** — the §5.3 validator gains **both-way id integrity** (inherits P4: every UC →
  id-or-`TODO`, AND every tagged Playthrough → a live UC, every non-`TODO` id exists) and **precondition coverage**
  (every UC's `seed`/`preconditions`/`entitlement` resolves to what the Playthrough seed provides, at validate-time
  not runtime). The Playthrough seed is covered by the same `datadna` conformance gate. ([`spec.md`](spec.md) §5.3.)
- **#19 Zero-edit hard-wall escape valve** — a 4th reporting map state
  **`unimplementable-without-platform-edit`** (distinct from `unimplemented`): the surface can't be driven without
  a platform edit → **escalate, do not edit the platform**, mirroring `coverage-protocol.md`'s re-scope trigger.
  ([`spec.md`](spec.md) §5.5.)
- **#20 dev-vs-demo login substrate** — named-hero login *is* the **demo-only Clerkenstein seat-switch** (M37
  roster → fake-FAPI → cockpit `?__clerk_identity=` handshake); a plain dev-N runs real Clerk with one fixed
  identity + the light `dev-min` set-dress. So hero-driven Playthroughs are **demo-N (or a Clerkenstein-injected
  dev-N)**, and the v0.2 "same suite on dev *and* demo" claim is corrected; dev-N roster + fake-FAPI wiring is
  flagged as an **open build item**. §5.6's shared foundation reuses the **existing cockpit-login helper**.
  ([`spec.md`](spec.md) §5.4, §6, §8.)

Beyond these decisions, `v0.3` applied the review's pure-clarifications in place: the **stale-status** flips (P3's
"key open tech decision" → resolved; §4's "open item" → settled-at-shape), the **feasibility recalibration** (the
landmark registry is **load-bearing** on antd v6's near-zero-a11y UI, promoted to a shared **per-surface
page-object layer** so re-pinning is O(surfaces) not O(tests); §5.1's auto-waiting softened with the
`RE_ASSERT_MAX_TRIES` poll), the **determinism corollaries** (P6 deterministic-by-construction seed; P2 asserts
OUTCOME STATE and forbids asserting the *value* of LLM output), the **P1b out-of-band continuation** carve-out
(and the async-projection-wait-is-not-a-backdoor clarification), the **file-upload static-fixtures** note (§5.4),
and the **seed-vs-platform-drift** first-class-diagnosis caveat (§5.4 seed carries the closure gene + the G14
value-validity checks).
