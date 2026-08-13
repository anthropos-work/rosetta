# iter-282 — decisions

## D-M257x-282-1 — the class is DRIFT, not copying, and only drift is fenceable

`ROUTE-M257x-h70` sizes at 172 (module, doc) pairs sharing a verbatim run. **Repairing 172 pairs is not
the deliverable and would not be an improvement** — a corpus that quotes its tooling is doing the right
thing, and de-duplicating it would remove the quotation without removing the coupling. What the three
paid-for instances actually share is narrower and mechanical: *the same sentence in both trees carrying
different numbers in the same slot.* That is the fence's subject. Copying is left alone.

## D-M257x-282-2 — direction-blind by construction

Harden pass 70 saw the corpus copying a stale docstring; pass 71 saw the corpus correct a claim whose
correction never reached the tooling. `prose_twin_guard` takes **no view on which copy is the original**
— it names the disagreeing sites and stops. A fence that named an original would close whichever half
happened to be found first, which is the defect pass 71 recorded.

## D-M257x-282-3 — three false-positive classes, excluded by CONSTRUCT, and each was measured

Not one of these was anticipated; all three came out of the sizing probe.

1. **An ordinal is not a claim.** `rule 8 — a check that skips…` and `exit 2, not 0 — a check that
   skips…` share their entire run and differ only at the window's edge, where the number belongs to the
   *preceding* sentence. Excluded by requiring shared context on **both** sides of a divergent slot.
2. **An equation written the other way round is agreement.** `2,989 + 51 = 3,040` and `3,040 = 2,989 +
   51` normalize identically with the numbers permuted. Compared as **multisets**, they agree — and they
   do agree. Excluded by comparing multisets, never ordered tuples.
3. **A template instantiated for two subjects is not a copy.** Three sibling service docs share a
   paragraph and fill in their own ports; a maxim about mutation batteries is restated with each
   battery's own size. **Both numbers are correct and no repair leaves both true.** Not mechanically
   separable from drift, so: an explicit waiver that must carry a reason.

## D-M257x-282-4 — the waiver is keyed on the SENTENCE, never on `file:line`

An anchored waiver expires the moment a paragraph is inserted above it, and it expires **in the accept
direction**: the fence goes RED at a pair a human already adjudicated, and the cheapest way to green it
is to re-waive at the new line. That is a ratchet, and it is the same rot this fence exists to catch,
committed by the fence. Keyed on the numeric-slot template, a reword **lapses** the waiver — correct,
because a reworded claim has not been adjudicated. Both directions are pinned as tests.

## D-M257x-282-5 — two tiers, because a silent exclusion is a defect

The shared-context rule buys precision and costs recall, and the cost is not hypothetical: the
`up-injected.sh` advisory-anchor divergence sits at a run's tail and is invisible to it. So the guard
prints a **REPORT tier** — divergence anywhere in a shared window, never fatal — beside the RED tier it
grades. The fence's own recall gap is published rather than implied. Two of the REPORT-tier findings
were repaired this iter anyway; the residual is routed, with its size.

## D-M257x-282-6 — TWO defects in the instrument, both found by triage rather than by review

- **The ordinal-leader stripper ate the integer part of a line-leading decimal.** `136.5 s npm ci`
  normalized to `5 s npm ci`, so the fence reported a file disagreeing **with itself** at two sites that
  say the same thing. Repaired by requiring the ordinal to be followed by whitespace; both directions
  pinned (the decimal survives; a genuine ordinal is still stripped).
- **One defect was reported once per overlapping window.** A sentence longer than the window yields many
  windows, and the sliding start makes each a slightly different site set — so the population count was
  a count of *windows* wearing a count of *defects*' clothes. Merged on site overlap, and the merged
  finding reports the **union** of sites: a window that starts a token later drops a site, and a finding
  naming two of three stale copies sends the repair back a second time.

## D-M257x-282-7 — the exclusion arms could not tell exclusion from refusal

Staging the only code file under `tests/fixtures/` left the code side with **no** numeric prose, so the
guard refused (exit 2) — and the first draft of those arms read that refusal as the exclusion working.
A ballast file makes the two outcomes distinguishable, and a third arm stages the same twin one
directory over and requires exit 1. **A fixture that cannot tell its two outcomes apart is not evidence
for either** — the third occurrence of a fixture-shaped defect in three iters.

## D-M257x-282-8 — the ratchets were held by rephrasing, never by bumping

`TEST_MODULE_LITERAL_CEILING` breached +2 on the new test module (`7700 seats` / `7900 rows` in a staged
fixture). Repaired by moving the numbers out of the literal text (`%d`-formatted) rather than raising the
ceiling: **the fixture is the subject here, and it must not also be a member of another census.** All
three ceilings read `exact +0` afterwards.

## D-M257x-282-9 — SIDE FINDING, escalated rather than absorbed: the demo writes to the PRODUCTION S3 bucket, and it is already documented

Surfaced by the orchestrator mid-iter from live use of `demo-2`: studio-desk's *"fine-tune in Advanced
mode"* attempted `s3:PutObject` against `s3://production-storage20240826131618541000000005/cms/…` and was
refused **403** by IAM. **Nothing was written.** Checked immediately, read-only:
`corpus/ops/safety.md:203` **already names this exposure precisely** — the private-bucket row is marked
*"unclassified — the guard does not cover it"*, states that compose hardcodes
`STORAGE_S3_BUCKET=production-storage…` on `backend`, that `PreflightEnv` forces only the **public**
bucket (`stack-seeding/isolation/audit.go:146`), and that a stack with working AWS credentials therefore
writes private uploads into the production private bucket. Its disposition is an **open escalated item**,
`DEF-M257x-iter80-storage-prod-bucket`, severity high.

So the corpus did not overstate at that row. What the live event adds is the first **exercised** proof of
the path, and it does not close the question of whether other sites in `safety.md` (or the skills) assert
the stronger *"a demo cannot write prod"* without that qualification. **Routed, not absorbed** — it is a
third line of investigation in this iter and the tripwire applies.
