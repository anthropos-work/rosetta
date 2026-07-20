# iter-10 — decisions

## D1 — the chain reproduces cold, with zero intervention

**Procedure.** Re-pin `billion` → `playbill-m236-latency-tz-fix` (the final tooling tag) ·
`rosetta-demo down 1 --purge` (containers + network + container-owned data) · cold
`./up-injected.sh 1 --public-host billion.taildc510.ts.net` · re-measure all three components.

**Result.** Bring-up green on its own tail verdict (`demo-patches: all applied (none refused, none
skipped)`, `frontend builds: ok (the running images are this run's)`, `autoverify demo-1: OK`), and:

| | cold reading |
|---|---|
| content-stories sweep | 29/29 |
| academy grid | 65 course links, 0 Draft chips |
| hero p95 | employee 1.22 s · manager 1.51 s, 5/5 ACCESS each |

**Why this decides the gate.** Iters 04–09 measured a stack mutated in place — manifest re-exported twice,
cockpit restarted three times, rext clone re-pinned twice. Any of those readings could have depended on
hand-applied state. They did not: the corrected manifest (skill-path manager views removed, academy on
`/courses/`) arrives from **published tooling consumed by tag**, and every fix this milestone landed is in
that tooling. Reproducibility here is satisfied by construction rather than by narration.

**iter-07's prediction confirmed.** The cockpit binds `0.0.0.0:17700` on a cold cycle; the `127.0.0.1`
bind was an artifact of restarting in place after `tailscale serve` had taken the port, exactly as that
iter's close predicted. One `ss` line, no investigation.

## D2 — the cold stack measured faster than the warm one; report both

Cold: employee **1.22 s** / manager **1.51 s**. Warm (iter-09): **3.15 s** / **2.71 s**.

The counterintuitive direction is explicable — the warm stack had been up for hours across three cockpit
restarts and two re-pins, while the cold one runs clean caches and this run's own frontend images.

**Decision.** Record both, and name iter-09's as the pessimistic pair. The temptation is to report the
better number now that the gate is met; the cost of doing so is that the next reader treats ~1.2 s as the
expected steady state and reads a perfectly normal 3 s as a regression. `latency-budget.md`'s standing
instruction — state the environment with every number — cuts this way too.

## D3 — "0 platform edits" verified per-clone, not asserted

Every Go service clone and `next-web-app`: **0 modified**. Two clones are not clean, and both are
accounted for rather than rounded away:

- **`cms`: `?? studio/`** — untracked, the `anthropos-studio-room` clone `make init-studio` creates. No
  tracked file is modified.
- **`ant-academy`: 4 modified** — `serverTenant.js` (the M230 `academy-fs-published-fallback`),
  `next.config.js` (the M212–215 `allowedDevOrigins` tailnet patch), and `public/catalog.json` +
  `public/content/index.md` (build-generated).

The academy patches are the **sanctioned live-patch mechanism**: ant-academy runs natively under
`next dev`, so its patches must persist for the process lifetime and revert on `--stop` — applied to the
demo's **ephemeral, gitignored** clone, never to `anthropos-work/ant-academy`. The bring-up's own
`demo-patches: all applied (none refused, none skipped)` covers them.

**The gate's meaning is "the canonical repos are never touched," and that holds.** Stating it as a bare
"0" would have been true-sounding and unauditable; naming the four files and why they are there is what
lets a reviewer disagree with the judgement if they want to.
