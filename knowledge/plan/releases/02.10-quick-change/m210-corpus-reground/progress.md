# M210 — Progress

Section checklist (closure = all boxes land).

- [x] Adopt + validate `origin/docs/skiller-in-app-merge` (architecture/subgraph half lands as-is) — 28 files clean-adopted (arch + service + ops + misc), each hunk verified vs the M208 merged facts + the re-synced stack-dev/app clone; commit 05932b1
- [x] Fix missed file `profile-completeness-spec.md` (43/44→44/44) — schema refs were already `public.*`; no literal 43/44 exists anywhere (exhaustively verified); made the one genuine merge-sweep fix (node-id home → `public`), did NOT fabricate a phantom count; commit 320c534 (see decisions.md §2)
- [x] Flip rext-facing tooling-doc bodies to `public.*` + delete interim notes: snapshot-spec, safety, recipe-snapshot-world, stories-spec, seeding-spec, coverage-protocol (+ directus-local) — the core work; 0 schema-qualified `skiller.<table>` remain describing tooling queries; commit fa382d5
- [x] Reconcile db-access ↔ tooling contradiction — adopted the colleague's `public.*` db-access re-point; now agrees with the §3-flipped tooling docs; commit b9b11a5
- [x] Sweep skill files: dev-up/reference, stack-snapshot/SKILL, stack-update/reference, db-query/SKILL — verified vs the re-synced compose (no skiller service, 11 graphql-profile containers, `SKILLER_RPC_ADDR=http://backend:8083`); superseded the colleague's stale stack-snapshot exit-4 note with an accurate M209-done note; commit ed8b30f (see decisions.md §5)
- [x] Update CLAUDE.md service catalog (skiller stub, 5→4 subgraphs, RPC addr) — adopted + added the explicit `SKILLER_RPC_ADDR` note + a Skiller "Archived/merged" entry for corpus-consistency; commit 06edb4f

## Done-bar — met
- Whole-corpus sweep: **0** `skiller.<table>` schema refs describing what the (re-grounded) rext tooling queries.
- db-access ↔ snapshot/seed docs agree on `public.*`.
- The colleague's architecture half landed; the missed file + skill files swept; CLAUDE.md catalog current.
- The email-asset PNGs on the colleague's branch were EXCLUDED (unrelated to the merge).

## Pre-flight audit
- Phase 0b KB-fidelity: **YELLOW** (proceed) at open — the pre-flip staleness was the milestone's own fix-list.
  After the flips, the corpus aligns with M209's landed `public.*` code → expected GREEN at close.
