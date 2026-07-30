# iter-20 — decisions

## D97 — `pt-orgadmin-role-create` is blocked by AUTHORIZATION, and every prior diagnosis was wrong

**The mutation fires.** iter-05 concluded the create-role dialog's `Save` "fails silently with an EMPTY
`alert` region", and routed the UC behind a live-LLM `Generate` leg. Nobody had looked at the network. Driven:

```
REQ  organizationJobRolesByName
REQ  createJobRole
ERR  200 {"errors":[{"message":"Failed to fetch from Subgraph 'backend'.","extensions":{
       "errors":[{"message":"unauthorized: forbidden","path":["createJobRole"],
       "extensions":{"code":"DOWNSTREAM_SERVICE_ERROR"}}],"serviceName":"backend"}}],"data":null}
```

`createJobRole` is sent, reaches `backend`, and is **DENIED**. The transport is HTTP **200** with a GraphQL
error body, which is why nothing surfaces: dialog stays open, `[role=alert]` count **0**, catalog count
**20 → 20**, the role absent on re-read. Watched 40 s.

**And the permission is nowhere in the policy.** The resolver's authz layer
(`app/internal/web/backend/graphql/graph/resolver_skiller_taxonomy_authz.go:75`) checks
`permission.OrgFeatureTaxonomyWrite`. Every taxonomy row in the running Sentinel policy:

| p_type | domain | role | feature |
|---|---|---|---|
| p3 | default | admin | `org:feature:taxonomy:read` |
| p3 | default | member | `org:feature:taxonomy:read` |
| p3 | default | candidate | `org:feature:taxonomy:read` |
| p3 | default | content_creator | `org:feature:taxonomy:read` |

**`org:feature:taxonomy:write` is granted to NO role.** Not to `admin`, not to anyone. The read counterpart
exists for four roles; the write counterpart exists for none. So `createJobRole` cannot succeed for any
seeded actor, and this is not a stale-enforcer-cache issue either — the runner reloads Sentinel on every
`--reset` (M203 iter-05) and did so here.

**Two claims from iter-05 D18 are now retracted:**

- *"Suggest skills transforms the dialog and the primary button becomes `Generate`"* — irrelevant, and the
  path was never needed. The dialog's "Core skills" mode choice has **"Start from scratch" ALREADY
  SELECTED** on open (measured: the first card carries `background: var(--color-primary-light); border:
  1.5px solid var(--color-primary)`, before *and* after clicking it). There is no LLM leg on the happy path.
- *"`Save` is enabled with the form incomplete"* — half right, and the wrong half was load-bearing. `Save`
  is correctly **DISABLED** while the fields are empty and enables once both are filled. It is not a
  validation problem at all.

## D98 — PRODUCT DEFECT: a forbidden mutation renders as nothing whatsoever

Independent of the Playthrough, and user-facing: an org admin clicks `Save`, the backend refuses, and the UI
shows **no error of any kind** — no alert, no toast, no inline message, the dialog simply stays open. The
only signal is in devtools. A user's reasonable conclusion is that the app is broken or hung.

Worth stating plainly because it is *why this took fifteen iters*: the failure is invisible at the exact
layer everyone was looking at. iter-05 stared at the form because the form was all the product would show
it.

## D99 — Granting ourselves the permission would MANUFACTURE the capability under test → escalate

The mechanical fix is one policy row: grant `admin` → `org:feature:taxonomy:write` in the demo's Sentinel
policy. The seeder already writes per-membership feature grants (the `g3` `FEATURE_JOB_SIMULATIONS` family,
171 rows), so it is technically an rext-owned, zero-platform-edit change.

**It is refused as an implementation decision**, because a Playthrough that passes only because we granted
ourselves a permission the product grants to nobody proves nothing about the product. That is exactly
iter-07's rule as sharpened at iter-17 — *the one move that can manufacture a state the application never
learns is the wrong move where "the app permitted it" is what you are proving* — one layer down from the
DOM, in the policy.

**So the disposition is a user decision, and it is one of two:**

| option | consequence |
|---|---|
| **(a) It is a platform authorization gap.** Report the defect; `org-admin.roles.UC1` becomes the milestone's **first** `unimplementable-without-platform-edit`. | Clause 3's landed half tops out at **org-admin 3 of 4**, with a written verdict. Zero `unimplementable` becomes one. The re-scope trigger (> 3) does **not** fire. |
| **(b) The demo's policy is legitimately incomplete** and production grants this some other way. | Then the missing grant is a **seed** gap, the seeder writes it, and the UC lands as an 8th mutating Playthrough. |

**What is needed to choose is one fact this iter could not obtain locally:** whether a real production org
has an `org:feature:taxonomy:write` grant. That is a **production Sentinel policy read** — outside this
milestone's local-only remit and requiring the standing sign-off rule. It is a single query.

The strong circumstantial reading is **(a)**: the product ships a "New Role" button, a create dialog, and a
`createJobRole` mutation that no role can execute. But it is circumstantial, and the milestone's own record
is full of confident readings that measurement overturned — including two in this iter. So it is put to the
user rather than assumed.
