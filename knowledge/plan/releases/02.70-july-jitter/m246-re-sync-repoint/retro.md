# M246 — Retro

## Summary
M246 was the **HARD go/no-go barrier** that opens v2.7 "july jitter". It re-grounds the demo + rext tooling
on the **CONSOLIDATED platform** (current `origin/main` — skillpath fully decommissioned into `app`, **3
subgraphs**, jobsim standalone) and defuses the imminent seeder break. Delivered: the **load-bearing seeder
re-point** (`skillpath.skill_path_sessions` → `public.skill_path_sessions`, 8 live sites + DNA + ~16 test
assertions, hard-cut per D-1), a **durable canonical `clones.pin.json`** (12 repos @ current origin/main;
skillpath excluded) + a copy-if-absent seam, and the **de-skillpathed LIVE bring-up path** (§3 Fate-1
expansion, D-3 — dropping skillpath from `up-injected.sh`'s `INJECT_SVCS`/`verify_svcs` once the compose
check proved it has no service). **Proven** by ONE cold `/demo-up` (LOCAL demo-2; billion untouched) that
came up GREEN — **561 rows in `public.skill_path_sessions`**, 3 subgraph images + 0 skillpath, health 200 +
casbin 1250, all probes pass. The ~386-commit `app` bump surfaced **no** migration/subgraph break (open
question #2 → NO). Emitted the **9-row confirmed-drift ledger** (D-01..D-09) that scopes M247–M254 against
the proven-green topology. Section milestone, closed-complete; **0 platform-repo edits.**

## Incidents This Cycle
- **None at close.** The close review found **0 must-fix / 0 should-fix** findings — the re-point is a
  mechanical schema-qualifier flip already asserted at every write-site and live-proven at 561 rows, and the
  clone-pin seam was correct. No regressions, no flakes (**5/5** sequential on TestCloneFreshnessM237 incl.
  the 2 new seam tests; full `test_tooling.py` 168/168; frontend_build 94/94; stack-injection 264/9-skipped;
  Go stack-seeding all pkgs ok).
- **P3 (non-M246, carried) — a pre-existing `ResourceWarning`** in `TestMigrateRaceGuard`
  (`test_tooling.py:1413`, an unclosed file in an unrelated test). Not M246-touched, benign, not a failure —
  carried, not fixed (touching unrelated code would violate the barrier's tight scope).

## What Went Well
- **The barrier proved its own worth on the load-bearing axis.** The seeder re-point's target
  (`public.skill_path_sessions` existing on the consolidated platform) was left as an *empirical* claim to be
  proven by the bring-up, not assumed — and the cold demo wrote **561 rows** to it while the legacy
  `skillpath.skill_path_sessions` stayed a 0-row husk. The strongest possible proof that the re-point is real
  and correct.
- **Scoped-as-a-barrier held.** M246 fixed the bring-up mechanism + re-pointed the seeder and *ledgered* the
  drift; it did NOT reach downstream to reconcile the corpus (M247) or fix any fidelity defect. The §3
  expansion (de-skillpath the live path) was genuinely REQUIRED for green (a stale skillpath clone would be
  built + verified), so it landed as Fate-1 — not scope creep — while the inert residue (the `INJECTED`-dict /
  `test_injection.py` / `exposure_claim_guard` skillpath hygiene) was ledgered, not force-fixed.
- **Hard-cut over dual-schema was the right call (D-1).** The demo builds from a single pinned clone set and
  version skew is handled by the per-stack rext **tag pin**, so a dual-schema fallback would have added a
  live schema probe + a new failure mode to serve a transition window the demo never occupies. The hard-cut
  also makes the dangerous skew direction fail LOUDLY (the close adversarial scenario).
- **Right-sized harden + clean close.** One honest harden pass deepened only the 2 genuinely-untested
  clone-pin seam branches (2/4 → 4/4) and added NOTHING to the mechanical re-point; the close confirmed that
  judgment (nothing else worth adding).

## What Didn't
- **Nothing blocking.** The only friction was the honest accounting of the 9 drift rows: several
  (D-02/D-03/D-04 rext hygiene, D-07 perf-patch anchor, D-08 login re-prove) are "in the domain of" a sibling
  but not literally enumerated in that sibling's `In:` list. Resolved cleanly by treating the durable,
  checked-in **drift ledger** (which M247 formally `depends_on`) as the Fate-2 tracking vehicle — no premature
  sibling-overview edits, no silent scope erosion.

## Carried Forward
- **D-01 corpus skillpath→app reconciliation → M247** (Fate-2; explicitly in M247's `In:` list — skillpath.md
  redirect + 4→3 subgraph re-point across ~30 files). **D-06** audit-prose → M247.
- **D-07 AI-readiness `loadMembers` perf-patch anchor went stale in the bump → M250** (its domain; the live
  render loop re-anchors it).
- **D-08 fake-FAPI http-vs-TLS cheap-win probe artifact + end-to-end login not re-run → M251 (probe) / M254
  gate-part (h) (live login re-prove).**
- **D-02/D-03/D-04 inert rext injection-generator skillpath hygiene → M247's triage** of the drift ledger
  (M247 doc-only, so the rext-code cleanup lands wherever triage assigns — M247-adjacent or M251).
- **D-05 stale `stack-demo/skillpath/` clone dir · D-09 academy peripheral → housekeeping / standing** (no
  milestone action required; never re-cloned/built / non-fatal by design).
- All tracked in `drift-ledger.md` + `audit-deferrals/deferral-audit-2026-07-23-m246-close.md` (verdict
  GREEN).

## Metrics Delta
- **Tests (rext, the milestone's real code):** Go `stack-seeding` all pkgs ok; `test_tooling.py` 168/168
  (+2 M246 seam tests); `test_frontend_build.py` 94/94; `stack-injection` 264 (9 skipped). **Flake 0** (5/5
  sequential). Seam branch coverage 2/4 → 4/4.
- **Live proof:** cold demo-2 GREEN — 561 rows `public.skill_path_sessions`, 3 subgraphs, 0 skillpath.
- **Code-of-record:** rext tag `july-jitter-m246-harden` @ `9b29f3a` (commit c9fbf6b) on origin; build
  commits 97585f5 + ee44b9a + 88bcdb8, harden c9fbf6b.
- **Platform-repo edits: 0.** Supply chain: no new deps. Full metrics: `metrics.json`.
