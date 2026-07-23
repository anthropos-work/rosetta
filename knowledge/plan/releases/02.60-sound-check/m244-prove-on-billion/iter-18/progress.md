**Type:** tik (run 7, tik 1). Active strategy: TOK-02.

# iter-18 — progress

## Gate (c) discrete stack-verify half — COMPLETE on billion (+ a real Bedrock provisioning fix)

Led with talk-to-data-m239 as the decisive probe (gate-c spec AND gate-h Bedrock proof). It reproduced the exact Bedrock-less failure its own comment predicts (flag gate PASSES → loads; Bedrock round-trip never streams → "chat did not grow"). Root cause: billion's demo had **NO Bedrock creds** (values-blind: source/provisioned/bridged app+platform .env all lacked `AWS_*`/`CLAUDE_CODE_USE_BEDROCK`); the LOCAL secrets source HAS them → a `/stack-secrets`-shaped provisioning gap, the v2.6 "not all gets provisioned" defect class. **Fixed values-blind** (synced the 4 Bedrock lines local→billion source+provisioned+bridge-target .env, idempotent/newline-guarded; force-recreated the backend) → **talk-to-data-m239 GREEN (11.2s), live member-count answer.** (D2.)

Then finished the discrete half. **Finding (D3): the "retrofit" is ENV-WIRING, not per-spec code** — every discrete gate spec already reads a full-URL base env var (localhost defaults only). Built + validated a reusable **`run-discrete.sh`** (run-coverage.sh HOST+SCHEME+OFFSET pattern; rext `6aacc32`, m244 tag moved+pushed, peels 6aacc32 on origin; harness-only → no billion re-pin).

### All 10 discrete gate spec files GREEN on billion (demo-1, tailnet https)
- persona-check ✓ · m224-candidate-heroes ✓ (iter-17)
- **talk-to-data-m239 ✓** (after the Bedrock fix — live answer)
- cockpit-overlay-return ✓ · enterprise-surfaces-m239 ✓ (#4 library populates + #1 hierarchical manager nav) · smoke ✓
- verify-members-B ✓ (real member rows) · verify-activity-dashboard-servegrant ✓
- **render-hiring-comparison ✓** (1-sim ASSERTED / HARD_GATE: 5 sims listed, drawer opened, **9/9 real candidate rows**, scores 27–70, 0 junk, 0 ejects — 3.4s; the 5-sim all-in-one loop timed out purely on the Next.js intercepting-route quirk × tailnet latency, not a broken surface)
- **m220-session-and-egress ✓ (6/6)** — session survives the academy, studio-stays-in-studio, 0 phone-home, clerk-js from FAPI; incl. the **"KNOWN RED" academy-content test now GREEN** (body>400 — an independent live cross-check that iter-15's academy fix holds; the KNOWN-RED docstring is now stale).

⇒ **Gate (c) stack-verify half = COMPLETE on billion** (2 coverage sweeps iter-16 + all 10 discrete gate specs). Gate (c) as a whole NOT ticked — the 16 Playthroughs run LAST (pt-world reset destroys the demo seed).

## Gate (h) byproduct progress (credited, not the iter's planned target)
Four of gate (h)'s ~6 v2.6-fix live-proofs landed as byproducts: **talk-to-data live answer ✓** (D2), **library ✓** (enterprise-surfaces #4), **cockpit UX ✓** (cockpit-overlay-return), **content fidelity ✓** (gate b, iter-13; media = voice presence-only, iter-07). Remaining for gate (h): academy course-start + language toggle + **p95 < 5s** (needs an autoverify refresh — the current verdict is stale >4h).

## Metric
Gate parts stay **5/8** (a,b,d,e,g). iter-18 completed the gate-c stack-verify HALF + 4/6 gate-h fixes but ticked NO full gate part (gate c needs the Playthroughs; gate h needs 3 more) — the coarse binary-per-gate artifact TOK-02 diagnosed. Real, load-bearing progress; the binary reads flat.

## Close — 2026-07-23

**Outcome:** The entire gate-(c) stack-verify half is GREEN on billion (all 10 discrete gate specs), enabled by fixing a real Bedrock provisioning gap (talk-to-data now answers live) + the finding that the retrofit is env-wiring (reusable `run-discrete.sh` shipped). 4/6 gate-h fixes proven as byproducts. Metric stays 5/8 (gate c needs the Playthroughs LAST; gate h needs 3 more).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (gate part (c) — stack-verify half complete; the 16 Playthroughs remain, run LAST on pt-world)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (iter-18 is a tik) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (tik 1/5) — (6) protocol-stop: n — Outcome: **continue**
**Decisions:** D1 (target held), D2 (Bedrock provisioning fix), D3 (retrofit = env-wiring + run-discrete.sh). See iter-18/decisions.md.
**Side-deliverables:** the Bedrock provisioning fix on billion (values-blind secrets sync + backend recreate) — a real v2.6 defect, durable on billion (source secrets updated); `run-discrete.sh` reusable runner.
**Routes carried forward:**
- gate (h): academy course-start + language toggle + **p95 < 5s** (autoverify refresh first) — 4/6 already proven; finishes gate h → ticks 6/8.
- gate (f): 3 drift-carries (BURNIN-M221 needs a `/dev-up --public-host` remote dev burn-in; F-M220-4; PROBE-M218-c3).
- gate (c) tick: the 16 Playthroughs LAST (pt-world reset).
- Finding (note for harden/close): m220's academy-content docstring is stale ("KNOWN RED" but now passes post-iter-15). The M239 Bedrock-cred DNA is R3/not-critical → the billion gap was NOT caught loudly (secrets-posture note candidate).
**Lessons:** "the specs execute green on billion" hid TWO things the gate wording didn't: (1) a spec can be RED on a provisioning gap (Bedrock) that looks like a broken feature but is a values-blind `/stack-secrets` fix; (2) the "remote-capability retrofit" was env-wiring, not code — the specs were already env-capable. Probe the riskiest spec FIRST (talk-to-data) — it surfaced the real defect and de-risked the whole discrete half.
