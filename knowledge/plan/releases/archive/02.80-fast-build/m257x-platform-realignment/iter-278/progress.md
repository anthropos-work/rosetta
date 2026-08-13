**Type:** tik — under `TOK-08` (census the mechanical classes; stop sampling them).

Opened 2026-08-10. Pre-registrations sealed in this iter's first commit (`fddc0fc2`), before any repair.

---

## Phase A — the census

`clone_drift_guard` at open: **RED**, one finding — `rosetta-extensions` at `0a8674e74`, **22 commits**
past the nearest of **12** cited shas, **44** citing sites. Guard family: **23 GREEN · 1 RED · 11 not-run**.

Three populations were enumerated, each with its denominator stated.

| population | measured | verdict |
|---|---|---|
| corpus sites citing an rext **sha** | **44** sites / **12** shas | **44 of 44 ref-scoped** (39 by the mechanical form test `(@\|at\|rext\|ref\|pinned\|tag)` immediately preceding; the other 5 scope in substance — *"the prior pin"*, *"landed at M226"*, *"iter-02 (…) rewrote all three sites"*, *"at `415240f` every line number below resolves to unrelated code"*, *"byte-exact in the pinned clone"*) |
| corpus citations of an rext **path** | **338** occurrences / **164** distinct paths (`rext_path_guard`'s own population, not a second parser's) | **all resolve at HEAD** |
| corpus sites citing a file the advance **touched** | **51** sites / **13** files (of 52 files touched) | **7 ref-pinned** (out of subject) · **13** unpinned with a line anchor · 31 path-only |

**PR-1 HELD, and it is the iter's central result.** `advance_impact_census`'s own contract puts a
block-pinned citation **out of subject** — §5 rules 41/44 make it true-at-that-ref, *"an advance cannot
falsify it"* — and `clone_drift_guard`'s docstring says D1 *"does NOT adjudicate truth-at-a-ref."* So the
44 sites were never a repair backlog, and **0 of 44 were renumbered.** This is iter-256's lesson honoured
rather than re-learned: that iter applied 27 "obvious" renumbering repairs and reverted them byte-for-byte,
**seven having moved *correct* citations onto comments.**

**PR-2 HELD.** One apparent dead path — `stack-verify/e2e/tests/probe-aireadiness-deeplink.spec.ts` at
`coverage-protocol.md:680` — is a **false positive of a path-must-resolve rule**: the sentence's own claim
is that the file *"was never committed — the file does not exist."* Graded for what the sentence claims,
not for what the path does. The crude first pass also flagged 24 `knowledge/…` sites that belong to *other
repos'* knowledge dirs; `rext_path_guard`'s population is the one quoted above.

## Phase B — what D1 actually asserts, and the direction problem

D1 is RED on a different proposition than the 44 sites: *a whole repo's worth of change has never been
looked at.* True — the 22 commits are this milestone's own iters 253–276, and every one of them wrote its
findings into `knowledge/plan/`, **never into `corpus/**`.** That is iter-277's Lesson 3 as a measurement.

**The 13 unpinned line anchors were measured and deliberately NOT repaired.** Comparing the cited line's
text at the last-cited ref (`d739952`) against HEAD: **0 of 13 held.** That reads like a 13-site repair
backlog and **it is not one** — the direction is undecidable mechanically, which the census module says in
its own docstring (*"the corpus's anchors are at MIXED refs … nothing mechanical can tell which one a given
citation is on"*). Two worked cases, one each way:

- `external_services.md:208` cites `gen_injected_override.py:698-699` and the **next line** pins it
  (*"both @ the demo's **pinned** rext `09d06070`"*). `block_ref` returns `ambiguous` because the pin is not
  on the anchor's own line. **Renumbering it would have falsified a correct ref-scoped claim** — exactly
  iter-256's seven.
- `safety.md:712` cites `up-injected.sh:1358-1385` for `bridge_bedrock_creds`, which is at `:1358` at
  `d739952` and `:1364` at HEAD — old clock. `media-substrate-spec.md:122` cites `:1364` for the *same*
  function — new clock. **Two sites, one function, two refs, no marker.**

`org-repos.md:195`'s `:321-364` matched **neither** ref (`directus_lines` opens at `:324` at HEAD, and
`:321` was a bare `"""` at `d739952`), so that anchor was already stale **before** this advance and is not
this iter's subject either way. Routed, not guessed — `D-M257x-122-5`.

## Phase C — the repairs (PR-3: **≥ 1** predicted, **2** found, both substantive)

**A — `corpus/services/cms.md` (clause-5 scope): a landed tooling fix published as live backlog.**
The block asserted, in the present tense, that *"the demo tooling still ENTERS this repo"* via a hardcoded
`_studio_repos="cms"` at `ensure-clones.sh:310`, and routed
`FIX-M257x-268-ensure-clones-hardcodes-cms-as-studio-fetcher` as **open**. Measured at HEAD: the hardcode,
the `make init-studio` special case and the cms *preference* are **all gone** (iter-270, `2833a64`);
`:314` calls `studio_consumer_names "$PLAT/repos.yml"` and **refuses the bring-up** when the set cannot be
derived; acquisition is a plain `git clone` for every consumer. The route is **closed**.

And iter-268's *"nothing is broken by it"* was wrong in a way worth carrying: iter-270 graded all **8**
platform-topology derivations in the demo bring-up path and found this one **failed OPEN** — an unreadable
`repos.yml` collapsed the consumer set to **`cms` alone**, dropping every live consumer and re-arming the
`/build/studio: not found` failure the phase exists to pre-empt. *A preference does not fail*, which is why
it survived four releases.

What survives its mechanism is preserved, per `D-M257x-129-4`: `stack-demo/` still carries **6** clones
`repos.yml` does not name (measured on disk: `cms`, `graphql-wundergraph`, `jobsimulation`, `messenger`,
`roadrunner`, `storage`; `stack-demo/cms/studio` still holds 18 entries), and
their remaining on disk is the measured, accepted state
`ROUTE-M257x-265-stack-demo-carries-six-dead-clones` **closed at iter-268** — whose deliverable was the
census, *nothing is deleted in this iter* being one of its sealed pre-registrations. **The first draft of
this block asserted that route open, in the corpus and in the ledger, and `route_disposition_guard` caught
it** — the same defect shape as iter-277's, one iter later: **a landed closure re-published as backlog.**

**B — `corpus/ops/platform-alignment.md` §8: the fence registry described a three-armed fence with 16
tests, and named the wrong file as the one the mechanism reads.** Three claims, each measured:

| claim as written | measured at rext `0a8674e74` |
|---|---|
| arms **(A) (B) (C)** | **four** — arm **D** landed at iter-257 (`2ff1547`), grading every `stack-*/clones.pin.json` **workspace copy** against the canonical |
| *"— 16 tests"* | **34** `def test_` functions (16 was correct at iter-222 `cdb87a1`, where the file was added; 23 at `d739952`) |
| *"`DEMO_ADVANCE_CLONES=pinned` checks each clone out at the ref **it** names"*, of the **canonical** pin | **false** — `pinned` reads `$DEMO/clones.pin.json`, the **workspace copy** (`ensure-clones.sh:184`), seeded from the canonical (`:206`) **copy-if-ABSENT** and, until iter-257, never reconciled |

The third is the substantive one, and its cost is measured, not hypothesised: **at iter-257 this box's
workspace copy named 11 repos to the canonical's 6** — the five phantom keys iter-222 removed, every one
with a directory on disk carrying a git checkout. **The fence that cleaned the canonical did not reach the
file the mechanism reads** — the same *"second registry one file over"* shape as finding A, in the same
iter, from the same advance.

**This document already contradicted itself, and the half that was right stayed right.** §7's advance rule
has said since iter-257 that `pinned` *"reads the **`stack-*/clones.pin.json` copy**, not the canonical
file."* Only §8's registry disagreed. The repair moves §8 onto §7's reading rather than inventing one.

## Phase D — re-measure

- `clone_drift_guard`: **RED → GREEN.** *"every cited clone's HEAD is a commit the corpus cites, and all 2
  go.mod-cited pin(s) match."*
- Guard family: **24 GREEN · 0 RED · 0 could-not-check · 11 not-run** (from 23/1). Re-run after the §5
  rule-80 addition: unchanged.
- **PR-4 HELD, graded by reading.** Three HEAD shas were written, and each sits on a **new-state** claim —
  *"the hardcode … are all gone"*, *"`pinned` reads `$DEMO/clones.pin.json`"*, *"**34** test functions"* —
  never on a sentence about the drift. That distinction is not decoration: this fence's own docstring books
  `FIX-M257x-iter107-drift-fence-satisfiable-by-prose`, having gone GREEN once because a paragraph
  *describing* the drift contained the sha. **Satisfying it that way was available here and was not taken.**
- **PR-5 HELD** — `route_disposition_guard` GREEN. The closed-route correction landed in `corpus/**`, which
  is not that guard's subject.

## Phase D2 — the suite, and the two failures it found were BOTH MINE

`stack-core` under pytest, from a venv **outside the tree** (§5 rule 80, written from this invocation):

    2 failed, 2217 passed, 1 skipped, 597 subtests passed in 3390.18s (0:56:30)   CONTENDED

**56m30s is not a baseline** — it sits alongside iter-277's **53m13s** on the same contended host, and
the corpus's cleaner-box figure (12m06s–27m34s) does not describe this machine.

**Both failures were caused by this iter, and that was PROVEN rather than assumed.** The two arms live
in `test_claim_census_substrate_m257x.py` and read `ROOT = Path("/Users/…/rosetta")` — a **hardcoded
absolute path**, so `ROSETTA_ROOT` cannot redirect them and the usual frozen-clone control is
unavailable. The derivation was therefore reproduced directly against both corpora
(`git clone --local --shared` of this repo at the pre-repair commit `fddc0fc2`):

| derivation | control @ `fddc0fc2` | live, post-repair |
|---|---|---|
| distinct-grain share | 291 of 707 | **292 of 708** |
| pair-grain share | 407 of 982 | **408 of 983** |
| `unresolved-wrong-repo-guarded` | 1 | **2** |

**The two arms partitioned my append correctly, and answering them the same way would have been the
error.**

- **The wrong-repo arm found a REAL DEFECT — in this iter's own prose.** I had written the derivation
  as `` `studio_consumer_names "$PLAT/repos.yml"` ``, and the census read `PLAT/repos.yml` as a path
  citation that *"names a directory path that exists nowhere under the repo or the clone set."* It was
  right: a shell expansion inside a code span is not a citation. **Repaired in the corpus** — the
  sentence now names the platform clone's `repos.yml` in prose — and the class went **2 → 1**. The
  comment was **not** bumped to 2.
- **The denominator arm was maintenance, not a defect.** With the prose fixed the numerators returned to
  **291 / 407** and only the denominators stood at **708 / 983** — because the append net-added one
  distinct line-pinned citation (`+demo-stack/ensure-clones.sh:314`, `+stack-core/lib/studio.sh:121`,
  `−rosetta-extensions/demo-stack/ensure-clones.sh:310`).

**This is the third consecutive firing of that arm on a corpus append rather than on a defect** (harden
pass 61: 293→291 / 410→407; harden pass 68: 706→707 / 981→982; iter-278: 707→708 / 982→983) and the
**second** time it has fired on the paragraph that describes it. The docstring now states the invariant
the three occurrences establish: *a corpus that gains a line-pinned citation moves these denominators
and nothing else — a numerator that moves without a parser change is the signal worth stopping for.*

`tests/test_claim_census_substrate_m257x.py`: **34 passed**, before and after the re-point.

## Phase D3 — the loop closed in this iter, which is the whole thesis

The docstring repair is an **rext edit**, and taking it advances the clone past the sha this iter had
just reconciled the corpus to. **That is not a reason to defer it — it is the coupling iter-277 named,
and the iter that creates the debt is the iter that pays it.** So the full loop ran here:

1. `rosetta-extensions` `8e2974f47` committed, tagged **`fast-build-m257x-iter-278`**, pushed, and the
   tag **verified on origin** with `git ls-remote --tags origin` — *tagging is not publishing*.
2. `stack-demo/rosetta-extensions` advanced to that tag (untracked report artifacts preserved).
3. All three claims re-verified **at the new sha before re-pointing**: 34 test functions; `PIN_FILE="$DEMO/clones.pin.json"` at `:184`; and `_studio_repos="cms"` surviving only as the **comment at `:310`** that records its removal, with `:314` holding the live derivation.
4. The corpus's three sha citations re-pointed `0a8674e74` → **`8e2974f47`**.
5. `.agentspace/rext.tag` bumped **`fast-build-m257x-iter-270` → `fast-build-m257x-iter-278`** — it was
   **already 6 iters stale before this iter touched it**, which is iter-01's git-ignored-pin root cause
   still live.

**No latent RED was left behind.** Advancing the clone without re-pointing would have re-armed D1 on the
next sync; re-pointing without advancing would have made the corpus cite a sha the clone is not at.

## Close — 2026-08-11

**Outcome:** `clone_drift_guard` **RED → GREEN**, earned by two substantive repairs rather than by
renumbering 44 sites or by prose about the drift. The census settled that the 44 sha-citing sites are
**ref-scoped and none was a defect**; the real debt was two claims the advance falsified — a **closed**
tooling fix that `corpus/services/cms.md` still published as live backlog, and a fence-registry row in
`platform-alignment.md` describing a three-armed fence that ships four, with 16 tests that are 34, and
naming the wrong file as the one the advance mechanism reads. Guard family **30 GREEN · 0 RED**.
**Clause 5 is NOT met and no `P` is claimed.**
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (1 tik) — (6) protocol-stop: n — (7) budget-exhausted: **y** — Outcome: **exit-7**

**Decisions:** `D-M257x-278-1` (a drift fence's RED names a review, never a renumbering),
`D-M257x-278-2` (paid with a new-state claim, never with prose about the drift), `D-M257x-278-3` (0 of
13 anchors held and 0 were repaired — the direction is not measurable), `D-M257x-278-4` (a fence's
registry entry decays like any other claim), `D-M257x-278-5` (rule 80 takes the weaker half of
iter-277's route on purpose), `D-M257x-278-6` (a route that must not be spent is named in full).

**Side-deliverables:** `§5` **rule 80** — the runner must not live inside its own subject; closes the
runbook half of `ROUTE-M257x-277-the-census-cannot-be-run-from-inside-its-own-tree`.

**Routes carried forward:**
- **`ROUTE-M257x-278-thirteen-unpinned-rext-anchors-are-on-undecidable-clocks`** — 13 of 13 unpinned
  anchors into advance-touched files differ from their text at the last-cited ref, and the direction is
  **not mechanically decidable**: `safety.md:712` and `media-substrate-spec.md:122` cite the *same*
  function at the old and new refs respectively, and `org-repos.md:195` matches neither. Repair needs an
  author reading each sentence, not a renumbering pass.
- **`FIX-M257x-278-clone-pin-guard-docstring-says-three-arms`** — `clone_pin_guard.py`'s header
  enumerates three arms while shipping four, and repeats the canonical-vs-workspace framing this iter
  retracted. The code is right; only the docstring is behind.
- **`FIX-M257x-278-census-substrate-tests-hardcode-an-absolute-ROOT`** — `TheBasenameShareIsDERIVED` and
  `TheModulesOwnCommentFiguresAreDERIVED` pin `ROOT = Path("/Users/marco/workspace/anthropos/rosetta")`,
  so they cannot be run against a frozen control tree at all. Rule 78's class, one layer in.
- **`ROUTE-M257x-278-rext-tag-SoT-was-six-iters-stale-unnoticed`** — `.agentspace/rext.tag` read
  `iter-270` while the consumption clone sat at `iter-276`. Git-ignored, so it never appears in a diff.
  Bumped here; nothing asserts it.
- **The fence half of `ROUTE-M257x-277`** stays open — *"the counters exclude virtualenvs"* is the
  stronger fix and was deliberately not taken (`D-M257x-278-5`).
- **Clause 5's semantic reading is still unmeasured** (last: iter-131, `P = 29 / N = 47`, a floor).
- Unchanged and not absorbed: `ROUTE-M257x-274-successor-half-is-uncovered`,
  `ROUTE-M257x-274-tie-order-is-unstable`, `FIX-M257x-269`,
  `ROUTE-M257x-270-directus-consumer-cms-key-outlived-its-rollback-path`, `FIX-M257x-266`,
  `FIX-M257x-265`, `ROUTE-M257x-h59`, `ROUTE-M257x-h65`.

**Lessons:**
1. **A fence's finding-count is a count of MENTIONS, not a workload.** *"44 citing sites"* named zero
   defects; the two real ones were in documents the fence never pointed at. Read what the assertion
   says, not how big its noun phrase is — iter-256 paid 27 reverted edits for the other reading.
2. **Re-pointing the citations an edit MOVES does not reconcile the claims it FALSIFIES.** iter-270
   re-pointed 7 citations and updated `setup_guide.md` correctly — and still left both findings here,
   because neither cited a line iter-270 touched. One was written by a *different* iter (268) about a
   defect 270 fixed; the other described a fence's *shape*. **Second-order debt is invisible to a
   citation sweep and is where the advance's real cost sits.**
3. **When a fix breaks a test, ask which one is wrong — and the answer can be BOTH, differently.** One
   arm had found a genuine defect in my prose; the other was asking for routine maintenance. Bumping
   both figures would have absorbed a real defect into a ratchet; repairing both in the corpus would
   have been impossible. **The partition was the finding.**
4. **A test with a hardcoded absolute ROOT cannot be controlled.** The standard frozen-clone control was
   unavailable, and the substitute — re-deriving the same functions against both corpora — worked only
   because the derivation was three lines. A fence that cannot be run against a control is a fence whose
   findings cannot be attributed.
