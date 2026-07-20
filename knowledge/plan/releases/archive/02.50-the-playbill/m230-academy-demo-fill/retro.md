# M230 "academy demo-fill" — Retro

## Summary
Built + runtime-proved the production-faithful academy fill (Option C: the `academy-fs-published-fallback` sha-pinned
demo-patch on the demo's own ephemeral clone) — the patched grid served 59 real skill-path cards, 0 Draft chips,
through the exact DB-authoritative code path. Closed **incomplete** (pragmatic): the fill works, but the FORMAL
cold-`/demo-up` card-count gate folds to M235/M236 (which do cold bring-ups anyway).

## Incidents this cycle
- **User-blocker (not a defect):** the formal gate needs a cold `/demo-up`, blocked locally by a drifted `next-web`
  clone (2 demopatch manifests would drift-refuse). Surfaced honestly by the bootstrap-tok's environment assessment
  rather than burning hours on a heavy local bring-up. Routed to M235/M236.

## What went well
- The bootstrap tok chose Option C (demo-patch, no prod-DB dependency) over Option B (snapshot surface) precisely
  because it could be PROVEN in this environment with zero platform edits — a good environment-aware strategy call.
- The runtime proof measured the RIGHT thing (a rendered-card count + 0 draft-chip assertion via the exact code
  path) — not the M53 port-serves + SSR-title check that let F4 slip in the first place.

## What didn't
- The formal cold-`/demo-up` gate couldn't land locally (next-web clone drift + heavy bring-up). The lesson is the
  release's own shape absorbs it: M235/M236 do cold bring-ups, so the formal academy render proves there for free.

## Carried forward
- Formal cold-`/demo-up` card-count proof → M235/M236. next-web re-anchor → M235/M236 demo-up prereq.
  `getPublicCatalogView` anonymous-routes 2nd manifest → M235 next-iter queue. See `carry-forward.md`.

## Metrics delta
- 14 rext unit tests (flake 3/3) + runtime proof (59 cards, 0 chips) · 0 platform edits · fix tag
  `playbill-m230-academy-fs-published`. See `metrics.json`.
