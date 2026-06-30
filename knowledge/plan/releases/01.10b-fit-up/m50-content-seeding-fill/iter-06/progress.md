# iter-06 — progress

**Type:** tik (under TOK-01) — the run-2 refined target: languages + cert-coverage fill + the D4/F1
manager-manifest strengthening. Per `coverage-protocol.md` Phase A–E.

## Work

- **Phase A (re-survey + calibrate).** Confirmed the run-1 `(0,0)` manager gate was BLIND to two M50-own
  annotation gaps: `world_languages`/`user_languages`/`membership_languages` all 0 rows; `user_certifications`=9
  (hero-only) across 340 members. Calibrated the live manager render: `/enterprise/members` shows
  Member/Email/Role/Job Role/**Location**/Status/**Join Date**/Content Access/**Last Activity Date**/Tags (the
  iter-02 fill renders; first row "Stockholm, Sweden" / "Jan 28, 2026"). The Workforce **Talent Pool** tab
  renders the languages + certifications as distribution CHARTS (cards "Languages spoken" + "Certifications",
  data in the chart SVG `<text>`) — so the manifest can assert them after a tab-click.

- **Phase C (fix — seeders).**
  - **NEW `MemberLanguagesSeeder`** (`stack-seeding/seeders/member_languages.go`): populates `world_languages`
    (a curated ISO-639-1 standard catalog, 16 entries) + per-member `user_languages` (1-3 distinct languages:
    a location-coherent native [level 5] + near-universal English + an occasional third), iterating the SAME
    population index space UsersSeeder writes (ALL members, not hero-only). The DB AFTER-INSERT trigger
    `on_insert_user_languages_insert_membership_languages` fans each `user_languages` row out to
    `membership_languages` — so we seed only `user_languages`. Registered in `main.go` (DAG level 2, after users).
    +4 unit tests (catalog + per-member coverage + determinism + native-coherence).
  - **EXTENDED `CertificatesSeeder`** to ~45% role-coherent supporting-population coverage (1-2 certs each,
    mirroring UsersSeeder's deterministic role assignment for bank/skill coherence; heroes still get their §B
    2-3; managers still excluded). Updated 2 existing cert tests + added 2 (member-coverage determinism + bound).
  - Full `stack-seeding` suite GREEN; `go build`/`go vet` clean.

- **Phase C (fix — harness + manifest strengthening, the D4/F1 reconciliation).**
  - `coverage-manifest.ts`: NEW `preAssert` field (a tab-click-before-assert step) + NEW `textMatch`
    realContent kind (an OR-expressive regex-over-visible-text assertion). Strengthened the manager manifest:
    `/enterprise/members` `members-roster` now requires the **Location** header; NEW `members-location-values`
    asserts a real seeded city renders (textMatch over the seed's city alphabet — robust to page sort);
    `/enterprise/workforce` gains `talent-languages` + `talent-certifications` sections (preAssert clicks the
    Talent Pool tab; assert the "Languages spoken"/"Certifications" cards + a seeded distribution token).
  - `section-assert.ts`: `runPreAssert` (the tab-click executor; non-fatal so a tab that won't open FAILs
    honestly) + the `textMatch` verdict path. Wired into `coverage.spec.ts`'s per-section loop.
  - Manifest unit spec updated for the new kind/field; all 17 GREEN.
  - Added 2 calibration probe specs (`calibrate-manager` + `calibrate-talent`) + a `.gitignore` for their
    transient output dumps. Fixed a SIDE-BUG in `run-coverage.sh`: it forwarded consumed positional args
    (N/vantage/key) to playwright as filename filters, so the new `calibrate-manager.spec.ts` matched "manager"
    and ran in parallel with the sweep — corrupted the run. Fixed with a guarded shift-consume loop.

- **Phase C (re-apply).** Re-seeded demo-1 (`--stack demo-1 --seed presets/stories.seed.yaml`): world_languages
  16, user_languages 747 (ALL 340 members), membership_languages 747 (trigger fan-out OK), user_certifications
  9→236 (158 distinct users ~46%). Audit clean (prod=false, no shared/external writes). **M17 idempotency PROVEN:**
  a 2nd seed wrote 0 new rows; DB counts unchanged.

- **Phase D (re-sweep).** Manager re-sweep with the STRENGTHENED manifest (cap=300):
  **`reachable=69 failingSections=0 escapes=0 personaFailures=0 notReached=0 frontier=EXHAUSTED gateMet=True`,
  cross-port studio-desk follow OK.** The strengthened sections ALL PASS real-content (not skipped):
  `talent-languages` (4329 chars, "Languages spoken" + English), `talent-certifications` (4329,
  "Certifications"), `members-roster` (3113, incl the "Location" header), `members-location-values`
  (**textMatch 20 ≥ 1** — 20 seeded cities render). The gate is MET on the manifest that PROVES the gaps —
  not the run-1 blind manifest.

## Close — 2026-06-30

**Outcome:** The two run-2 annotation gaps the run-1 gate was BLIND to are now FILLED and PROVEN.
`MemberLanguagesSeeder` (NEW) + the `CertificatesSeeder` member-coverage extension materialize on demo-1
(world_languages 16, user_languages 747 across ALL 340 members, membership_languages 747 via the DB trigger,
user_certifications 9→236 / 158 distinct users), M17-idempotent (2nd seed = 0 new rows). The manager coverage
manifest is STRENGTHENED (new `preAssert` tab-click + `textMatch` OR-assert harness capabilities) to assert the
members Location column + a real city value + the Talent-tab "Languages spoken" + "Certifications" charts. The
**manager re-sweep on the strengthened manifest is GREEN** (reachable=69, frontier-exhausted, (0,0), persona 0,
cross-port OK). The employee vantage is provably unaffected (0 employee sections use the new fields; the existing
text/count/both assertion paths are byte-unchanged — only new additive branches). **M42 gate MET both vantages
on the STRENGTHENED manifest** (employee run-1 green + manager run-2 green).
**Type:** tik
**Status:** closed-fixed
**Gate:** MET
**Phase 5 grading:** (1) gate-met: y — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: exit-1
**Decisions:** D1 (seed user_languages only — trigger fans out), D2 (ISO-639-1 standard catalog, not fabrication), D3 (~45% role-coherent cert coverage), D4 (preAssert + textMatch manifest strengthening = the D4/F1 reconciliation), D5 (side-bug: run-coverage.sh arg-forwarding).
**Side-deliverables (if any):** `run-coverage.sh` arg-forwarding fix (consumed positionals leaked to playwright as filename filters → the new calibrate-manager spec ran in parallel + corrupted a sweep) — a separate tooling-correctness fix (D5), committed with the harness changes.
**Routes carried forward (to close/harden + M53 — NOT gate-blocking the warm metric):**
- **COLD reset-to-seed acceptance** (the explicit exit_gate): a fresh `/demo-up` (all M50 seeders + fixes reproduce from tooling) + both-vantage sweeps on the strengthened manifest → confirm (0,0) on COLD. M53 owns this; close-milestone runs it or surfaces.
- Re-pin the consumption clone (`stack-demo/rosetta-extensions`) to the `fit-up-m50` tag at close (it carries the iter-04/05/06 fixes).
- AI-keys policy (F7) + academy menu-link/non-anonymous-session (F6): decision deliverables — for close/M51.
**Lessons:** (1) A green coverage gate is only as honest as its manifest ASSERTS — run-1's `(0,0)` passed blind to two M50-own gaps because the manifest never asserted languages/certs/member-fields. Always cross-check the gate's assertions against the milestone's intent, not just the (0,0). (2) Tab-gated content (in-page-on-click) + paginated grids need new harness primitives (`preAssert` tab-click + `textMatch` OR-assert) — keep them additive (new branches gated on new fields) so other vantages stay provably unaffected. (3) `run-coverage.sh` forwarding consumed positional args as playwright filename filters is a latent footgun — a new spec whose filename contains a vantage token (e.g. "manager") gets silently co-run; consume positionals before `"$@"`.
