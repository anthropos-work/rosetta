# iter-23 — progress

**Type:** tik (standard shape, `TOK-01` move 4 — *"close the honesty items last, deliberately, not as
leftovers"*)

The full evidence, the line-level cause and the sweep live in the milestone-level
[`decisions.md`](../decisions.md) § `DEFECT-M256-silent-forbidden-mutation` — **deliberately there and
not only here**, because a defect recorded inside the iter that found it is a defect that closes with
the iter.

## Phase A — reproduce it deliberately, and enumerate the channels

The defect was only ever visible because a demo was **missing** a grant; iter-21's fix removed the
symptom from every demo. So it was reproduced on purpose: back up the
`p3 admin → org:feature:taxonomy:write` row → `DELETE` it on **demo-2 only** → reload Sentinel → drive
the real create-role journey → restore → re-verify.

**The claim under test is a NEGATIVE** (*"nothing at all is surfaced"*), so the probe enumerated ten
channels rather than checking one. The single sharpest fact came out of the enumeration and would have
been missed by any narrower probe:

> **`[role=alert]` count 1 — text EMPTY.** The form has a mounted error slot that says nothing.

Also: the dialog **stays open with `Save` still enabled** (inviting a retry that fails identically), the
catalog total reads **49 → 49**, nothing lands, the console carries only an unrelated Clerk warning — and
there **is** an uncaught page error, so the failure exists as an unhandled rejection, not as UI.

## Phase B — the line-level cause: TWO defects, one symptom

Read-only source reads in `stack-demo/next-web-app`. Zero platform edits.

1. **`AddJobRole.tsx` handles exactly one error code and rethrows the rest** — `if (dup) {…} throw error;`
   from an `async` click handler, i.e. an unhandled rejection. `onClose()` sits after the try/catch, which
   is why the dialog stays open, and the empty `[role=alert]` is the **duplicate-warning slot**, never
   populated. *That is why iter-05 saw "an EMPTY alert region": the alert element belongs to a different
   error.* This `throw error;` is the **only one** in the whole `packages/ui` tree.
2. **The systemic half:** the app's global `mutations.onError` is `captureException` + PostHog — **no user
   surface at all**. Every mutation is silent to the user on failure unless it builds its own.
3. **And a dead contract that makes it look handled:** six mutations across four `hooks/organization/*`
   files declare `meta: { error: 'Failed to …' }` — and **no handler reads them**. There is **no
   `MutationCache`** anywhere (0 occurrences); the only consumer is `QueryCache.onError`, which reads
   `query.meta.error` and uses it as a **Sentry tag**.

## Phase C — the sweep

All four org-admin writes share outcome (2), and the settings write (`pt-orgadmin-setting-toggle`'s
mutation) additionally carries one of the **dead** `meta.error` strings. So the sweep's result is not
"three more forms are silent too" but something more pointed: **the org-admin writes' authors wrote
failure messages and the framework never wired them up.** That names a fix using a convention the
codebase already believes it has.

**And the limit of that sweep, stated rather than glossed.** This iter's own hypothesis 2 said *"measure
it; do not infer it from (1)"*, and **the sibling half was answered by source read, not by driving.**
Only `createJobRole` was actually refused live — the three siblings check different permissions, and
revoking each would have been three more revoke/restore cycles against a stack later iters depend on.
What the source establishes **definitively** is the dead-`meta.error` claim (a search for
`MutationCache` returns **0** results, so nothing can read a mutation's `meta`) and the global handler's
Sentry-only body. What it establishes only by **inference** is that a refused tags-create or
settings-toggle would look equally silent to the user. That inference is strong and it is not a
measurement, and the difference is exactly the kind this milestone refuses to let slide. **Routed as a
residual on the defect record, not claimed as measured.**

## Phase D — safety, and a fence proven for free

- Every write went to **demo-2's own Postgres**. Production was neither written nor read this iter.
- The grant row was backed up before the revoke and **restored byte-identically** (`diff` clean).
- `stackseed --policy-check --stack demo-2` → **rc 0 · `live=18 expected=18`**, re-verified after the
  restore *and* again after a full reset-to-seed.
- **Confirming suite run:** `187 passed`, **rc 0**, 1.5 m — the revoke/restore left no damage.
- The deliberately drifted cockpit fixture restored byte-identically (`e991b47a`).
- **Side benefit:** iter-21's `--policy-check` fence was watched **RED against a live stack** for the
  first time — `rc 1`, `live=17 expected=18`, naming `MISSING admin → org:feature:taxonomy:write
  (under-grant)`. It had only ever been proven against mutants. And its rc was captured **into a
  variable**: read off a pipe it reported `0`, which is the milestone's own pipe-discipline rule catching
  me in the act.

## Close — 2026-07-30

**Outcome:** the milestone's one **product** defect is captured with evidence good enough to act on
without re-deriving it — reproduced deliberately after its own fix had removed the symptom, enumerated
across ten channels, root-caused to two lines in two files, and swept across the org-admin write set.
D98 is **confirmed and sharpened**: it is not that there is no error handling, it is that there is one
error surface and it belongs to a different error. **No gate clause moves** — that was the plan.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** the milestone-level `DEFECT-M256-silent-forbidden-mutation` record in
[`../decisions.md`](../decisions.md)
**Side-deliverables:** the `--policy-check` fence's first live RED (recorded above, not a separate
commit — it is a measurement of existing code, not a change to it).

**Routes carried forward:**
- **The defect itself → the platform.** Recorded at milestone level so `/developer-kit:close-milestone`
  and `/developer-kit:close-release` route it. **Not fixed here** — it is a platform edit, and zero
  platform edits is the release's hardest constraint.
- `ONBOARD-M256-seat-append` + the 4 remaining onboarding UCs → the long pole, unchanged. **Next.**
- `D-v28-5` part (b) · `PT-M256-readiness-step-asserts` · `NEGCTL-M256-studio-pair` → unchanged.

## Lessons

1. **A fix can destroy the evidence for the defect it reveals.** This defect was visible *only* because a
   demo lacked a grant. Granting it — correct, necessary — made the refusal path unreachable by accident.
   Anything found in a broken state must be captured while the state exists, or reproduced deliberately
   and on the record. That generalises past this defect and is now in the protocol doc.
2. **A negative finding needs an enumeration, not a check.** "Nothing is surfaced" cannot be established
   by looking in one place. Ten channels produced the fact that reframed the whole report: the alert
   region is *present and empty*, which is a completely different bug from "there is no alert region".
3. **A message nobody reads is worse than no message.** Six `meta.error` strings look like a user-facing
   error contract, sit on exactly the writes this milestone covered, and are consumed by nothing —
   because the cache that would consume them was never instantiated. A reader of that code would
   reasonably conclude failures are surfaced. *The most convincing evidence that something is handled is
   code that was written to handle it and then wired to nothing.*
4. **The pipe discipline caught me mid-iter.** `stackseed --policy-check … | tail` reported rc `0` while
   the check had exited `1`. The milestone has a written rule about this and I broke it inside the same
   hour I relied on it. Rc into a variable, always.
