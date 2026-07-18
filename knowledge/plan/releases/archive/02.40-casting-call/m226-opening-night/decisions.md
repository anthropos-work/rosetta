# M226 — Decisions

_Implementation + strategy choices with rationale (incl. tok strategy revisions)._

| # | Decision | Rationale | Date |
|---|----------|-----------|------|
| — | _(intra-iter decisions live in `iter-NN/decisions.md`)_ | | |

---

## TOK-01: reprove-hiring-on-billion — 2026-07-17

**Tok type:** bootstrap (iter-01)

**Initial strategy:** Prove the 7-condition hiring gate on billion by **cutting billion's stale v2.3 substrate
over to the casting-call code-of-record, then running a bare default `/demo-up` and measuring from this Mac.**
Concretely, the first batch of tiks executes this order:

1. **Cold teardown of the stale demo.** billion currently runs a v2.3 `panorama` demo-1 (rext `cue-to-cue-v2.3.2`,
   up 2 days) with a live `tailscale serve` config on the base ports. Tear it down cleanly via the consumed
   tooling's teardown path (`/demo-down 1` / `up-injected.sh` teardown) — including the `tailscale serve` reset +
   the ant-academy respawner reap (the M221 F5/F5b/F12 discipline). Verify from this Mac that the base ports are
   freed and no survivor process remains (the M217 orphan-cockpit lesson).
2. **Substrate cutover.** Update billion's `~/panorama/stack-demo/rosetta-extensions` from `cue-to-cue-v2.3.2` →
   **`casting-call-m225-harden`** (the code-of-record; first confirm `sections`↔`harden` is a test-only diff so the
   runtime is unchanged), and refresh the billion platform clone to the casting-call refs (`ensure-clones.sh`
   bootstrap, which also pulls in the `apps/hiring` 2nd-app build path). This is what makes the hiring org + the
   4 hiring demo-patches available.
3. **Default cold bring-up.** Run a **default `up-injected.sh 1` (NO FLAGS)** on billion — remote reach is
   default-on (M220 D-DESIGN-3), so it auto-discovers `billion.taildc510.ts.net` + serves over the trusted cert.
   Run it **synchronously** in a blocking ssh and WAIT (a full casting-call rebuild with the 2-app hiring image is
   ~15–25 min; NEVER detach — a leaked build strands the demo-patches, the M221/M217 lesson). Heartbeat by
   file-write/process activity, not transcript idleness.
4. **Measure the 7-condition gate FROM THIS MAC** (the tailnet peer) — never on-host (on-host probing yields
   SSL-artifact false-REDs). Conditions: (1) hiring org present + `is_hiring=true` + exactly 5 mgr/45 cand;
   (2) recruiter comparison ≥40 non-junk rows per each of the 5 positions; (3) 2 candidate profiles usable;
   (4) reads as hiring; (5) **p95 click→ACCESS < 5 s** recruiter vantage (via `stack-verify/e2e/run-latency.sh`,
   gated on a fresh-green `autoverify.json`); (6) coexists with the 3 workforce orgs on the cockpit; (7) 0
   platform-repo edits (audit the demo clone is left git-clean).
5. **Attribute before fixing.** For any failing condition, name the surface — a re-surfaced render gate (R1), the
   45×5 whole-org-hydration latency (R4), a Clerkenstein/seed-wiring gap, or a memory/OOM stall — using the
   protocol's diagnostic discipline (read the arithmetic signature per `latency-budget.md`; the environment is
   part of every number). Fix via **tooling / a sha-pinned demo-patch re-proven live at final code**, tagged
   `casting-call-m226-*`. **A platform-repo edit is NEVER in bounds** — an un-patchable surface ESCALATES.

**Rationale:** This is the M215/M221 "prove-on-<VM>" pattern applied to the hiring release. Everything M226 proves
was already built + proven LOCAL (M222–M225); the only new information is what breaks on the live cross-machine
cold run — historically the last breakages are cross-machine (render origins, SSR-inlined URLs, mock coherence,
memory). The strategy front-loads the substrate cutover (billion is a full release behind) because that is the
real cost of the first tik, and orders measurement before fixing so gaps are attributed, not guessed. The billion
memory floor (7.3 GiB vs a documented 12 GiB, now carrying a 2nd UI container) is the sharpest live-only risk and
is watched from tik-01.

**Strategy class:** new-direction (bootstrap — no prior strategy to compare against).

**Distance-to-gate context:** Gate metric = the 7-condition live billion proof, reproducible on a cold
reset-to-seed. **Starting value: 0/7 conditions proven on billion at casting-call code** — the last billion proof
(M221) was the v2.3 panorama substrate, which has no hiring org and predates all of M222–M225. The hiring demo has
5 workforce-style conditions plus the recruiter-vantage p95 (a 3rd measured latency path, new to this milestone).

**Next-tik direction:** iter-02 (tik) executes steps 1–4 above: cold-teardown the stale v2.3 demo-1, cut billion's
rext (+ platform) over to `casting-call-m225-harden`, run a default cold `up-injected.sh 1` synchronously, then
take the FIRST 7-condition measurement from this Mac. Attribute every failing condition to its surface before any
fix. Handler: `PROVE-M226-iter02-first-cold-bringup`.
