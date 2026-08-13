# iter-55 — decisions

## D-M257x-55-1 — re-point to `0dab54d` inside this iteration, not as a follow-up

TOK-04's P3 says the detecting iteration re-points, in that iteration. Origin had moved one commit ahead of
the ref TOK-04 named, so the clone was fast-forwarded **before the first measurement**. Fetch + checkout
only; zero platform edits; tree clean.

The alternative — measure at `ef32d4c` and note the drift — is what the previous ten readings did, and it
is why none of them can be quoted. A number taken against a ref that is not the gate's ref does not answer
the gate's question.

## D-M257x-55-2 — DERIVE the profile rather than correct the literal

A corrected literal (`graphql` → `core` in five places) is a five-minute change that is correct today. The
platform's own roadmap says v9.0 is mid-flight and further folds follow, so its expected lifetime is short —
and its failure mode is silent: a stale profile name is not an error to compose, it selects an empty service
set and brings up nothing.

Deriving from the `backend` service's own `profiles:` row costs one small module and is correct at both refs
with no edit between them — which is the property under test, and the first thing the new test file asserts.
`platform-alignment.md` §5 rule 27's first branch, applied where it was cheapest.

## D-M257x-55-3 — `FALLBACK_PROFILE` is tolerated, and fenced rather than removed

`gen_injected_override.frontend_lines(n, offset)` is called without a `platform_dir` by the exposure guard
and several unit tests. Those callers need *a* profile name and have no clone to derive one from.

Rather than pretend otherwise, the constant is kept, documented as **prose-under-review** (§5 rule 27's
third branch: correct at `0dab54d`, will rot at the next rename), and fenced — the production path threads a
derived value, and `TestProductionPathIsDerived` proves it using a sentinel profile name
(`zzz-not-a-real-profile`) that no platform will ever use, so a green cannot be an accident of the two
values coinciding. The fence carries its own mutation self-test.

## D-M257x-55-4 — the clause-2 reading on the 44-hour-old stack is a CONTROL, not a restoration

It has the right shape (`passing=30 failing=0 unimplemented=1`). It was taken on a stack built at platform
`28c5f0d`, **three folds behind origin HEAD**, still running `cms`, `jobsimulation`, `roadrunner` and
`storage` containers. Reporting it as clause 2 restored would attach it to a ref it never touched.

That is exactly the defect TOK-04 was written about: the iter-37 reading's platform ref *existed only by
adjacency*. Repeating it here — with a wider adjacency — would be the milestone contradicting itself in the
same week it named the rule.

## D-M257x-55-5 — the teardown asks Docker, not the compose file

The purge failure has no compose-file-shaped fix, because the compose file is the thing that went stale. Any
fix phrased as "regenerate the override first" or "tear down with the old base compose" is another statement
about topology that can itself go stale. Docker's `com.docker.compose.project` label is the one source that
cannot.

Deliberately NOT chosen: leaning on `docker compose down --remove-orphans`. It was already in the command
and did not help — twice, for two different reasons. Once because the whole project was invalid, so nothing
ran at all; and once (observed live on the fixed run) because `storage` is still *declared* in the base
compose under `storage-legacy`, so it is not an orphan — while not being in the `core` profile, so it is not
selected either. **A container can be simultaneously not-an-orphan and not-selected, and fall through both.**

## D-M257x-55-6 — three stranded demopatches are recorded, not force-reverted

Killing the first bring-up mid-build left `stack-demo/next-web-app` with patches applied. `demopatch revert`
cleanly reverted five; three refused (`next-web-back-to-cockpit`, `next-web-studio-url`,
`next-web-public-website-url`) because their manifests' whole-file shas are stale against the current clone —
the known self-healing-freshness condition, which the chained pair compounds.

`demopatch --force-pristine` exists for this and performs `git checkout -- <path>`. **Not used**: this
session operates under an absolute ban on `git checkout --` in any tree, and a tool flag that performs the
banned operation is still the banned operation. Routed as `FIX-M257x-iter55-stranded-demopatch-revert`.

Not blocking: the apply path is idempotent and anchor-based, and it logged
`WARN … whole-file sha DRIFTED … but the anchor is intact (1x)` on exactly these manifests during the run,
which is the tolerated case.

## D-M257x-55-7 — close on Phase A + B, route Phase C forward

The scope-creep tripwire fired: the teardown defect was a third, unplanned line of investigation. Per the
tripwire, what is complete lands and the rest is routed with a named handler.

Clause 1 could not have been met by this iteration in any case — it requires **three consecutive** cold
cycles, and the first one available at the new ref is the one that surfaced the teardown defect. Reporting
fewer than three, or counting the aborted cycle as one, would be D-M257x-55-4's error in a different place.
