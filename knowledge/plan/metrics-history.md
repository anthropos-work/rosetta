# Metrics History

One row per shipped release (newest first). Aggregated at `/developer-kit:close-release` Phase 4b from the release's per-milestone `metrics.json`. Full per-release detail lives in `releases/{archive/}{VV.VV}-{codename}/metrics.json`.

| Release | Codename | Shipped | Framework tests | Mirror tests | Gates | Flakes | Supply-chain |
|---|---|---|---|---|---|---|---|
| v1.0 | body double | 2026-06-03 | 43 test + 3 fuzz (stdlib-only) | 123 test + 6 fuzz / 8 pkg / 3 DNAs | Go 22/22 · JS 9/9 · express 9/9 · drift 9/9 (all 100/100; triple-clean 3/3) | 0 | GREEN (zero external modules) |

_Baseline release — no prior to regress against._
