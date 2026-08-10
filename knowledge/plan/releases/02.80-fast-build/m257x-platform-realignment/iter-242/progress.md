**Type:** tik — under `TOK-08`, closing out `ROUTE-M257x-236-host-is-the-unreliable-witness`, which
iter-241 narrowed but explicitly left open (*"the other 30 family members were not audited"*).

# iter-242 — the family audited against a fresh clone set: 2 of 26 are sensitive, and both are honest

## The measurement

The whole guard family, run twice against the **same corpus** — once on this box's 13-clone workspace,
once on a clone set restricted to what a fresh bring-up creates (`repos.yml`'s four + the two
`clone_pin_guard` sanctions + the rext consumption clone) — and diffed member by member.

| | this box (13 clones) | fresh set (7 clones) |
|---|---|---|
| summary | **26 GREEN · 0 RED · 0 could-not-check · 5 not-run** | **25 GREEN · 0 RED · 1 could-not-check · 5 not-run · 1 off a fallback rung** |
| members whose verdict differs | — | **2** |

The two:

* **`platform_alignment_guard`** — GREEN on both, `11 of 109` → `27 of 109` citations unchecked. **It
  says so**, because iter-241 made it. This is the disclosure working end-to-end, one iter later.
* **`service_registry_guard`** — **GREEN → CANNOT-RUN (rc=2)**. It cannot read the compose file it grades
  and **fails closed**, which is exactly the contract.

> ### No member exits 0 on both runs while silently checking less. The family's summary line is NOT
> ### reach-blind — it moves on both axes and names the member that dropped to a fallback rung.

That is the reassuring answer, and it converts `ROUTE-M257x-236` from *"one member fixed, 30 unaudited"*
to *"audited empirically; 2 of 26 are clone-set-sensitive and both are honest about it."*

## The defect this iter DID find — and it found it by being bitten

**The first run of this measurement produced a byte-identical output file.** Same summary, same 26 member
verdicts, same `platform …` banner — naming the workspace the run was trying to avoid — and **no warning
of any kind**. `cmp` reported the two files identical.

The cause: `guard_family.py` does `Path(a.platform).resolve()`, and the clone set every platform-facing
member grades is `platform/..`. The restricted set had been built the natural way — **symlinks** into the
existing workspace — so `.resolve()` walked straight back to `stack-demo` and every member graded the
original tree.

**Had that reading been believed, this iter would have published "the family is completely
clone-set-independent" — the exact opposite of the truth**, on evidence that was internally perfect. It is
the run's sixth instrument-first finding and by some distance the most dangerous, because the wrong answer
arrived with no symptom at all.

**Landed:** a non-fatal `NOTE` whenever `--platform` resolves to a path whose parent differs from the one
given, naming both clone sets and saying that a symlinked alternative *does not take effect*. Non-fatal by
design — a symlinked workspace is legitimate; **the defect was never the resolution, it was that the
resolution was silent while changing the subject**. Two tests: the note fires on substitution, and stays
quiet on a real directory (`test_guard_family.py` 52 → **54**).

## Pre-registration — scored 3 confirmed / 2 refuted

| claim | prediction | result |
|---|---|---|
| `P-242-1` ≥ 2 members differ | ≥ 2 | **CONFIRMED** — exactly 2 |
| `P-242-2` ≥ 1 silently green while checking less | ≥ 1 | **REFUTED — 0.** Both sensitive members are honest |
| `P-242-3` ≥ 1 verdict-CLASS change | ≥ 1 | **CONFIRMED** — `service_registry_guard` GREEN → CANNOT-RUN |
| `P-242-4` the family summary is identical (reach-blind) | identical | **REFUTED** — it changes on both axes and names the fallback-rung member |
| `P-242-5` no member crashes on the restricted set | none | **CONFIRMED** |

**Third consecutive iter whose pessimistic pre-registrations were wrong about the tooling.** iters 240,
241 and 242 each opened expecting rot and found instruments that were already honest — 2 of 5, 1 of 5 and
2 of 5 predictions refuted, every refutation in the same direction. Five consecutive dirty *corpus* inputs
(235–239) trained an expectation that does not transfer to the *tooling*, and `P-242-4` is the clearest
case: it predicted that the sentence this milestone quotes as evidence 40+ times was meaningless, and the
sentence was fine.

## Close — 2026-08-10

**Outcome:** the family is audited against a fresh bring-up's clone set — **2 of 26 members are
clone-set-sensitive and both disclose or fail closed**, and the summary line is not reach-blind. The real
defect was the runner silently substituting a symlinked clone set, which is now disclosed and tested.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: y — Outcome: exit-7
**Decisions:** `D-M257x-242-1` (the substitution note is non-fatal, because symlinked workspaces are
legitimate and the defect was the silence) · `D-M257x-242-2` (`service_registry_guard`'s CANNOT-RUN on a
restricted set is CORRECT behaviour and is not "fixed").
**No `N`/`P` movement is claimed** — this iter took no graded seat.

**Suite state at close** — `stack-core` (pytest 8.4.2 / CPython 3.9.6, Python):
`tests/test_guard_family.py` **54 passed / 0 failed** (was 52). iter-241's `test_platform_alignment_guard.py`
66, iter-240's 14, iter-239's 17 — all still green. Guard family at platform reach on this box:
**26 GREEN / 0 RED / 0 could-not-check / 5 not-run**; on the restricted set: **25 GREEN / 0 RED / 1
could-not-check / 5 not-run**.

**Side-deliverables:** none.

**Routes carried forward:**
- `ROUTE-M257x-236-host-is-the-unreliable-witness` → **CLOSED by measurement.** Audited empirically across
  all 26 runnable members; 2 are clone-set-sensitive, both honest. Retained as history, not as backlog.
- `ROUTE-M257x-241-wider-citation-surface-is-ungraded` → open (107 corpus citations into the six
  frozen-legacy repos, outside the map, graded by nothing).
- `ROUTE-M257x-240-prereq-floors-live-in-three-parallel-blocks` → open.
- `ROUTE-M257x-239-stackseed-sentinel-reload-is-demo-only` → open.
- `ROUTE-M257x-238-claude-md-fences-are-unmaintained` → open, six-for-six.
- `ROUTE-M257x-238-container-vs-native-is-undrawn` → open.
- `ROUTE-M257x-237-critical-env-list-is-unfenced` → open.
- `ROUTE-M257x-236-disclosure-scope-is-document-level` → open.
- `ROUTE-M257x-235-fence-scope-is-unread` → open.
- `ROUTE-M257x-235-runnable-block-has-two-halves` → open.

**Lessons:**
1. **A negative result that arrives with no symptom is the most dangerous reading there is.** A
   byte-identical diff is *evidence of independence* and *evidence of a broken fixture*, and nothing in
   the output distinguished them. Before believing a null result, prove the instrument moved.
2. **Build a fixture out of the thing under test, not out of a convenience.** Symlinks were the obvious
   way to make a restricted clone set without copying gigabytes, and they are precisely what the code
   under test resolves away.
3. **An expectation trained on one subject does not transfer to another.** Five dirty corpus inputs made
   "it will be rotten" feel like knowledge; applied to the tooling it was wrong three iters running, in
   the same direction each time.
