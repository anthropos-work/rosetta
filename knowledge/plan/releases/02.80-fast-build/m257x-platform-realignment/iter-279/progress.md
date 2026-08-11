**Type:** tik — under `TOK-08`.

Opened 2026-08-11. Pre-registrations sealed in this iter's first commit.

## Phase A/B — the source, measured then repaired

`clone_pin_guard.py`'s header carried two false claims, both of which the corpus had copied:

| claim | measured at rext `8e2974f47` |
|---|---|
| *"**Three arms**, all mechanically decidable"* | **four** — arm D shipped at iter-257 |
| *"`DEMO_ADVANCE_CLONES=pinned` checks each clone out at the ref **it** names"*, of the canonical pin | reads `PIN_FILE="$DEMO/clones.pin.json"` (`ensure-clones.sh:184`), **not** `CANONICAL_PIN` (`:206`) |

**PR-2 HELD** — a repo-wide sweep for arm-count claims about this fence returns exactly one wrong site.
Six other modules say *"three arms"* or *"four arms"* about **themselves**, correctly.

Arm D is now documented where the other three are, with its asymmetry stated because that is the part a
reader cannot infer: **a phantom key FAILS, a differing value is DISCLOSED, an absent copy is NOT-RUN
with its reason printed.**

## Phase C — PR-1, and the two the vocabulary caught

**PR-1 HELD: 34 passed before, 34 passed after** — no test changed verdict, so the docstring carried no
assertion and this is a documentation-of-code change, as declared.

But the repair tripped `TheNounVocabularyIsMeasuredNotAssumed` **twice**, and the two were answered
differently (`D-M257x-279-2`): `numerators` → `_NOT_NOUNS` on the documented `iters`/`series` precedent;
`minutes` → **REFUSED**, because a duration really is a measurement noun and admitting it would make
every `N minutes` in the repo a tracked literal — a repo-wide behavioural change, routed rather than
made in passing to unblock a commit.

**And the first repair of the second finding re-created it** by quoting the offending clause verbatim
while explaining it — **the third consecutive iter caught this way** (`D-M257x-279-3`).

**This also exposes a defect iter-278 shipped:** the `numerators` finding came from *iter-278's own*
`claim_census_guard` docstring, which went out having re-run only the arm its edit obviously touched.
The whole-section run it needed had already finished — **against the pre-edit tree.** That is why this
iter ran the full suite *before* its last commit rather than after its last edit.

## Phase D — the loop, and one more closed-fix-published-as-open

Whole `stack-core` section, venv **outside** the tree, **on the tree that was committed**:

    2219 passed, 1 skipped, 598 subtests passed in 2105.13s (0:35:05)   CONTENDED

**Neither 35m05s nor iter-278's 56m30s is a baseline** — same box, same suite, same day, ~1.6× apart.

Shipped: rext **`e64a3cd3b`**, tagged **`fast-build-m257x-iter-279`**, pushed, **tag verified on origin**
(`git ls-remote --tags`); `stack-demo/rosetta-extensions` advanced; `.agentspace/rext.tag` bumped to it;
all four claims re-verified **at the new sha before** re-pointing (34 test functions · `PIN_FILE` at
`:184` · *"Four arms"* present · the studio hardcode absent from live code); corpus re-pointed
`8e2974f47` → **`e64a3cd3b`** at 3 sites.

**And re-pointing surfaced the same defect class one more time, in iter-278's own prose.** The §8 note
read *"the guard's own docstring **still** says three … a rext edit, **deliberately not spent** here"* —
false the moment this iter spent it. **That is the third instance in two iters of a landed fix being
published as open backlog** (the `cms.md` studio-fetcher block, `ROUTE-M257x-265`, and now this), and the
first where the stale sentence was written by the immediately preceding iter. Repaired with the closure
and the retracted cost stated.

`guard_family` **30 GREEN · 0 RED · 5 not-run** — **PR-3 HELD**.

## Close — 2026-08-11

**Outcome:** `FIX-M257x-278-clone-pin-guard-docstring-says-three-arms` **CLOSED** — the source of the
false claim iter-278 repaired in the corpus is now fixed in the tooling, so it cannot be copied forward
again. `D-M257x-278-6`'s stated reason for deferring it is **retracted**, because iter-278 paid that
exact cost hours later for a different fix. Whole section **2219 passed / 0 failed**; family **30 GREEN ·
0 RED**. **Clause 5 NOT met and no `P` is claimed.**
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (2 tiks) — (6) protocol-stop: n — (7) budget-exhausted: **y** — Outcome: **exit-7**

**Decisions:** `D-M257x-279-1` (fix the source, and retract the reason for not doing it),
`D-M257x-279-2` (a duration is a measurement noun — and that is not a change to make in passing),
`D-M257x-279-3` (quoting the defect committed it, for the third consecutive iter),
`D-M257x-279-4` (the commit message was mangled by the shell; recorded, never force-fixed).

**Side-deliverables:** none.

**Routes carried forward:**
- **`ROUTE-M257x-279-durations-are-unclassified-measurement-nouns`** — `minutes` is a genuine measurement
  noun in neither vocabulary list. Admitting it to `_MEASURED_NOUNS` makes every `N minutes` a tracked
  literal; that is plausibly right and is a repo-wide behaviour change. **Also exposes the arm's reach:**
  it scans `stack-core/` only, and `~20 minutes` sits un-graded in three `dev-stack`/`demo-stack` test
  files today — the vocabulary's zero is a zero *over `stack-core`*.
- **`ROUTE-M257x-278-thirteen-unpinned-rext-anchors-are-on-undecidable-clocks`** — unchanged.
- **`FIX-M257x-278-census-substrate-tests-hardcode-an-absolute-ROOT`** — unchanged, and it bit again this
  iter: the noun-vocabulary arm cannot be run against a control tree either.
- **`ROUTE-M257x-278-rext-tag-SoT-was-six-iters-stale-unnoticed`** — bumped twice in two iters now;
  nothing asserts it.
- The **fence half** of `ROUTE-M257x-277` — unchanged, deliberately.
- **Clause 5's semantic reading is still unmeasured** (last: iter-131, `P = 29 / N = 47`, a floor).
- Unchanged: `ROUTE-M257x-274-successor-half-is-uncovered`, `ROUTE-M257x-274-tie-order-is-unstable`,
  `FIX-M257x-269`, `ROUTE-M257x-270-directus-consumer-cms-key-outlived-its-rollback-path`,
  `FIX-M257x-266`, `FIX-M257x-265`, `ROUTE-M257x-h59`, `ROUTE-M257x-h65`.

**Lessons:**
1. **Repair the source or the claim comes back.** The corpus sentence iter-278 retracted was copied out
   of a docstring; fixing only the copy leaves the generator running.
2. **A deferral's stated COST is a claim, and it can be refuted by the deferring iter itself.** iter-278
   declined this on a cost it then paid the same session. Deferral reasons should be written so they can
   be checked, and retracted in place when they fail.
3. **These fences read prose, so an explanation must DESCRIBE the offending construct, never reproduce
   it.** Three consecutive iters have now been caught by their own footnotes.
4. **Run the whole suite against the tree you are about to commit, not the tree you started from.** A
   green from before the last edit grades a tree that no longer exists — which is how iter-278 shipped
   the finding this iter opened with.
