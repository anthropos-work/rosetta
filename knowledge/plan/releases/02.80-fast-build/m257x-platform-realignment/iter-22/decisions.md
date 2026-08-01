# iter-22 — decisions

## D-M257x-22-1 — two of the hand-off's 21 were REFUTED, and the refutation named a 22nd

**The hand-off said** (items #8, #10) that `services/jobsimulation.md:82` and `services/messenger.md:108`
were blockers because they carry `JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401`, and that the correction
was `http://backend:8083`.

**Measured against platform origin HEAD `2adcf71` before applying:**

    docker-compose.yml:52   JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401   (backend)
    docker-compose.yml:258  JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401   (messenger)
    docker-compose.yml:256  CMS_RPC_ADDR=http://cms:8091                       (messenger)
    docker-compose.yml:265  SKILLER_RPC_ADDR=http://backend:8083               (messenger)

Only `SKILLER_RPC_ADDR` was re-pointed. **The corpus was right and the correction was wrong** — applying it
would have replaced a true statement with a false one.

And it is not compose lag. `app/main.go:1196-1202` @ `5ba17044`:

> *cms-in-app M807: the in-app CMSService edge … **Additive + DORMANT**: external callers (messenger) keep
> hitting the standalone cms via `CMS_RPC_ADDR` **until the M809 re-point***

So the husk addresses are **deliberate, current, and load-bearing** until platform M809.

**Where the hand-off's error came from — and it is the 22nd blocker.** `services/backend.md:175` asserts
messenger *"points `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR` and `SKILLER_RPC_ADDR`
**all** at `http://backend:8083` locally."* Two of the four. That sentence is the refuting citation iter-21
trusted, so one false corpus line propagated into two false corrections in a hand-off authored to be
executed without re-derivation.

**Resolution:** #8 and #10 are **not** blockers and were not "corrected" — they were *annotated* with the
verification (`verified current at 2adcf71`, points at the husk, re-point is M809/M810) so the next reader
does not re-file them. `backend.md:175` was corrected, cited to `docker-compose.yml:255-265` +
`app/main.go:1199`. Net: 21 → 19 real blockers + 1 net-new = **20**.

**Lesson, and it generalizes past this milestone:** *re-derive the CORRECTION, not just the anchor.* An
enumerated hand-off makes the anchors cheap to verify and the corrections easy to trust — and the anchors
were all fine. The failure mode a mechanical hand-off invites is not a moved anchor; it is an inherited
falsehood wearing a `file:line` citation. Filed to `platform-alignment.md` §5.

## D-M257x-22-2 — "not in the local compose" was false in both merge banners, and it is the same shape as the class this milestone exists for

`cms.md:6-7` and `jobsimulation.md:7-8` both stated the merged service is *"not in the local compose."*
`docker-compose.yml:144` and `:83` @ `2adcf71` both still define the service **in the default `graphql`
profile**. Both sentences arrived from `main` in iter-21's merge — not ours, not regressions, but false.

The distinction that was collapsed: **merged-in-production is not removed-from-compose.** A service can be
`service_desired_count = 0` in prod terraform, folded into `app`, its subgraph gone — and still start a
container locally on every `make up`. The map already had a word for it (`running_but_unfederated`); the
service docs did not use it.

This matters beyond wording. The whole M257x class is *tooling writing to a schema the platform no longer
creates*. A reader who believes the container is gone will not think to ask **which of the two** answers an
RPC, and both containers are still reachable on the compose network. Corrected to name the husk, its
`file:line`, its state, and its M809/M810 teardown.

## D-M257x-22-3 — the router edges in the mermaid, and why a diagram is the last thing a term sweep reaches

`architecture_overview.md`'s topology diagram routed `Web --> GraphQL --> Gateway` as live topology, four
releases after `2adcf71` deleted the router from local dev. Three prior audits in this milestone missed it
because no drift *term* appears in a mermaid edge — `Web --> GraphQL` contains neither "router", nor "5050",
nor "cosmo".

Corrected by splitting the edge rather than deleting it: the router edges become `-.->|prod only|`, and
net-web / hiring / studio-desk gain solid `-->|local: :8082/graphql/query|` edges to the monolith. The
diagram now shows both columns, which is what the Request Flow section (:239) needed too — it had one flow
where the platform has two.

## D-M257x-22-4 — the roadrunner row: precision, not correction

`roadrunner.md:9-10` said *"there is no roadrunner service in production."* The map says `live-standalone` in
prod — `roadrunner/terraform/main.tf:19` still reads `service_desired_count = 1`, untouched since `87d8d44`
(2026-06-19, before the fold), while `repos.yml:29-31` says *"legacy — folded into app."*

Rather than pick a side, the banner now says **orphaned, not absent** and records that this is the one row
where prod and the platform's own declaration contradict each other. Deciding which is authoritative would
require a platform answer we do not have and are not permitted to force. Recording the contradiction where a
reader meets it is the honest move; the map already carries it as such.

## D-M257x-22-5 — the same drift is in `corpus/ops/**`, and it is OUT of clause 5 (routed, not fixed)

Measured while verifying the corrections, **not** fixed — clause 5's scope is exactly
`corpus/services/**` + `corpus/architecture/**`, and this is a 3rd line of investigation (scope-creep
tripwire). **13 occurrences** of the dead router in `corpus/ops/**`:

| file:line | what |
|---|---|
| `run_guide.md:18` | *"GraphQL Playground \| http://localhost:5050 \| API gateway (Cosmo Router)"* — the service table |
| `run_guide.md:125` | `curl -s http://localhost:5050/health` as a health check |
| `setup_guide.md:528` | *"already default to `localhost:5050` … which are correct"* |
| `platform_repo.md:88,139` | *"no usable `:5050` endpoint"* · `open http://localhost:5050` |
| `update_guide.md:223` | the same dead health curl |
| `verification.md:161` | GraphQL introspection at `:5050+offset` |
| `staging-clerk.md:96,132` · `staging-bringup.md:44,208,404,511` | staging URLs (**arguably fine** — staging may still front a router; needs a separate call) |

**This matters more than its scope suggests.** `run_guide.md` and `setup_guide.md` are the onboarding path —
`CLAUDE.md` points a new engineer (or agent) straight at them, and `/dev-up` executes them. A dead health-check
`curl` there fails *silently into a warning* rather than loudly, which is the exact shape of the
`|| echo 0` defect that stopped M257.

Routed forward: **`DOC-M257x-iter22-ops-guides-5050`**. The staging rows need a prior question answered (does
staging still run a router?) that local compose cannot settle — split them out rather than bulk-editing.
