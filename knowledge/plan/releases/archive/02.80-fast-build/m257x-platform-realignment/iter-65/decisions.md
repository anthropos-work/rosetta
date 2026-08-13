# iter-65 — decisions

## `D-M257x-65-1` — a citation must NAME its subject, and for a known token that is decidable

`anchor_construct_guard`'s docstring draws a line this fence family does not cross: *"Catching #17
requires deciding what a sentence claims."* True in general. **Not true when the subject is a named
token.** G6's two-sided record required the corpus to cite a read site, and tested it with
`if site in all_text` — a whole-corpus substring match. Any document mentioning `main.go:446` for any
reason at all closed the finding.

The rule now requires the site **and** the variable name in the same **block** — the unit
`_pin_window` established at iter-63 (a paragraph of contiguous non-blank lines; a table row is one
line and therefore its own). No claim is parsed; two tokens must co-occur in one paragraph.

**Live corpus GREEN** — the existing two-sided records already name their subject in-block, which is
the right outcome for a strengthening and is *not* evidence the rule works. The fixtures are: a site
cited in ANOTHER block does not close it, a site with no variable anywhere does not close it, the
variable alone does not close it, and wrapped prose / a table row each count as one block. Reverting
to the substring test takes 2 of them RED.

## `D-M257x-65-2` — G6's universe excluded the consumer side, and a FIXTURE found it

Writing the test `test_a_site_named_with_no_variable_anywhere_does_NOT_close_it` produced a *green*
result where RED was expected. The cause was not the citation rule at all:

```
universe = set(compose.rpc_addrs) | named_anywhere | set(env_example_names)
```

A variable the platform configures **nowhere** and the corpus names **nowhere**, yet `app` **reads**,
had no row in the split — so G6 could not see it. **That is the most-undocumented case there is**:
silent at boot, absent from every document, and invisible to the assertion whose entire job is to
catch it. The universe now includes `set(app_reads)`.

**Second time this milestone a FIXTURE has surfaced a reach hole the live corpus happened not to
exhibit** (iter-61 was the first). A fence's blind spots are not discoverable by running it on a tree
that does not contain them — which is the argument for fixtures with known answers, stated in the
protocol and now demonstrated twice.

## `D-M257x-65-3` — the `pms:87` anchor resolved to a row about something else

`service_taxonomy.md`'s Directus retraction appealed to `platform-migration-status.md:87` as *"the
corpus's own fenced source of truth"* for whether a Directus compose service ever existed. **That map
has no Directus row at all** — it maps *repos*, and Directus is an external service. The anchor
resolved (the line exists, carries content, and passed `anchor_construct_guard`) and named
`anthropos-studio-room`.

Adjudicated against the platform, which is the only thing that could settle it and which the same
paragraph already cites: `git show a2a3ee6^:docker-compose.yml`. The appeal to the map is removed and
the reason recorded in place, so the next reader does not re-derive it.

**The instructive part is that every mechanical check passed.** The line resolved, carried content,
and was faithfully re-pointed twice by iters 63–64's line-map re-point — each time preserving a
citation that was pointing at the wrong subject all along. **A re-point preserves intent; it cannot
audit it.**
