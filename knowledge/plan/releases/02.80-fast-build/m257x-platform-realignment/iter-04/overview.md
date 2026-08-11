---
milestone: M257x
iter: 04
iteration_type: tik
status: closed-fixed
created: 2026-07-31
---

# iter-04 — the first bring-up on this host

**Active strategy reference:** `TOK-01` ("instrument first, then follow"), still **step 1** — *unblock the
gate's instrument*. Every clause except (3) and (5) is measured **through a bring-up**, and no bring-up has
ever run on this box.

## Step 0 — re-survey (mandatory, run before targeting)

TOK-01's `Next-tik direction` named odysseus and a FATAL rext-pin mismatch. Both are stale (D-v28-15 machine
move; iter-02 found the pin already clean; iter-03 confirmed the guard MATCHING live). The re-survey measured
current state rather than inheriting it:

| probe | reading |
|---|---|
| `stack-demo/` clone set | **13 repos** + `clones.lock.json` + `clones.pin.json` — `HOST-M257x-stack-demo` genuinely DONE |
| `.agentspace/rext.tag` | `fast-build-m257x-iter-02`; consumption clone at `54bccf7` == that tag |
| target `.env` files | **all 5 ABSENT** — secrets never provisioned on this host |
| `docker ps` | Docker up, **0 containers** — no stack running, `demo-1` free |

So the residual before clause 1 is exactly what iter-03's addendum said: **provision secrets, then bring up.**

## Phase 0d pre-flight — PASS (and it changed nothing, which is the point)

Built `stacksecrets` from the pinned consumption clone and ran the demo-aware coverage check against
`.agentspace/secrets` **before** committing to the long operation:

    Overall:  65.5%
    Critical: 100.0%     <- the gate condition; exit 0

Critical coverage is the gate; it is met. The 65.5% overall is non-critical breadth (LiveKit, Stripe,
ElevenLabs, Sentry, FontAwesome, several AI providers) — features a demo does not need to come up. Recorded
so a later failure in one of those surfaces is read as *a known-absent optional key*, not a new defect.

## Cluster / target identified

The bring-up itself. Nothing else in the milestone can be measured until it runs once.

## Hypothesis

With a complete clone set, a matching pin, Docker present, and critical secrets at 100%, `demo-up 1` will
progress materially further than any prior attempt in this milestone (all three of which aborted before
reaching a container). Whatever it hits first **is** the finding — this is the first honest reading of how
much of rext still runs against platform origin HEAD, which is `overview.md` Open Question 5.

## Expected lift

Clause 1 requires **3 consecutive** cold cycles, so a single bring-up **cannot** move the 0/5 clause count.
The expected lift is diagnostic: the first measured failure surface, or a working stack.

## Phase plan

1. Provision secrets values-blind (dry-run, then real) — `stacksecrets provision`.
2. `demo-up 1` — no `--public-host` (no tailscale on this box: `HOST-M257x-toolchain`).
3. Read what happens. Route findings per §7/§8 of the protocol.

## Escalation conditions

- A guard that **refuses** with a decision only the user can make → user-blocker.
- A failure whose fix is a platform-repo edit → route forward; the v2.8 zero-platform-edit constraint holds.
- A failure inside rext with an in-scope fix → land it (Fate 1) if it fits the iter; else route with a handler.

## Acceptable close-no-lift outcomes

A documented first-failure characterization with its mechanism named and cited satisfies the protocol even
though the clause count stays 0/5. Per `platform-alignment.md` §9, *"whenever a bring-up fails oddly, check
signals 2 and 3 before debugging the tooling"* — a finding that the platform moved again is a first-class
result, not a miss.
