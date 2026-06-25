# M41 Retro — Profile depth seeding

## Summary
M41 closed the third (and last) `section` milestone of v1.10: a logged-in hero's `/profile` now shows a
believable **work history + education timeline** and a **deep, role-aligned skill set with a wide
claimed-vs-verified gap**. A new `ProfileSeeder` (rext `stack-seeding`, surface `"profiles"`) writes per
end-user hero the `companies`/`user_experiences`/`user_educations` fan-out (G3) and the ~60-skill
claimed-but-unverified `user_skills`/`user_skill_evidences` tail tied to the experiences (G5), plus a
preset `verified:` bump 8→~30 for thriving heroes — **zero platform-repo edits** (the timeline read path is
unchanged; M41 only supplies rows). Code @ tag `method-acting-m41` → `0346113`; the rosetta side carries the
doc-half (`seeding-spec.md` + `stories-spec.md`) + plan records. Built + hardened (Pass 1+2, 100%
per-function) + closed in one near-clean review: 8 findings, 0 blocking.

## Incidents This Cycle
None. 0 production bugs across the build + the harden Pass 1+2; 0 flakes (5/5 shuffled `-race` at close); 0
regressions (supply-chain GREEN, go.mod/go.sum byte-identical). The close's one new test
(`TestSeedClaimedTail_EmptyEducationsNoPanic`, the Phase-2c adversarial empty-`eduIDs` modulo scenario)
confirmed the in-seeder `len(eduIDs) > 0` guard already handles it — the code was correct; the test pins the
environmental assumption. The one harden-pass test self-correction (the determinism test initially compared
wall-clock audit columns) was resolved within that pass — not a product defect.

## What Went Well
- **The live-schema landmines were caught at build, not at close.** The overview's column guesses
  (`company(nullable)`, etc.) were wrong; the build verified the real demo-3 schema first — NOT-NULL company
  FK, DATE `from<=to OR to IS NULL`, lowercase `location_type` enum, json skills — and the
  `LIVE-SCHEMA CORRECTIONS` section made those the contract. Every emitted row was dry-insert-validated
  against the live schema in a rolled-back transaction. So the close found 0 doc *inaccuracies* (only
  completeness gaps). (M41-D2.)
- **The provenance-edge tie was the natural G3↔G5 join.** The unverified tail needed a non-NULL provenance
  edge (no `job_simulation_id`); tying it to the seeded experiences/educations satisfies the DB CHECK AND
  makes the claimed skills render *under* each work experience — one decision served two surfaces. (M41-D3.)
- **The never-clobber guard is the gap mechanic's safety net.** The claimed evidence UPSERT is a separate SQL
  with `ON CONFLICT … WHERE is_verified = false`, so a (skill,user) collision never overwrites a verified
  row — the verified side always wins. Pinned by a harden invariant test. (M41-D4.)
- **Harden drove both M41 files to 100% per-function and pinned the load-bearing invariants** (determinism,
  date-progression monotonicity, never-clobber, small/exact-pool degradation, no-fabrication blank-node-id
  skip) before close — so the close was a near-clean review.

## What Didn't
- Nothing material. Two minor surfaces: (1) the docs split — the never-clobber guard + the full `from<=to OR
  to IS NULL` CHECK were in `stories-spec.md` but not the `seeding-spec.md` summary (completed at close); (2)
  the `stack-seeding/README.md` test count had drifted (406, pre-M39) — reconciled to the ground-truth 496 at
  close (the handbook-count-drift fix). Both ~minutes to resolve.

## Carried Forward
- **DEF-M40-01 — KPI "AI simulations completed" = 0** → **M42e + M42m** (Fate-2, inherited from M40-D7). M41
  added no deferral of its own. The KPI reads `public.local_jobsimulation_sessions` directly (no CMS, no
  `user_skills`/`user_experiences` surface), so M41's rows can't move it either — it's owned by the
  per-vantage coverage milestones whose exit gate already encompasses a non-zero completed-KPI. Not aged-out;
  deferral re-audit GREEN.

## Metrics Delta
(from `metrics.json`)
- Go test funcs: stack-seeding **462 → 496** (+34; the ProfileSeeder — build +9, review +1, harden +23, close
  +1). Release total 1292 → **1326** (v1.9-close baseline 1248 → +78 across M39+M40+M41).
- Coverage: `profile.go` + `profile_write.go` **100.0% per-function** (75.9% avg → 100% at harden).
- Flake: **0** (5/5 sequential shuffled `-race` at close).
- Supply-chain: **GREEN** (0 new deps; go.mod/go.sum byte-identical across `81ef650^..method-acting-m41`).
- Platform-repo edits: **0** (the `/profile` timeline read path unchanged — M41 only supplies rows).
- Live validation (build+harden-recorded, demo-3): the full company→experience(current+closed)→education→
  claimed-user_skill→guarded-claimed-evidence chain dry-inserted in a rolled-back transaction; all shapes
  satisfy the live CHECKs/FKs/enums; rollback verified clean (0 probe rows leaked).
