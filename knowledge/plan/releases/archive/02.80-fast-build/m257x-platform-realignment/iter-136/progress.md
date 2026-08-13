**Type:** tik

# iter-136 — the eighth root mount, verified against the adjudicator who found it

Takes **two** items off iter-135's work list, **selected by consequence** per `TOK-08`'s carried
finding (*grade by consequence, not by class*), and leaves the rest routed rather than half-done.

---

## 1. Re-derived at source, because the reporter had already been wrong once

iter-135's `adj-F` upheld a seat's *number* while **refuting its diagnosis**: the seat proposed
`/ai-readiness/unsubscribe/:token` as the missing 8th root route, and that route **was already in the
corpus's table**. So this iter verified the adjudicator's candidate too, at `app` `ad9f3c498`:

```go
g := e.Group("/v1/labs", apiKeyAuthMiddleware(apiKeys, "labs:write"))
…
// Serve is OUTSIDE the write group — it has OPTIONAL auth (a public Lab's
// workspace is served to anyone; a tenant-private Lab requires a key with
// access). This is the URL the control-plane fetches at boot:
// CP_WORKSPACE_BASE_URL = <app>/v1/labs → GET <base>/<slug>/workspace.tar.gz.
e.GET("/v1/labs/:slug/workspace.tar.gz", h.ServeWorkspace)
```

`internal/web/backend/labs_admin.go:31-41`, wired **unconditionally** at `backend.go:301` (no
`colony.Development` guard, no feature flag). **Confirmed: the count is 8**, and the eighth is a
workspace-**tarball** endpoint mounted on the root `e`, outside the group whose
`apiKeyAuthMiddleware(…, "labs:write")` the corpus describes.

**Stated without over-claiming, deliberately.** It is **not** "an endpoint with no authentication": the
handler's own comment says a public Lab's workspace is public *by design* and a tenant-private one
requires a key with access — **the tenancy decision moved inside the handler**, which is exactly the
shape the group-level sentence fails to describe. `D-M257x-121-2` records this milestone publishing a
**new** absolute quantifier over a security surface inside a repair whose subject *was* absolute
quantifiers; the repair here therefore ships an **enumeration of all eight** rather than an adjective.

## 2. Repaired

- `security_compliance.md:250` — *seven* → **eight**; the table gains the route with its declaration
  site, its wiring site and the source's own words about the gate.
- `security_compliance.md` (the closing summary) — *"11 groups + 7 ungrouped"* → **8**, plus the
  full split: **2 open by design · 2 development-only · 3 self-authenticating in-handler · 1 optional-auth**.
- `architecture_overview.md:406` — *seven* → **EIGHT**, naming the route, pointing at the enumerated
  table and **not restating it** (the single-derivation rule). Its own history is now recorded: this
  line has been wrong about **two different counts** — *"6 Echo groups"* until run 82, *"seven"* root
  mounts until now.
- `iter-131/adjudicator-brief.md:78` — the example predicate. It read *"cms's production ECS state is
  unmeasurable **/** infrastructure was never in a clone set"*, **joining a false proposition to a true
  one with a slash** and inviting adjudicators to book the conjunction. Corrected to the causal form and
  annotated with the rule: **state predicates so that every conjunct is independently false, or split
  them.**

**The brief repair is the one that compounds.** `adj-1` (iter-131) and `adj-C` (iter-135) — two blind,
independent adjudicators one reading apart — made the *same* premise-vs-inference correction, and
`adj-C` traced it to that line. **An instrument that models a conflated predicate teaches the error it
exists to catch**, and every future reading pays for it.

## 3. Test gates

- **Guard family: 18 GREEN · 0 RED · 4 not-run** (commit-/input-scoped members, no `--range`/`--ledger`).
  Not a whole-family green; the runner says so.
- **Zero `rosetta-extensions` files changed.** No code-test gate applies and none is claimed.
- **Whole suite not re-run; §5 rule 60 requires saying so.** Nothing executable has changed since
  iter-132's clean run (`1 failed · 1208 passed`). **Stated as a gap, not characterised as covered.**

---

## Close — 2026-08-08

**Outcome:** the milestone's **4th security-surface understatement** is corrected — `app` mounts
**eight** routes on the root Echo, not seven, and the eighth serves a **workspace tarball under
deliberately optional auth**, wired unconditionally. **Verified at source rather than taken from the
adjudicator**, because the *seat* that first reported the miscount had named the wrong route. And the
**adjudicator brief's own example predicate** — which two independent adjudicators, one reading apart,
had to correct around — is fixed, so the next reading is not trained on the conflation.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5**, unchanged; no reading taken.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**no successor strategy is authorable — `TOK-08`'s sealed refutation branch bars one; running under the user's direct brief**) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: **y — 5 tiks this session** — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **exit-5**
**Decisions:** `D-M257x-136-1` — *a count claim is graded by re-enumerating, never by accepting the
reporter's candidate.* The seat named a route the table already contained; a repair driven by its report
would have changed nothing while appearing to close the finding. The adjudicator's candidate was
therefore verified too, at source. **Reported as a rule, because it is the third time on this milestone
that a correct FINDING arrived with a wrong DIAGNOSIS attached.**
**Side-deliverables:** none.
**Routes carried forward:**
- `FIX-M257x-iter135-adjudicated-live-defects` — **the remainder, and it is the next iter's obvious
  target**: `shared_libraries.md:77` · `security_compliance.md:156` · `clerk-integration.md:126` ·
  `backend.md:13`'s dangling *UNEVEN bullet* · roadrunner's prod state (5 anchors) · `sentinel.md:5` ·
  `dependency_map.md:9` · `ai-readiness.md:18-20` · `org-repos.md:227`,`:370`,`:43` ·
  `ai_architecture.md:111`,`:224` · `next-web-app.md:17`,`:186` · `external_services.md:368` ·
  `adj-E`'s five rotted anchors. **Each adjudicated and cited.**
- `FIX-M257x-iter135-bare-pin-blind-spot` · `FIX-M257x-iter133-two-fives-need-a-fence` ·
  `FIX-M257x-iter134-fence-family-has-no-shared-predicate-layer` ·
  `FIX-M257x-iter132-infrastructure-is-cloneable-so-clone-it` ·
  `FIX-M257x-iter132-suite-walltime-is-not-a-measurement` ·
  `FIX-M257x-iter131-predicate-sets-not-enumerated`.
- **CLOSED this iter:** `FIX-M257x-iter135-brief-teaches-the-error`.
**Lessons:**
1. **A correct finding can arrive with a wrong diagnosis, and the diagnosis is what a repair follows.**
   Re-enumerate; do not accept the reporter's candidate — however many reporters agree on the number.
2. **Fix the instrument before the next reading, not after it.** The brief's defect had already cost two
   adjudicators a correction each, one reading apart, and would have cost the next one too.
3. **When a work list is too big for an iter, split it by CONSEQUENCE and name the remainder.** Two items
   landed completely; the rest are routed with anchors. That is the tripwire applied *before* the creep.
