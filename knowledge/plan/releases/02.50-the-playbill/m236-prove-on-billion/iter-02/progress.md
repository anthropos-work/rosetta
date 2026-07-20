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

---

## RESUMED — 2026-07-20 (post-user-resolution)

USER-BLOCKER-M236-01 resolved by the user (B1–B5, `../decisions.md` → RESOLUTION). Phase 0b re-verdicted
**RED → DISCHARGED / proceed-as-YELLOW** (`../spec-notes.md` § iter-02 RE-VERDICT). Steps 2–5 executed.

## Phase plan status (final)

| # | Step | Status |
|---|---|---|
| 1 | Prune `billion` build cache | ✅ DONE (109 GB reclaimed, 40 G → 139 G free) |
| 2 | `git push origin main` + 13 `playbill-*` tags | ✅ **DONE** |
| 3 | Re-pin `billion:.agentspace/rext.tag` → `playbill-m235-hardened` | ✅ **DONE** |
| 4 | Check consumption clone out at the tag | ✅ **DONE** |
| 5 | Verify tags-on-origin / clone content / pin guard / headroom | ✅ **DONE** |

### Step 2 — publish

Re-ran the mandated iter-01 B2 freshness check first: `git ls-remote --tags origin | grep -c playbill` →
**0**, `origin/main` → `1d97861` (the M228 commit). Unchanged since iter-01; the publish plan was still
valid. Re-verified all 13 tags as ancestors of `main` and the push as a clean fast-forward, then pushed:

```
main:  1d97861..60eff14  (fast-forward, 20 commits)
tags:  13 × [new tag]  playbill-m230-… → playbill-m235-scrub-fix
```

Purely additive — 0 name collisions, 0 force, 0 rewritten history.

### Steps 3–4 — re-pin + checkout

`billion:/home/devops/panorama/.agentspace/rext.tag`: `casting-call-m228-hiring-scope-fix` →
`playbill-m235-hardened`. Consumption clone `stack-demo/rosetta-extensions` fetched the new tags and
checked out `playbill-m235-hardened` (`60eff14`), cleanly from `1d97861`.

### Step 5 — verification (all 4 binary outcomes from `overview.md`)

1. ✅ All 13 `playbill-*` tags resolve on origin (fetched successfully by the host).
2. ✅ Consumption clone at `playbill-m235-hardened` and carries `stack-seeding/contentsession/` +
   `stack-seeding/presets/content-manifest.json`.
3. ✅ **M217 FATAL pin guard PASSES** — pin SoT (`playbill-m235-hardened`) == clone checkout
   (`playbill-m235-hardened`).
4. ✅ Host headroom: 139 G free (29 % used) — ample for Phase C's cold UI-tier rebuild.

**Bonus verification (unplanned, in-scope):** recomputed the gate denominator from the *published*
manifest on the host — **31 landable pairs confirmed** against the shipped artifact (13+13 simulation,
2+2 skill-path-legacy, 2 ai-labs presence-only/not-landable, 1 skill-path-new). Detail + the 33-vs-31-vs-18
arithmetic trap in `../spec-notes.md`.

## Close — 2026-07-20

**Outcome:** `billion` can now obtain the feature under test. The v2.5 tooling is published (main +
13 tags), the host is re-pinned to `playbill-m235-hardened`, the M217 pin guard agrees, and the gate
denominator of **31** is independently confirmed against the published artifact. Primary metric unchanged
at **0/31** — by design; this iter bought *reachability*, not numerator.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n (the one that existed is RESOLVED) — (5) cap-reached: n (1 tik this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** INCOMPLETE-EXIT-2026-07-20 (superseded by this close); D1 (Phase 0b re-verdict RED→discharged); D2 (denominator confirmed against shipped artifact)
**Side-deliverables:** `billion` Docker build-cache prune — 109 GB reclaimed, free 40 G → 139 G. Host maintenance, no repo touched; recorded separately so it does not colour the close grade.
**Routes carried forward:**
- Phase C (cold bring-up on `billion` at the new tag) → **iter-03**, handler `PHASE-C-M236-iter03-cold-bringup`.
- B4/B5 doc work (`coverage-protocol.md` `skipPaths` reversal + amendment; content-stories backfill of `coverage-protocol.md` + `playthroughs.md`) → **a later iter, paired with the harness that needs it** (Phase H), handler `DOC-M236-iterTBD-protocol-backfill`. Not done here: amending the protocol before a live render exists would repeat exactly the author-it-blind mistake USER-BLOCKER-M235-02 identified.
**Lessons:**
- **Stopping on the Phase 0b RED was correct and cheap.** The publish itself was never in doubt; what the audit questioned was the *destination*. Resuming cost one round-trip; proceeding would have spent 3–4 iters of Phase-L work against a gate that could not be satisfied, with the failure mis-attributed to seeding defects. Generalizable: when a pre-flight audit attacks the *gate* rather than the *code*, the cost of stopping is one user round-trip and the cost of continuing scales with the milestone's remaining length.
- **Verify a denominator against the shipped artifact, not only the source-of-truth preset.** Reading the published manifest on the host cost one command and independently confirmed 31 — and made the 33 (naive) / 18 (product-level) miscount modes concrete rather than theoretical.
