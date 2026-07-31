**Type:** tok (bootstrap)

Iter-shape selection per `/developer-kit:build-mstone-iters` Phase 0 **rule 1** (the unconditional
iter-01 bootstrap-tok rule), not a protocol refinement.

## Work

1. **Phase 0b pre-flight KB-fidelity gate — RAN, returned RED, blocked strategy authoring.**
   The audit found M257's own `iteration_protocol_ref` (`corpus/ops/demo/build-budget.md`) asserting in
   six places that the v2.8 gate is measured on `billion` against a `666.29 s` baseline — both superseded
   by `D-v28-14` hours earlier. Report: [`../kb-fidelity-audit.md`](../kb-fidelity-audit.md).
2. **Independently re-verified the audit's three load-bearing claims** before escalating, rather than
   escalating on one reading. All three held — including the mechanical one (the baseline mirror fence
   pins `billion.json` and explicitly fences M257's own `overview.md` in the very band odysseus's p50 must
   be written into).
3. **Reconned odysseus read-only** (`spec-notes.md` F1–F4) and **resolved the audit's highest-value
   unknown in the gate's favour**: it is a containerd-image-store host, so the unpack leg is paid and L1
   keeps its full price. Two risks the audit could not have had, because they needed the host: a
   **truly-cold** Docker (0 images / 0 cache) and **zero swap**.
4. **Surfaced three contract decisions**, answered as `D120` / `D121` / `D122`.
5. **Cleared the RED** against those decisions — the host-parameterised fence, the minimal host
   corrections, the re-homed re-confirmation, the mandatory warm-up rule, the re-anchored §8.5 list, the
   corrected prerequisite, and the ready-to-execute fix list.
6. **Re-ran the gate: YELLOW, prior RED CLEARED** (11 CLOSED / 1 PARTIAL / 2 tracked-YELLOW, 0 blockers).
7. **Authored `TOK-01`** in the milestone-root `decisions.md`.

## Close — 2026-07-31

**Outcome:** `TOK-01: instrument before baseline, baseline before levers` authored against **verified**
knowledge — which required clearing a RED pre-flight first. The opening strategy is sequencing-driven
because the gate's distance is **unknown**: odysseus's baseline does not exist, and the release's own rule
forbids inheriting billion's.
**Type:** tok
**Status:** closed-fixed
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n *(bootstrap toks never fire this exit)* — (3) re-scope: n — (4) user-blocker: n *(the RED-blocker raised mid-iter was answered as D120/D121/D122 and is resolved)* — (5) cap-reached: n *(toks do not count toward the 5-tik cap)* — (6) protocol-stop: n — **Outcome: continue**
**Decisions:** `D120` / `D121` / `D122` (coordinator-answered, milestone-root) · `TOK-01` (milestone-root) ·
`USER-BLOCKER-2026-07-31` (raised and resolved within this iter)
**Side-deliverables:**
- The host-parameterised **baseline mirror fence** (rext `3f27bf0`) — 4 → 28 tests. It found **12
  un-hosted baseline claims** on its first live run, one of them `state.md:4`: M257's own gate line
  inheriting `666.29 s` with no host named. **The fence caught the exact defect its own docstring
  describes, in the release that wrote it.**
- Three **machine-readable** false claims corrected (`billion.json` role + `gated_baseline.note`,
  `laptop.json` storage_driver) — the last of which called L9 *"a billion-only phenomenon"*, the misread
  that would have mis-priced L1 on this host.
- `odysseus` went from **0 corpus occurrences to 21** in `build-budget.md`.
- The rext README's copy-pasteable campaign pointed at the **off-limits** host, and was unrunnable as
  written; both fixed.
- `name` added to the profile loader's required keys; two tests that hardcoded `("billion","laptop")` now
  glob, so `odysseus.json` will ship **validated** rather than untested.
**Routes carried forward:** none as deferrals. The Phase 0b YELLOW residuals are carried as
**`TOK-01` § Known-context #1–#9**, which is where a strategy's known-context belongs — not as a
deferral queue. Two are scheduled work: known-context #1 and #2 are iter-02 deliverables (b).
**Lessons:**
- **Phase 0b earned its place in the protocol.** Had the bootstrap tok authored strategy first, it would
  have priced levers against a baseline the release had disowned hours earlier — this release's own hunted
  defect, committed in the strategy rather than the code. The gate ran *before* strategy for exactly this
  reason and it fired.
- **Re-verify an audit's load-bearing claims before acting on them.** The audit was right on all three
  checked, but it was also wrong elsewhere in the same report (four §8.5 anchors, two line cites, one
  claim of a `provisional_fields` mechanism that does not exist). An audit is evidence, not a verdict.
- **Two agreeing weak signals are not one strong signal.** The F4 retraction: `ssh host 'cmd'` reported
  Go absent, and `~` having no Go dir plus atlas genuinely missing *felt* corroborating. Both were
  consistent with either hypothesis. The disproof needed a **different kind** of evidence — the
  filesystem — not more of the same kind. Recorded in `spec-notes.md` rather than overwritten, because
  the error is the lesson.
- **A fence written to protect a release can block it, and that is not a reason to weaken it.** D120
  parameterised the fence by host instead of retiring billion's baseline, which made it *stronger*: it
  promotes *state the environment with every number* from prose convention into a machine check.
