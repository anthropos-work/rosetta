---
milestone: M257x
iter: 07
iteration_type: tik
status: archived
opened: 2026-07-31
---

# iter-07 — `REPOINT-M257x-cms-similarity-writes`

**Active strategy reference:** `TOK-01` — *instrument first, then follow*, step 2: **"fix the mechanism, not
the symptom"**, and its governing rule *"derive it, or fence it. Never both-hand-maintain-and-trust it."*

## Step 0 — re-survey (mandatory; measured, not inherited)

The pre-compute at `progress.md` was re-measured against live `demo-1` before targeting. It holds, and the
re-survey added **one fact the pre-compute did not have, which changes the design**:

**The snapshot cache key does NOT contain the schema name.** `pg.SchemaVersionSQL` digests
`table_name || '.' || column_name || ':' || data_type`; the schema is only the `WHERE` filter. So the
staleness digest of a *narrowed row surface* is **schema-independent**. Measured on `demo-1`:

    digest over public  narrowed to the 4 similarity tables : 032c99ea47678187631c59c31b4ef059
    digest over cms     (same 4 tables)                     : <empty — cms holds 0 tables>
    cached manifest schema_version (captured 2026-06-29)    : 032c99ea47678187631c59c31b4ef059   <-- EXACT MATCH

Three consequences, none of them assumable:

1. **The 2026-06-29 capture is NOT stale** — and this is a *stronger* statement than the pre-compute's
   column-by-column comparison, because the digest covers name **and type**, ordered, for all 4 tables.
   The "re-capture freshness is a separate question" caveat is **answered: no re-capture is needed.**
2. **The cache will HIT the moment the probe reads the right schema.** The failure is entirely a
   *resolution* failure, not a data or freshness failure.
3. The re-point therefore has to move exactly two things — the **probe's** schema and the **replay's**
   schema — and must move **neither** the capture side (prod read, `D-M257x-7`).

**Write surface, measured by scan and split live-vs-comment (§7 rule 1).** Exactly **one** live site names
the schema: `stack-snapshot/simembeddings/simembeddings.go:44` `const Schema = "cms"`. Every other
occurrence in the tree (`dev-setdress.sh:140,340,342`, `repos_yml.sh:97`) is a **comment**. This is the
smallest write surface of the three folds — and the reason it is nonetheless the hardest is that the one
constant is read by both halves of the system.

## Cluster / target identified

`REPOINT-M257x-cms-similarity-writes` — the **last** `REXT_TRANSITIONAL_SCHEMAS` entry, i.e. the last thing
between here and a claimable **gate clause 4**. TOK-01's named next target; the re-survey confirms it is
still untouched and still the right one (no substitution).

## The design decision (the pre-compute deliberately left it open)

**Rejected — `ReplaySchema = "public"`.** Two lines, and it is *the same hand-maintained constant this
milestone exists to end*. It would be wrong again at the v9.0 fold, and it is precisely the shape
`platform-alignment.md` §2 forbids.

**Adopted — a derived, replay-time schema resolver.** For a surface's declared schema `S` and its table
set `T`, ask the **target** (never the source):

| what the target says | resolution |
|---|---|
| every table of `T` exists in `S` | use `S` unchanged — taxonomy + directus are untouched |
| no table of `T` is in `S`, and **exactly one** other schema holds **all** of `T` | remap to it, and **say so loudly** |
| **no** schema holds all of `T` | fail loud — the surface is unprovisioned on this stack |
| **more than one** schema holds all of `T` | fail loud, **naming the candidates** — never guess |

This is "follow the platform when it moves" implemented once, generically. It self-heals at the next fold,
it needs no per-surface edit, and every ambiguous case is a loud failure rather than a quiet guess.

**Two guards it must carry, from this milestone's own doctrine:**
- §5 rule 7 — *a probe must not be able to satisfy itself*. The resolver's answer is read from
  `information_schema.tables` **on the target**, and it may never fall back to the declared value when the
  lookup fails. A lookup error must propagate, not degrade to identity.
- §8 rule 3 — *pin the mechanism, not the contents*. The tests fence **where the schema comes from**, not
  the string `"public"`.

## Hypothesis

Resolving the replay-side schema from the target makes the `sim-embeddings` surface replay its 274/278/274/664
rows into `public.*` on a stack where the platform no longer creates `cms` — turning the current
`sim-embeddings replay skipped (rc=4)` into a load. `REXT_TRANSITIONAL_SCHEMAS` then goes `"cms"` → empty and
the no-growth fence's **shrink branch** fires for the second and last time, which is the deliberate act that
makes **gate clause 4** claimable.

## Expected lift

- `sim-embeddings` replay: `rc=4 skipped` → **loaded, 1490 rows across 4 tables** into `public`.
- `REXT_TRANSITIONAL_SCHEMAS`: `"cms"` → **empty** (gate clause 4 claimable).
- autoverify: **no regression** (the surface is currently skipped, so this is a net-add).

## Phase plan

1. **Phase A** — the resolver: a pure decision function + one catalog query, both unit-tested.
2. **Phase B** — wire it through the two consumers that must move (the CLI's probe, `replay.Run`) and
   **not** the one that must not (capture).
3. **Phase C** — fences, each mutation-verified RED.
4. **Phase D** — prove it LIVE on `demo-1`: replay the surface, count the rows, then **drop the `cms`
   schema from the stack** and re-run to prove the stack does not need it.
5. **Phase E** — pay the debt down; watch the shrink branch fire; close.

## Escalation conditions

- If the resolver would have to guess between two schemas on a real stack → that is a platform-shape
  finding, not a tooling detail: record it and escalate rather than allow-list.
- If dropping `cms` from the CREATE SCHEMA set breaks a running container (the `cms` container **is** still
  in the default local profile per `D-M257x-3`) → route the paydown forward and keep the re-point.

## Acceptable close-no-lift outcomes

- The resolver lands and is proven correct, but the live replay is blocked by a cache/stack precondition
  that is out of this iter's scope — provided the blocker is measured and named.
