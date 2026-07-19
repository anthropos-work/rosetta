# M230 — Progress

## Running ledger

- iter-01 (tok/bootstrap): TOK-01 authored — Option C (FS-as-published fallback via a sha-pinned demo-patch on the ephemeral ant-academy clone) chosen over Option B; seam + vehicle code-verified; baseline framed (0 real cards, F4). closed-fixed. — see iter-01/progress.md
- iter-02 (tik): built the `academy-fs-published-fallback` demo-patch (manifest + native apply/revert helper + `ant-academy.sh` wiring + 14 unit tests, all green; rext tagged `playbill-m230-academy-fs-published`) + **runtime-proven** (59 real cards, 0 draft chips, clone byte-clean) + updated `frontend-tier.md`. closed-fixed; Gate NOT MET (formal cold-/demo-up sweep = user-blocker). — see iter-02/progress.md

## Next-iter queue (the formal gate — needs a go-ahead)

- **Cold `/demo-up` + coverage sweep** (`ANT_ACADEMY` rendered-card descriptor, employee vantage) consuming rext `playbill-m230-academy-fs-published` → assert count ≥ floor, 0 draft chips, 0 prod-ejects. THE formal gate. (Fix is built + runtime-proven; only this heavy verification remains — see the milestone-root USER-BLOCKER record.)
- **next-web clone re-anchor** (2 drifted demopatch manifests: `next-web-public-website-url` + `next-web-studio-url`) — a cold-`/demo-up` prerequisite, unrelated to the academy patch.
- (faithful follow-on) a 2nd manifest for `getPublicCatalogView` (the anonymous `/library` + `/free` routes) — the employee-authed home-grid gate only needs `getServerCatalogView`.

## M230: GATE OUTCOME LEDGER (close Phase 9-iter) — closed-incomplete (pragmatic)

**Gate:** target = cold-`/demo-up` academy rendered-card count (no Draft chip, DB-authoritative, 0 ejects) ·
**achieved = MET-BY-PROXY** (standalone runtime proof: 59 real cards, 0 Draft chips, exact code path, byte-clean
revert; 14 unit tests, flake 3/3) · distance = the formal cold-`/demo-up` sweep only · **status = closed-incomplete
(user pragmatic-close mandate `PROVE-M230-close-on-runtime-proof`).**

**Iter ledger:** iter-01 (tok — TOK-01 chose Option C over B) · iter-02 (tik — built the `academy-fs-published-fallback`
demo-patch + runtime-proved 59 cards / 0 chips, rext tag `playbill-m230-academy-fs-published`). Both closed. No orphan
iters/commits.

**Routes carried forward (Fate-3, see `carry-forward.md`):** (1) formal cold-`/demo-up` card-count proof → **M235/M236**
· (2) local `next-web` clone re-anchor → **M235/M236** demo-up prereq · (3) `getPublicCatalogView` anonymous-routes 2nd
manifest → **M235** next-iter queue. All homed in `m235-prove-it-lands/overview.md § In` (Fate-3 annotation).

**Dropped:** none. **Escape-hatch (cross-release):** none.

**Protocol evolution:** the bootstrap tok's environment-aware option choice (pick the path that can be PROVEN here with
0 platform edits) worked well; the runtime proof measured a rendered-card count (not the M53 port/title check that let
F4 slip). **Deferral audit: GREEN** (3 Fate-3 carries, all homed; 0 repeat-defers; 0 escape-hatch).
