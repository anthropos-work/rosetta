# M258 — decisions

_Release-level binding decisions **D-v28-1 … D-v28-15** live in
[`../../../roadmap.md`](../../../roadmap.md) § Active — v2.8._

---

## TOK-01: measure the composition before engineering it — 2026-08-12

**Tok type:** bootstrap (iter-01)

**Initial strategy.** M258 composes two halves. **One is measured and one has never been measured at
all.** The strategy is therefore ordered by what is *knowable*, not by what is *buildable*:

> **Take the composition's second half as a measurement before treating the milestone as a wiring job —
> and publish its spread beside its p50 from the very first sample, because the only evidence that
> exists says this suite's timing is not decidable at n=3 on this host.**

Four tiks, strictly ordered. Each is a precondition of the next, and the order is the strategy:

1. **Re-pin, then measure both halves in one cold campaign.** Discharge `R0` (the pin names
   `fast-build-m257-iter-09`; `fast-build-m257-close` is what is on origin, and it carries the M257
   close's three fail-open repairs **to the gate instrument this milestone reads its own number from**).
   Then one cold `demo-down --purge` + `demo-up` on the **free `demo-1` slot**, followed by the full
   Playthrough batch. That single campaign yields the bring-up half **at the corrected pin** and the
   **first wall-clock the batch half has ever had**. Report `load1` and the environment with both.
2. **Wire the batch-gate at the tail hook** (`up-injected.sh:2810`) under `D-v28-3` semantics: runs to
   completion, never halts at first red, never retries, **one consolidated red set at batch end**, the
   stack left **UP** regardless, the bring-up exiting **non-zero and loudly** on a non-empty set.
3. **Land the world-contract restore leg** (resolution (b), decided below).
4. **The composed 3× cold campaign** against the gate, with the spread published beside the p50.

**Rationale.** The milestone's own shape note argues this is *"M257 plus one invocation line"* and
*"should close in 1–2 iters"* — and that reading is only safe if the composition arithmetic is known.
It is not. `overview.md` § *Budget honesty* asked M256 to measure the reset-to-seed leg; **it did not**,
and the Phase-0b audit found the gap had reached neither this plan nor the corpus. Wiring first and
measuring last would put the milestone's one genuine unknown at the **end** of the milestone, which is
the failure shape M257 spent three iters inside (a gate whose subject could not be measured, closing
"delta 0" three times while nothing was learnable). Measuring first costs one campaign and converts
"1–2 iters" from a hope into a prediction with a number behind it.

**The world contract — RESOLVED: (b) restore after.** `overview.md` names two admissible resolutions and
requires the choice at iter-01. **(a) pt-world-native is refuted by the gate's own text**, not by
preference: the gate requires the stack be left *"in a presenter-usable world"*, and the overview's own
description of (a) ends *"But it is not a presenter demo."* A resolution that cannot satisfy a gate
clause is not admissible for a milestone gated on that clause. **(b)** is chosen, and it is cheap for
reasons the overview already measured: `--reset` does not touch the snapshot-replayed taxonomy (no
catalog tables in `resetTables`, so the 78.0 s replay is not repaid), the stories seed measures 7.6 s,
and the manifests need no re-export (`--cockpit-export` takes no DB, `stackseed/main.go:172`; ids are
deterministic). **Restore leg ≈ 20–45 s**, to be measured in tik 1 rather than assumed.

**Host topology — the working assumption, disclosed.** The gate is taken against the **single-box
`--no-public-host`** mode, because that is the only mode in which *"one cold command"* is literally
satisfiable, and because `D-v28-15` puts dev/test **local to `macmini`** while `billion` is
demo-deployment-only. The overview's own note stands and is not being papered over: this proves the
composition **in a mode the presenter never uses**. The peer path is already available and unbroken
(`run-playthroughs.sh --reset-only` splits the DB half from the browser half, `:58-62`), so gating the
presenter mode instead is a re-cut the user can take later at low cost. **Disclosed, not decided
silently** — this is the one strategic assumption in `TOK-01` that a user might overturn.

**Strategy class:** `new-direction` (bootstrap — no prior strategy to compare against).

**Distance-to-gate context.** Gate: **p50 ≤ 480 s** over 3 consecutive cold reset-to-seed cycles, zero
standing red, 0 platform edits, presenter-usable world.

| half | value | provenance |
|---|---|---|
| bring-up | **286.99 s** (n=3, min 280.99 / max 303.44) | M257 iter-09, `macmini`, `gateable: true` |
| batch | **UNMEASURED** | — |
| headroom on the ceiling | **193.01 s** for batch + restore | arithmetic |

The 480 s premise **held** — M257 landed 286.99 s, inside the 240–300 s window the ceiling-sum was
conditioned on — so the composition is not re-litigated. What is unknown is whether the batch half fits
in 193 s. The only wall-clock evidence anywhere is **56.6 s over 18 specs** (M256 `progress.md`
iter-02); the shipped suite is **209 passed across 30 live Playthroughs**, so that figure prices a
suite an order of magnitude smaller and **must not be quoted as the batch half**.

**⚠️ The decidability caveat, carried from the first tik and not deferred.** M256 iter-12 escalated that
its suite timing is **not decidable at n=3 on this host** — six full-suite runs, unchanged specs, a
control subset spanning **0.5281× → 1.0762× (a 2.04× spread) with no trend** — and `D-v28-12` was
re-cut for exactly that reason. M258's gate is a **p50 over n=3** whose second half is that suite.
Every campaign under `TOK-01` therefore **publishes the spread beside the p50**. If the spread makes
n=3 undecidable, that is a user renegotiation **with measurements attached**, routed through the
milestone's declared `re_scope_trigger` — never an un-measured escalation.

**Known-context carried into every tik under this strategy** (from the Phase-0b audit,
`../kb-fidelity-audit.md`; these are context, not deferrals):

1. **`R0`** — the rext pin is one tag behind origin. Tik 1 opens by re-pinning.
2. **`C1`** — the batch half has no published wall-clock. Tik 1 produces the first.
3. **`C2`** — the 2.04× decidability caveat above.
4. **`F1`** — `FIX-M257-content-stories-pair-count` is **genuinely open and correctly described**
   (probed: `content-pairs.ts:115` has the `manager_presence_only` branch, the shell re-implementation
   at `run-content-stories.sh:145-152` does not, so it counts 47 against a pinned 45 and `exit 2`s).
   It gates the **content-stories sweep**, which is *not* a Playthrough — so it does **not** block the
   batch, and it must not be allowed to look like a batch blocker.
5. **`F2`** — `ptvalidate` is invoked nowhere outside its own tests; wiring it as a binding pre-flight
   discharges the M256-inherited permissive-gate hole. It runs static-only in seconds.
6. **The inherited lists are SUSPECT-UNROUTED until verified open.** M257's close proved the routing
   itself fails (M257x's carry-forward reached this `overview.md` zero times), and this audit found 13
   stale anchors in the same file. **Re-verify every inherited item against code before working it** —
   `F1` was re-verified this way and survived; a lesser check would have mis-graded it as already-fixed.

**Next-tik direction (iter-02, a tik).** Re-pin `.agentspace/rext.tag` → `fast-build-m257-close` and
verify the consumption clone follows. Then one cold campaign on the **free `demo-1` slot**
(`demo-2` and the 5-container dev stack are the **user's** — do not tear down, re-seed, restart or
reset either): `demo-down --purge` + `demo-up --no-public-host`, then the full Playthrough batch.
Deliverable: **the first measured batch half**, its **spread**, the restore-leg cost, and a composed
figure against 480 s — reported with `load1` and the environment, never as a clean baseline on a
host that is permanently contended.
