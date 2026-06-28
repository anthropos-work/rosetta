# Run-throughs — Next / Out-of-scope (parking lot)

> **Status:** Draft · spec-draft · 2026-06-28
> Companion to [`spec.md`](spec.md). Anything **out of scope for defining + first-building the Run-throughs
> capability** lands here: deliberate **non-goals** (different test classes) and **later** work.

## Purpose

Two kinds of thing park here:

1. **Non-goals** — test classes that are explicitly **not** Run-throughs, so the capability stays focused on
   *functional-flow truth* ([`spec.md`](spec.md) §7).
2. **Later** — work we agree is "after we've defined the capability," so the spec stays about *what it is*, not
   *what to test first*.

When something is genuinely Run-through scope, promote it into [`spec.md`](spec.md) — don't leave it here.

## How to use

- One row per item. Keep it short.
- **Kind** = `non-goal` (a different test class, never Run-throughs) or `later` (Run-through scope, just not now).

## Parked items

| # | Item | Kind | Why parked |
|---|------|------|-----------|
| 1 | **The actual products / stories / use-cases / tests** — declaring the manifest's real content and implementing the Run-throughs | later | The founding directive: *define the capability + principles + tech now; do **not** list/build tests*. This is the build phase, after the spec is agreed (ideally turned into a versioned plan via `/developer-kit:design-roadmap`). |
| 2 | **Visual-regression / pixel-diff testing** | non-goal | The **opposite** of P2 — Run-throughs deliberately do not care about pixels/copy. A separate capability if ever wanted. |
| 3 | **Performance / load / stress testing** | non-goal | Run-throughs assert *function*, not latency/throughput/capacity. |
| 4 | **Unit / integration testing** of platform code | non-goal | A platform-repo concern, below the user-facing flow; also outside Rosetta's zero-platform-edit boundary. |
| 5 | **API / contract testing** (wire-level) | non-goal | Run-throughs verify at the **user surface**, not the API. |
| 6 | **Security / penetration testing** + **accessibility *auditing*** | non-goal | Distinct disciplines. (Run-throughs *use* the a11y tree to locate; they don't *audit* it.) |
| 7 | **Cross-browser / device matrix** + **mobile-app (Expo) flows** | later | Start on one browser engine for determinism; broaden the matrix and add mobile once the core suite + principles are proven. |
| 8 | **CI integration + scheduled regression runs** | later | The suite is a regression reference; wiring it into a gate/cron is a follow-on once it's stable and the stack-binding (spec-progress #5) is settled. |
| 9 | **Multi-actor / concurrent-flow Run-throughs** (e.g. a manager assigns → an employee completes, in one story) | later | Stories can chain across actors; orchestrating two heroes inside one Run-through is an extension once single-actor flows are solid. |
| 10 | **A "semantic landmark" convention** (blessed already-exposed anchors), if pure-semantic proves too ambiguous | later | Default is pure-semantic (spec-progress #3); only promote a landmark convention into the spec if real surfaces force it — and only without any platform edit. |

> Add new rows above. Promote an item into scope only by moving it into [`spec.md`](spec.md) / a spec section.
