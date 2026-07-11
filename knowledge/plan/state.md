---
active_release: "v2.2 panorama — IN DEVELOPMENT (branch release/02.20-panorama; tag v2.2 at close); designed 2026-07-11"
active_branch: "release/02.20-panorama"
active_milestone: "M212 — the single host knob (planned, next up)"
last_closed: "v2.1 quick change — 2026-07-09 (tag v2.1, 4 milestones M208..M211) — the skiller-in-app re-ground"
phase: "in development — M212 up next (run /developer-kit:build-milestone)"
last_updated: "2026-07-11"
---

# State

**Active release:** **v2.2 "panorama"** — the **external-shareability release**: make dev/demo stacks reachable from
other machines on a **Tailscale** tailnet (run a stack on a Tailscale VM, e.g. `billion.taildc510.ts.net` on the
odyssey Proxmox host; a teammate with Tailscale up browses the demo end-to-end). Designed 2026-07-11 via
`/developer-kit:design-roadmap`; branch `release/02.20-panorama` cut from `main`; tag **`v2.2`** at close. **4
milestones M212 → { M213 ∥ M214 } → M215** (+ optional/deferrable **M216**). **Decisions:** HTTPS-everywhere under
one MagicDNS origin (Clerk needs a secure context); external access **opt-in, default off** (an explicit
`/demo-up --public-host` flag); **demo-first** (dev-path parity = optional M216). **Tooling + docs + an opt-in flag
only — zero platform-repo edits** (two platform-family files ride the EXISTING rext sha-pinned patch mechanism). The
**sanctioned re-proposal** of the v1.4 seed "external stack shareability (Tailscale/ingress)" dropped 2026-06-11.
Records: [`releases/02.20-panorama/`](releases/02.20-panorama/).

**Last shipped:** **v2.1 "quick change" — 2026-07-09, tag `v2.1`** (rext code-of-record `v2.1` =
`quick-change-m211`; `.agentspace/rext.tag` = `v2.1`). The **skiller-in-app re-ground** — the merged platform stands
up cold on both stacks via the re-grounded tooling. 4 milestones M208..M211. Records:
[`releases/archive/02.10-quick-change/`](releases/archive/02.10-quick-change/).

**Active milestone:** **M212 — The single host knob** (`section`, `planned`). Introduce `STACK_PUBLIC_HOST` (default
`localhost` → byte-identical when unset) surfaced as the opt-in `/demo-up --public-host` flag, threaded through every
rext emitter that bakes a browser-facing `localhost`/`127.0.0.1`. Not yet started — no milestone branch cut.

**Phase:** **in development.** The release branch is cut + all four milestone contracts are scaffolded under
`releases/02.20-panorama/`; the feasibility is established (workflow `wf_bea3be47` — config-only core + a 2-item
patch tail; Tailscale confirmed live both ends). Design-roadmap Phase 0 note: the deferral audit + KB blind-area
check were folded into the feasibility pass (the release **authors** its KB anchor `tailscale-serve.md` in M214, so
the one blind area is homed via a `Delivers →` line, not designed into).

**Next up:** **run `/developer-kit:build-milestone`** → opens **M212** (creates `m212/public-host-knob` from
`release/02.20-panorama`, accumulates commits). Then { M213 ∥ M214 } → M215 (the iterative acceptance gate). The
optional **M216** (dev-path parity) is roadmap-only until promoted.

**Push-gated KEEP (the user's manual step):** origin has NOT received `main` + tags `v1.10.1` + `v2.0` + `v2.1` +
the rext tags. The new `release/02.20-panorama` branch is likewise local. `.agentspace/rext.tag` is `v2.1`. An
administrative KEEP, not a deferral.

**Standing backlog (unscheduled, cross-release):** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes), DEF-M21-01
(`replayCmd` hermetic test), CAVEAT-1 (clean-box literal full `/dev-up` — belt-and-suspenders), M314b (prod
frozen-read hydration — a prod-team follow-up). Reserved **Playthroughs futures** M205–M207 stay in vision. All
tracked in [`roadmap-vision.md`](roadmap-vision.md).

## Recently shipped releases
- **v2.1 "quick change"** — **2026-07-09**, tag `v2.1`. The **skiller-in-app re-ground** (M208..M211). Records:
  [`releases/archive/02.10-quick-change/`](releases/archive/02.10-quick-change/).
- **v2.0 "opening night"** — **2026-07-02**, tag `v2.0`. The **Playthroughs** pillar (M201..M204); 10 live
  Playthroughs GREEN on cold reset-to-seed. Records:
  [`releases/archive/02.00-opening-night/`](releases/archive/02.00-opening-night/).
- **v1.10b "fit-up"** — **2026-07-01**, tag `v1.10.1`. Field-hardening backfill (M47..M53). Records:
  [`releases/archive/01.10b-fit-up/`](releases/archive/01.10b-fit-up/).

_(Earlier v1.x — v1.0 … v1.10 — full shipped table in [`roadmap-legacy.md`](roadmap-legacy.md) § Shipped releases.)_

## Headline numbers (inherited baseline — v2.1 final; v2.2 sets its own at close)
- **rext Go test funcs:** **1764** across 6 modules (v2.1 baseline). `go vet ./...` clean; triple-clean 3/3.
- **rext TS unit specs:** **103**. **Live Playthroughs:** **10** (6 employee + 4 manager) + 1 in-manifest TODO.
- **Flake:** **0**. **Supply-chain:** target **0 net-new deps** (a reverse-proxy component, if not a stdlib/OS
  package, is the one supply-chain item to weigh in M213). **Alignment gates:** **100%/100%** on all 5 Clerkenstein
  surfaces (v2.2 touches the pk/FAPI host + cert, NOT the Clerk contract — re-score at M213/M215).

## Branch model / shipped tags
**v2.2 IN DEVELOPMENT:** `release/02.20-panorama` cut from `main` 2026-07-11. Milestone branches `m{N}/{slug}` cut
from it per `/developer-kit:build-milestone`; merge back at `/developer-kit:close-milestone`; the release merges →
`main` + tags `v2.2` at `/developer-kit:close-release`.
**Shipped tags:** **v2.1** `v2.1` · **v2.0** `v2.0` · **v1.10b** `v1.10.1` · **v1.10** `v1.10` · **v1.9** `v1.9` ·
**v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` ·
**v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`. (Full shipped detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

_Last updated: 2026-07-11 (v2.2 "panorama" DESIGNED + PROMOTED to active development — the external-shareability /
Tailscale-serve release; branch `release/02.20-panorama`, tag `v2.2`; 4 milestones M212 → { M213 ∥ M214 } → M215
(+ opt M216); opt-in default-off, HTTPS-everywhere, demo-first; the sanctioned re-proposal of the dropped v1.4
Tailscale/ingress seed. Next: `/developer-kit:build-milestone` → M212.)_
