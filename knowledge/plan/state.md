# State

**Active version:** **v1.9 "storytelling" — IN DEVELOPMENT** (designed 2026-06-22 via
[`/developer-kit:design-roadmap`](roadmap.md); branch `release/01.90-storytelling`). The
**believable-demo-narrative release** — turn the placeholder seeder into a declarative **Stories & Heroes**
engine: each *story* is one org with a thriving/struggling/manager **hero** trio, seeded via the real
**verified-skill chain** (the 7-table jobsim→`user_skills`→`user_skill_evidences` fan-out) so the **skill
profile** + the org **Workforce dashboard** tell one coherent story — plus a standalone **presenter cockpit**
(log in as a hero + jump to the right screen). 5 `section` milestones **M34→M38** across `rosetta-extensions`
(`stack-seeding`/`clerkenstein`/`demo-stack`) + the rosetta corpus doc-half. **Tooling + docs only — zero
platform-repo edits.** **M34 ✅ + M35 ✅ + M36 ✅ shipped 2026-06-23** (the verified-skill spine + the multi-org
Stories engine + the Workforce-dashboard surfaces — both product Musts now done); M37 next.

**Active milestone:** **M37 — Clerkenstein multi-identity** — **PLANNED (not started).** A demo stack can
**switch the active browser identity** among the seeded heroes/orgs — the seat-switch the cockpit's "login as"
needs: a users/orgs registry in `clerk-frontend` (replacing the single `DefaultDemoUser`), an active-user
selection mechanism (token-injection vs a parameterized FAPI handshake — O11, spike both early), and an
**Alignment DNA** for the new multi-identity surface (must hold the 100%/100% Clerkenstein gates). Builds on the
existing `wip/clerkenstein-browser-login` branch. Code lands in the `clerkenstein` ext section (a different
section from M34–M36's `stack-seeding`). Depends on M35 (shipped — needs only the hero-identity list). The
verified spec ([`.agentspace/seeding_gaps.md`](../../.agentspace/seeding_gaps.md)) remains the authoritative
design for M37–M38. The overviews live under [`releases/01.90-storytelling/`](releases/01.90-storytelling/).
**Next up:** build M37 via `/developer-kit:build-milestone`.
**Last closed:** **M36 — Dashboard surfaces (Must #2) — 2026-06-23** (v1.9 "storytelling"; the 6
Workforce-dashboard seeders [`membership_skills` funnel + `tags`/teams + `target_roles` gap/mobility +
`succession` + `feedback` + `population_evidence` org-scale gap] + the assignments status-mix fix + the
closure-gene 4th surface; merged into `release/01.90-storytelling`). Detail in the `### M36` block of
[`roadmap.md`](roadmap.md).
**Phase:** **v1.9 in development — M34 + M35 + M36 closed (both product Musts done); M37 next, M38 after.**
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
- **Go test funcs:** **1174** total (`Test`+`Fuzz`). Per-module: `rosetta-extensions/alignment` 52 · clerkenstein
  223 · stack-seeding **406** (+59, M36: the 6 Workforce-dashboard seeders + the assignments status-mix fix +
  the named-skill resolver + 3 harden passes; 484 incl. subtests; integration tests opt-in behind
  `//go:build integration`) · stack-snapshot 333 · stack-secrets 160. (v1.9 grows stack-seeding [M34–M36 ✅]
  + clerkenstein [M37] + a new demo-stack cockpit surface [M38].) `go vet`+`gofmt`+`shellcheck` clean; flake 0
  (M36 flake gate 5/5). stack-seeding coverage: blueprint 100% · seeders 95.5% · dna 87.7%.
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

_Last updated: 2026-06-23 (**v1.9 M36 CLOSED** via `/developer-kit:close-milestone` — the org Workforce-dashboard
surfaces [6 seeders: `membership_skills` funnel + `tags`/teams + `target_roles` gap/mobility + `succession` +
`feedback` + `population_evidence` org-scale gap; + the assignments status-mix fix + the skillpath completed-share
+ the closure-gene 4th surface], merged `m36/dashboard-surfaces` → `release/01.90-storytelling`; close GREEN [4
findings, 0 blocking], deferral re-audit GREEN, stack-seeding 406 tests / blueprint 100% · seeders 95.5% · dna
87.7%, flake 0 [5/5]. ext tag `storytelling-m36` @ `11e15e3`. Both product Musts done. Next: build M37
[Clerkenstein multi-identity]. Prior: M35 closed 2026-06-23.)_
