# iter-285 — the cockpit advertised a world the database no longer held

**Type:** tik — under `TOK-09`.

## Defect 5 — it is not an AI-readiness defect, and the data was never wrong

Reported: logging in as **Dana Whitlock (manager)** lands on `/ai-readiness` showing the *"start now"*
boilerplate instead of the seeded dashboard — a regression against M219, which seeds **both** a `closed`
and an `active` cycle.

The first query refuted the premise:

```
SELECT status, count(*) FROM public.ai_readiness_cycles GROUP BY status;   -- active|1  closed|1
```

**Both cycles are seeded.** The second query found what actually differs:

| what the cockpit advertises | what `demo-2`'s database contains |
|---|---|
| Northwind Aviation · Meridian Talent · Cervato Systems · hero **Dana Whitlock** | Meridian Labs · Halcyon Retail · **Vertex Logistics** · Kestrel Hiring Group · Cervato Systems |

Those are two different worlds. The database holds **`pt-world`** — the Playthroughs' decoupled test seed
— and its AI-readiness cycles belong to **Vertex Logistics**, whose manager seat is **Nadia Ferrante**.
**Dana Whitlock is not in the database at all.** So the seat logs a presenter in as a hero who does not
exist, `/ai-readiness` finds no cycle for whatever org that resolves to, and renders the empty state.
Which reads exactly like a product regression.

**The mechanism, measured, not inferred:**

* the cockpit process started **Mon Aug 10 22:30:57**, reading `--manifest …/demo-2/cockpit-manifest.json`;
* that file was **rewritten at 23:21:39Z** — 51 minutes later — by a `--reset` re-seed to `pt-world`;
* `cockpit.py` holds its manifest **in a closure**, so it kept serving the pre-reset world from memory.

## And the guard for exactly this EXISTS, is correct, is tested — and was never passed the flag

`cockpit.py`'s own source says so, in a comment block written at **v2.8 M256** for this precise failure:

> *"`--reset` re-exports `fake-fapi-roster.json` but never `cockpit-manifest.json` … `handleHandshake`
> in the fake FAPI IGNORES an unknown `__clerk_identity` and establishes a session anyway, so a button
> naming a hero the roster no longer holds produces a SUCCESSFUL-LOOKING WRONG LOGIN — no UI error, no
> log line … So the cockpit no longer trusts its manifest absolutely."*

The remedy is `cockpit.py --roster`, which cross-checks the advertised heroes against who actually
exists. **`up-injected.sh` never passes it.** It exports the roster and threads `roster_flag` into
`gen_injected_override.py` — the **fake-FAPI's** mount, a different consumer with a different flag — and
the cockpit launch at `:2603` carries `--manifest --fapi-host --app-base --port` and four optional args,
none of them the roster. Every demo cockpit has therefore run `ROSTER_CHECK_ABSENT`, the state the module
itself documents as *"cross-check legitimately disabled."*

`run-playthroughs.sh` asserts the opposite in two places — *"cockpit.py's `--roster` cross-check makes a
stale in-memory manifest fail closed"* and *"the `--roster` cross-check still fails closed"*. Both
sentences are true about the guard and false about the deployment.

> **A check that is implemented, tested and never invoked is indistinguishable from one that passes.**
> This is the milestone's *parsed-but-never-read is invisible* class, one level up: **built-but-never-armed.**

**Fixed:** the cockpit launch now carries its own `cockpit_roster_arg`, deliberately a **different array**
from `roster_flag` (one array cannot serve two consumers whose flag names can diverge independently), in
the bash-3.2-safe `+alt` form the script uses everywhere else. Fenced by four arms in
`demo-stack/tests/test_cockpit.py` — the launch passes a roster, the variable is assigned, it is not the
override generator's array, and the empty case is `set -u` safe. **RED-proven** by removing the argument.

## Defect 3 — Log out is REPLACED, not joined

*"when you patch it to add the back to the cockpit, remove the logout option."* On a demo, Log out is a
dead end: there is no password to sign back in with and the cockpit **is** the way back. The swap rides
the **same fail-closed expression** as the item itself — `backToCockpitMenuItem ? null : mapItem(logOutMenuItem, 0)`
— so with `NEXT_PUBLIC_COCKPIT_URL` undefined the dropdown is byte-identical to production. `post_sha256`
recomputed. Three arms fence it, including the anti-vacuity one (an unconditional Log out must not
survive *beside* the conditional one) and the fail-closed shape; **RED-proven**.

## Verification

| scope | result |
|---|---|
| `demo-stack/tests/test_cockpit.py` + `test_back_to_cockpit_m249.py` | **245 → 248 passed** |
| the four roster-wiring arms, with the argument removed | **2 failed** — RED-proven, restored |
| the three log-out-swap arms, with the swap reverted | **3 failed** — RED-proven, restored |
| `bash -n up-injected.sh` | clean |

**NOT COVERED, stated:** no stack was brought up, re-seeded, restarted or torn down; **the live `demo-2`
cockpit still serves the stale manifest**, because the fix is read at launch. The AI-readiness read path
itself was never suspected after the first query and was not exercised.

## Close — 2026-08-11

**Outcome:** defect 5 is **not an AI-readiness defect** — `demo-2`'s database holds the Playthrough test
world while its cockpit serves the demo world from memory, and the cross-check written for exactly this
failure was never passed to the cockpit. Wired and fenced. Defect 3 landed with its fail-closed shape
preserved.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (3 tiks) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**

**Why `closed-fixed-partial`:** defect 3's planned scope was the back-to-cockpit patch **family**, and
only the `next-web` member landed. The two siblings are routed with named reasons, not silently dropped.

**Routes carried forward:**
- **`ROUTE-M257x-285-logout-swap-for-studio-and-academy`** — `studio-desk-back-to-cockpit` sits in a
  **sha chain** (`studio-desk-logout-url.pre_sha256` **is** its `post_sha256`), so editing it re-pins two
  manifests; `ant-academy-back-to-cockpit`'s anchor **stops immediately above** the logout row, so
  removing it needs the anchor extended over a block this iter did not read. Both are cheap-ish and
  neither is in the app the user was in.
- **`ROUTE-M257x-285-demo-2-cockpit-serves-a-stale-world`** — the live stack. Restarting the cockpit
  would make the seats match the DB (i.e. show `pt-world`), which is **not** what the user expects to
  see; restoring the demo world needs a re-seed. **Both are the user's call, not this iter's.**

**Lessons:**
1. **Refute the premise before diagnosing the mechanism.** One `GROUP BY` showed the AI-readiness data
   was correct, which moved the search a whole layer up and saved the iter.
2. **Built-but-never-armed.** A guard can be written for the exact failure, tested, quoted in two other
   scripts as the reason a hazard is handled — and never invoked. Grep for the flag at the **call site**,
   not for the feature in the module.
3. **An exported artifact and its consumer are two different questions.** `roster_flag` existed, was
   correct, and went somewhere else entirely.
