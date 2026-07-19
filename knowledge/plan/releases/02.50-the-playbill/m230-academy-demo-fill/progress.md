# M230 — Progress

## Running ledger

- iter-01 (tok/bootstrap): TOK-01 authored — Option C (FS-as-published fallback via a sha-pinned demo-patch on the ephemeral ant-academy clone) chosen over Option B; seam + vehicle code-verified; baseline framed (0 real cards, F4). closed-fixed. — see iter-01/progress.md
- iter-02 (tik): built the `academy-fs-published-fallback` demo-patch (manifest + native apply/revert helper + `ant-academy.sh` wiring + 14 unit tests, all green; rext tagged `playbill-m230-academy-fs-published`) + **runtime-proven** (59 real cards, 0 draft chips, clone byte-clean) + updated `frontend-tier.md`. closed-fixed; Gate NOT MET (formal cold-/demo-up sweep = user-blocker). — see iter-02/progress.md

## Next-iter queue (the formal gate — needs a go-ahead)

- **Cold `/demo-up` + coverage sweep** (`ANT_ACADEMY` rendered-card descriptor, employee vantage) consuming rext `playbill-m230-academy-fs-published` → assert count ≥ floor, 0 draft chips, 0 prod-ejects. THE formal gate. (Fix is built + runtime-proven; only this heavy verification remains — see the milestone-root USER-BLOCKER record.)
- **next-web clone re-anchor** (2 drifted demopatch manifests: `next-web-public-website-url` + `next-web-studio-url`) — a cold-`/demo-up` prerequisite, unrelated to the academy patch.
- (faithful follow-on) a 2nd manifest for `getPublicCatalogView` (the anonymous `/library` + `/free` routes) — the employee-authed home-grid gate only needs `getServerCatalogView`.
