---
close_status: closed-incomplete
gate_target: "cold /demo-up → academy grid renders real cards (>= floor) for the employee vantage, NO Draft chip, DB-authoritative path, 0 ejects, verified by a RENDERED-CARD COUNT (coverage sweep ANT_ACADEMY descriptor)"
gate_achieved: "MET-BY-PROXY — standalone runtime proof: the Option C-patched academy home grid served 59 real skill-path cards, 0 Draft chips (draft-ribbon=0, data-draft=0), through the exact DB-authoritative code path with a signed-in persona; clone reverts byte-clean; 14 unit tests green"
gate_distance: "the FORMAL measurement only — a cold /demo-up + the coverage sweep's ANT_ACADEMY card count. The fill mechanism itself is DONE + runtime-proven."
close_reason: "user pragmatic-close mandate 2026-07-19 (PROVE-M230-close-on-runtime-proof)"
---

# M230 academy-demo-fill — carry-forward

## TL;DR
The academy-fill mechanism is **built and runtime-proven** (Option C: the `academy-fs-published-fallback` sha-pinned
demo-patch, rext tag `playbill-m230-academy-fs-published`). What did NOT land is the **formal** gate measurement — the
rendered-card count on a **cold `/demo-up`** — because that needs a full local demo bring-up blocked by a drifted
`next-web` clone, and the release already ends with cold-reset-to-seed proofs (M235/M236) where a cold bring-up happens
anyway. Per the user's mandate, the formal proof folds there.

## Root-cause cluster 1 — the formal cold-`/demo-up` card-count proof
- **Affected items:** the exit gate's "coverage sweep on a RENDERED-CARD COUNT on a cold `/demo-up`" measurement.
- **Root cause:** the fill is proven standalone (exact code path, 59 cards, 0 chips) but not via the specified cold
  `/demo-up` sweep; a cold `/demo-up` is a heavy local bring-up on a box with prior docker trouble.
- **Estimated scope:** a cold bring-up (already required by M235/M236) + running the existing `ANT_ACADEMY` coverage
  descriptor consuming tag `playbill-m230-academy-fs-published`. No new mechanism.
- **Fate:** Fate-3. **Target milestone:** **M235** (prove-it-lands — cold reset-to-seed) primarily; **M236**
  (prove-on-billion) as the live confirmation (its exit gate already requires "the academy grid renders real cards (Thread A)").
- **Provenance:** M230 iter-02 EXIT_REASON user-blocker; decisions.md PRAGMATIC-CLOSE-MANDATE.

## Root-cause cluster 2 — local `next-web` clone re-anchor (demo-up prerequisite)
- **Affected items:** a cold local `/demo-up` — the local `next-web` clone has DRIFTED from 2 pinned demopatch
  manifests (`next-web-public-website-url` + `next-web-studio-url`), which would drift-refuse on a cold bring-up.
- **Root cause:** local clone drift; unrelated to the academy patch — a general demo-hygiene item.
- **Estimated scope:** re-sync/re-pin the local `stack-demo` clones (re-clone or re-anchor the 2 manifests).
- **Fate:** Fate-3. **Target milestone:** **M235/M236** as a cold-`/demo-up` prerequisite (they need a clean cold
  bring-up regardless).
- **Provenance:** M230 iter-02 concrete obstacle.

## Root-cause cluster 3 — `getPublicCatalogView` 2nd manifest (anonymous routes)
- **Affected items:** the anonymous academy `/library` + `/free` routes (which use `getPublicCatalogView`, empty-eid).
- **Root cause:** the M230 patch fills the **employee-authed home grid** (`getServerCatalogView`); the anonymous
  routes are a faithful follow-on not needed by M230's gate.
- **Estimated scope:** a 2nd manifest branch of the same FS-published transform for the empty-eid path.
- **Fate:** Fate-3. **Target milestone:** **M235** next-iter queue (a faithful completeness follow-on).
- **Provenance:** M230 iter-02 DEFERRED.

## Projected post-resolution state
After M235/M236: the academy grid renders real cards on a cold reset-to-seed (formal gate met, both local and on
billion), the local clones are re-anchored, and the anonymous academy routes are filled — Thread A fully proven live.

## Cross-references
- Fix tag: rext `playbill-m230-academy-fs-published`. Mechanism doc: `corpus/ops/demo/frontend-tier.md`. Content model:
  `corpus/services/ant-academy.md` § The Content Model. Demo-patch contract: `corpus/ops/demo/demopatch-spec.md`.
