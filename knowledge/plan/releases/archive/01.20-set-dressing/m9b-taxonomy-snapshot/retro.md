# M9b — Retro (Taxonomy snapshot, the first real surface)

## Summary
M9b proved the M9a framework on the **real ~2.1 GB public taxonomy** — the 10-table skiller catalog in FK
replay order, `org_id IS NULL` filters plus **parent-scoped predicates** for the column-less embedding/
translation/`job_role_skills` tables, the pgvector index rebuilt via `REINDEX` on replay (vectors verbatim,
dim 1536 in the manifest), the **manifest→fidelity bridge** + `datadna measure-snapshot` CLI, and the
`TaxonomySnapshotSeeder` DAG node. Coverage moved `waived-m7c → snapshot-seeded-m9b` across all 5 fidelity
operators. Built → 3 harden passes → close.

The close took **two attempts**: attempt 1 BLOCKED at Phase 1b on a RED deferral audit — the **offline
pg_dump-FILE reader** had been carried M9a→M9b without landing (a repeat/aged-out deferral), so the audit
required a fresh user fate decision. The user **DROPPED** it (M9b-D9). Attempt 2 recorded the drop, landed the
companion correctness fix (docs+code had been *claiming* an offline reader that never shipped), and merged clean.

## Incidents This Cycle
- **P2 (build, fixed):** the parent-scope leak probe counted customer-parented rows across the WHOLE live table
  → it would have FALSELY ABORTED a correct capture (`skill_embeddings` legitimately has customer-parented rows
  the filter excludes). Fixed: `buildParentLeakProbe` ANDs the capture filter (count leaks WITHIN the captured
  set, 0 by construction); extracted as a pure function + unit + fuzz tested (extensions `0404760`).
- **P2 (close, the blocker):** the offline pg_dump-FILE reader was an M9a→M9b repeat/aged-out deferral; close
  attempt 1's deferral audit went RED and required a user fate decision (resolved: DROP, M9b-D9).
- **P3 (close, doc/code lie):** `source.go`/`main.go`/`README.md` + `snapshot-spec.md` claimed an offline reader
  ships, while the CLI always required `--dsn` and exposed a dead `--dump <path>` flag. Fixed alongside the drop.
- No regressions; no flakes (2 fuzz targets, 0 crashers; stack-snapshot 5/5 shuffled green).

## What Went Well
- **The M9a framework held.** The real taxonomy surface dropped in via the existing capture/firewall/replay/
  store machinery; the one genuinely new piece was the parent-scoped predicate (M9b-D2) — the M9a empty-filter
  gap that the toy surface never exercised. The framework's seams (interface-based adapters) made the fidelity
  bridge + the harden-pass error-path tests cheap.
- **The deferral audit did its job.** It caught a real repeat-defer pattern (the offline reader, "belongs with
  the next surface" each time, with no next surface left for the read path) and forced a conscious decision
  instead of a silent drop-at-close.
- **The drop was the right call and made the docs honest.** Dropping the offline reader removed a feature that
  added no capability (identical snapshot) and no reliable speed gain (fiddly two-pass dump text-parsing vs a
  one-time Postgres bulk-restore), AND it forced the docs+code to stop lying about a reader that never existed.

## What Didn't / Constraints
- **Docs+code drifted from the truth during M9a.** The "dump-ingest ships offline" framing was aspirational at
  M9a and never corrected when the reader didn't land — a doc/code lie that survived until this close caught it.
  Lesson: a deferred *mechanism* must leave the docs describing what EXISTS, not what's planned.
- **The live capture path is still proven only catalog-deep + on `reference-toy` hermetically.** The real
  ~2.1 GB read (and the `REINDEX` cost) is exercised via `--dsn` against a restored dump or the safe primary
  read; M9b proves the *plan* + the firewall + the fidelity measure against prod catalog evidence, not a full
  GB-scale read on this machine (no live skiller PG in the loop here). That is the intended split.
- **Two-repo close choreography (again).** The code lives in the gitignored extensions clone (own git, tagged +
  pushed); the rosetta merge carries only docs + records. Tag-last continues to be the right discipline.

## Carried Forward
- **None new.** DEF-M9b-01 LANDED, DEF-M9b-02b DROPPED (M9b-D9 — cut, not deferred, not seeded to v1.3),
  DEF-M9b-03 (tag) resolved this close. The M9a-retro source carry-forward for the offline reader was CUT so it
  cannot re-surface. M10 (Directus content) + M11 (recipes/corpus) own their own scope; the v1.3 seeds
  (AI-content, shareability, more-mirrors, cloud-store) are unchanged in `roadmap-vision.md`.

## Metrics Delta (from metrics.json)
- **Go test funcs:** 556 → **635** (+39; stack-snapshot 128→167, stack-seeding 164→204).
- **Coverage (post-close):** firewall/reference/replay/source 100%; capture 95.5%; cmd/stacksnap 76.4% (up from
  72.3% at harden — the close regressions added source-resolution-branch coverage); dna fidelity-probe 100%.
- **Flake:** 0 (5× shuffled sequential, stack-snapshot — the close-touched module). **Race:** clean.
  **gofmt + vet:** clean. **Fuzz:** `FuzzSplitCSV` (1.7M) + `FuzzBuildParentLeakProbe` (766K), 0 crashers.
- **Review:** 4 findings (1 scope=the drop, 1 code-quality=remove `--dump`, 2 docs, 1 tests, 1 decision-triage),
  0 code defects beyond the build-phase parent-leak fix, deferral audit RED→GREEN via the user DROP.
