---
title: "KB Fidelity Audit — M258 «proven-live build»"
date: 2026-08-12
scope: milestone:M258
invoked-by: build-mstone-iters (Phase 0b, pre-bootstrap-tok)
---

## Verdict

**YELLOW** — proceed. No blind areas. Every load-bearing *substantive* claim this milestone would act on
was probed and holds. The findings are (1) systematic **line-anchor drift in the milestone's own plan**,
repaired inline, and (2) two **measurements that never propagated** out of a predecessor's milestone
record, now recorded at the destination.

Measured on `macmini`, `rosetta` @ `b8ae62c6` (branch `m258/proven-live-build`),
`rosetta-extensions` authoring copy @ `679a5f7` (= `origin/main` = tag `fast-build-m257-close`).

## Pre-flight rung zero (`corpus/ops/verification.md` §PRE-FLIGHT RUNG ZERO)

Run before anything else, because every downstream reading depends on it.

| rung | assert | result |
|---|---|---|
| 1 | work committed + tagged in the authoring copy | ✅ clean tree, `fast-build-m257-close` |
| 2 | the tag is **on origin** (`git ls-remote --tags origin`) | ✅ `15c1352…` → `refs/tags/fast-build-m257-close` |
| 3 | `.agentspace/rext.tag` names that tag | ⚠️ **NO — it names `fast-build-m257-iter-09`** (`8956e69`) |
| 4 | the consumption clone is checked out at it | not yet probed (follows from 3) |

**Finding R0 — the pin is one tag behind the published tooling.** `fast-build-m257-close` carries the
M257 close's three **fail-open repairs to `buildbench` itself** — the gate instrument M258's own p50
number is read from. A bring-up driven at the current pin consumes the **iter-09** instrument, i.e. the
one with the fail-opens still in it. This is the stale-pin class this release has hit repeatedly
(M236 lost an entire iteration to it). **Disposition:** stale-reference repair, owned by the first tik
that brings a stack up — re-pin before the first campaign, not after.

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| bring-up verification / autoverify | `corpus/ops/verification.md` | `stack-verify/live/{autoverify,verify}.sh`, `stack-verify/lib/*.sh` | PAIRED |
| the Playthrough batch + lifecycle | `corpus/ops/demo/playthroughs.md` | `playthroughs/e2e/run-playthroughs.sh`, `playthroughs/manifest/`, `playthroughs/cmd/{ptreport,ptvalidate}` | PAIRED |
| the world contract (reset vs presenter) | `corpus/ops/idempotency.md` + `playthroughs.md` §lifecycle | `stack-seeding/cmd/stackseed/main.go`, `demo-stack/up-injected.sh` | PAIRED |
| build budget / the p50 gate | `corpus/ops/demo/build-budget.md` | `demo-stack/buildbench.py`, `hostprofiles/*.json` | PAIRED |
| host topology / `--public-host` | `corpus/ops/demo/tailscale-serve.md` | `run-playthroughs.sh:92-105` | PAIRED |
| demo lifecycle / registry | `corpus/ops/rosetta_demo.md` | `demo-stack/rosetta-demo`, `up-injected.sh` | PAIRED |
| access latency | `corpus/ops/demo/latency-budget.md` | `stack-verify/e2e/run-latency.sh` | PAIRED |
| the content-stories sweep | `coverage-protocol.md`, `content-stories-spec.md` | `run-content-stories.sh`, `lib/content-pairs.ts`, `content-denominator.json` | PAIRED |

**No BLIND-AREA rows.** The milestone's `Delivers →` targets (`verification.md`, `playthroughs.md`) both
exist and both already cover their subject.

## Fidelity Findings

### F1 — `FIX-M257-content-stories-pair-count` — **ALIGNED, and genuinely open**

- **Source:** `overview.md` § *Inherited from the M256 close*
- **Expected:** the runner re-implements `buildPairs()` inline, **omits `manager_presence_only`**,
  computes **47** against the pinned **45** and `sys.exit(2)`s — the sweep refuses to start.
- **Actual:** confirmed on all four particulars. `lib/content-pairs.ts:115` **has** the
  `manager_presence_only` branch (landed `a5b1288`, v2.7 close); the shell re-implementation at
  `run-content-stories.sh:145-152` **does not** — its manager arm guards only `has_manager_view` +
  path + seat. `content-denominator.json` pins `expected_pairs: 45` with the arithmetic derived in
  full. So a served manifest carrying the two manager-presence-only voice cells counts **47 ≠ 45** and
  the runner exits 2 **before the sweep begins**.
- **Verdict:** ALIGNED. **Fix owner:** update code (the shell must gain the branch the TS has).
- **Note worth carrying:** the shell's own comment at `:131-133` explains this exact lockstep failure —
  for the *player* branch, at 49→47 — and then repeated it for the manager branch at 47→45. *The
  fourth inline copy of a counting rule drifted again, in the file that documents the drift.*

### F2 — `ptvalidate` is invoked nowhere outside its own tests — **ALIGNED**

- **Actual:** swept every `.sh`/`.go`/`.py`/`.ts` in the tree. Every hit outside `cmd/ptvalidate/` and
  `*_test.go` is **prose in a comment**, including `run-playthroughs.sh:330`, which names the
  both-way pre-flight as *"the honest fix"* without wiring it.
- **Verdict:** ALIGNED. Cheap to discharge: `ptvalidate --manifest-dir ./manifest --e2e-dir ./e2e` runs
  static-only in seconds (`datadna gate: SKIPPED (no --stack)`).

### F3 — the Playthrough denominator — **ALIGNED**, measured not quoted

- **Expected** (`playthroughs.md`, `CLAUDE.md`): 30 live Playthroughs + 1 verdicted TODO, 31 manifest
  use cases, 10 products.
- **Actual:** `ptvalidate` answers `manifest VALID: 10 product(s), 31 use case(s), 30 live
  Playthrough(s), 1 TODO`, `discovered 30 Playthrough test id(s)`.
- **Verdict:** ALIGNED. (A raw `grep '@pt:'` returns **35** — the extra 5 are the mutation-class fence's
  deliberately-split literals and prose citations. Grep is the wrong instrument here; the validator is
  the right one.)

### F4 — `verification.md`'s check-(h) claim — **ALIGNED**

- **Expected:** the container-liveness gap *"landed at M257 as `autoverify.sh` check **(h)**, which
  derives its expected set from `services.sh` and adds the injected trio."*
- **Actual:** `autoverify.sh:564` — `── 4. (h) CONTAINER LIVENESS — the stack's declared container set
  is all RUNNING (v2.8 M257)`. Present.
- **Verdict:** ALIGNED.

### F5 — the `DEMO_STORIES_PRESET` seam (world-contract resolution (a)) — **ALIGNED in substance**

- **Actual:** `up-injected.sh:261` `STORIES_PRESET="${DEMO_STORIES_PRESET:-…/stories.seed.yaml}"`, and
  all four downstream exports read `$STORIES_PRESET` (roster `:1981`, cockpit `:2615`, seed-manifest
  `:2622`, content `:2633`). The seam is real and resolution (a) is executable as described.
- **Verdict:** ALIGNED in substance; **anchors were wrong** — see F7.
- Note: `corpus/ops/demo/demo-up-defaults.md:47` cites `up-injected.sh:261` and is **correct**. The
  corpus was fresher than the milestone plan.

### F6 — M257's achieved numbers reached the corpus — **ALIGNED**

- `build-budget.md:396` carries `| total cycle p50 on macmini | 449.51 s | 286.99 s |`;
  `frontend-tier.md:805` carries the same transition. The doc is re-pointed at `macmini` throughout.
- **Verdict:** ALIGNED.

### F7 — **STALE: 12 line anchors in M258's own `overview.md`** (applied)

Every one is **in range** and lands on unrelated content, so no automated out-of-range lint could see
them. In **all twelve** the *substance* holds — this is drift, not a wrong contract. That is precisely
what makes it costly: an agent following `up-injected.sh:218` for the `DEMO_STORIES_PRESET` seam finds
a comment about `INJECT_SVCS` and may conclude the seam does not exist, when resolution (a) rests on it.

| claim | cited | actual |
|---|---|---|
| `TRUNCATE TABLE <t> CASCADE` | `stackseed/main.go:733` | `:839` |
| the `resetTables` list | `:44-125` | `:44-131` |
| roster export via `os.Create` (truncating) | `stackseed/main.go:225` | `:256` |
| `--cockpit-export` is "(no DB)" | `stackseed/main.go:150` | `:172` |
| re-seeds only `pt-world.seed.yaml` | `run-playthroughs.sh:96-98` | `:146` |
| the public-host / docker-proxy trap | `run-playthroughs.sh:56-72` | `:92-105` (`--reset-only` split at `:58-62`) |
| the `DEMO_STORIES_PRESET` seam | `up-injected.sh:218`, `:225` | `:254`, `:261` |
| Clerkenstein roster export | `up-injected.sh:1718` | `:1981` |
| cockpit manifest projection | `up-injected.sh:2256` | `:2615` |
| seed manifest export | `up-injected.sh:2263` | `:2622` |
| content manifest projection | `up-injected.sh:2274` | `:2633` |
| the bring-up **tail hook** | `up-injected.sh:2411` | `:2810` (the `autoverify.sh` invocation) |
| M204's pt-world-native precedent | `playthroughs.md:429` | `:777` |

**Fix owner:** update doc. **Applied** — all repaired in place. A stale reference is a repair, not a
re-plan, so no clause, target or number was touched.

## Completeness Gaps

### C1 — *critical* — the batch half of the composed budget has **no published wall-clock**

`overview.md` § *Budget honesty* required M256 to *"measure and report the reset-to-seed leg."* It did
not. No corpus doc publishes a suite wall-clock or a reset-leg figure; M256's `Gate Outcome Ledger`
grades clause 1 on a **per-test median**. The only wall-clock anywhere is **56.6 s** (M256
`progress.md` iter-02, n=3, *"reported, not gated"*) — and it was taken over **18 specs**, where the
suite closed at **209 passed across 30 live Playthroughs**. Quoting it as the batch half would price a
suite an order of magnitude smaller than the one M258 must run.

⇒ M258's `286.99 + ?` composition has **one measured half**. **Backfilled** into `spec-notes.md` and
flagged at `overview.md` § *Budget honesty*; measuring it is a first-tik obligation, not a doc fix.

### C2 — *critical* — M256 escalated that this suite's timing is **not decidable at n=3 on this host**

M256 `progress.md` iter-12: *"clause 1 is **NOT DECIDABLE at n=3 on this host**"* — six full-suite runs,
same box, unchanged specs, control subset spanning **0.5281× → 1.0762× (2.04×) with no trend**. The gate
was re-cut (`D-v28-12` → `D-v28-13`) because the threshold sat inside its own noise floor. That
escalation reached **neither** M258's plan **nor** the corpus.

M258's gate is a **p50 over n=3** whose second half is that same suite. **Backfilled** to `spec-notes.md`
+ `overview.md`. Recorded as **evidence, not a re-cut** — the obligation it creates is that the first
campaign publishes the batch half's **spread** beside its p50, so "inside the gate" can be told from
"sampled favourably".

### C3 — *incidental* — `build-budget.md:555` cites `apps/{web,hiring}/package.json:98/:92`

Platform-repo paths, not resolvable from the rext tree; my lint matched them against same-named rext
files and flagged them out-of-range. **Not a finding** — recorded so a later audit does not re-raise it.

## Applied Fixes

1. **`overview.md`** — 13 stale line anchors repaired (F7). Substance untouched.
2. **`overview.md` § Budget honesty** — a bounded ⚠️ note recording C1 + C2 with provenance, pointing at
   `spec-notes.md`. No clause, target or number altered.
3. **`spec-notes.md`** — written from stub: the Phase-0b verdict record, the topic→doc→code triple table
   (fast start for later audits), and the full derivation of C1 + C2 with per-figure provenance.

## Open Items (require user decision)

**None.** Every finding was either applied (F7, C1, C2 as annotations) or is in-scope milestone work with
a named owner (F1, F2, R0). C2 *could* become a user renegotiation — but only once measured; escalating a
noise floor before taking a sample would be exactly the un-measured escalation this release has twice
had overturned.

## Gate Result

**YELLOW — proceed with tracking.** The bootstrap tok authors `TOK-01` against the corrected plan, and
carries R0, C1 and C2 as known-context:

- **R0** — re-pin `.agentspace/rext.tag` to `fast-build-m257-close` before the first campaign, or the
  gate is read from an instrument with three known fail-opens in it.
- **C1** — the batch half is unmeasured; measure it before predicting the composition.
- **C2** — publish the spread beside the p50, or the n=3 reading is not falsifiable.
