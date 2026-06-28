# Playthroughs — Spec Progress (open points tracker)

> **Status:** Draft · spec-draft · 2026-06-28
> Tracker + decision log for the spec, [`spec.md`](spec.md). We work decisions **one at a time** and record each
> here.

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
| 5 | **The seed/world** the suite runs against | ✅ | A **dedicated Playthrough seed, decoupled from the demo seed** (may fork it as a starting point); runs on **both dev-N and demo-N**. Rationale: demo data changes faster than tests can chase + they serve different purposes. [`spec.md`](spec.md) §5.4. |
| 6 | **Where Playthroughs live** in rext | ✅ | A **dedicated `playthroughs` section** (own manifest, dedicated-seed wiring, lifecycle) on a **shared e2e foundation** (Playwright + Clerkenstein-login + locator/landmark helpers shared with `stack-verify`). [`spec.md`](spec.md) §5.6. |
| 7 | First **coverage targets** (which products/stories first) | ⏭️ | **After** this spec — the build backlog, not part of defining the capability. The user's directive was define capability/principles/tech now, **not** list/build tests. → [`next-release.md`](next-release.md). |
| 8 | Relationship to the M42 coverage sweep | ✅ | **Complementary, not redundant.** Coverage = *presence* (pages show real content); Playthroughs = *function* (the hero can complete the use case). [`spec.md`](spec.md) §6. |
| 9 | What this capability is **not** | ✅ | Not visual-regression (the opposite of P2), not perf/load, not unit/integration, not API-contract, not security/a11y-auditing. [`spec.md`](spec.md) §7. |

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
