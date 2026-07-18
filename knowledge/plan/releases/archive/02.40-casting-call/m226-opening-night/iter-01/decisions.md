# iter-01 — intra-iter decisions

| # | Decision | Rationale | Date |
|---|----------|-----------|------|
| D1 | billion ssh access is **`ssh devops@billion`**; stack root `~/panorama/stack-demo/`. | Tailscale-SSH `kirality` user mapping fails on billion ("failed to look up local user"); the M221/M217 records use `devops@billion`. Confirmed live in recon. | 2026-07-17 |
| D2 | Cut billion's rext over to the code-of-record tag **`casting-call-m225-harden`** (verify `sections`↔`harden` is test-only diff first). | REPROVE-at-final-code discipline (M221) — re-prove the pinned demo-patches at the *current* code. `-harden` is the code-of-record; `-sections` was the M225-proof/consumption tag. Hardening should be test-only (no runtime change) — confirm in tik-01 before pinning. | 2026-07-17 |
| D3 | The stale v2.3.2 `demo-1` must be **cleanly torn down** (serve reset + academy respawner reap) before the casting-call bring-up, not left alongside. | 7.3 GiB RAM can't hold two full stacks + two UI tiers; the stale demo also squats the base ports + serve config the fresh demo needs. Reuse N=1 for the fresh casting-call demo. | 2026-07-17 |
