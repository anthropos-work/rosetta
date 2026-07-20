# iter-02 — Decisions

## INCOMPLETE-EXIT-2026-07-20

**What got done:** step 1 of 5 — `billion`'s Docker build cache pruned, 109 GB reclaimed, free space
40 G → 139 G (80 % → 29 % used). The host capacity precondition for Phase C's cold UI-tier rebuild is
satisfied. No repo state changed; no platform repo touched.

**What's left:** steps 2–5 — publish `rosetta-extensions` `main` + the 13 `playbill-*` tags to origin,
re-pin `billion:/home/devops/panorama/.agentspace/rext.tag` → `playbill-m235-hardened`, check the
consumption clone out at that tag, verify the M217 FATAL pin guard.

**What blocked progress:** the Phase 0b `audit-kb-fidelity` verdict returned **RED (blocker)** mid-iter.
Per Phase 0b, RED blocks; per the skill's critical-decision list, a sub-agent audit-RED stops the loop and
wakes the user.

**Why I did not just push anyway.** The publish is safe in isolation and I had already verified it
(ancestors, fast-forward, zero collisions, green tests). The reason to stop is not the push — it is
**what the audit says about the destination.** Two of the five blocker classes attack M236's *exit gate*
itself:

- *"reachable only over the tailnet"* is **false by construction** — every demo container publishes on
  `0.0.0.0` on every bring-up (`safety.md:405`, measured at `:413`), and on Linux `tailscale serve`
  bypasses the host firewall so `ufw deny` does not help (`tailscale-serve.md:626-627`). The clause can
  only ever be true as a property of `billion`'s **network placement**, demonstrated by an explicit
  **off-tailnet probe** — which is a different deliverable from the one the gate implies.
- *"p95 click→ACCESS < 5 s"* is **unmeasurable today for every one of the 31 actions** — the content-player
  CTA emits no `data-login-as` (`cockpit.py:421-425`), and that attribute *is* the ACCESS predicate that
  `latency.ts:123-127` throws without, before t0; `run-latency.sh:42-47` additionally hard-rejects
  non-hero vantages with `exit 2`.

Publishing and proceeding would have committed 3–4 iters of Phase-L work toward a gate that cannot be
satisfied as written, and the eventual failure would have been attributed to seeding or render defects
rather than to the gate's own definition. **Stopping here costs one round-trip; not stopping costs the
milestone's credibility about what it proved.**

**Handling per Phase 4 Step 0:** no `## Close` section written, no `iter(M236/02):` commit, iter graded
with no close status. The iter's artifacts are committed under a `probe(M236/02):` intermediate prefix so
the working tree stays clean for the orchestrator while downstream tooling still reads iter-02 as
in-flight.

**Resumption:** steps 2–5 are unchanged and ready. Whoever resumes should re-run only the iter-01 B2
freshness check (`git ls-remote --tags origin | grep playbill` → expect 0) before pushing.
