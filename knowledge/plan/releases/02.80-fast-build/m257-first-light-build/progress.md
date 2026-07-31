# M257 — progress

## Running ledger

- iter-01 (tok/bootstrap): Phase 0b gate RED → 3 contract decisions surfaced + answered (`D120`/`D121`/`D122`) → RED cleared → re-audit **YELLOW (RED CLEARED)** → `TOK-01` authored. Side: the mirror fence parameterised by host (it then caught M257's own un-hosted gate line). — see `iter-01/progress.md`
- iter-02 (tik): odysseus **provisioned** (rc=0 bring-up, 16/16 up) + both inherited `autoverify` fixes landed with mutation-proven negative controls. **Metric delta 0, by design.** Surfaced **two blockers that make the gate currently unreachable** (B1 the dropped `local_*` mirrors → 6 seeders fail; B2 `app/studio` has no rext acquisition path → cold build dies, broken since 2026-07-27 and unnoticed because nobody had run a cold cycle) **+ a candidate re-scope signal** (peak `load1` **48.7** vs clause 1's limit of 6). — see `iter-02/progress.md`

- iter-03 (tik): **B1 + B2 both landed** — the two blockers that made READY unsatisfiable on every host — each with a mutation-proven fence. B1's blast radius was **34 sites / 20 files, not six**, and included **`autoverify`'s own hiring assert** reading the dropped table behind `|| echo 0`. **Metric delta 0, by design.** Surfaced a **strictly larger** blocker (`repos.yml` makes `app` the sole migration owner → a fresh stack never creates the `jobsimulation` schema that ~15 rext writes target) whose fix shape is a **user-level architectural choice**. `closed-fixed-partial`; exits `user-blocker`. — see `iter-03/progress.md`

## Next-iter routing

- **iter-02 (tik, under `TOK-01`)** — *Make odysseus a bench, and the instrument falsifiable.*
  (a) provision the host (PATH-fix the installed Go 1.26.5, install atlas, ssh-agent, snapshot cache;
  confirm the six rext Go modules build against 1.26.5); (b) land both inherited `autoverify` fixes —
  container-liveness scoped to the **real** hole (`fake-fapi`/`fake-bapi` have no `services.sh` row), and
  a fapi probe independent of the host TLS stack, *after* establishing empirically whether check (d) even
  warns on Linux; (c) **prove `autoverify` can go RED** via a deliberate negative control.
  **Expected metric lift: zero, by design** — grades on planned deliverables.
- **iter-03 (tik)** — the baseline campaign: **discarded warm-up cycle** (mandatory on this
  empty-cache host) then `n ≥ 3` → p50; check in `stack-core/hostprofiles/odysseus.json` with a
  **measured** `gated_baseline`; prove the HEADROOM assert falsifiable. **Decision point: is 360 s
  reachable, or is this the `re_scope_trigger`?**
- **iter-04+** — levers, largest-measured-second first, one per iter, re-measured at `n ≥ 3`:
  **L1** (+ the ISOLATION assert, landed *with* it since L1 changes exactly the layers that carry the
  baked key/origin) → **L3** → **L2** (+ union-apply + the hiring recruiter Playthrough re-verify, since
  `D-v28-7`'s "inert outlier" premise was false) → the small levers as the arithmetic demands.

## Carried known-context

`TOK-01` § Known-context #1–#9 (the Phase 0b YELLOW residuals + the recon findings). Not deferrals.
