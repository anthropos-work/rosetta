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
**the iteration protocol the coverage milestones followed** (M203 employee-vantage ∥ M204 manager-vantage — the
`iterative` milestones that grew the real journey coverage against this foundation to 10 live Playthroughs, and
then M219 `ai-readiness` → M225 `hiring` → M243 the assign-WRITE → M252 `studio` → M256 `org-admin` +
`onboarding` took it to **30 live Playthroughs, 1 verdicted TODO** (31 manifest use cases across 10
products) the corpus stands at today; see § "The iteration protocol" below). It is the *function*
sibling of [`coverage-protocol.md`](coverage-protocol.md)'s *presence* sweep. **The `spec.md` v0.3 draft it
graduates is FROZEN at 2026-06-28** (pre-M219): where the two disagree — notably §5.7's serial-concurrency
rationale — **this runbook wins**.

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
Product             a platform product / capability area  (Profile, Hiring, Workforce, Skill Paths, Academy, Studio, …)
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
| `actor.entitlement` | The actor's tier — anon / free / paying / enterprise / expired — a *declared* precondition (reachable surface is tier-gated). **⚠️ DECLARED-ONLY — no seeder materializes a tier (v2.8 M256 pre-flight).** The validator only checks the string is one of the world's declared `tiers:`; `blueprint.TierMix` is parsed, defaulted (`blueprint/stories.go` §`DefaultStoryTierMix`) and validated but **consumed by no seeder** — it never reaches a DB column, and `pt-world.seed.yaml` declares no `tier_mix` at all. So an entitlement **gate cannot be exercised today**, and a `blocked` outcome cannot be produced by a tier — which is exactly why M256 iter-11's `blocked` Playthrough is gated on an **org feature grant** (a real Casbin g3 row) instead. All live use cases were `entitlement: enterprise` until iter-11 declared one `entitlement: free`; that string still reaches no DB column, so it documents intent and nothing more. |
| `seed.world` + `seed.preconditions[]` | The named seeded world (`pt-world`) + extra named world-state the Playthrough seed provides (the validator resolves both — no silent "ideally"). |
| `engine` | For surfaces mid-migration, the engine this UC targets — `legacy` or `new-academy`. Omitted where there is one engine. |
| `flow` | The high-level steps that serve the goal — *what the user does*, not *which selectors*. |
| `outcome` | `success` (default) · `blocked` (a correct refusal — a gate / deny) · `error` (a correct validation failure). A `blocked`/`error` UC asserts the *refusal* is functional truth. **`blocked` is EXERCISED since v2.8 M256 iter-11** — `ai-simulations.access-denied.UC1` (`pt-aisim-org-feature-blocked`), the first and so far only one; see § "The `blocked` outcome" below for how a real refusal was produced. `error` is **still at ZERO**. Note the M256 pre-flight's finding still stands and is why the refusal is not tier-based: `actor.entitlement` is **declared-only** (no seeder writes a tier, and `ptvalidate`'s precondition check fail-opens on it), so a `blocked` outcome must come from a **real refusal surface** — an RBAC/Sentinel deny, a cross-org access attempt, or a validation error — never from an entitlement tier. |
| `expectations.intermediate[]` | Ordered, **labelled** outcome checkpoints along the flow; `intermediate[i]` binds 1:1 to the i-th asserted checkpoint, reported individually. |
| `expectations.final` | The goal achieved (or the correct refusal landed), observable to the user. |
| `playthrough` | The id of the test that proves it, OR the sentinel `TODO` while it is still a build-reference gap. |

The M202 **foundation manifest** ([`playthroughs/manifest/profile.yaml`](../../../.agentspace/rosetta-extensions/playthroughs/manifest/profile.yaml))
opened with one product (`profile`), one story, and the single proof-of-life use case
`profile.foundation.UC1` (login → /profile → assert hero identity). The M201 manifest corpus (the
user-curated 9-product / ~27-use-case surface) lands here product-by-product across the coverage milestones, each
validated by the same contract. **M203 (employee vantage) landed** the 3 employee-vantage products —
`profile.yaml` (identity + verified-skill + growth + timeline), `skill-paths.yaml`, `ai-simulations.yaml` — as
**6 live Playthroughs**; **M204 (manager vantage) landed** the manager products — `workforce.yaml` (funnel /
roster / succession) + `assignment-monitoring.yaml` (the per-member activity-dashboard drill-down) — as **4 more
live Playthroughs** (`pt-workforce-funnel`, `pt-workforce-roster`, `pt-workforce-succession`, `pt-activity-drilldown`).
**M219 (v2.3 "cue to cue") landed `ai-readiness.yaml`** — the AI-readiness product, as **4 more live
Playthroughs** (see below). **M225 (v2.4 "casting call") landed `hiring.yaml`** — the recruiter-vantage candidate
comparison, as **1 more live Playthrough** (`pt-hiring-recruiter-compare`; see below). **M243 (v2.6 "sound check")
landed the assign-WRITE half** — `assignment-monitoring.assign-and-track.UC1` (`pt-assignment-assign`), the one
net-new journey, which flips the last in-manifest `TODO`. **M252 (v2.7 "july jitter") added `studio-builders.yaml`**
— Product **"Studio"**, studio-desk's FIRST-EVER manifest entry, as **2 more** Playthroughs
(`pt-studio-advanced-generate` + `pt-studio-guided-generate`; see below). **M256 (v2.8 "fast build") opened the
`org-admin` product** — one of the M201 curated corpus's four un-homed clusters for five releases — as **2 more**
live Playthroughs (`pt-orgadmin-tag-create` + `pt-orgadmin-setting-toggle`, both WRITES read back through a full
reload) with 2 declared `TODO` carrying written diagnoses, and **added `skill-paths.save-for-later`** as **1 more**
(`pt-skillpath-bookmark`, a write + a delete, both read back), and **opened the `onboarding` product** — the
LAST whole surface in the M201 curated corpus that no e2e suite had ever touched, *the first thing every real
user does* — as **1 more** live Playthrough (`pt-onboarding-complete`, itself a **net-new, non-curated** use
case) with **all 5 CURATED onboarding use cases declared and carrying written verdicts**, and **added
`workforce-intelligence.organization-feedback`** as **1
more** (`pt-workforce-org-feedback` — un-homed for five releases until pricing it revealed its data was already
seeded), and **opened the `ai-simulations.access-denied` story** as **1 more** (`pt-aisim-org-feature-blocked`) —
**the suite's FIRST `outcome: blocked`**, i.e. the first Playthrough that proves the platform correctly says *no*
(see § below). M256 then **COMPLETED the org-admin product 4 of 4** (`pt-orgadmin-member-tag` at iter-17,
`pt-orgadmin-role-create` at iter-22 — the latter's parked draft *could not have passed*: the app navigates to
`/enterprise/roles/<serverAssignedId>?setup=true` after Save and the list paginates at 20/page, so the new role
is not on page 1 even after a **successful** create), and **landed the onboarding product's second use case**
(`pt-onboarding-aireadiness-guided` at iter-26 — see § "The day-0 readiness seat" below) **and its third**
(`pt-onboarding-hiring-candidate` at iter-27 — the suite's first Playthrough to drive an onboarding flow in the
**hiring** app; see § "The hiring-org day-0 candidate" below) **and its fourth**
(`pt-onboarding-org-prepared` at iter-29 — the org-prepared variant; see § "The org-prepared onboarding
variant") **and its fifth LIVE use case — the 5th of the 6 it declares in the manifest** (`pt-onboarding-individual`
at iter-32 — **the ORG-LESS user**, the LAST un-homed
use case in the M201 curated corpus and the one this milestone's own pre-flight audit had priced as
impossible; see § "The org-less seat" below). The corpus stands at **32 live Playthroughs, 1 verdicted TODO**
(33 MANIFEST use cases, 11 products). The 30/31/10 figures that stood here were proven live-GREEN on a local
`demo-2`, 0 flake over 3 consecutive cold reset-to-seed runs; **v2.8 M258 added the two below**, each proven
RED-then-GREEN against the live `demo-1` on `billion` rather than in a cold suite run:

- **`academy` (a NET-NEW product, `academy.read-a-chapter.UC1` → `pt-academy-chapter-module`).** ant-academy is
  a first-class demo surface — its own port, its own cockpit link, its own four demo-patches — and it had
  **zero** Playthroughs, which is exactly why it broke invisibly and stayed broken across several rounds of
  hand-checking. It asserts **liveness, not presence**: a click on a module CARD (a `<button onClick>` with no
  href fallback) must advance the reader. See [`tailscale-serve.md`](tailscale-serve.md) for the failure it
  covers — a page that SSRs perfectly and never hydrates, invisible to every HTML-level check.
- **`assignment-monitoring.nav-v2.UC1` → `pt-assignments-nav-v2`.** Click the *assign content* nav entry and
  assert it opens `/enterprise/assignments-list`. Its own story on purpose: the assign Playthroughs navigate
  **by URL**, so they pass while the menu points somewhere else entirely. `flag_enable_assignments_v2` has
  **three** call sites and was patched twice at the two nobody clicks.

**Onboarding's one remaining CURATED use case is a WRITTEN VERDICT, and that is a result rather than a shortfall.**
`enterprise-workforce-standard.UC1` (the self-import journey) is `disposition: will-not-build`: its only
advancing path scrapes a live public third-party profile from a site that blocks automation, so shipping it
would make a real person's profile a permanent fixture and **its RED would read as a product regression when
nothing about the product had changed**; the deterministic CV route is blocked by a measured product defect
we do not own. See § "A `TODO` must carry a WRITTEN VERDICT" above — the position is now in the artifact the
tooling reads, not only in prose.

> **Onboarding was thought unseedable, and it was a schema misreading (v2.8 M256 iter-07/08).** The milestone's
> own KB-fidelity audit concluded there was **no pre-onboarding state and none could be declared**, reasoning
> from `UsersSeeder` writing a membership for every seeded user. **Membership is not onboarding.** Onboarding
> completion lives in **`public.user_params.onboarding`** (a `jsonb` column — there is no onboarding table), and
> it is **NULL for all 191 seeded users**: the pre-onboarding state is the **DEFAULT**, already present for every
> hero, and `/onboarding` drives as-is. Had the audit's claim stood, 5 un-homed use cases would have been
> `unimplementable` and the milestone's re-scope trigger would have fired. **The general lesson: a
> "no pre-X state exists" claim is usually a claim about the wrong column — check the writer, then check the
> column, before pricing a cluster as impossible.**

> The 2 org-admin `TODO`s (`org-admin.roles.UC1` + `org-admin.members.UC1`) were **diagnosed, not merely
> unbuilt** — their specs were parked in `playthroughs/e2e/drafts/*.spec.ts.draft` (the `.draft` suffix keeps
> Playwright from collecting them, so diagnosed work is preserved **without a red suite**) with the measured
> evidence in `e2e/drafts/README.md`. **Both landed** (iter-17 / iter-22) and the product is 4 of 4. The
> remaining 4 `TODO`s are all **onboarding**, each carrying a written verdict.

### The day-0 readiness seat — *a seat that could not be DECLARED, and the green that would have hidden it*

`onboarding.enterprise-workforce-ai-readiness.UC1` needs a member who **enters** her org's guided AI-readiness
flow rather than resuming it. Its blocker was recorded twice (iter-08, iter-18) as *"needs an Org C stage-0
seat"* — a seat merely **absent** from the seed, i.e. a YAML append. **Measured at iter-24, it was not
declarable at all.** `stack-seeding/seeders/ai_readiness_funnel.go:aiReadinessStageFor` mapped a hero by persona
kind — manager → stage 0, struggling → stage 1, **everything else → stage 3** — so any end-user hero appended to
the readiness org arrived having **already completed** the journey.

**The dangerous version of that gap is not a failure but a PASS.** A Playthrough written against such a seat
drives its hero *past* the flow it means to watch her enter, and can satisfy itself on the completed surface — a
green over onboarding that nobody onboarded. iter-24 held it with a failing-on-purpose test whose message
carried the discharge instruction; **iter-26 discharged it** with a seeder capability, not a YAML edit:

- **`blueprint.Persona.AIReadiness`** (`""` | `"not_started"`) declares the funnel stage **explicitly**, checked
  *before* the trajectory-derived default, and read in **exactly one place**. It is deliberately **not** a third
  `trajectory` value: readiness stage is orthogonal to the life-arc (the diagnostic is a time-boxed cycle, and a
  member who joined mid-cycle has skills, a career and an activity history while having done none of its steps),
  and a new trajectory would have rippled through the **five** seeders that switch on `Trajectory` via a silent
  `default:` branch in each. The value is a **closed enum** — a typo would fall back to stage 3, i.e. produce
  exactly the already-completed seat, so it fails the seed loudly instead.
- **`pt-ai-onboard`** is the resulting seat, **appended last** to Org C (hero indices are declaration order —
  `seat_append_test.go`), exposed as the `ai-readiness-dayzero-member` precondition.

**Why the seat is worth a Playthrough at all is a measurement**, and it is also the trap: the funnel **names all
three steps at every stage** (3/3 for the day-0 seat *and* for the started one), so an assertion on step names
is satisfied by a hero who has already finished the step. The **step CARDS** discriminate —

| locator | `pt-ai-onboard` (stage 0) | `pt-ai-started` (stage 1) |
|---|---|---|
| step NAME × 3 | 3 / 3 | 3 / 3 ← discriminates **nothing** |
| step-1 `Start · ~8 min` | 1 | 0 |
| step-1 `Done` badge | 0 | 1 |
| `Update Skill Mapping` | 0 | 1 |
| step-2 **CARD** (the step is locked) | 0 | 1 |

— so the Playthrough's pre-state IS its negative control, and its finals read legitimately FALSE for the resumed
member. Two more measured facts the spec depends on: the Step-1 wizard's forward control **relabels** (`Next` on
screens 1–4, **`Go to the AI Simulation`** on screen 5 — the iter-18 relabel trap in a second surface), and
**paging `Next` alone persists nothing** — the terminal control is the click that writes. It also *navigates the
browser off `/home`*, which is why the read-back must re-navigate rather than read the post-modal client state.

**P6 boundary, declared:** steps 2 and 3 are a real ~30-minute AI simulation and a live AI interview, so "all
three steps" is not drivable inside a Playthrough budget. The use case is scoped to the flow being **entered from
day-0** and its **first step completed and persisted** with the next step observably unlocked — the same
discipline as `pt-aisim-chat-launch` (stops at `/start`).

### The hiring-org day-0 candidate — *the routing claim we deliberately do not make*

`onboarding.enterprise-hiring.UC1` is the only onboarding use case whose curated final spans two apps, and for
two releases its manifest note carried a warning: a final observing *"the member is in the hiring app"* might be
proving the **cockpit's** routing rather than **onboarding's**, and the discriminator had to be found before the
use case could be honestly asserted.

**iter-27 found it by reading the source, and the answer is worse than the warning.** The eject is
`apps/web/src/context/UserStatusContext.tsx:142-173`, and it fires on **`userHasAllHiringOrgs` alone** — there is
**no onboarding condition in it**. So the cross-app routing belongs to neither onboarding nor the cockpit: it is a
membership-shape redirect that happens on any `apps/web` page load, for a member who has onboarded or not.
Asserting it would put a green over a mechanism the use case is not about, so **the Playthrough does not assert
it** — and says so, in the spec header and in the manifest, rather than dropping the intermediate silently.

**What IS onboarding-owned is asserted.** The hiring app has its **own** onboarding route
(`apps/hiring/src/app/(authenticated)/(signedup)/onboarding`) whose `onClose` is `router.replace('/home')`. The
Playthrough drives the **hiring base** (3001+offset) — where the platform puts an all-hiring-org member anyway —
and proves the half onboarding performs.

**The seat needed no seeder change, and that is worth recording** because its routed blocker read as if it did.
Day-0 onboarding is the **default** (completion lives in `public.user_params.onboarding`, NULL for every seeded
user); an end-user hero in a hiring org is a **candidate**, not a member (`endUserHeroRole`, M224); and
`heroHiringStage` already pins a **struggling** candidate hero to `assignedOnly` — *"a pending assignment, not yet
on the scoreboard."* So `trajectory: struggling` on `pt-hiring-onboard` is load-bearing for its **hiring**
meaning, not its skills arc: a *thriving* candidate hero arrives **assessed**, having already taken the positions
— iter-26's stage-3 defect class in the hiring domain.

**Three measured NON-facts, recorded so nobody adds the obvious assertion and watches it fail on a working
product:**

1. Revisiting `/onboarding` in the **hiring** app after completing it does **not** redirect — it serves the flow
   again, unlike `apps/web`. `pt-onboarding-complete` proves persistence *by* that redirect; that read-back does
   not exist here.
2. **Her home is seed state.** The greeting, the tenant chrome and the assigned position are all present
   *before* she onboards — measured as a mutant that **skipped the write entirely and still passed**. The fix
   was to delete the spec's manual navigation and let the app's own `router.replace('/home')` be the
   observation; that one line is now the only thing the write can satisfy.
3. The surface does **not** distinguish taken from not-taken (an already-assessed candidate renders the identical
   startable link). The final asserts the **affordance** — her position, in her tenant, with a real title; the
   "not yet taken" half is a **seed guarantee**, true by construction and not a claim this surface supports.

Point 2 is the general lesson and it is not specific to hiring: **on a seeded world, "the outcome is present
after the action" is not evidence unless it was absent before.** The Playthrough that reads a page the write
never touched is the same green-but-wrong shape as an assertion satisfied by an empty state, arriving through the
seed instead of through the locator.

### A seeded hero is part of the TEST SUITE's contract — twice paid for

M256's iter-13/14 deliberately re-aimed the negative controls at seeded facts **by name**: the org's email
domain, its member magnitude, a hero's name, a hero's **role**. That is what replaced the vacuous structural
finals, and it worked. The price only became visible when the milestone started *adding* heroes:

- **iter-26** gave a new Org C seat the role its COMPLETED hero already held. That made the role a two-member
  role, and the org's succession **key-role card** for a two-member role turns out to be **non-deterministic**
  (4 of 5 page loads at occupancy 2 against 5 of 5 at occupancy 1). The casualty was the cross-tenant control's
  own **liveness floor**, reading as *"succession failed to compute for the contrast tenant"* — RED in 2 of 6
  runs. A 45 s timeout was tried first, on a diagnosis of "host stall", and failed too. **Fenced:** hero roles
  must be pairwise distinct within a story (`seed-facts-fence.unit.spec.ts`, self-test injects a collision).
- **iter-28** appended a seat to Org A. `pt-workforce-funnel` went RED in **all three** runs: its sharpened
  final asserts **Pat Ellis's member-spotlight card carries her seeded role**, and one extra Org A member
  displaced Pat from the spotlight entirely. Deterministic, not flaky.

**The rule, and it is cheap to follow: before appending a hero, check whether `e2e/lib/seed-facts.ts` names her
org.** `SEEDED_ORGS = [PT_ORG_A, PT_ORG_C, PT_ORG_D]` — so **Org B (`pt-halcyon-retail`) is the one pt org
nothing anchors on**, and it is the default host for a seat that exists to serve one Playthrough. Both
regressions were caught by the 3× gate and in both cases the right fix was **the seat, not the assertion**: the
assertions were sharp on purpose.

### The org-prepared onboarding variant — *when a probe sweep samples a constant*

`onboarding.enterprise-workforce-standard.UC2` needs a member whose org has already imported her profile, so
onboarding opens on a prepared summary instead of an empty import form. Its trigger was *"not yet identified"*
for two iters: iter-08 measured a hero with a populated profile, iter-18 measured heroes across **four** orgs
(A, C, D) — every one served the identical import step.

**It is one `useState` in the component both apps mount** (`packages/ui/src/Onboarding/OnboardingUser.tsx:135`):

```ts
const lastStep = reimport ? Import : steps?.[steps.length - 1]?.step;
const [managerImport] = useState(
  Boolean(lastStep === OnboardingStep.Import && organizationName && userStats));
```

`organizationName` and `userStats` are always supplied by the host page, so the only missing input is `steps` —
i.e. **`public.user_params.onboarding`, NULL for every seeded user.** **Which is exactly why probing could not
find it: every seat in the world has the same value for that column, so a four-org sweep was sampling a
constant.** The general rule — *when a probe sweep returns the same answer for every vantage, the input is not
one of the axes you are varying, and more vantages will not help* — is the sibling of the routing finding
above, from the other side.

Seeded by **`onboarding: org_prepared`** (`blueprint.Persona.Onboarding` + `OnboardingParamsSeeder`): one jsonb
row whose last step is `import` — **not** `done`, which would complete onboarding and redirect. Two things
worth carrying from building it:

1. **The insert alone silently did nothing.** `public.user_params` is populated **row-per-user at user-insert
   time** (191 rows within 300 ms of the users COPY, all NULL, written by nothing in the seeding fleet), so
   `ON CONFLICT (id) DO NOTHING` skipped the row with no error and the seat kept getting the plain import form.
   The seeder now **inserts-then-heals** and **fails the seed** if a declaring hero's row cannot be reached —
   because that no-op presents as *"the product does not show the prepared summary."*
2. **The missing `audit.Record` was caught by the isolation guard** on the first live run (*"surface reports 1
   row written but recorded NO audit entry"*). Two guards fired while building this and both were right.

**The journey, measured (iter-29):** `/onboarding` serves the prepared summary and **no import form at all**;
the relabelled **`Start`** control **confirms the role and advances** — `user_params.onboarding` goes
`[{import}]` → `[{import},{role}]` — landing on the skills screen with the declared role's **real taxonomy
skills** to keep or discard. `pt-onboarding-org-prepared` asserts exactly that, with the import form's absence
as a cross-vantage control against every other seat in the world.

### The org-less seat — *the actor the seeder could not express* (v2.8 M256 iter-32)

`onboarding.individual.UC1` — *"a solo user with no organization completes first-run setup"* — was the **last
un-homed use case in the M201 curated corpus**, and the one where the milestone's pre-flight audit had a
kernel of truth: `UsersSeeder` wrote a `public.memberships` row for **every** seeded user unconditionally, so
a genuinely org-less actor could not be expressed at all.

**Price a capability by DELETING, not by building.** iter-30 answered the load-bearing question — *can the app
even serve a user with no organization?* — by deleting a seeded hero's membership row on a demo (restored by
`--reset`). She logs in, `/onboarding` serves the flow, `/home` renders. Five minutes, against a question that
had been open for twenty-two iters, and the FK error the delete raised **named its own first four consumers**.

**But a capability has as many halves as the platform reads.** That delete left Clerk still carrying her
organization, so what it measured was *a user whose DB membership vanished*, not *a user who never had one*.
The seeded state needs both: no membership row **and** no Clerk org claim in the exported roster — because the
host page mounts `<OnboardingIndividual>` versus `<OnboardingUser>` off `useGetClerkOrganization`. A DB-only
org-less hero is served the ENTERPRISE flow, and a Playthrough written for the individual journey would
silently prove the ordinary one.

**Make the blank state REQUIRED, not permitted.** The capability (`org_membership: none`) also *inverts* the
end-user `verified > 0` rule: an org-less hero **must** declare `verified: 0`. A verified skill's fan-out is
org-scoped, so a "verified org-less hero" writes rows tying her to an org she is not in — and requiring zero
makes the whole persona/profile/activity chain skip her **because she declares nothing**, instead of because
every seeder learned a special case. Likewise a manager is refused outright: the manager vantage IS the
org-intelligence seat.

**The FK list is the entry fee; the org-scoped tail is the work — and they need DIFFERENT guarantees.**

| half | how it fails | how it is caught |
|---|---|---|
| **loud** — a row keyed on a membership that does not exist | the seed **stops** and names the constraint | a **source-scan fence**: membership ids are a deterministic `membershipUUID(prefix, i)`, so *the call sites ARE the FK surface* — enumerate them and require each to consult the org-less predicate |
| **quiet** — an `organization_id`-bearing row for a user with no org | nothing at all; it is simply wrong | **measurement**: reseed, then sweep the live DB for her uuid across every uuid column |

Both arrived inside one iter. The loud one was the succession seeder FK-ing a population session she no longer
had. The quiet one was two `public.job_simulation_sessions` rows carrying an `organization_id`, plus activity events,
skill-path sessions, assignments and bookmarks that made a *day-0* user look like a returning one. **Say which
half a fence covers.** A static scan cannot see an `organization_id` write, and a fence that implied otherwise
would be more dangerous than none — the half it misses is the half that fails silently.

**When a fence suggests a remedy, the remedy is still a judgement.** Two pre-existing fences refused this seat
before it could ship: the curated-pool fence rejected its first role (no curated family → the taxonomy's
alphabetical junk head), and the ladder-depth fence then rejected the replacement because a **65**-skill
claimed tail would have drawn that family dry. The fence advised growing the allow-list. The real finding was
that **an org-less day-0 user should claim nothing** — a hero with zero verified skills and 65 self-rated ones
is not a coherent person. The footprint got smaller instead of the allow-list getting bigger.

**For an irreversible write, reset before EVERY mutant.** Onboarding cannot be undone through the UI. Run
against a world the green drive had already consumed, one mutant went red at the wrong line and another
**PASSED** — a false pass indistinguishable from a weak assertion, caused entirely by state. The protocol is
**reset → mutate → run**, every time.

### An absence assertion needs a companion that proves WHEN it was read — not only WHERE

This is the **third** variation on one defect in this milestone, and the third confounder is the surprising one:

| iter | the absence assertion was satisfied by | the confounder |
|---|---|---|
| **12** | a **dead page** (an ablated response: `bodyLen` 2147 → 24, 0 buttons) — which satisfies *every* absence | the page is broken |
| **22** | a table reading *"No roles match your filters."* — the loading row **is** the empty row | the page is empty |
| **29** | a page that **had not hydrated yet** — `toHaveCount(0)` immediately after a navigation | the page is **not there yet** |

iter-29's case is the one most likely to recur, because the assertion *looked* like a textbook server-side
read-back on a fresh navigation and its reasoning was **correct**: once a `role` step is persisted,
`managerImport` cannot be true, so the route cannot re-open on the prepared summary. It was still worthless —
iter-27's standing **Q1 mutant** (*delete the action and see whether anything fails*) left only that read-back
and it **PASSED**.

**It was removed, not weakened.** An absence assertion a mid-hydration page satisfies is worse than no
assertion, because it reads as proof. The honest repair is a POSITIVE locator on the screen the reload actually
lands on — and when nobody has driven that screen, the right move is to route the half, not to assert it.

> **The rule, in one line: after a navigation, prove the page ARRIVED before you prove anything is missing from
> it.** iter-12's liveness-before-absence has a temporal reading, and this is it.

### The `ai-readiness` product (M219) — and why a *blind area* is the worst kind of gap

Until M219 the AI-readiness diagnostic — a shipped product, seeded into the demo since v1.10b — was covered by
**nothing**, on **either** vantage. No Playthrough, no coverage descriptor for the member half. What that
bought, in the demo the team was presenting:

- the **STARTED** hero — *the entire point of the persona* — rendered **no readiness surface at all**. The
  member funnel is gated on a `deadline`; the backend derives one **only from an ACTIVE cycle**; the seed wrote
  only a **closed** one; `AIReadinessHero` returns `null` without a deadline. Nothing failed. Nothing asserted.
- the **COMPLETED** hero silently degraded from her full result hero to a compact archived rail-card.
- **six** manager sub-sections (the whole Step-3 interview-findings block and the four under it, the per-person
  *Recommended actions*, the *Assessment sources*) were **absent from the page** — the frozen (closed-cycle)
  read returns them as `null`.
- and every demo pointer (cockpit `jump_to`, deep-link catalog, coverage manifest) resolved to
  `/enterprise/workforce/ai-readiness` — the **pre-v3.0 LEGACY orphan**: no nav entry, no tab, no redirect,
  reading a cycle-less endpoint with no cycles, no archetypes and no people. The sweep asserted it for four
  releases and **passed**, because the page *does* render. It renders the dashboard the product no longer ships.

**A surface that renders is not the same as the RIGHT surface** — and that distinction is only visible to a test
that names the route. The four Playthroughs cover both vantages and both cycle states:

| Playthrough | Hero (seat) | Surface | What it proves |
|---|---|---|---|
| `pt-aireadiness-member-done` | `pt-ai-completed` | **`/home`** | the COMPLETED member's result renders — her score + a recap of all 3 steps. Anchored on the mode-`done` title, so the silent archived-rail-card degradation is a **red test**, not a shrug. |
| `pt-aireadiness-member-progress` | `pt-ai-started` | **`/home`** | the STARTED member's in-progress funnel renders — the 3 steps + the cycle due-date. **This is the surface that rendered as literally nothing.** |
| `pt-aireadiness-manager-dashboard` | `pt-ai-manager` | **`/ai-readiness`** | the org score, the Knowledge × Usage archetype matrix, the team breakdown, and a **resolved** cycle (not the "no cycles yet" zero-state) — **and** that the manager is NOT on the legacy orphan. |
| `pt-aireadiness-manager-howwemeasure` | `pt-ai-manager` | **`/ai-readiness`** | the 3-step method **and** the Step-3 AI-interview **findings** — the blocks a frozen read returns as `null` and the page therefore omits entirely. |

> **The member surface has NO ROUTE OF ITS OWN.** `AIReadinessHero` + `AIReadinessRailCard` are **embedded in
> `/home`**. That is why route-crawling never found them, and it is the single fact any future work on this
> product must start from. The manager dashboard is **`/ai-readiness`** (the only readiness route the navbar
> links). `url-shapes.ts` carries both as patterns — `AI_READINESS_URL` is **origin-anchored** (`://host/ai-readiness`)
> precisely so it **refuses** the legacy `…/workforce/ai-readiness`; `LEGACY_AI_READINESS_URL` exists so a
> Playthrough can assert the manager did **not** land there.

### The `hiring` product (M225) — the recruiter journey, on a SECOND app

**M225 landed `hiring.yaml`** — the recruiter-vantage candidate comparison, the FOURTH product and the first whose
surface lives **in a different app**. The one Playthrough proves the recruiter journey end-to-end:

| Playthrough | Hero (seat) | Surface | What it proves |
|---|---|---|---|
| `pt-hiring-recruiter-compare` | `pt-recruiter` | **apps/hiring `/enterprise/activity-dashboard`** | login → the Results scoreboard renders the org's **shared positions** with a real, comparable **candidate cohort** (not an empty grid), and the org reads as **HIRING** (the "Results" re-skin), never the workforce "Activity" view. |

> **The recruiter surface is `apps/hiring`, not next-web (the M224 two-app demo).** `apps/web` **ejects** an
> all-hiring-orgs recruiter to the hiring app by design (`UserStatusContext`); the hiring app's symmetric guard
> keeps her in. So the recruiter Playthrough drives **`env.hiringAppBaseUrl`** (offset **3001**-port,
> `PT_HIRING_BASE_URL` override; `run-playthroughs.sh` exports it), never `appBaseUrl`. `HiringResultsPage`
> **reuses** the M224 render-probe's calibrated **tanstack-table** anchor (`tbody.tbody > tr.tr` — the scoreboard
> is a custom react-table, NOT AntD; `.ant-table-tbody > tr` matched zero). An **empty scoreboard is a FAILURE**
> (a cold snapshot cache / starved `SIMULATION_TYPE_HIRING` pool leaves `readHiringSimPool` empty — the same
> silent-failure the M225 autoverify hiring cheap-win fences at bring-up), never a pass.
>
> **Scope: recruiter only** (one GREEN Playthrough = the milestone gate). The candidate is "optional" per the
> milestone and is covered on the **presence** side by S2's candidate coverage manifests — the clean pillar
> split (`coverage-protocol.md` = presence; this doc = function).

### The assign-WRITE Playthrough (M243) — the first MUTATING manager journey, proven to LAND

**M243 (v2.6 "sound check") landed the sole remaining in-manifest `TODO`** — the WRITE half of the
assign-and-track story. It is the FIRST Playthrough whose action-under-test **mutates real state**, so it is
where the release's anti-toothlessness thesis is sharpest: a test that merely closes a modal proves nothing;
the assignment must be shown to actually **LAND**.

| Playthrough | Hero (seat) | Surface | What it proves |
|---|---|---|---|
| `pt-assignment-assign` | `pt-manager` | **`/enterprise/assignments` (Skill Paths tab)** | login → the manager assigns a skill path to a member with a deadline → the assignment is **written and read back**: the target member's inline "Assign Skill Path" affordance FLIPS to the assigned title, so the assignable-affordance count drops by exactly ONE. |

> **The read-back IS the proof (the anti-toothless bar).** The final assertion is the affordance-count delta,
> not a closed modal. The members table query is keyed `['assignments', …]` and the org-assign mutation
> (`app.createOrganizationAssignments` → `public.organization_assignments`) invalidates `['assignments']`, so the
> table **refetches from the backend** and the target member's cell flips from "Assign Skill Path" to the assigned
> title. That count can only drop if a real `organization_assignments` row landed AND is read back through the
> real members query — a write that silently failed leaves the modal open (a red `confirmAssign` hidden-wait) or
> the count unchanged (a red poll). **No new seed data was needed:** Org A (Meridian Labs, 40 members) pre-assigns
> skill paths to only a handful, so ~34 members are deterministic assign TARGETS; the backend **refuses a
> duplicate active assignment**, which is exactly why the target must be an unassigned member. The precondition is
> DECLARED + enforced — UC1 names `seed.preconditions: [public-catalog, org-unassigned-member]`, the latter added
> to `seed-worlds.yaml`'s pt-world capabilities in lockstep (so a future "assign to everyone" seed change trips
> `ptvalidate`, not a mystery SETUP failure).
>
> **antd-v6 Select lesson (page-object layer).** The catalog picker is an antd `rc-virtual-list` Select whose
> `role="option"` nodes carry the raw VALUE (a uuid) as their accessible name and are treated as **non-visible**
> by Playwright (the visible title/image render in separate child nodes). `getByRole('option').click()` is
> therefore unreliable; the page object commits the first real option by **keyboard** (`ArrowDown`+`Enter`) —
> robust to the virtual list and genuinely user-driven (P1). Recorded for any future antd-Select surface.

### The `studio` product (M252) — studio-desk's FIRST manifest entry, the builder GENERATE

**M252 (v2.7 "july jitter") added `studio-builders.yaml`** — Product **"Studio"**, the FIRST time **studio-desk**
enters the Playthroughs manifest. The surface is the demo's own studio-desk (`9000+offset`), and both journeys are
driven by the **org-admin manager hero** (`pt-manager`), who alone clears the studio role gate. The two
Playthroughs prove the two builder GENERATE flows reach their **completion boundary** — a real result rendered, no
`500`:

| Playthrough | Hero (seat) | Surface | What it proves |
|---|---|---|---|
| `pt-studio-advanced-generate` | `pt-manager` | **studio-desk `sim-advanced-builder`** | login → the **advanced** builder GENERATE runs to its completion boundary — the generated result renders, `POST /api/ai/completion` returns (no 500). |
| `pt-studio-guided-generate` | `pt-manager` | **studio-desk `sim-guided-builder`** | login → the **guided** (interview-flow) builder GENERATE runs to its completion boundary — the same completion assertion. |

> **GENERATE is a real `/api/ai/completion` call — a live LLM at the assertion boundary.** So, per **P2**
> (functional truth, not pixel truth) and the integration-dependent assertion-boundary rule (§ "The iteration
> protocol"), the Playthrough asserts the flow **reached completion** (result rendered / no 500), never the
> model's generated *text*. Two M252 facts make this possible on a demo: (1) the studio backend actually holds an
> AI provider key — the studio-desk clone's `.env` is now wired into the container via the injected-override
> `env_file` (see [`../../services/studio-desk.md`](../../services/studio-desk.md) § Demo AI wiring +
> [`frontend-tier.md`](frontend-tier.md)); and (2) the studio surface is reached by a
> **Clerkenstein-authenticated** hero — the **org-admin manager** logs in via the cockpit handshake and passes
> studio-desk's Studio-role gate (the studio is **not** auth-disarmed; there is **no** `MOCK_CLERK`). ⚠️ *The gate is no longer `checkEnterpriseAndAdmin` in `src/index.ts` — that file is deleted. Since the Next migration it is **edge middleware** in `proxy.ts`, default-deny, reading the server-side BAPI; the role set is unchanged.*
>
> **Surface base URL.** The studio Playthroughs drive **`env.studioBaseUrl`** — studio-desk's single-port
> `9000+offset` (`PT_STUDIO_BASE_URL` override; the runner exports it), never `appBaseUrl`. Same pattern as the
> M225 recruiter's `hiringAppBaseUrl` — a product on its own app/port gets its own base URL, single-sourced in
> the env layer.

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
  **⚠️ The old exemption here is RETRACTED (M257x iter-38/39).** This principle used to carve out simulation
  scores as exactly-assertable on the ground that *"scoring is deterministic rubric-based, NOT AI-scored."*
  That is a **conjunction, and only one conjunct holds**: the rubric *arithmetic* is deterministic, but
  **most of the per-check verdicts it aggregates are LLM-produced** — a model is asked whether each check is
  met. (*Most*, not all: deterministic `EngineTextDiff` checks are the minority exception, and "all verdicts
  are AI" is the opposite error.) So a simulation **score is not exactly assertable** in general: assert
  *range* / *presence* / *monotonicity*, exactly as for any other LLM-derived value, unless the sim under
  test is provably all-`EngineTextDiff`. See
  [`../../architecture/ai_architecture.md`](../../architecture/ai_architecture.md) § Evaluation System.
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

  > **P6 can forbid a journey that demonstrably WORKS — and the test to apply is misattribution, not novelty**
  > (v2.8 M256 iter-18). Onboarding's self-import use case was recorded as blocked on a missing résumé fixture.
  > Driven, the blocker was false: the LinkedIn source needs no fixture and the import genuinely completes on a
  > demo in ~15 s, populating a real career profile — and it even arrives with a same-surface negative control
  > (a non-resolving profile URL reaches the identical step and the forward control never enables). It was still
  > **refused**, because what makes it green is a scrape of a live third-party site that blocks automation. The
  > day that site says no, the Playthrough goes RED reading like a product regression. **That is misattribution
  > — the same defect class [`seed-facts-fence.unit.spec.ts`](../../../.agentspace/rosetta-extensions/playthroughs/e2e/tests/seed-facts-fence.unit.spec.ts)
  > exists to prevent, sourced from outside the building.**
  >
  > The distinction that decides it is **not** "does it touch the network" — the suite already runs a live-LLM
  > studio lane at a 300 s budget. It is **whose refusal produces the RED**: a metered API called with our own
  > credentials, designed to be called, is a dependency; a site that actively refuses automated clients is a
  > coin-flip wearing a dependency's clothes. When a journey is only reachable through the second kind, the
  > deliverable is a **verdict built on the measurement** — plus the working journey preserved as a
  > `.draft` so the next attempt starts from evidence rather than re-deriving it — not a shipped test.
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
   **The tag grammar is `@pt:([a-z0-9][a-z0-9._-]*)` and it lives in TWO places by necessity** — `ptvalidate`'s
   `discover.go` and `report/playwright.go` each carry a copy, because the Go sections don't import each other.
   A **twin lockstep test** in each package (`cmd/ptvalidate/pttag_lockstep_test.go` and
   `report/pttag_lockstep_test.go`) pins both copies to one canonical literal + a shared match corpus, so
   *change one → change both* is enforced rather than commented (M203 close TEST-G1). **Any new greppable
   per-spec tag must follow this pattern** — one canonical literal, fenced from both sides — not a third
   unfenced regex.
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
  (`PageObject`): the shared semantic-locator primitives (`goto`, `main`, `byRole`, `byText`) that enforce the
  discipline in one place — every shipped surface scopes to `main()` then disambiguates by visible text.
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
- **`main()` is not universal — scope to the surface, not reflexively to `<main>` (M204 iter-03 D2).** Most
  surfaces (profile, roster) render their content inside `<main>`, so `main()` is the right outer scope. But some
  do **not**: the activity-dashboard **drill-down detail** renders its per-member results table in a plain-div
  layout *outside* `<main>` (the page even carries two `<main>` elements; `table.closest('main')` is false). There,
  scoping to `main()` finds the wrong/empty region — the correct anchor is a **page-level** table locator
  disambiguated by the surface itself (we're on the segment-anchored drill-down route, under the "Simulation
  Results" heading, and the detail carries exactly one `<table>`). Still within §5.2 ("scope within a named
  *surface*, disambiguate by a visible landmark") — the discipline is surface-scoping, and `<main>` is only the
  most common surface, not the only one.
- **A control can RELABEL, so name the accessor for the INTENT, not the label you last saw** (v2.8 M256 iter-18).
  Onboarding's import step has one forward button that reads **`Next` and is DISABLED** while no import source is
  supplied and becomes **`Import` and ENABLED** the moment a URL is typed. A `/^Next$/` locator therefore matches
  an element that is *permanently disabled* on that path, so any wait on it can only time out — and the timeout
  says nothing about the label. iter-18 concluded *"Next is disabled, the import path is not drivable"* across
  **four consecutive probe passes**, one of which sat on it for **6.9 minutes**, before dumping every button's
  label and finding `"Import"|en` sitting right there. The fix is an intent-named accessor spanning both states
  (`forwardControl()` → `/^(Next|Import)$/`) with the label-specific ones kept for assertions that deliberately
  name one state; the fence is
  [`onboarding-locators.unit.spec.ts`](../../../.agentspace/rosetta-extensions/playthroughs/e2e/tests/onboarding-locators.unit.spec.ts),
  which **captures the shipped matcher** and executes it against both strings, so re-narrowing it goes RED.
  **The generalisation: when a wait on a control times out, dump every candidate's label before concluding the
  path is closed** — the sibling of iter-17's *"when a wall has been measured four times, check whether every
  attempt shared an assumption"* (there, four pointer attempts; here, four passes reading one label).
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
- **M204 adds the manager-journey surfaces** (the additive merge with M203 — each vantage adds its own page
  objects, no collision): `workforce-page.ts` (`WorkforcePage` — the WI SPA funnel + org-scale gap),
  `members-page.ts` (`MembersPage` — the roster), `activity-dashboard-page.ts` (`ActivityDashboardPage` — the
  per-content activity aggregates + the per-member drill-down), and `succession-page.ts` (`SuccessionPage` — the
  succession / at-risk / mobility route). Their `/enterprise/*` route shapes are single-sourced in `url-shapes.ts`
  under the same anchored-segment discipline (`WORKFORCE_URL`, `MEMBERS_URL`, `ACTIVITY_DASHBOARD_URL`,
  `ACTIVITY_DRILLDOWN_URL`, `SUCCESSION_URL` — each with a symmetric `isOn*`/`isIn*` predicate pinned by the
  single-source-agreement block). All four extend `PageObject` and use only find-only landmarks (`<main>`,
  headings, visible stat labels, scoped `svg`/`table tbody tr`), identical in shape to the M203 trio.
- **M243 adds the assign-WRITE surface**: `assignments-page.ts` (`AssignmentsPage` — the `/enterprise/assignments`
  Skill-Paths assign builder), with `ASSIGNMENTS_URL` + `isOnAssignments` single-sourced in `url-shapes.ts` under
  the same anchored-segment discipline. It carries the only MUTATING page-object methods so far (open the
  builder → keyboard-pick a catalog skill path → confirm) plus the read-back accessor (the assignable-affordance
  count). Same find-only landmark discipline: `<main>`, the "Assign Skill Paths" heading, the antd
  `dialog`/`combobox` roles, visible button text ("Assign Skill Path" / "Assign"), and a text-filtered
  `table tbody tr` row — no CSS/testid.

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
seeds **four orgs** distinct from the demo showcase orgs, spanning entitlement tiers +
multi-org-private content — Org A (the enterprise employee + manager), Org B (the free-tier entitlement actor),
from **M219** **Org C** (`narrative: ai-readiness`, size 40 — the AI-readiness diagnostic org with a
COMPLETED member, a STARTED member, and its manager), and from **M225** **Org D** ("Kestrel Hiring Group",
`narrative: hiring` + `is_hiring`, size 40 → 4 admin + 36 candidates — the recruiter comparison org, distinct
from the demo's "Meridian Talent" AND from this world's Org A "Meridian Labs" so the two worlds stay cleanly
separable). The `seed-worlds.yaml` index
([`seed/seed-worlds.yaml`](../../../.agentspace/rosetta-extensions/playthroughs/seed/seed-worlds.yaml)) is
**single-sourced with the preset** — every world id / roster seat / tier / capability the validator resolves
against is materialized by the seed. It is covered by the **same datadna conformance gate** as the demo seed
(above).

> **⚠️ Three corrections to the paragraph above (v2.8 M256 pre-flight).**
> 1. **"spanning entitlement tiers" is a DECLARATION, not seeded state.** `seed-worlds.yaml` declares
>    `tiers: [anon, free, paying, enterprise, expired]` and the capability `entitlement-gated`, and annotates
>    the `pt-free` seat "*entitlement-gate use cases — outcome: blocked*" — but **no seeder writes a tier**
>    (see `actor.entitlement` above). The `pt-free` seat *is* seeded as a user; it is simply **not tier-gated**,
>    and it was referenced by **0** of the 18 use cases. **RESOLVED DIFFERENTLY at M256 iter-11:** `pt-free` now
>    drives two Playthroughs (onboarding, and the `blocked` refusal) and **its gate is real** — but the gate is
>    an **org feature grant** (`sim_feature_disabled: true` → no g3 casbin row), not a tier. The tier remains
>    declaration-only; the seat finally has an enforced refusal behind it. `ptvalidate`'s precondition-coverage check resolves
>    `entitlement-gated` against the declared list, so it **passes without the gate existing** — a fail-open
>    in the one check that is supposed to forbid a silent "ideally".
> 2. **There is NO pre-onboarding user state.** `UsersSeeder` writes a `public.memberships` row for **every**
>    seeded user unconditionally, and no onboarding flag/field exists anywhere in `stack-seeding/`. Every
>    `pt-world` actor is a *post*-onboarding org member. An onboarding journey therefore needs a **net-new
>    seed capability** (a seeder + a `capabilities:` entry + a roster seat), not just a new Playthrough.
> 3. **`--reset` is whole-stack, not org-scoped.** `stackseed --reset` (`cmd/stackseed/main.go` §`doReset`)
>    takes **no org filter**: it `TRUNCATE … CASCADE`s each of the ~28 `resetTables` — `public.organizations`
>    and `public.users` included — for that stack's Postgres, guarded only by `--stack` + the N=0 `--force`
>    rule. It does **not** spare the demo's showcase orgs. (`pt-world.seed.yaml`'s own header comment claims
>    the opposite — "*not touched by pt-world's reset*" — and is **wrong**; the claim in this doc's lifecycle
>    section, "full FK-ordered TRUNCATE, per-stack only", is the accurate one.) Practical consequence: a
>    `--reset` Playthrough run on a shared demo **destroys the showcase world**; re-run that demo's own preset
>    to get it back.

> **Layering finding (M202-D4).** Seeding `pt-world` onto an *already-seeded* demo-1 collided: the stories model
> forces the FIRST story onto `LegacyOrgID` (the Clerkenstein default org), which on a seeded demo IS the
> showcase's default org — so a pt-org merged into it and duplicate-keyed on the showcase's pre-existing
> `user_skills`. The **zero-platform-edit, zero-fork fix**: `pt-world` carries a leading **anchor story** (size 0,
> no heroes) that harmlessly re-declares the demo default org, pushing the real pt orgs to story index ≥1 so they
> get their own deterministic `StoryOrgID`s and never collide. This is a genuine seeding-machinery constraint for
> a *second world on a shared stack* (the demo default-org slot is single-tenant), recorded for the coverage
> milestones to inherit.

> **A world's shape is DECLARED, and the declaration is enforced (M219).** `seed-worlds.yaml` is not
> documentation — `ptvalidate`'s **precondition-coverage** check resolves every use case's `seed.world`,
> `actor.hero`, `actor.entitlement` and `seed.preconditions[]` against it and **hard-fails** on anything the
> seed does not provide. That is why the AI-readiness product had to land its **three artifacts in lockstep** —
> the `pt-world` Org C, the `seed-worlds` capabilities (`ai-readiness-org`, `ai-readiness-active-cycle`,
> `ai-readiness-completed-member`, `ai-readiness-started-member`), and the manifest. A partial landing is not a
> head start; it is a **broken validator**. The capabilities are deliberately *distinct* rather than one lumped
> `ai-readiness`: each one, absent, breaks a **different** Playthrough, and the whole point of
> precondition-coverage is that a missing precondition surfaces at **validate-time** instead of as a SETUP
> failure masquerading as a capability break.

## The BAKED-IN lifecycle — the suite now runs at the tail of every bring-up (v2.8 M258)

Until M258 the suite was something you *went and ran*. It is now the last thing a bring-up does:
`up-injected.sh` invokes [`e2e/batch-gate.sh`](../../../.agentspace/rosetta-extensions/playthroughs/e2e/batch-gate.sh)
right after the `UP.` line, so a single cold command brings the stack up **and proves every journey on
it** — the ambition being that this is the *normal* way to bring a stack up, not a ceremony reserved for
a release gate.

The batch-gate contract (`D-v28-3`) and its **four** fail-closed ledger rules are documented once, in
[`../verification.md`](../verification.md) § *The layer ABOVE autoverify* — read it there rather than
here. What matters at **this** layer is what the gate does to the world the suite runs in:

- **The suite is driven UNSCOPED and un-retried**, so the run is **binding** (a `--grep`'d run is
  advisory — see the SCOPED split at the bottom of `run-playthroughs.sh`).
- **The reset that this suite needs is destructive to the presenter demo**, and the gate therefore owns
  the restore leg: `restore-presenter-world.sh` puts the stories world, the Clerkenstein roster and the
  cockpit/content manifests back, on **every** path where the reset ran — a red batch must not *also*
  cost the presenter the demo world. It ends in a **post-condition** that cross-checks the restored menu
  against the identities that can serve it (`check-cockpit-roster.py`), because each layer reporting its
  own success is exactly the state in which the original defect shipped: a stories roster beside a
  pt-world menu, every export `ok`, exit 0.
- **The restore's preset comes from the clone that OWNS the live stack**, and `DEMO_STORIES_PRESET` wins
  over the default — so a bare `restore-presenter-world.sh N` restores what the bring-up actually seeded.
  This is the one substitution the post-condition **cannot** catch, since the roster and the menu are both
  exported from the same preset and a wrong one yields a wrong-but-self-consistent pair.
- **The batch is skipped, and recorded as `skipped` (never `green`), wherever a documented knob removes
  something it needs.** On a `--public-host` stack, because such a demo cannot be browsed from its own
  host — and since `--public-host` is default-on, a bare `/demo-up N` skips while `/demo-up N
  --no-public-host` gates. **Also on `DEMO_NO_STORIES=1` / `DEMO_STORIES=0`** (no heroes, no cockpit, and
  no per-stack `stackseed`) **and on `DEMO_NO_UI=1`** (no browser surface at all). Running on either is
  not a measurement: it is an error, or 30 false REDs describing the operator's own configuration.

## The lifecycle — reset-to-seed + the serial-default runner

P1 mandates the action-under-test **mutates real state**; P6 demands *same inputs → same result*. These hold
together only if the world is **reset to the known seed between runs** — so the runner
([`e2e/run-playthroughs.sh`](../../../.agentspace/rosetta-extensions/playthroughs/e2e/run-playthroughs.sh)) does
a **per-suite reset-to-seed** on `--reset`:

- It runs the **real `stackseed --reset` path** (full FK-ordered TRUNCATE, per-stack only, honoring stackseed's
  own N=0 guard) **then** a fresh seed of `pt-world.seed.yaml`. **Additive re-seed is FORBIDDEN as a reset** — an
  `ON CONFLICT DO NOTHING` re-seed silently leaves stale state (the M42e "green-but-wrong" trap). See
  [`idempotency.md`](../idempotency.md) + [`seeding-spec.md`](../seeding-spec.md) for the `--reset` contract.
- **It also refreshes the Clerkenstein roster + restarts the fake services (v2.1 M211 iter-16).** The world the
  cockpit seat-switch logs into is *DB + identities* — but the fake-FAPI/BAPI resolve identities from a mounted
  `/roster/roster.json` baked at **bring-up** from that demo's preset. A reset that only swaps the DB leaves a
  stale roster, so a hero login for a `pt-world` seat on a demo brought up for **something else** (e.g. a
  stories/coverage demo) `400`s with `unknown_identity` — the whole suite red. So `--reset` re-exports the roster
  from THIS seed (`stackseed --roster-export --seed pt-world` — a pure function of the seed, no DB) to the
  `docker inspect`-discovered mount path, restarts `demo-N-fake-{fapi,bapi}`, and waits for the FAPI. This
  **completes the reset-to-seed** so the Playthroughs run on **any** demo, not only a `pt-world`-native one
  (M204 masked this by bringing its demo up `pt-world`-native). Non-fatal for a roster-native demo; zero platform edits.
- **…and the COCKPIT MANIFEST too — the other half of "the world" (v2.8 M256 iter-10).** The roster refresh
  above fixes *who the fake FAPI can mint*; the cockpit's `[Log in as]` menu is a **separate** projection
  (`cockpit-manifest.json`, see [`cockpit-spec.md`](cockpit-spec.md)) and it was never re-exported — so after a
  `--reset` the cockpit offered seats from the *previous* world and every selection **silently** fell back to
  the last-active seat. 23 Playthroughs stayed green while the human-facing cockpit was entirely stale.
  `--reset` now also runs `stackseed --cockpit-export --seed pt-world` to the manifest beside the roster mount
  (same pure-function-of-the-seed property — no DB read). **A running cockpit holds its manifest in a closure**,
  so refreshing the file is not enough on its own: `cockpit.py`'s `--roster` cross-check makes a stale
  in-memory manifest **fail closed** rather than offer a dead seat, and the next cockpit start picks the
  refreshed file up. Non-fatal — a demo with no cockpit manifest beside the roster is a valid shape and is
  skipped with a stated reason.
- The runner **refuses N=0** (the main dev stack) outright — a Playthrough run always targets a demo-N.
- **Gate-run prereq — the pinned `stackseed` must be on PATH (M204 iter-05 D1).** The runner shells out to bare
  `stackseed` (the pinned tooling the demo consumes), which is **not on the login PATH**. When running the gate
  against a demo from its **consumption clone**, prepend that demo's `bin/` —
  `stack-demo/rosetta-extensions/demo-stack/stacks/demo-N/bin` — so `run-playthroughs.sh --reset` resolves the
  pinned `stackseed`/`stacksnap`. This is a **gate-run environment prereq, not a runner code change**: the runner
  correctly delegates to the pinned CLI rather than hard-coding a path. (Running from the authoring copy instead,
  the CLIs are `go run`-able in place.)
- **Serial by default.** The runtime is a single shared `organization_id`-scoped Postgres, so two mutating
  Playwright workers would interfere — and Playwright defaults to *parallel*. The config
  ([`e2e/playwright.config.ts`](../../../.agentspace/rosetta-extensions/playthroughs/e2e/playwright.config.ts))
  therefore pins **`fullyParallel: false`, `retries: 0`** (a retry that masks a flake hides a
  Playthrough defect) and resolves `workers` through `resolveWorkers()` (`e2e/lib/stack-env.ts` §resolveWorkers)
  — **default 1**, overridable only by a `PW_WORKERS` that is a positive integer (a `0`/negative/NaN override
  **fails loud** rather than silently going parallel). The sanctioned throughput-reclaim paths —
  **stack-per-worker** (a stack each) or per-worker org/hero partitions in the seed — are opt-in via
  `PW_WORKERS`, never the day-one default.

  > **⚠️ Postgres is NOT the binding shared surface — the fake-FAPI seat is (v2.8 M256 pre-flight).** The
  > paragraph above is the *original* rationale and it is incomplete in a way that matters the moment anyone
  > tries to reclaim throughput. **Clerkenstein holds ONE active seat for the whole stack**: a single registry
  > `activeKey` (`clerkenstein/clerk-frontend/registry.go` §`Registry.activeKey` / `active()` / `Select()`),
  > one `signedIn`, one `sessID` (`clerk-frontend/server.go` §`type Server`). Every Playthrough login routes
  > through `hero-login.ts` → the shared `stack-verify/e2e/lib/cockpit-login.ts` §`selectSeat` →
  > `POST /v1/demo/select` → `handleSelectIdentity`, which re-points the seat **and** sets
  > `s.signedIn = false; s.sessID = ""` **globally**. The read path (`handleMe`, `handleToken`, `handleClient`,
  > `handleMeOrganizationMemberships`) **discards the `*http.Request`** and answers from `activeUserLocked()`
  > with **no cookie or token input** — so **`storageState` reuse does not isolate two seats either**, and
  > `handleSignOut` is stack-wide. Consequences, both load-bearing: (1) **per-worker org/hero seed partitions
  > alone are NOT sufficient** — worker 2's login signs worker 1's browser out mid-journey and its `/v1/me`
  > 401s; only **stack-per-worker** (a fake FAPI each) is safe today. (2) Any *third* parallelism path must
  > first make the seat per-client (a cookie/`__client`-scoped registry or one fake FAPI per worker) — a
  > Clerkenstein auth-model change with an alignment-DNA consequence, not a config flip. Two in-repo comments
  > already record this verdict (`stack-verify/e2e/tests/m224-candidate-heroes.spec.ts` §serial-mode preamble;
  > `stack-verify/e2e/tests/content-stories.spec.ts` §"SERIAL BY NECESSITY"), and the same limitation is
  > disclosed from the presenter side in [`cockpit-spec.md`](cockpit-spec.md) § *Limitation — one seat per
  > stack*. The frozen `spec.md` v0.3 draft (§5.7) carries the Postgres-only rationale and is **superseded on
  > this point by this section**.
- **The runner reconciles inline** (M204 iter-02). After the Playwright run it invokes `ptreport` over the
  manifest + this run's fresh JSON results and prints the four-state map — so a single `run-playthroughs.sh`
  invocation both *runs* and *reconciles*. **The reconciliation is BINDING on a full run and ADVISORY on a
  scoped (`--grep`) run — v2.8 M256 harden-final.** It used to be unconditionally non-fatal (`|| echo`), with
  the runner exiting on Playwright's status alone — but `report.go`'s `!ok` branch is the ONLY mechanism that
  notices a **declared Playthrough that never ran**, and **Playwright exits 0 when a spec file is simply
  absent**, so deleting a spec produced a fully green run (measured: `203 passed`, rc 0). A full run now exits
  non-zero when the gate is unmet; a scoped run stays advisory with a stated reason, because every un-grepped id
  correctly reports *"did not run"*. **Anything downstream that ran the suite and trusted a zero exit is now
  genuinely gated.** Full account, plus the `set -e` trap the first version of the fix fell into:
  [`../verification.md`](../verification.md) § *A gate whose exit code is discarded is not a gate*.
  **Reporter-override lesson (load-bearing):** a
  Playwright CLI `--reporter=…` flag REPLACES the config's *entire* reporter list — so the runner must **not**
  pass one, or it silently suppresses the config's `['json', {outputFile: ./report/last-run.json}]` reporter,
  leaving `last-run.json` stale and decoupling `ptreport` from the actual run (a green-but-wrong-reconciliation
  trap). The config declares `['list', 'json', 'html']`; letting that set fire keeps the console `list` output
  AND refreshes the JSON `ptreport` reads. (This fixed a latent M202/M203 wiring defect too.)

```bash
cd playthroughs/e2e
./run-playthroughs.sh 1              # run the suite against demo-1 (serial), no reset; reconciles inline
./run-playthroughs.sh 1 --reset      # reset-to-seed the pt-world FIRST, then run + reconcile
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
| **`unimplemented`** | `[TODO]` | A declared use case with no Playthrough yet (`playthrough: TODO`) — **and, since v2.8 M256 iter-31, one that MUST carry a written `verdict` block saying which kind of gap it is**; the detail line is that verdict, not a generic sentence (see § below). |
| **`unimplementable-without-platform-edit`** | `[BLOCKED-PLATFORM]` | The surface cannot be driven without a platform edit (a hard zero-edit wall — e.g. a hardcoded URL with no override). It **escalates, it does not edit the platform** — the P3 escape valve, mirroring the coverage sweep's re-scope trigger. Declared deliberately (with a rationale) in [`report/unimplementable.yaml`](../../../.agentspace/rosetta-extensions/playthroughs/report/unimplementable.yaml), never inferred from a failure. |

The four glyphs are deliberately **visually distinct** — a `pending`-vs-`unimplemented` ambiguity would hide a
real semantic distinction. `Report.AllGreen()` (nothing failing/unimplementable/unimplemented) is the
foundation-complete gate; `Report.NoRegressions()` (nothing `failing`) is the gate a *coverage* milestone runs —
a build-reference `TODO` gap must not fail the suite. Coverage = passing ÷ total declared.

### A `TODO` must carry a WRITTEN VERDICT — and the fence runs both ways (v2.8 M256 iter-31)

`unimplemented` used to be one state wearing two completely different meanings, and the map could not tell them
apart. Every use case without a Playthrough reported the **same sentence** — *"declared use case, no
Playthrough yet (build-reference gap)"* — whether it was genuinely unbuilt or had been **measured and
deliberately refused**. For M256's self-import use case that sentence was simply **false**: its only advancing
path scrapes a live public third-party profile from a site that blocks automation, so shipping it would make a
real person's profile a permanent fixture and **its RED would read as a product regression when nothing about
the product had changed**. The reasoning existed — in prose, in a story `note`. The four-state map, which is
what tooling and reviewers actually read, stated the wrong position about the one use case a release gate was
being closed around.

So a use case with no Playthrough now carries a **`verdict`** block, and `ptreport` renders it in place of the
generic sentence:

```yaml
playthrough: TODO
verdict:
  disposition: will-not-build      # CLOSED enum: will-not-build | not-yet-built
  measured_by: "M256 iter-18 (six probe passes, live) + D104"
  rationale: >                     # >= 80 chars, no placeholder spellings
    MEASURED then deliberately refused — the only advancing path scrapes a live third-party profile …
```

**`will-not-build`** is a measured refusal and MUST NOT name a handler; **`not-yet-built`** is a real gap and
MUST name the routing handler that will close it. That asymmetry is what stops the two blurring into *"TODO
with a paragraph attached"*. It is also the mechanical form of iter-30's **D117** — *a routed blocker must
carry the measurement that produced it, or be marked an estimate* — which is what `measured_by` is for: three
of one session's five iters found a routed blocker mis-stated, each written in good faith and each read as a
measurement by the next iter.

**The state model is UNCHANGED.** A `will-not-build` is not a fifth state and gets no new glyph — it is an
`unimplemented` use case that now says which kind it is. A new glyph would imply a state the reconciler does
not have.

**Fence it in BOTH directions.** A `TODO` without a verdict fails; **a use case with a LIVE Playthrough that
still carries one fails too.** The second half is not symmetry for its own sake: a verdict left behind on a use
case that has since been proven is a stale claim with **no expiry** — nothing in the artifact tells a reader it
is out of date, and it goes on asserting a blocker that no longer exists. (Same shape as M255's knob guard: a
doc-promised flag with no parser entry is a *false promise*; a parser flag with no doc row is
*undiscoverable*.) Landing a use case must therefore **force** the verdict's removal rather than leave it to
diligence.

**And fence it against being VACUOUS, or the schema just re-hosts the silence.** A presence check is satisfied
by `rationale: TODO`. Hence the 80-character floor, the placeholder blacklist, and a **closed** disposition set
with **no fallback member** — this milestone twice shipped a seed enum whose unrecognised value fell back to a
permissive default, and in both cases the fallback would have produced a Playthrough that looked like a product
regression. Prove the fence on the **shipped** manifest, not only on fixtures: five green unit tests once drove
a mock path the real client never used (iter-16), and a fence proven only against fixtures is a fence proven
against itself.

### Liveness before absence — now machine-checked, because three iters fixed it by hand (v2.8 M256 iter-31)

After a navigation, an absence assertion is **not evidence** until something positive has been observed on that
page. `toHaveCount(0)` and `not.toBeVisible()` are satisfied by a page that is dead, empty, **or simply not
there yet**. The rule is old; what iter-31 added is that a test now holds it
(`playthroughs/e2e/tests/liveness-before-absence-fence.unit.spec.ts`), because the same defect arrived three
times in three costumes:

| iter | costume | what satisfied the absence |
|---|---|---|
| **07** | **dead** | an ablated GraphQL response — `bodyLen 24`, 0 nav, 0 buttons. The whole mechanism was refuted. |
| **22** | **empty** | the roles table's placeholder row — its LOADING row and its EMPTY row are the same `<tr>` carrying the same sentence. *In this app the empty state occupies a row.* |
| **29** | **not there YET** | a `domcontentloaded` navigation's pre-hydration DOM. Nothing was dead and nothing was empty — **TIME** was the confounder, which is exactly why the two earlier fixes did not prevent it. |

*An absence assertion needs a companion that proves **when** it was read, not only **where**.* The witness can
be a polled body-length/nav-chrome floor (`assertPageIsAlive`), an ordinary `toBeVisible()` on a landmark, a
non-zero count, or a `waitForURL` — anything a dead, empty or unhydrated page could not satisfy. Note what does
**not** count: `expect(landedUrl).not.toMatch(/\/login\b/)` reads a string the harness already returned and says
nothing about whether the page rendered.

**Measure a fence before adopting it** (iter-15 D74): the scan reported **29 files · 62 navigation sites · 184
liveness witnesses · 37 absence assertions · 0 violations**, so the invariant was already true everywhere and
cost **zero edits** — the fence buys the *next* spec, not this one. And make it **fail-closed** on floors for
each of those counts: a scan that matches nothing passes every assertion, and a fence is the worst possible
place to commit this milestone's signature defect.

> **⚠️ AND THEN THE FENCE ITSELF ENFORCED HALF OF WHAT IT ADVERTISED** (v2.8 M256 harden-final). `.toBeVisible(`
> is a **substring of** `not.toBeVisible(`, and the LIVENESS branch was evaluated before ABSENCE and returned —
> so a `not.toBeVisible()` line was scored as a positive **witness**, which *disarmed* the state machine and
> licensed every absence after it on that navigation. The `not\.toBeVisible\s*\(` alternative in the ABSENCE
> pattern was unreachable dead code, and *"0 violations across 29 specs"* was true only of the `toHaveCount(0)`
> spelling. Separately, `toHaveCount\(\s*0\s*\)` required the `)` immediately after the zero, so
> `toHaveCount(0, { timeout: … })` escaped entirely — while the comment above it claimed that spelling was
> *"covered by the `0` alternative's tolerant whitespace"*. **Whitespace tolerance is not option-object
> tolerance.**
>
> Widening to `[,)]` then drags in `not.toHaveCount(0, …)`, which is the **opposite** claim ("at least one") and
> is spelled that way at three live sites — so **`not.` is the discriminator in BOTH directions**, and both
> patterns carry a lookbehind. ABSENCE is now tested **first** (fail-closed: a line that could read either way is
> an absence).
>
> **Both defects were LATENT in the corpus, and that is the transferable lesson.** Neither spelling appeared in
> any spec, which is exactly why a fence that only ever runs over the corpus could not surface either one — it
> was validated by whatever the corpus happened to contain. **Latent is not fixed; it is one ordinary edit away
> from live.** A classifier now gets its own synthetic self-test alongside the corpus scan.

### A fence must scan the thing it is NAMED for (v2.8 M256 harden-final)

The bounded-interaction fence exists because of iter-06 **D25**, whose subject is `pickFirstSkillPath` in
`assignments-page.ts` — written `for (let attempt = 0; attempt < 3; attempt++)`. Its loop pattern matched only
`for(;;)` and `while(true)`, so **the fence built to prevent D25 recurring had never once looked at D25's own
loop**. It was satisfied instead by an unrelated `for(;;)` two hundred lines above, in a different method.

Its regression pin had the same disease one level up: it asserted a **per-file loop count `>= 1`**, which any
loop in the file satisfies — so the D25 retry could be deleted outright with the pin green. *A per-file count
cannot pin a per-method invariant.* The pin now names four loops **individually, by the method that owns each**,
and checks the loop is inside that method.

**The generalisable rule:** when a fence names a defect, make it assert that it is scanning **that defect's
actual site**. A "not vacuous" floor that counts *anything* is satisfied by the wrong thing, and reads as
coverage.

### The standing mutant question, asked of the OLD Playthroughs too (v2.8 M256 harden-final)

`PT-M256-standing-mutant-Q1` — *"delete the action; does anything fail?"* — had only ever been asked of the
Playthroughs each iter had just written. Harden-final asked it of three **older, never-mutated** mutating
Playthroughs, each on a **fresh reset-to-seed world** (mandatory: the write is irreversible, and iter-32 had two
mutant runs confounded because the previous run had already consumed it):

| Playthrough | action deleted | verdict |
|---|---|---|
| `pt-orgadmin-setting-toggle` | `settings.toggle(SETTING)` | **RED** — *"the switch flipped in the UI"* |
| `pt-skillpath-bookmark` | the save click | **RED** — *"the save PERSISTED: after a full reload…"* |
| `pt-assignment-assign` | `confirmAssign()` (the WRITE) | **RED** — *"the assignable-affordance count drops by exactly one"* |

3 of 3 red. **9 of the 12 mutating Playthroughs remain unasked** — named, not implied.

### The fifth outcome the map has no glyph for — a PRODUCT DEFECT the suite finds (v2.8 M256 iter-23)

The four states above all describe **the suite's relationship to a use case**. None of them describes the
thing a Playthrough is ultimately for: *the product is broken, and here is the evidence.* That outcome has no
row, no glyph and no ledger — so when M256 found one, it lived in a single paragraph attached to a use case
that then went green, one manifest edit away from being deleted with the comment it sat in.

**Record a product defect where it cannot be tidied away**: the milestone's `decisions.md` (which
`/developer-kit:close-milestone` reads and routes) **and** its `progress.md` routing table, never only in the
manifest comment beside the UC. A defect discovered while covering a use case outlives the coverage work.

**And capture it while you can still reproduce it.** M256's case is the pattern: the defect was only visible
because a demo was *missing* a grant; the fix that unblocked the use case also **removed the symptom from the
demo**. Reproducing it deliberately (revoke on the demo DB → observe → restore, verifying the restore) is a
sanctioned demo-DB write and is how the evidence gets taken before it is unreachable.

**When the finding is a NEGATIVE — "nothing is surfaced" — enumerate the channels.** A probe that checks one
place and finds nothing has not proven absence; it has failed to look. M256's capture enumerated the GraphQL
response body, the dialog state, `role=alert`, `role=status`, antd `message`/`notification`/form-explain, the
browser console, uncaught page errors, the URL, and the post-state — and the enumeration is what produced the
sharpest single fact: **the `role=alert` region was PRESENT and EMPTY.** A form with a mounted, empty error
slot is not a form with no error handling; it is a form whose error handling covers a different error, and
that distinction is the whole bug report.

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

**Read the WRITER, not the declaration, before an iter commits to a precondition (v2.8 M256 iter-01).**
`ptvalidate`'s precondition-coverage check resolves a use case's `seed.world` / `actor.hero` /
`actor.entitlement` / `seed.preconditions[]` against the **names** `seed-worlds.yaml` declares — so it can
only catch a precondition the *index* omits, never one the *seeder* never writes. `actor.entitlement` is the
worked example (above): declared in every world, materialized by nothing, and green in the validator. So when
an iter plans a Playthrough around a precondition, the evidence it needs is **the seeder line that writes the
column**, not the capability entry that names it. A capability with no writer is a fail-open, and it surfaces
as a Playthrough asserting a behaviour the platform has no reason to exhibit.

**A route that gates on state is a FREE read-back — look for one first (v2.8 M256 iter-08).** The cheapest
mutating proof shape in the suite is not a label flip or a list delta; it is a route that **serves-or-redirects**
on the very state the Playthrough writes. `/onboarding` SERVES the first-run flow while
`user_params.onboarding` is unset and **REDIRECTS to `/home`** once it is set, so one URL supplies *both* halves —
the pre-state absence (which is the negative control, per the pattern below) and the persisted post-state — each
read on a fresh navigation, with no second surface, no toast, and no DB assert. When such a route exists, prefer
it. Corollary that bit once: if you assert the same pattern in **both** directions, its **segment anchoring is
load-bearing** — an `/onboarding-tour` look-alike match would make the "it persisted" half pass on the wrong
route, so the pin test asserts the rejection explicitly.

**Choose the seat for a persistent mutation; do not inherit it (v2.8 M256 iter-08).** Completing onboarding
cannot be undone through the UI, so driving it on `pt-employee` would have coupled that Playthrough to every
other one asserting on that hero. `pt-free` was registered in `seed-worlds.yaml` and driven by **0** use cases —
the pre-flight audit had recorded that as a *gap*, and it turned out to be the asset that makes an irreversible
write safe. Before writing an irreversible mutation, ask which seat nothing else reads.

**A mutating Playthrough's PRE-STATE read is its negative control — for free (v2.8 M256 iter-06).** A negative
control is the demand that a Playthrough be *demonstrably RED when its outcome is absent*, and the instinct is to
reach for a second stack, a mock, or a DOM ablation. None of that is needed for a Playthrough that WRITES. Read
the target state **before** the action and make the final assertion a **strict inequality or a strict negation**
against that reading — `.toBe(!before)`, `.toBe(before - 1)`, `toHaveCount(0)` before / non-zero after,
`progressAfter > progressBefore`. Such a predicate is **false by construction at the pre-state**, so the run
itself demonstrates the assertion discriminates the outcome rather than matching page chrome — against *real*
product state, not a simulated absence, and inside the same run. Prefer the delta form over the absolute
("the label is not yet `Continue (N%)`"): the absolute form false-REDs on a re-run against a world that was not
reset, and **a false red is exactly as dishonest as a false green**. iter-06 found all three pre-existing mutating
Playthroughs already had this property, unnamed — so the pattern was ratified, not invented.

**A STRUCTURAL final has no contrast vantage — so sharpen the final, don't hunt for a vantage (v2.8 M256
iter-13).** A read-only Playthrough's negative control comes from a **contrast vantage** — a hero or org for
whom the asserted outcome legitimately does not exist. iter-12 measured that this mechanism **cannot** apply to
a *structural* final: "a stat label is visible", "a chart count ≥ 1", "a Work section exists" are all satisfied
for **every** seeded member, because M44's profile-completeness seeder gives every member a career and skills
(measured: the manager reads `verifiedSkillsStat` 1 / `skillCharts` 10 / `workSection` 1 on her own profile).
That is correct, and it means **no suppression switch can exist** — do not go looking for one.

But the limit is on the **assertion**, not on the Playthrough. Those finals were structural *because they were
written structurally*, and the same surfaces carry hero-specific facts the seed pins deterministically. Re-aim
the final at the hero's own seeded data and the contrast vantage that could not falsify the structural version
falsifies the specific one. iter-13's three worked examples, all measured on both vantages first:

- **A seeded magnitude.** The rendered `Verified Skills` stat equals the hero's seeded `skills.verified`
  *exactly*, and `All Skills` equals seeded `verified + mapped` *exactly* (confirmed on two heroes
  independently) — so the claimed-vs-verified GAP itself is assertable, where "two stat labels rendered" was not.
- **A seeded identity in context.** The `"<role> at <org>"` line names two authored seed fields, so it is false
  for a different person in the same org *and* for the same role in a different org.
- **A COMPUTED outcome that needs the seeded history.** The closest-role recommendation renders only once a
  hero has enough *verified* evidence for the matcher to produce candidates — present for an 8-verified hero,
  absent for 3- and 2-verified ones. That proves the surface computed something from *her* history, where a
  chart count proves only that a chart drew.

Keep the old structural assertions as **intermediates**: they establish that the surface is there; they never
established whose it is. And **machine-link every number to the seed** — a magic `8` in a spec is a claim about
the seed file with no link to it, so a renumbered seed turns three Playthroughs RED naming a product regression
that never happened. A declared facts module plus a fence that PARSES the seed and reconciles is the pattern
(`playthroughs/e2e/lib/seed-facts.ts` + `tests/seed-facts-fence.unit.spec.ts`); the fence's **first** assertion
must be that the parse is not vacuous, because a reconciliation over an empty parse passes every comparison
silently. This does not violate P2: authored seed literals under reset-to-seed do not "vary across captures" —
what P2 forbids is generated content (bios, generated employer history, computed match percentages), and none
of that is asserted.

**The contrast vantage follows the SUBJECT of the final, not the product (v2.8 M256 iter-14).** iter-13's
worked examples are all about a **person**, so their contrast is another person in the same org. Applied one
product over, that seat is *wrong*: the four Workforce-Intelligence finals are read BY the manager and are
about **her org's aggregates**, so she cannot falsify her own dashboard. An org-aggregate final needs a
manager of a **second seeded TENANT**. Name what the final is *about* before hunting a vantage — measured
first, this is one line; discovered afterwards it is a wasted iter.

The same widening applies to what a spec may name. Alongside the per-hero facts, pin the **org's** authored
identity (`org.name`, `org.slug`, `org.size`) and reconcile it with the same fence. Three properties earned
their place on measurement:

- **A per-row org anchor beats a single row.** A roster renders 20 of 40 members, so *any* one member's row
  is a bet on sort order; the **org email domain** (`@<org.slug>.com`) is on every org-member row and holds
  on any page. Bound it honestly — 15 of 20 rows carry it, the rest being `Candidate`-role members on
  external addresses — so assert *present among the rows*, never *all rows*.
- **Say whether a magnitude is a strengthening or a discriminator, at the assertion.** "Overall Members
  equals the seeded `org.size`" catches a stat card rendering a constant and an accessor reading the wrong
  card — real value a visible LABEL cannot give — but if both seeded orgs are `size: 40` it discriminates
  **nothing**. Write which it is, or a reader banks the wrong one.
- **Mutate the FINAL as well as the control.** Two mutant groups: re-aim each control's absence assertion at
  the contrast vantage's own data (proves the *control* can fire), **and** drive each sharpened Playthrough
  on the contrast vantage (proves the *final* discriminates). The second group is the one that matters — the
  first alone can demonstrate healthy controls over still-vacuous Playthroughs, which is a green test about
  a green test.

**When a Playthrough RESISTS a negative control, suspect the ASSERTION before you suspect the world (v2.8
M256 iter-19 — the third instance, so it is now the rule).** iter-12 recorded three profile finals as
uncontrollable and iter-13 refuted it; iter-14 took the same move up to the tenant; iter-19 closed the last
non-studio gap the same way. `pt-hiring-recruiter-compare` had been priced for four iters as *"needs a
same-vantage control whose absence half is unmeasured"* — and the real situation was that its final,
`positionRows().count() > 0`, was **VACUOUS**: measured, the *identical* row anchor renders **20 rows** for
another tenant's manager on a **different app**, several of them badged the very type under test. A live,
in-demo, fully-populated page satisfied it. **Finding the control and finding the defect were one
measurement**, which is why the probe comes before the assertion every time.

Two refinements iter-19 adds to the recipe above:

- **A recorded rejection can be right and still be the wrong QUESTION.** The control file already recorded
  the obvious vantage as rejected — a Workforce-org manager driven at the hiring base **ejects the browser to
  production**, so the absence is true and meaningless — and iter-19 re-confirmed it. Keeping that note was
  correct; it was also why the gap survived four iters. What unlocked it was not overturning the rejection
  but **inverting the question**: not *"what does a non-recruiter see on the hiring board?"* but *"what does
  the hiring board's own locator set find on a non-hiring tenant's own grid?"* Same locators, live page,
  outcome legitimately absent.
- **Define a complement by EXCLUSION, never by enumeration.** The type-purity half asserts "no row badge is
  anything other than Hiring" (`hasNotText`), not `/^(Assessment|Training|Interview)$/`. An enumerated
  "everything else" is an assertion with an expiry date — it silently stops discriminating the day a new
  variant ships, which is the quietly-matches-nothing shape this milestone found 17+ times.
- **And a seed fact whose home is another MODULE still gets fenced.** The asserted shared-position count is
  not in the seed YAML at all — it is a code-owned Go constant in `stack-seeding`. That is the *easiest* kind
  of fact to leave dangling, because nothing in the consuming module's own directory changes when it drifts.
  The fence parses that file directly, fail-closed on a rename. Extend the reconciliation to the fact's real
  source; do not downgrade the assertion because the source is inconvenient.

**A styled string and a DOM string are different strings — `innerText` applies CSS `text-transform`,
`textContent` does not, and Playwright matches `textContent` (v2.8 M256 iter-19).** A type badge that
*renders* `HIRING` holds `Hiring` in the DOM, uppercased purely by `text-transform: uppercase`. iter-19's
probes read `innerText`, so the locator built from that reading was `getByText(/^HIRING$/)` — plausible on
screen, **impossible in the DOM** — and it failed the sharpened final *and its own negative control* on the
first live run. The general form: **a probe must read the same property the locator will match.** `innerText`
is the right tool for "what does the user see" and the wrong one for "what will my selector find"; the two
diverge on `text-transform`, `::before`/`::after` content, and visibility-collapsed whitespace. Match
case-insensitively when a badge's case is a styling decision.

**Two stat cards in the same app can have OPPOSITE shapes — measure, don't infer from the sibling page object
(v2.8 M256 iter-14).** The profile stat card renders label-and-value as one element (`"Verified Skills\n8"`),
so its accessor matches the label and parses the number from the same node. The Workforce dashboard's card is
the reverse: the label is its own `<span>` carrying **no number**, and the parent's `textContent` is
`"40Overall Members40 active"` — value first, label second, a *second* number after it, and no whitespace
anywhere. An accessor copied from the profile one returns `null` against a page that plainly renders the
stat, and the Playthrough fails on a working surface.

**An assertion that cannot tell "no data" from "a service is down" is the could-not-fail class wearing a
different hat (v2.8 M256 iter-15).** `pt-activity-drilldown` asserted `contentRows().first()` visible then
`count() > 0`. Measured: when the `jobsimulation` container is down, its grids render **20 `<tr>` whose
`textContent` is empty** — *indefinitely* (watched 40 s), with **no empty state and no error anywhere in the
UI** — and both assertions pass on that. The Playthrough still failed, three steps later, on a row-link wait,
reporting a timeout that blamed a locator. The sharpened assertion (rows that **carry text**) failed
immediately and named the real condition. A suite whose job is to detect breakage must not carry assertions
that report the wrong *cause*; that is the same defect as one that cannot fire, just louder and wronger.

Which leads to the operational half, worth knowing before you spend an hour on it: **a clean `Exited (0)` is
not a healthy container.** After an un-clean Postgres restart, `jobsimulation` and `cms` **self-terminate by
design** on their DB-health monitors (*"DB too many ping failures, shutting down"*) and nothing restarts
them. `docker ps` then reads 14 of 16 "Up". Disk was fine — this is **not** the
[`build-budget.md`](build-budget.md) M239-F1 ENOSPC trap, which is the first thing it resembles. Recovery is
a `docker start` of the two containers (no build, no compose, no teardown). **Check container liveness before
diagnosing a Playthrough**: the cheapest measurement, and it should be a bring-up cheap-win.

**A fence that needs a human to find its own misses is not yet a fence — and the widening must be MEASURED
(v2.8 M256 iter-15).** The bounded-interaction fence scoped itself to retry loops, enumerated its
out-of-scope set, and stated its own trigger for growing: *"if a straight-line site ever produces an opaque
hang, it becomes evidence and this boundary moves."* Its self-test proving it was **not** trigger-happy then
quoted `getByText(/How we measure/i).first().click()` verbatim as an example of a *safe* site — and that line
hung for **600 s** on a vantage where the tab does not exist (Playwright's action timeout defaults to `0`).

Two transferable moves. First, **find the property that distinguishes the harmful sites**, rather than
converting everything: here it is *the element may legitimately not exist on some vantage*, which is not
statically decidable — but a method whose **name** declares intent (`open*` / `switchTo*` / `expand*` /
`reveal*` / `drill*`) is exactly where that is true, so the author's own naming makes the rule decidable.
Second, **measure the blast radius before adopting the rule**: this one flags **7** sites, where "bound
everything in the page-object layer" would have been ~28 evidence-free edits. Seven is a boundary moving;
twenty-eight is a fence rotting into noise and then being switched off. Keep the old self-test, correctly
re-scoped, as the record of where the boundary was.

**A probe must use the predicate the CODE uses (v2.8 M256 iter-15).** A probe that counted elements whose
text *equalled* a step name read `0` on a surface where the accessor under test — which matches a **regex
substring** — reads `1`, and that briefly looked like a refutation of a correct earlier finding. An
exact-match probe over a substring accessor is not a stricter measurement, it is a **different question**.
Sibling of iter-14's rule about DOM shape: what was inferred rather than measured here was the *matching
semantics*.

**A settle predicate the empty state satisfies is not a settle predicate (v2.8 M256 iter-14).** A probe that
waited for `table tbody tr > 5` reported a grid as *populated* while it was rendering **20 rows with no cell
content**, and the conclusion drawn from that — a permanently empty surface, i.e. a free contrast vantage —
was wrong in the direction that ships. **The instrument is part of the measurement:** this is the same defect
as an assertion that cannot fail, committed in the probe rather than in the test. A settle predicate must be
FALSE in the state it is waiting to leave. (The finding that survived the retraction is worth keeping: an
`await expect(rows.first()).toBeVisible()` + `count() > 0` pair is satisfied by a skeleton grid, so assert
rows that carry **text**.)

**`\b` in a `hasText` regex is unreliable — `textContent` concatenates sibling nodes (v2.8 M256 iter-13).**
Playwright's `hasText` filter matches against an element's **`textContent`**, which joins sibling text nodes
with **no separator**. A work-timeline card therefore reads `…Meridian LabsFeb 2024 - Present (2 years)…`, in
which "Feb" is preceded by the "s" of "Labs" — so a pattern anchored `\b\w{3} \d{4}…` has no word boundary to
find and matches **nothing**. The identical constant is safe when consumed through `getByText`, which resolves
to leaf-ish elements whose text is not a concatenation. **Same regex, different consumer, different rules.**
iter-13 shipped this and caught it in the same hour, as a **false RED** on a page that plainly rendered the
element — a false red is exactly as dishonest as a false green, and it was caught only because the sharpened
finals were run and *watched*. When a landmark pattern is used as a container filter, drop the boundaries and
pin the concatenated shape in a unit test.

**A BOUND is not a RECOVERY — a retry loop over a mounted UI object needs both (v2.8 M256 iter-13).** M256's
harden pass established that every interaction inside a retry loop must carry an explicit timeout, so a stuck
attempt *yields to the next* (Playwright's action default is `0` — no timeout — bounded only by the test
budget). Necessary, and not sufficient. `pt-assignment-assign` then failed with every bound correct: the assign
modal is **ROW-SCOPED** (its title is *"Assign Skill Path to `<member>`"*, rendered by the member row's action
cell), so a members-table re-render **unmounts** it — and the modal had been opened 2.2 s after the first row
painted, while the table was still settling. From the trace: healthy at t+3.79 s, the Select's inner input
"not stable" ×3 then "detached from the DOM" at t+4.15 s, and **the dialog never returned**. The remaining time
decomposes exactly as the ladder's own bounds — 3 × 15 s of clicks against a locator that cannot resolve + a
20 s diagnostic + the spec's 15 s expect = **84 s**, the reported duration — reported as *"element(s) not
found"* on the submit button, three layers from the cause. **Bounding makes a stuck attempt yield; it does not
make a dead subject detectable.** So: check the subject still exists at the top of every attempt and
**re-establish it**; prefer not racing at all (a semantic settle — e.g. two equal reads of a row count ~1 s
apart — never a banned `networkidle` one); and note that recovery creates a *correctness* obligation, because a
re-opened modal may target a different row, so any identifier the assertion depends on must be read from the
instance that **accepted** the action, not from the one that was first opened.

Two method notes worth as much as the fix. **Read the artifact the failure already produced before proposing a
mechanism:** three plausible causes (a bloated Casbin policy — measured clean at `g3 = 171` for 191 memberships
with 0 orphans; an antd `maskClosable` re-click — it *throws* on the mask and the modal survives; an `Enter`
keypress with the dropdown closed — `aria-expanded` stays true) were each refuted by a bounded probe, and the
trace's own arithmetic handed over the fourth. **And prove a recovery deterministically rather than fishing for
the flake:** drive the exact failing state with a *real* user action (here the modal's own Cancel — `Escape` is
disabled on it, measured) and show the ladder recovers. Never manufacture the state by deleting DOM nodes —
iter-07's rule is that a control the application never learns about proves nothing.

### The `blocked` outcome — proving the platform correctly says *no* (v2.8 M256 iter-11)

For 23 Playthroughs across five releases the suite had **zero** non-`success` outcomes. That was not an oversight
in the specs — **it was a property of the SEED**. AI-Simulations access is a per-membership **g3
`FEATURE_JOB_SIMULATIONS`** grouping row in Sentinel's Casbin policy, added by an org-admin action
(`OrgAllowUserToUseFeature` → `AddNamedGroupingPolicy("g3", org, membership)`) and **never a default** — but the
`UsersSeeder` had written it for **every** membership since M42e iter-09, because a demo whose members cannot
launch a sim is a broken demo. Measured on `demo-2`: 20/20 · 40/40 · 40/40 · 40/40. **There was no refusal
anywhere in the world to drive**, so `blocked` was 0 by construction, and no amount of spec-writing could have
changed it.

**The fix is a seed opt-out, not a harness trick.** `StoryOrg.sim_feature_disabled: true` withholds the g3 grant
for one org (`blueprint` → `ResolvedStory.SimFeatureEnabled()` → the `UsersSeeder` guard), and Org B of
`pt-world` declares it. The refusal then comes out of the **running enforcer**: clicking *Start Simulation*
opens the deny dialog *"You cannot start AI Simulations in this organization / Please contact your administrator
at **Halcyon Retail** to request access."* and the URL **never advances** to `/sim/<slug>/start`. Nothing is
stubbed, intercepted, or faked. **A refusal faked in the harness proves nothing about the platform.**

**Assert a refusal from more than one direction.** A `blocked` outcome is the easiest outcome to satisfy by
accident: a page that failed to load also fails to show a launch confirmation. So the Playthrough pins four
things — the deny dialog is PRESENT, it **names the member's own organization** (so the assertion proves *which*
tenant was denied, not merely that something was denied — the M219 lesson in the negative direction), the launch
confirmation is ABSENT, and the URL is still the detail route. **A dead page satisfies exactly one of those.**

**The refusal and the launch are each other's negative control.** `pt-aisim-chat-launch` (Org A, granted) asserts
that same deny locator **ABSENT**; `pt-aisim-org-feature-blocked` (Org B, withheld) asserts it **PRESENT**. One
locator, two orgs, opposite verdicts, both live on every run — which is what makes the launch Playthrough's
`toHaveCount(0)` meaningful rather than vacuous: a locator that silently stopped matching anything would still
pass there, and the paired Playthrough is what goes red for it. This is the **cross-vantage** negative-control
mechanism, and it costs a manifest story plus a seed flag when the two vantages differ by **seeded state** rather
than by test code.

> **⚠️ `--reset` was not resetting the authz grants, and only a test that needed a grant to be ABSENT could see
> it (v2.8 M256 iter-11).** The first live run of the refusal Playthrough went **RED against a world that was
> never in its declared state**: `stackseed --reset` deleted only `g2` rows, so **`g3` accumulated forever** —
> 731 rows for 140 memberships on `demo-2`, **540 of them orphaned** from worlds already truncated. And because
> seeded membership ids are **deterministic**, a stale g3 row from a previous seed **silently re-granted** the
> feature to the freshly-seeded world: the org declared as *not* having AI Simulations came up granted **20/20**.
> The reset now deletes the seeded grouping policies **as a class** (`g2` + `g3`, pinned by
> `cmd/stackseed/reset_casbin_test.go`, never a `TRUNCATE` — the table also holds `init_policy.sql`'s global
> policy). **The general lesson: an additive leftover in a reset path is invisible for as long as every test
> wants the thing to be PRESENT.** The suite was green for five releases *because* every assertion was a success
> assertion — the first negative assertion found the leak on its first run. That is the argument for negative
> controls stated as a measurement rather than a principle.

**The mutation class of a Playthrough is a MEASUREMENT, not a reading (v2.8 M256 iter-06).** Every Playthrough
spec now carries a machine-checked `@pt-mutation: MUTATES | READ-ONLY | UNKNOWN` tag, plus a
`@pt-negative-control:` line whenever the class is `MUTATES`, fenced by
`playthroughs/e2e/tests/mutation-class-fence.unit.spec.ts` (one class **per `@pt:` id**, not per file — one spec
file holds two Playthroughs). `MUTATES` carries the strict definition: mutates state **AND reads it back**. A
spec that writes and only checks a toast, a closed modal, or in-page client state is `UNKNOWN` — which is why
that state exists and why it is not a synonym for "probably fine". The fence exists because the count was wrong
in **both** directions when it was merely read off the specs:
- `pt-aisim-chat-launch` was assumed to mutate ("clicks Start Simulation"). It writes **nothing**: reaching
  `/sim/<slug>/start` and rendering the launch confirmation created **0** `public.job_simulation_sessions` rows. The
  session is written past the welcome dialog, on the far side of the §5.8 live-AI boundary.
- `pt-skillpath-legacy` does mutate, but not observably where its own comment implied: `Start` writes a
  `public.skill_path_sessions` row that lands `progress=0, started_at=NULL`, and next-web's CTA needs one of
  those two to render anything but "Start". So the *enrolment* is invisible and the **step-completion** is the
  write worth reading back.
Neither correction was findable by reading. **Also**: the tag grammar is deliberately disjoint from `@pt:` —
`cmd/ptvalidate/discover.go` scans `@pt:(...)` and rejects any hit with no manifest use case as an ORPHAN, so a
first draft using `@pt:mutation` **failed validation**. The fence pins that disjointness against its own copy of
the Go regex.

**A cross-vantage negative control lives OUTSIDE the Playthrough it covers, and it discriminates ONLY an
org- or hero-SPECIFIC outcome (v2.8 M256 iter-12).** A mutating Playthrough gets its control free from its
pre-state read (above); a READ-ONLY one cannot, because its outcome is already present when it starts. The
control therefore comes from a **contrast vantage** — a hero or org for whom the asserted outcome
*legitimately does not exist* — with the Playthrough's **own final locator** run against it and required to
find nothing. Three properties make it honest, and each was learned the hard way:

- **It is not inside the Playthrough.** A second login roughly doubles that Playthrough's duration, and
  clause 1 gates the *median per Playthrough* — 16 in-test controls would break the speed clause in order to
  satisfy the honesty clause. So `playthroughs/e2e/tests/negative-controls.spec.ts` declares **no `@pt:` id**:
  not a Playthrough, not reconciled by `ptreport`, never in the median, and batched by vantage so N absences
  cost ONE login.
- **It asserts LIVENESS before absence, polled.** A dead page satisfies *every* absence assertion (iter-07's
  ablation: `bodyLen` 2147 → 24), so an absence is evidence only once the app is proven up. Polled because a
  bare `.count()` right after a `domcontentloaded` navigation reads the pre-hydration DOM and reports a
  working app as dead — a false RED inside the mechanism built to prevent false greens.
- **The coverage link is machine-checked and fail-closed.** The control file declares which Playthroughs it
  covers; the fence unions those links with the specs' own `@pt-negative-control:` lines, and **rejects** a
  link naming an id no Playthrough declares (phantom coverage — a rename is the easy way to create it) or a
  token that does not look like an id (a typo, or the tag written in prose).

**The limit is the load-bearing part.** A **structural** final — a stat label, a chart, a table's first row —
renders for *any* populated org or *any* seeded member, so no contrast vantage exists for it: measured,
Org A's manager reads `verifiedSkillsStat` 1, `skillCharts` 10, `workSection` 1, because the M44
profile-completeness seeder gives every member a career and skills. Writing contrast controls for those would
produce assertions that pass for any org — **re-introducing the exact vacuity iter-07 refuted, via the
mechanism adopted to replace it.** Their fix is instead to **sharpen the final to name real seeded data**,
which strengthens the Playthrough whether or not a control follows. Two vantages were also rejected outright
on measurement, recorded so they are not re-tried: the hiring Results view for a Workforce-org manager
**ejects the browser to production** (`app.anthropos.work/login`, `bodyLen` 162 — "absent" while not even in
the demo), and a Playthrough sitting on a known false green must not be given a control at all (it would
certify it).

> **A negative control does not only confirm an absence — it finds assertions that prove less than they
> appear to.** On its first run, the readiness control caught that
> `pt-aireadiness-manager-howwemeasure`'s step-name assertions **match on a non-readiness org**, because
> `/ai-readiness` without the feature renders a live **upsell** panel that names the very steps. Those
> sub-assertions are satisfiable by the not-enabled state — so a Playthrough that looked like it proved the
> method panel was partly proving the marketing copy.

**The negative-control COUNT is computed too (v2.8 M256 iter-11).** The same fence now reports
`@pt-negative-control registry: N of M Playthroughs carry a negative control` and names the uncovered ids, with a
no-regression floor. It exists for the reason iter-06's header already gives — *a gate whose metric is a prose
claim is not a gate* — applied to the figure that was the milestone's largest remaining gap: through iter-10 the
negative-control count was a prose number quoted from iter to iter. The floor is a floor and not an equality on
purpose: the target is *every* Playthrough, the count climbs across iters, and a fence that had to be edited on
every increment would be edited without being read. What it cannot do is go quietly backwards. Note a
**cross-vantage pair contributes two** — the relation is symmetric (each member asserts the same locator in the
opposite direction), so both sides carry the tag.

> **…and coverage is credited PER PLAYTHROUGH, not per file (v2.8 M256 harden pass).** The count read its
> `@pt-negative-control:` tag with a non-global regex — *first hit only* — and then credited **every** `@pt:` id in
> that file. iter-06 closed exactly this per-file-vs-per-ID hole for `@pt-mutation:`, one field over, and left it
> open on the number **clause 2 is scored by**. It survived twelve iters because it is *latent*:
> `studio-builder.spec.ts` is the only file holding two Playthroughs and it declares no control yet — so the very
> edit that closes those two controls is the edit that would have inflated the count by two and cleared the floor
> on one declaration. Now arity-checked and fail-closed: a file declaring fewer controls than it holds Playthroughs
> credits **none** (crediting one would pick a Playthrough arbitrarily), and the mismatch is its own named failure.
> **The general rule: when a count feeds a gate, the unit the parser credits must be the unit the gate counts.**
> Verified by injecting one tag into the two-Playthrough spec — the count held at 13, where the old parser read 15.

**Bound every interaction inside a retry loop, or the loop is decoration (v2.8 M256 iter-06).** A Playwright
action with no `timeout` inherits the **test** budget. `pt-assignment-assign` wrapped its antd-Select interaction
in a 3-attempt retry loop whose first `combobox.click()` was unbounded — so when the Modal's open animation and
its async option load kept re-mounting the inner `<input>` ("element is not stable" → "element was detached from
the DOM, retrying"), Playwright retried *silently for the full 240 s* and `attempt` never reached 1. Measured:
**245 s timeout inside the suite, 6.0 s green on an immediate solo re-run** — the signature of a retry loop whose
first attempt can eat the whole budget. Two-part fix: wait for the **form** to have mounted (the submit button
attaching is the cheapest semantic signal that the dialog body is rendered, not mid-animation), and give every
interaction an explicit `timeout` so a stuck attempt yields to the next one. Under `retries: 0` a flaky
Playthrough is a **defect**, so this class gets fixed, never re-run.

> **The class outlived its first fix, and the fence is what ends it (v2.8 M256 harden pass,
> `bounded-interaction-fence.unit.spec.ts`).** iter-06 fixed the *site*. The two retry loops written after it —
> `assignments-page.ts:openAssignBuilderForFirstAssignable` (iter-03) and `org-admin-page.ts:clickUntilDialog`
> (iter-04) — reproduced the shape exactly: an unbounded `click()` **inside** the loop and an unbounded `waitFor`
> **guarding** it, with only the inner `dialog().waitFor` bounded. Each declared a 30 s budget that could not be
> enforced from either position, and the class went on to cost **two more 240 s hangs** (iter-11 run 1, iter-12
> run 1).
>
> It stayed open partly because it was counted by the **spelling of the symptom**. iter-12 recorded "four unbounded
> `waitFor` calls remain … none is inside a retry loop, so none is proven harmful" — but D25's root cause was a
> `click()`, not a `waitFor`, and two of those four are the *guard* of a retry loop, which is the same
> unreachability through a different door. **A fence scoped to the spelling of the bug you already found is the
> mistake iter-03 corrected for `networkidle`** — the same lesson, one subsystem over.
>
> The fence's invariant is the loop's own contract: inside a `for(;;)`/`while` block that re-checks a deadline, and
> on the wait immediately guarding it, every interaction carries an explicit `timeout`. The bounded click sits
> **inside the `try`**, because D25's remedy is that a stuck attempt *yields to the next* — a bounded click outside
> the try aborts the loop on the first detach, which is the same unreachable-loop outcome by yet another road.
>
> **The exception boundary is enumerated, deliberately:** straight-line interactions elsewhere in the harness (28
> sites) are **out of scope** — there is no loop deadline for them to falsify, the test budget *is* their intended
> ceiling, and a blanket rule would be 28 edits with no evidence behind any of them. D25's sentence "give **every**
> interaction an explicit timeout" was scoped to the interactions *in that loop*; reading it harness-wide is how a
> fence becomes noise and then gets switched off. If a straight-line site ever produces an opaque hang, that is
> evidence and the boundary moves — with the measurement recorded, as D25's was. *A fence with known exceptions
> that are not written down is a fence that will rot.*

**A source-scan fence must report the right LINE, not just the right verdict (v2.8 M256 harden pass).** Both
Playthrough source-scan fences stripped comments by **deleting** block comments before scanning. Every offender's
line number was therefore shifted by the length of the file's own prose — and these files carry 70–90-line
headers. The fences were correct about *whether* and wrong about *where*: a live mutation at
`org-admin-page.ts:62` was reported as line 24, sending the reader to an innocent line. Blank block comments **in
place** (`m.replace(/[^\n]/g, ' ')`) rather than removing them. Cheap, and it is the difference between a fence
that is trusted and one that is argued with.

**One committed `.only` makes the whole suite report success on 1 of 30 (v2.8 M256 harden pass 2).** Both
Playwright configs in rext set `forbidOnly: !!process.env.CI`, and **nothing that drives either harness sets
`CI`** — not `run-playthroughs.sh`, not `run-coverage.sh`, not `run-latency.sh`. So a single `test.only` left in a
committed spec silently reduces the run to that one test and exits **0**. Measured: a 20-test unit run plus
`.only` on a third spec → `1 passed`, rc 0, no warning anywhere. For the Playthrough suite `ptreport` *would* flag
the other 29 as *"did not run"* — and at the time the runner swallowed that reconciliation into a deliberate
non-fatal `|| echo`, so the **run verdict** stayed green having checked one Playthrough. (**v2.8 M256
harden-final closed that half**: on a full run the `ptreport` gate is now **BINDING**, so 29 *"did not run"*
entries fail the run — see [`../verification.md`](../verification.md) § *A gate whose exit code is discarded is
not a gate*. `forbidOnly` remains the first line of defence.) For the `stack-verify` sweeps it is
worse: their denominators (**29/29**, then **47/47** — figures this corpus quotes) are read from what ran, so a
shrunken world reports as a complete one.

> **A fence under `tests/` cannot close this**, and that is the interesting part: `.only` stops the fence from
> running too. The guard has to live where Playwright evaluates it *before any test* — the config. Both are now
> `forbidOnly` **default-ON**, with a named escape (`PT_ALLOW_ONLY=1` / `PW_ALLOW_ONLY=1`) for a deliberate local
> focus run. **General rule: a check that lives inside the thing it checks cannot detect a failure mode that
> suppresses execution.**

**`ptreport` reconciled a run that never happened (v2.8 M256 harden pass 2).** `run-playthroughs.sh` handed
`ptreport` a path — `e2e/report/last-run.json` — and nothing asked whether that file came from the run just made.
Demonstrated: `npx playwright test --no-such-flag` exits 1 having written no report, and `ptreport` still printed a
complete four-state map off the file already on disk and applied its gate to it.

The runner's own history says this is not hypothetical: its **M204 iter-02** note records a CLI `--reporter=list`
override *replacing* the config's reporter list, leaving `last-run.json` stale for a whole milestone and
decoupling the four-state map — the map the milestone gate reads — from the actual run. **That fix removed the
cause and added no check.** Now two mechanisms, because they fail differently:

1. the runner **deletes** the results file before the run, so *"the reporter never fired"* is a loud missing-file
   error rather than a silent read of last time;
2. it records the run start and passes `ptreport --results-not-before <epoch>`, which **refuses** an older file.

The refusal is **exit 3, not the gate exit**: a stale file does not mean the Playthroughs regressed, it means this
reconciliation has no evidence, and sending the diagnosis the wrong way is the mistake `runDatadnaClosure`
documents one binary over. The flag takes an **integer epoch** deliberately — M236 lost half the world to an age
check that parsed a UTC timestamp as local time, and an integer has no timezone to get wrong.

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
  page-object layer + specs + the serial runner · `e2e/drafts/` — measured-but-unshipped specs, `.draft`-suffixed
  so Playwright cannot collect them) · `report/` (the four-state map) · `fixtures/`
  (version-controlled static input files fed to the real file chooser, spec §5.4 — **populated at v2.8 M256
  iter-18**: `synthetic-cv-sre.pdf` + `synthetic-cv-sre.docx`, a wholly invented CV whose employers and school
  occur nowhere in the seed, the taxonomy, or any real registry, so an assertion naming them can only be
  satisfied by *this file having been imported*. No **shipped** Playthrough consumes them yet — the self-import
  use case they exist for carries a `will-not-build` verdict and its CV route is blocked by a product defect
  upstream of the fixture; the two files ARE the evidence for that verdict). Section README:
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


---

## Content stories — where the (session × action) proof lives (v2.5 "the playbill" M236)

**Not a Playthrough, and deliberately so.** A Playthrough *plays a journey* — it logs in as a hero and
performs the actions that produce an outcome. A **content story** is the opposite direction: the session
was **already played** (cloned from a real production session by the M232 `ContentStorySeeder`), and what
must be proven is that its **result surface renders real content** for the player and manager vantages.

There is nothing to play, so there is no Playthrough. The proof lives in the **content-stories sweep**,
specified in [`coverage-protocol.md` § "Content stories — the (session × action) LANDS sweep"](coverage-protocol.md):

| | Playthroughs (this doc) | Content stories |
|---|---|---|
| Question | *can the hero DO the thing?* | *does the already-played story SHOW real content?* |
| Actor | a roster **hero**, playing forward | a non-hero **`content-player-<idx>`** seat, landing on a result |
| Entry | a journey's first surface | an **exact URL** from the seeded `content-manifest.json` |
| Harness | `playthroughs/e2e/` | `stack-verify/e2e/{tests/content-stories.spec.ts, lib/content-result-page.ts}` |
| Data | the decoupled `pt-world` seed, reset-to-seed | the demo's own content-story seed (source-pinned) |

**What IS shared:** the M37 cockpit seat-switch (`lib/cockpit-login.ts`) — the content-stories sweep uses
the same `loginAs()`, exploiting its **`landingPath`** option to enter directly on the result URL. As with
Playthroughs, the seat-switch was **reused, never forked**.

**Where they meet:** the manager-view trap — ⚠️ **it was a MIRROR trap until M257x iter-129 and the mirrors
are gone.** Both suites depend on the canonical `public.job_simulation_sessions` /
`public.skill_path_sessions` rows being seeded — a manager scoreboard reads them
(`content-stories-routes.md`). M236 found the mirror correctly populated (13/13) while the manager
scoreboard still rendered `No data` + `undefined undefined`, proving the mirror is **necessary but not
sufficient** for the manager vantage.
