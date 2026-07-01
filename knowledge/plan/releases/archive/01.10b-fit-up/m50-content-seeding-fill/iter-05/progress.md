# iter-05 — progress

**Type:** tik (under TOK-01) — Phase C (fix) + D (re-sweep). The fix surface is a stack-snapshot content rewrite.

## Work
- **Phase C (fix).** Added the post-replay Directus content-URL rewrite to `up-injected.sh` (NO_SETDRESS block):
  an idempotent demo-local `UPDATE` regexp_replacing any `https?://<subdomain>anthropos.work` host →
  `http://localhost:$((3000+OFFSET))` in `simulations.{public_landing_page_url,read_more_link}` +
  `skill_paths.public_landing_page_url`. +4 static-fence tests (frontend-build 55→59 GREEN). Committed to rext
  (`fix(M50/05)` ×2 — the base rewrite + the skill_paths/staging broadening). Applied to demo-1: 0
  `anthropos.work` URLs remain in any content field; cleared the cms Redis DB-5 sim cache (the M46 poison class).
- **Phase D (re-sweep, cap=300).** FINAL gate-VALID verdict (frontier exhausted, reachable=69):
  `failingSections=0 escapes=0 personaFailures=0 crossPortFailures=0 notReachedPages=0 gateMet=True` — **0 eject
  pages.** The `1bc8e23c` drill-down (the iter-04 ejector) now follows its rewritten content link to a demo-local
  404 (`localhost:13000/...`, NOT a prod eject — the gate's concern is the eject, not the 404). A mid-finalization
  read showed a transient `escapes=1` from a stale cms cache entry; the bounded re-assert + the cache clear
  resolved it to the clean FINAL verdict.

## Close — 2026-06-30

**Outcome:** The content-URL rewrite (simulations + skill_paths) + the cms cache clear eliminated the manager's
sole residual escape. **Manager gate MET** (reachable=69, frontier-exhausted, all 0). Combined with the
**employee gate MET** (iter-02), the **M42 coverage gate is GREEN on BOTH vantages on the warm demo-1** — the
protocol's primary metric `(failingSections, escapes)=(0,0)` reached on both vantages. All three M50 fixes are
reproducibly baked into the bring-up tooling. The milestone's explicit exit_gate ("on a COLD reset-to-seed
demo") is the remaining acceptance step — reserved for the heavy COLD pass (close/harden + the v1.10b M53
cold-rebuild milestone).
**Type:** tik
**Status:** closed-fixed
**Gate:** MET
**Phase 5 grading:** (1) gate-met: y — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: exit-1
**Decisions:** D1 (content-URL rewrite), D2 (broadened: staging host + skill_paths root), D3 (cms cache clear is load-bearing — M46 poison), D4 (closed-fixed; WARM both-vantage gate MET; COLD reset proof = M53).
**Side-deliverables (if any):** none.
**Routes carried forward (to close/harden + M53, NOT gate-blocking the warm metric):**
- **COLD reset-to-seed acceptance** (the explicit exit_gate): a fresh `/demo-up` (all 3 fixes reproduce from tooling) + both-vantage sweeps → confirm `(0,0)` both vantages on COLD. The v1.10b M53 "cold-rebuild acceptance" milestone owns this; close-milestone should run it (or surface to the orchestrator).
- **Manager manifest-strengthening (D4/F1):** assert the `/enterprise/members` Location column (+ the workforce Growth/Verification/Talent tab contents) so the member-field fill (iter-02) is gate-PROVEN — the data renders (visually confirmed) but the current manifest doesn't ASSERT it.
- The consumption clone (`stack-demo/rosetta-extensions`) carries the iter-04/05 fixes synced for the live verification — authoritatively re-pinned to the `fit-up-m50` tag at close.
- AI-keys policy (F7) + academy (F6): decision deliverables, NOT gate blockers (seeded content renders without live AI; academy AI chat documented-as-absent) — for close/M51.
**Lessons:** (1) The snapshot replays prod content VERBATIM, including absolute prod/staging URLs in content fields — a content-URL rewrite is a missing snapshot transform (the content-side analog of the app-constant injection link-rewriting). Match ANY anthropos.work subdomain (prod + staging) via regex, across ALL content roots that carry the URL field (simulations + skill_paths). (2) A Directus content change is only live once the cms Redis DB-5 sim cache is cleared (the M46 poison) — pair every in-place content rewrite with a cache flush; a fresh COLD /demo-up has no poison. (3) Read the FINAL report, not a mid-finalization snapshot — the harness's bounded re-assert can resolve a transient stale-cache escape after the first measurement.
