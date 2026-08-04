**Type:** tik — under `TOK-05` (*stop repairing claims; fence the predicates under them*), step 2 of
its ordering: **fence → citations → map state → read**. The citations step, carrying
`FIX-M257x-iter63-app-citation-residual` scope **B2**.

# iter-69 — the citation class was five defects wearing a count of ninety-six

## Phase A — re-derive before targeting, and check the ref before spending the iteration

Both re-derived at open, neither inherited.

**The class.** iter-63's own enumerator, re-run: **135 sites / 105 distinct / 22 files**, against the
briefing's 96 and iter-68's close of 105. It grew **96 → 105 during iter-68 itself**, with no
platform movement in that term.

**The ref (§7 rule 4b).** `platform` clone and origin both `0dab54d` — **level**. `app` clone
`b948604` v1.366.0, origin/main `9d00a313` **v1.367.0, 56 commits ahead**, tip dated 2026-08-04
10:56Z — **unchanged since iter-68 measured it eleven hours earlier**. So the adjudication ref is
settled and, for once, standing still.

## Phase B — the screen, watched RED before it was trusted

One mechanical shape: *a citation whose line CONTENT changed between the ref the corpus was written
against (`b948604`) and the ref the gate names (`origin/main`)*, with the **pin** read from the
corpus BLOCK — a one-line window is what made the last routed policy question turn out to be a
window bug (iter-68), and a pin routinely sits a sentence away on a wrapped line. It reuses
`platform_predicate_guard._REF_PINNED` verbatim, so the two instruments cannot disagree about what a
pin is.

| control | result |
|---|---|
| inverted mutant — pin detection disabled | candidates **2 → 61** |
| inverted mutant — content comparison neutered | candidates **2 → 0**, HELD 62 → 123 |
| **no-op positive control** (comment-only edit) | **byte-identical output — SURVIVED** |

## Phase C — and the answer was not 64

| class | n | defect? |
|---|---|---|
| identical at both refs | 62 | no |
| drifted, but the block **names its ref** | **59** | **no — a measurement** (`D-M257x-69-1`) |
| drifted and **UNPINNED** | **2** | **yes** |
| file absent at the ref | 3 | mis-rooted, not dead |

**Five, not sixty-four.** Verified rather than assumed: every ref-pinned mainline citation in the B2
files lands exactly on its construct at `b948604` — `backend.md:39`'s seven all do — and lands on an
unrelated `web.NewServer` argument at `origin/main`. The corpus is internally honest. **A pass that
"repaired" the 59 would have moved 59 correct claims onto a ref that moves again next week**, and
grown the class by every line number it rewrote.

## Phase D — the defects the screen could not see, found by reading (`D-M257x-69-2`)

Two of the three real defects were **structurally invisible** to the instrument that found the third.

**`shared_libraries.md:70` contributes ZERO citations to every count this class has ever reported.**
It names its antecedent as `` `app/main.go` `` — in backticks, one clause away — and then cites
`` `:1178` ``, `` `:1179` ``, … with no line number for the enumerator's `last_path` to latch onto.
Derived, the hidden class is **23 citations across 14 lines**. Six are on this one line, and
**five of the six were wrong at every ref**: the mux is `1187 / 1188 / 1196 / 1204 / 1213 / 1228`,
not `1178 / 1179 / 1187 / 1195 / 1204 / 1218`. Only JobSimulation was right — and `backend.md:39`,
which is pinned, had the correct set the whole time.

The same rule is **too greedy in the other direction**: `ai_architecture.md:96`'s bare `` `:15-17` ``
means *this document's* lines 15-17 (its own ⚠️ retraction), and inherit-the-last-path attributed it
to `app/internal/skillerai/ai.go`. One rule, a false negative in prose and a false positive in prose,
depending only on what preceded it.

The other two blind spots each hid a live defect: a path written without its clone root
(`anchor_construct_guard`: `unresolvable head 'internal' x6`) hid the Judge0 line, and the screen's
`app`-only universe hid a platform-rooted one.

**Repaired** — full adjudication in `decisions.md`:

- `shared_libraries.md:70` — the six handler line numbers, corrected + pinned, with the CMS
  handler's `if cmsRPCServer != nil` conditionality restored.
- `shared_libraries.md:79` — Judge0 `:118` → **`:123`** at `origin/main` (true at `b948604`,
  unpinned, false at the gate's ref); and `ROADRUNNER_RPC_ADDR` *"(`docker-compose.yml:118`)"*,
  which asserts the variable is in the compose when at `0dab54d` it is there **zero times**.
- `platform-alignment.md:872` — rule **32**'s own worked example, pinned. **Rule 33 applied to rule
  32's neighbour**, where it was needed and not written; the storage fold has since deleted both
  cited lines.
- `external_services.md` ×3 — `app/studio/**` is the **in-image** path from `anthropos-studio-room`,
  in no `app` commit at any ref. One scoping note; the claims were correct.

**The repair widened the fence, measurably.** `anchor_construct_guard`'s citations adjudicated at
`origin/main@9d00a31` went **43 → 49**: re-rooting `internal/…` to `app/internal/…` made six
citations reachable that the guard had been silently dropping.

## Phase E — G8, the construct no fence reached (`FENCE-M257x-iter68-profile-bullet`)

Eight service docs open with `* **Profile**: …`. iter-68 found **seven of the eight wrong**, all
seven naming `graphql`, while every fence read GREEN — the bullet is none of G1's three constructs.

**G8 is G7 inverted**: G7 reads a profile and checks the services beside it; G8 reads a **service —
from the doc's own file stem, derived** — and checks the profiles beside it. Three shapes, each
decidable against compose (`list` both-directions · `no-service` · `always-on`); anything else is
**UNREACHED, never an empty claim**.

Live: **8/8 reached, 0 unreached, `{list: 5, no-service: 2, always-on: 1}` — GREEN.** Green one
iteration after a hand repair is the right outcome and is **not** the evidence.

The evidence: **5 source mutants, all caught** (one-direction · shape-blind · stem-blind ·
always-on-unchecked · prose-graded-as-an-empty-claim), an **artifact** inversion that swaps two
services' compose membership and requires both correct bullets to become findings **naming the
mutated truth**, and a **no-op control that SURVIVED**. Every fixture is copied verbatim from the
live corpus and mutated by one fact — the harden-pass-16 rule, because **G3's reach was zero for
three iterations when its fixtures agreed with its pattern instead of with the corpus**.

## Phase F — gates

| gate | result |
|---|---|
| `platform_predicate_guard` (**8** assertions) | **OK** — G1 99 sites/8 tokens · G7 21/22 · **G8 8/8, 0 unreached** · G4 13 · G6 7 vars, 0 mid-fold @ `origin/main@9d00a31` |
| `platform_alignment_guard` | **OK** — 74 citations resolved, 0 unresolvable |
| `anchor_construct_guard` | **OK** — 124 anchors; **origin/main reach 43 → 49** |
| `markdown_structure_guard` · `corpus_index_guard` | OK — 112 files · 84 indexed docs |
| `tests/test_platform_predicate_guard.py` | **119** (was 108); `TestG8ProfileBullet` 10/10 |
| G8 mutation battery | **5 source mutants caught + 1 artifact inversion; no-op control SURVIVED** |
| `stack-core` suite | **753 tests, 1F** — `test_claim_twin_guard_iter48_answer_key::test_02…` , the perishable iter-48 fixture. **Baseline matched by IDENTITY**, not count (+11 from this iter) |
| `stack-injection` | **332 OK** (2 skipped) |
| `dev-stack` | **151 OK**, run **solo** |
| `demo-stack` · `stack-verify` | **not run — zero files touched in those sections.** `git status` in rext shows exactly two modified files, both `stack-core/`. Stated rather than claimed |

## Class re-measured AT CLOSE (the briefing's instruction)

| | open | close |
|---|---|---|
| citation sites | 135 | **141** |
| distinct | 105 | **109** (mainline 49 / non-mainline 60) |
| files | 22 | 22 |
| **UNPINNED-MOVED (the gate-relevant residual)** | **2** | **0** |
| file-absent (mis-rooted) | 3 | 3 — **now documented as in-image paths** |

**The class grew by 4 distinct while being repaired**, entirely from this iter's own corrected line
numbers. §5 rule 34's sibling, measured on both ends this time.

And re-measuring at close is what surfaced `D-M257x-69-5`: the enumerator **crashed** — its
`sorted({(path, a, b)})` compares `None` with `int` when one path carries both `X:N` and `X:N-M` at
the same start line, which no corpus contained until this iter wrote `app/main.go:1212-1214` beside
`app/main.go:1187`. **Every count this class has ever reported came off an instrument that was
correct by luck.** Fixed.

## Close — 2026-08-04

**Outcome:** B2 re-derived at the gate's ref and **it was 5 defects, not 64** — 59 of the "residual"
are ref-pinned measurements that each resolve exactly at their own pin, and repairing them would
have induced drift rather than removed it. The 5 are repaired; the **unpinned-and-moved class is 0**.
Reading found what the screen structurally could not: **`shared_libraries.md:70` contributes zero
citations to every count this class has ever reported** — its antecedent is backticked one clause
away with no line number — and **five of its six handler line numbers were wrong at every ref**,
contradicting the pinned-and-correct `backend.md:39`. **G8** ships for the `* **Profile**:` bullet,
the seventh reach limit of this milestone: **8/8 reached, 0 unreached, 5 source mutants + 1 artifact
inversion caught, no-op control survived.** The repair widened `anchor_construct_guard`'s origin/main
reach **43 → 49**. And the enumerator behind six iterations of load-bearing counts **crashes** on a
shape the corpus was one edit away from containing.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — 4 of 5, unchanged; clause 5 is still graded only by a reading that returns zero.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n (no attempt was
invalidated; the app advance was already measured at iter-68 and had not moved — trigger stays at
occurrence 1 of 2) — (4) user-blocker: n — (5) cap-reached: n (1 tik of 5) — (6) protocol-stop: n —
Outcome: **continue**.
**Decisions:** `D-M257x-69-1` (a ref-pinned citation is a measurement; the defect class is the
unpinned half — 5, not 64), `D-M257x-69-2` (the screen's reach, stated, because it hid two of the
three real defects — 23 citations across 14 lines are invisible to every count this class has
reported), `D-M257x-69-3` (the three defects, adjudicated against platform artifacts; the repair
widened the anchor guard's reach 43 → 49), `D-M257x-69-4` (G8 — the per-service profile bullet, G7
inverted), `D-M257x-69-5` (the enumerator behind every count in this class was correct by luck).
**Side-deliverables:** none. The screen and the enumerator fix are instruments this iter's own
measurement required, not unrelated work.
**Routes carried forward:**
- `FIX-M257x-iter69-pathless-antecedent` — the **23 citations across 14 lines** whose antecedent is
  a backticked path with no line number. `shared_libraries.md:70`'s six are repaired here; the
  other **17 across 13 lines** are underived and unread. The derivation is a ten-line script and is
  quoted in `decisions.md` `D-M257x-69-2`. **This is a prerequisite for the graded read** — a
  reading over a class the instrument cannot enumerate measures the instrument.
- `FENCE-M257x-iter69-citation-antecedent` — teach the shared citation parser that a backticked path
  with no line number IS an antecedent for a following bare `:N`, **and** that a bare `:N` with no
  path antecedent in prose means *this document*. One rule, both directions; both were live.
- `FENCE-M257x-iter68-citation-resolution` — **still open**, and now scoped by measurement: the
  assertion is *"every citation is HELD at the ref its block names"*, with `_REF_PINNED` deciding the
  ref. The screen is the prototype; it is scratch, not a guard.
- `FIX-M257x-iter58-mainline-shift` — `shared_libraries.md:70` was one of its outstanding sites and
  is closed here; the rest stand.
- Unchanged: `FIX-M257x-iter53-union-set` (**PENDING USER DECISION**) ·
  `FIX-M257x-iter56-assignment-flake` (**NOT DECIDED** — needs a failure *rate*) ·
  `CHECK-M257x-iter38-ai-act-classification` (needs an owner outside this milestone) ·
  `CHECK-M257x-iter57-anchor-guard-bare-class` · `FENCE-M257x-iter54-refs-block` ·
  `FIX-M257x-iter57-within-block-drift` · `CHECK-M257x-iter58-derive-preregistrations` ·
  `CHECK-M257x-iter52-second-ai-manager` · `-cold-daemon-registry` · `-grep-vs-failclosed` ·
  `-empty-stdout-class` · `-baseline-refs` · RF-2/3/7–13.
- **Closed here:** `FENCE-M257x-iter68-profile-bullet`; `FIX-M257x-iter63-app-citation-residual`
  scope **B2** (residual 0 at the gate's ref, with a derived denominator).

**Lessons:**

1. **Count the class by PREDICATE before you budget for it.** "64 unrepaired citations" and "5
   defects" are the same set. The difference is one question — *does the block name its ref?* —
   and asking it cost fifteen minutes and saved a repair pass that would have made the corpus worse.
2. **A GREEN screen over a class your instrument cannot enumerate is not a measurement.** Two of the
   three real defects were structurally unreachable by the tool that found the third. The screen was
   mutant-verified and honest about what it grades; that is exactly why its *reach* had to be written
   down beside its verdict.
3. **The oldest construct in a document is the least examined.** `shared_libraries.md:70` has been
   wrong through every reading of this milestone because no enumerator could see it and no reader
   thought to check a parenthetical. **Six citations, one line, five wrong.**
4. **Re-measure at close and the instrument will tell you something about itself.** The enumerator
   crashed only because this iter's repair wrote the one shape it could not sort. Six iterations of
   counts came off it; none of them had a right to be trusted.
5. **A fixture must be copied from the corpus, not written to the pattern.** Stated in harden pass
   16, applied here by construction: every G8 fixture is a live bullet with one fact changed. The
   rule earns its place because the alternative failure — G3's zero reach — was invisible for three
   iterations and looked like health.
6. **When a repair makes a fence reach further, say by how much.** 43 → 49 at origin/main is the
   difference between "I fixed a citation" and "I moved six citations inside the guard's reach", and
   only the second is a durable claim.
