**Type:** tik

# iter-132 — the hedge the read retired, and the fence that called the retraction a hedge

**Resumed, not opened.** A prior session committed `fix(M257x/132)` (`1012679` — one line in
`org-repos.md`) and stopped before creating this dir. That commit is this iter's first deliverable.

---

## 1. Step 0 — re-survey, and a correction to the route I inherited

`unreadable_repo_claim_guard` at open: **OK**, with its standing NOTE — *"9 site(s) still hedge about
`infrastructure` while 13 report having read it."* The target is live and unabsorbed.

**The route over-stated by one file, and it is the file every agent loads.**
`FIX-M257x-iter131-infrastructure-hedge-stale` reads *"11 sites **+ `CLAUDE.md`**"*, with the
aggravating note that *"`CLAUDE.md` publishes it too."* It does not — corrected at iter-124, and it
carries the settled reading twice (`:194-203`, `:259`). Its remaining `clone set` mentions are
`stack-demo`'s (`:123`, `:132`) and `customerio-sync`'s `go.mod` (`:293`), neither of which is this
predicate. → `D-M257x-132-4`.

## 2. Width, measured before repairing (§5 rule 57)

Four independent searches over `corpus/**` + `CLAUDE.md` + `README.md`:

| search | hits |
|---|---|
| `not measurable\|unmeasurable\|un-measurable` | 33 lines / 19 files |
| `clone set` ∧ `infrastruct` | 25 lines / 13 files |
| `assert neither\|report both\|do not assert either way` | 13 lines |
| union of the first two, line-level | **22 lines / 11 files** |

Triaged, **and the majority of the union is NOT the predicate** — which is the whole reason rule 57
asks for the measurement first:

| group | sites | verdict |
|---|---|---|
| **A** — cms's M810 prod state inferred UNMEASURABLE from the clone-set premise | **8** | **FALSE — repaired** |
| **B** — the production RPC address hedged on the same premise | **7** | **stale premise — repaired** |
| C — jobsimulation's **GitHub archive state** (lives in the org API, never in a clone) | 3 | TRUE — left |
| D — the router banners' *"Vercel runtime configuration, in no clone set"* | 3 | TRUE, different subject — left |
| E — already-corrected sites + quotations of the retired hedge | 5 | left |
| F — `backend.md:51` (`customerio-sync`), `staging-bringup.md:461` (`colony`) | 2 | different repos — left |

## 3. The premise was one `git clone` away, and nobody had spent it

The corpus hedges group B because *"no `.tf` file in any clone names
`http://backend.internal.anthropos:8081`"* and *"the deciding declaration lives in `infrastructure`,
which has never been in any clone set."* Both conjuncts true. **The conclusion drawn from them —
UNMEASURABLE — is not**, and settling it cost one `--depth 1 git clone`, well under a minute.

> **`HEAD` = `13c248e64935faa467806c7429a6fabf8b1d5a37` — the exact sha the corpus already cites 28
> times.** So this is not a new reading of a moving target; it is the *same* ref, re-obtained.

**PRODUCTION DOES NAME THE LITERAL, EXACTLY ONCE:**

| what | where |
|---|---|
| `cms_rpc_address = "http://${local.private_dns_services_2.backend}:8081"`, an input to `module "backend_euwest1"` | `terraform/production/services.tf:346` |
| `backend = "backend.internal.anthropos"` | `terraform/production/locals.tf:22` |
| **only `.tf` occurrence in the repo** — two independent searches (`rpc_addr` repo-wide; `8081` over every `*.tf`) | agree |

**So M809 landed in production too, and in the same shape as locally: the address the platform hands
`backend` for "cms" is `backend` itself.** There is no `skiller_rpc_address` and no
`jobsimulation_rpc_address` in production terraform at all — which also makes `skiller.md`'s old flat
assertion wrong about the *variable*, not only the tense.

**Two riders the same read produced, and both are the platform contradicting its own prose:**

1. `services.tf:352-355` explains the absent `jobsimulation_rpc_addr` input by pointing at
   *"`module.messenger_euwest1`"* — and `:618-621` **of the same file** says that module is deleted.
   Stale **in the platform's own terraform**.
2. `infrastructure`'s own narrative docs (`CLAUDE.md:85-86`, `knowledge/architecture.md:87-88`,
   `knowledge/service-dependencies.md:115-116`) all describe `module.messenger_euwest1` wiring **four**
   RPC addresses to backend. That module does not exist. **§6's rule, met in the wild for the second
   time: the platform's CONFIG is its documentation of record; its NARRATIVE docs are not.**

**Corroborations taken in the same pass, none of which moved a corpus claim:** `services.tf` is **666
lines** (as `org-repos.md` says); it declares **10 `module` blocks, 9 of them service repos** (as
`org-repos.md` says — the 10th is an ACM certificate, not a service); `module "jobsimulation_euwest1"`
**does** survive at `:475` and carries only bucket/SSM/atlas inputs, exactly as the corpus states.
**Upheld claims are counted as results here** (§5, iter-130's discipline).

## 4. The repair — 15 sites, 8 files

**Group A (8)** — the conflation, replaced by the settled verdict, each naming `13c248e6`:
`cms.md:16` · `backend.md:13` · `backend.md:86` · `external_services.md:175` · `storage.md:175` ·
`dependency_map.md:31` · `jobsimulation.md:12` · `platform-migration-status.md:89`.

**Group B (7)** — the RPC hedge, replaced by the measurement, derived **once** in `backend.md`'s *RPC
re-pointed, then un-set* bullet and cited from the rest: `backend.md:186` · `backend.md:327` ·
`cms.md:61` · `cms.md:218` · `jobsimulation.md:54` · `skiller.md:26` · `dependency_map.md:27`.

Plus `org-repos.md`'s `infrastructure` row, upgraded from the iter's own first commit's *"sound, not
re-checkable"* to **"when a claim turns on this repo, clone it — do not hedge it."**

**Every repaired sentence keeps the true conjunct.** `infrastructure` really is not in the standing
clone set; `make init` does not clone it and no `stack-*/` holds it, so the ~31 corpus lines citing it
are **sound but not re-derivable in place**. What is retracted is the inference to *unmeasurable* — and
it is retracted, not softened (`D-M257x-129-4`'s rule).

## 5. The fence fired on the repair — twice correctly, and once because it cannot see the difference

Re-running after the sweep, the NOTE read **"11 site(s) still hedge … while 14 report having read
it"** — **up from 9/13**. The repair appeared to make the corpus hedge *more*.

**Instrumented rather than argued.** Of the 11:

| bucket | count | what it is |
|---|---|---|
| marker **and** ref-pinned reading | **8** | **retractions** — they quote the retired wording in order to retract it |
| marker alone | 3 | 2 of them **paragraphs iter-132 had just written** |
| — | | |

- **The 2 were real defects of mine.** The new riders assert `module.messenger_euwest1` is deleted with
  no sha in that paragraph. **Fixed by carrying `13c248e6` into the riders**, not by touching the fence.
  → `D-M257x-132-3`. The 3rd is `platform-alignment.md:2618`, the protocol doc's own historical worked
  example, which is correct as it stands.
- **The 8 are the instrument's limit, not the corpus's.** `architecture_overview.md:227` and
  `platform-migration-status.md:93` are the corpus's *model* retractions and were being counted as live
  hedges, because the guard matches markers by substring and tests `marker` **before** `measured`.

**Fixed in the instrument, not in the prose** (`D-M257x-132-2`): a third bucket `mixed`, and a
`KNOWN_WEAKNESS` line saying in the guard's own voice that it **cannot** tell a quoted retraction from a
live hedge. Post-fix:

```
unreadable-repo-claim-guard: OK — all 26 `module.*_euwest1` mention(s) are satisfied: 1 by an
  unmeasurable marker ALONE, 8 by both a marker and a REF-PINNED reading (a RETRACTION, or a
  mixed paragraph), 16 by a REF-PINNED reading of `infrastructure` alone.
unreadable-repo-claim-guard: NOTE — 1 site(s) hedge ... with NO ref-pinned reading ...
unreadable-repo-claim-guard: KNOWN_WEAKNESS — 8 paragraph(s) carry BOTH ... Read the paragraphs.
```

**Live hedges: 9 → 1.** The remaining one is the protocol doc's worked example.

### Controls, accounted honestly

3 tests added. A **meta-mutation deleting the `mixed` bucket kills 2 of the 3.** The third survives its
own mutant — so it is **renamed** `test_MUTATION_…` → `test_FIXTURE_INTEGRITY_…` with the survival
stated in its docstring. *A control that cannot fire is this milestone's most-caught defect class;
naming one MUTATION while it survives is that defect with better branding.*

The real-corpus anti-vacuity floor was re-cut **from `hedged ≥ 9` to `hedged + mixed + measured ≥ 22`** —
because a floor pinned to one bucket would have to be re-cut every time a hedge is legitimately
retired, i.e. the ratchet would argue against the repair it exists to enable.

## 6. Test gates

- `tests/test_unreadable_repo_claim_guard.py` — **30 passed** (was 26; +4).
- Meta-mutation — **2 of 3 new controls killed**, as stated above.
- **Guard family: 18 GREEN · 0 RED · 4 not-run** (`anchor_offset_guard`, `repair_leak_guard`,
  `repair_reach_guard`, `value_change_guard` — commit-/input-scoped, no `--range`/`--ledger`). **Not a
  whole-family green, and the runner's own summary says so.** It went **RED once mid-iter, on this
  iter's own §5 rule 61 text**, which asserted `module.messenger_euwest1` was deleted without naming a
  ref — *the rule fired on its own worked example.* Repaired by citing `13c248e6`, not by touching the
  fence.

### The whole-suite claim (§5 rule 51) — and one run is DISCLOSED AS CONFOUNDED

```
cd .agentspace/rosetta-extensions/stack-core
/usr/bin/python3 -m pytest tests/ -q --tb=line -p no:cacheprovider --no-header --durations=3
```

**Run 1 — CONFOUNDED, quoted only as a finder:** `2 failed · 1206 passed in 1650.37 s (27:30)`. The
tree was edited *and committed* while it ran (the bucket-ordering fix and its regression test landed
mid-run), and a guard-family run executed concurrently — so neither its counts nor its **+56 % wall
time against iter-122's 1055.54 s** is a result. **It nonetheless did its job**, exactly as iters 111 /
121 / 122 / 128 disclosed theirs:

| failure | verdict |
|---|---|
| `test_claim_twin_guard_iter48_answer_key::test_02` | **the standing, documented RED** — re-attested by a full run, not carried |
| `test_fence_provenance::TestFamilyRefusesAnUnstateableTree::test_the_escape_accepts_and_records` | **an artifact of the confound** — the fence tree was DIRTY while it ran. **Re-run alone on the committed tree: `1 passed in 83.11 s`.** Not a defect; proven rather than argued |

**Run 2 — CLEAN**, taken after the rext tree was committed and pushed, nothing else running:

```
      ->  1 failed · 1208 passed  in  2077.16 s (34:37)   [rext 223e4a6, corpus at iter-132]
```

**The one failure is `test_claim_twin_guard_iter48_answer_key::test_02` — the standing, documented
RED, re-attested by a full run rather than carried.** `1208 passed` vs iter-122's `1146` = **+62**, of
which 4 are this iter's. **Run 1's second failure did not reproduce**, which is the point of taking
run 2.

**And rule 51's timing check FAILS, so it is reported as a failure rather than smoothed.** 2077.16 s is
**+96.8 %** against iter-122's 1055.54 s baseline — and the decisive number is not that one:

| run | wall | conditions |
|---|---|---|
| run 1 | **1650.37 s** | confounded — concurrent guard-family run, tree edited mid-run |
| run 2 | **2077.16 s** | clean, alone |

**The clean run was 26 % SLOWER than the contaminated one, within the same hour, on the same host.**
Whatever explains the +96.8 % is not concurrency from this session, and it is not test growth (+5.4 %
in count). `test_m257x_mechanical_fences_mutation_battery::test_01` alone moved **636.94 s → 866.25 s**
between the two runs and is **41.7 % of the clean total**. **This host is not currently a stable timing
substrate, and no wall-time claim from it should be treated as a measurement** — iter-128 already
reported +48.7 % against the same baseline and named its contaminants; this iter cannot name one.
Routed as `FIX-M257x-iter132-suite-walltime-is-not-a-measurement`. The COUNTS are unaffected and are
what this iter's gate rests on.

---

## Close — 2026-08-07

**Outcome:** the milestone's largest measured cluster (iter-131 P1, 19 of 80 blockers) is **repaired at
15 sites in 8 files**, and its premise turned out to be settleable rather than merely re-wordable — one
`--depth 1` clone at the sha the corpus already cites established that **production names
`http://backend.internal.anthropos:8081` exactly once**, retiring a *second* hedge nobody had routed.
The fence then caught **two unref-pinned assertions in the repair itself** (fixed) and revealed that
**8 of its 11 "hedges" were retractions** it cannot distinguish by substring — fixed in the instrument
with a disclosed third bucket, taking **live hedges 9 → 1**. **No reading was taken; no `N` movement is
claimed.**
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5**, unchanged. Clause 5 is met only by a reading that returns zero, and this
iter took no reading.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**no successor strategy is authorable — `TOK-08`'s sealed refutation branch bars one; this iter runs under the user's direct brief, as iters 121–131 did**) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (1 tik) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-132-1` … `D-M257x-132-4` (see [`decisions.md`](decisions.md))
**Side-deliverables:**
- `platform-alignment.md` §5 **rule 61** — *"not in the clone set" is a fact about our habits; it never
  entailed "not measurable"* — with both riders earned by this iter's own prose.
- Two **platform-side** stale-comment findings recorded in `backend.md` (a `services.tf` comment naming
  a module the same file says is deleted; three `infrastructure` narrative docs describing it). Zero
  platform edits — recorded in our corpus, not theirs.
**Routes carried forward:**
- `FIX-M257x-iter131-adjudication-independence` — **still the first item.** iter-132 did not touch it;
  it is the highest-value open route and needs independent agents, not a repair pass.
- `FIX-M257x-iter131-my-three` — P7 (§1 missing `library-unimported`), P19 (the over-claimed curl), P5
  (`architecture_overview.md:83` still lists `ai`). Untouched by this iter.
- `FIX-M257x-iter131-predicate-sets-not-enumerated`, `FIX-M257x-iter131-root-mount-count-underived` —
  untouched.
- **NEW — `FIX-M257x-iter132-infrastructure-is-cloneable-so-clone-it`:** ~31 corpus lines cite
  `infrastructure` and none is re-derivable in place. The repo clones in under a minute. **Either add it
  to a clone set or write down that we re-clone on demand** — the third option, hedging, is what rule 61
  now forbids.
- **NEW — `FIX-M257x-iter132-marker-fences-cannot-see-retractions`:** `unreadable_repo_claim_guard` is
  fixed; **the same substring-vs-retraction blindness plausibly affects every marker-matching fence in
  the family.** Nobody has checked the others.
- **NEW — `FIX-M257x-iter132-suite-walltime-is-not-a-measurement`:** the whole suite ran **1650 s
  confounded and 2077 s clean within one hour on the same host** — the clean run slower — against a
  1055 s baseline. §5 rule 51 asks for an expected wall time so an operator can say *"this is normal"*;
  on this host that check currently cannot discriminate. Either re-baseline against a measured host
  profile or state in rule 51 that the timing leg is suspended. **Counts are unaffected.**
**Lessons:**
1. **A hedge outlives its premise silently, because a hedge looks like diligence.** Eleven sites went on
   publishing a limit four days after one clone retired it, and the corpus cited that clone 28 times in
   the same tree.
2. **Measure the premise before you improve the wording.** The planned deliverable was prose; the
   valuable half was a `git clone` the plan did not contain.
3. **A fence that reads by substring will report your retraction as the disease** — and the fix is a
   disclosed bucket, never laundered prose. Second independent arrival at `D-M257x-121-4`'s ruling.
