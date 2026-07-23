---
iter: 10
milestone: M244
iteration_type: tik
iter_shape: cleanup
status: closed-fixed
active_strategy: TOK-01
created: 2026-07-22
---

# iter-10 — inherited carry: the reap-17700 standing-9 test-isolation fix (M239-D13)

**Type:** tik (cleanup shape — discharges an inherited carry; test-isolation debt, 0 production runtime code).
Under **TOK-01**. Chosen for the final tik of run 4 as a clean, completable, low-risk win: the gate-(d)
`/library` fix (iter-09 finding) is a diffuse multi-file ant-academy demopatch (rabbit-hole risk on the last
tik), and the corpus 29→47 / 39→40 doc reconciliation is premature before the live 47/47 proof (gate h). The
big gate-(h) cold reset-to-seed is deferred to a fresh run's focused attention.

## Target
`DEF-M239-01`-sibling **reap-17700** (M239-D13): `test_reap.py::test_a_RACED_listener_exits_silently` false-fails
on a box with a live demo-1 cockpit on :17700 (test-isolation collision, 0 product/reap.sh defect). Fix per the
D13 recipe: isolate the test from the ambient 17700 via a guaranteed-free port.
