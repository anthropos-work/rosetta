# iter-01 — intra-iter decisions

(Strategy lives in the milestone-root `decisions.md` as `TOK-01` + open decision `D1`.)

- **iter-01-D1 — colony replace granularity (to resolve in the first authn tik):** stub just
  `colony/authn`'s provider + the `orgclient`, or `replace` all of `colony`? `authn` is a package
  *inside* the `colony` module, so a `replace` swaps the whole module. Fallback (proven): vendor the
  whole `colony` as staging already does for its `vendor-colony/` v2-JWT patch. Decide once the authn
  twin's interface surface is extracted.
