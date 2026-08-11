# iter-270 — decisions

## `D-M257x-270-1` — the frozen-pin control is SPENT, deliberately, and this is the receipt

`D-M257x-258-1` held the rext pin at `fast-build-m257x-iter-101` so the platform-ref advance would be the
only changed variable in the demo proof. **That experiment concluded**: iter-258 (cold demo green on the
advanced refs, first attempt), iter-260 (three consecutive `--purge` + `up` cycles green) and iter-262 (a
dev stack from current `main`) all landed under the frozen tooling, and gate clause 1 is met.

Holding it any longer buys nothing and costs everything queued behind it — **six routes** carried the
words *"needs a tag + pin bump."*

**What the control bought, stated so it is not lost:** the advance is proven to build **under tooling that
predates it by 205 commits**. That is a stronger claim than "the advance builds" — it says the platform's
consolidation did not require a tooling change to be consumable. Every fix landing now is an improvement,
not a prerequisite, and that distinction was only available because the pin was held.

**What spending it costs:** the next demo/dev bring-up runs on tooling nobody has yet proven cold. That is
disclosed, not hidden — and it is the reason the two new fail-closed arms were each proven RED with their
precondition absent before shipping, rather than merely GREEN with it present.

## `D-M257x-270-2` — the `cms` key in `DIRECTUS_DATA_CONSUMERS` is KEPT, and the reason it can be kept shrank

Its comment justifies the key as rollback support: an older platform clone that still *defines* a `cms`
container must still be re-pointed at the per-stack Directus. This iter removed `cms` from
`INJECT_CANDIDATES` and from `INJECTED`, so such a clone would now get **no injected cms image** and would
fail earlier, for a different reason. **Half a rollback path reads as a supported one, which is worse than
none.**

It is nevertheless **not removed here**: that is a behaviour change with no evidence behind it in this
iter, and the class is not unwatched (`test_directus_consumer_derivation.py` grades the key against the
platform's own Go source on every run). Routed as
`ROUTE-M257x-270-directus-consumer-cms-key-outlived-its-rollback-path`, with the shrunk rationale written
into the constant itself so the next reader meets it.

## `D-M257x-270-3` — an OPERATING list must not name a corpse; a SCANNING domain must

Recorded because iter-270 shipped **both** halves within the hour, in adjacent files, and getting the
second one backwards would have silently disarmed a live fence:

- **operating** (`INJECT_CANDIDATES`, `INJECTED`, `REUSE_DEV`, `_studio_repos`) — naming a decommissioned
  service makes the tooling *act on* a corpse. Removed.
- **scanning** (`DIRECTUS_READER_DOMAIN`) — a husk clone still on disk that still *behaves* like a live
  consumer is exactly what the fence exists to notice. A domain that omits it cannot. Widened.

The concrete near-miss: the Directus-reader fence was passing `INJECTED` as its scan domain, so pruning
`cms` from an unrelated map instantly made *"cms stopped reading"* and *"cms was never looked at"* the
same green verdict. **A fence whose reach is a side effect of an unrelated map is not a fence.** The rule
is in `platform-alignment.md` §8.

## `D-M257x-270-4` — a stale citation the guard could not see until this iter moved a line

`corpus/ops/demo/build-budget.md:516` cited `up-injected.sh:490` and `:1008` for
`ctx="$DEMO/next-web-app"`. Measured at HEAD (**before** any iter-270 edit): line 490 was `return 0`,
line 1008 was a `log` line, and the real anchors were **565** and **1083** — wrong by ~75 lines, and
**green**, because `anchor_construct_guard` is a declared FLOOR: it detects *"resolves to nothing"*, never
*"resolves to the wrong construct."*

This iter's `+6` shift tipped `:1008` onto a bare `fi` and the guard fired. So the guard did its job — but
only by accident of which line the drift landed on. **A floor that only catches drift when the drift is
lucky is a floor, and it should be quoted as one.** Repaired to `:571` / `:1089` (re-derived, never
bumped). Seven corpus citations were re-pointed in total this iter (4 by the anchor-offset census, 3 by
the construct guard) plus 24 `Read at` rows auto-rewritten from the parsers by `demo_knob_guard --fix`.
