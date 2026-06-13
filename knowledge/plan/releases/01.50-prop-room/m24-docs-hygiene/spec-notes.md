# M24 — Spec notes

_Technical detail that doesn't belong in `overview.md` (code maps, contracts, edge cases). Accumulates during build._

## Pre-flight audits — §1 Stale local-Directus corrections (session start)
- **Verdict: GREEN** — report at `kb-fidelity-audit.md`, audited at sha `8807ac0` (HEAD at session start).
- Topic → doc → code triples (verified 2026-06-13):
  - Local-Directus reality → `external_services.md` (122/181/270-271/177-181/552/617) · `service_taxonomy.md:242-260`
    · `quick_ops.md:162` → `stack-dev/platform/docker-compose.yml` (**NO directus service**; only env pointers
    `DIRECTUS_BASE_ADDR=https://content.anthropos.work` lines 237-238). All the local-image/admin/compose claims are
    STALE — the M24 §1 deliverable.
  - Per-stack Directus end-state → `directus-local.md` (exists, 261 lines) · `snapshot-spec.md` known-state (389-413,
    already M22-current) · `safety.md` §2 → `stack-snapshot` / `dev-setdress.sh`.
  - Zero-critical-genes guard → `alignment/internal/compare/compare.go:247` (`pct` returns `100.0` when `d==0`) +
    `:92` (`rep.Critical = pct(critAligned, critTotal)`, `critTotal==0` ⇒ vacuous 100%) · `dna.go:169` `Validate()`
    (rejects per-capability invalid criticality only, never the aggregate "zero critical genes"). Bug CONFIRMED
    ABSENT-guard at the cited lines.
  - `/project-stats` scope → `developer-kit/skills/project-stats/stats.sh` (scans repo tree incl. gitignored
    `stack-*/` clones).

## Audit-reuse anchor
- Audit sha: `8807ac0`. Load-bearing knowledge docs for this milestone: the corpus docs listed above. Sections §1–§3
  edit these very docs, so each later section re-checks per the audit-reuse rule (the docs change between sections by
  design — but each section's edits are themselves the verified-correct truth, so no re-audit of an unrelated
  subsystem is needed; sections §4–§7 are ext-repo code, a different subsystem → no doc-fidelity dependency).
