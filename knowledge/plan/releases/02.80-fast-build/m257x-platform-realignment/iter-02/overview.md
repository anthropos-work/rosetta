---
milestone: M257x
iter: 02
iteration_type: tik
status: in-progress
date: 2026-07-31
---

# iter-02 — derive the migration tuple from `repos.yml`, and fence what cannot be derived

**Type:** tik, under `TOK-01` ("instrument first, then follow").

## Step 0 — re-survey (mandatory) — and it substituted the target

`TOK-01`'s **Next-tik direction** named `FIX-M257x-rext-pin` on **odysseus**. Re-measured on the new Mac
(D-v28-15), both halves of that direction are stale:

| TOK-01 claim (2026-07-31, odysseus) | measured now (new Mac) |
|---|---|
| `.agentspace/rext.tag` = `cockpit-deeplinks-v1`, **63 commits behind** `main` | `rext.tag` = **`fast-build-m257x-iter-01`** |
| the tag may not be on origin (*"tagging is not publishing"*) | **on origin**: `31d2b5df…  refs/tags/fast-build-m257x-iter-01` |
| pin is **N behind** `main` | `31d2b5df` **== `origin/main`** — 0 behind, 0 ahead |
| the consumption clone sits at a *different, newer* tag ⇒ FATAL `ensure-clones.sh:94-101` | **there is no `stack-demo/`** — the mismatch cannot exist |
| host = `odysseus` | odysseus **retired** (D-v28-15); host is this Mac |

**`FIX-M257x-rext-pin` is ABSORBED** by the machine move — the pin was re-created deliberately and is
self-consistent, current, and published. Nothing to fix. Re-doing it would be exactly the "re-derived from
scratch each time" waste this milestone exists to end.

**But the instrument is still blocked, by a bigger and different fact.** Surveyed this box:

    docker podman colima nerdctl lima orbstack   → ALL absent
    /Applications/{Docker,OrbStack}.app          → absent
    gh · psql · tailscale                        → absent
    go · atlas · node · pnpm · python3           → present

**There is no container runtime at all.** Gate clauses **1** (cold `demo-up` ×3), **2** (Playthrough suite)
and the **live half of 4** are host-blocked until a runtime is installed — an admin/licensing action, not a
tik's work. Routed forward as `HOST-M257x-container-runtime` (below); it does **not** change what code lands
in this iter, so per Phase 5 § 4 it is a routed-forward item, not a user-blocker.

**Substitution, under the same TOK-01, minimal deviation:** advance to TOK-01's **step 2** —
`FIX-M257x-migrate-tuple`, already a named carry-forward routed to "iter-02/03", and fully landable offline.

## Active strategy reference

`TOK-01` step 2 — *"Fix the mechanism, not the symptom. The symptom is 'rext writes `jobsimulation.*`'. The
mechanism is `migrate-demo.sh`'s hand-maintained 4-tuple that creates those schemas itself while ignoring
`repos.yml`'s `migrations:` flag. Derive the tuple from `repos.yml`."*

## Cluster / target identified

`corpus/ops/platform-alignment.md` §2 (the mechanism) + §8 (the fence). Three hardcoded sites across **two**
scripts — the twin `migrate-dev.sh` carries the identical tuple and iter-01 did not know that:

| site | content |
|---|---|
| `demo-stack/migrate-demo.sh:29` | `_have_svc` pre-flight loop over the 4-tuple |
| `demo-stack/migrate-demo.sh:83-85` | `CREATE SCHEMA` for `cms`, `jobsimulation`, `skillpath` |
| `demo-stack/migrate-demo.sh:106` | atlas loop over `app:public cms:cms jobsimulation:jobsimulation skillpath:skillpath` |
| `dev-stack/migrate-dev.sh:57-59` | same `CREATE SCHEMA` trio |
| `dev-stack/migrate-dev.sh:68` | same atlas 4-tuple |

Ground truth, fetched from **origin HEAD `1e8e7540`** this iter (not a stale clone — there is no local
`platform` clone on this box at all):

    app          migrations: true   schema: public     <-- the ONLY one of each
    cms          migrations: false  (schema: key DELETED)
    jobsimulation migrations: false (schema: key DELETED)
    roadrunner   migrations: false
    sentinel · storage · messenger · next-web-app · studio-desk · graphql-wundergraph  migrations: false
    skillpath    ABSENT ENTIRELY

**The 4-tuple is wrong on 3 of its 4 entries**, and `skillpath` — the canary §2 predicted — names a repo
`repos.yml` no longer contains.

## Hypothesis

Deriving the atlas `(repo → schema)` pairs from `repos.yml`'s **machine-readable fields only**
(`name`/`migrations`/`schema`, per `D-M257x-2`) disarms the §2 M810 time bomb: the pairs then come from the
file that declares the truth, so when the legacy repos leave the clone set the list shrinks *with* them
instead of `[ -d ] || continue` silently skipping 13 write targets at once.

The schema-**creation** list cannot be fully derived — `sentinel` is `migrations: false` **and** alive with
its own `sentinel` schema (protocol **Trap A**). So per the protocol's own rule — *"Derive it, or fence it.
Never both-hand-maintain-and-trust it"* — the residual becomes an **explicit, cited, fenced** list rather than
an invisible one.

## Expected lift

**No exit-gate clause flips.** This is clause 4's *mechanism* half; clause 4 additionally requires the write
re-point plus a live assert, both of which need a stack this box cannot yet run. Honest expected metric delta
on "clauses met": **0/5 → 0/5**. The measurable sub-progress is: hardcoded platform-service tuples in rext
**5 sites → 0**, and dead-schema creations **3 → 0 underived**.

Grading this iter as `closed-fixed` therefore depends on the *planned deliverables* landing, not on a clause
flipping — per Phase 4 Step 0, status grades planned scope.

## Phase plan

Per `platform-alignment.md`:
1. §4 signals 1–3 against origin HEAD (**done** in Step 0 — repo set, migrating repos, declared schemas).
2. §7.1 measure the write surface by scan, splitting live code from comments.
3. Implement the derivation in both scripts; §7.3 leave the history in a comment naming what the relation
   *was* and which change removed it.
4. §8 fence, **watched going RED before it is trusted**.

## Escalation conditions

- Derivation would require reading `repos.yml` **prose** → stop; `D-M257x-2` forbids it.
- Removing legacy schema creation would manifest a live breakage that cannot be verified on this box →
  do **not** ship a known-broken bring-up; keep the legacy creations behind an explicit fenced transitional
  list and route the removal to the re-point tik. (This is the anticipated outcome, not a failure.)
- Test gates RED → user-blocker.

## Acceptable close-no-lift outcomes

If the scan shows the tuple is already derived somewhere and the hardcoding is vestigial, recording that
falsification with evidence is a complete iter. Likewise if the derivation proves to need the live schema set
(protocol Trap A generalising further than expected), characterising that is the deliverable.
