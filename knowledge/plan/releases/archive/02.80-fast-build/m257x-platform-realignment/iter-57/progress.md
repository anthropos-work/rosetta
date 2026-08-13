**Type:** tik — under `TOK-04`, working clause 3 (P4: *derive, else fence, else declare prose-under-review*).

# iter-57 — clause 3: fence the map's CITATIONS, not just its membership

## Phase A — the measurement that named the gap

Both clause-3 fences were run against the shipped map **before any edit**, at platform `0dab54d`:

```
platform_alignment_guard: OK — platform-migration-status.md and repos.yml agree in both directions.
anchor-construct-guard:   OK — every resolvable anchor names a construct   (113 resolved)
```

**Both GREEN, over a map with five claims iter-55 had already falsified by hand.** Neither guard is
broken; they fence properties other than the one that failed:

- `platform_alignment_guard`'s assertion **D checks only that the evidence cell is non-empty.** A cell
  full of citations to lines that no longer exist is, to D, indistinguishable from a correct one.
- `anchor_construct_guard`'s `_QUALIFIED` regex requires a `/` in the path or a `.md` suffix. It was
  narrowed deliberately (its docstring records the 134-findings-all-ports over-match that forced it) —
  but the consequence had never been measured **for this file**.

Measuring it:

| citation class in the map | count | seen by a fence? |
|---|---|---|
| path-qualified (`app/main.go:573`, `storage/terraform/main.tf:19`) | 15 | yes |
| bare platform file (`docker-compose.yml:90`, `repos.yml:18-20`, `common.yml:22`) | 25 | **no** |
| bare continuation (`:178`, `:161`) | 12 | **no** |

**37 of 52 citations — 71% of the map's evidence — sat in a class no fence could see**, and both dead
citations iter-55 found by hand were in it. *"A fence over membership says nothing about prose"* was the
routed finding; this is **why**, and it is not a wording problem.

## Phase B — assertion F, and the draft that had to be thrown away

The first draft asserted *"the subject is named in the cited path or on the cited line."* Watched RED on
the real map it returned **22 findings — 7 of them its own false positives** (`app` cites the `backend:`
key; `jobsimulation` cites a line saying `jobsim`; `cms`/`roadrunner` cite a `repos.yml` header comment
that is *about* them without naming them; `postgresql`/`redis` cite `common.yml`).

Narrowing until only the known-bad ones fired was available and was **refused** — that is §5 Trap A,
fitting a fence to the answer key. The rule was replaced with one **the file itself defines**: for
`docker-compose.yml`, the cited line must sit inside the compose block of the row's own service, with
block boundaries *and* repo→service aliases parsed out of compose — `context: ${APP_BUILD_CONTEXT:-../app}`
is how the guard learns `backend` **is** the `app` repo, rather than being told.

**22 → 8 findings, 7 → 0 false positives.** Every survivor verified against the platform independently:

| finding | cited | actual | verified |
|---|---|---|---|
| `roadrunner` `JUDGE0_BASE_URL` | `:56` | `:59` | `grep JUDGE0` |
| `storage` service | `:90` (`volumes:`) | `:102` | block scan |
| `messenger` service | `:141` (`VERSION: dev`) | `:156` | block scan |
| `next-web-app` service | `:211` (`- .env`) | `:228` | block scan |
| `studio-desk` service | `:180` (`- REDIS_STREAMS_INDEX=4`) | `:197` | block scan |
| `customerio-sync` git context | `:121-123` (`- RPC_PORT=8301`) | `:136-138` | `grep context: git@` |
| `gotenberg` image | `:238-239` | `:255-256` | `grep gotenberg/gotenberg` |
| `gotenberg` profile | `:251` | `:268` | `grep profiles:` |

All eight resolve to **real, non-blank lines** — which is why `anchor_construct_guard`'s blank/closing-
delimiter classifier reports GREEN on every one of them, and why existence was never the property to check.

## Phase C — repair, and the claims the platform actually invalidated

Beyond the line numbers, `0dab54d` (*"run without the standalone storage; rename graphql -> core"*)
falsified the substance of two rows and one narrative line:

- **`storage`** — `fresh local stack: live-standalone` was **false**. The v9.0 fold has LANDED in compose:
  `STORAGE_RPC_ADDR` deleted, `STORAGE_S3_BUCKET`/`_PUBLIC_BUCKET` added, service moved to
  `profiles: [storage-legacy]`, kept startable only for rollback (two writers on one bucket otherwise).
  `make up` no longer starts it. → `live-standalone (opt-in profile)`.
- **`messenger`** — *"not started by the default `graphql` profile"* was false twice over: the profile is
  now **`core`** (`Makefile:10` `PROFILE ?= core`), and `0dab54d` also dropped messenger from `all`,
  because `backend` now consumes messenger's **own** Redis consumer group.
- **`gotenberg`** — *"default `graphql` profile"*, same rename.
- **§5's narrative** — *"v9.0 `storage` + `messenger`, PR #1103 open"* → its compose half has landed.

Three citations were **real drift the fence did not name** (`messenger`'s `:178`, `:159`, `:161`): all
inside messenger's *own* block, so the block rule passes them. Repaired by hand from derived line numbers,
and recorded as a stated limitation rather than left for a future reader to discover (`D-M257x-57-5`).

One survivor was **correct**: `roadrunner` cites `docker-compose.yml:59`, inside `backend`'s block,
because the claim *is* that Judge0 moved onto backend. Rather than loosen the rule, it gained a stated
clause — a cross-block citation is legal when the cell **names the block it points into** — held honest by
an inverted mutant that fires when the declared block is the wrong one (`D-M257x-57-2`).

## Phase D — GREEN, and the coverage the map now states out loud

```
platform_alignment_guard: assertion F resolved 49 citation(s) — 20 subject-checked,
                          28 range-only, 1 outside any service block; 0 unresolvable
platform_alignment_guard: OK — ... agree in both directions.                      RC=0
```

Reach is printed on **every** run, GREEN or RED, and a run that subject-checks nothing **refuses (exit 2)**
rather than reporting the map clean. The map's §4 now carries the standing table P4 asks for — what is
derived+fenced, what is fenced, what is resolution-and-range only, and what is explicitly
**prose-under-review** — because the third category only works if it is visible.

**Tests:** 30/30 in `test_platform_alignment_guard.py` (was 19), including the no-op control that must
survive, an inverted alias mutant, an inverted wrong-block mutant, the exit-2 positive control, and the
port/URL immunity case.

## Pre-registration, graded

| # | prediction | verdict |
|---|---|---|
| 1 | F names ≥ 2 findings on the unrepaired map | **HELD** — 8 |
| 2 | F names more than the 2 found by hand | **HELD** — 8 vs 2; the hand reading's recall on this file was **25%** |
| 3 | the repair is ≤ 3 rows + the narrative line | **REFUTED** — 7 rows + the narrative line |

Prediction 3's refutation is the useful one: the routed handler said *"5 sites in rows `:75`, `:76`,
`:180`"*, and the real surface was **more than twice that**, in rows nobody had looked at. A hand reading
scoped the repair to where the hand reading had looked.

---

## Close — 2026-08-03

**Outcome:** clause 3's prose is **repaired and now fenced**. Assertion F ships in
`platform_alignment_guard.py`, was watched **RED (8 findings) then GREEN**, and closes the class that made
the clause-3 downgrade possible: 71% of the map's citations were invisible to every existing fence. The
hand reading that routed this work had found **2 of 8**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: **y** (the user paused the session) — (5) cap-reached: n — (6) protocol-stop: n — Outcome: exit-4

**Gate reading at close, against platform `0dab54d`** (P3 re-checked at close, 17:21:30Z — origin
**unchanged**, so the measurement stands rather than being invalidated by construction):

| clause | reading | basis |
|---|---|---|
| 1 — 3 cold cycles green | **MET** | iter-56 evidence at these same refs; `demo-1` still up, untouched |
| 2 — full Playthrough suite | **MET** | `passing=30 failing=0` at these same refs |
| 3 — the migration-status map | **MET (this iter)** | 8 dead citations repaired; assertion F watched RED→GREEN; 20 compose citations now fenced; the un-fenceable residue **declared** prose-under-review in §4 |
| 4 — zero writes to a dropped schema | **MET** | unchanged; guards re-run |
| 5 — KB-fidelity | **NOT MET** | untouched; not re-cut |

**4 of 5**, up from 3 of 5. Clause 5 remains the only open clause and is unchanged — not re-cut, not
narrowed, not deferred.

**Decisions:** `D-M257x-57-1` … `D-M257x-57-6` (`iter-57/decisions.md`)

**Side-deliverables:** none. Every edit was inside the planned three-line scope.

**Routes carried forward:**
- `FIX-M257x-iter57-within-block-drift` → F catches cross-block drift only; `messenger`'s `:178`/`:159`/
  `:161` were real drift inside their own block and were found by hand. A within-block rule would need to
  assert what the line *says*, which is the line this fence family does not cross.
- `CHECK-M257x-iter57-anchor-guard-bare-class` → the same 37/52 blind spot may exist across the wider
  corpus, not just this map: `anchor_construct_guard` cannot see any bare `file:line` citation. **Measure
  the class corpus-wide before deciding whether it is worth fencing.**
- `FIX-M257x-iter56-app-ref-moved` → **still open, and now the next iteration's first act** (P3). `app`
  `v1.365.0 → v1.366.0`; not taken here because it changes what a bring-up consumes and `demo-1` is live
  clause-1/2 evidence.
- Unchanged and still open: `FIX-M257x-iter56-assignment-flake`, `FIX-M257x-iter56-preflight-fails-late`,
  `FIX-M257x-iter56-evidence-gitignore`, `CHECK-M257x-iter56-directus-race-uncertified`,
  `CHECK-M257x-iter56-stale-autoverify-twin`, `FIX-M257x-iter55-stranded-demopatch-revert`, the 81 drift
  sites / 21 files, `FIX-M257x-iter53-union-set` (**a user decision, `D-M257x-53-5`**),
  `FENCE-M257x-iter54-refs-block`, `CHECK-M257x-iter52-second-ai-manager`, RF-2/3/7–13, root `CLAUDE.md`,
  `CHECK-M257x-iter38-ai-act-classification`.

**Lessons:**
1. **A fence being GREEN is a statement about the fence, not about the file.** Two guards passed a map
   with five known-false claims. Neither was broken and neither was mis-configured — they simply fenced
   other properties. *"Is it fenced?"* is not a yes/no question; the answer is a coverage number, and
   until this iteration nobody had computed it for this file.
2. **Compute a fence's blind spot on the actual artifact, not from its docstring.**
   `anchor_construct_guard`'s narrowing is well-reasoned and well-documented, and its consequence here —
   71% of the map invisible — is nowhere in that reasoning, because the reasoning was about the corpus at
   large.
3. **When a fence's first draft has a false-positive rate, replacing the RULE beats tuning the
   THRESHOLD.** The 22-finding draft could have been narrowed to 8 by hand in minutes. Deriving the rule
   from compose's own block structure got the same 8 with zero false positives *and* an alias the guard
   was never told about — and it will still be right after the next compose edit, which a tuned threshold
   would not.
4. **A hand reading scopes the repair to where the hand looked.** The routed handler named 5 sites in 3
   rows; the mechanized reading found 8 in 7 rows, plus 3 more the fence structurally cannot see. This is
   the same recall problem iter-50 measured for corpus readings, reproduced on a single file.
