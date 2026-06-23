# State

**Active version:** **v1.9 "storytelling" — IN DEVELOPMENT** (designed 2026-06-22 via
[`/developer-kit:design-roadmap`](roadmap.md); branch `release/01.90-storytelling`). The
**believable-demo-narrative release** — turn the placeholder seeder into a declarative **Stories & Heroes**
engine: each *story* is one org with a thriving/struggling/manager **hero** trio, seeded via the real
**verified-skill chain** (the 7-table jobsim→`user_skills`→`user_skill_evidences` fan-out) so the **skill
profile** + the org **Workforce dashboard** tell one coherent story — plus a standalone **presenter cockpit**
(log in as a hero + jump to the right screen). 5 `section` milestones **M34→M38** across `rosetta-extensions`
(`stack-seeding`/`clerkenstein`/`demo-stack`) + the rosetta corpus doc-half. **Tooling + docs only — zero
platform-repo edits.** **M34 ✅ + M35 ✅ shipped 2026-06-23** (the verified-skill spine + the multi-org Stories
engine); M36 next.

**Active milestone:** **M36 — Dashboard surfaces (Must #2)** — **NEXT (not started).** The org
**Workforce-Intelligence dashboard** renders believably for a seeded story: `membership_skills` (a believable
verification funnel), tags/teams, `organization_target_roles`+`user_target_roles` (gap + two-way mobility),
succession feeders sized to clear the coverage gate, feedback (~2:1 positive), the assignments status-mix fix,
and the org-scale claimed-vs-verified gap + growth-arc distributions. Depends on M35 (the multi-org Stories
engine, now shipped). **Parallel opportunity:** M37 (Clerkenstein multi-identity) may start alongside M36 — it
needs only M35's hero-identity list. The verified spec
([`.agentspace/seeding_gaps.md`](../../.agentspace/seeding_gaps.md)) remains the authoritative design for
M36–M38. The M34–M38 overviews live under [`releases/01.90-storytelling/`](releases/01.90-storytelling/).
**Next up:** build M36 via `/developer-kit:build-milestone`.
**Last closed:** **M35 — Stories & Heroes model + multi-org — 2026-06-23** (v1.9 "storytelling"; the `stories[]`
blueprint + `EffectiveStories()` multi-org normalization + the thriving/struggling/manager hero trio +
trajectory logic + the `jobroleref` resolver; merged into `release/01.90-storytelling`). Detail in the `### M35`
block of [`roadmap.md`](roadmap.md).
**Phase:** **v1.9 in development — M34 + M35 closed; M36 next (M37 may run ∥ M36).**
**Paused:** _(none)_

**Carry-forward / user-authorized follow-ups (from v1.8 close, still open):** the live field-bake on a
freshly-emptied `stack-demo/`; pushing the ext tags (`understudy-m26` + `house-lights-m31`/`m32` +
`stage-door-m27`/`m28`/`m30` + `prop-room-m21..m25`) to `origin`. The **`wip/clerkenstein-browser-login`**
branch now has its design home — **v1.9 M37** builds on it.

## Recently shipped releases
- **v1.8 "understudy"** — **2026-06-15**, tag `v1.8`. Self-contained-demo release: a demo builds **entirely from
  `stack-demo`'s own clone set** (a box with only `stack-demo/` runs a demo end-to-end). Single `section`
  milestone **M26**. Code: `rosetta-extensions` @ tag `understudy-m26`. Records:
  [releases/archive/01.80-understudy/](releases/archive/01.80-understudy/).
- **v1.7 "house lights"** — **2026-06-15**, tag `v1.7`. Demo-UI-hardening: M31 mkcert FAPI cert (next-web stops
  blanking) + M32 studio-desk single-port/production fix. Ext tags `house-lights-m31`/`m32`.
- **v1.6 "stage door"** — **2026-06-14**, tag `v1.6`. Secret-provisioning: ingest a secret source → provision
  every repo's `.env` (values-blind) + a 6-repo/55-gene secret-coverage DNA + the `/stack-secrets` skill. M27→M30.

## Headline numbers (v1.9 baseline — inherited from the v1.8 close 2026-06-15; reset at each v1.9 milestone close)
- **Go test funcs:** **1115** total (`Test`+`Fuzz`). Per-module: `rosetta-extensions/alignment` 52 · clerkenstein
  223 · stack-seeding **347** (+45, M35: the multi-org Stories engine — `stories[]` blueprint + the hero trio +
  trajectory helpers + the `jobroleref` resolver + 3 harden passes; 425 incl. subtests; integration tests opt-in
  behind `//go:build integration`) · stack-snapshot 333 · stack-secrets 160. (v1.9 grows stack-seeding [M34–M36]
  + clerkenstein [M37] + a new demo-stack cockpit surface [M38].) `go vet`+`gofmt`+`shellcheck` clean; flake 0
  (M35 flake gate 5/5). stack-seeding coverage: blueprint 100% · seeders 98.0%.
- **Python tests:** **501** (demo-stack/tests 138 · stack-injection/tests 113 · …). Triple-clean 3/3.
- **Supply-chain:** **GREEN** (stdlib-only posture; 0 third-party deps added through v1.8).
- **Alignment gates (green since v1.0):** **100%/100%** on all 4 Clerkenstein surfaces — **M37 must hold these**
  while adding the multi-identity surface (a new measured surface, not a regression of the existing ones).

## Branch model
**v1.9 IN DEVELOPMENT:** `release/01.90-storytelling` cut from `main` at design time (2026-06-22) so milestone
branches (`m34/…` → `m38/…`) have a parent. Milestones merge into the release branch via
`/developer-kit:close-milestone`; the release merges into `main` + tags `v1.9` via `/developer-kit:close-release`.
Code lands in the `rosetta-extensions` `stack-seeding` / `clerkenstein` / `demo-stack` ext sections (authored in
`.agentspace/rosetta-extensions/`, consumed per-stack at a pinned tag).
**Shipped:** **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` · **v1.3b** `v1.3.1` ·
**v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.

_Last updated: 2026-06-23 (**v1.9 M35 CLOSED** via `/developer-kit:close-milestone` — the multi-org Stories &
Heroes engine [`stories[]` blueprint + `EffectiveStories()` normalization + the thriving/struggling/manager hero
trio + trajectory logic + the `jobroleref` resolver; #M34-D7 landed in full as D-M35-4], merged
`m35/stories-multi-org` → `release/01.90-storytelling`; close GREEN [2 findings, 0 blocking], deferral re-audit
GREEN, stack-seeding 347 tests / blueprint 100% · seeders 98.0%, flake 0. ext tag `storytelling-m35` @ `06d872c`.
Next: build M36 [M37 may run ∥ M36]. Prior: M34 closed 2026-06-23.)_
