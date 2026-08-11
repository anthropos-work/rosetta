**Type:** tik, under `TOK-05`. Resolves iter-89's `user-blocker` with the user's decision, re-derived.

# iter-90 — the demopatch asymmetry: journal the observed pre-state, and test the CONJUNCTION

## What was measured before anything was touched

iter-89 left the demo clones dirty **on purpose**, so the defect would reproduce without being re-created.
That state was spent only after it had bought a test, in the order the user mandated.

**The reproduction, live, on the shipped manifests — two commands:**

```
$ demopatch status stack-demo --manifest next-web-back-to-cockpit.yaml
patched
$ demopatch revert stack-demo --manifest next-web-back-to-cockpit.yaml
demopatch: revert REFUSE: target sha256=a28808e0… is neither pre nor post — manual drift; refusing to guess.
```

`status` says **patched**; `revert` says **neither pre nor post**. Both are correct, and together they are
the whole defect. Reproduced identically on all three next-web manifests (`next-web-studio-url`,
`next-web-public-website-url`, `next-web-back-to-cockpit`).

**The residue was independently re-derived, not taken on trust.** The captured diff
(`evidence/dirty-clone-residue.diff`, 45 lines) is exactly two hunks in `urls.ts` (the PUBLIC_WEBSITE_URL and
STUDIO_URL patches) and one in `NavbarTop.tsx` (labelled `// demo-patch: next-web-back-to-cockpit`).
**Nothing else. No human work.** The user's own inspection is confirmed.

## The three steps, in the mandated order

### 1. The conjunction test came FIRST, and was shown RED

`TestGuardConjunctions` — six tests. Run against the **unfixed** code:

| test | pre-fix |
|---|---|
| `g2_self_heal_THEN_g5_revert_leaves_the_clone_git_clean` | **RED** |
| `g2_self_heal_THEN_g5_revert_leaves_no_journal_residue` | **RED** |
| `g4_idempotent_reapply_THEN_g5_revert_still_restores_observed_pre` | **RED** |
| `a_CHAINED_pair_on_one_file_applies_LIFO_and_reverts_clean` | **RED** |
| `a_hand_edit_with_NO_journal_still_REFUSES` *(negative control)* | GREEN — and must stay so |
| `MUTATION_blinding_the_journal_…` *(mutation control)* | — |

**Mutant signature:** `revert REFUSE: target sha256=… is neither pre nor post — manual drift; refusing to
guess.` — 4 of 5 RED, the negative control correctly unmoved.

Why the pair was invisible before: `test_g2_drifted_sha_with_an_INTACT_anchor_SELF_HEALS` applies onto a
drifted base and **never reverts**; `test_g5_revert_on_drifted_refuses_without_force` reverts a **manually**
drifted target that was never applied. Neither composes the two. Both passed. The suite was 52-for-52 green
while the mechanism was broken.

### 2. The (b) fix

`apply` journals the observed pre-image before writing a byte; `revert` consults that journal first and
restores exactly those bytes. Revert no longer depends on a baseline that is guaranteed to go stale.

Three properties, each asserted by a test rather than asserted in prose:

- the journal lives in the **workspace root**, never inside a clone — a journal that dirtied the clone would
  defeat the promise it exists to keep;
- it is **consumed** on a successful revert and its directory removed once it empties — no per-apply leak;
- **no journal means no guessing** — an un-journalled drifted target still hits the baseline comparison and
  still refuses. Journaling made revert *exact*, not *blind*.

The mutation control is **permanent and test-side**: it rebuilds `demopatch` with `_journal_read` neutered
and asserts the battery goes RED again with the original signature. A first cut put that mutation behind an
env flag inside the tool; it was removed — **a production code path that exists only for its own test is a
backdoor.**

### 3. The clean, and the limitation named

Cleaned via the tool's own `--force-pristine`. Both files restored; `git status` empty.

**Journaling cannot retroactively revert those two files** — they were applied before any journal existed, so
there was nothing to restore *from*. The one-time manual clean is the honest consequence of (b), recorded as
a limitation rather than presented as the fix working. The mechanism now WARNs when `apply` meets an
already-patched target with no journal entry, so the condition is self-describing at the moment it is created
instead of costing two iterations to rediscover.

## The live proof — the fix run against the real drifted clone

The condition that stranded it, re-run end to end:

```
apply  next-web-studio-url        → SELF-HEALING (sha DRIFTED) … applied
apply  next-web-public-website-url → applied            [journal: 2 entries]
revert next-web-public-website-url → demo clone left git-clean     ← LIFO, as the trap does it
revert next-web-studio-url         → demo clone left git-clean
git status → (empty)               .demopatch-journal → removed
```

The chain's recomputed `post_sha256` is `ebab9e7e…` — **exactly** the sha the failing tests reported when the
clone was dirty. That closes iter-89's diagnosis: the dirty state was this chain, fully applied and never
reverted.

## The 2 residual suite failures changed MEANING and were deliberately NOT re-pinned

`TestRealManifest`'s two live-clone assertions failed at open and still fail — but not the same failure. The
live `urls.ts` sha moved `ebab9e7e…` (chain fully applied) → `6d6292ef…` (pristine, genuinely drifted from
the manifest's `0d4c3790…` baseline). The artifact was masking the real signal; the real signal is now
visible.

Not re-pinned, per iter-88's standing instruction (*adjudicate before touching*). And there is a prior
question: under M217 **the anchor is the contract and the whole-file sha is only a baseline**, so an
assertion that a shipped `pre_sha256` still matches a live, persistently-updated clone is asserting the
property the design deliberately stopped requiring. Routed as `CHECK-M257x-iter90-realmanifest-baseline`.

## The wider suite: 1054 tests, 6 failures, 0 of them mine

| failures | cause | mine? |
|---|---|---|
| 3 `test_migrate_race_live.*` | need a LIVE Postgres container; no stack is up | no — environmental |
| 2 `test_demopatch.TestRealManifest.*` | live `next-web-app` clone vs stale manifest `pre_sha256` | no |
| 1 `test_ant_academy…round_trip_on_the_real_next_config` | live `ant-academy` clone vs stale `pre_sha256` | no |

Checkable, not asserted: this iter's rext diff is exactly two files, and the ant-academy patch runs on a
**different vehicle** (`ant-academy.sh`, **0** `demopatch` references).

**The third instance turns this into a class.** Three manifests across two independent patch vehicles fail
the same way, because these clones are persistent and drift while their pinned baselines do not — the same
rot M217 removed from *apply*, still live in the *test layer*. `CHECK-M257x-iter90-realmanifest-baseline` is
widened accordingly: not "re-pin two shas" but *should a test assert a whole-file baseline against a live,
persistently-updated clone at all?*, answered once for all three.

## The guard-family correction — the class is real, the number did not reproduce

The user's standing rule (*a guard result is invalid unless the clone was fetched first*) was applied and the
result must be reported precisely: **every `stack-demo/` clone was fetched; nothing moved** except the rext
consumption clone; `platform` was already at origin HEAD `0c91421`. The family reads **13 GREEN · 0 RED**
identically before and after, so the reported `10 GREEN · 3 RED` did **not** reproduce here.

**The class is real anyway, and it is worse than a stale clone.** `platform_alignment_guard` resolves
citations at `origin/main` → `HEAD` → and then **silently at the worktree**, never inspecting which it used:

| reference | verdict |
|---|---|
| `CITE_REF=auto` (refs present) | **GREEN** — 90 resolved, 0 unresolvable |
| `CITE_REF=worktree` (the stale-clone fallback) | **RED** — 8 findings, 4 unresolvable |

And `unresolvable` is **printed but never graded** — the only positive control is `subject_checked == 0`, so
*partial* unresolvability is folded into GREEN. That is exactly the three-valued discriminator failure
iter-79 and harden pass 19 established: yes / no / **cannot-tell**, with cannot-tell laundered as no.
Routed to iter-91 as `FENCE-M257x-iter91-clone-freshness`, fix shape already measured.

## Close — 2026-08-05

**Outcome:** the demopatch asymmetry is repaired by journalling the observed pre-state (the user's option
(b), re-derived and upheld with one correction to its rationale), proven live on the real drifted clone that
stranded, and fenced by a 6-test conjunction battery with both a negative and a permanent mutation control.
The two dirty files are clean, and the limitation that they could not be journal-reverted is recorded rather
than hidden.
**Type:** tik
**Status:** closed-fixed — all three declared steps landed, in the mandated order
**Gate:** NOT MET — **4 of 5, unchanged.** No reading was taken this iter; clause 5 is graded only by a
reading that returns zero, and the corpus/guard preconditions for taking one are iter-91's scope.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (1 tik this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D-M257x-90-1 … D-M257x-90-6 (iter-90/decisions.md)
**Side-deliverables:** none — the `demopatch-spec.md` rewrite is part of landing the fix, not a side fix.
**Routes carried forward:**
- `FENCE-M257x-iter91-clone-freshness` → iter-91: make the stale-clone reading `UNMEASURED`, and grade
  `unresolvable`/`worktree(no-ref)` instead of printing them.
- `CHECK-M257x-iter90-revert-idempotency` → iter-91: the `RETURN` trap reverts manifests that were never
  applied; exit 1 into `/dev/null`. Adjudicate whether revert's no-op should be decided by `_classify` (the
  anchor) rather than a whole-file sha — the same asymmetry, one level further.
- `CHECK-M257x-iter90-realmanifest-baseline` → iter-91: adjudicate whether the two live-clone sha assertions
  should be re-pinned, re-scoped to the anchor, or retired.
- The 7-guard **conjunction-pair enumeration** (G1–G7) → iter-91.
- The **M810 `cms`-vs-`jobsimulation`** split sweep → iter-92.
**Lessons:**
- **When a test fails, check whether it encodes a real requirement before changing the code to satisfy it.**
  A double-revert test failed against the fix; the premise was checked against `up-injected.sh:741` and found
  false (one revert per trap, then `trap - RETURN`). It was replaced by the **chain** pair, which is on the
  shipped path and strictly more valuable. A design bent to satisfy a wrong test is worse than either.
- **A documented defect can be documented as benign and still be live.** This exact asymmetry was already
  written down in `demopatch-spec.md`, ending *"it is not currently harmful."* The reasoning behind that
  clause was true of the `app` build-scratch clone and false of the persistent `next-web` clone the paragraph
  was about. It then cost iter-88 and iter-89 two full iterations. **A known-and-dismissed finding deserves
  the same re-derivation as a new one** — the dismissal is a claim too.
- **Guards must be tested in PAIRS.** Generalised into the protocol doc (§8) and `demopatch-spec.md`.
