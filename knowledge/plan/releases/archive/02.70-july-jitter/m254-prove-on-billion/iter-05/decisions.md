# M254 · iter-05 — decisions

## D1 — (h)-latency p95 < 5 s MET both vantages (live, solo)
`run-latency.sh 1 {employee|manager}` from this tailnet peer against billion's HTTPS origins,
`LATENCY_SCHEME=https`, `LATENCY_HOST=billion.taildc510.ts.net`, `LATENCY_AUTOVERIFY_JSON`=scratch copy of
billion's fresh green verdict (06:25:41Z, 13 min old), `LATENCY_GATE_MS=5000`, 5 cold runs, SOLO (quiet system).
Employee (maya-thriving): p50 0.92 s, **p95 1.43 s**, 5/5 ACCESS. Manager (dan-manager): p50 0.74 s,
**p95 1.41 s**, 5/5 ACCESS. Both PASS the < 5 s gate. The per-run ERR_ABORTED anomalies are benign
RSC-prefetch aborts on navigation (the harness grades ACCESS = loader-gone + identity-present, not prefetch).

## D2 — (c)-render: web + hiring LIVE-confirmed; all 4 apps baked-to-stack
Live authenticated DOM checks (temp Playwright specs, cockpit login flow, removed after): **next-web** —
account dropdown (the "Maya" user button) carries "Back to Cockpit", `href` = `https://…:17700` ×2;
**hiring** (rae-recruiter, shared `packages/ui/NavbarTop`) — same, text=1, href → :17700.
Resolve-to-stack confirmed for **all 4** apps at build: `:17700` baked in next-web `.next/static` chunk,
hiring `.next/static`, studio `dist/public/assets/main-BPJCXUm9.js`, academy `.env.local
NEXT_PUBLIC_COCKPIT_URL`. All 4 back-to-cockpit demopatches applied at bring-up (iter03 log). Prod-eject side
(0 escapes/133 pages) already proven iter-04.

## D3 — (c)-academy DEFECT: Back-to-Cockpit item not durable on a long-running native academy (routed forward)
The academy user-menu opens live but carries NO Back-to-Cockpit item. Root cause: the running clone's
`code/src/components/UserMenu.jsx` sha = **exactly `pre_sha256`** (`d423203e…`) = fully **pristine** (the patch
was reverted out-of-band ~9 h into the run — the fresh-launch path DID apply it at bring-up + `next dev`
compiled it, so it rendered initially). The heal (`reapply_clone_patches`, ant-academy.sh:170) is only invoked
from the **reconcile branch** (ant-academy.sh:333) when `ant-academy.sh <N>` is re-invoked — which has not
happened since bring-up. So a demo that's been up a while (or through a coverage/autoverify cycle that reverted
the clone) silently drops the academy Back-to-Cockpit item. **Docker apps are immune** (the patch is baked into
an immutable image); the NATIVE academy serves the mutable clone → only as durable as the clone stays patched.
Fix surface: rext `ant-academy.sh` (durable re-heal — e.g. periodic reapply, or a heal invoked by the
verify/coverage cycle). Routed to the (f)/academy fix-iter; needs an academy re-invoke/re-bring-up to prove.
This is the "provisioned-wrong / fragile-provisioning" class the live-prove exists to catch (0 platform edits).

## D4 — (c)-studio render is (f)-coupled
studio-desk's back-to-cockpit patch is applied + `:17700` baked in `dist`, so the item is correctly BUILT.
Its LIVE render can't be confirmed until (f) is fixed — on `--public-host`, `:19000` bounces to `:13000/login`
(the studio shell never paints), so the studio menu is unreachable. Folds into the (f) fix-iter re-bring-up.
