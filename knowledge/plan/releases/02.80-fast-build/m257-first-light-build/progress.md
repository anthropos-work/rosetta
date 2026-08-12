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
