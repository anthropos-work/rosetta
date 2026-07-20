# Hardening Ledger — M230 academy-demo-fill

## Pass 1 — 2026-07-19 — final

**Iters hardened this pass:** all milestone-touched code (iter-01 tok · iter-02 tik).
**Tiks covered since prior pass:** all iters in milestone (first + only harden pass).

**Milestone-touched code footprint (rext `.agentspace/rosetta-extensions/`, tag `playbill-m230-academy-fs-published`):**
- `demo-stack/patches/academy-fs-published-fallback/academy-fs-published-fallback.yaml` — the sha-pinned demo-patch manifest
- `demo-stack/ant-academy.sh` — the apply-before-launch / revert-on-stop wiring + opt-out knob
- `demo-stack/tests/test_academy_fs_published.py` — 14 unit tests
(rosetta side: `corpus/ops/demo/frontend-tier.md` doc update — no code)

**Coverage / dimension scan (final-mode, against the demo-patch surface):**
- **Test depth (14 tests):** TestManifest (schema · marker-in-replacement-not-anchor · shas hex64 · anchor→post_sha),
  TestLadder (apply/revert roundtrip + idempotent · **drift-refuse**), TestStripTransform (removes draft tags · matches
  manifest replacement), TestLauncherWiring (helper defined · apply-before-launch · env-in-launch · revert-on-stop ·
  opt-out knob · executable). This is the full demopatch-spec G1–G7 guard surface for this patch.
- **Error paths:** the drift-refuse guard is exercised (a patch on a drifted clone REFUSES rather than silently
  shipping unpatched — the M217 defect class the demopatch mechanism exists to prevent).
- **Edge cases:** idempotent re-apply + revert-to-byte-clean are pinned (the clone must be left git-clean).
- **Integration / runtime:** iter-02's standalone runtime proof (patched grid → 59 real cards, 0 Draft chips, exact
  DB-authoritative code path, byte-clean revert) IS the integration test; the formal cold-`/demo-up` card-count is
  carried forward to M235/M236 per the pragmatic-close mandate.
- **Fuzzing / perf:** n/a (a manifest-driven patch, no non-trivial input surface or perf-sensitive path).

**Bugs surfaced + fixed inline:** none.
**Flakes stabilized:** none — flake gate 3/3 clean.
**Knowledge backfill:** the shipped mechanism + the corrected F4 attribution already landed in `frontend-tier.md`
(iter-02 DELIVERS); the demopatch contract lives in `demopatch-spec.md`.

**Stop condition:** **stabilized** — 14 comprehensive tests covering the full patch-guard surface, flake gate 3/3,
runtime-proven via the exact code path, no new gaps. Satisfies `/developer-kit:close-milestone`'s iterative
final-harden gate. (Milestone closes closed-incomplete/pragmatic on the runtime proof; the formal cold gate is
carried to M235/M236.)
