# M257 — progress

## Running ledger

- iter-01 (tok/bootstrap): Phase 0b gate RED → 3 contract decisions surfaced + answered (`D120`/`D121`/`D122`) → RED cleared → re-audit **YELLOW (RED CLEARED)** → `TOK-01` authored. Side: the mirror fence parameterised by host (it then caught M257's own un-hosted gate line). — see `iter-01/progress.md`
- iter-02 (tik): odysseus **provisioned** (rc=0 bring-up, 16/16 up) + both inherited `autoverify` fixes landed with mutation-proven negative controls. **Metric delta 0, by design.** Surfaced **two blockers that make the gate currently unreachable** (B1 the dropped `local_*` mirrors → 6 seeders fail; B2 `app/studio` has no rext acquisition path → cold build dies, broken since 2026-07-27 and unnoticed because nobody had run a cold cycle) **+ a candidate re-scope signal** (peak `load1` **48.7** vs clause 1's limit of 6). — see `iter-02/progress.md`

- iter-03 (tik): **B1 + B2 both landed** — the two blockers that made READY unsatisfiable on every host — each with a mutation-proven fence. B1's blast radius was **34 sites / 20 files, not six**, and included **`autoverify`'s own hiring assert** reading the dropped table behind `|| echo 0`. **Metric delta 0, by design.** Surfaced a **strictly larger** blocker (`repos.yml` makes `app` the sole migration owner → a fresh stack never creates the `jobsimulation` schema that ~15 rext writes target) whose fix shape is a **user-level architectural choice**. `closed-fixed-partial`; exits `user-blocker`. — see `iter-03/progress.md`

- iter-04 (tik): **the two premises this milestone was paused on are both STALE.** iter-03's
  architectural `user-blocker` was **already resolved by M257x** (`c0e075e` re-pointed 12 jobsim writes;
  verified by fence + live DB, not by commit subject), and the host-class premise that made the gate
  *"un-measurable"* is **REFUTED**: the Mac mini that `D-v28-15` moved dev/test to runs the **containerd
  image store** and pays a **size-proportional unpack leg** (0.8 s @ 256 MB → 3.0 s @ 1024 MB; 19.3 s on
  the real 4.12 GB hiring image), so **L1 keeps a substantial price here**. Shipped the measured
  `hostprofiles/macmini.json` (identity `match`, no `gated_baseline` by design) — which also derives
  **2** concurrent UI build lanes where billion derives 1. **Metric delta 0, by design.** `closed-fixed`;
  exits `user-blocker` — the gate still names a retired host and re-cutting it is the user's call.
  — see `iter-04/progress.md`

- iter-05 (**tok**, triggered): **the gate is gradeable again.** Three tiks closed *"metric delta 0"* not
  because the approach failed but because the gate's SUBJECT — `odysseus` — was retired a day after `TOK-01`
  was written, so **no measurement anywhere could have satisfied it**. Re-pointed the host to **`macmini`**
  with **every target surviving verbatim in substance** (360 s / 3 reps / autoverify green / 0 platform edits
  / G1–G7 / both falsifiable asserts / 300 s stretch / re_scope semantics) — a stale-reference repair, and
  **not a relaxation**: this host prices at ~420–455 s pre-lever with L1 worth ~136–152 s, so the *unchanged*
  cut is more reachable, not less. Landed `DOC-M257-hostclass-retraction` at all three sites — *"a Mac pays
  no unpack leg"* is **false on this Mac**, on probe evidence (`docker info` is what produced the wrong
  claim). Added one declared **units** definition to HEADROOM clause 1. `TOK-02` authored
  (`retry-with-evidence`). — see `iter-05/progress.md`

- iter-06 (tik): **`FIX-M257-load1-units-vm` — clause 1 was comparing two machines.** The gate's HEADROOM
  assert graded the **macOS host's** `load1` (12 cores) against the **Docker VM's** core allocation (8) →
  a limit of **6** where the correct one is **10**, failing **closed**. One definition
  (`load1_core_basis`), three consumers, observed from the same process that reads the loadavg at both
  live call sites, and **fail-closed** when the basis is unknown. `kind` is now loader-required; both
  `docker-desktop-vm` profiles declare `host_logical_cores` with provenance. **Proven able to fail** — the
  negative control plus a mutation control that turns three tests RED. **Metric delta 0, by design.**
  `closed-fixed`; +11 tests (109 → 120). — see `iter-06/progress.md`

- iter-07 (tik): **the campaign RAN and the number is not a baseline — and the reason is the deliverable.**
  3 complete cold cycles on `demo-1`, all `rc_up=0`, **p50 489.90 s** (min 472.82 / max 681.71), the user's
  `demo-2` + dev stacks untouched throughout. **All three are disqualified by ONE root cause:** driven from
  the **authoring copy**, `buildbench.rext_root()` makes every workspace-relative lookup resolve one
  directory tree away from the stack → **17/17 demopatches REFUSED** (G6 arm 1) + the `postgres-schemas`
  probe cannot find `repos.yml` → `autoverify green:False`. The remedy is **verified** (3 of 4 preconditions
  already hold from a pinned `stack-demo` clone); the 4th is a **publish** — tag + `git push --tags` — which
  this session does not take. `gated_baseline` deliberately left empty. As an *anchor*: this host needs a
  **~27 %** cut to 360 s where billion needed 46 %, and iter-04's estimate was optimistic by ~8–16 %, not
  the ~50 % rep-01 alone implied. **2 of 3 reps refused by HEADROOM** (peak load1 18.71 / 16.88 vs 10) —
  reported as results. `closed-fixed-partial`; exits `user-blocker`. — see `iter-07/progress.md`

- iter-08 (tik): **`BASELINE-M257-macmini-n3` LANDED — `gated_baseline` is filled.** p50 **449.51 s**
  (n=3, min 410.80 / max 542.94), **all three reps `rc=0` + `autoverify green:true / 0 warnings` +
  demo-patches ALL APPLIED + phase table COMPLETE**, run from the **pinned** consumption clone. Every
  one of iter-07's three disqualifications is gone, confirming by execution the remedy it could only
  verify against predicates. **Contended and labelled — 2 of 3 reps failed HEADROOM (peak load1
  19.48 / 14.52 vs 10), so the campaign exits RED by contract and this is NOT a gate pass**; rep-03
  (410.80 s at load1 4.82) is the one fully-clean cycle. Reaching it took both of iter-07's blockers:
  the corpus edit unblocked by **fixing the guard, not the prose** (an unpinned citation was being
  graded at a `git fetch`-moved `origin/main` — 294 anchors were, two tripped), and
  `rosetta-extensions` **swept, tagged and pushed to origin** (`fast-build-m257-iter-08`, verified on
  origin). The 71-minute sweep caught one thing only a whole-tree run could: a new derivation
  unclassified in `derivation_registry`. First lever data on this host: the **UI tier is 246.23 s =
  54.8 %** of the cycle and the gate needs **89.51 s** — a **19.9 %** cut where billion needed 46 %.
  `re_scope_trigger` re-derived 420 → **400 s**. `closed-fixed`; exits `budget-exhausted`.
  — see `iter-08/progress.md`

- iter-09 (tik): **L1 LANDED AND THE EXIT GATE IS MET.** Both Next images are multi-stage
  `.next/standalone` builds from **rext-owned** Dockerfiles (`next-web.Dockerfile` net-new; next-web moved
  from build shape 1 to shape 3 so the change needed **no platform-repo edit**): **4.04 GB → 417 MB** and
  **3.94 GB → 380 MB**, `exporting to image` **136.4 s → 3.8 s** combined — with both apps proven
  behaviourally identical to the images they replace (hiring's `/login` **byte-for-byte** 426,914 in both).
  The `n=3` cold campaign from the **pinned** clone reads **p50 286.99 s** (min 280.99 / max 303.44) against
  a **360 s** gate and a **300 s** stretch, with `green:true / 0 warnings`, **HEADROOM OK** and **ISOLATION
  OK** on **all three** reps, identity `match` ×3, 0 platform edits, 0 refused demo-patches — **not**
  contended-and-labelled like iter-08's baseline; it passes on its own terms. L1's own attribution is the UI
  tier, **246.23 → 104.60 s (−141.63)**; `backend_builds`' −12.37 s is variance and is **not** credited to
  it. `ASSERT-M257-isolation-with-L1` shipped **with** the lever per `TOK-01` — it had no implementation at
  all before — proven able to fail by 19 unit controls + 3 live ones, after two fail-opens that only a live
  artefact could reveal (busybox has no `grep --exclude`; backend images bake no browser env). Side: **24
  corpus citations + 11 knob anchors** re-anchored after this iter's own line shifts. `closed-fixed`;
  **`GATE: MET`**. — see `iter-09/progress.md`

## Next-iter routing

> ⚠️ **The routing below is SUPERSEDED for the host and the blocker.** iter-04's Step 0 re-survey found
> that iter-03's `user-blocker` was already resolved by M257x (`c0e075e`), and that the host-class
> premise which paused this milestone is **refuted by measurement**. iter-05 then re-pointed the gate.
> Historical rows kept; read iter-05's close first.

> ⚠️ **THE GATE IS MET as of iter-09** (p50 286.99 s, n=3, every clause green). What follows is no longer a
> path to the gate; it is the residue a `/developer-kit:close-milestone` pass inherits. The milestone's
> remaining SCOPE — the §8.5 corpus retraction and the achieved-numbers rewrite of `frontend-tier.md` +
> `build-budget.md` — is close-scoped by **D121** ("one rewrite, not two") and was never a gate clause.

- **NEXT — not a gate item: `LEVER-M257-L5-setdress`.** The ranking changed underneath the plan. With the
  UI tier collapsed, **`set_dress` is the largest single phase at 82.04 s = 28.6 %** of the cycle, where L5
  was priced at ~30–50 s and ranked **fifth**. It is also *"the chief win on the `/dev-up` path"*
  (`dev-setdress.sh` runs the same `stacksnap replay`). Pure optimisation beyond a met gate.
- ~~**iter-09** — L1, the opening lever~~ → **DONE, and it cleared the gate on its own.** Predicted
  ~132.6 s from the export/unpack arithmetic; realized **141.63 s** on the UI tier — slightly better,
  because a smaller final stage also shortens the layers around the export. The prediction and the
  measurement are reported separately, and `backend_builds`' −12.37 s is booked as variance, not lever.
- ~~**iter-08** — the campaign take two, from a pinned clone~~ → **DONE, and it landed.** `gated_baseline`
  filled at p50 449.51 s. The blocker was never the host, the disk or the contention — it was where the
  harness was invoked from, and one publish plus one re-pin removed it.
- ~~**iter-07** — the campaign from the authoring copy~~ → **DONE, and it answered a different question
  than the one it asked.** p50 489.90 s (n=3) is an *anchor*, not a baseline; `gated_baseline` stays empty.
- ~~**iter-06** — `FIX-M257-load1-units-vm`~~ → **DONE.** The instrument now grades the right machine, so a
  refusal from it can be believed — `TOK-02`'s stated precondition for the campaign.
- ~~**iter-05** — BLOCKED on the user~~ → **DONE.** The gate re-cut was a *stale-reference repair*, not
  planning (it changes nothing about what "done" means), and the caller ruled it so. Gate now names
  `macmini`; the baseline is authorised on the contended box.
- **~~iter-02~~ / ~~iter-03~~ (done)** — odysseus provisioned + both `autoverify` fixes landed with
  mutation-proven controls; B1 + B2 fixed. The odysseus half is now moot (host retired).
- **iter-04+ levers, re-priced for the host that exists:** **L1** keeps a substantial price here — this
  Mac mini pays a size-proportional unpack leg (measured 19.3 s on the 4.12 GB hiring image), so L1 is
  still the opening lever (+ the ISOLATION assert landed *with* it). **L2 changes character:** the
  derived `max_parallel_ui_lanes` is **2** on this host where it is **1** on billion, so the two compile
  legs *can* run concurrently here → **L3** → the small levers as the arithmetic demands.

## Carried known-context

`TOK-01` § Known-context #1–#9 (the Phase 0b YELLOW residuals + the recon findings). Not deferrals.

## M257: Final Review (close, 2026-08-12)

Findings from the close's scope / code-quality / adversarial / docs / tests / decision-triage passes.
**Every one addressed** — no partial fixes, no sign-off deferrals.

### Scope
- [x] The §8.5 corpus retraction (`D-v28-10`) — the milestone's one remaining declared deliverable, and it
      had not been started. Landed with achieved numbers, once (`D121`).
- [x] The achieved-numbers rewrite of `frontend-tier.md` **and** `build-budget.md`, both host-stated.
- [x] The grep gate `D-v28-10` promised and that did not exist — `demo_knob_guard` was named as the fence
      and structurally cannot see prose numbers.
- [x] The four M255-inherited items, unlanded across nine iters: 1 landed now, 3 re-fated with measured
      reasons (see `decisions.md` § deferral re-audit).
- [x] M257x's carry-forward had never reached M258's `overview.md` — the `BIND_HOST` failure repeated.

### Code Quality
- [x] [must-fix] `effective_disk_avail_gib` graded the **host** filesystem for a `docker-desktop-vm`
      profile when the VM probe failed — clause 3 holding the opposite policy to clause 1 on the identical
      question, invisible to clause zero because the fallback always returns a float.
- [x] [must-fix] `isolation_ok` computed in `build_report` and read by nothing — `red`, `gateable` and
      `print_report` all blind to the milestone's own new gate clause.
- [x] [must-fix] the RED reason string named six causes and not isolation.
- [x] [should-fix] `_image_bundle_pks` never got the fingerprint skip `_stack_minted_pk` got — hardening
      applied to one of two call sites of the same regex.
- [x] [should-fix] the named-file loop in `_stack_minted_pk` lacked the `except OSError` its own glob loop
      had — a `.env.demo` that is a directory takes the rep down mid-campaign.
- [x] [should-fix] `_git` fail-open recorded `rext_dirty: False` for "git could not run" — the provenance
      field iter-08's whole baseline rests on.
- [x] [should-fix] `is_ephemeral_clone` matched an **absolute** path substring; any checkout under an
      ancestor named `stacks` published a census of zero.
- [x] [should-fix] "it is a shared predicate now" was declared, not performed — three copies of one rule.
      All three now import it.
- [x] [should-fix] six stale "~3.7 GB / ~3 min" claims still driving operator text in `up-injected.sh`,
      two of which had been retracted in the same file.
- [x] [should-fix] the `DEMO_DISK_MIN_GIB` comment asserted the operator floor and the gate "cannot
      drift" — they already had (25 vs the gate host's 22). Corrected + routed.
- [x] [nice-to-have] the module docstring never mentioned ISOLATION; `_PORT_RE` silently truncated a
      6-digit run; `CLONE_DIR_PARTS` was a one-element tuple; `next-web.dockerignore` cited `Dockerfile.dev`.
- [x] [nice-to-have] `macmini.json` did not say its lane-peak / image-size fields describe the pre-L1 build.

### Adversarial review
- [x] *"Does a clause-3 fix retroactively fail the gate?"* — asked before fixing, answered by measurement
      (the reps recorded ~65 GiB, the VM's own figure), then re-verified after.
- [x] *"Does making an absent isolation block non-gateable retro-fail iter-08's baseline or billion's?"* —
      no: iter-08 was already RED on headroom, and billion's M255 campaign is already non-gateable on
      identity. Both checked rather than argued.
- [x] *"Is the new fence satisfiable by a corpus that quotes its own retractions?"* — it is not, absence-
      based; redesigned positionally before shipping.

### Documentation
- [x] 7 corpus files rewritten/swept; 7 citations re-anchored (one had slid onto a **different** construct
      — the guard's known blind spot, hit live); 13 knob anchors regenerated via `--fix`.
- [x] Seven sub-phase figures in the first draft were **fabricated** and caught against the raw
      `campaign.json` before commit.

### Tests & Benchmarks
- [x] +25 tests across `test_buildbench.py`, `test_isolation_assert_m257.py`, `test_union_apply_guard.py`
      and the net-new `test_section85_retraction_fence_m257.py`, each with its control.
- [x] The M255 mutation battery's `falsy-zero-disk-probe` mutant stopped matching its subject and the
      battery correctly went RED — re-pinned, plus a net-new mutant for the clause-3 arm.
- [x] `_ledger`, the fixture claiming to be "the shape `run_campaign` really writes", carried no isolation
      block. Repaired the fixture, not the assertions it was hiding.

### Decision Triage
- [x] D-M257-C1 (claim-keyed fence) → blended into `build-budget.md` + the fence's own docstring.
- [x] D-M257-C2 (the retraction that isn't) → blended into `frontend-tier.md` §12 GB prerequisite.
- [x] D-M257-C3/C4 (re-grade before believing; repair the fixture) → archive, maintainer-only.

## M257: Gate Outcome Ledger (Phase 9-iter) — `closed-on-gate`

### Gate

| | |
|---|---|
| **Target** | cold `demo-down --purge` + `demo-up` reaching `autoverify green:true / 0 warnings` at **p50 ≤ 360 s across 3 consecutive cycles on `macmini`**, **0 platform-repo edits**, **all 7 demopatch guards (G1–G7)**, and **two falsifiable asserts** (HEADROOM, ISOLATION) that FAIL the gate when tripped. Stretch **≤ 300 s** |
| **Achieved** | **p50 286.99 s** (n=3 — 286.99 / 303.44 / 280.99; min 280.99 / max 303.44) |
| **Distance** | **−73.01 s inside the gate, and −13.01 s inside the STRETCH** |
| **Status** | ✅ **`closed-on-gate`** — met on the gate's own terms, not by ruling |

**Every clause, graded:** `green:true / 0 warnings` **3/3** · **HEADROOM OK 3/3** (peak load1 5.47 / 6.16 /
8.72 against a limit of 10) · **ISOLATION OK 3/3** (8 images per rep, 0 failures, a real
`own_pk_fingerprint`, empty `foreign_pks` / `foreign_origins`) · host identity `match` **×3** ·
`rc_up=0` **3/3** · phase table complete **3/3** · **0 platform-repo edits** · **0 refused demo-patches**
(against iter-07's 17-of-17 refusal, the failure this campaign shape was built to exclude) ·
campaign `ok: true` / **`gateable: true`**.

**Both falsifiable clauses actually fired during the milestone, which is what makes the pass mean
something.** HEADROOM failed 2 of 3 reps on the iter-08 baseline campaign (peak load1 19.48 / 14.52) and
took that campaign RED by contract; ISOLATION was proven able to fail by 19 unit controls + 3 live ones,
after two fail-opens that only a live artefact could reveal.

**Re-verified twice after the fact, not assumed.** The final harden re-graded the three reps under
post-harden code (`rep_is_ok` True ×3); this close re-graded them again under the three fail-open fixes it
landed in the instrument — `rep_is_ok` True ×3, `ok`/`gateable` true, p50 **286.99 unchanged**, identity
`match` ×3. The headline verifies exactly against the raw per-rep ledgers.

### Iter ledger

**9 iters — 7 tiks + 2 toks**, all closed, all with a complete `iter-NN/{overview,progress,decisions}.md`.
No orphan iters; no orphan commits (15 on the milestone branch before the close).

| | |
|---|---|
| tiks | 02, 03, 04, 06, 07, 08, 09 |
| toks | **01** (bootstrap — `TOK-01`, *instrument before baseline, baseline before levers*) · **05** (triggered, after a 3-tik no-progress streak — `TOK-02`, `retry-with-evidence`) |
| levers landed | **1 of 8 priced (L1)** — and it cleared the gate alone, worth **−141.63 s** on the UI tier |
| metric-moving iters | **2 of 9** (08 established the baseline, 09 moved it). The other seven closed *"metric delta 0, by design"* — accurately, and three of them **could not** have moved it |

**The milestone's defining fact, and it is a process finding as much as a technical one:** iters 02–04 each
closed with zero metric delta, and `TOK-02` diagnosed why — the gate's **subject** (`odysseus`) was retired
one day after `TOK-01` named it, so **no measurement taken anywhere could have satisfied the gate**. A gate
that names a dead thing does not fail; it **abstains**, and abstention is invisible. iter-05 re-pointed it
at `macmini` with every target surviving verbatim — a stale-reference repair, not a relaxation, and the
`≤ 360 s` cut turned out to be *more* reachable on the new host, not less.

### Routes carried forward

**None required by the gate** — this is a `closed-on-gate` close, so **no `carry-forward.md` is produced.**
Optimisation beyond a met gate and instrument hygiene are routed Fate 3 to **M258** and recorded *at the
destination* in its `overview.md`; the per-item audit is in `decisions.md` § deferral re-audit. Headline:
**`LEVER-M257-L5-setdress`** — with the UI tier collapsed, `set_dress` is now the **largest single phase at
82.04 s = 28.6 %**, where the plan priced L5 at ~30–50 s and ranked it **fifth**.

### Dropped

`INVESTIGATE-M257-load1-48` — un-reproducible; its host is retired. Narrowed twice before dropping, with
the surviving hypothesis recorded in code and its companion suspicion answered.

### Protocol evolution

- **Check a gate for GRADEABILITY before checking it for satisfaction** (iter-05). Three iters reported an
  accurate "no delta" and none could say "and no delta is achievable".
- **A measurement on a contended box, LABELLED, beats waiting for a quiet one** (`TOK-02` step 3) — and
  labelling waives nothing: the HEADROOM clause still took that campaign RED.
- **Land each falsifiable assert WITH the lever that can trip it** (`TOK-01`) — ISOLATION shipped with L1,
  and the harden found the first evidence the pairing pays.
- **Predict, then measure, and keep the two apart** — L1 predicted ~132.6 s, realized 141.63 s;
  `backend_builds`' −12.37 s was booked as variance rather than credited to the lever.
