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
| **G5 — journalled self-revert** *(rewritten M257x iter-90)* | **`apply` records the OBSERVED pre-image; `revert` restores exactly it.** Revert consults that journal FIRST and, for a patch this tool applied, needs no baseline at all — which is what makes it exact on a **drifted** base, the normal state of a persistent clone. The journal lives in the **workspace root** (never inside a clone), is consumed on success, and its directory is removed once it empties. **No journal ⇒ no guessing:** a target with no entry falls through to `pre_sha256`/`post_sha256` and is still refused — *"manual drift; refusing to guess"*. `--force-pristine` falls back to `git checkout -- <path>` (a working-tree restore, never a history operation) and is the one-time recovery for clones patched before the journal existed. <br>⚠️ *This row used to read "swaps `replacement → anchor` and re-asserts `sha256 == pre_sha256`" — a **whole-file-sha** gate, while G2 one row up had already moved to the **anchor**. The two rows contradicted each other, and the contradiction was live: on any drifted clone the patch applied and would not come off. See §6's FIXED block.* <br>⚠️ **G5 is a capability, not a sweep — the recovery rung (R1) that invokes it sweeps EVERY manifest on disk (directory-driven since v2.6 M237 — 23 today; was a hardcoded 3). See §2.1.** |
| **G6 — demo-only scope** | The manifest must declare `scope: demo`, and the workspace must be a demo workspace. Note the **structural** check is the one that actually fires at fresh-build time — the unified registry has no `demo-N` row yet when patches are applied. |
| **G7 — apply post-condition** *(unnamed until this spec; made real at the M217 close)* | The write is **atomic** (`tmp` + `fsync` + `os.replace`) and the post-condition is verified against **the bytes that actually landed on disk**, not against the in-memory object. On mismatch the **pristine file is restored**. <br>*It was previously a tautology*: it re-hashed the same in-memory string `classify()` had just hashed, so it could not fail and its exit code was unreachable — while the real exposure (a truncate-in-place write with no rollback, leaving half-written source on a short write/ENOSPC/SIGINT) went unguarded. |

> **No write path bypasses G1 + G2.** `apply` runs both before it writes anything.

### 2.1 The R1 recovery rung sweeps every manifest on disk (`F-M236-CLOSE-2`, closed v2.6 M237)

**Read this before adding a patch.** G5 above describes what `revert` *can* do. The rung that actually runs it
unattended — **R1**, the pristine-ing pass in `demo-stack/ensure-clones.sh` — is now **directory-driven**: it
iterates **every** `patches/<name>/<name>.yaml` (all 23 today), not a hand-maintained list:

```sh
for _mf in "$HERE"/patches/*/*.yaml; do
  [ -f "$_mf" ] || continue
  "$DEMOPATCH" revert "$DEMO" --manifest "$_mf" --force-pristine …
done
log "demopatch R1: swept $_r1_swept manifest(s) from $HERE/patches/ (directory-driven — F-M236-CLOSE-2)"
```

**What this fixed.** Through v2.5, R1 iterated a **hardcoded three-entry array** (`next-web-studio-url`,
`next-web-members-pagination`, `app-targetrole-authz-skip`) — about 20% of the 14 manifests under
`demo-stack/patches/`. The list never grew as patches were added, so the other **11** had **no unattended
recovery at all**.

**What R1 exists to catch.** A patch is applied just before an image build and reverted by the build's `RETURN`
trap. If a run dies **after apply but before the trap** (`Ctrl-C`, an OOM, a failed build that exits hard), the
clone is left **carrying the patch**. R1 is the next-bring-up sweep that restores pristine. Without it, the
leftover persists — and because G2's anchor gate then finds the anchor **gone**, the next apply is *correctly*
refused, so it surfaces as a **silently skipped patch** rather than a loud failure. `revert --force-pristine`
only restores-to-pristine (a no-op on a clean path; it **never applies**), so sweeping every manifest is safe by
construction; a manifest `demopatch` legitimately refuses (e.g. an `app` patch whose real vehicle is a
`stack-injection` shell helper on the build-scratch clone — §4) logs a **non-fatal skip**.

> 🔴 **This is not hypothetical, and it is the failure mode that costs the most to diagnose.** Measured
> 2026-07-20: both boxes were carrying leftover patches, in **disjoint** sets — **5** in the local
> `next-web-app` clone, **2** in `billion`'s `ant-academy` clone. Two boxes, two different sets, neither
> detected by anything. A silently-refused perf patch on exactly this path shipped a **76 s members grid for
> four releases** (§6). *A patch that is refused because a previous crash left it applied looks identical to a
> patch that was never wired.* Since M237, a stranded patch **outside** the old three is swept too. Proven live
> on `billion` (2026-07-21): `demopatch R1: swept 14 manifest(s) … directory-driven` (a v2.6-era count; 23 today — see §5).

**Consequences for an author adding a patch (v2.6+):**

1. **Nothing to wire.** A new `patches/<name>/<name>.yaml` is swept by R1 automatically — the directory *is* the
   list. (`TestR1SweepM237` fences the R1 glob against the real `patches/` count so a naming break is caught;
   **`TestPatchInventory` (v2.6 M238) additionally pins the EXACT inventory total + per-repo breakdown against §5**
   so an add/remove/mis-file drift goes RED until the doc and the fence's constants move together.)
2. **`--force-pristine` is invoked for every manifest.** Recovery is unattended for all patches; the manual
   `demopatch revert … --force-pristine` (or `git checkout -- <path>` in the demo clone) remains available.
3. **When a patch appears not to have applied, check for a stranded prior apply first** — `demopatch status
   <workspace> --manifest <m.yaml>` reports `pristine | patched | drifted | absent`. `patched` before the build
   means it was stranded (R1 should now have swept it — check the `swept N manifest(s)` line).

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
| **`demopatch`** (the tool) | the **eleven** `next-web-app` patches (3 × `apps/web` + 2 × `apps/hiring` + 3 × `packages/ui` + 2 × `packages/core-js` + 1 × `packages/graphql`) **+ the five `studio-desk` patches** (M249's back-to-cockpit/logout/logo trio — the FIRST studio-desk source patches — + M253's first-paint pair; `stack-demo/studio-desk/…` is inside `DEMO_WS`, image-baked by `build_frontend_studio_desk`) | the target lives **inside** the demo workspace → G1/G6 pass |
| **`stack-injection/apply-app-*.sh`** | the two `app` patches | the target is the **build-scratch** clone (`stacks/demo-N/clones/app`), which is **outside** the demo workspace → **`demopatch`'s own G1/G6 correctly REFUSE it**. The shell helpers re-implement the same guard ladder against **the same canonical manifest** — the manifest stays the single source of truth; only the vehicle differs |
| **`stack-injection/apply-ant-academy-*.sh`** / **`apply-academy-fs-published*.sh`** | the **five** `ant-academy` patches (`ant-academy-dev-origins`, `academy-fs-published-fallback`, `academy-fs-published-public`, `academy-fs-published-chapter-body`, **`ant-academy-back-to-cockpit`** — M249, `apply-ant-academy-back-to-cockpit.sh`) | ant-academy runs **natively** (`next dev`), not baked into an image → each patch must **persist for the process lifetime** → apply-before-launch, revert-on-stop (one shell helper each, same guard ladder, same canonical manifest) |

**Exit codes differ by vehicle.** `demopatch` uses `1` (guard refuse) and `2` (manifest/OS error). The shell helpers
use a richer space: `0` applied-or-already-patched · `1` manifest/target missing · `2` **pre-sha drift** · `3` anchor
count ≠ 1 · `4` replacement was a no-op · `5` patched sha ≠ post · `6` post_marker absent.

### The chain rule

`next-web-public-website-url` targets the **same file** as `next-web-studio-url`, so **its `pre_sha256` IS studio's
`post_sha256`.** It must be applied **after** studio and reverted **before** it.

> ⚠️ **It therefore reads "DRIFTED" against a pristine file BY DESIGN. Do not "fix" this.** A unit test fences it.

**The chain runs on BOTH frontend builds (M224).** The `urls.ts` pair is applied by `build_frontend_next_web`
**and** `build_frontend_hiring` — the Studio nav link is in the **shared `packages/ui` NavBar** (`key: STUDIO_URL`),
so the hiring image ejects to `studio.anthropos.work` unless the same pair bakes into it. Each build carries its
own patch-set fingerprint (§5-bis): **next-web's is now derived from a `patches/next-web-*/*.yaml` GLOB** (v2.8
M258 — see the warning below), while hiring's is still over an explicit **7** and studio-desk's over its own
**5**. A test fences the hiring-side chain apply-order and the fingerprint union.

> **⚠️ This paragraph used to state next-web's count as "9 manifests", and a COUNT was the wrong shape to
> record — because the list producing it was hand-maintained, and it rotted exactly as this count did.**
> The fingerprint decides whether a cached image is reused. When the PostHog-bootstrap patches were added,
> nobody added them to that argument list, so the fingerprint did not move, the bring-up logged *"image
> already built … reusing (no rebuild)"*, and the new patches never reached the bundle. The fix was written,
> tagged and deployed **three times** while the browser served the same stale image — reading each time as
> *"the patch does not work"* rather than *"the patch never ran"*, which sent the debugging into the patch
> (and twice into the wrong call site of a three-call-site flag) instead of into the cache.
>
> next-web's fingerprint is therefore no longer a list to keep in sync: **a `next-web-*` patch dir that
> exists is an input.** Over-hashing costs a needless rebuild — slow, not wrong — which is the correct side
> to err on for a cache key. `stack-core/patchset_fingerprint_guard.py` (guard family) holds the invariant
> for **every** frontend build: the manifests a build APPLIES must be a subset of the manifests its
> fingerprint HASHES.
>
> **And validate the manifest itself.** Both bootstrap manifests also shipped with `post_marker` set to the
> patch ID rather than to text from the replacement, so `demopatch` refused them outright — **non-fatally**,
> the same silent-no-op shape. `stack-core/tests/test_all_demopatch_manifests_load.py` now loads every
> manifest on disk in the test suite, where a failure is loud and cheap.

> **⚠️ Corrected at v2.8 M255** (pre-milestone KB-fidelity audit). This paragraph previously said hiring's
> fingerprint was over a **"4-manifest union"** — a C1 mirrored-count that drifted at **M232** (the interview
> pair) and again at **M249** (Back-to-Cockpit) — and that the apply and revert orders *"are identical on
> both"*, with revert being **LIFO**. Derived from the source, **neither half is true**:
>
> | build | apply order | revert order | strict LIFO? |
> |---|---|---|---|
> | `next-web` (9) | studio · pubweb · pagination · ssr-origin · aireadiness · thirdparty · interview-container · interview-result · back-to-cockpit | pubweb · studio · pagination · ssr-origin · aireadiness · thirdparty · interview-result · interview-container · back-to-cockpit | **no** |
> | `hiring` (7) | studio · pubweb · rolemap · pagination · interview-container · interview-result · back-to-cockpit | interview-result · interview-container · pagination · rolemap · pubweb · studio · back-to-cockpit | **no** |
>
> **And it does not need to be.** The only order-sensitive relationship in either set is the `urls.ts`
> chain — *pubweb must revert before studio* — and **both builds get that right**. Every other manifest
> targets a file no other manifest in its set touches, so its position is free. The real invariant is
> therefore **"the chain reverts in the right order"**, not "the whole list is LIFO", and that is what
> `rosetta-extensions/stack-core/union_apply_guard.py` now asserts, in both builds, in both directions.
>
> This matters beyond pedantry: v2.8's **union-apply** parallelism rule (see
> [`build-budget.md`](build-budget.md)) is stated as *"apply the union once, build both in parallel, revert
> once LIFO"* — and it was about to inherit a "LIFO" that had no referent in the code.

> **~~A related sharp edge, worth knowing before you rely on revert.~~ FIXED at M257x iter-90 — and the
> paragraph that stood here was wrong about the one thing that mattered.** It described the defect
> accurately: G5 compared the file against the manifest's **recorded** `post_sha256` while apply wrote a
> **recomputed** post whenever the self-healing freshness gate fired, so the paired revert refused, silently,
> into the `RETURN` traps' `/dev/null`. Then it concluded *"it is not currently harmful."*
>
> **It was harmful, and it was live.** M257x iter-88 routed four separate demo-stack failures as up to three
> classes; iter-89 collapsed all four to this one cause with a one-minute probe — `git status` on
> `stack-demo/next-web-app`, which had been left dirty by exactly this refusal. Three shipped manifests
> reported `status: patched` and `revert: neither pre nor post` **in the same breath**. The reasoning that
> made it look benign — *the ephemeral clone is force-checked-out on the next bring-up* — is true of the
> `app` build-scratch clone and **false of the persistent `next-web` clone this paragraph was about.**
>
> **The fix (`demopatch`, M257x iter-90): apply JOURNALS the observed pre-image; revert restores exactly it.**
> Revert no longer consults a recorded baseline at all for patches this tool applied, so it is exact on a
> drifted base — which is the normal state of a persistent, pulled clone. The journal lives in the
> **workspace root** (never inside a clone, which would defeat the promise it exists to keep), is **consumed
> on a successful revert**, and its directory is removed once it empties. **No journal means no guessing:** a
> target that drifted with no apply behind it still falls through to the baseline comparison and still
> refuses, so revert never became a blind restorer.
>
> **Known limitation, stated rather than hidden.** A patch applied *before* the journal existed has no entry
> and cannot be journal-reverted; those clones need the one-time `revert --force-pristine`. `apply`'s
> idempotent no-op now **WARNs** when it finds an already-patched target with no journal entry, naming that
> consequence, so the condition is self-describing at the moment it is created instead of costing two
> iterations to rediscover.
>
> **The finding that outlives the fix, and the reason it is recorded here rather than in a changelog:**
> **G2 (refuse on drift) and G5 (always self-revert) cannot both hold once the base is allowed to move** —
> and the base is always allowed to move. Every test asserted them **separately**, and each passed; the suite
> was green while the mechanism was broken. **A specification with seven guards needs at least one test per
> PAIR that can interact, not one test per guard.** Those pairs now live in
> `demo-stack/tests/test_demopatch.py::TestGuardConjunctions`, with a mutation control that rebuilds
> `demopatch` with the journal blinded and asserts the battery goes RED again with the original signature.

### The `app` patches are never reverted — and that is correct

The `next-web` patches are reverted by a `RETURN` trap so the persistent clone is left git-clean — **but not
in LIFO order**, and this sentence used to say it was. The next-web trap
(`demo-stack/up-injected.sh:825` @ rext `023fd38`) reverts *pubweb · studio · pagination · ssr-origin ·
aireadiness · thirdparty · interview-result · interview-container · back-to-cockpit* against an apply order of
*studio · pubweb · pagination · ssr-origin · aireadiness · thirdparty · interview-container · interview-result ·
back-to-cockpit* (`:841`, `:846`, `:865`, `:877`, `:889`, `:901`, `:915`, `:950`, `:963`): **only two adjacent
pairs are inverted — the `urls.ts` sha-chain and the interview flag-gate pair; the other five manifests are
reverted in apply order.** That is the same
finding the §4 correction box above records — see its table for the hiring lane, whose revert
(the hiring `RETURN` trap, `:1388`) is likewise not LIFO.

> ⚠️ **These ten anchors were re-derived 2026-08-24 (v2.10 M272) and the previous set was WHOLLY stale.**
> The trap read `:811` and the applies `:827/:832/:851/:863/:875/:887/:901/:936/:949`; measured against
> `up-injected.sh` at rext `023fd38` the true sites are the ones above. Two of the old anchors landed on a
> bare `fi` — which is the only reason this was caught: `anchor_construct_guard`'s
> *anchor-on-block-terminator* arm fires on a terminator and stays silent on an anchor that merely points
> at the wrong real line. The other eight were **plausible-but-wrong**, the failure mode this file's own
> §"the anchor is the contract" warns about, and no fence saw them.
>
> 🔴 **The `studio-desk` lane's anchors are NOT re-derived here, because the code they described is gone.**
> This paragraph used to continue: *"The one lane that IS strict LIFO is `studio-desk`: apply back · logout ·
> logo · shell · nothirdparty (`:992`, `:997`, `:1017`, `:1032`, `:1035`), revert exactly reversed
> (`:976`)."* At `023fd38` the five `sd_*_manifest` variables are declared at `:1093-1101` and then used at
> **exactly one** site — the patch-set fingerprint at `:1111`. There is no `demopatch apply` for any of them
> anywhere in the file (the complete apply-site list is `841 846 865 877 889 901 915 939 950 963` for
> next-web and `1408 1411 1430 1447 1460 1468 1477 1509 1520` for hiring). So the claim is not merely
> mis-anchored — **whether that lane still applies its patches at all is an open question**, and inventing
> new line numbers for it would have manufactured a citation for behaviour nobody has verified. Left
> stated, not guessed.
The invariant is *"the chain reverts in the right order"*, not *"the list is LIFO"* — and all three lanes
satisfy it. `stack-core/union_apply_guard.py` asserts it for the two lanes that share a clone
(`_FUNCS = {"next-web": …, "hiring": …}`, `union_apply_guard.py:75`); studio-desk builds from its own clone
and is out of that guard's scope.

The `app`
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

**27 patches: 14 × `next-web-app` (4 × `apps/web` + 3 × `apps/hiring` + 3 × `packages/ui` + 3 × `packages/core-js` + 1 × `packages/graphql`) · 3 × `app` · 5 × `ant-academy` · 5 × `studio-desk`.**

> ⚠️ **RE-DERIVED FROM THE DIRECTORY 2026-08-27, and the previous line was stale by FOUR — not by the one
> this pass added.** It read `23 patches: 11 × next-web-app … 2 × app`; the truth on disk was 26 before
> `app-directus-upload-endpoint` and is 27 after. The `next-web-app` sub-breakdown by package was stale in
> three of its five terms. The fence caught the total because `TestPatchInventory` pins its OWN constants
> against this table — and the constants had been carried forward correctly while the prose had not, which
> is the inverse of the M253 failure recorded below. **Both mirrors move in one commit or neither is
> trustworthy.**

> **v2.7 "july jitter" M249 adds FIVE — the cross-app "Back to Cockpit" family + the FIRST-EVER `studio-desk` SOURCE patches:** `next-web-back-to-cockpit` (a `packages/ui` NavbarTop item — SHARED, so it bakes into BOTH the web and hiring images; `packages/ui` goes 2 → 3); the **three** `studio-desk` patches (`studio-desk-back-to-cockpit` + `studio-desk-logout-url` + `studio-desk-logo-url` — a NEW repo in this inventory, image-baked via a net-new `build_frontend_studio_desk` patch ladder + patch-set fingerprint, §5-bis); and `ant-academy-back-to-cockpit` (native-run, `ant-academy` goes 4 → 5). See §"Additive-UI injection" for the pattern the four "Back to Cockpit" items share.

> **v2.7 "july jitter" M253 adds TWO more `studio-desk` patches (3 → 5) — the first-paint pair on `app/core/main.ts`:** `studio-desk-shell-first-paint` (paint the `.page-skeleton` shell BEFORE the boot awaits) **CHAINED** with `studio-desk-no-thirdparty` (no-op `Sentry.init`/`posthog.init`). They ride the SAME `build_frontend_studio_desk` ladder M249 built; the patch-set fingerprint grows 3 → 5 manifests, so a pre-M253 studio image is detected stale + rebuilt. See [`latency-budget.md` §"studio-desk first-paint budget"](latency-budget.md).

> **Inventory reconciled to the `demo-stack/patches/` directory (15 manifests at v2.6 M238; 16 at M244, adding the anon-view `academy-fs-published-public`; 21 at v2.7 M249, adding the 5 cross-app "Back to Cockpit" patches; 23 at v2.7 M253, adding the 2 `studio-desk` first-paint `main.ts` patches).** This table had drifted from the
> `demo-stack/patches/` directory in **two** ways, both fixed here after a directory-vs-table sweep:
> 1. **The 5 `ant-academy` patches are NATIVE-RUN, not `demopatch`-tool patches** — ant-academy runs via `next dev`
>    from its clone (not an image), so each is applied by its **own** `stack-injection/apply-ant-academy-*.sh` /
>    `apply-academy-fs-*.sh` shell helper (apply-before-launch / revert-on-`--stop`), re-implementing the guard
>    ladder against the same canonical manifest (see §4 "Three apply vehicles"). This is why they were historically
>    absent from this inventory (which grew around the image-baked `demopatch` tool) — added the
>    **`academy-fs-published-*`** rows: `-fallback` (the catalog, M230), `-chapter-body` (the body, M238), and
>    `-public` (the anon /library + /free + home view, M244), one
>    FS-as-published behavior gated on `ACADEMY_DEMO_FS_PUBLISHED` (+ `DEMO_NO_ACADEMY_FILL` opt-out); see
>    [`frontend-tier.md`](frontend-tier.md) and [`../../services/ant-academy.md`](../../services/ant-academy.md).
> 2. **The 2 M232 `next-web-interview-flag-*` patches** (`packages/ui`, the interview-report flag gate — the M219
>    aireadiness-flag twin, for the content-stories interview sessions) were never added to the table. Added below.
>    *(**Landed v2.6 M238 harden — the standing hygiene gap is closed:** `demo-stack/tests/test_patch_inventory.py`
>    (`TestPatchInventory`) is the directory-driven fence. It enumerates every `patches/<name>/<name>.yaml`, loads
>    each through `manifest_loader` (valid + `scope=demo` + `id==dirname`), and pins the EXACT total (**27** at 2026-08-27) AND
>    the per-repo breakdown (`14 next-web-app · 3 app · 5 ant-academy · 5 studio-desk`) against this §5 table — so adding, removing,
>    or mis-filing a patch goes RED until BOTH this table and the fence's constants are updated together.)*

> **Mirrored-count discipline (v2.7 release-level note, C1).** v2.7 shipped **three** count-drifts: the
> patch-inventory total shipped RED at HEAD (M253 updated this table's header but not the fence's constants —
> caught at close, FIX-M254-h), the "live Playthroughs" count (16 → 18), and the AI-readiness KPI-tile count
> (4 → 5). **Rule:** any count mirrored in more than one doc, or backed by a test fence, must move with **all**
> its mirrors **and** its fence in the **same** commit — never header-only.

> **studio-desk is a first-class demopatch target now (v2.7 release-level note, C2).** Before v2.7, studio-desk
> carried **no** source patches. It graduated across three milestones: **M249** built the net-new
> `build_frontend_studio_desk` patch ladder + patch-set fingerprint (the first-ever studio-desk source patches —
> the "Back to Cockpit" + prod-eject trio); **M252** wired the studio backend's AI-provider `.env` into the demo
> container (enabling the studio-builder Playthroughs — see [`playthroughs.md`](playthroughs.md)); **M253** added
> the first-paint `main.ts` pair. studio-desk now holds **5** of the 23 patches.

> **The `apps/hiring` patches are M224 "the callback" (v2.4 "casting-call").** The demo now runs the
> **real Hiring app** as a second UI container (TOK-02 — the two-app demo), so a recruiter hero lands on the
> genuine `apps/hiring` candidate-comparison Results screen instead of a re-skinned workforce fake. **The HIRING
> image (`build_frontend_hiring`) bakes SEVEN patches today**, not two — and not the four this paragraph
> claimed until M257x: the **2 net-new** `apps/hiring` patches
> (`next-hiring-role-remap`, `next-hiring-members-pagination`), the **2 chained shared `urls.ts`** patches
> (`next-web-studio-url` → `next-web-public-website-url`), the **2 shared `packages/ui` interview flag-gates**
> (`next-web-interview-flag-container` / `-result`, M232) and the **shared `next-web-back-to-cockpit`** item
> (M249) — measured off the fingerprint call itself, `demo-stack/up-injected.sh:1141-1143` @ rext `8956e69`,
> which passes exactly those seven manifests to `next_web_patchset_fp`. The `urls.ts` pair is
> applied on the hiring build because the Studio nav
> link lives in the **shared `packages/ui` NavBar** (`key: STUDIO_URL`) — so an unpatched hiring image ejects the
> presenter to `studio.anthropos.work` exactly as `apps/web` did. Found + killed at iter-13 (the hiring image's
> client chunks were `docker exec`-grep-verified to carry **0** `studio.anthropos.work`; the trustworthy render
> probe of iter-12 had surfaced the eject the earlier broken probe hid). All seven ride `build_frontend_hiring`'s
> transient apply/revert, fenced by a **7-manifest patch-set fingerprint union** (§5-bis) that forces a
> rebuild if any of the seven moves. **The revert is not LIFO** — `up-injected.sh:1213` reverts
> *interview-result · interview-container · pagination · rolemap · pubweb · studio · back-to-cockpit* against
> an apply order of *studio · pubweb · rolemap · pagination · interview-container · interview-result ·
> back-to-cockpit*; only the `urls.ts` chain is inverted, which is the only ordering that matters (§4's
> correction table).
> **This is the SAME drift §4's M255 box already corrected on its own copy of the number** — it moved
> `4 → 7` there and left this paragraph on 4, which is precisely the mirrored-count failure the C1 note two
> boxes up defines. A correction that reaches one cell is not a correction.
> The 2 net-new `apps/hiring` patches are the **same class as a known `apps/web`
> patch** — the same monorepo (`next-web-app`), the same defect the web app already fixed, never mirrored onto
> hiring. *(**The rest of this paragraph is M224-era bookkeeping — it predates the M232 interview-flag + M238 academy-body additions.** At
> M224 the distinct-manifest total was **11**; the mechanism it records still holds — the chained `urls.ts` pair is
> counted once (under `packages/core-js`) yet applied on **both** frontend builds — but the **current
> directory-fenced total is 23**, per the §5 header above. The pre-M224 line read "8 patches / 5 × next-web-app";
> M224 corrected it to 11 with the `next-web-no-thirdparty` row, and M238 reconciled the whole table to the 15 on
> disk.)*

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
| `next-hiring-role-remap` | `next-web-app` · `apps/hiring/src/context/UserStatusContext.tsx` | **(M224 tik C) the recruiter reaches the hiring enterprise Results routes.** `apps/hiring` stores the Clerk org-role RAW (`role: userRole` = `org:admin`) where `apps/web` **remaps** it (`remapUserRole('org:admin') → 'admin'`). So an admin recruiter reads as **non-admin** in the hiring app, `EnterpriseWrapper` bounces her to the candidate Home, and **0 insights rows** render. The patch adds the same remap (nested, string-literal casts — `apps/hiring` imports `MembershipRoles` **type-only**). **NOT Clerkenstein** (`org:admin` is faithful to real Clerk RBAC), **NOT the seeder** (Rae is already `role='admin'`). Targets its **own** file — no chain. **⚠️ This patch is a WORKAROUND for a live PRODUCTION defect, filed 2026-08-27 as [PD-v29-B](../../../knowledge/plan/platform-defect-register.md) — `origin/main` still stores the role raw at `apps/hiring/src/context/UserStatusContext.tsx:174`, so every Hiring `/enterprise/*` route bounces a real org admin for real customers too. Retire the patch when the platform lands the remap; until then do not read its existence as "demo-only".** |
| `next-web-interview-flag-container` | `next-web-app` · `packages/ui/src/AISimulation/AISimulationResultContainer.tsx` | **(M232)** turns the INTERVIEW report **FETCH** on for a demo — a demo bakes no PostHog, so `posthog.isFeatureEnabled('flag_interview_*_report')` is falsy forever and the two report GraphQL fetches never fire. The M219 aireadiness-flag twin, for content-stories interview sessions. Applied on both frontend builds |
| `next-web-interview-flag-result` | `next-web-app` · `packages/ui/src/AISimulation/AISimulationResult/AISimulationResult.tsx` | **(M232)** turns the INTERVIEW report **RENDER** on — the render gate is a SEPARATE component that independently recomputes the same flag booleans (chained with the FETCH patch above). Same PostHog-unconfigured root cause |
| `next-hiring-members-pagination` | `next-web-app` · `apps/hiring/src/context/InsightsContext.tsx` | **(M224 tik D) the Results dashboard stops hanging on the loading spinner.** The exact **mirror of `next-web-members-pagination`**: `apps/hiring`'s InsightsContext fetches `useGetOrganizationMembers({ limit: 1000 })` — an unbounded whole-org fetch the activity-dashboard layout **blocks** on (`if (loading) return <BaseLoading/>`), and its `GET_MEMBERS` query resolves `targetRole` **per row** — so the per-sim scoreboards never mount. Caps the fetch `1000 → 30`. The **per-member Sentinel authz half of the wall needed NO new patch**: the hiring app hits the **same shared `app` backend** that already bakes `app-targetrole-authz-skip`, so `targetRole`'s per-object RPC is already dropped for this path too. Targets its **own** file — no chain |
| `ant-academy-dev-origins` | `ant-academy` · `code/next.config.js` | admits a `--public-host` demo's MagicDNS origin to `next dev` |
| `academy-fs-published-fallback` | `ant-academy` · `code/src/lib/serverTenant.js` | **(M230, native-run)** the empty demo home GRID renders REAL cards via an FS-as-published catalog fallback (no "Draft" chip), gated on `ACADEMY_DEMO_FS_PUBLISHED`. Applied by `apply-academy-fs-published.sh` |
| `academy-fs-published-public` | `ant-academy` · `code/src/lib/serverTenant.js` | **(M244, native-run)** the ANONYMOUS-view half — /library, /free/*, and the cross-port academy home (:3077) render REAL cards via the same FS-as-published fallback on `getPublicCatalogView` (`getBackendCatalogView(new Set())` — the public/empty eid set, so no tenant content leaks onto an anon route). **CHAINED** on `serverTenant.js` (its `pre_sha256` **is** `-fallback`'s `post_sha256`): applied AFTER `-fallback`, reverted BEFORE it. Same `ACADEMY_DEMO_FS_PUBLISHED` gate. Applied by `apply-academy-fs-published-public.sh` |
| `academy-fs-published-chapter-body` | `ant-academy` · `code/src/lib/serverChapterBody.js` | **(M238, native-run)** the BODY half — clicking "Start the course" renders the FS chapter body (locale-aware, unlocked, un-chipped) instead of the "You wandered off the trail" 404. Same `ACADEMY_DEMO_FS_PUBLISHED` gate. Applied by `apply-academy-fs-published-body.sh` |
| `next-web-back-to-cockpit` | `next-web-app` · `packages/ui/src/NavBar/NavbarTop.tsx` | **(M249)** a fail-closed **"Back to Cockpit"** item in the desktop account dropdown (reads `NEXT_PUBLIC_COCKPIT_URL` = 7700+OFFSET — a DIFFERENT port from the web app; renders only when set). **SHARED `packages/ui`, so it bakes into BOTH the web + hiring images.** The **additive-UI injection** pattern (a NEW menu element, not a URL rewrite — see the section below). Targets its own file — no chain |
> **⚠️ ALL FIVE `studio-desk` PATCHES BELOW ARE RETIRED (2026-08-17) — the Next migration.** They are
> sha-pinned to `app/core/main.ts`, `app/core/scaffold/userProfile.js` and `app/core/scaffold/pageWrapper.js`;
> `app/core/` **does not exist** after the migration, and every one of them reads `import.meta.env`, which
> resolves nowhere in a Next build. They do **not** "drift" — demopatch classifies the target as **absent**
> and **REFUSES at G2**, so running the ladder cost five guaranteed refusals per bring-up and wrote five
> misleading `REFUSED` lines. `build_frontend_studio_desk` no longer applies them; it writes `RETIRED` lines
> instead.
>
> **What replaced them.** The three URL patches are now SOURCE behaviour — `app/_lib/externalUrls.ts` reads
> `NEXT_PUBLIC_* || <prod host>`, fed by the build args the rext-owned Dockerfile declares. That is strictly
> better than re-pinning: the patches only ever ran on a **demo**, so a `dev-N` stack prod-ejected either way.
> The two first-paint patches need no successor — `AppShell.tsx` server-renders the chrome, so there is no
> empty-body blank to reorder and no `.page-skeleton` to inject (that CSS was deleted deliberately).
>
> **Their manifests stay in the build's cache fingerprint on purpose.** Retiring a patch and retiring its
> cache-key entry are two different decisions: drop the paths and the hash changes once, silently, after
> which any image built before the retirement matches forever and keeps serving the Vite bundle. Only the
> first decision has been taken.
>
> **Still missing, and deliberately not faked:** the **"Back to Cockpit" menu ITEM**. That patch was
> *additive UI*, not a URL swap, so nothing in source replaces it. `NEXT_PUBLIC_COCKPIT_URL` is passed to the
> build and is currently read by nothing — the seam is ready for whoever adds the item.

| `studio-desk-back-to-cockpit` | `studio-desk` · `app/core/scaffold/userProfile.js` | **(M249, the FIRST-EVER studio-desk SOURCE patch)** rewrites the user-menu **"Back"** control to THIS stack's app (`import.meta.env.VITE_WEB_APP_URL`, killing the `app.anthropos.work` prod-eject) **and** ADDS a fail-closed **"Back to Cockpit"** sibling (reads `VITE_COCKPIT_URL`). Image-baked via `build_frontend_studio_desk` (net-new patch ladder). **CHAINED** with `studio-desk-logout-url` (same `userProfile.js`; that patch's `pre_sha256` **is** this one's `post_sha256`) — applied FIRST, reverted LAST |
| `studio-desk-logout-url` | `studio-desk` · `app/core/scaffold/userProfile.js` | **(M249)** rewrites `handleLogout()`'s hardcoded `app.anthropos.work/logout` prod-eject to THIS stack's app (`import.meta.env.VITE_WEB_APP_URL || …`). **CHAINED** on `studio-desk-back-to-cockpit` (same file) — reads DRIFTED against a pristine `userProfile.js` BY DESIGN |
| `studio-desk-logo-url` | `studio-desk` · `app/core/scaffold/pageWrapper.js` | **(M249)** rewrites the header **logo** link's hardcoded `app.anthropos.work` prod-eject to THIS stack's app (`import.meta.env.VITE_WEB_APP_URL || …`). Standalone file — no chain |
| `studio-desk-shell-first-paint` | `studio-desk` · `app/core/main.ts` | **(M253, v2.7 "july jitter")** injects the `.page-skeleton` header+sidemenu DOM **synchronously right after `preloadCriticalCSS()` (~L97), BEFORE the boot awaits** (`clerk.load`/`l12nService.init`/`userService.canAccess`), so the shell paints from CSS+DOM with zero network instead of after a ~4.7 s blank. De-dup is automatic (`PageWrapper#init` wipes `document.body.innerHTML` then rebuilds). **CHAINED** with `studio-desk-no-thirdparty` (same `main.ts`; that patch's `pre_sha256` **is** this one's `post_sha256`) — applied FIRST, reverted LAST. Cuts skeleton-visible p95 4669 → 817 ms (demo-2 local). See `latency-budget.md` |
| `studio-desk-no-thirdparty` | `studio-desk` · `app/core/main.ts` | **(M253)** no-ops `Sentry.init` + `posthog.init` on the demo host (no reachable GlitchTip / no PostHog project on a Clerk-free demo; the imports stay referenced by later `captureException`/`identify`). **CHAINED** on `studio-desk-shell-first-paint` (same `main.ts`) — reads DRIFTED against a pristine `main.ts` BY DESIGN |
| `ant-academy-back-to-cockpit` | `ant-academy` · `code/src/components/UserMenu.jsx` | **(M249, native-run)** a fail-closed **"Back to Cockpit"** `<a href>` in the academy user menu (reads `process.env.NEXT_PUBLIC_COCKPIT_URL`, baked by `ant-academy.sh` `write_env_local`). Applied by `apply-ant-academy-back-to-cockpit.sh` (apply-before-launch / revert-on-`--stop`). Targets its own file — no chain |

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
needless rebuild is far cheaper than serving an unpatched demo to a customer — and since v2.8 M257's L1 that
rebuild is **53.31 s on `macmini`**, n=3 p50, where this line used to say "~3 min").

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

> **⚠️ Until v2.8 M258 this section described ONE of the three apply vehicles, and read as if it described
> all of them.** The gate below landed in the canonical `demopatch`/`apply_patch.py` tool at M217. The
> **five shell helpers** in `stack-injection/` (`apply-ant-academy-dev-origins.sh`, the three
> `apply-academy-fs-published*.sh`, `apply-ant-academy-back-to-cockpit.sh`) each carried a **hand-copied
> ladder that hard-refused on any whole-file drift** — the exact behaviour the box below explains why we do
> *not* keep. One guard, two vehicles, opposite semantics, and nothing compared them.
>
> It came due on 2026-08-14: ant-academy bumped, `next.config.js` drifted by unrelated bytes,
> `ant-academy-dev-origins` refused **non-fatally**, and the demo academy shipped **unable to hydrate**
> (see [`tailscale-serve.md`](tailscale-serve.md) § the ant-academy `allowedDevOrigins` patch). The helpers
> now share **one** ladder — `stack-injection/live_patch_ladder.py` — implementing this same gate, with
> `stack-core/tests/test_live_patch_ladder.py` pinning it.
>
> **Two things the shared ladder does that the canonical tool's `_classify` does not, and both were paid
> for in production:**
> - **The marker is never the "am I applied?" probe** — `academy-fs-published-fallback` and
>   `academy-fs-published-public` patch the **same file** and share the `post_marker`
>   `ACADEMY_DEMO_FS_PUBLISHED`, so a marker-first classifier reads the second as corrupt once the first
>   applies and refuses a patch the demo depends on. The **anchor** decides; the marker is only a
>   post-condition on the output.
> - **`patched` is decided by removing the replacement** and asking whether a pristine hunk remains, not by
>   counting the anchor — because for a *prepend-shaped* patch the anchor is a **substring of its own
>   replacement**, and a naive count calls every applied prepend-patch corrupt.
>
> If you add a sixth live-clone helper, **use the shared ladder**; do not copy a ladder again.

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

### 6-bis. The pre-flight nobody could run: *would the whole set still apply, at ref X?* (M257x iter-223)

`demopatch preflight` answers this for **one** manifest against the **checked-out** tree. There was no way
to ask it of the **set**, at an **arbitrary ref** — which is the question you actually have when deciding
whether to advance the clones. Until iter-223 the only way to find out was to run a bring-up and read the
warnings ~11 minutes in, and §5-bis is the record of what that costs: *a silently-refused perf patch
shipped a 76 s members grid for four releases.*

`rext stack-core/patch_anchor_guard.py --clones-root <stack-demo> [--ref <ref>]` censuses every manifest
at a named ref and asserts the two conditions `demopatch`'s own G2 refuses on — **the path exists**, and
**the anchor occurs exactly once**. It runs in seconds, needs no build, and is a member of the guard
family (`--ref` defaults to `HEAD`, so the family asks *"does the patch layer apply to what this stack
would build now"*; pass `--ref origin/main` for the pin-advance question).

**Sha drift is counted and printed and is never a finding**, which is the whole design decision. Since
M217 the gate self-heals: a drifted whole-file sha with the anchor intact `1×` is `pristine`, WARNed and
applied. A fence that reddened on drift would contradict the shipped mechanism and go RED on a set that
works — and a fence that cries wolf gets suppressed.

**The first run, and the number that matters is the second one:**

| measured at `stack-demo`, 2026-08-09 | `HEAD` (the pin) | `origin/main` |
|---|---|---|
| manifests | 23 | 23 |
| path missing | 0 | 0 |
| anchor occurs exactly once | **23 / 23** | **23 / 23** |
| `pre_sha256` no longer matches | 10 | 10 |

The clones were **not** at the same tree in those two columns — iter-222 measured `app` **+28**,
`next-web-app` **+12**, `ant-academy` **+9** behind `origin/main`. So the answer to *"would advancing the
clones break the patch layer?"* is **no**, and that is a real result about a real question.

**And the drift sets are IDENTICAL, member for member.** Those 10 baselines were already stale at the ref
the demo builds today; the platform advancing did not make them so, and advancing further would not make
them worse. At least one is stale **by design** — §6's chain case, where `next-web-public-website-url`'s
`pre_sha256` *is* studio's `post_sha256` and reads DRIFTED against a pristine file on purpose. Which of
the remaining nine are chain cases and which are simply un-repinned is **not** settled here, and the
guard is deliberately silent on it: it reports the count, never a verdict on the cause.

**Read the green over its reach.** This says every patch *applies*. It does not say any patch is still
*right* — an anchor can survive onto a line whose meaning moved — and it says nothing about two patches
interacting on one path, which is `union_apply_guard`'s subject.

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
7. **Nothing to register — R1 is directory-driven** (`demo-stack/ensure-clones.sh`) — see **§2.1**. Since v2.6
   M237 the old hand-maintained `PATCH_MANIFESTS` array is **gone**: R1 iterates `patches/*/*.yaml`, so a new
   `patches/<name>/<name>.yaml` gets unattended recovery **automatically** (the directory *is* the list) — the
   pre-M237 hazard (a run dying between apply and the `RETURN` trap strands the patch applied; every later build
   then *correctly* refuses to re-apply it because G2 finds the anchor gone; it presents as a patch that silently
   stopped working) no longer needs a manual registration step to avoid. Two tests fence it: `TestR1SweepM237`
   pins the R1 glob against the real `patches/` count, and `TestPatchInventory` (v2.6 M238) pins the EXACT
   inventory total + per-repo breakdown against §5.

---

## 8. Additive-UI injection — patching in a NEW UI element (v2.7 M249)

**Every patch above §7 REWRITES a value** — a URL, a flag predicate, a fetch limit. The value is already in the
source; the patch replaces it with an env-gated form. The M249 "Back to Cockpit" family is the **first** class
that **ADDS a new element to the rendered UI** (a menu item that did not exist), and the anchor/replacement/
fail-closed rules bend differently for it. This section is the pattern reference so the next additive patch does
not re-derive it (a genuine blind area before M249).

### The shape of the problem

A demo runs each sub-app (Workforce, Hiring, Studio, Academy) behind the presenter cockpit's *"become any
hero"* launcher, but once inside an app there was **no way back to the cockpit** — the account/user menu offered
only Settings + Log out. The fix injects a **"Back to Cockpit"** item into each app's menu, pointing at the
per-stack cockpit (`…_COCKPIT_URL` = **7700+OFFSET**, a *different* port from the app itself). There is no config
seam for "add a menu item", so it is a demo-patch (§1 ladder exhausted) — but an **additive** one.

### The four rules an additive-UI patch adds to §7

1. **Anchor on the ASSEMBLY point, not a value.** A rewrite anchors on the string it replaces; an additive
   patch anchors on the **list/markup where sibling items are assembled** (the account-menu array, the menu-
   options `innerHTML` block, the JSX above the logout row) and re-emits it with one new sibling spliced in. The
   anchor must still occur **exactly once** (G2) — pick the assembly point that is unique (e.g. the *desktop*
   `!hiddenSidebar` return block, not a `logOut` line that recurs in both the desktop and mobile branches).

2. **Fail-closed is a CONDITIONAL RENDER, not an `env || original`.** A rewrite stays behaviour-identical by
   keeping the original value as the `env || …` fallback. An additive element has no "original" — so it must
   render **only when its env var is set**, collapsing to *nothing* when unset, so an un-baked build (and any
   world where the patch were ever upstreamed) is **byte-identical** to today. Each framework has its idiom:
   - **React + antd (`next-web` NavbarTop):** build the item only when the env is set, and let the existing
     `lodash/_compact([...])` **drop the `null` slot** — `backToCockpitItem ? mapItem(backToCockpitItem, 0) :
     null`. Unset ⇒ `null` ⇒ dropped ⇒ identical array.
   - **Vanilla-JS template string (`studio-desk` userProfile):** a **nested** template that collapses to the
     empty string — `${import.meta.env.VITE_COCKPIT_URL ? \`<button …>…</button>\` : ''}`. Unset ⇒ `''`.
   - **React JSX (`ant-academy` UserMenu):** a ternary to `null` — `{process.env.NEXT_PUBLIC_COCKPIT_URL ? (<a
     …/>) : null}`. Unset ⇒ `null` ⇒ nothing.

3. **No new import, no new i18n key.** The replacement must compile with **only what is already in scope** — a
   demo-patch that adds an `import` line needs a *second* anchor (the import block), and one that adds a
   translation key would touch platform **message JSONs** (a platform edit the whole mechanism forbids). So:
   reuse an in-scope symbol (next-web spreads the in-scope `logOutMenuItem` for a valid `IconDefinition` —
   #M249-D1), use a plain **string literal** label (`'Back to Cockpit'` — a demo affordance, never shipped to
   real users, so an un-i18n'd string is correct here), and a FontAwesome class already loaded by the app.

4. **The env value is baked by the CALLER, offset-templated.** The manifest's `build_env` **documents** the
   line (`…_COCKPIT_URL=$SCHEME://$HOST:$((7700+OFFSET))`), but the bring-up bakes it explicitly into the app's
   env overlay (`up-injected.sh` → `apps/web`/`apps/hiring/.env.local`; the `.env.production.local` overlay for
   `studio-desk`'s `VITE_COCKPIT_URL`, since it is **not** a declared Dockerfile ARG; `ant-academy.sh`
   `write_env_local` → `code/.env.local`). The item is **inert without the bake, and the bake is inert without
   the item** — the same two-part contract as `next-web-no-thirdparty`.

### The rewrite half rides along (studio-desk)

The `studio-desk` lane is a hybrid: `studio-desk-back-to-cockpit` is *additive* (rules 1–4), but it **also**
rewrites the existing "Back" control's `app.anthropos.work` prod-eject to `import.meta.env.VITE_WEB_APP_URL || …`
(a §7-style rewrite), and its siblings `studio-desk-logout-url` / `studio-desk-logo-url` are **pure** rewrites.
Reading `import.meta.env.VITE_WEB_APP_URL` directly (not `config.WEBAPP_URL`) is deliberate (#M249-D2) — it
keeps the **original `app.anthropos.work` fallback** (so it is behaviour-identical when unset, per §7-4) and
needs no `config` import (rule 3), while reading the *same* env var `config.WEBAPP_URL` reads.
