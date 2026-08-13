**Type:** tik · **Active strategy:** `TOK-02` step 4 — *then levers, largest-measured-second first, one per
iter, each landing together with the falsifiable assert that can trip it* · **Declared multi-step shape**
(`overview.md` Phase plan): land L1 → land the ISOLATION assert → measure → close.

## Line 1 — L1 lands, and the mechanism was verified before it was trusted

Three premises had to hold on **this** host. Each was re-measured at open rather than quoted:

| premise | how it was checked | verdict |
|---|---|---|
| no app `next.config` sets `output` | `grep` over both configs + repo-wide for `output:` in any `next.config.*` | HOLDS (×5) |
| `NEXT_PRIVATE_STANDALONE` flips Next 16's frozen `defaultConfig` | read out of the **real installed** next@16.2.12 inside `demo-1-next-web` — `dist/server/config-shared.js:129`: `output: !!process.env.NEXT_PRIVATE_STANDALONE ? 'standalone' : undefined` | HOLDS |
| M255: turbo's `strict` env mode filters the var out before `next build` sees it | read this repo's own `turbo.json` `globalEnv` — **44 names, and `NEXT_PRIVATE_STANDALONE` is not one of them** | **LIVE HERE** |

The third is the one that matters most, because its failure is silent: without `--env-mode=loose` the turbo
task never receives the variable, `next build` exits **0**, and the image is the old ~4 GB one. So the
builder stage does not trust it — it asserts `test -d apps/<app>/.next/standalone` and fails the build with
a message naming the cause. **A lever whose failure mode is "green, and 68 s slower" needs a louder fence
than a code review.**

Two rext-owned Dockerfiles, no platform-repo edit:

- **`demo-stack/frontend/hiring.Dockerfile`** — already rext-owned since M224; made multi-stage in place.
- **`demo-stack/frontend/next-web.Dockerfile`** — **net-new**. next-web had been building from the platform
  clone's own single-stage `Dockerfile.dev`; a multi-stage form cannot land there without editing the
  platform repo, so it lands as the alternative `overview.md` § Hard constraints already names — *an
  rext-owned Dockerfile in the shape `hiring.Dockerfile` already sanctions*. The clone is now not even read
  for its build recipe; it is a build **context** and nothing else.

Three contracts the runner stage had to preserve, each of which would have failed silently:

1. **The baked `NEXT_PUBLIC_*` must be re-declared in the final stage.** `docker image inspect` reads
   `.Config.Env` of the FINAL stage only, and a multi-stage build starts that stage from a fresh base. Miss
   this and up-injected.sh's offset validator reads an **empty** endpoint, takes its documented
   *"unreadable ⇒ no regression, reuse"* branch, and the stale-offset check is disarmed forever.
2. **`/app/apps/<app>/.next/static` is a cross-file contract**, not an implementation detail — the
   minted-pk probe greps that exact path to prove the `.env.local` overlay was applied. The `COPY`
   destinations are byte-identical to it.
3. **`ENV HOSTNAME=0.0.0.0`.** standalone's `server.js` binds `process.env.HOSTNAME`. Measured: the L1
   container logs `Network: http://0.0.0.0:3000`, the old one `http://172.17.0.3:3000`.

## Line 2 — what L1 is worth on this host, measured

Baseline column is **iter-08 rep-03**, the headroom-clean rep of the landed `gated_baseline` (macmini, Apple
M4 Pro, arm64, Docker Desktop VM 8 vCPU / 11,948 MiB, containerd image store). L1 column is this iter, same
host, same clone (`next-web-app` @ `19423a1fb`, v2.137.3), `load1` 2.0–4.3 throughout.

| | next-web | hiring | **both** |
|---|---|---|---|
| image **before** | 4.04 GB | 3.94 GB | 7.98 GB |
| image **after** | **417 MB** | **380 MB** | **797 MB** |
| | *9.7×* | *10.4×* | *10.0×* |
| `exporting layers` before → after | 53.0 s → **1.4 s** | 50.1 s → **1.3 s** | 103.1 → 2.7 |
| `unpacking to …` before → after | 17.1 s → **0.6 s** | 16.0 s → **0.5 s** | 33.1 → 1.1 |
| **`exporting to image` before → after** | **70.2 s → 2.0 s** | **66.2 s → 1.8 s** | **136.4 s → 3.8 s** |

**≈ 132.6 s removed from a leg the gate needs 89.51 s out of.** The unpack leg is the half the corpus said
this host did not pay at all — the claim iter-05 retracted — and it is 33.1 s of the 136.4 s.

**What did NOT change, and must not be credited to L1:** the two compile legs (28.5 + 25.0 = 53.5 s) and the
two installs (14.4 + 14.9 = 29.3 s). Those are L4's and L3's. Attribution is the deliverable (`TOK-01`), so
the lever is priced on the leg it actually moves.

### The A/B's own defect, recorded because it changes how the number should be read

The intended design was warm-up → L1 → baseline, all three from one context so the arms shared a cache
state. The **baseline arm came back in 1.12 s**: identical to the warm-up in every layer, so BuildKit served
the entire build — including the export — from cache. **That arm measured nothing**, and its 1.12 s is not a
number about the old Dockerfile.

It does not damage the finding, because the two figures that carry it — **image size** and **the export /
unpack legs** — are properties of the artefact, not of cache state, and the "before" side is taken from
iter-08's real bring-up rather than from this arm. But *"the second arm was a cache hit"* is exactly the
kind of thing that silently becomes a quoted number later, so it is written down as a null result.

## Line 3 — the two apps serve identically, which is the claim that actually needed proving

A 10× smaller image is worthless if it does not run. Both images, side by side against the images they
replace, same host, same probe:

| | `/` | `/login` |
|---|---|---|
| next-web **L1** (417 MB) | 307, 78 B, 0.047 s | **200, 339,595 B**, 0.242 s |
| next-web **old** (4.04 GB) | 307, 78 B, 0.040 s | **200, 339,594 B**, 0.240 s |
| hiring **L1** (380 MB) | 307, 78 B, 0.041 s | **200, 426,914 B**, 0.250 s |
| hiring **old** (3.94 GB) | 307, 78 B, 0.043 s | **200, 426,914 B**, 0.281 s |

**hiring's authenticated login page is byte-identical across a 10× image reduction.** next-web differs by
1 byte, which is the build-date stamp `next.config.mjs` bakes (`APP_BUILD_DATE`).

**Two false alarms on the way, both mine, both instrument errors** — and this milestone's standing lesson is
*read the instrument before theorising about the subject*:

- *"the minted pk is missing from the L1 bundle"* — the probe was `grep -rql … | head`, and `-q` **silences
  `-l`**, so it printed nothing whatever it found. The validator's **exact** probe returns FOUND, 2 files,
  the same shape as the old image.
- *"the L1 image 500s"* — it did, on `Publishable key not valid`, because I had fabricated a pk. Clerk
  parses the key (`pk_test_<base64(host$)>`). Minting one with the tooling's **own** `mint_pk` produced the
  table above. **A fixture can be the defect**, and it was.

## Line 4 — `ASSERT-M257-isolation-with-L1`: the gate's second falsifiable clause, implemented

The gate has named two asserts that FAIL it when tripped since the milestone was written. HEADROOM has been
implemented and fenced since M255; **ISOLATION had no implementation at all**. `TOK-01` deferred it
deliberately (*land each falsifiable assert together with the lever that can trip it, never after*) and the
deferral went unrecorded until iter-07. L1 rewrites exactly the layers that carry the two values, so it
lands here.

`isolation_assert()` in `stack-core/buildbench.py`, wired into the rep entry beside `headroom` and read by
`rep_is_ok`, so a leaking rep **fails the campaign** rather than being noted. Two clauses:

- **`foreign_pk`** — every publishable key baked into a `demo-N` image's build output must be this stack's
  minted key. Reported values-blind (12-char prefix + sha256 tag), because a failure message should be safe
  to paste.
- **`foreign_origin`** — every origin in the image's baked `NEXT_PUBLIC_*`/`VITE_*` env must resolve to
  stack `n` under `base + N*10000`, and the message **names the stack it leaked from**.

Plus the arms that stop it passing by measuring nothing: **zero images**, **unknown own key**, and
**unreadable bundle** are each a FAILURE.

**Proven able to fail — 17 unit controls plus three against real images.** The live ones matter most:

| live control | result |
|---|---|
| the real L1 images graded as **stack 1** (their true owner) | **green**, having actually read the minted key |
| the same images graded as **stack 2** | **RED**, 6 × `foreign_origin`, each naming *"belongs to STACK 1"* |
| `alpine:3` — an image with no build output | **RED**, `unreadable_bundle` — not "clean" |
| **the exact live call path** — `_stack_minted_pk` + `_image_sizes(1)` + both probes, over all **8** real `demo-1-*` images | **green**, `own_pk_fingerprint pk_test_MTI3…61fbfaf4`, 0 foreign |

### The bug that would have failed every campaign, caught before one was run

`isolation_assert` is handed **every** `demo-N-*` image, and most are not UI images. The first cut demanded a
scannable bundle from all of them. Measured on the real stack: `next-web` bakes 3 browser vars, `hiring` 3,
`studio-desk` 4 (and keeps its output in `/app/dist`, not `.next`) — while `postgresql`, `sentinel` and
`app` bake **zero** and ship no build output at all. Three `unreadable_bundle` failures per rep, forever, on
a clause that had just been added to `rep_is_ok`. Under a speed campaign that reads as *"my lever broke the
stack"* — the trap `build-budget.md` already records for a mid-campaign ENOSPC presenting as `redis exited (1)`.

The fix is a **derived** predicate — *does this image bake browser constants* — not a hardcoded list of UI
image names, which would silently skip a fourth UI image nobody remembered to add.

**And the first version of that fix was wrong in the opposite direction, caught by this file's own
regression test.** It used the predicate to decide whether to scan at all — so a runner stage that FORGOT to
re-declare `NEXT_PUBLIC_*`, which is the single most likely way to get L1 wrong, would present as "not a UI
image" and never have its bundle read. **The defect L1 is most likely to introduce would have been invisible
to the assert landed to guard it.** The bundle is now always requested; the predicate decides only whether
its *absence* is an anomaly.

### The fail-open the live control caught, about sixty seconds after it was written

The first probe scanned all of `/app` and passed `grep --exclude=.env`. Run against the real image:

```
grep: unrecognized option: exclude=.env      # busybox — node:24-alpine has no GNU grep
```

It matched nothing, returned `[]`, and **the assert pronounced the image CLEAN having read nothing.** That
is the precise failure this release exists to retract, committed inside the assert written to enforce it.

Two fixes, both structural rather than local: the probe **depends on no grep flag beyond POSIX**, and it
returns **`None` for "no measurement" distinct from `[]` for "measured, clean"** — which `isolation_assert`
books as `unreadable_bundle`. A fixture-only suite could never have caught this: a fixture's `pks_in` never
touches busybox. It is now pinned by a regression test that says so.

## Line 5 — a real pre-existing finding, kept rather than swallowed

Scanning for the pk found the **platform's own real-Clerk publishable key** inside the next-web image, at
`/app/apps/web/.env` — a **committed** repo file. Before concluding anything about L1, the same probe was
run against `demo-1-next-web`, built by the **real tooling at iter-08, before L1 existed**:

| | built output (`.next`) — what the browser executes | source `.env` — carried as a file |
|---|---|---|
| L1 image | the **minted** key only | the real-Clerk key |
| pre-L1 iter-08 image | the **minted** key only | the real-Clerk key |

**So it is pre-existing, L1 is exonerated, and the demo is NOT wired to real Clerk** — `.env.local` wins at
build time, which is why the bundle is correct. But the file ships into every demo image, and it is one
build-order accident from the M218 iter-03 incident (a demo silently phoning production auth).

The clause is scoped to the **build output**, which is the gate's own word — *"**baked** publishable key"* —
and a source `.env` is an input carried as a file, not something baked. That scoping is not what makes the
finding go away: **it is routed forward with its own handler** (`FIX-M257-committed-env-ships-real-clerk-pk`)
rather than absorbed into a clause the gate did not write about. A permanently-red clause is a clause that
gets switched off.

*(Incidental L1 benefit: the standalone image does not ship `apps/web/.env.local` at all — one fewer
credential-bearing file in the runtime image, where the old image shipped both.)*

## Line 6 — the tests that encoded the old wiring, and the one that was a real signal

The change turned **17 tests** red. 16 were fixtures and harnesses provisioning a world that no longer
matches (`apps/web/package.json` now stands where `Dockerfile.dev` used to, and `$HERE/frontend/` now holds
two Dockerfiles the harness had to copy in beside the `.dockerignore` it already copied). **The 17th was a
signal about my change and not about the tests:**

```
AssertionError: 112408 not less than 76775 : the hiring role-remap patch must be wired BEFORE the hiring docker build
```

`test_demopatch_hiring_role_remap_wiring` located the hiring build by searching the **whole file** for
`docker build -f "$dockerfile"` — unique only for as long as hiring was the only rext-Dockerfile build. L1
gave next-web the same shape, so the search silently changed subject to the *next-web* build thousands of
characters earlier, and the assert failed on a file whose ordering was perfectly correct. **A locator that
changes subject when a second call site appears is the defect**; it is now scoped to
`build_frontend_hiring`'s own body.

One test was **rewritten rather than repaired**, and deliberately.
`test_unmodified_platform_dockerfiles_are_the_build_input` asserted the literal `-f "$ctx/Dockerfile.dev"`
and nothing else — conflating the **contract** (zero platform-repo edits) with one **wiring detail** that
happened to satisfy it. Keeping it would have made a test that *requires the thing this milestone set out to
change*. It is now `test_platform_repo_is_a_build_context_only` and grades the contract in both directions:
rext-owned Dockerfiles for the two Next images, the clone's own `Dockerfile.dev` still used where it is
genuinely the input (studio-desk), and the trap-removal that keeps the clone pristine. **Verified
end-to-end**: after the build, `git status --porcelain` in the clone is empty.

Net: **+3 tests** (`282 passed` in demo-stack's two suites, from 280) and **+19** in stack-core.

**Two more locators had the same disease and were still GREEN** — `test_demopatch_studio_url_wiring` and
`test_demopatch_members_pagination_wiring` both located "the next-web build" by whole-file search for
`-f "$ctx/Dockerfile.dev"`. After L1 the only remaining match is the **studio-desk** build, further down the
file; since studio-desk builds after next-web the orderings still held, so both tests kept passing while
grading a weaker proposition than their own messages state. Both are now scoped to
`build_frontend_next_web`'s body. **A green test that changed subject is worse than a red one**, because
nothing tells you.

## Line 7 — 24 corpus citations shifted underneath, and the fence caught 5

`buildbench.py` grew ~150 lines and `up-injected.sh` ~21, which moves every line beneath them. The corpus
holds **61** `file:line` citations into those two files:

| | count |
|---|---|
| still resolved to their intended content | 37 |
| silently pointing at **different content** | **24** |
| of those, caught by the pre-commit fence | **5** |
| knob anchors in `demo-up-defaults.md` (a separate standalone guard) | 11 broken |

The fence books a citation that lands on a **non-construct** (blank line, `fi`, `}`) — so a citation that
slides from one real construct onto a *different* real construct passes. One of them had been wrong since
**before** this iter: `build-budget.md` cited *"the argparse constructed at `buildbench.py:1464`"*, and at
HEAD `:1464` was `report["reclaim_attribution"] = …`.

All 24 repaired from a positional HEAD→worktree line map rather than by guessing offsets (several intended
lines are blank or a bare `#`, and two are duplicated between the next-web and hiring functions, so
content-matching would mis-resolve). Re-verified: **51 of 51 resolve to intended content, 0 stale**, fence
`OK — this tree publishes no adjudicated claim the baseline did not already record`. The 11 knob anchors
were regenerated with the guard's own `--fix`, which then reported both directions in agreement.
Routed forward as **`FIX-M257-anchor-guard-content-drift`** — the class is mechanically detectable with the
same ~15-line map.

## Line 8 — the sweep, published, and the two regressions that were mine

Scope derived from what the tag ships (`git status --short` → two sections), the same derivation iter-08
used. **demo-stack 9 failed / 1087 passed; stack-core 55 failed / 2251 passed** (35 min).

demo-stack's nine are **bit-for-bit iter-08's set**, and that is a measurement rather than a hope: none of
the four failing test files, and none of the demopatch manifests they read, appears in this iter's
`git status`. Six are whole-file `pre_sha256` baselines (the anchors still hold — `demopatch-spec.md`'s own
rule), three need a live postgres.

**stack-core was 55 against iter-08's 51, and the +4 had to be accounted for rather than waved at.** Two
were mine, both found and **fixed rather than absorbed**:

| | | |
|---|---|---|
| `DOCSTRING_LITERAL_CEILING` | 254 → **255** | my `pk_fingerprint` docstring said *"12-char prefix"*. Re-worded to *"a short prefix"*; the census now reads **254**, exactly iter-08's figure. **The ceiling was NOT raised** — the 254 > 240 breach is pre-existing and stays routed as `RATCHET-M257-literal-ceilings-breached`. iter-08's move was *"stop feeding them"*; I fed it one and took it back. |
| citation denominators | 708 → **710**, 983 → **985** | this iter's own re-anchoring added two net-new line pins. **Numerators unchanged at 291/407** — which is precisely the shape `claim_census_guard`'s docstring documents as maintenance-not-defect, now for the fourth time. Updated, with the provenance recorded in the same paragraph that predicts it. |

**And then the decisive check: does any remaining failure NAME my files?** `test_isolation_assert_m257.py`,
`next-web.Dockerfile`, `buildbench` — grepped across the whole re-run: **zero hits**. The residue is the
clusters iter-08 already triaged and routed, dominated by the 2.2 GB of gitignored stack scratch under
`demo-stack/stacks/` that every tree-walking census scans (`FIX-M257-sweep-scratch-pollutes-census`; `git
ls-files` → 0 tracked, so no clone at any tag can carry it).

**Published, rung zero honoured:** `main` pushed, tag `fast-build-m257-iter-09` pushed and **verified on
origin** with `git ls-remote` (`8ca2b5b3`) rather than assumed — the thing M236 lost a whole iteration to.
Consumption clone re-pinned to the tag; the box SoT at `/rosetta/.agentspace/rext.tag` re-pointed (previous
value preserved at `.agentspace/scratch/work-m257/rext.tag.before-iter09`), which is the file
`up-injected.sh`'s `REPO_ROOT` actually reads and the one iter-08 lost three fast-failed reps to getting
wrong.

## Line 9 — the campaign: **the gate is MET**

`n=3` cold `--purge` + `up` cycles on the free `demo-1` slot, driven from the **pinned** consumption clone at
the tag just published, launched at `load1` 3.06.

```
rep-01 total=286.99s up_rc=0 green=True warn=0 headroom=OK isolation=OK phases=complete  peak_load1 5.47
rep-02 total=303.44s up_rc=0 green=True warn=0 headroom=OK isolation=OK phases=complete  peak_load1 6.16
rep-03 total=280.99s up_rc=0 green=True warn=0 headroom=OK isolation=OK phases=complete  peak_load1 8.72

n=3   p50 286.99   min 280.99   max 303.44   host identity MATCH x3
```

**Every clause of the exit gate, checked one at a time:**

| gate clause | result |
|---|---|
| p50 ≤ **360 s** over 3 consecutive cold cycles | **286.99 s** ✅ |
| …on **`macmini`** | `host_identity: match` ×3 ✅ |
| `autoverify green:true / **0 warnings**` | `green=True warnings=0` ×3 ✅ |
| **0 platform-repo edits** | `git status --porcelain` empty in `next-web-app`, `studio-desk`, `platform` ✅ |
| all **7 demopatch guards** passing | `demopatch.log`: **0** REFUSED, **0** SKIPPED; `buildfail.log` 0 bytes ✅ |
| **HEADROOM** (falsifiable) | `ok=True` ×3 — peak load1 5.47 / 6.16 / 8.72 against a limit of 10 ✅ |
| **ISOLATION** (falsifiable) | `ok=True` ×3, 8 images inspected, 0 failures — *and it exists as of this iter* ✅ |
| **stretch ≤ 300 s** | **286.99 s** p50, and 2 of 3 individual reps ✅ |

**This one is not contended-and-labelled.** iter-08's baseline exited RED by contract because 2 of 3 reps
breached HEADROOM at peak load1 19.48 / 14.52. All three reps here passed every clause, so the campaign
stands on its own terms rather than on a disclosure.

### Where the 162.52 s went

| sub-phase | iter-08 p50 | iter-09 p50 | Δ |
|---|---|---|---|
| `ui_hiring` | 117.45 | **44.21** | **−73.24** |
| `ui_next_web` | 120.79 | **53.31** | **−67.48** |
| `backend_builds` | 16.26 | 3.89 | −12.37 |
| `ui_studio_desk` | 7.99 | 7.08 | −0.91 |
| `compose_up` | 44.43 | 43.65 | −0.78 |
| `secrets_provision` · `clones_and_inject` · `host_preflight` · `seed_tooling` | | | −0.44 combined |
| `autoverify` | 2.26 | 2.34 | +0.08 |
| `set_dress` | 81.61 | 82.04 | +0.43 |
| **TOTAL** | **449.51** | **286.99** | **−162.52** |

**L1's own attribution is the UI tier: 246.23 → 104.60 s, −141.63 s**, against the ~132.6 s the export/unpack
arithmetic predicted — slightly *better*, because a smaller final stage also shortens the layers around the
export, not only the export itself. The UI tier's share of the cycle falls from **54.8 % to 36.4 %**.

**`backend_builds`' −12.37 s is NOT L1's and must not be credited to it.** Its iter-08 range was 3.40–34.77
(a p50 of 16.26 sitting mid-spread); here it is 3.75–3.92. That is variance collapsing, not a lever.
Attribution is the deliverable, and the honest figure for L1 is 141.63 s.

**The next target is now `set_dress` at 82.04 s — 28.6 % of the cycle and the largest single phase.** That is
**L5**'s (the taxonomy replay), and the ranking has changed underneath the plan: L5 was priced at ~30–50 s
and ranked fifth.

## Close — 2026-08-12

**Outcome:** **L1 landed and the exit gate is MET.** The two Next images are multi-stage
`.next/standalone` builds from rext-owned Dockerfiles — **4.04 GB → 417 MB** and **3.94 GB → 380 MB**,
`exporting to image` **136.4 s → 3.8 s** combined — with both apps proven behaviourally identical to the
images they replace (hiring's `/login` is **byte-for-byte** 426,914 in both). The `n=3` cold campaign from
the pinned clone reads **p50 286.99 s** (min 280.99 / max 303.44) against a **360 s** gate and a **300 s**
stretch, with `green:true / 0 warnings`, **HEADROOM OK** and **ISOLATION OK** on **all three** reps, host
identity `match` ×3, 0 platform-repo edits and 0 refused demo-patches. `ASSERT-M257-isolation-with-L1`
shipped **with** the lever, as `TOK-01` requires — it had no implementation at all before this iter — and
was proven able to fail by 19 unit controls plus 3 live ones. Unlike iter-08's baseline this campaign is
**not** contended-and-labelled: it passes on its own terms.
**Type:** tik
**Status:** closed-fixed
**Gate:** MET
**Phase 5 grading:** (1) gate-met: **y** *(p50 286.99 ≤ 360 and ≤ the 300 stretch; green/0-warnings ×3; both falsifiable asserts OK ×3; 0 platform edits; 0 refused patches; identity match ×3)* — (2) triggered-tok: n — (3) re-scope: n *(the trigger reads p50 > 400 s after L1+L2+L3; L1 alone reached 286.99)* — (4) user-blocker: n — (5) cap-reached: n *(tik 1 of 5)* — (6) protocol-stop: n — (7) budget-exhausted: n — **Outcome: exit-1**
**Decisions:** see [`decisions.md`](decisions.md) (D1–D6)
**Side-deliverables:**
- **24 corpus citations re-anchored** after this iter's own line shifts, from a positional HEAD→worktree line map; 51/51 now resolve to intended content. One of them (`buildbench.py:1464`, cited as "the argparse") had been silently wrong **before** this iter.
- **11 knob anchors** in `demo-up-defaults.md` regenerated with `demo_knob_guard.py --fix` (the sanctioned path), which then reported both directions in agreement.
- `claim_census_guard` denominators 708→710 / 983→985, numerators unchanged — the documented maintenance shape, with provenance recorded in the paragraph that predicts it.
- `corpus/ops/demo/build-budget.md`: the image-isolation invariant section now describes a shipped assert rather than an intention; a net-new *"A/B-ing two Dockerfiles: the cache makes the second arm free"* section; and the harness block's two stale claims corrected (`macmini.json` exists; `profile_describes_host` does compare profile to machine).
- `corpus/ops/demo/frontend-tier.md`: next-web moved from build shape 1 to shape 3, with the `--env-mode=loose` hazard written down.
**Routes carried forward** (Fate 3, named handlers):
- **`LEVER-M257-L5-setdress`** → **the ranking changed underneath the plan.** `set_dress` is now the largest single phase at **82.04 s = 28.6 %** of the cycle, where L5 was priced at ~30–50 s and ranked fifth. The gate is met, so this is optimisation beyond it — but it is where the next second lives, and it is also **the chief win on the `/dev-up` path**.
- **`FIX-M257-anchor-guard-content-drift`** → `anchor_construct_guard` books a citation that lands on a *non-construct*, so one that slides onto a *different* construct passes silently: it caught 5 of the 24 this iter broke. Mechanically detectable with the same ~15-line positional line map.
- **`FIX-M257-committed-env-ships-real-clerk-pk`** → `apps/web/.env` is a committed repo file carrying the platform's own real-Clerk publishable key, and it ships into every demo image. Pre-existing (measured identically in the pre-L1 iter-08 image); the bundle correctly bakes only the minted key because `.env.local` wins at build time. Not a live defect, one build-order accident from the M218 iter-03 incident.
- iter-08's routes carry unchanged: `FIX-M257-sweep-scratch-pollutes-census` (still the dominant stack-core cluster), `RATCHET-M257-literal-ceilings-breached` (**deliberately not raised**; my +1 was taken back instead), `FIX-M257-demopatch-sha-baselines-drifted`, plus iter-05/06/07's tail.
**Lessons:**
- **Run a new assert against a real artefact, not only against fixtures you also wrote.** 15 unit tests were green over a bundle probe that read *nothing* — it passed `grep --exclude`, a GNU flag busybox does not implement — and a live control caught it in about a minute. A fixture's injected reader never touches busybox.
- **`[]` and `None` are different answers.** "Scanned and clean" and "could not scan" must not be the same value, or every probe failure reads as a pass.
- **Scope a new clause to what the gate actually says.** ISOLATION says *"**baked**"*; a committed `.env` is carried, not baked. Scoping there kept the clause from being permanently red for a condition it was not written about — and the finding was **routed with its own handler** rather than disappearing into the scoping.
- **A green test that changed subject is worse than a red one.** L1 made two whole-file locators ambiguous: one failed loudly (and was a signal about my change), two kept passing while grading a weaker proposition than their own messages state.
- **Editing a heavily-cited file is a corpus edit.** ~170 added lines moved 24 of 61 citations onto different content, and the fence — by construction — saw 5.
- **Predict, then measure, and keep them apart.** The export/unpack arithmetic predicted ~132.6 s; the UI tier realized 141.63 s. `backend_builds` also fell 12.37 s and that is variance, not L1 — crediting it would have made the lever look better than it is.
