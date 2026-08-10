**Type:** tik (under [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07))

# iter-244 — the seventh runnable input, and the first that grades the SUBJECT

## Phase A — the instrument, and the three times it was its own largest finding

The census was written, run, and **corrected three times before a single corpus line was repaired.** Each
correction is recorded because each is now a regression test, and because this milestone's standing rule is
that a census must enumerate its substrate before anyone believes it.

| pass | findings | of which the instrument's own defect | the defect |
|---|---|---|---|
| 1 | 10 | **6** | `(?:…\|js\|json\|…)` — leftmost-alternative matching truncated `data-dna.json` to `data-dna.js`, a file that does not exist. Fixed by ordering longest-first **and** adding `(?![A-Za-z0-9])`; both halves are needed |
| 2 | 4 | **2** | no `/` or `.` in the left lookbehind, so `app/knowledge/architecture.md` and `.claude/skills/stack-secrets/SKILL.md` matched on their **tails** |
| 3 | 26 | **23** | the section set was derived from the rext tree, which contains `knowledge/` — and `knowledge/…` in this corpus is overwhelmingly a **platform** path (`app/knowledge/…`, `infrastructure/knowledge/…`) or rosetta's own `knowledge/plan/**` |

Pass 3 is the instructive one: **deriving the subject from the tree was the right call and still produced
the worst pass**, because a directory that exists is not therefore a *referenced section*.
`corpus_citation_guard`'s docstring had already measured and named this exact class in 2026-08-06 — the
lesson was written down, in this repo, and re-derived anyway.

**Final substrate, stated with its denominator and its clone set** (iter-114 / iter-241):

| | |
|---|---|
| live documents scanned (`git ls-files`; `corpus/**` + `CLAUDE.md` + `README.md` + `.claude/skills/**`) | **114** |
| sections derived from the rext tree (`knowledge` excluded, with reason) | **11** |
| distinct rext paths referenced | **145** |
| occurrences of them | **300** |
| documents carrying ≥ 1 | **45** of 114 |
| clone set | `.agentspace/rosetta-extensions` @ `c2d9052`, worktree-clean |

## Phase B — the reading

**3 of 145 distinct paths do not resolve.** Two are defects; one is a disclosed absence.

```
corpus/services/clerkenstein.md:65   `alignment/dna/clerk-2.6.0.json:131`
corpus/services/clerkenstein.md:149  `alignment/dna/clerk-js-5.json`
corpus/ops/demo/coverage-protocol.md:680  stack-verify/e2e/tests/probe-aireadiness-deeplink.spec.ts
```

The third is excluded **by name and out loud**: its sentence's whole claim is *"was never committed — the
file does not exist."* A correct citation of an absence is not a broken citation — but per `§5`, a correct
exclusion is still a defect while it is silent, so the guard prints it on **every** run, green or red.

### The defect, and why it is worth an iter

The DNAs live at **`clerkenstein/alignment/dna/`**. The corpus dropped the `clerkenstein/` prefix — and
`alignment/` **is itself a real section**, so the wrong path names a real section and a real-looking
subpath. **It reads correct.**

And it was read. Verbatim, at `clerkenstein/alignment/dna/clerk-2.6.0.json:131`:

> *"…M219 landed the fix (Store.SeedOrgIdentity / LookupOrgEid, wired from the roster at cmd/fake-bapi),
> taking the Go surface 97.2% -> 100%."*

**The quote is exact. The line number is exact. The path cannot be opened.** Measured: **13 distinct iters
across 20 graded seat files** verified this one citation — iters 49, 50, 53, 76, 82, 97, 99, 101, 103, 109,
116, 119, 122 — recording it as *"ENUMERATED"*, *"is **exact**"*, *"✔"*. Every one confirmed the CONTENT.
None noticed the path did not resolve.

That is `TOK-08` reduced to a single citation: **a reading samples MEANING; only a fence censuses
RESOLUTION.**

### The structural finding — an unresolvable anchor and a WRONG one produce the identical record

`anchor_construct_guard` nominally has `alignment/dna/clerk-2.6.0.json:131` in its subject: it is a
path-qualified `file:line`. Run on this iter's opening tree it reported

```
anchor-construct-guard: 882 anchor(s) resolved across 114 file(s); 599 unresolvable
anchor-construct-guard: OK — every resolvable anchor names a construct     ← exit 0
```

**An anchor it cannot resolve is booked as OUT OF REACH, not as a finding**, so a citation naming a file
that does not exist joined the same 599-strong bucket as a citation whose clone is merely absent. The guard
is not wrong to do that — it cannot tell the two apart, because it does not know whether the clone is
present. This guard can, and only runs when it is: **that is the whole reason the class needed a second
instrument rather than a wider first one.**

## Phase C — repair

Both sites re-pointed to `clerkenstein/alignment/dna/…` in `corpus/services/clerkenstein.md`. The repair is
**not** "delete the citation" — the files are real, one directory deeper, and a test asserts both halves
(the target exists under the `clerkenstein/` prefix **and** does not exist without it, so the defect stays
a defect).

## Phase D — the fence

`stack-core/rext_path_guard.py` — **FENCE-M257x-iter244**, the **seventh** runnable-input fence and the
first that grades the runnable **subject** rather than an argument.

* **Subject derived, never listed** — sections come from the rext tree. CLAUDE.md's own enumeration of them
  was measured wrong at iter-129, so it is not a source. The **named limitation**: a section that is
  *renamed* drops out of the subject rather than firing; the derived set and its count print on every run
  so the reach is visible rather than assumed.
* **Every exclusion is stated in the docstring with the measurement that motivated it** — `knowledge/`
  (26 findings, 23 false), tail matches (23 false), `<rext>/`-prefixed runtime paths with placeholder
  segments, directory references.
* **Anti-vacuity, four ways, each exit 2 and never 0**: no `CLAUDE.md`, **no rext tree**, zero sections,
  zero references. The absent-tree case is the one that matters — the guard's subject is *"does this path
  exist in rext"*, which is unanswerable without rext, and `§9` iter-174's rule is that a capability probe
  failing OPEN disarms the check it guards.
* **17 unit tests**, `/usr/bin/python3 -m pytest` (CPython 3.9.6), all passing: mutation in **both**
  directions (a planted path goes RED; the same path goes green when the file appears), every-occurrence
  reporting, the `rosetta-extensions/`-qualified form, **two instrument-regression tests written directly
  from Phase A's own failures**, the disclosed-absence disclosure test, a *"the pardon is keyed on (path,
  citing file)"* hole test, a denominator-in-the-verdict test, four anti-vacuity tests, and an **answer key**
  that replays the real pre-repair `clerkenstein.md` out of this iter's own sealed probe commit and demands
  exactly the two findings.

Wired into `guard_family`: **26 GREEN → 27 GREEN / 0 RED / 0 could-not-check / 5 not-run**
(`--repo-root . --platform stack-demo/platform`).

### The fence's first catch was this iter's own prose

Worth recording in full, because it is the cleanest demonstration of induction this milestone has
produced. The Phase-D protocol write-up in `platform-alignment.md` **quoted the defective citation
verbatim as its example** — and `rext_path_guard` went RED **on the paragraph announcing
`rext_path_guard`**, at `platform-alignment.md:3109`, one edit after the class had been closed.

The repair was not an exclusion. The corpus already carries the doctrine — `retracted_pin_guard`'s whole
subject is that *a retraction must not reproduce the pin it retracts* — so the paragraph now **describes**
the wrong spelling instead of reprinting it. **`TOK-06`'s inflow rule, live: the repair loop feeds itself,
and the only thing that catches it is a fence that keeps running after the repair.**

## Phase E — the pre-registrations, graded after the last edit

**2 of 5 refuted.** Recorded as the point of sealing them, not as an embarrassment.

| id | prediction | outcome |
|---|---|---|
| **P-244-1** | ≤ 12 of 140 non-resolving | **HELD** — 3 of 145 (2 defects + 1 disclosed). Note the sealed denominator **140 was itself instrument-inflated**; the corrected one is 145 |
| **P-244-2** | ≥ 1 non-resolving path sits in a **runnable position** (fenced block, or after `rext `) | **REFUTED** — all three are backticked inline prose citations. The runnable-input frame found the class; the class itself turned out to be a **citation** defect, not a command defect |
| **P-244-3** | **0 of 27** pinned paths fail to resolve | **REFUTED** — **1 of 27**, and it is precisely the one `anchor_construct_guard` nominally covers. This refutation is what produced the structural finding above |
| **P-244-4** | the dominant failure mode is a rename (basename exists elsewhere) | **HELD, and sharpened** — both instances are **prefix truncation**: same basename, one directory deeper, one shared mechanism |
| **P-244-5** | **0** guards enumerate bare rext path references | **HELD** — verified against `anchor_construct_guard` (subject is `file:line`) and `corpus_citation_guard` (declared scope `corpus/…` only). Now 1 |

## Close — 2026-08-10

**Outcome:** the seventh runnable-input surface — **the path of the tool itself** — censused end-to-end for
the first time: **145 distinct rext paths / 300 occurrences across 45 of 114 live documents, 3
non-resolving**. Two were repaired (`alignment/dna/…` → `clerkenstein/alignment/dna/…`), one is a disclosed
absence now excluded **out loud**, and the class is fenced at zero by `rext_path_guard`
(FENCE-M257x-iter244), family **26 → 27 GREEN**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-244-1` (sections DERIVED from the tree, with the rename limitation stated rather than
patched) · `D-M257x-244-2` (`knowledge/` excluded from the section set on a 23-false-of-26 measurement) ·
`D-M257x-244-3` (the disclosed absence is excluded by **(path, citing file)** and printed on every run, not
silently skipped) · `D-M257x-244-4` (`anchor_construct_guard` is **not** widened — its could-not-resolve
bucket is correct for a guard that cannot know whether the clone is present; a second instrument that only
runs with the clone is the right shape).
**No `N`/`P` movement is claimed** — this iter took no graded seat.

**Suite state at close** — Python (`stack-core`, `/usr/bin/python3 -m pytest`, CPython 3.9.6):
`tests/test_rext_path_guard.py` **17 passed / 0 failed**. Guard family (`--platform`, from repo root):
**27 GREEN / 0 RED / 0 could-not-check / 5 not-run**.

**Side-deliverables:** none.

**Routes carried forward:**
- `ROUTE-M257x-244-unresolvable-and-wrong-share-one-bucket` → **new.** `anchor_construct_guard` reports
  **599 unresolvable of 1,481** and exits 0. `rext_path_guard` now separates wrong-from-uncheckable for the
  rext slice only. **The platform slice — `main.go` ×33, `infrastructure` ×18, `studioManager.go` ×9 — is
  the same shape and is still ungraded**, and it is larger.
- `ROUTE-M257x-241-wider-citation-surface-is-ungraded` → open — 107 corpus citations into the six
  frozen-legacy repos, graded by nothing.
- `ROUTE-M257x-240-prereq-floors-live-in-three-parallel-blocks` → open.
- `ROUTE-M257x-238-claude-md-fences-are-unmaintained` → open, six-for-six.
- `ROUTE-M257x-238-container-vs-native-is-undrawn` → open, three independent hits.
- `ROUTE-M257x-237-critical-env-list-is-unfenced` → open.
- `ROUTE-M257x-236-disclosure-scope-is-document-level` → open.
- `ROUTE-M257x-235-fence-scope-is-unread` → open.
- `ROUTE-M257x-235-runnable-block-has-two-halves` → open.
- `ROUTE-M257x-h59-range-anchors-are-ungraded` → open (490 range anchors, 24.9 % of the qualified
  population, in neither of `anchor_construct_guard`'s two counts).

**Lessons:**
1. **A citation can be verified thirteen times on its content and never once on its address.** Twenty
   graded seats read the quote at `alignment/dna/clerk-2.6.0.json:131`, confirmed it verbatim, and marked
   it ✔. Reading grades what a citation *says*; only a fence grades whether it can be *opened*.
2. **A path that names a REAL section and a real-looking subpath is the hard case.** A wrong path that
   looks wrong gets caught by anyone. `alignment/dna/…` survived because `alignment/` exists.
3. **Deriving a subject from the tree is right and is not sufficient.** A directory that exists is not
   therefore a referenced section — `knowledge/` cost the worst pass of the three, and the class had
   already been measured and written down in a sibling guard's docstring.
4. **"Cannot resolve" and "resolves to nothing" are different verdicts, and a guard that cannot tell them
   apart must say so rather than merge them.** The fix is not to widen the first instrument; it is to build
   a second one whose preconditions make the distinction decidable, and to keep the first one's honest
   could-not-reach count.
