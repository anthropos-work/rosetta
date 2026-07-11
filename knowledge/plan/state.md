---
active_release: "(between releases) — v2.2 panorama SHIPPED (tag v2.2); awaiting /developer-kit:design-roadmap"
active_branch: "main"
active_milestone: "(between releases)"
last_closed: "v2.2 panorama — 2026-07-12 (tag v2.2, 4 milestones M212..M215) — external-shareability over Tailscale"
phase: "between releases — awaiting /developer-kit:design-roadmap"
last_updated: "2026-07-12"
---

# State

**Active release:** **(between releases).** **v2.2 "panorama"** — the **external-shareability release** — SHIPPED
2026-07-12 (tag `v2.2`). No release is active; the next v2.x release awaits `/developer-kit:design-roadmap`.

**Last shipped:** **v2.2 "panorama" — 2026-07-12, tag `v2.2`** (rext code-of-record rolled to `v2.2` = `39e8013`
[the D-CLOSE-1/-2/-3 README reconcile + ADV-1 F12 comment atop `panorama-m215` @ `00ba6b6`]; `.agentspace/rext.tag`
bumped `v2.1` → `v2.2`). Make dev/demo stacks reachable from **another machine on a Tailscale tailnet** — run a demo
on a Tailscale VM (e.g. `billion.taildc510.ts.net` on the odyssey Proxmox host) and a teammate with Tailscale up
browses it end-to-end over one trusted HTTPS origin. **Opt-in, default-off** (`/demo-up N --public-host <magicdns>`);
HTTPS-everywhere under one MagicDNS origin (Clerk needs a secure context); `tailscale cert` (real LE cert, no CA
install) fronted by per-offset-port `tailscale serve`; CORS trio + the ant-academy sha-pinned patch; the fresh-Linux-VM
host prereqs the tooling pre-flights / auto-handles / fails-loud on. **The FIRST live remote Linux-VM deploy** — proven
end-to-end for BOTH hero vantages on a trusted cert, cold reset-to-seed reproducible; **unset knob byte-identical**.
**4 milestones M212 → { M213 ∥ M214 } → M215**, all merged `--no-ff` → `release/02.20-panorama` → `main`. **Tooling +
docs + an opt-in flag only — zero platform-repo edits, 0 net-new deps.** Records:
[`releases/archive/02.20-panorama/`](releases/archive/02.20-panorama/).

**Active milestone:** **(between releases).** M215 "prove it on odyssey" was the final v2.2 milestone (closed
2026-07-11 `closed-on-gate`); v2.2 is closed + tagged. No milestone is active until the next release is designed.

**Phase:** **between releases — awaiting `/developer-kit:design-roadmap`.** v2.2 "panorama" merged → `main` + tagged
`v2.2`; the `release/02.20-panorama` branch is deleted. The next v2.x release (or a promotion of optional **M216**
dev-path parity / a vision item) is a `/developer-kit:design-roadmap` away.

**Next up:** **`/developer-kit:design-roadmap`** to scope the next release. Candidate seeds carried in
[`roadmap-vision.md`](roadmap-vision.md): optional **M216** (dev-path Tailscale parity + an operator surface — the
declared v2.2 scope-flex lever, never scaffolded), the Playthroughs futures M205–M207, and the standing backlog.

**Origin sync (push at 2026-07-12 close):** before this close, origin carried `main` + shipped tags (`v1.10.1`,
`v2.0`, `v2.1`) for **both** rosetta and rosetta-extensions + the `release/02.20-panorama` branch. **This close
produced (to push):** rosetta `main` now carries the `release/02.20-panorama` merge (all four milestone merges + the
close commit) + tag **`v2.2`**; rext `main` carries the new close commit `39e8013` + the four per-milestone annotated
tags (`panorama-m212/213/214/215`) + the rolled release tag **`v2.2`** (`11fa0a62` → `39e8013`). Push both repos'
`main` + all new tags.

## Recently shipped releases
- **v2.2 "panorama"** — **2026-07-12**, tag `v2.2`. **External-shareability over Tailscale** (M212..M215); first live
  remote Linux-VM demo deploy, both vantages green on a trusted cert, cold reset-to-seed reproducible. Records:
  [`releases/archive/02.20-panorama/`](releases/archive/02.20-panorama/).
- **v2.1 "quick change"** — **2026-07-09**, tag `v2.1`. The **skiller-in-app re-ground** (M208..M211). Records:
  [`releases/archive/02.10-quick-change/`](releases/archive/02.10-quick-change/).
- **v2.0 "opening night"** — **2026-07-02**, tag `v2.0`. The **Playthroughs** pillar (M201..M204); 10 live
  Playthroughs GREEN on cold reset-to-seed. Records:
  [`releases/archive/02.00-opening-night/`](releases/archive/02.00-opening-night/).
- **v1.10b "fit-up"** — **2026-07-01**, tag `v1.10.1`. Field-hardening backfill (M47..M53). Records:
  [`releases/archive/01.10b-fit-up/`](releases/archive/01.10b-fit-up/).

_(Earlier v1.x — v1.0 … v1.10 — full shipped table in [`roadmap-legacy.md`](roadmap-legacy.md) § Shipped releases.)_

## Headline numbers (v2.2 "panorama" — FINAL, at close-release)
- **rext Go test funcs:** **1772** across 6 modules (v2.1 baseline 1764 + **8** in clerkenstein: `clerkjs_proxy` +7
  [M213] + `handshake` +1; M212 + M214 + M215 touched zero Go). `go test ./...` exit 0 + `go vet` clean all 6.
- **rext Python:** demo-stack **424** (0f — the M215 delta is **+41** net-new: host pre-flight / keyless ssh-agent /
  Linux data-dirs / atlas-fail-loud + the ~400-tag F3 regression fixture + the F12 `tailscale serve` teardown/up-path
  reset), stack-injection **147p/8s**, stack-core **97** — **668** total passed. **TS e2e** **124** (playwright
  `--list`; tree byte-identical to v2.1). **TS unit specs:** **103**. **Live Playthroughs:** **10** + 1 in-manifest TODO.
- **Flake:** **0** (triple-clean 3/3). **Supply-chain:** **0 net-new deps**, 0 reachable vulns (govulncheck, go1.25.12);
  13 pre-existing dependabot alerts all the unreachable `x/crypto` ssh-subpackage in clerkenstein → cleared next rext
  roll (`go get x/crypto@v0.52.0`). **Platform-repo edits:** **0**. rext code-of-record `v2.2` = `39e8013`.

## Branch model / shipped tags
**v2.2 SHIPPED:** `release/02.20-panorama` merged `--no-ff` → `main` + tagged `v2.2` at `/developer-kit:close-release`
(2026-07-12); the release branch is deleted. The next release cuts a new `release/{version}` from `main` at
`/developer-kit:design-roadmap`.
**Shipped tags:** **v2.2** `v2.2` · **v2.1** `v2.1` · **v2.0** `v2.0` · **v1.10b** `v1.10.1` · **v1.10** `v1.10` ·
**v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3**
`v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`. (Full shipped detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

## Standing backlog (unscheduled, cross-release)
- **v2.2 residuals → Fate-2 standing backlog:** F5/DEF-M215-01 (app demopatch sha re-anchor), F9/DEF-M215-02 (remote-VM
  snapshot-cache pre-stage/auto-sync), F11/DEF-M215-03 (seed hero-name cosmetic + committed remote-origin Playwright
  gate), F13/DEF-M215-04 (jobsimulation exits(1)) — all orthogonal to shareability, none on the proven journey path.
  Plus the next-rext-roll `x/crypto@v0.52.0` bump. Routed in
  [`releases/archive/02.20-panorama/m215-prove-on-odyssey/carry-forward.md`](releases/archive/02.20-panorama/m215-prove-on-odyssey/carry-forward.md).
- **Optional M216** (dev-path Tailscale parity + operator surface) — the declared v2.2 scope-flex lever; stays
  roadmap-only until promoted via `/developer-kit:design-roadmap`.
- **Older:** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes), DEF-M21-01 (`replayCmd` hermetic test), CAVEAT-1
  (clean-box literal full `/dev-up`), M314b (prod frozen-read hydration — a prod-team follow-up). Reserved
  **Playthroughs futures** M205–M207 stay in vision. All tracked in [`roadmap-vision.md`](roadmap-vision.md).

_Last updated: 2026-07-12 (`/developer-kit:close-release` — v2.2 "panorama" merged `--no-ff` → `main` + tagged `v2.2`;
release branch deleted; rext rolled to `v2.2` = `39e8013` + `.agentspace/rext.tag` bumped `v2.1`→`v2.2`; release dir
archived → `releases/archive/02.20-panorama/`. All 3 close gates: metrics GREEN, deferral GREEN, supply-chain
YELLOW-non-blocking. Between releases — awaiting `/developer-kit:design-roadmap`.)_
