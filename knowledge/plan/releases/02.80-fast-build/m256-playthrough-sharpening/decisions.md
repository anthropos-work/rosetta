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
