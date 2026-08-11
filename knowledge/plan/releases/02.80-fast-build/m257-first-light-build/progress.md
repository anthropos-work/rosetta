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

## Next-iter routing

> ⚠️ **The routing below is SUPERSEDED for the host and the blocker.** iter-04's Step 0 re-survey found
> that iter-03's `user-blocker` was already resolved by M257x (`c0e075e`), and that the host-class
> premise which paused this milestone is **refuted by measurement**. iter-05 then re-pointed the gate.
> Historical rows kept; read iter-05's close first.

- **iter-06 — `FIX-M257-load1-units-vm`, then `BASELINE-M257-macmini-n3`.** Under `TOK-02`. First fix the
  instrument: clause 1 grades **host** `load1` against a **VM-allocation** core count, so on this host it
  computes a limit of **6** where the correct one is **10** and **fails closed** — a refusal from it today
  cannot be told from a real one. Prove the new fail-closed arm RED with its precondition absent. Then the
  `n ≥ 3` cold campaign on a **free** demo slot, **taken CONTENDED and labelled so** (the box cannot be
  freed; `laptop.json` shows what waiting for quiet produces — a clause-1 refusal at load1 10.69 and no
  cycle number at all). Every rep carries its load1; a refusal is reported as a **result**. Heartbeat before
  touching the user's `demo-2` or the dev stack. Then levers, largest-measured-second first.
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
