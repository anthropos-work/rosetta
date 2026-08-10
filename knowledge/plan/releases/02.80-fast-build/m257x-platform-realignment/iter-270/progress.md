**Type:** tik, under `TOK-08`. Shape declared in `overview.md` (multi-step, 6 phases).

Pre-registrations sealed in this iter's first commit (`add716c`) before any measurement.

# iter-270 — spend the frozen pin: the hardcoded service lists, and the derivation that fails OPEN

## The control is spent, and this is what it bought

`D-M257x-258-1` froze the rext pin at `fast-build-m257x-iter-101` as an experimental control: hold the
tooling constant so the platform-ref advance is the single changed variable. **The question it existed to
answer is answered** — iter-258 (cold demo green on the advanced refs, first attempt), iter-260 (three
consecutive `--purge` + `up` cycles green) and iter-262 (a dev stack from current `main`) all landed under
the frozen tooling. Six routes were queued behind it saying *"needs a tag + pin bump."* This iter spends
it.

Measured at open (2026-08-10T18:19Z): `.agentspace/rext.tag` = `fast-build-m257x-iter-101`; the authoring
copy is **205 commits** past that tag and **12 commits** past `origin/main`, tree clean, on `main`.

## Pre-registrations — 3 of 5 held, and both refutations are the iter

| # | prediction | verdict | measurement |
|---|---|---|---|
| **PR-1** | the fail-open arm is **untested** | **REFUTED** | it is tested — as a **characterisation that pins the defect**, with the repair instruction written into the assertion message |
| **PR-2** | `INJECTED`/`REUSE_DEV` dead: pruning changes **zero bytes** | **HELD** | `diff` of the generated override before/after = **identical** |
| **PR-3** | the `cms` hardcode is the path that actually ran here | **HELD** | `stack-demo/cms/studio/requirements.txt` present, with its own `.git` |
| **PR-4** | the dev path has **zero** studio handling | **HELD** | `grep -c` over `dev-stack/` for `studio_required\|lib/studio.sh\|anthropos-studio-room\|STUDIO_REPO` → **0** |
| **PR-5** | the fail-open population is **≥ 3** | **REFUTED** | census: **8** platform-topology derivations, **2** defective. Under a looser grammar it would be 3 — the looser grammar is not used, see below |

**PR-1's refutation is the more useful one.** Harden pass 68 did not miss the fail-open; it *pinned* it,
deliberately, because repairing it needed the tag this iter finally spent — and it wrote the exit into the
assertion itself:

```python
self.assertNotIn("die ", head, "the fail-open was repaired — update this characterisation and the route")
```

A characterisation that names the condition under which it becomes wrong is a genuinely different artifact
from a stale test. It told the next iter what to do, and this iter did it.

**PR-5 is refuted, and refusing to rescue it is the point.** The prediction said *"falls back to a
**permissive/unfiltered** default."* Counting `verify_svcs` — which falls back to a **narrower** floor,
the opposite direction — would make the number 3 and the prediction "hold." A predicate widened until the
answer you already expected appears measures nothing (§5, iter-114). The honest count is **2**.

## The census — 8 platform-topology derivations, graded

Subject: every derivation of the **platform's own topology** (its `repos.yml` or its compose) inside the
demo bring-up path that guards on a precondition or swallows a failure and then continues with a
substitute.

| # | site | derives | on failure / absence | grade |
|---|---|---|---|---|
| 1 | `up-injected.sh:1690` `derive_inject_svcs` | compose build set | the **unfiltered candidate list**, naming 2 corpses | **FAIL-OPEN** ❌ |
| 2 | `up-injected.sh:2156` | default compose profile | `die` | fail-closed ✅ |
| 3 | `up-injected.sh:2681` `verify_svcs` | default service set | a 4-service floor, **disclosed in comment** | fail-narrow ⚠️ |
| 4 | `ensure-clones.sh:133` | repos.yml stub sweep | skipped (no-op) | fail-closed ✅ |
| 5 | `ensure-clones.sh:311` `_studio_repos` | studio consumers | **`"cms"` alone** — the corpse, every live consumer dropped | **FAIL-WRONG** ❌ |
| 6 | `ensure-clones.sh:548` | no-push remote sweep | skipped (no-op) | fail-closed ✅ |
| 7 | `migrate-demo.sh:40` | migration pairs | empty set + a LOUD warn | fail-closed ✅ |
| 8 | `migrate-demo.sh:117` | schemas to create | unguarded | n/a |

**The repair was already written down, one directory over.** `migrate-demo.sh:44-48` states the rule in
the platform's-own-words idiom this milestone keeps reaching for:

> *"What we must NOT do is silently fall back to a stale hardcoded list."*

Two of eight sites disobeyed it. Nothing else new had to be invented.

## What landed

**rext (all four routed sites, plus the dev half):**

1. `demo-stack/lib/studio.sh` → **`stack-core/lib/studio.sh`** (hard move, no alias), so the dev section
   can source it without depending on the demo section. Gains `studio_consumer_names` (derive-or-refuse)
   and `studio_acquire` (the acquisition loop, previously inline in `ensure-clones.sh`).
2. `ensure-clones.sh` — the `_studio_repos="cms"` hardcode and the `make init-studio` special case are
   **gone**. The consumer set is derived from `repos.yml` or the bring-up refuses; acquisition is a plain
   `git clone` for **every** consumer, which is exactly what `make init-studio` was.
3. `up-injected.sh` — `derive_inject_svcs` now **`die`s** on an underivable build set, and
   `INJECT_CANDIDATES` is `"app"`.
4. `gen_injected_override.py` — `INJECTED` → `{"backend": "app"}`, `REUSE_DEV` → `{"sentinel": …}`, and
   the summary line counts **emitted** images instead of `len(INJECTED)`.
5. `dev-stack/dev-stack` — a studio pre-flight in `cmd_up`, on the shared lib, with the same fail-closed
   contract (`FIX-M257x-262-dev-path-needs-the-studio-acquisition`). Placed **before `reg_allocate`**, not
   merely before `docker compose up`: `die` is `exit 1`, which does **not** fire the `ERR` trap that frees
   a reserved slot, so a refusal after allocation would leak N out of the unified registry — a failed
   bring-up that also costs you a stack number. Nothing in the pre-flight needs N.

**Both new fail-closed arms were proven RED where their precondition is absent, not merely green where
present** — the bar the last three harden passes kept missing:

| control | result |
|---|---|
| `studio_consumer_names /nonexistent/repos.yml` | **rc 1** + the refusal diagnosis |
| `studio_consumer_names <repos.yml declaring no repo>` | **rc 1** |
| `studio_consumer_names <real repos.yml>` | rc 0 → `app sentinel next-web-app studio-desk` (**no `cms`**) |
| `derive_inject_svcs` with a bogus `$PLAT` | **`die`, rc 1** |
| `derive_inject_svcs` with the real platform clone | rc 0 → `INJECT_SVCS=app`, both corpses filtered |
| `dev-stack up 1` with `repos.yml` removed | **rc ≠ 0** + the refusal (regression test, shipped as a pair with its silent-path twin) |

## Three things this iter found that were nobody's planned scope

**1. The direct anti-regression test was GREEN while the bug was present, and the reason is a one-line
window.** `test_no_script_gates_studio_on_a_service_name` requires a `"$svc" = <name>` comparison **and**
the word `studio` on the *same line*. The regression had re-entered in the two-line form shell actually
uses. Measured against the pre-repair file: the widened window (line + successor) → **1** hit; the old
one-line window → **0**. iter-268 found the hardcode by hand instead, four releases later.

**2. Correcting `INJECTED` silently narrowed an unrelated fence — and then widening its domain silently
disarmed it from the other side.** `test_directus_consumer_derivation.py` passed `INJECTED` as the map it
scans, so pruning `cms` made *"cms stopped reading Directus"* and *"cms was never looked at"* the same
green verdict. Giving it its own `DIRECTUS_READER_DOMAIN` then broke it again: the live corroboration
selected its clone root with `all(isdir(...) for repo in DOMAIN)`, so adding a husk cloned **nowhere on
this box** (`skiller`) made the selector return `None` and the test **SKIP**. Both repaired; the assertion
now grades `shipped ∩ present` and **names what was unmeasured**.

**3. Two test doubles modelled states no real clone can be in.** `test_injection.py`'s `_cfg` declared
**four** deleted services and asserted they receive injected images. `dev-stack`'s platform double had no
`repos.yml` at all — every `make init` clone has one — which is why the new arm was the first check in the
script to notice, and 23 registry/set-dress tests went red on a file none of them are about. The fix is to
make the double model a real clone, not to soften the arm.

## Corpus

- `corpus/ops/platform-alignment.md` — **three new §8 rules**: *an OPERATING list must not name a corpse; a
  SCANNING domain must*; *a one-line window over a multi-line construct is a check of a different thing*;
  *a summary that counts its own INPUT cannot disclose a filter*.
- `corpus/ops/setup_guide.md` + `CLAUDE.md` — the lib's path, and the correction that *"the dev path has no
  such step"* is now **half**-true: `dev-stack up N` acquires; the main `N = 0` `make init` + `make up`
  path still needs the manual clone. **Say which dev path you mean — the two now differ.**

## Verification (all CONTENDED — durations are not baselines)

| scope | result |
|---|---|
| `stack-injection` whole section | **335 passed** |
| `dev-stack` whole section | **155 passed** (incl. the new fail-closed regression PAIR) |
| `demo-stack` whole section | **1072 passed / 9 failed** — all pre-existing: 6 demopatch/next-web sha-drift against live clones (`D-M257x-258-3`: 5 stale + 1 chained) + 3 needing a live Postgres. **No manifest, `next-web` or `ant-academy` file is in this iter's diff** (`git status` over the rext tree: 0 matches) |
| `demo-stack` studio + tooling, re-run after the last edit | **198 passed** |
| `shellcheck -S warning` on all three edited scripts + the new lib | **rc 0** |
| `guard_family --repo-root . --platform stack-demo/platform` | **30 GREEN · 0 RED · 0 could-not-check · 5 not-run** (back to the pre-iter baseline after the citation repairs) |
| `stack-core` whole section | **2,209 passed · 8 failed · 3 skipped** (3,042 s) — **but that run READ A MIXED TREE**: it started 19:02Z and the corpus-citation repairs landed 19:05–19:11Z, so 7 of the 8 are stale-read artifacts of files the run had already opened, and the 8th (`test_fence_provenance`) grades the *dirty* fence tree. **All 8 re-run at the settled tree: 138 passed.** Reported this way rather than as "8 failures" because a verdict taken against a tree that no longer exists is not a verdict about this tree |
| the 6 corpus-grading modules, at the settled tree | **167 passed** |
| the 8 stack-core modules the mixed run reddened, at the settled tree | **138 passed** |

> **Lesson, and it is the third instance of the same shape this iter:** *a long test run is a MEASUREMENT OF A TREE, and editing the tree under it invalidates the reading — silently, and only for the files the run had not yet reached.* The 50-minute run cost more than it proved. Run the long sweep **after** the tree settles, or grade only the modules you re-ran.

## The tag, and the pin

| step | result |
|---|---|
| rext commit | `2833a64` |
| tag | `fast-build-m257x-iter-270` (annotated) |
| `git push origin main` | `2ff1547..2833a64` |
| `git push origin <tag>` | `* [new tag]` |
| **verified ON ORIGIN** (`git ls-remote --tags origin`) | `4e5fb251…` → **`2833a64…`** — the M236 pre-flight rung zero, run rather than assumed |
| `.agentspace/rext.tag` | `fast-build-m257x-iter-101` → **`fast-build-m257x-iter-270`** |
| commits spanned by the bump | **206** |

**Force-push was neither used nor needed**; the tag name is new (`git tag -l` and `ls-remote` both
returned 0 matches before cutting it), so the `F18` three-tags-disagreeing pattern is not repeated.

## Close — 2026-08-10

**Outcome:** The frozen-pin control is spent deliberately, and the four routes queued behind it landed
together. Every operating list in rext is free of decommissioned service names; the one derivation that
fell back to the unfiltered list now `die`s; the dev path has the studio acquisition it never had. Two
pre-registrations refuted, and both refutations carried more than the predictions would have.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: **y** — Outcome: exit-7

**Decisions:** `D-M257x-270-1` (the control is spent — and what it bought, recorded), `D-M257x-270-2`
(the `cms` key in `DIRECTUS_DATA_CONSUMERS` is kept, its rollback rationale halved and routed),
`D-M257x-270-3` (an OPERATING list must not name a corpse; a SCANNING domain must),
`D-M257x-270-4` (a citation wrong by ~75 lines was green — the construct guard is a declared FLOOR).

**Side-deliverables:** the `demo_knob_guard --fix` pass rewrote **24** `Read at` citations from the
parsers, and `build-budget.md`'s two anchors were **already** stale before this iter. Both are separate
from the planned scope and do not change the close status.

**Routes carried forward:**
- `FIX-M257x-269-force-append-grows-the-demo-env-without-bound` — declared out of scope in this iter's
  `overview.md`; rides the next tag.
- **`ROUTE-M257x-270-prove-the-spent-pin-cold`** — **new, and the highest-value item.** The pin now points
  at tooling **nobody has run cold**. A `demo-down --purge 2` + `demo-up 2` cycle on the new pin is the
  real test of every path this iter changed, and it should be iter-271's first act.
- `ROUTE-M257x-270-directus-consumer-cms-key-outlived-its-rollback-path` — new (`D-M257x-270-2`).
- `FIX-M257x-267-capture-the-succession-RESPONSE` (gate clause 2),
  `FIX-M257x-266-manual-path-drops-gates-the-automated-path-enforces`,
  `FIX-M257x-265-prose-deletion-instructions-are-out-of-D-reach`,
  `ROUTE-M257x-h59-rext-edits-fire-no-fence-anywhere`,
  `ROUTE-M257x-h65-fresh-checkout-class-needs-a-scheduled-remeasure` → open.
- `ROUTE-M257x-258-the-pin-is-157-iters-stale` → **CLOSED** by the bump.

**Lessons:**
1. **A characterisation that names its own exit is a different artifact from a stale test.** Harden pass
   68 could not fix the fail-open (it needed a tag) so it pinned it *and wrote the repair instruction into
   the assertion message*. That is what made this iter cheap. Prefer it to a TODO.
2. **Correcting one registry can silently narrow an unrelated fence.** Ask, of every constant you prune:
   *who else reads this as their DOMAIN?* Pruning `INJECTED` shrank a Directus fence's reach; widening
   that fence's domain then made its clone-root selector skip. Two opposite failures, one edit apart.
3. **A long test run measures the tree it started on.** 50 minutes produced 8 red modules that were all
   green at the settled tree. Let the tree settle first, or grade only what you re-ran.
