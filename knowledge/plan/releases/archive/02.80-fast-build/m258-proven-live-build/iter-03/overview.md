---
iteration_type: tik
status: in-flight
milestone: M258
iter: 03
active_strategy: TOK-01
created: 2026-08-12
---

# M258 iter-03 — unblock the batch half (the diagnosis was wrong; the blocker is real)

**Type:** tik · **Active strategy:** `TOK-01` step 1 — *measure the composition before engineering it.*

## Step 0 — re-survey (mandatory) — THE ROUTED DIAGNOSIS IS REFUTED

`progress.md` routed **`FIX-M258-iter02-inject-appends-and-swallows`** as iter-03's first job, on the
reading that *"all three UI images were wired to a **real Clerk app**"* because a 24th appended block
*"carries the foreign key"*. **Re-survey against code refutes the mechanism and the severity, and finds
the true blocker adjacent to it.** This is `TOK-01` known-context #6 (*re-verify every inherited item
against code before working it*) paying out on the milestone's own routing, one iter after it paid out
on M257x's.

**The key is not foreign. It is Clerkenstein-minted, for this machine.**
`mint_pk(host) = pk_test_ + base64(host + "$")` (`inject.py:29-37`). Decoded:

| block | decoded FAPI host | verdict |
|---|---|---|
| 1st (`.env.demo-1:7`, `pk_test_MTI3…`) | `127.0.0.1:15400` | loopback — an **earlier** bring-up's key |
| 24th (`:122`, `pk_test_bWFy…`) | `marcos-mac-mini.taildc510.ts.net:15400` | **this run's**, correctly minted |

`bWFy` is base64 `mar`, not a real-Clerk key. **No demo was wired to production auth.** iter-02's D7
quarantine (*"must not be browsed — its UI tier would talk to a real Clerk app"*) rests on a refuted
premise.

**Why the host differs — auto-discovery is DEFAULT-ON, and the campaign never opted out.**
`up-injected.sh:139` `FAPI_HOST="${STACK_PUBLIC_HOST:-127.0.0.1}"`; `:110-114` auto-discovers a MagicDNS
host unless `--public-host` **or** `--no-public-host` is given. `cycle.log:26` records it verbatim:
*"public-host AUTO-DISCOVERED — marcos-mac-mini.taildc510.ts.net … The demo will be reachable over the
tailnet on HTTPS."* Ports are bound `0.0.0.0` (`docker port`, confirmed). **So iter-02's D2/D7 claim that
`demo-1` is `--no-public-host` / localhost-bound is FALSE** — a record-level correction with a safety
face (§ Record corrections below).

**The ISOLATION RED is a FALSE POSITIVE, and its cause is a reader/consumer ordering disagreement.**
`_stack_minted_pk` (`buildbench.py:1753`) returns `_first_key(...)` — the **FIRST** `pk_` in
`.env.demo-N`. Every consumer of that file (compose, `next build`) takes **LAST-wins**. With 24
accumulated blocks the assert compares this run's correctly-baked images against **an earlier bring-up's
stale key** and reds all three. Its own docstring already condemns exactly this failure *shape* for the
fingerprint case — *"Wrong-and-loud is not fail-closed — it is a fence that cannot be believed"* — and
missed the ordering case one line away.

**And the actual blocker on the milestone's primary unknown is a fourth thing, not yet routed at all.**
`buildbench.py:1454-1456` builds `argv_up` and appends `--public-host H` only when given; it **never
passes `--no-public-host`**. So a campaign **cannot express the single-box mode `TOK-01` declares the
gate is taken against**, and silently runs public-host. That is fatal to the batch half:
`run-playthroughs.sh:88-108` documents that a `--public-host` demo **cannot be browsed from its own
host** (docker-proxy binds `0.0.0.0`, bypassing `tailscale serve` → `ERR_SSL_PROTOCOL_ERROR`, *"every
page renders a permanent loading spinner and every assert fails for a reason that has nothing to do with
the product"*). **iter-02's refusal to measure was right, for the wrong reason.**

**Substitution recorded:** `TOK-01`/`progress.md` named *append + swallow*; re-survey keeps the append
(real, P1) , **replaces the swallow-as-root-cause with the first-wins reader (P2)**, and **adds the
mode passthrough (P3)** as the item that actually gates the deliverable. Same strategy, corrected target.

## Cluster / target identified

One causal chain, four planned deliverables, then the measurement `TOK-01` step 1 still owes.

- **P1 — `inject.py`: strip-then-append.** The enabling condition (24 blocks). The fix shape is not
  invented: `up-injected.sh:2052` **already** does exactly this for its own `DESK_CLERK_` block in the
  same file — *"Strip any prior DESK_CLERK_ block so a re-up rewrites (idempotent)"* — which is why that
  block appears once and inject's appears 24 times.
- **P2 — `_stack_minted_pk`: last-wins for `.env.demo-*`.** The proximate cause of the false RED. The
  assert must be right independently of the writer, because `--purge` does not clear the stack dir
  (F-9), so existing dirs stay multi-block after P1.
- **P3 — `buildbench`: `--no-public-host` passthrough.** The precondition for measuring the batch half.
- **P4 — `up-injected.sh:2036`: surface `inject.py`'s stderr.** ⚠️ **The `|| true` is DELIBERATE and must
  NOT be "fixed"** — `:2033-2035` explains it guards a `set -e` death so the `[ -n "$PK_DEMO" ]`
  fail-loud guard on `:2037` can own the failure path. Only `2>/dev/null` is the defect: it discards
  inject's success report, whose text is *"minted {pk} (host={fapi_host}, round-trip OK)"* — **the one
  line that would have named the MagicDNS host at mint time.**

## Hypothesis

P1+P2 clear the false RED so a rep is `gateable`; P3 lets the cycle run in the mode the batch can
actually be driven from. With those, one cold cycle yields **the first measured batch half** — the
milestone's primary unknown and `TOK-01` step 1's outstanding deliverable.

## Expected lift

Primary metric: **the batch half exists as a number** (currently unmeasured). Secondary: bring-up half
at n=2 in the *correct* mode, comparable to M257's 286.99 s p50 — the 395.31 s n=1 sample was taken in
public-host mode, which pays `tailscale serve` + cert-mint legs M257's p50 never paid. **That is a
candidate explanation for part of the +108.32 s delta and will be stated as a hypothesis, not a finding,
unless the numbers support it.**

## Phase plan

- **Phase A** — land P1–P4, each with a regression test proven **RED with its precondition absent**.
- **Phase B** — test gate: the touched rext suites.
- **Phase C** — publish + re-pin (tagging is not publishing), then one cold `--no-public-host` cycle and
  **the full Playthrough batch**. Report `load1` + environment with every figure.
- **Phase D** — close: record the batch half, its caveats, and the record corrections.

## Escalation conditions

- Batch red set non-empty at batch end → **one consolidated escalation** (`D-v28-3`), never a halt.
- Composed p50 > 600 s → the declared `re_scope_trigger`, surfaced **with measurements**.
- Any *new* line of investigation beyond P1–P4 → route forward (tripwire), do not absorb.

## Acceptable close-no-lift outcomes

A measured batch half that **misses** the composed budget is a **result**, not a failure — provided its
spread is published beside it (`TOK-01`'s standing rule, `C2`). Likewise a HEADROOM refusal.

## Record corrections carried by this iter

1. **`demo-1` is tailnet-reachable, not localhost-bound** (iter-02 D2/D7). It is Clerkenstein-wired, so
   this is the *documented* demo posture (`safety.md` Part 3: unauthenticated, authz-weakened,
   tailnet-scoped) — **not a safety violation, and not a real-Clerk exposure**. Corrected in the open.
2. **The "must not be browsed" instruction is withdrawn as to its stated reason.** No production auth is
   involved.
