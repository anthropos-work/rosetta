# M218 — Decisions

_Implementation + strategy choices with rationale (incl. tok strategy revisions)._

| # | Decision | Rationale | Date |
|---|----------|-----------|------|
| — | _(intra-iter decisions live in `iter-NN/decisions.md`)_ | | |

---

## ✅ CLOSED — Fate-1 items owed by this milestone (all three LANDED, 2026-07-14)

> **Discharged by `/developer-kit:harden-mstone-iters --final` (pass 1).** Full evidence:
> [`hardening-ledger.md`](hardening-ledger.md). `/developer-kit:close-milestone` is unblocked.
>
> | # | Outcome | The proof (measured, not asserted) |
> |---|---|---|
> | **F1-1** | **LANDED** | `gate.sh` exit **2 (RED)** @ `8ebc89e^` — `GetUser` **0/2**, critical **88.2%**; both heroes came back as the stub `11111111-…`/`demo@anthropos.test`. `gate.sh` exit **0 (GREEN)** @ HEAD — `GetUser` **2/2**, critical **100.0%**. The gene fences the bug. |
> | **F1-2** | **LANDED** | `clear_stack_verdict()` runs **first** in `cmd_down` (before any step that could abort it), plus an unlink at bring-up start and a `ts` field. **7 regression tests — all 7 RED against `f296e5e`.** |
> | **F1-3** | **LANDED** | The check **did not exist** (not merely "did not bind"). Now: `consumed_surface` + a `Validate` rejection that makes `alignctl run` **refuse to score**, `alignctl dna coverage`, wired into `gate.sh`. Doc rewritten to state what it guarantees **and what it does not**. |
>
> **Two new blind spots surfaced while landing them** — `GET /v1/users/{id}/organization_memberships`
> (studio-desk's admin gate: no gene in any DNA either → now 2 critical genes) and **F-11**, the ORG-level
> twin of the user stub (routed forward, see below).

_Original statement of the three items, kept verbatim for the record:_

**Handler: `/developer-kit:harden-mstone-iters --final`** (runs after the gate fires, before
`/developer-kit:close-milestone`).

| # | Item | Why it is Fate 1 (land it completely, here) | Definition of done |
|---|---|---|---|
| **F1-1** | **The alignment blind spot — a `GetUser` gene with per-hero identity on the BAPI surface (`clerk-2.6.0`).** `/align-run` scored Clerkenstein **100% critical / 100% overall, 0 divergences — before *and* after** iter-04, while the fake BAPI was returning **the wrong human for every request**. Verified cause: `getUser` has **no gene in any of the 5 DNAs**, and the three genes that *do* name identity (`ExtractIdentity`, `Me`, `DeployIdentity`) all assert the variant **`universal-user`** — *the stub itself*. The goldens **ratified the defect**. `DistinctIdentity` (the only per-hero gene) exercises the **FAPI/JWT** path — the half that was already correct. | It is a **regression golden for a bug this milestone just fixed**, and a mirror that serves a stub identity to every hero scoring 100% makes the score untrustworthy for every future mirror. Small; exactly what a harden pass is for. **Deliberately not done inline in iter-05:** a code change **restarts the 5-cycle count** (iter-05 D13), so it is sequenced *after* the gate fires. | The gene exists, is `critical`, and is **red against `8ebc89e^` / green after**. If it passes on both, it is not fencing the bug. |
| **F1-2** | **Teardown must remove `autoverify.json`** (F-10). It is *not* removed on teardown, so a torn-down stack leaves a `green:true` verdict on disk for a stack that no longer exists — read by both `run-latency.sh`'s green gate and any grader. Caught live in iter-05: a **failed** bring-up (0 containers) still presented `{"green":true,"warnings":0}`. Same class as **F-6**. | A grader that can be handed a stale `green` cannot be trusted to refuse a broken stack — which is the entire premise of the M217→M218 hard barrier. | `cmd_down` unlinks it; a test fences it. |
| **F1-3** | **The capability-coverage check is not binding.** `alignment_testing.md:169–172` offers, as the *named mitigation* for a hollow score, "`/align-dna`'s capability-coverage check (**every consumed endpoint is present**)". `GET /v1/users/{id}` **is** consumed (next-web's server-side `currentUser()`, every authenticated render) and has no capability — so the safeguard did not bind. | The doc currently **over-claims a guarantee the tooling does not provide**. Either make it bind, or say so. | Either the check covers the fake BAPI's read surface, or `alignment_testing.md` states honestly that it does not. Joins `DOC-M218-audit-corrections`. |

_Full analysis + evidence: `iter-05/decisions.md` **D15**._

---

## D16 — Ship an honest 97.2% over a hollow 100% (the deliberately RED gene) — 2026-07-14

**Handler:** `/developer-kit:harden-mstone-iters --final`, pass 1.

**What surfaced.** Landing F1-1 exposed a **second** identity stub, one layer up. The demo roster carries each
hero's real internal **org** UUID (`org_eid`), and real Clerk reports it in
`organization.public_metadata.eid` (the platform syncs it there via `UpdateClerkOrganizationWithExternalId`).
Clerkenstein's `organizationWithEid` instead **fabricates** `"org_eid_" + orgID` for any org that isn't the
hardcoded demo org. It is the **ORG-level twin of the user-level stub that cost this milestone ~6 s per
render** — the same defect, the same blind spot, one field over. Nothing measured it. (**F-11**)

**The bind.** Three options, all bad in different ways:

1. **Fix it inline.** `resources.go` is the demo's **runtime** path. The milestone's exit gate is *a p95 over
   5 cold reset-to-seed cycles*, graded on the current binary — and **iter-05 D13 established that a code
   change restarts the count**. Fixing it post-gate means **shipping something other than what was measured**,
   which is precisely the sin iter-05 spent its whole budget eradicating (grading a stack that was not what it
   claimed to be). Rejected.
2. **Omit the field from the gene.** The gene goes green, the score stays 100%, and nobody ever hears about
   it. This is *exactly* how the headline bug survived four releases — **a silently-omitted field**. Rejected,
   emphatically.
3. **Land it as a RED gene.** The mirror is genuinely not 100% faithful, so **the score should not say 100%**.

**The decision: option 3.** `MembershipOrgIdentity/real-org-eid` ships **red**, `standard` weight. The Go
surface now reports **97.2% overall / 100% critical** — the gate (**≥95 / =100**) is still **MET**, and the
divergence is named in the report on **every single run** until someone lands the fix with a fresh battery:

```
FAIL MembershipOrgIdentity/real-org-eid  (exact, w2)
     value differs: source={"org_eid":"1d0e6c22-…"} mirror={"org_eid":"org_eid_org_seed_demo-1"}
```

**Why this is the right trade.** The entire thesis of this milestone is that **a 100% score which hides a lie
is worse than an honest 97%** — Clerkenstein scored 100%/100%/0-divergences while returning the wrong human
for every hero. Restoring a clean 100% by *looking away from* the next stub would reproduce the failure mode
in the very pass convened to end it. The score is now a measurement rather than a decoration.

**Routed forward:** `FIX-M219-bapi-org-eid` (needs a runtime change + a fresh 5-cycle battery).

**⚠ Reviewer's note.** This deliberately reduces a reported score from 100% → 97.2%. If that is unacceptable,
the honest alternative is **not** to omit the gene but to land the runtime fix *and re-run the 5-cycle cold
battery* — never to drop the assertion.

---

## D17 — The stale-verdict hazard is the *class*, not five bugs — 2026-07-14

M218 hit the **same failure class five times**: **F-6** (a 9-h-old `autoverify.json` graded a
production-Clerk-wired stack green), **F-9** (`--purge` returned `rc=1`, purged nothing, said nothing),
**F-10** (a green verdict for a stack with **zero containers**), plus two probe-level instances (`[ -e ]`
reading *permission-denied* as *absence*; `assertNotIn` passing on a **failed** command's empty output). It
had **already survived M217's hardening**.

**Decision:** stop fixing instances. Name the class and give it invariants —
[`corpus/ops/verification.md` → *THE STALE-VERDICT HAZARD*](../../../../../corpus/ops/verification.md):
**a status artifact that outlives the thing it describes, and is then read as evidence.**

1. **A verdict must not outlive its subject.** Destroy it on teardown **and** at the start of every bring-up —
   and destroy it **first**, before anything that could fail and abort the sequence (exactly how F-9 leaked).
2. **Absence must be the safe state.** A grader with no verdict **refuses to measure**. Nearly every instance
   is the same mistake in different clothing: treating *"nothing here"* as *"nothing wrong."*

**Corollary, applied to this pass's own tests:** a probe that can pass **without ever executing its
assertion** is a stale verdict in test form. The new regression tests count what they inspected and **fail if
the count is zero**.

**Sibling hazard, same family:** a **safeguard that exists only in prose** — which is precisely what F1-3
turned out to be (`alignment_testing.md` named a coverage check that had never been written).

---

## TOK-01: Reachability-first — fix the address, not the variable — 2026-07-13

**Tok type:** bootstrap (iter-01)

**Initial strategy:**
**Restore the SSR GraphQL call, then re-measure and let the next bottleneck name itself.** Concretely, in order:

1. **Build the latency harness first (iter-02), and take a real p95 baseline before any fix.** Playwright, real
   https origin, both vantages, per-leg attribution (click → handshake/303 → SSR → clerk-js → client-gate →
   data-query). Reuse `stack-verify/e2e/lib/cockpit-login.ts` (`selectSeat`/`loginAs`) for the handshake, but
   **parameterize its wait strategy** — do NOT inherit its `waitUntil:'networkidle'` (the milestone bans it; the
   audit found the collision). Gate every run on `autoverify.json` = green.
2. **Then fix C-1 by making the baked public URL REACHABLE from inside the container** — *not* by re-pointing the
   variable (that is a proven no-op, D1) and *not* by re-baking it (that breaks the browser). See **iter-01 D2**:
   compose-level host-aliasing in `stack-injection`, zero platform edits.
3. **Then re-measure and re-run the parked legs.** C-3 (router retries) is **unexercisable today** because the SSR
   fetch dies upstream of the router; the moment a login completes, the federation gets exercised for the first
   time and C-3 becomes measurable. **Expect a new bottleneck to appear** — the fix does not end the milestone, it
   *unblocks the measurement of the next layer*.
4. **Take the free, gate-free wins while in the area** — C-5 (vendor clerk-js + bound the unbounded timeout;
   alignment-invisible), Clerk telemetry off, ant-academy's real-Clerk-secret leak, `x/crypto` bump.
5. **Pay the doc debt as we go, not at the end** — `latency-budget.md` (the blind area), the post-303 login
   sequence, the M43-D5 correction (**4+ sites**, anchors corrected by the audit), the CI-inert correction
   (**`:232,233,239`**), and F-2/F-3/F-4.

**Rationale:**
The 4-leg experiment (iter-01) closed the arithmetic gap that made this milestone iterative in the first place. The
milestone's own framing — *"the confirmed cost budget does not sum to the symptom (~18 s vs 60–120 s), so writing a
fix now means guessing which fix to build"* — is now **resolved**: the budget sums. One defect (**C-1, with its
mechanism corrected**) accounts for **~112 s** on the tailnet demo, hits **both** vantages via the shared
authenticated layout, and is ~20× cheaper on a laptop, which is exactly why four releases of local measurement
never saw it.

The strategy is **"reachability-first"** because the *shape* of the fix is the milestone's real risk. The plan's
one-liner targets the **wrong variable** (`process.env`, which SSR never reads). The correct target is the
**address itself**: one build-time constant serves two consumers with incompatible reachability needs, so the only
move that satisfies both is to make the address the browser needs also **work from inside the container**. That
keeps the fix in `stack-injection` (rext-only), preserves the browser's correctly-baked URL, and honours the hard
zero-platform-edit constraint.

Harness-before-fix is retained (not skipped) even though the cause is already known — because the **gate is a p95
over 5 cold runs**, and there is currently **no instrument that can produce that number**. Without the harness we
could ship a correct fix and still be unable to *prove* the gate. The harness is also the milestone's declared
discharge of **DEF-M215-03(b)**.

**Strategy class:** new-direction

**Distance-to-gate context:**
- **Gate:** p95 click→ACCESS **< 5 s**, both vantages, over 5 consecutive cold reset-to-seed runs.
- **Starting value: UNMEASURED — no instrument exists.** The only numbers today are the user's report (**60–120 s**)
  and iter-01's component measurements (handshake **17 ms**; SSR connect-timeout **10.56 s** × 3 attempts + 6 s
  backoff ≈ **37.5 s** per blocking fetcher; ×3 fetchers ≈ **112 s**). **Establishing the baseline is iter-02's
  entire job** — the gate cannot be graded until an instrument exists.
- Expected shape: iter-03's fix should collapse the dominant term from ~112 s to ~0.1 s (94 ms measured on the
  working origin). **Whether that alone lands under 5 s is unknown** — the parked C-3 and the client-gate leg are
  unmeasured. Assume at least one more bottleneck.

**Next-tik direction (iter-02):**
Build the Playwright latency harness in a **new `stack-verify` surface** (not a Playthrough — Playthroughs declare
perf a **non-goal**). Deliverables: per-leg attribution, `autoverify.json` green-gating (at its **real** path — D3),
p95 over N runs, both vantages. Ship it, run it against the **current unfixed** `billion` demo, and **record the
honest pre-fix baseline** — that number is the milestone's headline and the M43-D5 correction's evidence.
**Write no fix in iter-02.**

---

## D18 — "Baseline" must resolve to the milestone's START ref, not to any earlier commit — 2026-07-14

**Handler:** `/developer-kit:close-milestone`, Phase 4.

**What surfaced.** The final harden pass declared `stack-injection::test_next_web_block_shape` *"pre-existing,
out of M218 scope — M217 footprint"*, on the evidence that it **"reproduced at baseline `f296e5e`"**.

`f296e5e` is **M218's own iter-05 commit**.

Re-measured against the **true** pre-M218 ref — the rext tag `cue-to-cue-m217`:

| ref | result |
|---|---|
| `cue-to-cue-m217` (true pre-M218 baseline) | **PASSES** (`1 passed, 131 deselected`) |
| M218 HEAD (`f849b5f`) | **FAILS** — `First extra element 13: '      - WUNDERGRAPH_SSR_ENDPOINT=…'` |

**M218 broke it, in iter-03** — its own C-1 fix added `WUNDERGRAPH_SSR_ENDPOINT` to the next-web block and
left the exact-shape fence pinned to the old block. The suite had been red since iter-03 and nothing caught
it. **It is the identical bug, one file over, from the identical cause, as the `demo-stack`
`test_tag_guard_present_for_both_frontends` fence that the same pass found and fixed.** The pass caught the
neighbour and misfiled the twin.

**The rule.** *"Reproduced at baseline ⇒ pre-existing"* is a sound inference **only if the baseline predates
the milestone.** Choosing a mid-milestone commit as "baseline" converts a regression into a clean bill of
health — which is **D17 exactly**: a status artifact (the verdict *"pre-existing"*) that outlives the thing it
described, and is then read as evidence.

**Binding for future harden/close passes:** resolve "baseline" to the milestone's **start ref** (the release
branch's merge-base, or the previous milestone's rext tag) — **never** to a commit that merely predates the
current *pass*. State the ref explicitly in the ledger, so the claim is checkable.

**Meta.** This is the **sixth** instance of D17's class in one milestone, and the first to appear **inside the
pass convened to name it**. That is not irony; it is the evidence that the class is real and that naming a
hazard does not inoculate you against it. Only a probe does.

---

## D19 — Ship the `x/crypto` bump, and say what it does and does not prove — 2026-07-14

**Handler:** `/developer-kit:close-milestone`, Phase 7 (the rext roll).

**The bind.** `x/crypto@v0.52.0` (13 dependabot alerts, **all govulncheck-UNREACHABLE**) was deferred at the
v2.2 close, re-fated **"Fate-1 → M218's rext roll"** in the v2.3 roadmap, and restated in M217's retro. **It
still had not landed.** This close **is** the rext roll — so deferring again would make it a **repeat +
aged-out** deferral, which the deferral audit treats as a **RED blocker**.

Against that: bumping it **rebuilds the clerkenstein binaries**, i.e. it perturbs the artifact the 5-cycle
cold latency battery was graded on — the same objection that routed **F-11** to M219 (**D16**).

**The decision: land it — and be precise about the difference.** F-11 is a **behavioural** change to the demo's
runtime identity path. This is an **indirect, transitive crypto library** with no code path on the login flow.
The claim is not *"it's fine, trust us"*; it is **measured**:

- The **56-gene alignment gate scores IDENTICALLY pre- and post-bump** — `97.2% overall / 100% critical`,
  `rc=0`, same single divergence. The gate includes every `critical` handshake/identity gene. **That is the
  instrument that would detect a behavioural change in the mirror, and it did not move.**
- Full suites: Python **887/0**, Go **0 failures across 6 modules**, flake gate **5/5**.

**What this does NOT prove, stated plainly:** the **5-cycle cold latency battery was not re-run** after the
bump. A crypto transitive has no latency surface, and re-running the battery costs a full remote 5-cycle
campaign — but the honest statement is *"behaviour-neutral by the alignment gate and the suites"*, **not**
*"re-graded against the gate"*.

**Why that residual is acceptable:** **M221 "prove it on billion" re-runs the p95 gate on the VM, with no
flags, by design.** The perturbed binary gets re-graded there as a matter of course. Leaving 13 open security
alerts to avoid a rebuild that the next milestone re-measures anyway is the worse trade.

---

## RELEASE-SCOPE-DEFER (v2.3 close-release Phase 1b — 2026-07-15)

_Recorded at `/developer-kit:close-release`. **User signed off (2026-07-15): accept → v2.4.** This is the
originating milestone for one of the four v2.3 tail carries; landing spot is `roadmap-vision.md` under v2.4. The
full four-item disposition is in `m221-prove-on-billion/audit-deferrals/deferral-audit-2026-07-15-m221-close.md`._

**RELEASE-SCOPE-DEFER: PROBE-M218-c3-rerun — Cosmo federation cms/Directus 403 re-check (DEF-M221-08).**
Originating in M218 (the C-3 router-403 re-check on the content path), inherited through to M221, not reached.
(Naming note: M218's close briefly labelled it `PROBE-M221-c3-rerun`; the majority/canonical form used across
M221 + `state.md` is **`PROBE-M218-c3-rerun`** — same item.)
- **Fate-1 (land now) FAILS.** A router-403 re-check needs the **live box** (a running federated stack with the
  cms/Directus content path exercised) — it cannot be discharged from the repo at close.
- **Fate-2 (drop) FAILS.** A real verification still owed on the content path; worth keeping tracked.
- **Fate-3 (defer) is correct.** Non-gate: the v2.3 headline gate is on click→ACCESS, independent of this
  content-path probe. → **v2.4** (re-run the router-403 check against a live stack).
