# Parent-side re-derivation, run BEFORE reading any auditor report

Measured 2026-08-02 against `stack-demo/app/internal/data/ent/schema/` (the `app` clone sibling to the
platform clone; platform origin HEAD `2adcf71`, re-fetched unchanged at iter open).

The milestone's standing rule is **re-derive, don't re-match** — an anchor that still matches proves the
quote is there, not that the replacement is true. iter-33's tenancy fence had already been wrong twice in
opposite directions, both times failing *toward* "isolation is handled", so its current numbers were the
single highest-value thing to re-derive independently of the auditors.

## The counts

    ls *.go                                    →  139   (FILE count)
    grep -l 'ent.Schema' *.go                  →  135   (actual schema types — 4 files are not schemas)
    grep -l 'OrganizationMixin{}' *.go         →   30
    grep -l 'OrganizationIDMixin{}' *.go       →    7
    schemas with NEITHER mixin                 →   98
      ...of those declaring "organization_id"  →   18

## Verdict on `security_compliance.md:66-77` (the twice-wrong tenancy fence)

| claim in the doc | re-derived | verdict |
|---|---|---|
| "only **30** use `OrganizationMixin{}`" | 30 | ✅ exact |
| "**Seven** use `OrganizationIDMixin{}`" | 7 | ✅ exact |
| "a further **~18** declare a plain `organization_id` … no mixin and no policy at all" | 18 | ✅ exact |
| the 9 named example files | all 9 present in the derived 18-file list | ✅ exact |
| "of **139** Ent schemas" | 139 is the **file** count; **135** declare `ent.Schema` | **minor** — already logged under `DOC-M257x-iter33-corpus-minors` |

**The fence's substantive numbers hold on independent re-derivation.** A first attempt at this
measurement returned **22**, not 18 — because it subtracted only `OrganizationMixin` users, leaving the
7 `OrganizationIDMixin` schemas in the pool. That is a *different denominator* from the one the doc's
sentence actually declares ("no mixin **and no policy at all**"). On the doc's denominator — schemas
carrying **neither** mixin — the count is 18, exact.

Recording the near-miss because it is the shape of error this milestone keeps finding, and here it ran
in the flattering direction: **the doc was right and my first re-derivation was wrong.** Had I graded on
that first read, I would have filed a blocker against a correct claim — a false positive is as costly as
a false negative when the whole clause turns on a blocker count.

Full 18-file list, for any future re-check:

    academy_feedback · admin_audit_log · ai_readiness_diagnose_narrative · ai_readiness_recommendation
    api_key · assignment_invitation_link · interview_aggregated_report · job_role_skill_suggestion_cache
    job_simulation_session · jobsimulation_feedback · lab_session · org_membership
    org_membership_invitation · org_sim_link · org_subscription · organization_feature
    organization_settings · profile_history
