# Playthroughs — Next / Out-of-scope (parking lot)

> **Status:** Draft · spec-draft · 2026-06-28
> Companion to [`spec.md`](spec.md). Anything **out of scope for defining + first-building the Playthroughs
> capability** lands here: deliberate **non-goals** (different test classes) and **later** work.

## Purpose

Two kinds of thing park here:

1. **Non-goals** — test classes that are explicitly **not** Playthroughs, so the capability stays focused on
   *functional-flow truth* ([`spec.md`](spec.md) §7).
2. **Later** — work we agree is "after we've defined the capability," so the spec stays about *what it is*, not
   *what to test first*.

When something is genuinely Playthrough scope, promote it into [`spec.md`](spec.md) — don't leave it here.

## How to use

- One row per item. Keep it short.
- **Kind** = `non-goal` (a different test class, never Playthroughs) or `later` (Playthrough scope, just not now).

## Parked items

| # | Item | Kind | Why parked |
|---|------|------|-----------|
| 1 | **The actual products / stories / use-cases / tests** — declaring the manifest's real content and implementing the Playthroughs | later | The founding directive: *define the capability + principles + tech now; do **not** list/build tests*. This is the build phase, after the spec is agreed (ideally turned into a versioned plan via `/developer-kit:design-roadmap`). |
| 2 | **Visual-regression / pixel-diff testing** | non-goal | The **opposite** of P2 — Playthroughs deliberately do not care about pixels/copy. A separate capability if ever wanted. |
| 3 | **Performance / load / stress testing** | non-goal | Playthroughs assert *function*, not latency/throughput/capacity. |
| 4 | **Unit / integration testing** of platform code | non-goal | A platform-repo concern, below the user-facing flow; also outside Rosetta's zero-platform-edit boundary. |
| 5 | **API / contract testing** (wire-level) | non-goal | Playthroughs verify at the **user surface**, not the API. |
| 6 | **Security / penetration testing** + **accessibility *auditing*** | non-goal | Distinct disciplines. (Playthroughs *use* the a11y tree to locate; they don't *audit* it.) |
| 7 | **Cross-browser / device matrix** | later | Start on one browser engine for determinism; broaden the matrix once the core suite + principles are proven. *(Distinct from the mobile blocker — row 11.)* |
| 8 | **CI integration + scheduled regression runs** | later | The suite is a regression reference; wiring it into a gate/cron is a follow-on once it's stable and the stack-binding (spec-progress #5) is settled. |
| 9 | **Multi-actor / concurrent-flow Playthroughs** (e.g. a manager assigns → an employee completes, in one story) | later | Stories can chain across actors; orchestrating two heroes inside one Playthrough is an extension once single-actor flows are solid. *(Distinct from §5.7's N-separate-Playthroughs concurrency, which is decided.)* |
| 10 | **A "semantic landmark" convention** (blessed already-exposed anchors), if pure-semantic proves too ambiguous | later | Already largely the decided policy (the §5.2 registry, load-bearing on the antd-v6 UI); kept here only as the place to formalize a *broader* convention if surfaces force it — and only without any platform edit. The **AI / media-widget tier is the known forcing case**, not speculative (§5.2). |
| 11 | **Mobile-app (Expo) flows** | later | Not a smooth "broaden the matrix" follow-on: the `@anthropos/mobile` app is **paused / workspace-excluded / not stack-served**, *and* Playwright (the locked tool) **cannot drive a native Expo app** (that needs Appium/Detox/Maestro, precluded by "extend the M42 foundation"). Only Expo's **web target** is Playwright-reachable. Revisit if/when the app is un-paused. |
| 12 | **Mirror engines for integration-dependent legs** — voice (LiveKit), recording (Chime), payments (Stripe test-mode), email (Brevo sink), and a deterministic / mocked-AI sim-completion path | later — needs a mirror engine | Clerkenstein mocks **only Clerk**; these legs have no mirror, so a live sim is asserted only at the **launch/completion boundary** (§5.8). A full in-widget Playthrough needs a per-engine mirror first — significant, deferred build work. |
| 13 | **Language-switch / fallback flow class** (drive the UI to switch locale; verify English fallback for untranslated content) | later | The test *locale* is pinned to English as a known-state axis (§5.4); proving the *language-switching behaviour itself* is a distinct flow class, parked until the core suite is proven. |

> Add new rows above. Promote an item into scope only by moving it into [`spec.md`](spec.md) / a spec section.
