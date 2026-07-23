# iter-09 — decisions (local)

- **D1 — measure a RE-SEED-INDEPENDENT gate part against the current billion.** The gate-(h) cold reset-to-seed
  (re-pin m244 + re-bake + re-seed) is the critical path for gates b-live/c/f/h, but it is a large high-stakes op
  (F1–F12 host minefield, never kill a mid-build) best given a fresh run's focused attention. Gate (d) (the anon
  academy render) is INDEPENDENT of the m244 content-story changes, so it is verifiable against the current m243
  billion NOW — a cheap way to advance a gate part without the big op. This is why iter-09 targeted gate (d).

- **D2 — root cause characterized; the fix is an ant-academy demopatch, routed not landed.** The anon `/library`
  grid renders empty because it reads a TENANT-FILTERED catalog (`ServerCatalogContext`), which is empty for an
  anonymous demo user (no tenant entitlement) — by design, `publicSource.js` refuses to import the FS catalog to
  avoid leaking cross-tenant content to the browser. `/free` renders because it reads a public/tier source. The
  fix — feed the FS catalog to the anon `/library` grid on a demo — is an ant-academy code/config change, so it
  must route through a sha-pinned demopatch (0 platform edits), a substantial follow-up. Given the remaining
  budget + that iter-08 already consumed a deep demopatch investigation, this iter MEASURES + characterizes +
  routes rather than rabbit-holing into a second demopatch build. closed-no-lift with a documented falsification
  is a first-class outcome (the protocol working), NOT a deferral against the three-fate rule.

- **0 platform edits.** The measurement is read-only (Playwright render against billion). The routed fix is a
  demopatch, never an edit to the ant-academy repo.
