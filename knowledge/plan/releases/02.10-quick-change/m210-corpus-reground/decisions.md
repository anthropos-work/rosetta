# M210 — Decisions

_Implementation choices with rationale, logged as they are made._

## Adoption mechanic: selective per-section adoption (not cherry-pick)
The colleague's `origin/docs/skiller-in-app-merge` is one docs commit (`e3d4692`) + 4
separate PNG commits (`_email-assets/sales-newsletter-331/*.png`). Chose **selective
per-section adoption** (`git checkout origin/docs/skiller-in-app-merge -- <files>` +
manual reconcile) over `git cherry-pick e3d4692` because:
- build-milestone wants **per-section commits**; the colleague's monolithic docs commit
  spans all 6 M210 sections at once — a cherry-pick would fight the section discipline.
- M208 already re-framed `backend.md` (fact-sheet) + `skiller.md` (stub banner) **differently**
  from the colleague → the cherry-pick would conflict on those and need manual reconcile anyway.
- The rext-facing docs need a real **flip to `public.*`**, not the colleague's interim
  "Present-state note" annotation → I supersede those hunks, not adopt them.
- PNGs are excluded (unrelated to the merge).
Attribution preserved via commit messages crediting `origin/docs/skiller-in-app-merge`.

## KB-1 (Phase 0b): pre-flip staleness is the deliverable, not a blocker
Audit YELLOW. 33 stale `skiller.<table>` refs on the release branch vs M209's landed `public.*`
rext code (89 refs verified in `.agentspace/rosetta-extensions`). Not a blocker — the milestone
overwrites these; ground truth is the code, not the stale docs. Tracked as progress.md sections.

## KB-2 (Phase 0b): profile-completeness "43/44" literal does not exist
Orchestrator END-STATE #2 says "member count 43/44 → 44/44" in `profile-completeness-spec.md`.
**No literal `43/44` exists** anywhere in corpus/skills, nor in the file's git history
(`git log -S "43/44"` empty). The file's schema refs are ALREADY `public.*`. The only ratio is
`340/341` (a completed avatar-coverage bug narrative, line 72 — deliberately historical). The one
skiller mention is line 105 "the closure gene governs skiller node-ids" — a permitted
conceptual/historical reference (not a tooling-queries-`skiller.<table>` claim). Resolution logged
in Section 2 below after evidence-based inspection.

## §2 resolution: profile-completeness-spec.md — the "43/44 → 44/44" is not a real edit
Exhaustive search (corpus + `.claude/` + the rext authoring copy `.agentspace/rosetta-extensions`)
found **NO** `43/44`, `44/44`, `43 of 44`, standalone `43`/`44` count, or member-roster count that the
skiller merge shifts. The file's DB refs were **already `public.*`** (public.users, user_skills,
user_certifications, user_projects, world_languages, user_languages, membership_languages) — the
member-owned tables always lived in `public`, not the skiller schema, so no schema flip was owed.
The ONE genuine merge-sweep miss was the prose at the M50 languages bullet: "the closure gene governs
**skiller node-ids**, not ISO codes" — a conceptual reference. Fixed it to "the closure gene governs
the **taxonomy node-ids** (the skills/roles catalog — in the `public` schema since the v2.1 skiller→app
merge)". This aligns the file fully with the merged world (no lingering bare "skiller" a reader could
misread as a live schema). Per the three-fate rule + the audit "don't fabricate" principle, I did NOT
invent a phantom 43/44→44/44 count. The file is now merge-swept.

## §5: superseded the colleague's stale stack-snapshot/SKILL.md interim note
The colleague's `stack-snapshot/SKILL.md` hunk added an interim warning: the taxonomy
surface "still targets the legacy `skiller` schema... hits **exit 4**... until re-pointed
at `public`". That note was written PRE-M209. M209 **already re-pointed** the surface —
verified in `.agentspace/rosetta-extensions/stack-snapshot/taxonomy/taxonomy.go` ("post
skiller-in-app merge, v2.1 M209; was 'skiller'") + capture/replay tests use `public.skills`.
So I did NOT adopt the colleague's exit-4 note; I wrote an accurate M209-done note (surface
re-pointed at `public`; replay/capture work post-merge; recapture → M211). The other 4 skill
files (dev-up/reference, stack-update/reference, update-knowledge, db-query) were correct →
adopted. Container/migration/RPC/subgraph facts verified against the re-synced
`stack-dev/platform/docker-compose.yml`: no skiller service, 11 graphql-profile containers,
`SKILLER_RPC_ADDR=http://backend:8083` (×4), 12 app services.
