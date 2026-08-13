# iter-247 — decisions

## `D-M257x-247-1` — invert the polarity: fence ABSENCE, not presence

`ROUTE-M257x-237` proposed the must-exist reading. The substrate refutes it in one measurement: of **327
distinct `UPPER_SNAKE` tokens / 1,213 occurrences**, the most-cited are variables the platform
**DELETED** (`CMS_RPC_ADDR` ×38, `SKILLER_RPC_ADDR` ×19, `STORAGE_RPC_ADDR` ×18). A must-exist fence
would redden the corpus's most careful writing, because **absence is the point of those sentences**.

Absence claims have the opposite property: exactly one reading, about the platform, falsifiable by the
platform at any moment.

## `D-M257x-247-2` — grade FAMILY globs, not single variables

Measured: 12 variables carry an absence phrase over 29 sites, **6 of the 12 are actually present**, and
**all six are co-location artifacts** — the phrase belongs to a neighbouring clause. Narrowing to
one-env-token lines leaves 2 sites and *still* mis-attaches, because the real subject of *"zero
`*_RPC_ADDR` variables anywhere in compose"* is the **glob**.

`*_RPC_ADDR` is asserted absent at **23 sites across 15 documents**. Grading the glob covers all of them
with one assertion, and it is the claim that would actually break.

## `D-M257x-247-3` — a claim must name its own file; a read-axis claim is a category error

Only a sentence naming `compose` / `docker-compose.yml` / `.env_example` is graded — of 41 family-absence
sites, **23 do and 18 do not**, and the 18 are mostly frontend/rext (`DEMO_NO_*`, `BUNNY_*`, `VITE_*`).

Separately, *"read by nothing"* is about **code**. `GRAPHQL_SCHEMA_FOR_GEN` is declared in `.env_example`
**and** read by nothing; both true, only one about this file. Refused with that reason rather than
silently skipped.

## `D-M257x-247-4` — a bare `no` is not an absence quantifier

It produced the guard's only two findings and both were false: *"no rebuild needed"* and *"with no
`NEXT_PUBLIC_*` / env / compose override"* (the absence of an override **seam** for one value). One of the
two sentences asserts the **opposite** of absence. Dropping it took the guard to zero and removed no real
claim — `*_RPC_ADDR`'s sites say *"zero"*.
