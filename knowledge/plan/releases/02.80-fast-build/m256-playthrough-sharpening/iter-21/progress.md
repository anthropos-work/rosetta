---
iter: 21
iteration_type: tik
status: closed-fixed
opened: 2026-07-30
closed: 2026-07-30
---

# iter-21 — progress

## What landed

Three config-fidelity gaps, stacked behind one silent failure, each found by **driving the write path**
rather than reasoning about it — and each fixed in a different rext section. Code: rext `ef41a19`
(tag `fast-build-m256-iter-21`, on origin).

| # | gap | section | shape |
|---|---|---|---|
| **F1** | replay never restored the identity sequences | `stack-snapshot` | the root cause |
| **F2** | the platform's sanctioned `p3` grant was never applied to any stack | `stack-seeding` | the D99 subject |
| **F3** | `SKILLER_AZURE_OPENAI_{KEY,ENDPOINT_URL}` absent from the secret DNA | `stack-secrets` | the middle blocker |

**F1 is the generalisable one and it is bigger than this use case.** Replay clears with
`TRUNCATE … RESTART IDENTITY` — deliberately, *"so re-loaded rows keep stable ids"* — then COPYs the rows
back **with their explicit ids**, and nothing ever put the sequences back above them:
`job_role_embeddings` 18 920 rows / max id 21 274 / **sequence at 4**; `skill_embeddings` 42 790 / 43 583 /
**sequence at 1**. Those are the **only two identity columns in `public` and BOTH were broken**, so on
**every demo ever built** every taxonomy write duplicate-keyed — creating a role and creating a custom
skill alike. The new replay Phase 5 discovers sequence-backed columns from the **target's live catalog**
(whole-class, so a future migration is covered without a re-capture) and `setval`s each to
`max+1, is_called=false` — the shape that is also correct on an empty table, where `true` would silently
burn id 1.

**And the fence, which is worth more than the use case.** `stackseed --policy-check` compares a stack's
live `p3` surface against the checked-in expected set and reports **both directions**: MISSING is the
under-grant that caused this, EXTRA is the over-grant. That is the mechanical form of the judgement
iter-20 had to make by hand — *a Playthrough over a permission we granted ourselves is green about our own
grant*. A gate in the verb, an advisory in the seed path (D-M255-1: one assert, two consumers, two
contracts).

## Verification

Closed by the coordinator after an API 529 killed the sub-agent mid-cycle (the sixth agent death of the
session; rext was already committed and pushed, so nothing was lost).

- **3× cold reset-to-seed gate: `181 passed`, `rc 0`, every run.** Exit codes captured per run, not off a pipe.
- **`stackseed --policy-check --stack demo-2` → `live=18 expected=18`, rc 0 — re-run AFTER all three cold
  resets.** That is the load-bearing check: the grant survives reset-to-seed because the **seeder** writes
  it, not because anyone inserted a row by hand.
- Write path proven end-to-end before the crash: the role was created and the app navigated to its detail
  page. Not "the mutation returned 200" — the thing happened.
- Drifted cockpit fixture backed up and restored, sha `99e2f315` verified (the three resets re-export it).

## Close

**Status:** `closed-fixed`.
**Decisions:** D99 mechanism **corrected** (recorded in milestone `decisions.md`) — the seeding fleet never
wrote a `p3` row in its life; all 17 come from the platform's `init_policy.sql`, which **deliberately**
withholds `taxonomy:write` (platform `c6096d1`) and ships `local_superadmin_grants.sql` as the sanctioned
grant whose stated use case is verbatim *"Testing flows that require taxonomy:write"* — and which **nothing
has ever applied to a demo or dev stack**. The demo was faithful to `init_policy.sql` and unfaithful to
production.
**Side-deliverables:** none — every change fell inside the planned scope.

**Routes carried forward:**
- `PT-M256-orgadmin-role-create` → the three blockers are cleared and the write path is proven; **landing
  the Playthrough itself, with its negative control, is still open** (the crash landed between the fix and
  the test). Next iter.
- `DEFECT-M256-silent-forbidden-mutation` → **still owed, and now more urgent, not less.** With the grant in
  place the silent-failure path stops being exercised on a demo, so the evidence is perishable. Sweep the
  other org-admin writes for the same shape and record it before it becomes unreproducible.
- onboarding routes, `ONBOARD-M256-seat-append` first (append only — `personaUserIndexFor` indexes by
  declaration order); `D-v28-5` part (b); `PT-M256-readiness-step-asserts`; `NEGCTL-M256-studio-pair` — all
  unchanged.

## Lessons

1. **Clearing one blocker is how you find the next one.** Three gaps sat in series behind a single
   silent refusal, and only the first was visible. Each was found by driving the write to completion —
   the second and third did not exist as hypotheses until the one in front of them was gone.
2. **"Faithful to the seed file" and "faithful to production" are different properties, and only one was
   ever checked.** The demo matched `init_policy.sql` exactly. It still misrepresented production, because
   production also carries a row from a file nobody runs. A fidelity check against the wrong reference
   passes.
3. **The correction strengthened the refusal it corrected.** iter-20 declined to grant itself the
   permission under test. Under the coordinator's *wrong* mechanism ("our seeder dropped it") the fix would
   have been patching around an invented divergence; under the true one it is applying the platform's own
   row, from the platform's own file, for the platform's own stated use case. **Getting the mechanism right
   changed what the fix means, not just how it reads.**
4. **A whole-class fix costs barely more than the instance.** Reading sequence-backed columns from the live
   catalog instead of naming two tables covers every future migration without a re-capture. The two tables
   were the entire population *today*; the catalog query is what keeps that true tomorrow.
