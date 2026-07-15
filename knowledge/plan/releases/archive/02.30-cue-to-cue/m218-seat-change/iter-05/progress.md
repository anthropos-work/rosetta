# iter-05 — Progress

## THE GATE IS MET — over 5 *actual* cold reset-to-seed runs

_`billion` · demo-1 · rext **`cue-to-cue-m218-iter05`** (`3d6227a`) · measured **from the tailnet** (the
presenter's real vantage) · **5 samples/vantage/cycle, gate ARMED** (`LATENCY_GATE_MS=5000`) · harness reused
unchanged (rext `stack-verify/e2e/`, iter-02 — **not forked**) · **0 code changes across the battery**._

Every cycle: `/demo-down 1 --purge` → `/demo-up 1 --public-host …` (**all images rebuilt from zero**) → **proven
cold** → **proven green** → 5 cold logins/vantage (cookies cleared per sample), both heroes.

| cycle | proven cold? | autoverify | employee p50 | **employee p95** | ACCESS | manager p50 | **manager p95** | ACCESS | verdict |
|---|---|---|---|---|---|---|---|---|---|
| **1** | ✅ initdb re-ran | `green`, 0 warn | 1035 ms | **2413 ms** | 5/5 | 1164 ms | **1417 ms** | 5/5 | ✅ PASS |
| **2** | ✅ initdb re-ran | `green`, 0 warn | 1037 ms | **2246 ms** | 5/5 | 1147 ms | **1767 ms** | 5/5 | ✅ PASS |
| **3** | ✅ initdb re-ran | `green`, 0 warn | 1008 ms | **2239 ms** | 5/5 | 1137 ms | **1320 ms** | 5/5 | ✅ PASS |
| **4** | ✅ initdb re-ran | `green`, 0 warn | 1041 ms | **2311 ms** | 5/5 | 1140 ms | **1375 ms** | 5/5 | ✅ PASS |
| **5** | ✅ initdb re-ran | `green`, 0 warn | 1015 ms | **2243 ms** | 5/5 | 1112 ms | **1371 ms** | 5/5 | ✅ PASS |

**Worst p95 across all 10 vantage-runs: 2413 ms** (employee, cycle 1) vs the **< 5000 ms** gate — a **2.07×**
margin. **50/50 logins reached ACCESS.** Every armed gate assertion passed (`1 passed` per Playwright run).

> **exit_gate:** *"p95 click→ACCESS < 5 s … for BOTH `maya-thriving` (employee → `/profile`) AND `dan-manager`
> (manager → `/enterprise/…`), measured over **5 consecutive cold reset-to-seed runs**."* → **MET, as written.**

Raw evidence: `battery-results.json` (per-cycle p50/p95/ACCESS/autoverify/cold-witness).

## What the honest cold number cost

iter-04 reported employee p95 **1456 ms**. The five *genuinely* cold runs sit at **2239–2413 ms** — **~1.6×
slower, consistently.** That gap is exactly the warm database iter-04 was unknowingly measuring on (D12/D13): a
cold cycle pays for a fresh `initdb` → migrate → snapshot replay → seed, cold page cache, and a freshly-built
image. iter-04's number was *optimistic*, not wrong-in-kind — the gate is met either way, but only the cold
number is the one the gate asked for.

**Manager stays faster than employee** (1320–1767 ms vs 2239–2413 ms), consistent with iter-03/iter-04.

## Findings (this iter added no performance work — it added three findings)

| # | Finding | Status |
|---|---|---|
| **F-9** | **`/demo-down --purge` never purged, on any Linux host.** `rm -rf "$stack/data"` cannot unlink postgres's UID-1001/0700 cluster dir → `Permission denied` → `set -euo pipefail` **aborted `cmd_down`** → the DB was never wiped, the images were never removed, the registry slot leaked, and it returned a bare `rc=1` with no diagnosis. **`billion`'s postgres cluster still carried the `PG_VERSION` `initdb` wrote on 2026-07-11**, through every "cold" bring-up — **including iter-04's**. No demo on that host had been reset-to-seed in days. | **FIXED** (rext, no platform edit) + 5 regression tests, all green **on the Linux host**. Proven in the field: first `--purge` after the fix → `data dir GONE`, images `0`, containers `0`. |
| **F-10** | **A torn-down stack can hand you a stale `green`.** `autoverify.json` is not removed on teardown. Caught live: a **failed** bring-up (0 containers) still presented `{"green":true,"warnings":0}` — read by both the battery's gate and `run-latency.sh`'s own green gate. Same class as **F-6**. | **Worked around** in the battery (delete before each bring-up). **Durable fix routed Fate-1** → `decisions.md` F1-2. |
| **D15** | **The alignment blind spot.** `/align-run` scored Clerkenstein **100%/100%, 0 divergences — before *and* after** iter-04, while the fake BAPI returned **the wrong human for every request**. Cause verified: `getUser` has **no gene in any of the 5 DNAs**; the three genes that *do* name identity all assert the variant **`universal-user`** — *the stub itself*. The goldens **ratified the defect**. The corpus's own named mitigation (`/align-dna`'s "every consumed endpoint is present") **did not bind**. | **Routed Fate-1** with a named handler + red-before/green-after DoD → `decisions.md` **D15** + milestone `decisions.md` **F1-1/F1-3**. |

## Escalations

**None.** No fix needed a platform-repo edit. The hard constraint held: all changes in `rosetta-extensions`
(demo tooling) + `rosetta` (knowledge). **Zero platform-repo edits.**

## Next

The gate is met **as written**. Next: **`/developer-kit:harden-mstone-iters --final`** — which **owes the three
Fate-1 items** in the milestone `decisions.md` (F1-1 the `GetUser` identity gene, F1-2 teardown unlinks
`autoverify.json`, F1-3 the coverage-check honesty) — then `/developer-kit:close-milestone`.
