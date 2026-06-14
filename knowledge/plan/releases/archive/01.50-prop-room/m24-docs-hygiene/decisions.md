# M24 — Decisions

_Implementation decisions with rationale, numbered `M24-D1`, `M24-D2`, … . Empty at scaffold; filled during build._

## M24-D1 — the Go toolchain pin is an explicit `toolchain` directive in each go.mod (§4)

NEW-5 asked to bump the toolchain `go1.25.3 → go1.25.11+`, lazily (no dedicated rebuild). The 4 modules had a
**language** `go` directive but **no `toolchain`** directive, so nothing pinned the *toolchain* (the language
version floats the minimum, not the build toolchain). Added an explicit `toolchain go1.25.11` to all 4 go.mod
(alignment, clerkenstein, stack-snapshot, stack-seeding) and tightened the one CI workflow (`clerkenstein
alignment.yml`) from a floating `"1.25"` to `"1.25.11"`. The pin takes effect on the next natural `go build` (Go
auto-downloads go1.25.11) — no rebuild session run here, per the user. Parse-verified with `GOTOOLCHAIN=local go
mod edit -print` so no download was triggered while bumping.

## M24-D2 — the zero-critical-genes guard is defence-in-depth (Validate + GateMet), not one site (§6)

The bug (NEW-11): `compare.pct(0,0)` returns `100.0`, so a DNA with **zero critical genes** scored a perfect
`critical %` and cleared any `--gate-critical` for free; `dna.Validate` never caught it. Rather than only patch
`pct` (which would silently turn a vacuous 100 into a 0 and still let a thresholdless run pass), the fix guards at
both honest layers: (1) **`dna.Validate` rejects** a capabilities-bearing DNA with no critical gene — the
authoritative load/lint-time gate (a DNA that can't be critically gated is malformed); (2) **`GateMet` refuses** a
non-zero critical threshold when the report's new `CriticalGenes` count is 0 — the scoring-time defence for any
Report built around the hole. The report carries `CriticalGenes` so a 100 with zero critical genes is
self-evidently vacuous. All 5 shipped alignment DNAs have a critical gene, so the live gate is unaffected.

## M24-D3 — the `/project-stats` scope fix lands in the developer-kit `stats.sh` (where the script actually is), gitignore-aware (§7)

The milestone's "Repo split" assumed the stats fix was a `rosetta-extensions` item, but investigation found the
`/project-stats` skill is the shared **developer-kit plugin** `stats.sh` (there is no stats tooling in
rosetta-extensions). Its `detect_src_roots()` falls back to `.` for a doc-corpus repo like rosetta, and its
`PRUNE_PATHS` excluded `node_modules`/`.git`/… but **not** the gitignored `stack-*/` per-stack workspaces (cloned
platform repos), so the code-size scan counted ~2M foreign lines (measured: 9,235 code files → the `stack-*/`
clones). Fixed at the real source — the developer-kit plugin (not a platform repo, not the rosetta corpus): added
`*/stack-*/*` to `PRUNE_PATHS` (the named class) **and** a general `drop_gitignored` filter on `list_code_files`
(any gitignored clone is dropped — the correct general behaviour, robust to future workspaces; degrades to a
pass-through when git is absent). Verified: rosetta's inflated code count collapses to the gitignore-respecting
truth.

**Surfaced (out of §7 scope, recorded not dropped):** `stats.sh`'s `count_files`/`count_lines` (the
Knowledge/Plan/Journal doc-counter) report 0 for rosetta because rosetta's docs live under `corpus/` (the counter
only looks at `knowledge/`), and the `eval find … $PRUNE` is fragile under some shells. This is a **pre-existing**
developer-kit doc-counter limitation, independent of the `stack-*/` scanning bug §7 names, and not a v1.5
deliverable — noted here so it isn't silently lost.

## Adversarial review (close-milestone Phase 2c)

**Scenario — incidental-mention false-negative (the README-index guard, §5/b).** The guard's contract is "a
doc's filename appears (token-bounded) anywhere in its directory's README". Adversarial question: does an
*incidental* mention — the filename inside an unrelated URL, a code fence, or prose that isn't an index row —
make the guard treat an otherwise-unindexed doc as "referenced", a false negative? **Probed live:** a README that
mentions `config.md` only inside `https://example.com/docs/config.md` yields exit 0 for a sibling `config.md`.

**Response — by design, no fix.** This is the guard's deliberate lenient-by-filename scope (documented in its
module docstring, §"SCOPE"): it catches the *recurring real miss* (a doc with NO mention at all — the failure
class that bit v1.3/v1.3b/v1.5) without policing *where* the mention sits, which would make it brittle across the
corpus's heterogeneous README styles (tables / bullet lists / prose). Tightening to "must be an index row" would
trade a rare, low-harm false-negative for frequent false-positives and editorial coupling. The complementary
risk — a LONGER filename's tail masquerading as a shorter one (`setup.md` inside `dev-setup.md`) — IS guarded
(the token-bounded match, the real bug fixed in harden Pass 1, ext `191d650`). Recorded so a future reviewer sees
the incidental-mention case was considered, not missed.
