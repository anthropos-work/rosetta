---
milestone: M257x
---

# M257x — spec notes

## The recurring class this milestone must end

Three occurrences, one shape: **the platform consolidates a service into `app`, and rext keeps writing to the
schema that service owned.**

| release | service | how it surfaced |
|---|---|---|
| v2.1 | skiller → app | seeder broke |
| v2.7 | skillpath → app | seeder broke again; corpus asserted skillpath live Tier-1 in ~30 files |
| **now** | jobsimulation (+ cms, roadrunner "own no local schema") | **latent** — only because our clones are stale |

Each time the fix was re-derived from scratch. The deliverable that breaks the cycle is not the re-point — it
is the **fence** (clause 4) plus the **written procedure** (`corpus/ops/platform-alignment.md`).

## Measurement discipline inherited from this release

- **State the environment with every number** — the baseline mirror fence is now parameterised by host and
  FAILS a baseline-shaped claim that names no host (D120).
- **Prove a check can go RED before trusting it.** M256 found 43 checks that reported success without checking;
  M257 found the gate's own health check reading a dropped table behind a swallowed error.
- **A cold cycle is the only honest test.** B1 and B2 were both invisible to warm cycles for four days.
