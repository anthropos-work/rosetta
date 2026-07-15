# M218 "seat change" — Retro

## Summary

**The gate is met: p95 click→ACCESS 2413 ms / 1767 ms vs < 5000 ms**, over 5 genuine cold reset-to-seed
cycles, both vantages, 50/50 ACCESS. From a baseline of **39.45 s / 38.30 s** — **~16×** on the honest cold
number. **Zero platform-repo edits.**

The user reported *"1 or 2 minutes to access the platform."* The corpus said it was *"~2–5 s, which we can't
shorten."* **Both were wrong, and the corpus was wrong in the more expensive way** — because a claim that
something is unfixable is a claim that stops anyone from looking. That claim had stood, unmeasured, for
**four releases**.

## Incidents this cycle

| # | What | Class |
|---|---|---|
| **F-9** | **`/demo-down --purge` had NEVER purged on any Linux host.** Postgres's cluster dir is UID-1001/0700, so `rm -rf` died `Permission denied`; `set -euo pipefail` then aborted `cmd_down` **silently** — DB never wiped, images never removed, registry slot leaked, bare `rc=1`. `billion`'s postgres still carried the `PG_VERSION` `initdb` wrote **3 days earlier**. **The exit gate was not merely ungraded — it was ungradeable.** | P1 · stale state |
| **F-6** | A **9-hour-old** `autoverify.json` graded a **Clerkenstein-DEWIRED** stack green — the browser's clerk-js was talking to the **REAL Clerk app**. The cache-validator read image ENV; the pk is *inlined into the bundle*. | P1 · stale verdict |
| **F-10** | A **torn-down** stack (zero containers) still served `{"green":true}` to every grader. | P2 · stale verdict |
| **D15** | Clerkenstein scored **100% / 100% / 0 divergences** while its fake BAPI **returned the wrong human for every hero**. `GET /v1/users/{id}` had **no gene in any of the 5 DNAs**; the three genes that *did* name identity all asserted **the stub itself**. **The goldens ratified the defect.** | P1 · hollow measurement |
| **F-11** | The **ORG-level twin** of the same stub, found while fixing the first. | P2 · same class |
| **MF-1** (close) | The **F-9 fix reproduced F-9's collateral leak**: `die` on a failed purge still skipped `docker rmi` / `reg_del` / `ureg_release` — **the exact three items F-9's own test file enumerates as its damage.** | P1 · same class |
| **MF-2** (close) | A **failed frontend build graded GREEN** — compose started a stale image; nothing proved the running image was this run's build. **M218 made that verdict a gate input**, so a pre-existing false-green became **load-bearing**. | P1 · same class |
| **MF-3** (close) | A **SKIPPED demo-patch graded GREEN, on the gate path.** A stack running **without the headline SSR fix** printed *"✓ demo-patches: none refused"* and was **measurable**. | P1 · same class |
| **D18** (close) | The hardening pass **misfiled an M218 regression as "pre-existing"** by using *M218's own iter-05 commit* as the "baseline". | P1 · same class |

**Nine incidents. One class.** A status artifact that outlives the thing it describes, and is then read as
evidence — plus its corollary, *absence read as success*. It is now a named hazard in
[`corpus/ops/verification.md`](../../../../../corpus/ops/verification.md) with two invariants: **a verdict must
not outlive its subject**, and **absence must be the safe state** (a grader with no verdict **refuses to
measure**).

## What went well

- **Harness before fix.** iter-01 was a **4-leg experiment with zero code written**, and it paid for itself
  immediately: it proved the milestone's *own planned one-line fix was a no-op* (`NEXT_PUBLIC_*` is
  build-inlined; SSR never reads `process.env`). We would have shipped a fix that changed nothing and
  "confirmed" it against a warm stack.
- **Arithmetic as a diagnostic.** iter-01's static prediction (37.5 s) and iter-02's measured SSR body
  (37.533 s) agreed **to within 33 ms**. Later, *a blackhole and a refusal turned out to be six seconds apart
  in signature* (`3 × 10.5 s + 6 s` vs `3 × 33 ms + 6 s`) — **the magnitude named the bug class before a line
  of code was read.**
- **Refusing the convenient number.** iter-04 had a passing gate, a 3.4× margin, and every incentive to
  declare victory. iter-05 spent its entire budget proving that gate **wasn't real** — and found F-9.
- **D16: shipping an honest 97.2% over a hollow 100%.** The alternative — omit the field, keep the green —
  is *precisely* the mechanism by which the headline bug survived four releases.

## What didn't

- **We named the hazard and then kept walking into it.** D17 was written *during* this milestone. The close
  then found **four more instances** — one of them **inside the hardening pass convened to name it**, and one
  **in my own regression test**, which false-passed because its function slice swallowed the definition it was
  asserting on. **Naming a hazard does not inoculate you against it. Only a probe does.**
- **The headline fix had NO TEST until the final harden pass.** The ~37.5 s → 2.4 s collapse rested on a
  **two-part chain across two subsystems** (a demo-patch reads `WUNDERGRAPH_SSR_ENDPOINT`; `gen_injected_override.py`
  supplies it) — **neither half works alone, and neither had a fence.** Per-iter regression tests, each scoped
  to one iter's diff, *structurally cannot* hold such a chain together.
- **Docs corrections were half-landed.** M218 corrected the "CI gates alignment" claim — and replaced it with
  a **different false claim** ("rext has no `.github/workflows` at all"; **it does, and it's git-tracked** — it
  simply never runs, because Actions only reads the *repo root*). It corrected the coverage-check claim — and
  the correction **over-claimed the new check** (the gate passes `--if-declared`, so 4 of 5 surfaces are
  *warned about*, not *enforced*). **D16 made the score honest in the code and left 7 docs saying 100%.**
- **A `Delivers →` item silently didn't ship** (the clerk-js caching/timeout contract), and a brand-new corpus
  doc shipped **undiscoverable** — `latency-budget.md` had exactly one inbound link, mid-body, and appeared in
  no index.

## Carried forward

| Item | Destination |
|---|---|
| **F-11** — the BAPI fabricates the org eid (ships as a **deliberately RED gene** until fixed) | **M219** |
| `expressrun` is **UNMEASURABLE** — rc=2, **no score**, and nothing treats that as failure | **M219** |
| The demopatch freshness gate **skips** without the clone (*absence read as success*) | **M219** |
| Clerk telemetry off · **F-5** ad-tech egress · **C-5** vendor clerk-js + bound the **unbounded `Timeout: 0`** · ant-academy's **real** Clerk secret | **M220** |
| **F-7** — `NEXT_PUBLIC_BACKEND_API_URL`, a **measured 10.5 s blackhole** from inside the container. **A loaded gun**: dormant *only* because every reader is client-side. | **M221** |
| **C-3** — cms/Directus **403s** on the CONTENT path (exercisable for the first time now that logins complete) | **M221** |

**Every receiving `overview.md` was EDITED** — Fate 3 means the sibling's plan changes, not that a string
changes in ours. **Zero escape-hatch deferrals; nothing leaves v2.3.**

## Metrics delta

| | M217 | **M218** |
|---|---|---|
| **p95 click→ACCESS** | *(unmeasured — no instrument existed)* | **2413 / 1767 ms** vs < 5000 ms ✅ |
| Python tests | 867 (0 fail) | **887** (0 fail) |
| Go test funcs | 1750 | **1784** (+34, same method) |
| Alignment (Go surface) | 100% / 100% *(hollow)* | **97.2% / 100% critical** *(honest)* |
| Flake count | 0 | **0** (5/5 sequential) |
| Platform-repo edits | 0 | **0** |
| Bugs fixed | 32 | **21** (build 6 · harden 4 · **close 11**) |

## The one line worth keeping

**Do not write "we can't fix this" about something you have never measured.** It cost four releases and a
16× regression that nobody was allowed to see — because the documentation said there was nothing to look at.
