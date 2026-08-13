# iter-02 — intra-iter decisions

## D1 — the row-less-container gap is 3, not 2; and the trio gets a check-local list, not `services.sh` rows

The audit named `fake-fapi` + `fake-bapi`. **`hiring-app` has no row either**
(`gen_injected_override.py:353`) — 13 rows against 16 compose services, which is exactly the
"14 of 16 Up" M256 measured; the audit's own arithmetic only balances with it.

**Rejected: giving the trio `services.sh` rows.** Rows are emitted for **every** project, so `dev-N` and
the main dev stack — which have no Clerkenstein and no hiring app — would false-`down` on every unscoped
verify. That is the M247 skillpath class the corpus **just** removed. Neither fits the table's probe
kinds either (the fake FAPI is TLS-only, so an `http` row would reproduce Defect 2 *inside* `verify.sh`).
The liveness check derives its base set from `service_rows()` so it **cannot drift** from the probe set,
and appends the trio check-locally, gated as the override emits them. The absence is pinned by a test
asserting it is **by design**.

## D2 — check (d) does not fire on Linux, and the fix landed anyway

`TOK-01` known-context #2 forbade assuming in either direction. Resolved by code-read: the failing
client is macOS system `curl` on LibreSSL, while `openssl s_client` on the **same** LibreSSL handshakes
the **same** leaf fully — a `curl` defect, not a LibreSSL or cert one. On a `--public-host` demo the leaf
is a real Let's Encrypt cert from `tailscale cert`.

**So it was never a gate risk on the gate host.** Landed regardless, because it is a real
developer-workstation defect *and* because it removes a host-toolchain dependency from a gate input —
a gate whose verdict depends on which `curl` the operator happens to have is not a gate. The fix is
strictly additive (rung B is the unchanged pre-M257 leg), so it cannot regress the Linux path.

## D3 — B1 and B2 are routed forward, not escalated as user-blockers

Both make the gate **currently unreachable**, which is severe — but Phase 5 §4's user-blocker test is
whether an answer is needed *from the user* that changes what code lands. Neither qualifies: re-pointing
six seeders at canonical tables is rext `stack-seeding` work, and giving `app/studio` an acquisition path
is rext `demo-stack` work. Both are squarely inside this milestone's tooling-only remit and its
**zero-platform-repo-edit** constraint. Phase 5 §4 lists *"surfaced future work / new findings discovered
mid-iter"* explicitly as **NOT** a user-blocker: route forward, continue. They become iter-03's planned
scope and its precondition.

## D4 — the `load1 48.7` reading is recorded as a CANDIDATE re-scope signal, not a fired one

If peak `load1` really reaches 48.7 on this host, HEADROOM **clause 1** (`≤ cores − 2 = 6`) cannot pass
and the gate is unreachable **as written** — a genuine re-scope conversation. But it is **one sample,
taken by an ad-hoc host probe, not by `buildbench`'s sampler**, during a phase (5 parallel Go builds)
that is not a UI lane. Grading a re-scope trigger as fired on an un-probed sample would be the mirror
image of the un-probed-lift dishonesty Phase 3's self-check forbids. **Routed as
`INVESTIGATE-M257-load1-48`, blocking the campaign** — iter-03 answers it with the real instrument
before anything is concluded.

## D5 — the manual `app/studio` unblock was accepted for THIS host, and explicitly does not count as a fix

Copying `cms/studio` → `app/studio` is a gitignored path, no tracked file touched, byte-identical in
shape to what rext's sanctioned `make init-studio` produces for `cms` — so it holds the
zero-platform-repo-edit line. It was accepted because it unblocks measurement now. It is **not**
reproducible by the tooling, so it is recorded as a **hack with a named handler**
(`FIX-M257-app-studio-acquisition`), not as B2's resolution. A stack that only works because someone
ran a manual `cp` is not a stack the gate can be measured on.
