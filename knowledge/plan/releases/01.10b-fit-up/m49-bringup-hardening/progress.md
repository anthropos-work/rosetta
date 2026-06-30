# M49 — progress

## Section checklist
_Checked off as each In-scope deliverable lands. Close when all boxes are ticked._
_Each section = one rext code fix (committed in `.agentspace/rosetta-extensions` @ `main`, rolled into tag `fit-up-m49`) + its corpus doc truth-up (committed in `rosetta` @ `m49/bringup-hardening`)._

- [x] **§1 rext-tag-sot** (#1) — `.agentspace/rext.tag` SoT + `lib/rext_tag.sh` reader + non-fatal pin guard in `ensure-clones.sh`; reconciled **4** conflicting prose pins (the overview's 3 + a stray in `verification.md`) to read from it. Pin = `fit-up-m49`. Docs: `rosetta_demo.md`, `demo-up/SKILL.md`, `frontend-tier.md`, `verification.md`. 96/96 tests green.
- [x] **§2 env-guard-order** (#3) — moved the `.env`-presence guard below the M30 provision block (now checks `BASE_ENV`); a stack-demo-only box provisions-then-checks instead of aborting. Regression test added. Doc: `rosetta_demo.md` bring-up-order step 2. 149/149 tests green.
- [x] **§3 invitation-hmac-secret** (#4) — added `INVITATION_HMAC_SECRET` as a critical/required platform/.env gene + `DemoGeneratedKeys` demo overlay (so demo pre-flight treats it satisfied) + up-injected auto-gen (openssl, values-blind, idempotent, non-fatal). Doc: `secrets-spec.md` (split now 40/8/8 + 13-crit, 56-gene; AI-keys policy noted as M50). 6 count-literals reconciled corpus-wide. Go + 98 demo-stack tests green.
- [x] **§4 ant-academy-clone** (#5) — ensure-clones.sh phase (d2) clones ant-academy EXPLICITLY (non-fatal), NOT via repos.yml (ephemeral platform clone = non-durable + platform edit). Static + 4 functional tests. Docs: `ant-academy.md` (3 spots), `CLAUDE.md` corrected (the M48 "M49 adds repos.yml entry" prediction → the actual explicit-clone fix). 103 demo-stack tests green.
- [x] **§5 disk-preflight-down-cleanup** (#6) — `preflight_disk_headroom()` (non-fatal, mirrors the RAM check, offers prune, `DEMO_DISK_MIN_GIB`/`DEMO_DISK_AVAIL_KB`) + `cmd_down --purge` removes the stack's `demo-N-*` images (scoped, non-fatal; plain `down` keeps them). 5 functional + 2 static tests. Docs: `frontend-tier.md` + `demo-up/SKILL.md`. 162 tests green.
- [x] **§6 nonfatal-frontend** (#7) — before `compose up`, scale any absent-image frontend to 0 (`--scale svc=0`, set-u-safe) so a failed UI build no longer aborts backend + set-dress + verify + cockpit. Static pins (image check, scale-to-0, NO_UI gate, ordering). Doc: `frontend-tier.md` (the "non-fatal" claim is now true). 163 tests green.
- [x] **§7 demopatch-reanchor** (#8) — re-anchored `next-web-studio-url` `pre_sha256` (b3d62db→e961aeae) + `post_sha256` (9f27e25→be0c803a) to the current v2.89.0 source (the hunk is byte-identical; only file-level hashes moved). Computed via the tool's manifest_loader; VERIFIED apply→revert end-to-end on a throwaway copy. Doc: `frontend-tier.md`. 46 demopatch tests green.

### Close gate
- [x] All 7 rext fixes committed in the authoring copy + **tagged `fit-up-m49`** (annotated, on rext HEAD `1035efd`); the on-box `.agentspace/rext.tag` pins `fit-up-m49`.
- [x] Both working trees clean (verified at build-end).
- [x] Static/code-level verification only (no live `/demo-up` — that's the orchestrator's live-verify gate + M53). bash -n + shellcheck -x clean; secret-DNA JSON valid; Go (secretdna+stack-secrets) + 163 demo-stack + 46 demopatch tests green; demopatch re-anchor verified apply→revert.

## M49: Hardening

### Scope manifest (Phase 1)
The M49 testable surface lives in the rext authoring copy (`.agentspace/rosetta-extensions`). Per the orchestration
constraints, the harden deepening focuses on the **Go/testable units**; the bash bring-up scripts are already
covered by `bash -n` + `shellcheck -x` + the demo-stack Python suites + the from-cold live-verify gate (PASSED).

| Module / file | Stack | Existing tests | Coverage at harden-start |
|---|---|---|---|
| `stack-secrets/secretdna/demo.go` (#4 `DemoGeneratedKeys` overlay) | Go | `demo_test.go`, `demo_harden_test.go` (M28) | 99.0% pkg; `Shape` 88.9% (one M49 branch uncovered) |
| `stack-secrets/secretdna/secret-dna.json` (#4 critical gene) | Go (JSON) | `secret_dna_json_test.go` (gene asserts ✓) | n/a — fully asserted |
| `demo-stack/lib/rext_tag.sh` (#1 reader) | bash | `test_tooling.py` | live-verify + shellcheck + py |
| `demo-stack/{up-injected,ensure-clones,rosetta-demo,ant-academy}.sh` (#3/#5/#6/#7) | bash | `test_tooling.py`(163), `test_frontend_build.py`, `test_ant_academy.py`(17) | live-verify + shellcheck + py |
| `demo-stack/patches/next-web-studio-url/*.yaml` (#8) | manifest | `test_demopatch.py` (46) | live-verify apply→revert + py |

**Gap surfaced:** `demo.go:105` — the `IsDemoGenerated(key) → ShapeOpaque` branch of `mintedSource.Shape`.
`Shape` is only reached via `FormatMatch`, which only runs for a gene naming a `format:*` operator; the
`INVITATION_HMAC_SECRET` gene names only key-present + nonempty, so `MeasureForStack` never calls `Shape` on it.
The branch is defensive (a future `format:*` on a demo-generated gene must report opaque, not falsely pass `pk`/`sk`).
Fate-1 LAND NOW: a direct `mintedSource.Shape` unit test, mirroring `TestMintedShape_OpaqueKeys` for the minted side.

### Pass 1 — 2026-06-30
**Coverage delta (milestone-touched files):**
- `secretdna` package statements: 99.0% -> 99.3% (residual 0.7% = `catalog.go`/`dna.go`/`introspect.go` — non-M49, out of scope)
- `demo.go` (the #4 overlay, M49-touched): `Shape` 88.9% -> 100.0%; every other func already 100%

**Tests added:**
- `demo_harden_test.go` (rext): 1 unit — `TestMintedShape_DemoGeneratedKeyIsOpaque` pins the M49 `IsDemoGenerated -> ShapeOpaque` branch (a direct `mintedSource.Shape` probe; proves a future `format:*` on a demo-generated gene would correctly FAIL, not falsely pass `pk`/`sk`). Negative-control verified (fails on branch removal).
- `test_tooling.py` (rext, `RextTagSoT`): 3 — `test_crlf_line_ending_does_not_leak_a_trailing_cr` (the bug regression, below), `test_comment_or_blank_only_file_returns_empty`, `test_hash_glued_to_token_strips_from_the_hash`.

**Bugs fixed inline:**
- **#1 rext_tag.sh CRLF carriage-return leak** (commit `51cc701`, rext): the reader's `awk` captured a trailing `\r` into the picked token, so a CRLF-edited `.agentspace/rext.tag` (a Windows editor / any `\r\n`-writing tool) yielded `fit-up-m49\r` instead of `fit-up-m49`. The leaked CR would fail `git checkout <ref>` in `ensure-clones.sh` with a baffling "pathspec did not match" (the ref carries an invisible CR). Fixed with `gsub(/\r/,"",$1)` (awk's `\r` is portable across GNU/BSD, unlike a `sed s/\r//` some BSD seds skip). Regression test fails on the pre-fix reader (verified via negative-control).

**Flakes stabilized:** none observed.

**Knowledge backfill:** `corpus/ops/rosetta_demo.md` — noted the `.agentspace/rext.tag` reader is CRLF-tolerant (strips a trailing CR so a Windows-edited pin resolves as a clean git ref). Keeps the documented "bare one-line tag string" promise honest (commit `197d6ee`, rosetta).

### Pass 2 — 2026-06-30
**Coverage delta (milestone-touched files):** Go M49 surface (`demo.go`) holds at 100%; no Go delta. Python +1 test on the #8 surface.

**Tests added:**
- `test_demopatch.py` (rext, `TestRealManifest`): 1 — `test_real_manifest_hashes_reproduce_against_live_clone`. `TestRealManifest`'s docstring claimed the shipped `next-web-studio-url` manifest "reproduces its pinned post_sha256" but the test asserted only structural facts, never the hash bytes. The new regression makes the promise real: when the local `stack-demo/next-web-app` clone is present it proves full internal consistency against the REAL pristine source — `pre_sha256` matches the live v2.89.0 source (the #8 re-anchor is correct), the anchor occurs exactly once, the single swap reproduces `post_sha256` (apply's pre-write postcondition), and the revert reproduces `pre_sha256` (G5 self-revert). Catches a FUTURE re-anchor that mis-computes `post_sha256` statically, without a full bring-up; skips cleanly when the clone is absent (CI). Negative-control verified (fails on a corrupted `post_sha256`).

**Bugs fixed inline:** none (pure deepening).

**Flakes stabilized:** none.

**Knowledge backfill:** none KB-worthy in this pass (the #8 hash-reproduction fact is a test-internal invariant, already prose-described in `frontend-tier.md`/`rosetta_demo.md` from the build phase).

### Pass 3 (stabilization re-scan) — 2026-06-30
Full six-dimension re-scan of the remaining M49 surface (#3 `.env`-guard order, #5 ant-academy clone, #6 disk pre-flight + `down --purge` image cleanup, #7 non-fatal frontend). All are well-covered by the existing Python static + functional fences (the #6 disk-math non-numeric edge, the `down --purge` per-stack image scoping, the exact `-ge` floor branch) plus the from-cold live-verify gate (PASSED). The remaining candidate boundaries (exact-at-floor) would be shallow tests exercising an already-covered code path — not added, per the no-shallow-tests rule. No new worthwhile gap surfaced.

### Stop condition
Loop terminated after Pass 2's deepening + a clean Pass 3 re-scan: the Step 2b scan found nothing new worth adding, the M49 Go surface is at 100% (residual package % is non-M49 code), and the flake gate is clean (3 consecutive sequential runs of all 5 new tests, zero flakes). Final tally: 5 new tests (2 unit Go + JSON-structure already complete / 3 bash-functional regression+edge / 1 manifest-hash regression), 1 real bug fixed inline (the rext_tag CRLF leak) with a negative-control-verified regression, 1 knowledge backfill. Both repos clean. The bash bring-up scripts (#3/#5/#6/#7) stay covered by shellcheck + the Python suites + the live-verify gate per the milestone's static-only harden scope; `/developer-kit:close-milestone` Phase 4 runs as defense-in-depth.

