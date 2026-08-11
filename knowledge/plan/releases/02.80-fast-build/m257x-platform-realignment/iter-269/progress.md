**Type:** tik, under `TOK-08`. Route: `FIX-M257x-262-demo-env-append-is-not-idempotent`.

# iter-269 — the duplication is a designed consequence of two correct decisions

## The measurement (values-blind — no value was read, echoed or logged)

`stack-demo/platform/.env` after 31 bring-ups:

| reading | result |
|---|---|
| total lines | **471** |
| distinct keys | **18** |
| occurrence histogram | **13 keys × 31** · 5 keys × 1 |
| keys whose **value varies** across occurrences | **0** |
| keys blank in **every** occurrence | **`DIRECTUS_TOKEN`** |

**iter-262 called it an "18-key block"; it is a 13-key block, with 5 singleton keys alongside.** Small,
and corrected because the next reader will diff against it.

## The writer, and why the routed fix was the wrong fix

`up-injected.sh:1538` runs `provision … --force` **unconditionally** on every demo bring-up. `--force`
skips the copy-if-absent check, and the merge stays **append-only** — so each bring-up adds one block,
without bound and with no reaper.

**Neither half is a defect, and that is exactly why this survived.**

- Append-only is what makes the tool values-blind. `provision/io.go:173-175`: *"an existing line is never
  re-read for its value or rewritten, so provision can never corrupt or echo a value already in the
  target."*
- `--force` is deliberate, and `up-injected.sh:1522` states its second purpose: it *"overwrites stale keys
  **AND blanks the `DIRECTUS_TOKEN` family via last-wins** (the strip-on-non-prod class)."*

**So compose's last-wins resolution is LOAD-BEARING.** The demo's `DIRECTUS_TOKEN` blanking is delivered
*by being appended last*. That refutes iter-262's routed instruction — *"find the writer and make it
replace-or-skip"* — on its own terms: replace-in-place would either **re-read an existing value** (breaking
values-blindness) or **drop the trailing blank** (re-arming `DIRECTUS_TOKEN` on a demo, the fix16/17 class
`secrets-spec.md` exists to prevent). A real repair must keep **both** properties and prune **older**
duplicates rather than stop appending.

## The hazard, stated rather than assumed

Today all 31 copies agree, so nothing is broken. The hazard is conditional and worth naming because the
file is *designed* to accumulate: **with N copies of a key the last one wins**, so the moment any writer
appends a differing value — a stale source, a partial run, a hand-edit — the file silently prefers
whichever landed last. Symptom: the classic *stack boots, catalog empty*. **Diagnose a suspect `.env` by
reading the LAST occurrence of a key, never the first.**

## Pre-registration grading

| PR | prediction | outcome |
|---|---|---|
| **PR-1** | all blocks byte-identical, no key varies | **HELD** — 0 varying keys |
| **PR-2** | `DIRECTUS_TOKEN` blank in all 31 | **HELD** — and it is the *only* always-blank key |
| **PR-3** | the writer is the bring-up path, **not** `stacksecrets provision` | **REFUTED** — the writer **is** `provision`, and it appends **by design**, for a safety reason. The bring-up's contribution is passing `--force` every time |
| **PR-4** | last-wins makes this a correctness hazard | **HELD, and stronger than predicted** — last-wins is not a latent hazard, it is the **mechanism the demo depends on** to blank `DIRECTUS_TOKEN` |
| **PR-5** | the fix needs a tag + pin bump | **HELD** — and the fix *shape* changed, which is the more useful half |

## Side-deliverable — the fence family caught iter-265's own guard

The full `stack-core` suite (**2,175 passed / 28 failed**, 37 min, run in the background across this call)
found that **iter-265 shipped `decommissioned_instruction_guard.py` without enrolling it**, turning four
registry fences RED across 24 tests. Every one was the family working as specified, and all are now green
(rext `cf4da42`):

- **`guard_family.INVOCATIONS`** (11 tests) — *"a guard that cannot be selected is not a guard."*
- **`fence_provenance`** (3) — every member stamps the tree its verdict came from, or the verdict cannot
  be re-checked from a sha.
- **`frozen_expectation_census`** (7) — docstring ratchet **238 → 239**, comment ratchet **226 → 228**,
  each with a recorded reason naming the module. The comment ratchet converged in **two passes**, because
  the reason-comment written for the first bump *is itself a comment* — anticipated by harden pass 63, and
  both passes recorded, since a ratchet documenting only its first pass reads later as a blind bump
  wearing a reason. Its own sub-fence then caught that the block's last arrow target no longer equalled
  the constant.
- **`derivation_registry.DECISIONS`** — `derive_decommissioned` classified `DECLINE:tree-scan`: the set is
  derived from the migration-status map, which `platform_alignment_guard` already fences against
  `repos.yml` both ways; freezing it here would freeze the thing that guard exists to keep honest.

**A second, unrelated RED was found and closed in the same pass:** `blocking_state_guard` was RED because
**iter-259's `user-blocker: y` had never reached `deferrals-audit.md`** — the user lifted that prohibition
mid-iter-261 and iter-262 proved the dev half, but the resolution lived only in the iter stream. Recorded;
the guard is green. **It was RED before this run and would have blocked the close gate.**

## Close — 2026-08-10

**Outcome:** The duplication is explained, and the routed fix is refuted on its own terms: append-only
protects values-blindness, `--force` delivers the `DIRECTUS_TOKEN` blank via last-wins, and *replace-or-skip
would break one or the other*. `secrets-spec.md` now carries the mechanism, the measured numbers, the
unbounded-growth statement and the read-the-LAST-occurrence diagnostic. Separately, the fence family caught
iter-265's unenrolled guard and a pre-existing close-gate blocker; both are green.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: **y** — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: exit-5

**Decisions:** `D-M257x-269-1` (last-wins is load-bearing, so replace-or-skip is the wrong repair).

**Side-deliverables:** rext `cf4da42` (family enrollment + provenance stamp + two ratchet re-pins + the
DECISIONS entry) and the `deferrals-audit.md` iter-259 row. Both are separate from this iter's planned
scope and do not change its close status.

**Routes carried forward:**
- `FIX-M257x-262-demo-env-append-is-not-idempotent` → **CLOSED**, superseded by
  `FIX-M257x-269-force-append-grows-the-demo-env-without-bound` — the fix must keep values-blindness AND
  the trailing blank, and prune **older** duplicates rather than stop appending. Needs a tag + pin bump.
- `FIX-M257x-268-ensure-clones-hardcodes-cms-as-studio-fetcher`,
  `FIX-M257x-262-dev-path-needs-the-studio-acquisition` (tooling half),
  `FIX-M257x-267-capture-the-succession-RESPONSE`,
  `FIX-M257x-266-manual-path-drops-gates-the-automated-path-enforces`,
  `FIX-M257x-265-prose-deletion-instructions-are-out-of-D-reach`,
  `ROUTE-M257x-258-the-pin-is-157-iters-stale` → open. **Four of these are one tag away**; they should
  land together.

**Lessons:**
1. **A new fence member is not tested by the tests it ships with.** iter-265 ran its guard, its own 12
   tests, and the four fences it believed adjacent — all green — while the family it had just joined was
   RED in 24 tests. Run the *whole* suite when adding a member to a registry-governed family.
2. **Two correct decisions can compose into a defect, and the composition has no owner.** Append-only is
   right. `--force` is right. Unbounded growth is neither's bug, which is why nobody's tests fail.
3. **A routed instruction can encode a wrong fix.** *"Make it replace-or-skip"* was written before anyone
   read `io.go`'s rationale. Re-derive the fix at the moment you land it, not from the routing.
