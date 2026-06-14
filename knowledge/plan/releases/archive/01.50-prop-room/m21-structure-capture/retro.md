# M21 — Retro

_Structure capture — close the collection-schema gap. The 1st + **foundational** milestone of v1.5 "prop room", closed 2026-06-13 (closed-on-gate)._

## Summary

M21 closed the M10 **collection-schema gap** that made every stack read public content live from prod. The snapshot
now carries the content-model **structure** (the user-collection table DDL + **PRIMARY KEYs** + sequences + the
`directus_collections` registration + the public-policy `directus_permissions` read grant), captured atomically from
the same sanctioned `--dsn` as the rows. `stacksnap` **auto-provisions** a bootstrapped-gap stack before the row
replay, so the directus replay **exits 0** and a booted Directus **serves a captured public simulation anonymously
over HTTP** — the exit gate, MET by tooling at iter-08. The defining safety move was a new firewall
**structural-metadata admissibility class** (`AssertStructuralMetadata`): admit `directus_*` system tables as
"structure, not tenant data" iff they carry zero tenant-scope columns — **extend, never loosen**. 8 iters (1 bootstrap
tok + 7 tiks, 0 triggered toks — every tik advanced the 6-stage pipeline). The close was a clean shape: 5 findings
(1 scope · 1 should-fix code · 1 docs · 0 adversarial-new · 1 fate-routing), all addressed Fate-1.

## Incidents This Cycle

- **None at close.** Flake gate 5/5 deterministic (shuffled, full module); the two final harden passes fixed **0**
  production bugs — the M21 code held under error-path + edge-case enumeration (the firewall carve-out, the
  assert-then-read ordering, the auto-provision gap-guard all behaved as designed under the deepened tests). The
  close's cross-cutting code review found **0** must-fix and **1** should-fix (an error-label consistency nit), and
  the 5 simulated adversarial scenarios were **all already test-pinned** — no behavioral defect anywhere.
- **One P2-class regression was caught + fixed DURING the iters (not at close), worth recording:** iter-07's
  adversarial review caught that `tryAutoProvision` fired on ANY cache miss, not just a true bootstrap gap — a
  diverged target would hit the non-idempotent `CREATE TABLE`s → raw exit 1, regressing the clean exit-5 divergence
  path. Fixed inline (gate the apply on a zero-user-collection probe; #M21-D12). The general lesson — any
  auto-provision-on-cache-miss must gate its mutation on the precondition it assumes — is captured in the decision.

## What Went Well

- **The iterative shape was the right call — the empirical quirks were real.** The load-bearing finding was the
  **PRIMARY-KEY rule**: Directus silently 403s a PK-less collection *even for admin*, while the schema digest (over
  column *types*) still converges and the row COPY still works (#M21-D9). A fixed `In:` checklist would have shipped a
  column-only DDL that converged the digest, passed the replay, and then silently failed to serve. The gate — "serves
  anonymously" — caught it; doc-state wouldn't have. The staged-pipeline metric (furthest stage passing, 0–6) gave
  honest per-tik progress against a binary gate.
- **The firewall was extended, not loosened — and proven so.** The carve-out admits a previously-inadmissible class
  (`directus_*` registry tables) under a strict, explicit predicate (zero tenant-scope columns), runs **assert-then-
  read** (admissibility on the introspected column set BEFORE any row is materialized), and the final harden pass
  pinned that it composes with the row-side `AssertPlan`/`AssertCaptured` gates without widening or narrowing them.
  `directus_access` (has a `user` column) stays excluded; `directus_policies` stays bootstrap-provided.
- **Reuse-not-reinvent held.** Structure rode the existing machinery: an **additive** manifest `Structure` artifact
  (the `Predicate`-field precedent — no format bump), the serve-rows ride the existing `CapturesStructure` capture
  path + the iter-07 `ExecScript` apply path with **zero apply-side code change** (they ARE part of the structure
  SQL). The privilege-visibility alignment (#M21-D10 — capture intersects `pg_catalog` with the digest's
  `information_schema` view) generalizes to any digest-keyed capture against a least-privilege read role.
- **The operator-gated prod read was handled cleanly.** When self-contained structure sources proved exhausted
  (#M21-D6 — the cms structs are a lossy app-view, a hand-built reference can't invent prod types), the milestone
  surfaced the prod-read decision to the operator rather than guessing; the read stayed bounded / read-only /
  public-only behind the firewall.

## What Didn't

- **The committed net-new doc never got authored during the iters — caught at close.** `corpus/ops/directus-local.md`
  was M21's `Delivers →` net-new deliverable, but the 8 iters all banked their empirics into the milestone's
  `spec-notes.md`/`decisions.md` and never promoted them to the corpus. The close caught it (scope review + the
  per-unit-handbook contract) and landed it Fate-1 — but it's a reminder that an iterative milestone's *doc*
  deliverable needs the same per-iter attention as its code, or it pools in the planning dir until close. (The
  README index row was missing too — the exact recurring-miss class v1.5/M24's index-row guard targets.)
- **`directus_files` rode along unfinished across all 8 iters.** It was in M21's `In:` scope but never blocked the
  gate (anonymous serve of simulations works without it). Re-fated at close: it is asset-ref plumbing that belongs
  with M23's "keep the asset plane on prod" work, not the structure-serve gate — **Fate-3 annotated to M23**. The
  lesson: an `In:` item that the gate doesn't require can quietly never-land across an iterative milestone; the close
  deferral audit is where it gets honestly re-homed.

## Carried Forward

- **`directus_files` ref capture** → **M23** (Fate-3, applied to M23's `overview.md` `In:`/delivers). The asset-ref
  plumbing — wire the dead `media.go` filter/columns into `directus.Surface()` so content rows resolve their image
  UUIDs to the prod-public asset URLs. Refs only; blob BYTES stay backlog (DEF-M10-01).
- **The 20 dangling relations** → **M23** (Fate-2, already owned by M23's "Referential closure" In-item). No plan edit.
- **HARDEN-M21-AP1-replaycmd-conn-seam** → a tracked follow-up build iter. A hermetic `replayCmd` wiring test needs an
  injectable-connector refactor (>50 lines, the load-bearing replay path) — exceeds a harden/close inline boundary.
- **HARDEN-M21-serve-live-integration** → live-integration backlog. The serve-row render is hermetically unit-tested +
  hand-validated live; only the *automated* live harness is routed — becomes near-free once M22 boots a live Directus.
- **DEF-M10-01** (S3 media blob bytes + cloud `SnapshotStore`) → backlog, re-signed fresh at v1.5 design with its
  user-facing sting removed by the real-images-via-prod-links posture. Orthogonal to M21, not aged. Deferral audit
  GREEN (0 repeat / 0 chronic / 0 aged-out).

## Metrics Delta

- **Go `stack-snapshot`:** 231 → **290** test+fuzz funcs (+59 — the structure-capture core + auto-provision + the
  serve half + 21 final-harden tests). Whole-suite total across the 4 Go modules: **795** (+59 vs the v1.3b-close 736).
- **Coverage (M21-touched):** directus **100%** · firewall **100%** · manifest 98.4% · capture 98.9% · cmd/stacksnap
  80.1% · pg 47.0% (the cmd/stacksnap + pg residual is exclusively the live-DB connection wrappers + the
  replayCmd/replayAdapter real-conn pass-throughs = the route-forward AP-1 conn-seam item — reachable surface saturated).
- **Flake:** 0 (5/5 shuffled, full module). **Build + vet:** clean.
- **Ext repo:** tag `prop-room-m21` @ `835d940` (set by the orchestrator post-close).
- Full machine-readable breakdown: [`metrics.json`](metrics.json).
