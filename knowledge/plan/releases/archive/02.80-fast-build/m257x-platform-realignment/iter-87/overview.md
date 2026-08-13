---
iteration_type: tik
status: archived
opened: 2026-08-05
---

# iter-87 — the platform moved again: absorb `0c91421` in the iter that detected it

**Type:** tik, under [`TOK-05`](../decisions.md#tok-05-stop-repairing-claims-fence-the-predicates-under-them--2026-08-04).

The shape is not chosen — it is prescribed. TOK-04 P3 / §5 rule 26: **"the iter that detects the move
re-points it, in that iter."** iter-54 absorbed a three-commit move in under an hour and recorded that
*"the expense is in deferring it, not in doing it."*

## Step 0 — re-survey (mandatory), and what it overturned

Gate **4 of 5** at open. `rosetta ae5c1db`, `rext ac30b9b` (authoring copy, on `main`), both clean,
both local == remote.

**The platform had moved, and our ground truth was pre-drift.** `stack-demo/platform` sat at `0dab54d`,
**2 commits behind** `origin/main`:

```
0c91421 2026-08-05T16:21:19+02:00  Merge PR #26 chore/drop-support-service-containers
838d907 2026-08-05T16:14:25+02:00  chore(compose): drop the storage, messenger and customerio-sync containers
```

Advanced to `0c91421` (fast-forward, clone left clean, **nothing committed into it** — re-pointing our own
working copy is not a platform edit). Full derivation in [`ground-truth.md`](ground-truth.md).

| signal | `0dab54d` | `0c91421` (origin HEAD) |
|---|---|---|
| `repos.yml` entries | 6 | **4** — `storage`, `messenger` removed |
| compose services declared | 8 | **5** |
| effective topology (with `include:`) | 10 | **7** |
| profiles present | 8 | **5** — `storage-legacy`, `messenger`, `customerio-sync` all gone |
| `core` selects | 5 | **5** (unchanged) |
| `STORAGE_S3_BUCKET` on `backend` | `:82` | **`:82` — persists** |

Every figure in the hand-off re-derived and **confirmed**. One hand-off figure did **not** survive, and it
is the load-bearing one — see the next section.

## The three claims this iter tests, and how each came out

**1. Does `platform_alignment_guard` go RED on the `repos.yml` membership change?** The hand-off named
this the decisive question, and said a fence that failed to fire would be "MORE IMPORTANT than the
repair." **It fired.** Assertion B, **2 for 2**, unaided, naming both departures in the guard's own voice:

```
[B departure] the map claims messenger is in repos.yml, and it is not — a service left the clone set
              and the map still asserts it
[B departure] the map claims storage is in repos.yml, and it is not — …
```

Run with a **control**: a detached worktree at `0dab54d` in the same tree gives **0** assertion-B
findings; the advanced checkout gives **2**. The delta is the event, not accumulated debt. This is the
fence's **second** unaided catch of a live membership change (§5 rule 27 records the first, 3 for 3).

**2. Is the hand-off's "13 GREEN · 0 RED · 3 not-run" reproducible?** **No — and the reason is a new
mechanism, not an error of counting.** Measured at the *identical* `0dab54d` checkout, after a `git
fetch`: **10 GREEN · 3 RED · 3 not-run**. The three REDs are the citation-resolving guards
(`anchor_construct_guard`, `platform_alignment_guard` assertion F, `repair_postcondition`), and by the
iter-68 `CITE_REF=auto` contract they resolve every citation at **`origin/main`**, not at the checkout.
So they were armed the moment `git fetch` landed `0c91421` on the remote-tracking ref — **the fetch, not
the checkout, is what turned them RED**. An unfetched clone makes a citation fence read GREEN. New
protocol rule; see the close.

**3. Does the re-scope trigger fire?** Graded explicitly below. **It does not** — and the count it is
graded against is **not** the one the hand-off carried.

## The re-scope trigger — graded, not left open

The hand-off (via `state.md:6`) says the trigger stands at **occurrence 1 of 2**. **`state.md` is stale by
33 iters.** The milestone's own decision record is authoritative and says otherwise:

- `iter-53/decisions.md:44` — **`D-M257x-53-6`: "Occurrence 2 of 2. `EXIT_REASON` corrected from
  `user-blocker` to `re-scope-trigger`."** The trigger **already fired**, at iter-53.
- `decisions.md:665` (TOK-04) — *"`re_scope_trigger`, occurrence 2 of 2, plus a direct user ruling."*

So the real history, derived from the platform's own commit dates:

| # | date | commit | outcome |
|---|---|---|---|
| 1 | 2026-07-31 | `2adcf71` WunderGraph drop | recorded (iter-12) |
| 2 | 2026-08-03 | `ef32d4c` prune merged services | **TRIGGER FIRED** (iter-53) → escalated → remedy = **TOK-04**, a pinning-and-tracking policy |
| 3 | 2026-08-05 | `0c91421` drop support containers | **this iter** |

**Grading: NOT FIRED.** The recorded condition is *"**TWO CONSECUTIVE** full-alignment attempts are
invalidated."* Between occurrence 2 and occurrence 3, **33 iters ran with no platform movement at all**
(no commit between 2026-08-03 and 2026-08-05). Two invalidations separated by 33 clean iters are not
consecutive, so the predicate is false on its own words.

And the second half, which matters more than the arithmetic: **the trigger's prescribed remedy already
exists and is demonstrably working on this very event.** The trigger says the answer is *"a
pinning-and-tracking POLICY (how we choose a platform ref, how we notice it moved, who re-points), not
more alignment work."* That policy was built as TOK-04 P1/P2/P3 after occurrence 2. On occurrence 3 it
performed: the move was **noticed by a fence within hours**, the ref is **stated in the artifact**, and
**the detecting iter re-points it** — which is this iter. Firing a trigger whose remedy is already in
place and functioning would be ceremony.

## Cluster / target identified

One planned line with a multi-step shape (TOK-05's predicate unit): **absorb `0c91421` across every
predicate it falsified.** The fences enumerate the set — 38 findings, all of them consequences of the
2-commit move:

| source | findings | predicate |
|---|---|---|
| `platform_alignment_guard` B | 2 | *service X is in `repos.yml`* |
| `platform_predicate_guard` G1 | 3 tokens / **28 sites** | *profile token X selects something* |
| `platform_predicate_guard` G8 | 3 | *service X declares `profiles:`* |
| `platform_predicate_guard` G10 | 1 | *compose declares N services* |
| `platform_predicate_guard` G2 | 1 | *`repos.yml` lists N repos* |
| `platform_predicate_guard` G4 | 9 | *the platform sets `<VAR>` locally* |
| `anchor_construct_guard` + F | 15 tree-wide / 19 in-map | *this line number names this construct* |

Six predicates, not 38 claims. Exactly the unit TOK-05 exists for.

**iter-86's repair is among the falsified.** It repaired the corpus to say `storage` sits in
`profiles: [storage-legacy]` and `messenger` in `profiles: [messenger]`. Both were correct at the ref it
measured and are false at origin HEAD. That is §5 rule 33 — *a ref-pin is a DATE, not an exemption* —
arriving as a live event rather than as doctrine.

## Hypothesis

Advancing the clone and repairing by predicate returns the guard family to its pre-move verdict on the
checkout-dependent assertions (B, G1–G10) and clears the citation drift the fetch armed.

## Expected lift

Guard family from **9 GREEN · 4 RED** to **13 GREEN · 0 RED** (3 not-run without `--range`/`--ledger`).
Gate stays **4 of 5** — clause 5 is met only by a reading that returns zero, and no reading is taken here.

## The clone-advance rule (`D-M257x-87-1`) — decided, and derived rather than preferred

10 of 13 clones are behind origin (`app` 93 — the hand-off's figure, **not** iter-86's stale 60;
`next-web-app` 41, `rext` 34, `storage` 20, `messenger` 7, `ant-academy`/`jobsimulation` 4,
`cms`/`sentinel`/`studio-desk` 2, `graphql-wundergraph`/`roadrunner`/`platform` 0).

> **FETCH ALL; ADVANCE ONLY WHAT A DERIVED SET READS.**
>
> Every clone is **fetched** in any iter that takes a measurement. A clone's **checkout** is advanced only
> when the checked-out tree is an input to a derived legal set or to a build.

Derived, not chosen — the measurement is in the guards' own code:

- `anchor_construct_guard.resolve()` and `platform_alignment_guard.cited_text()` both default to
  `CITE_REF=auto`, whose ladder is **`origin/main` first**, HEAD second (iter-68,
  `FENCE-M257x-iter68-citation-resolution`). A **fetched** clone is therefore already graded at origin
  HEAD *no matter where its HEAD sits*.
- Only two things read the checked-out tree: `platform_predicate_guard` (the dir passed to `--platform`)
  and `platform_alignment_guard` assertion B (the `repos.yml` path passed as argv). Both are `platform`.

**Consequence, and it dissolves the deferral the hand-off anticipated:** advancing `app` by 93 commits
would move **nothing**, because no fence reads `app`'s checkout. The 65 `app/*:N` citations in the corpus
are **already being graded at origin HEAD, in this reading**. The "large wave" is not deferred — it is
*in* the measurement, and it costs **2 anchors** (`app/main.go:471`, `app/main.go:637`), not 65.

**What is deferred under the rule:** nothing citation-bearing. The `stack-demo/rosetta-extensions`
consumption clone (+34) is a **pin**, not a citation target; advancing it is a bring-up act governed by
§7 rule 4 and belongs to a clause-1/clause-2 iter. Routed as Fate 2 (already covered — clause 1/2 work).

## Phase plan

1. Advance the platform clone; re-derive ground truth with a control. **(done at open)**
2. Repair by predicate, tree-wide per §5 rule 19 — every site of each of the six, including `CLAUDE.md`
   and `.claude/skills/**`, which rule 19's corollary names as the highest-propagation sites.
3. Re-point the drifted citations (§7 rule 4 second half / `D-M257x-59-3`).
4. Repair `state.md`'s stale re-scope-trigger record.
5. Add the fetch-vs-checkout finding to the protocol doc as a numbered rule.
6. Re-run the guard family; commit; push both repos and verify with `ls-remote`.

## Escalation conditions

- A guard that should have caught the move and did not → that is the finding, and it outranks the repair.
- A third platform commit landing mid-iter → re-grade the trigger (consecutive, this time).
- The carve-out (`storage.md:55/:154/:181`, `DEF-M257x-iter80-storage-prod-bucket`) is **not** touched.

## Acceptable close-no-lift outcomes

A fence proving unable to see the event would be a first-class result and would close the iter on the
finding rather than the repair.
