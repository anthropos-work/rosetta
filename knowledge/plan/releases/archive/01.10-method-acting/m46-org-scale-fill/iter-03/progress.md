**Type:** tik (#2, under TOK-01)

# M46 · iter-03 — per-story batch distribution

Implemented deliverable #2: generated members route into their own story's org.

## What landed (rext `stack-seeding/`)
- `blueprint/batch.go`: new `BatchMember.StoryIndex int`, set at expansion from `st.Index`. It is pure
  seed-time routing metadata — NOT part of the mother prompt, so the cache key is unchanged (no reseed).
- `seeders/generated_batch.go`: the seed loop no longer resolves `stories[0]` once. It resolves each
  member's story context (`org / key-prefix / email-domain / company`) PER member from `m.StoryIndex`
  (with a safe fallback to story 0 for an out-of-range index). The per-story key prefix gives two stories'
  members distinct deterministic ids even for the same `(batchID, memberIndex)`; the membership's
  `organization`, the email domain, and the company all become per-story.

## Tests (fixtures-first, no key/cost) — `seeders/generated_batch_test.go`
2 new tests, green + the full stack-seeding suite green:
- `PerStoryRouting` — a 2-story blueprint (a batch each) → story-1 members in `LegacyOrgID`, story-2 in
  `StoryOrgID(s2)`; 2 DISTINCT orgs; per-org email domains (orgone.com / orgtwo.com); 4 distinct user ids.
- `PerStoryFillComposition` — deliverable #1 (fill) + #2 (per-story) compose: each story's FILL batch fills
  ITS OWN org to ITS OWN Size from one descriptor each (OrgOne 4−1 hero = 3, OrgTwo 2 = 2).

## Re-measure
Gate primary metric (the M42 semantic sweep on a generated org) NOT measured this iter — it moves at the
real-run gate-proving tik (#5). Progress here is a structural deliverable landed toward the gate, composing
cleanly with iter-02.

## Close — 2026-06-26

**Outcome:** per-story batch distribution (deliverable #2) landed + unit-proven (2 new tests + full suite
green). Each story's generated members now land in its own org/domain; composes with the iter-02 auto-fill
(each org fills to its own Size).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (org-scale gate proven on the real-run sweep, tik #5; this iter lands a deliverable the
gate depends on)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (2 tiks this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** (none beyond TOK-01)
**Side-deliverables:** none.
**Routes carried forward:** deliverable #3 (preview/dry-run mode), #4 (`--gen-batches` fence + throughput/429
verification), #5 (the real-run gate-proving) remain on TOK-01's plan.
**Lessons:** Carrying the routing key (`StoryIndex`) on the member but keeping it OUT of the cache-keyed
mother prompt is the clean way to add per-story routing without touching the $0-reseed invariant — routing
is seed-time identity (CODE-owned), not generation content (AI-owned). The per-story key prefix
(`storyKeyPrefix`) already existed for the curated population; reusing it for the generated members keeps
the id-namespacing consistent across curated + generated members of the same story.
