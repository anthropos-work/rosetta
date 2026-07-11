---
active_release: "v2.2 panorama — IN DEVELOPMENT (branch release/02.20-panorama; tag v2.2 at close); designed 2026-07-11"
active_branch: "release/02.20-panorama"
active_milestone: "(between milestones — v2.2 'panorama' ALL 4 core milestones M212..M215 CLOSED+merged; ready for /developer-kit:close-release)"
last_closed: "M215 prove-on-odyssey — 2026-07-11 (v2.2 final milestone, closed-on-gate) | v2.1 quick change — 2026-07-09 (tag v2.1)"
phase: "in development — ALL 4 v2.2 milestones closed (M212, M213, M214, M215); next: /developer-kit:close-release"
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

**Active milestone:** **(between milestones)** — **M215 — Prove it on odyssey** (`iterative`, complexity **large**,
the FINAL v2.2 milestone) **CLOSED 2026-07-11 `closed-on-gate`**, merged `--no-ff` into `release/02.20-panorama`; the
`m215/prove-on-odyssey` branch is deleted. The FIRST remote Linux-VM demo deploy over Tailscale, driven end-to-end
from a DIFFERENT tailnet machine for BOTH vantages (employee `maya-thriving` → `/profile`; manager `dan-manager` →
`/enterprise/workforce`) on a genuinely trusted Let's Encrypt cert (`verify=0`, no CA install), 0 console errors, 0
localhost/prod ejects, assets rendering — reproducibly on a clean cold reset-to-seed one-shot; unset knob
byte-identical. The **user-directed propagation close-gate is SATISFIED** (every deployment finding F1/F2/F4/F6/F8/F9/
F12 in tools+KB+skills; the NEW `corpus/ops/demo/tailscale-serve.md` runbook stands a fresh Linux VM up unaided). rext
code-of-record FROZEN at tag `panorama-m215` @ `00ba6b6` (re-tag reserved for close-release). Close: demo-stack **424**
(+41) / stack-injection **147p/8s**, shellcheck clean; 4 findings (0 must-fix; 2 docs fixed; ADV-1 non-blocking; plan
backfill); deferral audit GREEN; 4 documented non-blocking residuals (F5/F9/F11/F13) → standing backlog Fate-2. Full
closure narrative in [`roadmap.md`](roadmap.md) § M215. **Next:** **`/developer-kit:close-release`** (v2.2 "panorama").

**Phase:** **in development — ALL 4 core milestones closed (M212, M213, M214, M215); ready for
`/developer-kit:close-release`.** The release branch carries all four `--no-ff` merges: the knob foundation + the
HTTPS-over-tailnet auth surface + the origins/links/patch tail + the live cross-machine acceptance (M215 proved it on
`billion` — both vantages green on a trusted cert, cold reset-to-seed reproducible).

**Next up:** **`/developer-kit:close-release`** for **v2.2 "panorama"** — the full release-level review + merge. It
rolls the rext **`v2.2`** tag (from the three per-milestone tags `panorama-m212/213/214` + the M215 tag
`panorama-m215`), bumps `.agentspace/rext.tag` from `v2.1` → `v2.2`, reconciles the three rext-README index residuals
**D-CLOSE-1/-2/-3** (+ the ADV-1 F12 comment/placement one-liner) in the single rext commit, then merges
`release/02.20-panorama` → `main` + tags **`v2.2`**. The **4 M215 residuals** (F5 demopatch re-anchor / F9 remote
snapshot-cache / F11 seed hero-name / F13 jobsimulation) are documented, non-blocking, Fate-2 → standing backlog
(routed in `releases/02.20-panorama/m215-prove-on-odyssey/carry-forward.md`). Optional **M216** (dev-path parity)
stays roadmap-only until promoted.

**Origin sync (2026-07-11):** origin has `main` + all shipped tags (`v1.10.1`, `v2.0`, `v2.1`) for **both** rosetta
and rosetta-extensions, **and** the `release/02.20-panorama` branch (pushed 2026-07-11, `765528d`). **Local ahead of
origin (push at release close):** the rosetta `release/02.20-panorama` now carries **all four merges — M212 + M213 +
M214 + M215** (+ their close commits); rext has **four new local annotated tags — `panorama-m212` @ `770f81b`,
`panorama-m213` @ `b9f41dd`, `panorama-m214` @ `99c86b7`, and `panorama-m215` @ `00ba6b6`** (the per-milestone
code-of-record; the box-level `.agentspace/rext.tag` stays `v2.1` until `/developer-kit:close-release` bumps it →
`v2.2` + reconciles the D-CLOSE-1/-2/-3 rext-README residuals + the ADV-1 F12 note). No push performed by the close.

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

## Headline numbers (v2.2 — ALL 4 milestones closed; baseline v2.1 final)
- **rext Go test funcs:** **1771** across 6 modules (v2.1 baseline 1764 + **7** M213 `clerkjs_proxy` funcs; M212 +
  M214 + M215 touched zero Go). `go vet ./...` + `gofmt` clean.
- **rext Python (v2.2-touched suites):** stack-injection **155** (147p/8s), demo-stack **424** (0f) — the M215 delta
  is **+41 demo-stack** net-new funcs (host pre-flight / keyless ssh-agent / Linux data-dirs / atlas-fail-loud + the
  ~400-tag F3 regression fixture + the F12 `tailscale serve` teardown/up-path reset); stack-injection unchanged.
  **TS unit specs:** **103**. **Live Playthroughs:** **10** + 1 TODO.
- **Flake:** **0** (5/5 gate). **Supply-chain:** **0 net-new deps held**. **Alignment gates:** **100%/100%** on all 5
  Clerkenstein surfaces (M215 **proved the remote-auth foundation live** on `billion` — both vantages green over a
  trusted cert). rext code-of-record FROZEN at `panorama-m215` @ `00ba6b6`.

## Branch model / shipped tags
**v2.2 IN DEVELOPMENT:** `release/02.20-panorama` cut from `main` 2026-07-11. Milestone branches `m{N}/{slug}` cut
from it per `/developer-kit:build-milestone`; merge back at `/developer-kit:close-milestone`; the release merges →
`main` + tags `v2.2` at `/developer-kit:close-release`.
**Shipped tags:** **v2.1** `v2.1` · **v2.0** `v2.0` · **v1.10b** `v1.10.1` · **v1.10** `v1.10` · **v1.9** `v1.9` ·
**v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` ·
**v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`. (Full shipped detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

_Last updated: 2026-07-11 (M215 "prove it on odyssey" CLOSED `closed-on-gate` + merged `--no-ff` into
`release/02.20-panorama` — the FINAL v2.2 milestone. The first remote Linux-VM demo over Tailscale, driven end-to-end
from a 2nd tailnet machine for BOTH vantages on a trusted LE cert [verify=0, no CA install], 0 console errors, 0
localhost/prod ejects, reproducible on a cold reset-to-seed; unset knob byte-identical. Propagation close-gate
SATISFIED [F1/F2/F4/F6/F8/F9/F12 in tools+KB+skills; the NEW corpus/ops/demo/tailscale-serve.md runbook]. rext
code-of-record @ tag panorama-m215 @ 00ba6b6 [FROZEN — re-tag at close-release]; demo-stack 424 / stack-injection
147p8s / shellcheck clean; 4 close findings [0 must-fix, 2 docs fixed, ADV-1 non-blocking, plan backfill]; deferral
audit GREEN; 4 documented residuals F5/F9/F11/F13 → standing backlog. ALL 4 v2.2 milestones now closed. Next:
/developer-kit:close-release.)_
