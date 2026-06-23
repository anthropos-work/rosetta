# State

**Active version:** **(between releases)** — v1.9 "storytelling" SHIPPED 2026-06-23 (tag `v1.9`, merged
`--no-ff` → `main`). No version is currently in development; the next awaits a `/developer-kit:design-roadmap`
run. v1.9 was the **believable-demo-narrative release** — it turned the placeholder seeder into a declarative
**Stories & Heroes** engine (per-story org + a thriving/struggling/manager **hero** trio, seeded via the real
**verified-skill chain**) so the individual **skill profile** + the org **Workforce dashboard** tell one
coherent story, plus a standalone **presenter cockpit** (login-as a hero + jump-to the right screen) on
**Clerkenstein multi-identity**. 5 `section` milestones **M34→M38**; **tooling + docs only — zero
platform-repo edits**.

**Active milestone:** **(between releases)** — no milestone in progress.
**Next up:** **`/developer-kit:design-roadmap`** — design the next version (none staged behind v1.9 in
`roadmap-vision.md`). The unscheduled backlog (DEF-M10-01, DEF-M21-01..04, M25-D9) is candidate input. (M33
ant-academy liveness was resolved by the post-v1.9 demo-hardening patch — see below.)
**Last closed:** **v1.9 "storytelling" (release) — 2026-06-23** via `/developer-kit:close-release` — reviewed
all 5 milestones (M34–M38) as one PR: **GREEN, 0 blocking findings**. 9 review sweeps clean (supply-chain /
scope / deferral-re-audit / code-quality / docs / KB-consolidation / tests / metrics-regression / decisions);
the release-level docs coherence review returned **0 findings**. Deferral re-audit **GREEN** (0 open, 0
repeat, 0 aged-out, **0 escape-hatch**; both v1.9 items landed Fate-1 in-release — #M34-D7→D-M35-4, M38-D7→
M38-D8). Triple-clean gate 3/3 (local). Merged `release/01.90-storytelling` → `main` + tagged `v1.9`.
**Phase:** **between releases — awaiting `/developer-kit:design-roadmap`.**
**Paused:** _(none)_

**Carry-forward / user-authorized follow-ups (still open):** the live field-bake on a freshly-emptied
`stack-demo/`; the literal browser-pixels Stories & Heroes cockpit end-to-end (now the default `/demo-up`,
a deliberate demo step);
**pushing the ext tags + `main` + the `v1.9` tag to `origin`** (the orchestrator's separate post-close step —
the close merged + tagged LOCALLY only). Ext tags pending push: `storytelling-m34..m38` + the prior
`understudy-m26` / `house-lights-m31`/`m32` / `stage-door-m27`/`m28`/`m30` / `prop-room-m21..m25`.

**Post-v1.9 tooling-hardening pass (between releases) — SHIPPED.** A demo-hardening patch landed on the
demo-stack / dev-setdress tooling (`rosetta-extensions` @ tag `storytelling-postfix-1`): **`DEMO_STORIES` is
now default-on** (a bare `/demo-up N` seeds the multi-org Stories & Heroes world + serves the presenter
cockpit; `DEMO_NO_STORIES=1` opts back out to the legacy structural small-200 + single-identity fake-fapi +
no-cockpit demo); the **M33** ant-academy + cockpit "dead on a later visit" bug is **resolved** (both
host-native daemons now launch session-detached via `demo-stack/detach.sh::launch_detached` instead of bare
`nohup`, so they survive the launching session/task ending — ant-academy is otherwise a non-fatal skip when
its Font Awesome Pro deps aren't installed); the per-stack **Directus boot now health-gates** on the stack's
own offset `/server/health` before returning so the bring-up-tail autoverify can't race its ~30s
re-introspect; and the **prod-Directus content note is guarded** (printed only on `DEMO_NO_LOCAL_CONTENT=1`,
since the default demo boots a per-stack Directus serving the captured catalog locally). **Tooling + docs
only — zero platform-repo edits; all 5 Clerkenstein alignment gates still 100%/100%.**

## Recently shipped releases
- **v1.9 "storytelling"** — **2026-06-23**, tag `v1.9`. Believable-demo-narrative release: the placeholder
  seeder becomes a declarative **Stories & Heroes** engine (per-story org + a hero trio via the real
  verified-skill chain) so the skill profile + the Workforce dashboard tell a story, plus a presenter cockpit
  on Clerkenstein multi-identity. 5 `section` milestones **M34→M38**. Headline: zero platform-repo edits; all
  **5** Clerkenstein alignment gates 100%/100%; supply-chain GREEN (0 new deps); Go 1027→**1248** (stack-seeding
  259→444). Code: `rosetta-extensions` @ tags `storytelling-m34..m38`. Records:
  [releases/archive/01.90-storytelling/](releases/archive/01.90-storytelling/).
- **v1.8 "understudy"** — **2026-06-15**, tag `v1.8`. Self-contained-demo release: a demo builds **entirely from
  `stack-demo`'s own clone set** (a box with only `stack-demo/` runs a demo end-to-end). Single `section`
  milestone **M26**. Code: `rosetta-extensions` @ tag `understudy-m26`. Records:
  [releases/archive/01.80-understudy/](releases/archive/01.80-understudy/).
- **v1.7 "house lights"** — **2026-06-15**, tag `v1.7`. Demo-UI-hardening: M31 mkcert FAPI cert (next-web stops
  blanking) + M32 studio-desk single-port/production fix. Ext tags `house-lights-m31`/`m32`.

## Headline numbers (v1.9 "storytelling" — final close values, 2026-06-23)
- **Go test funcs:** **1248** total (`Test`+`Fuzz`). Per-module: `rosetta-extensions/alignment` 52 ·
  clerkenstein **259** (250 `Test` + 9 `Fuzz`) · stack-seeding **444** · stack-snapshot 333 · stack-secrets 160.
  Verified at close: stack-seeding `go test ./...` 8 pkgs `ok`, clerkenstein 14 pkgs `ok`; `go vet`+`gofmt`
  clean. Baseline v1.8 1027 → **+221**.
- **Python tests:** **283** across the two v1.9-touched suites (demo-stack/tests **166**, stack-injection/tests
  **117** [8 opt-in skipped]). Both green; no suite decreased (untouched rext suites carry forward at v1.8
  counts).
- **Flake:** **0** — triple-clean release gate 3/3 (stack-seeding incl `-race` + clerkenstein, shuffled; Python
  re-verified).
- **Supply-chain:** **GREEN** — `go.mod`/`go.sum` diff `storytelling-m34~1..m38` EMPTY (byte-identical to v1.8);
  0 new third-party deps; all deps MIT/BSD/Apache.
- **Alignment gates (green since v1.0):** **100%/100%** on **all 5** Clerkenstein surfaces — M37 added the
  multi-identity `clerk-multi-1` (9 genes) and held the 4 existing (Go 22/22, JS 9/9, deploy 7/7) green (the
  `clerk-express-1` node-CI gate is an env prereq, not a regression).

## Branch model
**v1.9 SHIPPED:** `release/01.90-storytelling` merged `--no-ff` → `main` + tagged `v1.9` (LOCAL — origin push
is the orchestrator's separate step). The release branch is deleted. Code-of-record lives in the
`rosetta-extensions` authoring copy (a SEPARATE repo) at tags `storytelling-m34..m38` — close-release merged
ONLY the rosetta doc-half branch.
**Shipped:** **v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` ·
**v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.

_Last updated: 2026-06-23 (**v1.9 "storytelling" SHIPPED** via `/developer-kit:close-release` — reviewed
M34–M38 as one PR, **GREEN/0 blocking**; release-level docs coherence 0 findings; deferral re-audit GREEN [0
open, 0 escape-hatch]; supply-chain GREEN [0 new deps]; all 5 Clerkenstein alignment gates 100%/100%; Go
1027→1248. Merged `release/01.90-storytelling` → `main`, tagged `v1.9`, branch deleted — origin push pending
[orchestrator's step]. **Next: `/developer-kit:design-roadmap`.** Prior: v1.8 "understudy" SHIPPED 2026-06-15.
**Post-v1.9 (between releases): demo-hardening patch SHIPPED** (`rosetta-extensions` @ `storytelling-postfix-1`)
— `DEMO_STORIES` default-on + `DEMO_NO_STORIES` opt-out; M33 ant-academy/cockpit session-detach fix; Directus
boot health-gate; guarded prod-Directus note. Tooling + docs only, zero platform-repo edits, 5 alignment gates
still 100%. M33 dropped from the backlog.)_
