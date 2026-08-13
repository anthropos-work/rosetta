**Type:** tik — under `TOK-05` (*stop repairing claims; fence the predicates under them*), step 2 of 4.

# iter-63 — the citations, and the hole a ref-pin leaves

## Phase A — re-measure the citation class, before touching anything

Both inherited figures failed, and for the same reason: **each measured a subset and named it the
class** (`D-M257x-63-3`).

| reading | figure | what it actually counted |
|---|---|---|
| iter-58 (protocol §7 rule 4 records *"22 of 23"*) | 21–22 of 22–23 moved | **raw sites** of the string `main.go:N` — pooling `app/`-qualified, bare, and `cmd/*/main.go` forms, three different files under one name |
| iter-61 | 5 of 16 distinct still landing | **distinct** citations of the two app-mainline forms — the right unit, but blind to the bare `:N` **continuation** construct (`` `app/main.go:446`, `:524`, `:992` ``) |

Derived enumeration over **both** constructs, resolved against the app clone on disk at `b948604`
v1.366.0:

```
104 citation SITES / 86 DISTINCT citations across 22 corpus files land in `app`
 18 of those are the app MAINLINE (`main.go`) — the routed class
  5 HELD · 13 MOVED
```

So **18** is the mainline denominator and **86** the denominator of the class §7 rule 4 actually
names (*"every corpus citation whose path lands in the advancing repo"*). 16 and 23 are both readings
of a subset. The 68-citation non-mainline remainder is routed WHOLE.

## Phase B — repair, adjudicated against platform artifacts

All 13 moved mainline citations re-pointed with the ref written beside each one. The interesting part
is what adjudicating against artifacts surfaced that the citation work was not looking for
(`D-M257x-63-5`):

> `docker-compose.yml:171-183` @ platform `0dab54d` sets **all four** `*_RPC_ADDR` to
> `http://backend:8083`, under compose's own comment *"cms + jobsimulation are folded into app: all
> four RPC edges are the one backend mux"*. **M809 has landed.**

`messenger.md`, `cms.md`, `jobsimulation.md`, `dependency_map.md` and `backend.md:195` all asserted
the two-of-four split as **current** — three of them emphatically (*"current, not stale"*).
`platform-migration-status.md:76` had it right the whole time, naming `docker-compose.yml:174`/`:176`.
**The fenced map was right and the prose was wrong**, which is what the map is for.

Also repaired in the same sweep, the anti-repair language §5 rule 31 predicts: `platform-alignment.md`'s
own *"merged-in-production is not removed-from-compose"* corollary and its colony-pin enumeration both
stood in the present tense about husk containers `d11a403` deleted. The corollary is **kept — as a
PHASE, with the phase's end dated.**

## Phase C — the induced half, applied to this iter's own edits

§7 rule 4 was written for a pin advance; it applies identically to a corpus repair (`D-M257x-63-4`,
now §5 rule 34). This iter's edits moved **9 intra-corpus citations across 6 files**, two of them in
root `CLAUDE.md`. `anchor_construct_guard` caught exactly **one** — the one that happened to land on a
blank line. The other eight landed on content and would have read as correct. Re-pointed from the
`git diff -U0` line map, in the same commit.

## Phase D — the ref-pin decision, and the fence

`CHECK-M257x-iter60-stale-pin-exemption` is **three mechanisms wearing one name** (`D-M257x-63-1`),
and only the third is the one the briefing described:

| # | mechanism | instance | what it really was |
|---|---|---|---|
| 1 | pin crosses a **row** boundary | `shared_libraries.md:41-42` | a window bug, not a policy question |
| 2 | pin crosses a **cell** boundary | `service_taxonomy.md:98` (one row, two clauses) | the same bug one level finer |
| 3 | pin cited **as evidence of currency** | `service_taxonomy.md:55`, `messenger.md:108` | the policy question |

**Decision: a ref-pin is a DATE, not an exemption.** Not an expiry (a threshold, and §4 Trap A) — *age
is not the variable, tense is*. Not a mandatory two-sided citation (right as a repair, wrong as a
requirement; it would force every historical sentence to restate the present). The rule is: **a pin's
scope is the claim's own block — a markdown CELL in a table, a wrapped sentence in prose; a pin naming
the ref the checker derived from exempts nothing; and a block that asserts currency cannot be pinned
into silence.** Blast radius measured *before* adopting: 19 corpus blocks are both pinned and assert
currency; exactly **1** also carried a checked construct.

And the fence was still not enough, because of `D-M257x-63-2`: `service_taxonomy.md`'s Services table
names its profile column **fourth**, and G1 required it **first**. Those six rows were not exempt —
they were **unreachable**, never once looked at. *Fifth time in this milestone that a GREEN reading
turned out to be a reach limit.*

**Watched RED, then GREEN.** Widened + re-scoped guard on the live corpus: **17 illegal profile sites
across 2 files, every one real — zero guard-own findings** (iter-60 was 16/37 its own, iter-61 13/35;
column identity beats substring matching). Repaired to **GREEN**: `service_taxonomy.md`'s Services
table rewritten from the artifact (ten compose services, default `core` selects five, storage moved
out of the default selection into `storage-legacy`), rows 98/99 flipped from *"YES — still starts"* to
*"NO — gone at `0dab54d`"*, the tier summary corrected, and `shared_libraries.md`'s colony pin split
corrected **three-way → two-way**.

One finding was the fence catching **its own repair's quotation** of the retired construct — rephrased,
and routed as a class to watch.

## Phase E — gates

| gate | result |
|---|---|
| `platform_predicate_guard` | **OK** — reach 90 profile sites / 8 tokens, 13 RPC claims, ref-pinned skips **2/2/2** (was 2/3/7), guard ref `0dab54d…` resolved |
| `anchor_construct_guard` | OK — every resolvable anchor names a construct |
| `markdown_structure_guard` | OK — 112 published files |
| `repair_postcondition` (pre-commit) | OK on both commits |
| `tests/test_platform_predicate_guard.py` | **54 tests** (was 35), all pass |
| mutation battery | **4 inversion mutants, all caught**: revert-to-2-line-window (3 tests RED), disable-currency (2), own-ref-exempts-again (1), profile-column-must-be-first (3) |
| `stack-core` suite | **664 tests, 1F** — the perishable iter-48 fixture, the single expected failure. Back at baseline |
| `dev-stack` suite | **151 tests, OK** |
| `stack-injection` | OK |

Two things the suites caught that inspection did not:

- **`test_test_collection_fence` went RED on my own edit** — I appended the new test classes *after*
  `if __name__ == "__main__": unittest.main()`, so they were invisible to direct execution. The fence
  is there for exactly that and it fired; guard moved to the end of the file.
- **`dev-stack` must be run ALONE.** Its first run reported **6 failures**, all in
  `test_dev_public_host` — and every one vanished on a solo re-run (151/OK). `stack-core`'s m220
  battery **spawns nested `dev-stack` runs**, and its own log says so (*"a nested run was re-taken
  under contention"*). Running the two suites concurrently is self-inflicted flake, not a finding.

## Close — 2026-08-04

**Outcome:** the routed citation class re-measured (**18 mainline of an 86-citation `app` class**, not
16 and not 23 — both prior figures were subset readings), 13 moved citations repaired, and
`CHECK-M257x-iter60-stale-pin-exemption` **answered and fenced** — a ref-pin is a date, not an
exemption, and the hole was three mechanisms of which two were a window bug. Repairing it exposed a
sixth reach limit (the profile column must be found by header, not position) whose 17 sites were
**unreachable rather than exempt**. Compose refuted four service docs at once: **M809 has landed.**
**Zero guard-own findings** in the widened fence — the first widening this milestone with a clean
signal-to-noise ratio.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — 4 of 5, unchanged; clause 5 is still graded only by a reading that returns zero.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** `D-M257x-63-1` (a ref-pin is a DATE, not an exemption — three mechanisms, not one),
`D-M257x-63-2` (the profile column is found by HEADER, not position), `D-M257x-63-3` (both prior
citation sizings were subset readings), `D-M257x-63-4` (a corpus repair moves the corpus's own line
numbers), `D-M257x-63-5` (compose refutes four service docs; the fenced map was right).
**Side-deliverables:** none.
**Routes carried forward:**
- `FIX-M257x-iter63-app-citation-residual` → the 68 non-mainline `app`-resolving citations, routed
  WHOLE with their mechanical grading and its known instrument artifacts named.
- `CHECK-M257x-iter63-quoting-a-retired-token` → the fence flags a *quotation* of a retired construct;
  rephrasing sufficed for one site, a quotation discriminator if the class grows.
- `DOC-M257x-iter59-storage-mid-fold` (the map's 8th `mid-fold` token) ·
  `CHECK-M257x-iter60-g6-citation-subject` · `FIX-M257x-iter53-union-set` (**PENDING USER DECISION**) ·
  `FIX-M257x-iter56-assignment-flake` (**NOT DECIDED** — needs a failure *rate*) ·
  `CHECK-M257x-iter38-ai-act-classification` (needs an owner outside this milestone) ·
  `CHECK-M257x-iter57-anchor-guard-bare-class` · `FENCE-M257x-iter54-refs-block` ·
  `FIX-M257x-iter57-within-block-drift` · `CHECK-M257x-iter58-derive-preregistrations` ·
  `CHECK-M257x-iter52-second-ai-manager` · `-cold-daemon-registry` · `-grep-vs-failclosed` ·
  `-empty-stdout-class` · `-baseline-refs` · RF-2/3/7–13 · root `CLAUDE.md`.

**Lessons:**

1. **Measure the mechanism before designing the policy.** The briefing offered three shapes for a
   pinned claim — expiry, two-sided citation, something else. Two of the three instances turned out
   not to be a policy question at all: the exemption reached further than the claim it was written
   for. **A "policy hole" is often a window bug wearing a policy's name.**
2. **Age is not the variable — tense is.** *"`fdfa189` removed intelligence"* is pinned, past-tense
   and true forever; *"at `2adcf71` `CMS_RPC_ADDR` reads `http://cms:8091`"* is pinned, present-tense
   and was false within days. Any expiry would have hit the first and missed the second.
3. **A fence can be silent for two different reasons and only one of them is an exemption.** Six of
   the 17 sites were *unreachable*, not *exempt* — the fence had never looked at that column. The two
   are indistinguishable from the outside, and only one is fixed by changing a rule.
4. **§7 rule 4 is about line numbers, not about pins.** The corpus cites itself ~200 times; every
   repair that changes a line count is an advance. This iter's own edits moved 9 citations and the
   anchor guard caught 1 of them.
5. **A mutant that does not actually invert the rule proves nothing.** The first draft of the
   window mutant fell through to the prose branch, so the row rule survived it and the battery
   reported a pass that meant nothing. §8 rule 5's corollary, met in person: write the true revert,
   re-run, and read which tests go RED — not merely *that* some do.
6. **Two suites that run each other cannot run at once.** `dev-stack` failed 6 tests beside
   `stack-core` and passed all 151 alone; `stack-core`'s m220 battery spawns nested `dev-stack`
   runs. A concurrency-induced RED reads exactly like a regression.
7. **Emphatic language keeps marking the falsest claim.** *"current, not stale"* (three sites) and
   *"YES — still starts"* (two rows) were all wrong. Second occurrence in this milestone after §5
   rule 31's anti-repair fortification. **Confidence in the prose is not evidence; the artifact is.**
