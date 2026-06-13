# M22 — Decisions

_Implementation decisions with rationale, numbered `M22-D1`, `M22-D2`, … . Empty at scaffold; filled during build._

## KB items (from the Phase 0b fidelity audit, 2026-06-13 — GREEN)

- **KB-1** — `directus-local.md:40` historical aside cites `provision.go:108` ("previously dead-ended at the print-only placeholder"); post-M21 that line holds the content-schema step Detail. The "was" framing is accurate but the numeric anchor drifted. Fix in M22 Phase 5 (the lifecycle-section rewrite touches this exact area) — prefer a prose/symbol reference over the numeric anchor. Incidental, non-blocking. **RESOLVED (§6, `0ab823a`):** the anchor now reads "the print-only placeholder in `provision.go`'s `ProvisionPlan`" (prose/symbol, no line number).

## Decisions

- **M22-D1 — collateral stale-claim retirement (Fate 1).** §6 surfaced that M22's executed provisioning makes the "print-only for BOTH stack types" / "no stack type has an automated per-stack Directus" claims in `snapshot-spec.md`, `safety.md`, and `demo/README.md` false. These are load-bearing claims the milestone invalidates, so they were fixed in the same §6 docs commit (Fate 1 — land now), reframed as: **executed on a `--local-content` stack** (demo default / dev opt-in), **print-only prod-read** otherwise; the **M23 cutover** (re-point `DIRECTUS_BASE_ADDR`) stays honestly future.
- **M22-D2 — studio-desk verify-port test repair (Fate 1).** §4's demo-stack suite run surfaced that `test_verify.py`'s `TestFrontendTierRegistration` still asserted studio-desk port **9100** while `services.sh` had already moved to the single user-facing port **9000** (committed earlier in `6c5f516` "single-port verify", which updated `test_injection.py` but missed `test_verify.py`). A stale, failing assertion; corrected to 9000 as part of §4. Also fixed two `test_frontend_build.py` chain-contract tests that ran the §2 set-dress block under `set -u` without binding `NO_LOCAL_CONTENT`.
