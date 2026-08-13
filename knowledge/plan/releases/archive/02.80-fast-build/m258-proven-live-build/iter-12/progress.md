# M258 iter-12 — progress

**Type:** tok (bootstrap flavor, user-directed — see `overview.md` for why not `triggered`)

Measured 2026-08-12 10:06–10:12Z on `macmini` (Apple M4 Pro, arm64, Docker Desktop VM, **overlayfs**),
`load1` **19.76**, three stacks resident (`demo-1` 11 · `demo-2` 11 · dev 5). Free host disk **173 GiB**.

## Phase A — re-survey: verify, don't inherit

iter-11's reclaim **held**, and this is the confirming re-read, not a repeat of its claim:

```
TYPE            TOTAL   ACTIVE   SIZE       RECLAIMABLE
Images          31      22       23.83 GB   1.754 GB (7 %)
Containers      27      27       280.8 MB   0 B
Local Volumes    6       6       0 B        0 B          ← was 184 / 5.297 GB at iter-11
Build Cache    123      14       27.45 GB   19.28 GB
```

Volumes are still **6**, still **0 B**. So the 5.297 GB did not come back in the ~10 minutes since, and
`D55`'s producer model survives its first re-check: nothing has *started* a Postgres container since.

## Phase B — the axis nobody has measured: what this costs the HOST

Every figure above is **logical, inside the VM**. On this host Docker is a VM whose disk is one sparse
file, and the honest question — *how many bytes of the user's SSD does this cost* — has never been asked
in this milestone.

```
~/Library/Containers/com.docker.docker/Data/vms/0/data/Docker.raw
    apparent  = 137.44 GB     ← the VM's virtual CAPACITY. Not consumption.
    allocated =  50.68 GB     ← what the SSD actually gives up
```

**Docker's own logical total is 51.56 GB** (23.83 images + 0.28 containers + 27.45 cache). Allocated is
**50.68 GB**. They agree to within **1.7 %**, and that is the finding:

> **On this host the sparse file TRACKS reclaim — so in-VM reclaim really does return SSD.**

This had to be measured because the opposite is the common case: a VM disk that only ever grows, where
`docker system prune` frees megabytes inside the VM and **zero** bytes on the host, and every reclaim
figure in this milestone would have been fiction. It is not fiction here. But note the second trap sitting
right beside the first (`D53`'s sibling): `ls -l` on `Docker.raw` reads **137.44 GB**, 2.7× the truth.

**Total M258-attributable disk = 50.68 GB (Docker) + 12.8 GB (host tree) ≈ 63.5 GB.**

## Phase C — the host tree, broken down

| path | size | note |
|---|---|---|
| `stack-demo/` | **7.1 GB** | of which `stacks/` **4.2 GB** + `ant-academy/` **2.4 GB** |
| `stack-dev/` | 3.8 GB | the user's dev stack clone set |
| `.agentspace/` | 1.9 GB | authoring clone + its 263 MB leftover `stacks/demo-1` |

Inside one stack dir (`demo-1`, 2.2 GB):

| | size | reclaimable on teardown? |
|---|---|---|
| `data/` | **1.9 GB** | yes — it IS the database |
| `clones/` | 220 MB | **yes, and today it is NOT reclaimed** — the F-9 defect |
| `bin/` | 37 MB | yes |
| `fake-fapi` + `fake-bapi` | 18.5 MB | yes |

⚠️ **One honest correction to iter-11.** It listed the orphan `stacks/demo-4/` beside the 4.2 GB as
though both were space findings. `demo-4/` is **8.0 KB**. It is a real hygiene defect and a real
instance of F-9, but it is **not** a space finding and must not be quoted as one. The 4.2 GB is real;
the orphan is 8 KB.

## Phase D — the image axis, and the trap re-verified

`D53` re-confirmed from a second direction. The four `m257-*:probe` leftovers read **8.88 GB** in the
`docker images` SIZE column; `system df` says **1.754 GB** is reclaimable across *all* 9 unreferenced
images. The mechanism is visible in the tags: `m257-warmup-next-web` and `m257-old-next-web` are both
4.04 GB and both are pre-L1 siblings of `demo-2-next-web:latest` (4.04 GB, **in use**).

**The L1 win, and the one image it never reached:**

| image | demo-2 (pre-L1) | demo-1 (post-L1) | ratio |
|---|---|---|---|
| next-web | 4.04 GB | **417 MB** | 9.7× |
| hiring | 3.94 GB | **380 MB** | 10.4× |
| **studio-desk** | **1.7 GB** | **1.7 GB** | **1.0× — untouched** |

Post-L1, studio-desk is the **largest UI image by 4×**. Its layers say exactly why:

```
1.04 GB  RUN npm ci                       ← 61 % of the image
 63.2 MB RUN npm run build:server && build:frontend   ← what the 1.04 GB exists to produce
 61.8 MB COPY . .
 ~162 MB node:24-alpine base
```

`Dockerfile` is **single-stage** — `npm ci` installs all **63** dependencies (30 prod + **33 dev**:
typescript, vite, rollup and their trees) and then **ships them**, to run a `CMD ["npm","start"]` whose
`start` is `node dist/index.js`. A plain Node process. The toolchain that produced the 63.2 MB is
carried forever beside it. This is verbatim the shape L1's own header records for next-web
(*"2,630 MB of node_modules shipped to serve a 241 MB build output"*).

## Phase E — the strategy, and the reason it is not just "delete things"

The user's constraint (`D58`, his own words: *"not at compromise of time … account for this on the cache
consideration"*) is not a tiebreaker to apply at the end. It **partitions the space**, and the partition
is the strategy. `TOK-02` is authored in the milestone-root `decisions.md`; its spine:

| class | coupling to time | measured basis | policy |
|---|---|---|---|
| **A** | **zero** | reclaiming costs no build second | take all of it |
| **B** | **favourable** | export/unpack is size-proportional at **5.73–8.05 s/GB** on this host class (`build-budget.md:65,167`) | this is where the leverage is |
| **C** | **adverse** | one 356.8 MB eviction cost **173 s** (`build-budget.md`) | **forbidden as a default** |

**Class C is 19.28 GB — the single largest reclaim on the box, and the one we do not take.** That is the
constraint doing real work rather than being quoted.

### The sharpened design rule the pricing produced

Pricing studio-desk on **both** axes changed its fix. The naive multi-stage — a fresh runner that runs
`npm ci --omit=dev` — wins the space but **buys it with time**, because a second install is a second
install. L1 escaped this only because `next build` *emits* `.next/standalone` as a build product, so its
runner copies and never installs. studio-desk has no standalone emitter, so the equivalent is:

> **prune-and-copy, never re-install:** `npm prune --omit=dev` in the builder, then `COPY --from=builder`
> the already-populated tree.

Same space win, no second resolution, no second network fetch. **The constraint is what found this** — a
space-only reading would have shipped the re-install form and made cold builds slower.

⚠️ **And the time claim is trimmed to what the arithmetic supports.** iter-11 routed studio-desk as
*"the two axes converge"* on **1.7 GB × 2 and 115.35 s cold**. The space win is large and real
(~1.2–1.3 GB/stack). The **time** win is the export/unpack leg only: ~1.25 GB removed × 5.73–8.05 s/GB =
**≈ 7–10 s**, not 115 s. The other ~105 s is `npm ci` + `tsc` + `vite build`, which a multi-stage does
not remove. **Predicting a 115 s win here would be exactly the "figures written from memory" failure this
release keeps paying for.** Target: **≈ 7–10 s and ~1.25 GB**, both to be measured, not asserted.

## Close — 2026-08-12

**Outcome:** `TOK-02` authored — space partitioned by its **coupling to time** into three classes, each
priced from a measurement, with the largest single reclaim on the box (the 19.28 GB build cache)
deliberately placed **out of bounds**. Two axes nobody had measured were opened: the **host** cost of
Docker (`Docker.raw` allocated **50.68 GB**, tracking in-VM reclaim to 1.7 % — so this milestone's
reclaim figures are real SSD and not VM fiction) and the per-stack-dir breakdown (`data/` 1.9 GB,
`clones/` 220 MB surviving teardown). studio-desk's 1.7 GB is now **attributed to a layer**: `npm ci` is
1.04 GB — 61 % of the image — for a 63.2 MB build output, from a **single-stage** Dockerfile whose
runtime is a bare `node dist/index.js`.
**Type:** tok
**Status:** closed-fixed
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n *(never, by ruling — `D52`)* — (2) triggered-tok: **n** *(this is a
bootstrap-flavored user-directed tok; the 3-no-prog trigger did NOT fire and toks of this flavor do not
terminate the call)* — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n *(0 tiks this session)*
— (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**

**Decisions:** D59–D63

**Side-deliverables:** none (strategy work only; 0 code modified).

**Routes carried forward:** the three `TOK-02` tiks, in order — `TIK-A` studio-desk (Class B),
`TIK-B` the Class A free wins, `TIK-C` `END-M258-one-stack`. Full ordering + rationale in `TOK-02`.

**Lessons:**

- **A constraint applied early changes the design; applied late it only vetoes.** Pricing studio-desk on
  both axes *before* choosing the fix produced prune-and-copy over re-install. A space-only reading would
  have shipped a slower cold build and called it a win.
- **Ask what a reclaim costs the HOST, not the VM.** Every figure in this milestone was logical-inside-
  the-VM. They happen to be true here (1.7 % agreement) — but that had to be measured, and on a host
  whose sparse file did not TRIM, all of them would have been fiction.
- **The trap has siblings.** `D53` was `docker images` SIZE (~5× over). Its siblings are `ls -l` on
  `Docker.raw` (2.7× over) and `docker system df` not seeing host bind mounts at all (4.2 GB invisible).
  *A number's units are a claim about which question it answers.*
