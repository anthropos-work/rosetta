**Type:** tik (run 6, tik 3). Active strategy: TOK-02 (gate (d) first, per iter-14's coupling finding).

# iter-15 — progress

## The fix — a chained anon-view academy demopatch (native, NO Docker re-bake)
Root cause (confirmed in `serverTenant.js`): `getServerCatalogView()` (authed, line 145) IS patched by M230's `academy-fs-published-fallback` (`getBackendCatalogView(eids)`); `getPublicCatalogView()` (anon `/library` + `/free/*` + the cross-port home, line 164) reads `getBackendCatalogView(new Set())` and was left unpatched → empty on a Clerk-free demo. The existing patch's own note flagged this as "a distinct line … out of scope."

Authored **`academy-fs-published-public`** (rext `demo-stack/patches/`) — mirrors the sibling on the `new Set()` line (public-only, no tenant leak), env-gated on the SAME `ACADEMY_DEMO_FS_PUBLISHED` (+ `DEMO_NO_ACADEMY_FILL` opt-out), strips `_draft`/`_origin` (no Draft chip). **CHAINED** on serverTenant.js: `pre_sha256` = the sibling's `post_sha256` (e0f48e81), `post_sha256` = 41cc2d7e. New native helper `stack-injection/apply-academy-fs-published-public.sh` (the pattern of the body helper) + `ant-academy.sh` wiring (apply AFTER FSPUB, revert BEFORE it — reverse chain order). Applied via `next dev` HMR, not a Docker re-bake.

## Verification (local + live)
- **Local chain test** (stack-dev clone): pristine → FSPUB(post1 e0f48e81) → FSPUBLIC(post2 41cc2d7e) → idempotent no-op → revert FSPUBLIC(post1) → revert FSPUB(pristine), git-clean. Chain-order guard REFUSES (exit 2) applying FSPUBLIC on a pristine file. shas computed by applying both patches to the pristine file (deterministic).
- **Live on billion**: re-pinned the rext clone (`stack-demo/rosetta-extensions` b38ad75 → 8391843) + applied FSPUBLIC to the academy clone (serverTenant.js e0f48e81 → 41cc2d7e). `next dev` HMR picked it up.
- **Anon academy renders REAL cards** (browser, tailnet):
  - home `/`: **483 chapter cards** (was empty in iter-14), keyless=false;
  - `/library/`: **37+ visible cards, bodyLen 92998, 0 console errors** (stable 1–20 s);
  - `/free`: **42 cards**, "AI Academy" present;
  - **Draft-chip count 0** on all — the iter-09 `/free` 2-Draft-chip sub-finding RESOLVED (FSPUBLIC strips `_draft`).
- **⇒ Gate (d) MET** — the anon `/library`+`/free` twin both render real cards.

## Gate (d) DISCHARGED → metric 4/8 → 5/8.

## New finding routed to gate (c) — the coverage cross-port has a SECOND, independent blocker
The gate-(c) coverage cross-port to the academy home STILL fails (reachable rose 59→62, but 1 crossPortFailure). It is NO LONGER emptiness: the academy home now renders 483 real cards. The `ANT_ACADEMY_HOME_SECTION` descriptor (`coverage-manifest.ts:712`) requires `mustInclude: ['AI Academy']` in the `main, body` region, but the redesigned academy home's VISIBLE heading is **"Academy"** (+ "AI by Role" + course cards), not the literal "AI Academy" (the string is only in the non-visible `<title>`/curl-HTML). So the marker is STALE vs the real render — a marker-vs-reality mismatch, not a defect. The M219 note that authored this marker (`coverage-manifest.ts:693-711`) targeted the keyless-blank trap (which is NOT occurring — keyless=false, 483 cards). Recalibrating this marker (to a token present in a real render AND absent in a keyless/blank/empty page — while keeping the anti-keyless teeth) is gate-(c) work, routed to iter-16. NB the earlier `/library` browser-blank was a PROBE ARTIFACT (the `/library`→`/library/` redirect race); `/library/` direct renders fine.

## Close — 2026-07-23

**Outcome:** Gate (d) DISCHARGED — the anon academy `/library`+`/free` twin renders real cards live on billion (home 483 / library 37+ / free 42 cards, 0 Draft chips) via the new chained FSPUBLIC demopatch. Primary metric **4/8 → 5/8**. The `/free` Draft-chip sub-finding resolved.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (overall a–h; gate part (d) now MET — 5/8)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (tik 3/5) — (6) protocol-stop: n — Outcome: **continue**
**Decisions:** D1, D2, D3 (iter-15/decisions.md).
**Side-deliverables:** none.
**Routes carried forward:**
- iter-16 = **gate (c)**: recalibrate the `ANT_ACADEMY_HOME_SECTION` coverage marker to the academy's real render (keep the anti-keyless teeth) → re-run coverage to green; then the discrete stack-verify specs; then the 16 Playthroughs LAST (pt-world reset).
**Lessons:** (1) a tok-named "preferred fix" is a hypothesis — but so is a prior iter's coupling claim: iter-14 said gate (d) would unblock the coverage cross-port; it unblocked the EMPTINESS but revealed a SECOND blocker (a stale marker). Fixing one blocker often uncovers the next. (2) `next dev` HMR applies a server-lib demopatch with no Docker re-bake — a whole class of academy fixes is far cheaper than the next-web re-bake path. (3) curl-HTML ≠ browser-innerText: "AI Academy" is in the `<title>` (curl sees it) but not the visible heading (the browser marker doesn't) — measure the gate the way the gate measures.
