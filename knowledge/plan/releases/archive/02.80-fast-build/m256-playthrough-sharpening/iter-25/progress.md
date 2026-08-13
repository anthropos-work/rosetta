# iter-25 — progress

**Type:** tik (standard shape, `TOK-01` move 4)

## Phase A/B — the D81 contract, failing-test-first

**RED before a line of fix**, on the real defect: after an explicit sign-out, a bare handshake left
`/v1/me` at **200**. The contract implemented:

- a handshake carrying `__clerk_identity` is an **explicit login** (the cockpit's [Log in as] and every
  Playthrough) → establishes, always, and clears the flag;
- a **bare** handshake after an explicit sign-out → must **not** establish;
- a **bare** handshake never preceded by a sign-out → **must** still establish, because that is the demo's
  first-visit path and `autoverify` handshakes bare.

That third clause is why this is a sign-**OUT** flag and not "establish only when asked". And `signedOut`
is deliberately **not** the negation of `signedIn`: a seat switch drops `signedIn` on purpose so the new
seat starts clean.

**Mutants (P1–P5), each restore byte-identical:**

| # | mutation | result |
|---|---|---|
| P1 | the establish becomes unconditional again | RED — the original defect |
| P2 | a seat switch sets the sticky flag | RED — cockpit + every Playthrough login break |
| P4 | only an explicit login establishes | RED — first visit **and** `autoverify` break |
| P3 | `establishLocked` stops clearing the flag | **PASSED — and that was DATA** |
| P5 | the seat-select clearing removed | RED — *"stranded on /login"* |

**P3 is worth the space.** It should have been RED and was not, so per iter-17 **D84** it was treated as
evidence rather than as a broken mutant. The reason was real: a form sign-in sets `signedIn`, and a later
bare handshake merely *declines* to establish — harmless while already signed in. The sequence that IS
exposed adds a seat switch: `sign out → form sign-in → /v1/demo/select → bare handshake == stranded`. So
the mutant found a **coverage hole, not a dead line**; a dedicated test now drives that 4-step sequence and
P3 re-run is RED. *The line is load-bearing and proven so rather than kept on faith.*

## Phase C — the LIVE proof, and it caught what five unit tests did not

Tag pushed → `stack-demo/rosetta-extensions` re-pinned → `fake-fapi` cross-compiled from **the stack's own
clone** (the consumption-copy policy) → image rebuilt → container recreated with `--no-deps --no-build`.
**No `/demo-up`, no `/demo-down`, no compose down, no `--purge`.**

**The first live run FAILED at step 5:** after an explicit sign-out, clicking a cockpit hero landed on
`/login`. The guard meant to make logout work had broken logging back **in**.

Cause, and it is a coverage shape rather than a logic error: `loginAsHero` POSTs `/v1/demo/select` and then
lets the middleware handshake — and **that handshake is BARE**. `handleSelectIdentity` dropped `signedIn`
without clearing the sticky flag, so the guard correctly declined. A seat **selection** is an explicit
**login intent**, so it now clears the flag. *A sticky flag needs a test per ENTRY DOOR — the handshake
with `__clerk_identity`, the sign-in form, and `/v1/demo/select` — and the uncovered door was the one every
presenter and every Playthrough uses.*

**Re-tagged, rebuilt, re-run — every reading flipped:**

| step | pre-fix (iter-16, measured) | now |
|---|---|---|
| 1 — after `/logout` | landed on **`/home`** | **`/login`** ✅ |
| 2 — Clerk cookies | all present | all present — clerk-js still leaves them; the **server** no longer honours them, which is the point |
| 3 — `/v1/me` | **200** | **401** ✅ |
| 4 — re-visit `/profile` | served the logged-out hero (`"Pat Ellis"`=1) | **`/login`**, hero absent ✅ |
| 5 — explicit login after sign-out | n/a | **works on the FIRST attempt**, lands the selected hero ✅ |

Note step 2: the fix is deliberately **server-side**. clerk-js's cookie behaviour is unchanged and
unchangeable from here; what changed is that the mock no longer treats a leftover `__client_uat` as a
licence to resurrect a session.

## Phase D — no-regression gate

Every Playthrough logs in through the path this iter changed, so the gate is the real check:

| run | result | rc |
|---|---|---:|
| A | `187 passed` (1.7 m) | **0** |
| B | `187 passed` (1.4 m) | **0** |
| C | `187 passed` (1.3 m) | **0** |

**0 flake**, rc captured per run into a variable. `ptreport` **26/31 passing, 0 failing, 0 unimplementable**.
`stackseed --policy-check` rc 0. 16 containers Up, 0 exited. Drifted cockpit fixture restored
**byte-identically** (`e991b47a`).

## Close — 2026-07-30

**Outcome:** **D-v28-5 is DISCHARGED and PROVEN LIVE** — the gate's one non-Playthrough clause. Both halves
(iter-16's sign-out route + this iter's D81 stickiness) now hold on one rebuilt `fake-fapi`, with all four
of iter-16's measurements flipped and the presenter able to log back in on the first click. No Playthrough
added, per the user's explicit call.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: **y** — (5) cap-reached: n — (6) protocol-stop: n — Outcome: **exit-4**

> **The grading is corrected, and the correction is the honest one.** This was first written as
> `cap-reached / exit-5`. That is **wrong**: this session closed **4** tiks (iter-22, 23, 24, 25), not 5, so
> the cap did not fire. What actually fires is § 4's rare-but-real clause — the session's context budget is
> genuinely spent (8 full suite runs, 3 container rebuilds, ~20 mutants, four iters of records), and the
> next routed item (`ONBOARD-M256-stage0-capability`) is a **seeder capability plus a live reseed plus a
> full gate** that cannot be completed and verified in what remains. Starting it would leave exactly the
> half-state the protocol forbids committing. Recorded rather than rounded up, because a milestone that
> spent 25 iters removing checks that could not fail should not close a session on a mis-graded exit.
**Decisions:** D81 implemented (specified at iter-16); the seat-select door recorded above.
**Side-deliverables:** none.

**Routes carried forward:**
- `ONBOARD-M256-stage0-capability` + the 4 onboarding UCs → clause 3's remaining gap. **Next.**
- `PT-M256-readiness-step-asserts` → **already discharged at iter-15**; the brief listed it as open.
- `NEGCTL-M256-studio-pair` → unchanged (deliberately withheld behind `FIX-M256-studio-false-green`).
- **`FIX-M256-autoverify-fapi-libressl` re-confirmed live:** host `curl` returns **HTTP 000** against the
  fake-FAPI's mkcert leaf on macOS (LibreSSL cannot handshake it), so the proof had to be driven in a
  browser. Already routed; this is a second independent sighting.

## Lessons

1. **Five green unit tests and one live run disagreed, and the live run was right.** The tests covered each
   door separately and never their composition, which is the only path a real user takes. *A stateful flag
   needs a test per entry door, and the doors are enumerable — so enumerate them instead of testing the two
   you thought of.*
2. **A mutant that passes is data.** P3 was the second time this milestone a "broken mutant" turned out to
   be a true report about missing coverage (iter-17 D84 was the first). The instinct to fix the mutant is
   the wrong one; the instinct to ask what path it proves is missing is the right one.
3. **A fix to a login path must be proven by logging in.** The unit suite could not have caught the
   seat-select door because the door is a *sequence of HTTP calls the harness makes*, not a function. That
   is precisely the class iter-16 D82 refused to accept a green suite for — and it was right for a second,
   different reason than the one it gave.
