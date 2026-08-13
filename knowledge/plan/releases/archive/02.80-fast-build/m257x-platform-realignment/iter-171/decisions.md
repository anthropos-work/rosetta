# iter-171 — decisions

## `D-M257x-171-1` — the runner disagreement is a SHIPPED DEFECT, not a harness artifact

**Decision:** dispose `SURVEY-M257x-iter170-cockpit-runner-dependence` as a **defect in production code**,
repair it, and fence the property — not as a tight-test-window to be relaxed.

**How it was settled, and the method is the transferable part.** The two hypotheses on the table were both
plausible and both wrong-shaped: *"the 2 s bind window is too tight"* and *"3.14 is flaky."* Neither is
falsifiable by argument. What settled it in **one call** was dumping the blocked thread's frame:

```
cockpit.py:1768  httpd = ThreadingHTTPServer((a.host, a.port), handler)
  socketserver.py:457  self.server_bind()
    http/server.py:150  self.server_name = socket.getfqdn(host)
      socket.py:817    hostname, aliases, ipaddrs = gethostbyaddr(name)
```

The bind was not slow; it was **waiting on the resolver**. Measured immediately after, same machine, same
address:

| interpreter | `socket.getfqdn("127.0.0.1")` | cold | warm |
|---|---|---|---|
| `/usr/bin/python3` 3.9.6 (Apple system — the fleet runner) | `'1.0.0.127.in-addr.arpa'` | **0.005 s** | 0.001 s |
| `python3` 3.14.6 (homebrew — the working interpreter) | `'localhost'` | **35.005 s** | 0.000 s |

**Two interpreters, one address, four orders of magnitude — and they do not even agree on the answer.** The
alphabetically-first binding test class in `test_cockpit.py` is `TestContentTabMainWiring`, so under
unittest it is the process's first binder and pays the whole 35 s inside a 2 s window; under pytest
(definition order) `TestMainServes` binds first, and every later test in the process finds the OS resolver
warm. **Nothing about the harness differed. The order of the first bind did.**

**Why it is a production defect and not a test curiosity.** `up-injected.sh:2601-2604` polls the cockpit's
`/healthz` **25 times at 0.2 s** — connection-refused returns instantly, so the real window is ≈ **5 s**. A
35 s bind therefore does not read as a slow cockpit; it reads as

> `⚠ presenter cockpit FAILED to come up on :7700 — /healthz never answered (non-fatal, but there is NO
> working cockpit).`

and, worse, the `tailscale serve` front for the cockpit sits **inside the `cockpit_ok = 1` branch** — so on a
`--public-host` demo the entry point every presenter opens is left on plain HTTP while the bring-up reports
it dead, and it then starts serving normally half a minute later. **A false negative, not a delay.**

## `D-M257x-171-2` — repair the POPULATION by property, and delete the alias the fence caught

**Decision:** ship `cockpit.ThreadingHTTPServer` — CPython's `server_bind` with exactly one substitution,
`socket.getfqdn(host)` → `host` — and route **every** construction site in `rosetta-extensions` through it,
including the twelve in tests.

**Two things here were not obvious, and both are `TOK-08`-shaped.**

1. **The tests were part of the population, not the observers of it.** Ten sites in `test_cockpit.py` and
   two in `test_roster_invariant.py` imported `ThreadingHTTPServer` **from `http.server`** — so they bound a
   server the cockpit does not run, and a test suite that "passed" was exercising a different bind path than
   production. Both modules now take the class from `cockpit`. Repairing only the two failing tests would
   have left eleven sites paying a cost the fence claims is gone.

2. **The fence RED-ed on its own author's first attempt, and it was right.** The initial repair kept the
   stock class as a module-level alias (`from http.server import ThreadingHTTPServer as
   _StdThreadingHTTPServer`) purely as a base. The population check flagged
   `cockpit.py::_StdThreadingHTTPServer` — a name in the section bound to a server that still resolves.
   Carving it out by its leading underscore was available and **rejected**: `§5` rules 70/71 say a fence
   pinned to a *spelling* is not pinned to a *property*. The alias was deleted instead
   (`class ThreadingHTTPServer(http.server.ThreadingHTTPServer)`), so the fence keeps **zero exemptions**.

**The repair preserves what it replaces** (iter-158: *a proposed repair is a hypothesis*). `server_name` and
`server_port` are still populated — asserted — so every downstream reader of them still holds. What is lost
is the fully-qualified *form* of `server_name`, which is read by the CGI environment builder alone and this
process builds no CGI environment.

## `D-M257x-171-3` — the fence poisons the resolver rather than timing the bind

**Decision:** the predicate is **"was the resolver consulted?"**, implemented by making `socket.getfqdn`
raise, and the mutation control is **CPython's own class**, which must still raise under the same poison.

**Why not a timing assertion.** Three reasons, and each one alone is sufficient:

- **The cache defeats it.** The second lookup is ~0 s on both interpreters. Any timing test written after
  the first run in a process measures nothing — the bug is *invisible warm*, which is exactly how it
  survived to iter-171.
- **`§5` rule 51's timing leg fails on this host.** Wall-time is not a usable measurement here; counts and
  predicates are.
- **A timing test would cost the suite the seconds this iter exists to reclaim.** The poison is free.

**Both controls fire, and were shown to fire, not asserted to.** The stock-class mutation control raises
(so the property test is not vacuous); the population predicate flags a planted `http.server` class (so the
census is not vacuous); the candidate prefilter asserts a **non-empty** population, so a prefilter that
silently stops matching turns the fence RED rather than green — `§9`'s *a census returning ZERO must prove
its instrument*.

**The census, with its denominator stated** (`TOK-08` mandates this, and iter-114's rule): across **all of
`rosetta-extensions`** — not just `demo-stack/` — exactly **4** `.py` files mention `HTTPServer`, carrying
**3** reachable `HTTPServer`-derived classes and **13** construction sites (1 production + 12 in tests). All
13 now bind through the one fixed class. The two `socketserver.TCPServer` sites in `test_ant_academy.py` are
**out of the population by property, not by exemption**: `TCPServer.server_bind` calls `getsockname()` and
stops — only `HTTPServer` adds the lookup. The scan walks the whole monorepo precisely so that the next HTTP
server, in whatever section, cannot land unfixed, and an **unimportable** candidate is reported as an
unknown rather than counted as clean.

## `D-M257x-171-4` — a 43-line insert moved 6 corpus pins, and 2 of them were ALREADY WRONG

**Decision:** repair all six `cockpit.py:NN` pins the corpus carries, and record that **two of them were
stale before this iter touched the file** — do not let the repair commit absorb a pre-existing defect
silently.

**How it surfaced.** The repair inserted 44 lines and deleted 1 near the top of `cockpit.py`, shifting every
construct below by **+43**. The pre-commit fence (`anchor_construct_guard`) refused the commit with two
sites, which is the mechanism working. But the arithmetic did not check out, and that is the finding:

| corpus site | pin | pre-edit line held | verdict |
|---|---|---|---|
| `frontend-tier.md:499`, `ant-academy.md:394` | `cockpit.py:812` | the client-side `document.cookie = 'e2e_persona='` | **correct** → +43 → `:855` |
| `frontend-tier.md:499`, `ant-academy.md:395` | `:1496` | the `/go` `Set-Cookie` header | **correct** → +43 → `:1539` |
| `ant-academy.md:395` | `:327` | the comment naming both paths | **correct** → +43 → `:370` |
| `latency-budget.md:44` | `cockpit.py:1214` | a docstring tail — *"never emits inert academy JS"* | **ALREADY WRONG** |
| `latency-budget.md:44` | `:882` | `dyn.style.display = 'none';` — JS, not a CTA | **ALREADY WRONG** |

The two `latency-budget.md` pins claim to name *"`data-login-as` is emitted on hero cards only"* and *"the
Content-stories tab's seat CTAs carry a bare `href`"*. Those constructs are at **`:1374`** and **`:1036`**
today; the pins were off by ~160 and ~150 lines respectively **before iter-171 existed**. Repaired to the
constructs, verified by reading them.

**The transferable half — the guard detects a STRUCTURAL landing, not a WRONG one.** Both bad pins survived
every prior run because neither landed on a blank line or a closing delimiter; one landed on a docstring
tail, the other on a live JS statement. **What made the fence fire was my shift pushing `:882` onto a `}`**
— i.e. it fired on a *coincidence of the new offset*, not on the staleness that had been there all along. A
pin that rots onto plausible-looking code is invisible to it. That is
`SURVEY-M257x-iter163-anchors-with-no-quoted-literal` measured on live examples rather than argued, and it
is why `§5` rule 73's *grade at the grain of the claim* applies to anchors too: the claim is *"this line
emits `data-login-as`"*, so the check is *"does the pinned line contain that literal"* — not *"is the pinned
line syntactically unremarkable."*
