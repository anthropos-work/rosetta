# M255 — the three spikes, with their raw evidence

**Host:** `billion.taildc510.ts.net` (8 vCPU / 7.3 GiB / x86_64 / Linux 6.8.0-134 / Docker 29.6.2, containerd
image store) unless stated otherwise. **State the environment with every number** — `latency-budget.md`'s rule,
and it is load-bearing here: two of the three spikes produced host-specific answers.

**Raw artefacts:** `billion:/home/devops/panorama/m255/spike-a/{spike-a.log,build-A.log,build-B.log}` ·
laptop `.agentspace/scratch/work-m255/laptop/{run.log,build.log,samples.tsv}`.

---

## Spike (a) — the L1 multi-stage experiment · **THE BARRIER DECIDER** · verdict **GO**

**Question.** Does shipping `.next/standalone` instead of the whole build tree materially cut the export leg?

**Method.** Two builds from the *same* clone (`stack-demo/next-web-app` @ `e6c818502`, unpatched, build
context only — the tree was git-clean before and after), on the rext-owned `hiring.Dockerfile`, which rext
owns outright so no demopatch and no platform-edit question arises. **A first**, so its cache seeds B's
builder stage and the delta isolates to the export leg. Both tagged to fresh names so neither reuses the
live `demo-1-hiring` image.

| | **A** — shipped single-stage | **B** — multi-stage `.next/standalone` |
|---|---|---|
| final image | **4.84 GB** | **379 MB**  (−92.2 %) |
| `pnpm install` | 29.2 s | *(cache hit from A)* |
| `pnpm turbo build` | 45.5 s | 43.7 s |
| **export step** | **146.8 s**  (layers 108.9 + unpack 37.8) | **2.9 s**  (layers 1.9 + unpack 0.8) |
| total wall | 225 s | **49 s** |

**Export leg: 146.8 s → 2.9 s. A −143.9 s (−98 %) cut on ONE of the two Next.js images.** The annotated
baseline pays 141.9 s (next-web) + 136.7 s (hiring) = 278.6 s of export across the pair, so L1's headline
estimate of ~200–250 s is if anything *conservative*.

**Two seams made it work, both rext-owned, zero platform edits:**

1. **`ENV NEXT_PRIVATE_STANDALONE=1`.** Next 16's frozen `defaultConfig` reads
   `output: !!process.env.NEXT_PRIVATE_STANDALONE ? 'standalone' : undefined`, and `apps/hiring/next.config.mjs`
   sets no `output` (verified — it passes through `withMDX` → `withNextIntl` → `withSentryConfig`, none of
   which touch `output`). **D-M255-2 resolves to the private-API path; the `next.config.mjs` demopatch
   fallback is not needed.**
2. **`turbo … --env-mode=loose`** — and this one was *not* in the design note. **Turbo 2's default env mode is
   `strict`**, which filters out any variable not listed in `turbo.json`'s `globalEnv`. `NEXT_PRIVATE_STANDALONE`
   is not in that list, so without `--env-mode=loose` the variable never reaches the `next build` child and
   standalone silently no-ops — a **green build that produces the old 4.84 GB image**. The prototype guards
   against exactly that with a `RUN test -f .../standalone/apps/hiring/server.js` step that fails the build
   loudly; keep it in the M257 lever.

**Functionally viable, not merely small.** The 379 MB image boots (`▲ Next.js 16.2.7 · ✓ Ready in 0ms`) and,
given the Clerk publishable key its compose block already supplies, answers `307 →
https://billion.taildc510.ts.net:15400/sign-in?redirect_url=…` — the correct Clerkenstein-wired middleware
redirect. A bare `docker run` with no env returns 500 (`Missing publishableKey`), which is the spike's own
omission, not the shape's: `gen_injected_override.py` already sets `CLERK_PUBLISHABLE_KEY` on the
`hiring-app` service.

**Carry into M257.** The runtime stage must receive the runtime env that `next start`-from-source got for
free from the on-disk `apps/hiring/.env.local`; standalone does not carry it. Either `COPY` the overlay into
the runtime stage or rely on the compose env block — verify which, don't assume.

---

## Spike (d) — plateau or I/O ceiling?  ·  verdict **NEITHER — it is a single-stream serial leg**

**Question.** The annotated run peaked at **load1 4.90 of 8 cores** and concluded *"the box was never
CPU-saturated"*. If BuildKit is instead **I/O-throttled**, L2's parallel win is far smaller than the naive
~200 s and M257's gate must be re-cut.

**Why the existing evidence could not answer it.** Linux `load1` counts tasks in **uninterruptible sleep**,
not just runnable ones — so a run blocked on disk inflates load1 *without using CPU*. A load of 4.90/8 is
therefore consistent with both "half the cores are idle" and "everything is stuck on the disk". The
annotated sampler recorded **no disk metric at all**, so the question was unanswerable from it. `buildbench`'s
sampler now records disk **`%util`** and write throughput from `/proc/diskstats` alongside load1.

**MEASURED, first full cycle with the new sampler** (billion, 2026-07-27, 65 samples @10 s, cold-images):

| metric | value |
|---|---|
| peak load1 | **3.75** of 8 cores · avg 1.99 |
| **peak disk `%util`** | **63.4 %** · avg **21.0 %** |
| **samples at ≥ 90 % util** | **0** |
| peak disk write | 199.3 MB/s |
| peak memory / swap | 5,887 MB / **4,338 MB** |

**Verdict: neither resource is saturated.** Not CPU (3.75 of 8), and — refuting the I/O-ceiling hypothesis —
not the disk either (63 % peak, never above 90 %, at a 10 s average). The export leg is **serial and
single-stream**: `unpigz` at ~40 % of one core, ~34 MB/s sustained on a device that demonstrably reaches
199 MB/s. Nothing is *full*; one thing is *sequential*.

*(Caveat, stated because it bounds the claim: a 10 s average cannot see a sub-second saturation spike. What
it rules out is a **sustained** ceiling, which is the hypothesis that mattered.)*

**Three consequences for M257, in order of importance:**

1. **The naive "L2 buys ~200 s" is wrong, and L1 is why.** L2's value was overlapping two ~140 s export legs.
   Spike (a) removes them: after L1 the hiring image costs ~49 s of which ~44 s is `next build`. **Overlapping
   the pair then buys at most ~45 s, not ~200 s.** Sequence L2 *after* L1 and re-cut the gate accordingly.
2. **The two COMPILE legs cannot overlap on this host anyway.** The headroom assert derives
   `max_parallel_ui_lanes = 1` on `billion` (0.8 × 7,500 MiB budget − 1,500 MiB idle) ÷ 3,900 MiB measured
   per-lane peak. Two lanes swap — and this run already peaked at **4.3 GiB of swap** with one.
3. **But the EXPORT legs could.** Export is low-memory (a decompressor plus writes) and demonstrably leaves
   both CPU and disk headroom. If L1 does not land, "overlap the exports, not the compiles" is the shape of
   the win — and it is a smaller, safer change than parallelising whole builds.

## Spike (e) — host-vs-peer topology for M258 · verdict **achievable, behind a one-line bind change**

> **Prior art, credited:** the *symptom* was already documented — `tailscale-serve.md` § "NOT from the VM itself" (M219) carries the error string, three-port measured evidence, and a retraction of the old *"(or the VM itself)"* claim. This spike is the **increment**: the mechanism is a **bind collision**, not an unavoidable property of the loopback path — which is what turns a documented limitation into a one-line fix. The correction is folded back into that section.

**Question.** A `--public-host` demo reportedly *cannot be browsed from its own host* — and `--public-host`
is default-on for the demo path (D-DESIGN-3). So is M258's *"one cold command on billion"* achievable as
literally worded?

**The symptom reproduces exactly as documented.** From `billion`, against its own live `demo-1`:

| probe | result |
|---|---|
| `https://billion.taildc510.ts.net:13000/` (next-web, the presenter's origin) | **`SSL routines::wrong version number`**, 0.10 s |
| `https://billion.taildc510.ts.net:15050/graphql` (the router — the leg that kills every page) | **same failure**, 0.02 s |
| `http://billion.taildc510.ts.net:13000/` (plain HTTP, same name) | **307** — works |
| `http://127.0.0.1:13000/` | **307** — works |
| `ss -lntp` on :13000 | `0.0.0.0:13000` → **`docker-proxy`**, and **no `tailscaled` listener at all** |

**But the stated mechanism is not quite the actionable one.** The corpus says the connection *"hits the
kernel socket and bypasses `tailscale serve`"*. The sharper cause, proven by controlled A/B below, is that
**a wildcard `0.0.0.0:P` bind stops `tailscaled` from creating its `100.x:P` listener in the first place** —
you cannot bind a specific address on a port a wildcard bind already holds. Peers still reach `tailscale
serve` (their traffic arrives through the WireGuard tunnel and never touches the kernel's port table); the
node's own traffic never enters the tunnel, so it lands on docker-proxy's plain-HTTP socket.

**The controlled A/B.** Two trivial HTTP servers on `billion`, identical except for their bind address, each
fronted by `tailscale serve --https=<port>`:

| backend bind | `tailscaled` listener created? | node-local `https://<magicdns>:<port>/` |
|---|---|---|
| `127.0.0.1:14998` (loopback only) | **yes** — `100.110.136.3:14998` | **HTTP 200, `ssl_verify=0` (trusted), 40 ms** |
| `0.0.0.0:14997` (wildcard) | **no** (IPv4) | **fails** — `wrong version number`, `ssl_verify=1` |

**So the host CAN browse its own trusted `tailscale serve` HTTPS origin — as long as nothing else holds the
wildcard bind on that port.**

**The seam is one line.** `up-injected.sh:146`:

```sh
if [ -n "${STACK_PUBLIC_HOST:-}" ]; then BIND_HOST="0.0.0.0"; else BIND_HOST=""; fi
```

Publishing on `127.0.0.1` instead of `0.0.0.0` when a public host is in play would let `tailscaled` take the
tailnet address, and the node would reach the *same* origins the presenter does.

**Recommendation for M258's gate text.** *"One cold command on billion"* is **achievable as literally
worded**, and does **not** need to fall back to `--no-public-host` (which would prove the composition in a
mode the presenter never uses). It is contingent on the `BIND_HOST` change, which is rext-only and
zero-platform-edit. Two riders M258 must carry:

- **It is also an exposure REDUCTION**, and a large one. [`safety.md`](../../../../corpus/ops/safety.md) §3
  Part 3 currently discloses that **every demo container is published on `0.0.0.0` — all interfaces — on
  every `demo-up`, flag or no flag**. Loopback-plus-`tailscale serve` removes the non-tailnet LAN surface
  entirely. Two problems, one line.
- **It makes `tailscale serve` load-bearing for all access.** Today a serve failure degrades to
  "reachable over plain HTTP"; after the change it becomes "not reachable at all from off-box". The existing
  cert/serve fallback warnings (and M255's new `certwarn.log`) are the mitigation, and M258 should assert on
  them rather than discover it live.
- **IPv6 asymmetry, noted in passing:** `tailscaled` bound `[fd7a:…]:14997` even in the wildcard case,
  because the IPv4 wildcard does not shadow the v6 address. A v6-capable client may therefore already reach
  serve on a port whose v4 side is shadowed — an inconsistency worth not relying on either way.

*(All probe ports were torn down after the experiment; `tailscale serve status` was left as the bring-up
configured it.)*
