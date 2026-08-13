# iter-09 — decisions

## D1 — next-web's multi-stage build lands as an **rext-owned Dockerfile**, not a demopatch

**The choice.** L1 needs a multi-stage build for both Next apps. `hiring` already builds from an rext-owned
`hiring.Dockerfile` (M224), so it was edited in place. `next-web` built from the platform clone's own
`Dockerfile.dev`, and making *that* multi-stage would be a platform-repo edit — forbidden outright by this
milestone's hard constraints. Two sanctioned alternatives exist and `overview.md` names both: a sha-pinned
`demopatch`, or **an rext-owned Dockerfile in the shape `hiring.Dockerfile` already sanctions**.

**Chose the rext-owned Dockerfile.** A demopatch would have been the weaker of the two here:

- A demopatch's drift-refuse gate is a `pre_sha256` over the file it patches. `Dockerfile.dev` is a file the
  platform team edits — its last change (`M257x iter-146`) corrected the GraphQL default — so every such
  edit would refuse the patch and silently drop L1 back to a 4 GB image. `demopatch-spec.md`'s own war story
  is a silently-refused patch shipping a 76 s members grid for four releases; pinning this milestone's
  largest lever to that mechanism would be re-enacting it.
- The rext file has **no upstream to drift against**, and it is the shape the sibling build has used for
  four releases without incident.

**Cost, stated:** a genuine fork. `Dockerfile.dev` and `next-web.Dockerfile` can now diverge, and a platform
change to the former (a base-image bump, a new build ARG) will not reach the demo. That is the same standing
liability `hiring.Dockerfile` already carries, and the header records the provenance so a reader can diff
them.

## D2 — the ISOLATION clause is scoped to the **build output**, and the `.env` finding is ROUTED, not folded in

Scanning for baked keys surfaced the platform's own real-Clerk publishable key inside the next-web image, at
the **committed** `/app/apps/web/.env`. It is in the pre-L1 iter-08 image too, so it is pre-existing.

**Two ways to handle it, and the wrong one is seductive:** let the clause fire on it (it *is* a foreign key
in a built image) — or scope the clause and pretend nothing was found.

**Did neither.** The clause is scoped to the build output because that is the gate's own word — *"no built
image contains another stack's **baked** publishable key"* — and a committed `.env` is a build INPUT carried
as a file, not something baked; measured, the bundle carries **only** the minted key, in both the pre- and
post-L1 images. Scoping there is also not a loosening: an image that carried another stack's overlay would
bake that stack's key **into the bundle**, and the bundle is what is scanned.

**And the finding is kept** as `FIX-M257-committed-env-ships-real-clerk-pk`, with its evidence, rather than
absorbed. Letting the clause fire on it instead would have made the gate permanently red for a condition it
was not written about — and a permanently-red clause is a clause that gets switched off within an iter.

## D3 — the bundle probe returns `None` for "no measurement", distinct from `[]` for "measured, clean"

**Forced by a live control, about a minute after the first version was written.** The probe passed
`grep --exclude=.env`; the images are `node:24-alpine`, whose **busybox** grep does not implement it
(`grep: unrecognized option: exclude=.env`). It matched nothing, returned `[]`, and the assert reported a
CLEAN image having read nothing at all.

Two structural fixes rather than a flag swap: the probe now **depends on no grep flag beyond POSIX**, and
the "could not scan" outcome is a **distinct return value** that `isolation_assert` books as
`unreadable_bundle`. Pinned by a regression test naming the busybox cause.

**Recorded because of how it was caught.** 15 unit tests were green over the broken probe — correctly, since
a fixture's injected `pks_in` never touches busybox. Only running the assert against a **real image** found
it. *A capability probe that fails OPEN disarms the check it guards*, and the only reliable way to know which
way a probe fails is to run it against the real thing.

## D4 — a test was REWRITTEN, not repaired, and the distinction was checked before doing it

`test_unmodified_platform_dockerfiles_are_the_build_input` asserted `-f "$ctx/Dockerfile.dev"`. L1 makes that
literal false for next-web, so the test had to change — which is exactly the situation where *"a test can
REQUIRE the defect"* must be checked before touching it.

**Checked, and it was the second case, not the first.** The invariant the test's own docstring and class name
protect is **zero platform-repo edits**; `-f "$ctx/Dockerfile.dev"` was one *wiring detail* that satisfied it.
L1 satisfies it more strictly — the clone is no longer read even for its build recipe. So the test now grades
the contract in both directions (rext-owned Dockerfiles for the Next pair, the clone's own for studio-desk,
the trap-removal that leaves the clone pristine), and the end-to-end `git status --porcelain`-is-empty test
that actually proves the invariant was left untouched and still passes.

**The counter-case, for contrast:** `test_demopatch_hiring_role_remap_wiring` also went red, and there the
correct action was the opposite — the assert was *right* and my change had made its locator ambiguous. It was
re-scoped, not rewritten, and nothing it grades was weakened.

## D5 — editing a heavily-cited file broke 24 corpus citations, and only 5 of them were DETECTED

Adding ~150 lines to `buildbench.py` and ~21 to `up-injected.sh` shifted every line beneath them. The corpus
carries **61** `file:line` citations into those two files. Measured after the edit:

| | count |
|---|---|
| citations that still resolved to their intended content | 37 |
| citations silently pointing at **different content** | **24** |
| of those, **caught by the pre-commit fence** | **5** |
| knob anchors in `demo-up-defaults.md` broken (a separate, standalone guard) | 11 |

**The fence caught 5 of 24 — and that is not a fence defect so much as a limit worth naming.**
`anchor_construct_guard` books a citation when it lands on a **non-construct** — a blank line, a closing
`fi`, a `}`. A citation that shifts from one real construct onto a *different* real construct still lands on
a construct, so it passes. Nineteen did exactly that, including one that had been silently wrong **since
before this iter**: `build-budget.md` cited *"the argparse constructed at `buildbench.py:1464`"*, and at HEAD
`:1464` was `report["reclaim_attribution"] = …`. It resolved to a construct, so nothing ever complained.

**All 24 repaired**, with the numbers derived rather than guessed: a positional HEAD→worktree line map
(`difflib.SequenceMatcher`, `equal` blocks only) maps each intended line to its new number, which is exact
where content-matching is not — several intended lines are blank or a bare `#`, and two are duplicated
between the next-web and hiring functions. Verified after the fact by re-resolving all 61: **51 resolve to
their intended content, 0 stale.** The 11 knob anchors were regenerated with the guard's **own** `--fix`
(the sanctioned path, not a hand-edit), which then reported *"OK — the defaults table and the parsers agree,
both directions."*

**Routed forward: `FIX-M257-anchor-guard-content-drift`.** The whole class is mechanically detectable — the
line map above is ~15 lines and turns "did this citation land on *a* construct" into "did it land on *its*
construct". Not landed here: it is a guard change, this iter is already a two-line-shape iter under the
scope-creep tripwire, and it wants its own controls.

## D6 — the A/B's baseline arm is reported as a NULL RESULT

The third build of the A/B returned in 1.12 s: byte-identical to the discarded warm-up, so BuildKit served
the whole build including the export from cache. **That arm measured nothing.**

It is written into `progress.md` as a null result rather than quietly dropped, and the "before" column is
taken from iter-08 rep-03's real bring-up instead. The lever's two load-bearing figures — image size and the
export/unpack legs — are properties of the artefact and not of cache state, so the finding stands; but a
1.12 s number sitting in a scratch log is precisely the sort of thing that gets quoted later as *"the old
build took a second"*.
