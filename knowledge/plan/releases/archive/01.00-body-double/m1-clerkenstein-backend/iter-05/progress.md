**Type:** tik

# iter-05 — standard orgclient → GATE MET

## Close — 2026-06-03

**Outcome:** the 9 standard orgclient genes (InviteMember + BulkInviteMembers `shape`, RevokeInvitation `error_class`, the 3 metadata writes) + goldens. **`alignctl run --gate-overall 95 --gate-critical 100` exits 0: overall 68.4% → 100.0%, critical 100.0% (22/22 genes aligned). THE EXIT GATE FIRES.**
**Type:** tik
**Status:** closed-fixed (planned scope landed; gate reached)
**Gate:** **MET** (overall 100.0% ≥ 95% AND critical 100.0% = 100%)
**Metric delta:** overall 68.4 → **100.0%** (+31.6); critical 100.0% (held). +9 genes aligned.
**Phase 5 grading:** (1) gate-met: **y** — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (4 tiks) — (6) protocol-stop: n — **Outcome: exit-1 (gate-met)**
**Decisions:** none new.
**Routes carried forward:** the **injection tik** (`go.mod replace` whole-colony for authn + the fake-Clerk-API-server for orgclient per M1-D2) + `corpus/services/clerkenstein.md` — these are the post-gate deliverables `/developer-kit:harden-mstone-iters` + `/developer-kit:close-milestone` will surface (the alignment gate fired, but injection + the Delivers→ doc remain milestone scope).
**Lessons:** the whole consumed Clerk backend surface aligns offline with a tiny disarmed mirror (authn JWT + an in-memory orgclient) — the hard part was never the mirror, it's the *injection* (authn = clean go.mod replace; orgclient = needs the fake-API-server, M1-D2). The alignment framework (M0) made "is the mirror faithful?" a one-command 100%.
