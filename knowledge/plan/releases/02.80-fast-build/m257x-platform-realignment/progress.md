---
milestone: M257x
---

# M257x — progress

## Running ledger

_(iter closeouts append here, newest last)_

- iter-01 (tok/bootstrap): all 5 open questions answered vs platform origin HEAD; authored the absent `corpus/ops/platform-alignment.md` and executed its procedure; found the class's root cause (**pinning disables drift detection** — 11/11 clones `behind: null` while the log says "provably fresh") and its local mechanism (**`migrate-demo.sh`'s hand-maintained 4-tuple** creates the legacy schemas itself, bypassing `repos.yml`); refuted 5 inherited/audited claims by measurement, one of which inverted a planned guard — see iter-01/progress.md

## Routes carried forward

| item | why | target |
|---|---|---|
| `DECIDE-M257-jobsim-schema-ownership` | The exit blocker that created this milestone: platform says cms/jobsimulation/roadrunner own no local schema; rext writes ~15 `jobsimulation.*` tables. **Inherited from M257 iter-03 — this milestone owns it now.** | iter-01+ |
| `FIX-M257-feedback-score-approximation` | Benign between a mirror and its source; **not** benign between two tables claiming to be the same row. | M257x |
| `DOC-M257-studio-in-app` | Corpus says studio-room is CMS-only in 5 places; nothing records `app` embeds it. | M257x |
| `FIX-M257-stacksnap-directus-sequences`, `FIX-M257-directus-coldstart-order` | Carried from M257 iter-02, both platform-shape-dependent. | M257x |
