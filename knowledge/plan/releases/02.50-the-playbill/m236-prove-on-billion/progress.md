# M236 — Progress

## Running ledger

- iter-01 (tok/bootstrap): TOK-01 "publish-then-prove" authored; baseline measured live — gate denominator is **31** landable (session × action) pairs, currently **0/31**, blocked by an unpublished-tooling gap (`billion` pins the M228 tag; 0 of 13 `playbill-*` tags are on origin) — see iter-01/progress.md

- iter-02 (tik): **IN-FLIGHT, NOT CLOSED** — Phase P step 1/5 landed (`billion` build cache pruned, 109 GB reclaimed, 40 G → 139 G free); steps 2–5 (publish + re-pin) halted by a RED Phase 0b verdict — see iter-02/progress.md

## Next-iter queue

- **BLOCKED — USER-BLOCKER-M236-01** (see decisions.md): the exit gate contains two clauses that cannot be proven/measured as written (tailnet-only reach; p95 click→ACCESS for content seats), a cited page-object that does not exist, a documented coverage rule M236 must reverse, and a hollow `iteration_protocol_ref`. **User decision needed on gate re-scope before Phase P resumes.**
- **iter-02 resumption (Phase P steps 2–5):** publish `rosetta-extensions` `main` + 13 `playbill-*` tags; re-pin `.agentspace/rext.tag` → `playbill-m235-hardened`; verify the M217 FATAL pin guard. Verified safe and ready; re-run the B2 freshness check first.
- **Phase 0b gate:** CONSUMED — verdict RED, recorded in spec-notes.md, full report in kb-fidelity-audit.md.
