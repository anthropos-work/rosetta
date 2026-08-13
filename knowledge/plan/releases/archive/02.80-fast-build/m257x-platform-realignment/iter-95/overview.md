# iter-95 — THE READING (readings #17 + #18), at platform `0c91421`

**Shape:** measuring pass. **No repair happens inside this iter** — whatever the read surfaces is
booked and routed. That separation is the only reason `140 → 43` meant anything.

## Why this iter is a reading and not a harden pass

Run 59's brief settled it: **no fourth harden pass. The instrument is frozen. The reading is the
whole job.** Runs 54–58 modified the **guard family** — clause 3's instrument. Clause 5's instrument
is the **graded READ**, a different object, and it has never been touched. iter-80 established that
distinction and recorded that this milestone had conflated the two before; it did so again across
five iters and three harden passes.

## The three preconditions, each re-derived at this open (not inherited)

| precondition | how it was re-derived | verdict |
|---|---|---|
| **corpus repaired at platform `0c91421`** | `git -C stack-demo/platform fetch origin` then `ls-remote origin HEAD` → `0c91421dfdb08dc75f17f1aabfb61394070e770b`, equal to the checkout and to `origin/main`; tree clean | **HOLDS** |
| **guard family GREEN, 0 RED, on fetched clones** | `rext stack-core/guard_family.py --repo-root … --platform … [--range …]` run twice (bare, and with `--range`/`--verify-remote`) | **HOLDS — 14 GREEN · 0 RED** over 17 members |
| **READ instrument byte-identical** | `shasum -a 256` on `instrument/briefing-iter76-AS-RUN.md` → `3858ec536ea613ee00dfa8b383f858486564e23815581ef5135e85ec18e52eb0`; `git log --follow` on the path → **exactly one commit ever** (`012edd2`, iter-76) | **HOLDS** |

**One correction to the run brief, in the safe direction.** The brief stated the guard family at
**15 GREEN**. The measured figure is **14 GREEN · 0 RED · 3 input-gated**, and that is not a
regression — the pass-22 ledger itself records `14 GREEN · 0 RED · 0 could-not-check · 3 not-run`
over 17 members. The brief's 15 is a transcription slip; the instrument and the ledger agree with
each other and with this re-derivation. The load-bearing half — **0 RED, and no member laundering a
non-run as green** — holds exactly.

The 3 input-gated members are `repair_leak_guard`, `value_change_guard` (both need a `--range` with
prose actually in scope) and `repair_reach_guard` (needs a per-run repair ledger). Given a range,
the first two correctly report **CANNOT-RUN, "Nothing was checked; this is not GREEN"** rather than
green-over-nothing — the iter-94 anti-vacuity property working as designed.

## The instrument — frozen, and NOT touched by this iter

| knob | value | evidence |
|---|---|---|
| briefing | `instrument/briefing-iter76-AS-RUN.md`, **byte-identical** | sha256 `3858ec53…`; delivered to seats as a verbatim copy whose sha was re-checked after copying |
| seats | **7 per reading** — full-read partitions A–F + adversarial diff seat G | unchanged since iter-41 |
| readings | **2**, identical partition (#17, #18) | the seat-variance control (§5 rule 23) |
| file set | `corpus/architecture/*.md` + `corpus/services/*.md` — clause 5's stated scope | **40 files / 10,108 lines** |
| partition method | files sorted by line count **descending**, snake-dealt A→F then F→A | re-executed; deals a different hand than #13–#16 because file sizes moved |
| blindness | seats barred from `knowledge/plan/**`, `.agentspace/scratch/`, and each other's output | restated as absolute bars in the delivered packet |

**The partition deals a different hand than iter-76's and iter-82's, and that is the method working,
not drifting.** The scope has grown `9,544 → 9,712 → 10,108` lines as the milestone repaired; a fixed
method over a changed corpus necessarily re-deals. iter-76 and iter-82 each recorded the same
consequence rather than engineering it away.

### Partition (computed by the fixed method — 40 files / 10,108 lines)

| seat | lines | files |
|---|---|---|
| **A** | 1886 | external_services · cms · backend · ai-labs · academy-backend · skiller · TEMPLATE |
| **B** | 1698 | ai-readiness · platform-migration-status · jobsimulation · sentinel · customerio-sync · services/README · architecture/README |
| **C** | 1648 | alignment_testing · clerkenstein · ai_architecture · messenger · coursebuilder · gotenberg · db-backup |
| **D** | 1632 | service_taxonomy · architecture_overview · security_compliance · roadrunner · next-web-app · dependency_map · intelligence |
| **E** | 1638 | studio-room · hiring · graphql-wundergraph · storage · clerk-integration · frontend_architecture |
| **F** | 1606 | ant-academy · studio-desk · chronos · shared_libraries · askengine · skillpath |
| **G** | (diff) | adversarial diff-read `8d6bb6c..HEAD -- corpus/ CLAUDE.md .claude/` — **54 files, +1642 −680** |

Seat G's base `8d6bb6c` is iter-82's close — the last graded reading — so G's scope is *every corpus
repair made since the last graded read*.

## Ground truth — re-derived at this open, and the one thing that IS different

The frozen briefing's sha table was captured at iter-76 and is **stale for `platform`**. The seats
were given a superseding addendum (clearly marked; the briefing text above it untouched) carrying
both the checkout **and `origin/main`** per clone — the structural half of
`CHECK-M257x-iter76-seat-ref-discipline`'s fix, adopted at iter-86.

| clone | checkout | `origin/main` | behind |
|---|---|---|---|
| `platform` | `0c91421d` | `0c91421d` | 0 |
| `app` | `b948604f` | `2035f9a4` | **93** |
| `next-web-app` | `bb3313bc` | `8297c684` | **41** |
| `storage` | `4ce8ece5` | `9f8cb532` | **20** |
| `messenger` | `fa47850d` | `e9421c68` | **7** |
| `ant-academy` | `9c3843cd` | `22df69dd` | **5** |
| `jobsimulation` | `462343b0` | `82cb66ec` | **4** |
| `cms` | `ca50c817` | `f38c0c4a` | **2** |
| `sentinel` | `88bc5592` | `f2c46190` | **2** |
| `studio-desk` | `14a5442a` | `41ee3575` | **2** |
| `roadrunner` | `87d8d443` | `87d8d443` | 0 |
| `graphql-wundergraph` | `60c229f3` | `60c229f3` | 0 |
| `rosetta-extensions` (authoring) | `6130bfd8` | `6130bfd8` | 0 |

**Only `platform` moved** (`0dab54df` → `0c91421d`). Every other checkout is identical to iter-86's
sheet, so the ref surface is unchanged except for the one repo whose move this milestone exists to
absorb.

## COMPARABILITY — stated explicitly, because silence is not acceptable here

**This is a DECLARED RE-BASELINE of the raw series, and a CONTINUATION of the adjudicated one —
with one honest caveat.**

- The **raw** (booked, pre-adjudication) series was **already declared discontinuous at iter-86**
  (`D-M257x-86-2`), out loud, for the seat-ref sheet. It does not become more discontinuous here.
- The **adjudicated** series — `140` (iter-76/80) → `43` booked / `40` upheld (iter-82/84) — was
  explicitly left **untouched** by iter-86. This reading extends it.
- **The caveat, and it is real:** this is the first reading taken at platform `0c91421` rather than
  `0dab54d`. The ground truth for every compose / `repos.yml` / profile claim genuinely moved
  underneath the corpus between readings. A finding count taken against a different ground truth is
  not strictly the same measurement as one taken against the old — so the adjudicated series is
  continuous **in instrument** (briefing, seats, method, scope definition, grading rule all
  identical) and **discontinuous in ground truth** for the `platform`-derived predicates specifically.
  Findings that are *not* platform-derived remain directly comparable.

Stating both halves is the point. A single number quoted without this paragraph would be the exact
defect this milestone has spent 95 iters learning to name.

## PRE-REGISTERED PREDICTIONS — written before ANY seat report was read

Written after launch (the seats were already reading, blind) but **before a single report existed on
disk or was opened**. That is the property that makes a pre-registration meaningful; the launch order
is not.

1. **Per-reading count.** Each of `N₁₇`, `N₁₈` lands in **[0, 12]**. The gate needs zero; the
   milestone's own track record says a small repair list is likelier.
2. **Neither reading returns zero.** Stated unsoftened because a prediction written to be safe is
   not a prediction — and because this is the prediction the gate most wants falsified.
3. **Union exceeds both.** `|#17 ∪ #18| > max(N₁₇, N₁₈)` — each reading books at least one blocker
   the other misses. Falsified if either reading returns 0.
4. **Per-pass recall stays below 60 %.** The 43–51 % single-pass figure has held across every paired
   measurement in this milestone.
5. **A platform-derived class dominates.** More booked findings trace to the `0dab54d → 0c91421`
   ground-truth move (compose services, `repos.yml`, profiles) than to any other single cause.
6. **The upheld rate stays high.** ≥ 80 % of booked findings survive adjudication, continuing
   iter-80's 92.1 % and iter-84's 93.0 % — i.e. the instrument is still not crying wolf.

## Escalation conditions, written in advance

- **More than ~15 union blockers** → measure and route; **do not repair inside this iter**.
- **A seat that cannot state a `wc -l` for an assigned file** → that seat's reading is void and is
  reported as void rather than counted.
- **Any finding inside a class a prior iter closed by adjudication** → that close is **retracted**,
  not defended.

## Held by instruction, and disclosed rather than folded in

- **`DEF-M257x-iter80-storage-prod-bucket`** — `storage.md:55/:154/:181`. Still the user's, still
  open. If the read books these anchors they are marked **held-by-instruction** and their
  contribution to `N` is stated separately, both ways, so the number can be read either way.
- **The five post-freeze items** (UNMEASURED reach keys graded as exit 0 · unreported waivers ·
  crashed guards rendering as RED with tracebacks itemised as findings · `demo_knob_guard` lacking a
  vacuity control · demopatch's silent fallback on a corrupt journal) stay **disclosed in the pass-22
  ledger and routed**. Freezing with known routed items named out loud is the honest way to freeze.
- **`cms` deleted its build workflow on 2026-08-04** stating the ECR *"is decommissioned"*, pointing
  opposite to what the corpus records for M810's cms half. The deletion lands in `infrastructure`,
  which is **in no clone set**, so it is **reported, not asserted** — and
  `unreadable_repo_claim_guard` (GREEN this open) is the fence that already labels that class as
  not-a-measurement.
