---
iter: 02
milestone: M211
iteration_type: tik
status: closed-fixed
created: 2026-07-08
---

# iter-02 — tik: recapture→replay loads public.* into the warm merged stack

**Type:** tik — first tik of M211, under **TOK-01** ("warm-first cache-migrate, then cold-prove").

## Step 0 — Re-survey
Re-ran the baseline probe: warm `public.skills` = 0 rows, all 10 public taxonomy tables present, skiller
schema absent, no skiller container. TOK-01's next-tik target (sub-condition (b)) is still the right + open
next thing. No substitution.

## Active strategy reference
**TOK-01** — its next-tik direction named sub-condition (b): taxonomy replay loads `public.*` (~42,790)
into the WARM merged stack, via the user's cache-migration.

## Cluster / target
Sub-condition **(b)** of the composite gate: recapture→replay loads `public.*` (taxonomy replay rc 0,
~42,790 public skills). This is the biggest data-prerequisite blocker; every downstream sub-condition
(seed, verify, coverage, playthroughs) depends on a real public taxonomy being present.

## Hypothesis
The skiller→app merge was a pure schema-prefix move, so the existing real 42,790-row cache (captured under
`skiller.*`) IS the real public taxonomy. Re-keying it to `public.*` + keying it under the merged-schema
digest → the re-grounded replay loads it faithfully. The one merged-schema addition (`ts_search`, a
GENERATED tsvector) auto-computes because the explicit-column-list COPY correctly omits it.

## Expected lift
+1 gate sub-condition (b): 1/6 → 2/6 met.

## Phase plan (executed)
1. Empirical column-match (D1 gate): `\d public.skills` vs the 15 cached columns.
2. Build stacksnap from the re-grounded authoring copy.
3. Probe the merged-public digest (dry-run capture) → the new cache key.
4. Execute the cache-migration (re-key manifest schema/payload/filter/public_via + hardlink payloads under
   `public.X.copy` + `schema_version`=new digest).
5. Replay into the warm stack; measure rc + `public.skills` count + `ts_search` population.
6. Re-pin consumption (`.agentspace/rext.tag` → `quick-change-m209`).

## Escalation conditions
- Unreconcilable column drift → `user-blocker` (user decides prod-DSN vs gate-partial). **Did NOT fire** —
  the only drift is the added GENERATED `ts_search`, which COPY omits and PG auto-computes.

## Outcome
**Target MET.** Replay rc 0; 330,261 rows loaded (10 tables); `public.skills` = 42,790 (all public,
`organization_id IS NULL`); `ts_search` populated on all 42,790; both pgvector indexes rebuilt; sample
skills carry real names + valid node_ids (genuine taxonomy, not fabricated). See progress.md close.
