# iter-213 — decisions

## `D-M257x-213-1` — `classify()` strips route ids before grading

One line: `segment = ID_RE.sub(" ", segment)`. The function's own docstring already claimed it graded
what the segment *records*; it was handed the segment with the ids in it and `REOPEN_RE` /`PARTIAL_RE`
match ordinary English words (`retract`, `refut`, `supersede`, `arm`, `half`, `slice`). The repair makes
the docstring true rather than changing what it promises.

**Deliberate consequence, fenced:** a verdict word glued INTO a slug
(`FIX-M999-iter01-thing-CLOSED`) is stripped with the id and grades `other`. That is correct — **a
route may not smuggle its disposition into its name** — and it is asserted, not assumed.

## `D-M257x-213-2` — the exposure is 1, not 14, and this iter's own S4 was wrong

S4 pre-registered *"every bare re-listing supplies its own excuse."* **Falsified by this iter's own
first staged control**, which went RED on the plain-slug leg: a bare re-listing grades `other` both
before and after, and `other` never triggered the contradiction rule — which fires on **`open` after
`closed`**. Partition of the 14 changed segments:

| transition | n | what it moves |
|---|---:|---|
| `reopen → other` | 13 | the census **display** only |
| `reopen → open` | **1** | the **rule** — the only segment that could ever have been suppressed |

The one is `FIX-M257x-iter177-ledger-carries-a-retracted-retraction` at iter-178 — *"unchanged; owner =
the next harden pass"*, a route carried **unchanged** and booked as **re-opened** because its slug says
`retracted-retraction`. It was never closed beforehand, so **no contradiction was actually suppressed**:
`violations()` is 0 before and 0 after. `§5` — *grade at the grain of the claim*; the headline is the
mechanism, and the live exposure is **1 segment, latent**.

The falsification is kept as an executable arm
(`test_a_BARE_relisting_is_other_for_BOTH_slugs_which_is_why_the_arm_above_uses_open`) paired with the
control it corrected, so the two cannot drift apart.

## `D-M257x-213-3` — the probe methodology, corrected mid-iter

The first measurement of S2 returned **40 events across 26 ids** — **~2.9× overstated**. It compared a
verdict computed on the FULL segment against one recomputed on the **240-character truncation** stored
in `events[rid]`. Same class as iter-209's hand-written slugger: **when the question is about the
INPUT, hold the machinery fixed.** The corrected probe re-walks the sources and classifies both ways on
one string; every figure published by this iter comes from that probe.
