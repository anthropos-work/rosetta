# The demo-patch mechanism (`demopatch`)

**The sanctioned zero-platform-edit escape hatch.** When a demo needs a fix that has **no env / config / compose
seam** — the value is *baked into platform source* — `demopatch` patches the demo's **own ephemeral clone** just
before the image build, lets the image bake the fix, and reverts the clone. The canonical `anthropos-work` repos are
**never touched**, and the rule *"the platform stays read-only"* survives intact.

> **Status:** v1.0 — authored M217 / v2.3 "cue to cue" (2026-07-13). Born M42m / v1.10 "method acting".
> **Code:** `rosetta-extensions/demo-stack/patches/` (`demopatch`, `manifest_loader.py`, one dir per patch) +
> `rosetta-extensions/stack-injection/apply-*.sh` (the two other apply vehicles — see §4).
> **Related:** [`safety.md`](../safety.md) · [`rosetta_demo.md`](../rosetta_demo.md) ·
> [`coverage-protocol.md`](coverage-protocol.md) (the fix-surface routing table that *routes* work here) ·
> [`frontend-tier.md`](frontend-tier.md)

---

## 1. Why it exists

Some demo-believability and demo-perf fixes are **platform-bound**: the value is a compile-time constant in platform
source, with no `NEXT_PUBLIC_*` / env / compose override. The founding case (M42m): next-web's `STUDIO_URL` is a
`NEXT_PUBLIC_NODE_ENV` ternary that bakes `https://studio.anthropos.work` into the bundle. A demo's "Studio" left-nav
link therefore **ejects the presenter to production** — the exact prod-eject the coverage gate forbids — and the only
clean fix is a one-line source edit, which the zero-platform-edit line forbids.

The resolution: **patch the demo's own gitignored, throwaway clone.** The image carries the fix; the clone is left
git-clean; the canonical repo never sees it. The mechanism's entire purpose is that its guards make *"touch only a
demo's ephemeral clone, never anything else"* **mechanically enforced** rather than merely intended.

**A demo-patch is the LAST resort, and it is DISCLOSED.** The order of preference is:

1. an **env / compose / injection** fix (rext-owned, no source touched) — always try this first;
2. a fix in **rext's own code** (Clerkenstein, the injector, the seeder);
3. **a demo-patch** — only when the value is genuinely baked into platform source;
4. **escalate.** Never edit a platform repo.

Every patch manifest carries a **DISCLOSED** header stating what the *real* platform fix would be, so that patching a
demo never quietly erases a genuine platform finding.

---

## 2. The guards

Seven guards. Six are named in the tool's own contract; the seventh (the apply post-condition) is real but was
unnamed until this spec.

| Guard | Enforces |
|-------|----------|
| **G1 — hard path-assert (demo-clone only)** | The target is canonicalized with `realpath` (which resolves `..` **and every symlink**, killing symlink escapes) and must resolve **inside the stack workspace**; the path must equal the manifest's `repo`/`path` **exactly** — no globs, no traversal. The manifest-derived path is *re-*canonicalized and re-checked, so a `repo: ../stack-dev/…` manifest is refused. The loader independently rejects `..`/absolute paths at parse time. |
| **G2 — the ANCHOR gate** *(rewritten M217-close)* | **The anchor is the contract; the whole-file sha is only a baseline.** The **anchor must occur EXACTLY ONCE**: zero → refuse (*the code being patched is gone*); two or more → refuse (*ambiguous — refusing to choose a hunk*). A **drifted whole-file sha with an intact anchor is NOT a refusal** — it self-heals (§6). Counting a target as *already patched* is a **coherence** probe, not a marker sniff: the whole replacement must be present **and** the anchor gone; otherwise the target is **PARTIALLY PATCHED or CORRUPT** and is refused. **Both vehicles enforce this identically** — `demopatch` and `apply_patch.py` were converged at the M217 close, because leaving `demopatch` on the old sha gate would have shipped the identical rot on the three next-web patches. |
| **G3 — never-commit / working-tree-only** | The tool never runs `git add/commit/push/tag` — **a unit test greps its own source for any mutating git verb**. The only `git checkout` is the `-- <path>` working-tree form, isolated in one function precisely so the grep can whitelist it. After writing, it asserts the file is modified **and unstaged**; if not, it refuses *and reverts its own write*. |
| **G4 — idempotent re-apply** | The demo clone **persists** across `/demo-up`. An already-patched target is a no-op, exit 0. **"Already patched" is G2's COHERENCE probe** — *the whole replacement present **and** the anchor gone* — **not a post-sha match.** <br>⚠️ *This row used to read "post-sha **and** marker", i.e. exactly the whole-file-sha check that §6 spends a section explaining ROTS. It contradicted G2 in the same table. Corrected at the M219 close; the two rows now describe one mechanism.* |
| **G5 — content-anchored self-revert** | `revert` swaps `replacement → anchor` and then **re-asserts** `sha256 == pre_sha256`. Already-pristine is a no-op. A file matching *neither* pre nor post is refused — *"manual drift; refusing to guess"*. `--force-pristine` falls back to `git checkout -- <path>` (a working-tree restore, never a history operation). |
| **G6 — demo-only scope** | The manifest must declare `scope: demo`, and the workspace must be a demo workspace. Note the **structural** check is the one that actually fires at fresh-build time — the unified registry has no `demo-N` row yet when patches are applied. |
| **G7 — apply post-condition** *(unnamed until this spec; made real at the M217 close)* | The write is **atomic** (`tmp` + `fsync` + `os.replace`) and the post-condition is verified against **the bytes that actually landed on disk**, not against the in-memory object. On mismatch the **pristine file is restored**. <br>*It was previously a tautology*: it re-hashed the same in-memory string `classify()` had just hashed, so it could not fail and its exit code was unreachable — while the real exposure (a truncate-in-place write with no rollback, leaving half-written source on a short write/ENOSPC/SIGINT) went unguarded. |

> **No write path bypasses G1 + G2.** `apply` runs both before it writes anything.

---

## 3. The manifest

A deliberately tiny **strict YAML subset** — parsed by a hand-written loader, **not PyYAML** (rext's stdlib-only
supply-chain rule). Top-level `key: scalar` and `key: |` literal blocks only; nested maps, flow collections, and
anchors are errors.

**All ten keys are mandatory.** There are no optional keys, and a present-but-empty value fails. **A duplicate key is refused at load** (M217-close): the loader was previously *last-wins*, so a manifest with two `pre_sha256:` lines let `--repin` rewrite the first while the loader returned the second — and, far worse, a duplicate `anchor:` could **steer which hunk gets replaced in platform source**. *An ambiguous manifest is not a manifest.*

| Key | Meaning |
|-----|---------|
| `id` | the patch id; by convention `patches/<id>/<id>.yaml` |
| `repo` | the demo clone dir under the workspace root (e.g. `next-web-app`) |
| `path` | the file inside that clone |
| `pre_sha256` | sha256 of the **whole pristine file**, 64 **lowercase** hex |
| `post_sha256` | sha256 of the **whole file after** the single replacement |
| `anchor` | block scalar — the **exact** pre-image hunk. Must occur **exactly once** |
| `replacement` | block scalar — the post-image hunk |
| `post_marker` | a substring present only in the patched form — the positive idempotency probe (G4). Rejected at load if absent from `replacement` |
| `build_env` | a build-time env line the **caller** appends to the `.env.local` overlay, offset-templated (`$((9000+OFFSET))`). Stored verbatim; the caller expands it. Source-only patches set it to an inert comment (it is mandatory, so it cannot be omitted) |
| `scope` | must be literally `demo` (G6) |

**Tabs survive.** The loader dedents by **spaces only**, so the literal tab bytes of Go source are preserved verbatim
inside a block scalar. The Go manifests depend on this.

**The design rule visible in every manifest:** the replacement is **behavior-identical when the env var is unset**
(prepend `process.env.X ||`, keep the original as the fallback). That is what lets a *dynamic* value (an offset port,
a MagicDNS host) coexist with a *static* `post_sha256`.

---

## 4. Three apply vehicles (the most under-documented fact)

Not every patch is applied by `demopatch` itself, and this surprises people.

| Vehicle | Patches | Why |
|---------|---------|-----|
| **`demopatch`** (the tool) | the **eight** `next-web-app` patches (6 × `apps/web` + 2 × `apps/hiring`) | the target lives **inside** the demo workspace → G1/G6 pass |
| **`stack-injection/apply-app-*.sh`** | the two `app` patches | the target is the **build-scratch** clone (`stacks/demo-N/clones/app`), which is **outside** the demo workspace → **`demopatch`'s own G1/G6 correctly REFUSE it**. The shell helpers re-implement the same guard ladder against **the same canonical manifest** — the manifest stays the single source of truth; only the vehicle differs |
| **`stack-injection/apply-ant-academy-dev-origins.sh`** | `ant-academy-dev-origins` | ant-academy runs **natively** (`next dev`), not baked into an image → the patch must **persist for the process lifetime** → apply-before-launch, revert-on-stop |

**Exit codes differ by vehicle.** `demopatch` uses `1` (guard refuse) and `2` (manifest/OS error). The shell helpers
use a richer space: `0` applied-or-already-patched · `1` manifest/target missing · `2` **pre-sha drift** · `3` anchor
count ≠ 1 · `4` replacement was a no-op · `5` patched sha ≠ post · `6` post_marker absent.

### The chain rule

`next-web-public-website-url` targets the **same file** as `next-web-studio-url`, so **its `pre_sha256` IS studio's
`post_sha256`.** It must be applied **after** studio and reverted **before** it.

> ⚠️ **It therefore reads "DRIFTED" against a pristine file BY DESIGN. Do not "fix" this.** A unit test fences it.

**The chain runs on BOTH frontend builds (M224).** The `urls.ts` pair is applied by `build_frontend_next_web`
**and** `build_frontend_hiring` — the Studio nav link is in the **shared `packages/ui` NavBar** (`key: STUDIO_URL`),
so the hiring image ejects to `studio.anthropos.work` unless the same pair bakes into it. The apply-order (studio →
public-website) and revert-order (LIFO) are identical on both; each build carries its own patch-set fingerprint
(§5-bis) — next-web's over its manifest set, hiring's over the **4-manifest union** (the 2 `apps/hiring` patches +
this shared pair). A test fences the hiring-side chain apply-order + LIFO revert + the 4-manifest fingerprint union.

### The `app` patches are never reverted — and that is correct

The `next-web` patches are reverted by a `RETURN` trap (LIFO) so the persistent clone is left git-clean. The `app`
patches have **no revert**: the build-scratch clone is **force-checked-out at the newest `v*` tag on every bring-up**,
which discards the previous run's injections wholesale.

### Opt-outs

`DEMO_NO_PATCH=1` (the next-web set) · `DEMO_NO_AUTHZ_SKIP=1` · `DEMO_NO_AIREADINESS_LOADMEMBERS_BOUND=1`.

### Caller convention: default-on, **non-fatal**

A refused patch **warns and continues** — it never aborts a good bring-up.

> ⚠️ **This is exactly how two perf patches rotted silently for four releases** (see §6). Non-fatality is right; a
> refusal that is **invisible** is not. A refusal must be **loud**.

> ✅ **M219 closed the last silent hole in this** (`TEST-M219-freshness-gate-skips`). The freshness preflight was
> wrapped in `if [ "${DEMO_NO_PATCH:-0}" != 1 ] && [ -d "$DEMO/app/.git" ]` **with no `else`** — so on a box with
> no app clone the entire block, *including its only success line*, was skipped and **nothing was printed at
> all**. The bring-up log was byte-identical to one where the gate had run and passed: zero anchor-drift
> protection, silently. Both branches now speak — a deliberate `DEMO_NO_PATCH=1` says so, and a missing clone
> says **"NOT RUN … this is NOT a pass"**. Separately, three unit tests skipped themselves with a message
> deferring to a *"live-verify gate"* that **does not exist** (grep-confirmed); they now report themselves as
> **coverage holes, not passes**. A skip is not a pass — the same rule as the alignment surfaces'
> *absence-of-a-score* (`alignment_testing.md`).

---

## 5. The patch inventory

**11 patches: 8 × `next-web-app` (6 × `apps/web` + 2 × `apps/hiring`) · 2 × `app` · 1 × `ant-academy`.**

> **The `apps/hiring` patches are M224 "the callback" (v2.4 "casting-call").** The demo now runs the
> **real Hiring app** as a second UI container (TOK-02 — the two-app demo), so a recruiter hero lands on the
> genuine `apps/hiring` candidate-comparison Results screen instead of a re-skinned workforce fake. **The HIRING
> image (`build_frontend_hiring`) bakes FOUR patches**, not two: the **2 net-new** `apps/hiring` patches
> (`next-hiring-role-remap`, `next-hiring-members-pagination`) **plus the 2 chained shared `urls.ts`** patches
> (`next-web-studio-url` → `next-web-public-website-url`), applied on the hiring build too because the Studio nav
> link lives in the **shared `packages/ui` NavBar** (`key: STUDIO_URL`) — so an unpatched hiring image ejects the
> presenter to `studio.anthropos.work` exactly as `apps/web` did. Found + killed at iter-13 (the hiring image's
> client chunks were `docker exec`-grep-verified to carry **0** `studio.anthropos.work`; the trustworthy render
> probe of iter-12 had surfaced the eject the earlier broken probe hid). All four ride `build_frontend_hiring`'s
> transient LIFO apply/revert, fenced by a **4-manifest patch-set fingerprint union** (§5-bis) that forces a
> rebuild if any of the four moves. The 2 net-new `apps/hiring` patches are the **same class as a known `apps/web`
> patch** — the same monorepo (`next-web-app`), the same defect the web app already fixed, never mirrored onto
> hiring. *(The distinct-manifest total is unchanged at **11**: the chained `urls.ts` pair is shared — counted once
> under the 6 × `apps/web` — and merely applied on **both** frontend builds. The prior count line read "8 patches /
> 5 × next-web-app" — it undercounted by the `next-web-no-thirdparty` row; corrected here.)*

| id | target | what it does |
|----|--------|--------------|
| `next-web-studio-url` | `next-web-app` · `packages/core-js/src/constants/urls.ts` | the Studio nav link stops ejecting to `studio.anthropos.work` |
| `next-web-public-website-url` | **same file — CHAINED** | the sim drill-down stays demo-local |
| `next-web-ssr-graphql-origin` | `next-web-app` · the SSR GraphQL origin → `WUNDERGRAPH_SSR_ENDPOINT` | **(M218) THE fix for the 38-second login.** The SSR pass fetched the **public MagicDNS origin from inside the container**, where the tailnet IP **blackholes** (ts-input drops the SYN-ACK on the docker bridge) → ~37 s per authenticated render. **Only manifests on a `--public-host` demo.** ⚠️ **This row was missing from this inventory for a full release** — the highest-impact patch on the box, absent from the doc that calls itself the contract |
| `next-web-members-pagination` | `next-web-app` · `InsightsContext.tsx` | the enterprise members fetch `limit: 1000 → 30` |
| `next-web-aireadiness-flag-gate` | `next-web-app` · `components/ai-readiness/data/useAiReadinessActive.ts` | **(M219)** the **member** readiness surface never mounts on a demo: a demo bakes no PostHog, so `useFeatureFlagEnabled()` is `undefined` **forever** and the code demands `=== true`. Treats *"PostHog unconfigured"* as *"no rollout gate"*; the ORG boolean still decides. **Behaviour-identical wherever PostHog IS configured.** Targets its **own** file — does **not** chain with the `urls.ts` pair |
| `app-targetrole-authz-skip` | `app` · `internal/roles/roles.go` | short-circuits a per-member Sentinel RPC on the **read** path → members grid **76.7 s → 0.51 s**. Mutations still enforce |
| `app-aireadiness-snapshot-loadmembers` | `app` · `internal/workforce/ai_readiness.go` | bounds the frozen-read member hydration to the ~199 snapshot users instead of the whole org → the **180 s** AI-readiness read completes. **Data-identical** |
| `next-web-no-thirdparty` | `next-web-app` · `apps/web/src/app/layout.tsx` | **(M220 S6/g) stops the demo phoning home.** The root layout hardcodes **four** third-party scripts with **no env seam of any kind** — `plausible.io`, `analytics.bellasio.com`, `uptime.betterstack.com`, and `<GoogleTagManager gtmId='GTM-PXRTBZK'/>` (which itself loads **Google Analytics, DoubleClick, Google Ads and LinkedIn Ads**). They fire on **every page load**, so a presenter demoing to a customer silently ships that customer's page views to **seven** third parties, from a demo the corpus calls self-contained. The patch wraps all four in one build-time env gate (`NEXT_PUBLIC_DISABLE_THIRD_PARTY_SCRIPTS`, baked to `1`); every tag is preserved byte-for-byte inside it, so the behaviour is **identical when the var is unset**. Targets its **own** file — no chain. *The plan named only the 4 GTM ad networks; reading the file found 3 more vendors on top — the D17 signature again.* |
| `next-hiring-role-remap` | `next-web-app` · `apps/hiring/src/context/UserStatusContext.tsx` | **(M224 tik C) the recruiter reaches the hiring enterprise Results routes.** `apps/hiring` stores the Clerk org-role RAW (`role: userRole` = `org:admin`) where `apps/web` **remaps** it (`remapUserRole('org:admin') → 'admin'`). So an admin recruiter reads as **non-admin** in the hiring app, `EnterpriseWrapper` bounces her to the candidate Home, and **0 insights rows** render. The patch adds the same remap (nested, string-literal casts — `apps/hiring` imports `MembershipRoles` **type-only**). **NOT Clerkenstein** (`org:admin` is faithful to real Clerk RBAC), **NOT the seeder** (Rae is already `role='admin'`). Targets its **own** file — no chain |
| `next-hiring-members-pagination` | `next-web-app` · `apps/hiring/src/context/InsightsContext.tsx` | **(M224 tik D) the Results dashboard stops hanging on the loading spinner.** The exact **mirror of `next-web-members-pagination`**: `apps/hiring`'s InsightsContext fetches `useGetOrganizationMembers({ limit: 1000 })` — an unbounded whole-org fetch the activity-dashboard layout **blocks** on (`if (loading) return <BaseLoading/>`), and its `GET_MEMBERS` query resolves `targetRole` **per row** — so the per-sim scoreboards never mount. Caps the fetch `1000 → 30`. The **per-member Sentinel authz half of the wall needed NO new patch**: the hiring app hits the **same shared `app` backend** that already bakes `app-targetrole-authz-skip`, so `targetRole`'s per-object RPC is already dropped for this path too. Targets its **own** file — no chain |
| `ant-academy-dev-origins` | `ant-academy` · `code/next.config.js` | admits a `--public-host` demo's MagicDNS origin to `next dev` |

---

## 5-bis. The image cache had no idea which patches were in it (M220 S6/g)

**A patch that applies perfectly can still never reach the demo.** This is the mechanism behind this document's
own war story — *"a silently-refused perf patch shipped a 76 s members grid for four releases"* — and it is not
about refusal at all.

`build_frontend_next_web` **reuses** a cached `demo-N-next-web` image when two things still match: the baked
offset endpoint, and the minted publishable key. **Neither has any relationship to the demo-patch set.** So an
image built *before* a patch was added — or before a patch's sha was re-pinned — passes both checks and is
reused. The patch is applied to the clone, dutifully reverted afterwards, and **never reaches the image**. The
bring-up reports success. The bundle is unpatched. Nothing anywhere says so.

> **It was about to happen again.** The `demo-1-next-web` image on `billion` already carried a matching endpoint
> and pk, so the first bring-up after adding `next-web-no-thirdparty` would have **reused it** and served a
> bundle still phoning home to all four vendors — *while grading green*.

**The fix: a PATCH-SET FINGERPRINT.** The sha256 of the manifest set (each manifest's own sha256, plus the
`DEMO_NO_PATCH` opt-out) is baked into the image as a **label** (`demo.patchset`) and compared on reuse. A label
is image metadata, so it needs **no Dockerfile edit** — the zero-platform-edit line holds (the repo stays a build
*context* only). Change a patch, re-pin a hash, add a manifest, or flip `DEMO_NO_PATCH` ⇒ the label moves ⇒
**rebuild**. An image with **no** label predates the fingerprint and is treated as a mismatch (fail-safe: a
needless ~3 min build is far cheaper than serving an unpatched demo to a customer).

It fired on its first live run:

```
next-web: cached image demo-1-next-web was built with a DIFFERENT demo-patch set
  (<none: predates the fingerprint> != cee1e4ff…) — removing + rebuilding so the current patches
  actually bake into the image.
```

**The rule this adds to §4's ladder:** *applying a patch is not shipping it.* Adding a manifest to the apply
ladder and forgetting it in the fingerprint call re-opens the same hole one level up — so a fence asserts the two
sets agree.

---

## 6. The freshness gate — and why the whole-file sha rots

**The failure this section exists to prevent** (found M217, after it had been live for four releases):

> Both `app` perf patches **silently refused on every single run**. They were pinned against app v1.295/v1.315; the
> box was building v1.337. The applier printed the exact sha mismatch to stderr — and **the caller piped stderr to
> `/dev/null`**. The demo shipped with a 76-second members grid and a 180-second AI-readiness read, and *nothing
> said so*. Four bring-up logs carried the warning. Nobody saw it.

### Why a static pin cannot work

`pre_sha256` hashes the **whole file**. But the demo builds the scratch clone at **"the newest `v*` tag on this
box"** — so *any* unrelated edit anywhere in that file, in any app release, breaks the pin. Worse:

> **`internal/workforce/ai_readiness.go` is not byte-identical between app v1.334.1 and v1.337.0.** Two boxes on two
> app tags ⇒ **no single committed whole-file pin can be correct on both.** The manifest schema cannot express the
> truth, and a one-shot re-pin cannot fix it.

Meanwhile **the anchor survives every tag tested**, occurring exactly once. The *semantic* target is stable; only the
whole-file proxy rots.

### The freshness preflight — it runs BEFORE the clone (M217-close)

A dedicated preflight runs **before the inject loop**: it resolves the app tag this box will build, reads each
patch target **straight out of git** (`git show <tag>:<path>` — no clone, no checkout, no build), and runs the
gate in `--check` mode. **A broken anchor aborts there**, in seconds, instead of minutes into a build that has
already done `make init`, the secret provision, a clone, and a `checkout -f`.

> This was **promised and checked off in M217's own plan, and never built** — the gate ran *inside* the loop.
> It is built now. (Finding that a checked box described code that did not exist is exactly the class of
> false-claim this milestone's first section exists to delete.)

The preflight honours the same `DEMO_NO_*` opt-outs as the appliers: a deliberate no-patch run is never blocked
by a gate.

### The gate (decided M217 — the self-healing gate)

**The anchor is the contract; the sha is a baseline.**

- A **freshness preflight** runs **before** the inject loop. For each patch it resolves the tag *this box will
  actually build*, hashes the target, and compares.
- **Whole-file sha drifted, but the anchor still occurs exactly once** → the patch is still semantically valid.
  **Recompute `pre`/`post` for this box, report the drift LOUDLY, and apply.** The demo comes up green on any box at
  any app tag.
- **The anchor is gone (0×) or ambiguous (2+×)** → the code being patched has genuinely changed. **ABORT, loudly.**
  This is a real semantic break and a human must look at it.
- **G7 still holds**: the post-condition is verified against the *recomputed* post-sha, so a bad swap still cannot be
  written.
- `--repin` rewrites the manifest's recorded baseline. The escapes (`DEMO_NO_*`) bypass the preflight — a
  deliberate no-patch run must not be blocked by a gate.

> **Why not keep the hard sha gate?** It would abort a bring-up on every app release, and — because the boxes are on
> different tags — a pin committed from one box would abort the other. It was protecting against "something else in
> the file changed", which for a **perf-only, read-path, data-identical** shortcut in a **demo** is a proxy, not a
> real protection. The anchor is the thing that carries meaning.

### Re-pin runbook

1. The preflight fails loud and prints the paste-ready corrected `pre_sha256` / `post_sha256`.
2. `apply_patch.py --repin` (or paste the two lines).
3. Commit + tag rext.

#### `--repin` works on an ALREADY-PATCHED target — and that matters

The natural workflow puts you there: you run a bring-up, see the **SELF-HEALED** notice with the corrected pins,
and *then* want to record them. But by that point the build-scratch clone **is patched**.

> **This used to silently do nothing** and print *"already patched (idempotent no-op)"* — so the operator
> believed the manifest had been updated when it had not. (Found in M217's hardening pass.)

`--repin` now **recovers the pristine form** by reversing the swap (the same content-anchored move G5's revert
makes) and **round-trip verifies** it: re-applying the patch to the recovered body must reproduce the current
file **byte-for-byte**. Only then does it write the pin.

- **Drift *outside* the patched hunk** (the common case — `app` churns elsewhere constantly) round-trips
  cleanly → it re-pins.
- **A hand-edit *inside* the patched hunk** does not round-trip → it **REFUSES (exit 1)**.
  **We do not write a pin we cannot prove.** Re-checkout the clone and try again.

`--repin` **never touches the target file** — only the manifest.

---

## 7. Adding a new patch

1. **Exhaust the alternatives first** (§1). A demo-patch is the last resort before escalation.
2. Write the manifest header as a **DISCLOSED** note: what the real platform fix would be, and why the demo cannot
   wait for it. The platform finding stays in the corpus.
3. Choose an anchor that occurs **exactly once** and is **semantically load-bearing** — it is now the gate.
4. Make the replacement **behavior-identical when its env var is unset**.
5. Pick the vehicle (§4) by where the target lives: inside the demo workspace → `demopatch`; the build-scratch clone
   → a `stack-injection/apply-*.sh` helper; a natively-run app → the apply/revert helper form.
6. **Add a live-clone pin test.** The absence of one for the two `app` patches is precisely what let the drift ship.
