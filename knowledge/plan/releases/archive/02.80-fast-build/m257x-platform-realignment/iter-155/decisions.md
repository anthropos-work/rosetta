# iter-155 — decisions

## `D-M257x-155-1` — "observed on a live stack" conflated three ingredients with three provenances

iter-153 declared `hiring-app` / `fake-fapi` / `fake-bapi` ungradeable on the rule that *a probe row needs
a container name, a host port and a health target **observed on a live stack***. The rule was right about
one ingredient and wrong to bind the other two to it:

| ingredient | provenance | derivable here? |
|---|---|---|
| container name | the layout constant `container_for_project()` already owns it | **yes** |
| host port | `gen_injected_override.py` at offset 0 | **yes** |
| health target | requires observing what the service answers on | **no** |

Because the registry's `docker` probe kind needs **only** a container name, all three services get a real
row that asserts something true — *the container exists and is running* — without inventing a target.

**The general form: when a blocker names a conjunction, check each conjunct's provenance separately.** A
single un-derivable ingredient had been suppressing two derivable ones, and the cost was three services —
including the two whose failure kills every login on a demo — reading as `✓ pass`.

## `D-M257x-155-2` — `fake-bapi` is why the line is drawn at `docker` and not guessed at `http`

`fake-bapi` publishes `127.0.0.1:5401 → 443` in-container: TLS with a minted cert. An `http` probe against
it would report `down` for the wrong reason — a **false failure dressed as a finding**, which is worse
than the silence it replaced, because it trains a reader to discount the row.

So the readiness half stays explicitly unclaimed and routed
(`FIX-M257x-iter155-injected-service-readiness-needs-a-live-stack`), and the fence asserts the kind is
`docker` **with the reason in the failure message**, so a future promotion has to arrive with an
observation rather than by relaxing an assertion.

## `D-M257x-155-3` — the empty array is a claim; deleting it would be silence

`STACK_INJECTED_SERVICES_NOT_PROBED` is now empty. It is **kept**, because it is the arrival surface: the
next service the injection generator grows will have no row, and the guard's arm needs somewhere to send
it. An empty declared set says *"we know of none"*; a deleted array says nothing at all, and saying
nothing is precisely what iter-153 created the array to stop. Fenced — the array must still be
*declared*, and must be empty while every injected service has a row.

## `D-M257x-155-4` — my own fence hard-coded the platform side, and its own arm caught it

The first draft of `test_the_generator_emits_every_injected_service_and_no_further_one` hard-coded the
platform-side service set as the five default-profile services. It went RED naming `next-web-app` and
`studio-desk` as undeclared arrivals — **they are platform services** (defined in the platform compose
under non-default profiles) that the generator re-emits with demo images and offset ports.

A hand-written set, inside the fence written to end hand-written sets, in the milestone about
hand-written sets. Now derived by delegating to `platform_topology.py` — the module that owns compose
parsing — rather than re-implementing a second parser. The distinction the fix encodes is real and worth
stating: **`rext-injected` means the platform compose does not define the service at all, not that the
generator re-emits it.**

## `D-M257x-155-5` — the THIRD consecutive fence pinned to a state of the world, and this one was mine

iter-153 re-pointed harden pass 35's fence; iter-154 re-pointed `dev-stack`'s contract test and wrote
`§5` **rule 71** about the shape. One iter later, rule 71 fired on **the fence iter-154 shipped alongside
it**: `test_the_ui_tier_enters_the_scope_and_the_unrowed_services_are_announced` asserted, as literals,
that the three services are *not* in the scope and that the tail prints `UNGRADED`. Both were true when
written; iter-155's **correct** change made both false, producing a RED indistinguishable from a
regression.

Re-pointed the same way as its two predecessors — to the property, which never changed: *whatever the
union reports unprobeable must be announced; whatever it reports probeable must be in scope.* The
expectation is now derived from `scope-union.sh` at test time, so the test is correct on both sides of
iter-155. The `else` branch is new and deliberate: with nothing unprobeable, the tail must **not** warn —
a warning nobody can act on trains readers to skip the real ones.

**What three instances in three iters say that one did not:** this is not a mistake a checklist catches.
Writing the rule did not stop its author from breaking it in the same commit. The durable defence is
structural — **derive the expectation from the same source the code derives from** — and that is now
rule 71's prescribed repair rather than a suggestion to be careful.
