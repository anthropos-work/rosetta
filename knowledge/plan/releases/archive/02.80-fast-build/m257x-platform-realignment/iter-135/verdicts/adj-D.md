# Adjudication D — seats r33-D, r34-D (M257x iter-135, re-adjudicating iter-131)

**Scope line.** I read the adjudicator brief and exactly two seat reports (`iter-131/raw/r33-D.md`,
`iter-131/raw/r34-D.md`). I opened **nothing else** under `knowledge/plan/**` — no other iter dir, no
`progress.md` / `decisions.md` / `adjudication.md`, no other adjudicator's output, no iter-132/133/134
material. Every verdict below rests on evidence I opened myself in `corpus/**` and in the platform
clones at the brief's ground-truth shas. Read-only throughout; this file is my only write.

Two seats, **7 claimed blockers** (r33-D B1–B3, r34-D B1–B4). r33-D B2 ≡ r34-D B2 and r33-D B3 ≡ r34-D B4
(same proposition, same anchor), so the 7 bookings carry **5 distinct propositions**, of which 4 are upheld.

**Corpus baseline used for "true as booked".** The iter-131 read was sealed at `a532493`
(*"probe(M257x/131): the read is SEALED — pre-registration committed before any seat is dealt"*). I
diffed each anchored corpus file `a532493` → `HEAD`:

| file | sealed → HEAD |
|---|---|
| `corpus/architecture/security_compliance.md` | **byte-identical** (diff exit 0) — both anchors still live |
| `corpus/architecture/shared_libraries.md` | **byte-identical** (diff exit 0) — both anchors still live |
| `corpus/services/clerk-integration.md` | changed only at `:107-120`; `:40` / `:44` identical |
| `corpus/architecture/service_taxonomy.md` | `:175` and `:525` **rewritten** (repair `9c86e0f iter(M257x/133)`) |

So exactly one blocker in my set is **since-repaired** (r34-D B1); the rest are still live in the tree.

## Counts table

| | n |
|---|---|
| **claimed** | **7** |
| **UPHELD** | **6** (of which since-repaired: 1) |
| REJECTED — `wrong-tree` | 0 |
| REJECTED — `misread` | 0 |
| REJECTED — **`true-at-its-ref`** | **1** |
| REJECTED — `retraction-not-contradiction` | 0 |
| REJECTED — `minor-not-blocker` | 0 |
| REJECTED — `not-in-scope` | 0 |
| **REJECTED total** | **1** |
| **CANNOT-SETTLE** | **0** |
| **wrong-tree count** | **0** |
| **DISTINCT-PREDICATES-IN-MY-SET** | **4** |

**Upheld rate: 6/7 = 85.7 %** (by distinct proposition: 4/5 = 80 %).

---

## Verdict table

| seat | B# | anchor | verdict | rejection class | predicate (if upheld) | class | multi-pin | repair-induced (sha) |
|---|---|---|---|---|---|---|---|---|
| r33-D | B1 | `corpus/architecture/security_compliance.md:250` (+ `:265`, `:293`) | **UPHELD** (narrowed — see disagreement §1) | — | P-D1 | platform-drift / arithmetic-count | yes | **yes** — `3785b47 iter(M257x/129)` |
| r33-D | B2 | `corpus/architecture/shared_libraries.md:77` | **UPHELD** | — | P-D2 | platform-drift (wrong-construct line anchor) | yes | **yes** — `65921bd iter(M257x/123)` |
| r33-D | B3 | `corpus/architecture/security_compliance.md:156` | **UPHELD** | — | P-D3 | intra-corpus-citation | yes | **yes** — `f723101 fix(M257x/120)` |
| r34-D | B1 | `corpus/architecture/service_taxonomy.md:175` (+ `:525`) | **UPHELD (since-repaired)** | — | P-D4 | platform-drift + self-contradiction | yes | no — pre-read toucher was `f8be5a1 fix(M257x/108)`; the `iter(M257x/133)` touch is the post-read repair |
| r34-D | B2 | `corpus/architecture/shared_libraries.md:77` | **UPHELD** (dup of r33-D B2) | — | P-D2 | platform-drift (wrong-construct line anchor) | yes | **yes** — `65921bd iter(M257x/123)` |
| r34-D | B3 | `corpus/architecture/shared_libraries.md:6` (+ `:38`) | **REJECTED** | **`true-at-its-ref`** | — | — | yes | (yes — `65921bd`, moot) |
| r34-D | B4 | `corpus/architecture/security_compliance.md:156` | **UPHELD** (dup of r33-D B3) | — | P-D3 | intra-corpus-citation | yes | **yes** — `f723101 fix(M257x/120)` |

---

## Upheld predicates, deduplicated within my assignment

**P-D1** | *"`app` mounts exactly SEVEN routes on the root Echo instance outside any group, and only two of
them are open by design."* — measured **eight**, and the eighth serves content to unauthenticated callers.
| anchors: `corpus/architecture/security_compliance.md:250`, `:253-261` (the 7-row table), `:265`
(*"11 groups + 7 ungrouped root mounts"*), `:293-294` (*"seven further routes … two of them open by design"*)
| class: **platform-drift / arithmetic-count**

**Evidence I opened.** At `app` `ad9f3c498` — the ref the passage names — with a positive control in the
same pass (`e.Group(` → 12 lines, matching the doc's own "11 real groups + 1 comment"):

```
git grep -nE '(^|[^.a-zA-Z_])e\.(GET|POST|PUT|DELETE|PATCH|Any|Add|File|Static)\(' ad9f3c49 -- '*.go' | grep -v _test.go
```

returns 11 lines; 3 are in `cmd/labsdemo/main.go` (a separate `main` — the standalone dev server, correctly
excluded), leaving **8** on the app server instance. Seven are the doc's table. The eighth is
`internal/web/backend/labs_admin.go:40` → `e.GET("/v1/labs/:slug/workspace.tar.gz", h.ServeWorkspace)`.
I confirmed the `e` is the same root instance, not a group: `backend.go:301`
`AttachLabsAdminRoutes(e, labs.Catalog, labs.Workspace, apiKeyManager)`, and `web.go:124`
`backend.Attach(srv.e, …)`. `labs_admin.go:36-39` states in the platform's own words:
*"Serve is OUTSIDE the write group — it has OPTIONAL auth (a public Lab's workspace is served to anyone;
a tenant-private Lab requires a key with access)."* So the *"two of them open by design"* clause is also
short: a third route serves public Lab workspaces to any caller.

I widened the search beyond the seat's to make sure eight is not itself an undercount: `.e.<verb>(` → 0
hits; `RouteNotFound` → 0 hits; and every function in the tree taking a `*echo.Echo`
(`academy_embeddings_admin.go:36`, `emailpreview/handler.go:64`, `labs_admin.go:29`) names its parameter
`e`, so the original pattern covers them. `workspace.tar.gz` occurs **0** times in
`security_compliance.md` (positive control: `labs_admin.go` occurs once, at the `/v1/labs` group row) —
the route is nowhere in that document.

**P-D2** | *"the `analytics-go` → Brevo tracking manager is wired at `app/main.go:507-508`."* — at every
readable ref, `:507-508` is the storage-in-app comment block; the wiring is `:494-495`.
| anchors: `corpus/architecture/shared_libraries.md:77`
| class: **platform-drift (wrong-construct line anchor)**

**Evidence I opened.** The cell at `:77` names no ref of its own — the `3eaadae6` pin in the banner at `:6`
scopes *"`app`'s actual org-private module requirements … `app/go.mod:14-18`"*, and the body table
containing `:76`/`:77` starts after the banner closes at `:38-40` — so it grades at the ground-truth
checkout. At `ad9f3c49`, `main.go:507-508` reads:

```
507  // public clients, cmsStorage and jobsimwiring.Wire). Per-namespace clients are
508  // derived at each consumer via internalstorage.NewClient / NewPublicClient.
```

— the middle of the storage-in-app block `:503-518`. The real wiring is
`main.go:494 trackingManager := tracking.New(os.Getenv("BREVO_KEY"))` /
`:495 paymentsManager := payments.New(…, trackingManager)`. I checked the three other refs this corpus
uses: `b948604f:507-508` is a jobsim-in-app `BACKEND_USERS_RPC_ADDR` comment; `9d00a313:507-508` is an
AI-Readiness auto-assign comment; `2035f9a4:507-508` is the same storage comment. The anchor names the
wrong construct at all four. **Everything else in the cell verifies**: `go.mod:14 analytics-go v0.3.1` ✓;
`handler.go:302-316` is the `m.analyticsManager.Track(analytics.Event{…})` call ✓; the switch at
`:283-300` is on `entSub.Status` with exactly **seven** cases ✓ (I enumerated them:
`SubIncomplete/IncompleteExpired/Trialing/Active/PastDue/Canceled/Unpaid`).

**P-D3** | *"the `\"only\"` absolute-quantifier defect this corpus cites as its own precedent lives at
`clerk-integration.md:40`."* — it lives at `:44`; `:40` is the org-invitations bullet and carries no
quantifier at all.
| anchors: `corpus/architecture/security_compliance.md:156`
| class: **intra-corpus-citation**

**Evidence I opened.** `security_compliance.md:153-159` (the iter-120 Layer-2 correction box) reads
*"The same class as the `clerk-integration.md:40` \"only\" and the `cms.md` `bash -c` inversion: an
absolute quantifier over a security surface, published unhedged."* At **both** the sealed state and HEAD
(`:40` byte-identical in both), `corpus/services/clerk-integration.md:40` is
`- **Org invitations** — backend create/revoke (+ a hand-rolled bulk call); frontend list/accept.` —
no *"only"*, no quantifier, no security surface. `grep -n '"only"\|used to say'` over the whole 210-line
file returns exactly two hits: `:44` (*"Sign-in tokens — **five** live minting sites … this bullet used
to say \"only\", and it was false"*) and `:64` (the iter-120 correction box). The intended subject is
`:44`, four lines down.

I weighed `minor-not-blocker` and rejected it: the citation is illustrative, but this same file at
`:282-283` names this exact class as a defect in the corpus's own words (*"a citation that stops one line
short of its own subject is exactly the wrong-construct class `anchor_construct_guard` does not detect"*),
and books it. A corpus cannot grade that class as a defect when it finds it in others and as cosmetic
when it is its own. Held to its own standard, this is a blocker.

**P-D4** | *"the set of private Go modules a stack actually builds against is THREE — colony, proto,
taxonomy."* — it is **five**: `analytics-go`, `colony`, `proto`, `storage`, `taxonomy`.
| anchors: `corpus/architecture/service_taxonomy.md:175`, `:525`
| class: **platform-drift + self-contradiction** (the corpus contradicts itself one file away)

**Evidence I opened.** `git show a532493:corpus/architecture/service_taxonomy.md` — the text as the seat
read it — `:175`: *"the live private-module set a stack builds is **colony, proto, taxonomy**"*; `:525`:
*"**5 libraries, 3 imported by a service a stack builds** — colony, proto, taxonomy"*. Measured at
`app` `ad9f3c49`, `go.mod:14-18`:

```
14  github.com/anthropos-work/analytics-go v0.3.1
15  github.com/anthropos-work/colony       v0.35.2
16  github.com/anthropos-work/proto        v1.210.0
17  github.com/anthropos-work/storage      v0.15.2
18  github.com/anthropos-work/taxonomy     v1.2.0
```

— five, all direct, zero `// indirect`, `go.sum:64-73` carries exactly those five (two lines each), and
the *"pulled at Docker build via `GH_PAT`/`GOPRIVATE`"* mechanism the same sentence describes applies to
all five identically: `Dockerfile:13` and `Dockerfile.dev:13` both set
`ENV GOPRIVATE=github.com/anthropos-work/*`. The sentence's own predicate ("what a stack builds") is
therefore false by two members, and both omitted members are load-bearing (`storage` is a direct require
of the service; `analytics-go` carries the Stripe→Brevo event path).

**Since-repaired.** `HEAD` now reads *"The live private-module set a stack builds is FIVE — `analytics-go`,
`colony`, `proto`, `storage`, `taxonomy`"* with an explicit *"This sentence said \"colony, proto, taxonomy\"
until M257x iter-133"* retraction, and `:525` likewise. The claim was **true as booked** against the sealed
text; I mark it `UPHELD (since-repaired)` and do not reject it for being fixed.

---

## Rejections, with the evidence I opened

**r34-D B3 — `shared_libraries.md:6` is pinned to `app` `3eaadae6`, "a sha that exists in no clone" →
REJECTED, class `true-at-its-ref`.**

The seat's local observations are all correct and I reproduced every one: `git -C stack-demo/app cat-file
-t 3eaadae6` → *"fatal: Not a valid object name"*; `git rev-parse v1.371.1` → unknown revision; the
clone's newest tag is `v1.369.0` and `git describe --tags ad9f3c49` → `v1.369.0-7-gad9f3c498`. But the
seat's *predicate* — that the pin is unresolvable, and therefore the banner's measurement unsupportable —
does not survive the one command that tests it. `git ls-remote origin` (read-only, no local write):

```
3eaadae68e6c5969f7b917574ee3d74d4edf4315	refs/heads/main
3eaadae68e6c5969f7b917574ee3d74d4edf4315	refs/tags/v1.371.1
```

The sha is **real**, it is **`origin/main`**, and it is **tagged `v1.371.1`** exactly as the banner says.
What the seat found is a **clone-set limit, not a measurement limit** — the identical distinction this
corpus itself had to learn at iter-123/124 for `infrastructure`, where *"NOT MEASURABLE from our clone set"*
was retracted the moment the repo was cloned. The brief is explicit that *"a pin is a date, not an excuse —
if the claim is true at its named ref, it is TRUE, however stale"*, and the frozen ground-truth clone being
behind `origin` is a property of the reading, not a defect in the document.

Substantively the banner also survives independent re-derivation at the checkout: at `ad9f3c49`,
`go.mod:14-18` is the five modules the banner tabulates, all direct with zero `// indirect`; the file is
**295** lines and its single `replace` at `:295` is
`github.com/getsentry/sentry-go/echo => …v0.44.1` — *"not an org module"*, exactly as `:7-8` claims. So
every value the banner asserts is byte-true at a ref the seat *could* open, and the ref it names is real
at origin. Nothing here is false.

---

## Cannot-settle

None. All seven blockers settled on evidence I opened.

(For completeness: the seats' own *"what I could not settle"* lists — `db-backup`, `infrastructure`,
`colony`/`proto`/`taxonomy` internals, GitHub archive state, the AI-Labs repo — carry no blockers in my
assignment, so nothing there required a verdict from me. I did not book them and I did not launder them.)

---

## Disagreements with how the seats framed their predicates

**1. r33-D B1 — I narrow the predicate by one sub-claim.** The seat books three defects in one blocker:
(i) the ungrouped count is 7 and should be 8; (ii) *"two of them open by design"* should be three; and
(iii) *"its `/v1/labs` row asserts a group-level API-key + scope gate that this route does not carry."*
**(iii) does not stand.** The row at `security_compliance.md:212` names its subject precisely — *"`/v1/labs`
| `internal/web/backend/labs_admin.go:31` (mounted `backend.go:301`)"* — and `labs_admin.go:31` really is
`g := e.Group("/v1/labs", apiKeyAuthMiddleware(apiKeys, "labs:write"))`. The row is a true statement about
the **group**; the serve route is a **sibling on the root**, not a member of that group. Booking (iii)
would make the repair wrong: the fix is to add an eighth row to the root-mount table (and correct
*"two … open by design"*), **not** to weaken an accurate group row. I uphold P-D1 on (i)+(ii) only.

**2. r34-D B3 — the framing is the error, not the observation.** *"Exists in no clone"* was framed as
*"unsupportable"*. Those are different predicates: the first is about the reading's frozen ground truth,
the second about the corpus. `git ls-remote` separates them and clears the corpus. This is the highest-value
disagreement in my set, and it is sharpened by an **intra-seat contradiction**: seat D booked this as a
BLOCKER in reading #34 while, in reading #33 over the *same* file set, it explicitly declined to book it
(*"§5 rule 49 forbids refuting another observer's report of a concurrently-writable surface with my own —
here older — snapshot"*) and named the settling command verbatim: *"one `git ls-remote origin
refs/tags/v1.371.1` between readings."* I ran it. It settles in the corpus's favour. **Reading #33 was
right and reading #34 was wrong, on identical evidence** — a same-seat test-retest divergence, not a
corpus defect.

**3. The reciprocal miss, in the other direction.** Reading #33 **did not book** the largest platform-drift
defect in my set (P-D4, `service_taxonomy.md:175`/`:525`) even though it read that file in full (530 lines,
declared in its own positive-control table) *and* had the refuting measurement in hand — its "positively
cleared" section states *"`app`'s org-private requires at `ad9f3c49` are analytics-go `:14`, colony `:15`,
proto `:16`, storage `:17`, taxonomy `:18`"*. The contradiction with `:175`'s *"colony, proto, taxonomy"*
sat between two sections of one report. So across the pair: #34 booked one thing #33 correctly refused,
and #33 missed one thing #34 correctly caught — the two readings' disagreements are **symmetric**, and
neither is attributable to a different tree, ref, or search failure. Both seats had the same evidence.

**4. A framing note on P-D2 that neither seat drew.** Both seats correctly treat the `:77` cell as
ungoverned by the banner's `3eaadae6` pin. I checked the structural basis rather than assuming it: the
banner block runs `:3-38` and closes at `:40` (*"This document covers **five** internal library repos"*),
and its stated scope is `app/go.mod:14-18`. The `:76`/`:77` rows sit in a different, later table. So the
cell grades at the checkout, and r33-D's hedge to *medium* confidence (*"the anchor may well resolve at
`3eaadae6`"*) was unnecessary caution — the storage-in-app comment block occupies `:503-518` continuously
across `2035f9a4` and `ad9f3c49`, so a jump of the tracking wiring from `:494` to `:507` in the handful of
commits to `3eaadae6` would require ~13 lines of insertion in exactly the wrong place. r34-D's *high*
confidence is the better-calibrated of the two on the same evidence.

---

## Counts

```
UPHELD=6 REJECTED=1 (of which wrong-tree=0) CANNOT-SETTLE=0
DISTINCT-PREDICATES-IN-MY-SET=4
```
