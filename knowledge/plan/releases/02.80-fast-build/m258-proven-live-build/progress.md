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

- iter-02 (tik, `closed-fixed-partial`): **`R0` discharged** (rext re-pinned to `fast-build-m257-close`;
  the third pin copy proven **inert** by path arithmetic, not assumed). Bring-up half re-measured at the
  corrected pin: **395.31 s** (n=1, `load1 2.26`, contended + labelled) — `rc=0`, green, HEADROOM OK,
  identity MATCH, phases complete. **The batch half could not be measured, and the blocker is the
  finding: ISOLATION went RED on the first campaign** — all three UI images carry a **non-minted**
  publishable key. Both causes the assert names were **refuted** (fresh build, overlay present); the
  real mechanism is a third — `.env.demo-1` holds **24 appended Clerkenstein blocks** and this run's is
  the one carrying the foreign key, so **last-wins** wired the UI tier to a real Clerk app.
  `inject.py:89` appends instead of rewriting; `up-injected.sh:2036` runs it `2>/dev/null || true`.
  **Not caused by the re-pin** (its whole `up-injected.sh` diff is comments + `log` strings) and
  **`demo-2` is clean** (last key minted). — see `iter-02/progress.md`

## Next-iter routing

- **iter-03 (tik, under `TOK-01`)** — **`FIX-M258-iter02-inject-appends-and-swallows` first**: it gates
  the milestone's primary unknown. Two defects, one symptom — `inject.py:89` appends a Clerkenstein
  block per bring-up (24 in `.env.demo-1`, so the file's history outranks its intent) and
  `up-injected.sh:2036` swallows the mint's stderr and tolerates its failure. **Fix both** — repairing
  the append alone would mask the swallow. Then re-run the cold cycle and **take the batch half**, which
  is still the deliverable `TOK-01` step 1 owes.
- Then, unchanged: wire the batch-gate at `up-injected.sh:2810` under `D-v28-3` semantics → land the
  world-contract restore leg (b) → the composed 3× cold campaign, **spread published beside the p50**.
- ⚠️ `demo-2` (11 containers) and the 5-container dev stack are the **user's**: do not tear down,
  re-seed, restart or reset either. **`demo-1` is left UP as iter-03's reproduction and must not be
  browsed** — its UI tier would talk to a real Clerk app (`iter-02/decisions.md` D7).

### Also routed from iter-02 (smaller, same tik or later)

- **`CHECK-M258-iter02-studio-desk-is-the-untouched-leg`** — `ui_studio_desk` **115.35 s** is the
  largest UI leg and the one L1 never touched (L1 multi-staged the two *Next* apps). Named suspect for
  the n=1 vs M257-n=3 delta and a candidate lever if the composed budget needs room. **Confirm against
  n≥3 before any claim.**
- **`ROUTE-M258-iter02-isolation-names-two-causes-not-three`** — the refusal text offers two
  explanations and both were refuted; a refusal naming the wrong cause sends the reader at the build.
- **`ROUTE-M258-iter02-headroom-defaults-to-billion`** — bare `assert-headroom` grades against
  `billion.json`; the host must be named every time (cluster 4's shape, second entry point).
- **`ROUTE-M258-iter02-purge-did-not-clear-the-stack-dir`** — 24 accumulated blocks means the stack dir
  survived repeated `--purge` cycles: `verification.md`'s **F-9** instance.

## Carried known-context

`TOK-01` § *Known-context* #1–#6 — `R0` (stale pin) · `C1` (batch half unmeasured) · `C2` (n=3
decidability, 2.04× spread) · `F1` (`FIX-M257-content-stories-pair-count`, verified open; gates the
**content-stories sweep**, not the batch) · `F2` (`ptvalidate` unwired) · the **SUSPECT-UNROUTED** rule
for every inherited M257x / M257 item. Not deferrals.
