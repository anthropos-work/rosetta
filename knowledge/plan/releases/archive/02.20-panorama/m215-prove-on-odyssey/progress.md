# M215 — progress

> **Direct-drive iterative milestone.** M215 is live shared-infra work (a real Linux VM over Tailscale), so it
> was driven directly rather than via background sub-agents. The canonical iter record is
> [`iter-01/findings.md`](iter-01/findings.md) (tik-1 → tik-2 → both-vantages → cold-reset capstone, findings
> F1–F13). This ledger summarizes; `decisions.md` records the strategy + residual fates.

## Running ledger

- **iter-01 (2026-07-11) — bootstrap tok + tiks, exit gate MET for the core.**
  - **tik-1 (trimmed de-risk):** backend + Clerkenstein only (`DEMO_NO_UI=1 DEMO_NO_LOCAL_CONTENT=1`) on billion with
    `--public-host`. Surfaced + fixed the host-prereq/rext set: F1 (tailscale operator), F2 (host Go), F3 (`git tag
    | head` SIGPIPE→141 — REXT FIX), F4 (keyless ssh-agent for buildx `ssh: default` — REXT FIX), F5 (app demopatch
    sha-drift, non-fatal note), F6 (Linux bind-mount perms — REXT FIX), F7/F8 (host atlas; migrate fail-loud — REXT
    FIX). Result: UP + healthy + migrated (91 tables) + seeded (201 users/1 org); a **remote Mac reached it over
    Tailscale with a trusted LE cert** (backend/router/FAPI all `verify=0`). F9 noted (empty taxonomy — no snapshot
    cache on VM).
  - **tik-2 (real remote login, employee):** added next-web + studio-desk (RAM held, 5.7 GB free). A headless
    Chromium on a **different** tailnet machine drove the full cockpit→FAPI-handshake→`/profile` login as
    `maya-thriving` — landed authenticated, profile fully rendered, `ignoreHTTPSErrors:false`, 0 console errors, 0
    functional request failures. **Employee-vantage exit-gate proof.** F11 noted (seed hero-name cosmetic).
  - **manager vantage:** logged in as `dan-manager` → `/enterprise/workforce` ("Workforce Intelligence"), fully
    rendered with real seeded structural data, 0 console errors. **Both vantages proven.**
  - **cold reset-to-seed capstone:** synced the fixed rext (`panorama-m215`) to billion, ran a **wiped-DB one-shot**
    `--public-host` bring-up. The three auto-fixes fired as designed (pre-flight OK / F4 ssh-agent / F6 data-dirs, no
    manual steps). Surfaced + fixed **F12** (teardown didn't reset `tailscale serve` → re-deploy port-conflict — REXT
    FIX: per-port serve-reset on teardown + defensive up-path pre-reset). Re-ran clean: 14 containers, serve fronting
    5 ports, seed 12,245 rows / 541 users / 9-hero roster, `/api/health 200`, login-ready from the Mac. **"Cold
    reset-to-seed reproducibility" gate item MET.** F13 noted (jobsimulation container exits(1) — separate
    demo-service fix, off the proven path).
  - **Gate delta:** core exit gate **MET** — both vantages proven live over Tailscale (trusted cert, 0 ejects,
    assets rendering) + cold reset-to-seed reproducible one-shot + unset knob byte-identical. Residuals F5/F9/F11/F13
    documented + routed (non-gate-blocking). rext code-of-record FROZEN at tag `panorama-m215` @ `00ba6b6`.
  - **propagation close-gate satisfied:** every deployment finding (F1/F2/F4/F6/F8/F9/F12) landed in tools (rext) +
    KB + skills; a fresh reader can stand up a remote demo unaided from `corpus/ops/demo/tailscale-serve.md`. See
    [`propagation-checklist.md`](propagation-checklist.md).

## M215: Final Review (close)

### Scope
- [x] Gate-distance (iterative): core exit gate MET; residuals F5/F9/F11/F13 documented + Fate-2 routed (carry-forward.md).
- [x] Iter-ledger audit: iter-01 has findings.md (the canonical record); running ledger + decisions.md backfilled at close.
- [x] Inherited Fate-2→M215 (DEF-M213-02/03/04) all landed Fate-1 / resolved-documented (deferral audit GREEN).
- [x] No stray TODO/FIXME/HACK in the M215-touched rosetta source.

### Code Quality
- [x] [rext, frozen] diff `41a28aa..00ba6b6` reviewed cross-cuttingly — GREEN. One scheme/guard pattern; macOS/dev path byte-identical (Linux-only / missing-prereq-only / public-host-only guards); no dead code, no resource leak (keyless agent reaped). Not re-tagged (close-release's job).
- [x] [nice-to-have] ADV-1 recorded: F12 up-path defensive pre-reset sits after compose-up — teardown reset (primary) covers the normal cycle; up-path pre-reset is defensive-only. Reconciliation routed to the rext close-release re-tag (rext frozen).

### Documentation
- [x] [should-fix] tailscale-serve.md "findings F1–F11" → F1–F12 (F12 teardown-reset is baked into the runbook); the file-pointer → F1–F13 (full finding set).
- [x] [nice-to-have] tailscale-serve.md §-anchor to setup_guide corrected to the full heading "…(for a remote/VM demo over Tailscale)".
- [x] All 9 cross-references resolve; setup_guide + clerkenstein anchors present; skills (demo-up/stack-secrets/dev-up) carry the Linux/remote content.

### Tests & Benchmarks
- [x] rext demo-stack **424 passed** (independently re-run), stack-injection **147p/8s**; shellcheck clean — matches the propagation-checklist + findings claims. (No rosetta-side tests — docs+plan deliverable.)

### Decision Triage
- [x] Strategy + residual fates recorded in `decisions.md` (D-STRAT-1, DEF-M215-01..04, ADV-1); the mechanism/why-it-works already blended into `tailscale-serve.md` + `clerkenstein.md` (M213/M214). No net-new knowledge blend needed at M215 close.
