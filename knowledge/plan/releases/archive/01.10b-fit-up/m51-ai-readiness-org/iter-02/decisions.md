# iter-02 — decisions (local)

**D1 — `narrative: ai-readiness` as the enablement discriminator (not a new YAML field).**
- Context: only the 3rd org (the showcase) should get the `organization_settings` `ai_readiness` gate row; Cervato
  + Solvantis must stay un-enabled. The seeder needs a per-story signal to decide which orgs opt in.
- Options: (a) reuse the existing free-form `StoryOrg.Narrative` field with value `ai-readiness` as the
  discriminator; (b) add a new `StoryOrg.AIReadiness bool` (or an enablement block) to the blueprint schema.
- Choice: (a). The `narrative` field already exists, is free-form (no enum validation), and semantically IS "what
  this org's story is about" — `ai-readiness` is exactly that. Zero schema change, zero migration risk to the
  KnownFields-strict parser, and it reads naturally in the YAML.
- Why not (b): a new boolean is more schema surface for one consumer, and a future org could plausibly want the
  dashboard on for a different narrative — but YAGNI; if that arises, the discriminator generalizes to a set. The
  narrative value is the honest single source ("this org runs the AI-readiness diagnostic").

**D2 — `OrgSettingsSeeder` writes only the AI-readiness gate today, but is named for the general table.**
- Context: `public.organization_settings` is a general per-org feature-flag table (setting/is_enabled/options).
  Nothing in the fleet writes it today. M51 needs exactly the `ai_readiness` row.
- Choice: name the seeder `OrgSettingsSeeder` (surface `org-settings`) for the general table, but scope its M51
  body to the AI-readiness gate. Future org-settings rows (other features) extend the same seeder rather than
  proliferating one-seeder-per-setting. Deterministic id from `(org-id, setting)`; idempotent via
  `CopyRowsIdempotent ON CONFLICT (id)`, backed by the DB's `UNIQUE(setting, organization_id)`.

**D3 — Re-export the roster + cockpit manifest in-place (Phase C re-apply) to give the new manager a seat.**
- Context: the live demo-1 roster (`fake-fapi-roster.json`) + cockpit manifest were exported at demo-up time from
  the 2-story preset; they don't know the new Northwind heroes, so the manager-vantage sweep can't log in as Dana.
- Choice: re-export both from the authoring copy against the updated 3-story preset (a pure blueprint projection,
  no DB — the same id-derivation the seeders use, so the roster seats match the seeded rows), write to demo-1's
  stack-dir paths, restart fake-fapi + fake-bapi (the coverage-protocol "Wrong identity / org on a surface"
  re-apply step). The manager seat Dana resolves with `org_role=admin` (required for the enterprise admin gate).
- Why this is in-scope (not a re-scope): the roster export is the documented re-apply for an identity/seat gap and
  is the only way the strand-1 deliverable (a renderable, sweepable 3rd-org manager seat) is actually testable.
