**Type:** tik (standard shape; §9 iter-type refinements consulted, none selected).

# iter-178 — the `N of M` class, measured then dispositioned

**Controlling strategy:** `TOK-08` — *census the mechanical classes; stop sampling them.*

## Phase A — census + split

`derived_count_guard` has printed *"NOT REACHED: the `N of M` prose shape"* on every run since iter-173.
Nobody had counted the class. Measured at corpus `794b167`:

| surface | occurrences |
|---|---|
| `corpus/**` + `CLAUDE.md` — all forms | 61 |
| `corpus/**` + `CLAUDE.md` — bolded `**N of M**` | 18 |
| **clause-5 surface** (`corpus/architecture/**` + `corpus/services/**`) | **9, over 8 lines** |

Split by derivability first (§8 iter-173): **4 DERIVABLE · 3 OBSERVED · 1 HISTORICAL**, 8 dispositions
covering 9 occurrences (`shared_libraries.md:257` states its claim twice on one line — once as the count,
once as its scope disclaimer — and they are one claim).

## Phase B — re-derive the DERIVABLE half

Substrate first (`D-M257x-122-4`): read from `stack-demo/`, `app @ ad9f3c498` — the ref the prose cites.

| claim | derivation | verdict |
|---|---|---|
| `1 of 43` (×2 sites, a twin) | `graph/schemas/*.graphqls` → 43; `academy.graphqls` present | **TRUE** |
| `31 of 135` | struct types embedding `ent.Schema` → **135 exactly**; live `OrganizationMixin{}` → **29**; +Membership +Organization = **31** | **TRUE** |
| `6 of 7` | 7 repos on disk carry a `go.mod`; 6 require `taxonomy`; `roadrunner` alone does not | **TRUE** |

**Four of four hold. The instrument proved itself doing it** (§9 iter-149): the schema count surfaced a
**fourth** `Policy()` declaration nobody's arithmetic had mentioned — `user.go` — which had to be
adjudicated rather than waved past. It is correctly excluded: `User`'s policy is `FilterSameUserRule` /
`DenyNotSameUser`, a **per-user** filter, not a per-organization one. That independently corroborates
iter-52's refutation of the `32` reading, reached from the code rather than from the ledger.
→ `D-M257x-178-4`.

## Phase C — Arm D

Into `derived_count_guard` rather than a new fence module (`D-M257x-178-3`): a new module would drag four
registries behind it and three of those four have been caught rotting in this milestone.

* **Keyed `<path>::<claim text>`, no line number** — an edit above a known site is not a fake finding.
* **Both directions** — an undispositioned claim is RED, and a disposition matching no live claim is RED.
* **Never verifies `M`** — the printed clause now says what it does reach and what it does not, in one
  sentence, on every run including green.
* **The fail-open is closed by a PAIR** (`D-M257x-178-5`): the guard records `arm D surface: present |
  absent` on every run *and* a test asserts the surface is present on the real tree.
* **The population size is asserted, not remembered** (`D-M257x-178-6`).

## Phase D — run

| suite | runner | result |
|---|---|---|
| `tests/test_derived_count_guard.py` | `/usr/bin/python3` 3.9.6 | 17 passed (0.78 s) |
| `tests -k "frozen or derivation or registry or census or postcondition or suite"` | same | 347 passed · 1 skipped · 1,207 deselected (160.40 s) |
| **`stack-core/tests` (whole section)** | same | **1,553 passed · 2 skipped · 0 failed** (1332.26 s) |

**The delta reconciles exactly:** iter-177 closed at **1,548 / 2 / 0**; this run is **1,553 / 2 / 0** —
`+5`, which is this iter's five new tests and nothing else. **Scope stated (r60/66):** `stack-core` only.
The other four `rosetta-extensions` sections were not run and nothing is claimed about them.

## Close — 2026-08-09

**Outcome:** the `N of M` prose class — declined by `derived_count_guard` and unmeasured for five iters —
is now **counted** (61 corpus-wide, **9 over 8 lines** on the clause-5 surface) and **dispositioned**.
Arm D asserts, both directions, that every clause-5 `N of M` carries a written `DERIVABLE:` /
`OBSERVED:` / `HISTORICAL:` disposition; it never verifies `M`, and says so on every run. The four
DERIVABLE ones were re-derived at the refs the prose itself cites and **all four hold** — and the census
proved its own instrument on the way, surfacing a fourth `Policy()` declaration (`user.go`) that had to
be adjudicated and is correctly excluded, independently corroborating iter-52's refutation of the `32`
reading from the code rather than from the ledger.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (tenth consecutive `closed-fixed`; **no
`P`/`N` reading taken, so the metric is UNMEASURED, not unmoved** — `§9`, and `TOK-08` declares the
class-by-class sweep order in advance) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n —
(6) protocol-stop: n — (7) budget-exhausted: **y** — Outcome: **exit-7** (BETWEEN ITERS, tree clean)
**Decisions:** `D-M257x-178-1` … `D-M257x-178-7` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none.

**Routes carried forward:**
- `SURVEY-M257x-iter173-derived-count-guard-reach` — **narrowed and half-closed.** The *disposition* half
  is fenced on the clause-5 surface; the *value* half is untouched and stays declined, for the reason
  iter-173 gave. What is new is that the decline is now a measurement.
- `SURVEY-M257x-iter178-n-of-m-outside-clause-5-is-52-sites` — **NEW.** Arm D's surface is the clause-5
  roots; the other **52** occurrences (`corpus/ops/**`, `corpus/tools/**`, `CLAUDE.md`) are enumerated
  but undispositioned. Widening the arm is a defensible future iter; doing it *silently*, so a green over
  61 sites reads as a statement about the gate's surface, is the conflation `F4` forbids
  (`D-M257x-178-7`).
- `FIX-M257x-iter177-ledger-carries-a-retracted-retraction` — unchanged; owner = the next harden pass.
- `FIX-M257x-iter177-derivation-registry-decline-rationale-is-false` — unchanged; open, with its
  measurement recorded at iter-177.
- `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` (narrowed at iter-177) ·
  `FIX-M257x-iter174-accept-registers-one-registry-of-two` · `FIX-M257x-iter173-ledger-denominator` ·
  the observed half of `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` — unchanged; open.
- The standing queue, unchanged.

**Lessons:** **a NOT-REACHED clause is a measurement or it is a mood.** Disclosing a blind spot is better
than hiding one, but a disclosure with no number cannot be ranked, worked, or noticed going stale — this
one sat for five iters not because it was hard but because nothing said it was nine. The move that
unlocked it does not solve the hard part: **split the VERIFICATION of the value from the DISPOSITION of
the claim, and fence only the half that is decidable.** Two corollaries paid for directly: assert the
population SIZE (the class was fenceable *because* it was nine — sixty-one would have been a different
decision), and when a fixed subject surface creates a fail-open, close it with a **pair** — the guard
records the surface state on every run *and* a test asserts the surface is present on the real tree.
Written into `platform-alignment.md` §8 in this iter's commit.
