# iter-125 — decisions

## `D-M257x-125-1` — a filed defect is re-derived AFTER filing too, and the correction goes first

`D-M257x-121-1` made re-derivation mandatory **at filing time**. This iter tested the entry that rule
produced, and found the rule does not go far enough: **3 of the entry's 4 claims re-derived verbatim and
the fourth was wrong** — `services.tf:47-57` names the root module's *inputs*, not the container's
environment, so `SECRET` and `KEY` were attributed to an anchor that does not carry them.

**Decision: the correction is stated FIRST in the entry, and the entry now cites the task definition
rather than the module's call site.** A platform engineer who opened `services.tf:47-57` looking for
`SECRET` would have found nothing and stopped reading — **a wrong anchor inside a true report is more
expensive than a missing one**, because it spends the reader's trust before the finding is reached.

**And the re-derivation paid for itself in the same pass:** it found `KEY` (`directus/terraform/main.tf:111-114`)
injected as a **plain `environment` value** interpolating an SSM parameter's `.value` — so that secret is
also materialised into the task-definition JSON in clear. **A second, independent exposure that nobody was
looking for**, found only because the environment was read from where it is actually assembled.

## `D-M257x-125-2` — disclosure is a PLACEMENT problem, and a census is not a placement

iter-123 wrote the AKB comparison in full, correctly, with numbers — in `org-repos.md` § 11, a repo
census. **Nothing a reader would see changed.** A reader who needs the taxonomy figures opens
`shared_libraries.md`; an engineer who is about to install the plugin opens `toolchain_overview.md`.
Neither opens a 93-row repo inventory.

**Decision: a finding is placed where the DECISION is made, not where the investigation happened.** The
same text landed in three places — the figures section, the install line, and the register — and only the
second of those is load-bearing for the actual harm, which is *an engineer installing a plugin that
serves a refuted figure into their editor on every Anthropos repo*.

**This generalises past AKB** and is booked as `platform-alignment.md` §5 **rule 55**. It is the sibling
of iter-124's rule 54: 54 says a correction must reach every site publishing the predicate; 55 says a
*disclosure* must reach the site where someone acts on it.

## `D-M257x-125-3` — the two AKB verdicts are kept apart, and "AKB is wrong" is REFUSED as a summary

*"18K roles"* is **REFUTED** — the public subset alone measures 22,470, and public ⊆ total, so the true
count is at or above it. *"60K skills"* is **UNVERIFIED, not refuted** — the same capture is public-only,
cannot see org-private skills, and therefore neither supports nor rules out 60,000.

**Decision: never publish a merged verdict.** A summary reading *"AKB's taxonomy figures are wrong"* is
**false on the skills row**, and this milestone has spent iterations on exactly this failure — a
conclusion that survives while the quantifier under it does not (`D-M257x-121-2`). The register and both
corpus placements carry the split.

**Related, and deliberately hedged:** `public.job_role_embeddings` holds 18,919 rows, which is a plausible
source for a mis-transcribed "18,000 roles". **Nothing here can measure what AKB's author read**, so it is
filed as a lead for the owner and labelled a hypothesis. That is the `hedge` fate used correctly — for a
proposition genuinely unreachable from here, not manufactured for a fact somebody could measure.

## `D-M257x-125-4` — the direction of the errors is recorded, because a ranking would be false

**AKB was RIGHT and this corpus WRONG on the WunderGraph router's production residue, in a fenced table.**
The reason is structural and worth more than the instance: **AKB reads `infrastructure`, which this corpus
had never cloned.** iter-123 cloned it and confirmed AKB's reading.

**Decision: every placement states this alongside the taxonomy contradiction.** A document that lists only
the other corpus's errors is advocacy, and the next person to compare them will find the omission and
discount the rest. The two have **different blind spots, not a ranking** — and this corpus's blind spot
was a *clone-set* limit it had been calling a *measurement* limit for four iterations.

**Found while writing it:** `org-repos.md` § 11 item 1 still read *"Unresolved"* — one screen below § 3 of
the same file, which had settled it. **A document contradicting itself across two of its own sections**,
inside the very section that documents a disagreement between corpora. Repaired. Same class as rule 54,
smallest possible reach: one file.
