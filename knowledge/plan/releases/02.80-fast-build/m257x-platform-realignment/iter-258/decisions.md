# iter-258 — decisions

## `D-M257x-258-1`: the rext pin stays at `fast-build-m257x-iter-101`, deliberately

**Status:** taken at iter open, before the bring-up, and recorded because the alternative is that a
default settles it silently.

Pre-flight rung zero **passes** — the pinned tag resolves on origin (`0011c10a`, among 470 origin tags),
so the M236 *"tagging is not publishing"* failure is not present. What IS present is staleness: the pin is
**157 iters** behind the milestone, and the authoring copy is **5 commits ahead of `origin/main`**,
unpushed and untagged, therefore invisible to any stack.

**Accept the stale pin.** Reasons in descending weight:

1. **Experimental control.** `demo-1` went green (`0` warnings, 2026-08-06T10:21:56Z) under **this exact
   tooling version**, at app `ad9f3c498`. Holding tooling constant makes the platform-ref advance the
   single changed variable. Bumping the pin changes two at once, and an unattributable failure is
   precisely the condition this route exists to remove.
2. **It would not change what gets built.** iter-257 measured `DEMO_ADVANCE_CLONES` default **`0`**, with
   nothing outside `ensure-clones.sh` setting it — a default bring-up applies **no pin**, building the
   clones as checked out. The advanced `clones.pin.json` is not consumed either way.
3. **Publishing a tag mid-experiment is a side-effect, not a control.**

**The gap is disclosed, not closed.** iter-256's `clones.pin.json` advance and iter-257's
`clone_pin_guard` arm D are absent from the tag this demo clones. **Conditional obligation:** if the
bring-up fails *inside rext tooling* rather than inside platform code, the stale pin is a candidate cause
and the failure must be re-tested at a fresh tag **before** any platform verdict is published. Routed as
`ROUTE-M257x-258-the-pin-is-157-iters-stale` regardless of outcome.

## `D-M257x-258-2`: the registry that decides allocation is not the registry in the obvious place

Two files named `registry.json` describe stack slots and they **disagree in both directions**:

- `demo-stack/stacks/registry.json` — records **slot 2**, omits the live **`demo-1`**. The pre-M12
  demo-only legacy file, kept for demo provenance (`rosetta-demo:56-57` says so).
- `stack-core/.stacks/registry.json` — records **`demo-1` only** (`adopted: true`, `status: up`). This is
  `DEFAULT_REGISTRY` at `stack_registry.py:51`, the unified dev+demo allocator's record.

Only the second decides whether an N is free. The first is what a reader finds when they look under
`demo-stack/`, and it would have told this iter that slot 2 was taken and slot 1 was free — **exactly
inverted from the truth**. Recorded rather than repaired: the legacy file has a declared purpose, and
its content is provenance, not allocation. What was missing is the sentence saying so where a reader
looks.

## `D-M257x-258-3`: the demopatch drift measures BASELINE STALENESS, not the advance — corrected in-iter

This iter's own first draft of `progress.md` wrote that the six drifted demopatches drifted *"because of
the advance — these files moved between `ad9f3c498` and `3eaadae68`."* **The manifests refute it**, and
the refutation was taken before the claim reached a commit:

| patch | pinned baseline | clone actually at |
|---|---|---|
| `app-targetrole-authz-skip` | v1.295.0 | v1.371.1 |
| `app-aireadiness-snapshot-loadmembers` | `app@3df8536` (v2.7 pin) | v1.371.1 |
| `next-web-ssr-graphql-origin` | v2.108.0 | v2.137.3 |
| `next-web-studio-url` | v2.106.1 | v2.137.3 |

Every baseline predates iter-256's advance by tens of minor versions, so the drift cannot be attributed
to it. **The withdrawn claim was the convenient one** — it would have made the advance look consequential
in a section where the iter's whole job is to characterise the advance.

Two rules fall out, and the second is the general one:

1. **The corrected reading is stronger than the withdrawn one.** The anchors held across a ~76-minor
   gap, which says more for the anchor design than a 28-commit gap ever could. A correction is not
   automatically a downgrade.
2. **A drift signal names the distance between a PIN and a TREE — it never names the last thing that
   moved the tree.** Reading it as "the most recent change did this" is the same error class as reading
   a service repo's `service_desired_count` as production state: an input mistaken for an outcome.

Chain-rule sub-finding, confirmed by direct manifest read rather than by citing the spec:
`next-web-studio-url.post_sha256` **==** `next-web-public-website-url.pre_sha256` (`fe15aa715a17…`), both
on `packages/core-js/src/constants/urls.ts`. The second chains on the first, so its DRIFTED line is
**by design** and inherits the first's drift. Counting it as an independent stale baseline would
double-count a declared dependency — so the honest tally is **5 stale baselines + 1 chained**, not 6.
