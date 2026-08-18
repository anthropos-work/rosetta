# Dev-Stack Identity — making a real Clerk operator exist to the platform

**What this is.** A dev stack's database gets its **schema** from Atlas and its **global Sentinel
policy** from `migrate-dev.sh`. It does **not** get an identity. Nothing in the documented dev
bring-up creates the `organizations` / `users` / `memberships` rows for the human who is about to
sign in, nor the per-membership Casbin grants that make that human *authorized* rather than merely
*authenticated*.

That is a deliberate gap on the main dev stack, not an oversight: `dev-setdress.sh` **hard-refuses
`N=0`** so auto-seeding can never touch a developer's primary stack (`corpus/ops/safety.md` §2.5 —
one of the two independent n=0 guards). But the gap has no documented filler for the ordinary case,
and that is what this page fixes.

## The failure this produces, and why it does not look like an identity problem

**You sign in successfully.** Clerk is a separate system and it is perfectly happy: your account
exists, your org exists, your session is real. The JWT it mints carries an `eid` claim (your
platform user UUID) and an `org.eid` claim (your platform org UUID) — **UUIDs of rows a fresh dev
database has never held**.

So the shape of the failure is: *authentication succeeds, authorization silently does not*. What you
actually see:

| Symptom | What is really missing |
|---------|------------------------|
| Pages render their shell, then panels say "couldn't load" or stay empty | `users` / `organizations` / `memberships` rows for the JWT's UUIDs |
| A GraphQL field returns `forbidden` while its **siblings answer normally** | the per-membership Casbin **`g2`** grant |
| Every authorized route 403s | both of the above |

The middle row is the one that costs an afternoon. In the skills explorer, `categories` and
`specializationsByCategory` render and **`skillsBySpecialization` returns `forbidden`** — so it reads
as a broken third-level query. The cause is that the taxonomy resolver checks the org feature **only
when a request carries an organization**
(`app/internal/web/backend/graphql/graph/resolver_skiller_taxonomy_authz.go`: it returns `nil` for a
nil org), and only the third-level call passes one. Same missing row, three different-looking
outcomes.

## Why the Casbin grants are a separate thing from the membership row

Creating a `memberships` row with `role = 'admin'` is **not** enough, and this is the part that
surprises. Sentinel's matcher for an org-feature check is (`app/internal/sentinel/casbin.go`):

```
m3 = g2(r3.org, r3.sub, p3.sub_role) && (r3.org != r3.sub)
     && ('default' == p3.org || r3.org == p3.org) && r3.feat == p3.feat
```

The role is read from a **`g2` grouping row**, not from `memberships.role`. A member with no `g2`
row matches **no policy at all**, whatever the application table says.

And the global policy load does not write one. `migrate-dev.sh` pipes `init_policy.sql`, which
populates `p`, `p2`, `p3` and `p5` — the *model* (which role may use which feature) — and **zero `g`
rows**. Grouping rows are written **per membership**, at membership-creation time, by
`app/internal/sentinel/policy_manager.go`. Skip that code path and you skip the grants.

> **A useful check:** `select p_type, count(*) from sentinel.casbin_rules group by 1`. Only `p*`
> types and no `g2` means no member of any org is authorized for anything, however populated the
> `memberships` table looks.

The three rows a membership is supposed to carry, all written by `app/internal/bootstrap`'s
`JoinOrg` + `SetDefaultOrgFeatureCredits`:

| Row | Written by | Meaning |
|-----|-----------|---------|
| `g2  <org> <user> <role>` | `OrgAddUserToRole` | this user holds this role in this org |
| `g3  <org> <membership id>` | `OrgAllowUserToUseFeature` | this membership may use org features |
| `p6  <org> FEATURE_JOB_SIMULATIONS <credits>` | `SetDefaultOrgFeatureCredits` | the org's credit allowance |

## The three ways to get an identity, and which one applies

| Path | What it does | When it applies |
|------|--------------|-----------------|
| **Clerk webhook** (`corpus/ops/webhook_setup.md`) | Clerk pushes user/org events; the platform creates the rows | You are willing to run a public tunnel, and you are exercising the sync itself |
| **`app`'s `cmd/bootstrap-org` + `cmd/bootstrap-user`** | Creates a Clerk user **and** the DB rows **and** the grants, in one shot | You want a **fresh, tool-made** account. It **mints a new Clerk user**, so it cannot adopt an account you already have |
| **`rext dev-stack/dev-identity.sh`** | Adopts an **existing** Clerk account: reads its ids from Clerk and writes the DB rows + grants | You already sign in as yourself and just need the dev DB to know who that is |

The third is the one this page adds, because it was the case with no answer.

## `dev-identity.sh`

```bash
rext dev-stack/dev-identity.sh --email you@anthropos.work            # the main dev stack (N=0)
rext dev-stack/dev-identity.sh 2 --email you@anthropos.work          # dev-2
rext dev-stack/dev-identity.sh --email you@anthropos.work --print    # resolve + report, write NOTHING
```

**Everything it writes is derived, nothing is typed.** It reads the ids straight out of Clerk:

| Value | Where it lives in Clerk | Becomes |
|-------|------------------------|---------|
| user UUID | the user's first-class **`external_id`** | `public.users.id` — the JWT's `eid` claim |
| org UUID | the org's **`public_metadata.eid`** | `organizations.id` — the JWT's `org.eid` claim |
| role | the membership role, `org:admin` | `admin` (the policy spelling — the namespace is stripped) |

> ⚠️ **The two ids live in different places, and the symmetry you expect is wrong.** A Clerk *user*
> carries the platform UUID in `external_id`; a Clerk *org* carries it in `public_metadata.eid` and
> its `external_id` is **null**. Reading `external_id` off the org — the obvious guess — yields
> nothing, and a tool that pressed on would key an org row to a UUID no JWT will ever name.

It then writes the identity rows and the three grant rows, and publishes the platform's own
`sentinel:policy:invalidate` Redis channel so a **running** backend reloads. That last step matters:
Casbin holds policy in memory and loads it at boot, so rows written underneath a live backend are
invisible until it restarts — which makes a correct fix look like no fix at all.

**Safety.** Every write goes through `docker exec <stack postgres container> psql`. The target is a
**container name derived from `N`**, never a DSN and never a host, so the tool structurally cannot
reach a remote database. It is idempotent (guarded inserts; a second run is a verified no-op),
values-blind (`CLERK_SECRET_KEY` never reaches stdout/stderr, including via a traceback), and it
**refuses rather than guesses** when Clerk does not carry an id it needs.

## Verifying it worked

```bash
# the grants exist
docker exec anthropos-postgresql-1 psql -U postgres -c \
  "select p_type, v0, v1, v2 from sentinel.casbin_rules where p_type in ('g2','g3','p6')"
```

Then sign in and open a surface that needs an **org-scoped** read — the skills explorer's third
level is the sharpest one, because it is the field that fails while its siblings pass:

```
/skills?mode=explorer&categoryId=<category>&specializationId=<specialization>
```

Skills render ⇒ the `g2` grant is live. `forbidden` ⇒ it is not.

---

*Measured 2026-08-18 on `dev-0`: deleting the `g2`/`g3`/`p6` rows reproduced
`skillsBySpecialization → forbidden` with 0 skills rendered; re-running `dev-identity.sh` restored 24
skills with no restart.*
