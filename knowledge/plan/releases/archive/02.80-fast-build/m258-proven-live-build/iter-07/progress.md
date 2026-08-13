# M258 iter-07 — progress

**Type:** tik · **Active strategy:** `TOK-01` step **4** — *the composed 3× cold campaign against the
gate, with the spread published beside the p50.* Steps 1–3 are discharged (bring-up 247.79 s, batch
129/160 s, gate wired + restore leg landed).

## Phase A — the re-pin, and four things verified rather than inherited

### A1. Rung zero, in both directions

`fast-build-m258-iter-06` is **on origin** (`5f3a3815`), and the authoring copy's HEAD is the **same
sha** with a clean tree. The consumption clone was at `fast-build-m258-iter-03`, and the gap was
**exactly the feature under test**: all four batch-gate files absent, no hook in `up-injected.sh`
(`D21`). Re-pinned; all four present, hook at `:2844`. Had this been skipped, the campaign would have
measured a bring-up **with no batch gate in it** — and the phase table would still have summed.

### A2. The user's stacks — and a hazard the re-pin walked into

`resolve_stack_dir` (docker mount, never `$EXT_ROOT`) shows the two clones own **different** stacks:
**demo-1 → authoring** (mine), **demo-2 → consumption** (⚠️ the **user's**). So the re-pin changed the
clone that owns demo-2. **Verified harmless rather than assumed** (`D23`): `demo-stack/.gitignore:8`
ignores `stacks/`, confirmed by `git check-ignore -v`, so a checkout provably cannot reach
`stacks/demo-2/`. demo-2 measured **11 containers, up 35 hours**; dev stack **5** — before and after.

It also means the campaign's teardown would be **cross-clone** (`rosetta-demo down` runs from
`rext_root()` = consumption, while demo-1's `docker-compose.injected.yml` is in the authoring clone).
Pre-empted: tear demo-1 down from its **owning** clone immediately before launch, so all three reps are
consistent rather than rep 1 being a special case.

### A3. `D27` — the routed fix was a hypothesis, and it was settled WITHOUT a campaign

iter-06 closed `green: false` on one `autoverify` probe and routed *"run from the consumption clone"*
forward. Instead of spending a 3-rep campaign finding out, it was reduced to arithmetic plus one query
(`readiness.sh:65,71`):

| input | expression | authoring | consumption |
|---|---|---|---|
| `_rl` | `<rext>/stack-core/lib/repos_yml.sh` | ✅ | ✅ |
| `_ry` | `<rext>/../platform/repos.yml` | ❌ absent | ✅ present |

Then both halves live: the derivation yields **`extensions`, `sentinel`, `public`** (+`directus` when
that container runs), and demo-1's Postgres reports **all four**. So the probe will not merely assert —
it will **pass**. **`green` flips `false → true` for topology reasons alone**, and the last blocker to a
gate-usable rep is discharged *before* the campaign rather than by it.

### A4. Clause 4 graded early, because a boolean survives contention

**0 platform-repo edits: PASS** (`D26`). platform / app / sentinel / next-web-app / studio-desk all
clean (`app`'s `?? studio/` is the CI-pulled studio-room, untracked, not an edit). `ant-academy` carries
4 modified files — **pre-existing sanctioned demo-patches**, self-identifying in source and gated behind
`ACADEMY_DEMO_FS_PUBLISHED` / `NEXT_PUBLIC_COCKPIT_URL`; the canonical `anthropos-work` repos are
untouched, which is the `demopatch-spec.md` contract. Surfaced one route
(`ROUTE-M258-iter07-demopatch-G5-does-not-revert-the-native-clone`) — ant-academy runs **natively**, so
it has no ephemeral clone for G5's revert leg to discard, and G2 drift-refuse can later read DRIFTED
against a clone a previous run patched. **Not repaired**: reverting tracked files is a forbidden op, and
this clone is what the **user's demo-2** serves from.

## Phase B — the launch window (the host, stated with the number)

The gate's clause 3 is the only **timing** clause; 1, 2, 4 and 5 are **booleans**. That distinction is
what makes this host workable at all: *a boolean survives contention, a timing does not.*

`headroom_assert` grades clause 1 as `load1 ≤ host_logical_cores − 2` = **10** on this 12-core host, and
it is enforced **twice** (`D24`): `pre_rep_assert` **instantaneously before each rep** — on failure
`run_campaign` **returns 1**, ending the campaign before any teardown — and `headroom_assert` **post-rep
on `peak_load1`**, feeding `rep_is_ok`. So a campaign launched into a high load yields either an
immediate refusal or reps the harness discards.

Measured at open: `assert-headroom --profile macmini` → **FAIL, load1 16.58 vs 10** (disk 66.9 GiB free,
ample). The contention is **third-party and named** (`D25`): 8 × `mdworker_shared` + `mds` (Spotlight,
peaking **227 % CPU**) plus **two** `a8-cart-runner` processes at ~100 % each belonging to a *different*
user project (`hyperspace/anima8/…/m270a2-iter132/`). `R=16` runnable against 12 cores, `U=0` —
CPU-bound, not I/O-blocked. Neither is mine to stop; the standing rule is that this host **cannot be
freed**.

The campaign was therefore **held rather than launched-and-reported** — a refusal is a result, and a p50
over reps the instrument discards is not a number.

### B1. The "before", captured so the campaign has something to be compared against

Taken on the live `demo-1` **before** any teardown, so clause 5 has a falsifiable baseline rather than a
recollection:

| artifact | reading |
|---|---|
| `autoverify.json` | `{"green": false, "warnings": 1}` ← the `postgres-schemas` refusal `D27` explains |
| `batch-gate.json` | `verdict green` · `red_count 0` · `batch_seconds 160` · `restore_seconds 7` · `total 31` · `passing=30 unimplemented=1` |
| `cockpit-manifest.json` | 4 stories · 13 deep links |
| **post-condition** | `check-cockpit-roster.py` → **✓ all 12 cockpit seats resolve in the 35-identity roster** (rc=0) |

So `demo-1` is **verifiably presenter-usable at iter-07 open**, and the single field the campaign must
move is `autoverify.green` → `true`.

### B2. The window never opened — measured, not asserted

~30 minutes of foreground polling (`D28`). `load1` vs a limit of **10**:

```
07:28 20.34 · 07:31 23.62 · 07:33 11.93 · 07:34 32.90 · 07:37 40.61
07:40 13.37 · 07:42 43.49 · 07:44 45.92 · 07:45 62.88 · 07:46 55.95
```

**Minimum 11.93; never below the limit; trending UP.** The campaign was not launched, on two independent
grounds — the second of which is the load-bearing one:

1. Every rep would be discarded by `rep_is_ok`, and `demo-1`'s presenter world would be spent to produce
   nothing. (The launch script now asserts headroom **before** the teardown so this cannot happen.)
2. **A batch run at `load1` 40–60 would have manufactured FALSE REDS.** The suite runs with
   **`retries: 0`** by contract (`D-v28-3` clause 2), so browser timeouts under load are recorded as
   failures, and clause 4 then **escalates the red set to the user for renegotiation**. That hands over a
   verdict describing the *host* as though it described the *product* — the precise failure
   `batch-gate.sh`'s own header warns of: *"a false RED … TRAINS ITS OPERATOR TO DISBELIEVE THE GATE."*

## Phase C — the gate, graded honestly: 4 of 5 clauses, and the 5th is one quiet window away

| # | clause | kind | verdict |
|---|---|---|---|
| 1 | one cold command brings the stack up **and** drives the full batch | boolean | ✅ **proven at iter-06** (hook at `up-injected.sh:2844`, 215 passed, 31 use cases) |
| 2 | zero standing red | boolean | ✅ **red set EMPTY** (`batch-gate.json`: `red_count 0`, `passing=30 unimplemented=1`) |
| 3 | composed **p50 ≤ 480 s** over 3 consecutive cold cycles | **timing** | ⬜ **NOT TAKEN** — instrument refuses (`D28`); projection 414.15 s stands as a *projection* |
| 4 | 0 platform-repo edits | boolean | ✅ **PASS** (`D26`, measured across all six peer clones) |
| 5 | stack left presenter-usable | boolean | ✅ **12/12 seats resolve in the 35-identity roster**, verified live this iter |

The one clause that is a **timing** is the one clause outstanding — which is exactly the shape the
standing host rule predicts: *a boolean survives contention, a timing does not.*

## Close — 2026-08-12

**Outcome:** **The last blocker between this milestone and a gate-usable rep is discharged — and it was
discharged by measurement, not by a campaign.** iter-06 routed forward the hypothesis *"run from the
consumption clone"*; rather than spend a 3-rep campaign testing it, `D27` reduced it to path arithmetic
plus one query and proved that `autoverify`'s `postgres-schemas` probe **will assert and will pass**
there — so `green` flips `false → true` for topology reasons alone. The consumption clone was re-pinned
to the pushed tag (the gap was **exactly the feature under test**: four missing files and no hook), the
launch is scripted and shellcheck-clean with a headroom-before-teardown guard, and **4 of the gate's 5
clauses are proven**. The 5th is a timing, and the host never offered a window: `load1` minimum **11.93**
against a limit of **10** across ~30 minutes, trending up to **62.88**.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n *(the trigger reads a
composed **p50 exceeding 600 s after 3 tiks**; no p50 was taken at all, and firing on an unmeasured
quantity is the category error iter-04 and iter-06 both refused)* — (4) user-blocker: **y** *(the
milestone's headline number needs ~30 min of a host at `load1 < 10`; this box is saturated by a
**different user project's** parallel campaign, which is not mine to stop — see below)* —
(5) cap-reached: n *(1 tik)* — (6) protocol-stop: n — (7) budget-exhausted: n *(context and turns are
not the constraint; the host is)* — Outcome: **exit-4**

> **Why `user-blocker` and not `budget-exhausted`.** A budget stop resumes by re-invoking. This does
> **not**: the next invocation meets the same saturated box and produces the same refusal. The
> constraint is external, measured, and **the user is the only party who can arbitrate it** — by
> scheduling the campaign when `anima8`'s campaign is idle, or by accepting the measurement on a
> different host. Surfaced **with measurements, not as a question**, exactly as the milestone's
> re-scope discipline requires.

**Decisions:** D21–D28 (this iter's `decisions.md`)

**Side-deliverables:**

- **`.agentspace/scratch/work-m258/launch-iter07-campaign.sh`** — the campaign, scripted and validated
  (`bash -n` + shellcheck clean). Asserts the user's stacks resident (`demo-2`=11, dev=5) and **refuses
  if they are not**; asserts headroom **before** the teardown; tears `demo-1` down from its **owning**
  clone; runs 3 reps from the consumption clone. iter-08 runs one command.
- **`ROUTE-M258-iter07-demopatch-G5-does-not-revert-the-native-clone`** — new, with evidence (`D26`).

**Routes carried forward:**

- **`TOK-01` step 4 — the composed 3× cold campaign** (iter-08). **Every precondition is now
  discharged**: tag on origin, consumption clone re-pinned and carrying the gate, `postgres-schemas`
  proven satisfiable, script written. **It needs only ~30 minutes of `load1 < 10`.**
- **`LEVER-M257-L5-setdress`** — untouched reserve, still the largest bring-up phase (80.42 s at
  iter-06). Only needed if the composed p50 lands over 480 s.
- Unchanged and still open: `FIX-M258-iter03-guard-scans-its-own-scratch` (+ its same-family sibling
  `test_fence_provenance::test_the_escape_accepts_and_records`) ·
  `ROUTE-M258-iter02-isolation-names-two-causes-not-three` ·
  `ROUTE-M258-iter02-headroom-defaults-to-billion` · `ROUTE-M258-iter02-purge-did-not-clear-the-stack-dir`.
- ⚠️ `demo-2` (11) and the dev stack (5) verified resident before **and** after every operation.
  **`demo-1` was never torn down** — it is UP and presenter-usable (12/12 seats).

**Lessons:**

- **A routed fix is a hypothesis — and the cheapest way to test one is often not to run the experiment
  it names.** iter-06 routed "run from the consumption clone" as a *campaign* precondition. It was
  settled with two path expressions and one `psql` query, before any host was needed. The campaign was
  never the way to learn whether the campaign would work.
- **Refusing to generate a verdict is sometimes the deliverable.** Running the batch under load 40–60
  with `retries: 0` would not have produced a slow green — it would have produced a **red set describing
  the host**, escalated to the user as though it described the product.
- **Guard the irreversible step with the gate that governs it.** `buildbench` asserts headroom *after*
  the point where it has already torn the stack down; the launch script moves that assert **in front of**
  the teardown, so a marginal window can never cost a good demo.
- **When two clones of one tree exist, ask the running system which one it is using — for OWNERSHIP too,
  not just for paths.** `D19` established the rule for a file path; `D23` found the same hazard one level
  up: the two clones own *different stacks*, and the user's demo-2 belongs to the one this iter re-pinned.
