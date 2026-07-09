# M208 Retro — Re-sync & merged-schema ground-truth

**Closed:** 2026-07-08 · `closed-complete` · `section` · complexity medium.
**Outcome:** the load-bearing foundation of v2.1 (mirrors v1.10b M47). Re-synced both stacks to the merged
platform (`app` `a848cccb→c3c45e01` v1.334.1 — the 86-commit merge pull; `platform` `5e1ae6b→0808b92` — skiller
gone from compose/repos.yml/Make; sibling set current), removed both vestigial `stack-*/skiller/` clones, retired
the #1 release risk GREEN via a live containerized de-risk on stack-dev, and pinned the authoritative merge
fact-sheet in `corpus/services/backend.md` (+ a stub banner on `skiller.md`). Zero platform-repo edits.

## Summary
The committed rosetta diff is **100% documentation** (fact-sheet + milestone records + two sibling-overview
routing edits + state.md; zero `.go`/`.ts`/`.sh` files). The substantive engineering was the **operational**
re-sync + live de-risk on the gitignored `stack-*/` workspaces — proven, not committed to the corpus. The
fact-sheet is deliberately minimal (the grading contract M209/M210/M211 measure against); the full pre-merge
corpus body-flip is M210's chartered scope, not M208's.

## Metrics delta (from `metrics.json`)
- **Tests:** N/A — docs+ops milestone, zero code/test surface in the committed diff. Inherited v2.0 baseline
  unchanged (rext Go 1745 funcs, 10 live Playthroughs). HARDEN correctly N/A.
- **Live de-risk:** cold `make up` rc=0 · 4-subgraph compose · **no skiller container** ·
  `SKILLER_RPC_ADDR=http://backend:8083` · clean-slate `reset-db`+`migrate` builds the full `public` taxonomy with
  no skiller schema on a clean DB · prod `public.skills WHERE organization_id IS NULL` = **42,790**. Verdict GREEN.
- **Close findings:** 0. **Deferral audit:** GREEN (5 single, 0 repeat, 0 escape-hatch). **Flake:** 0 (no tests).

## Incidents this cycle
- **No P0/P1/P2. No regressions. No flakes.** Two bring-up findings surfaced during the de-risk and were routed
  forward (both user-accepted), not incidents against M208's own committed deliverable:
  - **Finding 1 (M25-D9 class):** a clean cold `make reset-db` doesn't bootstrap the `extensions` schema
    (pgvector + `pg_trgm` + resolvable `gin_trgm_ops`) before `make migrate`, so the merged vector/trigram
    migrations fail on an empty DB; plus a PG-readiness race. Did NOT fall out as a trivial Fate-1 (a
    bring-up-tooling requirement) → **Fate-3 M211** (overview pinned) + M209 Risk-2 cross-ref.
  - **Finding 2:** stack-dev's hand-assembled `.env` lacks `INVITATION_HMAC_SECRET` (backend `Exited(0)` on the
    containerized cold run) — a per-stack `.env` completeness gap, not merge-caused → **Fate-2 M211 /
    `/stack-secrets`**.

## What went well
- **The #1 release risk retired early.** The 86-commit `app` pull + migration re-run (the fit-up M47 ⚠ class) was
  the biggest unknown; the live containerized de-risk proved the merged `public` schema migrates clean from an
  empty DB with no skiller schema — so M209/M210/M211 now grade against a *proven* merged state, not a hypothesis.
- **The firewall predicate held.** `public.skills WHERE organization_id IS NULL` (42,790) empirically confirms the
  snapshot public-predicate survives the merge — de-risking M209's recapture safety ahead of time.
- **Clean scope discipline.** The minimal fact-sheet (not M210's body-flip) + the parked-then-restored native-dev
  override + the reverted dev-secret probe all kept M208 inside its charter with zero platform edits.

## What didn't (go as smoothly)
- **The opportunistic M25-D9 didn't land as Fate-1.** The overview hoped it might fall out on the re-migrate path;
  it surfaced only on the clean-slate run and proved to be a bring-up-tooling requirement, not a one-line tweak —
  correctly re-fated to M211 rather than scope-crept into M208.
- **The clean-slate authoritative run couldn't verify the backend subgraph LIVE** (Finding 2 kept backend down);
  the complementary existing-volume run covered that gap (router serves the 4-subgraph federation end-to-end).

## Carried forward
- **Finding 1** (extensions-bootstrap + PG-readiness) → **M211** (Fate-3; `overview.md` pinned) + M209 Risk-2 xref.
- **Finding 2** (`INVITATION_HMAC_SECRET` dev `.env` gap) → **M211 / `/stack-secrets`** (Fate-2).
- **KB-1/2/3** (pre-merge corpus prose: backend.md consumer claims / skiller.md standalone body / 5→4 subgraphs) →
  **M210** (Fate-2, the chartered corpus body-flip).
- **Consumption-clone re-pin / rext tag** stays `v1.10.1` until M209 tags rext `v2.1` (push-gated KEEP).
