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
