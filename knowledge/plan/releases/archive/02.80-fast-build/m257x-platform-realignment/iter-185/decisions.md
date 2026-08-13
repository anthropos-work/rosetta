# iter-185 — decisions

## `D-M257x-185-1` — the class has a second member, and its costliest miss is `go.mod`

iter-184's rule (*a fence's POPULATION is a registry too*) predicted more instances. Swept `stack-core`
for **population-defining** literals rather than predicate literals; `predicate_enumerator.CITATION_RE`
is one — thirteen file extensions, typed, deciding what the enumerator can see at all.

Measured over `corpus/**.md` + `CLAUDE.md` in **file context**: eight undeclared extensions with ≥2
citations — **`mod` 51**, `jsx` 20, `hcl` 6, `gitignore` 6, `example` 5, `ini` 4, `dev` 3, `txt` 2 — and,
once the fence was written (it has no threshold; the census did), three more singletons: `vue`, `css`,
and one authority.

**`mod` is the one that matters.** `app/go.mod:14-18` alone is **12 citations** — the anchor CLAUDE.md's
shared-libraries banner rests on, the claim about which org modules `app` actually requires. A predicate
anchored there was outside the reach denominator that `TOK-07` / iter-114 made a declared thing.

## `D-M257x-185-2` — deriving was TRIED and REFUTED, and that is why the tuple stays

iter-184's rule prefers derivation. It was attempted and does not work offline:

| candidate rule | verdict |
|---|---|
| a `/` in the stem means a path | **REFUTED** — `.anthropos` has **14** with-slash hits, all `https://` authorities |
| not inside a `scheme://` token | better, still leaves `api.clerk.com:443`, `backend.internal.anthropos:8083` |
| resolve the path on disk | ground-truthed, but needs the clone set; the enumerator runs corpus-only |

So iter-184's **fallback clause** applies verbatim: *if deriving is impossible, the declaration is
itself a registry and needs the both-directions treatment every other registry gets.* Shipped:
`test_every_cited_extension_is_DECLARED` (a cited-but-undeclared extension is RED, naming it with
examples) and `test_every_DECLARED_extension_actually_occurs` (a declared extension that never occurs is
RED — the direction that let iter-184's own fence carry `PROBE` and `TASK`, two kinds that had never
existed).

**Note the arms found more than the census did**: the census used a ≥2 threshold and the fence has none,
so it immediately named `vue`, `css` and `de`. *A fence with a lower threshold than the survey that
motivated it is doing its job.*

## `D-M257x-185-3` — the authority carve-out is NAMED, and it is fenced against becoming a blanket

`com`, `net`, `org`, `work`, `anthropos`, `local`, `internal`, `io`, `de` are `host.tld:port`, not
`file.ext:line`; admitting them would put port numbers into a line-number census. Each is a live corpus
token, so this is a reasoned exclusion, not a silent skip (`§5` rule 8) — `de` exists for exactly one:
`u422950.your-storagebox.de:23`, the Hetzner Storage Box SSH port in `db-backup.md`.

**And the carve-out is itself fenced**, because a reason list is how a blanket exclusion disguises
itself: no tail may be in both lists, and the carve-out may not grow past the class it carves out of.

## `D-M257x-185-4` — the reach delta is PUBLISHED, not absorbed

The iter's own escalation condition. Widening the class changes what the enumerator sees, and a
population change with downstream numbers attached must be stated:

**1,276 → 1,376 citation tokens (+100, +7.8 %)** across **978 → 1,041** corpus lines carrying at least
one citation. Any reach percentage this milestone published against a corpus-derived citation
denominator was computed over the smaller one.
