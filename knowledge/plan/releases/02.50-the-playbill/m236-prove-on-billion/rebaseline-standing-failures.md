# Re-baseline of the standing demo-stack test failures

**Status:** measured, dated, and handed to `/developer-kit:close-release` for a fate decision.
**Measured:** 2026-07-20, at the M236 close continuation, under an explicit user decision (see
`decisions.md` → CLOSE-D2 → **RESOLUTION**).
**Fate:** **NOT decided here.** The user's instruction was *re-baseline now, decide at release close*. This
document is the input to that decision — deliberately a characterisation, not a fix.

> **Read this before acting on the old label.** The carried item said *"14 pre-existing demo-stack test
> failures, 6 of them `pre_sha256` pin drift."* The count was right. **The diagnosis was wrong**, and the
> remedy it implied — re-anchor the drifted pins — would have caused real damage. See §3.

---

## 1. Headline

| | count |
|---|---|
| Carried label (10 milestones, 2 releases) | 14 |
| Reproduced before any change | **14** — the label's count is accurate |
| After a clean, stable-`main` clone set | **8** |
| Of those 8: real product/tooling defects | **0** |
| Of those 8: stale tests asserting deliberately-changed behaviour | **7** |
| Of those 8: environment-conditional test with a missing skip guard | **1** |
| Of the original 14: `pre_sha256` pin drift | **0** |

**Six of the fourteen were not test failures at all** — they were a dirty clone reporting itself as a test
failure. **The remaining eight are all test-side debt.** Nothing in the set indicts shipped behaviour, and
nothing in it is a v2.5 regression.

## 2. How it was measured

Reproduced with the documented invocation, `cd rosetta-extensions/demo-stack && pytest tests/`, on macOS
(Darwin 25.1.0), against the local `stack-demo/` clone set — which is what these tests read (they resolve
`<rosetta>/stack-demo/next-web-app/<path>`; they do **not** require a running stack).

Per the user's condition, every platform repo was first advanced to its **stable `main`**. No repo exposes a
distinct "stable" ref, so `main` was used directly, as instructed.

**Resolved refs (both boxes; reproducible baseline):**

| repo | ref | sha | date |
|---|---|---|---|
| `app` | `main` | `aa2574541` | 2026-07-20 |
| `cms` | `main` | `93e6aa354` | 2026-07-17 |
| `graphql-wundergraph` | `main` | `5d9c7568e` | 2026-07-17 |
| `jobsimulation` | `main` | `5d3003f9f` | 2026-07-06 |
| `messenger` | `main` | `d41029217` | 2026-07-20 |
| `next-web-app` | `main` | `61d72e24d` | 2026-07-20 |
| `platform` | `main` | `0808b9224` | 2026-07-07 |
| `roadrunner` | `main` | `87d8d4438` | 2026-06-19 |
| `sentinel` | `main` | `88bc55929` | 2026-06-19 |
| `skillpath` | `main` | `2d0fc2d27` | 2026-06-19 |
| `storage` | `main` | `769660542` | 2026-06-19 |
| `studio-desk` | `main` | `f6320f865` | 2026-07-03 |
| `ant-academy` | `main` | `a43420bdd` local / `495f3c243` on `billion` | 2026-07-17 / 07-11 |

> **One disclosed exception.** `ant-academy` on `billion` could **not** be fast-forwarded: its working tree
> carries `code/public/catalog.json` and `code/public/content/index.md`, both **tooling-generated** by the
> academy-fill step and regenerated on every bring-up. Advancing it would have required discarding
> working-tree files, which is outside what this session is permitted to do. It is a native, non-fatal
> peripheral and 5 commits behind. Disclosed rather than quietly rounded to "all repos at main".

## 3. The pin-drift diagnosis is refuted

The carried item asserted 6 failures were `pre_sha256` **pin drift** — i.e. the demopatch manifests had gone
stale against a moved platform source, and needed re-anchoring.

Every manifest `pre_sha256` was compared against three readings of its target file: the **worktree**, the
**stale `HEAD`**, and **`origin/main`** — the last two read as blobs via `git show`, so nothing was mutated to
get this table.

| manifest | pre matches `HEAD`? | pre matches `origin/main`? | pre matches worktree? |
|---|---|---|---|
| `next-web-ssr-graphql-origin` | ✅ | ✅ | ❌ |
| `next-web-studio-url` | ✅ | ✅ | ❌ |
| `next-web-aireadiness-flag-gate` | ✅ | ✅ | ❌ |
| `next-web-no-thirdparty` | ✅ | ✅ | ❌ |
| `next-web-members-pagination` | ✅ | ✅ | ❌ |
| `next-web-interview-flag-container` | ✅ | ✅ | ✅ |
| `next-web-interview-flag-result` | ✅ | ✅ | ✅ |
| `next-web-public-website-url` | *by design* | *by design* | *by design* |

**Every pin is correct at both refs.** The targeted files did not move at all across the 202-commit gap. The
only disagreement is with the **worktree**, which was carrying **leftover applied patches** from a build that
never completed its revert (root cause: `F-M236-CLOSE-2` — R1's pristine sweep covers 3 manifests of ~15).

`next-web-public-website-url` matches nothing and is **correct**: its `pre_sha256` *is*
`next-web-studio-url`'s `post_sha256` — the documented **chain rule** — so it reads "DRIFTED" against a
pristine file by design. `demopatch-spec.md` states this; a re-anchoring pass would have "fixed" it into
being genuinely broken.

**Why this matters more than the count.** The remedy implied by the label was *re-anchor the drifted pins*.
Executed against these clones, that would have re-pinned five manifests to **patched** content — baking a
demo patch into the manifest's definition of pristine, permanently, and silently disarming the drift
detector for those files. **The deferral's own proposed fix was the dangerous action.** Reverting the clone —
the actual fix — took one command and changed no checked-in file.

Confirmation: after `demopatch revert --force-pristine` on the six affected manifests, all 6 clone-dependent
failures pass. Independently, the live cold bring-up on `billion` at stable `main` **applied every patch
cleanly**, which only happens if the pins match `main`.

## 4. The 8 that remain

All are test-side debt. None indicts shipped behaviour.

### 4a. Stale tests — deliberately-changed behaviour (7)

**`test_cockpit.py` — the academy link (4)**
`TestAcademyLink::test_academy_link_renders_per_hero_when_base_set`,
`TestAcademyCatalogEntryEdges::test_render_academy_entry_fields_are_escaped`,
`TestAcademyCatalogEntryEdges::test_render_defaults_academy_path_persona_label_when_absent`,
`TestServedPanelWithAcademy::test_root_serves_academy_link`

These assert the **M53 F6 per-hero `[Academy]` link** — one `class="btn academy"` per hero card. The hero
card now renders exactly **one** unified CTA (`class="btn login"`, *"Log in as"*), which is the documented
M43 design (*"the one unified [Log in as] CTA per hero … no more separate [Jump]"*).

The marker `class="btn academy"` **still exists in `cockpit.py`** — but it now belongs to the **M234
Content-stories tab**, emitted per *session* under `product.app_base == "academy"`, not per hero. The string
was **re-used for a different feature**, so the tests fail on a count mismatch (`0 != 2`) rather than on a
missing symbol, and a grep for the marker makes the feature look present. That is why this set reads as
mysterious and has survived ten milestones of triage.

**`test_cockpit.py::TestOverlayJs` (2)**
`test_inflight_window_is_30s` asserts the literal `30000` is present. The 30-second in-flight window was
**deliberately removed** as a bug fix, and the code carries a paragraph explaining why: at the pre-M218
38–40 s login the window made sense; at the measured 2.4 s login it re-showed a spinner over an
already-loaded workspace on every back/reload. The test asserts the bug.

`test_localstorage_access_is_guarded` asserts `js.count("try {") >= 3`; there are now 2, because one
`try` block was removed with the 30 s re-show. **Both** `localStorage` accesses (`setItem`, `removeItem`)
remain correctly guarded — the property under test holds; the *assertion* was written as a count and is
over-specified.

**`test_host_prereqs_m215.py::TestF12ServeResetGenerator::test_public_host_emits_per_port_off_for_all_browser_ports` (1)**
Expects `[13000, 13077, 15050, 17700, 18082, 19000]`; the generator emits that plus **`13001`** — the
hiring-app port (`3001 + 1×10000`). A port was legitimately added to the browser set; the expected list was
never updated. The generator is right, the fixture is stale.

### 4b. Environment-conditional, missing skip guard (1)

**`test_purge.py::TestPurgeDataDir::test_purges_container_owned_0700_data_THE_BUG`**
Self-declaring: the assertion message reads *"expected on Docker Desktop/macOS; this test is meaningful on a
Linux host"*. It stages a container-owned `0700` dir and asserts the host user **cannot** remove it; on
Docker Desktop/macOS the host user **can**, so the precondition fails and the test **fails instead of
skipping**.

This one is **host-dependent, and therefore so is the failure count.** On this macOS box the standing set is
8; the same commit on a Linux host is expected to be **7**. The behaviour it guards was observed working on
`billion` during this session — the teardown logged *"purging container-owned data (UID 1001 / 0700 — the
host user cannot unlink it)"* followed by *"data purged"*.

> **Consequence for the ledger:** *"N failures"* is not a well-defined number for this suite without naming
> the host OS. Any future carry must state the platform, or it will drift again for the same reason it
> drifted the first time.

## 5. Recommendation to `/developer-kit:close-release`

Offered as input; the fate is the release close's to take.

1. **Fate 1 (LAND) is now cheap and is the recommended default.** All 8 are test-side edits with no product
   risk: delete or invert 2 overlay assertions, re-point 4 academy assertions at the M234 semantics (or
   retire them with the M53 feature they describe), add `13001` to one expected list, and convert the purge
   test's precondition into a `skipUnless(Linux)`. None requires a live stack or a platform edit.
2. **Fix `F-M236-CLOSE-2` first** — extend R1's sweep from 3 manifests to the full `patches/` set. Otherwise
   the clone-dependent failures return the next time a build is interrupted, and the next reader re-derives
   this document.
3. **Do not re-anchor any `pre_sha256`.** §3 shows why. If a pin ever *does* read drifted, first confirm the
   clone is pristine.
4. **Consider `F-M236-CLOSE-1` separately and at higher priority** than the 14. That a demo rebuilds images
   from never-updated source is a standing correctness problem for every demo the project ships, and it is
   the upstream generator of this whole class. It is not part of the 14 and should not inherit their fate.

## 6. Provenance

- Before-reading: `14 failed, 716 passed, 9 subtests passed in 244.41s`
- After clean stable-`main` clones: `8 failed, 722 passed, 9 subtests passed in 230.99s`
- Host: macOS (Darwin 25.1.0); suite is host-sensitive — see §4b
- Related findings: `decisions.md` → `F-M236-CLOSE-1` (never-updated clones),
  `F-M236-CLOSE-2` (incomplete pristine sweep)
