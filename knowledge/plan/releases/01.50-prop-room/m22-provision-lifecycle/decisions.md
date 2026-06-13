# M22 — Decisions

_Implementation decisions with rationale, numbered `M22-D1`, `M22-D2`, … . Empty at scaffold; filled during build._

## KB items (from the Phase 0b fidelity audit, 2026-06-13 — GREEN)

- **KB-1** — `directus-local.md:40` historical aside cites `provision.go:108` ("previously dead-ended at the print-only placeholder"); post-M21 that line holds the content-schema step Detail. The "was" framing is accurate but the numeric anchor drifted. Fix in M22 Phase 5 (the lifecycle-section rewrite touches this exact area) — prefer a prose/symbol reference over the numeric anchor. Incidental, non-blocking. **RESOLVED (§6, `0ab823a`):** the anchor now reads "the print-only placeholder in `provision.go`'s `ProvisionPlan`" (prose/symbol, no line number).

## Decisions

- **M22-D1 — collateral stale-claim retirement (Fate 1).** §6 surfaced that M22's executed provisioning makes the "print-only for BOTH stack types" / "no stack type has an automated per-stack Directus" claims in `snapshot-spec.md`, `safety.md`, and `demo/README.md` false. These are load-bearing claims the milestone invalidates, so they were fixed in the same §6 docs commit (Fate 1 — land now), reframed as: **executed on a `--local-content` stack** (demo default / dev opt-in), **print-only prod-read** otherwise; the **M23 cutover** (re-point `DIRECTUS_BASE_ADDR`) stays honestly future.
- **M22-D2 — studio-desk verify-port test repair (Fate 1).** §4's demo-stack suite run surfaced that `test_verify.py`'s `TestFrontendTierRegistration` still asserted studio-desk port **9100** while `services.sh` had already moved to the single user-facing port **9000** (committed earlier in `6c5f516` "single-port verify", which updated `test_injection.py` but missed `test_verify.py`). A stale, failing assertion; corrected to 9000 as part of §4. Also fixed two `test_frontend_build.py` chain-contract tests that ran the §2 set-dress block under `set -u` without binding `NO_LOCAL_CONTENT`.

## Adversarial review (close Phase 2c, 2026-06-13)

The close adversarial pass simulated an external reviewer against the milestone's safety-critical module
(`dev-setdress.sh`'s executed-provision flow + the `stack-verify` directus probes). Each scenario below is the
*failure mode considered* — **every one was already test-pinned** at build/harden time; the pass surfaced **zero
new findings** and required zero fixes. Recorded here so future reviewers see what was examined.

- **Write-before-firewall ordering.** Could a provision write (CREATE SCHEMA / bootstrap / replay) land before
  the `provision-plan --check-env` gate fires? No — the gate `die`s at `snapshot_step` lines 221-223 *before*
  `provision_directus_step` runs (line 228); there is no write path between them. Pinned by
  `test_executed_firewall_gate_is_load_bearing`, `test_prod_directus_env_aborts_before_replay`,
  `test_provision_recipe_failure_aborts_before_replay`.
- **`set -u` trip on `DIRECTUS_PROVISIONED`.** With local content OFF (the var's branch never sets it true),
  does the later `[ "$DIRECTUS_PROVISIONED" = 1 ]` read trip `set -u`? No — it's initialized to `0`
  unconditionally before the branch. Pinned by `test_no_snapshot_with_local_content_skips_provision_no_setu_trip`.
- **Degrade path writing to prod.** Could the non-fatal degrade (a bootstrap/schema failure) ever write to prod?
  No — the firewall validated the env prod-safe *before* any provision, so a degrade only ever falls back toward
  the safe prod-*read* path. Pinned by `test_bootstrap_failure_degrades_to_prod_read_nonfatal`,
  `test_create_schema_failure_degrades_nonfatal`, `test_capture_never_runs_even_on_the_cache_miss_degraded_path`.
- **False-safe static gate.** If `--check-env` returned 0 on a mis-wired override that actually points at prod,
  is there a runtime backstop? Yes — the `autoverify` no-prod-read env assert reads the *container's* live
  `DB_CONNECTION_STRING` and warns if it resolves to a prod host, and fails-open-safe (skip, not false-leak) when
  the DSN is unreadable. Pinned by `test_no_prod_read_assert_catches_a_prod_dsn`,
  `test_no_prod_read_assert_skips_when_dsn_unreadable`.
- **Half-bootstrap re-provision.** A crash leaving some `directus_*` tables but no registry — does a re-run skip
  onto the broken schema? No — the guard probes the `directus_collections` *sentinel* (complete-bootstrap
  marker), not a blanket count, so a half-bootstrap re-bootstraps to converge. Pinned by
  `test_half_bootstrap_rebootstraps_to_converge`, `test_bootstrap_probes_the_directus_collections_sentinel`.
- **Silent serve-nothing.** A Directus UP (`/server/health` 200) but serving an empty catalog (registry never
  populated) — caught? Yes — the "registered collections > 0" cheap-win (the casbin-assert analog), robust to a
  non-numeric psql result. Pinned by `test_registered_collections_zero_is_caught`,
  `test_collections_nonnumeric_is_caught_not_crashed`.
- **N=0 base-port write.** Could `--local-content` provision the developer's primary box (N=0) without the guard?
  No — N=0 is refused before any provision unless `--force`. Pinned by
  `test_n0_local_content_refused_without_force_before_any_provision`.
