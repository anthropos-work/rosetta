# M256 · iter-02 — decisions

## D6 — `pt-studio-advanced-generate` is a FALSE GREEN, and the gate's "irreducible LLM lane" has no referent

**Symptom.** The Playthrough the corpus documents as *"a real ~2–3 min live-LLM round-trip"* — the one
D-v28-9 built a whole gate carve-out around — completed in **1.52 / 1.19 / 1.26 s**.

**Evidence it never waited for the generation.** Timestamps from `docker logs -t demo-2-studio-desk-1`
against run 1 (suite start `11:17:45Z`, process exit `11:19:19Z`):

```
11:18:59.098  [AIService] Using Azure OpenAI … (tier: instant)
11:18:59.120  [AIService] Using Azure OpenAI … (tier: thinking_fast)
11:18:59.821  [AIService] Azure OpenAI failed … → Using OpenAI (fallback)  (instant)
11:19:00.044  [AIService] Azure OpenAI failed … → Using OpenAI (fallback)  (thinking_fast)
11:19:02.211  [AIService] Completed in 2.381s
11:19:19.007  [AIService] Completed in 18.962s      ← after the whole suite had finished
```

The generation was **still in flight** when the test passed, and its `thinking_fast` leg landed **19 s
later**, at the very moment the suite process exited.

**Root cause, in one line.** `playthroughs/e2e/lib/studio-builder-page.ts` §`advancedDesignerRendered`:

```ts
return this.byText(/Simulation Advanced Builder|Scenario Characters|Mission Tasks/i).first();
```

The first alternative — **`Simulation Advanced Builder`** — is the **route's own page header**. It renders the
instant `/sim-advanced-builder` opens, before any draft exists. With `.first()`, the matcher resolves to the
header, so the "completion boundary" assertion is really a **route-arrival** assertion. The `currentUrl()`
check that follows asserts the same arrival a second time. The two genuinely post-draft alternatives
(`Scenario Characters`, `Mission Tasks`) are never the ones that fire.

**Why this matters beyond one spec — three claims are wrong as a result:**
1. `corpus/ops/demo/playthroughs.md` § the `studio` product: *"the advanced builder GENERATE runs to its
   completion boundary — the generated result renders"*. It does not; the route opens.
2. **D-v28-9's premise**: *"studio-advanced is an irreducible ~2–3 min live-LLM round-trip — plausibly
   ~120 s of the 228 s baseline on its own"*. On this host it is **1.26 s**, so the 228 s figure it was
   inferred from cannot be explained this way.
3. **The exit gate's own carve-out** — *"with the irreducibly LLM-bound studio lane excluded from the median
   and budgeted separately"* — currently excludes the **two fastest tests in the suite**. That makes clause 1
   **harder**, not easier (3.326 s excluded vs 3.067 s included), so the gate stays **conservative** and no
   re-cut is required. The carve-out becomes meaningful only once the spec really waits.

**This is a clause-2 deliverable, not a bug report.** Clause 2 requires *every Playthrough passes a negative
control (demonstrably RED when its outcome is absent)*. `pt-studio-advanced-generate` is the milestone's
first proof that the clause was worth writing: a negative control would have caught this on the day it
landed. **Fate 3 → the clause-2 tik**, as `FIX-M256-studio-false-green`, with the doc correction
(`DOC-M256-llm-lane-premise`) written **once**, against the fixed behaviour.

**Not fixed in iter-02.** The iter's planned scope is a baseline, and a baseline measured on changed code is
not a baseline. The scope-creep tripwire applies.

## D7 — The measurement protocol is PINNED, because run-to-run spread ≈ the gate's own target

Per-run medians over the 16 non-studio Playthroughs: **3.935 / 2.989 / 3.228 s**. The coldest-to-warmest
spread is **18 %** against clause 1's **21 %** target — so *which run you quote* very nearly decides the
gate. `pt-aireadiness-manager-dashboard` alone swung **17.19 s → 1.86 s → 1.62 s** (9×) as next-web's route
warmed.

**Pinned, and to be recomputed identically at the end:**
- **Statistic:** median across the 16 non-studio Playthroughs of each Playthrough's **median across 3
  consecutive `--reset` runs**. (Cross-check, median of per-run medians: 3.228 s — 3 % lower.)
- **n = 3**, consecutive, `--reset` before each, `workers: 1`, `retries: 0`.
- **Run 1 is the first run after bring-up and is INCLUDED.** Discarding it would look like rigour and
  behave like a thumb on the scale: the warm-up is real work a real run pays, and the end-of-milestone
  measurement will pay it too. Symmetry is what makes a *relative* gate valid.
- **Environment stated with every number**, and no comparison to billion (D-v28-12).

## D8 — Re-target iter-03: the median's driver is the per-test LOGIN, not the `networkidle` inheritance

TOK-01 move 2 named the residual `networkidle` as the clause-1 lever. The baseline says that is the right
**correctness** fix but the wrong **timing** target on this host, and points at a better one.

**Why `networkidle` is cheap here.** `page-object.ts` records that `networkidle` "passed on localhost (fast,
sparse requests settle) and DEADLOCKED over the tailnet", and M244 iter-23 proved the deadlock only over the
tailnet. The measurement agrees: `pt-aireadiness-manager-dashboard` — a polling surface, a networkidle
inheritor, and the class that deadlocked on billion — costs **1.86 s** warm here. There is no local
seconds-scale `networkidle` tax to reclaim.

**Why the login is the driver.** `pt-profile-identity` is the suite's minimal journey — login → `/profile` →
assert one name — and it costs **3.50 s**, which is **above** the 3.326 s median. Every Playthrough pays that
same handshake (`selectSeat` → `POST /v1/demo/select` → a protected-route `goto` → clerkMiddleware 307 → FAPI
handshake → 303 → cookie → render). So the login is very nearly the **whole** cost of the median test, and it
is the one cost shared by all 18 — exactly the shape that moves a **median** rather than a tail.

**iter-03 therefore leads with `storageState` reuse (TOK-01's L3), not L1.** The constraint from iter-01 D1
stands and shapes the design: Clerkenstein holds **one global seat** and `handleMe` reads it with no cookie
input, so a reused `storageState` does not re-point the server-side seat. Reuse must therefore be
**seat-grouped and serial** — group the 18 tests by their 6 seats, `selectSeat` once per group, and reuse
that group's storage state within it. `pt-profile-identity` is retained as the one test that still performs
the full handshake, so the handshake itself stays proven.

L1 + L2 + the widened fence still land in iter-03 — they are correct, cheap, and they protect the tailnet
path M258 will drive — but they are landed as **correctness with a measured (likely small) local delta**,
not as the clause-1 lever. Claiming otherwise would be exactly the un-probed-lift dishonesty the protocol's
self-check forbids.
