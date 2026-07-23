# iter-10 — decisions (local)

- **D1 — fix the DEFAULT, not just the one call site.** The D13 recipe says "give the RACED test (and any
  sibling that relies on the hardcoded 17700 default) a `_free_port()`." Only the RACED test currently uses the
  default, but defaulting `_reap_with_stubs(port=None)` → `_free_port()` fixes it AND hardens against any future
  caller that forgets the port arg (it can never again pin itself to the ambient cockpit port). The three
  explicit-port callers are unchanged. This is the minimal-surface, maximum-isolation form of the recipe.

- **D2 — prove the fix with :17700 HELD.** This workstation has no live cockpit on 17700, so the bug does NOT
  reproduce here (17700 is free → the RACED test passes with OR without the fix locally). To prove the fix is
  load-bearing, I SIMULATED the collision: bound `0.0.0.0:17700` locally and ran the RACED test — it PASSES with
  the fix (uses a free port), whereas the hardcoded-17700 form would have refused (reap sees 17700 held → 1).
  A host-state-dependent test bug must be proven under the state it depends on.

- **scope: reap-17700 only.** The other 8 durable standing failures (6 `test_cockpit.py` academy+overlay +
  `test_host_prereqs_m215` + `test_purge`) are the M238-D5 standing-8 class (a different failure mode), already
  Fate-2 → M244. Not in this iter's reap-17700/D13 scope; they stay routed. Expanding to them on the cap tik
  would be scope creep.

- **0 platform edits.** A rext test-only change (demo-stack/tests/test_reap.py).
