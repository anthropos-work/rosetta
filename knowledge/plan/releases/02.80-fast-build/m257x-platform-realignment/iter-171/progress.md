**Type:** tik — under `TOK-08` (*census the mechanical classes; stop sampling them*), applied to a
**property** rather than a claim class.

# iter-171 — the bind that waits on a name nobody reads

## Phase A — reproduce, then dump the stack instead of theorising

iter-170 left exactly one runner disagreement it could not attribute to imports: two
`test_cockpit.TestContentTabMainWiring` tests fail under 3.14/unittest and pass under 3.9.6/pytest. Both
reproduce at HEAD, **including in isolation** — so it is not test-order pollution, which was hypothesis one.

The two readings on the table were *"the 2 s bind window is too tight"* and *"3.14 is flaky."* Neither is
falsifiable by argument, and both are the comfortable answer. What settled it was one `sys._current_frames()`
dump of the blocked worker thread:

```
cockpit.py:1768      httpd = ThreadingHTTPServer((a.host, a.port), handler)
  socketserver.py:457  self.server_bind()
    http/server.py:150   self.server_name = socket.getfqdn(host)
      socket.py:817        hostname, aliases, ipaddrs = gethostbyaddr(name)
```

**The bind was not slow. It was waiting on the resolver** — and `serve_forever` is not reached until it
answers. Measured immediately, same machine, same address, same instant:

| interpreter | `socket.getfqdn("127.0.0.1")` | cold | warm |
|---|---|---|---|
| `/usr/bin/python3` 3.9.6 (Apple system — the fleet runner) | `'1.0.0.127.in-addr.arpa'` | **0.005 s** | 0.001 s |
| `python3` 3.14.6 (homebrew — the working interpreter) | `'localhost'` | **35.005 s** | 0.000 s |

Four orders of magnitude, from a call whose result nothing reads. And the two interpreters **do not even
agree on the answer** — one falls back to the `in-addr.arpa` form, the other resolves `localhost`.

**Why one runner sees it and the other does not, exactly.** The cost is paid **once per box** (the OS
resolver caches), by whichever server binds **first**. Under `unittest`, classes load in `dir(module)` order
— alphabetical — so `TestContentTabMainWiring` is the first binder and eats the whole 35 s inside its 2 s
window. Under `pytest`, collection is in **definition order**, so `TestMainServes` (line 614, ~2000 lines
earlier) binds first, absorbs the cost invisibly, and every later test finds the resolver warm. **Nothing
about the harness differed. The order of the first bind did.**

## Phase A′ — the production consequence, which is a false negative and not a delay

`up-injected.sh:2601-2604` polls the presenter cockpit's `/healthz` **25 times at 0.2 s**. Connection-refused
returns instantly, so the real window is ≈ **5 s** — seven times shorter than the bind. What a 35 s bind
produces is therefore not a slow cockpit:

> `⚠ presenter cockpit FAILED to come up on :7700 — /healthz never answered (non-fatal, but there is NO
> working cockpit).`

And the `tailscale serve` front for the cockpit sits **inside** the `cockpit_ok = 1` branch — so on a
`--public-host` demo (default-on since v2.3 M220) the one page a presenter actually opens is left on plain
HTTP while the bring-up declares it dead, and it then begins serving normally half a minute later.

## Phase B — the census, with its denominator stated

`TOK-08`'s report shape: enumerated population, how many were already false, and the fence's reach with its
denominator named.

| measured across **all of `rosetta-extensions`** (11 sections) | count |
|---|---|
| `.py` files mentioning `HTTPServer` | **4** — all in `demo-stack/` |
| reachable `HTTPServer`-derived classes | **3** |
| construction sites | **13** — 1 production (`cockpit.py:1768`, post-repair) + 12 in tests |
| sites that were paying the lookup **before** this iter | **13 of 13** |
| `server_bind` overrides anywhere in the monorepo | **0** |

**The tests were part of the population, not observers of it.** Ten sites in `test_cockpit.py` and two in
`test_roster_invariant.py` imported `ThreadingHTTPServer` **from `http.server`** — binding a server the
cockpit does not run. A suite that "passed" was exercising a different bind path than production.

**Excluded by property, not by exemption:** the two `socketserver.TCPServer` sites in `test_ant_academy.py`.
`TCPServer.server_bind` calls `getsockname()` and stops; only `HTTPServer` adds the lookup. Clerkenstein's
Go listeners are out for the same kind of reason — `net/http` resolves nothing at listen.

## Phase C — the repair, applied to the population

`cockpit.ThreadingHTTPServer` is CPython's own `server_bind` with **exactly one substitution**,
`socket.getfqdn(host)` → `host`, so `server_name` and `server_port` are still populated and every downstream
reader still holds (asserted, per iter-158: *a proposed repair is a hypothesis*). Both test modules now take
the class from `cockpit`, so all 13 sites bind through it.

**The fence RED-ed on its author's first attempt, and it was right.** The initial repair kept the stock class
as a module-level alias purely as a base; the population check flagged `cockpit.py::_StdThreadingHTTPServer`
— a name in the tree bound to a server that still resolves. Carving it out by its leading underscore was
available and **rejected** (`§5` rules 70/71: a fence pinned to a *spelling* is not pinned to a *property*).
The alias was deleted instead. **The fence keeps zero exemptions.**

## Phase D — the fence, both controls shown firing

`demo-stack/tests/test_bind_no_reverse_dns.py`, 7 tests, plain `unittest` — **no pytest import, no fixtures,
no `parametrize`**, so it runs on *both* interpreters (rule 75: a fence only one runner can run is a fence
with an unstated scope).

The predicate is **"was the resolver consulted?"**, implemented by making `socket.getfqdn` **raise** — not by
timing the bind. Three independent reasons: the OS cache makes any warm timing read zero (which is how this
survived to iter-171); `§5` rule 51's timing leg is unusable on this host; and a slow stub would cost the
suite the very seconds the iter reclaims.

| control | fires on | shown |
|---|---|---|
| mutation (property) | CPython's **own** `ThreadingHTTPServer` under the same poison must raise | ✅ |
| mutation (population) | a planted stock class in a scanned module must be flagged | ✅ |
| anti-vacuity (prefilter) | <3 candidate files ⇒ RED, not green | ✅ asserted |
| anti-vacuity (reach) | <3 reachable server classes ⇒ RED | ✅ asserted |
| fail-closed | an **unimportable** candidate is an unknown, never a pass | ✅ asserted |
| contract-preservation | `server_name` / `server_port` still populated | ✅ |
| hostile-resolver ceiling | bind completes < 5 s against a stub that hangs 30 s | ✅ |

The scan walks the **whole monorepo**, not `demo-stack/` — so the next HTTP server, in whatever section,
cannot land unfixed.

## Phase E — re-measure, both runners named

`stack-core/suite_census.py --only demo-stack --runner both`, **35 modules**:

| runner | GREEN | ENV-GATED | RED | TIMEOUT | tests |
|---|---|---|---|---|---|
| unittest (3.14.6) | **31** | 4 | **0** | 0 | 1073 |
| pytest (3.9.6) | **31** | 4 | **0** | 0 | 1062 |

**Runner disagreement across `demo-stack`: 1 module → 0. ACTIONABLE REDs: 0.** The 4 ENV-GATED are the
declared set (`test_demopatch`, `test_migrate_race_live`, `test_ssr_origin_chain`, `test_ant_academy`) —
unchanged, and six of their tests remain `FIX-M257x-iter145-sha-baseline-drift`, still a freshness signal.

`test_cockpit` + `test_roster_invariant` directly: **239 tests OK under both interpreters** (was: 2 failures
under 3.14). The new fence: **7 OK under both**.

**What this did NOT cover** (`§5` rule 60 — a scoped green is evidence about its scope alone): the other
four Python sections were not re-run. The change is confined to `demo-stack/` and the census proves no other
section holds a member of this class, but that is an argument about *reach*, not a measurement of *those
suites*. The 11-test count difference between runners (1073 vs 1062) is pre-existing and unexplained here.

## Close — 2026-08-08

**Outcome:** the last unexplained runner disagreement is **closed as a shipped defect, not an artifact** —
CPython's `HTTPServer.server_bind` waits on a **reverse-DNS lookup** (measured **35.005 s** cold on 3.14.6
vs **0.005 s** on 3.9.6, same address) that nothing reads, on the critical path of the presenter cockpit's
bind, against a `/healthz` gate ≈ **5 s** wide — so it presented as *"there is NO working cockpit"* and
skipped the cockpit's `tailscale serve` front. Censused across **all of `rosetta-extensions`** (4 files, 3
classes, **13 of 13** construction sites affected), repaired at the population, and fenced by a poisoned
resolver with **both** controls shown firing. Runner disagreement across `demo-stack`: **1 → 0**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (third consecutive `closed-fixed`; no no-prog
streak, and **no `P`/`N` reading was taken, so the metric is UNMEASURED, not unmoved** — `§9`) —
(3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (1 tik so far this run) — (6) protocol-stop: n —
(7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-171-1` … `D-M257x-171-4` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none — the `latency-budget.md` signature row and the `§5` rule 76 entry are the
iter's own lesson landed where it cannot rot (the protocol-evolution rule), not unrelated fixes.

**Routes carried forward:**
- `SURVEY-M257x-iter170-cockpit-runner-dependence` — **CLOSED by root cause.** Not a harness assumption; a
  blocking resolver call in shipped code. Every disagreement iter-170 found is now attributed.
- `SURVEY-M257x-iter171-runner-test-count-gap` — **NEW.** The same 35 modules yield **1073** tests under
  unittest and **1062** under pytest. Pre-existing, unexplained, and by rule 75's own logic an unstated
  scope: 11 tests execute under one runner and not the other, and nobody has named which.
- `SURVEY-M257x-iter171-anchor-guard-detects-structure-not-staleness` — **NEW, and measured on live
  examples.** The 43-line insert moved 6 corpus pins into `cockpit.py`; the pre-commit anchor guard caught
  2. But **2 of the 6 were already wrong before this iter** (`latency-budget.md:44`'s `:1214` / `:882`, off
  by ~160 and ~150 lines), and they had survived every prior run because they landed on a docstring tail and
  a live JS statement rather than on a blank line or a closing delimiter. The guard fired only because the
  new offset happened to push one onto a `}`. **A pin that rots onto plausible-looking code is invisible to
  it.** Concretises `SURVEY-M257x-iter163-anchors-with-no-quoted-literal`. All 6 repaired (`D-M257x-171-4`).
- `FIX-M257x-iter170-two-modules-cannot-run-on-the-modern-interpreter` — unchanged; still open.
- `FIX-M257x-iter145-sha-baseline-drift` — unchanged; still the freshness signal, still 6 tests.
- The standing queue, unchanged.

**Lessons:** **an unexplained runner disagreement is a shipped defect until someone proves it is an
artifact.** iter-170 earned *name the runner*; this iter is the corollary that gives it teeth. "That runner
is weird" is the reading that guarantees nothing is learned — and here the "harness assumption" turned out
to be CPython's own bind path, running identically in production, on the demo's entry point, behind a health
gate seven times too narrow to see it.

Two method notes worth more than the fix. **Dump the frame before theorising:** one
`sys._current_frames()` call named the culprit where two plausible wrong answers were already on the table.
And **the cache is the reason this lived so long** — the second measurement is free, so every warm reading
reports the bug as absent; a fence for any resolver/lookup class must *defeat* the cache (poison the call)
rather than *time* it. Booked as `§5` rule 76 and as the fourth arithmetic signature in
`corpus/ops/demo/latency-budget.md`.
