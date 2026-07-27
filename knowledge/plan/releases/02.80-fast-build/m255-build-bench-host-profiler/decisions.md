# M255 — decisions

Release-level binding decisions **D-v28-1 … D-v28-11** live in
[`../../../roadmap.md`](../../../roadmap.md) § Active — v2.8.

## To be recorded during the build

- **D-M255-1 (required, item 3):** reconcile the headroom assert's *"fail loudly"* against the codebase's
  standing **never-block-a-bring-up** pre-flight contract (`up-injected.sh:279`, and the advisory-since-M19
  `preflight_vm_ram`/disk shape). Expected resolution: the assert gates **buildbench and the M257 gate**, and
  does **not** block an operator's bring-up — but this must be written down, because M255 also retracts the
  advisory pre-flight as "cosmetic", and the two statements have to be consistent.
- **D-M255-2 (expected, spike a):** whether L1 lands via `ENV NEXT_PRIVATE_STANDALONE=1` (a Next-**private**
  API, zero config edits) or via the `next.config.mjs` demopatch fallback.
