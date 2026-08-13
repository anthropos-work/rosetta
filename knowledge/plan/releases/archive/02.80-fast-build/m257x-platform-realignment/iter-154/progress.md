**Type:** tik — under `TOK-08` (census the mechanical classes; stop sampling them).

# iter-154 — the two bring-ups computed their verify scope by hand-appending literals

## Phase A — the census, and a second refuted precondition

iter-153 routed `FIX-M257x-iter153-bringup-scope-tuple-is-hand-written` forward because it *"does need a
live demo to grade."* Phase 1's re-survey applied iter-153's own `D-M257x-153-3` to that sentence and it
did not survive either: the verify tail is a bash block, extractable and executable exactly as iter-147
extracted `derive_profile` with `awk`. **Two consecutive iters have now refuted the blocker on their own
routed item** — the second one written by iter-153 itself, hours earlier.

And the census found the shape at **both** bring-ups, not the one the route named:

| site | derives | then hand-appends | gated on |
|---|---|---|---|
| `up-injected.sh:2686-2689` | platform set | `next-web-app studio-desk` · `directus` | `NO_UI` · `NO_LOCAL_CONTENT` |
| `dev-stack:352` | platform set | `directus` | `local_content` |

### The history is the finding (`D-M257x-154-1`)

- **iter-55** replaced the demo side's hand tuple with `platform_topology.py` and did **not** carry the
  repair to its dev twin.
- **harden pass 3** replaced the dev side's hand tuple with the same derivation — and left **both** sides'
  conditional hand-appends standing.

The *base* set was derived twice; the *conditional* set was never derived at all. Each repair wrote down
in its own comment that the sibling had the same defect, and neither wrote a fence over the pair. `§5`
rule 69 in its plainest instance: **an observation about a twin is not a fence over it.**

### The conditional was right; the SET was wrong (`D-M257x-154-2`)

`[ "$NO_UI" = 1 ] || verify_svcs="$verify_svcs next-web-app studio-desk"` encodes a true fact with a false
enumeration. Measured against `gen_injected_override.py` at platform `0c91421`, **`--no-ui` drops three
services** — the two named plus **`hiring-app`**, the surface `pt-hiring-recruiter-compare` plays through.

**No amount of reading that line finds this.** It is internally consistent, its comment is correct and its
gate is correct; the defect is visible only by asking the generator what it emits. The site was never
censused because it *looked* derived.

## Phase B — the repair, both sides, one commit

Both bring-ups already write the stack's own override **before** they verify (`up-injected.sh:1933`,
`dev-stack:128`), so both now read it back through `scope-union.sh` (iter-153) instead of naming services:

- **`up-injected.sh`** — the two hand-appends replaced by the union, which also **logs** the services the
  probe registry has no row for (*"running and UNGRADED, not absent"*).
- **`dev-stack`** — the same, in the same commit, because the sibling set was already enumerated.

Non-fatal on every branch (`D-M257x-154-4`): missing script, failing script, or no override each leave
`verify_svcs` exactly as the platform derivation left it — pre-iter-154 behaviour minus the tuple. One
consequence recorded rather than left implicit: `scope-union.sh` is invoked **by path**, so its executable
bit is load-bearing on both bring-ups; a lost mode bit would silently un-union every stack's scope with no
error, which is the quiet failure this whole thread is about. Asserted.

## Phase C — the fence

`stack-core/tests/test_bringup_verify_scope_m257x.py`, **15 tests**, all classes above the `__main__`
guard. It lives in `stack-core` because its subjects span three sections and **a twin fence that lives
inside one of the twins has a home-field advantage**.

Both tails are **EXTRACTED from the real scripts and EXECUTED** — never grepped (`D-M257x-148-3`). Arms:

- **anti-vacuity zero** — the anchors resolve and both blocks are non-empty and call the union. Without
  this, every behavioural assertion below could be silently testing the empty string;
- **neither tail names a service**, checked twice: inside the extracted block, and across the **whole**
  verify tail from the platform derivation to the `autoverify.sh` invocation — because a repair that
  merely *moves* the tuple ten lines up would pass the narrow arm;
- **the demo tail unions the real override** — the UI tier enters the scope, the three unrowed services do
  **not**, and the tail **announces** them;
- **the flag conditionals survive as emergent properties** — `--no-ui` drops the UI tier and keeps
  directus; `--no-local-content` drops directus and keeps the UI tier — asserted **without the tail
  knowing any of their names**;
- **the dev tail** gets directus from a `--local-content` override and **not** from a prod-read one (the
  false-`down` the M22 gate existed to prevent);
- **non-fatal on every branch**, including a deliberately unresolvable `$HERE`.

**Controls, all run:** a mutation proving the no-literal arm detects a re-introduced tuple; a mutation
proving the demo scope arm reads *the override* and not something else that happens to say `next-web-app`
(an override without the UI tier must not produce one); an anti-vacuity control proving the extracted
block modifies `verify_svcs` at all (a block that never assigns it would pass every negative assertion for
free); a subjects-exist + executable-bit control.

**RED-PROOF against the REAL pre-fix bring-ups recovered from `HEAD`: 8 failures + 2 errors of 15.**

### The second consecutive fence found pinned to a SPELLING (`D-M257x-154-3`)

`dev-stack`'s own contract test `test_verify_scopes_directus_only_when_local_content_on` failed — it
asserted the **literal source line** `[ "$local_content" = 1 ] && verify_svcs="$verify_svcs directus"`.
iter-153 had re-pointed harden pass 35's disclosure fence for the identical reason one iter earlier.

Both fences protected a real property and encoded it as the current *spelling* of the code implementing
it. Re-pointed, not deleted: the property moves to the behavioural arms above, and what stays in the
body-text contract is the structural half such a test can honestly assert — *it calls the union*, and *no
verify-scope line names a service*.

> **A fence written by quoting the line it guards will fail on the day that line is IMPROVED, and its
> failure is indistinguishable from the day that line is BROKEN.** That is not an argument for fewer
> fences; it is an argument for writing them against behaviour. `§5` gains **rule 71**.

## Phase D — gates

| gate | result |
|---|---|
| new fence, direct run **and** pytest | **15 / 15** |
| RED-proof vs pre-fix bring-ups (from `HEAD`) | **8 failures + 2 errors of 15** |
| `dev-stack` full section | **151 passed · 0 failed** (1 failed before the re-point; that one was this iter's own, graded not bypassed) |
| `demo-stack` full section | **9 failed · 1,055 passed** — the nine identical **by name** to iter-147's baseline (6 sha-drift + 3 host-postgres, both already routed). **0 regressions** |
| `stack-core` targeted (`test_bringup_verify_scope` + `demo_knob_guard` + `dev_flag_guard`) | **53 passed** |
| guard family | **20 GREEN · 0 RED · 4 not-run** (accepted) |

`stack-core` and `stack-injection` full sections **not re-run, and saying so** (`§5` rule 60): no
`stack-core` runtime code was touched (one test file added), and `stack-injection` was not touched at all.
The two `stack-core` guards that read the edited bring-ups — `demo_knob_guard` (parsers) and
`dev_flag_guard` — were run directly and are green.

## Close — 2026-08-08

**Outcome:** both bring-ups now read their verify scope back from the override they themselves wrote,
instead of hand-appending service names. The measured defect: the demo tuple named **two** of the three
services `--no-ui` drops, so `hiring-app` was in no stack's verify scope on any path.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5**, unchanged. **No `N` reading was taken, so no `N` movement is claimed**
(`§9`'s guard-rail 1, in its required words).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**iters 135–154 took no reading, so the metric is UNMEASURED not unmoved — `§9`'s iter-type refinement; `TOK-08`'s sealed refutation branch bars an agent-authored successor in any case**) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (**2 tiks this session**) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-154-1` … `D-M257x-154-5` (iter-154/decisions.md).
**Side-deliverables:** none. The `dev-stack` contract-test re-point is planned scope — the repair made its
assertion false, and grading it was the iter's obligation, not a discovery.
**Routes carried forward:**
- `FIX-M257x-iter153-bringup-scope-tuple-is-hand-written` — **CLOSED by this iter**, at both sites, and
  its stated blocker refuted (the second consecutive route whose blocker did not survive the re-survey).
- `FIX-M257x-iter153-stack-injected-services-have-no-rows` — **still open and now louder**: both bring-ups
  print the unrowed set on every run, so `hiring-app` / `fake-fapi` / `fake-bapi` are now visibly ungraded
  rather than invisibly so. Adding rows still needs a live stack.
- `SURVEY-M257x-iter154-other-fences-may-be-pinned-to-spellings` — **NEW**, and the generalisation of
  `D-M257x-154-3`. Two fences in two iters asserted a **quoted source line** rather than a behaviour.
  Neither was found by review; both were found by improving the code they guarded. The population is
  every fence in the family that reads its subject as **text** — but **do not close this on a grep for
  `assertIn(` against source strings**: iter-152's route warns that anchoring is a mechanism, not the
  property. Grade whether the assertion would survive a *correct* rewrite of its subject.
- Unchanged and still queued: `SURVEY-M257x-iter152-half-up-services-are-ungradeable` ·
  `SURVEY-M257x-iter152-other-guards-may-read-prose-as-data` ·
  `SURVEY-M257x-iter150-partition-completeness-elsewhere` · `FIX-M257x-iter145-sha-baseline-drift` ·
  `-iter145-migrate-race-needs-a-host-postgres` · `-iter145-green-but-stale-graphql-mentions` ·
  `SURVEY-M257x-iter144-orphan-arm-is-the-residual` · `FIX-M257x-iter144-correction-vs-retraction-unfenced` ·
  `FIX-M257x-iter143-wrong-head-is-unfenced` · `-iter143-scope-derivation-by-grep` ·
  `-iter143-appending-to-the-protocol-doc-rots-the-ledger` · `FIX-M257x-iter142-value-change-articles` ·
  `-iter142-path-arm-window` · `-iter142-tier-b-underflag` · `FIX-M257x-iter135-adjudicated-live-defects` ·
  `-iter140-receipts-not-checkable-here` · `-iter140-receipt-fence` · `-iter138-anchor-rot-fence` ·
  `-iter134-fence-family-has-no-shared-predicate-layer` · `-iter133-two-fives-need-a-fence` ·
  `-iter131-predicate-sets-not-enumerated`.
**Lessons:**
1. **A site that LOOKS derived can still carry a hand-written half.** Both tails opened with a correct
   `platform_topology.py` call, which is exactly why neither was censused for 30 and 150 iters
   respectively. Grade the whole expression, not the first assignment.
2. **When a repair's own comment names its twin, the fence is due in that commit** — this class was
   half-repaired twice, each time with the sibling's defect written down beside it.
3. **A fence that quotes the line it guards fails identically on improvement and on breakage** — `§5`
   rule 71. Twice in two iters; both re-pointed to behaviour rather than deleted.
