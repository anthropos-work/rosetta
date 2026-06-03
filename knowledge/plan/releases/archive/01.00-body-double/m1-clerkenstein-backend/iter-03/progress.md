**Type:** tik

# iter-03 — the authn twin

Built Clerkenstein's authn-provider twin and produced the first real alignment score.

## Close — 2026-06-03

**Outcome:** authn twin built — HS256 mint/verify (one universal key) + `GetUser` extracting the platform claims into a `clerkUser` that **implements the real `colony/authn.{Provider,User,Organization}`** (compile-time `var _ authn.Provider` assertion; builds offline against cached colony v0.34.1). Runner + hand-authored authn goldens. **`alignctl run`: VerifyToken 4/4 ok → score 0% → 21.1% overall / 30.8% critical (4/22 genes).**
**Type:** tik
**Status:** closed-fixed (planned scope — authn twin + first score — landed; the 4 critical authn genes align)
**Gate:** NOT MET (21.1% overall / 30.8% critical < 95% / 100%)
**Metric delta:** overall 0 → **21.1%** (+21.1); critical 0 → **30.8%** (+30.8). 4 genes aligned (VerifyToken ×4).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (2 tiks) — (6) protocol-stop: n — **Outcome: continue → iter-04**
**Decisions:** iter-01-D1 RESOLVED — Clerkenstein implements the real `colony/authn` interface; zero-platform-change injection = **replace the whole `colony` module** (a Clerkenstein-patched colony, the staging `vendor-colony/` precedent), deferred to the injection tik.
**Side-deliverables:** none.
**Routes carried forward:** iter-04 (tik) — the **critical orgclient** methods (CreateOrganization, CreateMembership, ChangeRole, DeleteOrganizationMembership) + goldens → drive **critical → 100%**. iter-05 — the standard orgclient methods (InviteMember, BulkInviteMembers, RevokeInvitation, Update×3) → overall ≥95%. Injection tik (go.mod replace + skip-worktree) + `corpus/services/clerkenstein.md` after the gate.
**Lessons:** the authn alignment is genuinely offline (mint a token from the gene claims → GetUser → compare the extracted identity to the golden), so it needs no live Clerk — confirming TOK-01's easy-side-first ordering. The colony interface compiles offline from the module cache (GOPROXY=off GOSUMDB=off), so the "true drop-in" guarantee holds without GH_PAT.
