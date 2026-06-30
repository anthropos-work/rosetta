# Profile-completeness spec — the "complete profile" rubric

> **Scope (v1.10 "method acting" M44):** the **DATA-DENSITY** rubric for a *believable, fully-populated*
> profile — for **every** character a presenter can land on: end-user heroes, **manager** heroes, **and** the
> bulk fill-members. M41 gave the logged-in hero a timeline; **M44 makes the WHOLE roster read as real people**.
> **DATA DENSITY ONLY** — zero platform / next-web edits (no "% complete" widget); the deliverable is **rows**,
> mapped to their **seeding surface** + a **Playwright acceptance assertion** (the M42e/M42m coverage gate). It
> EXTENDS, never contradicts, [`stories-spec.md`](stories-spec.md) (the verified-skill chain) +
> [`seeding-spec.md`](../seeding-spec.md) (the isolation boundary) + [`coverage-protocol.md`](coverage-protocol.md)
> (the semantic believability gate).

## The rubric — three layers

A profile is "complete" when all three layers are populated and **self-consistent** (the coverage gate's
persona clause). Each layer maps to seeding surfaces and, where a character page renders it, a Playwright
assertion in the manifest (`stack-verify/e2e/lib/coverage-manifest.ts`).

| Layer | What it means | Seeded by |
|---|---|---|
| **Identity** | who this person *is* — real name, real-photo avatar (menu == profile), org name + logo, role + title | `UsersSeeder` (name/email/avatar/`user_basic_info` role), `OrgSeeder` (name/logo), `roster.go` (menu claims) |
| **Content** | the *artifacts* of a career — work history, education, certifications, projects, verified + claimed skills | `ProfileSeeder` (timeline + claimed tail), `PersonaSeeder` (verified chain), **`CertificatesSeeder`** + **`ProjectsSeeder`** (M44 §B) |
| **Semantic** | the layers *cohere* — role ↔ skills ↔ history ↔ certs/projects read as ONE believable person; the claimed-vs-verified gap is coherent with the trajectory | the closure gate (every node-id real), the trajectory shaping (§A), role-coherent banks (§B) |

## Per-vantage completeness — member vs manager

The two presenter vantages (the two coverage manifest namespaces, employee + manager) have DIFFERENT
completeness bars. M44 fills both.

### End-user (member) vantage — `/profile`, `/home`, the Library

| Component | Seeding surface | DB write | Playwright assertion (manifest section) |
|---|---|---|---|
| Real name + org-domain email | `UsersSeeder` (`users.go`) | `public.users` firstname/lastname/email | persona self-consistency: profile name == menu name |
| Real-photo avatar (menu == profile) | `photoAvatarDataURI` (`avatar.go`) → `users.picture` + `roster.go` Picture | `public.users.picture` (data URI) — **NB: `/enterprise/members` reads a different column, `memberships.picture_url`; see §D gotcha** | persona: avatar is a real photo, menu == profile (`lib/persona-assert.ts`) |
| Role + title | `UsersSeeder` `user_basic_info` backfill | `public.user_basic_info.job_role_id/job_title` | `/profile` header section: role label non-empty |
| Work history (3 jobs) | `ProfileSeeder` (`profile.go`) | `public.user_experiences` (+ `companies`) | `/profile` Experience section: `count >= floor` |
| Education (1 degree) | `ProfileSeeder` | `public.user_educations` | `/profile` Education section: real content |
| **Certifications (2-3 hero; ~45% of members carry 1-2 — v1.10b M50)** | **`CertificatesSeeder`** (`certificates.go`, M44 §B1; member-coverage extended in v1.10b "fit-up" M50) | `public.user_certifications` | `/profile` Certifications: `count >= 1`; **manager Workforce → Talent Pool "Certifications" chart** |
| **Projects (3-4)** | **`ProjectsSeeder`** (`projects.go`, M44 §B2) | `public.user_projects` | `/profile` Projects section: `count >= 1` |
| **Spoken languages (every member: 1-3 — v1.10b M50)** | **`MemberLanguagesSeeder`** (`member_languages.go`, v1.10b "fit-up" M50) | `public.world_languages` (catalog) + `public.user_languages` → `public.membership_languages` (via the DB AFTER-INSERT trigger) | **manager Workforce → Talent Pool "Languages spoken" chart** (`/profile` also reads `user_languages`) |
| Verified skills + Skill-Spotlight chart | `PersonaSeeder` (the §3 7-table chain) | `public.user_skills` (is_verified) + `user_skill_evidences` | `/profile` Skill Spotlight: chart renders (>=2 datapoints) |
| Claimed-but-unverified tail (the gap) | `ProfileSeeder` claimed tail | `user_skills` (is_verified=false) + evidences (anthropos NULL) | persona: the claimed-vs-verified gap is visible |
| **Self-rating completeness (§A)** | `PersonaSeeder` `user_level` branch | `user_skill_evidences.user_level` (set or NULL) | thriving hero: a self-assessment shows; struggling: it reads incomplete |

### Manager vantage — the manager's OWN `/profile` (M44 §C) + the org dashboards

A manager PRIMARILY demos the org-intelligence surfaces (the M36 dashboards), but **her own `/profile` must
no longer be empty**. M44 §C removes the pre-M44 `IsManager` skips and gives her a **modest** personal profile.

| Component | Seeding surface | Manager-specific shape |
|---|---|---|
| Verified skills | `PersonaSeeder` (unskipped, §C) | a FLAT 3-8 skills, L1-L2 band (34-52), no growth arc — populated but modest |
| Self-rating | `PersonaSeeder` `user_level` branch | self-rated (`user_level` set — a manager is a normal member on her own page) |
| Work history | `ProfileSeeder` (unskipped, §C) | a leadership ladder: "Engineering/Sales Manager" ← "Team Lead" ← "Senior X" |
| Education | `ProfileSeeder` | 1 degree (the standard depth) |
| Claimed tail | `ProfileSeeder` | a SMALL flat tail (`managerClaimedTail` = 8), not the deep ~60 |
| Org dashboards | the M36 dashboard seeders (unchanged) | the funnel / gap / teams / succession she rides on |

### Bulk-member depth — `/enterprise/members` (M44 §D)

The manager's `/enterprise/members` roster must read as **real people**, not heroes amid blank profiles.

| Component | Seeding surface | Shape |
|---|---|---|
| Avatar (EVERY member) | `photoAvatarDataURI` (`users.go:156`) → **`memberships.picture_url`** (the column this surface reads) **and** `users.picture` | a real-photo data URI per member |
| Shallow career (EVERY member) | `ProfileSeeder` member pass (`seedMemberProfile`, §D) | 3 short-tenure experiences + 1 education + a flat <=6-skill claimed tail; role mirrors the membership row |

> **§D avatar GOTCHA (M44 §D fix1, load-bearing — target the RIGHT column):** `/enterprise/members` does
> **NOT** read `public.users.picture`. The `EnterpriseMembersTable` → `memberEmployeeColumn.tsx` renders
> `<NamedAvatar src={member.pictureUrl}>`, and the GraphQL `pictureUrl` field resolves in the platform `app`
> service as `_Membership_pictureUrl → obj.PictureURL` = the **`public.memberships.picture_url`** column. The
> M44 §D build filled only `users.picture` (340/341), so the members list rendered **silhouettes** even though
> `users.picture` was full. The fix: `UsersSeeder` now writes the SAME `photoAvatarDataURI(uid)` into
> **`memberships.picture_url`** for every member (heroes + fill + managers) — on a fresh seed via the COPY, and
> on a RE-seed via an idempotent UPDATE backfill (`backfillMembershipPictures`, since `CopyRowsIdempotent`'s
> `ON CONFLICT (id) DO NOTHING` can't heal an existing row). Render-verified: 0 silhouettes, every row a photo.
> **Future avatar work: `/profile` + the menu read `users.picture`; `/enterprise/members` reads
> `memberships.picture_url`. Fill BOTH.** (The shallow-career half of §D — the timeline pass over the non-hero
> slots — is unchanged; the `/enterprise/members` assertion is the manager-manifest section's cardinality +
> per-member avatar.)

## How M44 maps to the seeding surfaces (the build summary)

- **§A trajectory-aware self-rating** — `PersonaSeeder` writes `user_skill_evidences.user_level` only for a
  **self-rated** hero (`Persona.EffectiveSelfRated()`: struggling = false → `user_level` NULL; everyone else =
  true). A thriving hero shows a completed self-assessment; a struggling hero "hasn't self-rated" (the claimed
  side is absent) while her 2-3 verified skills still render the chart.
- **§B1 `CertificatesSeeder`** (surface `"certificates"`) + **§B2 `ProjectsSeeder`** (surface `"projects"`) —
  2-3 `user_certifications` + 3-4 `user_projects` per end-user hero, role-coherent banks, idempotent COPY,
  closure-clean skills, managers skipped (the §C path owns manager profiles). **Both surface as top-level
  `TimelineGroupedItems.certifications` / `.projects`** on `/profile`.
- **§C manager personal data** — the `IsManager` skips at `persona.go` + `profile.go` are removed; the manager
  gets a modest flat verified set + a manager-track timeline (above).
- **§D bulk-member depth** — every non-hero member gets a shallow career (above).

### v1.10b "fit-up" M50 — the Talent-tab fill (languages + org-wide cert coverage)

The field review flagged the manager's Workforce → **Talent Pool** tab: *"no language spoken + Certification
are really low numbers."* Two surfaces were genuinely empty/sparse on a clean demo, both filled in M50:

- **`MemberLanguagesSeeder`** (surface `"member_languages"`, `member_languages.go`) — the language tables
  (`world_languages` / `user_languages` / `membership_languages`) were 0 rows DB-wide. It populates
  `world_languages` with a **curated ISO-639-1 standard catalog** (16 EU-professional-weighted entries — a
  published standard list is a factual reference, NOT a fabricated taxonomy node-id; the closure gene governs
  skiller node-ids, not ISO codes) then writes **per-member `user_languages`** (every member, 1-3 distinct
  languages: a location-coherent native at level 5 + near-universal English + an occasional third). **It seeds
  ONLY `user_languages`** — the DB AFTER-INSERT trigger `on_insert_user_languages_insert_membership_languages`
  fans each row out to `membership_languages` (the column the org Talent tab reads). Idempotent COPY on
  `id` (deterministic per user+language) → a re-seed conflicts-and-skips, so the trigger does not re-fire.
  Runs in the DAG after `users`.
- **`CertificatesSeeder` member-coverage extension** — certs were **hero-only** (the "really low numbers" gap).
  M50 extends the seeder to give a deterministic **~45% of the supporting population** 1-2 role-coherent certs
  (heroes still carry their 2-3; managers still excluded — the §C path owns them), so the org reads as a
  credentialed workforce. The cert bank + skills stay role-coherent via the SAME deterministic role assignment
  `UsersSeeder` writes (no fabrication — an empty taxonomy pool → the general bank + no skills tag).
- **The manager coverage manifest is strengthened to PROVE these** (the D4/F1 reconciliation): before M50 the
  manager M42 gate passed `(0,0)` BLIND to languages/certs/member-fields (the manifest never asserted them). M50
  adds a `preAssert` tab-click capability + a `textMatch` (OR-over-alternatives) assertion to
  `stack-verify/e2e/`, and the manager manifest now asserts: `/enterprise/members` Location column + a real
  seeded city value, and `/enterprise/workforce` → Talent Pool's "Languages spoken" + "Certifications" charts.
  The gate is MET only when this STRENGTHENED manifest passes — proving the gaps closed, not passing blind.

## The isolation + closure invariants (inherited, unchanged)

Every M44 write surface is **`PerStackIsolated`** (per-stack Postgres — the same zero-pollution class as the
rest of the fleet; see [`seeding-spec.md`](../seeding-spec.md) §"production-isolation boundary"). Every
skill node-id / role node-id comes from the **same replayed taxonomy resolvers** the verified chain uses
(`resolveJobRoleRefs` / `resolveNamedSkillRefs`) — an absent taxonomy degrades (blank skills/role, tail
skipped) and **never fabricates** a node-id, so the closure gene (`datadna measure-closure`) stays GREEN.

## Live-schema corrections (the overview's column guesses were WRONG)

M41's spec-notes warned "the live demo schema wins". M44 verified the cert/project schemas against the live
demo-3 DB; the overview's guesses were wrong on every count:

- **The certificates table is `public.user_certifications`** (NOT `user_certificates`). Columns:
  `certification` (text NOT NULL — the cert *name*, NOT `cert_name`), `issued_by` (text NOT NULL), `from`
  (DATE NOT NULL — issued-at), `to` (DATE — expires-at), `skills` (json), `field_of_study` / `credential_id`
  (varchar). **No `created_at`/`updated_at`** (unlike most of the fleet). **No `organization_id`.**
- **`public.user_projects`**: `title` (text NOT NULL — NOT `project_name`), `from` (DATE NOT NULL —
  started-at), `to` (DATE — ended-at, NULL = ongoing, NOT `end_date`), `description` / `url` (varchar),
  `skills` (json), two nullable related-FK edges. **No `organization_id`.**

Both are also FK targets for `user_skills` provenance edges (`user_skill_certification` / `user_skill_project`)
in the `user_skills_check_foreign_keys` CHECK — so claimed skills *could* attach to them, though M44 keeps
certs/projects standalone (their own `skills` json) and ties the claimed tail to experiences/educations.

**Language tables (verified against demo-1, v1.10b M50):**
- **`public.world_languages`** — the PUBLIC reference catalog: `id` (uuid), `code` (varchar(2) NOT NULL UNIQUE —
  ISO-639-1), `name` (text NOT NULL). Empty on a clean demo (the snapshot didn't carry it) → M50 seeds it.
- **`public.user_languages`** — `user` (uuid FK → users), `level` (int — CEFR-ish 1..5, 5=native), `id` (uuid),
  `world_language_user_languages` (uuid FK → world_languages), `created_at`/`updated_at`. UNIQUE
  `(user, world_language)`. Carries an **AFTER-INSERT trigger** that writes the matching
  `membership_languages` row per the user's memberships — so seed `user_languages` only.
- **`public.membership_languages`** — the trigger-written org-side mirror (`membership_languages` FK → memberships,
  `world_language_membership_languages` FK → world_languages, `level`). The manager Talent-tab languages chart
  reads this. Never seeded directly (the trigger owns it).
