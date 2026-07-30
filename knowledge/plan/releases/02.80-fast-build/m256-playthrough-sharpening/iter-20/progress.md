**Type:** tik · shape: standard (single target: `org-admin.roles.UC1`, the last org-admin UC)

# iter-20 — fifteen iters of form debugging, and the answer was in the network tab

## Phase A — drive the dialog, and watch the half nobody had watched

iter-05 parked this UC with a diagnosis about the FORM: *"`Save` is enabled with the form incomplete and
fails silently with an EMPTY alert region"*, plus a route through an LLM `Generate` leg. Three passes:

| pass | question | answer |
|---|---|---|
| 1 | what is in the dialog, and what does `Save` do? | Two textboxes + an unselected-looking **"Core skills"** mode choice. `Save` is **correctly DISABLED** when empty, enables when both fields are filled. Clicking it: dialog stays open, `[role=alert]` **0**, catalog **20 → 20**, role absent. 40 s. |
| 2 | choose "Start from scratch" (the non-LLM mode) | **No change at all** — and the card HTML shows why: it was **ALREADY SELECTED** on open (`--color-primary-light` background, `--color-primary` border, before *and* after the click). There is no LLM leg on the happy path. |
| 3 | **capture the network** | `createJobRole` **IS SENT**, and comes back `200` with `"unauthorized: forbidden"`, `path: ["createJobRole"]`, `DOWNSTREAM_SERVICE_ERROR`. |

**It was never a form problem. It is an authorization denial** (D97), and the resolver names the permission:
`resolver_skiller_taxonomy_authz.go:75` checks `permission.OrgFeatureTaxonomyWrite`. The running Sentinel
policy holds `org:feature:taxonomy:read` for **four** roles and `org:feature:taxonomy:write` for **none** —
so `createJobRole` cannot succeed for any seeded actor. Not a stale enforcer cache either: the runner
reloads Sentinel on every `--reset`, and did.

**Two of iter-05's claims are retracted** (D97): the `Generate` route was never needed, and `Save`'s
enablement is correct, not premature.

## Phase B — not implemented, deliberately

The mechanical fix is one policy row (`admin` → `org:feature:taxonomy:write`), and the seeder already writes
per-membership feature grants, so it is technically available. **It is refused as an implementation
decision** (D99): a Playthrough that passes only because we granted ourselves a permission the product
grants to nobody proves nothing about the product. That is iter-07's rule as sharpened by iter-17 — *the one
move that can manufacture a state the application never learns is the wrong move where "the app permitted
it" is what you are proving* — one layer below the DOM, in the policy.

So the iter ends on the question rather than on a green test. **The choice needs one fact this iter cannot
obtain locally:** whether a production org carries an `org:feature:taxonomy:write` grant. One query, behind
the standing sign-off rule.

## Phase D — the tree is unchanged, and verified so

No shipped code was modified this iter (the probe was removed; `git status` clean before the run). One
confirming cold reset-to-seed run rather than three, because there is nothing new to de-flake:
**181 passed, rc 0**, `ptreport` **25 passing / 0 failing / 6 TODO / 0 unimplementable**. Drifted cockpit
fixture restored + sha-verified **`99e2f315`**.

## Close — 2026-07-29

**Outcome:** **`org-admin.roles.UC1`'s fifteen-iter-old diagnosis is replaced by its real root cause, and it
is not the one anybody was looking for.** The create-role `Save` does not fail on the form — the
`createJobRole` mutation fires and Sentinel **refuses** it (`unauthorized: forbidden`), because
`org:feature:taxonomy:write` is granted to **no role in the policy** while its `:read` counterpart is
granted to four. The UI surfaces this as **absolutely nothing**, which is a user-facing product defect in
its own right (D98) and the reason the form got blamed for fifteen iters. The UC is **not landed**, because
the only way to land it is to grant ourselves the permission under test — a user decision, put to the user
with both options and the single fact needed to choose (D99).
**Type:** tik
**Status:** closed-no-lift
**Gate:** NOT MET — clause 2 mutating **7** (≥5 MET), `blocked` **1/1 MET**, negative controls **23 of 25**
(terminal); clause 3 verdict half **COMPLETE**, landed half **org-admin 3 of 4**, **onboarding 1 of 5**,
`unimplementable` still **0** *pending the D99 decision*; clause 1 leg half **N/A**, flake half **MET**
(181 passed, 0 failing). **D-v28-5** part (b) open.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (2 no-prog tiks — iter-18 and this one; the streak needs 3, and iter-19 lifted the metric between them) — (3) re-scope: n (the trigger is > 3 `unimplementable`; this is a candidate FIRST, and the trigger explicitly does not fire on one) — (4) user-blocker: **y — an architectural/authorization question whose answer changes what code lands in THIS iter** (D99: grant the permission and land an 8th mutating Playthrough, or report a platform defect and record the milestone's first `unimplementable`). Phase 5 §4 names this case exactly: *"is this a documented divergence or a real bug?"* — (5) cap-reached: n (3rd tik of this invocation) — (6) protocol-stop: n — Outcome: exit-4
**Decisions:** D97 (**authorization, not the form** — `createJobRole` fires and is denied; `taxonomy:write`
granted to nobody; two iter-05 claims retracted), D98 (**PRODUCT DEFECT** — a forbidden mutation renders as
nothing at all, which is why the form was blamed), D99 (**granting ourselves the permission would
manufacture the capability under test** — escalated with two options and the one missing fact).
**Side-deliverables:** none — no shipped code changed.
**Routes carried forward:**
- `PT-M256-orgadmin-role-create` → **BLOCKED ON A USER DECISION (D99).** Not a build task any more. Option
  (a) report the platform authz gap and record the milestone's first `unimplementable`; option (b) confirm
  the demo policy is incomplete and seed the grant. **Needs one production Sentinel policy read** — does a
  real org carry `org:feature:taxonomy:write`? — which is behind the standing sign-off rule.
- `DEFECT-M256-silent-forbidden-mutation` → **NEW (D98).** Report to the platform: a refused GraphQL
  mutation produces no user-visible error on the create-role surface. Worth a sweep — if the pattern is
  shared by other org-admin writes, every one of them fails invisibly.
- `NEGCTL-M256-studio-pair` + `FIX-M256-studio-false-green` → unchanged; the latter is now the only route to
  clause 2's 25 of 25.
- All of iter-18's onboarding routes stand, `ONBOARD-M256-seat-append` first.
- `D-v28-5-cockpit-logout` part (b) → unchanged.
**Lessons:**
1. **When a UI write fails silently, read the network before you read the form.** Three iters of this
   milestone's own history (iter-05's diagnosis, iter-17's routing, iter-20's first two passes) were spent
   inside a dialog whose problem was one HTTP response away. The product showed a form, so everyone debugged
   a form.
2. **A `200` is not a success.** GraphQL errors ride inside a 200 body, so a network panel filtered on
   status shows nothing wrong. The probe that found this matched on the response *body*, not the code.
3. **"Already selected" looks exactly like "my click did nothing".** Pass 2 clicked the mode card and read
   no change, which is ambiguous between *the click failed* and *it was already chosen*. The card's own
   inline style resolved it in one line. Dump the state, not the diff.
4. **The escalation is the deliverable, and it should be answerable in one reading.** Two options, their
   consequences for the gate, and the single missing fact that decides between them — not "blocked, please
   advise."
