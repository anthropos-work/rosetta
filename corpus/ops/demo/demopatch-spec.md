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
| **G2 — pre-patch drift-refuse + single-occurrence anchor** | `sha256(target)` must be one of `{pre_sha256, post_sha256}`; a third hash means the clone **drifted** → refuse. `post_sha256` alone is not enough to count as *patched* — the `post_marker` must also be present, so a hash collision on a partial apply cannot masquerade as success. If pristine, the **anchor must occur EXACTLY ONCE**: zero → refuse ("content drift"); two or more → refuse ("ambiguous — refusing to choose a hunk"). |
| **G3 — never-commit / working-tree-only** | The tool never runs `git add/commit/push/tag` — **a unit test greps its own source for any mutating git verb**. The only `git checkout` is the `-- <path>` working-tree form, isolated in one function precisely so the grep can whitelist it. After writing, it asserts the file is modified **and unstaged**; if not, it refuses *and reverts its own write*. |
| **G4 — idempotent re-apply** | The demo clone **persists** across `/demo-up`. An already-patched target (post-sha **and** marker) is a no-op, exit 0. |
| **G5 — content-anchored self-revert** | `revert` swaps `replacement → anchor` and then **re-asserts** `sha256 == pre_sha256`. Already-pristine is a no-op. A file matching *neither* pre nor post is refused — *"manual drift; refusing to guess"*. `--force-pristine` falls back to `git checkout -- <path>` (a working-tree restore, never a history operation). |
| **G6 — demo-only scope** | The manifest must declare `scope: demo`, and the workspace must be a demo workspace. Note the **structural** check is the one that actually fires at fresh-build time — the unified registry has no `demo-N` row yet when patches are applied. |
| **G7 — apply post-condition** *(unnamed until this spec)* | After the in-memory swap and **before writing a byte**: the patched text's sha must equal `post_sha256`, and the `post_marker` must be present. |

> **No write path bypasses G1 + G2.** `apply` runs both before it writes anything.

---

## 3. The manifest

A deliberately tiny **strict YAML subset** — parsed by a hand-written loader, **not PyYAML** (rext's stdlib-only
supply-chain rule). Top-level `key: scalar` and `key: |` literal blocks only; nested maps, flow collections, and
anchors are errors.

**All ten keys are mandatory.** There are no optional keys, and a present-but-empty value fails.

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
| **`demopatch`** (the tool) | the three `next-web-app` patches | the target lives **inside** the demo workspace → G1/G6 pass |
| **`stack-injection/apply-app-*.sh`** | the two `app` patches | the target is the **build-scratch** clone (`stacks/demo-N/clones/app`), which is **outside** the demo workspace → **`demopatch`'s own G1/G6 correctly REFUSE it**. The shell helpers re-implement the same guard ladder against **the same canonical manifest** — the manifest stays the single source of truth; only the vehicle differs |
| **`stack-injection/apply-ant-academy-dev-origins.sh`** | `ant-academy-dev-origins` | ant-academy runs **natively** (`next dev`), not baked into an image → the patch must **persist for the process lifetime** → apply-before-launch, revert-on-stop |

**Exit codes differ by vehicle.** `demopatch` uses `1` (guard refuse) and `2` (manifest/OS error). The shell helpers
use a richer space: `0` applied-or-already-patched · `1` manifest/target missing · `2` **pre-sha drift** · `3` anchor
count ≠ 1 · `4` replacement was a no-op · `5` patched sha ≠ post · `6` post_marker absent.

### The chain rule

`next-web-public-website-url` targets the **same file** as `next-web-studio-url`, so **its `pre_sha256` IS studio's
`post_sha256`.** It must be applied **after** studio and reverted **before** it.

> ⚠️ **It therefore reads "DRIFTED" against a pristine file BY DESIGN. Do not "fix" this.** A unit test fences it.

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

---

## 5. The patch inventory

| id | target | what it does |
|----|--------|--------------|
| `next-web-studio-url` | `next-web-app` · `packages/core-js/src/constants/urls.ts` | the Studio nav link stops ejecting to `studio.anthropos.work` |
| `next-web-public-website-url` | **same file — CHAINED** | the sim drill-down stays demo-local |
| `next-web-members-pagination` | `next-web-app` · `InsightsContext.tsx` | the enterprise members fetch `limit: 1000 → 30` |
| `app-targetrole-authz-skip` | `app` · `internal/roles/roles.go` | short-circuits a per-member Sentinel RPC on the **read** path → members grid **76.7 s → 0.51 s**. Mutations still enforce |
| `app-aireadiness-snapshot-loadmembers` | `app` · `internal/workforce/ai_readiness.go` | bounds the frozen-read member hydration to the ~199 snapshot users instead of the whole org → the **180 s** AI-readiness read completes. **Data-identical** |
| `ant-academy-dev-origins` | `ant-academy` · `code/next.config.js` | admits a `--public-host` demo's MagicDNS origin to `next dev` |

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
