# Release Review: v2.2 "panorama"

**Date:** 2026-07-12
**Milestones:** M212 (public-host knob) · M213 (auth over the tailnet, TLS/proxy/pk) · M214 (origins & links, CORS + patch tail) · M215 (prove-on-odyssey — the first live remote Linux-VM deploy, iterative closed-on-gate)
**Method:** 7-agent parallel close-release workflow (`wf_bc0d8656`), two-repo (rosetta docs/plan @ `release/02.20-panorama` + rext code @ `panorama-m215`=`00ba6b6`, rolled to `v2.2`=`39e8013` at close).

**Headline:** clean release — the **external-shareability** capability shipped **and proven live on a real Linux VM over Tailscale** (both hero vantages, trusted LE cert, cold reset-to-seed reproducible). **Tooling + docs + an opt-in flag only — zero platform-repo edits, 0 net-new deps** across all four milestones. Metrics + deferral gates GREEN + verified; supply-chain YELLOW-non-blocking (pre-existing, unreachable). **No must-fix blocker.** All findings were should-fix docs (landed this close) or nice-to-have.

## Scope (Phase 1/1b) — GREEN
- [x] All 4 roadmap promises delivered (opt-in host knob / `tailscale cert` HTTPS / CORS+patch tail / live remote-deploy proof); all 4 `--no-ff` merges present on `release/02.20-panorama` (`52a3664`→`3530dd0`→`82dd3dc`→`2fcab2c`). All 5 roadmap top-risks retired (secure-context, mixed-content, RAM-on-billion, build-bake tax, patch-tail drift).
- [x] **0 Fate-3-undelivered**; M214 discharged the two items M212/M213 routed into it (DEF-M212-01 CORS emission, DEF-M213-01 recipe+links) as Fate-1 landings.
- [x] **Nothing dropped, nothing unaccounted.** **M216** (dev-path parity + operator surface) is the release's **declared optional / scope-flex lever** — never scaffolded, stays roadmap-only until promoted; a conscious optional, not a drop.
- [x] **M215 propagation close-gate SATISFIED** — every deployment finding (F1/F2/F4/F6/F8/F9/F12) landed in tools(rext)+KB+skills; `corpus/ops/demo/tailscale-serve.md` stands a fresh Linux VM up unaided.

## Supply chain (Phase 0) — YELLOW (no blocker)
- [ ] [should-fix, next rext roll] `cd clerkenstein && go get golang.org/x/crypto@v0.52.0 && go mod tidy` — clears **13 open dependabot alerts** (7 crit / 2 high / 4 med), **all** the same indirect `golang.org/x/crypto v0.51.0` in the `ssh`(+`ssh/agent`) subpackage. **govulncheck call-graph confirms all 13 UNREACHABLE** — clerkenstein is a Clerk mock that pulls x/crypto only via go-jose (JWT), never imports ssh; internal-only tooling, not distributed, not an SSH server. Published against the already-pinned v0.51.0 after v2.1 shipped; v2.2 added no dep and didn't touch it. Not RED (0 reachability + internal-only).
- [x] **0 reachable vulns** (govulncheck across all 6 Go modules on go1.25.12 — cleaner than v2.1, the toolchain advance retired v2.1's reachable ECH leak); npm 0 vulns (2 e2e devDeps, Playwright 1.61.1); Python 0 third-party deps (all stdlib/local); 0 GPL/LGPL/AGPL (only MPL-2.0 file-level internal `golang-lru/v2`). `dependencies.lock` written.

## Code Quality (Phase 2) — GREEN
- [x] M215 rext diff (`41a28aa..panorama-m215`, 3 commits) clean, correctly guarded — **every new behavior gated Linux-only / public-host-only / missing-prereq-only, so the macOS/dev path stays byte-identical**; non-fatal where required; no dead code; no resource leaks. Spot-checks passed: F4 `trap reap_ssh_agent EXIT` is the sole EXIT-trap owner; `$OFFSET` semantics consistent across serve + F12 pre-reset; F8 atlas-error classifier benign-safe (`_have_svc`-gated); `gen_tailscale_serve.py --reset` mirrors the serve emitter.
- [x] [RESOLVED this close] **ADV-1 (F12)** — the up-path defensive pre-reset sits after `compose up` (can't pre-empt a leftover-listener bind conflict; the **teardown** reset is the primary guard, + the runbook's manual `tailscale serve reset` for the by-hand case). Landed as a **5-line comment-only** reconcile in `up-injected.sh` (rext `39e8013`) — verified NO logic change; demo-stack pytest re-run 424 passed, shellcheck clean.

## Documentation (Phase 3/3b) — was SHOULD-FIX, all landed this close
- [x] **`tailscale-serve.md` finding-range contradiction** (L17 "F1–F12" vs See-also "F1–F13"; table omits F13; `propagation-checklist.md:27` mislabelled) → standardized: the runbook table = the F1–F12 **host/deploy** set; F13 (jobsimulation crash) explicitly flagged **out-of-scope** with a numbering note (F10 unused); See-also + propagation-checklist reconciled.
- [x] **Cross-ref §-anchor label truncation** (2 files) → `rosetta_demo.md` + `clerkenstein.md` now match the full heading "Linux host prerequisites (for a remote/VM demo **over Tailscale**)".
- [x] **Top-level index omission** → `tailscale-serve.md` added to the `CLAUDE.md` "Demo Environments" list **and** to the `corpus/ops/README.md` "Available Operations" table (was reachable only 2 hops in).
- [x] [nice-to-have, landed] Stale "tracked for M215" in `tailscale-serve.md` reworded to "an accepted future enhancement (cockpit ships deliberately plain-HTTP)".
- [x] Plan records (`state.md`, `roadmap.md`, per-milestone `overview.md`) reflect the correct pre-close state; the SHIPPED flip / archive / `v2.2` tag / `rext.tag` bump are Phase-10 close work (below).

## Tests & Benchmarks (Phase 4/4b) — GREEN
- [x] Go **1772** test funcs (+8 vs v2.1's 1764, all in clerkenstein: clerkjs_proxy +7 M213 + handshake +1 — reconciled exactly), `go test ./...` exit 0 + `go vet` clean on all 6 modules. Python **668** passed (demo-stack 424 / stack-injection 147 +8 skip / stack-core 97). TS e2e **124** (playwright `--list`; tree byte-identical to v2.1 — the number difference vs v2.1's "103" is a count-method difference, not a code change). 10 live Playthroughs + 1 in-manifest TODO, re-proven live over Tailscale at M215.
- [x] **All 4 regression rules PASS:** no test-count decrease (Go +8, TS flat, Python monotonic-up), no coverage drop >2pp, **flake 0** (double-run identical + milestone triple-clean), 0 net-new deps, 0 platform-repo edits. `metrics.json` written.

## Decision Consolidation (Phase 5) — was YELLOW hygiene, landed this close
- [x] **`D-SCHEME-1` ID reused for two distinct decisions** — M213's (HTTPS-everywhere access-scheme *policy*) and M214's (`browser_scheme` http/https *predicate* implementing it). Substantively consistent + individually milestone-qualified in citations, but `tailscale-serve.md` read as a dup → disambiguated to `M213-D-SCHEME-1` / `M214-D-SCHEME-1` in the corpus cross-refs.
- [x] **No semantic cross-milestone conflict.** The two chained decisions that could collide are explicitly reconciled: D-PK-1 "REFINES" D-IMPL-1 (codec stays permissive, validation moves up); M214 D-SCHEME-1 flips only the terminal emission of M212's wired-but-unemitted `gen_injected_override` seam, byte-identical in between. All decisions blended into skill/corpus docs or deliberately archived; no orphans.
- [x] M215's stub `progress.md`/`decisions.md` (direct-drive milestone; canonical record in `iter-01/findings.md`) were backfilled at close — a recorded process lesson, not an open gap.

## Resolution (Phase 7 — all findings landed)
**Corpus/planning (rosetta, this close):** `tailscale-serve.md` finding-range reconcile (F1–F12 scope + F10/F13 numbering note) + `D-SCHEME-1`→`M213/M214-D-SCHEME-1` disambiguation + "tracked for M215" reword · §-anchor "over Tailscale" fix in `rosetta_demo.md` + `clerkenstein.md` · `propagation-checklist.md:27` range reconcile · `tailscale-serve.md` indexed into `CLAUDE.md` + `corpus/ops/README.md`.
**Rext (`panorama-m215` → `39e8013`, tagged `v2.2`=`11fa0a62`):** D-CLOSE-1 (demo-stack README test-count 50→128, measured) · D-CLOSE-2 (stack-injection README `gen_tailscale_serve.py` +`--reset` row) · D-CLOSE-3 (`apply-ant-academy-dev-origins.sh` + patches README index) · ADV-1 (F12 up-path pre-reset 5-line clarifying comment, no logic change). Re-verified: demo-stack pytest 424 passed, shellcheck clean.
**Verification:** both trees clean; triple-clean demo-stack suite 3/3 (Phase 8c); Go 6/6 green + `go vet` clean; govulncheck 0 reachable.
**Deferred (documented, non-gate-blocking standing backlog):** F5/DEF-M215-01 (app demopatch sha re-anchor) · F9/DEF-M215-02 (remote-VM snapshot-cache pre-stage/auto-sync) · F11/DEF-M215-03 (seed hero-name cosmetic + committed remote-origin Playwright gate) · F13/DEF-M215-04 (jobsimulation exits(1)) — all orthogonal to shareability, none on the proven journey path. Plus the next-rext-roll `x/crypto@v0.52.0` bump (Phase 0).

## Verdict — ✅ GO
Merge `release/02.20-panorama` → `main`, tag **v2.2**. All three blocking gates pass (metrics GREEN, deferral GREEN, supply-chain YELLOW-non-blocking); full roadmap promise delivered; scope ledger clean (0 escape-hatch, 0 dropped, 0 unaccounted); rext already re-tagged `v2.2` with the D-CLOSE trio + ADV-1 discharged and verified green.
