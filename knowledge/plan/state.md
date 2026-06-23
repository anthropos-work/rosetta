# State

**Active version:** **v1.9 "storytelling" — IN DEVELOPMENT** (designed 2026-06-22 via
[`/developer-kit:design-roadmap`](roadmap.md); branch `release/01.90-storytelling`). The
**believable-demo-narrative release** — turn the placeholder seeder into a declarative **Stories & Heroes**
engine: each *story* is one org with a thriving/struggling/manager **hero** trio, seeded via the real
**verified-skill chain** (the 7-table jobsim→`user_skills`→`user_skill_evidences` fan-out) so the **skill
profile** + the org **Workforce dashboard** tell one coherent story — plus a standalone **presenter cockpit**
(log in as a hero + jump to the right screen). 5 `section` milestones **M34→M38** across `rosetta-extensions`
(`stack-seeding`/`clerkenstein`/`demo-stack`) + the rosetta corpus doc-half. **Tooling + docs only — zero
platform-repo edits.** **M34 ✅ shipped 2026-06-23** (the verified-skill spine); M35 next.

**Active milestone:** **M35 — Stories & Heroes model + multi-org** — **NEXT (not started).** One declarative
`stack.stories.yaml` seeds **multiple orgs**, each with its thriving/struggling/manager **hero trio** at
vantage-appropriate fidelity: the `stories[]` blueprint (supersedes the org-centric `stack.seed.yaml` for demo
stacks), multi-org `OrgID`/`orgClerkID` threaded through the 4 consuming seeders + Clerkenstein org-claim
alignment, `PersonaSeeder` scaled from M34's one hero to the locked 2-story × 3-hero roster, the trajectory
logic, and supporting-population fidelity. Depends on M34 (the verified-skill spine, now shipped). M35 also
picks up the two M34-routed roster guards (#M34-D7: the `len(Personas) <= Size` validation + index-collision
warning, and the short-role-pool top-up product call). The verified spec
([`.agentspace/seeding_gaps.md`](../../.agentspace/seeding_gaps.md)) remains the authoritative design for
M35–M38. The M34–M38 overviews live under [`releases/01.90-storytelling/`](releases/01.90-storytelling/).
**Next up:** build M35 via `/developer-kit:build-milestone`.
**Last closed:** **M34 — Verified-skill chain (vertical slice) — 2026-06-23** (v1.9 "storytelling"; the
verified-skill 7-table chain + G14 fix + closure gene; merged into `release/01.90-storytelling`). Detail in the
`### M34` block of [`roadmap.md`](roadmap.md).
**Phase:** **v1.9 in development — M34 closed; M35 next.**
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
- **Go test funcs:** **1070** total (`Test`+`Fuzz`). Per-module: `rosetta-extensions/alignment` 52 · clerkenstein
  223 · stack-seeding **302** (+43, M34: PersonaSeeder + TaxonomyRefs + G14 + closure gene + 2 harden passes;
  381 incl. subtests; integration tests opt-in behind `//go:build integration`) · stack-snapshot 333 ·
  stack-secrets 160. (v1.9 grows stack-seeding [M34–M36] + clerkenstein [M37] + a new demo-stack cockpit surface
  [M38].) `go vet`+`gofmt`+`shellcheck` clean; flake 0 (M34 flake gate 5/5). stack-seeding coverage: seeders
  96.6% · `dna/seed_closure.go` 100%.
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

_Last updated: 2026-06-23 (**v1.9 M34 CLOSED** via `/developer-kit:close-milestone` — the verified-skill chain
[PersonaSeeder 7-table fan-out + G14 fix + TaxonomyRefs + closure gene], merged `m34/verified-skill-chain` →
`release/01.90-storytelling`; close GREEN [5 findings, 0 blocking], deferral re-audit GREEN, stack-seeding
381 tests / 96.6%, flake 0. ext tag `storytelling-m34` @ `8eb603b`. Next: build M35. Prior: 2026-06-22 v1.9
DESIGNED.)_
