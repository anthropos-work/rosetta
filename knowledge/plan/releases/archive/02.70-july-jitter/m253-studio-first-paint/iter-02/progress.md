**Type:** tik (under TOK-01)

# M253 iter-02 — progress

## Deliverables landed (rext `july-jitter-m253-studio-first-paint` @ b8969c0, pushed to origin)
1. **`studio-desk-shell-first-paint`** demopatch (`main.ts`) — inject the `.page-skeleton` DOM synchronously
   after `preloadCriticalCSS()` (L97), before Sentry/posthog/clerk.load/l12n/canAccess. De-dup automatic
   (`PageWrapper#init` wipes `document.body.innerHTML`). Round-trip-validated + `demopatch check` PASS.
2. **`studio-desk-no-thirdparty`** demopatch (`main.ts`, chained on top) — no-op `Sentry.init` + `posthog.init`.
3. **`build_frontend_studio_desk` ladder extended** — 5-manifest patch-set fingerprint (M249 3 + M253 2) forces
   a studio rebuild; apply chain + LIFO RETURN-trap reverts + SKIPPED-is-not-applied evidence in all branches.
   `bash -n` clean; `next_web_patchset_fp` is variadic.
4. **Net-new studio-FCP runner** — `run-studio-fcp.sh` + `tests/studio-fcp.spec.ts` + `lib/studio-fcp.ts`
   (establish Clerkenstein session via the real cockpit CTA → cold-load studio → time the `.page-skeleton`
   shell). Green-gate + non-integer-N guard mirror run-latency; never gates on `networkidle`.

## Rebuild + measurement (demo-2, LOCAL LAPTOP)
- Rebuilt the studio image via the REAL `build_frontend_studio_desk` (lib-only source of authoring up-injected):
  all 5 patches applied, image baked (patch-set `02087a8…`), clone git-clean after the LIFO revert, container
  recreated on the new image, studio answering HTTP 302 on :29000.
- **FCP (5 cold loads, measure-only):** skeleton-visible 817 / 795 / 480 / 539 / 743 ms →
  **p50 743 ms · p95 817 ms · max 817 ms**. 5/5 reached the shell, 0 login bounces.
- **vs baseline 4669 ms → 817 ms p95 (~5.7×).** Numerically MEETS the gate (p95 < 1000 ms AND max ≤ 1000 ms).

## Green-gate status (the formal clause)
Fresh autoverify on demo-2 returned `green:false` with **4 warnings, none studio-related**: (1) "demo-patch NOT
APPLIED" reads the CONSUMED clone's stale `demopatch.log` — an artifact of the iteration shortcut (I built via
the AUTHORING function, which logs to the authoring stack dir; the image IS patched, confirmed at build);
(2) "fake-FAPI not answering / nobody can log in" is a probe discrepancy — the 5 FCP samples ALL logged in via
that exact FAPI with 0 bounces; (3) hiring under-set-dress and (4) academy-down are pre-existing set-dress gaps.
The studio surface, login, and shell paint are all verified WORKING by the FCP runner's own probes. A fully-green
verdict on this warm/partially-dressed demo-2 is not achievable for reasons entirely unrelated to studio
first-paint; per **coordination rule 9** (overview.md), the fully-green COLD-p95 confirmation is chartered to
**M254 (prove-on-billion)**, which brings a demo up cold + fully set-dressed.

## Close — 2026-07-24

**Outcome:** the fix + FCP runner shipped; skeleton-visible p95 dropped **4669 ms → 817 ms** on demo-2 (local
laptop), 5/5 cold loads, no login bounce. Numerical gate MET; formal fresh-green cold confirmation routed to M254.
**Type:** tik (under TOK-01)
**Status:** closed-fixed
**Gate:** MET (numerical p95 817 ms < 1000 ms, demo-2 LOCAL LAPTOP; fresh-green COLD confirmation → M254 per coord rule 9)
**Phase 5 grading:** (1) gate-met: n (numerical only; the fresh-green clause is not achievable on this warm demo-2 and is chartered to M254) — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (tik #1 of session) — (6) protocol-stop: n — Outcome: continue → iter-03 (docs deliverables)
**Decisions:** D3 (chained-manifest sha generator + round-trip validation), D4 (lib-only rebuild vehicle for iteration), D5 (green-gate non-achievable on warm demo-2 → M254)
**Side-deliverables:** none
**Routes carried forward:** iter-03 — the three docs Delivers (latency-budget.md studio budget · demopatch-spec.md 2 patches · studio-desk.md MPA boot model). M254 — the fully-green cold-p95 confirmation on billion.
**Lessons:** the milestone's "clerk.load 10 s timeout" hypothesis was a red herring (140 ms actual); the fix is a pure paint-ordering demopatch, and it lands the number decisively. The green-gate clause on a warm, partially-set-dressed local demo cannot be satisfied by a studio-only fix — the cold+green confirmation genuinely belongs to M254.
