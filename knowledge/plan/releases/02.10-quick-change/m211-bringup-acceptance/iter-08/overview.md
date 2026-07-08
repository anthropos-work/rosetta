---
iter: 08
milestone: M211
iteration_type: tik
iter_shape: bring-up
status: planned
created: 2026-07-08
---

# iter-08 — tik: cold `/demo-up` GREEN on the merged platform

**Active strategy reference:** TOK-01 "Warm-first cache-migrate, then cold-prove both stacks" — move (4)
"prove COLD". The warm inner loop is green (5/6, iters 02-06); now the first COLD proof: a full `/demo-up`
from a torn-down state on the merged platform with the re-grounded tooling.

**Step 0 — Re-survey.** iter-06/07 route: run the full COLD `/demo-up` (the well-automated path — its
`migrate-demo.sh` already bootstraps extensions+schemas+casbin+PG-wait, and it's the UI-tier substrate for
(e) coverage + Playthroughs). Pre-flight confirmed (this iter's open): stack-demo/platform + stack-demo/app
are at MERGED code (skiller service/repo removed, skiller-in-app merge landed); no live demo containers;
secrets dir present (5 entries); Docker RAM ~9.7 GB (below the 12 GB UI recommendation — non-fatal warn,
freed by the warm-dev teardown); the rext consumption clone `stack-demo/rosetta-extensions` is STALE at
v1.10.1 (pin = quick-change-m209) and the tag is local-only (not pushed).

**Cluster / target identified:** cold `/demo-up` (headline gate item + (e) substrate). Prereq: re-pin
`stack-demo/rosetta-extensions` → quick-change-m209 via a LOCAL fetch from the authoring copy (the tag isn't
on GitHub; local never-pushed tag ops are allowed), then tear down the warm dev (ONE stack at a time), then
`up-injected.sh 1` reap-safe (detached process + in-turn polling).

**Hypothesis:** the merged platform stands up cold via the re-grounded demo tooling — 4-subgraph compose /
no skiller container (a), cache-first replay loads public.* ~42,790 (b), Stories & Heroes seed closure green
(c), auto-verify merged-assertion passes (d), 0 skiller residue (f) — all COLD, proving the warm results
weren't warm-state artifacts.

**Expected lift:** the cold `/demo-up` GREEN flips the demo half of the "both stacks GREEN cold" headline +
proves a-d,f COLD on the demo. Metric: 5/6 → cold-demo-proven (the (e) sub-condition still needs the
coverage+Playthroughs sweeps in tik-09/10 against this live demo).

**Phase plan:** re-pin consumption clone → tear down warm dev → launch `up-injected.sh 1` detached → poll
in-turn (reap-safe, heartbeats) → on DONE, read the auto-verify tail + confirm a-d,f cold → `rosetta-demo
status`.

**Escalation conditions:** if the cold build OOMs on the ~9.7 GB Docker VM (below the 12 GB rec) and can't be
made green without raising the VM (a user-machine setting I can't change) → surface as a resource blocker
(user-blocker). If a surface needs a platform-repo edit → escalate `unimplementable-without-platform-edit`.

**Acceptable close-no-lift outcomes:** if the cold build surfaces a tooling fix-loop (route to rext/corpus,
re-measure) that can't complete this iter → land what's ready, route the rest forward, close on planned-scope
outcome.
