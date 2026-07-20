# iter-02 — tik: Phase P (publish + re-pin)

**Type:** tik
**State:** **IN-FLIGHT — NOT CLOSED.** Halted mid-Phase by a blocking Phase 0b verdict. No `## Close`
section and no `iter(M236/02):` commit by design (Phase 4 Step 0: budget/blocker-exhausted iters exit as
user-blocker *without* closing; `/developer-kit:harden-mstone-iters` Phase 1 step 2 therefore correctly
ignores this iter).

## Phase plan status

| # | Step | Status |
|---|---|---|
| 1 | Prune `billion` build cache | ✅ **DONE** |
| 2 | `git push origin main` + 13 `playbill-*` tags | ⛔ **NOT EXECUTED** — halted |
| 3 | Re-pin `billion:.agentspace/rext.tag` → `playbill-m235-hardened` | ⛔ not reached |
| 4 | Check consumption clone out at the tag | ⛔ not reached |
| 5 | Verify tags-on-origin / clone content / pin guard / headroom | ⛔ not reached |

## Step 1 — host capacity precondition: DONE

`docker builder prune -af` on `billion`:

```
before:  /dev/sda1 193G 153G  40G  80% /      Build Cache 588 entries  109GB (107.6GB reclaimable)
after:   /dev/sda1 193G  55G 139G  29% /      Build Cache   0 entries    0B
```

**109 GB reclaimed; free space 40 G → 139 G (80 % → 29 % used).** The host now has ample headroom for
Phase C's cold UI-tier rebuild. This was the one step of iter-02 that is *independent* of the audit
verdict — it is host maintenance, touches no repo, and is beneficial regardless of how the blocker
resolves. It is recorded as a landed side-deliverable, **not** as grounds to grade the iter.

## Why the iter halted before step 2

The Phase 0b `audit-kb-fidelity` verdict returned **RED (blocker)** while step 1 was running. Phase 0b is
explicit: *RED → blocked; fill blind areas before proceeding*, and a sub-agent audit-RED is a critical
decision that stops the loop and wakes the user.

Critically, this is **not** a case where the block is merely procedural and the work is obviously still
right. The audit found that **M236's exit gate, as written, contains clauses that cannot be proven or
measured** — so executing the publish and marching into Phase C would have driven the milestone toward a
gate it cannot satisfy, and the failure would have surfaced 3–4 iters later attributed to the wrong cause.
Full findings in [`../spec-notes.md`](../spec-notes.md) § Pre-flight audits and
[`../kb-fidelity-audit.md`](../kb-fidelity-audit.md).

The push itself remains **safe and ready** — verified at iter-01: 13 tags all ancestors of `main`, clean
fast-forward `1d97861..60eff14`, 0 name collisions on origin, `stack-seeding` 16/16 packages green. Nothing
about the publish is in doubt; what is in doubt is **what M236 is publishing toward.**

## Metric

**0 / 31, unchanged.** Expected — this iter was a reachability precondition, never a numerator move.
Reachability itself is also unchanged: the tooling is still unpublished, so `billion` still cannot obtain
the feature under test.
