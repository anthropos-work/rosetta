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
