# iter-23 decisions

## D-M257x-23-1 — the colony-pin correction was INCOMPLETE, and it was applied anyway (in corrected form)

The handed correction named 4 of the 6 live services carrying a colony pin. Rather than apply it verbatim or
discard it, the row was rewritten from the measurement: `app` + `messenger` `v0.35.2`, `cms` + `jobsimulation`
`v0.35.1`, `sentinel` + `storage` `v0.34.3` — with the `cms`/`jobsimulation` pair explained as
still-running merged husks, so the reader does not read the third pin as a contradiction.

**Why this is worth a decision entry:** iter-22's rule ("re-derive the correction") was written against
corrections that are *false*. This one was *true and partial*, which the rule as written does not catch — a
verify-the-quoted-values check passes it. §5 now says: re-derive the ENUMERATION, not just the values.

## D-M257x-23-2 — `hiring.md` was re-grounded at its THESIS, not only at its rows

The doc's stated purpose is to name the one table that feeds the recruiter score, and that table is dropped.
Six row-level corrections would have left a doc whose opening argument points at a table that does not exist,
while its tables point at one that does — internally inconsistent, and worse than either alone.

Chosen: a `⚠️ RE-GROUNDED` banner immediately after the existing warning, stating the three facts that changed
(score source, schema, subgraph count), plus a `History` block at the headline section. **Rejected:** deleting
the mirror history. A seeder author who inherits a pre-drop blueprint needs to recognise the old shape to know
it is old; a doc that silently presents only the new shape cannot tell them anything.

## D-M257x-23-3 — the Directus consumer-list finding is ROUTED, not landed

`backend`, not `cms`, has been the Directus reader since cms-in-app; rext re-points only `cms`. Measured live
(`docker inspect demo-1-backend-1` → `https://content.anthropos.work`, empty token).

This is a third line of investigation on a two-line iter, so the tripwire applies: it is routed as
`FIX-M257x-iter23-backend-directus-not-repointed`. It also needs work this iter cannot do — an rext source
change, an inverted test (the existing one asserts the pre-merge shape and would fail on the fix), a tag
pushed to origin, and a cold cycle to prove.

**Explicitly NOT concluded:** that it causes `FIX-M257x-iter15-directus-versions-403` or the
`library_category` shape drift. Both are plausible downstream of an anonymous read against *prod* Directus,
and iter-19 proved them independent of the *serving* defect — which is a different defect and does not settle
this one. The next tik measures before attributing; this milestone has already refuted one attribution made
one iter after it was asserted.

## D-M257x-23-4 — `staging-clerk.md`'s allowed-origins lists left verbatim

They record what the Clerk instance holds, not what the platform serves. Editing `http://localhost:5050` out
would make the record inaccurate about Clerk. Annotated instead — the entry is now dead weight (nothing
listens on `:5050`), and `:8082` is what a cross-origin browser GraphQL call needs allowed.

The general rule this instantiates: **a transcript of external state is corrected by annotation, never by
edit.** Only claims the corpus makes in its own voice get rewritten.
