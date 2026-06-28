# Run-throughs — Spec Progress (open points tracker)

> **Status:** Draft · spec-draft · 2026-06-28
> Tracker + decision log for the spec, [`spec.md`](spec.md). We work decisions **one at a time** and record each
> here.

**Legend:** 🔴 not decided · 🟡 discussing / proposed · ✅ decided · ⏭️ deferred (→ [`next-release.md`](next-release.md))

| # | Topic (plain English) | Status | Decision |
|---|------------------------|:------:|----------|
| A | What is the capability? | ✅ | A manifest-driven suite of **e2e functional-flow tests** that **pretend to be the human** and prove the platform completes real user journeys. Defined in [`spec.md`](spec.md) §1. |
| B | The model — how the manifest is structured | ✅ | Four levels: **Product → Story → Use Case → Run-through**, mirroring seeding's stories model. A **Use Case** = goal + flow + intermediate & final expectations. [`spec.md`](spec.md) §2, §4. |
| C | One use case = one test | ✅ | **1:1** — every declared use case maps to exactly one Run-through, traceable both ways by a stable id (P4). [`spec.md`](spec.md) §3. |
| D | The manifest doubles as build **and** regression reference | ✅ | Yes — manifest-first (P5): a use case can be declared *before* its Run-through exists (a build-reference gap), and the suite is the regression reference once green. Coverage = passing ÷ declared. [`spec.md`](spec.md) §1.2, §5.5. |
| E | Resilient-to-UI-churn is the cardinal principle | ✅ | **Functional truth, not pixel truth** (P2): assert on goal/outcome/state, never exact copy/DOM/CSS/layout. Slight UI or copy change must NOT break a test. [`spec.md`](spec.md) §3. |
| F | Tooling | ✅ | **Playwright** — semantic locators (role/label/accessibility-tree), auto-waiting, tracing; extends the existing M42 `stack-verify/e2e` harness. [`spec.md`](spec.md) §5.1. |
| G | Document the principles so building + extending stay aligned | ✅ | Done — [`spec.md`](spec.md) §3 (P1–P8) is the alignment contract; P8 makes the doc itself the mechanism. |
| 1 | **The brand name** for the tests / the feature | 🟡 | **Proposed: "Run-throughs"** (a theatre run-through = a full performance to prove the show flows; fits Rosetta's theatre lineage + the stories/heroes model; inherently flow-not-words). Alternatives: **Playthroughs** (gaming — complete the content following the story), **Encore** (regression-flavoured — perform again on cue), **Odysseys** (journey-flavoured). Lock with a find-replace. |
| 2 | Pretend-to-be-human boundary — backdoors? | ✅ | The **action under test** uses no API/DB backdoor (P1). Backdoors allowed for **setup/teardown only** (seed/reset). [`spec.md`](spec.md) §3 P1. |
| 3 | **Pure-semantic vs. a "semantic landmark" convention** (given zero-platform-edit) | 🔴 | Default = **pure semantic** (role / accessible-name / a11y-tree); **no platform `data-testid` dependency** (zero-platform-edit). Open: whether to also bless a lightweight set of already-exposed **semantic landmarks** as load-bearing anchors — WITHOUT any platform edit. [`spec.md`](spec.md) §5.2. |
| 4 | The authoritative **manifest schema** + file layout + validation | 🔴 | Shape sketched ([`spec.md`](spec.md) §4) — `products[] → stories[] → use_cases[]{id,goal,actor,seed,flow,expectations,runthrough}`. The exact schema, on-disk layout, and a validator are TBD. |
| 5 | **Stack binding** — which seeded world the suite runs against | 🔴 | Reuse the seeded **stories & heroes** world (§2 symmetry) — open whether a **dedicated** Run-through seed/preset or the existing demo stories world; and how a manifest `seed:` resolves to a real stack. [`spec.md`](spec.md) §5.4. |
| 6 | **Harness relationship** — extend M42 vs. a new section | 🔴 | Lean: **extend `stack-verify/e2e`** (it owns the Playwright foundation); promote to a dedicated `run-throughs` rext section only if it grows to warrant it. [`spec.md`](spec.md) §5.6. |
| 7 | First **coverage targets** (which products/stories first) | ⏭️ | **After** this spec — the build backlog, not part of defining the capability. The user's directive was define capability/principles/tech now, **not** list/build tests. → [`next-release.md`](next-release.md). |
| 8 | Relationship to the M42 coverage sweep | ✅ | **Complementary, not redundant.** Coverage = *presence* (pages show real content); Run-throughs = *function* (the hero can complete the use case). [`spec.md`](spec.md) §6. |
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

### Point 1 — brand name (proposed 2026-06-28)

Proposed **"Run-throughs."** In theatre a *run-through* is a complete, uninterrupted performance done to prove
the show flows end-to-end — exactly what these tests do, and it sits naturally in Rosetta's theatre-lineage
codenames and its stories/heroes vocabulary; it also connotes *flow*, not exact words, reinforcing P2. The whole
spec is written with it, but it's a **find-replace away** from any alternative (*Playthroughs*, *Encore*,
*Odysseys*). To be **locked by the user.**

### Point 3 — semantic locating under zero-platform-edit (open 2026-06-28)

The hard constraint: Rosetta makes **zero platform-repo edits**, so we **cannot** add `data-testid` hooks for
the tests. That actually *reinforces* P2/P3 — we locate by what the **user** perceives (ARIA role, accessible
name, label, the accessibility tree), never by markup. The open sub-question is whether, on top of pure-semantic
locating, we sanction a small set of **already-exposed semantic landmarks** (headings/roles/labels we agree to
treat as load-bearing) as a documented fallback — strictly without any platform change. Default stands at
**pure-semantic**; revisit if real surfaces prove too ambiguous to locate reliably.

### Point 7 — first coverage targets (deferred 2026-06-28)

Per the founding directive, **defining the capability comes first; listing and building tests does not.** Which
products/stories/use-cases to declare first is the build backlog — parked in [`next-release.md`](next-release.md)
until the capability spec is agreed (and ideally until `/developer-kit:design-roadmap` turns it into a versioned
plan).
