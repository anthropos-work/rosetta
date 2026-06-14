# M30 — Retro

_Field-bake: build a compliant secret dir from stack-dev + prove it. The 4th + FINAL milestone of v1.6 "stage
door". Two-repo (rosetta record + corpus reconciliation; ext field fixes @ tag `stage-door-m30`). Closed
2026-06-14._

## Summary

M30 proved the whole v1.6 secret-provisioning mechanism end-to-end on a real stack. **Part 1:** assembled a
compliant, gitignored `.agentspace/secrets` dir from current stack-dev (5 repo `.env` files cp'd into the
DNA-driven reader layout, values-blind; ant-academy source filled with the shared Clerk publishable key via a
values-blind line-append), and `check` scored **Critical == 100%** on both dev and demo (exit 0). **Part 2
(live, with the user's go-ahead):** a fresh **demo-3** was brought LIVE from that assembled source — provision
wrote 26 / blanked 2 / skipped 0, the bring-up pre-flight scored Critical **100%**, and the demo came up with
**17 containers** (backend tier + UI tier: next-web + studio-desk + ant-academy native), the full UI-inclusive
auto-verify exit 0. The **observable-behavior gate** (provision → Critical 100% → stack UP) is **MET LIVE**.

The bake caught + fixed **2 real release bugs** Fate-1 (parallels v1.5 M25's 4): (1) `sentinel/DB_CONNECTION`
was critical/required but is compose-injected config (hardcoded `environment:` entry, always overrides
`env_file`; sentinel never reads it from a `.env`, and no `sentinel/.env` exists on stack-dev) → reclassified
`waived-config` + a regression assertion (was falsely failing the gate at Critical 84.6%); (2) the demo
bring-up only *checked* coverage but never *provisioned* — and `preflight.sh` resolved its source path one
level too shallow (doubled `.agentspace/.agentspace/secrets` → the demo-aware gate **always silently skipped,
exit 2**) → added a non-fatal provision step (`DEMO_NO_PROVISION=1` opt-out) that writes stack-demo's per-repo
`.env` from the source + repoints the run's base env, and corrected the path to `EXT_ROOT/../..`. Tag
`stage-door-m30` @ `29c922b`.

The close reconciled the rosetta milestone record + `corpus/ops/secrets-spec.md` to the executed live bake
(the docs had been written for a "Part 2 held" outcome that was then overtaken by the live run): 4 findings,
all Fate-1 (3 docs + 1 decision-triage), deferral audit GREEN.

## Incidents This Cycle

None at close. No P0/P1/P2 incidents, no flakes (Go 5/5 `-race -shuffle`), no regressions. The 2 field bugs
the bake caught are the milestone's *deliverable* (a field-bake exists to surface them), not close-cycle
incidents — both were fixed Fate-1 during the bake. One P2-class **documentation-drift** finding at close:
`progress.md`/`spec-notes.md`/`secrets-spec.md` were authored describing a "Part 2 held" outcome + only 1 bug,
but Part 2 ran live and a 2nd bug was found — the record was stale relative to the executed work. Reconciled in
Phase 7 (no code impact; doc-only).

## What Went Well

- **A live field-bake earns its keep.** The dry-run (`provision --dry-run` planning 26/2/0) was green, but the
  *live* bring-up exposed two defects a dry-run structurally cannot: the bring-up never actually provisioned
  (it ran from the operator's dev env), and the pre-flight's default source path was doubled so its gate
  silently skipped. Both are exactly the "scored green but never moved the bytes / never ran the gate" class
  that only an end-to-end live run surfaces.
- **Defense-in-depth on the prod `DIRECTUS_TOKEN` held under live verification.** Provision writes the strip
  family BLANK on the non-prod target AND the injection override strips it at compose-emit — and the live
  check confirmed the len-32 prod token was armed in **ZERO** containers (cms blank; studio-desk = a local
  len-27; graphql = the router len-129, not prod). The fix16/17/M28 non-rearm class survived a real bring-up.
- **Values-blind survived assembly + provision + live run.** The compliant dir was assembled by `cp` /
  line-append only; provision stdout is key NAMES + counts only; no value entered any output, log, or commit,
  and `.agentspace/secrets` stayed gitignored (verified). The whole v1.6 hard-safety invariant held end-to-end.
- **The waived-config catch is a clean conceptual win.** `sentinel/DB_CONNECTION` is the textbook case the
  DNA's waived class exists for — a key the runtime gets from compose `environment:`, never from a `.env`.
  Reclassifying it (vs hand-creating a `sentinel/.env` no runtime reads) keeps the gate honest at 100% without
  weakening the anti-vacuous-100 guard (12 required+critical genes remain).

## What Didn't

- **The milestone record was written ahead of the live run and went stale.** Part 1 ran first and the docs
  were committed describing Part 2 as "held / PENDING" + 1 bug; the live Part 2 (+ the 2nd bug) then overtook
  them but the record wasn't updated in the build commits. Caught + fully reconciled at close (Phase 7), but
  it's a reminder that when a milestone executes in two passes, the record should be reconciled in the pass
  that completes it, not left for the close to catch.

## Carried Forward

- **Nothing M30-originated.** As the final v1.6 milestone, M30 routes nothing forward in-release — the
  build-from-stack-dev validation was M30's own scope and it landed (live).
- **Inherited release-level backlog** (DEF-M10-01 / DEF-M21-01 / M25-D9) — orthogonal to secret provisioning;
  re-signed at the v1.5 close, carried unchanged. These surface again at the **v1.6 `/close-release`**
  re-audit.
- **v1.6 "stage door" is ready for `/developer-kit:close-release`** — all 4 milestones M27–M30 are closed.

## Metrics Delta

(from `metrics.json`)
- **Findings:** 4 (0 scope · 0 code-quality · 3 docs · 1 decision-triage) — all Fate-1.
- **Field bugs caught + fixed Fate-1:** 2 (sentinel waive-config; demo provision-wiring + preflight path).
- **Go tests:** 1027 → **1027** (+0 — the M30 ext regression is a sub-assertion inside an existing test func).
- **Python tests:** 459 → **459** (+0); demo-stack's own suite 99 pass.
- **Flake count:** 0 (Go 5/5 `-race -shuffle`).
- **Observable-behavior gate:** MET LIVE — demo-3, 17 containers UP, Critical 100%, prod DIRECTUS_TOKEN in 0 containers.
- **Ext code:** 2 field-fix commits on `m30/field-bake` @ tag `stage-door-m30` (head `29c922b`); orchestrator finalizes the ext side.
- **Deliverables:** the proven (gitignored) `.agentspace/secrets` reference dir + the field-bake record + the corpus reconciliation + the skill-doc tag bump.
