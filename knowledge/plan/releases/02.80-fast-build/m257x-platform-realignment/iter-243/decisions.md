# iter-243 — decisions

## `D-M257x-243-1` — keep the NAME, do not re-derive it

Two repairs were available: re-derive the family from something else at the call site, or stop throwing
it away. The second is the only one that fixes the class.

`blueprint.ParseStackN` returns an **offset** — it takes everything after the first `-` and discards the
prefix. Every function downstream of it that needs to know *which family* a stack belongs to has to
either be handed the name or invent one, and inventing one is exactly what produced this defect. The fix
threads the name through; `ParseStackN` is untouched, because returning the offset is its job.

## `D-M257x-243-2` — an unqualified `--stack 1` keeps the historical `demo-N`

`ParseStackN` accepts a bare offset (a name with no `-` parses whole), and until iter-239 the corpus's
own runnable examples taught exactly that spelling. With no family in the name there is nothing to
honour, so the fallback is the pre-existing `demo-N`.

That makes the demo path **byte-identical** — which was this iter's stated escalation condition: the
demo path is the proven one and this is a dev-path correction. Pinned by two of the five table cases.

## `D-M257x-243-3` — `P-243-4` is REFUTED and the class is NOT generalised

The hard-coded `demo-` shape occurs at **exactly one** site in the whole `stack-seeding` section. There
is no pattern to sweep, so none is swept and no fence is built for a population of one.

Recorded because the reflex under `TOK-08` is to census every class, and a census whose population is a
single known instance is ceremony — `§5`'s *measure a hazard's size, or "the same problem exists
elsewhere" is only a mood* (iter-168), applied in the direction that says **stop**.
