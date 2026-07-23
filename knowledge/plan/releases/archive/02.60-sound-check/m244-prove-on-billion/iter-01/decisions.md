# iter-01 — decisions (local)

_(the strategy decision TOK-01 lives in the milestone-root `decisions.md`; intra-iter notes here)_

- **D1 — run the bring-up detached-as-remote-foreground, not on-host-detached.** billion-safety forbids "detached on-host scripts." The cold bring-up (20–50 min) will be driven as a foreground remote process under a single ssh session (backgrounded on THIS workstation via `run_in_background`, polled + heartbeated), so there is one driver and nothing detached on billion. Never kill a mid-build.
- **D2 — KB-fidelity YELLOW is proceed-with-tracking.** The two stale counts (denominator 29→49, spec 39→40) are narrative-only; the implementation reads files. Corpus reconciliation deferred to close so the live 49/49 result is in hand before rewriting corpus historical-example prose.
