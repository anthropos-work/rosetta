**Type:** tik

# iter-10 — the academy's 30 s 500, and the check that read the wrong route

## What the re-survey overturned

iter-09 closed `no-lift` on purpose and handed off a **mechanism** plus a **fix menu**. The mechanism was
wrong, and it was wrong in a way that would have cost the fix: it said `localhost` resolves to `::1` before
`127.0.0.1`, so an IPv4-only listener cannot be self-dialled. Every candidate in the menu followed from that
story. Three measurements killed it before a line of code was written:

| probe | result | what it rules out |
|---|---|---|
| `curl http://127.0.0.1:13077/` (never utters `localhost`) | **500 / 30.061 s** | any client-side resolution story |
| `-H 127.0.0.1` + `NODE_OPTIONS=--dns-result-order=ipv4first` | **500 / 30.123 s** | the direct repair for a resolution-order bug |
| `-H ::1` (makes the first-resolved address the bound one) | **500 / 30.1 s** | iter-09's own candidate #1 |
| `-H 0.0.0.0` (control) | 200 / 0.686 s | — |

**The real mechanism is a string comparison**, and the evidence is a response header:

```
-H 127.0.0.1 :  HTTP/1.1 500 …  x-middleware-rewrite: http://localhost:13077/     ← absolute ⇒ proxied
-H localhost :  HTTP/1.1 200 …  x-middleware-rewrite: /                            ← relativized ⇒ served
```

Three files of `next@16` disagree about what a loopback host is called:

- `server/web/next-url.js:15-20` — `REGEX_LOCALHOST_HOSTNAME` normalizes **every** loopback hostname
  (`127.x.x.x`, `[::1]`, `localhost`) to the literal `"localhost"`. The middleware's `request.nextUrl` origin
  is therefore *always* `http://localhost:$PORT`, regardless of bind or `Host` header — which is exactly why
  dialling `127.0.0.1` directly did not change the outcome.
- `server/lib/router-utils/resolve-routes.js:117` — the router's `initUrl` is built from the **raw**
  `opts.hostname`, i.e. the literal string handed to `-H`. No normalization.
- `shared/lib/router/utils/relativize-url.js` — `getRelativeURL` keeps a rewrite **absolute** unless
  `relative.origin === baseURL.origin`. Plain string equality.

Mismatched, the app's own in-place rewrite is classified as an **external** proxy target
(`router-server.js:377` → `proxy-request.js`) and the dev server proxies **to itself** until
`http-proxy`'s `proxyTimeout` default of `30_000` ms. **The flat 30.0 s is that constant** — the launcher's
comment had attributed exactly this shape to Turbopack cold-compile and budgeted 120 s for it, a wait that
could never succeed.

## The fix, and why it does not trade the security away

`-H 127.0.0.1` → **`-H localhost`**: the one loopback literal that is its own next-normalized form.

The orchestrator's constraint was binding — *do not re-open the exposure to make a gate check green*. This
fix does not: `localhost` resolves **only** to a loopback address on every host family, so the M221
de-exposure is unchanged. Proven three ways rather than argued:

- `exposure_claim_guard.py` run live → `academy:localhost` read as loopback, `OK — the docs state the
  exposure truthfully, for BOTH stack families`. (`_LOOPBACK_BINDS` already contained `localhost`.)
- `curl http://192.168.1.126:13077/` — this box's **routable** address — `code=000`, not reachable.
- `lsof` on the running server: `TCP [::1]:13077 (LISTEN)`. One address, loopback.

And it is **resolver-independent by construction**: the comparison next performs is on the string *we*
supply, so no host's resolver ordering can re-break it. That is the property the DNS-shaped fix menu could
never have had.

### Two-OS honesty

The origin-equality half is a string operation in the same JS bundle on any OS; it is argued from cited
`file:line`, not from a single-host measurement. The genuinely OS-dependent half — *which* loopback address
`listen(port,'localhost')` picks — was measured on **both** families:

| host | `localhost` resolution order | `listen('localhost')` binds |
|---|---|---|
| darwin/arm64, node 26.5.0 (this box) | `::1`, then `127.0.0.1` (verbatim) | `[::1]` |
| **linux/arm64**, node 26 (debian slim, container) | `::1`, then `127.0.0.1` (verbatim) | `[::1]` |

**Scope of the claim, stated plainly:** the full academy render was validated only on darwin/arm64. On Linux
only the bind-address half was measured, in a Debian-slim container — not on `billion`, which this box cannot
reach (`HOST-M257x-toolchain`: no tailscale). Both loopbacks are equally un-routable, and every rext consumer
of this port dials `$HOST` (default `localhost`) — `up-injected.sh:604`, `cockpit.py --academy-base`,
autoverify's academy check — so nothing dials the IPv4 literal. A future consumer that does would break on
both families measured here, and the bind comment says so.

## The finding that outgrew the fix

**`autoverify` never detected this at all**, and iter-09's parting claim that the academy was "the only
genuine ✗ blocking clause 1" is wrong in both directions. Baseline autoverify against the broken academy:
**2 FAILED, neither of them the academy** (both are the routed-forward evidence-log-path bug). It printed:

```
✓ AI Academy renders its catalog on :13077/library (real course cards)
```

…over a demo whose landing page was a 500. Check (f) probed **only** `/library/`, which answered **200 in
9 ms** while `/` answered **500 in 30.0 s** — `/library` is public and short-circuits in Clerk's middleware
*before* the rewrite that loops; `/` does not. **The one route the check read was the one route the defect
spared.** A presenter clicking "AI Academy" got the 500; the verifier that measures the gate said ✓.

That is the same class as everything else this milestone has found, and it is the reason M221's tightening
survived four releases: it shipped its exposure fence — which correctly stayed **green**, because the bind
*was* loopback — and no check at all for whether the thing still worked.

## What landed

1. **`demo-stack/ant-academy.sh`** — `bind_args=(-H localhost)`, with the mechanism written into the file
   (the three next `file:line`s, the four measurements, the exposure argument, the OS-dependent residual).
2. **`ant-academy.sh`'s readiness probe** — was `curl -fsS --max-time 3`: `-f` collapses a 5xx into silence,
   and the per-attempt window is **shorter than the 30 s failure it watches**, so it printed *"alive but
   NEVER ANSWERED"* over a server that was answering. It now captures the HTTP status **and** curl's exit
   code and names the state it measured — *absent* / *hung* / *answering wrong*, three different repairs.
3. **`autoverify` check (f)** — now probes `/` as well as `/library/`, reports the three states distinctly,
   and names the repair in the warning rather than only the symptom.

## The defect this iter committed, and how it was caught

The live negative control — running the **new** check against the **old** bind — reported *"does not answer
at all"* about a server that had demonstrably answered 500. Cause: the first attempt got the 500 after 30.1 s,
the retries piled up behind the server's own in-flight self-proxies and blew the 35 s per-attempt cap
(`curl` exit 28 → `000`), and the check read the **last** attempt. **The exact conflation the check exists to
end, one level down, committed by the check itself** — and it was only visible because the control was run
live rather than reasoned about. Fixed by ranking: a real code observed **once** outranks any later timeout;
`000` is reported only when nothing else was ever seen. Re-proved live afterwards.

## Live proof (demo-1, this box)

| | before | after |
|---|---|---|
| `GET /` | **500 in 30.077 s** | **200 in 0.205 s** (0.060 s warm) |
| launcher verdict | `✗ … NEVER ANSWERED on :13077 within 120s` | `started + SERVING on :13077` |
| body carries `AI Academy` | no (error page) | yes |
| `/library/` course cards | 37 | 37 (unchanged — it was never broken) |
| autoverify academy check | `✓ renders its catalog` (**false green**) | `✓ renders its catalog` (**true**) |
| autoverify vs the broken bind | *(check did not exist)* | `⚠ answers HTTP 500 on :13077/ — the LANDING PAGE is broken … must be launched with -H localhost` |
| autoverify FAILED | 2 | 2 (both `FIX-M257x-autoverify-evidence-log-path`, routed to iter-11) |
| exposure on `192.168.1.126:13077` | not reachable | not reachable |

## Tests

7 fences added/updated; **5 mutations RED-proven**, each mutant `bash -n`-clean and *collected before it was
run* (iter-07's §8 rule 5):

| mutation | fence that went RED |
|---|---|
| `-H localhost` → `-H 127.0.0.1` | `test_bind_defaults_to_loopback…` + `test_the_loopback_literal_is_the_one_next_normalizes_to_ITSELF` |
| `-H localhost` → `-H 0.0.0.0` | `HostNativeBindTest::test_the_real_shipped_ant_academy_binds_loopback` |
| per-attempt window `10` → `3` | `test_readiness_probe_can_observe_the_failure_it_reports` |
| autoverify home probe `/` → `/library/` | all 3 `TestAutoVerifyAcademyLandingPage` cases |
| *(fixture)* 500-then-000 attempt sequence | `test_a_500_seen_ONCE_outranks_later_timeouts` |

The subtlest one is deliberate: the exposure guard's pin was **widened** from `== "127.0.0.1"` to the
loopback **set**, and the new fence in `test_ant_academy.py` owns *which* loopback literal is required.
Pinning the exact literal in both places would have made the exposure fence go RED at a change that does not
touch exposure — an exposure alarm with no exposure in it, which is how a fence gets ignored.

`demo-stack/tests/test_ant_academy.py` + `stack-injection/tests/test_exposure_claim_guard.py`:
**112 passed, 1 failed** — the failure is the pre-existing `test_apply_revert_round_trip_on_the_real_next_config`
already logged under `CHECK-M257x-live-clone-suites-red`, reproduced on a pristine control clone in iter-06.
`shellcheck -S warning` clean on both edited scripts.

## Close — 2026-07-31

**Outcome:** the academy's four-release-old 30 s 500 is fixed at its real (origin-string) mechanism with the
M221 de-exposure intact and re-proved three ways; and the two checks that should have caught it — the
launcher's readiness probe and autoverify's academy check — were both measuring something other than what
they reported, and now measure it.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D-M257x-14 (the loopback literal is a *program-semantics* choice, not only an exposure one) ·
D-M257x-15 (a check ranks the most informative observation, not the last)
**Side-deliverables:** none — both landed lines were planned scope.
**Routes carried forward:**
- `FIX-M257x-autoverify-evidence-log-path` → **iter-11**. Unchanged and now the *only* autoverify ✗; fully
  specified in iter-09's pre-compute. It is what stands between here and a clean `green:true / 0 warnings`.
- `CHECK-M257x-live-clone-suites-red` → unchanged (1 of its 7 reproduced here).
- `HOST-M257x-toolchain` → unchanged; it is why the Linux half of this iter is a container measurement and
  not a `billion` one.

**Lessons** (promoted to `platform-alignment.md` §5 as **rule 11**, same commit):

- **A probe must exercise the surface whose health it claims.** Related to rule 7 but distinct: this probe
  measured something real, it just measured the wrong thing and reported a conclusion about the thing it had
  not touched.
- **Every security tightening ships with a paired "does it still work?" check.** M221 shipped its exposure
  fence and nothing else. The exposure claim was true and the app was broken — for four releases.
- **A boolean probe hides the state it cannot express**, and a per-attempt timeout shorter than the failure
  it watches guarantees it. Capture the status and the exit code, and name the state you measured.
- And the one this iter learned about itself: **a check must be run against the broken state, live.** The
  ranking defect was invisible to three unit tests and obvious on the first live negative control.
