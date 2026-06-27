# Hardening Ledger — M42m Manager-coverage

> The M42m CODE lives in the SEPARATE rext authoring repo (`.agentspace/rosetta-extensions` @ tag
> `method-acting-m42m-iter05`); this corpus ledger records the harden passes. Test commits + the final
> tag land in rext; this file is the planning-side record close-milestone greps for the final-mode entry.

## Pass 1 — 2026-06-26 — final

**Iters hardened this pass:** all milestone-touched code (final-mode cumulative scope: iters 03-05 —
demopatch + manifest_loader + the next-web-studio-url manifest; feedback.go org-feedback mirror;
stack-verify/e2e manager namespace; up-injected.sh + ensure-clones.sh).

**Tiks covered since prior pass:** all iters in the milestone (first + only harden pass — the ledger
did not exist before this entry).

**Mode:** final (gate-met inference + explicit `--final`). The latest iter (iter-04) `progress.md`
carries `**Gate:** MET`, reproduced on a fresh zero-manual demo-up (manager 70(0,0,0,0) EXHAUSTED
gateMet:true + employee 59(0,0,0,0) EXHAUSTED gateMet:true, no regression); iter-05 (R1b sweep) is the
closing iter. This is the cumulative-scope sweep before `/developer-kit:close-milestone`.

**Coverage delta on touched files** (measured as a finder, not a goal — adversarial/edge/error grids the
per-iter per-symptom tests didn't sweep):
- `demo-stack/patches/demopatch` (THE highest-value target — a tool that patches platform source):
  guard test count 18 → 43 (+25). Every adversarial REFUSE edge now pinned; ZERO source change (the six
  guards held against the full adversarial grid).
- `demo-stack/patches/manifest_loader.py` (the patch-SPEC parser): the strict-subset parser edges now
  fenced (path/abs/marker/uppercase-sha rejects; block-scalar internal-blank + `|-`/`|+` variants;
  indented-line / non-key-line / block-key-without-pipe rejects; comment-skip; quoted-scalar strip).
- `stack-seeding/seeders/feedback.go` (the M42m org-feedback mirror fix): feedback test count 5 → 8 (+3)
  — the second-COPY (mirror) error arm + empty-population no-op + the 1:1 mirror invariant.
- `stack-verify/e2e/lib/coverage-manifest.ts` + `tests/coverage.spec.ts` (the manager namespace): a new
  PURE-LOGIC unit spec (17 tests) — previously the manifest decision logic was only exercised by the live
  Playwright sweep (needs a demo up); now pinned in CI with no stack.

**Tests added:**
- demopatch → `demo-stack/tests/test_demopatch.py`: +25 (TestAdversarialGuards: G1 dir-symlink-escape /
  `..`-traversal / repo→sibling-stack(../stack-dev) refuses; G2 0-anchor refuse / partial-apply
  post-hash-without-marker → drifted-not-patched / replacement-disagrees-with-post_sha256 refuse;
  G5 drifted-refuse-without-force / pristine-idempotent-noop / force-pristine-restores-drift; G3
  stage→revert cycle; absent-target / status-absent / missing-manifest exit-2 / unknown-verb. Plus
  TestManifestLoader: +11 parser-edge fences).
- feedback → `stack-seeding/seeders/feedback_test.go`: +3 (MirrorCopyErrorPropagates, EmptyPopulationNoOp,
  MirrorMatchesFeedbackOneToOne).
- TS harness → `stack-verify/e2e/tests/coverage-manifest.unit.spec.ts`: +17 (manifestFor selection,
  pageDescriptorFor normalization + exact match, per-page/per-section structural integrity, unique paths,
  all-manager-pages-calibrated, org-feedback ≥2 floor, manager fan-out sample-rule regexes match the
  /user + activity-drilldown explosion and NEVER a dashboard page).

**Bugs surfaced + fixed inline:**
- One TEST-author slip caught + fixed inline: the new `MirrorMatchesFeedbackOneToOne` used the wrong
  mirror column indices (status/completion are localSessionCols 5/6, not 6/7) → fixed in the same commit
  (791376a). NO production-code bug surfaced — the demopatch guards and the feedback mirror both held
  against the full adversarial/edge grid (zero source change across the whole harden session).

**Flakes stabilized:** none (no flakes). Flake gate: 3 consecutive clean runs each of the demopatch new
classes, the feedback new tests (`-count=1`), and the TS unit spec.

**Audits (all GREEN):**
- G3 never-mutate-git source grep — explicit: no `git add/commit/push/tag/reset/stash/rebase/merge/...`
  verb in `demopatch` source (the single sanctioned `git checkout -- <path>` is whitelisted to one fn).
- Supply-chain — zero `go.mod`/`go.sum`/`package.json`/lockfile/`requirements.txt` change in the entire
  M42m footprint (`53574ae..HEAD`); demopatch + manifest_loader import stdlib + two LOCAL rext modules
  only (no PyPI dep).
- Closure — `datadna measure-closure` is a LIVE-stack operator (`--stack demo-N --dsn base`); its CODE is
  an unchanged M34 gene and the feedback seeder's refs use the closure-safe `linkedRef`/`resolveContentRefs`
  path. The tool compiles + `cmd/datadna` unit tests pass; closure was verified PASS on the live fresh
  demo-up at gate acceptance (not re-runnable in this stack-less harden context).
- Alignment — N/A: zero `clerkenstein/` + zero `alignment/` change in the M42m footprint (confirmed).
- Zero CANONICAL platform-repo edits — the harden session is purely additive test code (3 test files,
  +496 lines, zero source/canonical edit); the v1.10 hard line holds.

**Full-suite verification (all GREEN):** Go seeders `go test -race ./seeders/` ok; Python
`unittest tests.test_demopatch` 43 ok; TS `playwright test --list` 24 specs / 7 files transpile + the unit
spec 17 pass; shell `bash -n` + `shellcheck -S warning` clean on up-injected.sh + ensure-clones.sh;
`gofmt -l` clean + `go vet ./seeders/` clean.

**Knowledge backfill:** none required — the harden surfaced no new edge-case/error-path semantics or
precision boundaries not already documented in the milestone's specs (stories-spec.md / frontend-tier.md /
the demopatch README) + the demopatch source's own inline guard docs. The guard semantics the tests pin
are already authoritatively described in the demopatch module docstring (the six-guard contract).

**Stop condition:** stabilized — coverage delta this final pass is a one-shot cumulative sweep (no prior
pass to delta against; the dimension scan found no further untested branch after the demopatch/manifest/
feedback/TS-harness grids), all suites + the three flake gates GREEN, zero production bug surfaced, every
audit GREEN. This satisfies `/developer-kit:close-milestone`'s iterative-milestone final-harden gate.
