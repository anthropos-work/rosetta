---
iteration_type: tik
status: closed-fixed
milestone: M253
iter: 2
active_strategy: TOK-01
created: 2026-07-24
---

# M253 iter-02 — tik (the two demopatches + FCP runner + ladder + first measurement)

**Active strategy reference:** TOK-01 (shell-before-awaits + no-thirdparty demopatches on the M249 ladder + a
net-new FCP runner). This is the first tik under TOK-01.

**Cluster / target identified:** TOK-01's next-tik direction — author both demopatch manifests, the FCP runner,
extend the `build_frontend_studio_desk` ladder + fingerprint, rebuild the studio image on demo-2, measure.

**Hypothesis:** injecting the `.page-skeleton` DOM synchronously before the boot awaits drops
skeleton-visible from the ~4.7 s baseline to well under the 1000 ms gate.

**Expected lift:** skeleton-visible p95 from 4669 ms → < 1000 ms.

**Phase plan:** author (3 lanes: patches / FCP runner / ladder) → rebuild studio image on demo-2 → measure
5 cold loads → grade.

**Escalation conditions:** demopatch G2 refuse (re-author anchor); FCP still ≥ 1 s after the reorder (triage
the residual per latency-budget arithmetic); a login bounce (session/cookie fault) → user-blocker.

**Acceptable close-no-lift outcomes:** none expected — this is a landing tik.
