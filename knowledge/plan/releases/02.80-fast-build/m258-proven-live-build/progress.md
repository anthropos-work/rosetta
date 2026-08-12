# M258 — progress

## Running ledger

- iter-01 (**tok**, bootstrap): Phase 0b gate **YELLOW** — 0 blind areas, but **13 stale line anchors in
  M258's own `overview.md`** (all in-range, all landing on unrelated content; substance held in every
  case) repaired, and **two never-propagated measurements** recorded at the destination: the batch half
  of the composed budget has **no published wall-clock** (M256 was asked and did not report it), and
  M256 **escalated that this suite's timing is not decidable at n=3 on this host** (2.04× spread, no
  trend) — against a gate that is a p50 over n=3. `TOK-01` authored: *measure the composition before
  engineering it*. **World contract RESOLVED → (b) restore after**, refuting (a) on the gate's own text.
  Rung zero found the rext pin **one tag behind origin** (`R0`). `F1` re-verified against code and
  **survived** — it read as already-fixed and is open. — see `iter-01/progress.md`

## Next-iter routing

- **iter-02 (tik, under `TOK-01`)** — re-pin `.agentspace/rext.tag` → `fast-build-m257-close`
  (discharges `R0`; the current pin predates the three fail-open repairs to the very instrument this
  gate reads), then **one cold campaign on the free `demo-1` slot**: `demo-down --purge` +
  `demo-up --no-public-host`, then the full Playthrough batch. Deliverable: **the first measured batch
  half**, its **spread**, the restore-leg cost, and a composed figure against 480 s — each reported with
  `load1` and the environment. ⚠️ `demo-2` (11 containers) and the 5-container dev stack are the
  **user's**: do not tear down, re-seed, restart or reset either.
- Then, in order: wire the batch-gate at `up-injected.sh:2810` under `D-v28-3` semantics → land the
  world-contract restore leg (b) → the composed 3× cold campaign.

## Carried known-context

`TOK-01` § *Known-context* #1–#6 — `R0` (stale pin) · `C1` (batch half unmeasured) · `C2` (n=3
decidability, 2.04× spread) · `F1` (`FIX-M257-content-stories-pair-count`, verified open; gates the
**content-stories sweep**, not the batch) · `F2` (`ptvalidate` unwired) · the **SUSPECT-UNROUTED** rule
for every inherited M257x / M257 item. Not deferrals.
