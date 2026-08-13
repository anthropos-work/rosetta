# iter-222 — decisions

## D-M257x-222-1 — the user redirected the milestone's target selection (2026-08-09)

**Verbatim:** *"the goal remains alignment and be able to build a working stack with the new platform
repos (only the remaining ones that are still part of it)."*

**Binding consequence, recorded so it survives this run's context:** iters target **the platform, and
the claims rosetta/rosetta-extensions make about it**. Instrument work — registries, censuses of
censuses, number-matchers, test runners — is de-ranked, not cancelled (`§8` iter-110: grade a strategy
leg by leg). `TOK-08` still holds as the *method* (census, do not sample); the redirect changes the
*subject set* the method is pointed at.

**"Only the remaining ones that are still part of it" is machine-readable**, and this iter uses it as
such: it is `repos.yml` @ platform `origin/main`, plus the two sanctioned non-`repos.yml` clones
(`platform` itself, and `ant-academy` per M49 #5).

## D-M257x-222-2 — pin freshness is DISCLOSED, not auto-advanced

The canonical pin is 3-of-6 behind `origin/main` (app 28, next-web-app 12, ant-academy 9). Advancing it
is **not** this iter's repair. The pin is M246's *"reproducible barrier"* — the topology gate clauses 1+2
were proven against. Bumping it silently would mean the tooling claims a proven topology it has never
built. The repair is to make the freshness **derivable and dated** instead of **asserted as current**.

Routed forward as `ROUTE-M257x-222-pin-advance-needs-a-reproof` — the advance is legitimate work, and it
is gate-clause-1 work (a cold bring-up at the newer topology), not manifest work.

## D-M257x-222-3 — the fetch finding is a DISCLOSURE, not a regression

`anchor_construct_guard` + `repair_postcondition` went RED after a `git fetch origin` in
`stack-demo/app` moved `origin/main` from `ad9f3c498` (= the clone's own HEAD) to `3eaadae68`. No corpus
file changed; no guard changed. **Grade the direction:** the guards stopped being blind. The 9 anchors are
routed to iter-223, not repaired here, because the first question — is each site ref-pinned (`§5` rules
41/44, in which case the guard is reading a ref-scoped claim at the wrong ref) or unpinned (corpus rot)? —
is itself the iter.

**Provenance of the 5 dropped pin keys**, preserved so the repair loses nothing:

| repo | dropped sha |
|---|---|
| cms | `93e6aa354a96a4b5c9b52ead4dd4c27351bbb4a1` |
| jobsimulation | `5d3003f9f133df9dd68acd21b3e336ae27824cd4` |
| storage | `7696605425a42e41738dfa1fd7413b65eb785689` |
| messenger | `d41029217828a5ac737920cac82f22ab363e82a6` |
| roadrunner | `87d8d44382ef07a9f165869530cbac9e5e0a4332` |

The workspace copy `stack-demo/clones.pin.json` is **left as found**: `ensure-clones.sh` never clobbers an
operator's workspace pin, and it lives in a git-ignored ephemeral workspace. It is re-seeded from the
canonical pin on any box that does not have one.
