# iter-95 ground truth — re-derived at this open, 2026-08-05

**This sheet is authoritative for refs.** A seat that grades a claim against a ref not on this sheet
must say which ref it used and why. Delivered to all 14 seats as a marked addendum below a
**byte-identical** copy of the frozen briefing.

| clone | checkout | `origin/main` | behind |
|---|---|---|---|
| `platform` | `0c91421d` | `0c91421d` | 0 |
| `app` | `b948604f` | `2035f9a4` | **93** ⚠️ |
| `next-web-app` | `bb3313bc` | `8297c684` | **41** ⚠️ |
| `storage` | `4ce8ece5` | `9f8cb532` | **20** ⚠️ |
| `messenger` | `fa47850d` | `e9421c68` | **7** ⚠️ |
| `ant-academy` | `9c3843cd` | `22df69dd` | **5** ⚠️ |
| `jobsimulation` | `462343b0` | `82cb66ec` | **4** ⚠️ |
| `cms` | `ca50c817` | `f38c0c4a` | **2** ⚠️ |
| `sentinel` | `88bc5592` | `f2c46190` | **2** ⚠️ |
| `studio-desk` | `14a5442a` | `41ee3575` | **2** ⚠️ |
| `roadrunner` | `87d8d443` | `87d8d443` | 0 |
| `graphql-wundergraph` | `60c229f3` | `60c229f3` | 0 |
| `rosetta-extensions` (authoring, `.agentspace/`) | `6130bfd8` | `6130bfd8` | 0 |
| `rosetta` (the corpus under audit) | `1977224` | — | tree clean |

## What changed since iter-86's sheet, and what did not

**Only `platform` moved:** `0dab54df` → `0c91421d`. Every other **checkout** is byte-identical to
iter-86's sheet. The `origin/main` column has advanced for most clones (`app` 60 → **93** behind), so
the *staleness* is worse even where the checkout is unchanged — which is exactly why the column
exists and why it was given to the seats.

## The two rules that go with the sheet

1. **Grade at the ref the claim names — UNLESS the sentence asserts currency**, in which case no
   neighbouring pin rescues it. A pin is a **date**, not an excuse.
2. **A corpus citation into `rosetta-extensions` grades against the AUTHORING copy** (`6130bfd8`)
   unless the citing block pins a ref. The `stack-demo/` clone (`ab81527a`, 47 behind) is a
   different, older clone of the same repo.

## Precondition derivations (each run at this open, not inherited)

| precondition | command | result |
|---|---|---|
| platform == origin HEAD | `git -C stack-demo/platform fetch origin` + `ls-remote origin HEAD` | `0c91421dfdb08dc75f17f1aabfb61394070e770b` — equal to checkout **and** `origin/main` |
| guard family | `guard_family.py --repo-root … --platform …` (bare, and with `--range`/`--verify-remote`) | **14 GREEN · 0 RED · 3 input-gated** over 17 members |
| READ instrument frozen | `shasum -a 256` + `git log --follow` | `3858ec53…6eb0`; **exactly one commit ever** (`012edd2`, iter-76) |

The verbatim copy handed to the seats was sha-checked **after** copying and matched the source
byte-for-byte.

### On the guard-family count

The run brief said **15 GREEN**; the measurement is **14**. This is a transcription slip in the
brief, not a regression: the **pass-22 hardening-ledger entry itself records** `14 GREEN · 0 RED ·
0 could-not-check · 3 not-run` over 17 members. Instrument, ledger and this re-derivation agree.

The 3 input-gated members are `repair_leak_guard` and `value_change_guard` (need a `--range` whose
diff actually touches published prose) and `repair_reach_guard` (needs a per-run repair ledger).
Supplied a range, the first two report **CANNOT-RUN — "Nothing was checked; this is not GREEN"**
rather than green-over-nothing. That is the iter-94 anti-vacuity property working, and it is the
reason the honest ceiling here is 14 and not 17.
