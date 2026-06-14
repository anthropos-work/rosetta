**Type:** tik (sixth tik under TOK-01) — code-ification, apply side (multi-step planned shape).

# M21 iter-07 — progress

## Work done (Go, rosetta-extensions authoring copy)
1. **Capture wiring** — `capture.Surface.CapturesStructure` (directus opts in) + an optional, type-asserted
   `StructureCapturer` interface (existing capture fakes untouched). `capture.Run` captures the structure after the row
   firewall passes, stashes it on the same all-or-nothing path, stores `_structure.sql`, and sets `manifest.Structure`.
   The directus `captureAdapter` implements it via `directus.CaptureStructure`; `captureConn` gains `QueryRowString`.
2. **Apply wiring** — `replayCmd`, on cache MISS, calls `tryAutoProvision`: finds the surface's structure-bearing
   cached snapshot, checksum-verifies the structure payload, applies it via the new `pg.ExecScript` (simple-protocol
   multi-statement), re-probes the digest, re-resolves → hit → replays rows. Redefined exit semantics: empty schema →
   4, gap-without-structure → 5, gap-with-structure → 0. (`autoprovision.go`, testable: `provisionConn` interface.)
3. **Recipe truth** — `provision.go` ProvisionPlan content-schema step no longer "NOT YET AUTOMATED"; the replay step
   advertises auto-provision. Test updated (`TestProvisionPlan_ContentSchemaIsAutoProvisioned`).
4. **Robustness fix (M21-D11)** — `structureSeqSQL` re-keyed from ownership to **default-reference** (`pg_attrdef` →
   sequence), so it captures exactly the sequences the CREATE TABLEs reference regardless of source ownership.
   Live-validated: 8 on prod (owned) AND 8 on the standalone-sequence integration source.
5. **Tests** — new `capture_structure_test` (4) + `autoprovision_test` (5, incl. corrupt-structure + apply-error +
   no-structure no-op + other-surface scoping) + the recipe-truth test. Full 12-package suite green.

## Live two-harness integration test (real prod content, 10128 rows) — ALL GREEN
- **Capture** source→test-store: `stacksnap capture --surface directus` → exit 0, 9 tables / 10128 rows + structure
  artifact `_structure.sql` (60 stmts = 8 seq + 26 tables + 26 PKs), `manifest.structure` set.
- **Replay auto-provision** test-store→fresh bootstrapped target: "provisioned directus structure into dev-6
  (b4cb55bc → 6cd35278); loading content…" → exit 0, 26 user tables, digest **converged to 6cd35278**, simulations=304,
  roles=953.
- **Idempotent re-run** → exit 0, simulations=304 (not doubled; no re-provision since digest matches).
- **Exit-code matrix:** empty schema → **4**; gap + structure-less store (the real pre-M21 cache) → **5**;
  gap + structure-bearing store → **0**.
- Evidence: `/tmp/m21iter07-evidence.md` (transient); full detail in this file.

## Adversarial review (ultracode)
Ran the `m21-iter07-review` workflow (4 dimension reviewers → adversarial verify; 12 agents, 13 raw findings).
- **2 confirmed should-fix** converged on ONE real regression — **the gap-guard** (F1 + AUTOPROVISION-NO-GAP-GUARD):
  auto-provision fired on ANY directus cache-miss, so a diverged/schema-skewed target (has user tables, digest ≠
  captured) collided on `CREATE TABLE` → raw exit 1, **regressing** the pre-M21 clean exit-5 divergence path. **FIXED**
  → `tryAutoProvision` now gates on the GAP precondition (0 user collections); a diverged target → no-op → exit 5
  (M21-D12). Live-confirmed: diverged target now → exit 5 with the divergence message, not exit 1.
- **3 correct dismissals** (good adversarial rejections): ExecScript IS atomic (simple-protocol implicit transaction);
  directus cannot produce generated/view/partitioned collections.
- **Nits fixed inline:** premature "loading content…" log; scoped the structureSeqSQL identity-column comment.
- **Routed to harden (`/developer-kit:harden-mstone-iters`):** AP-1 (a hermetic replayCmd-wiring test needs a conn
  injection seam — the wiring is integration-tested live meanwhile); AP-2 (multi-snapshot lexical-first determinism —
  stale-but-valid); AP-3 (exit-4-boundary regression guard); the user-defined-type / firewall-ordering direct-test nits.

## Close — 2026-06-12

**Outcome:** `stacksnap` now CAPTURES the directus schema structure and AUTO-PROVISIONS it on replay (the gate's
"stacksnap applies the captured structure" clause, schema half — stages 3-4 automated). Validated live end-to-end +
the full exit matrix (empty→4, gap+no-structure→5, diverged→5, gap+structure→0) + idempotency. An adversarial review
caught + I fixed a divergence-path regression (the gap-guard, M21-D12).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (schema apply automated; the SERVE rows — registration + permissions — are iter-08 → flips gate met)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (tik) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (1 tik this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** M21-D11 (default-reference seq capture), M21-D12 (auto-provision gap-guard); iter-local in `decisions.md`.
**Side-deliverables:** none separable (all in-scope code-ification).
**Routes carried forward (→ iter-08 under TOK-01):**
  - `STRUCT-M21-iter08-serve` — capture + apply the directus_collections registration + the public read permissions
    (the serve half, M21-D9) + the firewall structural-metadata admissibility class → a stacksnap-provisioned stack
    serves anonymously with no hand SQL → **flips the gate met**.
  - To `/developer-kit:harden-mstone-iters`: AP-1 (replayCmd-wiring hermetic test + conn seam), AP-2 (multi-snapshot
    tie-break), AP-3 (exit-4-boundary guard), UDT/identity guards, firewall-ordering direct test.
  - Carried: `directus_files` ref capture; M23 referential closure.
**Lessons:** an adversarial review workflow on a non-trivial wiring diff earned its keep — it caught a real
divergence-path regression (gap-guard) that the green unit + integration tests missed (neither exercised apply-onto-a-
non-empty-schema). For any "auto-provision / auto-heal on cache-miss" code, gate the mutation on the precondition the
comments assume (here: target is a true gap), or a skewed input silently degrades a clean error into a raw failure.
