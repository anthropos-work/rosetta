# iter-22 — decisions

## D103 — the drafted spec's final could not have passed, and only driving the journey said so

`e2e/drafts/orgadmin-role-create.spec.ts.draft` had sat, written and reviewed, since iter-04. Its
final was *"the dialog hides, then the list underneath grew by exactly one row."* Every clause of
that sentence is false about the running product, and each was refuted by a 9-second probe:

| the draft assumed | measured |
|---|---|
| the dialog closes and the list is underneath | the app **NAVIGATES** to `/enterprise/roles/<id>?setup=true` **1 526 ms** after Save |
| the list shows the catalog | it **PAGINATES at 20 per page** (41 roles, pager "1 2 3") |
| the new role appears in it | `titleRows=0` on a full reload **after a successful create** — it is not on page 1 |
| `roleCount()` reads the catalog | it read **1** against a catalog of **41** |

So the draft would have gone **RED on a working product**, and the fifteen-iter-old belief that "the
create-role form doesn't work" would have been *confirmed by the test written to check it*. iter-20
retired one wrong diagnosis of this UC by reading the network; this iter retired a second by driving
the DOM. **A parked spec is a hypothesis with good formatting.**

## D104 — the loading row IS the empty row, so the settle predicate was satisfied by nothing at all

Sampled at 500 ms intervals over the surface's first six seconds:

```
t=0 ms     rows=0
t=500 ms   rows=1   "No roles match your filters."     ← the placeholder, mid-fetch
t=1000 ms  rows=20  "5G Security Consultant | …"       ← the real catalog
```

`roleRows()` was `main table tbody tr`, and `roleCount()` waited for the first row to be *visible*
then counted. The placeholder is a `<tr>` and it is visible, so the accessor returned **1** — and the
Playthrough's `roles-surface-populated` intermediate is `roleCount() > 0`. **That intermediate was
satisfied by a table displaying "No roles match your filters."**

This is iter-14 **D71** (*a settle predicate the empty state satisfies is not a settle predicate*) and
iter-15 **D76** (*an assertion that cannot tell "no data" from "a service is down"*) — and D76 is not
hypothetical here: when `demo-2-jobsimulation-1`/`cms` self-terminate `Exited 0`, jobsimulation
surfaces render exactly this shape.

**Proven live before the fix was written**, on the same surface filtered to a ghost term:

```
OLD predicate → rows=1, first visible=true  →  `> 0`  PASSES on an empty catalog
NEW predicate → rows=0                      →  `> 0`  correctly FAILS
```

Fixed by **subtracting the placeholder by its TEXT** — not by `tr.ant-table-placeholder` and not by
`tr[data-row-key]`, both of which were available and both of which are the class/attribute anchors
this module's locator discipline forbids. And the same row is then reused *positively* as
`emptyCatalogRow()`, the **liveness witness** iter-12's rule demands before any absence assertion:
the sentence the app renders when it has answered and found nothing.

## D105 — `ROLES_URL` matches the detail route, so "still on the list" was satisfiable by a role page

`ROLES_URL = /\/enterprise\/roles(?:[/?#]|$)/` is segment-anchored on `roles`, which is correct for
rejecting `/enterprise/roles-import` — and it also matches `/enterprise/roles/J-PTROL1-39EA`. So
`isOnRoles(detailUrl)` is `true`, and an assertion that the admin is on the role **list** is satisfied
by a role **detail** page. Latent until this iter because no Playthrough had ever left the list.

`ROLE_DETAIL_URL` + `isOnRoleDetail` + `extractRoleId` split them, and the landing assertion now says
both things: on the roles route **and not** on a detail route. The trap itself is asserted in the
fence (test 4) so the reason the second shape exists cannot be forgotten and "simplified" away.

## D106 — three independent finals, because each one alone is weaker than it looks

The manifest's final was one sentence; it is now three, and the reason is that each candidate final
has a specific way of being true without the write having landed:

| final | what it rules out | what it does NOT rule out alone |
|---|---|---|
| (a) a fresh navigation to `/enterprise/roles/<id>` serves the role | client-side-only state; a closed modal | that the id came from anywhere the client chose |
| (b) a full reload + the surface's own search names it in exactly 1 row | a fabricated id; stale client cache | a pre-existing role of the same name (per-run-unique title covers this) |
| (c) the catalog TOTAL grew by exactly one | a listed-but-uncounted row; a search index artifact | nothing further — but it needs the total to *exist*, hence `totalRoles()` returning `null` rather than `0` |

(c) is what pagination forced and it is the strongest of the three: the stat card counts the whole
catalog, so it cannot be defeated by sort order. **`totalRoles()` returns `null` when the card is
absent, never `0`** — "the card did not render" and "the org has zero roles" are different facts and
only one of them is a bug, and a silent `0` would make the delta assertion compare two unknowns.

## D107 — the negative control is in-line, and it is the pre-state read ASSERTED

Per iter-06 **D22**, a mutating Playthrough's pre-state read is its control when the final is a strict
delta. That only counts if the pre-state read is an *assertion*, not a baseline variable — a `before`
that is merely captured proves nothing about absence. So the control is explicit:

1. search the per-run-unique title;
2. **assert the surface is ALIVE** — `emptyCatalogRow()` has count 1, i.e. the app answered the query
   (iter-12: prove liveness before absence, or a dead page passes);
3. **assert the role is ABSENT** — `roleNamed(title)` has count 0.

Watched RED: inverting step 3 to `toHaveCount(1)` fails with `Expected: 1, Received: 0`.

No contrast vantage is needed and none would be honest here: the outcome is not org- or
hero-specific, it is *newly created*, so the absence is genuinely available on the same vantage a
moment earlier. That is the cheapest control shape in the suite and it exists only for writes.

## D108 — ten mutants, and the one that mattered most was about the instrument

| # | mutation | result |
|---|---|---|
| M1 | skip the write entirely | RED |
| M2 | the pre-state control expects the role PRESENT | RED — `Expected: 1, Received: 0` |
| M3 | the total expected unchanged | RED — `Expected: 43, Received: 44` |
| M4 | the list read-back searches a GHOST title | RED |
| M5 | **live**: the OLD settle predicate on an empty catalog | **PASSES** (rows=1, visible) — the defect, demonstrated |
| M6 | the detail read-back uses a wrong role id | RED |
| M7 | drop `roleRows()`'s `hasNotText` | RED (fence tests 1 + 2) |
| M8 | `emptyCatalogRow()` subtracts instead of selecting | RED (fence test 3) |
| M9 | `ROLE_DETAIL_URL` made as loose as `ROLES_URL` | RED (fence test 4) |
| M10 | `totalRoles()` falls back to `0` instead of `null` | RED (fence test 5) |

M3 is the most informative *green*: the total moved **43 → 44** as a direct consequence of the write,
which is the whole claim of the Playthrough stated as one integer. M5 is the most informative
*failure*, and it is not a mutation of the subject at all — it is a measurement of the instrument, and
the instrument was the thing that was broken.

Every mutation used a `cp` backup and every restore was verified byte-identical (the standing ban on
`git checkout <file>`, kept).

## D109 — a stray unhandled rejection is a green suite with an error beside it

The fence's first run reported `5 passed` **and** `1 error was not a part of any test`: `searchRoles`
is `async`, was fired without an await against the Recorder fake, and its `waitFor` rejected after the
test had finished. The suite was green; the run was not clean. Fixed by giving the Recorder inert
action stubs and driving the async accessor to completion. Recorded because "passed, with an error
printed underneath" is precisely the kind of verdict this milestone exists to stop trusting.

## Routes carried forward

- **`DEFECT-M256-silent-forbidden-mutation`** — unchanged and still owed. Deliberately not opened here
  (it needs the grant revoked, which would break the write path this iter landed). **iter-23.** Its
  evidence is reproducible on demand — we hold the grant — so it is schedulable, not perishable.
- **`ONBOARD-M256-seat-append`** and the four remaining onboarding UCs — the long pole, unchanged.
- **`D-v28-5` part (b)**, **`PT-M256-readiness-step-asserts`**, **`NEGCTL-M256-studio-pair`** — unchanged.
- **A note for whoever writes the next org-admin-adjacent spec:** the roles surface's stat card is the
  pagination-proof read-back shape, and `OrgTagsPage`'s name-lookup is the other. Between them they
  cover every list-plus-write surface in the product; a page-1 row-count delta covers none of them.
