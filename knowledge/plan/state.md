# State

**Active version:** **v1.9 "storytelling" — IN DEVELOPMENT** (designed 2026-06-22 via
[`/developer-kit:design-roadmap`](roadmap.md); branch `release/01.90-storytelling`). The
**believable-demo-narrative release** — turn the placeholder seeder into a declarative **Stories & Heroes**
engine: each *story* is one org with a thriving/struggling/manager **hero** trio, seeded via the real
**verified-skill chain** (the 7-table jobsim→`user_skills`→`user_skill_evidences` fan-out) so the **skill
profile** + the org **Workforce dashboard** tell one coherent story — plus a standalone **presenter cockpit**
(log in as a hero + jump to the right screen). 5 `section` milestones **M34→M38** across `rosetta-extensions`
(`stack-seeding`/`clerkenstein`/`demo-stack`) + the rosetta corpus doc-half. **Tooling + docs only — zero
platform-repo edits.** **M34 ✅ + M35 ✅ + M36 ✅ + M37 ✅ shipped 2026-06-23** (the verified-skill spine + the
multi-org Stories engine + the Workforce-dashboard surfaces + Clerkenstein multi-identity); **M38 (the
presenter cockpit) is the one remaining milestone.**

**Active milestone:** **M38 — Presenter cockpit (the LAST of v1.9).** A **standalone served panel** (rext
`demo-stack`, offset port — not an in-app overlay, preserving the zero-platform-repo-edit line) that lists
stories→heroes with **[Login as]** + **[Jump to section]**, reading the **same** `stack.stories.yaml`: it
wires M37's multi-identity seat-switch + the seeder roster into a clickable login-as-a-hero cockpit. **In:**
the panel; the deep-link catalog (enumerated next-web routes per vantage — **O9**); login-as wired to M37;
jump-to. Depends on M37 (shipped). The verified spec
([`.agentspace/seeding_gaps.md`](../../.agentspace/seeding_gaps.md)) remains the authoritative design. The
overviews live under [`releases/01.90-storytelling/`](releases/01.90-storytelling/).
**Next up:** build M38 via `/developer-kit:build-milestone` (the last milestone — then `/developer-kit:close-release`).
**Last closed:** **M37 — Clerkenstein multi-identity — 2026-06-23** (v1.9 "storytelling"; the users/orgs
registry replacing `DefaultDemoUser` + the server-authoritative active-seat selection [O11 → parameterized
FAPI handshake] + the 5th Alignment DNA `clerk-multi-1` [9 genes, 100%/100%] + the `wip/clerkenstein-browser-login`
reconcile/retire; merged into `release/01.90-storytelling`). Detail in the `### M37` block of
[`roadmap.md`](roadmap.md).
**Phase:** **v1.9 in development — M34–M37 closed (both product Musts + the seat-switch capability done); M38
(the presenter cockpit, the LAST milestone) remains.**
**Paused:** _(none)_

**Carry-forward / user-authorized follow-ups (from v1.8 close, still open):** the live field-bake on a
freshly-emptied `stack-demo/`; pushing the ext tags (`understudy-m26` + `house-lights-m31`/`m32` +
`stage-door-m27`/`m28`/`m30` + `prop-room-m21..m25` + the v1.9 `storytelling-m34..m37`) to `origin`. The
**`wip/clerkenstein-browser-login`** branch was reconciled (note folded into `architecture.md`) + **retired**
at M37 close — no longer open.

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
- **Go test funcs:** **1248** total (`Test`+`Fuzz`). Per-module: `rosetta-extensions/alignment` 52 · clerkenstein
  **259** (250 `Test` + 9 `Fuzz`) · stack-seeding **444** (M38: +2 close — the `roleForHero` vantage-faithful
  `org_role` + the three-write-lockstep regression [`TestBuildRoster_OrgRoleVantageFaithfulAndLockstep`,
  `TestRoleForHero`]; the M38 roster/cockpit producers + the O9 catalog land here) · stack-snapshot 333 ·
  stack-secrets 160. `go vet`+`gofmt`+`shellcheck` clean; flake 0 (M38 flake gate 5/5 Go + Python). M38 coverage:
  the milestone-touched producers ~100% (`roster.go`/`cockpit.go`/`cmd/stackseed`); `roleForHero` all branches.
- **Python tests:** **283** counted across the M38-touched surfaces (demo-stack/tests **166** [M38: +1 close —
  the cockpit empty-key defensive-skip test] · stack-injection/tests **117** [8 opt-in skipped]). All green.
- **Supply-chain:** **GREEN** (stdlib-only posture; 0 third-party deps added through v1.9 M37).
- **Alignment gates (green since v1.0):** **100%/100%** on **all 5** Clerkenstein surfaces — M37 added the
  multi-identity `clerk-multi-1` (9 genes) and held the 4 existing ones (Go 22/22, JS 9/9, deploy 7/7) green
  (the `clerk-express-1` node-CI gate drives the genuine `@clerk/express` SDK — runs where npm deps are
  installed; an env prereq, not a regression).

## Branch model
**v1.9 IN DEVELOPMENT:** `release/01.90-storytelling` cut from `main` at design time (2026-06-22) so milestone
branches (`m34/…` → `m38/…`) have a parent. Milestones merge into the release branch via
`/developer-kit:close-milestone`; the release merges into `main` + tags `v1.9` via `/developer-kit:close-release`.
Code lands in the `rosetta-extensions` `stack-seeding` / `clerkenstein` / `demo-stack` ext sections (authored in
`.agentspace/rosetta-extensions/`, consumed per-stack at a pinned tag).
**Shipped:** **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` · **v1.3b** `v1.3.1` ·
**v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.

_Last updated: 2026-06-23 (**v1.9 M37 CLOSED** via `/developer-kit:close-milestone` — Clerkenstein multi-identity:
the users/orgs registry replacing `DefaultDemoUser` + the server-authoritative active-seat selection [O11 →
parameterized FAPI handshake] + the 5th Alignment DNA `clerk-multi-1` [9 genes, 100%/100%] + the
`wip/clerkenstein-browser-login` reconcile/retire; merged `m37/clerkenstein-multi-identity` →
`release/01.90-storytelling`; close GREEN [1 finding, 0 blocking — a stale "four DNAs" corpus claim fixed],
deferral re-audit GREEN, clerkenstein 250 tests + 9 fuzz `-race` / `registry.go` 100%, flake 0 [5/5], all 5
alignment gates 100%/100%. ext tag `storytelling-m37` @ `52c1be0`. Next: build M38 [the presenter cockpit, the
LAST milestone]. Prior: M36 closed 2026-06-23.)_
