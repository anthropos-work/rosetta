**Type:** tik — under [`TOK-05`](../decisions.md), step 1 of `D-M257x-59-5`'s ordering (**fence** →
citations → map state → read).

# iter-60 — the sibling predicate guard, and the profile class it closes

## Phase A — re-derive every denominator (nothing inherited)

Platform clone and origin both at `0dab54d`, re-verified at open **and again at close**. All six
denominators reproduced from the artifacts, then **cross-checked against `docker compose --profile X
config --services`** so the numbers are not merely what my own parser believes:

| # | derived | value | cross-check |
|---|---|---|---|
| G1 | compose services (`include:` resolved) | **10** | `--profile all` → 8 + `storage` + `messenger` |
| G1 | the always-on **floor** (no `profiles:` key) | **3** | every selection contains exactly these |
| G1 | legal profile tokens | **8** | — |
| G2 | `repos.yml` entries | **6** | — |
| G3 | `PROFILE ?=` default and its selection | **`core`** → **5** | `--profile core` → the same 5 |
| G4 | `*_RPC_ADDR` compose sets | **4**, all `http://backend:8083` | — |
| G5 | repos with `migrations: true` | **1** (`app`) | — |
| G6 | `STORAGE_RPC_ADDR` | unset in config / **6** app reads | — |

**Three inherited numbers re-derived; two corrected** — `D-M257x-60-7`. The briefed *"17 files / 30
occurrences"* measures **26 live docs / 56 lines**; `cmd/academyImport/main.go:235` is the `is
required` return and the `Getenv` is at **`:231`**.

## Phase B — the fence, watched RED then GREEN

`rosetta-extensions/stack-core/platform_predicate_guard.py` — a **new sibling**, not a widening
(`D-M257x-60-1`). Inputs `repos.yml` + `docker-compose.yml` (`include:` resolved) + `Makefile`; six
assertions, each run **both directions**.

**Resolving `include:` is load-bearing, not tidy.** Two thirds of the floor lives in `common.yml`, so a
`docker-compose.yml`-only parser computes `floor = {sentinel}` and concludes — wrongly, and quietly —
that a dead token starts nothing. Shipped as its own regression test.

| run | findings |
|---|---|
| first draft, live tree | **37** — of which **16 were the guard's own** |
| after replacing four rules | **21**, 0 false positives |
| after the corpus repair | **2** |
| after pinning the protocol doc's two historical quotes | **0 — GREEN** |

`D-M257x-60-3`: all four were fixed by **replacing the rule with one derived from the artifact's
structure**, never by excepting a name — `--profile` needs a compose driver in-window (`buildbench
--profile billion` selects a *host* profile); a token must match compose's own token shape (`PROFILE=P`
is a metavariable); a repo-count's modifier slot is case-sensitive (`§ 2 Clone repos` is a section
number); a name adjacent to `.`/`/` is an identifier, not a repo. A fifth rule was **dropped rather
than fixed** — per-name attribution of `migrations:` flags in free prose is not mechanically decidable
(English puts the list *before* "are `migrations: false`" and *after* "have `migrations: true`"), so
G5 checks the enumeration construct plus the count invariant and **names the other 22 of 26 lines as
UNREACHED**.

**28 tests**: derivation · RED→GREEN · **three INVERSION mutants** (swap which token compose declares,
invert the migration flags, invert the RPC value — a guard hardcoding *"graphql is dead, core is
alive"* survives all three; a derived one cannot) · a **no-op positive control that survives**
re-running with identical reach · one regression per replaced rule · G6-unmeasured-is-not-zero.

`D-M257x-60-4`: G6 was re-cut mid-iter. As first written it fired on the *platform* being mid-fold,
which would have made it **permanently RED** — and a permanently-RED fence is one that gets disabled.
It now grades the **corpus**: a mid-fold variable is a finding only when no document cites any of its
real read sites.

## Phase C — the predicate class repaired, and a third failure mode nobody had recorded

`D-M257x-60-2`. The briefing described two behaviours; `docker compose --profile X config --services`
at `0dab54d` splits into **three**, and the missing one is a hard error:

| class | tokens | behaviour |
|---|---|---|
| works | `core` `backend` `all` `storage-legacy` `customerio-sync` | rc=0, selects beyond the floor |
| **silent no-op** | `graphql` `cms` `jobsimulation` `roadrunner` `storage` | **rc=0**, starts the floor only |
| **hard-fail** | `frontend` `studio-desk` `messenger` | **rc=1**, `depends on undefined service "backend"` |

`make up PROFILE=frontend` and `make up PROFILE=studio-desk` are **documented commands in
`setup_guide.md`'s table and both exit 1**. Of the six profile rows `CLAUDE.md` carried, **one**
(`all`) was accurate.

**11 files repaired, +241/−74.** The two that were fortified against repair:

* `cms.md` asserted the husk *"still starts"* and that messenger *"is still pointed at it"*
  (`CMS_RPC_ADDR=http://cms:8091`) — **M809 has landed**; all four compose values read
  `http://backend:8083` and there is no cms container.
* `platform-alignment.md` §5 carried iter-22's refutation as standing guidance — *"Applying the
  correction would have replaced two true statements with false ones"* — so the protocol doc **forbade
  the repair that is now required**. `D-M257x-60-6`, promoted to §5 rule 31.

`D-M257x-60-5`: the fence went **RED on this iteration's own repair text**. Warnings written as
`` `make up PROFILE=cms` does NOT fail `` re-introduce the construct G1 counts, and a copy-pasteable
invocation for a silent no-op is indistinguishable from an instruction. Every warning was rewritten to
**name the token without writing the invocation**.

**Also corrected: the protocol doc's own number.** §2 said a renamed profile is *"a successful command
that starts **zero** containers."* Measured: **three**. Zero would at least be unambiguous; three
presents as a partially-working stack and sends the reader debugging the application.

### A tag-succession note, recorded rather than tidied away

`fast-build-m257x-iter-60` was annotated, pushed and **verified on origin** (pre-flight rung zero) as
soon as the guard's 27 tests were green — before the live tree revealed the G6 re-cut
(`D-M257x-60-4`). A pushed tag is immutable and **force-push is forbidden**, so the follow-up commit
carries a successor tag, `fast-build-m257x-iter-60.1`, likewise verified on origin
(`git ls-remote --tags origin`). **The pin points at `.1`** — the tag that actually contains the
shipped guard. The lesson is small and worth keeping: *tag when the iteration's code is final, not
when its first green appears* — "tagging is not publishing" has a sibling, **publishing is not
finishing**.

## Phase D — the protocol-doc updates TOK-05 withheld

Written now, as `D-M257x-59-5` routed:

* **§5 rule 29** — *a reading names INSTANCES; only a derivation can name a PREDICATE.* Why the
  ten-reading series had a fixed point rather than a slope; extends rule 19 one level down.
* **§5 rule 30** — *grade on "does it still SELECT something", not "does it still parse"*, with the
  enumerate-the-floor corollary and the do-not-spell-a-dead-command corollary.
* **§5 rule 31** — *a refutation expires exactly like the claim it refuted*, and is more dangerous when
  it does, because anti-repair language reads as already-adjudicated.
* **§5 rule 32** — *re-derive the hand-off's numbers, including the orchestrator's.*
* **§6** — *the platform's CONFIG is its documentation of record; its NARRATIVE docs are not*, with the
  three self-documenting artifacts, the PR-#14 *two-documents-that-agree-are-not-two-witnesses* result,
  and the mid-fold two-sided corollary.
* **§7 rule 4** — the **citation-safety half** (`D-M257x-59-3`): schema-safety and citation-safety are
  unrelated properties; iter-58's advance was correctly vetted for the first and still moved 22 of 23
  citations, with a 4.5% fence catch rate.

`corpus/services/storage.md` carries the first **two-sided mid-fold record** — config side, compose
side, `repos.yml` side, consumer side, each cited. The map's 8th state token stays iter-62's.

## Close — 2026-08-04

**Outcome:** the sibling predicate guard shipped GREEN after being watched RED at 37 → 21 → 2 → 0, with
28 tests including three inversion mutants and a surviving no-op control; the `graphql`-profile
predicate class repaired across **11 files (+241/−74)** to **zero fence findings**; a **third**
profile failure mode (hard-fail, rc=1) recorded for the first time; two fortified-wrong claims
repaired, one of them the protocol doc's own instruction not to repair it; **six protocol-doc rules**
written (§5 29–32, §6, §7 rule 4b). Two inherited denominators corrected by measurement.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — 4 of 5. Clause 5 remains the only open one; this iter fences three predicates under
it and does not claim to have read it to zero.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** `D-M257x-60-1` (the guard, six assertions, both directions) · `D-M257x-60-2` (the
three-class failure taxonomy) · `D-M257x-60-3` (four rules replaced, one dropped) · `D-M257x-60-4` (G6
grades the corpus, not the platform) · `D-M257x-60-5` (do not spell a dead command) · `D-M257x-60-6` (a
refutation expires) · `D-M257x-60-7` (three inherited numbers re-derived, two corrected)
**Side-deliverables:** none — every edit was in the iter's planned scope.
**Routes carried forward:**
- `FIX-M257x-iter58-mainline-shift` (**21 of 22**) → **iter-61**, under §7 rule 4's new citation half.
- `DOC-M257x-iter59-storage-mid-fold` → **iter-62**: the map's 8th state token. The *measurement* landed
  here (`storage.md`'s two-sided block, G6-fenced); the **vocabulary + assertion-C change** did not.
- `CHECK-M257x-iter60-stale-pin-exemption` → **new**. A ref-pin exempts a claim from G4, so a stale
  claim can immunise itself by citing an old ref — `jobsimulation.md:95` is exempted correctly by rule
  while reading as current guidance. The guard now *prints* the refs that bought each exemption;
  turning "pinned to a superseded ref" into a finding is not built.
- `CHECK-M257x-iter60-g6-citation-subject` → **new**. G6's two-sided test matches the read site's
  `path:line`, not its subject; a citation landing on the right line for another reason would satisfy
  it. The one that closed it was hand-checked.
- `FIX-M257x-iter53-union-set` (46 vs 35) → **PENDING USER DECISION**, untouched.
- `FIX-M257x-iter56-assignment-flake` → still **NOT DECIDED**; needs a failure *rate*.
- `CHECK-M257x-iter38-ai-act-classification` → needs an owner **outside** this milestone.
- Unchanged and still open: `-cold-daemon-registry` · `-grep-vs-failclosed` · `-empty-stdout-class` ·
  `-baseline-refs` · `CHECK-M257x-iter58-derive-preregistrations` · `FIX-M257x-iter57-within-block-drift` ·
  `CHECK-M257x-iter57-anchor-guard-bare-class` · `FENCE-M257x-iter54-refs-block` ·
  `CHECK-M257x-iter52-second-ai-manager` · RF-2/3/7–13.

**Lessons:**

1. **A fence's first RED includes the fence's own errors, and telling them apart is the work.** 37
   findings, **16 of them the guard's**. Every one was removable by replacing the rule with one derived
   from the artifact's structure — and the replacements cost less than the exception lists would have.
   The tell is whether the fix names a *thing* (`billion`, `P`) or a *property* (needs a compose
   driver; must match the artifact's token shape).
2. **Some rules should be dropped, not fixed — and the drop must be REPORTED.** Per-name attribution of
   `migrations:` flags in free prose is not mechanically decidable in English. G5 now names 22 of 26
   lines as UNREACHED. A fence that quietly passes what it cannot read is worse than one that admits it.
3. **A fence that can never go green is not a fence.** G6's first draft fired on the *platform* being
   mid-fold — a legitimate state the developer is still working through. Grading the *corpus* instead
   (does it record both sides?) makes the same finding actionable and closable.
4. **Do not write a runnable spelling of a broken command, even to warn about it.** The fence caught
   this iteration doing exactly that in its own corrections. Name the token; do not write the
   invocation.
5. **The most dangerous claim is a correctly-refuted one, after the platform moves again.** iter-22's
   refutation was right at `2adcf71` and became a standing instruction not to make the repair that
   `0dab54d` requires — protected, not exposed, by its own emphasis.
