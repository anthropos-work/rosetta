# iter-16 progress

**Type:** tik (production-fix). Active strategy: **TOK-10**. P4 — REAL-photo avatars (menu==profile) +
org logo. **Two planned parts: (1) avatar real photos, (2) org logo.** Part (2) landed; part (1) is a
USER-BLOCKER (licensing/consent).

## Phase B — source + license-vet the real-photo avatar set (USER-BLOCKER)
- The user decision (design-plan §USER DECISIONS #2) is **LICENSED real-person STOCK PHOTOS — NOT
  synthetic, NOT illustrated** — CC0 / public-domain / permissively-licensed REAL portraits.
- The run prompt's HARD rule: *"Use ONLY clearly license-clean real photos — NEVER copyrighted/
  unlicensed/identifiable-without-consent images. If you cannot find a clearly license-clean real-photo
  source, STOP and surface it as a user-blocker."*
- **Investigation (web search, build-box network):** the available real-portrait sources split into two
  classes, and NEITHER is clean on BOTH axes the prompt requires:
  - **CC0 / Unsplash / Pexels / Pixabay real portraits** — the COPYRIGHT license is clean (free for
    commercial use, no attribution), BUT they carry **NO model release** → the depicted REAL person's
    CONSENT to being shown as a fictional employee ("Maya Chen, DevOps Engineer @ Cervato Systems") is
    UNKNOWN/unverifiable. That is exactly the **identifiable-without-consent** case the prompt forbids.
  - **Released stock (Adobe/Getty/Shutterstock with a model release)** — the standard model release
    explicitly **prohibits "sensitive/misleading use"**: depicting a recognizable real model as a
    fictional persona, or in a way implying false identity/endorsement, is a forbidden use under those
    releases (and the assets are licensed, not free/CC0 — violating the "never copyrighted" clause).
  - **FFHQ / face datasets** — the well-documented biometric/consent problem (MegaFace fallout); the
    depicted real people did not consent to commercial fictional-persona reuse.
- **Conclusion:** a source that is BOTH (a) clearly license-clean AND (b) consent-clean for
  fictional-persona depiction of a REAL identifiable person **does not exist** among the available
  sources. The two user-facing constraints (a REAL photo AND no identifiable-without-consent) cannot be
  jointly satisfied by a stock-photo source. Per the prompt: STOP, do NOT use questionable images,
  surface a **user-blocker** (SEVERITY: blocker). D1.

## Phase C — org logo (the CLEAN half — lands now)
- `orglogo.go` (NEW): `orgLogoDataURI(name)` — a deterministic, offline-safe, **license-clean** monogram
  data URI (a brand-tinted rounded tile + the org's white initials, e.g. "CS" for Cervato Systems).
  Mirrors avatar.go's data-URI shape. A GENERATED mark is the honest choice for a FICTIONAL org (no
  trademark to infringe, no real org whose consent is at stake — the design-plan default decision #4). D2.
- `org.go`: OrgSeeder now writes `logo_url` (the column already existed, NULL before) — lights the
  in-app org logo (next-web reads `organizations.logo_url`).
- Tests: `seeders_test.go` OrgSeeder column-count + the logo-is-a-data-URI assert updated; NEW
  `orglogo_test.go` (initials derivation + decodable SVG + determinism + distinctness). `go test ./...`
  GREEN; vet clean.
- **NOT done (coupled to the blocked avatar):** the Clerkenstein userRes/orgRes IMAGE threading (the
  top-MENU avatar + the top-menu org glyph). The userRes avatar image is blocked on the avatar-licensing
  decision, and the userRes+orgRes image threading rides the SAME `resources.go` change + alignment-golden
  re-capture — so doing only the orgRes half would leave a half-wired Clerkenstein change + a golden churn
  for a feature that can't ship its companion. Held as ONE unit behind the user decision. D3.

## Phase D — re-measure (demo-3 reset + re-seed with the new binary)
- Reset + clean re-seed (51734 rows, isolation clean prod=false). Both orgs now carry the monogram logo:
  `Cervato Systems` + `Solvantis` → `data:image/svg+xml;base64,…` (450-char SVG, decodes to a valid "CS"/
  "SO" monogram tile). Maya still DevOps Engineer, 12 verified + 18 claimed intact.
- The **in-app org logo** (`organizations.logo_url`) renders. The **top-menu** org glyph + the menu avatar
  remain unfixed (the Clerkenstein image threading is held behind the avatar blocker).

## Close — 2026-06-25
**Outcome:** P4 part (2) org logo landed (in-app `organizations.logo_url` monogram, both orgs). P4 part (1)
real-photo avatars + the coupled Clerkenstein userRes/orgRes top-menu image threading are a **user-blocker**:
no source is BOTH license-clean AND consent-clean for fictional-persona depiction of a real identifiable
person.
**Type:** tik (production-fix)
**Status:** closed-fixed-partial
**Gate:** NOT MET (P4 part 2 of P0–P8; the avatar half + P5–P8 remain)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: y (avatar
real-photo licensing/consent — no clean source) — (5) cap-reached: n — (6) protocol-stop: n — Outcome: exit-4
**Decisions:** D1 (avatar-licensing user-blocker), D2 (org-logo monogram), D3 (Clerkenstein image threading
held with the avatar) — see ./decisions.md + milestone-root decisions.md AVATAR-LICENSING-BLOCKER.
**Routes carried forward:** P4 part (1) avatar real photos + Clerkenstein userRes/orgRes image threading →
USER DECISION (3 options in the milestone-root decisions.md). P5 (Sentinel-reload), P6/P7/P8 — later.
**Lessons:** "license-clean" for a REAL-person avatar is a TWO-axis test — copyright AND model-consent for the
specific (fictional-persona) use. A CC0/Unsplash photo passes the copyright axis but FAILS the consent axis
for depicting a real identifiable person as a fictional employee. For a FICTIONAL org, a generated monogram
passes both axes trivially (no real entity, no trademark, no consent) — which is why the org-logo half is
clean while the avatar half is blocked.
**rext:** commit `21a1ad9`, tag `method-acting-m42e-iter16`.
