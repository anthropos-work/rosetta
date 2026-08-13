# iter-161 — decisions

## `D-M257x-161-1` — a golden over a SYNTHETIC input is not a frozen copy of a LIVE derivation

This iter's `overview.md` graded `test_platform_predicate_guard.py:193` as a literal asserted *against
the live platform*, and declared that reading a hypothesis. At source it is
`self.platform = write_platform(self.root)` — **a synthetic compose the test writes into a tmpdir.**

The distinction is the whole boundary of the class:

- a literal that duplicates what a **live** derivation returns is a **frozen copy** — the platform moves,
  the copy does not, and the test goes RED for something it is not about;
- a literal that states the expected output of an input **the test itself constructs** is a **golden** —
  it is the test's whole content, and deriving it would assert that the derivation agrees with itself.

**Measured precision cost: 2 of 10 candidates (20 %) are the second shape.** The instrument compares
values and does not know the provenance of the input that produced them — and a good synthetic fixture
*deliberately* mirrors the real platform, so it will match. Recorded as a stated limit with an exemption
each, and routed as `SURVEY-M257x-iter161-golden-vs-frozen-needs-input-provenance` rather than papered
over with a heuristic.

## `D-M257x-161-2` — a successful repair can destroy the proof that the instrument works

**The most surprising finding of the iter, and it arrived from the direction nobody watches.**

iter-160's labeled set read the **working tree**, because iter-155's instance was live there — 8 sites.
This iter repaired those 8 sites, and the proof immediately returned
`MISMATCH — a prediction failed; re-derive the taxonomy.`

Nothing was wrong with the instrument. The **subject** had changed. And the two obvious responses —
delete the label, or flip its expectation to BLIND — would each have produced a census reporting **zero**
with **no surviving demonstration that it can fire at all.** That is `§9`'s unfalsifiable-instrument
failure, reached not through a bad guard but through a *good repair*.

**A labeled set that reads the working tree decays the moment you use it.** Every instance now names the
commit that carries it and the proof reads that blob: `@4adc595` fires at 8 sites, `@HEAD` is silent —
which is simultaneously the proof and the ratchet's RED-proof, from one instrument on two refs.

This milestone has audited every guard for *can it fire*. It had never audited one for *will it still be
able to show that once the tree is clean*.

## `D-M257x-161-3` — the exemption window stays tight, and the marker goes adjacent

The first exemption silently did not take: the marker sat at the top of a five-line comment and
`exemption_for` looks back three lines.

**The window is not widened.** A window wide enough to be convenient is wide enough for an exemption to
drift onto a later, different assertion and excuse it silently — which is a worse failure than the one it
would save, because it manufactures a false green in the same file where someone was being careful. The
convention is **prose above, marker adjacent to the assertion**, now demonstrated in both exempted sites.

Paired with the exemption-quality fence: a declaration with no reason, or a reason under 20 characters,
fails — and so does the disappearance of every exemption, which would mean the mechanism is no longer
being exercised by anything.
