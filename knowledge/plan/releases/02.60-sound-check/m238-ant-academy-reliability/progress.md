# M238 — Progress

Section milestone. Checklist stub from the roadmap In-list.

## Sections

- [x] **#3 Start→404** — wire a chapter-body demo path. **DONE:** the FS-as-published `academy-fs-published-chapter-body` demopatch (rext `demo-stack/patches/` + `apply-academy-fs-published-body.sh` + `ant-academy.sh` wiring), gated on the same `ACADEMY_DEMO_FS_PUBLISHED` env var as the catalog patch. 16 tests green. **Proven live on billion**: a chapter 404→200 (real title/body). (decision D1)
- [x] **#2 language error** — **DONE (verdict, not a code fix):** NOT a distinct code bug. Locale is a `?lang=` query param (no `/it` route); the switcher is a sound EN↔IT toggle; the chapter-language 404 is the SAME backend-null path as #3, **fixed by the #3 patch** (locale-aware FS body). Confirmed fresh + live on billion (`?lang=it` chapter → 200). (decision D2)
- [x] **Academy presence/coverage sweep** — extend to assert chapter-body render. **DONE:** `ANT_ACADEMY_CHAPTER_SECTION` + a general `mustNotInclude` negative marker + the academy-chapter probe (catalog.json slug → chapter → `?lang=it` re-render) in `coverage.spec.ts`; 139 e2e unit tests green; live premise validated on billion. Full billion sweep run → M244 (Fate 2, D4). (decision D3)
- [x] **Delivers** — `corpus/services/ant-academy.md` (chapter-body demo path + #2 language verdict) + `corpus/ops/demo/frontend-tier.md` (the BODY half) + `corpus/ops/demo/demopatch-spec.md` (KB-1: the 2 academy-fs-published rows, inventory 11→13) + `corpus/ops/demo/coverage-protocol.md` (the chapter-body + language sweep extension) updated.

## Live verify (billion, 2026-07-21)
#3 + #2 proven on demo-1: `/chapters/<slug>/` 404→200 (title flips "Not Found"→real chapter), `?lang=it` → 200, `/it` a non-route (expected). See spec-notes.md § LIVE VERIFY.

## Audits
Phase 0b KB-fidelity: **YELLOW** (`kb-fidelity-audit.md`) — central premise ALIGNED; one incidental gap KB-1 (fixed in §4).
