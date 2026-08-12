# M258 iter-06 — progress

**Type:** tik · **Active strategy:** `TOK-01` steps **2 + 3** — *measure the composition before
engineering it.* Step 1 is discharged (bring-up 247.79 s, batch 129 s), so the strategy's own
precondition for treating the rest as a wiring job is satisfied.

## Phase A — the wiring

Three new files in `rosetta-extensions`, one hook, and one anchor. Zero platform-repo edits.

| file | what it is |
|---|---|
| `playthroughs/e2e/batch-gate.sh` | the `D-v28-3` batch gate: drive the full suite, emit ONE consolidated red set, restore the world, leave the stack UP, exit non-zero and loud on red |
| `playthroughs/e2e/restore-presenter-world.sh` | the world contract, resolution **(b) restore after** — reset-then-seed the stories preset, roster, cockpit + content manifests |
| `playthroughs/e2e/check-cockpit-roster.py` | the restore's **post-condition** (added mid-iter — see Phase C) |
| `playthroughs/e2e/stack-paths.sh` | resolve a stack's **LIVE** dir from the docker mount (added mid-iter — see Phase C) |
| `demo-stack/up-injected.sh:2839` | the hook, **after** the `UP.` line so "left UP regardless" is visible, not merely true |
| `stack-core/buildbench.py` | the `batch_gate` phase anchor + its own requirement class |

**The two decisions the wiring forced** are `D15` (the batch is default-ON but **skips on a
`--public-host` stack**, recorded `skipped`, never `green` — running it there would red all 30
Playthroughs on a TLS artefact) and `D17` (without a phase anchor the batch is **billed to autoverify**:
the table still sums, so nothing flags it — a *wrong* number, not a missing one).

## Phase B — the fences

- **The knob fence went RED in the direction it exists for** — `DEMO_NO_BATCH` *"has no row in
  `demo-up-defaults.md`"* (**UNDISCOVERABLE**). Row added; the three count mirrors moved **31 → 32**
  (`demo-up-defaults.md`, `demo/README.md`, `CLAUDE.md`). Now **OK, both directions**.
- **8 stale anchors repaired, all pre-existing.** The same fence reported `DEMO_NO_VERIFY` citing
  `:2762` against a parser reading `:2770`, and six siblings, every one off by **+8**. My edit is
  **purely additive from line 2821**, and `DEMO_NO_VERIFY` sits at 2770 in the *committed* tree too — so
  these were the `FIX-M257-anchor-guard-content-drift` family, not my drift. Each true line was
  **re-derived from the parser**, never assumed from the uniform offset.
- Corpus deliveries (both declared by `overview.md`): `verification.md` § *The layer ABOVE autoverify*
  (the contract, the skip, the world contract, the three fail-closed ledger rules) and
  `playthroughs.md` § *The BAKED-IN lifecycle*.
- `corpus_citation_guard` **OK** · `markdown_structure_guard` **OK** · `demo_knob_guard` **OK**.
  `corpus_index_guard` is RED on `node_modules/` + an M257x instrument dir — **pre-existing, none of it
  mine** (its only "batch" hits are `@graphql-tools/batch-execute`).

## Phase C — the controls, and what the LIVE run found that no control could

**16 controls** (`playthroughs/manifest/batch_gate_test.go`) execute the real script against a mirrored
temp tree, so they assert behaviour rather than spelling. The red direction is proven, not assumed:
non-zero exit, the loud block, **every** failing id named, and the artifact carrying the whole set. Three
fail-closed ledger rules each have a control — a **stale** report, an **empty** ledger, and a **non-zero
runner with an empty red set** ("the suite failed and nothing is failing" is not a green). Both mutants
tried were caught (red-branch `exit 0`; staleness guard disabled).

### The live finding — `D19`

The direct live run on `demo-1` went **GREEN** (215 passed, 31 use cases reconciled,
`failing=0 passing=30 unimplemented=1`, red set EMPTY, restore 7 s) **and the stack it left behind was
not presenter-usable.**

`rosetta-extensions` has **two** clone roles, so `demo-stack/stacks/demo-1/` exists **twice**. The live
one was the consumption clone's; the authoring copy's was stale since Aug 11 — with a working
`bin/stackseed`, a parseable manifest and plausible logs. `restore-presenter-world.sh` derived its stack
dir from **its own location**, so the roster went to the live path (`docker inspect`) while the cockpit
and content manifests went to **the stale one**. Measured: a **35-identity stories roster beside an
11-seat pt-world menu** — the exact stale-menu failure the restore exists to prevent, reintroduced by the
restore, while the gate printed *"presenter world restored"* and exited **0**.

**No unit test could have found this** — every fixture mirrors one tree. Two fixes landed:

1. **`stack-paths.sh::resolve_stack_dir`** — ask **docker** where the stack's files are. The container
   cannot be wrong about which directory the stack reads; a path derived from the script's own location
   is an assumption, the mount is evidence. `run-playthroughs.sh` already did this and *said so in a
   comment*; the rule existed and was simply not reused.
2. **`check-cockpit-roster.py`** — a post-condition, because the defect **passed every step-level
   check**: each export truthfully reported success and the damage lived in the *relationship between two
   artifacts*. It cross-checks the files **on disk** and fails on any advertised seat with no roster
   identity — which does not 404 but signs the visitor in as **whoever was last active**.

Re-run live after the fix: all three manifests to the live dir, **✓ all 12 cockpit seats resolve in the
35-identity roster**, restore **6 s**. `demo-1` verified back to 4 story orgs / 591 users / 12 presenter
heroes. The tests then had to be made **hermetic** — the resolver worked so well it found the *real*
daemon and wrote into an actual stack, so the harness now stubs `docker` (which keeps the resolver on the
code path rather than bypassing it).

## Phase D — the composed cold cycle: THE HOOK FIRES, AND THE ANCHOR EARNS ITS KEEP

`buildbench run 1 --reps 1 --profile macmini --no-public-host --label m258-iter06-composed`, launched
06:47:55Z at `load1` 5.71. `CAMP_RC=1` — the documented *"report is RED"* code.

| | |
|---|---|
| `total_s` | **545.22 s** |
| `up_rc` | **0** (the hook ran and the gate returned green) |
| `phases_complete` | **true** · `missing_anchors: []` · not-applicable: `serve_and_egress` (correct — single-box) |
| headroom / isolation | **OK / OK** |
| `green` | **false** — see below; **not** a platform red |

### The phase table — `D17` proven on a real run

| sub-phase | s | | sub-phase | s |
|---|---|---|---|---|
| host_preflight | 9.45 | | ui_hiring | 68.53 |
| secrets_provision | 1.34 | | compose_up | 44.30 |
| clones_and_inject | 1.91 | | set_dress | **80.42** |
| backend_builds | 96.52 | | **autoverify** | **2.41** |
| seed_tooling | 2.50 | | **batch_gate** | **166.36** |
| ui_next_web | 55.37 | | ui_studio_desk | 7.90 |

Sub-phases sum to **537.01** against `total_s` 545.22 — the 8.21 s difference is teardown + pre-state,
exactly as the model predicts. **`autoverify` reads 2.41 s and `batch_gate` 166.36 s**: without the anchor
those two would have been a single **168.77 s "autoverify"**, and the table would still have summed. That
is the wrong-number-not-missing-number failure `D17` describes, now measured rather than argued.

`batch-gate.json` corroborates from the other side: `verdict green`, `red_count 0`, **`batch_seconds 160`
+ `restore_seconds 7` = 167 ≈ the 166.36 s the anchor attributed**, `total 31`,
`passing=30 unimplemented=1`.

### Why `green: false`, precisely — and why it is NOT a platform red

`autoverify` failed **1 of 14** probes:

```
✗ postgres-schemas  fail: cannot derive the expected schema set (missing …/stack-core/lib/repos_yml.sh
                    or …/../platform/repos.yml) — refusing to assert a hand-maintained list
```

Measured cause: **`.agentspace/platform` does not exist**, while `stack-demo/platform/repos.yml` does. The
probe needs a **sibling `platform/` checkout**, and this cycle was driven from the **authoring copy** so
that it would exercise *this iter's uncommitted code*. So the probe is behaving exactly as
`verification.md` § *A guard that cannot find its subject must not exit 0* requires — **refusing rather
than asserting**. It is a topology artifact of an authoring-copy bring-up, not a defect in the stack, the
batch, or this iter's changes.

**The consequence is that this rep is correctly NOT gate-usable**, and the finding it produces is the
plan for the next iter: **the gate campaign must run from the CONSUMPTION clone at a pushed tag** — which
is the *tagging-is-not-publishing* discipline (`verification.md` pre-flight rung zero) arriving from a
direction nobody had named.

### The composed arithmetic — stated as arithmetic, not as a gate reading

| half | this cycle | note |
|---|---|---|
| bring-up | **370.65 s** | ⚠️ **cold-cache**: this clone had not built recently — `backend_builds` 96.52, `ui_hiring` 68.53 vs iter-05's 45.32 |
| batch + restore | **166.36 s** | batch 160 + restore 7 |
| **composed** | **545.22 s** | vs the **480 s** ceiling |

**545.22 s is NOT a gate number and must not be reported as one.** It is n=1, the rep is `green: false`
(so `rep_is_ok` excludes it), and its bring-up half is a cold-cache outlier against iter-05's **gateable
247.79 s**. It also **does not fire the `re_scope_trigger`**, which reads *"the composed **p50** exceeds
600 s after 3 tiks"* — this is one unusable sample, and firing on it would repeat precisely the category
error iter-04 refused.

Against the gateable bring-up half the projection is **247.79 + 166.36 = 414.15 s**, inside 480 by ~66 s —
**arithmetic across two runs, offered as a projection and not as a measurement.** The batch half also grew
(**160 s** here vs iter-04's **129 s**), so `C2`'s spread caveat now has a second sample on the batch side
and still no p50.

## Close — 2026-08-12

**Outcome:** **The batch gate is wired, and the milestone's gate is now measurable for the first time.**
`TOK-01` steps 2 **and** 3 landed together (`D16` — a batch without a restore leg is a regression, not a
partial delivery), proven live end-to-end: the hook fires from `up-injected.sh`, the suite runs to
completion (215 passed, 31 use cases, **red set EMPTY**), the presenter world is restored (**7 s**, vs the
20–45 s the plan assumed), the stack is left UP, and the phase table attributes **`batch_gate` 166.36 s**
beside a **2.41 s** `autoverify` instead of silently merging them. **The live run also found a real defect
in this iter's own restore leg** (`D19`) that every unit test passed over.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n *(the gate is a p50 over 3 cold cycles; this iter makes it
measurable and takes n=1, whose rep is correctly not usable)* — (2) triggered-tok: n — (3) re-scope: n
*(the trigger reads a p50 after 3 tiks; one unusable sample is not it)* — (4) user-blocker: n —
(5) cap-reached: n *(1 tik this session)* — (6) protocol-stop: n — (7) budget-exhausted: **y**
*(between iters, tree clean)* — Outcome: **exit-7**

> **Why exit-7 and not another tik.** iter-07 is the composed **3× cold campaign** — ~30 min of
> foreground polling on top of the two long live operations this session already drove (a 2.5-min batch
> and a 9-min cycle). The milestone's own `iteration_protocol` opens with **"fresh agent per run —
> context does not survive an 11-minute foreground op cleanly"**, and the campaign produces this
> milestone's *headline* number. Taking it on a heavily-consumed context is precisely the condition that
> rule exists to refuse. This is a clean between-iters stop, not an early one: the work landed, both
> trees are committed, and the tag is on origin.
**Decisions:** D15–D20 + the D18 addendum (this iter's `decisions.md`)
**Side-deliverables:**

- **`FIX-M257-anchor-guard-content-drift` — a live instance found and repaired.** `build-budget.md:538`
  cited `buildbench.py:2124` for `assert-headroom`'s instantaneous `os.getloadavg()[0]`; at the committed
  sha that line is `formatter_class=argparse…`. **It was already wrong before this iter** and passed
  silently because it lands on a *construct*, which is exactly the blind spot the routed item names.
  Re-pointed to `:2227` (verified: the site that feeds `headroom_assert`). The sibling `:375 → :393` was
  **mine** — my buildbench edits shifted it — and the fence caught that one because it landed on a blank
  line. Both re-derived from source, never offset-guessed.
- **8 pre-existing stale knob anchors** repaired in `demo-up-defaults.md` (Phase B).

**Routes carried forward:**

- **`TOK-01` step 4 — the composed 3× cold campaign** (iter-07). Now unblocked. **New precondition
  discovered this iter:** run it from the **consumption clone at a pushed tag**, not the authoring copy —
  an authoring-copy bring-up has no sibling `platform/`, so `autoverify`'s `postgres-schemas` probe
  refuses and every rep grades `green: false` and unusable, whatever the timings say.
- **`LEVER-M257-L5-setdress` — still the reserve, and still the largest bring-up phase** at **80.42 s**
  here (81.23 at iter-05, 82.04 at M257): unmoved across three measurements. If the composed p50 needs
  room, that is where it is.
- **`FIX-M258-iter03-guard-scans-its-own-scratch`** — unchanged and still open; **re-confirmed
  pre-existing** this iter (2 failures), and `test_fence_provenance::test_the_escape_accepts_and_records`
  is **the same family**: proven pre-existing by running both its RED members (`dev_flag_guard`,
  `demo_knob_guard`) against a pristine `HEAD` extract. Both only fire on a box that has run a demo.
- Unchanged: `ROUTE-M258-iter02-isolation-names-two-causes-not-three` ·
  `ROUTE-M258-iter02-headroom-defaults-to-billion` · `ROUTE-M258-iter02-purge-did-not-clear-the-stack-dir`.
- ⚠️ `demo-2` (11) and the dev stack (5) verified resident before and after every operation.

**Lessons:**

- **A path derived from the script's own location is an assumption; the docker mount is evidence.** The
  iter's worst defect was a stack directory resolved from `$EXT_ROOT` on a box with two rext clones. The
  correct rule already existed one file away, *in a comment*, and was not reused. When two copies of a
  tree can exist, ask the running system which one it is using.
- **A defect can pass every step-level check and still be real.** Each export reported success; the damage
  was in the *relationship between two artifacts*. That is what a post-condition is for — and why it
  compares files on disk rather than trusting the steps that wrote them.
- **Live proof finds what fixtures cannot, because a fixture mirrors one world.** 16 controls, two
  mutation checks and a clean shellcheck all passed over a bug that a single real run exposed in seconds.
- **A phase with no anchor is not un-measured, it is MIS-attributed** — and the table still sums, so
  nothing complains. Adding a step to an instrumented pipeline means adding its anchor in the same change.
- **The instrument refusing is a result.** `green: false` here is a probe declining to assert a
  hand-maintained list it cannot derive — the correct behaviour — and reading it as "the stack is broken"
  would have sent the next iter hunting a defect that does not exist.

## Isolation of the user's stacks

`demo-2` **11 containers** and the dev stack **5**, verified resident **before and after** every
operation in this iter. `demo-1` is mine.
