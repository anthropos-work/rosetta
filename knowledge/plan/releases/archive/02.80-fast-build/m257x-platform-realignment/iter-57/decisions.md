# iter-57 — decisions

## D-M257x-57-1 — the subject rule was REPLACED after measuring its own false-positive rate

The first draft of assertion F asserted *"the subject is named in the cited path or on the cited line."*
Run against the real map it produced **22 findings, of which 7 were the rule's own false positives** —
`app` cites the `backend:` key (an alias), `jobsimulation` cites a line reading `jobsim`, `cms` and
`roadrunner` cite a `repos.yml` header comment that is *about* them without saying so, `postgresql` and
`redis` cite `common.yml`.

A 32% false-positive rate is not a calibration detail. `anchor_construct_guard`'s docstring records the
fate of exactly this shape — 134 findings, "essentially all of them ports" — and a fence that noisy is
disabled on first contact, which is worse than no fence because it also consumes the credibility of the
next one.

**The forbidden repair was the obvious one:** narrow the rule until only the known-bad citations fire.
That is §5 Trap A — tuning a fence to the answer key — and it would have produced a guard that passes
today's map and asserts nothing about tomorrow's.

Instead the rule was replaced with one **the file itself defines**: for `docker-compose.yml`, the cited
line must sit inside the compose block of the row's own service, where block boundaries *and* the
repo→service aliases are parsed out of compose (`context: ${APP_BUILD_CONTEXT:-../app}` is how the guard
learns that `backend` is the `app` repo — it is not told). Findings went **22 → 8, false positives 7 → 0**,
and the rule now describes a property rather than a list.

**Honest note on how the scope was reached:** the first run informed the redesign. What makes the result
defensible is not that it was chosen blind — it was not — but that the replacement is justified by a
*stated structural property* (compose is the one cited file with self-describing, name-keyed blocks, and
it is the file whose line numbers shift on every platform compose edit) rather than by which specific
findings survive it. Every one of the 8 survivors was verified against the platform independently.

## D-M257x-57-2 — an out-of-block citation is legal when the row DECLARES the block

After repair, one finding remained: `roadrunner` cites `docker-compose.yml:59`, which is inside
`backend`'s block. That citation is **correct** — the whole claim is that `JUDGE0_BASE_URL` *moved onto
backend*.

Rather than loosen the rule (tuning), the rule gained a stated clause: a cross-block citation passes when
the evidence cell **names the block it points into**. This cannot launder drift, because the declared name
is compared against the block *compose itself* says owns the line — declaring the wrong service still
fires, and `test_F_declaring_the_WRONG_block_still_fires` is the inverted mutant that holds it to that.

The effect is to convert a silent cross-block citation into a **stated** one, which is the same move P4
asks of prose generally.

## D-M257x-57-3 — `citations=False` exists for one caller and is deliberately not a CLI flag

`RealMapWithoutACloneSet` runs the shipped map against a reference **fabricated from the map's own rows**,
so that assertions C and D stay unskippable on a box with no platform clone. Assertion F is not a map-alone
property — it needs the real clone — and running it against a fabricated one would measure the fabrication.

So `check()` takes `citations: bool = True`, and **`main()` does not expose it**. A fidelity check you can
quietly switch off is one that will be; the milestone has four instances of exactly that shape.

## D-M257x-57-4 — the prose-under-review category is now VISIBLE in the map

TOK-04's P4 records that the map "must carry that third category **visibly** — it currently does not,
which is exactly why a false §5 narrative read as authoritative as a fenced table row."

§4 now carries a standing table: what is derived+fenced (membership, both directions), what is fenced
(vocabulary; 20 compose citations), what is **resolution-and-range only** (28 citations into terraform /
`repos.yml` / `common.yml` / `app/…`, where there is no derivable notion of *whose* line it is), and what
is **prose-under-review** (the prod column, the PR/rollback narrative, §5's ordering).

The guard prints its reach on every run, GREEN or RED, and **refuses (exit 2)** when it subject-checks
nothing — so the coverage claim in the map is one the tool re-states each time rather than a sentence that
can quietly stop being true.

## D-M257x-57-5 — within-block citation drift is a STATED limitation, not a fenced property

The block rule catches **cross-block** drift. A citation that drifts but stays inside its own service's
block still passes — `messenger`'s `:178`, `:159`, `:161` were all real drift and all inside messenger's
own block, so F did not name them. They were repaired anyway (the correct lines were derived), but they
were found by hand, not by the fence.

This is written into the guard's docstring rather than left for a future reader to discover. The class the
fence closes is the one that fired: compose is edited, every line below shifts, citations slide into the
neighbouring service.

## D-M257x-57-6 — rext is tagged and pushed; the consumption pin is NOT advanced

`fast-build-m257x-iter-57` is cut and pushed to origin — per the root `CLAUDE.md` rule that **tagging is
not publishing**, and the M236 precedent where 0 of 13 tags on origin cost an entire iteration.

`.agentspace/rext.tag` stays at `fast-build-m257x-iter-56`, deliberately:

- The pin controls **what a bring-up consumes**. `demo-1` is up and IS clause 1's and clause 2's evidence,
  measured at that pin.
- Assertion F is a corpus-fidelity guard. **No bring-up path consumes it** — it runs from the authoring
  copy and from CI. Advancing the pin would change the bring-up's inputs for zero benefit to this
  iteration's deliverable, while making the live evidence describe a configuration that no longer exists.

The pin advance belongs to the iteration that next takes a cold cycle, which is where its effect can
actually be observed.
