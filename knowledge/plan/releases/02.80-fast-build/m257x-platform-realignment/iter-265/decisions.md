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
