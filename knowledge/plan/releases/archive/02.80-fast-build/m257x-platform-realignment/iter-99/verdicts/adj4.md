# Adjudicator 4 — verdicts for seats E, F, G (reading #22)

Every booking below was **re-derived from the platform clones**, opening the cited file at the ref the
claim itself names, and reading around the line. No seat's evidence and no prior verdict was used as
proof. Refs used: platform `0c91421d`, app `b948604f` (+ `origin/main` `2035f9a4`, `5ba17044`,
`3e5bc33ef` where a claim names them), next-web-app `bb3313bc`, ant-academy `9c3843cd`.

---

## Seat E

### E B1 | `corpus/architecture/service_taxonomy.md:405` | UPHELD | IN-SCOPE | Compose sets a second service address, `GOTENBERG_URL`; the edge count is not one

- evidence: `stack-demo/platform` @ `0c91421d`, `docker-compose.yml:57` on `backend`'s own env block
  (nine lines below the cited `:48`) is `- GOTENBERG_URL=http://gotenberg:3200` — an in-compose
  service, addressed by service name + port. `stack-demo/app` @ `b948604f`: `main.go:244`
  (`GotenbergURL: os.Getenv("GOTENBERG_URL")`) feeds `internal/converter/gotenberg.go:31`
  (`http.NewRequestWithContext(ctx, "POST", gotenbergURL+"/forms/libreoffice/convert", &body)`) with a
  dedicated `gotenbergClient` at `:13` — a synchronous, cross-container HTTP call.
- the passage is not shielded by scope: **this same document** puts Gotenberg in the *Tier 1: Core
  Backend Services* table (`service_taxonomy.md:81`, profile `core, backend, all`) and in the `core`
  five-container selection (`:68`, `:464`, `:487`). Under the file's own taxonomy the section heading
  *"Core Services ↔ Core Services"* has two synchronous cross-process edges, not one.
- I verified the clause the sentence gets **right** before booking the one it gets wrong:
  `grep -c '_RPC_ADDR'` over `0c91421d:docker-compose.yml` = **0**, and at `0dab54d` the four
  (`BACKEND_USERS_`, `CMS_`, `JOBSIMULATION_`, `SKILLER_`) all sat on the deleted `messenger` block.
  The `*_RPC_ADDR` half is true; *"a single service address"* is false at the ref the sentence names.
  (`REDIS_ADDR=redis:6379` at `:66` and the two `postgresql://…@postgresql:5432` conns at `:93-94`
  are further in-compose service addresses; gotenberg is the load-bearing one because the document
  itself calls it a core backend service.)

### E B2 | `corpus/architecture/service_taxonomy.md:101-102` | UPHELD | IN-SCOPE | `make up-all` never ran two Brevo pushers; backend's own was never on locally

- evidence: `stack-demo/app` @ `origin/main` `2035f9a4`: `env_guards.go:92-111` — `resolveSubsystemSwitch`
  returns `(false, nil)` for an empty value when `deployed == false`; `deployedEnvironment()`
  (`:37-44`) returns **false** for `ENVIRONMENT=development`. `main.go:286`/`:394-396` build the
  customerio-sync manager only under that switch.
- the window the sentence describes is closed on both ends. `internal/customeriosync/sync.go` was
  ADDED at `3e5bc33ef` (2026-08-04) and **already** gated there —
  `3e5bc33ef:main.go:387` reads `if deployedEnvironment() && os.Getenv("BREVO_KEY") != ""` — before
  `3df469da8` (same day) replaced that with the named switch. Platform `0dab54d:docker-compose.yml:56`
  sets `ENVIRONMENT=development` on `backend`, so the in-process pusher was OFF at every ref between
  the fold and the container's deletion at platform `838d907` (2026-08-05).
- and it self-contradicts three lines up in the same blockquote (`:98-100`): *"The last two stay **OFF**
  on a developer machine behind `MESSENGER_ENABLED` / `CUSTOMERIO_SYNC_ENABLED`, which compose
  deliberately does not set"* — corroborated verbatim by `0c91421d:docker-compose.yml:84-92`
  (*"…which default to OFF on a developer machine (ENVIRONMENT=development is what makes unset mean
  off)"*). If backend's own is off, `make up-all` started **one** pusher, not a second alongside it.

### E B3 | `corpus/architecture/security_compliance.md:158` + `:222-224` | REJECTED | IN-SCOPE | US DR row vs "no customer data in US by default" is reconcilable, not incompatible

- evidence: I opened both. `:153-160` is the *Backup & Disaster Recovery* table — `| **DR site** | US
  AWS region |`, with `| **Full backups** | Every 6 hours → S3, Azure, Hetzner (Germany) |` at `:155`
  and `| **Primary region** | EU-West-1 (Ireland) |` at `:157`; no US storage is asserted for any
  backup target. `:222-224` reads *"No customer data stored in US **by default**"*. A disaster-recovery
  site is by construction a non-default (post-failover) locus, and *"by default"* — the same hedge
  carried at `:7` and `:195` — absorbs it. The two passages can both be true; rule 5 requires
  passages that assert **incompatible** things.
- the second half of the booking mis-reads the enumeration's scope. `:200-206` routes the reader to
  `external_services.md:602-607` for ways **a request** leaves the EU — inference traffic — not for an
  inventory of data-at-rest locations. A storage site is out of that list's frame, so it is not
  "omitted" from it.
- nothing here is measurable from any clone (`infrastructure` and `db-backup` are in no clone set), so
  neither side can be shown false; and the seat's own confidence is LOW on exactly that ground.
- class: mis-read — the "by default" hedge and the request-scoped enumeration reconcile the two
  passages; no measurable falsehood and no strict contradiction.

---

## Seat F

### F B1 | `corpus/architecture/dependency_map.md:59` | UPHELD | IN-SCOPE | `SKILLER_STREAM` is 6 Go occurrences across 3 files at `2035f9a4`, not 4

- evidence: measured at the ref the sentence itself names. `git -C stack-demo/app grep -c
  SKILLER_STREAM 2035f9a4 -- '*.go'` → `main.go:2`, `subscriber_merge_test.go:3`,
  `subscriber_wiring.go:1` = **6 occurrences / 3 files**. `git grep -l` at the same ref lists exactly
  those three.
- three-instrument check before accepting the denominator: `git grep -a` (binary-as-text) at
  `2035f9a4 -- '*.go'` returns the same three files, and a byte-scan for NUL over every `.go` blob at
  that ref (`git show … | tr -dc '\000' | wc -c`) finds **zero** NUL-bearing Go files, so no source is
  being skipped. Off the `*.go` pathspec the set is **6** files (adding two `knowledge/**.md` and
  `terraform/main.tf`); dropping only the markdown gives 4 files but **7** occurrences. No reading of
  the corpus yields "6 across 4".
- rule 33 offers no shelter: the sentence pins `2035f9a4` and says *"where"*, so it is graded there.
  The two neighbouring assertions in the same cell are correct and I confirmed both —
  `2035f9a4:main.go:1276` is `apiKeyManager,`, and at `b948604f` `SKILLER_STREAM` has exactly one Go
  occurrence.

### F B2 | `corpus/services/ant-academy.md:63` | UPHELD | IN-SCOPE | Progress writes are posted from the client harness, not the beacon fallback route

- evidence: `stack-demo/ant-academy` @ `9c3843cd`. The cited route's own header,
  `code/app/api/academy/beacon/route.js:1-18`, states the mechanism in the opposite order: *"POST
  /api/academy/beacon — the on-unload last-ditch write flush … the real academy mutations go to the
  CROSS-ORIGIN WunderGraph supergraph with a Clerk Bearer token — and sendBeacon can't set an
  Authorization header … a best-effort last-ditch flush for a write that would otherwise be lost if
  the tab closes mid-retry"*, declaring itself fire-and-forget and fail-closed at `:16-18`.
- the real write path, re-derived independently: `code/src/progress/store.js:26-27` imports
  `UPSERT_CHAPTER_PROGRESS` / `SET_LAST_ACTIVITY` and fires them at `:162` and `:210` through the
  requester injected by `code/src/progress/useProgressSync.js:53-60`
  (`setProgressRequester(async (document, variables) => { const client = await graphqlClient(); return
  client.request(document, variables) })`) — i.e. straight to the supergraph. `store.js:139-146` names
  itself *"THE PROGRESS WRITE SEAM (backend-authoritative, immediate write-through)"*. Every
  in-session write is posted from there; the beacon route is the exception the sentence presents as
  the rule.
- collateral, not booked separately: `code/src/graphql/query/academyProgress.js:60-62` records
  *"UPSERT_CHAPTER_PROGRESS_BATCH removed … Progress is now an immediate per-chapter write"*, so the
  sentence's `upsertChapterProgress[Batch]` names a client mutation that no longer exists.

---

## Seat G

### G B1 | `corpus/services/hiring.md:80-81` | UPHELD | IN-SCOPE | `manager.go:485` is `}` and `:448` a blank line; the constructs are at `:535-537` / `:450-453`

- evidence: `stack-demo/app` @ `b948604f`, `internal/organization/manager.go` — `:448` is blank
  (`:450` is `switch org.IsHiring`, `:453` `antRole = enum.RoleCandidate`); `:485` is the closing `}`
  of `forceUserToOrganizationInClerk`. The hard-error is 50 lines further down, in a different
  function: `:535` `if !org.IsHiring {`, `:536` `m.logger.Error("organization is not hiring")`, `:537`
  `return fmt.Errorf("organization is not hiring")`.
- I enumerated the predicate rather than trusting the anchor: `grep -rn "organization is not hiring"
  internal/` returns exactly two non-test production sites — `siminvitationlink.go:63` (the doc's
  `:62` is its `if !org.IsHiring` guard, correct in substance) and `manager.go:536,537`.
- ref-discipline considered and rejected as a shelter. The claim's own block (`:71-88`, list item 1 of
  *The org-type gate*) names **no** ref; the only `app` pin in the file is the header blockquote at
  `:17` (*"RE-GROUNDED — v2.8 M257x iter-23 … `app` @ `5ba17044`"*), 60+ lines away in a different
  section. `platform-alignment.md:906-910` (§5 rule 33) is explicit — *"a pin's scope is the claim's
  own block — a markdown CELL in a table, a wrapped sentence in prose"*. The passage additionally
  carries its own currency assertion (`:85`, *"This passage claimed the opposite until M257x
  iter-52"*), which rule 33's fourth part says cannot be pinned into silence.
- I did verify the anchors were right at `5ba17044` (`:448` = `switch org.IsHiring`, `:485` =
  `if !org.IsHiring`), so this is genuine drift, not a mis-transcription — and the sentence's other
  app anchors (`resolver_cms_queries.go:95,210,258,295`, `resolver_queries.go:1034/1035/1053`) resolve
  identically at both refs, so they neither support nor refute a doc-wide pin.

### G B2 | `corpus/services/hiring.md:302-303` | UPHELD | IN-SCOPE | Nothing gates Workforce Intelligence on `isHiringOrg`; the gate is `isAdmin` + an `apps/hiring` prop

- evidence: `stack-demo/next-web-app` @ `bb3313bc`. I enumerated **every** `isHiringOrg` site in
  `packages/ui/src/NavBar/useNavbarSections.tsx` — `:153` (default), `:165` (type), `:321`, `:326-328`
  (member surfaces), `:339-340` (library trim to `librarySimulationsMenuItem`), `:415` (feedback
  label), `:460` (the "Results" relabel), `:580`. **None** touches the workforce entry.
- the actual gate: `enterpriseWorkforceMenuItem` (`:388-395`, `label: tNavbar('workforceIntelligence')`,
  `key: WORKFORCE_URL`) is included at `:546` as `showWorkforce ? … : null`, inside
  `intelligence: orgVisible.groups.intelligence ? … : []` (`:543-550`).
  `orgSectionVisibility({ isAdmin, showStudio })` returns `intelligence: isAdmin`
  (`packages/ui/src/NavBar/orgGroups.ts:48-64`) — **no `isHiringOrg` parameter at all**.
  `showWorkforce` defaults to `true` (`useNavbarSections.tsx:158`, `NavbarItems.tsx:30`,
  `NavbarLeft.tsx:26`, `NavbarTop.tsx:54`) and is passed `false` in exactly two places, both in
  `apps/hiring` (`.../(verified)/template.tsx:167` and `:248`) — a build-time constant, not the flag.
  `apps/web` never passes it.
- absence checked with more than one instrument: `grep -rn isHiringOrg` over the whole repo (174 hits,
  69 files, node_modules/.next excluded) shows no workforce site, and `grep -rn isHiring` over
  `apps/web/src/app/(authenticated)/(verified)/enterprise/workforce/` returns **0** — the route itself
  carries no `is_hiring` guard either. The bullet sits under `:289` *"What the flip changes"*, so the
  attribution is exactly the error that section exists to prevent.

### G B3 | `corpus/services/hiring.md:38` | UPHELD | IN-SCOPE | The cited "twin" `service_taxonomy.md:52` is the studio-room subprocess correction, not the schema claim

- evidence: `corpus/architecture/service_taxonomy.md:51-52` reads *"`app/internal/cms/studio/studioManager.go:119`
  execs `studio/gen.py`. Same correction as [`dependency_map.md`](./dependency_map.md)'s
  content-generation flow, which had it right all along."* — the Desk→Backend→Room direction
  correction, with nothing about schemas. The claim `hiring.md:30-37` is corroborating (*"No `app`
  migration touches the `jobsimulation` schema at all"*) is at **`service_taxonomy.md:62`**:
  *"**one schema, `public`, owned by `app`** … the `cms`, `jobsimulation` and `skillpath` schemas are
  legacy husks"*.
- the sibling anchor is fine, which is what makes it drift rather than invention:
  `corpus/architecture/dependency_map.md:78` does say *"…or directly to the **`public`** schema (the
  legacy `jobsimulation` schema is non-authoritative)"*. So "the twins" resolves half-right — rule 34's
  signature, a corpus repair that moved a line without re-pointing the citing site.

---

## Deduplication

No two bookings share a predicate. E B1 (a second compose-set service address / a second synchronous
cross-process edge) and F's clearance of the same sentence are the same *anchor* read two ways, not two
anchors on one predicate; I re-derived it and E is right on the `GOTENBERG_URL` clause. G B1, G B2 and
G B3 all land in `hiring.md` but are three distinct predicates (drifted source anchors · a false
mechanism attribution · a drifted intra-corpus citation) at three distinct anchors.

---

**BOOKED=8 UPHELD=7 REJECTED=1 IN-SCOPE-UPHELD-BLOCKERS=7**
