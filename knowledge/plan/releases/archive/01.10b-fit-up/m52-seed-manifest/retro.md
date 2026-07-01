# M52 — retro

## Summary
M52 consolidated the scattered seed+generation direction into ONE checked-in, auditable
`seed-generation-manifest.yaml` (all 3 orgs' population + the mother prompt + batch config + snapshot
sources; cache/generated data excluded), projected + honesty-gated from the canonical presets and served
by the cockpit [Download]. 4 sections landed cleanly (S1 `go:embed` prompt extraction byte-identical → the
M45 cache preserved; S2 the NEW `manifest` pkg + `--manifest-export` verb; S3 cockpit repoint; S4 the NEW
`corpus/ops/demo/seed-manifest-spec.md`). Closed `section`-shaped with all 5 `In:` items delivered as
Fate-1; the one carry (up-injected.sh end-to-end glue) is a confirmed Fate-2 → M53.

## Incidents This Cycle
None (no P2 flakes, no regressions). The build was already correct at close entry — the 3 harden passes
found no production-code bugs, and the close review's 12 findings were all quality/robustness deepening
(no correctness defect that had shipped broken). Flake gate 5/5 Go + 5/5 Python, clean.

## What Went Well
- **The honesty-gated projection (D2) paid off at close.** Because the manifest is projected from the
  canonical presets and pinned by `TestManifest_CanonicalFileMatchesProjection`, the close review could
  reason about drift precisely — and the F1 dedup (one canonical `blueprint` helper) closed the ONE gap
  the gate structurally couldn't catch (projection-vs-seeder drift), making the single-source property total.
- **The adversarial pass earned its keep.** 3 of its 4 findings became real robustness fixes on code that
  had just shipped: F4 (orphan gen-id was SILENTLY dropping generation intent — the exact failure the
  "single auditable file" goal exists to prevent, now WARNED), F3 (the cache-key golden was blind to the
  `{{else}}(none)` branch a hero-less fill batch renders), F5 (an empty-but-readable manifest served a
  hollow download). None were caught by the population-axis-only tests before.
- **Byte-identical extraction held.** The `go:embed` move preserved every M45 cache key with zero churn —
  the two cache-key goldens now fence both template branches against a future silent re-key.

## What Didn't
- **The primary cache-key golden shipped single-branch.** The build+harden goldens pinned only the
  reserved-names-PRESENT render, leaving the `(none)` branch (the real hero-less-fill case) unfenced until
  the close adversarial pass. A branchy template needs a golden per branch — a lesson for future prompt work.
- **A stale doc block survived S4's reconciliation.** cockpit-spec.md's Served-endpoints table was updated
  but the earlier "For PMs" prose (line 63) still claimed the download saves the JSON menu — the S4 pass
  reconciled the table it edited, not a sibling paragraph. Section-scoped reconciliation misses cross-block
  claims; a whole-file read catches them.

## Carried Forward
- **up-injected.sh end-to-end glue → M53** (Fate-2, DEF-M52-01): the manifest-export + `--seed-manifest`
  wiring has no shell unit harness; M53's cold `/demo-up` exercises it (its `In:` asserts the [Download]
  returns the complete inlined manifest). Confirmed, no plan edit.
- **Consumption-clone re-pin to `fit-up-m52` + `.agentspace/rext.tag` bump → M53** (push-gated KEEP).

## Metrics Delta
(from `metrics.json`) rext stack-seeding Go **749 → 786** (+37; NEW `manifest` pkg 100% stmt) · demo-stack
Python **299 → 313** (+14; cockpit `--seed-manifest` endpoint + fallback) · TS e2e unit 33 (unchanged) ·
flake **0** (5/5 Go `-shuffle` + 5/5 Python). No new dependency in either stack (`go:embed` is stdlib).
rext tag **`fit-up-m52`** @ `36d7430`. Close fixes: rext `c62623d`/`99dc098`/`36d7430` + rosetta `e523d93`.
