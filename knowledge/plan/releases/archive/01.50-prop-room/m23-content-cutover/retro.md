# M23 — Retro

## Summary
M23 cut the **data plane** over to the per-stack Directus: `cms`'s `DIRECTUS_BASE_ADDR` now re-points to the
in-network `http://directus:8055` (the prod token stripped on the way) while the **asset plane stays on prod**
(`DIRECTUS_PUBLIC_BASE_ADDR` unchanged) so browser images stay real — the data-plane-local / asset-plane-prod
split. studio-desk gets the per-stack instance + a locally-minted static admin token (stamped via Directus's
`ADMIN_TOKEN` bootstrap env, gated by a new `ValidateProvisionable` present-token check layered on the prod-safety
`Validate`). The `directus_files` ref capture (DEF-M21-03/Fate-3) landed as a new REFERENCED-SUBSET firewall
admissibility kind (reverse-reference closure, admit-iff `Filter==ReferencedFilesFilter`) with a `ClearByDelete`
DELETE-before-TRUNCATE for the external `directus_settings` FK. Referential closure is now **measured** by a
cross-surface gene — and the load-bearing data-design finding was that **full-taxonomy capture
(`organization_id IS NULL`) was already the state**, so closure is maximal by construction; the only residual is a
content ref pointing at a non-public node, which the firewall must never capture. 6 sections, all Fate-1.

## Incidents This Cycle
- **0 code incidents / 0 regressions / 0 flakes.** The 2-pass hardening sweep surfaced 0 production bugs (every
  gap was a test gap); the close adversarial pass (4 scenarios) found 0 new findings (all already test-pinned);
  the flake gate ran 5/5 clean on all four touched suites.
- **1 prod DATA-quality residual surfaced (not an M23 code incident).** The new closure gene found exactly 1
  dangling content→taxonomy ref in prod: `K-AIFUNX-E658`, referenced by 2 public published sims but existing only
  as a customer-scoped skill. It is **uncloseable by tooling** (capturing the customer node = a tenant-firewall
  breach; editing prod = forbidden). The gene's job was to make this VISIBLE rather than silently ship an empty
  picker — it did. Fated an operator-owned prod data correction (KNOWN-ISSUE), outside tooling scope.

## What Went Well
- **The closure problem dissolved instead of needing solving.** The roadmap budgeted M9b-style decision time for
  "closure-at-capture vs full-taxonomy capture". On inspection, full-taxonomy capture was already the captured
  state — so closure is maximal by construction and the only work was to MEASURE it (the cross-surface gene),
  not to subset-and-close. The simple fallback the corpus already named turned out to be the shipped reality.
- **The cms-only cutover model kept the env-sprawl to one service.** Because `cms` is the only platform service
  that talks to Directus directly (jobsimulation reads it via cms RPC; next-web via the router), re-pointing one
  `DIRECTUS_BASE_ADDR` cuts the whole data plane over — no per-service env explosion. The 4 service docs now
  state this dependency truth explicitly.
- **2 inherited M21 deferrals RESOLVED in-milestone.** DEF-M21-03 (directus_files) + DEF-M21-04 (referential
  closure) both landed Fate-1 — the deferral ledger shrank rather than grew. The strongest possible close-audit
  outcome.
- **Adversarial surface pre-pinned.** Every scenario the close pass could construct (a filter slipping the
  admit-iff gate; a valid token masking a prod BaseAddr; a two-dest scan on empty/error; the `directus_settings`
  FK breaking TRUNCATE) was already covered by a named test at build/harden time.

## What Didn't
- **One stale future-tense doc claim slipped past §6.** `snapshot-spec.md`'s M13 section still said "The M23
  *cutover* … remains future" after M23 landed the cutover — caught + fixed Fate-1 at close. The §6 docs commit
  updated the known-state + media + fidelity-gate sections but missed the older M13-section narrative. A reminder
  that "retire the old claim" has to sweep *every* section that references the changed mechanism, not just the
  current-state block.
- **The decision-backref tags were absent from the corpus.** M23-D1..D5 content was blended into the docs during
  §6, but the `(#M{N}-D{K})` traceback tags the corpus convention uses (M10-D4, M13-D2, M18-D1…) weren't added
  until close. Minor, but the tags are the convention — they should land with the blend.

## Carried Forward
- **DEF-M21-02** (automated serve-live-integration harness) → **M25 field-bake** (the live observable-behavior
  gate; unchanged — M23's gene runs against the replayed pair but the automated harness is M25 scope).
- **DEF-M21-01** (replayCmd conn-seam) → tracked tooling-debt follow-up (untouched by M23; M23 only added a
  struct field to `replay.go`, not the seam refactor).
- **DEF-M10-01** (S3 blob bytes + cloud store) → backlog (unscheduled), re-affirmed refs-only by M23-D3.
- **K-AIFUNX-E658** (the prod data residual) → operator-owned prod data correction (re-tag or remove the bad
  skill ref on the 2 public sims) — outside tooling scope; the gene continues to measure + name it.

## Metrics Delta
- **Go test funcs:** 795 → **844** (M23-own **+33** across stack-snapshot 290→316 + stack-seeding 252→259; the
  remaining +16 is a counting-method reconciliation on untouched modules). Coverage gaps closed:
  `CrossSurfaceDangling` 0→100%, `ValidateProvisionable` 80→100%.
- **Python (touched suites):** stack-core 61→**69** (+8); stack-injection **110** (8 env-gated skip). The other
  3 suites untouched.
- **Flake:** 0 (5/5 — Go shuffled, Python sequential). **go vet / shellcheck / py_compile:** CLEAN.
- **Review:** 6 findings (0 scope / 0 code-quality / 0 adversarial-new / 1 docs Fate-1 / 0 tests / 5
  decision-triage). **Deferral audit:** GREEN (2 inherited RESOLVED, 0 repeat / 0 aged / 0 blockers).
- Full machine-readable record: [`metrics.json`](metrics.json).
