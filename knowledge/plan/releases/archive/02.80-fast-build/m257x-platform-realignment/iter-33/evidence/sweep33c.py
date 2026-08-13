#!/usr/bin/env python3
"""sweep33c.py — M257x iter-33, the ADVERSARIAL-VERIFICATION pass.

Fixes the 3 blockers the corrections THEMSELVES introduced + the 3 the sweep MISSED,
found by a second read-only audit over the 13 changed files. Same contract as
sweep33{,b}.py: exactly-once anchors, two-phase, non-idempotent.
"""
import sys
from pathlib import Path

ROOT = Path("/Users/marco/workspace/anthropos/rosetta/corpus")

EDITS = [
    # === SELF-INFLICTED 1 — the "never mention organization" over-reach ===
    ("architecture/security_compliance.md",
     "> carries the privacy `Policy()` (`mixin.go:126`). Seven use `OrganizationIDMixin{}`, explicitly *\"a plain\n"
     "> nullable organization_id column\"* with **no policy**, and the rest never mention organization at all.",
     "> carries the privacy `Policy()` (`mixin.go:126`). Seven use `OrganizationIDMixin{}`, explicitly *\"a plain\n"
     "> nullable organization_id column\"* with **no policy** — **and a further ~18 declare a plain\n"
     "> `organization_id` field with no mixin and no policy at all** (`org_membership.go`, `org_subscription.go`,\n"
     "> `organization_settings.go`, `organization_feature.go`, `api_key.go`, `lab_session.go`,\n"
     "> `interview_aggregated_report.go`, `admin_audit_log.go`, `job_simulation_session.go`, …). **Those are the\n"
     "> rows most likely to be missed by an audit**: they look org-scoped and are not policed. The remainder\n"
     "> (the taxonomy, and other global reference data) carry no org column by design."),

    # === SELF-INFLICTED 2 — the anchors DO still resolve ===
    ("services/ai-readiness.md",
     "recognition pattern; the anchors below are pre-`dae0fb2f7` and no longer resolve:",
     "recognition pattern. **The three orphaned COMPONENTS are deleted; the surrounding anchors below still\n"
     "resolve at HEAD** — `urls.ts`, `useNavbarSections.tsx`, the e2e spec, `WorkforceNewClient.tsx` (still\n"
     "omitting readiness) and `useWorkforceAIReadiness.ts` (still cycle-less) are all live, so they remain\n"
     "checkable evidence rather than history:"),

    # === SELF-INFLICTED 3 (minor) — present tense about a surface that is gone ===
    ("services/ai-readiness.md",
     "**not** select between the two manager trees. The manager dashboard gates purely on the GraphQL\n"
     "`aiReadinessEnabled` boolean plus `isEnterprise` nav visibility.",
     "**not** select between manager trees — it never did, back when there were two. The (one) manager\n"
     "dashboard gates purely on the GraphQL `aiReadinessEnabled` boolean plus `isEnterprise` nav visibility."),

    # === SELF-INFLICTED 4 — the splice that orphaned a column list ===
    ("services/hiring.md",
     "   `completion_status` (values `passed`/`failed`/`pending`/`SIMULATION…`) — **spelled correctly in the DB**\n"
     "   (`app/terraform/migrations/20260722104506.sql:12`, `ent/schema/job_simulation_session.go:39`); the\n"
     "   `completition` misspelling survives only in the GraphQL sort-field enum\n"
     "   (`enum.InsightsSortFieldCompletitionStatus`) and a JSON tag, **never as a column name**,\n"
     "   `organization_id`, `tenant_id` (NULL or `=org`), `validation_version`, `anticheat_summary` (optional).",
     "   `completion_status` (values `passed`/`failed`/`pending`/`SIMULATION…`), `organization_id`,\n"
     "   `tenant_id` (NULL or `=org`), `validation_version`, `anticheat_summary` (optional).\n"
     "   ⚠️ **The column is spelled `completion_status` — correctly** (`20260722104506.sql:12`,\n"
     "   `ent/schema/job_simulation_session.go:39`). The `completition` misspelling exists only in the GraphQL\n"
     "   sort-field enum (`enum.InsightsSortFieldCompletitionStatus`), its GraphQL member and a JSON tag;\n"
     "   `insightsSortColumn` (`intelligence.go:885-886`) maps it back to `FieldCompletionStatus`, so it\n"
     "   **never reaches SQL**."),

    # === MISSED 1 — the tenancy claim left standing three files away ===
    ("architecture/architecture_overview.md",
     "The platform uses **shared database, shared schema** with `organization_id` on every table. Data isolation is enforced at three layers:\n\n"
     "1. **Database**: `organization_id` foreign key on all tables; Ent ORM policies auto-filter queries",
     "The platform uses **shared database, shared schema**, with `organization_id` on **org-scoped** tables\n"
     "(**not** on every table — the taxonomy and other global reference data carry none by design). Data\n"
     "isolation is enforced at three layers:\n\n"
     "1. **Database**: `organization_id` on org-scoped tables; Ent privacy policies auto-filter **only the 30\n"
     "   schemas using `OrganizationMixin{}`** — see\n"
     "   [Security & Compliance → Layer 1](./security_compliance.md#layer-1-database) for the measured split"),

    # === MISSED 2 — "the mirror's score column" ===
    ("services/hiring.md",
     "Ent table `public.job_simulation_sessions`, `field.Float32(\"score\")` (`local_jobsimulation_session.go` no longer exists) |",
     "Ent table `public.job_simulation_sessions`, `field.Float32(\"score\")` — **the score column, read at\n"
     "`intelligence.go:1820` and assigned at `:1846`. Not a mirror: `local_jobsimulation_session.go` no longer\n"
     "exists** |"),

    # === MISSED 3 — "the 2-table pair", twice ===
    ("services/hiring.md",
     "the scoreboard scores from the 2-table pair (+ membership + the Casbin gate) alone.",
     "the scoreboard scores from the **single** `job_simulation_sessions` row (+ membership + the Casbin gate)\n"
     "alone — the write-set used to be a PAIR and is now one row, since the mirrors were dropped."),

    ("services/hiring.md",
     "> rows are also M224+ (the M223 scoreboard needs only the 2-table pair).",
     "> rows are also M224+ (the M223 scoreboard needs only the single `job_simulation_sessions` row —\n"
     "> formerly a 2-table pair, until the mirrors were dropped)."),
]


def main() -> int:
    findings, applied = [], 0
    by_file: dict[str, list[tuple[str, str]]] = {}
    for fname, old, new in EDITS:
        by_file.setdefault(fname, []).append((old, new))

    staged: dict[Path, str] = {}
    for fname, pairs in by_file.items():
        path = ROOT / fname
        if not path.exists():
            findings.append(f"[missing] {fname}")
            continue
        text = path.read_text()
        for old, new in pairs:
            n = text.count(old)
            if n != 1:
                findings.append(
                    f"[anchor] {fname}: expected EXACTLY 1 occurrence, found {n} -> {old[:70]!r}")
                continue
            text = text.replace(old, new, 1)
            applied += 1
        staged[path] = text

    if not findings:
        for path, text in staged.items():
            path.write_text(text)

    if findings:
        print("SWEEP FAILED — no file written:", file=sys.stderr)
        for f in findings:
            print("  " + f, file=sys.stderr)
        return 1
    print(f"sweep33c: OK — {applied}/{len(EDITS)} edits applied across {len(by_file)} file(s)")
    return 0


if __name__ == "__main__":
    sys.exit(main())
