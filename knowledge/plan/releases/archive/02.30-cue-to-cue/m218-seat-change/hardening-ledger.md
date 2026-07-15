# Hardening Ledger — M218 seat change

## Pass 1 — 2026-07-14 — final

**Iters hardened this pass:** all milestone-touched code (final mode, cumulative scope — iter-01 … iter-05)
**Tiks covered since prior pass:** all iters in milestone (no prior harden pass)

**Scope.** The true M218 rext footprint is **17 files** (`3fab10c^..f296e5e`) — `clerk-backend/` +
`cmd/fake-bapi/` (iter-04), `rosetta-demo` + `up-injected.sh` + `test_purge.py` (iter-05), the
`next-web-ssr-graphql-origin` demo-patch + `gen_injected_override.py` (iter-03), the `stack-verify/e2e/`
latency harness (iter-02), `cockpit.py`. Plus, in rosetta: `alignment_testing.md`, `verification.md`,
`cockpit-spec.md`, `latency-budget.md`. (`dev-stack/` appeared in an earlier, wrongly-widened diff — it is
**M217's**, not this milestone's.)

### The three owed Fate-1 items — ALL LANDED

| # | Item | Outcome | Proof |
|---|---|---|---|
| **F1-1** | `GetUser` per-hero-identity gene on the BAPI surface | **LANDED** | `gate.sh` exit **2 (RED)** @ `8ebc89e^` — GetUser **0/2**, critical **88.2%**, both heroes returned the stub `11111111-…`/`demo@anthropos.test`. `gate.sh` exit **0 (GREEN)** @ HEAD — GetUser **2/2**, critical **100.0%**. Measured, not asserted. |
| **F1-2** | Teardown must unlink `autoverify.json` (F-10) | **LANDED** | `clear_stack_verdict()` called **first** in `cmd_down` + unlink at bring-up start + `ts` field. **7 regression tests, all 7 RED against `f296e5e`**, all green after. |
| **F1-3** | The capability-coverage check is not binding | **LANDED** | It did not merely fail to bind — **it did not exist**. `alignctl dna` was `list\|diff\|validate`; the only "coverage" was an *eyeball* step in `align-dna/SKILL.md:63`. Now: `consumed_surface` schema + `Validate` rejection (so `alignctl run` **refuses to score**) + `alignctl dna coverage` + wired into `gate.sh`. Doc rewritten to state what it actually guarantees **and what it does not**. |

**How F1-1 avoided being theatre.** The runner never mounts a roster, so an in-process gene falls back to the
stub — it would have been **green against the broken *and* the fixed mirror**. It also could not have compiled
against the pre-fix store (`Store.SeedUserIdentity` did not exist), and **a build error is not a red gene**.
The one seam byte-identical across the fix is the demo's own: `FAKE_FAPI_ROSTER` → `cmd/fake-bapi` →
`GET /v1/users/{id}`. The gene drives that binary as a subprocess, which also fences the **wiring**.

**Coverage delta on touched files** (coverage used as a *finder*, not a goal — see below for what it found):

| file | before | after |
|---|---|---|
| `clerkenstein/clerk-backend/{server,resources,store}.go` (`getUser` path) | unit tests only (iter-04); **no alignment gene** | **4 critical genes** + 3 unit tests; the `getUser` HTTP path is now gated |
| `clerkenstein/cmd/fake-bapi/main.go` (roster→identity seeding) | **0 tests** | fenced end-to-end by the roster-driven genes (the binary itself is booted) |
| `alignment/internal/dna` + `cmd/alignctl` | no coverage check at all | **17 new tests** (9 validate + 8 CLI/exit-code) |
| `demo-stack/rosetta-demo` (`cmd_down`) | purge fenced (iter-05); **verdict lifecycle unfenced** | **7 tests** (behaviour + wiring + ordering) |
| `demo-stack/patches/next-web-ssr-graphql-origin/` + `gen_injected_override.py` | **0 tests — the milestone's headline fix** | **12 tests**, incl. the live-clone freshness gate |

**Tests added:** 37 (17 Go alignment · 12 Go clerkrun · 19 Python — 7 verdict-lifecycle, 12 SSR-chain · 1 F-6 fence).
_Files: `coverage_test.go` ×2, `bapi_http_test.go`, `test_verdict_lifecycle.py`, `test_ssr_origin_chain.py`, `test_frontend_build.py`._

**Final suite state:** Python **592 passed / 0 failed** (was **1 failing** — see below) · Go **0 failures**
across all 3 modules · Go alignment gate **rc=0**.

**Bugs surfaced + fixed inline:**
- **`demo-stack/tests` was RED — and M218 itself broke it.** `test_tag_guard_present_for_both_frontends`
  pinned `docker image rm -f "$img"` at **2**; **iter-03's own F-6 fix (`0cfe23e`) added a third** rebuild
  trigger (next-web's minted-PK check) and never updated the fence. Verified failing identically at baseline
  `f296e5e`, and my diff changes none of the three counts — so it has been red since iter-03 and nothing
  caught it. Count corrected to 3, **with the reason written down** so it isn't "fixed" back down. (`f849b5f`)
- **The F-6 guard itself had no test** — the guard that catches a **Clerkenstein-DEWIRED demo talking to
  production Clerk**. Now fenced (the validator must grep the *built bundle* for the minted pk, since an
  ENV-level check cannot see a build-inlined constant). **Mutation-proven:** swap the bundle grep for an
  ENV-only check → the assertion fails. (`f849b5f`)
- **`gate.sh` could not run at all on a clean box** — `GOPROXY=off` + a `go.mod` `toolchain` directive made
  `GOTOOLCHAIN=auto` try to *download* the toolchain and die on `checksum database disabled` (a red herring;
  the local Go already satisfied the directive). Pinned `GOTOOLCHAIN=${GOTOOLCHAIN:-local}`. (`dd65ad0`)
- **My own first ordering test matched `docker compose` inside a comment** and failed a correct file. Fixed the
  test (assert on comment-stripped code) — and added `checked > 0` guards so an assertion that inspects
  nothing **fails** instead of passing. (`fc32baa`)

**Cross-iter integration finding (the defining work of final mode).** The milestone's **headline fix** — the
SSR GraphQL origin, worth ~37.5 s of the 39.45 s → 2.4 s collapse — is a **two-part chain across two
subsystems** (the demo-patch reads `WUNDERGRAPH_SSR_ENDPOINT`; `gen_injected_override.py` supplies it) and
**had zero tests**. Neither half works alone. Per-iter regression tests, each scoped to one iter's diff,
*structurally cannot* hold such a chain together. Worse, `demopatch` **refuses a drifted patch silently** —
`demopatch-spec.md`'s own warning is that this shipped a 76 s members grid for four releases. Landed 12 tests
(`ef6eefc`), mutation-proven: half-2 deleted → 2 fail; public origin → 2 fail; anchor drift → gate fails.

**Flakes stabilized:** none found. The new BAPI genes boot real subprocesses on ephemeral ports — flake gate
run **3×** consecutively: clean 3/3 (Go alignment, Go clerkrun, both Python suites, and `gate.sh` rc=0).

**Knowledge backfill:**
- `corpus/architecture/alignment_testing.md` — the capability-coverage section **rewritten** (it described a
  check that did not exist), with an explicit *what it does NOT guarantee*; the **CI-inert correction**
  (there is **no** `.github/workflows` in rext — the "weekly CI workflow" never existed); and the
  golden-ratifies-the-mirror lesson (prefer genes a stub **cannot** satisfy).
- `corpus/ops/verification.md` — **THE STALE-VERDICT HAZARD** promoted to a first-class named hazard (below).
- `.claude/skills/align-dna/SKILL.md` — step 7 "*List to eyeball coverage*" → **run `alignctl dna coverage`**.
- `rosetta-extensions/clerkenstein/knowledge/alignment.md` + `alignment/README.md` — new gene counts, the
  97.2% score and why, and the `expressrun` dependency-gate honesty note.

**The named hazard (the cross-cutting pattern).** The same failure class hit **five times in one milestone** —
F-6, F-9, F-10, plus two probe-level instances (`[ -e ]` reading permission-denied as absence; `assertNotIn`
passing on a failed command's empty output). It had already survived M217's hardening. Fixing a sixth instance
in isolation would have missed the point, so it is now documented as a first-class hazard:
**a status artifact that outlives the thing it describes, and is then read as evidence.** Two invariants:
(1) a verdict must not outlive its subject — destroy it on teardown *and* at the start of every bring-up, and
destroy it **first**; (2) **absence must be the safe state** — a grader with no verdict must refuse to measure.
Corollary for the checks themselves: a probe that can pass **without executing its assertion** is a stale
verdict in test form.

**Stop condition:** `stabilized` — the three owed Fate-1 items are landed and independently proven
(red-before/green-after, not asserted); the cumulative-scope dimension scan found one further gap (the
unfenced headline chain), which is now closed; the flake gate is clean 3/3; and the remaining surfaced items
are all **routed forward with named handlers** (below), none fixable inline without violating the
graded-artifact boundary.

### Routed forward (NOT fixable inline — recorded, not silently dropped)

| handler | item |
|---|---|
| `FIX-M219-bapi-org-eid` (**F-11**, new) | The BAPI fabricates `organization.public_metadata.eid` as `"org_eid_"+orgID` for any non-demo org instead of the **real** org UUID the roster carries — the **ORG-level twin** of the user-identity stub. **Not fixed here:** `resources.go` is the demo's **runtime** path and the 5-cycle cold gate was graded on the current binary (iter-05 **D13**); changing it post-gate would ship something other than what was measured. Landed instead as a **deliberately RED, permanently-visible** standard gene (`MembershipOrgIdentity/real-org-eid`) → Go surface **97.2% / 100% critical**, gate still MET. **A silently-omitted field is how the headline bug survived four releases.** |
| `DOC-M218-audit-corrections` | **F1-3 discharged** (the coverage claim) **+ the CI-inert correction discharged**. **Remaining:** the `clerkenstein.md:3-4` header. |
| `TEST-M219-expressrun-dep-gate` | `expressrun` is **UNMEASURABLE** on a box without `@clerk/express` `node_modules` (rc=2, no score) — reproduced identically at baseline `f296e5e`, so **pre-existing**. iter-04's "all 5 surfaces 100%" is therefore **not reproducible here**; 4 of 5 were re-measured this pass. |
| `TEST-M219-freshness-gate-skips` | The demo-patch live-clone freshness gate **skips** when `stack-demo/next-web-app` is absent (the established `test_demopatch.py` tradeoff). It ran here — but a box without the clone gets **no** anchor-drift protection. Itself an instance of *absence read as success*. |
| ~~_(pre-existing, out of M218 scope — **M217** footprint)_~~ **← ❌ WRONG. CORRECTED AT THE CLOSE, 2026-07-14.** | ~~`stack-injection` `test_next_web_block_shape` (1 failure)~~ **is an M218 REGRESSION, not pre-existing.** See the correction immediately below. The `dev-stack/tests/test_dev_stack.py` half of this row **stands**: those failures are **environment-driven** (this box's `.agentspace/secrets` lacks critical keys, so the secret-coverage pre-flight aborts *before the N-guard is reached* — which incidentally shows `guard_n` should run **before** the pre-flight, a dev-stack ordering smell, M217 footprint). |

#### ❌ Correction (close-milestone Phase 4, 2026-07-14) — this pass misdiagnosed its own regression

`stack-injection::test_next_web_block_shape` was classified above as *"pre-existing, M217 footprint"* on the
evidence that it **"reproduced at baseline `f296e5e`"**. **That baseline was wrong.** `f296e5e` is *M218's own
iter-05 commit* — it is not the pre-M218 baseline. The true pre-M218 baseline is the rext tag
**`cue-to-cue-m217`**, and the close re-measured against it:

| ref | result |
|---|---|
| `cue-to-cue-m217` (**true** pre-M218 baseline) | **PASSES** — `1 passed, 131 deselected` |
| M218 HEAD (`f849b5f`) | **FAILS** — `First extra element 13: '      - WUNDERGRAPH_SSR_ENDPOINT=http://graphql:8080/graphql'` |

⇒ **M218 broke it, in iter-03.** iter-03's C-1 fix added `WUNDERGRAPH_SSR_ENDPOINT` to the next-web block in
`gen_injected_override.py` and left the exact-shape fence pinned to the old block. The suite had been red since
iter-03 and nothing caught it — **the identical bug, one file over, from the identical cause, as the
`demo-stack` `test_tag_guard_present_for_both_frontends` fence this very pass found and fixed.** The pass caught
the neighbour and misfiled the twin.

**Why the misdiagnosis is the point, not a footnote.** "Reproduced at baseline ⇒ pre-existing" is only sound if
the baseline predates the milestone. Choosing a *mid-milestone* commit as "baseline" turns a regression into a
clean bill of health — **a status artifact that outlives the thing it describes, and is then read as
evidence.** That is **D17, the stale-verdict hazard**, in its sixth instance of this milestone — surfacing
inside the very hardening pass convened to name it.

**Fixed at the close:** the fence now asserts the `WUNDERGRAPH_SSR_ENDPOINT` line, **with the reason written
into the test** so it cannot be "fixed" back down by deleting the line. `stack-injection`: **186 passed / 0
failed** (was 185 passed / **1 failed**).

### What the coverage sweep found that per-iter tests structurally could not

Two of this pass's four inline fixes were **tests that were already red or absent inside the milestone's own
footprint**, and neither could have been caught by the iter loop's per-symptom discipline:

- `demo-stack/tests` had been **failing since iter-03** — iter-03 added a rebuild trigger and left the fence
  pinned at the old count. A per-iter regression test fences *the symptom it just fixed*; nothing re-runs the
  **neighbouring** fence.
- The **headline fix** (SSR origin) and the **F-6 guard** were both **entirely unfenced**, because each spans
  a chain (patch ⇄ injected env; ENV-check ⇄ built bundle) that no single iter's diff contains.

That is the argument for final mode existing at all: cumulative scope catches what iter-diff scope cannot see.
