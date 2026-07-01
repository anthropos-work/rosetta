# iter-06 — decisions

_Intra-iter implementation decisions (context → choice → why)._

## D1 — Seed only `user_languages`; let the DB trigger fan out `membership_languages`
`public.user_languages` carries an AFTER-INSERT trigger that INSERTs the matching `membership_languages`
row for every membership of the user. Seeding `user_languages` alone is therefore sufficient AND correct;
directly seeding `membership_languages` too would double-insert + collide with its own unique key. The
manager Talent tab reads `membership_languages`, so the trigger is load-bearing — verified 747→747 fan-out
on demo-1.

## D2 — `world_languages` = a curated ISO-639-1 STANDARD catalog (not a snapshot capture)
`world_languages` was empty on demo-1 (the snapshot didn't carry it). A published ISO-639 standard list is a
factual reference, not a fabricated taxonomy node-id — the closure gene governs skiller skill/role node-ids,
not ISO language codes (orchestrator-sanctioned). 16 EU-professional-weighted entries, deterministic ids
(COPY-idempotent on `code`).

## D3 — Cert coverage = ~45% of the supporting population, role-coherent (not 100%, not hero-only)
A believable org has a meaningful-but-not-universal share of credentialed members. Hero-only (9/340) read as
"Certification really low" (the field-review complaint); 100% would read fake. 45% (deterministic per-member
gate) + 1-2 certs each (vs heroes' 2-3) keeps the heroes the more-credentialed protagonists. The bank/skills
are role-coherent via the SAME deterministic role assignment UsersSeeder writes (no fabrication — empty pool →
general bank + no skills tag). Re-seed: 9→236 rows (158 distinct users).

## D4 — Strengthen the manifest with a `preAssert` tab-click + a `textMatch` OR-kind (the D4/F1 reconciliation)
The run-1 `(0,0)` passed BLIND — the manager manifest never asserted languages/certs/member-fields. The Talent
tab content renders only on a tab-click (in-page, not initial paint), and the members Location column needs an
OR-over-cities assertion (the grid paginates; no single city guaranteed on page 1). Added the minimal harness
capabilities: `preAssert` (a non-fatal tab-click-before-assert; a tab that won't open FAILs honestly) +
`textMatch` (regex-over-visible-text, the OR complement to `text`'s AND-of-substrings). The strengthened
manifest now PROVES: members Location header + a real seeded city + the Talent-tab "Languages spoken" (English)
+ "Certifications" cards. This is the gate the run-2 target requires — `(0,0)` on the strengthened manifest.

## D5 — Side-bug: `run-coverage.sh` forwarded consumed positionals to playwright as filename filters
Surfaced by adding `calibrate-manager.spec.ts`: `"$@"` (still holding `N vantage key`) reached
`playwright test`, which treats positionals as filename filters → "manager" matched the new calibrate spec →
it ran IN PARALLEL with the sweep on the same demo, corrupting the run. Fixed with a guarded shift-consume loop
(stops at the first `-flag`, so genuine extra flags still forward). A real tooling-correctness fix, recorded as
a side-deliverable (separate from the iter's planned scope).
