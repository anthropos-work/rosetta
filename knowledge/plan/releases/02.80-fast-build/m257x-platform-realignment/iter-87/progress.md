**Type:** tik, under `TOK-05`. Shape prescribed by TOK-04 P3 / §5 rule 26 — *the iter that detects the
move re-points it, in that iter.*

## What happened

The platform moved a third time, mid-milestone, and our ground truth was pre-drift. `stack-demo/platform`
sat 2 commits behind `origin/main`; `838d907` (merged `0c91421`, 2026-08-05, PR #26) **deleted the
`storage`, `messenger` and `customerio-sync` compose services outright** — not to a rollback profile, out
of the file — and removed `storage` + `messenger` from `repos.yml`.

The clone was advanced to origin HEAD, ground truth re-derived with a control, the guard family run, and
the corpus repaired by predicate. Details in [`ground-truth.md`](ground-truth.md) and
[`decisions.md`](decisions.md).

## The question the hand-off called decisive: did the membership fence fire?

**Yes — assertion B, 2 for 2, unaided, on a tree nobody had touched.**

```
[B departure] the map claims messenger is in repos.yml, and it is not — a service left the clone set
              and the map still asserts it
[B departure] the map claims storage is in repos.yml, and it is not — …
```

Run with a control that holds everything else fixed: a detached worktree at `0dab54d` gives **0**
assertion-B findings, the advanced checkout gives **2**. The delta is the event, not accumulated debt.
This is the fence's second unaided catch of a live membership change and the **first entry in §1's fold
table found by an instrument rather than by a breakage.**

**And the derived layer did better than the fence: it needed no repair at all.** §2's time bomb was
forecast to fire on *"the day they leave the clone set"* — 13 write targets failing 42P01 at once. This
was that day. Measured across the move, `repos_yml_migration_pairs` (`app:public`) and
`repos_yml_schemas_to_create` (`extensions sentinel public`) are **identical at both refs, and identical
correctly**, with zero human action. Third consecutive platform change absorbed unaided (`D-M257x-87-6`).

## The finding the hand-off did not anticipate

Its opening reading of *"13 GREEN · 0 RED · 3 not-run"* is **not reproducible, and the reason is a
mechanism, not a miscount.** Re-measured at the **identical** `0dab54d` checkout after a `git fetch`:
**10 GREEN · 3 RED**. The three citation-resolving guards read `origin/main` by the iter-68 `CITE_REF=auto`
ladder, so **the fetch armed them, not the checkout**. A citation fence pointed at an unfetched clone reads
GREEN. Landed as §5 **rule 41**, and it is what makes the clone-advance rule derivable rather than
preferential (`D-M257x-87-1`).

Second instrument finding: the family runner reported `lines[-1]` as a RED guard's headline, so the
21-finding alignment RED was summarised by a `gotenberg` citation nit while the two `[B departure]` lines
— the actual news — were invisible in the only view that claims to speak for the whole family. Repaired to
*"N finding(s); first: …"*, derived from the producer's own ordering. §5 **rule 42**, +5 tests
(`D-M257x-87-4`).

## The repair — 38 findings, six predicates, five disjoint packets

| guard | findings at open | at close |
|---|---|---|
| `platform_alignment_guard` (B departure ×2 + F ×19) | 21 | **0** |
| `platform_predicate_guard` (G1 ×3 tokens/28 sites, G8 ×3, G10, G2, G4 ×9) | 17 | **0** |
| `anchor_construct_guard` | 15 | **0** |

Repaired by predicate tree-wide, never by file (§5 rule 19), across `CLAUDE.md` + `.claude/skills/**` +
`corpus/architecture/**` + `corpus/services/**` + `corpus/ops/**`. Substantive corrections, not just
re-anchoring:

- **`customerio-sync` changed STATE** — `live-standalone` → **`merged-into-app`**, cited to the platform
  stating the fold inline on its own `backend` block plus `app`'s `CUSTOMERIO_SYNC_ENABLED` gate. No plan
  doc had named this third service.
- **`storage`'s prod ECS block is DELETED, not scaled to zero** — `storage/terraform/main.tf` is 18 lines;
  the paired *"each `service_desired_count = 0`"* sentence was **half false** and was split (messenger's
  half survives, at `:29`).
- **The `messenger → backend` RPC edge is gone, not re-pointed** — the messenger block was the only thing
  setting all four `*_RPC_ADDR` variables. Compose now sets **exactly one** service address
  (`AUTHORIZATION_ADDRESS`, `docker-compose.yml:48`) and **zero** `*_RPC_ADDR`. `backend → sentinel` is the
  only cross-process edge left.
- **iter-86's own repair was among the falsified** — it had just written that `storage` sits in
  `profiles: [storage-legacy]` and `messenger` in `profiles: [messenger]`. Correct at the ref it measured;
  false one day later. §5 rule 33 arriving as a live event rather than as doctrine.

Carve-out honoured: `storage.md`'s three carved claims verified **present verbatim**, and no diff touches
the production-bucket literal — `DEF-M257x-iter80-storage-prod-bucket` is untouched and still undecided.

## The suite found two more, and one of them is this milestone's own class

The rext suite is **35 files, 34 OK, 1 FAILED** — the failure being the documented perishable iter-48
fixture (TOK-05's recorded `1F/610` baseline), unrelated to this iter. Two files failed *because of the
move* and were fixed here:

- **`test_platform_alignment_guard`** — two anti-vacuity asserts read
  `assertGreaterEqual(len(declared), 5, "repos.yml parsed suspiciously few repos")`. `repos.yml` now has
  **4**, so an anti-vacuity guard failed for the one reason it must not be sensitive to. **The count was
  the wrong instrument**: what it stands in for ("the parse returned something real") is a property of the
  PARSER, while the number is a property of the PLATFORM — and the platform's number is exactly what this
  suite exists to track. **§2's hand-maintained list, one level in, inside the fence's own tests, where no
  fence was watching.** Replaced with a positive control (§5 rule 2): non-empty **and containing `app`**,
  which is structurally guaranteed (the only repo with `migrations: true`, the only one with a `schema:`
  key) and therefore cannot drift as services fold. Mutation-verified against a synthetic `repos.yml`
  holding only `sentinel`.
- **`test_service_doc_status_fence`** — correctly caught that `messenger.md` had no top-of-file
  merged-into banner after the packet rewrite. Added, matching the house shape used by `storage.md` and
  `customerio-sync.md`.

The first of those is the better finding: **a fence's test suite is a place hand-maintained platform
constants hide**, because the suite is not itself fenced.

## And the commit-scoped guards said the commit had not finished — twice

Run with `--range` against the iter's own commit (the vantage iter-86 built them for):

- **`repair_leak_guard` RED, 19 sites** — claims I had rewritten in one file while the identical claim
  stood in a twin. §5 rule 19's cross-file drift, caught at commit time, in my own repair. All 19 closed;
  the closing pass then found **5 more that its own rewordings created** and closed those by restoring the
  shared wording verbatim rather than propagating a reword to four more files. **`repair_leak_guard` is
  now GREEN — "the commit left no old form standing."**
- **`value_change_guard` RED, 7 sites — all adjudicated NOISE, by reading them.** Two are about **Clerk**
  credentials where the corrected site was about **Brevo** (different subject entirely); one is a section
  *heading* (`### Update All Repositories`) caught by the fuzzy 90-token window; four are the bare
  cross-reference verb `see` → `read`. This is exactly `CHECK-M257x-iter86-value-change-weak-form`, the
  3-token form iter-86 measured at **0/2 precision** — and the waiver mechanism deliberately **cannot
  express "false positive"**, so they stay RED rather than being laundered. Renaming a heading to satisfy
  a token window would be fitting the text to the guard.

**The runner fix broke its own rule and the fix caught it.** `guard_family.py`'s new headline counted
finding-SHAPED lines — correct for a guard printing one line per finding (alignment: 21/21), wrong for one
printing a site line plus two indented details: `value_change_guard` reported **7 site(s)** and the headline
said **14**. A wrong count is worse than no count, and it was wrong *in the same iteration that wrote
"always print the cardinality" into the rule*. §5 rule 20's *"the author of a newly-written rule violated
it while writing it"* — evidence **for** the rule, not against it. Corrected to take the cardinality from
the guard's own summary line; live output now reads `7 site(s)`. One residual is **pinned as a limitation
in a test** rather than guessed at: for a site+detail block the headline shows the detail, not the path.

## Close — 2026-08-05

**Outcome:** the third mid-milestone platform move absorbed in the iter that detected it — clone advanced
to origin HEAD `0c91421`, ground truth re-derived with a control, 38 fence findings across six predicates
repaired to **13 GREEN · 0 RED · 3 not-run**, and two instrument defects fixed (a citation fence that reads
GREEN until you fetch; a family summary that hid the event it caught). The membership fence fired unaided,
2 for 2, and the derived layer needed no repair at all.
**Type:** tik
**Status:** closed-fixed — every declared line landed; the carve-out held; nothing routed away
**Gate:** NOT MET — **4 of 5, unchanged.** Clause 5 is graded only by a reading that returns zero, and no
reading was taken in this iter. **The reading series is NOT affected**: no repair here was scored against
the instrument, so the raw series stays discontinuous-at-iter-86 and the adjudicated series stays
untouched. The next paired reading will be the first taken at `0c91421`, and it must say so.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this was a tik) — (3) re-scope: **n, graded
explicitly against the recorded condition** — occurrence 3 landed mid-iter, but the trigger requires TWO
**CONSECUTIVE** invalidated attempts and 33 clean iters separate it from occurrence 2, which already fired
at iter-53 and whose prescribed remedy (TOK-04's pinning-and-tracking policy) performed on this very event
(`D-M257x-87-2`) — (4) user-blocker: n — (5) cap-reached: n (1 tik this session) — (6) protocol-stop: n —
**Outcome: continue**
**Decisions:** `D-M257x-87-1` (clone-advance rule) · `-2` (re-scope grading + the stale `state.md`) · `-3`
(M810 sweep as side-deliverable) · `-4` (family RED headline) · `-5` (hand-off figures re-derived) · `-6`
(the time bomb retired)
**Side-deliverables:** the **M810 sweep** — platform M810 has already landed for `jobsimulation`
(`6092c6d2` destroyed the ECS service/task-def/ECR) while `cms` has not moved; the corpus asserted it as
future work in ~14 passages across 11 files. Landed as Fate 1 with its own commit rather than routed,
because the map had already been corrected and half-repairing is worse than not repairing (§5 rule 19).
Does not upgrade the close status. Also: `state.md` repaired after **73 iters** of drift.
**Routes carried forward:**
- `stack-demo/rosetta-extensions` pin advance (+34) → **Fate 2**, clause-1/clause-2 bring-up work. Not a
  citation target; governed by §7 rule 4.
- `roadrunner`'s prod state (`service_desired_count = 1` against a "folded" claim) and `customerio-sync`'s
  prod terraform → **Fate 3**, a future iter. Both need the **`infrastructure`** repo, which has never
  been in any clone set; the map records them as open questions rather than as claims.
- The `M710` vs `M810` disagreement on which milestone drops the legacy `jobsimulation` schema
  (`hiring.md:33` and `platform-alignment.md` say M710; the terraform comment and the map say M810) →
  **Fate 3**, a separate predicate, flagged not adjudicated.
- `platform-alignment.md:1200`'s citation to `backend.md:175` was **already stale before this iter** →
  **Fate 3**.
**Lessons:** two, both generalised into the protocol doc in this commit.
1. **§5 rule 41** — a check that resolves against a remote-tracking ref is only as current as your last
   fetch, and it reads GREEN until you fetch. Fetch every clone when you measure; advance a checkout only
   when a derived set reads the tree. A large "behind" count is not a large repair: `app` at 93 commits
   and 65 citations surfaced as **2** RED anchors, because they were already being graded at origin HEAD.
2. **§5 rule 42** — a RED summary must name the event, and "the last line of the output" does not. State
   how many, show the first, derive both from the producer's ordering.

And one worth carrying that is not a rule: **the milestone's own state file drifted 73 iters and fed a
wrong number into a run brief.** The re-scope trigger had fired 34 iters before anyone reading `state.md`
would have known. An orchestration file is a claim like any other, and nothing was fencing it.
