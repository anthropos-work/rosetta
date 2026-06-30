# iter-04 — progress

**Type:** tik (under TOK-01) — Phase C (fix) + D (re-sweep). The fix surface is a demopatch (not a seeder).

## Work
- **Phase C (fix).** Authored the `next-web-public-website-url` demopatch (make hardcoded `PUBLIC_WEBSITE_URL`
  read `NEXT_PUBLIC_PUBLIC_WEBSITE_URL`, prod fallback) + wired it into `up-injected.sh` (chained AFTER
  next-web-studio-url on the same urls.ts; LIFO revert) + 2 new tests (self-consistency + chained-hashes-
  reproduce-against-live-clone). demopatch suite 47→49 GREEN, frontend-build 55 GREEN, bash -n clean. Committed
  to rext (`fix(M50/04)`). Built into demo-1: rebuilt `demo-1-next-web` (all 3 patches baked via the lib-seam
  `build_frontend_next_web`; consumption clone left git-clean; `--no-deps` recreate, seed intact). Confirmed
  `localhost:13000` baked into the bundle.
- **Phase D (re-sweep, cap=300).** gate-VALID (frontier exhausted, reachable=68):
  `failingSections=0 personaFailures=0 crossPortFailures=0 escapes=1 → GATE NOT MET`. The demopatch WORKED for
  its class (6 AD sim drill-downs = 0 eject; the constant-built links are now demo-local) BUT the residual
  escape is a DIFFERENT class: `directus.simulations.public_landing_page_url` (+ `read_more_link`) — a hardcoded
  `https://anthropos.work/library/job-simulations/<slug>/` REPLAYED FROM PROD by the snapshot (28 sims have a
  prod `public_landing_page_url`, 14 a prod `read_more_link`). Routes to iter-05 (a stack-snapshot content
  rewrite).

## Close — 2026-06-30

**Outcome:** The `next-web-public-website-url` demopatch landed + tested + built into demo-1 + verified-correct
for the JS-constant link class (cleared the constant-built ejects). But the manager gate did NOT move
(escapes 1→1) because the re-sweep revealed the SOLE residual escape is a SEPARATE class — replayed Directus
content URL fields (`public_landing_page_url`/`read_more_link`), not a JS constant. The residual routes to
iter-05 (a `stack-snapshot` content rewrite) with a definitive root cause + fix surface.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET (employee MET+valid; manager failingSections=0/persona=0/crossPort=0/escapes=1 — the residual escape is the content-field class)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (demopatch built+verified on demo-1), D2 (correct fix, but gate didn't move — residual is a content-field class), D3 (fix surface = stack-snapshot content rewrite → iter-05), D4 (closed-fixed-partial).
**Side-deliverables (if any):** none.
**Routes carried forward:**
- **iter-05 (the residual fix):** a `stack-snapshot` content rewrite of `directus.simulations.public_landing_page_url` / `read_more_link` (+ any other `anthropos.work`-bearing replayed content field) → demo-local host, during replay (or a demo-local post-replay idempotent UPDATE). Re-replay/re-up + manager re-sweep → escapes 0 → manager GATE MET → full M50 gate met on warm demo-1.
- Manager manifest-strengthening (D4/F1): assert the `/enterprise/members` Location column (+ the workforce tab contents) so the member-field fill is gate-proven — pair with re-seeding the member-field fix.
- The member-field fill's gate-verification rides on the above.
- AI-keys policy (F7) + academy (F6): decision deliverables, future tiks.
**Lessons:** (1) An escape host that EQUALS a JS constant's prod value isn't necessarily BUILT from that constant — verify the actual source (the env baked correctly + other same-class pages were clean, yet the escape persisted → it was replayed CONTENT, a Directus field, not the constant). Diagnose the link's true source (DB content field vs JS constant vs hardcode) before assuming the fix surface — the iter-03 triage conflated two escape classes sharing one host. (2) A demopatch + frontend rebuild + re-sweep is a ~25-min heavy loop on the 9 GiB VM; the lib-seam `build_frontend_next_web` drive (mint PK → rm image → build → --no-deps recreate) is the faithful frontend-only rebuild that avoids a full cold `/demo-up`. (3) The snapshot replays prod content VERBATIM, including prod absolute URLs in content fields — a content-URL-rewrite is a missing snapshot transform (the content-side analog of the injection link-rewriting for app constants).
