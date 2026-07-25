# M253 — Decisions

_(Implementation decisions with rationale, D-numbered, recorded during build.)_

## TOK-01: shell-before-awaits + no-thirdparty demopatches on the M249 ladder — 2026-07-24

**Tok type:** bootstrap (iter-01)
**Initial strategy:** Cut the studio-desk first-paint blank by painting the `.page-skeleton` shell (dark
header + sidemenu) **synchronously, before the boot awaits**, instead of after them. Concretely, two sha-pinned
demopatches on the M249 `build_frontend_studio_desk` ladder:
  1. **`studio-desk-shell-first-paint`** — insert the `.page-skeleton` DOM (header + sidemenu + content [+
     footer]) synchronously right after `preloadCriticalCSS()` (main.ts ~L97), BEFORE Sentry.init / posthog.init
     / `clerk.load()` / `l12nService.init()` / `userService.canAccess()`. The CSS for those classes is ALREADY
     injected by `preloadCriticalCSS`, so the shell paints from CSS+DOM with zero network. De-dup is automatic:
     `PageWrapper#init` wipes `document.body.innerHTML` (pageWrapper.js L122) then rebuilds its own skeleton, so
     the early shell is seamlessly replaced — no double skeleton.
  2. **`studio-desk-no-thirdparty`** — no-op `posthog.init` (main.ts L117) + `Sentry.init` (L104) on the demo
     (they fire remote calls on a non-localhost host and serve no purpose on a Clerk-free demo; removing their
     synchronous setup also tightens the paint that follows).
Extend the M249 patch ladder + patch-set fingerprint in `up-injected.sh build_frontend_studio_desk` (add both
manifests so the fingerprint forces a studio rebuild). Author a net-new studio-FCP runner in
`stack-verify/e2e/` (run-latency measures next-web/hiring ACCESS, not studio FCP). Iterate on demo-2 (local
laptop); billion untouched (cold-p95 confirmation is M254).
**Rationale:** the baseline timeline proves the blank is the sequential awaits BEFORE the shell is built, not
the shell build itself. clerk.load was 140 ms (not its 10 s timeout); the ~3.9 s cost is inside canAccess. So
the correct, minimal, zero-platform-edit fix is a paint-ordering change (demopatch), independent of the
canAccess 404. This is the additive-DOM injection pattern M249 authored, applied to a new anchor.
**Strategy class:** new-direction
**Distance-to-gate context:** gate = first-meaningful-paint (`.page-skeleton` header+sidemenu visible) < 1000 ms
+ no blank > 1 s, p95 over 5 consecutive cold loads, on a fresh-green autoverify.json. Baseline (demo-2, local
laptop): skeleton-visible **4669 ms** (single confirming cold load). Need to cut ~3.7 s by moving the paint
ahead of the awaits.
**Next-tik direction:** iter-02 — author both demopatch manifests + the FCP runner, extend the ladder +
fingerprint, rebuild the studio image on demo-2 with the patches baked, and measure 5-cold-load p95 FCP.

## Close-time decisions (2026-07-24)

### D-close-1 — FCP runner secondary `browser FCP` field reads `n/a` → Fate-3 (future harden, not blocking)
The `run-studio-fcp.sh` harness gates on the PRIMARY metric — skeleton-visible **wall-clock** time (works,
5/5 cold, p95 817 ms). Its secondary `browser FCP` field (chromium First-Contentful-Paint) reads `n/a`
because paint-timing isn't populated in the measure path. This is a nice-to-have hardening candidate
(a `PerformanceObserver` in the page-context probe), **not** a gate defect — the primary metric is
authoritative and passes. **Fate-3 / future harden**; owed to no current milestone. Recorded so the note
survives the close.

### D-close-2 — decision triage: all M253 decisions → archive (maintainer-only)
iter-01 D1 (inline KB-fidelity verdict) · D2 (dominant-await = canAccess) · iter-02 D3 (chained-manifest
sha generator) · D4 (lib-only rebuild vehicle) · D5 (green-gate non-achievable on warm demo-2 → M254) ·
TOK-01 (bootstrap strategy) → **archive**. Their load-bearing platform facts (the per-leg boot baseline,
the paint-reorder fix, the chained patch pair, the MPA/empty-body boot model) were already blended into the
3 corpus doc Delivers during iter-03 (`latency-budget.md` / `demopatch-spec.md` / `studio-desk.md`) — verified
accurate + non-duplicated at close.

### D-close-3 — close docs fix: demopatch-spec.md §2.1 stale count 21 → 23
M253's canonical §5 inventory reconcile (21 → 23) left an earlier illustrative present-tense count in §2.1
("R1 sweep iterates … all 21 today") stale. Fixed to "all 23 today". Historical counts (14 / "the other 11" /
"swept 14 manifest(s)") are dated evidence and correctly unchanged.
