# M255 — decisions

Release-level binding decisions **D-v28-1 … D-v28-11** live in
[`../../../roadmap.md`](../../../roadmap.md) § Active — v2.8.

---

## D-M255-1 — the headroom assert fails LOUD in buildbench and stays ADVISORY in the bring-up

**Required by item 3.** M255 asserts headroom and also retracts `preflight_vm_ram()` as *"cosmetic"*. Those two
statements have to be reconciled, or the release reads as demanding the very thing it is retracting.

**Resolution — they are different contracts, for different consumers, and both are correct:**

| consumer | what it is | on a headroom failure |
|---|---|---|
| `up-injected.sh` pre-flight | an **operator's** bring-up | **warns, continues** (`:302`, `:341`) |
| `buildbench` pre-rep + post-rep | a **measuring instrument feeding a release gate** | **aborts the rep, exits 1** |

*"Never block a genuinely good bring-up on a soft heuristic"* is right for a human at a terminal. It is wrong
for a gate: **a number measured on a host without headroom is not a number.** So the same three clauses are
evaluated in both places and only one of them is fatal.

**And the retraction is narrower than "pre-flights are useless".** `preflight_vm_ram()` declares its `bytes`
and `gib` **`local`**, assigns no global and returns no verdict — *nothing branches on it*. It is cosmetic in
the sense of **inert**, which is not the same as *tolerant*. M255 makes the disk pre-flight **correctly sized**
(20 → 25 GiB, derived from measurement) while leaving it non-fatal, and puts the fatal twin where a fatal
twin belongs. `demo-stack/tests/test_tooling.py:437` independently forbids making the operator-facing one
fatal; that test is not in tension with this decision, it *is* this decision.

---

## D-M255-2 — L1 lands via `ENV NEXT_PRIVATE_STANDALONE=1`; the demopatch fallback is NOT needed

**Resolved by spike (a), with a correction the design note did not anticipate.**

`NEXT_PRIVATE_STANDALONE=1` works: Next 16's frozen `defaultConfig` reads
`output: !!process.env.NEXT_PRIVATE_STANDALONE ? 'standalone' : undefined`, and no app `next.config` sets
`output`. Measured on `hiring.Dockerfile`: **4.84 GB → 379 MB**, export **146.8 s → 2.9 s**, and the image
boots and Clerkenstein-redirects correctly. **The `next.config.mjs` demopatch fallback is not required.**

**But a second seam is, and it was not in the plan.** **Turbo 2's default env mode is `strict`**, which filters
any variable not listed in `turbo.json`'s `globalEnv`. `NEXT_PRIVATE_STANDALONE` is not in that list, so
without **`--env-mode=loose`** on the turbo invocation the variable never reaches the `next build` child and
standalone **silently no-ops** — producing a green build and the old 4.84 GB image. The prototype guards it
with a `RUN test -f .../standalone/apps/hiring/server.js` step that fails the build loudly. **M257 must keep
both the flag and the guard.**

---

## D-M255-3 — D-v28-7's "inert" waiver is REFUTED; union-apply is a real (beneficial) change for hiring

**Surfaced by the Phase-0b KB-fidelity audit, confirmed in code.**

D-v28-7 waived `next-web-ssr-graphql-origin` as *"inert for the hiring image **by its own manifest header**"*.
The manifest header contains **zero** occurrences of "hiring"; it says the patch is behaviour-identical **when
`WUNDERGRAPH_SSR_ENDPOINT` is unset**. The plan converted a statement about a *variable* into a statement
about an *image*. The variable **is set on the hiring container**
(`stack-injection/gen_injected_override.py:367`), and `apps/hiring` really imports the patched module
(`apps/hiring/src/app/api/bunny/recording/[sessionId]/route.ts` → `createServerGraphQLClient`).

**Decision: keep union-apply, replace the waiver's reasoning.** The change is a strict **improvement** — that
route currently resolves its SSR origin from the build-inlined *public* endpoint and therefore carries the
same blackhole M218 measured at 37.5 s — and it is a pure prepend to an existing `||` chain, so it cannot
regress a path where the variable is unset. But *beneficial* is not *inert*: **M257 must re-verify the hiring
recruiter Playthrough after flipping union-apply on.** Recorded in `union_apply_guard.py`'s waiver table
(where a stale or reason-less waiver now fails the guard), in `roadmap.md`'s D-v28-7, and in `build-budget.md`.

---

## D-M255-4 — "revert once LIFO" had no referent; the real invariant is the `urls.ts` chain

`demopatch-spec.md` §4 claimed both frontend builds revert **LIFO** and in **identical** order. Derived from
the source: **neither is strict LIFO, and the two orders differ.**

| build | strict LIFO? |
|---|---|
| `next-web` (9 manifests) | no — reverts mostly in APPLY order, with the chained pair and the interview pair swapped |
| `hiring` (7 manifests) | no — reverts mostly in reverse, with `back-to-cockpit` reverted **last** |

**And it does not matter.** Every manifest except the `urls.ts` pair targets a file no sibling in its set
touches, so its position is free. The **only** order-sensitive relationship is
`next-web-public-website-url.pre_sha256 == next-web-studio-url.post_sha256` → *studio applies first, pubweb
reverts first* — and **both builds already get that right**.

**Decision: state the narrow invariant and fence it, rather than inherit a wrong general one.** The union-apply
rule is now *"apply the union once, build both in parallel, revert once — with the `urls.ts` chain reverted
pubweb-before-studio"*, and `union_apply_guard.py` asserts the chain order in **both** builds and in **both**
phases (RED-proven against a swapped apply and a swapped revert). `demopatch-spec.md` §4 is corrected,
including its **"4-manifest union"** → **7** (a C1 mirrored count that drifted at M232 and again at M249).

---

## D-M255-5 — the corpus called the barrier's own shape "forbidden"; a third build shape is now documented

`frontend-tier.md`'s hard-line box stated the model as exhaustive (*"their Dockerfiles consumed UNMODIFIED"*)
and its closing section routed `output:'standalone'` to a **"forbidden"** upstream PR. Spike (a) — the
barrier's decider — uses precisely the shape those two passages exclude, and it is **not new**: rext has owned
`demo-stack/frontend/hiring.Dockerfile` since M224. `grep -ci hiring frontend-tier.md` returned **0**.

**Decision: document the third shape now, in this milestone.** A GO verdict that contradicts the corpus is not
a verdict. `frontend-tier.md` now enumerates three shapes (platform Dockerfile as-is · platform source
demo-patched · **rext-owned Dockerfile, platform repo as build context only**) and records that shape 3 is
*stronger* on the zero-platform-edit line, not weaker. The closing section's stale claims are annotated with
what has actually since shipped (the SSR origin landed as a demo-patch at M218; standalone needs no PR at
all). **Not** the full §8.5 numeric rewrite — that stays with M257 (D-v28-10), so `frontend-tier.md` moves
once with achieved numbers.

---

## D-M255-6 — spike (d) refutes BOTH hypotheses; L2 is a ~45 s lever, not a ~200 s one

Measured with the new disk sampler: **peak load1 3.75/8, peak disk `%util` 63.4 %, zero samples ≥ 90 %.**
Neither CPU-saturated nor I/O-ceilinged. The export leg is **serial and single-stream** — `unpigz` at ~40 % of
one core moving ~34 MB/s on a device that reaches 199 MB/s.

**Consequence, which M257 must absorb before it commits a gate:** L2's value was overlapping two ~140 s export
legs, and **L1 deletes them**. After L1 the hiring image costs ~49 s, of which ~44 s is `next build`. So
**L2 buys ≲ 45 s, not ~200 s** — and the two *compile* legs cannot overlap on `billion` regardless, because
the headroom assert derives `max_parallel_ui_lanes = 1` (0.8 × 7,500 MiB − 1,500 MiB idle) ÷ 3,900 MiB
measured per-lane peak, and this run already peaked at **4.3 GiB of swap** with a single lane. The *export*
legs remain overlappable (low-memory, headroom on both axes) if L1 slips.

---

## D-M255-7 — spike (e): M258's "one cold command on billion" is achievable as worded

The symptom was already documented (`tailscale-serve.md` § "NOT from the VM itself", M219). The **mechanism**
was not: it is a **bind collision**, not an unavoidable property of the loopback path. A wildcard `0.0.0.0:P`
bind stops `tailscaled` from creating its own `100.x:P` listener; peers are unaffected because their traffic
never consults the host's port table. Controlled A/B on `billion`: a loopback-bound backend on the same port
is reachable from the node at its own trusted MagicDNS HTTPS origin in **40 ms, `ssl_verify=0`**; a
wildcard-bound one is not.

**Decision: M258 does NOT need to fall back to `--no-public-host`** (which would prove the composition in a
mode the presenter never uses). The seam is one line — `up-injected.sh:146`'s `BIND_HOST` — it is rext-only
and zero-platform-edit, and it **also removes the non-tailnet LAN exposure** `safety.md` §3.1 discloses.
**Not changed in M255** (out of scope: item list has no exposure-model change, and it needs its own
verification pass). Two riders for M258: it makes `tailscale serve` load-bearing for all off-box access, and
IPv4/IPv6 behave differently under the collision.

---

## Items surfaced during the build and their fates (three-fate rule)

| item | fate | where |
|---|---|---|
| D-v28-7's "inert" premise refuted | **Fate 1 — landed** | waiver rewritten + guard clause + `roadmap.md`/`overview.md` corrected |
| `demopatch-spec.md` §4's "4-manifest union" (→7) and "LIFO" claims | **Fate 1 — landed** | §4 corrected with a derived table |
| `frontend-tier.md` excludes the rext-owned Dockerfile shape | **Fate 1 — landed** | three-shape table added; the stale "forbidden" list annotated |
| `frontend-tier.md` §8.5 full numeric rewrite (9 structural sites, incl. the tier definition and the port table) | **Fate 2 — already owned by M257** (D-v28-10) | annotated in place so M257 rewrites the *model*, not just the numbers |
| `DEMO_DISK_MIN_GIB` is an unfenced 8-site mirrored change | **Fate 1 — landed** | all live sites updated; the pinning test now DERIVES the floor and cross-checks the host profile |
| the defaults table's `Read at` column: **28 of 29 citations drifted**, unfenced | **Fate 1 — landed** | guard clause (5) + a `--fix` regenerator + 4 tests |
| the mirrored knob COUNT ("all 27" vs 29 parsed), unfenced | **Fate 1 — landed** | guard clause (4) across 3 prose mirrors + 7 tests |
| `buildbench`'s sampler shadowed `Thread._stop` | **Fate 1 — landed** | renamed + 2 regression tests (it cost one 11-minute cycle) |
| the `BIND_HOST` change that would let a host browse its own demo | **Fate 3 — attached to M258** | rationale + riders in D-M255-7 and `tailscale-serve.md` |
| G5 refuses revert whenever the self-healing freshness gate fired (traps swallow it) | **Fate 1 — documented** | `demopatch-spec.md` §4; not currently harmful (the clone is force-checked-out next bring-up) but "the trap left it clean" is an assumption, not a guarantee |

**Zero cross-release deferrals.** No escape-hatch entry was needed.
