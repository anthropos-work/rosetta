# iter-108 — the DERIVED repair ledger (never hand-assembled, §5 rule 19)

Derived by `repair_reach_guard.read_ledger()` over `iter-103/raw/` — **the same code path that
grades the repair's reach**, so the repair's input and its grader cannot disagree.

- ledger files parsed: **14** of 14 seat reports
- booked blocks: **48** (1 unparseable)
- distinct primary anchors: **31**

> The guard keys on each block's FIRST anchor, so secondary anchors inside a predicate
> (`:598`, `:431`, `:445`, `:294`, `:196`, …) are repaired but do not appear as separate rows.
> `N = 33` counts distinct anchors INCLUDING secondaries; this table is the reach denominator.

| # | anchor | booked by |
|---|---|---|
| 1 | `corpus/architecture/README.md:21` | r25-E B3, r26-E B3 |
| 2 | `corpus/architecture/ai_architecture.md:95` | r25-E B2, r26-E B2 |
| 3 | `corpus/architecture/ai_architecture.md:99` | r25-E B1, r26-E B1 |
| 4 | `corpus/architecture/architecture_overview.md:59` | r25-F B3, r26-F B3 |
| 5 | `corpus/architecture/architecture_overview.md:307` | r26-F B4 |
| 6 | `corpus/architecture/dependency_map.md:19` | r25-A B1, r26-A B1 |
| 7 | `corpus/architecture/external_services.md:565` | r25-A B2, r26-A B2 |
| 8 | `corpus/architecture/frontend_architecture.md:59` | r25-D B3, r26-D B1 |
| 9 | `corpus/architecture/service_taxonomy.md:47` | r26-D B3 |
| 10 | `corpus/architecture/service_taxonomy.md:166` | r26-D B4 |
| 11 | `corpus/architecture/service_taxonomy.md:509` | r25-D B2 |
| 12 | `corpus/architecture/shared_libraries.md:57` | r25-G B1, r26-G B1 |
| 13 | `corpus/architecture/shared_libraries.md:85` | r25-G B2, r26-G B2 |
| 14 | `corpus/architecture/shared_libraries.md:128` | r25-G B3 |
| 15 | `corpus/services/academy-backend.md:125` | r25-F B1, r26-F B1 |
| 16 | `corpus/services/ai-readiness.md:429` | r25-B B1 |
| 17 | `corpus/services/ai-readiness.md:436` | r25-B B2 |
| 18 | `corpus/services/ant-academy.md:121` | r25-F B4 |
| 19 | `corpus/services/backend.md:148` | r25-E B6, r26-E B6 |
| 20 | `corpus/services/clerk-integration.md:103` | r25-E B5, r26-E B5 |
| 21 | `corpus/services/clerkenstein.md:3` | r26-D B5 |
| 22 | `corpus/services/clerkenstein.md:276` | r25-D B1, r26-D B2 |
| 23 | `corpus/services/cms.md:55` | r26-B B3 |
| 24 | `corpus/services/hiring.md:24` | r25-G B4 |
| 25 | `corpus/services/jobsimulation.md:50` | r25-C B1, r26-C B1 |
| 26 | `corpus/services/jobsimulation.md:146` | r25-C B2 |
| 27 | `corpus/services/jobsimulation.md:208` | r26-C B2 |
| 28 | `corpus/services/next-web-app.md:22` | r25-E B4, r26-E B4 |
| 29 | `corpus/services/sentinel.md:85` | r25-F B2, r26-F B2 |
| 30 | `corpus/services/skiller.md:19` | r26-B B1 |
| 31 | `corpus/services/studio-room.md:67` | r26-E B7 |
