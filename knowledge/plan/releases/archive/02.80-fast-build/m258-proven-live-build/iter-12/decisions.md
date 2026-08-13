# M258 iter-12 — decisions

## D59 — The Docker VM's sparse file tracks reclaim here. The reclaim figures are real SSD.

`Docker.raw` is **apparent 137.44 GB / allocated 50.68 GB**, against a docker-logical total of
**51.56 GB** — agreement to **1.7 %**. So on this host, freeing bytes inside the VM frees bytes on the
SSD, and iter-11's 5.297 GB was a real 5.297 GB.

This was **not** safe to assume. The common failure is a VM disk that only grows: `prune` frees space
inside the VM, the host file never shrinks, and every reclaim number in the milestone is fiction. It had
to be measured before any space strategy could be trusted, and it is now the foundation `TOK-02` rests
on. Sibling of `D53`: **`ls -l` on `Docker.raw` reads 137.44 GB — 2.7× the truth.** Quote `du`, or
`system df`, never apparent size.

## D60 — Space partitions by its COUPLING TO TIME, and the partition is the strategy.

The user's constraint is not a tiebreaker applied after choosing targets; applied first it *sorts* them:

- **Class A — coupling zero.** Orphaned volumes, dead images, leftover stack dirs, a stack that is going
  away anyway. Reclaiming costs no build second. Take all of it.
- **Class B — coupling favourable.** Image size. Export/unpack is size-proportional at a measured
  **5.73–8.05 s/GB** on this host class (`build-budget.md:65,167`), so a smaller image is *also* a faster
  build. The constraint does not bind here; it rewards.
- **Class C — coupling adverse.** The build cache. **19.28 GB reclaimable — the largest single reclaim on
  the box — at a measured 173 s per 356.8 MB evicted.** Forbidden as a default (`D58`); `--filter
  until=24h` never `-af`; any policy argued on both axes with measurements.

The whole strategy is: **spend A and B, which is most of the win and none of the cost; never spend C by
default.**

## D61 — studio-desk must PRUNE-AND-COPY, not re-install. The time constraint chose the fix.

The obvious multi-stage gives the runner a fresh base and `npm ci --omit=dev`. It wins the space and
**buys it with time** — a second dependency resolution and fetch on every cold build. L1 avoided this
only because `next build` *emits* `.next/standalone` as a build product, so its runner copies and never
installs (`next-web.Dockerfile:118`). studio-desk has no standalone emitter.

The equivalent that keeps the constraint: **`npm prune --omit=dev` in the builder, then
`COPY --from=builder` the pruned `node_modules`.** One resolution, one fetch, same final tree.

Recorded because the space-only reading reaches the wrong Dockerfile, and it would have been graded a
win.

## D62 — studio-desk's time prize is ~7–10 s, not 115 s. Trim the claim to the arithmetic.

iter-11 routed studio-desk as *"the two axes converge here: 1.7 GB × 2 and 115.35 s cold"*, which reads
as though multi-staging recovers 115 s. It does not. The layer census says the image is
`npm ci` **1.04 GB** + build output **63.2 MB** + source **61.8 MB** + base **~162 MB**; a prune-and-copy
runner removes the devDependency share of that 1.04 GB, call it **~1.25 GB**. At the measured
**5.73–8.05 s/GB** that is **≈ 7–10 s** off the export/unpack leg.

The remaining ~105 s of the 115.35 s is `npm ci` + `tsc` + `vite build` — work the runtime image's shape
cannot remove. **Space win large, time win modest and positive.** Stated now so the tik is graded against
7–10 s and not against a number nobody derived.

## D63 — iter-11's `demo-4` orphan is 8 KB, not a space finding.

iter-11's close listed the orphan `stacks/demo-4/` alongside the 4.2 GB of host-side stack dirs, which
reads as though it contributes. It is **8.0 KB**. It remains a true instance of the F-9 defect
(`ROUTE-M258-iter02-purge-did-not-clear-the-stack-dir`) and stays routed as a *hygiene* item — the
correction is only that it must never be quoted as space. The 4.2 GB is real and unaffected.
