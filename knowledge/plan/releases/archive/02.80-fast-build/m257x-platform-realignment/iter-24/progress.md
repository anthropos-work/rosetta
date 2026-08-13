**Type:** tik

# iter-24 — the per-stack Directus was serving a consumer that no longer reads

## The diagnosis, and why it needed three probes rather than one

iter-23 routed this as a *candidate* cause. The re-survey upgraded it to a proven one, and the upgrade is the
methodologically interesting part: **one probe could not have done it.**

| probe | result |
|---|---|
| `backend`'s log on the standing `demo-1` | **96** Directus lines, all `403 FORBIDDEN` — `directus_versions` (killing `publicSkillPaths`, `getSkillPath`, `getOrCreateSkillPathSession`) and `library_categories` (killing `libraryCategories`) |
| the **local** per-stack Directus, anonymous | `library_categories` **200** · `skill_paths` **200** · `task_sub_checks` **200** · `/versions` **200** |
| **prod** `content.anthropos.work`, anonymous | `library_categories` **403** · `/versions` **403** |

The local instance answers the exact collection `backend` is refused, and **`backend`'s 403 set matches
*prod's* answers**. So this is not a grant missed on the replay — it is a client pointed at the wrong server.
§5 rule 7: *a probe must not be able to satisfy itself.* iter-18's probe asked "does the local Directus
serve?" and correctly answered yes; that question cannot distinguish these two worlds. Asking **both ends**
can, and the asymmetry is the whole finding.

This also explains iter-19's result rather than contradicting it: the 403 **is** independent of iter-18's
*serving* defect — because it is a **pointing** defect, one layer up.

## The mechanism

`DIRECTUS_DATA_CONSUMERS = ("cms",)`, in **both** twins. Correct at M23 (v1.5), when `cms` was the only
platform service that spoke to Directus. cms-in-app moved the reader: `app/cms_reader_switch.go` swaps
`backend`'s cms content reader to the **in-process** cms server once Directus is configured — *"a DIRECT
domain call — no proto round-trip … and no internal traffic to a standalone cms"* — and `app/main.go:971-973`
`log.Fatalf`s without `DIRECTUS_BASE_ADDR`.

**Nothing errored, and that is the point.** A stale *schema* name fails loudly at 42P01. A stale **service
name in a consumer list** fails silently: the list still names a real, running container that still starts and
still holds the variable. The read simply happens somewhere else. Promoted to protocol §5 at iter-23 with its
one-command check (`docker inspect <container> | grep <VAR>`).

`cms` **stays** in the list — its container is still started by the default `graphql` profile as a
merged-into-`app` husk (rollback path until platform M810) and messenger still addresses it over
`CMS_RPC_ADDR`. Dropping it would re-create the same split one service over.

## The tests were arguing for the defect — third occurrence this milestone

`test_only_cms_is_repointed_not_other_services` asserted `backend` **must not** carry the re-point. True
pre-fold; then a test pinning the defect as a contract (§8 rule 3), which would have failed on the fix.
Replaced by three tests that split the proposition — backend IS re-pointed · cms STAYS re-pointed ·
non-readers are not — so a future fold moves one assertion, not a conjunction.

**And the dev twin's fixture could not name its own subject.** It called the app service `"app"`; the platform
compose service is **`backend`** (stack-injection's own map spells it `{"backend": "app"}` — compose service →
repo dir). So the dev consumer list could have been wrong about `backend` and every dev test would still have
passed. `backend` added to the fixture; `"app"` kept as a genuine non-consumer so the negative assertions keep
a subject. Both twins fixed in one pass — §5 rule 9, the rule iter-12 paid a cold cycle for.

## Mutation battery — 6/6, with TWO declared-GREEN controls

`.agentspace/scratch/work-m257x/iter24/mutate24.py`. Pre-control and post-control green on both target suites;
every mutant restored from the in-memory original, never from git.

| mutant | expected | got |
|---|---|---|
| M1 demo drops `backend` | RED | RED |
| M2 **dev** drops `backend` | RED | RED |
| M3 demo drops `cms` | RED | RED |
| M4 demo over-reaches to `jobsimulation` | RED | RED |
| **M5 tuple reorder (demo) — CONTROL** | **GREEN** | **GREEN** |
| **M6 tuple reorder (dev) — CONTROL** | **GREEN** | **GREEN** |

The reorder is semantically identical to a membership test, so a suite that reddens on it is asserting the
literal rather than the proposition. M2 is the one that matters most: without the fixture fix it would have
survived, and four REDs would have "proven" a fence that never looked at the dev twin.

## Live proof — measured, not inferred

rext committed, tagged **`fast-build-m257x-iter-24` (`f9ac72f`)**, pushed, **verified on origin**
(`git ls-remote --tags`); `.agentspace/rext.tag` + the `stack-demo` consumption clone both re-pinned to it.

Regenerated `demo-1`'s override **from the pinned consumption clone's generator** — the emitted `backend:`
block now carries `DIRECTUS_BASE_ADDR=http://directus:8055`, and a `diff` against the pre-fix file shows
**that single added line** (plus the roster flags this invocation did not pass). Recreated only the `backend`
container (§5 rule 15 — diagnose before paying for a full cold cycle):

```
before: 96 Directus log lines, every one 403
after : 0
libraryCategories                    -> 8+ real categories from the replayed catalog
publicSkillPaths(options:{limit:3})  -> 3 real skill paths, by title
```

Both are the exact fields that threw `Directus FORBIDDEN 403` minutes earlier. This closes
**`FIX-M257x-iter15-directus-versions-403`** (58 occurrences, carried since iter-15) at its real mechanism —
which was never `directus_versions` permissions at all.

## Close — 2026-08-01

**Outcome:** the per-stack Directus is now wired to the service that actually reads it; `backend`'s 403 class
goes **96 → 0** and two previously-dead content queries return real replayed data, proven live on `demo-1`
through the pinned consumption clone.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n (platform origin `2adcf71`
re-checked at open and close, unchanged) — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n —
Outcome: continue
**Decisions:** D-M257x-24-1 … D-M257x-24-3 (this iter's `decisions.md`)
**Side-deliverables:** none — all four planned phases landed.
**Routes carried forward:**
- **The clause-2 re-measure** (next tik) — the Playthrough suite with a real `--reset` on this stack. The
  number is **deliberately not predicted**: iter-19's `diff` proved the ten failures span at least four
  causes and this fixes one. Predicting the lift here would be the attribution error this milestone has
  already made once.
- **`FIX-M257x-iter19-playthrough-runner-path`** — do it *before* the re-measure, not during: the runner's
  bare `stackseed` is not on PATH, so the suite cannot reset itself, and hand-supplying the path mid-measure
  is how iter-15 ended up comparing two different worlds.
- `DOC-M257x-iter23-rext-stale-session-comment` — still open; batch it with the next rext change.

**Lessons:**
1. **Ask both ends.** A one-sided probe ("does the server serve?") is satisfiable by a broken system. The
   asymmetry between *what the local instance answers* and *what the client is refused* is what identified
   the defect, and neither half alone says anything.
2. **A test fixture that cannot name the subject cannot test it.** The dev twin's `"app"`-vs-`backend`
   mismatch made an entire class of defect invisible on that side — and the mutation battery would have
   certified the fence anyway, on the strength of the demo twin's REDs.
3. **A stale service NAME fails silently where a stale schema name fails loudly.** 42P01 is a gift. A
   consumer list that names a container which still exists gives you nothing.
