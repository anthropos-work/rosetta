# M218 — Decisions

_Implementation + strategy choices with rationale (incl. tok strategy revisions)._

| # | Decision | Rationale | Date |
|---|----------|-----------|------|
| — | _(intra-iter decisions live in `iter-NN/decisions.md`)_ | | |

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
