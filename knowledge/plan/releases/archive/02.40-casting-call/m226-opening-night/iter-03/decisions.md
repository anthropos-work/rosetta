# iter-03 — intra-iter decisions

| # | Decision | Rationale | Date |
|---|----------|-----------|------|
| D1 | Apply Finding-1's fix via a **surgical `tailscale serve` re-apply** (consume the tag + regenerate/apply the serve plan), NOT a full re-bring-up. | The serve front is a separate step from the image build; the running demo + its baked demo-patches are untouched. Cheap (seconds), unblocks C2/C3/C5 measurement immediately. The reproducible DEFAULT-bring-up proof (serve fix in the default path) is routed to iter-04/05. | 2026-07-17 |
| D2 | Make the candidate-heroes render spec **remote-capable** (`CANDIDATE_HOST`/`CANDIDATE_APP_SCHEME`/`CANDIDATE_OFFSET`), mirroring the render harness's `RENDER_HOST`. Committed+tagged `casting-call-m226-c3-remote` (9396adc). | The C3 harness hardcoded localhost — the same remote-capability gap as Finding-1. Needed to measure C3 from the peer. Defaults keep localhost byte-identical (regression-safe). Runs from this Mac (not consumed on billion). 0 platform edits. | 2026-07-17 |
