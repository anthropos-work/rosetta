**Type:** tik — under `TOK-08`, testing whether iter-184's newly-named class has a second member. See
[`platform-alignment.md` §8](../../../../../corpus/ops/platform-alignment.md).

# iter-185 — `go.mod:14-18` is cited 12 times and the enumerator could not see it

## Phase A — census

iter-184 named the class *a fence's POPULATION is a registry too, and it is the one nobody audits*. Swept
`stack-core` for **population-defining** literals rather than predicate literals.
`predicate_enumerator.CITATION_RE` is one: thirteen file extensions, typed, deciding what the enumerator
can see at all. Measured over `corpus/**.md` + `CLAUDE.md`, in file context:

| extension | file-context citations | was declared? |
|---|---|---|
| **`mod`** | **51** | **no** |
| **`jsx`** | 20 | no |
| `hcl` · `gitignore` | 6 · 6 | no |
| `example` · `ini` · `dev` · `txt` | 5 · 4 · 3 · 2 | no |
| `vue` · `css` | 1 · 1 | no — **found by the fence, not by this census** |

**`app/go.mod:14-18` alone is 12 of the 51** — the anchor CLAUDE.md's shared-libraries banner rests on,
the claim about which org modules `app` actually requires. Any predicate anchored there sat outside the
reach denominator `TOK-07` / iter-114 made a declared thing (`D-M257x-185-1`).

## Phase B — deriving was tried, and refuted

iter-184's rule prefers derivation, so derivation was attempted before the tuple was touched. A `/` in
the stem does **not** separate a file from an authority — `.anthropos` has **14** with-slash hits, every
one from an `https://` URL — and even excluding `://` leaves `api.clerk.com:443` and
`backend.internal.anthropos:8083`. Resolution on disk would be ground truth but needs the clone set,
which this enumerator does not take.

So iter-184's **fallback clause** applies verbatim, and this iter is its first exercise
(`D-M257x-185-2`).

## Phase C — fence, both directions

- **cited ⇒ declared** — RED, naming the extension with example tokens.
- **declared ⇒ occurs** — RED. This is the direction that let iter-184's own fence carry `PROBE` and
  `TASK`, two route kinds that had never existed.
- **the carve-out is not a blanket** — no tail may be both a file extension and an authority, and the
  authority list may not outgrow the class it carves out of. A reason list is how a blanket exclusion
  disguises itself.
- the concrete miss pinned: `app/go.mod:14-18` and a `.jsx` path match; `api.clerk.com:443` does not.

**The arms found more than the census that motivated them** — the census used a ≥2 threshold, the fence
has none, so it immediately named `vue`, `css` and `de`. Each was dispositioned individually: two are
files, one is the Hetzner Storage Box SSH port in `db-backup.md`, named with its reason (`§5` rule 8).

## Phase D — measure, and publish the delta

The iter's own escalation condition was *if this changes what the enumerator finds, publish it*:

| | old class (13) | new class (24) | delta |
|---|---|---|---|
| citation tokens matched | 1,276 | **1,376** | **+100 (+7.8 %)** |
| corpus lines carrying ≥1 citation | 978 | **1,041** | +63 |

Any reach percentage published against a corpus-derived citation denominator was computed over the
smaller one.

| gate | result |
|---|---|
| `test_predicate_enumerator`, both runners | **34 / 34** under `unittest` 3.9.6 **and** 3.14.6 (was 29) |
| the arm, RED-proven pre-repair | **yes** — named 8 extensions, then 3 more once the tuple was extended |
| touched scope + both registry guards + the iter-183/184 fence | **113 passed · 0 failed** (119.82 s) |
| new `*_guard.py` | **0** — arms went into the existing test module; the README triple does not move |

Not covered: the rest of `stack-core` (1,594 P at iter-183), the 7 batteries, the four other rext
sections.

## Close — 2026-08-09

**Outcome:** iter-184's class has a second member and it was costing reach. `CITATION_RE` declared 13
file extensions and the corpus line-pins **11 more** — `mod` at **51 citations**, of which
**`app/go.mod:14-18` is 12**, the anchor under CLAUDE.md's shared-libraries banner. Class extended,
**fenced both directions**, authority carve-out named and itself fenced against becoming a blanket, and
the reach delta published rather than absorbed: **+100 citation tokens, +7.8 %**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (seventeenth consecutive `closed-fixed`; **no
`P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n — (5)
cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-185-1` … `D-M257x-185-4` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none.

**Routes carried forward:**
- `SURVEY-M257x-iter185-other-declared-populations-unaudited` — **NEW.** Two members of iter-184's class
  are now found and repaired, both by hand-sweeping for population literals. **Nobody has enumerated
  how many exist.** A population literal has no syntactic marker distinguishing it from a predicate
  literal, which is exactly why the class is invisible; the sweep that found these two was judgement.
- `SURVEY-M257x-iter185-reach-percentages-predate-the-citation-widening` — **NEW, with its number.**
  The citation denominator moved **1,276 → 1,376 (+7.8 %)**. Any published reach ratio computed against
  a corpus-derived citation denominator predates it. Cousin of
  `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix`; the defective subset is not
  regex-decidable for the same reason.
- `SURVEY-M257x-iter184-the-standing-queue-wildcard-is-unbounded` ·
  `SURVEY-M257x-iter183-only-ONE-registry-property-is-asserted` (half closed) ·
  `SURVEY-M257x-iter183-segment-grammar-refuses-28-multi-id-segments` ·
  `SURVEY-M257x-iter179-thirty-battery-tests-unrun` (owner: the next harden pass) ·
  `SURVEY-M257x-iter180-relation-grammar-supports-only-equality` · `FIX-M257x-iter173-ledger-denominator` ·
  `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` ·
  the observed half of `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` ·
  `SURVEY-M257x-iter178-n-of-m-outside-clause-5-is-52-sites` — unchanged; open.
- The standing queue, unchanged.

**Lessons:** **a fence with a lower threshold than the survey that motivated it is doing its job** — the
census used ≥2 and the fence has none, so it named three more the moment it ran. Write the assertion
against *all* of the population, then disposition the tail; do not carry the survey's threshold into the
fence, because a threshold is a sampling decision and a fence is a census. And the corollary that made
this iter cheap: **try the derivation first and record its refutation** — iter-184's rule prefers
deriving a population, and here deriving genuinely does not work (a `/` does not separate a file from an
authority). *Recording why* is what makes the fallback a decision rather than laziness. Written into
`platform-alignment.md` §8 in this iter's commit.
