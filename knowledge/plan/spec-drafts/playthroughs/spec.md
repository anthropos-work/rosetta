# Playthroughs — Functional-Flow E2E Testing Spec

> **Status:** Consolidated draft `v0.2` · spec-draft · 2026-06-28
> **Companions:** [`spec-progress.md`](spec-progress.md) (decision tracker + log) · [`next-release.md`](next-release.md) (out-of-scope / parking lot)
> **Brand:** *"Playthroughs"* — a complete **play through** the product's content, following the story, the way a player completes a game; locked 2026-06-28 (spec-progress #1).

This spec defines a **new Rosetta capability**: a manifest-driven suite of high-level, end-to-end functional
tests that **pretend to be the human** and prove the Anthropos platform actually *does its job*. It defines the
**capability, the model, the principles, and the tech approach** — it does **not** enumerate or build the actual
products / stories / use-cases / tests (that is the work this spec governs).

---

## 1. Overview

### 1.1 North star — what it is

A **Playthrough** is an automated actor that *is the user*. It logs in as a seeded hero, sets out with a goal,
and plays through a real journey across the platform — start to finish, the way a person would — then proves the
platform delivered the outcome. The capability is the **canonical, living set of these journeys**: the
platform's user-facing functionality, continuously **proven to actually work**.

**How it feels:** confidence that *"the product does its job,"* cleanly decoupled from *"the pixels are
identical."* We ship and refactor freely; a Playthrough breaks **only when a capability breaks** — never when a
button moves, a layout reflows, or copy is reworded.

This is the natural sibling to Rosetta's existing work. Seeding (**stories & heroes**) populates a believable
world; the M42 coverage sweep proves every page a hero reaches **shows** real content. Playthroughs prove the
hero can **do** the things that world is for — they verify **function**, not just presence.

### 1.2 The goal

A **deterministic, manifest-driven** suite of e2e tests, each mapped **1:1 to a declared use case**, that:

- **prove** the platform's functional flows work end-to-end, played as a human,
- are **resilient to UI / copy churn** — they capture *functionality at work*, not markup or wording,
- double as the **build reference** (what functionality must exist) **and** the **regression reference** (what
  must keep passing as the platform evolves).

The manifest is the contract. The tests are its deterministic, automated proof — one test per use case, each
traceable back to the declaration it satisfies.

### 1.3 Scope of this spec

In scope: the **capability definition, the manifest model & vocabulary, the principles, and the tech approach**
(tooling, resilient-locating discipline, the manifest format, the actor/environment, traceability, where it
lives).

Explicitly **out of scope here:** enumerating the real products, writing the stories/use-cases, or implementing
any test. Those are governed by this spec and happen after it is agreed. (And the *non-goals* — what
Playthroughs deliberately are **not** — are in §7 + [`next-release.md`](next-release.md).)

---

## 2. The model & vocabulary

Playthroughs are declared in a **manifest** with a four-level hierarchy — deliberately mirroring the seeding
**stories** model so the two share vocabulary and substrate:

```
Product             (1) a platform product / capability area  (Hiring, Workforce, Skill Paths, Simulations, Academy, …)
└─ Story            (2) an interconnected flow of product use — a coherent journey (may span products)
   └─ Use Case      (3) one GOAL + the platform FLOW that serves it + the INTERMEDIATE & FINAL expectations
      └─ Playthrough (4) the deterministic e2e test that PLAYS the use case as a human and ASSERTS its expectations
```

- **Product** — a platform product or capability area under test. The top-level grouping.
- **Story** — an *interconnected flow of product use*: a coherent journey a real user takes, possibly spanning
  several products. The same notion of "story" that seeding uses; where it helps, a Playthrough story **reuses a
  seeded story's heroes** as its actors.
- **Use Case** — the **atomic unit of functional truth**. It declares: a **goal**, the **actor** who pursues it,
  the **preconditions / seed** it assumes, the **flow** (the high-level steps that serve the goal), and its
  **expectations** — both **intermediate** (checkpoints along the flow) and **final** (the goal achieved).
- **Playthrough** — the deterministic, automated e2e test that **plays one use case** as a human and asserts its
  expectations. **One use case ↔ one Playthrough**, traceable both ways by a stable id.

**The manifest** is the **single source of truth**: it declares Products → Stories → Use Cases and is, at once,
the **build reference** (the functional surface that must exist) and the **regression reference** (the surface
that must keep passing). Coverage is simply: *use cases with a passing Playthrough ÷ declared use cases.*

**Symmetry with seeding.** Playthroughs reuse seeding's **machinery and stories model** — the seeded world is the
level, the heroes are the players, the use cases are the objectives they complete. But the Playthrough world is a
**dedicated, decoupled seed** (§5.4): the demo seed can be the *starting point*, yet test data and demo data are
kept **separate** — they serve different purposes and change at different rates.

---

## 3. Principles

> These are the load-bearing contract. They exist so that *building* and *extending* the suite stay aligned — a
> new Playthrough that violates a principle is wrong even if it passes. Every Playthrough and every reviewer
> holds to all of them.

- **P1 — Be the human.** Drive the **real UI** as a user would. The action *under test* uses no API / DB / admin
  backdoor — if a user does it by clicking, the Playthrough clicks. Verify **user-observable outcomes**.
  *(Backdoors are allowed only for **setup/teardown** — seeding the world, resetting state — never for the
  behavior being proven.)*

- **P2 — Functional truth, not pixel truth (the cardinal rule).** Assert on the **goal achieved** —
  capability, outcome, resulting state — **never** on exact copy, DOM structure, CSS, layout, or coordinates. A
  slight UI or wording change **must not** break a Playthrough. If a Playthrough fails, a *capability* failed.

- **P3 — Implementation-agnostic, zero platform coupling.** Rosetta's hard rule is **zero platform-repo edits** —
  so Playthroughs **cannot** rely on test hooks added to the platform (no bespoke `data-testid` contract). They
  locate by **semantics**: ARIA **role**, **accessible name**, **label**, the **accessibility tree** — the
  contract a *user* perceives — not brittle CSS/text/structure, and never the platform's internals (services,
  endpoints, tables). We verify what the human experiences. *(The tension this creates — pure-semantic vs. a
  negotiated landmark convention — is the key open tech decision; see §5.2 + spec-progress #3.)*

- **P4 — One use case ↔ one Playthrough.** Each test proves **exactly one** declared use case, is **isolated**
  (its outcome doesn't depend on another test's side effects), and is **traceable both ways** via the use case's
  manifest id.

- **P5 — Manifest-first.** The **use case is declared first** — goal + flow + expectations — independently of its
  test. The Playthrough *implements the declaration*. The manifest can list a use case **before** its Playthrough
  exists (a known gap), which is exactly what makes it a **build reference** as well as a regression one.

- **P6 — Deterministic, repeatable, seeded.** A Playthrough binds to a **known stack state** (a seeded demo/dev
  world). Same inputs → same result, every time. No flakiness, no time/order/network coupling that isn't
  controlled. A flaky Playthrough is a defect in the Playthrough.

- **P7 — Stories compose; use cases prove independently.** A story is an *interconnected* flow, so its use cases
  may **chain** (one's final state is the next's precondition) — but each use case must still be **independently
  verifiable** from a declared seed, so a single Playthrough can run and prove its one truth.

- **P8 — The spec is the alignment contract.** New products / stories / use cases extend the **manifest** under
  these principles; this document is how building and extending stay honest as the suite — and the platform —
  grow. When a principle and a convenient shortcut conflict, the principle wins (or the principle changes, here,
  on purpose).

---

## 4. What a Use Case declares

A use case is the unit of functional truth. It declares (and its Playthrough asserts) the following — note that
every expectation is written at the **outcome / semantic** level, per **P2**:

| Field | Meaning |
|---|---|
| `id` | Stable identifier; the 1:1 link to its Playthrough. |
| `goal` | The user-meaningful outcome being pursued ("shortlist a candidate for a role"). |
| `actor` | The seeded hero who performs it (reuses the seeding roster). |
| `seed` / `preconditions` | The world state assumed — ideally a named seeded story/preset. |
| `flow` | The high-level steps that serve the goal — *what the user does*, not *which selectors*. |
| `expectations.intermediate[]` | Outcome checkpoints along the flow ("the role now lists the candidate as *applied*"). |
| `expectations.final` | The goal achieved, observable to the user ("the candidate appears in *Shortlisted* with the chosen stage"). |
| `playthrough` | The id of the test that proves it (may be `TODO` while it's still a build-reference gap). |

**Illustrative shape only** (a made-up use case to show the manifest's form — **not** a use case to build):

```yaml
products:
  - id: hiring
    name: "Hiring"
    stories:
      - id: hiring.shortlist-flow
        name: "A recruiter shortlists a strong candidate"
        actor: dan-manager                 # a seeded hero
        seed: { preset: stories }           # the seeded world it runs against
        use_cases:
          - id: hiring.shortlist-flow.UC1
            goal: "Recruiter moves an applied candidate to the shortlist for a role"
            flow:
              - "open the role's candidate pipeline"
              - "pick an applied candidate"
              - "move them to the shortlist stage"
            expectations:
              intermediate:
                - "the candidate is reachable from the role's pipeline"      # outcome, not 'the 3rd row'
              final:
                - "the candidate now appears under the shortlisted stage"     # capability, not copy
            playthrough: pt-hiring-shortlist-uc1
```

(The schema is sketched here to fix the *shape*; the authoritative format + field rules are an open item —
spec-progress #4.)

---

## 5. Tech approach

### 5.1 Tooling — Playwright (and why)

**Playwright** is the principle-enabling tool, and it's already the foundation of Rosetta's M42 coverage harness
(`rosetta-extensions/stack-verify/e2e`). It gives us exactly what P2/P3/P6 require:

- **Semantic locators** — `getByRole` / `getByLabel` / `getByText` over the **accessibility tree**, so tests bind
  to what the user perceives, not the markup.
- **Auto-waiting & web-first assertions** — kills the timing flakiness that P6 forbids.
- **Tracing / video / screenshots** — first-class failure diagnosis.
- **Cross-browser + headless/headed**, and a real-browser fidelity that "is the human" honestly.

Playthroughs **extend the existing M42 Playwright foundation** rather than introducing a parallel stack.

### 5.2 Resilient locating — how P2 / P3 are enforced in code

A strict **locator discipline** turns the principles into mechanics:

- **Prefer, in order:** `getByRole(role, { name })` → `getByLabel` → `getByPlaceholder` → tolerant
  `getByText`/accessible-name matching (substring / i18n-aware), → last resort a **stable landmark**.
- **Forbid:** raw CSS / nth-child / XPath / class-name / coordinate selectors, and any assertion on exact copy,
  DOM shape, or styling.
- **Assert on state & outcome:** that the **goal-state is reachable and visible** (the right data landed in the
  right place), expressed through accessible semantics — not "the button says X at position Y".

**The locator policy (decided — spec-progress #3):** zero-platform-edit means we **cannot** add `data-testid`
hooks, so locating is **pure-semantic by default** (role / accessible-name / label / the a11y tree — what the
*user* perceives). For genuinely ambiguous surfaces (unlabelled data grids, duplicate names, custom widgets), a
**thin Rosetta-side landmark registry** supplies stable anchors — a region heading, a unique visible label, a
parent role to scope within — defined once, reused, and fixed in one place if the platform shifts. These are
anchors we **find**, never hooks we **add**: no platform edit, no platform dependency.

### 5.3 The manifest format

A **YAML manifest** (sibling-in-spirit to `stack.seed.yaml`), declared as **human-readable intent**: a use case's
`flow` and `expectations` are plain-language steps and outcome statements (*what*, never *how*) — the
**Playthrough code implements the mechanics**. This keeps the manifest a readable contract and honors P2.
**Layout: one file per product** (`hiring.yaml`, `workforce.yaml`, …), use cases nested under their stories. A
**light validator** enforces structure — unique ids, and every use case maps to a Playthrough id or an explicit
`TODO` (the build-reference gap). The exact field list refines during the first build. (Decided — spec-progress
#4.)

### 5.4 The actor & environment

Playthroughs run against a stack — **dev-N or demo-N, the same suite for both** — seeded to a **known state**,
logging in via **Clerkenstein** as the hero the use case names. The seed is a **dedicated Playthrough world,
decoupled from the demo seed**: built on the same seeding machinery and stories model (and it may be *forked*
from the demo seed as a starting point), but maintained on its own. Why decoupled: the demo evolves for showcase
reasons faster than tests can chase, and test data needs stable, purpose-built preconditions — different data for
different purposes. Deterministic seed → deterministic Playthrough (P6). (Decided — spec-progress #5.)

### 5.5 Mapping, traceability & reporting

Every Playthrough is **tagged with its use-case id**; a report reconciles the manifest against results into a
**three-state map** per use case — `passing` / `failing` / `unimplemented` (a declared use case with no
Playthrough yet). That map **is** the coverage dashboard and the regression reference: it answers both *"what
functionality is proven?"* and *"what is declared but not yet built?"*

### 5.6 Where it lives

In **`rosetta-extensions`**, as a **dedicated `playthroughs` section** — its own manifest, dedicated-seed wiring,
and lifecycle/reporting — built on a **shared e2e foundation** (the Playwright + Clerkenstein-login +
locator/landmark helpers it shares with `stack-verify`, so the plumbing isn't duplicated). Authored + **tagged**
like all rext tooling and consumed per-stack at a pinned tag, per the corpus's tooling policy. (Decided —
spec-progress #6.)

---

## 6. Relationship to existing Rosetta capabilities

| Capability | Relationship |
|---|---|
| **Seeding (stories & heroes)** | The **substrate** + shared vocabulary. Playthroughs act on the seeded world; reuse its heroes & stories. |
| **M42 semantic-coverage sweep** | **Complementary.** Coverage proves *presence* (every reachable page shows real, believable content); Playthroughs prove *function* (the hero can complete the use case end-to-end). Presence vs. function; Playthroughs are the deeper, goal-level guarantee. |
| **Clerkenstein** | The **actor's identity** — Playthroughs log in as a seeded hero through it. |
| **Demo / dev stacks** | The **target environment** the suite runs against. |
| **Alignment tests** (mirror-engine fidelity) | A **different axis** — alignment measures how faithfully a mirror reproduces a source engine; Playthroughs measure functional truth of the real platform. |

---

## 7. Out of scope for this capability (non-goals)

Playthroughs are **functional-flow truth** and nothing else. The following are different test classes,
deliberately **not** this capability (parked in [`next-release.md`](next-release.md)):

- **Visual-regression / pixel-diff testing** — the *opposite* of P2; we intentionally do not care about pixels.
- **Performance / load / stress testing** — Playthroughs assert *function*, not latency/throughput.
- **Unit / integration testing** — a platform-repo concern, below the user-facing flow.
- **API-contract testing** — Playthroughs verify at the user surface, not the wire.
- **Security / penetration testing** and **accessibility *auditing*** — distinct disciplines (we *use* the
  a11y tree to locate; we do not *audit* it).

---

## 8. Open / to-confirm

Tracked in [`spec-progress.md`](spec-progress.md). **All in-scope points are decided** — the name
(*Playthroughs*), the **locator policy** (§5.2), the **manifest** (prose-intent, per-product, §5.3), the
**seed/world** (dedicated + decoupled, dev *and* demo, §5.4), and the **home** (a dedicated `playthroughs`
section on a shared e2e foundation, §5.6).

The only remaining item is **deferred to after this spec**: *which products / stories / use-cases to declare
first* — the build backlog, a natural `/developer-kit:design-roadmap` job.
