---
iter: 234
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
---

# iter-234 — do the corpus's `file:NN` anchors LAND on the text they quote?

## Step 0 — Re-survey (mandatory before targeting)

Three iters have walked up the citation ladder and each closed the rung below this one:

| iter | rung | verdict |
|---|---|---|
| 230 | do the corpus's **shas** exist? | 132/132 measurable resolve |
| 231 | do its **current-ref claims** agree? | 4 agree · 1 unmeasurable · 1 stale |
| 232 | do its cited **paths** still exist? | 0 of 111 gone |
| 233 | is the **clone set** those readings run against healthy? | 15 clones · 4 flagged · 0 broken |

Every one of those grades a **container**: a sha, a path, a repo. **None of them has ever asked whether
the cited LINE carries the text the corpus says it carries.** `cite224.py` (iter-224) came closest and
stopped at *"is the file long enough for line NN to exist"* — a file-length check, not a content check,
and it ran over the six **archived** repos only, not the ones a stack builds from.

So the rung is open, it is the one the milestone exists for — *is the corpus describing the platform as it
is?* — and it is **mechanically decidable for a real subset**: wherever the corpus writes a `file:NN`
anchor **and** a backticked literal in the same breath, the literal either appears at line NN or it does
not. No sentence has to be interpreted. That is `TOK-08`'s co-quotation shape exactly.

**Active strategy reference:** `TOK-08` — census a mechanical class exhaustively; a reading samples, a
fence censuses.

**Scope follows the user's redirect** (corpus claims about the platform + a working stack, not the
instruments that grade them): the population is the **live** platform repos a stack actually builds from —
`app`, `platform`, `sentinel`, `next-web-app`, `studio-desk`, `ant-academy` — read at the **`stack-demo`
clone set**, which is *the tree a demo builds* (harden pass 54: reading at `origin/main` graded a tree
nothing builds and flipped the verdict).

## The two clocks, and both get read

Harden 54 established that a corpus claim only means something against the tree that gets built. But many
of these sites **name their own sha in the window** (`@ 2035f9a4`, `at b948604f`). Those two are different
questions and this iter grades both:

- **clone-HEAD reading** — does the anchor land on the tree a `demo-up` would build *today*?
- **stated-sha reading** — where the site names a ref, does the anchor land *at the ref it names*?

A site that lands at its own stated sha and misses at clone HEAD is **correct and dated**, not wrong. A
site that misses at *both* is a defect. Collapsing the two is the failure mode iter-69 named (*a pin is a
date*).

## Hypothesis

Corpus anchors were written over ~8 months against a moving `app`; line numbers are the most fragile thing
a document can carry. The exact-hit rate should be high but **not** 100 %, and the misses should split into
**offset drift** (the literal is in the file, at a different line) and **content mismatch** (the literal is
not in the file at all) — only the second class is an alignment defect.

## Predictions — SEALED BEFORE MEASUREMENT

| id | prediction |
|----|-----------|
| `P-234-1` | ≥ 150 corpus sites carry a `file:NN` anchor into one of the six **live** repos |
| `P-234-2` | ≥ 60 of those are **co-quoted** — a backticked literal in the same corpus line — and therefore mechanically gradable |
| `P-234-3` | the exact-LANDS rate over the gradable set, read at the clone tree, is **≥ 70 %** |
| `P-234-4` | ≥ 10 gradable anchors **MISS** at the cited line (offset drift is real, not hypothetical) |
| `P-234-5` | ≥ 1 MISS is a **content mismatch** — the quoted literal is nowhere in the cited file — i.e. a claim about the platform that the platform does not carry |
| `P-234-6` | ≥ 1 site that MISSES at clone HEAD **LANDS at the sha it names**, proving the two clocks are not interchangeable |

## Expected lift

No `N`/`P` reading is claimed — this iter takes no graded seat. Deliverable: the gradable-anchor
population with its denominator, the LANDS/NEAR/MISS partition read at **both** clocks, every MISS
classified drift-vs-mismatch, and repair of any content mismatch found in a live present-tense claim.

## Phase plan

1. Derive the anchor population from `corpus/**` + `CLAUDE.md` (strict regex; the `app/internal/cms/...`
   false-positive class iter-224 already characterised must stay excluded).
2. Partition gradable (co-quoted) vs ungradable, and **report the ungradable count** — a census that hides
   its own reach is not a census (`§8`).
3. Grade at clone HEAD and at each site's stated sha where one exists.
4. Prove the instrument: the census must return a **non-zero** on a control it cannot have fitted.
5. Classify every MISS; repair content mismatches in live claims; route the rest.

## Escalation conditions

- Instrument returns 0 gradable sites → the regex is the finding, not the corpus (`§9`: a census returning
  ZERO must prove its instrument).
- A content mismatch turns out to need a platform edit → route, never edit the platform.

## Acceptable close-no-lift outcomes

A measured 100 % LANDS rate with a proven-non-vacuous instrument is a complete iter: it closes the class
and refutes `P-234-4`/`P-234-5` on the seal, which is the mechanism working.
