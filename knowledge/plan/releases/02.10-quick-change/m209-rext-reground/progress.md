# M209 — Progress

Section checklist (closure = all boxes land).

- [ ] snapshot: flip `taxonomy/taxonomy.go:43 Schema "skiller"→"public"` + 2 PublicVia test assertions
- [ ] snapshot: narrow `pg.SchemaVersionSQL` digest to enumerated surface tables (Risk 1)
- [ ] snapshot: verify capture SELECT column list vs merged prod (Risk 2: small_embedding3 / extensions. opclasses)
- [ ] snapshot: recapture public taxonomy from merged-prod + ~42,763-row assertion + keep AssertPublicOnly
- [ ] seeding: re-point 5 real-SQL files skiller.*→public.* (keep `org_id IS NULL`) + isolation.go + data-dna.json
- [ ] seeding: rename 111 fake-Conn test string-matchers in lockstep; reword services/ai comments
- [ ] small: readiness.sh schema probe + services.sh container probe + up-injected.sh INJECT_SVCS + 5→4 note
- [ ] build + `go test ./...` green; tag rext `v2.1`; prep consumption re-pin
