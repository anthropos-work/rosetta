# iter-01 — progress

**Type:** tik → **tok (bootstrap)** — authors TOK-01, the milestone's first strategy (per
`build-mstone-iters` Phase 0 rule 1 + coverage-protocol § Iter type selection: "the bootstrap tok resolves the
overview's open questions + takes the baseline reading framing; it does NOT terminate the call").

## What this iter did
1. Loaded the domain context (ant-academy.md § Content Model, the overview's Option ladder, both protocol docs).
2. **Code-verified the Option C seam** against the REAL academy source (`stack-demo/ant-academy/code/src/lib/`):
   `getServerCatalogView()` = `const view = (await getBackendCatalogView(eids)) ?? emptyCatalogView(); return draftsEnabled() ? mergeDrafts(view, eids) : view`.
   The M7 cutover deliberately removed a pre-existing FS-as-published fallback at exactly the `?? emptyCatalogView()`
   expression — Option C restores it (demo-only, un-chipped).
3. **Confirmed the F4 root cause in code:** `ant-academy.sh` sets `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` 0 times.
4. **Confirmed Option C's vehicle exists + is proven:** `demo-stack/patches/demopatch` + `manifest_loader.py` +
   an existing **`ant-academy-dev-origins`** patch (a precedent for patching the ephemeral academy clone).
5. Probed infra feasibility: a cold /demo-up is feasible here (demo-1 images 41h old; docker healthy; disk fine).
6. Ran Phase 0b KB-fidelity → **GREEN**.
7. Authored **TOK-01** (Option C) in the milestone-root `decisions.md`.

## Close — 2026-07-19

**Outcome:** TOK-01 authored — Option C (FS-as-published fallback via a sha-pinned rext demo-patch on the
ephemeral ant-academy clone) chosen over Option B, code-verified seam + vehicle, baseline framed (0 real cards, F4).
**Type:** tok (bootstrap)
**Status:** closed-fixed  (the iter's deliverable — the initial strategy + baseline framing — landed)
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (bootstrap toks never fire this) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (0 tiks) — (6) protocol-stop: n — Outcome: continue (bootstrap tok continues into iter-02 within the same call)
**Decisions:** TOK-01 (milestone-root decisions.md); no intra-iter D-entries.
**Side-deliverables:** none.
**Routes carried forward:** none (iter-02 direction is the TOK-01 Next-tik field).
**Lessons:** The Option C seam and vehicle are both already present + proven (the `?? emptyCatalogView()` fallback
and the `ant-academy-dev-origins` academy-patch precedent), so the fix is a bounded, low-infra-risk demo-patch —
not a new snapshot surface. The heavy risk is entirely in PROVING it (a cold /demo-up), not in building it.
