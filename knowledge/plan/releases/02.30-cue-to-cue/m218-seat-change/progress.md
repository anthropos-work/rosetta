# M218 — Progress

_Iterative milestone: a **running ledger**, not a section checklist. `/developer-kit:build-mstone-iters` appends one
entry per iter and creates `iter-NN/` dirs as it goes._

## Running ledger

| iter | kind | what changed | gate metric | outcome |
|------|------|--------------|-------------|---------|
| iter-01 | tok (bootstrap) | The 4-leg experiment, zero code. **C-1 confirmed as root cause but its MECHANISM was wrong in the plan — SSR reads the BUILD-INLINED public URL, which connect-times-out (10.56 s) from inside the container; the planned runtime-env one-liner is a NO-OP.** Budget now sums to the symptom. C-2 refuted (M217 fixed it — field proof). C-5 refuted as cause, retained as a free win. C-3 parked (unexercisable until a login completes). TOK-01 authored. | baseline **UNMEASURED** — no instrument exists (iter-02 builds it) | closed-fixed — see `iter-01/progress.md` |
| iter-02 | tik | **The latency harness** (rext `stack-verify/e2e/`, new surface — discharges DEF-M215-03(b)) + **the first baseline ever taken.** Found + fixed a measurement trap: Next.js **streams** RSC, so `page.on('response')` (headers) reports a *fast* document while the body blocks for 37 s — now measured via `response.finished()`. **iter-01's static prediction (37.5 s) and the measured SSR body (37.533 s) agree to within 33 ms.** | **employee p95 39.45 s · manager p95 38.30 s** vs the **< 5 s** gate → **~7.9× OVER** (3/3 reached ACCESS — slow, not broken) | closed-fixed — see `iter-02/progress.md` |
| iter-03 | tik | **The 38-second login is GONE.** Landed the real C-1 fix — the `next-web-ssr-graphql-origin` demo-patch (a **server-only** `WUNDERGRAPH_SSR_ENDPOINT`, deliberately not a `NEXT_PUBLIC_*` name so it is a **real runtime read**) + its runtime value in `gen_injected_override.py`. Also **F-6** (side): the next-web image cache-validator was blind to the **minted pk** (it reads image ENV; the pk is *inlined into the bundle*), so an out-of-band rebuild had left the stack **Clerkenstein-DEWIRED** — browser clerk-js talking to the **REAL Clerk app**, login broken, and the stale 9-h-old `autoverify.json` still saying green. Fixed by validating the pk **in-bundle**; it forced the self-healing rebuild on its first run. | **employee p95 39.45 → 7.90 s · manager p95 38.30 → 7.00 s** (cold, green, 6/6 ACCESS) — a **5× collapse**; **1.6× over** the 5 s gate (was 7.9×) | closed-fixed — see `iter-03/progress.md` |
| iter-04 | tik | **THE GATE IS MET.** Root-caused the ~6.10 s residual: Clerkenstein's fake **BAPI served a hardcoded STUB user to EVERY hero** (`// Disarmed: any id → the demo user` — true when a demo had one user, false since the M35 Stories & Heroes model). So `currentUser().externalId` (BAPI) disagreed with the JWT identity (FAPI) → `app`'s `userPreferences` resolver refused the mismatch → next-web's `retry: 2` / 2 s+4 s ladder burned ~6 s on **every** authenticated render. Fixed by making the BAPI **roster-aware**. `/align-run` ~~across **all 5 surfaces: 100% critical / 100% overall, 0 divergences**~~ — nothing moved. ⚠ **The "all 5 surfaces" claim is FALSE and is corrected at the close:** `expressrun` is **UNMEASURABLE** on this box (no `@clerk/express` `node_modules` → rc=2, **no score**), so only **4 of 5** were ever measured — and the harden pass reproduced that identically at baseline `f296e5e` (pre-existing, not a regression). The Go surface also now reads **97.2% / 100% critical**, by design (**D16**). Also landed the **M43-D5 correction** + the **`latency-budget.md`** blind area. | **employee p95 1.46 s · manager p95 1.40 s** vs the < 5 s gate — **5/5 ACCESS both vantages, gate armed, ~~cold~~+green**. From 39.45/38.30 s = **27× faster**. ⚠ **The "cold" claim was false** — iter-05 proved the DB was **warm** (F-9): this cycle is **not counted** toward the gate. | closed-fixed — see `iter-04/progress.md` |
| iter-05 | tik (verification) | **THE GATE IS MET *AS WRITTEN*.** iter-04 graded 5 cold *logins* on **ONE** stack = **1 of the 5** required cold reset-to-seed **cycles**. Ran the battery — and found the gate was not merely ungraded but **ungradeable**: **F-9**, `/demo-down --purge` **never purged on any Linux host** (postgres's UID-1001/0700 cluster dir defeats `rm -rf`; `set -euo pipefail` then **aborted `cmd_down`**, so the DB was never wiped, the images never removed, the registry slot leaked, and it returned a bare `rc=1`). **`billion`'s postgres still carried the `PG_VERSION` `initdb` wrote on 2026-07-11** — every "cold" bring-up for *days*, iter-04's included, reused the **same DB and same images**. Fixed in rext (+5 regression tests, green on Linux); count restarted at 0; ran **5 genuine cold cycles**, each **proven cold** (initdb re-ran) + **proven green** before being measured. Also **F-10** (a torn-down stack still serves a stale `green`) and **D15** (the alignment blind spot: Clerkenstein scored **100%/100%** while the fake BAPI returned **the wrong human for every request** — `getUser` has **no gene**, and the three "identity" genes all assert the **stub**). | **5/5 cycles PASS** · worst p95 **2413 ms** (employee) / **1767 ms** (manager) vs **< 5000 ms** — **2.07× margin**, **50/50 ACCESS**. Honest cold ≈ **1.6× slower** than iter-04's warm-DB number. | **closed-fixed — GATE MET (as written)** — see `iter-05/progress.md` |

## Next iter

**None — the milestone's exit gate is MET *as written*** (iter-05: 5/5 cold reset-to-seed cycles, both vantages,
worst p95 **2413 ms** vs < 5000 ms).

Next: ~~`/developer-kit:harden-mstone-iters --final`~~ **DONE (pass 1, 2026-07-14)** →
**`/developer-kit:close-milestone`**.

> ✅ **The final harden pass LANDED all three owed Fate-1 items** — see
> [`hardening-ledger.md`](hardening-ledger.md) and `decisions.md` → *CLOSED — Fate-1 items owed by this
> milestone*. **`/developer-kit:close-milestone` is unblocked.**
>
> - **F1-1** — the `GetUser` per-hero identity gene: **RED (`gate.sh` exit 2, critical 88.2%) against
>   `8ebc89e^`, GREEN (exit 0, critical 100%) after.** Measured on both sides — it fences the bug, it is not
>   theatre. Also surfaced a **second** uncovered endpoint (`…/organization_memberships`, studio-desk's admin
>   gate) → 2 more critical genes.
> - **F1-2** — `autoverify.json` is now unlinked on teardown **and** at bring-up start (`ts` added as
>   defence-in-depth). **7 regression tests, all 7 RED pre-fix.**
> - **F1-3** — the capability-coverage check **did not exist at all** (`alignctl dna` was
>   `list|diff|validate`; the "check" was an *eyeball* step in a skill). Now binding: `alignctl run` **refuses
>   to score** a DNA with an uncovered consumed endpoint. Corpus corrected.
>
> **Two things a reviewer should look at first:** (1) the Go alignment surface now scores **97.2% / 100%
> critical** — not 100% — because a **deliberately RED** gene (**F-11**, the ORG-level twin of the user stub)
> now tells the truth about a real divergence rather than omitting the field (**D16**); (2) the milestone's
> **headline fix had no test at all** until this pass (the SSR-origin chain — 12 tests added, mutation-proven).

## M218: Gate Outcome Ledger (close, 2026-07-14) — `closed-on-gate`

### Gate

| | |
|---|---|
| **Target** | p95 click→ACCESS **< 5000 ms**, BOTH vantages, over **5 consecutive cold reset-to-seed runs** |
| **Achieved** | **2413 ms** (employee `maya-thriving`) · **1767 ms** (manager `dan-manager`) — **worst case across the 5 cycles** |
| **Distance** | **2.07× under** the gate, worst-case. **50/50** logins reached ACCESS. |
| **Status** | ✅ **MET** — `closed-on-gate`. No carry-forward is *owed*; the 9 routed items below are net-new findings, not gate debt. |
| **Baseline** | **39.45 s / 38.30 s** → a **~16×** improvement **on the honest cold number**. |
| **Environment** | `billion.taildc510.ts.net` (Linux VM, 7.3 GiB RAM), demo-1, measured **over the tailnet origin**. *A latency number without its environment is not a measurement.* |

> ### ⚠ The gate number is **2413 ms**, not 1456 ms. Say the honest one.
>
> iter-04 reported **p95 1.46 s / 1.40 s** and declared the gate met. **That number was measured on a WARM
> database** and **is not the gate number.** iter-05 found **F-9**: `/demo-down --purge` had **never purged on
> any Linux host** — postgres's cluster dir is UID-1001/0700, so `rm -rf` died `Permission denied`, and
> `set -euo pipefail` then aborted `cmd_down` *silently*. `billion`'s postgres still carried the `PG_VERSION`
> `initdb` wrote on **2026-07-11**: every "cold" bring-up for days — **iter-04's included** — had reused the
> same DB and the same images.
>
> So iter-04 graded **5 cold logins on ONE stack** = **1 of the 5** required cold reset-to-seed **cycles**,
> and that one cycle wasn't cold either. The count was **restarted at 0**, and **5 genuine cycles** were run,
> each **proven cold** (`initdb` re-ran) and **proven green** *before* being measured. The honest cold number
> is **~1.6× slower** than the warm one. **It still clears the gate by 2×.**

### Iter ledger

| iter | kind | closed | commit |
|---|---|---|---|
| iter-01 | tok (bootstrap) | ✅ | `2c1e950` |
| iter-02 | tik | ✅ | `07caacd` |
| iter-03 | tik | ✅ | `2f37522` |
| iter-04 | tik | ✅ | `27aa7ca` |
| iter-05 | tik (verification) | ✅ | `8354b06` |
| harden (final, pass 1) | — | ✅ | `7c6a48d`, `d460d89` |
| close | — | ✅ | this commit |

**5/5 iters have a dir + a closed `progress.md`. No orphan iters. Every commit maps to an iter or to
harden/close. No orphan commits.**

### Routes carried forward (9 — all **Fate 3**, all with the receiving milestone's `overview.md` **EDITED**)

None of these block the gate; none are gate debt. **All 9 share one root cause for not landing in M218:**
the exit gate is *a p95 over 5 cold cycles graded on a specific binary*, and **iter-05 D13 established that a
demo-runtime change restarts the count**. Landing them post-gate would ship **something other than what was
measured** — the precise sin iter-05 spent its whole budget eradicating.

| → | item | why not here |
|---|---|---|
| **M219** | **F-11** `FIX-M219-bapi-org-eid` — the BAPI fabricates the org eid | runtime change + needs a fresh 5-cycle battery. Ships as a **deliberately RED gene** (D16). |
| **M219** | `TEST-M219-expressrun-dep-gate` — `expressrun` is UNMEASURABLE (rc=2, **no score**) | pre-existing; a missing dep must fail **loud**, not score-absent |
| **M219** | `TEST-M219-freshness-gate-skips` — the demopatch freshness gate **skips** without the clone | *absence read as success*, the D17 class |
| **M220** | Clerk telemetry off (`CLERK_TELEMETRY_DISABLED`) — **never wired**, 0 grep hits | injected-env change; M220 owns `/demo-up` defaults |
| **M220** | **F-5** ad-tech egress (GA + DoubleClick + Google Ads + LinkedIn) on every authenticated load | same surface |
| **M220** | **C-5** vendor clerk-js + bound the **unbounded `Timeout: 0`** (`server.go:187`) | an **unbounded internet dependency in the login path** of a "self-contained" demo → `safety.md` **Part 3** |
| **M220** | ant-academy is handed the **REAL** `CLERK_SECRET_KEY` (`ant-academy.sh:146`) | `safety.md` violation, fix16/17 class; pure env |
| **M221** | **F-7** `NEXT_PUBLIC_BACKEND_API_URL` — a **measured 10.5 s blackhole** from inside the container | **a loaded gun**: dormant only because every reader is client-side. Prove it **on the box**. |
| **M221** | **C-3** cms/Directus **403s** on the CONTENT path (now exercisable for the first time) | threatens M221's gate item (2): *"the full replayed catalog — no SKIPPED surface"* |

### Dropped

**None.**

### Escape-hatch (release-scope-breaking) deferrals

**ZERO.** Nothing leaves v2.3.

### Known issue carried (flagged for sign-off, NOT a regression)

The **`dev-stack` suite fails on this box** — environmental (needs a live Postgres on `:15432`; this box's
`.agentspace/secrets` also lacks critical keys, so the secret-coverage pre-flight aborts *before* the
N-guard). **Identical at v2.2 and at M217. M218 does not touch `dev-stack/`.**

### Protocol evolution

The milestone's own iteration protocol held, with one correction that is worth carrying forward:
**"reproduced at baseline ⇒ pre-existing" is only sound if the baseline predates the milestone.** The final
harden pass used `f296e5e` — *M218's own iter-05 commit* — as a "baseline", and on that basis misfiled a
regression **this milestone caused** as "pre-existing, M217 footprint". Re-measured against the true pre-M218
ref (`cue-to-cue-m217`), the test **passes there and fails at HEAD**. Future harden/close passes must resolve
"baseline" to **the milestone's own start ref**, not to any commit that merely predates the current pass.

## Carry-forward queue (none block the gate)

| handler | item |
|---|---|
| **`PROBE-M221-backend-api-url-twin`** (was `PROBE-M218-…`; **Fate-3 → M221**, its `overview.md` edited) | **F-7:** `NEXT_PUBLIC_BACKEND_API_URL` bakes to `https://billion…:18082` — **measured as a 10.5 s blackhole from inside the container**, the exact C-1 shape. Dormant **only** because every reader is client-side (D10). **A loaded gun.** |
| **`FIX-M220-telemetry-egress`** (was `FIX-M218-…`; **Fate-3 → M220**, its `overview.md` edited) | **F-5:** the demo attempts **Google Analytics + DoubleClick + Google Ads + LinkedIn Ads** on every authenticated load (+ the in-scope Clerk-telemetry off). |
| **`FIX-M220-c5-clerkjs`** (was `FIX-M218-…`; **Fate-3 → M220**, its `overview.md` edited) | **C-5:** vendor clerk-js + bound the **unbounded** `Timeout: 0` (`server.go:187`). Alignment-invisible ⇒ gate-free. |
| **`PROBE-M221-c3-rerun`** (was `PROBE-M218-…`; **Fate-3 → M221**, its `overview.md` edited) | **C-3:** now exercisable. The router **is** logging cms/Directus **403s** (`getSkillPaths`, `_entities JobSimulation`) on the CONTENT path — not the login path; affects data-settle. |
| ~~`DOC-M218-audit-corrections`~~ ✅ **CLOSED** | **DONE for M43-D5 + `latency-budget.md`** (iter-04) **+ the CI-inert correction + F1-3's coverage claim** (harden pass — and the CI claim was *worse* than inert: **rext has no `.github/workflows` at all**). ~~**Remaining:** the `clerkenstein.md:3-4` header.~~ ✅ **DONE at the close** — plus the header's *self-contradicting* repo line, the **undelivered** clerk-js caching/timeout contract, and the **7 docs still claiming 100%/all-five**. |
| ~~`HARDEN-M218-F1-1`~~ | ✅ **LANDED** (harden pass 1). RED @ `8ebc89e^` / GREEN @ HEAD, proven both sides. |
| ~~`HARDEN-M218-F1-2`~~ | ✅ **LANDED** (harden pass 1). 7 tests, all RED pre-fix. |
| **`FIX-M219-bapi-org-eid`** (**F-11**, new) | The BAPI fabricates `organization.public_metadata.eid` as `"org_eid_"+orgID` instead of the roster's **real** org UUID — the **ORG-level twin** of the user stub. Needs a **runtime** change + a **fresh 5-cycle battery** (iter-05 D13), so it could not land post-gate. Shipped as a **deliberately RED gene** so the score stops lying (**D16**). |
| `TEST-M219-expressrun-dep-gate` | `expressrun` is **UNMEASURABLE** without `@clerk/express` `node_modules` (rc=2, *no score*) — **pre-existing** (identical at baseline `f296e5e`). So iter-04's "all 5 surfaces 100%" is **not reproducible on this box**; 4 of 5 were re-measured. |
| `TEST-M219-freshness-gate-skips` | The demo-patch live-clone freshness gate **skips** when the `stack-demo/next-web-app` clone is absent — so a box without it gets **no anchor-drift protection**. Itself an instance of *absence read as success*. |
| _(also-in-scope, from overview)_ | ant-academy real-Clerk-secret leak → **Fate-3 → M220** (`overview.md` edited). `x/crypto@v0.52.0` → ✅ **LANDED at the close** (Fate-1) — it was a **repeat deferral** (v2.2 → M218's rext roll) and this close **is** the rext roll. Bumped 0.51.0 → 0.52.0; the alignment gate scores **identically** pre/post, which is how we know it is behaviour-neutral. |

## M218: Final Review (close, 2026-07-14)

The close review (deferral audit + 2 parallel cross-cutting scans + adversarial pass + full-suite re-run)
found **17 items**. **All landed as Fate 1.** Every one of the four most serious is the **same class the
milestone itself named as D17** — *a status artifact that outlives the thing it describes, and is then read
as evidence* — and every one **survived both the iter loop and the final hardening pass**.

### Scope
- [x] **A REGRESSION, MISFILED AS "PRE-EXISTING".** `stack-injection::test_next_web_block_shape` was RED. The
      hardening ledger booked it *"pre-existing, out of M218 scope — M217 footprint"* on the evidence that it
      *"reproduced at baseline `f296e5e`"*. **`f296e5e` is M218's own iter-05 commit.** Re-measured against
      the true pre-M218 ref (`cue-to-cue-m217`): **PASSES there, FAILS at HEAD.** iter-03's C-1 fix added
      `WUNDERGRAPH_SSR_ENDPOINT` to the next-web block and left the exact-shape fence pinned to the old one.
      The suite had been red since iter-03. → **fixed**, with the reason written into the test; the ledger's
      false claim **corrected in place**. *"Reproduced at baseline ⇒ pre-existing" is only sound if the
      baseline predates the milestone.*
- [x] **The 9 carry-forward handlers were tagged `*-M218-*` on a milestone that is closing** — the
      "future milestone (unspecified)" anti-pattern in disguise. All re-fated to **M219 / M220 / M221**, and
      **each receiving `overview.md` was EDITED** (Fate 3 means the sibling's plan changes, not that a string
      changes here).
- [x] **`x/crypto@v0.52.0` — a REPEAT + AGED-OUT deferral.** v2.2 → "Fate-1 → M218's rext roll" → still not
      landed. This close **is** the rext roll. → **bumped** (0.51.0 → 0.52.0).

### Code Quality
- [x] **[must-fix] MF-1 — `--purge` failure LEAKED the registry slot and the images.** iter-05 fixed F-9's
      *silence* and kept F-9's *abort*: `die` on a failed purge exits **before** `docker rmi`, `reg_del` and
      `ureg_release` — **the exact three items `test_purge.py:18-20` enumerates as F-9's collateral damage**.
      The milestone documented the hazard and then walked into it, one branch over. → recorded and re-raised
      **last**; the teardown always reclaims.
- [x] **[must-fix] MF-2 — a FAILED FRONTEND BUILD graded GREEN.** `build_frontends()` returned 0 regardless,
      so compose started whatever `demo-N-next-web` image was on disk — **stale, unpatched, possibly
      Clerkenstein-dewired** — and nothing proved the running image was this run's build. **M218 is what made
      `autoverify.json` a GATE INPUT**, so a pre-existing false-green became **load-bearing**: the latency
      battery could have been graded on a stale image. → `buildfail.log` + an autoverify assert.
- [x] **[must-fix] MF-3 — a SKIPPED demo-patch graded GREEN, ON THE GATE PATH.** Only the *refusal* branches
      wrote `demopatch.log`; the **four skip branches wrote nothing**, and autoverify keys on non-emptiness.
      **A stack running WITHOUT the SSR-origin patch — this milestone's headline fix, ~37.5 s per
      authenticated render — printed "✓ demo-patches: none refused", graded green, and was gradeable.** →
      every non-applying branch now leaves evidence; autoverify asks *"is this demo patched"*, not *"was
      anything refused"*.
- [x] **[should-fix] `cmd_up` never cleared the verdict** — F-10's *"cleared at the START of every bring-up"*
      invariant, which `clear_stack_verdict()`'s own contract asserts **unconditionally**, held on **one of
      two** bring-up verbs.
- [x] **[should-fix] The `ts` field was write-only.** Added as "defence in depth" so a grader "can SEE that
      the verdict predates the stack" — **and nothing read it.** A safeguard that exists only in prose: the
      **F1-3 class, one layer down.** → `run-latency.sh` now **ages** the verdict and refuses a stale green.
- [x] **[should-fix] `stopBAPIServers()` left a fired `sync.Once` pointing at a deleted binary** — a second
      in-process `run()` would exec a removed file and turn every BAPI gene red. Latent (today's fixture
      declares no BAPI capabilities), and the same family: state outliving its subject.
- [x] **[should-fix] `purge_data_dir` discarded the container's diagnostics** — a docker-down / image-unpullable
      failure was undiagnosable. The *"reason piped to `/dev/null`"* pattern, again.
- [x] **[should-fix] Playwright run output was COMMITTED** (`latency-out/*.json`) against the section's own
      "never commit run artifacts" convention — every measurement dirtied the tree.
- [x] **[nice-to-have] `LATENCY_RUNS` coerced to `NaN`** → the sample loop never ran → **0 samples measured**,
      surfacing as `expect(0).toBe(NaN)`. A gate that measures nothing must fail loudly.
- [x] **[nice-to-have] `cockpit.py`**: a mangled comment block, and a stage whose copy ("this can take a few
      seconds") now contradicts the measured p95 directly above it.

### Documentation
- [x] **M218's own "CI-inert correction" was ITSELF FALSE.** It replaced *"a weekly CI workflow gates it"*
      with *"rext has **no** `.github/workflows` at all"*. **The workflow EXISTS and is git-tracked** at
      `clerkenstein/.github/workflows/alignment.yml` — it simply **never runs**, because GitHub Actions only
      reads `.github/workflows` **at the repo root**. **"Inert" ≠ "absent"**, and a reader who greps for the
      file finds it. → corrected, with the mechanism named.
- [x] **The coverage-gate section OVER-CLAIMED — the exact failure mode it was written to fix.** It said an
      undeclared surface **exits 2**. `gate.sh` passes **`--if-declared`**, which downgrades that case to a
      **warning**. So **4 of the 5 surfaces are *warned about*, not *enforced***. → stated plainly.
- [x] **`alignment_testing.md` never recorded the 97.2% score** (grep `97.2` → **0 hits**). D16's whole point
      was that the score stop lying — and the corpus went on saying 100%. → the per-surface score table, the
      deliberately-RED gene, and the `expressrun`-is-unmeasurable caveat all recorded.
- [x] **7 further docs still claimed "100% / 100% on all five surfaces"** (`clerkenstein.md`, rext
      `CLAUDE.md` / `README.md` / `scope.md` / `topic-map.md` / `kb-index.md`, `demo-stack/GUIDE.md`). One
      also said **"two DNAs"** (there are five) and two said **"all four"** surfaces. → all corrected.
- [x] **The `clerkenstein.md:3-4` header** (the milestone's own last owed doc item) — stale to **2026-06-24**,
      omitting M213/M217/M218 — **and its repo line was refuted by a block 11 lines below it** ("its own git"
      vs *One monorepo, two clone roles*).
- [x] **The clerk-js proxy's caching/timeout contract — a `Delivers →` item that NEVER LANDED.** grep
      `Timeout|cache|unbounded` in `clerkenstein.md` → **0 hits**. → documented: bare `http.Get` =
      `DefaultClient` = **`Timeout: 0`, unbounded**; **no server-side cache**; **on the login path**;
      0.2 s healthy / **~127 s if egress blackholes**.
- [x] **`latency-budget.md` was UNDISCOVERABLE** — a brand-new corpus doc with exactly **one** inbound link,
      mid-body. Not in `demo/README.md`, not in `CLAUDE.md`. → indexed in both.
- [x] **M43-D5 was never formally superseded**, though the milestone's `Delivers →` said *"re-open it
      formally"*. The archived decision still asserted *"the real ~2–5 s latency is **unavoidable**"* with no
      marker — **the very claim that stopped four releases from investigating.** → superseded, in the archive
      and in `roadmap-legacy.md`.
- [x] **The milestone's own artifacts repeated the false "all 5 surfaces 100%" claim** (`progress.md` iter-04
      row; `iter-04/decisions.md`). → both corrected in place.
- [x] `stack-verify/README.md` (the e2e unit gained a whole **latency harness** and its handbook never
      mentioned it — "two layers" → **three**) and `alignment/README.md` (`alignctl dna coverage`, the
      milestone's headline tooling addition, absent from its own unit's handbook).

### Tests & Benchmarks
- [x] **+12 fences** (`demo-stack/tests/test_close_findings.py`) — **ALL 12 RED against `f849b5f`**, 12/12
      green after. Not theatre: measured on both sides.
- [x] **One of those fences initially FALSE-PASSED pre-fix** — its `cmd_up` slice ran to `cmd_down` and
      **swallowed the `clear_stack_verdict()` DEFINITION**, so it matched a definition instead of a call.
      **D17 in test form, inside the fence written against D17.** Caught *only* by running the file against
      the pre-fix tree — which is the entire reason it is now a real fence.
- [x] Full-suite re-run: **Python 887 passed / 0 failed** (JUnit-counted, never grepped) · **Go 0 failures
      across 6 modules**, 1784 test funcs · **alignment gate rc=0** (97.2% / 100% critical) · **flake gate
      5/5 clean, sequential**.

### Decision Triage
- [x] **D16** (honest 97.2% > hollow 100%) → blended into `corpus/architecture/alignment_testing.md` +
      `corpus/services/clerkenstein.md` (#M218-D16).
- [x] **D17** (the stale-verdict hazard) → already in `corpus/ops/verification.md`; **reinforced by five more
      instances found at the close** (MF-1/2/3, the `ts` field, the `sync.Once`).
- [x] **D18** (the baseline-resolution rule) → new; recorded below and in the Gate Outcome Ledger's *Protocol
      evolution*.
- [x] iter-level decisions (D1–D15, TOK-01) → stay in `iter-NN/decisions.md` as maintainer archive.
