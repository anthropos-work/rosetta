# M44 — Profile completeness · Retro

## Summary
The whole-roster DATA-DENSITY completeness pass: §A trajectory-aware self-rating, §B the NEW
`CertificatesSeeder` + `ProjectsSeeder`, §C the manager personal-data unskip, §D bulk-member shallow
career. Closed GREEN with **1 finding / 0 blocking** (a stale rext-tag reference in `seeding-spec.md`'s
Status line, fixed in-close). All four fills render-verified on a live demo-3. Zero canonical platform
edits; 0 new deps; alignment N/A. The first milestone of the M43→M46 extension, and the gate that
unblocks M45.

## Incidents This Cycle
- **No P0/P1/P2 incidents, no regressions, no flakes.** The flake gate ran 5× sequential shuffled on the
  seeders pkg — 5/5 clean.
- **One build-phase render-miss (caught + fixed before close, NOT a post-close regression):** the §D
  avatar fill targeted only `users.picture`, but `/enterprise/members` renders
  `memberships.picture_url` — so the members list showed silhouettes even at `users.picture` 340/341.
  Caught by a live render-check, fixed in the build addendum (`UsersSeeder` now fills BOTH columns — COPY
  on fresh seed + an idempotent `backfillMembershipPictures` UPDATE heal, since `CopyRowsIdempotent`'s
  `ON CONFLICT (id) DO NOTHING` can't heal an existing row). Hardened with a heal-failure-propagation
  test + the cross-column-match invariant. The gotcha is now documented load-bearing in
  `profile-completeness-spec.md` §D ("future avatar work: fill BOTH").

## What Went Well
- **Live-schema-first paid off again.** M41's spec-notes warning ("the live demo schema wins") was heeded:
  the overview's column guesses were wrong on every count (`user_certifications` not `user_certificates`;
  `certification`/`title` not `cert_name`/`project_name`; no `created_at`/`organization_id`) — caught by
  verifying against the live demo-3 DB before writing the seeders, not after.
- **Reuse held the closure invariant.** Both new seeders reused `resolveJobRoleRefs`/`resolveNamedSkillRefs`
  + `CopyRowsIdempotent` + the `PerStackIsolated` audit pattern — no new mechanics, closure stayed GREEN,
  supply-chain byte-identical.
- **Deferral discipline clean.** 0 deferrals; the two `Out:` items routed Fate-2 to already-planned M45/M46
  (their `In:` lists already own them) rather than punted.

## What Didn't
- **The avatar-column gotcha cost a build-addendum + a re-seed idempotency trap** (the first re-seed didn't
  heal existing rows because the COPY is insert-or-skip; needed the explicit UPDATE backfill). Two cross-cut
  surfaces (`users.picture` for `/profile`+menu vs `memberships.picture_url` for `/enterprise/members`) read
  the same logical "avatar" — a non-obvious platform split that only a live render-check surfaced. Now
  documented so future avatar work targets both.
- **A 1-off test-count drift** in the M42m close (state.md quoted stack-seeding=537; ground-truth at the
  tag is 538) — corrected and recorded honestly in the M44 metrics/state, not papered over.

## Carried Forward
- **None from M44.** The `Out:` items (LLM-generated content, deep per-fill-member narratives) are not
  carry-forward — they are owned by M45/M46 by design (Fate-2, already in their `In:` lists). M44 unblocks
  M45 (the engine reuses M44's certificate/project + bulk-member surfaces + the trajectory-aware
  self-rating).

## Metrics Delta (from metrics.json)
- **Go test funcs (stack-seeding):** 538 → **567** (+29). Module total 1376 → 1406.
- **Coverage (seeders pkg statements):** 96.5% → **97.5%** (+1.0); all M44-introduced seeder funcs at 100%.
- **Flakes:** 0 (5× sequential shuffled clean).
- **Supply-chain:** GREEN (0 new deps; `go.mod`/`go.sum` byte-identical).
- **Alignment:** all 5 Clerkenstein gates 100% (N/A change — zero clerkenstein touch).
- **Platform edits:** 0 canonical.
