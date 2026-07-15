---
iter: 5
milestone: M218
iteration_type: tik
status: closed-fixed
opened: 2026-07-13
closed: 2026-07-13
---

# iter-05 (tik) — finish the battery: prove the gate over 5 *actual* cold reset-to-seed runs

**Active strategy:** **TOK-01** — *"reachability-first"*. The bottlenecks are gone (iter-03, iter-04). This iter
adds no performance work. It exists to **discharge the gate as written**, and to check the one thing iter-04
asserted but did not measure.

## Why this iter exists

iter-04 declared **GATE MET** on p95 **1.46 s** (employee) / **1.40 s** (manager) vs a **< 5 s** gate. The
numbers were real. The **grading was not complete**. The milestone's `exit_gate` says, verbatim:

> p95 click→ACCESS < 5 s … for BOTH `maya-thriving` (employee → `/profile`) AND `dan-manager` (manager →
> `/enterprise/…`), **measured over 5 consecutive cold reset-to-seed runs.**

iter-04 measured **5 cold logins per vantage on ONE stack** — i.e. **1 of the 5 required cycles**. It disclosed
this honestly. This iter runs the rest.

**This is not pedantry.** **F-6** (iter-03) proved bring-ups are non-deterministic: a rebuild silently baked the
**real Clerk publishable key**, so the demo phoned **production auth** with login broken — while its
`autoverify.json` still reported `green`. The 5-cycle rule exists to catch exactly that. One cycle cannot.

## What the battery found before it could even start

**The gate was not just ungraded — it was, on this host, *ungradeable*.**

- **F-9 — `--purge` never purged.** `rosetta-demo down N --purge` was a bare `rm -rf "$stack/data"`. postgres's
  cluster dir inside the bind mount is owned by **UID 1001, mode 0700**, so on a **Linux** host `rm -rf` dies
  `Permission denied` — and under `set -euo pipefail` that **aborted `cmd_down`**, skipping the image removal and
  the registry release too. `--purge` returned a bare `rc=1` having purged **nothing**.
- **Consequence:** `billion`'s postgres cluster still carried the `PG_VERSION` `initdb` wrote on **2026-07-11**,
  through every subsequent "cold" bring-up — **including iter-04's**. Fresh containers, two-day-old database, and
  the same cached images every time. **No demo on that host had been reset-to-seed in days.** iter-04's cycle was
  therefore **not** a cold reset-to-seed run, and is **not counted**.

Fixed in rext (no platform edit) — see **D13**. The fix is proven in the field: the first `--purge` after it
reported `data dir GONE (purged) OK`, images `0`, containers `0`.

Because that is a **code change, the 5-cycle count restarts at 0** (D13). All five cycles below run on the fixed,
re-pinned tooling, and each one is **proven cold** before it is allowed to be measured.

## The battery (what one cycle is)

Per cycle, in order — a cycle that fails any gate **stops the battery** (a finding, not a retry):

1. `rm -f autoverify.json` on the host — **F-10 / D14**: a torn-down stack must not hand us its predecessor's
   `green`. (Caught live: a *failed* bring-up left `{"green":true,"warnings":0}` on disk for a stack with **0
   containers**.)
2. `/demo-down 1 --purge` → `/demo-up 1 --public-host billion.taildc510.ts.net` (cold: initdb → migrate →
   snapshot replay → seed; **all images rebuilt**, so every cycle re-bakes the minted pk — the F-6 surface).
3. **GATE A — proven cold:** `PG_VERSION`'s mtime must be **newer than the cycle start** (i.e. `initdb` re-ran).
   *Assumed*-cold is what produced this whole finding; coldness is now **witnessed**.
4. **GATE B — proven green:** `autoverify.json` must be `{"green":true,"warnings":0}` — and must have been
   written by **this** bring-up.
5. **Measure:** 5 cold logins per vantage (cookies cleared per sample), **both** heroes, gate **armed**
   (`LATENCY_GATE_MS=5000`), from the **tailnet** — the presenter's real vantage. Harness reused unchanged
   (rext `stack-verify/e2e/`, iter-02); **not forked, not rewritten.**

## Result — **THE GATE IS MET, as written**

**5/5 cycles passed**, every one **proven cold** and **proven green**, both vantages, gate armed:

| cycle | **employee p95** | **manager p95** | ACCESS |
|---|---|---|---|
| 1 | **2413 ms** | **1417 ms** | 5/5 · 5/5 |
| 2 | **2246 ms** | **1767 ms** | 5/5 · 5/5 |
| 3 | **2239 ms** | **1320 ms** | 5/5 · 5/5 |
| 4 | **2311 ms** | **1375 ms** | 5/5 · 5/5 |
| 5 | **2243 ms** | **1371 ms** | 5/5 · 5/5 |

**Worst p95 across all 10 vantage-runs: 2413 ms** vs the **< 5000 ms** gate — a **2.07× margin**. **50/50 logins
reached ACCESS.**

**The honest cold number is ~1.6× slower than iter-04's** (employee 2239–2413 ms vs 1456 ms) — that difference
*is* the warm database iter-04 was unknowingly measuring on. The gate is met either way; only the cold number is
the one the gate asked for.

Full table + raw evidence: `progress.md`, `battery-results.json`.

## Escalation conditions

- A cycle that is not provably cold, or not provably green → **stop**, record, root-cause. Never re-roll for a
  luckier stack.
- Any cycle with p95 ≥ 5 s on either vantage → the gate is **NOT met**; report it and stop.
- Any fix that would need a platform-repo edit → escalate (none did).

## Also in this iter

**D15 — the alignment blind spot.** `/align-run` scored Clerkenstein **100% critical / 100% overall, 0
divergences** *both before and after* iter-04 — while the fake BAPI was returning **the wrong human for every
request**. Verified why: the BAPI's `getUser` has **no gene in any of the five DNAs**, and the three genes that
*do* mention identity (`ExtractIdentity`/`Me`/`DeployIdentity`) all assert the variant **`universal-user`** — *the
stub itself*. The goldens **ratified the defect**. Routed **Fate 1** to the final harden pass with a named
handler and a red-before/green-after definition of done. See `decisions.md`.
