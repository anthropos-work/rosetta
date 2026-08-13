**Type:** tik — under `TOK-05` (*stop repairing claims; fence the predicates under them*), step 2 of
its ordering: **fence → citations → map state → read**. This is the citations step, carrying
`FIX-M257x-iter63-app-citation-residual`.

# iter-68 — three guards with no ref, and a fold that landed while we were counting

## Phase A — re-derive the class before repairing it (the briefing's own instruction)

Two numbers moved before a single repair was made.

**The routed 68 is 64.** Re-running iter-63's *own* enumerator against today's corpus at the same
app ref: **123 sites / 96 distinct / 22 files — 32 mainline + 64 non-mainline**, against iter-63's
recorded 104 / 86 / 22 = 18 + 68. §5 rule 34 says a corpus repair moves the corpus's own line
numbers; **this is its sibling — a corpus repair also enlarges the corpus's own citation class**
(`D-M257x-68-3`). iters 63–67 wrote ten net-new citations, six of them in the two-sided `mid-fold`
row, and every one of those six died with the row.

**And `app` moved.** Origin/main advanced **56 commits to `9d00a313` v1.367.0** at 10:56Z that
morning, mid-iteration; the demo's clone is pinned to `b948604` v1.366.0.

## Phase B — §7 rule 4b applied *before* deciding anything (`D-M257x-59-3`)

Measured read-only (`git show`; no checkout, no risk to the green stack), same 96-citation universe,
graded at both refs:

| binding | ref | HELD | MOVED | GONE | DEAD | UNNAMED |
|---|---|---|---|---|---|---|
| pooled | `b948604` | 42 | 25 | 7 | 4 | 18 |
| pooled | `9d00a313` | **17** | 49 | 8 | 4 | 18 |

**25 of the 42 citations that hold at the pin break at origin HEAD** — 60%, in one working day. The
gate reads *against platform @ **origin HEAD***, and a cold `make init` clones `app` at its main, so
the adjudication ref for the repair is `9d00a313`. Repairing against the pin would have bought 45
claims that were already false.

## Phase C — and then the precondition nobody had seen (`D-M257x-68-1`)

The repair could not be evaluated, because **three guards were reading a different file than the one
being adjudicated against, and none of them said so.**

| | at `b948604` (the demo's pinned build ref) | at `origin/main` `9d00a313` |
|---|---|---|
| `STORAGE_RPC_ADDR` — G6 | **mid-fold**, 6 app read sites | **unconfigured**, no reader |
| `app/main.go` — assertion F | **1361** lines | **1569** |
| `app/internal/storage/service.go` | absent | present |
| the corpus — `anchor_construct_guard` | **RED, 4 findings** | **GREEN** |

Same guard, same corpus, opposite verdicts. **P1 — state your refs — is not satisfied by a number
whose ref is a checkout.** All three now resolve *and* read at a named ref (`auto` prefers
`origin/main`, falls back to `HEAD`; `worktree` by name; a caller-named ref that does not resolve is
UNMEASURED, never substituted), and all three print the provenance. Two corollaries were live
defects, not hypotheticals: **existence is decided at the same ref the content is read at** (a file
*born* in the advance read as `unresolvable head 'app'`), and **an untracked worktree file is in no
ref** (the box's untracked `studio/` — three `app/studio/…` citations belong to studio-room, not to
`app` at any commit).

## Phase D — the repair, adjudicated against platform artifacts

**Release 09.00 "support-in-app" has landed** (`D-M257x-68-2`). Storage and messenger are both
`service_desired_count = 0` in prod, both served in-process by `app`, both still cloned and still
startable as rollback paths. `app` does not merge messenger's handlers onto its own subscribers — it
**takes messenger's Redis consumer group over**, so the cursor survives and there is no gap.

Repaired: the map's `app` / `storage` / `messenger` rows (the `app` row now names **six** in-process
domains with six distinct wiring call sites, each verified to land on its construct at `9d00a313`),
`storage.md`, `messenger.md`, root `CLAUDE.md`, and the protocol's own `mid-fold` corollary — which
was standing in the present tense about a state that no longer has an instance.

**The eighth vocabulary token lived four iterations.** Built at iter-64 for `storage`; its only
instance was gone by iter-68. The token stays and the map now records that **no row carries it** —
a different statement from the token not existing, and the honest one.

**And a defect class the fence cannot see** (`D-M257x-68-5`): eight service docs carry a
`* **Profile**: …` bullet, which is none of G1's three constructs. **Seven of the eight were wrong** — and **all seven named `graphql`**, the profile `0dab54d` renamed `core` (re-derived per bullet; the first draft of this line said *five*). Two of them also named a profile for a service with no compose entry at all; `storage` also named a `storage` profile that never existed. Only `sentinel.md`'s was right. All seven repaired from the artifact.
**Seventh time in this milestone that a GREEN reading turned out to be a reach limit.**

## Phase E — the routed CHECK, answered (`D-M257x-68-4`)

`CHECK-M257x-iter63-quoting-a-retired-token` fired on my own edit — the third instance in three
iterations, and this time on the sentence *recording the rename*. It is **not a policy question**:
G1's negation discriminator was right and its **window was one line**, while the denial wrapped, with
a blockquote marker between the particle and the phrase. **The second window bug of this milestone
wearing a policy's name** — and `D-M257x-63-1`'s lesson had already said so in those words.

Window widened to the block (adjacency unchanged at two words); markers stripped as layout. The one
remaining site was a *postfix* denial and was **rephrased to the prefix form rather than taught to
the fence** — fitting the rule to that sentence is §4 Trap A.

## Phase F — gates

| gate | result |
|---|---|
| `platform_predicate_guard` (7 assertions) | **OK** — 99 profile sites / 8 tokens; app consumer side measured **@ `origin/main@9d00a31`** |
| `platform_alignment_guard` | **OK** — 74 citations resolved, **0 unresolvable** (was 2); and **4-findings RED** under `CITE_REF=worktree`, which is the demonstration |
| `anchor_construct_guard` | **OK** — 124 resolved; **4-findings RED** under `CITE_REF=worktree` |
| `markdown_structure_guard` · `corpus_index_guard` | OK — 112 files / 84 indexed docs |
| `tests/test_platform_predicate_guard.py` | **77** (was 67) |
| `tests/test_platform_alignment_guard.py` | **41** (was 36) |
| `tests/test_iter45_mechanical_fences.py` | 36 |
| mutation battery | **6 mutants across 3 guards, all caught** — ignore-the-ref (5 RED / 3 RED / 4 RED), silent-fallback (1), no-marker-strip (1), drop-adjacency (2), line-window (2). **One mutant SURVIVED its first test** — the existence check, whose test exercised `_exists_at_ref` directly instead of through `resolve_citation`; re-written against the call site with a file that lives at a ref and in no checkout, then RED |
| guard runtime | reading at a ref shells out per citation: `anchor_construct_guard` went **~1 s → 10.9 s**, multiplied by every mutant the battery runs. Memoised the ref lookups — **1.8 s**. A guard slow enough to skip is a guard that stops being run |
| `stack-core` suite | **701 tests, 1F** — the perishable iter-48 fixture, the single expected failure. Back at baseline (was 682/1F at iter-67; +19 from this iter's three test classes). **The first run was 2F**, and the second was a real catch: `service_doc_status_fence` — *a doc whose service the map calls merged must OPEN by saying so* — fired on `storage.md`, whose banner said the CALLERS had merged and never that storage had. Fixed, not waived |

## Close — 2026-08-04

**Outcome:** the routed citation class re-derived (**64, not 68** — a corpus repair grows its own
citation class), §7 rule 4b's delta measured across an app advance that landed mid-iteration
(**25 of 42 holding citations break at origin HEAD, 60% in one working day**), and the precondition
that made the repair evaluable at all: **three guards were adjudicating against a checkout and none
of them said so** — the same corpus is GREEN at origin HEAD and 4-findings RED at the pinned build
ref. Release 09.00 landed while we were counting: **storage and messenger are both folded**, prod
compute stopped on both, and the eighth vocabulary token built four iterations ago now has no
instance. Seven of eight `* **Profile**:` bullets were wrong in a construct no fence reaches. The
routed retired-token CHECK turned out to be **the second window bug of this milestone wearing a
policy's name**.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET — 4 of 5, unchanged; clause 5 is still graded only by a reading that returns zero.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n (**not** fired: the app advance did not invalidate an attempt, because rule 4b caught it *before* the repair was spent; the trigger stays at occurrence 1 of 2) — (4) user-blocker: n — (5) cap-reached: n (1 tik of 5) — (6) protocol-stop: **y** — Outcome: exit-6.
The harden cadence is at **10 tiks** with this one (58, 60–68; iter-59 was a tok), which is the threshold, and the milestone's orchestration owns that pass — the iter loop never spawns it. Closing here so `/developer-kit:harden-mstone-iters` runs before the next tik, which matters more than usual: this iteration shipped **three** guard changes and a memoisation, and the harden pass is what audits AST/call-site shapes like these.
**Decisions:** `D-M257x-68-1` (a guard that resolves against a checkout has no ref, and its verdict
is not a measurement), `D-M257x-68-2` (release 09.00 landed; `mid-fold`'s only instance lived four
iterations), `D-M257x-68-3` (the routed figure was stale when routed, and our own repairs staled it),
`D-M257x-68-4` (the retired-token CHECK is a window bug, not a policy question), `D-M257x-68-5` (the
`* **Profile**:` bullet — a defect class in a construct no fence reaches).
**Side-deliverables:** none — the three guard changes are the iter's planned scope, not side work:
the repair was not evaluable without them. The memoisation is a defect of this iter's own making,
repaired in it.
**Routes carried forward:**
- `FIX-M257x-iter63-app-citation-residual` — **still open, re-scoped.** The class is now **105
  distinct citations**; the fold half landed here. The non-fold remainder (`ai_architecture.md`,
  `external_services.md`, `backend.md:39`, `coursebuilder.md`, `cms.md`, `skillpath.md`,
  `shared_libraries.md:79`, `jobsimulation.md`, `content-stories-routes.md`) is **B2**, unrepaired.
- `FENCE-M257x-iter68-profile-bullet` — widen G1 to the `* **Profile**: …` bullet construct. Design
  is in `D-M257x-68-5`; 8 sites, all repaired, so it locks a correct state.
- `FENCE-M257x-iter68-citation-resolution` — the class-level fence TOK-05 actually wants: re-resolve
  every corpus citation into a clone at the adjudication ref and grade HELD/MOVED/GONE/DEAD/UNREACHED.
  The three guards now share the primitive; the assertion does not exist yet.
- Unchanged: `FIX-M257x-iter53-union-set` (**PENDING USER DECISION**) ·
  `FIX-M257x-iter56-assignment-flake` (**NOT DECIDED** — needs a failure *rate*) ·
  `CHECK-M257x-iter38-ai-act-classification` (needs an owner outside this milestone) ·
  `CHECK-M257x-iter57-anchor-guard-bare-class` · `FENCE-M257x-iter54-refs-block` ·
  `FIX-M257x-iter57-within-block-drift` · `CHECK-M257x-iter58-derive-preregistrations` ·
  `CHECK-M257x-iter52-second-ai-manager` · `-cold-daemon-registry` · `-grep-vs-failclosed` ·
  `-empty-stdout-class` · `-baseline-refs` · RF-2/3/7–13.
- **Closed here:** `CHECK-M257x-iter63-quoting-a-retired-token`.

**Lessons:**

1. **A guard's reference is part of its verdict, and it must be printed.** Three guards, one
   iteration, same defect: `read_text()` on a checkout. The failure is invisible by construction —
   both answers look like answers. If a fence resolves anything against a clone, the first question
   is *which ref*, and the second is *does it say so*.
2. **Measure the advance before you spend the iteration.** §7 rule 4b cost about fifteen minutes and
   redirected the whole iter: repairing 45 citations against the pinned ref would have been careful,
   correct-looking work that was false before it was committed.
3. **A routed count is a snapshot of a population you are still growing.** 68 → 64 with no platform
   movement at all in that term — our own six `mid-fold` citations were born and died between the
   routing and the repair. Re-derive at open; never repair to an inherited number.
4. **Write the lesson down and you will still not apply it.** `D-M257x-63-1` said *"a policy hole is
   often a window bug wearing a policy's name"*, and the very next routed policy question was a
   window bug. The lesson needs to be a **first check**, not a closing remark.
5. **Rephrase before you teach a fence English.** The one postfix denial could have bought a
   postfix-negation rule; it bought a rewritten sentence instead. Every discriminator added to
   accommodate one sentence is a hole the next hundred can walk through.
6. **A fence built one iteration after its class was repaired is GREEN, and that is the point** —
   but a construct a fence has *never reached* is not GREEN, it is unread. Seven of eight bullets
   were wrong and every fence in the corpus said OK.
7. **A fence you did not think about is the one that reads your work honestly.**
   `service_doc_status_fence` caught that `storage.md`'s opening banner announced *the callers*
   merging and never that **storage** had. I had written the fold three times in that file and still
   not said it in the one place a reader lands.
8. **Making a guard correct made it ten times slower, and that is a defect too.** Reading at a ref is
   a `git rev-parse` per citation; the anchor guard went ~1 s → 10.9 s and the mutation battery
   multiplies it. Memoised back to 1.8 s in the same iteration. Correctness that nobody waits for is
   not shipped.
