# M218 — Decisions

_Implementation + strategy choices with rationale (incl. tok strategy revisions)._

| # | Decision | Rationale | Date |
|---|----------|-----------|------|
| — | _(intra-iter decisions live in `iter-NN/decisions.md`)_ | | |

---

## OPEN — Fate-1 items owed by this milestone, with a named handler

_Milestone-scope, because the handler is a milestone-scope actor. **Do not close M218 with these open.**_

**Handler: `/developer-kit:harden-mstone-iters --final`** (runs after the gate fires, before
`/developer-kit:close-milestone`).

| # | Item | Why it is Fate 1 (land it completely, here) | Definition of done |
|---|---|---|---|
| **F1-1** | **The alignment blind spot — a `GetUser` gene with per-hero identity on the BAPI surface (`clerk-2.6.0`).** `/align-run` scored Clerkenstein **100% critical / 100% overall, 0 divergences — before *and* after** iter-04, while the fake BAPI was returning **the wrong human for every request**. Verified cause: `getUser` has **no gene in any of the 5 DNAs**, and the three genes that *do* name identity (`ExtractIdentity`, `Me`, `DeployIdentity`) all assert the variant **`universal-user`** — *the stub itself*. The goldens **ratified the defect**. `DistinctIdentity` (the only per-hero gene) exercises the **FAPI/JWT** path — the half that was already correct. | It is a **regression golden for a bug this milestone just fixed**, and a mirror that serves a stub identity to every hero scoring 100% makes the score untrustworthy for every future mirror. Small; exactly what a harden pass is for. **Deliberately not done inline in iter-05:** a code change **restarts the 5-cycle count** (iter-05 D13), so it is sequenced *after* the gate fires. | The gene exists, is `critical`, and is **red against `8ebc89e^` / green after**. If it passes on both, it is not fencing the bug. |
| **F1-2** | **Teardown must remove `autoverify.json`** (F-10). It is *not* removed on teardown, so a torn-down stack leaves a `green:true` verdict on disk for a stack that no longer exists — read by both `run-latency.sh`'s green gate and any grader. Caught live in iter-05: a **failed** bring-up (0 containers) still presented `{"green":true,"warnings":0}`. Same class as **F-6**. | A grader that can be handed a stale `green` cannot be trusted to refuse a broken stack — which is the entire premise of the M217→M218 hard barrier. | `cmd_down` unlinks it; a test fences it. |
| **F1-3** | **The capability-coverage check is not binding.** `alignment_testing.md:169–172` offers, as the *named mitigation* for a hollow score, "`/align-dna`'s capability-coverage check (**every consumed endpoint is present**)". `GET /v1/users/{id}` **is** consumed (next-web's server-side `currentUser()`, every authenticated render) and has no capability — so the safeguard did not bind. | The doc currently **over-claims a guarantee the tooling does not provide**. Either make it bind, or say so. | Either the check covers the fake BAPI's read surface, or `alignment_testing.md` states honestly that it does not. Joins `DOC-M218-audit-corrections`. |

_Full analysis + evidence: `iter-05/decisions.md` **D15**._

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
