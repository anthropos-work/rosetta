# iter-18 — decisions

## D88 — the 8th merge is `mid-fold`, not `merged-into-app`, and the reason is the schema

**Measured**, platform `766df6c` / `app` `c52dbc51e` / `sentinel` `f2c46190`, all three verified equal
to `origin/main` **at the remote** (`git ls-remote`, 2026-08-12T12:07Z) — not merely to a local ref.

`app/internal/sentinel/` is the Casbin PDP, ported from the standalone repo at tag **v0.24.2**
(`app/internal/sentinel/doc.go:10`; `f2c46190` is that repo's `origin/main` today, so port source and
repo head are the same commit). Wired **once**, at `app/main.go:305`, `log.Fatalf` on failure, under
the source's own *"There is no switch and no RPC path: app IS the PDP."*

The map's §1 defines `merged-into-app` as **all three** of: app owns the code and calls it
unconditionally · the tables live in `public` · the standalone is scaled to zero. Only the first
holds:

- **The tables did NOT move.** `sentinel.casbin_rules` is still in the `sentinel` schema, reached via
  `SENTINEL_DB_CONNECTION` (`docker-compose.yml:25`, `search_path=sentinel`) and migrated by `app`'s
  own `make migrations-sentinel` (`app/Makefile:80-81`, `atlas migrate diff --env sentinel`). **This
  is the first of the eight folds where that is true**, and it is the whole reason the row cannot say
  `merged-into-app`.
- **The standalone is not at zero.** `sentinel/terraform/main.tf:19` `service_desired_count = 1`, and
  unlike the `cms`/`messenger`/`roadrunner` class that count **is instantiated**: `module
  "sentinel_euwest1"` is one of the ten declarations in `infrastructure/terraform/production/services.tf`
  at `13c248e6` (2026-08-07). ⚠️ **Not re-measured here** — `infrastructure` is in no clone set and
  assertion F reports it unclonable. What *is* measured in-tree is `app`'s own pin of the same
  transition (`app/sentinel_wiring_test.go:57`, `TestNoRPCPathSurvives`). The platform names the
  teardown **M1103** at `docker-compose.yml:85`.

`mid-fold` had **no holder** since `storage`. The token existed precisely so this state would not have
to be rounded to one of its neighbours.

## D89 — the RPC edge is not gone-and-replaced; the whole listener is gone

The corpus said, in **eight** places, that `backend → sentinel` was *"the only cross-process
Connect-RPC edge left in a local stack."* It is not re-pointed — it does not exist:

- `AUTHORIZATION_ADDRESS` occurs **0** times across `docker-compose.yml`, `common.yml` and `repos.yml`
  at `766df6c`; in `app`'s whole Go tree it occurs **once**, inside a test asserting its absence.
- `app` deleted its Connect-RPC server (`app/main.go:1310`, *"NO RPC SERVER"*) — the port-8081 mux
  that carried Users / Organizations / Skiller / JobSimulation / CMS / lab. **`app/rpc.go` no longer
  exists** either (`a85e8308d`, *"delete rpc.go"*; pinned by `app/rpc_removal_test.go`).
- So a `core` stack now has **no cross-process Connect-RPC edge at all** — the count went 1 → 0, and
  every doc that carried the qualifier *"it is not the only cross-process **edge**"* keeps its point:
  `backend → gotenberg` is plain HTTP and Judge0 is a direct URL.

**Live confirmation, measured on `demo-3`'s own compose network** (a throwaway curl container, image
removed afterwards): `backend:8081` → no listener · `backend:8083` → no listener · `backend:8084` →
**404 (alive)**. Compose publishes 8081/8082/8083 and does **not** publish 8084, so **two published
ports bind nothing and the one live extra surface (the meta server) is unreachable from the host.**
`RPC_PORT=8083` (`docker-compose.yml:46`) is dead config.

## D90 — quoting a retracted citation RE-ARMS it, and can silently launder the drift it retracts

Found by doing it. Writing *"this used to say `docker-compose.yml:102-103`"* — backticked, in the
`path:line` form — makes `platform_alignment_guard`'s citation parser read it as a **live** citation.
Worse: a retraction almost always also names the block the stale line drifted **into** ("…it pointed
into `studio-desk`'s block"), which satisfies assertion F's **declared-cross-block escape**. The fence
then goes green *over the retracted citations*.

Measured on this very edit: the guard reached `rc=0` while three retracted citations were still being
graded, and `anchor_construct_guard` independently caught the same shape in
`service_taxonomy.md`. **Rule adopted and written into the map: a retracted citation is spelled
without the `path:line` form** (`` `repos.yml` `` lines 14-17). The rule is in the map's own banner so
the next editor meets it before repeating it.

## D91 — the fence caught the departure exactly as designed, and had simply never been run since

`platform_alignment_guard` against `766df6c`: **`rc=1`, 17 findings** — one `[B departure]` (*"the map
claims sentinel is in repos.yml, and it is not"*) plus **16 citation failures**. Repaired to **`rc=0`**.

Two things worth separating. **The membership arm did its job**: direction B is the arm that has fired
in anger every time, and it fired again. **The citation arm did the harder job**: 8 of the 16 were
citations that still land *inside* the file (`repos.yml` 28→13 lines, `docker-compose.yml` 190→164) but
in **another service's block** — sentinel's citations landed inside `backend`, studio-desk's inside
`next-web-app`, next-web-app's inside `gotenberg`. A human reader cannot see that class at all.

**Read exit codes directly, never through a pipe** (the release's standing rule, and it mattered here:
a first reading of `rc` through `tail` would have reported this fence as passing).

## D92 — `make bootstrap-dev` is broken in the platform, and we report rather than fix

`platform/Makefile:148-149` hard-requires `../sentinel/init_policy.sql`; `:164` runs `docker compose
restart sentinel`; `:165` waits on *"sentinel RPC on localhost:8087"*. None of the three exists at
`766df6c` — `make init` does not clone the repo, and there is no such service. The target fails at its
own guard with *"Run 'make init' first"*, advice that cannot help.

Recorded in `platform_repo.md` and `services/sentinel.md`. **Not fixed: 0 platform-repo edits.**

## D93 — the invalidation channel is the fact that explains iter-16, and it was documented nowhere

`sentinel:policy:invalidate` (`app/internal/sentinel/watcher.go:55`) — Redis **Pub/Sub**, chosen over
app's Watermill consumer-group plumbing because invalidation must **fan out** and an XREADGROUP `>`
entry goes to exactly one consumer.

This is the mechanism behind M258 iter-16's 15-red set, and the corpus had no page that would have
predicted it. Any tool that writes policy rows in **raw SQL** — as our seeders do — bypasses casbin's
write path, publishes nothing, and leaves every replica serving its boot policy: `forbidden` at HTTP
200. Written into `services/sentinel.md` as a standing warning to Rosetta's own tooling, not just as
platform trivia.

## D94 — `docker-desktop-vm` is a G1 FALSE POSITIVE, and the fence is what needs fixing

`platform_predicate_guard` stays **`rc=1`** on one finding: *"`docker-desktop-vm` is documented as a
profile at 2 site(s) but no service declares it"* (`build-budget.md:197`, `:551`). It is **not** a
compose profile — it is a **host** profile, the `hostprofiles/*.json` concept M255 introduced. G1's
noun-phrase detector (`_PROSE_PROFILE`) has three discriminators — negation, postfix negation, ref-pin
— and **no domain discriminator**, so every *"a `X` profile"* in the corpus is read as a compose token.

**Pre-existing** (both sites are outside this iter's diff; the only hunk in that file is line 101).
**Deliberately NOT fixed here** (`ROUTE-M258-iter18-g1-reads-host-profiles-as-compose-profiles`): the
repair is an rext change to a guard's detector, and loosening a detector without RED-proof mutants is
the *"a capability probe that fails OPEN disarms the check it guards"* failure this release has already
paid for once. It needs its own tik with tests, a tag and a push. **The prose is correct and was not
edited to make the fence green** — which is this release's rule, applied in the direction that costs
something.

## D95 — what was corrected, and what was deliberately left alone

The census is **71 files / ~410 mentions**, and most of those mentions are still **true**: the
`sentinel` *schema*, the `sentinel.casbin_rules` table, the Casbin model, and the historical narrative
all survive the fold. The sweep therefore targeted the **structural** claims — the ones that describe a
topology the platform no longer has — in six mechanical classes: the always-on floor (3 → 2), the
`core` container count (5 → 4), `repos.yml` membership (4 → 3), the cross-process RPC edge (1 → 0), the
Tier-1 Go-service list (2 → 1), and the service doc's own shape.

**26 files edited.** Ten fences green. What is left is prose that reads as history and is marked as
history — not a residue of wrong claims.


## D96 — attribute a suite result with a pristine extract, never with a straddling run

The `stack-core` suite takes ~31 min. The run started as a "baseline" **straddled the edits** — tests
that ran early read the old corpus, tests that ran late read the new one — so its 58 failures are not
evidence about either tree. That is a methodology error, and it is recorded rather than quietly
re-run, because it is easy to make and invisible in the output.

The fix was a **pristine `git archive HEAD` extract** in the scratchpad (the technique iter-13 used),
with `stack-demo` / `stack-dev` symlinked and the rext clone **copied** rather than symlinked — a
symlink makes `Path(__file__).parents[4]` resolve back to the live tree and the "pristine" run silently
grades the edited corpus. Then the same 18 modules were run on both trees:

| tree | failures |
|---|---|
| pristine HEAD | **46** |
| post-edit | **47** |
| **delta introduced** | **1**, named below |

**Positive control, and it is the important half:** the pristine extract reproduces
`platform_alignment_guard` at **`rc=1`, 17 findings, `[B departure]` included** — identical to the
reading that opened this iter. So the RED was pre-existing, the repair is the delta, and the substrate
is valid. A comparison without that control would prove nothing. Extract deleted afterwards (459 MB,
host tree, not Docker).

## D97 — the fence's negative control named a live service, and expired with it

The one introduced failure:
`test_service_doc_status_fence::test_the_banner_detector_can_actually_say_no`. Its job is §5 rule 7 —
*the probe must not be able to satisfy itself* — and it did it by running `has_banner()` against a doc
that should NOT read as a banner. The specimen was hard-coded `sentinel.md`, with the premise in its
own docstring: *"the map calls `sentinel` live-standalone, and it has no banner."*

**Both halves expired on the same day.** `766df6c` folded sentinel; this iter moved the map row and
gave the doc the banner the fence demands. The control then failed **while the detector was working
perfectly** — it said "banner" about a doc that has one.

The lesson is the module's own: *"what CAN be fenced is a derivation."* The control's specimen is now
derived from the same map the fence already reads (`live_services()`, the complement of
`gone_services()`), with an anti-vacuity floor — if no live-and-unbannered doc can be derived, the
control says it has no specimen rather than passing. **RED-proven**: mutating `has_banner` to
`return True` kills it (2 failed); unmutated it is 4/4. Five specimens derive today (`next-web-app`,
`studio-desk`, `db-backup`, `ant-academy`, `gotenberg`).

Shipped as `fast-build-m258-iter-18`, **verified on origin** with `git ls-remote` (rung zero), and
`.agentspace/rext.tag` re-pinned to it in the same iter — the `D71` half-re-pin class, not repeated.

⚠️ **`sentinel` is still in `live_services()`** and that is correct: its **prod** cell is `mid-fold`,
which is not a gone-state, and the fence is deliberately one-way (a banner on a live row is not a
finding — the `roadrunner` precedent). It is excluded from the control by carrying a banner, which is
the filter doing its job rather than a special case.
