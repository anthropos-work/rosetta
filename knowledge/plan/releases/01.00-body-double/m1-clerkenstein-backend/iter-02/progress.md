**Type:** tik

# iter-02 — author the Clerk Alignment DNA

Stood up the `clerkenstein` workspace (gitignored `anthropos-demo/`, own git) and authored the Clerk
Alignment DNA from the consumed surface.

## Close — 2026-06-03

**Outcome:** Clerk DNA `dna/clerk-2.6.0.json` authored + **validated** by `alignctl dna validate` — **22 genes / 11 capabilities (13 critical, 9 standard)**. Covers authn (VerifyToken × {valid-session, valid-no-org, expired, malformed}) + the 10 orgclient methods. `anthropos-demo/` added to rosetta `.gitignore`; clerkenstein repo git-init'd + committed (1 commit). D1 resolved (hybrid).
**Type:** tik
**Status:** closed-fixed (planned scope — workspace + validated DNA — landed; this is a DNA-authoring iter shape: the genome is the deliverable, not yet a score)
**Gate:** NOT MET (alignment score still 0% — no mirror/goldens yet; iter-02 delivers the prerequisite genome)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n (D1 resolved; colony v0.34.1 + clerk-sdk-go confirmed in the module cache → the authn-twin build is feasible) — (5) cap-reached: n (1 tik) — (6) protocol-stop: n — **Outcome: continue → iter-03**
**Decisions:** D1 RESOLVED (hybrid, milestone-root `decisions.md`); iter-01-D1 (colony replace granularity) still open, to resolve in iter-03 by reading the cached `colony/authn` Provider interface.
**Side-deliverables:** none.
**Routes carried forward:** iter-03 (tik) — build the authn-provider twin (satisfy `colony/authn`'s Provider, local JWT mint/verify, one universal credential) + its `--target source|mirror` runner + locally-captured authn goldens → first real `alignctl run` score on the 6 critical authn genes.
**Lessons:** capabilities returning server-generated values (org ids, invitation ids/timestamps) can't align under `exact`/`normalized` (the mirror's generated id ≠ the source's) — they use `shape` (CreateOrganization/minimal-valid, InviteMember/valid, BulkInviteMembers/*). Most orgclient methods return only `error`, so `error_class` is the dominant operator; authn claim extraction is deterministic given the token → `exact`.
