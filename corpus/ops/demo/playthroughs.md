# Playthroughs — the functional-flow e2e runbook (the Playthroughs pillar)

**A Playthrough is an automated actor that IS the user.** It logs in as a seeded hero, sets out with a goal,
and plays a real journey through the platform — start to finish, the way a person would — then proves the
platform actually delivered the outcome. The **Playthroughs** capability is the canonical, living set of these
journeys: the platform's user-facing functionality, continuously **proven to work**.

This runbook **graduates** the consolidated capability spec
([`knowledge/plan/spec-drafts/playthroughs/spec.md`](../../../knowledge/plan/spec-drafts/playthroughs/spec.md),
v0.3) into an operational corpus reference: the model, the vocabulary, the per-surface page-object layer, the
dedicated-seed + reset-to-seed lifecycle, the serial-default runner, and the four-state reporting map — as they
are actually built, in the rext **`playthroughs` section** (v2.0 "opening night" M202 "Foundation"). It is also
**the iteration protocol the coverage milestones follow** (M203 employee-vantage ∥ M204 manager-vantage — the
`iterative` milestones that grow the real journey coverage against this foundation; see § "The iteration
protocol" below). It is the *function* sibling of [`coverage-protocol.md`](coverage-protocol.md)'s *presence*
sweep.

> **Read alongside:** [`coverage-protocol.md`](coverage-protocol.md) (the M42 Playwright sweep this is built
> on + the presence-vs-function split), [`stories-spec.md`](stories-spec.md) (the Stories & Heroes seed model +
> the roster the actors reuse), [`seeding-spec.md`](../seeding-spec.md) (the seeding machinery the dedicated
> seed reuses + the production-isolation boundary), [`idempotency.md`](../idempotency.md) (the `--reset`
> contract + the N=0 guard the reset-to-seed lifecycle honors), and
> [`../../services/clerkenstein.md`](../../services/clerkenstein.md) (the M37 seat-switch the hero login rides).
>
> All the harness code lives in the gitignored `rosetta-extensions` monorepo, **section `playthroughs/`**
> (authored + tagged in the authoring copy, consumed per-stack at a pinned tag) — **zero platform-repo edits**
> (the hard line). An un-drivable surface **escalates**, it never edits the platform.

## For PMs — what a Playthrough proves

Rosetta already proves a demo world **looks** real: seeding populates a believable org, and the M42 coverage
sweep proves every page a hero reaches **shows** real content ([`coverage-protocol.md`](coverage-protocol.md)).
A Playthrough proves the hero can **do** the thing that world is for — it verifies **function**, not just
presence. It is the deeper, goal-level guarantee.

The feel we're buying: **confidence that "the product does its job," cleanly decoupled from "the pixels are
identical."** We can ship and refactor freely; a Playthrough breaks **only when a capability breaks** — never
when a button moves, a layout reflows, or copy is reworded. If a Playthrough goes red, a *capability* failed.

Every journey is declared first, in a plain-language **manifest** (a use case: a goal, the flow that serves it,
the outcomes to expect), and each declared use case is proven by exactly one automated test. The manifest
doubles as the **build reference** (what functionality must exist) and the **regression reference** (what must
keep passing as the platform grows). Coverage is simply: *use cases with a passing Playthrough ÷ declared use
cases.*

## The model & vocabulary — Products → Stories → Use Cases → Playthroughs

Playthroughs are declared in a manifest with a four-level hierarchy, deliberately mirroring the seeding
**stories** model so the two share vocabulary and substrate:

```
Product             a platform product / capability area  (Profile, Hiring, Workforce, Skill Paths, Academy, …)
└─ Story            an interconnected flow of product use — a coherent journey (may span products)
   └─ Use Case      one GOAL + the platform FLOW that serves it + the INTERMEDIATE & FINAL expectations
      └─ Playthrough the deterministic e2e test that PLAYS the use case as a human and ASSERTS its expectations
```

- **Product** — a platform product or capability area under test. The top-level grouping. One YAML file per
  product (`profile.yaml`, `hiring.yaml`, …).
- **Story** — an *interconnected flow of product use*: a coherent journey a real user takes, possibly spanning
  products. The same notion of "story" seeding uses; where it helps, a Playthrough story **reuses a seeded
  story's heroes** as its actors.
- **Use Case** — the **atomic unit of functional truth**. It declares a **goal**, the **actor** who pursues it
  (+ their entitlement tier), the **preconditions / seed** it assumes, the **flow** (the high-level steps that
  serve the goal), and its **expectations** — both **intermediate** (ordered checkpoints along the flow) and
  **final** (the goal achieved).
- **Playthrough** — the deterministic, automated e2e test that **plays one use case** as a human and asserts
  its expectations. **One use case ↔ one Playthrough**, traceable both ways by a stable id.

The **model is code** — the Go schema is
[`playthroughs/manifest/manifest.go`](../../../.agentspace/rosetta-extensions/playthroughs/manifest/manifest.go)
(`Product` → `Story` → `UseCase`, with `Actor`/`Seed`/`Expectations`/`Outcome`/`Engine`). `Load` reads one
product-file; `LoadDir` reads a directory of them and merges in sorted (deterministic) order. The manifest is
**human-readable intent**: a use case's `flow` and `expectations` are plain-language statements of *what*, never
*how* — the Playthrough **code** implements the mechanics.

### What a Use Case declares

| Field | Meaning |
|---|---|
| `id` | Stable identifier; the 1:1 link to its Playthrough. |
| `goal` | The user-meaningful outcome being pursued ("a hero logs in and sees her own identity"). |
| `actor.hero` | The seeded roster seat the actor logs in as (reuses the seeding roster), OR a free-form descriptor for a not-yet-seeded actor (a build-reference gap). |
| `actor.entitlement` | The actor's tier — anon / free / paying / enterprise / expired — a *declared* precondition (reachable surface is tier-gated). |
| `seed.world` + `seed.preconditions[]` | The named seeded world (`pt-world`) + extra named world-state the Playthrough seed provides (the validator resolves both — no silent "ideally"). |
| `engine` | For surfaces mid-migration, the engine this UC targets — `legacy` or `new-academy`. Omitted where there is one engine. |
| `flow` | The high-level steps that serve the goal — *what the user does*, not *which selectors*. |
| `outcome` | `success` (default) · `blocked` (a correct refusal — a gate / deny) · `error` (a correct validation failure). A `blocked`/`error` UC asserts the *refusal* is functional truth. |
| `expectations.intermediate[]` | Ordered, **labelled** outcome checkpoints along the flow; `intermediate[i]` binds 1:1 to the i-th asserted checkpoint, reported individually. |
| `expectations.final` | The goal achieved (or the correct refusal landed), observable to the user. |
| `playthrough` | The id of the test that proves it, OR the sentinel `TODO` while it is still a build-reference gap. |

The M202 **foundation manifest** ([`playthroughs/manifest/profile.yaml`](../../../.agentspace/rosetta-extensions/playthroughs/manifest/profile.yaml))
opened with one product (`profile`), one story, and the single proof-of-life use case
`profile.foundation.UC1` (login → /profile → assert hero identity). The M201 manifest corpus (the
user-curated 9-product / ~27-use-case surface) lands here product-by-product across the coverage milestones, each
validated by the same contract. **M203 (employee vantage) has landed** the 3 employee-vantage products —
`profile.yaml` (identity + verified-skill + growth + timeline), `skill-paths.yaml`, `ai-simulations.yaml` — as
**6 live Playthroughs, 0 TODO**; **M204** adds the manager vantage.

## The principles (the alignment contract)

These are the load-bearing rules a new Playthrough — and every reviewer — holds to. A Playthrough that violates
one is wrong even if it passes.

- **P1 — Be the human.** Drive the **real UI** as a user would. The action *under test* uses no API / DB / admin
  backdoor. Backdoors are allowed only for **setup/teardown** (seeding the world, resetting state) — never for
  the behavior being proven. (The one mid-flow carve-out: an **out-of-band artifact** — an email-confirm link, a
  provider webhook — may be advanced via a controlled non-UI mechanism, *provided the final assertion still lands
  on the user-observable outcome*.)
- **P2 — Functional truth, not pixel truth (the cardinal rule).** Assert on the **goal achieved** — capability,
  outcome, resulting state — never exact copy, DOM structure, CSS, layout, or coordinates. Assert on the
  **outcome state the flow produced**, never on pre-seeded specifics that vary across captures. **AI-generated
  content is on the forbidden list** — for LLM output assert *structure / presence / range*, never the value.
  (The inverse is load-bearing: simulation **scoring is deterministic rubric-based**, 0–100, NOT AI-scored — so
  the *computed outcome* IS assertable exactly; only the generated content around it varies. See
  [`../../architecture/ai_architecture.md`](../../architecture/ai_architecture.md) § Evaluation System.)
  Copy-immunity is *within a locale*, not across — the test locale is pinned to `en`.
- **P3 — Implementation-agnostic, zero platform coupling.** Zero platform-repo edits means we **cannot** add
  `data-testid` hooks. Playthroughs locate by **semantics**: ARIA role, accessible name, label, the a11y tree —
  the contract a *user* perceives. For the surfaces the real UI leaves ambiguous, a **find-only landmark
  registry** supplies stable anchors (anchors we *find*, never hooks we *add*).
- **P4 — One use case ↔ one Playthrough.** Each test proves exactly one use case, is isolated, and is traceable
  both ways via the use case's manifest id.
- **P5 — Manifest-first.** The use case is declared first — goal + flow + expectations — independently of its
  test. The manifest can list a use case **before** its Playthrough exists (`playthrough: TODO`), which is what
  makes it a build reference as well as a regression one.
- **P6 — Deterministic, repeatable, seeded.** A Playthrough binds to a **known stack state**. Same inputs → same
  result. The seed must carry no live-LLM content (or be fully cache-pinned) and be pinned to a taxonomy capture
  version. When a Playthrough mutates the world, P6 holds **only if state is reset to the known seed between
  runs** (§ reset-to-seed below). A flaky Playthrough is a defect in the Playthrough.
- **P7 — Stories compose; use cases prove independently.** A story's use cases may chain, but each must still be
  independently verifiable from a declared seed.
- **P8 — The spec is the alignment contract.** New products / stories / use cases extend the manifest under these
  principles.

## The tech approach

### The manifest + the light validator

The validator ([`playthroughs/manifest/validator.go`](../../../.agentspace/rosetta-extensions/playthroughs/manifest/validator.go),
run by [`cmd/ptvalidate`](../../../.agentspace/rosetta-extensions/playthroughs/cmd/ptvalidate/main.go)) enforces,
at validate-time (never a runtime surprise), three checks:

1. **Unique ids** — every product / story / use-case id is unique across the corpus (and no empty id — an empty
   id can't be a stable 1:1 link).
2. **Both-way id integrity** (inherits P4): (a) every use case resolves to a live Playthrough id **or** an
   explicit `TODO`; (b) every tagged (non-`TODO`) Playthrough id resolves to an **existing** e2e test, **and**
   every e2e test tagged `@pt:<id>` maps back to a use case (**no orphan tests**, no double-tagged id). Direction
   (b) is enabled by discovering the live registry of `@pt:` tags from the e2e specs
   ([`cmd/ptvalidate/discover.go`](../../../.agentspace/rosetta-extensions/playthroughs/cmd/ptvalidate/discover.go)).
3. **Precondition-coverage** — every use case's `seed.world`, `actor.hero`, `actor.entitlement`, and
   `seed.preconditions[]` resolves to something the dedicated seed **actually provides** (the `seed-worlds.yaml`
   index below), so a UC can never name a precondition the seed lacks and fail at *setup*, masquerading as a
   capability break. Closed enums (`outcome`, `engine`) are validated regardless of the seed index.

On top of the static half, `ptvalidate --stack demo-N` runs the **datadna closure gate** on the dedicated seed
as a subprocess (`datadna measure-closure` — the same conformance gate the demo seed is held to), so the
Playthrough seed is not a blind spot. The Go section imports **no** `stack-seeding` code — it invokes the
decoupled offset-port CLI, preserving the module boundary (§ decision M202-D2).

```bash
# static shape only (fast CI lint of the manifest):
go run ./cmd/ptvalidate --manifest-dir ./manifest
# full static validation (both-way integrity + precondition-coverage):
go run ./cmd/ptvalidate --manifest-dir ./manifest --e2e-dir ./e2e/tests --seed-worlds ./seed/seed-worlds.yaml
# + the datadna closure gate against a live seeded stack:
go run ./cmd/ptvalidate --manifest-dir ./manifest --e2e-dir ./e2e/tests --seed-worlds ./seed/seed-worlds.yaml --stack demo-1
```

### The per-surface page-object / locator layer

Locators are a **shared per-surface page-object layer every Playthrough imports** — the load-bearing
maintainability property: a UI / antd / copy shift is absorbed by editing the per-surface page object, **not N
Playthrough files** — re-pinning is **O(surfaces), not O(tests)**.

- The base is [`e2e/lib/page-object.ts`](../../../.agentspace/rosetta-extensions/playthroughs/e2e/lib/page-object.ts)
  (`PageObject`): the shared semantic-locator primitives (`goto`, `main`, `region(headingText)`, `byRole`,
  `byText`) that enforce the discipline in one place.
- The **locator discipline** (P2/P3, enforced by convention): prefer `getByRole(role,{name})` → `getByLabel` →
  `getByPlaceholder` → tolerant `getByText` → last resort a **stable landmark** (a region heading, a unique
  visible label, a parent role to scope within). **Forbid** raw CSS / nth-child / XPath / class-name / coordinate
  selectors, and any assertion on exact copy, DOM shape, or styling. The discipline is
  **"scope-within-a-named-region, then disambiguate by visible text"** — never a bare `getByRole('row')` against
  200 look-alike rows.
- **The registry is load-bearing, not a thin exception.** The real platform UI is antd v6 with almost no a11y
  surface (a handful of `aria-label`s, **0** `data-testid`). Anchor types are pinned to what antd actually gives
  us: the page `<main>`, `h1`–`h4` region headings, visible button text, and domain text (org / role / person
  names). Not class names, not nth-child.
- The **starting surface** (M202) was `/profile`:
  [`e2e/lib/profile-page.ts`](../../../.agentspace/rosetta-extensions/playthroughs/e2e/lib/profile-page.ts)
  (`ProfilePage`) — it owns the "how do I find the hero's name on /profile" knowledge (`heroName(name)` scoped
  within the identity region, `exact:false`); the test owns only the "assert her name is there" intent. M203 grew
  `ProfilePage` with the Skills/Career-tab accessors and added the skill-path + simulation surfaces (next bullet).
- **M203 adds the employee-journey surfaces**: `skill-path-page.ts` (`SkillPathPage`), `simulation-page.ts`
  (`SimulationPage`), plus the profile Skills/Career tabs on `ProfilePage`. Their **route-shape decision logic**
  (am-I-in-the-chapter-player vs still-on-detail; did-the-sim-reach-`/start` vs opened-detail) and the
  ProfileSeeder **timeline dated-range** landmark are extracted into pure, browser-free predicates in
  [`e2e/lib/url-shapes.ts`](../../../.agentspace/rosetta-extensions/playthroughs/e2e/lib/url-shapes.ts) that the
  page objects delegate to — so the resolution logic is unit-testable without a live stack
  (`tests/url-shapes.unit.spec.ts`). **Route-shape discipline (M203 harden truth):** anchor the terminal segment
  (`/chapter(?:[/?#]|$)`, `/start(?:[/?#]|$)`, `/profile/skills(?:[/?#]|$)`), **never a bare `\b`** — a bare
  word-boundary false-matches look-alike sibling segments (`/chapter-list`, `/start-now`, `/profile/skills-summary`,
  since `-` is a word boundary), a green-but-wrong hazard. Every route shape is single-sourced in `url-shapes.ts`
  (M203 close consolidated the last three inline `/profile/skills` `\b` copies into the anchored `SKILLS_TAB_URL`),
  so a re-pin is O(surfaces), not O(tests).

### Named-hero login — the cockpit seat-switch, reused

"Logging in as a seeded hero" is **not** environment-neutral — it *is* the M37 multi-identity seat-switch
(roster export → fake-FAPI → the `?__clerk_identity=` handshake), which is **demo-stack** tooling. The hero login
[`e2e/lib/hero-login.ts`](../../../.agentspace/rosetta-extensions/playthroughs/e2e/lib/hero-login.ts) **reuses
the existing cockpit-login helper** (`loginAs` from `stack-verify/e2e/lib/cockpit-login.ts`) — it does **not**
fork it, so the handshake mechanics stay single-sourced (a fix there is a fix here). It fails loud on a `/login`
bounce, so a Playthrough that silently ran unauthenticated cannot false-pass (P1).

Consequence (from the spec): **hero-driven Playthroughs run on demo-N** (or a dev-N explicitly
Clerkenstein-injected). A plain dev-N runs real Clerk with one fixed identity and only the light `dev-min`
set-dress — the stories *seed* can run on dev, but the seat-switch *login* cannot. Wiring a dev-N roster +
fake-FAPI so dev-N gains the seat-switch is a carried open build item (spec §5.4).

### The dedicated, decoupled seed

Test data ≠ demo data. The Playthrough world is a **dedicated preset decoupled from the demo seed**, built on
the same seeding machinery **unchanged (M202-D3)** (a `stack.stories.yaml` consumed by `stackseed`):
[`seed/pt-world.seed.yaml`](../../../.agentspace/rosetta-extensions/playthroughs/seed/pt-world.seed.yaml). It
seeds **two private orgs** distinct from the demo showcase orgs, spanning entitlement tiers + multi-org-private
content. The `seed-worlds.yaml` index
([`seed/seed-worlds.yaml`](../../../.agentspace/rosetta-extensions/playthroughs/seed/seed-worlds.yaml)) is
**single-sourced with the preset** — every world id / roster seat / tier / capability the validator resolves
against is materialized by the seed. It is covered by the **same datadna conformance gate** as the demo seed
(above).

> **Layering finding (M202-D4).** Seeding `pt-world` onto an *already-seeded* demo-1 collided: the stories model
> forces the FIRST story onto `LegacyOrgID` (the Clerkenstein default org), which on a seeded demo IS the
> showcase's default org — so a pt-org merged into it and duplicate-keyed on the showcase's pre-existing
> `user_skills`. The **zero-platform-edit, zero-fork fix**: `pt-world` carries a leading **anchor story** (size 0,
> no heroes) that harmlessly re-declares the demo default org, pushing the real pt orgs to story index ≥1 so they
> get their own deterministic `StoryOrgID`s and never collide. This is a genuine seeding-machinery constraint for
> a *second world on a shared stack* (the demo default-org slot is single-tenant), recorded for the coverage
> milestones to inherit.

## The lifecycle — reset-to-seed + the serial-default runner

P1 mandates the action-under-test **mutates real state**; P6 demands *same inputs → same result*. These hold
together only if the world is **reset to the known seed between runs** — so the runner
([`e2e/run-playthroughs.sh`](../../../.agentspace/rosetta-extensions/playthroughs/e2e/run-playthroughs.sh)) does
a **per-suite reset-to-seed** on `--reset`:

- It runs the **real `stackseed --reset` path** (full FK-ordered TRUNCATE, per-stack only, honoring stackseed's
  own N=0 guard) **then** a fresh seed of `pt-world.seed.yaml`. **Additive re-seed is FORBIDDEN as a reset** — an
  `ON CONFLICT DO NOTHING` re-seed silently leaves stale state (the M42e "green-but-wrong" trap). See
  [`idempotency.md`](../idempotency.md) + [`seeding-spec.md`](../seeding-spec.md) for the `--reset` contract.
- The runner **refuses N=0** (the main dev stack) outright — a Playthrough run always targets a demo-N.
- **Serial by default.** The runtime is a single shared `organization_id`-scoped Postgres, so two mutating
  Playwright workers would interfere — and Playwright defaults to *parallel*. The config
  ([`e2e/playwright.config.ts`](../../../.agentspace/rosetta-extensions/playthroughs/e2e/playwright.config.ts))
  therefore pins **`workers: 1`, `fullyParallel: false`, `retries: 0`** (a retry that masks a flake hides a
  Playthrough defect). The sanctioned throughput-reclaim paths — **stack-per-worker** (a stack each) or per-worker
  org/hero partitions in the seed — are opt-in via `PW_WORKERS`, never the day-one default.

```bash
cd playthroughs/e2e
./run-playthroughs.sh 1              # run the suite against demo-1 (serial), no reset
./run-playthroughs.sh 1 --reset      # reset-to-seed the pt-world FIRST, then run
./run-playthroughs.sh 1 --grep pt-profile-identity   # a single Playthrough by @pt tag
```

## The four-state reporting map

A report ([`report/report.go`](../../../.agentspace/rosetta-extensions/playthroughs/report/report.go),
`Reconcile`, run by [`cmd/ptreport`](../../../.agentspace/rosetta-extensions/playthroughs/cmd/ptreport/main.go))
reconciles the manifest against a run's results into a **four-state map** per use case — the coverage dashboard
AND the regression reference:

| State | Glyph | Meaning |
|---|---|---|
| **`passing`** | `[PASS]` | The Playthrough is green. |
| **`failing`** | `[FAIL]` | The Playthrough is red — a capability failed (or, per P6, seed-vs-platform drift; diagnose). A declared-but-absent test is `failing`, never a silent pass. |
| **`unimplemented`** | `[TODO]` | A declared use case with no Playthrough yet (`playthrough: TODO`) — the build-reference gap. |
| **`unimplementable-without-platform-edit`** | `[BLOCKED-PLATFORM]` | The surface cannot be driven without a platform edit (a hard zero-edit wall — e.g. a hardcoded URL with no override). It **escalates, it does not edit the platform** — the P3 escape valve, mirroring the coverage sweep's re-scope trigger. Declared deliberately (with a rationale) in [`report/unimplementable.yaml`](../../../.agentspace/rosetta-extensions/playthroughs/report/unimplementable.yaml), never inferred from a failure. |

The four glyphs are deliberately **visually distinct** — a `pending`-vs-`unimplemented` ambiguity would hide a
real semantic distinction. `Report.AllGreen()` (nothing failing/unimplementable/unimplemented) is the
foundation-complete gate; `Report.NoRegressions()` (nothing `failing`) is the gate a *coverage* milestone runs —
a build-reference `TODO` gap must not fail the suite. Coverage = passing ÷ total declared.

## The proof of life (M202)

The trivial proof Playthrough
([`e2e/tests/profile-identity.spec.ts`](../../../.agentspace/rosetta-extensions/playthroughs/e2e/tests/profile-identity.spec.ts),
tagged `@pt:pt-profile-identity`, use case `profile.foundation.UC1`) proves the whole plumbing end-to-end —
the cockpit seat-switch hero login + the page-object layer + the dedicated seed + a single user-observable
assertion, all against a live demo stack:

```
login as the seeded hero  →  open /profile  →  assert the hero's own name (Pat Ellis) renders.
```

It is deliberately the smallest real journey: it PLAYS the flow as the human (P1 — the login is the real
seat-switch, the navigation the real app) and asserts a user-observable OUTCOME (P2 — her name landed on her
profile), immune to copy/layout churn around it. It **passes green on demo-1** (M202 close).

## The iteration protocol (for M203/M204)

The coverage milestones (M203 employee-vantage ∥ M204 manager-vantage) are `iterative`: they grow the real
journey coverage against this foundation. Each iteration follows this loop — the same measure → triage → fix →
re-measure shape [`coverage-protocol.md`](coverage-protocol.md) established for the presence sweep, applied to
*function*:

1. **Declare** the next use case(s) in the manifest under the right Product/Story (from the M201 manifest
   corpus) — goal + flow + expectations, `playthrough: TODO` until built (P5). Run `ptvalidate` — the manifest
   must stay valid (unique ids, both-way integrity, precondition-coverage) at all times.
2. **Extend the seed** if a new precondition is needed — add it to `pt-world.seed.yaml` **and** `seed-worlds.yaml`
   in lockstep (they are single-sourced), and keep the datadna closure gate green. Never name a precondition the
   seed lacks.
3. **Add the page object** for any new surface (O(surfaces), not O(tests)) under the locator discipline; add the
   Playthrough spec tagged `@pt:<id>` and point the use case's `playthrough` at it.
4. **Run** `run-playthroughs.sh N --reset` (reset-to-seed, serial) → **reconcile** with `ptreport` → read the
   four-state map.
5. **Triage** each non-`passing` state: `failing` → fix the Playthrough or diagnose a real capability
   regression (suspect seed-vs-platform drift before concluding a regression on a short-circuited precondition,
   per P6); `unimplemented` → build the next Playthrough; `unimplementable-without-platform-edit` → **escalate,
   do not edit the platform** (declare it in `unimplementable.yaml` with a rationale). Fixes land in
   `rosetta-extensions` (the page-object layer, the seed, or the manifest) — **never** a platform edit.
6. **Re-measure.** The milestone's gate is `NoRegressions()` (nothing `failing`) at the vantage's declared
   use-case set; `unimplemented` gaps are the honest build-reference remainder, tracked in the map.

**Integration-dependent flows** (the assertion boundary): a live-AI or opaque-media leg (voice/LiveKit,
recording/Chime, payments/Stripe, email/Brevo — Clerkenstein mocks **only** Clerk) is **not** driven turn-by-turn
inside the widget. It asserts at the **launch / completion boundary** (the flow launched + reached an
interactive state, the outcome artifact materialized), which is the only thing provable under P6 with a live LLM
in the loop. Chat / code / document sim modalities are playable as-is. The mirror engines for the other legs are
carried as `later — needs a mirror engine` items (spec §5.8).

**Seed-then-reload for authz-gated features (M203 iter-05).** A feature whose access is gated by **Sentinel**
(a casbin policy — e.g. `FEATURE_JOB_SIMULATIONS`, which the AI-sim launch reads via
`userMembership.organizationFeatures` → the g3 grouping policy) is only effective **after the running Sentinel
enforcer RELOADS**. The seed writes the g3 grant into `sentinel.casbin_rules`, but the enforcer **caches its
policy in-memory** — a freshly-seeded grant is invisible to a running stack until an explicit `Reload` RPC. So
`run-playthroughs.sh --reset` calls Sentinel's `Reload` after re-seed (idempotent, non-fatal, zero platform
edits — it drives Sentinel's own RPC). **General rule:** any seed that writes casbin policy for a *running*
enforcer must pair with a post-seed Sentinel Reload, or the authz-gated surface false-denies despite a correct
DB grant.

## Where it lives + hard constraints

- **Section:** `rosetta-extensions/playthroughs/` — `manifest/` (Go model + validator) · `cmd/ptvalidate` +
  `cmd/ptreport` (the CLIs) · `seed/` (the dedicated preset + the seed-worlds index) · `e2e/` (the Playwright
  page-object layer + specs + the serial runner) · `report/` (the four-state map) · `fixtures/`
  (reserved for version-controlled static fixtures fed to the real file chooser, spec §5.4 —
  populated by M203/M204's upload flows; empty at the M202 foundation, whose proof Playthrough
  needs no fixture). Section README:
  [`playthroughs/README.md`](../../../.agentspace/rosetta-extensions/playthroughs/README.md).
- **Mixed toolchain (M202-D1):** Go for the manifest/validator/report (matching the seeding module's
  `datadna`/`stackseed` CLI family + the datadna-gated requirement) + TypeScript for the Playwright layer
  (matching the M42 e2e foundation). One section, two languages, each matching its reuse target.
- **Built ON the shared foundation, never forked:** the M42 e2e foundation
  (`stack-verify/e2e/lib/{cockpit-login,section-assert,empty-states,coverage-manifest}.ts`) + the seeding
  machinery (`stack-seeding/` — `stackseed --reset` + the `datadna` closure gate).
- **Zero platform-repo edits.** Authored + tagged in the authoring copy (`.agentspace/rosetta-extensions/`),
  consumed per-stack at a pinned tag. An un-drivable surface escalates via
  `unimplementable-without-platform-edit`; it never edits the platform.
- **Production-safe + isolated.** The dedicated seed rides the seeding isolation guard (structurally impossible
  for a non-prod stack to write a shared/prod store) and the reset-to-seed path honors the `--reset` contract +
  the N=0 guard. See [`../safety.md`](../safety.md).
