**Type:** tik — under `TOK-08` (census the mechanical classes; stop sampling them).

# iter-153 — the probe scope was derived from the wrong artifact

## Phase A — the census

The scope question is *"what services does this stack run?"* It is answered in three places, from two
different artifacts, and they disagree. Measured at platform `0c91421`, default demo flags, against a
**real** `gen_injected_override.py` emission (not a fixture, not a commit message):

| site | artifact it reads | answer | how it gets there |
|---|---|---|---|
| `up-injected.sh:2682-2690` | platform compose **+ 3 hand-written literals** | **8** | derive, then append `next-web-app studio-desk` under `NO_UI` and `directus` under `NO_LOCAL_CONTENT` |
| `stack-verify/reports/generate.sh` | platform compose **only** | **5** | derive |
| the stack's own generated override | — *it IS the artifact* | **11** | — |

```
platform_topology.py services  -> postgresql redis sentinel backend gotenberg              (5)
gen_injected_override.py       -> backend gotenberg postgresql redis sentinel
                                  directus next-web-app studio-desk
                                  hiring-app fake-fapi fake-bapi                          (11)
```

**So `/test-platform` — the tool whose entire purpose is *"what is actually working"* — probed five of
the eleven services a default demo runs, and printed `✓ pass`.**

### The flag matrix, because the hand-written tuple's premise had never been measured

| flags | services the generator emits | n |
|---|---|---|
| *(default)* | 5 platform + directus + next-web-app + studio-desk + hiring-app + fake-fapi + fake-bapi | **11** |
| `--no-ui` | 5 + directus + fake-fapi + fake-bapi | 8 |
| `--no-local-content` | 5 + next-web-app + studio-desk + hiring-app + fake-fapi + fake-bapi | 10 |
| both | 5 + fake-fapi + fake-bapi | 7 |

Two things fall out that no reading had produced:

1. **`--no-ui` drops THREE services, and the bring-up's hand-written tuple names two.** `hiring-app` —
   the surface `pt-hiring-recruiter-compare` plays through — is in the UI tier and in nobody's scope.
2. **`fake-fapi` / `fake-bapi` are unconditional in every combination and have no probe row at all.**
   They are Clerkenstein's fake Clerk front/back APIs: if either is down, **every login on the stack
   fails**, so the presenter's cockpit is dead behind a green report.

### The precondition that was refuted first

harden pass 35 routed this fix forward as Fate 3 because it *"cannot be verified without a live demo."*
Phase 1's mandatory re-survey tested that claim and it is **false** — `gen_injected_override.py` is pure
line-oriented text emission and `stack-demo/platform` is a real clone at `0c91421`, so the stack's own
override and its whole flag matrix are producible here with no docker and no bring-up. The fence that
grades all of it runs in **1.9 s**. (`D-M257x-153-3`.)

## Phase B — the repair

Three files, one idea: **derive from the artifact that decides the fact.**

- **`lib/target.sh`** gains `target_stack_override_file()` + `target_stack_override_services()` —
  the same layout constant `target_resolve_stack_dir()` already encodes, one file deeper
  (`demo-N → docker-compose.injected.yml`, `dev-N → docker-compose.dev.yml`). It reads the generated
  **artifact**, never `demo-stack`'s code; that is the whole of the coupling, and the escalation
  condition this iter's `overview.md` set was not reached.
- **`lib/scope-union.sh`** (net-new) answers *"what MORE does this stack run than the platform
  declares?"* in three lines: probeable / unprobeable / everything the override declares. A **separate
  process**, because `lib/services.sh` carries `set -euo pipefail` and sourcing it into a report
  generator contractually forbidden from aborting (`D-M257x-148-1`) would arm `-e` for the rest of the
  run. **Line 3 is load-bearing** — without it, *"the override adds nothing probeable"* and *"there is no
  override"* both produce an empty line 1, and the report would make the same confident statement about a
  stack it read and one it never located (`D-M257x-153-5`).
- **`reports/generate.sh`** unions the probeable set into `STACK_SERVICES`, and the disclosure now has
  **three branches**, each of which says something different and true: unioned-with (naming the added
  services), read-an-override-that-adds-nothing, and **found-no-override-at-all**. Where the intersection
  with the registry drops a service, the line **names it** — *"they are running and ungraded, not
  absent."*
- **`lib/services.sh`** gains `STACK_INJECTED_SERVICES_NOT_PROBED` — the third declaration array, and the
  first whose opposite side is the **stack** rather than the platform. iter-152's two arrays both take
  the platform compose as their counterparty; that is the wrong counterparty for a registry whose job is
  grading a stack. Rows were **not invented** for the three (`D-M257x-153-6`).

Live behaviour, `generate.sh live` against a synthetic stack root with a real generated override:

```
probe scope DERIVED from …/platform/docker-compose.yml: `postgresql redis sentinel backend gotenberg`,
UNIONED with the services this stack's own generated override adds: `directus next-web-app studio-desk`.
NOT probed, because the probe registry has no row for them: `hiring-app fake-fapi fake-bapi` — they are
running and ungraded, not absent.
```

## Phase C — the fence, and its controls

`stack-verify/tests/test_scope_union_m257x.py`, **16 tests**, every class **above** the `__main__` guard
(harden pass 35 appended two classes below its own and a direct run silently collected 6 of 16 while
printing OK — that is the failure this file is written not to repeat; it collects 16 both ways, verified).

**The answer key is GENERATED, never written down.** Every expectation is produced by running the real
generator against a real platform clone at test time. A hand-written service tuple is the exact defect
this milestone exists to end, and a fence that hard-codes the answer it checks re-creates it one layer up.
Where no platform clone exists the generated-key tests SKIP **with the reason named**; they never fall
back to a literal.

Arms:

- **the gap is real** — the platform set must be a *strict* subset of the override's set, and the flag
  matrix is asserted (including the arm that pins `--no-ui` dropping a service the bring-up's tuple never
  names, carrying an **inversion instruction** so closing that route cannot silently delete it);
- **partition completeness** — platform-scope ∪ probeable ∪ unprobeable must be **exactly** the
  override's service set. A service in none of the three is one the scope silently forgets;
- **two readers, on purpose** — the shell `awk` and a Python parser must agree on line 3; if both used
  the same parser an `awk` bug would be invisible;
- **the declaration equals the measured gap in BOTH directions** — arrival (an undeclared unrowed
  service) and departure (a stale declaration), `D-M257x-152-1`;
- **the disclosure is executed, not grepped** (`D-M257x-148-3` is this file's own history);
- **read-nothing vs found-nothing** are distinguishable, including the deliberately-constructed case of
  an override declaring only platform services.

**Controls, all run:** two mutation controls (a net-new unrowed service must surface as unprobeable
**and** must turn the declaration arm RED — the second exists because the first alone would not prove the
arm that matters); three anti-vacuity controls (the registry-row reader and the declaration reader must
each return a non-empty set, or every "has a row" / "has no row" assertion is vacuous; and the `awk`
parser must not cross a top-level mapping boundary, or every partition assertion checks a padded set); a
subjects-exist control.

**RED-PROOF against the REAL pre-fix code recovered from `HEAD`, not a reconstruction: 4 of 16 fail.**
The other 12 are measurement and control arms about the generator and the registry, which pre-fix code
does not change — they are **regression pins, not gap proofs**, and are recorded as such rather than
counted toward the delta (harden pass 35's own honesty rule, applied to this iter).

### The fence this repair obsoleted, and why it was re-pointed rather than deleted

Closing the gap made harden pass 35's `test_the_derived_disclosure_names_the_services_it_excludes` fail:
it asserted `generate.sh`'s **source** contains three service-name literals, and the repair removed them
deliberately. Deleting it would have retired a real property along with an obsolete spelling of it — the
precise shape the last harden pass booked against itself. Re-pointed at the **stronger** pair: the
disclosure block carries **no** service-name literal (mutation-controlled: re-introducing one fails), and
the real script when **run** still names what it left out (proven RED against pre-fix `HEAD`). Its
sibling test is unchanged and still load-bearing. (`D-M257x-153-4`.)

## Phase D — gates

| gate | result |
|---|---|
| new fence, direct run **and** pytest | **16 / 16**, both collect 16 |
| re-pointed `test_probe_scope_m257x.py` | **16 / 16** |
| **`stack-verify` full section** | **270 passed · 0 failed** — 244 at iter-152's close, +16 this fence, +10 from harden pass 35's two classes; **0 regressions**. The first run of this iter was **269 passed · 1 failed**, that one failure being pass 35's own now-obsolete assertion; re-pointed, then re-run clean (6 min 07 s) |
| `stack-core` blast radius (`service_registry_guard`, `guard_family`) | **69 passed** |
| `service_registry_guard` live | **ALIGNED** — 12 rows (7 graded, 5 declared absent) vs 7 services / 10 published ports (3 declared unprobed) |
| guard family | **20 GREEN · 0 RED · 4 not-run** (accepted) |

`stack-core`, `demo-stack`, `dev-stack`, `stack-injection` test sections **not re-run, and saying so**
(`§5` rule 60) — zero files touched in them; the two `stack-core` guards that read the edited
`services.sh` were run directly and are green.

## Close — 2026-08-08

**Outcome:** `/test-platform`'s probe scope now derives from the artifact that decides what a stack runs
instead of the one that merely constrains it — **5 of 11 services probed → 8, with the remaining 3 named
as ungradeable rather than dropped**. The routed-forward blocker (*"needs a live demo"*) was **refuted in
the re-survey**, which is why this landed in-iter instead of being carried again.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5**, unchanged. **No `N` reading was taken, so no `N` movement is claimed**
(`§9`'s guard-rail 1, in its required words).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**iters 135–153 took no reading, so the metric is UNMEASURED not unmoved — `§9`'s iter-type refinement; `TOK-08`'s sealed refutation branch bars an agent-authored successor in any case**) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (**1 tik this session**) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-153-1` … `D-M257x-153-7` (iter-153/decisions.md).
**Side-deliverables:** none — the re-point of `test_probe_scope_m257x.py` is planned scope (the repair
made its assertion false), not a side discovery.
**Routes carried forward:**
- `FIX-M257x-h33-derive-includes-stack-override` — **CLOSED by this iter**, and its stated blocker
  refuted rather than worked around.
- `FIX-M257x-iter153-bringup-scope-tuple-is-hand-written` — **NEW.** `up-injected.sh:2688-2690` hand-writes
  `next-web-app studio-desk` + `directus`; measured, `--no-ui` drops a **fourth** service (`hiring-app`)
  the tuple never names. Fenced here (with an inversion instruction), not repaired — editing the
  bring-up's verify tail is a second line of investigation on a path that does need a live demo.
- `FIX-M257x-iter153-stack-injected-services-have-no-rows` — **NEW.** `hiring-app` / `fake-fapi` /
  `fake-bapi` are declared unprobed because a probe row needs a container, port and health target
  **observed on a live stack**. The Clerkenstein pair is the one worth prioritising: their death takes
  every login down.
- `SURVEY-M257x-iter152-half-up-services-are-ungradeable` — still open, and this iter is evidence for it
  at the level above: the registry's denominator problem is not only ports-vs-services, it is also
  platform-vs-stack.
- Unchanged and still queued: `SURVEY-M257x-iter152-other-guards-may-read-prose-as-data` ·
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
1. **A routed-forward item's stated BLOCKER is a claim, and the re-survey is where it gets tested.** This
   one said *"cannot be verified without a live demo"*; one command refuted it and the fix took an iter.
   A wrong blocker carried long enough stops being re-read.
2. **Over-broad is loud; under-broad is silent.** Both halves of iter-148's defect had the same cause.
   The false-`down` half was fixed in one iter because it printed four failures; the unprobed half
   printed nothing and would have read as a clean bill of health indefinitely. Ask both questions.
3. **A gap-disclosure fence must be re-pointed when the gap closes, never deleted** — otherwise the repair
   silently retires a real property along with an obsolete spelling of it.
4. `§5` gains **rule 70** with all three corollaries (the deciding artifact; naming what an intersection
   drops; retiring a gap-disclosure fence).
