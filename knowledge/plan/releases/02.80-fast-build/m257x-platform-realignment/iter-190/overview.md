---
iter: 190
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-09
controlling_strategy: TOK-08
---

# iter-190 — the census keyed on a SHARED constant, and the interesting pair is the one that shares none

**Active strategy reference:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).

## Step 0 — re-survey, then ENUMERATE

iter-189 routed `SURVEY-M257x-iter189-the-parity-question-is-unasked-for-every-other-dual-reader` with a
selector — *a function whose docstring or name claims to be the same derivation as another.* Re-surveyed:
grepping that phrasing across `stack-core` returns **one** hit, the pair iter-189 already repaired. A
selector that finds only its own founding case is not an enumeration, so it was replaced by a
**mechanical** one and run over all 42 modules:

> a module-level constant used by BOTH a filesystem-reading function and a git-reading function.

**Result: 6 pairs, all in `platform_predicate_guard`** — `_GO_GETENV` (iter-189's, the instrument's own
proof it fires), `_INCLUDE_HEAD`, `_LIST_ITEM`, `_TOP_KEY`, `_REF_PINNED`, `_REPOS_YML_ENTRY`.

**All 6 agree today**, measured against the real clones: `parse_compose` → 7 services and
`compose_counts_at(None)` → `(5, 7)`; `repos_yml_history` ⊇ `parse_repos_yml` with 0 current entries
missing. A zero, therefore — and `§9` says a census returning zero must prove its instrument.

## Cluster / target identified

**Proving the instrument is what found the defect, and it is the pair the census structurally cannot
see.** `compose_counts_at` recognises a compose service with `_COMPOSE_SERVICE_KEY`; `_parse_one_compose`
recognises the same construct with `_SVC_KEY`. They are **two readers of one file that share no
constant** — so the shared-constant census misses them by construction (iter-184's class, in the
instrument written this iter).

| recogniser | pattern | first character |
|---|---|---|
| `_SVC_KEY` (`parse_compose`, the G1/G7/G8 topology) | `^  ([A-Za-z][A-Za-z0-9_.-]*):` | **letter only** |
| `_COMPOSE_SERVICE_KEY` (`compose_counts_at`, G10) | `^  (?P<name>[A-Za-z0-9_.-]+):` | any of `[A-Za-z0-9_.-]` |

Measured over a 9-name candidate table: **5 disagreements** (`3d-render`, `_internal`, `-legacy`,
`.hidden`, `9front`). Compose's own service-name charset admits a leading digit, so `_SVC_KEY` is the
**wrong** one — and its direction is under-count: a service it cannot see is absent from the topology G1
grades profile membership against, so claims about it read as UNREACHED rather than graded, while G10
counts it. Latent on today's platform (every service name is letter-initial); mechanical, and one
constant to close.

## Hypothesis

Unifying the two recognisers and shipping the dual-reader enumeration as a **fence with a declared parity
registry** (both directions) closes the class as `§8` prescribes — by an enumeration that keeps running —
and the *shared-constant* blind spot is closed by an arm that compares the two recognisers **as
predicates over a name table**, which needs no shared constant at all.

## Expected lift

No `P`/`N` reading (`§9`). 1 latent divergence closed; the 6-pair enumeration shipped with a
declared/derived partition; ≥6 arms; mutation-proven.

## Phase plan

- **A** — enumerate (done); measure each pair's agreement (done); measure the recogniser divergence (done).
- **B** — unify the service recogniser on compose's real charset.
- **C** — fence: the enumeration + declared parity registry, both directions · the two recognisers agree
  as predicates · the `§9` instrument control.
- **D** — mutation-prove; both runners; the guard's own 184-test suite.
- **E** — publish the `§8` rule; route residuals.

## Escalation conditions

- If widening `_SVC_KEY` makes `parse_compose` see a *new* service in any real clone, this is a **live**
  under-count in a published topology, not a latent one — grade it as such and stop.
- If the guard's own suite goes RED on the widening, the narrow charset was load-bearing somewhere;
  keep it, name where, and fence the difference instead.

## Acceptable close-no-lift outcomes

- The two recognisers turn out to be deliberately different (a comment or test says so) → record the
  falsification, keep both, and ship only the enumeration.
