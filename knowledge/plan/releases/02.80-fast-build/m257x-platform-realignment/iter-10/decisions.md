# iter-10 decisions

## D-M257x-14 — the loopback literal is a PROGRAM-SEMANTICS choice, not only an exposure one

**Context.** M221 tightened `next dev`'s localhost bind `0.0.0.0` → `127.0.0.1`. Reviewed as an exposure
change, fenced as an exposure change, correct as an exposure change — and it silently changed what the app
does, because `next@16` compares its middleware-rewrite origin against the router's base origin **by string
equality** while normalizing only one of the two sides' loopback hostnames.

**Decision.** Bind `-H localhost`, and record *why* the literal matters in the file itself.

**Alternatives rejected, each on measurement rather than argument.**

| candidate | why not |
|---|---|
| revert to `-H 0.0.0.0` | Re-opens the exposure M221 closed. The orchestrator's constraint forbade it and it was never attempted. |
| `-H ::1` (iter-09 candidate #1) | Measured: still **500 / 30.1 s**. Loopback, but not next's normalized form — and it additionally makes IPv4-loopback dials fail. |
| `NODE_OPTIONS=--dns-result-order=ipv4first` | Measured: **no change**. It repairs the mechanism iter-09 hypothesised, which is not the mechanism. |
| declare the origin via `experimental.trustHostHeader` | Would make `initUrl` `https://…` — wrong scheme, and it is a platform-repo config edit (forbidden). |
| patch `next` in `node_modules` | Not durable, not a rext-owned file, and the string fix makes it unnecessary. |

**Why this one is structural rather than lucky.** The equality next performs is on the string *we* pass, so a
host whose resolver orders `localhost`'s addresses differently cannot re-break it. Contrast every DNS-shaped
candidate, whose correctness would have been a property of the machine it was tested on. *Prefer a design
that cannot express the bug over a check that catches it.*

**Residual, disclosed.** `listen('localhost')` picks one loopback address, and which one is OS-dependent
(measured `[::1]` on darwin/arm64 **and** linux/arm64, node 26). Both are un-routable, so exposure is
unaffected; but a consumer that dialled the IPv4 literal `127.0.0.1:<academy port>` would fail. None does
today — every rext consumer dials `$HOST` (default `localhost`). The bind comment states the requirement for
future ones.

---

## D-M257x-15 — a check reports the most informative observation, not the last one

**Context.** The new autoverify academy probe retried three times and reported `$ahome_code` from the final
attempt. Run **live** against the broken bind, attempt 1 returned `500` after 30.1 s and the retries piled up
behind the server's own in-flight self-proxies, blew the 35 s cap, and returned `000`. The check then
announced *"does not answer at all"* about a server that had demonstrably answered — the exact conflation the
check was written to end, one level down, committed by the check.

**Decision.** Rank the observations: a `2xx/3xx` wins and stops the loop; **any real status code outranks a
later timeout**; `000` is reported only when nothing else was ever seen.

**The general form** (now `platform-alignment.md` §5 rule 11): *a real status code observed once is a fact;
a timeout is only the absence of one.* A retry loop that keeps the last value throws away its own evidence,
and does so most often exactly when the system is unhealthy — i.e. when the evidence matters.

**How it was caught, which is the part worth keeping.** Three unit tests passed against the defect, because
each stubbed a *stable* response. It took one live negative control — the new check pointed at the old bind —
to see it. **A check must be run against the broken state, live**, not only against a fixture of it.

---

## D-M257x-16 — the exposure pin was WIDENED on purpose

`test_the_real_shipped_ant_academy_binds_loopback` asserted `== "127.0.0.1"`; it now asserts membership in
`_LOOPBACK_BINDS`. That guard owns exactly one proposition — *is the academy exposed?* — and every member of
that set answers it identically. **Which** loopback literal is required is a different proposition and is
fenced where it belongs (`test_ant_academy.py::test_the_loopback_literal_is_the_one_next_normalizes_to_ITSELF`),
with next's own regex quoted verbatim so the fence states the upstream contract it depends on.

Had the exact literal stayed pinned in both places, the exposure fence would have gone RED at a change that
does not touch exposure. An exposure alarm with no exposure in it is how a fence gets read as noise.
