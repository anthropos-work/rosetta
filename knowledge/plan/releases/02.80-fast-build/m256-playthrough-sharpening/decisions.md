# M256 — decisions

Release-level binding decisions **D-v28-1 … D-v28-12** live in
[`../../../roadmap.md`](../../../roadmap.md) § Active — v2.8. This file carries the milestone's own
**strategy chain** (`TOK-NN`) plus milestone-level decisions. Intra-iter decisions live in each
`iter-NN/decisions.md`.

---

## TOK-01: cheap-lever speed, then the cluster that discharges two clauses at once — 2026-07-28

**Tok type:** bootstrap (iter-01)

**Initial strategy.** Four moves, strictly ordered, each landing before the next opens:

1. **Measure the denominator first, on this box, and change nothing until it exists.** A relative gate
   (D-v28-12: median per-Playthrough ≤ **0.79×** a same-stack pre-work baseline) is unfalsifiable without a
   measured starting point. n=3 on the local `demo-2`, environment stated with the number, recorded in
   `progress.md`. Report **median per non-LLM Playthrough** (the gated metric), the **suite wall-clock**
   (reported, not gated), and the **studio lane** separately (budgeted separately per D-v28-9).
2. **Take the per-test latency lever, not the parallelism lever.** Clause 1 is a **per-test median**;
   worker count cannot move a per-test median (iter-01 D1). The real lever is the residual
   **`networkidle`**: **12 of 18** login call sites omit `waitUntil` and inherit `cockpit-login.ts`'s
   `'networkidle'` default on an app whose own helper doc records that `networkidle` "resolves late and for
   the wrong reason"; plus **8 further unfenced violations** in the harness (2 page-object `goto` overrides
   + **6 unbounded `waitForLoadState`** sites the Phase-0b audit found). M254 iter-10 measured 13 min →
   3.8 min from exactly this class. Land the fix **with a widened fence** — the existing
   `home-login-networkidle.unit.spec.ts` guards only `/home`-landing specs, which is precisely why the
   other holes survived. Ship the machine-checked per-spec **`MUTATES` / `READ-ONLY` / `UNKNOWN`** tag in
   the same iter, following the `@pt:`-grammar **twin lockstep test** shape the audit identified as prior
   art (never a third unfenced regex).
3. **Land org-admin before onboarding, because org-admin discharges two clauses with one body of work.**
   All four curated org-admin UCs declare a persist-then-observe final → 4 mutating Playthroughs → **5**
   with `pt-assignment-assign`, which is exactly clause 2's `≥ 5 mutating` floor, while being half of
   clause 3's scope (D-v28-4). Onboarding is **seed-blocked** (audit F5 answers the overview's Open
   Question 1: **no** — `UsersSeeder` writes a membership for every user unconditionally; there is no
   pre-onboarding state, and none can be *declared* into existence), so it is ordered second and its cost
   is a **seeder + capability + roster seat**, not just specs. A seed wall must not be allowed to starve
   the clauses org-admin already discharges.
4. **Close the honesty items last, deliberately, not as leftovers.** Negative controls; the `blocked`
   outcome from an **RBAC/Sentinel deny** rather than an entitlement tier (iter-01 D4, refuted in-iter by
   the Phase-0b audit — `actor.entitlement` reaches no DB column, and `ptvalidate`'s precondition check
   **fail-opens** on it); the D-v28-5 cockpit Back-to-Cockpit / logout double-click fix; and a **written
   verdict for every remaining uncovered curated UC** including the 5-release-old M206/M207 reservations.

**Rationale — why this is the right opening move.** The milestone's own plan named parallelism as the
headline lever and the plan review already broke half of it. iter-01 broke the other half: **the D-v28-12
re-cut dissolved the requirement entirely**, because a median per-test metric is indifferent to worker
count. That removes the largest, riskiest item (a refactor of an Alignment-DNA-gated mirror engine) from
the critical path and frees the whole milestone budget for work that actually moves the three clauses. What
replaces it is cheaper *and* better-evidenced: a defect class with a measured precedent (M254 iter-10), a
known target list (12 + 8 sites), and an existing fence to widen rather than invent. The ordering rule
throughout is **discharge-per-unit-of-work**, which is why org-admin — the only cluster that serves two
clauses simultaneously — goes before the coverage cluster that is seed-blocked.

**Strategy class:** `new-direction` (bootstrap — no prior strategy to compare against).

**Distance-to-gate context.** Gate metric: **median per-Playthrough**, target **≤ 0.79× baseline**, on the
**post-coverage** suite (denominator 18 → ~27). Starting value: **not yet measured on this host** — that is
iter-02's entire job, and it is the first thing that happens. The only comparable prior number is billion's
228 s / 18 tests (~6.4 s per non-LLM test), which **must not be quoted as comparable** to any M256 number
(D-v28-12; the absolute billion re-measure is routed to M258). Clause 2 starts at **1** mutating
Playthrough and **0** `blocked` / **0** negative controls; clause 3 starts at **0** onboarding and **0**
org-admin.

**Known context carried from the Phase-0b audit (verdict YELLOW, `kb-fidelity-audit.md`).**
- **F4** `actor.entitlement` is declared-only → the `blocked` outcome needs a different refusal surface.
- **F5** no pre-onboarding state exists → onboarding needs a seeder, not just tests.
- **F6** `--reset` is **whole-stack** (`doReset` takes no org filter; it truncates
  `public.organizations`/`users`), and `pt-world.seed.yaml`'s header comment claims the opposite. Every
  `run-playthroughs.sh --reset` on `demo-2` therefore **destroys the showcase world** — acceptable, because
  `demo-2` is dedicated to this milestone, but it must be stated, not discovered.
- **Gap 4** 8 unfenced `networkidle` violations (2 `goto` + 6 unbounded `waitForLoadState`) — folded into
  move 2 above.
- **Gap 7** `run-content-stories.sh` recomputes a 47-vs-pinned-45 pair count and `sys.exit(2)`s, so the
  content-stories sweep **refuses to start**. Not M256's suite; **Fate 3 → M257/M258**, which compose it.

**Next-tik direction (iter-02).** Measure the baseline and nothing else. `run-playthroughs.sh 2 --reset`
from the authoring copy with `stack-demo/rosetta-extensions/demo-stack/stacks/demo-2/bin` prepended to
`PATH` (the M204 iter-05 gate-run prereq), n=3, `PT_HOST=localhost`, `PT_APP_SCHEME=http`. Record: per-test
durations from `report/last-run.json`, the median over the **non-LLM** subset, the studio lane separately,
the suite wall-clock, and the environment. Do **not** change harness code in iter-02 — a baseline measured
on already-modified code is not a baseline.

---

## USER-BLOCKER — 2026-07-29 (iter-20): `org-admin.roles.UC1` needs a disposition, not more debugging

Surfaced per Phase 5 § 4 (*"an architectural question whose answer changes the iter's planned fix shape"* —
its canonical example is literally *"is this a documented divergence or a real bug?"*). Full evidence in
[`iter-20/decisions.md`](iter-20/decisions.md) D97–D99.

**The finding.** The create-role `Save` was recorded for fifteen iters as a form/no-op defect. It is an
**authorization denial**: `createJobRole` is sent and Sentinel returns `unauthorized: forbidden` inside an
HTTP **200** GraphQL error body. The resolver checks `permission.OrgFeatureTaxonomyWrite`; the running policy
grants `org:feature:taxonomy:read` to **four** roles and `org:feature:taxonomy:write` to **none**.

**Why it is not simply fixed.** One policy row would land the Playthrough, and the seeder is already the
right place to write per-membership feature grants. But granting ourselves the permission whose enforcement
is the thing under test **manufactures the capability** — iter-07's rule as sharpened at iter-17, one layer
below the DOM. A green there would be a green about our own grant.

**The two options, and what each costs the gate:**

| | consequence for the exit gate |
|---|---|
| **(a) platform authorization gap** — report it; the UC becomes the milestone's **first** `unimplementable-without-platform-edit` | clause 3's landed half tops out at **org-admin 3 of 4** with a written verdict; `unimplementable` 0 → 1 (the re-scope trigger is **> 3**, so it does **not** fire) |
| **(b) the demo's policy is incomplete** and production grants it another way | the seeder writes the grant; the UC lands as an **8th** mutating Playthrough and org-admin completes **4 of 4** |

**The single fact that decides it:** does a real production org carry an `org:feature:taxonomy:write` grant?
One Sentinel policy read — which is a **production** read and therefore behind the standing sign-off rule, so
it was not taken.

**Recommendation, held loosely:** (a). The product ships a "New Role" button, a create dialog and a
`createJobRole` mutation that no role in the policy can execute, and the refusal is invisible in the UI —
which reads like a gap rather than a demo artefact. It is circumstantial, and this milestone's record is full
of confident readings that measurement overturned (two of them inside iter-20 itself), so it is put to the
user rather than assumed.

**Separately and unconditionally: `DEFECT-M256-silent-forbidden-mutation`** (D98) should be reported whichever
option is chosen — a refused mutation rendering as no user-visible error at all is a defect on its own terms,
and worth sweeping across the other org-admin writes.

## D99 — RESOLVED: the demo's seeded policy was incomplete; the platform is fine (2026-07-30)

**The question.** iter-20 found `pt-orgadmin-role-create` blocked by authorization, not the form:
`createJobRole` fires, Sentinel refuses it inside an HTTP 200, and the demo grants
`org:feature:taxonomy:write` to **no role at all**. Landing the UC would mean granting ourselves the
very permission whose enforcement the test exists to check — so the iter escalated instead of
forcing a green. Two candidate readings: **(a)** our seed is missing a grant real orgs have, or
**(b)** nobody has it anywhere and the feature is unreachable, a platform bug.

**Decided by reading production, not by argument** (read-only; casbin policy is configuration, not
tenant data, so it sits inside `db-access.md`'s boundary — and the user's constraint was *don't write*
to the platform DB, which a read does not):

| | `…taxonomy:read` | `…taxonomy:write` |
|---|---|---|
| **production** | admin, candidate, content_creator, member | **admin** |
| **demo-2** | admin, candidate, content_creator, member | **NONE** |

**(a) is correct. The delta is exactly one row** — `p3 | admin | org:feature:taxonomy:write`. The
seeder replicated the four read grants and dropped the single write grant.

**Bounded, and checked for siblings:** `org:feature:taxonomy:write` is the **only** `*:write` action in
production's entire `p3` surface, so there is no wider sweep hiding behind this. One row, one fix.

**Consequences.**
1. The seeder writes the missing grant (a **demo**-DB write, never prod), and `org-admin` can complete 4/4.
2. **iter-20's escalation was right and its hypothesis was wrong** — exactly the outcome escalating is
   *for*. Had it guessed (b) it would have filed a platform bug that does not exist; had it granted
   itself the permission blind, it would have papered over a real seed-fidelity defect.
3. **A fidelity finding that outlives this UC:** every demo built to date has misrepresented what an
   org admin can do. A presenter demonstrating role creation would have hit a silent refusal. That is
   a seeder-fidelity bug, not a Playthrough bug, and it is the same class as this milestone's others —
   a surface that *looks* correct because nothing compared it against the real thing.
4. `DEFECT-M256-silent-forbidden-mutation` stands on its own and still routes to the platform: a
   refused mutation renders **no user-visible error**. That is a real product defect independent of
   who holds the grant, and the demo's missing grant is precisely what surfaced it.

### D99 — CORRECTION to the mechanism (2026-07-30, from iter-21's measurement)

**The verdict stands; my stated mechanism was wrong and is retracted.** I wrote that "the seeder
replicated the four read grants and dropped the single write grant." **The seeding fleet has never
written a `p3` row in its life.** All 17 come from the platform's `sentinel/init_policy.sql`.

What is actually true, and it is a better justification than mine:

- `init_policy.sql` **deliberately withholds** `taxonomy:write` — platform commit `c6096d1`, *"drop
  default admin taxonomy:write, add on-demand grants file"* — and ships
  `sentinel/local_superadmin_grants.sql` as the sanctioned way to add it, whose stated use case is
  **verbatim** *"Testing flows that require taxonomy:write"*.
- **Nothing has ever applied that file to a demo or dev stack.** So the demo was **faithful to
  `init_policy.sql` and unfaithful to production**. Not a dropped row — an un-applied sanctioned file.

**Why this matters beyond pedantry:** it removes the last doubt about iter-20's refusal. Applying the
grant is *not* the manufactured-capability move it declined to make — **the row is the platform's own,
the platform's own file names this exact use case, and production has it.** Under my wrong mechanism
we would have been patching our seeder to paper over a divergence we invented; under the true one we
are applying a grant the platform ships for precisely this purpose.

**The generalisable finding also grew.** The blocked write was hiding two more, and #3 is the largest:
replay does `TRUNCATE … RESTART IDENTITY` (deliberate — "so re-loaded rows keep stable ids") then
COPYs rows back **with explicit ids**, and never puts the sequences back above them. Those two tables
are the **only** identity columns in `public` and **both** were broken, so on **every demo ever built**
every taxonomy write duplicate-keyed — creating a role and creating a custom skill alike. The fix
discovers sequence-backed columns from the **target's live catalog** rather than a hardcoded list, so a
future migration is covered without a re-capture.

**And the fence is worth more than the UC:** `stackseed --policy-check` compares a stack's live `p3`
surface against the expected set **in both directions** — MISSING is the under-grant that caused this,
EXTRA is the over-grant, which is the mechanical form of the judgement iter-20 had to make by hand.
*A Playthrough over a permission we granted ourselves is green about our own grant.*

---

## DEFECT-M256-silent-forbidden-mutation — CAPTURED (iter-23, 2026-07-30)

**Routes to the PLATFORM (`next-web-app`). No platform code was edited. Read-only source reads only.**

A mutation the backend REFUSES is, from the user's side, indistinguishable from a mutation that was never
sent. iter-20 found this while root-causing `org-admin.roles.UC1` and recorded it in one paragraph; iter-21's
fix then **removed the symptom from every demo**, so iter-23 reproduced it deliberately (revoke the
`p3 admin → org:feature:taxonomy:write` row on **demo-2 only**, drive the journey, restore the row and
re-verify) and enumerated every channel a user or operator could learn from.

### What a refused `createJobRole` shows the user — MEASURED, all channels

| channel | reading |
|---|---|
| HTTP status | **200** — the error rides inside it (`errors[].extensions.errors[].message = "unauthorized: forbidden"`, `code: DOWNSTREAM_SERVICE_ERROR`, `data: null`) |
| `[role=alert]` | **count 1, text EMPTY** — the slot is mounted and says nothing |
| `[role=status]` | none |
| antd `message` / `notification` / `form-item-explain` | **all empty** |
| the dialog | **stays open, `Save` still ENABLED** — it invites a retry that will fail identically |
| the URL | unchanged |
| the state | catalog total **49 → 49**, delta **0**, no row for the title |
| the browser console | one unrelated Clerk dev-keys warning — **nothing about the failure** |
| uncaught page error | **YES** — `Failed to fetch from Subgraph 'backend' … unauthorized: forbidden` |

**So D98 is confirmed and sharpened.** It is not that the app has no error handling; it is that the app has
**one** error surface here and it is reserved for a different error.

### The line-level cause, and it is TWO defects with one symptom

**(1) The form handles exactly one error code and rethrows the rest.**
`packages/ui/src/JobRoles/Form/AddJobRole.tsx` `handleSubmit`:

```ts
} catch (error) {
  const dup = duplicateJobRoleInfo(error);
  if (dup) { setServerDuplicate(dup); return; }   // ← the ONLY handled shape
  throw error;                                    // ← everything else, into the void
}
```

`throw error` from an `async` function invoked by a click handler is an **unhandled promise rejection** —
React renders nothing for it, which is exactly the `pageerror` above. `onClose()` sits *after* the try/catch,
which is why the dialog stays open. And the empty `[role=alert]` is the **duplicate-warning slot**
(`setServerDuplicate`), never populated. That is why iter-05 recorded *"an EMPTY alert region"* and why the
form was blamed for fifteen iters: the alert element it saw belongs to a different error.

Notably `throw error;` from a catch appears **exactly once** in the whole `packages/ui` tree — this form. The
rethrow is not the systemic part.

**(2) The systemic part: the app has NO default user-visible failure surface for any mutation.**
`apps/web/src/providers/Query.provider.tsx`:

```ts
mutations: { onError: (error) => { captureException(error); PosthogClient.captureException(error); } }
```

Sentry and PostHog — **no user surface**. So every mutation in the app is silent to the user on failure
unless it builds its own inline surface.

**(3) And a dead contract that makes it look handled.** Six mutations across four `hooks/organization/*`
files declare `meta: { error: 'Failed to enable organization setting' }` — human-readable failure sentences.
**No handler reads them.** There is **no `MutationCache`** anywhere in the codebase (0 occurrences); the only
`meta.error` consumer is `QueryCache.onError`, which reads `query.meta.error` and uses it **as a Sentry tag,
not as a message**. So the strings are inert, and they are inert on precisely the **org-admin write set** —
including `useUpdateOrganizationSetting`, the mutation behind `pt-orgadmin-setting-toggle`.

### The sweep the route asked for

All four org-admin writes share outcome (2). `useCreateJobRole`'s `onError` is Sentry-only; the other three go
through the same global handler; and the settings write additionally carries one of the dead `meta.error`
strings. **The authors of the org-admin writes wrote failure messages and the framework never wired them
up** — which is a more useful bug report than "the create-role form is silent", because it names a fix that
uses a convention the codebase already believes it has.

### The limit of the sweep — stated, not glossed

**Only `createJobRole` was refused LIVE.** The three sibling org-admin writes check different
permissions, and revoking each would have meant three more revoke/restore cycles against a stack later
iters depend on. So:

- **Measured:** create-role's ten channels; the empty `[role=alert]`; the uncaught page error; the
  49 → 49 no-op.
- **Definitive by source:** the dead `meta.error` strings (a repo-wide search for `MutationCache` returns
  **0**, so nothing can read a mutation's `meta`), and the global handler's Sentry-only body.
- **Inference, NOT measurement:** that a refused tags-create / member-tag / settings-toggle would look
  equally silent. The inference follows from the global handler, and it is still an inference.

Whoever picks this up: driving one sibling refusal would close that gap in one revoke.

### Suggested fix shape (for the platform team — NOT applied here)

Add a `MutationCache` with an `onError` that reads `mutation.meta.error` and renders it (the convention six
mutations already assume), and replace `AddJobRole`'s `throw error` with the same path. That turns six dead
strings live and gives every future mutation a default surface. Both are platform edits and are out of scope
for this milestone by its hardest constraint.

### Safety

Every write was to **demo-2's own Postgres**; production was not written and not read in this iter. The grant
row was backed up before the revoke and **restored byte-identically** (`diff` clean), with
`stackseed --policy-check --stack demo-2` returning **rc 0 · live=18 expected=18** afterwards. Side benefit:
iter-21's `--policy-check` fence was watched **RED against a live stack** for the first time — `rc 1`,
`live=17 expected=18`, naming `MISSING admin → org:feature:taxonomy:write (under-grant)`. It had previously
only been proven against mutants.

## D103 — clause 2's negative-control clause gets an explicit carve-out (coordinator, 2026-07-30; user may overrule at close)

**The problem.** Clause 2 reads *"every Playthrough passes a negative control"*. At **24 of 26** with two
deliberate exclusions, that sentence can never literally read "met" — so the clause as written is
unsatisfiable, which is the **third** time this milestone has found one of its own gate clauses
unmeetable as phrased (clause 1 twice, re-cut by `D-v28-12` then `D-v28-13`).

**The two exclusions are principled, not leftovers.** The studio pair sits behind
`FIX-M256-studio-false-green`: **a negative control over a known false green would certify the false
green.** Landing them before the false green is fixed would make the metric *worse* while making the
number *better* — precisely the trade this milestone exists to refuse.

**Decision.** Clause 2's control requirement is **MET at 24 of 26 with the studio pair recorded as an
explicit carve-out**, discharged by `FIX-M256-studio-false-green` rather than by a control. The carve-out
is named, reasoned, and countable — not a silent shortfall.

**Why a carve-out and not a re-cut of the number:** "24 of 26" as a target would be arbitrary the moment
another UC lands (the denominator moves with clause 3 — it went 24→25 at iter-17 and 25→26 at iter-22).
The stable statement is *"every Playthrough has a control except those whose control would certify a
known defect, each named"*.

**Flagged for the user at close.** This is the coordinator's call to avoid stalling the loop on a wording
question; it is recorded so closure ratifies or overrules it deliberately rather than inheriting it.

## D104 — clause 3's onboarding half: 4 landed + 1 reasoned verdict is the honest form (coordinator, 2026-07-30; user ratifies at close)

**The third clause in this milestone that cannot be met as written.** Clause 3 asks for onboarding's
**5 UCs LANDED**. One of them — `onboarding.enterprise-workforce-standard.UC1`, the self-import journey —
**should not be landed as a Playthrough at all**, so the sentence is unsatisfiable for a reason that has
nothing to do with effort.

**Why that UC must not be a Playthrough** (iter-18 measured it, and the refusal has held for twelve iters):

- The only path that actually advances **scrapes a live public LinkedIn profile** on a site that blocks
  automation. Landing it would make a **real person's profile a permanent test fixture**, send outbound
  traffic from a demo on every gate cycle, and be flaky by construction.
- Worse for the suite's meaning: **its RED would read as a product regression** when nothing about the
  product had changed. A test whose failure lies is worse than no test.
- The deterministic alternative is blocked by a **measured product defect**, not by us: the CV upload
  POSTs **200** for a valid PDF and a docx alike while the forward control never enables (100 s).

**Decision.** Clause 3's onboarding half is **MET at 4 landed + 1 written verdict**. The verdict is a
first-class outcome here, not a shortfall: this milestone holds **31/31 written verdicts and 0
`unimplementable`**, and the whole point of the verdict mechanism is that *"we measured this and it should
not be built"* is a result.

**Symmetry with `D103`, deliberately.** Both carve-outs share one shape: **the honest count is lower than
the number a green dashboard would show, and in both cases the missing item is missing because landing it
would certify something false** — a control over a known false green (D103), a Playthrough over a scraped
third-party profile (D104). Clause 1 was re-cut twice for the same underlying reason (`D-v28-12` →
`D-v28-13`). **Three of this milestone's three gate clauses turned out unmeetable as first written**, which
is itself the milestone's most transferable finding: *a gate authored before the work is a hypothesis about
the work.*

**Flagged for ratification at close** — recorded so closure decides deliberately rather than inheriting it.

## PLATFORM-M256-onboarding-step-not-resumed — the org-prepared onboarding flow cannot resume (iter-31, routes to the PLATFORM)

**Recorded at MILESTONE level, per the `playthroughs.md` rule iter-23 wrote** (*a product defect the suite finds
has no state, no glyph and no ledger — record it where it cannot be tidied away*), because its natural home
would otherwise be a comment inside a spec that is green.

**The symptom.** A member whose org pre-filled her profile confirms her role; the confirmation **persists**
server-side (`public.user_params.onboarding` gains a `role` step — verified in the DB every time). She reloads
`/onboarding` and is back on **step one**, progress `0`, `[Skip] [Start]`, with no trace of what she confirmed.
The screen is **byte-identical to the pre-state**. Six fresh navigations across three browser sessions, hours
apart, all agree. She can never advance past the first step across a page load.

**The mechanism — one array, two consumers, opposite ends.** Read at source and corroborated by the six
observations (a source read, so stated as such, not as a debugger session):

```
packages/graphql/src/hooks/onboarding/useGetOnboardingStatus.tsx:25-27
    result.onboarding.steps?.sort((a, b) => sorterFn({ first: b.updatedAt, second: a.updatedAt }))
    → the array handed to the component is sorted NEWEST-FIRST
packages/ui/src/Onboarding/OnboardingUser.tsx:130-132
    const lastStep = reimport ? Import : steps?.[steps.length - 1]?.step;
    → takes the LAST element of a newest-first array, i.e. the OLDEST step ever taken
```

So `lastStep` is the *first* step the user ever completed. `managerImport` (`lastStep === Import && …`) is true
again, and the initial step (`lastStep || Import`) is `Import` again — forever. The **host page reads the same
array from index 0** (`const [firstStep] = onboardingSteps`, `apps/web/.../onboarding/page.tsx:141-143`) and is
therefore *correct*, which is why completion (`done`) does redirect properly and nothing else looked wrong.

**Why nobody had seen it, and the transferable part.** It is invisible for a NULL or single-element `steps`
array — `length-1 == 0`, so both readings coincide — and that is **every one of the 191 seeded users** and every
hero any earlier iter could reach. Only a **multi-step** array exposes it, and the only multi-step user in
existence is the seat `pt-onboard-prepared` that **iter-28 minted to reach the surface in the first place**.
*The seed capability built to reach a surface turned out to be the only way to see a defect on it* — the same
shape as iter-11, where withholding one grant exposed four releases of leaked `g3` rows.

**A second, independent defect on the same journey** (recorded here so it is not lost with the spec comment):
the prepared flow **cannot be completed on a demo.** One `Next` past the skills screen reaches *"Add more
skills"*, which renders **"We're having trouble loading your skills at the moment. You can skip this step and
try re-importing your profile later."**, and its `Next` is **inert** — clicked five times, identical screen,
progress stuck at 100. `useClusterizeSkills` is the surface behind it. Consequence for the suite: the
`/onboarding` → `/home` completion read-back that `pt-onboarding-complete` relies on is **unreachable from this
seat**, so no Playthrough should be written expecting it.

**Consequence for clause 2, stated plainly.** `pt-onboarding-org-prepared`'s persistence half is **not
assertable through the UI**, and the reason is this defect rather than a harness weakness. The routed handler
`ONBOARD-M256-prepared-persistence` asked for *"a POSITIVE locator on the ROLE screen a reload lands on"* — there
is no such screen. What ships instead is a positive, hydration-proof assertion labelled for exactly what it
proves (the confirmation left her flow intact — it neither ejected her nor completed it), watched RED three ways,
and explicitly **not** presented as the write's read-back. The write itself stays proven by mutant S1 against the
screen the click reaches.

**Not fixed here.** Both are platform edits, out of scope by this milestone's hardest constraint. **Zero platform
files were modified** — the platform tree was read only.

## HARDEN-CAP-ACCEPTED-D105 — the final harden's un-stabilized cap is accepted, residuals named (user, 2026-07-30)

**The final pass stopped at `cap reached without stabilization`**, and the user was asked and chose
**"accept and close, residuals named"** over spending a fourth pass.

**The reasoning, recorded so the acceptance is legible later.** The class yield never dropped:

| pass | found |
|---|---|
| harden-1 (incremental) | **6** |
| harden-2 (incremental) | **11** → stabilized |
| harden-3 (final) | **9** → cap, un-stabilized |
| inside the iters themselves | **~9** more |

**~35 checks that could not fail, and a flat yield.** A flat yield does not mean "nearly done" — it means
the seam is **broad**, so a fourth pass would very likely find another handful and a fifth after that.
There is no natural stopping point inside this milestone, and the milestone's actual deliverable — a
sharpened Playthrough suite — is complete.

**What makes the acceptance honest rather than a shrug:** the residuals are **enumerated, not implied.**
9 standing mutants named Playthrough-by-Playthrough (mechanical, ~30 min of machine time, each needing
its own reset because the write is irreversible) plus **11 lower-severity findings recorded with
file-level specificity** so none needs re-discovery. The next milestone inherits a **list**, not a
surprise.

**The two findings that justify the whole exercise, and they are about the fixes themselves:** the
**liveness fence counted `not.toBeVisible()` as proof the page was alive** — an absence assertion serving
as the liveness witness, the exact defect it exists to prevent — and the **bounded-interaction fence had
never scanned the retry loop it is named after**. Three capabilities' entire implementations could be
deleted with the suite green, including **all of snapshot Phase 5 under the end-to-end test**, and a
**declared Playthrough could be deleted story-and-spec without turning the run red — four ways at once**.

*The fences built to catch the class were themselves instances of the class.* That is this milestone's
most transferable finding, and it is why the honest stop condition was recorded rather than rounded up.

---

## RATIFIED at close (2026-07-30) — D103, D104, and the iter-31/32 deviation

All three were recorded during the iters as coordinator calls **flagged for ratification at close**, and all
three are now **RATIFIED** under the user's standing delegation for this milestone (*"as long as we don't
touch the platform code and don't write on the platform db, for me i trust you call"*), the user having been
shown all three with their full reasoning and having separately chosen the harden-cap option when asked.
**The user's right to overrule any of them at release close is preserved.**

They are ratified **as written**, not softened:

| | what was ratified |
|---|---|
| **D103** | Clause 2's control requirement is **MET at 28 of 30**, the studio pair being an explicit, named carve-out discharged by `FIX-M256-studio-false-green` rather than by a control — because **a control over a known false green would certify the false green**. |
| **D104** | Clause 3's onboarding half is **MET at 4 landed + 1 written verdict** against the **CURATED 5**. The unlanded one is the self-import journey, whose only advancing path scrapes a live third-party profile, so **its RED would misreport as a product regression**. (The *manifest* denominator reads 6 declared / 5 live, because M256 also added one net-new non-curated use case. The two denominators are deliberately distinct — see `coverage-verdicts.md` §"The arithmetic".) |
| **iter-31/32** | The routed persistence repair was **impossible** — there is no Role screen for a reload to land on — and the coordinator had relayed iter-29's **source read as if it were an observation**. What shipped is a positive, hydration-proof assertion **labelled for exactly what it proves**, plus the platform defect routed. The absence-only assertion was **not** restored. |

**And the finding the three of them add up to, stated plainly because it is this milestone's most transferable
one: all three of the milestone's three gate clauses turned out unmeetable as first authored.** Clause 1 was
re-cut **twice** (`D-v28-12` → `D-v28-13`) after its threshold was measured to sit **inside its own 2.04×
noise floor**; clauses 2 and 3 each needed a carve-out. *A gate authored before the work is a hypothesis
about the work* — and the honest response to a falsified hypothesis is to re-cut it in the open, with the
measurement that falsified it, rather than to report the flattering reading. The flattering reading was
available every time and refused every time (most sharply at `D-v28-13`, where the original-16 subset read
`0.7063×` — inside the gate — and was rejected as a hand-picked denominator).

---

## DEFERRAL GATE — RED, resolved item by item (close, 2026-07-30)

`/developer-kit:audit-deferrals --scope=milestone` returned **RED** on three grounds: **8 items carried a
target that no longer exists** (`a later tik of M256` / `next iter` / `iter-16+` — the gate is MET and no
further iters will run), **10 repeat/chronic deferrals** (one re-typed into the next routing table ~10 times),
and **4 items routed to destinations that cannot hold them** (*"a future release milestone"*, *"whoever next
edits those specs"*, *"a future v2.8 milestone"*).

**The dead-target class is the priority, and it is the M255 failure repeating.** That close routed four items
to *"M255 harden resume"* — not a milestone — and its own retro says the routing *"should have been rejected
when written"*. A Fate-3 target that does not resolve is not a fate. **Every item below now names a real
milestone, or is dropped with a reason.**

### LAND-NOW (Fate 1 — landed in this close)

| item | what landed |
|---|---|
| `DOC-M256-claudemd-pt-count` | The count reconciliation, exhaustively: `playthroughs.md:14-16` and `demo/README.md:216-220` still read **18 live Playthroughs** while the same file's own §line read 30. Both corrected to **30 live + 1 verdicted TODO (31 manifest use cases, 10 products)**, plus `CLAUDE.md`'s onboarding sub-count given its **denominator** (4 of the curated 5). Zero `18 live Playthrough` hits remain corpus-wide. |
| `DOC-M256-ptworld-reset-comment` | `pt-world.seed.yaml`'s header claimed the showcase world is *"not touched by pt-world's reset"*. `doReset` takes **no org filter** (audit F6). A false comment is read as evidence by the next iter — the D118 rule this milestone wrote down. |
| The 4 platform defects' **durable home** | Structural gap the audit found: there was **no platform-defect register anywhere in this repo**, so all four defects lived only in a milestone `decisions.md` that archives at release close. Created [`../../../platform-defect-register.md`](../../../platform-defect-register.md) with each defect's `file:line`, so nobody re-derives them. |

### LAND-NEXT (Fate 3 — re-attached to a REAL milestone, its `overview.md` edited)

| item | target | why THIS milestone |
|---|---|---|
| `FIX-M256-studio-false-green` + `NEGCTL-M256-studio-pair` + `DOC-M256-llm-lane-premise` (one bundle, by their own routing) | **M258** | The **longest chronic in the milestone (~10 pushes)** and the reason D103 needs a carve-out at all. M258 composes this suite into the bring-up and claims a stack *"proves itself"* — a known false green inside the suite is a direct claim on M258's gate, which makes M258 the owner rather than a convenient parking space. Fix shape is measured and known (`progress.md:981`): assert a POPULATED section, not the empty scaffolding the matcher fires on at +2.1 s. |
| `PT-M256-standing-mutant-Q1` — the **9 remaining** standing mutants | **M258** | Mechanical (~30 min machine time, no design decision), each needing its own reset because the write is irreversible. Named Playthrough-by-Playthrough in `hardening-ledger.md` §residuals. Same reasoning: M258 runs the suite, so the suite's effectiveness is its business. `HARDEN-CAP-ACCEPTED-D105` accepted these as **named** residuals — that promise is only kept if they have a real home, which they now do. |
| The **11 lower-severity harden-3 scan findings** | **M258** | Recorded with file-level specificity at `hardening-ledger.md:532-544` so none needs re-discovery. Routed as a batch to the milestone that owns suite quality. *Stated risk:* M258 is `complexity: medium` and expects to close in 1–2 iters, so this batch may need re-fating there rather than absorbed silently — which is exactly why it is written into its `overview.md` rather than left in a ledger. |
| Spec-side import of the five enrolled heroes | **landed here instead** | Its stated target was *"whoever next edits those specs"* — not a milestone. It turned out to be the close review's own must-fix (the constants had **zero importers**), so it landed now. |
| `FIX-M256-demo2-service-self-termination` | **M257** | **Gate-relevant, which is the reason and not a convenience.** M257's gate reads *"reaches `autoverify green:true / 0 warnings`"*. On this failure `docker ps` shows 14 of 16 "Up", the app surfaces no error, and every jobsimulation surface renders 20 content-free rows — so M257 could **declare a green gate on a half-dead stack**. |
| `FIX-M256-autoverify-fapi-libressl` | **M257** | 31 iters with the target never advancing. It warns *"NOBODY CAN LOG IN"* on a working stack, and M257's gate reads autoverify's warning count — so a spurious warning is a gate problem. **Deliberately not landed here:** a TLS-probe fix that cannot be verified against a real bring-up is a fix on trust, and M257 brings stacks up repeatedly. |
| `FIX-M257-content-stories-pair-count` | **M258** | The sweep `sys.exit(2)`s before starting (47 recomputed against a pinned 45, omitting `manager_presence_only`), so it refuses to run at all. M258 composes the verification batch. The `FIX-M257-` prefix is an artifact of when it was found, not a routing decision. |
| `ptvalidate` is invoked nowhere outside its own tests | **M258** | Structural, and it is the honest home for the **permissive half of the runner's gate** this close left open: an all-matching `--grep` is still graded advisory, and the correct fix is a both-way `ptvalidate` pre-flight, which already implements the id-level question shell cannot. |
| `BIND_HOST` / `D-M255-7` | **M258** | **M255's own Fate-3 routing was declared and never applied** — `grep BIND_HOST` in M258's `overview.md` returned 0 hits. Applied now, with a backref. The M255 failure mode again: a routing recorded in a closing milestone's decisions is not a routing until the *target's* doc says so. |

### DROP (with reason)

| item | reason |
|---|---|
| `PT-M256-resume-fixture-pair` | Its premise dissolved and nobody noticed for ~23 iters. It existed to make two use cases **share the cost of one checked-in résumé fixture**; the fixture landed at iter-18 (`playthroughs/fixtures/synthetic-cv-sre.{pdf,docx}`), and its second member — `onboarding.enterprise-workforce-standard.UC1` — is now a `will-not-build` **verdict** (D104/D122). A pairing with one member and no cost to share is not a work item. `profile-skills.import.UC1` keeps its own verdict (A2) and its own future. |

### Fate 2 — already owned elsewhere, confirmed, no edit needed

| item | owner | verified |
|---|---|---|
| `MEASURE-M256-clause1-sampling` | **discharged** by `D-v28-13` (user, 2026-07-29) — clause 1 re-cut to gate the leg, not the aggregate — with the comparable **absolute** billion re-measure inherited by **M258** (`m258/overview.md:112-123`, explicitly *"reporting only — this does NOT become a fourth gate clause"*). | quoted at source |
| M255's 4 harden items | **M257** | `m257/overview.md:130-151` § *Inherited from the M255 close* names all four. |
| `frontend-tier.md` §8.5 numeric rewrite | **M257** | `m257/overview.md:100-113`. |

### KEEP-DEFERRED-WITH-SIGNOFF — **3 items, and each needs the USER's signature, not a re-route**

These are the only three I did not fate myself, because each is a **roadmap decision** rather than a routing
one, and assigning this fate to clear a gate is precisely the move the audit exists to catch. **Listed in the
close report for the user.**

| item | why it needs a signature |
|---|---|
| `PERF-M256-parallel-lane` | Its target was *"a future release milestone"* — not a milestone. It needs a **cookie/`__client`-scoped Clerkenstein registry or one fake-FAPI per worker**: a real build against an Alignment-DNA-gated mirror engine, not a lever. **No v2.8 gate clause needs it** (clause 1 gates the leg; M258's re-scope trigger already covers a wall-clock overrun). Naming v2.9 is a roadmap act. |
| `PT-M257-self-evaluation` | Re-homing a **5-release-old M206 vision reservation** is explicitly a roadmap decision, recorded as a *recommendation not an action* at iter-09 D39. Its current `M257` target is a **mis-route** — M257 is a build-speed milestone that will not author Playthroughs. |
| `PT-M257-talk-to-data` | Blocked on `ask_*` migrations **plus live Bedrock credentials**. A credential is not something a milestone can fix, and it belongs in the separately-budgeted integration lane. Its `M257+` target is likewise a mis-route. |
