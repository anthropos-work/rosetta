---
iter: 21
iteration_type: tik
iter_shape: standard
status: archived
opened: 2026-07-30
---

# iter-21 — the missing grant, and a fence so the next one fails loudly

**Active strategy reference:** `TOK-01` move 4 (*"close the honesty items last, deliberately, not as
leftovers"*) — this is the seed-fidelity half of `PT-M256-orgadmin-role-create`, unblocked by the D99
resolution.

## Step 0 — re-survey (mandatory)

Re-measured before targeting, and it **substituted the mechanism** the routing named while leaving the
target intact:

| checked | reading |
|---|---|
| demo-2 live `p3` surface | **17 rows**, `admin` has `taxonomy:read`, nobody has `taxonomy:write` |
| `sentinel/init_policy.sql` `p3` count | **17** — the demo's surface is the platform default, **unmodified** |
| does the rext seeding fleet write `p3` at all? | **NO.** `resetCasbinPTypes = {g2, g3}`; every seeder writes groupings only |
| who applies `sentinel/local_superadmin_grants.sql`? | **NOBODY** — no rext script, no platform Make target, on dev or demo |

**So D99's *conclusion* holds and its *mechanism* does not.** D99 recorded *"the seeder replicated the four
read grants and dropped the single write grant."* The seeder never wrote a `p3` row in its life. The real
shape is finer and more useful, and it is on the record in the platform's own git history:

- platform commit **`c6096d1`** (2026-04-23) *"drop default admin taxonomy:write, add on-demand grants file"*
  deliberately removed the row from the default seed, leaving an inline `NOTE` where it used to be, because
  the capability *"should not be a universal default"*.
- the same commit added **`sentinel/local_superadmin_grants.sql`**, whose stated use case is verbatim
  *"Testing flows that require `taxonomy:write`"* — a first-class, platform-provided, local-only grant file.
- **production carries the row** (the D99 read). Whether grandfathered from before the drop or applied on
  demand, prod's admins have it.
- **no demo or dev stack has ever applied that file.**

**Which resolves the honesty objection D99 raised, rather than overriding it.** iter-20 refused to grant
itself the permission under test, correctly, because inventing a grant would manufacture the capability.
Applying the platform's **own sanctioned grant file**, row-identical to production's, is the opposite move:
it stops the demo misrepresenting production. The row is not ours. It is the platform's, it names this exact
use case, and prod has it.

## Cluster / target identified

`PT-M256-orgadmin-role-create`'s **seed-side blocker**, plus the generalisable fence the user asked for. The
Playthrough itself and its negative control are **iter-22** — deliberately, not as a punt: this milestone's
own scope-creep tripwire counts a 4-line iter as the failure mode, and the live probe below is exactly the
measurement that de-risks writing the spec.

## Hypothesis

1. Seeding the single row `('p3','default','admin','org:feature:taxonomy:write','','','')` makes
   `createJobRole` **succeed** for a seeded org admin. **Probe this FIRST, by hand, before writing a line of
   seeder** — a seeder plus a fence shipped against an unverified premise is how iter-05 spent fifteen iters
   inside the wrong dialog.
2. A checked-in expected `p3` surface plus a live diff catches the next dropped grant **at seed time**
   instead of fifteen iters later as *"the form doesn't work"* — **and catches the opposite drift too**: an
   *extra* live grant is the mechanical form of exactly the hazard iter-20 refused by hand.

## Expected lift

No gate clause moves this iter (the UC lands in iter-22). Planned deliverables:

- a seeded on-demand `p3` grant, idempotent, `--reset`-safe, provenance-cited;
- a **bidirectional** `p3` fidelity fence — a gate in its own verb, advisory in the bring-up (D-M255-1's
  two-consumer rule);
- the live probe result recorded either way.

## Phase plan

- **A** — hand-insert the row on demo-2 + reload Sentinel + drive `createJobRole`. Revert the row. (If the
  mutation still fails, the seeder is premature and this iter re-scopes to reporting why.)
- **B** — `PolicyGrantsSeeder` + the fence + `--policy-check`, each assertion watched RED.
- **C** — tag rext, re-pin `stack-demo/rosetta-extensions`, reset-to-seed demo-2, confirm the row arrives by
  the seeder and the fence reads clean.
- **D** — suite run for no-regression; restore the drifted cockpit fixture + sha-verify `99e2f315`.

## RE-SCOPE, declared mid-iter (Phase A) — the premise was refuted TWICE, and the iter follows the cause

Phase A's escalation condition fired, twice. Recorded here rather than in a close note, because it changes
the iter's planned scope and the scope-creep tripwire is owed an explicit accounting.

**The grant is necessary and NOT sufficient.** With the row hand-inserted and Sentinel reloaded,
`createJobRole`'s `unauthorized: forbidden` **is gone** — and two further blockers stand behind it, each
measured, each a *config-fidelity* gap of the same class, each invisible in the UI:

| # | measured error | root cause | fix home |
|---|---|---|---|
| 1 | `unauthorized: forbidden` | the on-demand `p3 admin → org:feature:taxonomy:write` grant is applied by nothing | rext `stack-seeding` |
| 2 | `can't generate skill embedding: … azure client EU is not set` | `SKILLER_AZURE_OPENAI_{KEY,ENDPOINT_URL}` (`app/main.go:491-499`, declared in `app/terraform/main.tf`) are set on **no** demo or dev stack and are **absent from the 56-gene secret DNA** | rext `stack-secrets` |
| 3 | `duplicate key value violates unique constraint "job_role_embeddings_pkey"` | the snapshot replay does `TRUNCATE … RESTART IDENTITY` then COPYs rows **with explicit ids**, and never advances the sequence: `job_role_embeddings` max id **21274** / seq at **4**; `skill_embeddings` max **43583** / seq at **1** | rext `stack-snapshot` |

**#3 is the generalisable one and it is bigger than this UC.** Those are the **only two** identity columns in
`public`, and **both** are broken — so on every demo built to date, *every* taxonomy write dies on a
duplicate primary key: creating a custom role (this UC) and creating a custom skill
(`app/internal/skilltaxonomy/skill.go:80`, the sibling call). The replay code even carries a note that
identity columns are handled nowhere (`directus/structure.go:66-71`) — for the directus surface, which has
none. The taxonomy surface has two.

**So the iter follows the cause instead of shipping one third of a fix.** Planned scope becomes the three
fidelity fixes that make a demo's taxonomy write path work, each proven RED-then-GREEN, plus the `p3` fence.
The Playthrough and its negative control stay in **iter-22** — landing them together is what keeps gate
clauses 2 and 3 in lockstep.

## Escalation conditions

- Phase A's probe shows the mutation still denied after the grant → **do not ship the seeder**; record the
  second cause and route the UC forward with the new evidence.
- The fence cannot be made to fire RED in both directions → it is not a fence; do not claim it as one.

## Acceptable close-no-lift outcomes

Phase A refuting hypothesis 1 is a first-class close-no-lift: it would mean the authorization story has a
second layer, which is worth more than a seeder built on a guess.
