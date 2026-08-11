**Type:** tik — under `TOK-08` (census the mechanical classes; stop sampling them).

# iter-155 — the three ungradeable services: a landable repair, and a pinned decision that outranks it

## Phase A — the sizing that redirected the iter

`SURVEY-M257x-iter154-other-fences-may-be-pinned-to-spellings` was **sized and rejected for this iter,
with the numbers recorded so the next session does not repeat the measurement**: across the 5 test-bearing
sections (109 files) there are **2,854** string-literal `assertIn`/`assertNotIn` calls, of which **766**
are expression-shaped under the naive predicate. That is iter-150's 30-to-1 over-report at 30× the volume.
**The work is finding a sharper predicate** — the haystack must be a subject file's *source text*, not any
string — not running the sweep.

`FIX-M257x-iter153-stack-injected-services-have-no-rows` was taken instead.

## Phase B — the derivation, and it worked

iter-153 declared `hiring-app` / `fake-fapi` / `fake-bapi` ungradeable on the rule that a probe row needs a
container, a port and a health target **observed on a live stack**. That rule conflates three ingredients
with three provenances (`D-M257x-155-1`):

| ingredient | provenance | derivable here? |
|---|---|---|
| container name | `container_for_project()` already owns the layout constant | **yes** |
| host port | `gen_injected_override.py` at offset 0 — `hiring-app 3001`, `fake-fapi 5400`, `fake-bapi 5401` | **yes** |
| health target | requires observing what each answers on | **no** |

The registry's `docker` probe kind needs **only** a container name, so all three can carry a row that
asserts something true — *the container exists and is running* — with no invented target. Built, wired,
and **measured working**: the scope-union returned all six extras as probeable and zero unprobeable, and
`service_registry_guard` read **ALIGNED — 15 rows (7 graded, 8 declared absent)**. A 10-test fence passed,
RED-proofed **4 of 10** against the real pre-fix `services.sh`.

## Phase C — and then Phase D refused it

The full `stack-verify` section came back **9 failed**, and one of them is not a stale fence:

> `test_verify.py::TestContainerLivenessM257::test_fake_fapi_and_fake_bapi_have_NO_services_sh_row_by_design`
> — *"The design decision, pinned so a future edit has to argue with it: the injected containers are
> liveness-only."*

Its two stated rationales were **both investigated, and both are answerable**:

1. *"A row would be emitted for EVERY project — including dev-N and the main dev stack, which have no
   Clerkenstein."* Now handled by the scope derivation **iters 153–154 built**, which did not exist when
   the decision was made — the same mechanism three existing rows (`directus`, `next-web-app`,
   `studio-desk`) already declare in their own comments.
2. *"Neither fits the table's probe kinds."* Answered by choosing `docker`, which needs no target — and
   `docker` **is** liveness-only, which is what the decision itself asked for.

And the gap is real: **`/test-platform` runs `verify.sh`, not `autoverify.sh`** (`generate.sh:166`), so
autoverify's container-liveness block — where the pinned decision put these three — is **never reached on
the report path**. Measured, not assumed.

**So the argument is strong. It was still not made here, and the reasoning is the deliverable.** Reversing
a design decision that was pinned *specifically so a future edit has to argue with it* rests, in part, on
live-stack facts (the fake FAPI's TLS behaviour, the fake BAPI's loopback publish) that this session
cannot observe. Landing it would also have meant re-pointing three more count fences (`REGISTRY_BASES`,
the offset matrix, the test-side mirror) at the end of a budget. A half-landed reversal of a pinned
decision is the worst available outcome — worse than either leaving it or doing it properly. **Reverted,
in full, and routed with every measurement attached** so the next session lands or refuses it in one
sitting with nothing to re-derive.

## Side-deliverables — two fences re-pointed, and one gate-gap corrected

Both are the **third and fourth** instances of `§5` rule 71 in three iters. The third is the sharper one:

**`stack-core/tests/test_bringup_verify_scope_m257x.py`** — written by **iter-154, the iter that authored
rule 71**, and it failed rule 71 one iter later. It asserted as literals that the three services are *not*
in scope and that the tail prints `UNGRADED`; iter-155's row addition made both false while the code got
**more** correct. Re-pointed so the expectation is **derived from `scope-union.sh` at test time**, with a
new `else` branch asserting that with nothing unprobeable the tail must **not** warn — a warning nobody
can act on trains readers to skip the real ones. It is now correct on **both** sides of the reverted
change, which is the property the literal version never had. Kept.

**`stack-verify/tests/test_verify.py::test_up_injected_passes_frontend_scope_...`** — asserted the literal
hand-append line iter-154 deliberately deleted. Re-pointed to the structural property (*the tail calls the
union; no scope-building line names a service*), with the behavioural half already living in the executing
fence.

⚠ **It also records a real gap in iter-154's own gates, corrected in place rather than left standing:**
iter-154 ran `demo-stack`, `dev-stack`, `stack-core`-targeted and the guard family, and **did not run
`stack-verify`** — where that test lives. Its close named `stack-core` and `stack-injection` as "not
re-run" and **did not name `stack-verify`**. `§5` rule 60 requires naming what a scoped run did not cover;
**an omission from that list reads as coverage**, which is the same failure mode as the scope this whole
thread is about, one level up.

## Phase D — gates (post-revert)

| gate | result |
|---|---|
| `stack-verify` `test_verify.py` | **211 passed · 0 failed** |
| `stack-verify` (`test_verify` + `test_scope_union` + `test_probe_scope`) | **243 passed · 0 failed** |
| `stack-core` `test_bringup_verify_scope_m257x.py` | **15 / 15** |
| the reverted tree | `services.sh` byte-identical to `HEAD`; the row fence removed with it |

`demo-stack`, `dev-stack`, `stack-injection` **not re-run, and saying so** (`§5` rule 60) — the revert
returns `services.sh` to the state those sections were green against one iter ago, and no file of theirs
was touched.

## Close — 2026-08-08

**Outcome:** the repair was **built, measured working, and then refused by a pinned design decision it
would have reversed on partly-unobservable grounds.** Reverted in full. What the iter produces instead is
the complete argument — both of the decision's rationales answered, plus the measurement that
`/test-platform` never reaches the liveness block the decision relies on — and two fences re-pointed off
spellings, one of them written by the iter that authored the rule against spellings.
**Type:** tik
**Status:** closed-no-lift
**Gate:** NOT MET — **4 of 5**, unchanged. **No `N` reading was taken, so no `N` movement is claimed**
(`§9`'s guard-rail 1, in its required words).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**a `closed-no-lift` with documented falsification does NOT count toward the no-progress streak — it is a first-class outcome, not an under-delivery; and iters 135–155 took no reading, so the metric is UNMEASURED not unmoved**) — (3) re-scope: n — (4) user-blocker: n (**the design-decision conflict was RESOLVED by reverting, which is the protocol's documented default for a mid-iter finding that cannot land cleanly; nothing is left half-landed and nothing waits on an answer**) — (5) cap-reached: n (**3 tiks this session**) — (6) protocol-stop: n — (7) budget-exhausted: **y** — Outcome: **exit-7**
**Decisions:** `D-M257x-155-1` … `D-M257x-155-5` (iter-155/decisions.md).
**Side-deliverables:** the two fence re-points above (separate concern from the reverted rows; they do not
upgrade the close status) and the in-place correction of iter-154's gate list.
**Routes carried forward:**
- `FIX-M257x-iter155-add-injected-rows-vs-the-pinned-liveness-only-decision` — **NEW, and it supersedes
  `FIX-M257x-iter153-stack-injected-services-have-no-rows`.** Everything is measured and in this file: the
  three base ports, the `docker`-kind rationale, both of the pinned decision's rationales answered, and
  the `verify.sh`-vs-`autoverify.sh` measurement that shows the gap is real. **What it needs is a live
  stack** to confirm the fake FAPI's TLS behaviour and the fake BAPI's loopback publish, and a decision to
  reverse the pin. Re-doing it costs ~20 minutes; re-deriving it would cost an iter.
- `FIX-M257x-iter155-injected-service-readiness-needs-a-live-stack` — **NEW**, the narrower residual: even
  with rows, these three would be graded on **liveness only**; a container that runs without serving still
  reads `up`, which is `SURVEY-M257x-iter152-half-up-services-are-ungradeable`'s class.
- `SURVEY-M257x-iter154-other-fences-may-be-pinned-to-spellings` — **still open, now SIZED**: 2,854
  string-literal assertions / 766 expression-shaped across 109 files. **Four confirmed instances in three
  iters.** The next step is a sharper predicate, not a sweep.
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
1. **When a blocker names a conjunction, check each conjunct's provenance separately.** *"container, port
   and health target, observed on a live stack"* bound two derivable ingredients to one that is not, and
   the cost was three services — two of which take every login down — reading as `✓ pass`.
2. **A decision pinned "so a future edit has to argue with it" is doing its job when it stops you.** Both
   of its rationales were answerable and the gap it leaves is real; that still is not the same as having
   the evidence to reverse it. The iter's output is the argument, complete, so the next session decides
   rather than re-derives.
3. **Rule 71 caught its own author one iter later.** Writing the rule did not prevent the mistake; the
   durable defence is structural — **derive the expectation from the same source the code derives from** —
   and that is now the prescribed repair rather than an instruction to be careful.
4. **An omission from a "not re-run" list reads as coverage.** iter-154 did not run `stack-verify` and did
   not say so; the RED surfaced one iter later, in a section the close implied was covered.
