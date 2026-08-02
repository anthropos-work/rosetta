# iter-38 — the multi-tenancy fence, re-derived and PREDICATE-tested (not re-matched)

`security_compliance.md`'s Layer-1 fence has been wrong four times, in both directions. iter-34's
fifth generation ships its own derivation command. Rule: **run it, don't quote it** — and then test the
PREDICATE, because every prior failure reproduced the counts exactly and still asserted something false
(§5 rule 17).

## 1. The counts, re-derived from `stack-demo/app` at platform origin `2adcf71`

    $ cd app/internal/data/ent/schema
    $ ls *.go | wc -l                                     -> 139
    $ ALL=$(grep -l 'ent.Schema' *.go | sort); echo "$ALL" | wc -l    -> 135
    $ grep -l 'OrganizationMixin{}'   *.go | wc -l        -> 30
    $ grep -l 'OrganizationIDMixin{}' *.go | wc -l        ->  7
    $ comm -23 <(echo "$ALL") <(grep -lE 'Organization(ID)?Mixin\{\}' *.go | sort) \
        | xargs grep -l '"organization_id"' | wc -l       -> 18

**Every figure in the doc reproduces exactly: 139 / 135 / 30 / 7 / 18.** So did every earlier generation's,
which is precisely why the counts are not the test.

## 2. The PREDICATE test — the step the four earlier failures skipped

The claim is a conjunction: *"16 carry an `organization_id` with no policy of any kind"*, i.e. **neither an
own `Policy()` NOR a policy-bearing mixin**. Tested per file across all 18:

| file | own `Policy()` | mixins |
|---|---|---|
| `org_membership.go` | **YES** | PrimaryKey, CreatedAt, UpdatedAt |
| `academy_feedback.go` | no | PrimaryKey, CreatedAt, UpdatedAt, **UserMixin** |
| the other 16 | no | PrimaryKey / CreatedAt / UpdatedAt / DeletedAt only |

And the second conjunct — that those plumbing mixins carry no policy — checked rather than assumed:

    PrimaryKeyMixin  Policy defined? 0      UserMixin            Policy defined? 1
    CreatedAtMixin   Policy defined? 0      OrganizationMixin    Policy defined? 1
    UpdatedAtMixin   Policy defined? 0      OrganizationIDMixin  Policy defined? 0
    DeletedAtMixin   Policy defined? 0

Both exclusions are also right about the KIND of policy, which is the distinction the fourth failure got
backwards:

- `Membership.Policy()` (`org_membership.go:172-188`) is **organization**-scoped —
  `DenyMismatchedOrganization` on create, `AllowCurrentOrgEdgesOrSkipRule`, terminating in
  `privacy.AlwaysDenyRule()`. So *"31 auto-filter by ORGANIZATION = 30 + Membership"* holds.
- `UserMixin.Policy()` (`mixin.go`) is **owner**-scoped — `FilterOwnerRule()` + `DenyIfNoUserInContext()`,
  by *user*, not by organization. So `academy_feedback.go` is correctly excluded from the org count and
  correctly described as user-scoped rather than org-scoped.

## Verdict

**The fifth generation of this fence is CORRECT** — counts reproduce, and the predicate holds under a
per-file test of both conjuncts. This is the first generation to be verified by predicate rather than by
denominator.

## One MINOR, recorded rather than filed as a blocker

The `OrganizationIDMixin{}` seven are *also* unpoliced (the mixin declares no `Policy()`, confirmed above),
so the total unpoliced-but-org-columned set is **23**, not 16. The text does say so immediately above
("Seven use `OrganizationIDMixin{}` … with **no policy**"), and the 16 are qualified as *"most likely to be
missed"* — which is defensible, since the seven at least carry a mixin name an auditor can grep for. Not
false, so not a blocker; worth a clarifying clause.
