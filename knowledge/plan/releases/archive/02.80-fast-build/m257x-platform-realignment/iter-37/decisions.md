# iter-37 — decisions

## D-M257x-37-1: a demo-scoped FALLBACK is not an alias, and the DNA's rule is about the source

**Choice.** The injected override gives `backend`
`SKILLER_AZURE_OPENAI_KEY=${SKILLER_AZURE_OPENAI_KEY:-${AZURE_OPENAI_KEY:-}}` (and the endpoint twin).
The rejected alternatives were (a) declaring an alias family in the secret DNA, and (b) leaving the
Playthrough failing and escalating for a dedicated key.

**Why not the alias family.** `secret-dna.json` marks these keys **DISTINCT-SIMILAR** from
`AZURE_OPENAI_*` with an explicit *"do NOT auto-alias"*, and that is correct: in **production** they
address a different Azure resource, and an alias family would make `/stack-secrets` write one live
credential into the other's slot on every stack — including a staging one. **The rule governs the
SOURCE.** This change governs one demo's compose wiring, where there is no second Azure resource and
never has been, and it cannot reach the source at all.

**Why not escalate.** `D-v28-3` forbids leaving a standing red, and the credential in question is one
"nobody has to hand" (the DNA's own words when it declined to classify the gene `critical`). Escalating
would have converted a fixable demo-wiring gap into an indefinite gate blocker.

**The ordering is the whole safety argument, so it is a test, not a comment.** The operator's own
`SKILLER_*` value wins whenever set; the fallback applies only to its absence. Mutant N2 inverts the order
and the suite kills it.

## D-M257x-37-2: resolve in Compose, never in the generator

**Choice.** Emit the `${…:-${…:-}}` expression and let Compose resolve it at parse time.

**Why.** The alternative — have `gen_injected_override.py` read the environment and write the resolved
value — is three lines shorter and would put **a live Azure credential into a file that tests read and
logs echo**, and into a generator that `/stack-secrets` guarantees is values-blind. Compose interpolation
keeps the value entirely out of the tooling: the generator never sees it, the override file never carries
it. Verified as a property, not as an intention — mutant N3 makes the generator resolve the value and the
suite goes RED on the "must be an EXPRESSION, never a resolved value" assertion.

Nested-default support was **measured** before being relied on (`${A_UNSET:-${B_SET:-literal}}` against a
throwaway compose project resolved to `B_SET`'s value), rather than assumed from documentation.

## D-M257x-37-3: prove the fallback with the .env keys REMOVED, and force the recreate

**Choice.** The live proof required (i) deleting the hand-added `.env` keys used for the initial
experiment and (ii) `--force-recreate`.

**Why.** A plain `docker compose up -d backend` reported `Running` and did **not** recreate — correctly,
because the resolved config was byte-identical whether the value came from `.env` or from the fallback.
Reading the container's env at that point would have measured the container created for the *experiment*,
and reported the fallback working when nothing had exercised it. The same "the check measured the
previous state" shape as iter-17's withdrawn cycles.
