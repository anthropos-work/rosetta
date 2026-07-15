# Release Retro — v2.3 "cue to cue" (M217 → M221)

**Shipped 2026-07-15.** The presenter-speed release: *a presenter swaps heroes in under 5 seconds, on a demo
that comes up green, fully-loaded, and remotely reachable by default.* 5 milestones
(**M217 → { M218 ∥ M219 ∥ M220 } → M221**), tooling + docs only, **zero platform-repo edits**. Consolidates the
five milestone retros ([M217](m217-clean-stage/retro.md) · [M218](m218-seat-change/retro.md) ·
[M219](m219-readiness-renders/retro.md) · [M220](m220-cue-sheet/retro.md) · [M221](m221-prove-on-billion/retro.md)).

## What shipped

The trigger was a live presenter defect — *"I click a user, then it takes 1 or 2 minutes to access the
platform."* The investigation found three uncomfortable things at once: **Clerkenstein was innocent**, the walls
had **already been measured in this repo and nobody looked**, and the corpus **asserted the opposite in four
places** (*"~2–5 s, which we can't shorten"* — booked as M43-D5 with **zero deferrals recorded**, so it never
entered a ledger across four releases). v2.3's answer:

- **The headline gate — click→ACCESS < 5 s — set at M218 and re-proven live at M221**, 8/8 on `billion` over the
  tailnet with **no flags**: login p95 **maya-thriving 2.11 s / dan-manager 1.31 s** vs a **39.45 s / 38.30 s**
  baseline — a **~18×** collapse. ACCESS := the authenticated shell interactive with the hero's identity present
  (D-DESIGN-1); the heavy member grid is reported, never gated.
- **AI-readiness renders filled, on the current surfaces** (M219): the demo pointers had all aimed at an
  **unlinked legacy orphan**, and four sections read a table **no seeder ever wrote**.
- **The demo comes up full and remotely reachable by default** (M220): remote reach flipped **default-on for the
  demo path** (D-DESIGN-3, superseding v2.2's D-DESIGN-1), the "2 orgs" doc lie corrected to the true **3**, and
  `safety.md` gained **Part 3 — the exposure axis**.
- **Proven on the live box** (M221), not the model — and the `billion` demo is **left live** as the final
  deliverable.

## The D17 thread — the release's defining lesson

**D17: *a status artifact that outlives the thing it describes, and is then read as evidence.*** Named at the
M218 close. It is not one of v2.3's bugs — it is the **shape of almost all of them**. It bit **~24 times across
all five milestones** (M219's own running count reached 14 through M219; the release total is ~24), and the
single most important finding of this release is what it did *after* it was named:

> **Naming the class did not inoculate against it. It recurred inside the very passes convened to fix it.**

- **M218** wrote D17 *during* the milestone, then its close **found four more instances — one inside the
  hardening pass convened to name it, one in its own regression test** (a function-slice that swallowed the
  definition it was asserting on). "Nine incidents. One class."
- **M219** hit it **five more times — twice in its own `progress.md`, once inside its own fence, once in the
  harness convened to grade the milestone** (`run-coverage.sh` re-printed the *previous* run's numbers, "GATE:
  MET ✅" and all, and nearly graded a rebuild on the old broken stack's data). And **17 existing tests were
  guarding the bug** — asserting the junk-skill fallback *as the contract*, so any correct fix was guaranteed to
  fail them.
- **M220** was "the fences kept catching themselves": the exposure guard silently skipped its own repo; the
  ordering fence matched a flag inside its **own comment**; the S7 mutation battery graded a **featureless tree**;
  the HARD-INVARIANT fence asserted a **re-typed copy of itself** (mutating the shipped `SCHEME` to `https` left
  all 23 tests green). And the deepest: the dev-stack suite's *"environmental"* excuse **sat in this project's
  own headline numbers for four releases**, re-characterised by every milestone that touched it — **one missing
  env var underneath** (`DEV_SETDRESS_USE_STUB_BINS=1`).
- **M221** caught its own theatre at the close: `test_reap.py` printed **"Ran 21 tests … OK"** on a direct run
  while silently omitting 20 — a false all-clear on the very suite that fences the milestone's reap work.

**The keeper, stated once for the release:** ***a named hazard is not a fence; only an executable probe binds.***
Everywhere a fence could *execute* the thing rather than *grep* for it, it caught real bugs; everywhere it
described the thing, the description drifted from the code. The generalized sentences the milestones minted:

- *"the doc says N" ≠ "the code ships N"*
- *"it serves" ≠ "it renders" ≠ "the session survives"*
- *"the flag exists" ≠ "the flag works"* · *"the test is named X" ≠ "the test tests X"*
- *"it resolves" ≠ "it has skills"* · *"a pid exists" ≠ "the service is up"*
- *"the binary is installed" ≠ "it mints a cert"* · *"tailscaled handed me a string" ≠ "it is a hostname"*
- *an errored command is not "zero results"* · *a uniform result across unrelated mutations is a constant, not a result*
- ***"we deferred this before" is not a reason to defer it again — a re-characterised excuse is where a chronic bug hides.***

The corpus home for the hazard is [`corpus/ops/verification.md`](../../../../../corpus/ops/verification.md) ("THE
STALE-VERDICT HAZARD"), with two invariants: **a verdict must not outlive its subject**, and **absence must be
the safe state** (a grader with no verdict *refuses to measure*).

## Cross-milestone patterns

**1. Graded ≠ shipped — and v2.3 closed it.** M219 shipped its 5 greens on a **seed-path commit that landed
after the graded tag** (D13 restarts the count; *"the code that graded is not the code that shipped"*). M221's
answer was **fix-then-prove**: every off-box fence was RED- or mutation-proven *before* the live battery, so the
battery graded reality. The lesson from M219 was explicitly closed at M221.

**2. Prove it on the live box, not the model.** M221's F1 store-root shadow took **two live cycles**: the unit
fence modelled *existence*, the box surfaced *population* (an empty `snapshots/` subdir still shadowing the real
cache → `public.skills` back to 0), and Next 16's **detached** worker defeated the reap's parent-walk. *"The box
did what a model could not."* Harden then pinned a third depth as loud-not-silent. Direct-drive iteration on
shared live infra (the M215 "prove-on-odyssey" analogue) — speculative iteration doesn't pay there.

**3. The host-isolation lock prevented DATA corruption but not BOOKKEEPING corruption.** M219 was damaged by an
**orchestration** error — two batteries run concurrently against the single demo host, one purging the stack
mid-measurement (→ `GUARD-M221-host-isolation`). M221 landed the per-N host lock **first**, and it protected
iter-05's live cycle. But the same milestone then hit the *bookkeeping* version: the orchestrator prematurely
committed an iter-05 close (`766c029`) telling the **stale r3 story** while the agent was still re-proving at r4.
The host-lock isolated the shared *resource*; it did not isolate the shared *record*. Reconciled **loudly** in
`3c64af1`, never silently — the honest handling is the point.

**4. Adversarial execution over self-review; measurement over the plan.** M217: agents that *ran* the code found
20 bugs where self-review found 4 (three of them introduced by hardening pass 1, which had declared victory).
M219: two of the milestone's own premises (the `CycleID==nil` blocker, the "recompute never completes" claim)
were **refuted by looking**, and the planned demo-patch was **withdrawn** — *a patch that fixes a bug that does
not exist is indistinguishable, in a status report, from one that works.* M218: a claim that something is
unfixable is a claim that **stops anyone from looking** — that is exactly how the headline latency bug survived
four releases.

**5. Fences graded by mutation, not by re-running them green.** The release's standing rule — *a fence that
passes against both the pre- and post-fix code is theatre* — caught its own fences being exactly that (M219
instance 3; M220's H-1 self-copy; M220's 3-of-17 theatre mutants). M220's 17-mutant battery was itself first a
D17 (claimed-but-never-committed); once committed it found the milestone's HARD INVARIANT asserted against a copy
of itself.

## Carry-forwards → v2.4 (4 non-gate tail carries, user signed off 2026-07-15)

None is a gate condition of v2.3 (the 8/8 headline gate is independent of all four); each has a **structural**
reason it could not land in-release — a **platform-repo** edit or **live infra** — not "no time." Full
disposition: [`m221-prove-on-billion/audit-deferrals/deferral-audit-2026-07-15-m221-close.md`](m221-prove-on-billion/audit-deferrals/deferral-audit-2026-07-15-m221-close.md);
landing spot: `roadmap-vision.md` §v2.4; `RELEASE-SCOPE-DEFER` decisions in each originating milestone's
`decisions.md`.

- **F4** — academy grid renders 0 cards (catalog serves 2,705). The fix is a render-path change in the
  **`ant-academy` platform repo** → out of the zero-platform-edit scope by construction. Documented cosmetic gap.
- **BURNIN-M221-dev-public-host** — the dev-path `--public-host` (built M220) is fenced + mutation-proven
  byte-identical but never live-burned. Needs live infra.
- **F-M220-4** — `ant-academy.sh` re-runnable on a live public-host demo (serve↔bind contention). Needs live infra.
- **PROBE-M218-c3-rerun** — the Cosmo cms/Directus 403 re-check on the content path. Needs the live box.

*(The genuine release chronic — the dev-stack suite across 5 milestones and 2 releases — was NOT carried: it was
investigated-not-re-deferred and **LAND-NOW'd at M220**. That is the pattern the `audit-deferrals` gate exists to
catch, and it was caught one milestone before the close.)*

## Metrics delta

| | v2.2 (like-for-like) | **v2.3** | Δ |
|---|---|---|---|
| **p95 click→ACCESS** | *(unmeasured; corpus claimed "~2–5 s, unfixable")* | **2.11 s / 1.31 s** vs < 5 s, live on `billion` | ~18× vs the 39.45/38.30 s baseline |
| Go test funcs | 1749 | **1831** | +82 (same `grep -c '^func Test'` method) |
| Python tests | 668 (3 sections) | **1341** (5 sections, JUnit collected, 0 fail, 16 skip) | +673; shared-3 668→1105; **whole tree completes for the first time (M220)** |
| TS e2e | 124 | **151** (69 + 82, `playwright test --list`) | +27 (M219's 94 was a different measure — no false regression) |
| Alignment (Clerkenstein Go) | 100/100 | **100% / 100% critical (27/27)** | held — M218's hollow 100 made honest (97.2), M219 fixed the RED gene + retained it as a fence |
| Flake | 0 | **0** | triple-clean 3/3 |
| Net-new direct deps | 0 | **0** | only `x/crypto v0.51→v0.52` (indirect, M218) |
| Platform-repo edits | 0 | **0** | rosetta diff 108 files, all docs/planning |

## The one line worth keeping

**A named hazard is not a fence — only an executable probe binds.** v2.3 named D17 at M218 and then walked into
it ~24 more times, including inside the passes built to stop it. What consistently held was not the *name* but the
*probe*: the fence that ran the emitters, the grader that deleted its report first, the battery graded by
mutation, and — above all — **the cold reset-to-seed on the live box that no model could stand in for.**
