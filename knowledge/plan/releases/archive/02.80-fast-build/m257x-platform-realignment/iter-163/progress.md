**Type:** tik — under [`TOK-08`](../decisions.md): *intra-corpus mis-citation is mechanical; build a
fence that enumerates every instance, run it to zero, keep it green.*

# iter-163 — a rot detector keyed on a delimiter measures luck

## Phase A — the predicate, sharpened three times against its own output

**Draft 1 — 346 "findings", and none of them a measurement.** Pairing every backticked literal on a
line with every citation on that line is a cross-product. `CLAUDE.md:332` alone contributed 17 rows:
one mega-line, 5 citations, 12 literals. **The number was an artifact of the pairing, not of the
corpus** — the same shape iter-159 had to break to get from 2,854 to 961.

**Draft 2 — pairing discipline: 24.** Exactly one citation on the line; exactly one literal, the
nearest, within 60 characters; ≥ 12 chars; not itself a path.

**Draft 3 — the anchor names a CONSTRUCT, not a line: 16.** A citation of a function declaration is
not wrong because the quoted call sits in its body. Reusing `anchor_construct_guard._block_bounds`
absorbed 8 of the 24 — every one of them correct.

**Denominator, stated (iter-114's rule):**

| stage | count |
|---|---|
| single-citation lines in the corpus | **540** |
| …that resolve, at the ref their own block names | **442** |
| …carrying a quoted literal beside the citation | **193** (249 carry none) |
| …whose literal occurs **somewhere** in the target — *adjudicable* | **137** (56 absent ⇒ undecidable) |
| …surviving all three acceptance clauses | **16 → graded at source** |

## Phase B — all 16 graded at source, and they split four ways

| verdict | n | what it means |
|---|---|---|
| **REAL — repaired** | **4** | the anchor does not carry its subject |
| **exempt, declared with a reason** | **9** | the match is real and is not a defect |
| **absorbed by a clause the grading justified** | **3** | 2 range-siblings + 1 sentence boundary |

### The four repairs — each re-derived against the SUBJECT, never by bumping the offset

| site | was | is | why |
|---|---|---|---|
| `platform-alignment.md:229` | `docker-compose.yml:43` | **`:18`** | `:43` is `- "8083:8083"`, a port. `search_path=sentinel` is in the sentinel `DB_CONNECTION` at `:18` |
| `platform-alignment.md:1054` | `setup_guide.md:486` | **`:504`** | `:486` is a `psql` command; the `migrations: true` enumeration is 18 lines later |
| `platform-migration-status.md:161` | `secrets-spec.md:309` | **`:344`** | `:309` is an unrelated API-key table row; the `../hyper-studio/.env.example` borrow is at `:344` |
| `ai-readiness.md:28` | *two names, three anchors* | **three names** | **the anchors were right and the prose was short a name** — `:349` is `LoadMembers`, `:353` `LoadMembersByUserIDs`, `:357` `BaseMembers` |

The last one is worth its own line: **an anchor census found a missing enumerand.** The instrument was
built to catch a citation pointing at the wrong line; what it caught here is prose that named two
constructs and cited three. Nothing about anchors was wrong.

### The three clauses the grading earned — each with its live instance

1. **Sibling anchors include RANGES.** *"read at `cors.go:24` and applied at `:78-82` under
   `if !environment.IsProduction()`"* — the corpus attributes the literal explicitly. **Both live
   instances are ranges**, so a `` `:(\d+)` ``-only rule caught neither.
2. **Do not pair across a sentence boundary.** *"…the `flag_use_azure_us` caveat below.
   [`external_services.md:567`] carries the…"* — a full stop is the corpus's own statement that the
   subject changed. Proximity in characters is not proximity in argument.
3. **A positional anchor list.** ``members.go:349`/`:353`/`:357`` — N names, N lines, matched by
   order. Ordering is precisely what this census does not read, so any line in the run is a
   legitimate home. **One live instance, and it is stated as one** rather than generalised.

**No clause was tuned until a known instance fired.** The `_block_bounds` under-reach (2 sites) could
have been "fixed" by widening a window by 2 — that is Trap A, and instead both are **declared
exemptions naming the helper's defect**, with the helper fix routed.

## Phase C — zero

```
anchor-subject-census: 0 unexempted finding(s) (9 graded exempt) over 137 adjudicable pairs —
from 540 single-citation lines, 442 resolved, 193 carrying a quoted literal
(249 without one, 56 whose literal is absent from the target and so undecidable)
```

## Phase D — the fence

17 tests. Every clause in **both** directions — including the one that matters most, *the same
literal WITHOUT the full stop still pairs*, because a sentence clause that over-reaches silently
disables the census. Plus the live ratchet, an **exemption-quality** gate (a declaration with no
reason is not one), a **stale-exemption** gate (an exemption matching nothing is fiction carried as
reassurance), and two anti-vacuity controls (empty root ⇒ `CANNOT RUN`, CLI exit **2**, never 0).

One defect in this iter's own work, caught by its own fence: the module located the corpus root by
counting `..` from its own file — **the hand-maintained tuple again**, one iter after iter-162. It
now walks up for the two things only the corpus root has.

## Gates

- `test_anchor_subject_census_m257x` — **17 passed, 0 failed** (net-new).
- `test_repair_postcondition` (live-tree grading over the 4 repaired docs) — **28 passed, 0 failed**.
- `test_anchor_construct_denominator` + `test_anchor_offset_guard` — **54 passed, 0 failed**.

**NOT re-run, named in full (`§5` rule 60):** the `stack-core` suite in full (~20–35 min — this iter
ADDED two files and modified no existing rext file), and **demo-stack, dev-stack, stack-injection,
stack-verify, stack-seeding, stack-snapshot, stack-secrets, alignment, playthroughs, clerkenstein** —
untouched.

## Close — 2026-08-08

**Outcome:** the decidable slice of *"the anchor resolves but names the wrong thing"* now has an
instrument, a denominator and a zero. **540 → 442 → 193 → 137 adjudicable pairs; 16 graded at
source; 4 real defects repaired; 9 exempt with reasons; 3 absorbed by clauses the grading itself
justified.** The four repairs were each re-derived against the subject, and one of them was not an
anchor defect at all — the prose named two constructs and cited three.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**no `N` reading taken, so the metric is
UNMEASURED not unmoved (`§9`); a successor strategy remains FORBIDDEN by `TOK-08`'s sealed rule**) —
(3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7)
budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-163-1` … `D-M257x-163-3` (see [`decisions.md`](decisions.md))
**Side-deliverables:** none.
**Routes carried forward:**
- `FIX-M257x-iter138-anchor-rot-fence` — **CLOSED for the decidable slice.** The semantic remainder
  (an anchor with no quoted literal beside it — 249 of 442 resolved lines) is explicitly out of reach
  and stays routed as `SURVEY-M257x-iter163-anchors-with-no-quoted-literal`.
- `FIX-M257x-iter163-block-bounds-under-reaches-by-two` — **NEW.** `_block_bounds` closed a Go
  function 2 lines early at both `askengine.md:111` and `skillpath.md:80`. Two graded exemptions
  exist only because of it.
- `FIX-M257x-iter163-block-ref-attaches-the-wrong-sha` — **NEW.** When a block names several shas,
  `block_ref` picks one; at `frontend_architecture.md:39` it picked a two-releases-old compose for a
  citation that meant HEAD. One graded exemption exists only because of it.
- `FIX-M257x-iter163-anchor-guard-does-not-know-shell-keywords` — **NEW.** `up-injected.sh:2494` is a
  bare `fi`; the closing-delimiter clause knows `}` and `)` and no shell keyword, so a content-free
  shell anchor reads as content.
- `SURVEY-M257x-iter163-generic-literals-are-unadjudicable` — **NEW.** `MIN_LEN` is a weak proxy for
  specificity; `jobsimulation` passes it and proves nothing.
- Unchanged and still queued: `SURVEY-M257x-iter162-a-literal-has-a-ROLE-the-census-cannot-see` ·
  `SURVEY-M257x-iter162-small-derivations-are-coincidence-prone` ·
  `SURVEY-M257x-iter161-golden-vs-frozen-needs-input-provenance` ·
  `SWEEP-M257x-iter159-grade-the-961-haystack-candidates` ·
  `SURVEY-M257x-iter160-inexact-copies-are-invisible-to-an-equality` ·
  `FIX-M257x-iter160-b2-over-strict-direction-still-unfenced` ·
  `SURVEY-M257x-iter158-noise-classifier-is-narrow-by-choice` ·
  `SURVEY-M257x-iter156-other-reporting-layers` · `SURVEY-M257x-iter152-half-up-services-are-ungradeable` ·
  `FIX-M257x-iter155-add-injected-rows-vs-the-pinned-liveness-only-decision` ·
  `FIX-M257x-iter145-sha-baseline-drift` · `-iter145-migrate-race-needs-a-host-postgres` ·
  `-iter145-green-but-stale-graphql-mentions` · `-iter143-wrong-head-is-unfenced` ·
  `-iter143-scope-derivation-by-grep` · `-iter143-appending-to-the-protocol-doc-rots-the-ledger` ·
  `-iter144-correction-vs-retraction-unfenced` · `SURVEY-M257x-iter144-orphan-arm-is-the-residual` ·
  `FIX-M257x-iter142-path-arm-window` · `-iter142-value-change-articles` · `-iter142-tier-b-underflag` ·
  `FIX-M257x-iter135-adjudicated-live-defects` · `-iter140-receipts-not-checkable-here` ·
  `-iter140-receipt-fence` · `-iter134-fence-family-has-no-shared-predicate-layer` ·
  `-iter133-two-fives-need-a-fence` · `-iter131-predicate-sets-not-enumerated`
**Lessons:** **when a class is "too semantic to fence", look for the slice that carries its own
evidence.** Nobody could fence *"does this sentence's claim match this line"* — but the corpus
routinely quotes the literal it is talking about, and a quoted literal turns an interpretation
problem into a lookup. The reach is honest precisely because the denominator says how much it does
not cover: **137 adjudicable of 442 resolved.** Second, and sharper: **the pairing is the instrument.**
Three of this iter's four sharpenings changed nothing about *what* is compared and everything about
*which two things* get compared — 346 → 24 → 16 → 0, with no clause ever tuned until a known instance
fired.
