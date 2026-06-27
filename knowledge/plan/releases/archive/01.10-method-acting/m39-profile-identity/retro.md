# M39 — Retro

## Summary
M39 — the first milestone of v1.10 "method acting" — delivered the three highest-leverage, lowest-effort
**profile-identity** fixes so a logged-in hero (Maya Chen on demo-3) reads as a real person on her profile page:
the **right org name**, a **real role + title**, and a **real face**. **G1** threads each hero's story org
name/slug through the seeder roster (`RosterIdentity`/`BuildRoster`) into Clerkenstein's FAPI org resource
(`RosterEntry` → `DemoUser` → `orgMemberships()`), single-sourced through a new `orgSlugFor` helper so the
roster-carried org and the seeded `public.organizations` row can never disagree — a `DisallowUnknownFields`
**paired struct change** with a no-roster `"Clerkenstein Demo Org"` default fallback (so the single-identity
path stays byte-identical). **G2** backfills `public.user_basic_info` (`job_role_id` + `job_title` + `summary` +
`location`) — **the table the /profile header actually reads**, not `memberships` — from the same resolved hero
role, via an **idempotent IS-DISTINCT-FROM-guarded UPDATE** of the trigger-created row; one UPDATE lights the
header AND the role-gap/role-readiness widgets, no-fabrication preserved (NULL without taxonomy). **G4** replaces
the DiceBear *initials* HTTP URL (a 2-letter disc + an online fetch a sealed demo box can't reach) with a
self-authored **parametric SVG face generator** emitted as an **offline base64 data URI** — deterministic by
uuid, license-clean (authored in-repo, nothing to vendor), ~1 KB. Tooling + docs only — **zero platform-repo
edits.** Close GREEN: 5 findings, 0 blocking (all decision-triage reference tags), merged into
`release/01.10-method-acting`. Code-of-record: `rosetta-extensions` @ tag `method-acting-m39` (c360b4e).

## Incidents This Cycle
None. No P2 flakes, no regressions. The 3 harden passes surfaced **zero bugs** — the build phase's
implementation was correct; the passes only closed test gaps (the G2 backfill error/edge paths, the G1 no-org /
mixed-fallback paths, the G4 SVG well-formedness across seeds + the `FuzzAvatarDataURI` fuzz at 714K execs / 0
crashes) and pinned the cross-write agreement invariants (G2 header-role == membership-role; G1 seeded-org ==
roster-org slug/name; G4 byte-identical avatar across re-seed). The close's adversarial review (Phase 2c) weighed
two further scenarios — AR-1 (the G2 backfill's silent-0-rows case) and AR-2 (a G1 same-name slug collision) —
and both resolved to no-fix-needed (AR-1 depends on the documented per-row AFTER-INSERT trigger invariant with
the loud-error path already covered; AR-2 is a pre-existing, unchanged `OrgSeeder` property). The flake gate ran
5/5 clean (shuffled, both modules). The 5 close findings were all decision-triage reference tags, not defects.

## What Went Well
- **Single-sourcing `orgSlugFor` (D2) made an equality provable, not coincidental.** Extracting the slug rule
  (was inline in `org.go`) into one helper both `OrgSeeder` and `BuildRoster` call means the top-bar org and the
  DB org are derived identically — a class of "the slug drifted" bug eliminated by construction, and the
  `TestG1_SeededOrgSlugAndNameMatchRoster` integration test pins it.
- **The G2 idempotency guard (D5) reused the established pattern.** The `IS DISTINCT FROM` guard makes a no-change
  re-seed a true 0-row no-op, mirroring the COPY ON-CONFLICT + casbin WHERE-NOT-EXISTS patterns the M17 re-run
  contract already established — validated live on demo-3 (1st UPDATE 1 row, 2nd 0).
- **G4's "author the asset, don't vendor it" call (D7) was the strictly cleanest answer.** A self-authored SVG
  generator is offline (data URI, no fetch), deterministic (hashed), and license-clean (no third-party asset at
  all) — it dodged the bundled-photo-set's sourcing/vendoring/license burden AND the DiceBear URL's offline
  failure in one move, and `FuzzAvatarDataURI` proved the generator total over arbitrary input.
- **The two-repo discipline held cleanly.** Code in the rext authoring copy @ tag `method-acting-m39`; the
  rosetta doc-half on `m39/profile-identity`; go.mod/go.sum byte-identical (supply-chain GREEN); all 3 offline
  alignment gates 100%/100% throughout — exactly the v1.9 M34–M38 pattern.

## What Didn't
- **The alignment gate runner needed a `GOTOOLCHAIN=local` workaround at close.** The installed `go1.25.11`
  satisfies the `go.mod` toolchain directive, but the gate scripts export `GOSUMDB=off`, which makes Go's
  (no-op) toolchain-switch machinery try to verify the already-installed toolchain module and fail. Setting
  `GOTOOLCHAIN=local` skips the switch. Minor friction, noted in `metrics.json` for the next gate run; not a
  code or alignment issue.
- **A momentary mis-run of the multi-identity gate.** The first close attempt at gate 3 used the JS/FAPI runner
  (`jsfapirun`) instead of the multi runner (`multirun`), producing spurious `mirror=null` "failures"; corrected
  by reading the CI workflow's gate invocation. The gate itself was always 9/9 — operator error, not a regression.

## Carried Forward
- None from M39 itself (0 deferrals; the 2 Out-list items are Fate-2, already owned by release siblings — see below).
- **Fate-2 (confirmed, no edit):** work/education history + skill depth → **M41**; the library + activity-feed
  serve-grant → **M40**.
- **Standing backlog (unscheduled, orthogonal to v1.10, unchanged):** DEF-M10-01, DEF-M21-01, M25-D9.

## Metrics Delta (from metrics.json)
- **Go test funcs:** stack-seeding 444 → **462** (+18); clerkenstein 259 → **264** (+5). Baseline = v1.9 final close.
- **Coverage:** stack-seeding/seeders 96.1%, clerk-frontend 86.7%; all M39-touched functions at 100% (residual
  package sub-100% is pre-M39 error branches, out of scope). No >2pp drop.
- **Flake:** 0 (5× shuffle, both modules).
- **Alignment gates:** 100%/100% on all 3 offline surfaces (Go 22/22, JS 9/9, multi 9/9).
- **Supply chain:** GREEN — 0 new deps; go.mod/go.sum byte-identical.
