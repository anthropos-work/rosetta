---
iter: 03
milestone: M256
iteration_type: tik
status: closed-fixed
opened: 2026-07-28
---

# M256 · iter-03 — the shared login + navigation path (the clause-1 lever)

**Type:** tik · **Active strategy:** `TOK-01` move 2, **as re-targeted by iter-02 D8**.

## Step 0 — re-survey

TOK-01 move 2 named the residual `networkidle` as the clause-1 lever. iter-02's baseline **falsified that as
a timing target on this host** (`networkidle` settles fast on localhost; it only deadlocked over the tailnet)
and identified the real driver: the **login handshake every test pays**. `pt-profile-identity` — login →
`/profile` → assert one name — costs **3.50 s**, *above* the 3.326 s median.

**Substitution recorded (not a re-scope):** TOK-01's strategy holds ("take the per-test latency lever, not
the parallelism lever"); only the named lever within it moves from L1 → **L3 (`storageState` reuse)**, with
L1/L2 retained as **measured correctness**. Baseline re-read from `../progress.md` § Baseline; no harness
code has changed since it was taken.

## Cluster / target identified

**The one cost shared by all 18 Playthroughs** — `hero-login.ts` → `cockpit-login.ts` `loginAs` — plus the two
unfenced `networkidle` navigations that live in the same path. A shared cost is the only kind that moves a
**median** rather than a tail.

## Hypothesis

If a Playthrough can start from an already-authenticated browser state for its seat instead of re-running the
full handshake (`selectSeat` → protected `goto` → 307 → FAPI handshake → 303 → cookie → SSR), the per-test
cost falls by most of the handshake, on **every** test — moving the median toward the ≤ 2.628 s target.

**The constraint that shapes the design (iter-01 D1, confirmed by the Phase-0b audit).** Clerkenstein holds
**one global seat**: `handleMe` resolves `activeUserLocked()` with **no cookie input**, and
`handleSelectIdentity` sets `signedIn = false; sessID = ""` **globally**. Two consequences:
- Restored cookies alone do **not** identify a hero — the *server's* active seat does. So reuse must be
  **seat-grouped and serial**, and a group must not be re-entered after the seat moves on.
- A seat switch **invalidates** any cached session, so the handshake must be paid **once per seat** (6), not
  once per test (18).

**The hazard this creates, and the guard it needs.** If ordering ever interleaves seats, a restored state
would render *a* hero — the wrong one — and a render-presence assertion would still pass. That is a
false-green of exactly the class iter-02 D6 found. So any reuse **must** carry an identity assertion that
fails loud when the resolved hero is not the expected one.

## Expected lift

Phase A measures the login's share before anything is built. If the handshake is ≥ ~2 s of a 3.5 s minimal
test, the lever plausibly reaches the target on its own; if it is small, this iter says so and re-targets
rather than building machinery for a lift that is not there.

## Phase plan

- **Phase A — diagnostic probe (no production change).** Time the handshake's legs directly: `selectSeat`
  POST, the protected-route navigation, and a direct navigation with an already-warm session. This is the
  protocol's measure-before-fix step and the honesty gate on the claim.
- **Phase B — implement**, only if Phase A supports it: seat-grouped reuse in the shared login layer, with
  the identity guard above, `pt-profile-identity` retained as the one test that still performs the **full**
  handshake so the handshake itself stays proven.
- **Phase C — L1/L2 + the widened fence.** Pin `waitUntil: 'domcontentloaded'` on the remaining login call
  sites and the two page-object `goto` overrides (`skill-path-page.ts:31`, `simulation-page.ts:36`), and widen
  `tests/home-login-networkidle.unit.spec.ts` from `/home`-landing specs to **every** login call site + page
  object `goto` + the 6 unbounded `waitForLoadState` sites the audit found.
- **Phase D — re-measure** under D7's pinned protocol (n=3, `--reset`, cold run included) and report the
  delta honestly, attributing it to the leg that moved.

## Escalation conditions

- Phase A shows the handshake is a small share → close with the falsification, re-target, do **not** build.
- The identity guard cannot be made to fail on a mismatched seat → **do not ship the reuse**. An unguarded
  speed-up that can silently run the wrong hero is worse than a slow suite; that is the milestone's thesis.
- Any Playthrough turns red → per **D-v28-3** the batch runs to completion and one consolidated red set
  escalates at the end.

## Acceptable close-no-lift outcomes

Phase A falsifying the login hypothesis is a **complete** iter (`closed-no-lift` with documented
falsification): it converts a guess into a measurement and redirects the milestone's remaining budget.
