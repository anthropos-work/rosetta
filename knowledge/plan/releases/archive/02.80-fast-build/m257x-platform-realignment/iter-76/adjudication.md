# iter-76 adjudication — `FIX-M257x-iter76-read-union`

**Status: COMPLETE.** 152 booked blockers across all 14 seats, every one graded. Adjudicated in
run 52 (iter-80), four parallel adjudicators, one per seat-group, each re-deriving from the
platform/service clones rather than from any prior verdict.

> **Binding condition 1 of the routing — *"adjudicate before repairing"* — is now discharged.**
> Condition 3 (*"not closed until the G5/G2 reach hole is closed"*) was discharged by iters 77–79.
> Condition 2 (*"repair by PREDICATE, not by claim"*) is what §3 below exists to make possible.

## Verdict

| | r13-A..D | r13-E..G | r14-A..D | r14-E..G | **total** |
|---|---|---|---|---|---|
| booked | 32 | 45 | 36 | 39 | **152** |
| **UPHELD** | 31 | 38 | 34 | 37 | **140** |
| REJECTED | 1 | 7 | 2 | 2 | **12** |
| UNSETTLED | 0 | 0 | 0 | 0 | **0** |

**Upheld rate 92.1 %.**

### This is the seventh routed count in the milestone, and the first that did NOT collapse

The six before it collapsed on derivation — **64→5, 23→1, 21→0, 92→0, 4→3, 145→3** — and that
track record is exactly why iter-76 routed its own 77/75 rather than repairing it, and why the
adjudication was made a binding precondition. The prior was *"it will collapse."* It did not:
**152 → 140.**

That prior was well-earned and stating it is not a criticism of it. But it now has a counter-example,
and the counter-example is the one that decides the gate. **The instrument was not crying wolf.
Clause 5 is genuinely far from met, and the ~150 is a work item after all** — with the single
documented exception below, which was the ~150's one *systematic* false-positive class and which
turned out to account for **4 findings, not most of them**.

### What the 12 rejections were

| mechanism | n | note |
|---|---|---|
| **271↔387 line-count** (block pins `2adcf71`, where `docker-compose.yml` is 387 lines) | 4 | The class iter-76 predicted would dominate. It did not. In all but these 4 the citing block names **no** ref, so §5 rule 33 offers no rescue — *and* the claim itself is false, not merely the anchor |
| **pin in a subordinate clause** | 2 | `graphql-wundergraph.md:177` — *"the `graphql` profile, which **since `2adcf71`** contains no router"*. True at that ref. Adjudicators split on this one; the predicate lands anyway via an unpinned sibling |
| **ref-relative truth read as self-contradiction** | 2 | Two statements 58 lines apart, both true, at `app 9d00a313` and `b948604` respectively. Adjacency in a file is not co-reference |
| **emphatic scalar scoped by its own argument** | 1 | *"reads exactly one table"* inside a paragraph refuting a *different* table. Literally false, fails the *would-misdirect* limb |
| **off-by-≤8 onto a blank line / code fence** | 3 | True, but self-evidently wrong to anyone who opens the file. Minor-grade |

## The repair unit — 11 predicates

The 140 upheld dedupe hard. All four adjudicators converged on substantially the same list
independently, which is the property that makes repair-by-predicate viable.

| # | predicate | ground truth @ `0dab54d` | weight |
|---|---|---|---|
| **P1** | *"the `cms` / `jobsimulation` / `roadrunner` containers still exist, and the default profile starts them"* | No such compose services. `d11a403` deleted all three | **dominant — ~47 findings, 6+ files** |
| **P2** | *"`repos.yml` has 9 entries / still lists roadrunner+jobsimulation / `repos.yml:17-19` is the jobsim entry"* | **6** entries; `:17-19` is sentinel/storage | ~12 |
| **P3** | *"compose declares **nine** services"* | **8** declared, **10** effective with `include: common.yml`. Nine is a count of nothing | 5 *(partly repaired)* |
| **P4** | *"`graphql` is a live profile / the default"* | `Makefile:10 PROFILE ?= core`; `graphql` is in no `profiles:` key and silently selects only the 3-service floor | ~10 |
| **P5** | *"`core` starts nine containers / six Go services"* | `--profile core` → `postgresql redis sentinel backend gotenberg` = **5 containers, 2 Go services** | ~8 |
| **P6** | *"`storage` is in the default set"* | `profiles: [storage-legacy]` (`docker-compose.yml:134`) | 3 |
| **P7** | **stale compose line-anchors** in blocks that pin nothing (`:83`, `:144`, `:281`, `:311`, `:337-341`, `:352`, `:361`, `:362`, `:45/:99/:160`, `:70-80`, `:213-217`) | Resolve at `2adcf71`; land on unrelated constructs at `0dab54d` | ~14 |
| **P8** | **the `external_services.md` re-point is short by 8–9 lines** — 9 sites still cite `:546`/`:570`/`:578-588` | Retraction is `:554`; four-ways list `:578-581`. A prior repair applied **+1 of the +9** | 9 |
| **P9** | *"`STORAGE_RPC_ADDR` is read by `main.go` at `9d00a313`"* | **0** read sites there (3 hits, all comments). True at no ref. **Load-bearing — this is the evidence that moved `storage` off `mid-fold`** | 3 |
| **P10** | **wrong commit attribution** — `pms:74` promotes `d11a403`'s own factually-wrong commit message (*"its repos.yml entry was already gone"*) into a conclusion; `architecture_overview.md:181` attributes the husk removal to `2adcf71` | `git show d11a403 -- repos.yml` shows **that commit** deleting the entry; at `2adcf71` all three husks still started | 4 |
| **P11** | **false scalars / sets against source** — 4→5 usage KPIs · 5→6 snapshot operators · "four bring-up patches"→five · phantom `python-docx` · "the last migration in the repo" (two post-date it) · "clerkenstein ships in its own repo" (it is a monorepo section) · roadrunner port defaults 10400/10401 (binary defaults are 8080/8081) | per-source | ~12 |

## One finding is an operational hazard, not a documentation defect

**`storage.md:55,:154,:181` say local private storage is sandboxed to `/tmp`.**
`docker-compose.yml:82` @ `0dab54d` sets `STORAGE_S3_BUCKET=production-storage2024…` **on `backend`**,
and `app 9d00a313 main.go:463→471` reads it straight into `NewManager`. So **local private writes
land in a production bucket**, while the doc reassures the reader about precisely that manager.

This sits outside clause 5's *"would misdirect real work"* framing — it would misdirect real work
*into production*. Routed as **`DEF-M257x-iter80-storage-prod-bucket`**, severity **high**. It is a
platform/compose fact rather than a corpus fact, so the corpus-side repair is to *stop reassuring*,
not to re-anchor. Raised to the user in the run report; **not** actioned here, because the platform
side of it is a platform-repo question and this milestone is zero-platform-edit.

## Method notes the repair must inherit

1. **The seat headers misreport the corpus ref.** All 14 declare `1937e1f`; the tree actually read is
   `b8a1fb0`/`2829b6e` (established by matching every `wc -l` positive control). Grading at the
   declared ref would have spuriously rejected findings for *"the quote does not resolve."*
2. **Seat-supplied line *contents* are unreliable even where the verdict holds** — three separate
   cases where the seat quoted the wrong text for a line whose finding was nonetheless correct.
   **The repair must re-derive every line from source, never from these notes.**
3. **Re-pointing a P7 anchor is not always the repair.** `dc:70-80` really *is* backend's
   `depends_on` incl. `cms`+`storage` at `2adcf71`. These are not drifted line numbers; they are
   **whole facts that `d11a403` deleted**. Re-pointing them would produce a correctly-cited false
   statement — §4 Trap A wearing a citation.
4. **Demo docs must be graded against the injected override, not platform compose.**
   `gen_injected_override.py:420` rewrites both frontends with `profiles: !override [core]`, so a
   `corpus/ops/demo/**` profile claim can be true where the base compose says otherwise. One
   rejection came from exactly this.
5. **A document-level *"RE-GROUNDED against `<ref>`"* banner reads like a pin but is not one under
   rule 33.** This recurs and wants an explicit ruling before the repair, or it will manufacture
   findings out of coherent, correctly-dated anchor sets.

## Prior state of this file

Before this run it held one adjudicated pair (r13-F B13/B14, both REJECTED on the 271/387
mechanism). Both were re-derived independently here and both verdicts stand.
