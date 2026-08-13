**Type:** tik (under `TOK-01: instrument first, then follow`, step 3 — *land the fences, each watched going
RED, before trusting any green*).

# iter-16 — the two bring-up verdicts that grade themselves green

## What was measured before anything was written

The routed-forward findings RF-4 and RF-1 (`hardening-ledger.md`) were both re-read at source before being
treated as real, and **RF-4's scope was wrong in the ledger**. It is filed as a `dev-stack` finding; it is
not scoped to dev. `demo-stack/up-injected.sh` calls `dev-stack/dev-setdress.sh` **verbatim** via
`--stack-type demo` (the deliberate M20 convergence invariant — `test_reuses_the_dev_setdress_engine_via_stack_type_demo`
exists to keep a demo-only fork from being written). So the `*)` arm of `dev-setdress.sh`'s replay-rc case
statement is not *like* the site that produced gate clause 1's false greens. **It is that site.**

That reframes the iter. RF-4 stopped being "a dev-stack tidy-up" and became the mechanism behind the three
compromised clause-1 verdicts, and it was landed as such.

## RF-4 — the set-dress verdict was a constant

`dev-setdress.sh` accumulates a per-surface summary (`SNAP_SUMMARY`) and, since fix16, that field has been
**honest**: it read `directus=skipped(error)` on every one of clause 1's three cold cycles. The problem was
never the field. It was the **sentence the field sat inside** and the **exit code that followed it**:

```
==> demo-1: set-dressed (content:local-content, snapshot: taxonomy=replayed … directus=skipped(error), … seeded).
```
…followed by `exit 0`.

Three readers were misled by one line. The operator's eye reads the verdict word, not the sixth field of a
parenthesis. The caller reads the exit code — and `up-injected.sh:2262` *already* wrapped the call in
`if ! …; then log "⚠ set-dressing did not fully complete…"`, so **the caller-side handling this needed
already existed and had been dead code for its entire life**, unreachable for the one state it was written
for. And the gate read `autoverify green:true`, which is a different instrument entirely and was measuring
something else.

Three changes, and the shape of them is the point:

1. **rc 4 and rc 5 are DOCUMENTED degradations; anything else is an unclassified ERROR.** fix16 gave the
   unprovisioned stack rc=4 and the cache miss rc=5, each with a named operator fix printed alongside. Those
   still report `set-dressed`, and correctly — the stack really is dressed to the extent the environment
   allowed. Any *other* rc means stacksnap tried, failed, and we do not know why. That is not the same
   finding and no longer shares a verdict: it prints `replay FAILED`, is counted, and the surface is
   **named**.
2. **The verdict word is now a function of the per-surface outcomes**, not a constant. `SNAP_ERRORS != 0`
   prints `⚠ set-dress INCOMPLETE — N surface(s) FAILED to replay: <names>`.
3. **A distinct exit code, 3.** `die` already exits 1 ("aborted — nothing seeded"). 3 means *"the pass ran to
   completion and the seed floor landed, but a surface failed for a reason we could not classify."* The
   distinction is load-bearing: the seed floor is the M20 atomicity contract and it still runs, so this is a
   **report, not an abort**, and `up-injected.sh` keeps the M13/M18 non-fatal contract by warning.

The caller was given a matching arm that names rc=3 distinctly rather than folding it into the generic
warning — whose named fixes (*provision the stack* / *capture the cache once per release*) are precisely the
irrelevant advice for a replay error — and says the consequence out loud: **the failed surface is EMPTY,
which downstream reads as "the content layer 403s", not as a replay error.** That sentence is the one that
would have saved iter-15 its permissions detour.

### Three tests were asserting the tolerant verdict on the intolerable case

The most interesting part of RF-4 was not the script. Three existing tests used `stacksnap exit 1` and
asserted `returncode == 0`, and **their own comments show they knew**:

- `test_cache_miss_is_non_fatal_seed_still_runs` — *"here a generic exit 1 — fix16 reserves 4/5 for the
  classified cases"*. It names 1 as the unclassified code in the same breath as asserting the classified
  verdict on it. Re-pointed to rc=5, which is what its name always claimed to be testing.
- `test_demo_seed_is_the_atomicity_floor_after_a_replay_miss` — same re-point; the atomicity proposition it
  exists for is unchanged and holds on both paths.
- `test_replay_error_is_tolerated_seed_is_the_floor` — the load-bearing one. Its stated rationale was
  *"cache-miss vs firewall-error are the same exit 1 to the script; both degrade gracefully."* **That was
  true when it was written and stopped being true at fix16**, which split the graceful cases out into 4 and
  5. After that split, exit 1 is specifically neither of them. Renamed to
  `test_replay_error_still_runs_the_seed_floor` — the half of its claim that was always correct.

This is §8 rule 3 exactly: **a fence that pins the current shape of the drift converts the bug into a
contract.** Three cold cycles were graded green over `directus=skipped(error)` with all three of these
passing. The suite was not silent about the defect; it was *arguing for it*.

The `never-capture-on-a-degraded-path` prod-safety test was deliberately **kept** at exit 1 rather than
re-pointed: the unclassified error is exactly where a naive auto-repair would be most tempted to reach for a
capture, so it is the stronger place to pin that invariant. Only its expected exit code moved.

## RF-1 — the dev migrate loop discarded the evidence it pointed at

`migrate-demo.sh:150-177` has had output capture + exit classification + `mig_fail` + a refusal to report OK
since **M215 F8**, which fixed it *because* the old grading masked a total migration failure (atlas absent
or erroring → 0 tables → every seeder fails downstream, far from the cause). **The dev twin never got it**,
and it was strictly worse in two ways:

- Every non-zero atlas exit was graded `"$r had migration warnings (non-fatal — see atlas output)"` —
  pointing the operator at an atlas output that `>/dev/null 2>&1` had **thrown away**. The diagnosis did not
  merely go unclassified; it did not exist.
- The absent-clone branch logged `✗` and `continue`d **without recording a failure at all**, so the closing
  line `done — … the derived migration set applied.` was reachable after migrating **nothing**.

Ported whole. `atlas` exits 0 when there is nothing to do, so on the happy path every service still reports
`ok` byte-identically — the change is invisible until something is actually wrong, which is the correct
blast radius.

### The parity fence could not have caught this, and that is its own finding

`test_parity_with_migrate_demo_on_the_load_bearing_guards` exists to keep the two scripts in step. It was
green throughout, because it walks a **hand-maintained enumeration** of guard strings — and none of M215
F8's four guards were in it. A parity fence over a hand-maintained list is the milestone's own §2 class
(*a hand-maintained list that must track something that moves*) applied to the guards themselves. The four
M215 F8 strings are now in the list; the deeper shape is recorded as a route rather than silently absorbed.

And `test_absent_service_clone_is_skipped_not_fatal` had a **stale rationale**, which is why it was green
over the defect. It reasoned from an isolated cold-init proof that points `DEV_CLONES` at a tree with no
service dirs, and concluded an absent clone must be non-fatal. That inference stopped holding at **iter-02**:
once the migration set became *derived*, the cold proof's tree has no `platform/repos.yml` either, so
`MIG_PAIRS` is empty and the loop body is never entered — **the `continue` protects nothing there**. The two
propositions are not in tension and are now both pinned: the loop still continues (so one run names every
absent clone, not just the first) *and* the script refuses to report OK.

## Fences watched going RED — 11 mutants, all as predicted

Not one of these fences was trusted green. Every mutant was `bash -n`-checked before its verdict was read
(§8 rule 5 — *a mutation that doesn't compile isn't a passing test of your safety net*), and the unmutated
control was run **before and after** the pass so a restore failure could not be mistaken for a result.

| # | mutation | verdict |
|---|---|---|
| M1 | the unclassified rc stops being counted (`+ 1` → `+ 0`) | RED |
| M2 | the report prints but `exit 3` becomes `:` — a caller's `if !` still cannot see it | RED (×2 tests) |
| M3 | the failure is COUNTED but the surface is not NAMED | RED |
| M4 | rc=1 re-absorbed into the rc=4 documented-degradation arm | RED |
| M5 | atlas's output is `>/dev/null 2>&1`-discarded again | RED |
| M6 | an ABSENT clone stops setting `mig_fail=1` | RED |
| M7 | the "refusing to report a half-migrated dev stack" exit is removed | RED |
| M8 | `sd_rc = 3` → `= 99`: rc=3 falls through to the generic warning | RED |
| M9 | **no-op control** — a statement that changes nothing | **GREEN (survived, as required)** |
| M10 | `\|\| sd_rc=$?` removed: under `set -e` a failed set-dress aborts the bring-up | RED (×3 tests) |
| M11 | the `EMPTY` consequence dropped from the rc=3 message | RED |

M9 is deliberate and is the one that makes the other ten mean something: a fence that kills a mutant which
changed nothing is not measuring the mutant. It survived, as it must.

## The +1 the demo-stack suite caught — and why it was the right kind of failure

Landing RF-4 took `demo-stack` from its recorded 7 failures to **8**. The extra one was
`test_setdress_is_non_fatal`, and it is worth naming because the fix is not "update the string":

The test asserted `assertIn("if ! env ", self.BODY)` — it string-matched the **mechanism**, not the
**proposition**. Non-fatality was never in question and is unchanged; but `if !` collapses every non-zero
exit into one boolean, and telling rc=3 apart from everything else is the entire content of RF-4, so the
wrapper had to become `env … || sd_rc=$?`. A fence written against the proposition would not have moved.
Rewritten to assert what it actually cares about — the invocation's exit is *captured into a variable*,
never left bare, and never turned into `|| exit` / `|| die` — and scoped to the **extracted block** rather
than the whole file, because the surrounding script contains `|| exit` for genuinely fatal steps and a
whole-file `assertNotIn` would be asserting something false about a different part of the script (§8 rule 4,
the same trap iter-13's compose fence hit from the other direction).

Both sections are back on their recorded baselines exactly.

## Close — 2026-08-01

**Outcome:** The two bring-up verdicts that reported a state they had not measured now derive that state
from what they measured. The site behind gate clause 1's three false greens is fixed **and fenced**, so
iter-17's cold cycles are measured with an honest instrument instead of re-running the original mistake with
a better probe bolted on the side. Three existing tests were found asserting the tolerant verdict on the
intolerable case — the suite was arguing for the defect, not silent about it.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (2 of 5 clauses; clause 1 must be re-proven with the serving probe in the consumed tag — this iter is its precondition, not its proof)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n (no platform commit landed during this iter; trigger stays at occurrence 1 of 2) — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D-M257x-16-1 … D-M257x-16-4 (iter-local `decisions.md`), plus `HARDEN-CAP-ACCEPTED` recorded at milestone root.

**Metric:** clause 2 unmeasured this iter **by design** — the iter's declared expected lift was on clause 1's
*re-provability*, not on any clause's count. Suites: `dev-stack` **OK 125** (baseline 122; +3 net-new tests,
0 regressions) · `demo-stack` **7F of 1030** (baseline 7F of 1029; +1 net-new test, back to baseline
exactly after the `test_setdress_is_non_fatal` re-point).

**Side-deliverables:** none. Both lines were planned; the `test_setdress_is_non_fatal` re-point is part of
landing RF-4, not a separate discovery.

**Routes carried forward:**
- `CHECK-M257x-iter16-parity-fence-hand-maintained` — `test_parity_with_migrate_demo_on_the_load_bearing_guards`
  walks a hand-maintained guard list and was green throughout RF-1's four-guard drift. The list is patched;
  the *shape* (a parity fence that cannot notice a guard nobody added to it) is the milestone's own §2 class
  pointed at itself. → later tik.
- `RF-2 … RF-3`, `RF-5 … RF-12` — the remaining 10 routed-forward harden findings stand unchanged in
  `hardening-ledger.md`.
- Clause 1's re-prove (cold ×3 with `probe_directus_serves_content` in the consumed tag) → **iter-17**, and
  it now has its precondition.

**Lessons:**
1. **A fence's own comment can contain the refutation of the fence.** Two of the three re-pointed tests said,
   in their comments, that exit 1 was the *unclassified* code — and then asserted the classified verdict on
   it. When a test comment explains why the value it uses is the wrong one, read that as a finding, not as
   documentation.
2. **Check a routed-forward finding's SCOPE at source before trusting the ledger's filing.** RF-4 was filed
   under `dev-stack`; the file is shared with the demo path by an explicit convergence invariant, which made
   it the clause-1 site rather than a tidy-up. The ledger records where a finding was *found*, not
   everywhere it *reaches*.
3. **A no-op mutant belongs in every mutation pass.** Ten RED verdicts say nothing about whether the fence is
   discriminating until one mutant that changed nothing is shown to survive. (Already promoted to §8 rule 5
   at iter-11 as *a mutant that changes nothing is not a survivor*; this iter is its first use as a
   deliberate positive control rather than a disqualification rule — protocol addendum below.)
4. **When a fence breaks because the mechanism changed, ask whether it was ever asserting the proposition.**
   `test_setdress_is_non_fatal` failed for the right reason and was repaired by re-pointing it at the claim
   rather than at the new string.
