# iter-78 — progress

**Type:** tik, under `TOK-05`. Planned deliverable: **settle `CHECK-M257x-iter76-compose-service-count`
by derivation, repair the sites, and fence the predicate — or report the fence as a measured
negative.**

---

## Phase A — derive both counts, and the history

The unsettled question was *"8 vs 9 vs 10"*. It has three answers because **two of them count
different things and the third counts nothing.**

```
docker-compose.yml alone : 8   (backend customerio-sync gotenberg messenger
                                next-web-app sentinel storage studio-desk)
+ include: common.yml    : 10  (adds postgresql, redis — the always-on floor)
```

Across every ref the corpus cites, `docker-compose.yml`'s own count is:

| ref | date | compose-only | effective |
|---|---|---|---|
| `236771f` | 2026-07-29 | 12 | 14 |
| `b56d731` | 2026-07-31 | 12 | 14 |
| `2adcf71` | 2026-07-31 | 11 | 13 |
| `d11a403` | 2026-08-03 | 8 | 10 |
| `0dab54d` | 2026-08-03 | 8 | 10 |

**Nine was never a count of anything.** The `graphql` service was deleted at `2adcf71` (12 → 11) and
cms/jobsimulation/roadrunner at `d11a403` (11 → 8); the sequence never passes through 9.

## Phase B — adjudicate every site

Ten sites match a *"N services"* shape with a compose context. Adjudicated individually:

| site | states | verdict |
|---|---|---|
| `architecture_overview.md:191` | 8 | **correct** — `docker-compose.yml` alone |
| `dependency_map.md:31` | 10 | **correct** — the effective topology |
| `platform-alignment.md:194` | 14 | **correct** — a *dated* measurement, *"At 2026-07-31 these read: …"*, and at `b56d731` the effective count was 12 + 2 = 14 |
| `external_services.md:296` | nine | **RED** |
| `cms.md:35` | nine | **RED** |
| `jobsimulation.md:35` | nine | **RED** |
| `platform_repo.md:59` | 11, and names `graphql` | **RED** — stale by two folds |
| `staging-bringup.md:370` | 14 | **RED**, different predicate — running containers under `--profile all` |
| `staging_from_dump.md:323` | ~14 | **RED**, same |
| `service_taxonomy.md:57`/`:427`, `external_services.md:171`, `update_guide.md:206` | 3 / 7 / — | **correct** — the floor, the application services, no count at all |

`|select(all)|` derived for the two staging sites: **8**, and `all` no longer contains `messenger`
or `storage` — `0dab54d` dropped both, because running either alongside `app` means two consumers on
one Redis group or two writers on one bucket.

## Phase C — repair, by predicate

Six sites. Each now states **which set it counts**, which is the actual defect — the numbers were a
symptom of an unstated qualifier:

- `external_services.md:296`, `cms.md:35`, `jobsimulation.md:35` — *"declares **eight** services —
  ten in the effective topology, once `include: common.yml` adds the `postgresql`/`redis` floor"*.
- `platform_repo.md:59` — 11 → **8 declared / 10 effective**, and the `graphql` entry removed from
  the service list with its deletion ref.
- `staging-bringup.md:370`, `staging_from_dump.md:323` — 14 → **8**, with the reason `all` shrank.

## Phase D — the fence, measured before it was built

**The broad construct was rejected on its own numbers.** *"A number followed by `services`, in a
compose context"* reaches **14 live sites** and would fire on 9, of which 4 are true — **44%
precision** — because this corpus counts the floor (3), the application services (7), *"the last two
subgraph services"* and plain narrative (*"the eleven services that were perfectly fine"*) with the
same two words.

**G10 ships the narrow construct:** a **declaration verb** (`declares`/`defines`) with a **compose
subject in the same block**. Measured: **4 sites, 4 true, 100% precision**, and it misses nothing of
the class.

Two details, each measured rather than assumed:

- **Number words are in the pattern.** All three false sites write *"declares **nine** services"* —
  a digits-only pattern reads **none** of the live defects.
- **The window is the block.** Line-scoped, the construct reached **2 of 4** — two live sites wrap
  their sentence. Block-scoped it reaches all four at unchanged precision. **The third window bug of
  this milestone** (`_pin_window` iter-63, `_NEGATED` iter-68, this).

**G10 asserts the PAIR, not a value** (`D-M257x-78-1`) — 8 and 10 are both real, the corpus contains
one correct document of each, and a fence that picked a side would have gone RED on a correct
document and "repaired" it into a different truth.

**And it inherited iter-77's ref rule, which immediately paid.** `external_services.md:296` states
*"platform `0dab54d`'s compose declares nine services"* — about the guard's own ref — inside a block
whose leftmost pin is `b948604`, an **app** sha. Under the plain exemption the sentence was
unchecked; resolving the ref *in the platform clone* (a sha from another repo cannot date a platform
file) grades it and it goes RED. Second confirmed instance of that mechanism.

### Positive controls (§8 rule 5)

- **No-op that SURVIVES** — the *effective* count is accepted exactly as the file-local one is; both
  name a real set, so a correct document stating either stays GREEN.
- **INVERTED mutant** — a count matching neither goes RED and the finding says *"N is neither"*.
- **Reach controls** — a count with no declaration verb, and one with no compose subject, are **not
  reached at all** (`service_count_claims == 0`), so the 44% rule is provably not what shipped.
- **The wrapped-sentence control** — a claim split across two lines is still reached.

## Phase E — re-measure

```
G10 4 compose-service-count claim(s) (0 read at a historical ref)
platform_predicate_guard: OK
```

Nine new tests (140 → **149**). Five corpus guards GREEN.

## Close — 2026-08-05

**Outcome:** the milestone's one **explicitly unsettled** denominator is settled — not by choosing
between 8, 9 and 10, but by deriving that **8 and 10 count different real sets and 9 counts
nothing**, at any ref in the file's history (12 → 12 → 11 → 8 → 8). Three documents asserted the
nine. Six sites repaired so each states *which set it counts*, which was the real defect; two of them
(`platform_repo.md:59` naming a `graphql` service deleted two folds ago, and the staging pair
counting 14 running containers where `--profile all` now selects 8) had gone stale in a way no
existing assertion could see. **G10** fences the predicate at **100% measured precision** after the
obvious construct was measured at **44%** and replaced rather than thresholded.

**Type:** tik
**Status:** closed-fixed — every planned phase ran; the fence cleared its pre-registered precision
bar and shipped, and the alternative was measured and recorded rather than assumed.
**Gate:** NOT MET — 4 of 5, unchanged. Clause 5 is not re-cut.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n
— (5) cap-reached: n (2 tiks of 5) — (6) protocol-stop: n — Outcome: **continue**.
**Decisions:** `D-M257x-78-1` (two real sets and a count of nothing — so the fence asserts the pair),
`D-M257x-78-2` (the broad construct measured at 44% and replaced, not thresholded; number-words and
the block window each measured), `D-M257x-78-3` (iter-77's cross-repo-pin class confirmed live —
two instances, both false).
**Side-deliverables:** none — every change is in the iter's planned scope.
**Routes carried forward:** `CHECK-M257x-iter77-cross-repo-pin` **upgraded** (the class is real: 2
of the 145 confirmed to date a platform claim, both false; the general resolve-in-repo fix is the
natural next target) · `CHECK-M257x-iter78-running-vs-declared` (a *named* profile's container count
stated in prose is fenced by nothing; G3 covers only the default) · all iter-77 routes unchanged.
**`CHECK-M257x-iter76-compose-service-count` is CLOSED.**

**Lessons:**

1. **When two measurements disagree, ask what each counts before asking which is wrong.** iter-76
   recorded 8-vs-10 as a disagreement between its grep and its parser. Neither was wrong; the
   question was underspecified, and the third number — the one both agreed was suspicious — was the
   only real defect.
2. **A fence for a quantity with several legal answers asserts the SET.** Picking one would have
   turned a correct document RED and repaired it into a different truth.
3. **Check the window before declaring a class unreachable.** Third time this milestone. A construct
   that looks like it has 50% recall may simply be reading one line of a two-line sentence.
4. **The corpus spells its integers.** Three live defects were invisible to a digits-only pattern.
