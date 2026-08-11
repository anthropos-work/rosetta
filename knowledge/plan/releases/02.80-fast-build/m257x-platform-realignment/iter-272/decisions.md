# iter-272 — decisions

## D-M257x-272-1 — the succession failure is a seed/expectation mismatch, not a defect; and the fix is ours

**Context.** `pt-workforce-succession` is the one failing Playthrough standing between the suite and gate
clause 2. iter-261 offered two causes; iter-267 refuted both by SQL and relocated the fault above the data
layer, leaving a residual that still included *"the projection arithmetic"* — which, being platform code,
would have collided with v2.8's **0 platform edits** constraint.

**Measured (twice, deterministic, reset-to-seed each time).** The wire carries: 12 roles, `successors: []`
on **12 of 12**, `readyCount`/`developableCount` **0/0** on 12 of 12, `topTalents` empty, and `atRisk` with
**27** entries — 27 of the org's 28 members. Pat Ellis is the **single** exclusion, and appears in the
payload only as an `incumbents` entry of `DevOps Engineer` at **`fit: 87`**.

**Decision — no defect is claimed, because the code entails the observation.**
`succession.go:741-753` skips incumbents, skips zero-skill members, and requires `fit ≥ 40`;
`readinessBucket` calls `≥ 70` *ready*. Given that, `successors: []` everywhere **entails** that every
non-incumbent is below the fit-40 floor for every role. So:

- she is **not at-risk** because she is the only member whose skills match her role;
- she is **not a successor** because incumbents are excluded from their own role — by design;
- **no one else** qualifies anywhere, because everyone else is below the floor.

The assertion is **unsatisfiable on this seed**. The projection is behaving correctly and the renderer is
faithful (`0 ready` is what it was given).

**Consequence that changes the next iter's shape:** the fix surface is **ours** — the pt-world seed, or the
Playthrough's expectation — and no platform edit is needed or permitted.

**Recommendation, recorded so the next iter inherits the reasoning and not just the choice:** prefer the
**seed**. Changing the expectation to assert her as an incumbent turns the Playthrough green tomorrow, but
the `successors` / `topTalents` / `readyCount` half of this surface is now **proven empty on every
pt-world reset** — greening it without populating it would certify a computation that has never produced a
non-empty result. That is the green-but-wrong shape this milestone exists to catch.

## D-M257x-272-2 — a server-rendered surface is invisible to a trace's network log

**Context.** The route named the measurement as *"the `GetSuccession` HTTP response, captured from a
logged-in manager session"*, and PR-3 predicted the retained Playwright trace would contain it.

**Measured.** The trace holds **5** client-side GraphQL calls — `billing`, `userPreferences`,
`organizationSettings`, `setUserPreferences`, `userMemberships`. **The succession query is not among
them**, because next-web fetches the projection **server-side**; the response never crosses the browser
boundary. **PR-3 is refuted as stated.**

**Decision.** For a server-rendered surface, capture the **RSC payload resource** in the trace, not the
network log. The payload carried the entire projection — `roles`, `successors`, `atRisk`, `topTalents`,
`allMembers` — which is strictly more than the single response the prediction asked for.

**The distinction is kept deliberately.** PR-3's *conclusion* ("no new instrument is needed") held while
its *mechanism* was wrong. Recording that as a plain "held" would have taught the next iter to look for a
GraphQL call that does not exist. The escalation condition in this iter's `overview.md` — *do not
reimplement the login flow to get a response* — is what kept the refutation cheap: the answer was already
in the artifact, one directory over.
