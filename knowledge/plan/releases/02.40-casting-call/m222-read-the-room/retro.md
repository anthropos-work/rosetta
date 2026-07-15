# M222 "read the room" — Retro

_Closed 2026-07-15. Section milestone, shipped clean (0 incidents, 0 regressions). v2.4 "casting call", milestone 1 of 5._

## Summary

The release's **hard go/no-go barrier**, cleared: **GO**. M222 authored the missing hiring READ-model doc
(`corpus/services/hiring.md`) and **proved by rendering** that the recruiter candidate-comparison surface
(`/enterprise/activity-dashboard`) lives in the **dockerized `apps/web`**, renders from seedable data, and survives
the `is_hiring` flip — refuting BA-3 (the "it's `apps/hiring`-only, so showing it needs a platform edit" escalation
trigger) and retiring R2. It landed the `is_hiring` gate in rext (`stack-seeding`) — the load-bearing one-value
thread from a hardcoded `false` to `st.IsHiringOrg()`, gated behind both an explicit flag and a `narrative: hiring`
discriminator — with a RED-provable test. No seeder, no funnel: foundation only, correctly deferred to M223.

## Incidents This Cycle

None. No P2 flakes, no regressions, no false-greens. The gate test was proven RED-provable at close (reverting
`org.go` to hardcoded `false` fails it), so it is not a string-fence.

## What Went Well

- **Prove-by-rendering beat inference.** The whole release hinged on whether the comparison surface was demo-servable;
  M222 traced the read-path end-to-end on the live `billion` substrate rather than reasoning from route files. That
  trace surfaced the **mirror-table trap** (D2) — the score is `public.local_jobsimulation_sessions.score`, NOT
  `jobsimulation.sessions.score` — which is the single most valuable artifact for M223/M224: seeding the obvious
  table would have rendered an EMPTY scoreboard (the M219 render-gate-bypasses-the-seed class), and the whole
  release would have chased a blank page.
- **The doc was traced, not asserted.** Every file:line claim in `hiring.md` re-verified GREEN against the READ-ONLY
  platform clones at close (resolver at `:1088`, Casbin gate at `:1089`, delegation at `:1134`, the intelligence.go
  read-path, the Ent Float32 score, both `isEnterprise` definitions). D17-clean: no status artifact outliving the
  code it describes.
- **The gate is honest.** `manifest.Org.IsHiring` uses `omitempty` so existing presets stay byte-identical — and
  that honesty-gate property is defended by a serialization-level test that counts `is_hiring:` key occurrences, not
  just the struct field.

## What Didn't

- Nothing material. Two tiny close-time items: a doc-fidelity tighten (the `useGetClerkOrganization` quote dropped
  the `organization?.` optional-chaining the source has) and the required Phase-2c adversarial record. Both trivial.

## Carried Forward

- **The HiringSeeder + candidate-assessment funnel → M223** (Fate-2, owned by M223 `overview.md`). M222 landed the
  gate only, by design.
- **Clerkenstein `publicMetadata.isHiring` dual-write wiring → M224** (Fate-2; D3 records the dual-write contract).
- **`directus.job_position` replay DROPPED from M223 Scope.In** (Fate-3, applied): the captured snapshot has 0
  `job_position` rows and the scoreboard never reads the entity — the 5 "positions" ARE 5 real captured
  `SIMULATION_TYPE_HIRING` sims. Recorded in D4 + M223 `overview.md`.

## Metrics Delta

- **Go test funcs (all 6 modules):** 1831 → **1838** (+7 — the is_hiring gate: 4 blueprint + 1 manifest + 2 seeders).
- **stack-seeding module suite:** 965 pass / 0 fail / 0 skip; `go vet` + `gofmt` + `-race` clean.
- **Flake:** 0 (5/5 random-order). **Platform-repo edits:** 0. **Net-new deps:** 0. **Deferral audit:** GREEN.
- Python (1341) and TS e2e (151) untouched — carried from v2.3 close.
