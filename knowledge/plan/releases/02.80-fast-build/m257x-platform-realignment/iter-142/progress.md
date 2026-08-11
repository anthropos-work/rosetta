**Type:** tik — under `TOK-08` (*census the mechanical classes; stop sampling them*). Closes
`FIX-M257x-iter141-retraction-idiom-sweep`, the route iter-141 opened one iter earlier and the newest on
the queue. Three planned lines (census → repair → fence), declared in this iter's `overview.md`.

# iter-142 — the retraction idiom, censused — and the census's own blind spot, found by the fence family

## The class, and why it is censusable

`§5` rule 63(c′) named it at iter-141 from three incidents: **the corpus's own correction idiom turned
fences RED three times in five iters, in three different files, always on a pin whose own sentence
existed to retract it.** *"it was `:274` at `<sha>`"* keeps the retracted number **live in the text**,
in exactly the token shape every resolver binds — so the next insertion above its target rots it, and
**a fence matching on FORM cannot tell the quotation from the assertion.**

It is censusable under `D-M257x-140-2` — *a class is censusable iff its subject carries its own HEAD* —
because **the subject is the citing sentence itself**, not a resolved target. That is precisely what
iter-138's retracted census lacked, and it is why this one could be run to completion.

## ⚠️ Read this first: the census was WRONG before it was right, and a fence found it

The first pass of this iter matched only **bare** pins (`` `:274` ``), enumerated 922 of them, repaired
44, and reported the class **clean at 0**. Then the iter's own commit was put through the commit-scoped
guards, and **`repair_leak_guard` went RED on it**: `shared_libraries.md:77` still published *"This
cited `main.go:507-508` until M257x iter-138"* — a **twin of a `CLAUDE.md` site this fence had just
flagged and repaired**, missed for one reason only, that its pin carried a path.

**A fence whose regex defines its own denominator will report a clean census over the population it
happens to match**, and it will sound exactly like a complete one. Nothing inside the fence could have
caught this. What caught it was a *different* guard, on a *different* axis, run over the commit.

Five more path-qualified sites were behind it. The numbers below are the corrected, two-arm census.

## The census — stated with its denominator (iter-114's rule)

| quantity | value |
|---|---|
| source documents scanned (`corpus/**` + `README.md` + `CLAUDE.md`) | **94** |
| **backticked line pins enumerated — the denominator** | **2,185** |
| of those, path-qualified (`` `main.go:507-508` ``) | 1,311 |
| of those, bare (`` `:274` ``) | 874 |
| excluded as ports (`>= 10000`, bare form only) | 23 |
| excluded inside fenced blocks (transcripts, not prose) | 3 |
| **live findings — the class** | **50, across 20 files** (44 bare + 6 path-qualified) |
| still in a retraction clause after the sweep | **2, both in the teaching-exempt protocol doc** |
| **hand-audited precision, BARE arm, over its whole population** | **44 / 44 = 100 %** |
| **hand-audited precision, PATH arm, over its whole population** | **6 / 6 = 100 %** after the window fix; **1 / 5** before it |

**Every finding was read in context before a single line was repaired.** iter-138 published a
mechanical number and iter-139 audited it to **0-for-12**; this iter inverted the order, and that
inversion is the transferable part — **but the inversion protects PRECISION and does nothing for
RECALL**, which is precisely what the path-qualified miss demonstrates. Auditing every finding tells you
nothing about the findings the predicate never produced.

## What the audit changed — the two false REDs that shaped the predicate

The first draft required a retraction verb plus **any** supersession marker anywhere in the clause. Over
46 candidates it produced **2 false positives**, both from markers attached to something other than the
pin:

| site | why it is NOT a retraction |
|---|---|
| `external_services.md:473` | *"copies **were** deleted **at `915da06`**, which left the comment at `` `:19-20` `` in their place"* — the sha anchors the **deletion**; the pin is **live** |
| `hiring.md:221` | *"assigned at `` `:1846` ``. Not a mirror: `local_jobsimulation_session.go` **no longer** exists"* — *"no longer"* is about a deleted **FILE**, ~55 characters past a live pin |

Both weak markers (`at <sha>`, `no longer` / `had moved`) were demoted to a **Tier B** that must sit
**within 30 characters AFTER the pin** — where the genuine idiom puts them (*"it was `:100` at
`0dab54d`"*) and where the false positives did not have them. `until` / `wrote it` / `not pinned` stay
Tier A (anywhere in the clause) because they cannot govern anything but the pin. Result: **44 of 44**.
Both false-RED classes ship as named regression tests.

## The repair — 44 tokens gone, 0 evidence lost, 0 lines moved

The rule is (c′)'s: **retract by describing the artifact, never by reproducing it** — and this iter
sharpens it to **fence the TOKEN, not the digit**, which is what makes the repair non-destructive.
*"rotted +8"*, *"iter-102 added +23 and +16 to the old numbers instead of re-measuring"*, *"it stood ten
lines earlier in the file"* say everything the pin said and are invisible to every resolver.

**All 20 repaired files came out line-count FLAT — rewritten in place, added minus removed = 0 for
every file, verified with `git diff --numstat`.** (The one file with a net delta is
`platform-alignment.md`, which GAINED rule 63(c‴) — an addition, not a repair.) That is not decoration: iter-141's own repair of this class turned a fence RED
by inserting above a pin. A sweep that removes anchor rot must not induce any, and this one provably
did not.

Repaired, bare arm (17 files): `CLAUDE.md` · `architecture/{ai_architecture, dependency_map,
external_services, service_taxonomy}.md` · `services/{academy-backend, ai-readiness, backend,
clerk-integration, cms, gotenberg, graphql-wundergraph, hiring, jobsimulation, next-web-app, skillpath,
studio-desk}.md`. Path arm (6 sites): `architecture/{shared_libraries, dependency_map,
platform-migration-status}.md` · `services/{hiring, storage, coursebuilder}.md`. Plus
`service_taxonomy.md:213`'s cross-doc pin into `studio-desk.md`, which **`anchor_offset_guard` flagged
as rotted by this very commit** and which is now named by construct.

Two sites are worth naming because they are the class arguing with itself:

* **`graphql-wundergraph.md:138`** held **four** retracted pins in one paragraph — a paragraph whose own
  closing sentence reads *"the third pin in this one paragraph to do so, which is why all of them are
  now named rather than numbered."* It was still numbering them.
* **`clerk-integration.md:126`** was **iter-141's own repair of this very class**, and it reproduced
  **two** pins while doing it.

## The fence — `FENCE-M257x-iter142-retracted-pin`

`rosetta-extensions/stack-core/retracted_pin_guard.py`, registered in `guard_family.py` (`tree` class,
takes `--repo-root` for the same reason `corpus_citation_guard` does — its subject is not under its own
tree). **The guard family grows 22 → 23.**

* **Mutation control that fires** — three mutations (the `until` form, the `at <sha>` form, the
  *"named, not pinned"* form) each proven to turn it RED, **with the unmutated control proven GREEN in
  the same class**, so a RED is attributable.
* **A path-qualified arm with its own narrower window**, and the asymmetry is **measured, not
  designed**: Tier-A-anywhere scores 44/44 on bare pins and **1 of 5** on path pins. Requiring the
  marker *after* the pin gives 6/6 true and drops 4/4 false. ⚠️ **That window is a tuned constant on a
  five-site denominator** — stated here rather than buried, and routed forward.
* **Anti-vacuity control written against the guard's SUBJECT** (§8, iter-94), and it has an unusually
  sharp arm: **`census.clauses` over the real corpus can never legitimately reach 0**, because
  `platform-alignment.md` rule 63(c′) has to spell the idiom in order to teach it. If the matcher ever
  stops matching, the guard's own green becomes impossible rather than merely wrong.
* **The teaching exemption is by PATH and is counted, not hidden** — the two exempt instances are
  reported in the census line every run.

## Test gates

| gate | result |
|---|---|
| **New fence's own module** (3 mutations + 2 anti-vacuity arms + **6** false-positive regressions across both arms) | **21 passed / 0 failed** |
| **Scoped rext suites** — chosen by what this iter CHANGED, per rule 63(d): `guard_family`, `corpus_citation_guard`, `retracted_pin_guard`, `anchor_construct_denominator`, `anchor_offset_guard`, both mechanical-fence mutation batteries, `repair_postcondition` battery | **209 passed / 0 failed** (699.66 s) |
| **Guard family**, `--repo-root` + `--platform stack-demo/platform @ 0c91421` | **19 GREEN · 0 RED · 4 not-run** — up from 18 GREEN pre-iter, the +1 being this iter's own guard. Not a whole-family green; the runner's own summary says so |
| **The three COMMIT-scoped guards** (`anchor_offset`, `repair_leak`, `value_change`) over the iter's own range `fc7d71c3..HEAD` — three of the four the family reports as `not-run` | **3 RED on the first commit, two of them real misses**: the path-qualified twin (the headline) and a cross-doc pin this commit rotted. The third was measured and is a **shingle collision, disclosed not obeyed**. All resolved; **re-run GREEN — `anchor_offset` 25 citations graded of 57 seen, 0 CANNOT-TELL, 0 unmeasured; `repair_leak` GREEN; `value_change` GREEN**. **This is the gate that found what the tree-state fence structurally could not** |
| **Pre-iter baseline**, same invocation | **18 GREEN · 0 RED · 4 not-run** — identical to iter-141's close, so the delta is attributable |
| **Whole suite** | **NOT re-run — `§5` rule 60 requires saying so.** Stated as a gap, not characterised as covered. The 8 scoped modules are the change-derived set, and this is the **first** iter since 132 to touch `rosetta-extensions` at all, so iter-132's clean whole-suite run **no longer stands** on this rext tree |
| **Suite wall-time** | not quoted as a measurement — `FIX-M257x-iter132-suite-walltime-is-not-a-measurement` stands. The 699.66 s above is a duration, not a benchmark |

**One duplicate test run was killed mid-flight and is disclosed rather than dropped:** a first
invocation was launched piped to `tail`, produced no readable output, and a second was launched to a
file. Both were briefly alive together; the first was killed so two concurrent runs could not confound
each other's temp state. **The 209/0 above is the file-based run, taken alone.**

## Close — 2026-08-08

**Outcome:** the retraction-idiom class is **censused, repaired to zero, and fenced** — and the census's
own blind spot was found and closed inside the iter. **2,185 line pins enumerated over 94 documents; 50
live instances across 20 files (44 bare + 6 path-qualified)**, all hand-read before repair, all converted
from reproduction to description with **zero net line movement per file**. The headline finding is not
the sweep: **the first pass matched only bare pins and reported the class clean at 0, and
`repair_leak_guard` went RED on the commit that repaired it.** A fence whose regex defines its own
denominator reports a complete-sounding census over whatever it happens to match; auditing every finding
protects precision and says nothing about recall. Guard family **22 → 23 members, 19 GREEN · 0 RED**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5**, unchanged. **No `N` reading was taken, so no `N` movement is claimed**
(`§9`'s guard-rail 1, in its required words); `TOK-08` sequences the read after the mechanical sweep and
this is one class of that sweep, not its completion.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**iters 135–142 took no reading, so the metric is UNMEASURED not unmoved — `§9`'s iter-type refinement; and `TOK-08`'s sealed refutation branch bars an agent-authored successor in any case**) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (**1 tik this session**) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-142-6` (**the headline** — a fence's regex is its denominator, and only a guard
on a different axis can see the omission) · `D-M257x-142-1` (audit the predicate BEFORE the repair, not
after the publication — 44/44 vs iter-138's 0/12; it buys precision, not recall) · `D-M257x-142-2` (a ref-qualified historical pin is in the class, because the
fence does not read the qualification) · `D-M257x-142-3` (fence the TOKEN, not the digit — the repair
keeps the evidence) · `D-M257x-142-4` (a sweep against anchor rot must be line-count flat, and this one
was measured to be).
**Side-deliverables:** none.
**Routes carried forward:**
- **`FIX-M257x-iter142-value-change-articles` (NEW)** — `value_change_guard` treated a one-token
  **article** change (`the` → `that`) as a corrected *value*, and flagged a site that shares only the
  house idiom *"a pin is a pin, a branch name is not"*. Measured false; resolved here by keeping the
  article rather than by editing prose that was never wrong. **An article is not a value** — the guard
  should not gate on a change confined to function words.
- **`FIX-M257x-iter142-path-arm-window` (NEW)** — the path arm's *"marker must be within 40 characters
  after the pin"* is a **tuned constant on a five-site denominator**. It scores 6/6 there and the
  alternative scored 1/5, but that is a fit, not a derivation. Derive it or publish it as a disclosed
  floor.
- **`FIX-M257x-iter142-tier-b-underflag` (NEW)** — Tier B's 30-character reach **under-flags a chained
  pin**: on `external_services.md:209`, `` `:84` `` fired and its sibling `` `:669-670` `` did not,
  because the sha sits ~35 characters past the first pin. Under-flagging is the correct direction for a
  fence and the site was repaired anyway (both pins are gone), **but the reach is a tuned constant and
  nothing measures it.** Derive it, or state it as a disclosed floor.
- **`FIX-M257x-iter142-whole-suite-owed` (NEW)** — this iter changed `rosetta-extensions` for the first
  time since iter-132, so the standing *"iter-132's clean whole-suite run covers this tree"* claim is
  **spent**. A whole-suite run is owed before the next rext-touching iter quotes coverage.
- `FIX-M257x-iter135-adjudicated-live-defects` — non-citation remainder: `org-repos.md:227`,`:370`,`:43`
  · `ai_architecture.md:111`,`:224` · `next-web-app.md:17`,`:186` · `external_services.md:368`.
- `FIX-M257x-iter140-receipts-not-checkable-here` (13 of 22) · `FIX-M257x-iter140-receipt-fence` ·
  `FIX-M257x-iter138-anchor-rot-fence` (re-specified: head resolution first) ·
  `FIX-M257x-iter134-fence-family-has-no-shared-predicate-layer` ·
  `FIX-M257x-iter133-two-fives-need-a-fence` · `FIX-M257x-iter132-suite-walltime-is-not-a-measurement` ·
  `FIX-M257x-iter131-predicate-sets-not-enumerated`.
**Lessons:**
0. **A fence's regex IS its denominator.** The census said 922 pins and 0 residual; both were true of
   the population the regex matched and neither was true of the class. **No control inside the fence
   could have caught it** — not the mutation control, not the anti-vacuity arm, both of which passed.
   What caught it was `repair_leak_guard`, a *different* guard on a *different* axis, run over the
   commit. Two consequences worth carrying: **(a)** a census's reach claim must name the FORMS it
   matches, not just the count it returns; **(b)** the commit-scoped guards are not a formality after a
   sweep — on this iter they were the only thing between a real miss and a published "0".
1. **Audit the predicate before the repair, not after the publication.** Same milestone, same month:
   iter-138 published first and was audited to 0-for-12; iter-142 audited first and repaired 44 of 44.
   The predicates were comparably mechanical. **The ordering was the whole difference.**
2. **A fence's precision lives in its WEAKEST marker.** Two markers out of a dozen — both attaching to
   something other than the pin — were the entire false-positive rate. Ask of each marker: *can this
   govern anything but the thing I am matching?* If yes, it needs a reach limit, not a place in the set.
3. **Fence the token, not the digit.** Stated as a repair rule it sounds pedantic; measured, it is what
   let 44 retractions keep every number they were making a point with while losing the form that rots.
4. **A sweep against rot must be line-count flat, and you can prove it in one command.**
   `git diff --numstat` per file, added minus removed. iter-141 induced a RED with its own repair of
   this class; this iter could not have.
