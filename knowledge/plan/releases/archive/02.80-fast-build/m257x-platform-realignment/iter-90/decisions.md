# iter-90 — decisions

## `D-M257x-90-1` — (b) re-derived and UPHELD, but not for the reason offered

The user chose **(b) journal the observed pre-state at apply; revert restores exactly it**, and asked for it
to be re-derived rather than assumed. It is upheld. The re-derivation changed one thing worth recording.

The argument *offered* for (b) was that it "keeps both G2 and G5 true at once." Measured, that is right but
imprecise: **(b) does not reconcile G2 and G5 — it removes the term they disagreed about.** G2 and G5 were
never in conflict over the *anchor*; they were in conflict over the **baseline**, which apply had stopped
consulting (M217) and revert had not. (b) wins because it deletes revert's dependence on a recorded
baseline, not because it arbitrates between two guards.

That distinction is load-bearing for the next person: it says the fix belongs in **revert's source of truth**,
not in a reconciliation layer between guards, and it predicts that (a) — inverting the anchor — would have
been a second consumer of the same missing information rather than a fix.

The three rejected options were re-derived too, and the user's reasoning holds:

- **(c) strict apply** would restore consistency by giving up self-healing. Refuted directly by this iter's
  own live run: applying to the real clone printed `sha DRIFTED … SELF-HEALING`, i.e. the shipped clone is
  drifted **right now**, so (c) would refuse on the normal path. That is the silently-refused-patch failure
  that shipped a 76 s members grid for four releases.
- **(a) anchor-inverse revert** re-derives the pre-state from the post-state. Ambiguous wherever the anchor
  text recurs or the applied region was itself touched — and the chain case makes that concrete: two
  patches on one `urls.ts`, where the second's input IS the first's output.
- **(d) accept + `--force-pristine`** makes a destructive `git checkout --` the routine path. Rejected as a
  default — but see `D-M257x-90-3`: it is exactly right as the *one-time recovery* for clones patched
  before the journal existed.

## `D-M257x-90-2` — the CONJUNCTION test came first, and the mutation control is permanent

Ordered as the user mandated: the test was written and shown RED **before** the fix, and the dirty clones
were not spent until it existed.

**Mutant signature, pre-fix (4 of 5 conjunction tests RED, negative control GREEN):**

```
revert REFUSE: target sha256=… is neither pre nor post — manual drift; refusing to guess.
```

The same signature was captured **live** on all three shipped next-web manifests before anything was
touched — `status` reporting `patched` while `revert` reported `neither pre nor post`, in the same breath.
Two commands is the whole reproduction.

The control is now **permanent and test-side**: `test_MUTATION_blinding_the_journal_makes_the_conjunction_
tests_fail` rebuilds demopatch with `_journal_read` neutered and asserts the cycle refuses again with the
*original* signature. A first cut put that mutation behind an env flag inside `demopatch` itself; it was
removed. **A production code path that exists only for its own test is a backdoor**, and this contract is
checkable without one.

## `D-M257x-90-3` — the limitation is stated, not worked around

**Journaling cannot retroactively revert the two files iter-89 left dirty.** They were applied before any
journal existed, so there was nothing to restore *from*. The one-time `--force-pristine` was the honest
recovery and it is recorded as such, not dressed up as the fix working.

The mechanism now says this out loud rather than leaving it to be discovered: an already-patched target with
no journal entry makes `apply`'s G4 no-op emit a WARN naming the consequence (*"`revert` will fall back to
the recorded baseline and may REFUSE; `revert --force-pristine` is the recovery"*). The failure mode that
cost iter-88 and iter-89 two full iterations is now self-describing at the moment it is created.

## `D-M257x-90-4` — the double-revert test was WITHDRAWN after measurement, not weakened after failure

An early conjunction test asserted that a second `revert` must be an idempotent no-op. It failed against the
fix: the journal is consumed on success, so a second revert has no state and falls through to the baseline
refusal.

The tempting move was to contort the design until the test passed. Instead the premise was checked, and it
was wrong: `up-injected.sh:741` reverts each manifest **once** from a `RETURN` trap that then does
`trap - RETURN`. Back-to-back revert with no intervening apply is **not on the shipped path**, so the test
was asserting a requirement that does not exist.

It was replaced by the pair that IS on that path and is strictly more valuable: **the chain** — two patches
on one `urls.ts`, applied studio→pubweb and reverted pubweb→studio, on a drifted base. That conjunction is
invisible to any per-guard test because the interaction is between two **invocations**, not two guards.

Recorded because the reasoning generalises: **when a test fails, check whether it encodes a real
requirement before you change the code to satisfy it.** A test can be wrong, and a design bent to satisfy a
wrong test is worse than either.

Routed forward, not dropped: `CHECK-M257x-iter90-revert-idempotency` — the trap does invoke `revert` on
manifests that were never applied (a refused apply still gets a revert), which exits 1 into `>/dev/null` and
`|| true`. Benign today; it means `demopatch.log` cannot distinguish "never applied" from "failed to come
off". Adjudicate whether revert's no-op should be decided by `_classify` (the anchor) rather than by a
whole-file sha — which would be the same asymmetry fix applied one level further.

## `D-M257x-90-5` — the guard-family "13 GREEN" was re-measured on a FETCHED clone and STANDS

The user's standing correction is that a guard result is invalid unless the clone was fetched first, having
found a `13 GREEN · 0 RED` baseline that re-read as `10 GREEN · 3 RED` after fetching.

Re-derived at open, and the finding must be reported precisely because the *number* did not reproduce while
the *class* did:

- Every clone in `stack-demo/` was fetched. **Nothing moved** except the rext consumption clone
  (`ac30b9b → 7844e97`); `platform` was already at origin HEAD `0c91421`.
- The family reads **13 GREEN · 0 RED · 2 could-not-check · 1 not-run**, identical before and after. So on
  *this* box the baseline is valid, and I could not reproduce the 3 RED.

**The class is real anyway, and it is worse than a stale clone.** `platform_alignment_guard`'s assertion F
resolves citations at `origin/main`, falling back to `HEAD`, and then — when neither ref resolves — silently
to the **worktree** (`cited_text`, provenance `worktree(no-ref)`). The guard never inspects that provenance.
Measured, the two references are different files and give different verdicts:

| citation reference | verdict |
|---|---|
| `CITE_REF=auto` (refs present) | **GREEN** — 90 resolved, 0 unresolvable |
| `CITE_REF=worktree` (the stale-clone fallback) | **RED** — 8 findings, 4 unresolvable |

So a clone that cannot see `origin/main` is graded at a different reference than the exit gate names, and is
told nothing. Worse, `unresolvable` is **printed but never graded**: the only positive control is
`subject_checked == 0`, so *partial* unresolvability is folded into a GREEN verdict. That is precisely the
three-valued discriminator failure iter-79 and harden pass 19 established — yes / no / **cannot-tell** — with
cannot-tell laundered as no.

**Routed to iter-91** (`FENCE-M257x-iter91-clone-freshness`), with the fix shape already measured: assert
ref-freshness before running and report `UNMEASURED` rather than a verdict, and grade `unresolvable > 0` and
the `worktree(no-ref)` fallback instead of printing them. Not done here — it is a second line of
investigation and this iter's scope was declared as three steps.

## `D-M257x-90-6` — the 2 residual suite failures CHANGED MEANING, and are not re-pinned

`TestRealManifest`'s two live-clone assertions failed at iter open and still fail. They are **not the same
failure**:

| | at iter open | after the clean |
|---|---|---|
| live `urls.ts` sha | `ebab9e7e…` | `6d6292ef…` |
| what it means | the chain was **fully applied** and never reverted | the clone is **pristine**, and the manifest's `pre_sha256` (`0d4c3790…`) is genuinely stale |

The first was an artifact masking the second. That `ebab9e7e…` is exactly the `post_sha256` this iter's live
chain-apply recomputed is the closing evidence for iter-89's diagnosis.

**Not re-pinned**, deliberately, per iter-88's standing routing instruction (*adjudicate before touching; do
not re-pin first*). And there is now a prior question: under M217 the anchor is the contract and the
whole-file sha is only a baseline — so an assertion that the shipped `pre_sha256` still matches a live,
persistently-updated clone is asserting the property the design deliberately stopped requiring. Whether that
test should be re-pinned, re-scoped to the anchor, or retired is an adjudication, not a sha edit.

Routed: `CHECK-M257x-iter90-realmanifest-baseline`.

## `D-M257x-90-7` — the wider suite: 6 failures, 0 of them mine, and the baseline drift is SYSTEMIC

The whole `demo-stack` suite — **1054 tests** — was run to check the fix against more than its own file.
**6 failures, and every one pre-existing.** Attributed rather than counted:

| failures | tests | cause | mine? |
|---|---|---|---|
| 3 | `test_migrate_race_live.*` | need a LIVE Postgres container; no stack is up | no — environmental |
| 2 | `test_demopatch.TestRealManifest.*` | live `next-web-app` clone vs stale manifest `pre_sha256` | no — `D-M257x-90-6` |
| 1 | `test_ant_academy…test_apply_revert_round_trip_on_the_real_next_config` | live `ant-academy` clone vs stale manifest `pre_sha256` | no |

The attribution is checkable, not asserted: this iter's diff touches exactly two files
(`demo-stack/patches/demopatch`, `demo-stack/tests/test_demopatch.py`), and the ant-academy patch runs on a
**different vehicle** — `ant-academy.sh`, which contains **0** references to `demopatch`.

**The finding is that the third instance makes it a class.** `D-M257x-90-6` framed the stale live-clone
baseline as a question about two assertions in one test file. It is not: **three manifests across two
independent patch vehicles** all now fail the same way, because every one of these clones is persistent and
drifts while the pinned baseline does not. That is the same rot M217 removed from *apply* — still live in
the *test layer*, which is exactly where it would be least visible.

So `CHECK-M257x-iter90-realmanifest-baseline` is widened: the adjudication is not "re-pin two shas" but
**"should a test assert a whole-file baseline against a live, persistently-updated clone at all, given the
anchor is the contract?"** — answered once, for all three, across both vehicles.

Note the shape this iter has now hit twice: **whole-file-sha thinking survives in the places nobody
rewrote.** M217 fixed it in apply; iter-90 fixed it in revert; it is still in the tests.
