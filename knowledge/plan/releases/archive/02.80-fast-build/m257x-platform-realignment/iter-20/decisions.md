# iter-20 decisions

## D-M257x-20-1 — the map carries SEVEN states, not the gate's four

The exit gate names `{live-standalone, merged-into-app, decommissioned, net-new}`. The protocol
(`platform-alignment.md` §6) names seven, adding `running_but_unfederated`, `external` and `library`. The map
uses the seven, and the fence enforces the seven.

**Why the superset is the right call and not scope creep:** three rows at origin HEAD are *only* expressible
with the extra states, and collapsing them would make the map wrong in exactly the way this milestone exists
to fix. `cms` and `jobsimulation` are `merged-into-app` in production **and** still start as containers in a
fresh local stack — that is `running_but_unfederated`, and calling it either "merged" or "live" is a false
claim on one side. `colony`/`proto`/`ai`/`authn`/`taxonomy` are not services at all; forcing them into
`live-standalone` would put five non-processes on the traffic path.

## D-M257x-20-2 — `roadrunner`'s contradiction is RECORDED, not resolved

`repos.yml:29-31` says roadrunner is *"legacy — folded into app; backend calls Judge0 directly"*, and
`roadrunner/terraform/main.tf:19` still reads `service_desired_count = 1` — a file untouched since
`87d8d44` (2026-06-19), before the fold. From this box, whether that terraform is *applied* is unknowable.

The map therefore writes `prod = live-standalone` with the contradiction spelled out in the evidence cell,
rather than picking the reading that makes the table tidy. **Trap C in the protocol is precisely that the
platform's declarations lag its code** — a map that silently resolves a contradiction in favour of the
prettier row is how the corpus got wrong about skillpath for a release.

## D-M257x-20-3 — the fence has no default `repos.yml` path

`PLATFORM_REPOS_YML` is required; the guard exits **2** (not 0, not 1) when it is absent. Defaulting to
`stack-dev/platform/repos.yml` would have been convenient and would have reproduced the milestone's founding
defect: **this milestone exists because a stale local `repos.yml` was read as ground truth.** §4 Trap A in one
line — a fidelity check against the wrong reference passes.

Same reasoning gives the guard three exit codes rather than two: an unreadable input, an unparseable
`repos.yml` (0 repos), a missing fence marker, and an empty fenced table all raise rather than report OK. An
empty table makes every "every row must…" assertion vacuously true, which is the shape of the 43 checks M256
found reporting success without checking.

## D-M257x-20-4 — completeness is derived from git history, not from memory

"Every service the platform has ever had" is unfalsifiable as a promise, so the map states the two commands
that generate its row set (`git log -p --follow` over `repos.yml` and over `docker-compose.yml`) and invites
the reader to re-run them. Running them found **five services no one on our side has ever named** —
`nats`, `web-app`, `chromedp`, `simulator`, `realtime` — including `simulator`, which `84862d1` (2024-05-29)
replaced with `jobsimulations`: the first ancestor of `app/internal/jobsimulation/`.

Written from memory, this table would have had 24 rows and looked complete.

## D-M257x-20-5 — the GitHub `archived` flag is promoted into the protocol's signal 6

Measured this iter: `jobsimulation` and `skillpath` archived 2026-07-31, `graphql-wundergraph` 2026-07-30,
`skiller` 2026-07-01 — each within days of its fold. One API field answers a question that otherwise costs a
terraform read, a compose read and a code read.

It is a **confirmation, never a precondition** — `cms` is folded and unarchived, `roadrunner` is declared
folded and unarchived. And it cuts the other way: `chronos` is **not** archived while the corpus called it
archived, a corpus error none of the other five signals can see. Recorded in `platform-alignment.md` §4,
including the `curl` + `GH_PAT` form for a box with no `gh` (this one).
