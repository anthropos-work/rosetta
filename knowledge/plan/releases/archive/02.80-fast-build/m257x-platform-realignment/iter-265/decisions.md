# iter-265 — decisions

## Pre-registrations — SEALED BEFORE THE ENUMERATION RUNS

Sealed in this iter's first commit, corpus at `cf9c469`. Each is falsifiable and each is at genuine risk;
the iter grades all five at close whether they held or not.

**PR-1 — the studio sub-population exceeds the first grep.**
A proper enumeration (one *not* keyed on the string `Dockerfile.dev`) finds **≥ 5** corpus sites that
document the studio-runtime build requirement as a `cms` matter. The first grep found 4
(`cms.md:271`, `staging-bringup.md:428`, `staging_from_dump.md:475`, `setup_guide.md:772`).
*Risk:* the population may genuinely be 3–4, in which case iter-264's route named nearly all of it.

**PR-2 — the whole mechanical slice is between 15 and 70 sites.**
The slice = every corpus site issuing an *operational instruction* (a fenced command line, or an imperative
to edit/clone/run something) that names a decommissioned service repo — `cms`, `jobsimulation`,
`roadrunner`, `skiller`, `skillpath`, `storage`, `messenger`, `customerio-sync`, `chronos`, `intelligence`,
`graphql-wundergraph`. *Risk:* falsifiable in both directions; a slice under 15 makes the "class" a handful
of instances, one over 70 means the slice is too coarse to be the unit of repair.

**PR-3 — the defect concentrates in the OPS GUIDES, not the service docs.**
Of the 10 `cd <decommissioned-repo>` fenced blocks measured at open (all 10 in `corpus/services/*.md`),
**≥ 8 already carry an explicit historical/archived marker**. *Risk:* the reverse is entirely possible —
the service docs could be the rot and the guides clean — and that would invert the repair target.

**PR-4 — no existing guard fires on any member of this class.**
Running the corpus guard family over the current tree yields **0** findings that name a member.
*Risk:* real and specific — `fence_command_guard.py` grades whether `cd <dir>` resolves, and `cms` is
**not in `repos.yml`**, so a `cd cms` fence may already be firing. If it is, the fix shape changes from
*build a fence* to *make an existing check run* (iter-263's shape).

**PR-5 — the class is not studio-only.**
At least one member is a **non-studio** requirement that migrated with a merged service while its
documentation stayed behind (an env var, a make target, a build step, a port, a clone).
*Risk:* if refuted, `FIX-M257x-264-cms-md-past-tense-dependency` was a single-instance route wearing a
class's clothes, and this iter closes it as an instance — a legitimate and useful outcome.

## Escalation clause (pre-registered)

If **PR-4 is refuted** — an existing guard already fires on this class — the iter's fix shape changes
mid-flight from *author a fence* to *make the existing check run and be read*, and the change is recorded
as a decision rather than silently absorbed. This is the same re-shaping iter-263 performed on iter-262's
`D-M257x-262-3`, and it is pre-registered here so that it is a followed rule rather than a retrofit.

## D-M257x-265-1 — the requirement migrated; three copies of its remedy INVERTED

`app/Dockerfile:45-46` hard-COPYs `/build/studio`; `make up` cannot build `backend` without it. Three
corpus troubleshooting entries — `setup_guide.md:771`, `staging_from_dump.md:473`, `staging-bringup.md:428`
— instructed the operator to **delete or comment out those lines**, with
`RUN pip install --no-cache-dir -r studio/requirements.txt` **byte-identical** between the obsolete remedy
and the live Dockerfile. Every clause was wrong for a current stack: the failing image is `app`, not `cms`;
deleting the lines is a platform-repo edit this release forbids; and the Go binary does not run without the
Python runtime, because `app/internal/cms/` hosts the embedded studio-room pipeline `cms` used to own.

**The mechanism is indexing.** A troubleshooting entry is filed under its **symptom**, and the symptom
(`COPY … studio` fails) outlived the fold that moved its cause. An operator greps the error text, not the
repo name, so an obsolete remedy stays reachable for exactly as long as the error message is stable.
Repaired at all four sites (the three above + `cms.md`'s past-tense filing, which iter-264 left open
deliberately so the class could be enumerated rather than its last member repaired).

## D-M257x-265-2 — a marker fence can be GREEN because of the defect's own words

`decommissioned_instruction_guard.py` assertions A–C require a historical marker near any instruction
naming a decommissioned service. Controlled read-only against the pre-repair tree, they fired on **1 site
of 17** and on **none of the three that mattered**: the obsolete remedy says *"the `studio/` submodule
**has been removed from** `cms/main`"*, and `removed from` is in the marker vocabulary.

**A fence whose green is produced by the defect's own phrasing is worse than no fence** — it converts an
unwatched class into a watched-and-clean one. Assertion **D** was added in the same iter: *no corpus
instruction to DELETE code may name a line that is LIVE in the clone set*, a **relation** between two
artifacts rather than the **presence** of a word. D fired on 2 of the 3 known instances, cited to
`stack-dev/app/Dockerfile:46`; the third quotes no code and is a stated reach limit, routed.

The generalisable half: **grade a new enumeration against the instances you already hold.** D's first
regex reached 1 of 3 and would have published `1 finding, 0 remaining` as a complete sweep. The known set
is the only thing that made the shortfall visible, and it will not always be in hand — which is why the
A–C weakness is pinned as a test (`test_marker_assertion_is_satisfied_by_the_defects_own_words`) rather
than fixed silently.
