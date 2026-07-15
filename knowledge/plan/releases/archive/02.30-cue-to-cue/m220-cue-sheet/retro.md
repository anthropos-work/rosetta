# M220 "cue sheet" — Retro

## Summary

**The user's ask was three things. Two were ALREADY TRUE — the docs lied.**

- *"pull all the data"* — **already default-on.** The failure was a **cold cache** (M217's job), not a missing
  default.
- *"seed the 3 story orgs"* — **already default-on.** The preset ships 3. **The docs said "2 orgs" at 11 sites,
  live for four releases.** That single lie is why the user believed the seeding ask was unmet.
- *"remote access opt-out for demo"* — **the one genuine flip. Landed (S3).** A bare `up-injected.sh 1` now
  auto-discovers the tailnet host through a 6-rung capability ladder and serves over a trusted cert. Proven live
  from a tailnet peer (hero → `/profile` 200; manager → `/enterprise/workforce` 200). Opt out with
  `--no-public-host`. Dev stays **opt-in** (S7).

**The invariant that mattered most, and holds:** a box with no Tailscale falls back to **byte-identical
localhost behaviour** — proven by *genuinely removing* `tailscale`, not by mocking it. `SCHEME` and `BIND_HOST`
both derive from the same `-n $STACK_PUBLIC_HOST` predicate, so a **half-satisfied** public path is strictly
worse than localhost: every URL bakes `https://` against plain-HTTP listeners and the demo does not load at all.

**Zero platform-repo edits.** rext code-of-record: **`cue-to-cue-m220-final`** (live-graded on `billion` at
`-r6`).

## Incidents this cycle

D17 was the spine — *a status artifact that outlives the thing it describes, and is then read as evidence* —
and **several instances were found by the tooling in ITSELF.**

| # | What | Class |
|---|---|---|
| **S0-lie** | The plan said the "2 orgs" lie was at **4** sites; the KB audit found **7**; the **fence found 11**. A prose count never checked against the tree — the exact defect S0 existed to fix. | P1 · D17 |
| **S3-comment** | The ordering fence matched `--no-cockpit` inside a **comment** above the command; deleting the flag from the command still passed. | P2 · fence theatre |
| **S6/g cache** | The image cache had no idea which demo-patches were baked in (keyed on endpoint + pk only), so a stale image matched, the new patch **never baked**, and the bring-up graded green — the mechanism behind demopatch-spec's own 76-second war story. **Fired on its first live run.** | P1 · D17 |
| **S6 "reuse" tests** | Five `test_tag_guard_reuses_*` silently exercised the **rebuild** path — asserting the opposite of their own names while the suite looked fine. | P1 · the test was the bug |
| **D22 cert claim** | rext asserted `tailscale cert` "re-issues on re-run" — **measured false** (identical serial, 0.01 s, 0 ACME). Had it been true, default-on would burn a per-tailnet Let's Encrypt slot per `demo-up`. A doc lie that would have rate-limited a whole team. | P1 · false doc |
| **S7-battery** | The first S7 mutation battery `restore()`d with `git checkout` against **uncommitted** work, so mutants ran against a tree where the feature was **absent** and "went RED" for the wrong reason. The tell: unrelated mutants all reported an **identical** failure count. D17 inside the tool built to enforce D17. | P1 · fence theatre |
| **H-1** (harden) | **The HARD INVARIANT was unfenced.** The fence RE-TYPED the `SCHEME`/`BIND_HOST` derivation inside the test — beneath a docstring claiming it did not — so mutating the shipped `SCHEME` to unconditional `https` left all 23 tests green. A paraphrase can't disagree with itself. | P1 · fence asserted a copy of itself |
| **H-2** (harden) | The battery's baseline assertion fired on its OWN first run: 14 tests "RED" on an **unmutated** subject (a `/tmp` staging bug). Without it, all 7 dev mutants would have been logged RED — S7's original bug, reproduced live, caught by the guard built for it. | P1 · caught in the act |
| **H-3** (harden) | **`|| true` in `resolve_dev_public_host` was load-bearing and untested** (mutant V4 survived). A broken *tailscale* can't exercise it — the ladder is exit-0-always — only a broken ladder *process* can. | P1 · untested fallback |
| **D31** (close) | The dev-stack suite's "environmental" failure — **carried 5 milestones across 2 releases, re-characterised every time** — was **one missing env var** (`DEV_SETDRESS_USE_STUB_BINS=1`). The "145 % CPU" was the Go compiler. **D17 living in state.md's headline numbers.** | P1 · chronic D17 |
| **D32** (close, adversarial) | Rung 3 demanded a **dot**, not a **hostname**: `a.ts.net;rm -rf /`, `a.ts.net --bogus`, `-oProxyCommand.ts.net` all cleared 6 rungs and were returned as the public host — baked into the pk + argv. The fail-safe path was itself unsafe. | P1 · presence-probe in a capability ladder |
| **studio-desk (j)** | **NOT a defect (D15).** M219's *"302 → /login"* was a **cookieless curl** — the correct answer to an unauthenticated request, mis-read. A real manager clicking "Anthropos Studio" stays in Studio. **The retro does NOT repeat the false claim.** | P2 · false premise, corrected |

## The D17 thread — the release's spine, and it kept pointing at the tooling itself

Every P1 above is one shape: **a claim that outlived its evidence.** What made M220 different is *where* they
were found — **the fences kept catching themselves.** The exposure guard silently skipped its own repo (an
exclusion matched as a path substring). The org fence's `heroes?` regex missed bare `hero`. The ordering fence
matched a word in its own comment. The S7 battery graded a featureless tree. The HARD-INVARIANT fence asserted
a re-typed copy of itself. And the deepest one, D31, sat in the project's **own headline numbers** for four
releases, re-authored by every milestone that touched it and investigated by none — until the blocking
`audit-deferrals` gate forced the second look.

**The generalized keepers:**
- *"the doc says N" ≠ "the code ships N"*
- *"it serves" ≠ "it renders" ≠ "the session survives"*
- *"the flag exists" ≠ "the flag works"*
- *"the test is named X" ≠ "the test tests X"*
- *"the binary is installed" ≠ "it mints a cert"*
- *"tailscaled handed me a string" ≠ "it is a hostname"*
- *an errored command is not "zero results"*
- *a uniform result across unrelated mutations is a constant, not a result*
- **and the new one: *"we deferred this before" is not a reason to defer it again* — a re-characterised excuse
  is where a chronic bug hides.**

## What went well

- **The one genuine flip landed and is live-proven**, both vantages, on a trusted LE cert, cold reset-to-seed
  reproducible, 0 ejects. The open LE-rate-limit question was **settled empirically before shipping**, not
  trusted.
- **The fences are real because they were graded by mutation.** 17 mutants, 17 RED, and the three that were
  theatre when first written were caught by the mutation run — including the milestone's own HARD INVARIANT.
- **The chronic deferral was resolved instead of re-deferred.** The suite that "could never be run whole" — and
  was written into M221 as a release-close blocker — now runs; the whole rext Python suite completes for the
  first time in the release. One env var, after five milestones of excuses.
- **Two of the user's three asks were honestly reported as ALREADY-DONE**, with the real root cause (a doc lie,
  a cold cache) named — not re-implemented to look busy.

## What didn't

- **M220 shipped 5 new modules with zero README coverage** — the ladder and all four corpus guards. The exact
  undiscoverability defect S7 itself found in `--inject`, one level up, in the same milestone. Caught at close.
- **The academy is still half-done**: the session survives (S5) but the catalog renders empty (F-M220-2), and
  it binds `0.0.0.0` even on a localhost demo (F-M220-5, the S0 lie one layer up, invisible to a container-only
  exposure fence). Both honestly RED, both routed to M221 — the 400-char content floor was **not weakened**.
- **The demo-stack README test count had drifted to a stale 424** (actual 576) — a small D17 of its own,
  reconciled against the JUnit XML.

## Carried forward

| Item | Fate | Target |
|---|---|---|
| `FIX-M221-academy-empty-catalog` (F-M220-2) | Fate 3 | M221 |
| `FIX-M221-academy-loopback-bind` (F-M220-5) + extend the exposure fence to host-native listeners | Fate 3 | M221 |
| `F-M220-4` — `ant-academy.sh` re-runnable on a live public-host demo | Fate 3 | M221 |
| `BURNIN-M221-dev-public-host` — the dev opt-in is fenced, not live-proven | Fate 3 | M221 |
| `FIX-M221-devstack-test-spin` | **DISCHARGED** (D31) — retraction written into M221 overview | — |

## Metrics delta

- **Python tests:** **1215** (0 fail, 16 skip, 39 subtests) — demo-stack 583 · stack-injection 246 · stack-core
  152 · **dev-stack 120 (countable for the FIRST time)** · stack-verify 114. **The whole rext Python suite
  completes for the first time in the release** (D31). Like-for-like on the 4 sections M219 counted: 1095 vs
  903 = **+192**.
- **Go test funcs:** **1827** (+6 vs M219's 1821, same method). 0 failures, 6 modules, `go vet` clean.
- **Coverage (harden pass):** exposure_claim_guard 57→89 · tailscale_autohost 92→99 · gen_tailscale_serve 81→94
  · story_org_count_guard 62→91 · demo_knob_guard 80→91 · dev_flag_guard 71→74.
- **Mutation battery:** 17 mutants, 17 RED, committed (was: uncommitted, unfalsifiable).
- **Demo-patches:** 8 (+`next-web-no-thirdparty`). **Platform-repo edits: 0.**
- **Flake:** 0 (5 sequential randomized runs).
- **Deferrals:** YELLOW, 0 blocking, 0 escape-hatch; 1 chronic **resolved**.

## The one line worth keeping

**A re-characterised excuse is where a chronic bug hides. "Environmental" survived four releases because every
milestone that touched it rewrote the label instead of reading the code — one env var underneath the whole
time.**
