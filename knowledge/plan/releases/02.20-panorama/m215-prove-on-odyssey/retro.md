# M215 — retro (prove it on odyssey)

## Summary
The v2.2 "panorama" **iterative acceptance closer** — and the **FINAL** milestone of the release. It executed the
external-shareability surface **live** for the first time: a demo brought up on the odyssey `billion.taildc510.ts.net`
Linux VM with `--public-host`, driven **end-to-end from a DIFFERENT tailnet machine** (a remote Mac). Both vantages
were proven — employee (`maya-thriving` → `/profile`, the M41 ProfileSeeder depth rendered) and manager
(`dan-manager` → `/enterprise/workforce`, real seeded structural data) — over a **genuinely trusted** Let's Encrypt
cert (`ignoreHTTPSErrors:false`, `verify=0`, no per-machine CA install), with **0 console errors, 0 functional
request failures, 0 localhost/prod ejects**, assets rendering. A clean **cold reset-to-seed** one-shot bring-up with
the fixed tooling then proved reproducibility. The **core exit gate is MET**. The milestone was driven directly
(live shared-infra work, not background sub-agents — D-STRAT-1); the canonical run record is `iter-01/findings.md`
(findings F1–F13). Zero platform-repo edits throughout.

The user-directed **propagation close-gate** was satisfied: every deployment finding (F1/F2/F4/F6/F8/F9/F12) landed
in **tools (rext) + KB + skills**, and a fresh reader can stand up a remote demo on a new Linux VM unaided from the
NEW `corpus/ops/demo/tailscale-serve.md` runbook alone.

## Incidents This Cycle
No P1/P2 in the shipped work; no flakes (the frozen rext suite re-runs identical — demo-stack 424p, stack-injection
147p/8s). The milestone's **value was the live findings themselves** — the run surfaced a real host-prereq + rext-bug
set (F1–F13) that a macOS-only path never exposes:
- **F3** (rext bug) — `git tag --list | head -1` SIGPIPE→141 under `set -o pipefail` + `set -e` aborted the bring-up
  on a many-tag repo (`app` ~337 v-tags). Fixed to a pipe-less `git for-each-ref --count=1` + a ~400-tag regression
  fixture. A genuine latent bug caught only by a full-clone Linux run.
- **F12** (rext gap) — the demo teardown didn't reset `tailscale serve`, so a re-deploy port-conflicted
  (`address already in use`). Fixed: per-port serve-reset on teardown + a defensive up-path pre-reset. Surfaced on
  the cold reset-to-seed capstone (exactly the reproducibility gate item doing its job).
- F5 (app demopatch sha-drift, non-fatal), F9 (empty taxonomy — no snapshot cache on VM), F11 (seed hero-name
  cosmetic), F13 (jobsimulation exits(1)) — all documented residuals, none on the proven journey path.

## What Went Well
- **The make-or-break secure-context bet paid off live.** M213's HTTPS-everywhere + `tailscale cert` design meant a
  **remote** browser got a genuinely trusted cert with **no CA install** — the thing `mkcert` fundamentally cannot do
  for a second machine. `ignoreHTTPSErrors:false` held; Clerk's Web-Crypto secure-context requirement was satisfied.
- **The config-only core survived contact with reality.** No new platform edit was invented during the live loop —
  every breakage routed to the M212 knob / M213 TLS / M214 patch surface (or a host-prereq), exactly as the "config
  + a 2-item patch tail" feasibility claimed. The two shipped demopatches carried the MagicDNS value cleanly.
- **The auto-fixes fired first-try on the cold run.** On the wiped-DB one-shot, the three host auto-fixes
  (pre-flight OK / keyless ssh-agent / pre-created data dirs) fired as designed with **no manual steps** — the
  propagation work turned one-off manual unblocks into baked tooling.
- **The propagation close-gate forced durable coverage.** Because this was the FIRST Linux + remote deployment, the
  user-directed gate meant every finding landed in all three surfaces (rext + KB + skills), not just a fix note — so
  the next Linux/remote build is covered. The runbook is followable end-to-end (verified).
- **RAM fit was proven, not assumed.** The ballooned VM concern (risk 3) was burned down empirically — 16 GB swap
  added, ~5.7 GB free held with next-web + studio-desk up.

## What Didn't
- **The direct-drive milestone left its plan artifacts as stubs.** Because M215 was driven live (not via the standard
  `build-mstone-iters` loop), `progress.md` (running ledger) and `decisions.md` were empty at close — the real record
  lived only in `iter-01/findings.md`. Both were backfilled at close (running ledger + D-STRAT-1 + residual fates +
  ADV-1). Lesson: a direct-drive iterative milestone should still land a one-line ledger entry + the strategy
  decision as it goes, so the close isn't reconstructing them.
- **Content surfaces need the snapshot cache pre-staged on a remote VM (F9).** Identity/profile/dashboard/workforce
  render fully, but taxonomy/library are sparse without `.agentspace/snapshots` on the box. Operational, not a code
  defect — documented in the runbook (Step 4); an auto-sync helper is a future enhancement.
- **The F12 up-path defensive pre-reset placement is imperfect (ADV-1).** It sits after compose-up, so it can't
  prevent a backend-bind conflict from a leftover serve listener; the teardown reset (primary) + the documented
  manual `tailscale serve reset` cover the real cases. rext is frozen, so a one-line comment/placement reconciliation
  routes to the close-release rext re-tag.
- **python3.14 tooling friction recurred (non-blocking).** The box default `python3` (3.14) has no usable pytest; the
  suites run under `python3.12`, as at M212/M213/M214. Noted again.

## Carried Forward
All Fate-2, non-gate-blocking; full routing in [`carry-forward.md`](carry-forward.md); audit GREEN in
[`audit-deferrals/deferral-audit-2026-07-11.md`](audit-deferrals/deferral-audit-2026-07-11.md).
- **F9** — remote VM taxonomy/library snapshot-cache pre-stage + a future cache-auto-sync → **standing backlog**.
- **F5** — the two `app` demopatches sha-drift re-anchor → **standing backlog** (demo-service maintenance; rext frozen).
- **F13** — jobsimulation service-command crash → **standing backlog** (separate demo-service fix; not remote/Linux-specific).
- **F11** — seed hero-name cosmetic + a committed repeatable remote-origin Playwright gate → **standing backlog**.
- **D-CLOSE-1/-2/-3** (three rext-frozen README index reconciles) + the **ADV-1** F12 comment/placement one-liner →
  **v2.2 `/developer-kit:close-release`** (the single rext commit when rext re-tags + `.agentspace/rext.tag` bumps).

## Metrics Delta
- **Tests (rext, frozen at `panorama-m215` @ `00ba6b6`):** demo-stack 383 (M214) → **424** (+41: +25 host-prereqs
  incl. the ~400-tag F3 regression fixture, +16 F12 serve-reset); stack-injection unchanged at **155** (147p/8s).
  Independently re-run at close (`python3.12 -m pytest`), exit 0, matching claims. M215 touched **zero Go** (rext Go
  stays 1771) and **zero TS**. shellcheck clean (0.11.0, `-S style`).
- **Gate:** `closed-on-gate` — core exit gate MET (both vantages live + cold reset-to-seed reproducible + unset knob
  byte-identical). 4 documented non-blocking residuals routed.
- **Flake:** 0. **Supply-chain:** 0 net-new deps. **Platform-repo edits:** 0. **rext close edits:** 0 (frozen tag).
- Full machine-readable record: [`metrics.json`](metrics.json).

---
**M215 is the FINAL v2.2 milestone.** Next: `/developer-kit:close-release` — rolls the rext `v2.2` tag, bumps
`.agentspace/rext.tag` from `v2.1`, reconciles D-CLOSE-1/-2/-3 + ADV-1, merges `release/02.20-panorama` → `main`,
tags `v2.2`.
