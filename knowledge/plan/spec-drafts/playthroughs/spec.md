# Playthroughs — Functional-Flow E2E Testing Spec

> **Status:** Consolidated draft `v0.3` · spec-draft · 2026-06-28
> **Companions:** [`spec-progress.md`](spec-progress.md) (decision tracker + log) · [`next-release.md`](next-release.md) (out-of-scope / parking lot)
> **Brand:** *"Playthroughs"* — a complete **play through** the product's content, following the story, the way a player completes a game; locked 2026-06-28 (spec-progress #1).
> **v0.3** folds in an in-depth review (2026-06-28): a **test-data lifecycle** (§5.7), an **integration-dependent-flows** posture (§5.8), a **negative/error-path** outcome model (§4), **engine coexistence** (§4.4), and a cluster of determinism/locator/login recalibrations across §3 / §5. See [`spec-progress.md`](spec-progress.md) for the decision rows.

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

  - **P1b — Out-of-band continuation (the only mid-flow carve-out).** Some real flows hand off through an
    artifact that arrives **out of band**: an email-confirm or org-invite link (Brevo, not started by default), a
    provider webhook (a Clerk sync). This is neither the action-under-test nor seed/reset, so P1's setup/teardown
    word doesn't cover it. Where a flow needs an out-of-band **artifact**, advance it via a **controlled non-UI
    mechanism** (a harness-captured link, a Clerkenstein-synthesized webhook), **provided the final assertion
    still lands on the user-observable outcome** — *the carve-out is the **continuation**, never the capability.*
    Separately and more simply: **waiting for an async platform projection to settle** (the verified-skill
    pipeline is async by construction) is **not** a backdoor — it is ordinary P2 outcome-assertion plus
    Playwright web-first waiting.

- **P2 — Functional truth, not pixel truth (the cardinal rule).** Assert on the **goal achieved** —
  capability, outcome, resulting state — **never** on exact copy, DOM structure, CSS, layout, or coordinates. A
  slight UI or wording change **must not** break a Playthrough. If a Playthrough fails, a *capability* failed.

  - **Assert on the OUTCOME STATE the flow produced**, never on pre-seeded specifics that vary across captures —
    a generated member name, a catalog title or count, a snapshot-versioned value. Those are inputs, not
    outcomes; binding to them re-imports the churn P2 promises immunity from.
  - **AI-generated content is explicitly on the forbidden list.** Never assert the **exact value** of LLM output —
    an actor's conversational turn, a generated document, post-session insight prose, a voice transcript. For
    AI content assert **structure / presence / range** (a result *exists*, a score is *in band*), never the
    value. Note the inverse is true and load-bearing: simulation **scoring is deterministic rubric-based** (0–100,
    NOT AI-scored, for EU-AI-Act compliance — see [`ai_architecture.md`](../../../../corpus/architecture/ai_architecture.md)
    §Evaluation System), so the *computed outcome* IS assertable exactly; only the generated *content* around it varies.
  - **Copy-immunity is within a locale, not across locales.** P2 immunizes against wording churn *in the rendered
    locale*; a **locale change** is a separate, pinned axis (§5.4) — accessible names and labels are translated, so
    an assertion is implicitly bound to the test locale. Switching language is not a "slight wording change."

- **P3 — Implementation-agnostic, zero platform coupling.** Rosetta's hard rule is **zero platform-repo edits** —
  so Playthroughs **cannot** rely on test hooks added to the platform (no bespoke `data-testid` contract). They
  locate by **semantics**: ARIA **role**, **accessible name**, **label**, the **accessibility tree** — the
  contract a *user* perceives — not brittle CSS/text/structure, and never the platform's internals (services,
  endpoints, tables). We verify what the human experiences. *(The tension this creates is **resolved** in §5.2 /
  spec-progress #3: **pure-semantic by default + a find-only landmark registry** for the surfaces the real UI
  leaves ambiguous. On the actual platform UI — antd v6, near-zero a11y attributes — that registry is **load-
  bearing**, not a thin exception; see §5.2.)*

- **P4 — One use case ↔ one Playthrough.** Each test proves **exactly one** declared use case, is **isolated**
  (its outcome doesn't depend on another test's side effects), and is **traceable both ways** via the use case's
  manifest id.

- **P5 — Manifest-first.** The **use case is declared first** — goal + flow + expectations — independently of its
  test. The Playthrough *implements the declaration*. The manifest can list a use case **before** its Playthrough
  exists (a known gap), which is exactly what makes it a **build reference** as well as a regression one.

- **P6 — Deterministic, repeatable, seeded.** A Playthrough binds to a **known stack state** (a seeded demo/dev
  world). Same inputs → same result, every time. No flakiness, no time/order/network coupling that isn't
  controlled. A flaky Playthrough is a defect in the Playthrough.

  - **Deterministic *by construction*.** The Playthrough seed must contain **no live-LLM content** (or be **fully
    cache-pinned** — see [`cache-spec.md`](../../../../corpus/ops/demo/cache-spec.md)) and must be pinned to a
    **specific taxonomy capture version**, since the AI-generated set-dress (M45's `GeneratedBatchSeeder`, live
    gpt-4o-mini cache-cold) and a re-replayed snapshot otherwise produce *different* worlds. Determinism that
    depends on a moving substrate is not determinism.
  - **Conditional on the reset model (§5.7).** The moment a Playthrough mutates the world (most write-flow use
    cases do), "same inputs → same result" holds **only if state is reset to the known seed between runs**.
    P6's guarantee is therefore *conditional on* the §5.7 reset running — and on **never** treating an additive
    re-seed as that reset (it silently leaves stale state).
  - **Seed-vs-platform drift is a first-class diagnosis.** When a Playthrough whose *precondition* is a
    short-circuited seed (e.g. a directly-written verified skill) fails, **suspect seed-vs-platform drift before
    concluding a capability regressed** — the read-time projection the seed bypasses can drift (the G14 / column-
    spelling class, see [`seeding-spec.md`](../../../../corpus/ops/seeding-spec.md)). Prefer Playthroughs that
    **drive the real pipeline** where feasible; the §5.4 seed must carry the same drift detection the seeding
    suite uses. *(The function-under-test is already drift-immune via P1/P2 — this caveat is scoped to the
    precondition subset.)*

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
| `actor.entitlement` | The actor's **entitlement tier** (anon / free / paying / enterprise / expired) — a *declared* precondition, because what a user can reach is tier-gated (§4.3). |
| `seed` / `preconditions` | The world state assumed — a **named** seeded story/preset the Playthrough seed provides (the validator resolves it, §5.3 — no silent "ideally"). |
| `engine` | For surfaces mid-migration, the engine this UC targets — `legacy` or `new-academy` (§4.4). Omitted where there is one engine. |
| `flow` | The high-level steps that serve the goal — *what the user does*, not *which selectors*. |
| `outcome` | The **expected shape** of the result: `success` (default) · `blocked` (a correct refusal — gate / deny) · `error` (a correct validation failure). A `blocked`/`error` UC asserts the *refusal* is functional truth (§4.2). |
| `expectations.intermediate[]` | Ordered outcome checkpoints along the flow ("the role now lists the candidate as *applied*"). Each entry has a **stable label** and binds 1:1 to the i-th asserted checkpoint in the Playthrough (§4.1). |
| `expectations.final` | The goal achieved (or the correct refusal landed), observable to the user ("the candidate appears in *Shortlisted* with the chosen stage"). |
| `playthrough` | The id of the test that proves it (may be `TODO` while it's still a build-reference gap). |

### 4.1 How intermediate expectations bind

`expectations.intermediate[]` is **ordered and labelled**, and the binding is mechanical so two authors implement
it the same way: **`intermediate[i]` ↔ the i-th asserted checkpoint** in the Playthrough, matched by its **stable
label**, and **reported individually** in the result map. A **failed intermediate fails the Playthrough AND aborts
the remaining flow** (a broken step-3 makes step-4's assertion meaningless). This is what makes P5 traceability
hold *mid-flow*, not just at the final assertion — and it lets §5.5 surface *which* checkpoint broke.

### 4.2 Outcome shape — success, blocked, error

A use case's *correct* outcome is not always "the goal achieved." `outcome` declares which it is:

- **`success`** (default) — the capability delivered the goal.
- **`blocked`** — the platform **correctly refuses**: an employee is denied a manager-only action (a Sentinel
  deny), an under-tier actor hits an entitlement gate and is redirected, a duplicate is prevented. The Playthrough
  asserts the *refusal* is the functional truth (per **P2**: a correct refusal is as much "the platform doing its
  job" as a success — a regression reference that only proves happy paths is blind to authz/validation regressions).
- **`error`** — a correct, user-visible **validation failure** (a malformed input is rejected with the right
  guard). Distinct from a Playthrough *defect*: this is the platform behaving correctly.

> *Negative example (shape only — not a use case to build):* an **employee** actor (`outcome: blocked`) attempts a
> manager-only roster edit; the **final** expectation is *"the action is refused — the user is shown a deny / kept
> out of the manager view,"* not an error in the test. Enumerating which negative UCs to build stays in the
> backlog ([`next-release.md`](next-release.md)); §4 fixes only the **model**.

### 4.3 Entitlement tiers

A flow's reachable surface is **tier-gated**, so the actor's **entitlement tier** is a first-class *declared*
precondition (`actor.entitlement`), not an implicit "logged-in user." This makes a whole deterministic, seedable
behaviour class expressible: *"a free user cannot open paid content"* (`outcome: blocked`), *"an enterprise user
reaches the org dashboard."* It pairs with **private-path isolation** — *"a private org-X skill path is invisible
to an org-Y user"* — which the §5.4 seed must support by spanning **tiers** and **multiple orgs with private
content**. (This fixes the *model*; which tier/isolation UCs to build is deferred — [`next-release.md`](next-release.md).)

### 4.4 Engines, surfaces & coexistence

Some products are **mid-migration** between engines (legacy skill paths on `next-web-app` + `cms` + Directus + `app`
→ the new AI-Academy engine on `ant-academy` + `app`; see
[`path-migration/spec.md`](../path-migration/spec.md)). The Playthroughs stance:

- **A goal-level use case is engine-agnostic** (per **P2**) — *"complete a skill path and verify a skill"* is a
  capability, not an engine. A legacy→new move **must not break a green Playthrough**: that green test *is* the
  "nothing lost at sunset" regression net.
- Where a UC must pin an engine (parity-testing both during coexistence), it declares `engine: legacy |
  new-academy` plus the **entitlement-tier precondition** the engine gates on.
- **Migration policy:** **re-point, don't delete.** A passing legacy Playthrough is retired **only after** the
  new-engine equivalent is green — it is the proof that the migration lost no function.

**Illustrative shape only** (a made-up use case to show the manifest's form — **not** a use case to build):

```yaml
products:
  - id: hiring
    name: "Hiring"
    stories:
      - id: hiring.shortlist-flow
        name: "A recruiter shortlists a strong candidate"
        actor: { hero: dan-manager, entitlement: enterprise }   # a seeded hero + declared tier
        seed: { preset: stories }           # the named seeded world it runs against (validator resolves it)
        use_cases:
          - id: hiring.shortlist-flow.UC1
            goal: "Recruiter moves an applied candidate to the shortlist for a role"
            outcome: success                 # default; `blocked` / `error` for correct refusals
            flow:
              - "open the role's candidate pipeline"
              - "pick an applied candidate"
              - "move them to the shortlist stage"
            expectations:
              intermediate:
                - { label: pipeline-reachable, expect: "the candidate is reachable from the role's pipeline" }
              final:
                - "the candidate now appears under the shortlisted stage"     # outcome state, not copy
            playthrough: pt-hiring-shortlist-uc1
```

(The schema is sketched here to fix the *shape*; the **exact field set is intentionally settled during the first
build** — decided at the shape level, spec-progress #4.)

---

## 5. Tech approach

### 5.1 Tooling — Playwright (and why)

**Playwright** is the principle-enabling tool, and it's already the foundation of Rosetta's M42 coverage harness
(`rosetta-extensions/stack-verify/e2e`). It gives us exactly what P2/P3/P6 require:

- **Semantic locators** — `getByRole` / `getByLabel` / `getByText` over the **accessibility tree**, so tests bind
  to what the user perceives, not the markup.
- **Auto-waiting & web-first assertions** — handle the **ordinary paint timing** that would otherwise flake. This
  is **necessary but not sufficient at org scale**: large grids paint slowly, so the M42 harness already needed a
  bounded re-assert poll (`RE_ASSERT_MAX_TRIES=6` in `section-assert.ts`) on top of auto-waiting. Playthroughs
  **inherit that bounded-poll waiting discipline**; auto-waiting does **not** by itself satisfy P6 on heavy
  surfaces, and it does nothing for response-*content* nondeterminism or minute-scale async (see §5.8).
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
*user* perceives). For ambiguous surfaces a **Rosetta-side landmark registry** supplies stable anchors — a region
heading, a unique visible label, a parent role to scope within — defined once, reused, and fixed in one place if
the platform shifts. These are anchors we **find**, never hooks we **add**: no platform edit, no platform dependency.

**Recalibration — the registry is load-bearing, not a thin exception.** The real platform UI is **antd v6 with
almost no a11y surface** (a handful of `aria-label`s across ~480 components, **0** `data-testid`, manager rosters
rendered as ~1k structurally-identical `@rc-component/table` rows), and the M42 harness itself locates via raw CSS
(`table tbody tr`, `.ant-card`). So on this UI **the ambiguous surface *is* the UI**: any *act-on-a-specific-row*
flow needs the registry. Concretely:

- **Pin anchor types to what antd actually gives us:** the page `<main>`, `h1`–`h4` region headings, **visible
  button text**, and domain text (org / role / person names). Not class names, not nth-child.
- **The discipline is "scope-within-a-named-region, then disambiguate by visible row text"** — never a bare
  `getByRole('row')` against 200 look-alike rows.
- **Cost is real:** expect a registry entry per non-trivial step, and re-pinning when antd or the layout shifts.
  M42 reuse is **read-only**; *act-on-a-named-element* is **net-new** work. (See spec-progress #3.)

**The AI / media-widget surface class is distinct.** A LiveKit voice room (what is its accessible name? how does an
actor "speak"?), a Monaco / CodeMirror code editor (poor a11y tree, virtualized DOM), a document canvas — **neither
pure-semantic nor a find-an-anchor registry can *drive* these**, and P3 forbids any platform `data-testid`. For
these, a Playthrough asserts at the **launch / completion boundary** (the widget *launched*, the completion outcome
*landed*) rather than driving the in-widget interaction — reconciled with P1's setup/teardown exemption and detailed
per third party in §5.8. This is the **known forcing case** for the registry, not a speculative "if pure-semantic
proves too ambiguous."

### 5.3 The manifest format

A **YAML manifest** (sibling-in-spirit to `stack.seed.yaml`), declared as **human-readable intent**: a use case's
`flow` and `expectations` are plain-language steps and outcome statements (*what*, never *how*) — the
**Playthrough code implements the mechanics**. This keeps the manifest a readable contract and honors P2.
**Layout: one file per product** (`hiring.yaml`, `workforce.yaml`, …), use cases nested under their stories. The
exact field list refines during the first build. (Decided — spec-progress #4.)

The **light validator** enforces, at validate-time (not as a runtime surprise):

- **Unique ids** across the manifest.
- **Both-way id integrity** (inherits **P4**): (a) every use case resolves to a live Playthrough id **or** an
  explicit `TODO` (the build-reference gap); **and** (b) every tagged Playthrough resolves to a live use-case id
  (**no orphans**), and every non-`TODO` id resolves to an **existing** test.
- **Precondition coverage:** every use case's `seed` / `preconditions` (and `actor.entitlement`) resolves to
  something the **Playthrough seed actually provides** — so a UC can never name a precondition the seed lacks and
  fail at *setup*, masquerading as a capability break (which would violate P2/P6/P7). This is a validate-time error.

The Playthrough seed, reusing the seeding machinery, is itself covered by the **same `datadna` conformance gate**
as the demo seed (so it is not a blind spot — see [`seeding-spec.md`](../../../../corpus/ops/seeding-spec.md)).

### 5.4 The actor & environment

Playthroughs run against a stack seeded to a **known state**, logging in as the hero the use case names. The seed
is a **dedicated Playthrough world, decoupled from the demo seed**: built on the same seeding machinery and stories
model (and it may be *forked* from the demo seed as a starting point), but maintained on its own. Why decoupled:
the demo evolves for showcase reasons faster than tests can chase, and test data needs stable, purpose-built
preconditions — different data for different purposes. Deterministic seed → deterministic Playthrough (P6).
(Decided — spec-progress #5.)

**Named-hero login is the demo-only Clerkenstein seat-switch.** "Logging in as a seeded hero" is **not**
environment-neutral — it *is* the M37 multi-identity mechanism (roster export → fake-FAPI → the cockpit
`?__clerk_identity=` handshake), which is **demo-stack** tooling (see
[`clerkenstein.md`](../../../../corpus/services/clerkenstein.md)). A plain **dev-N runs real Clerk with one fixed
identity** and only the light `dev-min` set-dress — **not** the stories/heroes roster. Consequently:

- **Hero-driven Playthroughs are demo-N (or any dev-N explicitly Clerkenstein-injected).** The stories *seed* can
  run on dev via `stackseed`, but the seat-switch *login* cannot — so the suite is **not** "the same on dev and
  demo" today for hero-driven flows.
- **Open build item:** wiring a dev-N roster + fake-FAPI so dev-N gains the seat-switch. The gap to close is
  **identity/auth**, not seeding (spec-progress #5, #10).

**The seed must span the declared axes.** Beyond the roster, the Playthrough seed provides what §4 declares as
preconditions: **multiple entitlement tiers** (anon / free / paying / enterprise / expired) and **multiple orgs
with private content** (so private-path-isolation UCs — *"org-X's private path is invisible to org-Y"* — are
seedable). And it must carry the **same drift detection the seeding suite uses** — the **closure gene** + the
value-validity checks that caught the **G14** class (valid enum/token/status, real public node-ids) — so a
short-circuited precondition (e.g. a directly-written verified skill) doesn't quietly diverge from the platform's
read-time projection (cf. P6's drift-diagnosis corollary; [`seeding-spec.md`](../../../../corpus/ops/seeding-spec.md)).

**Pin a canonical test locale.** The world is seeded and rendered in **English (`en`)** as part of the known
state. Per **P3**, locators resolve against accessible names / labels / text — which are **translated** — so the
locale is a **pinned axis**, not free: §5.2 locators are written against `en`, and copy-immunity (P2) means
*within-locale* wording churn, never a language switch. (Language-switch / fallback as its own flow class is an
explicit in-scope-or-parked decision — parked, [`next-release.md`](next-release.md).)

**Static fixture assets.** File-chooser flows (resume / LinkedIn import, a document-sim upload) need a **real binary
on disk** the seeder cannot substitute. Playthroughs may carry a small set of **version-controlled static fixtures**
(a sample resume PDF, a doc-sim input) in the rext `playthroughs` section — distinct from seed data — fed to the
**real** file chooser via Playwright `setInputFiles` (the action stays P1-honest).

### 5.5 Mapping, traceability & reporting

Every Playthrough is **tagged with its use-case id**; a report reconciles the manifest against results into a
**four-state map** per use case:

- **`passing`** — the Playthrough is green.
- **`failing`** — the Playthrough is red (a capability failed — or, per P6, seed-vs-platform drift; diagnose).
- **`unimplemented`** — a declared use case with **no Playthrough yet** (the build-reference gap).
- **`unimplementable-without-platform-edit`** — the surface **cannot** be driven without a platform edit (a hard
  zero-edit wall — e.g. a hardcoded URL with no per-URL override). Distinct from `unimplemented`: it **escalates,
  does not edit the platform**, mirroring the coverage sweep's re-scope trigger (see
  [`coverage-protocol.md`](../../../../corpus/ops/demo/coverage-protocol.md)). This is the P3 zero-edit escape valve.

That map **is** the coverage dashboard and the regression reference: it answers both *"what functionality is
proven?"* and *"what is declared but not yet built (or blocked)?"* The map **optionally surfaces per-checkpoint
pass/fail** (per §4.1), so a failure localizes *which* intermediate step broke rather than just "the Playthrough
went red."

### 5.6 Where it lives

In **`rosetta-extensions`**, as a **dedicated `playthroughs` section** — its own manifest, dedicated-seed wiring,
and lifecycle/reporting — built on a **shared e2e foundation** it shares with `stack-verify`, so the plumbing isn't
duplicated:

- **Login:** reuse the **existing cockpit-login helper** (the M37 roster → fake-FAPI → `?__clerk_identity=`
  handshake), not a generic "Clerkenstein-login-anywhere." Same import pattern the cockpit already uses.
- **Locators as a per-surface page-object layer.** The §5.2 registry is **not** an ambiguous-surface fallback
  bolted onto each test — it is a **shared per-surface page-object / locator layer that every Playthrough imports**
  (extending the existing `cockpit-login` import pattern from login to locators). This buys the key maintainability
  property: **a UI / antd / copy shift is absorbed by editing the per-surface registry, not N Playthrough files —
  re-pinning is O(surfaces), not O(tests).** Inherit the coverage sweep's **centralized-anchor** discipline, not
  its inline-raw-CSS one.
- **Static fixtures** (§5.4) live here too, version-controlled alongside the suite.

Authored + **tagged** like all rext tooling and consumed per-stack at a pinned tag, per the corpus's tooling
policy. (Decided — spec-progress #6.)

### 5.7 Test-data lifecycle — reset & isolation

P1 mandates that the action-under-test **mutates real state**; P6 demands *"same inputs → same result."* These two
only hold together if the world is **reset to the known seed between runs**. M42's coverage harness never needed
this (it is **read-only**); Playthroughs do, so it is decided here:

- **Reset model: per-suite reset-to-seed**, via the **dedicated seed's own machinery** (the same `stackseed`
  path that built the world rebuilds it). A run starts from the seeded baseline; after a mutating suite the world
  is reset before the next run. *(Granularity is per-suite, not per-test, to keep cost bounded; a per-test reset is
  a later tightening if isolation demands it.)*
- **Additive re-seed is FORBIDDEN as a reset.** An `ON CONFLICT … DO NOTHING` re-seed silently leaves stale state
  (the M42e "green-but-wrong" trap) — it is **not** a reset. Use the real `--reset` path, honoring its contract:
  whole-fleet granularity + the **N=0 `--force` guard** (see
  [`idempotency.md`](../../../../corpus/ops/idempotency.md) + [`seeding-spec.md`](../../../../corpus/ops/seeding-spec.md)).
- **P6 is conditional on this reset running.** A Playthrough that mutates state and skips the reset is **not**
  deterministic — the determinism headline is unfounded without §5.7.
- **Concurrency / isolation: default serial.** The runtime is a **single shared `organization_id`-scoped Postgres**
  (one container set per stack N), so two mutating Playthrough workers interfere — and Playwright defaults to
  **parallel** workers, which would bite day one. Default therefore: **serial** (`workers: 1`,
  `fullyParallel: false`). P4's "isolated" is a *logical* property; this is its *mechanism*. Sanctioned
  throughput-reclaim paths, when speed is needed: **stack-per-worker** (a stack each, via the unified registry) or
  **per-worker org/hero partitions** in the seed. *(Reset and isolation are the same problem — both are "return the
  world to a known state.")*

### 5.8 Integration-dependent flows — the assertion boundary

The platform's most product-defining journeys depend on **third-party engines**, and **Clerkenstein mocks ONLY
Clerk** — there is no mirror for LiveKit / Chime / Stripe / Brevo. So the signature AI-simulation journey (voice +
recording + an AI conversation) is either un-Playable or silently degraded **invisibly** unless we name the posture
explicitly. The stance, **per third party**:

| Third party / leg | Posture | Playthrough handling |
|---|---|---|
| **Chat / code / document** sim modalities | **playable as-is** | Drive the real UI; assert the outcome state. |
| **Voice** (LiveKit) · **recording** (Chime) | **needs a mirror first** | Park as `later — needs a mirror engine` ([`next-release.md`](next-release.md)). Until then, assert at the **boundary** (below). |
| **Payments** (Stripe) | **needs a mirror / test-mode** first | As above (Stripe **test-mode** is the likely mirror). |
| **Email** (Brevo) | **needs a sink** first | As above; pairs with P1b out-of-band continuation (a captured confirm/invite link). |

**Assertion-boundary-per-flow-class.** A **live-AI or opaque-media** flow is **not** driven turn-by-turn inside the
widget. It asserts at the **launch / completion boundary**: the flow **launched** and reached an interactive state,
and the **outcome artifact materialized** (a session row, a result, a recorded-completion marker) — *not* by
scripting the in-widget conversation. This is the only thing **provable under P6** with a live LLM in the loop.

**Corrections honored here** (do not write the un-corrected version):

- **Scoring is deterministic rubric-based** (0–100, EU-AI-Act, **NOT** an AI scorer — see
  [`ai_architecture.md`](../../../../corpus/architecture/ai_architecture.md) §Evaluation System). The determinism
  risk is the **live voice / recording / AI-conversation legs** and the **per-session computed result timing**, not
  "an AI grader varies." The computed *score* is assertable exactly (assert structure / range, never the AI prose).
- **AI / voice creds ARE provisioned** on **both dev-N and demo-N** (`/stack-secrets` — OpenAI / Azure / Anthropic
  + the LiveKit pair, Critical == 100%); only `DIRECTUS_TOKEN` (the content write-token) is stripped. A hero is
  **not** cred-blocked from a sim. The barrier is determinism + the missing mirrors, not missing wiring.
- A **deterministic / mocked-AI** sim-completion path (to assert "completed with a specific evaluation") is a
  **`later`** item ([`next-release.md`](next-release.md)) — not provable under P6 with a live LLM. The sanctioned
  shortcut to a *seeded* completed state is the PersonaSeeder short-circuit, which is an **example of P1's existing
  setup/teardown allowance**, not a new exception.

### 5.9 Manifest currency & ownership

The manifest is a **build reference**, so it must not silently go stale as the (fast-moving) platform grows new
surfaces. The honest stance:

- **Owner + trigger:** the manifest is reviewed for new surfaces on the **`/stack-update` / release cadence** — the
  same beat that syncs a stack to new platform code is when a new capability gets a use case (or an explicit `TODO`).
- **How a new surface enters:** as a new use case under the right Product/Story, `playthrough: TODO` until built —
  which is exactly the build-reference gap P5 is for.
- **Deliberately NO automated drift gate** (the accepted-risk line). Unlike seeding's **`datadna`** gate — which can
  introspect *"what structural data must exist"* — there is **no introspectable schema for "what user-facing
  capabilities exist,"** so an *added* capability with no use case cannot be auto-detected (coverage would read
  100% while the new function is unproven). The *removed* direction is largely self-catching (a vanished capability
  reds its Playthrough via P2). This asymmetry is an **accepted risk**, mitigated by the cadence review above —
  recorded as such (spec-progress).

---

## 6. Relationship to existing Rosetta capabilities

| Capability | Relationship |
|---|---|
| **Seeding (stories & heroes)** | The **substrate** + shared vocabulary. Playthroughs act on the seeded world; reuse its heroes & stories. |
| **M42 semantic-coverage sweep** | **Complementary.** Coverage proves *presence* (every reachable page shows real, believable content); Playthroughs prove *function* (the hero can complete the use case end-to-end). Presence vs. function; Playthroughs are the deeper, goal-level guarantee. |
| **Clerkenstein** | The **actor's identity** — named-hero login *is* the M37 seat-switch (roster → fake-FAPI → cockpit handshake). Reuses the **existing cockpit-login helper** (§5.6). |
| **Demo / dev stacks** | The **target environment**. Hero-driven Playthroughs run on **demo-N** (or a Clerkenstein-injected dev-N); a plain **dev-N** is real Clerk + one identity + `dev-min` — so it's **not** the same suite on dev today for hero flows (§5.4, an open build item). |
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

Tracked in [`spec-progress.md`](spec-progress.md). **All in-scope capability points are decided** — the name
(*Playthroughs*), the **locator policy** (§5.2), the **manifest** (prose-intent, per-product, §5.3), the
**seed/world** (dedicated + decoupled, §5.4), the **home** (a dedicated `playthroughs` section on a shared e2e
foundation, §5.6), and — **newly decided in v0.3 by this review** — the **test-data lifecycle** (per-suite
reset-to-seed + serial-by-default, §5.7), the **integration-dependent-flows** posture (§5.8), the
**negative/error outcome** model (§4.2) and **engine coexistence** (§4.4), the **entitlement-tier / private-path**
model (§4.3), the **intermediate-binding** convention (§4.1), the **i18n test-locale** pin (§5.4), and the
**manifest-currency** stance (§5.9).

**Carried as open build items** (decided *in shape*, the work is deferred):

- **dev-N hero login** — wiring a dev-N roster + fake-FAPI so hero-driven Playthroughs run on dev-N as well as
  demo-N. Today: demo-N (or a Clerkenstein-injected dev-N) only (§5.4).
- **Mirror engines** — voice (LiveKit), recording (Chime), payments (Stripe test-mode), email (Brevo sink), and a
  deterministic / mocked-AI sim-completion path. Parked `later — needs a mirror engine` (§5.8,
  [`next-release.md`](next-release.md)).

**The single deferred-after-this-spec item** is *which products / stories / use-cases to declare first* — the build
backlog, a natural `/developer-kit:design-roadmap` job (spec-progress #7).
