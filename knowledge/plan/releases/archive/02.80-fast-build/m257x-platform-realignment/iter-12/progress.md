---
milestone: M257x
iter: 12
---

# iter-12 — progress

**Type:** tik

## Phase 1 — `FIX-M257x-vmram-gib-unit`

Opened first because it prints on every bring-up this iter was about to run. Both halves of the pre-flight
were wrong about units, and one of the two defects **I claimed** was refuted by my own mutation battery
before it shipped.

**What was real.**

1. **It reported the floor, not the measurement.** `gib=$(( bytes / 1024 / 1024 / 1024 ))` turned this box's
   measured `12528664576` B — **11.67 GiB** — into the line *"Docker VM RAM = 11 GiB"*. An operator cannot
   reconcile `11` with anything Docker Desktop displays, and it printed on every bring-up.
2. **The remediation named a value that cannot clear the floor** — the worse half. It said *"raise the VM to
   12 GB"* while asserting **12 GiB**. Docker Desktop's memory slider is **decimal GB**; setting exactly the
   documented 12 GB yields ~11.2 GiB, which is *under* the floor. Follow the instruction to the letter and
   the warning never clears, with nothing on screen explaining why. **The doc says one unit and the code
   asserts another** — `platform-alignment.md` §2, in our own tooling.

The message now names both units and states the floor in the slider's own, **derived** from whatever
`DEMO_VM_MIN_GIB` is in force (a second hand-written constant is the defect this milestone exists to end):
`20 GiB` → *"about 21.4 GB on the slider"*, mutation-verified.

**What was NOT real, and how it was caught.** The draft also claimed the floored *comparison* was a defect
(*"floor(12.9)=12 also passes a 13 GiB floor"*). Mutation `N2` reverted the comparison to the floored form
and the test **stayed GREEN** — because for an integer floor `m`, `floor(x) >= m` **iff** `x >= m`. The
comparison was always correct. Comparing bytes is a readability change, not a repair.

That correction is written **into the code comment**, not silently dropped: a comment claiming a fix that
fixed nothing is how a false contract gets pinned, which is §8 rule 3's whole subject. The test that pinned
the false claim was replaced by one pinning what nothing had tested — **the boundary itself**: exactly at the
floor is OK, one byte under warns. Both halves mutation-verified (`-gt`, and an always-OK compare).

**Measured:** 5 tests, 5 mutants RED (each parsing, each collecting exactly 1) + an unmutated control GREEN.
`demo-stack/tests/test_frontend_build.py` + `test_tooling.py` **275/275**.

**The predicted cost, paid.** The fix added 23 lines to `up-injected.sh` and restaled **14** `file:line`
citations in `demo-up-defaults.md` — exactly the class iter-05 hit. Repaired with the guard's own `--fix`,
then re-run: *"OK — the defaults table and the parsers agree, both directions."*

**An inherited item closed as REFUTED rather than fixed.** `DOC-M257x-claude-md-knob-count` said `CLAUDE.md`
claims 27 env knobs where the parsers expose 30. `CLAUDE.md:307` and `corpus/ops/demo/README.md:153` **both
already say 30**. The only surviving "27" is `coverage-protocol.md:1056`, and it is a **historical account of
what v2.5's close corrected** — accurate as written. Nothing to fix.

## Phase 2 — clause 1: cold cycles

_(in progress — see the Close section for the measured result)_

### Cycle 1 — RED, and the fix had been applied to the wrong twin

`demo-down 1 --purge` → `demo-up 1` failed on directus `exit(1)`. The root cause was not new: **iter-05's
cold-start fix existed, on the DEV twin, and had never been applied to the DEMO twin it was actually
measured on.** Its test passed because it tested the twin that was fixed — the "check that reports without
measuring" class again (D3). Fixed and fenced; rext `fast-build-m257x-iter-12b` (`9a4bf35`), both `12a` and
`12b` verified on origin, `.agentspace/rext.tag` and the `stack-demo` consumption clone re-pinned.

### Cycle 1 (second attempt) — killed at 60 s, because the ref was wrong

Before letting the rebuilt cycle stand, the gate's own wording was checked against the clone: *"against
platform @ **origin HEAD** (never a pinned pre-drift commit)."* `stack-demo/platform` was **3 commits
behind** — and the bring-up's freshness check had not caught it, because **it compares the checkout against
`clones.pin.json`, not against origin**. It printed `PIN-DRIFT` naming a pin that was staler still
(`28c5f0d`), which reads like a stale *pin* rather than a stale *clone*. §3 of the protocol doc, exactly:
pinning disables drift detection.

The cycle was killed rather than run to completion — 18 minutes measuring the wrong ref buys nothing.

### What origin HEAD actually did: it deleted the GraphQL router

`2adcf71` (2026-07-31 15:58, **during this milestone**) merged *"drop the WunderGraph router; point local dev
at backend."* The `graphql` service, its `repos.yml` entry and its clone are gone; GraphQL is now served
**directly by `backend`** at `:8082/graphql/query` — note the **path** change, `/graphql` → `/graphql/query`,
which is the half that a hostname-only re-point would silently miss. `.env_example` even leaves a note that
`NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` "is historical."

**The `graphql` *profile* survives** on sentinel/backend/jobsimulation/cms/storage/roadrunner/gotenberg — so
`COMPOSE_PROFILES=graphql` still selects the backend tier and nothing about the profile wiring warns. Only
the *service* vanished.

**Measured, not inferred** (`evidence/iter-12-router-drop.md`):

1. `gen_injected_override.py` against the origin-HEAD compose → **RC=0**. It does not notice. It still emits
   `depends_on: graphql` and two `WUNDERGRAPH_SSR_ENDPOINT=http://graphql:8080/graphql`.
2. `docker compose config` on base+override → **RC=1**:
   `service "hiring-app" depends on undefined service "graphql": invalid compose project`.

So **clause 1 is not attemptable at origin HEAD**: the bring-up dies at project validation, before a single
build. No number of cycles at this ref could have produced a green one, and a cycle at the *stale* ref would
not have counted.

## Close — 2026-08-01

**Outcome:** clause 1 stays **0 of 3** — and the reason is now measured rather than suspected: origin HEAD
deleted the GraphQL federation router mid-milestone, so `docker compose config` rejects the demo project
outright (RC=1, undefined service `graphql`). Two real fixes landed on the way (the VM-RAM pre-flight's
unit/remediation defect; iter-05's cold-start fix finally applied to the DEMO twin), plus 14 restaled
citations repaired and one inherited item closed REFUTED.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n (occurrence **1** of the
two-consecutive-invalidations trigger, recorded in D6 so the next one fires) — (4) user-blocker: n (the
re-point is rext-only work under granted autonomy; no platform edit is implied) — (5) cap-reached: n —
(6) protocol-stop: n — Outcome: continue
**Decisions:** D1–D7 (iter-12/decisions.md)
**Side-deliverables:** the 14 `file:line` citation repairs in `demo-up-defaults.md` forced by the RAM fix's
23-line shift (committed separately as `8be95b3`); `DOC-M257x-claude-md-knob-count` closed REFUTED.
**Routes carried forward:**
- `FIX-M257x-iter13-router-drop-repoint` → **iter-13**. Re-point rext off the deleted router. Known sites:
  `stack-injection/gen_injected_override.py` (the `REUSE_DEV` image map, the `depends_on: graphql` block, and
  **both** `WUNDERGRAPH_SSR_ENDPOINT` emissions — hostname **and** path), `stack-injection/gen_tailscale_serve.py`
  (port 5050), `demo-stack/up-injected.sh` (the SSR origin chain + its comments),
  `demo-stack/clones.pin.json`, `stack-verify/repos/run.sh`,
  `stack-core/union_apply_guard.py` + `demo-stack/tests/test_ssr_origin_chain.py` and their tests.
  **Assume this list is incomplete until measured** — that is this milestone's standing rule.
- `FENCE-M257x-iter13-compose-service-exists` → **iter-13**. A fence asserting every compose service rext
  emits a `depends_on` on is defined by the platform's own compose at the ref in use. Same class as clause
  4's schema fence, one axis over. Must be **watched going RED against `2adcf71`** before it is believed.
- `FIX-M257x-iter13-freshness-vs-origin` → **iter-13**. The clone-freshness check compares the checkout to
  `clones.pin.json` instead of to **origin**, so a stale *clone* is reported as a stale *pin*. The gate says
  origin HEAD; nothing currently measures distance to it.
**Lessons:**
1. **Check the ref before spending the cycle.** The gate named origin HEAD in its first clause; three
   18-minute cycles were about to be run against a clone 3 commits behind it. The check cost 20 seconds.
2. **A freshness check that compares to a pin cannot detect a stale clone** — it can only detect a stale pin,
   and it reports the two identically. Promoted to the protocol doc as §3's operational corollary.
3. **The generator's RC=0 is the finding, not the compose RC=1.** compose caught this one because the
   dependency was structural. The two `WUNDERGRAPH_SSR_ENDPOINT` values are *strings* — nothing would have
   caught those, and they would have presented as the latency-budget "blackholing address" signature
   (≈ 3 × 10.5 s + 6 s) rather than as an error.
4. This is the **fourth** occurrence of the milestone's founding class (skiller → app, skillpath → app,
   jobsimulation → app, now router → app) and the **first one caught by us before it caused a mystery** —
   because the milestone existed and its gate named a ref.
