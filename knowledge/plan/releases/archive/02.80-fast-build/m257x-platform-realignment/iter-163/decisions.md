# iter-163 — decisions

## `D-M257x-163-1` — the class is not "too semantic to fence"; the SLICE THAT CARRIES ITS OWN EVIDENCE is fenceable

`anchor_construct_guard` closed this question in its own docstring: catching an anchor that lands on
the wrong *ordinary code* "requires deciding what a sentence claims, which is the line this whole
fence family does not cross." That is correct as stated, and it stopped the class dead for eighteen
iters.

**What it missed is that the corpus usually quotes the thing it is talking about.** When a backticked
literal sits beside a citation, the question stops being *what does this sentence claim* and becomes
*is this string at that line* — a lookup. The reach is partial and the census says so with a
denominator on every run: **137 adjudicable pairs of 442 resolved single-citation lines.** The
remaining 249 carry no quoted literal and are explicitly out of reach.

**Not claimed:** that anchor rot is now fenced. What is fenced is the slice where the corpus supplied
its own evidence. The semantic remainder is routed under its own name, not folded into this zero.

## `D-M257x-163-2` — the PAIRING is the instrument, and three of four sharpenings only changed which two things get compared

The first draft printed **346 findings** and every one of them was arithmetic: pair every literal on
a line with every citation on that line and a mega-line with 5 citations and 12 literals contributes
17 rows. **346 → 24 → 16 → 0**, and across those steps *what* is compared never changed. What changed
was the pairing:

- exactly **one** citation on the line, exactly **one** literal — the nearest, within 60 chars;
- the anchor names a **construct**, so the enclosing block counts (`_block_bounds`), not a ±window;
- the corpus's **own** attribution wins — `` `:97` `` *and* `` `:78-82` ``, and a positional run
  ``:349`/`:353`/`:357``;
- a **full stop** between the two ends the pairing, because a sentence boundary is the corpus stating
  that the subject changed.

**Every clause has a live instance and is tested in both directions**, including the direction that
matters most — *the same literal WITHOUT the full stop still pairs*. A sentence clause that
over-reaches would silently disable the census while leaving it green, which is the failure mode this
milestone has now found in six instruments.

**And no clause was tuned until a known instance fired.** The two `_block_bounds` under-reaches could
have been absorbed by widening a window by exactly 2. That is Trap A. They are declared exemptions
naming the helper's defect instead, and the helper fix is routed.

## `D-M257x-163-3` — repairs are re-derived against the SUBJECT; one of the four was not an anchor defect at all

`§5` rule 22 (iter-22): re-derive the correction, not just the anchor. Applied literally — none of
the four repairs was an offset bump:

- `docker-compose.yml:43` → **`:18`**: `:43` is a published port; `search_path=sentinel` lives in
  sentinel's `DB_CONNECTION`.
- `setup_guide.md:486` → **`:504`**: `:486` is a `psql` invocation; the `migrations: true`
  enumeration is 18 lines below it.
- `secrets-spec.md:309` → **`:344`**: `:309` is an unrelated API-key row; the
  `../hyper-studio/.env.example` borrow is at `:344`.
- `ai-readiness.md:28`: **the anchors were correct and the prose was short a name.** It read
  *"`LoadMembersByUserIDs` / `BaseMembers` are at `:349`/`:353`/`:357`"* — two constructs, three
  lines. `:349` is `LoadMembers`, which the sentence never mentioned. The repair adds the name.

That fourth one is the decision worth recording: **an instrument built to find rotted anchors found a
missing enumerand instead.** A census that compares two things will surface disagreements of a shape
its author did not have in mind, and grading every survivor at source (iter-158) is what lets that
show up as a finding rather than as a mechanical "bump the offset" repair — which, here, would have
been a *correct-looking* edit that destroyed a true anchor and left the real defect in place.
