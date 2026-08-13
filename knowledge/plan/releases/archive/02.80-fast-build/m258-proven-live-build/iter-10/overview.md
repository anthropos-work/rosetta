---
milestone: M258
iter: 10
iteration_type: tik
status: closed-fixed
opened: 2026-08-12
---

# M258 iter-10 — the census walks the demo's ephemeral clone, and calls it rext

**Type:** tik · **Active strategy:** `TOK-01` (bootstrap). This milestone's claim is that a stack comes
up **and proves itself**; the trustworthiness of the instruments doing the proving is therefore M258's
business, which is why M257x routed *"the whole-section instrument is partially blind on this host"*
(cluster 2) and *"fences that are not RED-proven"* (cluster 3) here.

## Step 0 — re-survey before targeting (mandatory)

`TOK-01` step 4 (the gateable campaign) remains the milestone's only outstanding gate work, and it is
**armed, not stale**: `autoarm-campaign.sh` has been sampling every 15 s since 09:40:37Z against a fresh
`campaign-iter09/` dir at the `fast-build-m258-iter-09` pin. Measured at 09:46:42Z: **`load1 21.28`**,
minimum since arming **17.86**, against a threshold of 5.0. No window; nothing to re-run.

Target substituted within the strategy, from iter-09's own routing:
**`ROUTE-M258-iter09-literal-ratchets-scan-the-demo-clone`**. Needs no host window.

## Cluster / target identified

iter-09 measured the three literal ratchets against a pristine `git archive HEAD` extract and found the
working-tree excess was **not rext's**:

| ratchet | pristine HEAD | working tree | from `stacks/` |
|---|---|---|---|
| DOCSTRING | 248 | 258 | **+10** |
| COMMENT | 236 | 237 | +0 |
| TEST_MODULE | 657 | 672 | **+9** |

The source is `demo-stack/stacks/demo-1/clones/app/studio/**` — **the platform's own Python, inside a
demo's ephemeral clone**. Measured scope: `demo-stack/stacks/` is **264 MB / 4,560 files / 92 `.py`**,
against **220 `.py`** in rext proper. Roughly **30 % of the Python this census reads is foreign**, and
both `stacks/` dirs (`dev-stack/stacks/` too) are gitignored scratch by their own sections' rules
(`dev-stack/.gitignore:2`, `demo-stack/.gitignore:8`).

**This is the third consumer of one root cause** — after `test_decommissioned_instruction_guard` and
`test_fence_provenance::test_the_escape_accepts_and_records`, both open since iter-03/iter-06 and both
described as *"fires on any box that has ever run a demo."*

## Hypothesis

The repo already solved this exact problem once and enforced the solution in **one walker only**.
`derivation_registry.census_pruned` (M257x harden pass 69) exists because an environment at
`stack-core/.venv-check/` added **+43 / +42** to two ratchets *sitting exactly at their ceilings*, and
its docstring records the lesson as *"the runner must not live inside its own subject."* A per-stack
clone dir is the same class: **a tree the repo creates, inside itself, that is not its subject.**

So the fix should be a **one-name addition to the existing checked-in registry** (`_CENSUS_SKIP`),
not new machinery — and it should be component-exact, matching the two ephemeral `stacks/` dirs and
nothing else in the tree (verified: no other directory named `stacks` exists).

## Expected lift

**No movement on the composed p50** — again an instrument iter. The deliverable is that the three
ratchet figures become **measurements of this repo** rather than of whatever a demo last cloned, so
`RATCHET-M257-literal-ceilings-breached` can be graded honestly. Expected: working-tree figures fall to
the pristine-HEAD values (**248 / 236 / 657**), which is the falsifiable prediction.

## Phase plan

- **Phase A** — add the prune, prove it against the pristine-extract prediction in **both** directions
  (foreign rows gone; rext rows unchanged, so it prunes the clone and not the subject).
- **Phase B** — fence it, including the negative control (a planted foreign file is pruned; a real rext
  file at a similar depth is not).
- **Phase C** — check the other two consumers of the same root cause and report honestly whether this
  fix reaches them. **Do not claim it does without running them.**

## Escalation conditions

- The prune changes any count that is **not** attributable to `stacks/` → stop; over-pruning a census is
  strictly worse than the pollution, because it fails GREEN.
- A ratchet needing its ceiling *raised* to pass → never. Pay down or route.
- The campaign fires mid-iter → let it run; it consumes the immutable pinned tag and cannot be
  disturbed by authoring-clone edits. **Do not re-pin while a campaign is running.**

## Acceptable close-no-lift outcomes

- The prune lands and the ratchets are **still** breached on the pristine numbers (248 > 240,
  657 > 653). That is expected — this iter fixes the *measurement*, not the debt — and saying so
  plainly is the deliverable.
