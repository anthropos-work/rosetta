# iter-22 — progress

**Type:** tik (standard shape, `TOK-01` move 3 — its last unit of work)

## Phase A — drive the journey before writing a line

Three probe passes, all outside `tests/` so no fence, no `@pt:` reconciliation and no gated median
ever sees a probe. Every question answered by observation:

| # | question | measured |
|---|---|---|
| Q5 | does the write succeed with iter-21's three fixes in? | **YES** — `createJobRole` → `200 {"data":{"createJobRole":{"id":"J-PTROL1-39EA","name":"PT Role …"}}}` |
| Q2 | what happens after Save? | the app **NAVIGATES** to `/enterprise/roles/<id>?setup=true`, **1 526 ms** after the click |
| Q1 | does the list paginate? | **yes — 20 per page**; 41 roles; pager "1 2 3" |
| Q3 | is the new role on page 1 after a reload? | **NO** — `titleRows=0` after a *successful* create |
| Q7 | what does the table look like while loading? | `t=0 rows=0` · `t=500 rows=1 "No roles match your filters."` · `t=1000 rows=20` |
| Q6 | is there a catalog TOTAL? | **yes** — a stat card, `textContent` `"42Roles"` (value-first, no separator) |
| Q8 | does a search read-back survive a full reload? | **yes** — 1 row, badged `Custom` |
| Q9 | does the detail route serve the role by its server id? | **yes** — a fresh navigation renders the title as its only heading |
| Q10 | is a never-created title absent under the same search? | **yes** — `namedRows=0`, with the empty-state row present as the liveness witness |
| Q13 | which input filters the table? | of **three** inputs in `<main>`, the one with placeholder `"Search roles…"` (the others are antd Select internals, one being Page Size) |

**The drafted spec was refuted in every clause** (D103) and **the page object's settle predicate was
satisfied by the empty state** (D104). Both are recorded as decisions because both are defect classes
this milestone hunts, found in its own instrument.

## Phase B — land it, and watch every assertion go RED

**Shipped:**

- `tests/orgadmin-role-create.spec.ts` — `pt-orgadmin-role-create`, MUTATES, with an **in-line
  negative control** (liveness-then-absence) and **three independent finals** (D106): the
  server-assigned id served on a fresh navigation · a full reload + the surface's own search naming it
  in exactly one row · the catalog **total** grown by exactly one.
- `lib/org-admin-page.ts` § `OrgRolesPage` — rebuilt against measurement: content rows **subtract the
  placeholder by text**; `emptyCatalogRow()` reuses it as a positive liveness witness; `totalRoles()`
  reads the stat card and returns **`null`, never `0`**, when it is absent; `createRole()` returns the
  **server-assigned id** instead of waiting for a hidden dialog; `searchRoles()` deliberately does not
  settle (the caller polls the semantic outcome).
- `lib/url-shapes.ts` — `ROLE_DETAIL_URL` + `isOnRoleDetail` + `extractRoleId` (D105).
- `tests/orgadmin-locators.unit.spec.ts` — **5 tests**, fail-closed on a vacuous capture, pinning the
  subtraction, the liveness witness, the URL split and null-not-zero.
- `manifest/org-admin.yaml` — `org-admin.roles.UC1` `TODO → pt-orgadmin-role-create`, its final re-cut
  to the three measured facts, and the eighteen-iter history kept (it is a defect-class record).

**10 mutants, each watched** (D108): M1 skip-write · M2 control inverted (`Expected 1, Received 0`) ·
M3 total unchanged (`Expected 43, Received 44`) · M4 ghost read-back · **M5 live proof the OLD settle
predicate PASSES on an empty catalog** · M6 wrong detail id · M7–M10 the new fence's own four tests.
All restores verified byte-identical (`cp` backups, never `git checkout`).

## Phase C — the gate

`playthroughs/e2e/run-playthroughs.sh 2 --reset`, **3 consecutive cold reset-to-seed runs**, rc
captured per run into a variable (never off a pipe):

| run | result | rc | wall-clock |
|---|---|---:|---:|
| 1 | `187 passed` | **0** | 1.6 m |
| 2 | `187 passed` | **0** | 1.6 m |
| 3 | `187 passed` | **0** | 1.6 m |

**0 flake.** Baseline was `181 passed` → **+6** (1 Playthrough + 5 fence tests).

`ptreport`: **26/31 passing (83.9 %), 0 failing, 0 unimplementable, 5 TODO** — and **all 5 remaining
TODOs are onboarding**, which is the arithmetic statement that org-admin is complete.

Registries, computed by the fence (never narrated):
- `@pt-negative-control`: **24 of 26** (was 23 of 25 — numerator and denominator both moved)
- `@pt-mutation`: **MUTATES=8** READ-ONLY=16 UNKNOWN=2 (was 7)

`stackseed --policy-check --stack demo-2` → `live=18 expected=18 · OK`, re-run **after all three
resets** — iter-21's grant survives because the seeder writes it, not because the stack remembers it.

**The deliberately drifted cockpit fixture was preserved.** Backed up before the runs, and the
post-reset re-export (`d03781b2…`) replaced with the fixture (`e991b47a…`), **verified byte-identical**
and re-checked structurally: **12 advertised heroes** across 4 stories.

### Clause 1 — REPORTED, not gated (D-v28-13)

Environment stated with the number: `Kirality-Mac-Pro-6.local`, darwin 25.1.0, Docker VM 9.70 GiB,
`demo-2` offset 20000, localhost/http, `--no-public-host`. **Not comparable to billion's 228 s.**

| statistic | value |
|---|---|
| median per non-studio Playthrough (median of each id across 3 runs, n=24 ids) | **2.500 s = 0.7517×** |
| the same **excluding** the new Playthrough | **2.500 s = 0.7517×** — identical |
| observed range across the 24 | 1.300 s – 9.000 s |
| `pt-orgadmin-role-create` itself | 9.0 / 9.0 / 8.6 s |

Two honest notes. **The new Playthrough is the slowest non-studio one in the suite** (5 navigations
and a real backend write), and it did not move the median at all — one id in 24. And **0.7517× sits
inside the spread D-v28-13 published** (0.5281×–1.0762× on code no iter has touched since iter-03), so
it is reported as one sample, not as a verdict. **iter-22 landed no speed mechanism, so clause 1's leg
half has nothing new to measure; its flake half is MET** — and met by 0 flake over 3 runs with rc
captured per run.

## Close — 2026-07-30

**Outcome:** clause 3's org-admin cluster is **COMPLETE, 3/4 → 4/4** — the ninth product's last use
case, un-homed for five releases — landed with its own negative control so clause 2 moved with it
(controls 23/25 → **24 of 26**, mutating 7 → **8**). The drafted spec it replaced **could not have
passed on a working product**, and the page object it consumed had an intermediate satisfied by an
empty table.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D103, D104, D105, D106, D107, D108, D109 (this iter's `decisions.md`)
**Side-deliverables:** none — every change fell inside the planned scope.

**Routes carried forward:**
- `DEFECT-M256-silent-forbidden-mutation` → **iter-23**, deliberately sequenced (it needs the grant
  revoked, which would break the write path this iter landed). Reproducible on demand, so schedulable.
- `ONBOARD-M256-seat-append` + the 4 remaining onboarding UCs → the long pole, unchanged.
- `D-v28-5` part (b) · `PT-M256-readiness-step-asserts` · `NEGCTL-M256-studio-pair` → unchanged.

## Lessons

1. **A parked spec is a hypothesis with good formatting.** This one was written at iter-04, reviewed,
   and referenced by four later iters as "the Playthrough is WRITTEN and lives in drafts/". Every
   clause of its final was false about the running product, and one probe pass found all of them. The
   cost of *not* probing would have been a RED on a working feature — and, given the fifteen-iter
   belief that this form was broken, that RED would have been believed.
2. **The instrument's defect was worse than the subject's, and it was invisible in green.** An
   intermediate reading "the catalog has content rows" was satisfied by a table saying "No roles match
   your filters." — which is also what the surface renders when its backing service has
   self-terminated. Nothing about a green run could have shown that; only asking *what does this read
   when the answer should be no* did.
3. **When pagination defeats your read-back, look for the total.** The surface publishes its own
   catalog size, and a `+1` on that number cannot be defeated by sort order, page size or index lag —
   whereas the row-count delta the draft chose is a bet on all three. The stat card was two DOM nodes
   away from the assertion that could not work.
4. **A URL shape that matches too much is a passed assertion about the wrong page.** `ROLES_URL` was
   right for its original job and wrong the moment a journey left the list, and the failure mode is an
   assertion that *passes* — the same shape as iter-19's grid rows in another tenant's app.
