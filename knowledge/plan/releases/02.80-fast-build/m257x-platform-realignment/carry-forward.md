---
milestone: M257x
title: "platform re-alignment — carry-forward"
release: v2.8 "fast build"
close_status: closed-incomplete
close_authority: "USER RULING — `TOK-09` (decisions.md, iter-283, 2026-08-11). NOT a gate-met close."
gate_target: "5 clauses (see overview.md `exit_gate`)"
gate_achieved: "clauses 1–4 MET and proven; clause 5 OUT OF SCOPE BY USER RULING"
gate_distance: "1 clause — and it is a SCOPE RULING, not a measurement. Clause 5 was never met and is NOT being declared met."
last_measured_clause_5: "iter-131 — `P = 29` / `N = 47`, a FLOOR, not a reading of the tree as it stands"
created: 2026-08-11
last_updated: 2026-08-11
---

# M257x — carry-forward

## ⚠️ Read this before anything else: what "closed" means here

**M257x did not close on its exit gate.** It closed because the **user narrowed the definition of done on
2026-08-11** and the narrowed definition was met. The ruling is recorded verbatim as
[`TOK-09`](decisions.md) and its standard is three limbs:

> *"as long as rosetta knows the whole platform architecture, their repos/components/products and knows
> how to build them for me is in line with the goal of the milestone and you can close it"*

| limb | the proof, not the claim |
|---|---|
| architecture | the merge map is machine-checked against the platform's own `repos.yml` **in both directions** — a service entering *or* leaving the clone set turns a guard RED |
| register | all **93** `anthropos-work` repos enumerated, each with a home and a verdict |
| retired services | verified against **code**, not prose — the `roadrunner` and `cms` corrections both came from reading source and `infrastructure`, and both refuted a standing corpus claim |
| buildability | the demo rebuilt **cold ×3**; the dev stack built from current `main` |

**Clause 5 (KB-fidelity GREEN over `corpus/services/**` + `corpus/architecture/**`) is OUT OF SCOPE by that
ruling. It must never be reported as "measured clean", "met", or "green".** Its last measurement was at
iter-131 — `P = 29` / `N = 47`, and that `P` is a **floor**, not a reading of the tree as it stands today.
The scope completed with ~1h45m of the user's clock unspent.

## TL;DR — what the next milestone inherits

**Eleven routed items in two lanes, plus one block fate over a long tail, plus one disclosure.** Nothing
requires the escape hatch: M258 exists, is unstarted, and is in this release, so every item has an
in-release home. **Do NOT route any of this to M257** — its own `exit_gate` still names `odysseus`, retired
by `D-v28-15`, so M257 would inherit a gate that cannot be met as written.

The single most important line in this file: **the two tooling fixes this close landed
(`up-injected.sh`'s preflight and `blocking_state_guard`) are in the `rosetta-extensions` authoring copy
and are NOT on a pushed tag.** A stack consumes rext only at a tag fetched **from origin** (the M217 FATAL
pin guard). Until someone tags and `git push --tags`, those fixes are **unreachable to any stack** — the
exact failure M236 lost a whole iteration to, and which `CLAUDE.md` warns about in its own words: *tagging
is not publishing.*

---

## Cluster 1 — the production-bucket pointer is CONTAINED IN CODE, not on either running stack

**Affected items:** `ROUTE-M257x-h72-dev-path-carries-the-prod-bucket` · `ROUTE-M257x-284-demo-2-is-live-and-uncontained`

**Root cause.** Platform compose **hardcodes** `STORAGE_S3_BUCKET=production-storage…` on `backend`. The
tooling overrode only the *public* bucket, and the store was registered as per-stack-isolated when its
default pointer was production. On **2026-08-11** a demo's studio-desk attempted `s3:PutObject` against
that production bucket and was refused **403 by an IAM policy on an account we do not control** — **not**
by this design. Nothing was written.

**State at close — and this is the part that must not read as fixed:**

* iter-284 closed the pointer at **both** sites (the injected compose override the container reads, and
  `PreflightEnv`), re-classed `s3-private` `PerStackIsolated → SharedPollutionRisk`, and corrected the four
  unqualified *"cannot write prod"* claims in `safety.md`. **That containment is proven by a UNIT TEST on
  the emitter and on NO RUNNING STACK.**
* **`demo-2` predates the fix by nine hours and still carries both pointers.** So does the dev stack: the
  container-side strip is **demo-only** — `/dev-up` applies only `gen_override.py`, which emits no strip.
  Measured live with `docker inspect`, not inferred.
* What bounds the dev path today is **the absence of credentials, not a pointer override.** An operator
  whose `platform/.env` carries AWS keys re-opens on dev what iter-284 closed for demos.

**Estimated scope:** small on the demo path (already emitted; needs a tag + a bring-up to take effect),
medium on the dev path (a real design call — the remedy costs dev uploads a broken-thumbnail trade, which
is an operator's decision, not an agent's).

**Fate:** **LAND-NEXT → M258.** Not the escape hatch: `TOK-09` pre-authorises the routing, it is disclosed
at `corpus/ops/safety.md` §2.1/§2.2 and now at the guarantee itself, and it is bounded by default.

**Provenance:** `decisions.md` `TOK-09` priority 1 · iter-284 · hardening-ledger pass 72 · this close's
`deferrals-audit.md` §13.

---

## Cluster 2 — the whole-section instrument is partially blind on this host

**Affected items:** `ROUTE-M257x-h73-suite-census-unittest-path-is-dead-on-this-interpreter` ·
`ROUTE-M257x-h73-battery-stage-stdlib-set-is-interpreter-versioned` ·
`ROUTE-M257x-h73-readme-documents-a-python3-that-has-no-pytest` *(this last one **LANDED at close**)*

**Root cause — one defect, eight symptoms, proven both ways.** `suite_census`'s second runner shells
`python -m unittest <section-relative.dotted.name>`. **None of the five rext `tests/` directories has an
`__init__.py`**, and on `python3.12` unittest's loader cannot import a namespace-package submodule by
dotted name. Measured in a scratch tree: identical module, `FAILED (errors=1)` without `__init__.py`, `OK`
with it. The census reads the resulting `_FailedTest` as *"Ran 1 test"* + RED, so the population arm
cascades. **Consequence, stated precisely: the two-runner cross-check — the thing that makes
`both_runners_report_the_same_executed_count` meaningful — is currently a ONE-runner check that reports
RED.**

A second, independent interpreter fact: `test_battery_stage`'s stdlib derivation cross-checks
`sys.stdlib_module_names` and misses the **PEP 594** removals (`spwd`, `msilib`, `_msi`, `ossaudiodev`) —
a property of the interpreter, not of this tree.

**This is why the final harden pass stopped at `cap reached without stabilization`.** Coverage could not
stabilize because the instrument measuring it is itself partially blind here. **That is the finding, not a
gap in what was measured.**

**Estimated scope:** medium. The fix is a **contract change** to `run_one` (or five `__init__.py` files
that would change pytest's import semantics tree-wide) — a design decision, which is why a harden pass
correctly declined to make it.

**Fate:** **LAND-NEXT → M258.** M258's gate is *"UP, and every journey verified"*; a census that cannot
run one of its two runners is squarely in its path.

**Provenance:** hardening-ledger pass 73.

---

## Cluster 3 — fences that are not RED-proven, and two that no census can see

**Affected items:** 22 of 28 milestone-added guard/census modules carry **no mutation battery** ·
`stack-injection/exposure_claim_guard.py` and `demo-stack/backend_api_url_server_reader_guard.py` belong
to **no family census** and carry **no provenance stamp** · `guard_family.run_one` has **no `timeout=`** ·
`buildbench`'s single `Popen` has no timeout and no cleanup path · 21 test files run **zero tests and exit
0** when executed directly · root discovery is duplicated **four** times with a **divergent** fourth copy ·
`up-injected.sh:1634` still iterates a **hand-maintained** two-element app-patch tuple where
`patch_anchor_guard` derives the same population from disk · the same "empty subject" condition is graded
**rc=1 by one guard and rc=2 by five**.

**Root cause.** This milestone built its fence family fast and under a clock, and the family's own
promise — *"run the WHOLE guard family, and name every member"* — holds **inside `stack-core/` and stops
silently at its boundary**. The repo's convention is that a fence must be **proven able to go RED**;
`repair_postcondition`'s registry and `guard_family.union()` both glob `stack-core/` only, so a fence
living in a sibling section is structurally invisible to the runner that claims to name every member.

**Estimated scope:** large — 22 mutation batteries is a milestone's worth of work on its own, and several
of the others (the rc=1/rc=2 split, the shared root-discovery helper) are contract decisions.

**Fate:** **LAND-NEXT → M258**, and expect it to be **triaged rather than exhausted**. The four highest-risk
modules are named, because each already produced an unearned green **inside this milestone**:
`clone_pin_guard` (`e64a3cd`), `platform_alignment_guard` (`f5aad69`, `1149af4`),
`route_disposition_guard`, and `claim_census_guard` (three separate false-verdict fixes: `ec08f4d`,
`7ecfcb0`, `bedfcd0`).

**Provenance:** this close's Phase 2 + Phase 4 reviews.

---

## Cluster 4 — the benchmark harness cannot grade the host the release names

**Affected items:** `hostprofiles/` holds only `billion.json` and a `laptop.json` describing a **different,
retired** Mac · `demo-stack/tests/test_frontend_build.py:683-689` hardcodes `billion.json` · `buildbench`
asserts **nothing about elapsed build time**.

**Root cause.** `D-v28-15` (2026-07-31) retired `odysseus` and moved dev/test **local to the new Mac** —
and that supersession reached `knowledge/` **35 times and this corpus zero times** until iter-226. So
`build-budget.md` still argues throughout for a gate host that no longer exists, and **no profile has ever
been measured for the host that replaced it**, which makes gate clause 1 of the *build budget* not
gradeable today. Separately, the release is called **"fast build"** and the harness ships **no cycle-time
threshold at all**: a regression from the 666.29 s p50 baseline to 900 s passes every gate it has. Its
**headroom** assertions, by contrast, are exemplary — three clauses, each independently RED-proven, plus a
control proving a `None` input fails rather than silently skipping.

**Estimated scope:** small for the missing profile (measure the host); small for the hardcoded fence
(iterate `hostprofiles/*.json` — `test_baseline_mirror_fence.py` already does exactly this and the pattern
was not reused); medium for a wall-clock threshold, which needs a negotiated target.

**Fate:** **LAND-NEXT → M258.** M258's gate is a **p50 ≤ 480 s** number; it cannot be graded by a harness
that asserts no time target, on a host with no profile. This is a **precondition of M258's own gate**, not
an adjacent nicety.

**Provenance:** iter-225, iter-226; this close's Phase 4 review. Two stale `buildbench.py` line citations
in `build-budget.md` were re-derived **at this close** (`:1205`/`:1217` → `:1472`/`:1484`; `:1470` →
`:1538`).

---

## Cluster 5 — residual per-item content work, deliberately not absorbed

**Affected items:** `ROUTE-M257x-285-demo-2-cockpit-serves-a-stale-world` ·
`ROUTE-M257x-286-next-web-manifest-baselines-have-drifted` · `ROUTE-M257x-282-prose-twin-REPORT-tier-residual` ·
`ROUTE-M257x-282-intra-tree-prose-twins` · `ROUTE-M257x-h73-demo-stack-live-arms-red-while-a-demo-is-up` ·
`ROUTE-M257x-289-blocking-guard-checks-presence-not-disposition`

**Root cause.** These are the **scope-creep tripwire working as designed.** `TOK-09` set a **closed** list
and instructed that anything found outside it is *"recorded and routed, never absorbed."* Each of these was
found after that ruling. They are not process failures; routing them is the compliant outcome.

Two deserve a sentence each:

* **`h73-demo-stack-live-arms-red-while-a-demo-is-up`** is arguably **not a defect at all** — nine
  live-clone/live-container tests assert against a real clone, and the real clone is in use by the demo the
  user is validating on. What *is* a defect is that **nothing in the suite's output says so**, so a reader
  sees nine reds and cannot tell an in-use clone from a broken patch chain.
* **`289-blocking-guard-checks-presence-not-disposition`** was found **by running the fence at this close**.
  `blocking_state_guard` checks that a blocking grading is **named** in the deferral audit — not that its
  recorded **disposition is current**. It reported iter-119 as "represented" while the file still called it
  *"🔴 OPEN — the milestone is holding on it"*, **164 iters after `TOK-09` closed it**. That is *green over
  something never checked*, inside the fence built to catch exactly that.

**Estimated scope:** small each.

**Fate:** **LAND-NEXT → M258.**

**Provenance:** iters 282, 284–286; hardening-ledger pass 73; this close's deferral re-audit.

---

## Block fate — the long tail of carried tokens

**The mechanical sweep counts 487 distinct tokens, 215 carried across ≥3 distinct iter directories.** That
number must be read for what it is: **a marker count, not a state count.** The iter tree is append-only and
closure is written as prose *beside* a token, never as a deletion — so **a token that stops being restated
is indistinguishable from one that was fixed.** No per-item state claim is made for the 215.

**The real pattern is not re-deferral. It is going quiet.**

**Block fate: → M258**, as one decision, recorded here so it is a conscious routing rather than an unfated
drop. The named quiet carriers worth a first look are
`CHECK-M257x-iter52-second-ai-manager` (24 restatements, no closure), `ROUTE-M257x-235-*` ×2 (21 each),
`ROUTE-M257x-236-disclosure-scope` (20), and `DEF-M257x-iter101-briefing-rext-tree` (17, self-priced at
"~30 min").

---

## Projected post-resolution state

With clusters 1–5 and the block fate resolved in M258, the release enters its closer with:

* both stack families containing the production-bucket pointer **on a running stack**, not only in an
  emitter unit test;
* a two-runner section census that actually runs **both** runners, so
  `both_runners_report_the_same_executed_count` becomes a check rather than a claim;
* a benchmark harness with a **measured profile for the host it runs on** and a **wall-clock assertion**,
  which is what makes M258's `p50 ≤ 480 s` gate gradeable at all;
* the fence family RED-proven past `stack-core`'s boundary.

**Clause 5 is not on this list, and that is deliberate.** It is out of scope by user ruling. If a future
milestone wants it, it needs a **new** scope decision from the user — not an inheritance from this file.

## Cross-references

* [`overview.md`](overview.md) — the 5-clause `exit_gate` as written
* [`decisions.md`](decisions.md) § `TOK-09` — the user ruling this close rests on, verbatim
* [`deferrals-audit.md`](deferrals-audit.md) §12–§13 — the machine-derived blocking-state sweep
* [`hardening-ledger.md`](hardening-ledger.md) passes 72–73 — the final harden and its
  `cap-without-stabilization` stop condition
* [`progress.md`](progress.md) § *Ledger completeness* — the 8 iters reconstructed at this close
* [`../m258-proven-live-build/overview.md`](../m258-proven-live-build/overview.md) — the target milestone
* [`corpus/ops/safety.md`](../../../../corpus/ops/safety.md) — both guarantees, now qualified
