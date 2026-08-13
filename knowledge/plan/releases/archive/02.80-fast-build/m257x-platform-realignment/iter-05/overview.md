---
milestone: M257x
iter: 05
iteration_type: tik
status: closed-fixed
created: 2026-07-31
---

# iter-05 — clear the autoverify ✗s that stand between here and clause 1

**Active strategy reference:** `TOK-01`, still **step 1** (*"unblock the gate's instrument"*). iter-04 got
the instrument to run end-to-end for the first time; it now reports **3 FAILED checks**. Those three ARE
the remaining distance to clause 1 (`green:true / 0 warnings`, ×3 consecutive cold cycles), so clearing
them is the same strategy continued, not a new one.

## Step 0 — re-survey (measured at iter-05 open, not inherited from iter-04's close)

| probe | reading |
|---|---|
| `demo-1-directus-1` | `Exited (1) 10 minutes ago` — still down, stack otherwise up |
| directus logs | `Error: connect ECONNREFUSED 172.18.0.5:5432` at startup |
| schemas actually in the DB | `auth · cms · directus · extensions · jobsimulation · public · sentinel` — **no `skillpath`**, exactly as origin `repos.yml` dictates |

The second row settles the skillpath check by measurement: the schema is **correctly** absent, so the probe
demanding it is the thing that is wrong.

## Cluster / target identified — a declared 2-step shape

Both targets are rext-owned, both are clause-1 blockers, and both are cheap. Declaring the multi-step shape
up front so the scope-creep tripwire counts against **this** plan, not a single-target one.

1. **`FIX-M257x-autoverify-skillpath-schema`** — the `postgres-schemas` probe fails with
   `missing schemas: skillpath`. iter-02 correctly removed `skillpath` (absent from origin `repos.yml`,
   zero rext writes); this probe still requires it. A check pinning the *contents* of a list that is
   supposed to follow the platform — `platform-alignment.md` §8 rule 3, third instance in this milestone.

2. **`FIX-M257-directus-coldstart-order`** — carried since **M257 iter-02** as "platform-shape-dependent"
   and never reproduced. It just reproduced: directus raced postgres, got `ECONNREFUSED`, exited 1, and has
   `restarts=0`, so nothing recovered it. The per-stack Directus is the entire point of `--local-content`,
   and while it is down `directus HTTP 000000` is a guaranteed ✗ on every cycle.

## Hypothesis

The skillpath probe's expected-schema set can be derived from the same `repos.yml` source iter-02 already
built (`stack-core/lib/repos_yml.sh`) instead of being hand-listed — fixing the instance and the class in
one move. Directus needs a readiness dependency and/or a restart policy in the rext-generated compose
override; it is a startup race, not a configuration error, because the same container image and env work
once postgres is up.

## Expected lift

**2 of the 3 failing autoverify checks cleared.** Clause count stays 0/5 — clause 1 needs three consecutive
fully-green cold cycles and this iter does not attempt one. Anything short of "both checks pass on a live
re-verify" is a partial.

## Phase plan

1. Derive the expected-schema set for the probe from `repos.yml`; mutation-verify it goes RED.
2. Fix the directus startup race in the rext-owned compose override; restart and confirm it serves.
3. Re-run `autoverify` against the **running** stack and count the ✗s.
4. Tag + push rext (tagging is not publishing), re-point the consumption clone.

## Escalation conditions

- If the directus failure turns out to need a platform-repo edit → route forward; the v2.8 zero-edit
  constraint holds.
- If clearing the skillpath probe requires deleting an assertion rather than re-deriving it → stop and
  record, because an assertion deleted is a fence removed (§7 rule 2: re-point to the canonical target,
  never to nothing).

## Acceptable close-no-lift outcomes

A measured falsification — e.g. directus's exit is not a race but a genuine config/schema break whose fix
is out of scope — closes this iter honestly provided the mechanism is named and cited.
