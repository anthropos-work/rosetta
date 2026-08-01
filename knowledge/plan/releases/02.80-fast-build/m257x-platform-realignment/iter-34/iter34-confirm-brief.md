# Clause-5 CONFIRMING pass — addendum to `iter-33/iter33-groundtruth.md`

Read `../iter-33/iter33-groundtruth.md` **first and in full**. Everything in it still holds: the platform
clone is at origin HEAD **`2adcf71`**, re-fetched and confirmed unchanged at 2026-08-02 00:33 CEST. The
grading rule, the authoritative-source rule, the 14 derived facts and the mandatory full-read method are
all unchanged and binding.

This addendum adds only what is different about **this** pass.

## What is different: a corrective sweep already ran on this corpus

At iter-33 a sweep closed **25 blockers across 13 of these 40 files**:

    corpus/architecture/architecture_overview.md      corpus/services/ai-readiness.md
    corpus/architecture/dependency_map.md             corpus/services/backend.md
    corpus/architecture/external_services.md          corpus/services/clerkenstein.md
    corpus/architecture/security_compliance.md        corpus/services/cms.md
    corpus/architecture/service_taxonomy.md           corpus/services/hiring.md
    corpus/architecture/shared_libraries.md           corpus/services/jobsimulation.md
                                                      corpus/services/storage.md

**You are not confirming that sweep's work. You are auditing this corpus as it stands, from scratch.**

Two things follow, and they pull in opposite directions — hold both:

1. **Do NOT treat swept text as trusted.** iter-33 measured its own repair pass and found a **24 %
   self-inflicted rate**: 6 of 25 blockers were *created or left by the corrections themselves*. One was
   a tenant-isolation fence that claimed the non-mixin Ent schemas "never mention organization at all"
   when 33 do — it contradicted itself inside its own blockquote and failed **in the dangerous
   direction**. Another was a claim retracted-with-measurements in one file and left standing verbatim
   three files away. **New text is the least-audited text in the corpus.**
2. **Do NOT treat unswept text as clean.** 27 of these 40 files were never edited. Pass 1 reported only
   **1 blocker across 18 of them** — a density an order of magnitude below every other group. That is
   as likely to be under-detection as cleanliness. Give the unswept files the same full read.

## Where the errors actually live

Every one of the ~40 `file:line` anchors the sweep wrote **verified exact**. The defects were entirely
in the **surrounding prose** — the sentence before the anchor, the summary line at the top of the
section, the table cell that restates the claim in different words. An anchored check does not look
there. **Read the paragraph, not the citation.**

The class that pass 1 found is **derived-fact rot**, not status rot. Every doc gets *who is merged*
right (that layer is machine-fenced by `ServiceDocStatusFence`); what rots underneath is:

- **table names** the platform dropped or renamed — e.g. `public.sessions` → `job_simulation_sessions`;
  `local_jobsimulation_sessions` / `local_skill_path_sessions` **dropped** at
  `app/…/20260729133514.sql:58-62`;
- **package paths** that were split out — e.g. `internal/workforce` → `internal/aireadiness`;
- **work described as "routed forward" that already shipped**;
- **counts and cardinalities** stated as facts (schema counts, table counts, subgraph counts).

None of it uses merged/live/gone vocabulary. A grep cannot find it. Only a read can.

## Your job, precisely

For every file in your assigned group:

1. `wc -l <file>` — record the number.
2. Read it **top to bottom, in full**. No grep, no skim, no sampling.
3. For every claim that could misdirect real work, verify it against the platform clone at
   `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform` (read-only) or against
   `corpus/architecture/platform-migration-status.md`.
4. Grade: **BLOCKER** (false at HEAD *and* acting on it would misdirect real work) / **minor**
   (true-but-confusing, stale line number, dead link, imprecise count) / **not-a-finding**.

**Read-only. Do not edit any file anywhere, in any repo.**

## Report format

Open with a **positive-control table** — one row per assigned file: path, `wc -l`, and the last line
number you actually read. A file you did not finish must be reported as **UNREAD**, not guessed.

Then, per finding: `path:line` · verbatim quote of the false text · what is actually true at HEAD ·
the platform `file:line` or migration-map row that proves it · grade · a one-line suggested correction.

End with your own counts: **N blockers, M minors**, and one sentence on whether your group read as
recently-repaired, never-touched, or mixed.

**Do not tune your count toward any expectation.** A pass that reports zero because zero was hoped for
is worthless; so is one that inflates minors into blockers to look thorough. The grading rule is the
only arbiter.
