---
iter: 77
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-04
---

# iter-77 — adjudicate the union, then close the reach hole at its cause

**Active strategy:** `TOK-05` — *stop repairing claims; fence the predicates under them.*
`D-M257x-59-1` (repair by predicate), `D-M257x-59-5` (fence first, then citations).

## Step 0 — re-survey (mandatory)

Re-run at open, from the corpus root against `--platform stack-demo/platform`
(`0dab54d`; `stack-dev/platform` **does not exist on this machine** — the guard's refusal-to-default
is what surfaced that):

```
6 repo(s), 10 compose service(s), floor [postgresql redis sentinel], 8 legal profile(s),
default `core` selects 5, migrating ['app']
reach — G5 24 migration claim(s) = 1 enumerated + 21 free prose UNREACHED + 2 ref-pinned;
        G2 3 repo-count claim(s)                                            → OK
```

TOK-05's named target (`FIX-M257x-iter76-read-union`) is **live and unabsorbed**. No substitution.

## Cluster / target identified

`FIX-M257x-iter76-read-union` under its three binding conditions: **adjudicate before repairing** ·
**repair by predicate** · **not closed until the G5/G2 reach hole is closed**.

## Hypothesis

The briefing poses the design question directly: *can free prose be fenced at all, or should the
corpus restate those 21 claims in a form that can be checked?* The hypothesis this iter tests is a
**third answer**: that the reach hole is **not a prose problem**, and that both framings share a
false premise — that the 21 unreached claims are the failure. The prediction is that **the one claim
G5 does reach is read WRONG**, because the guard's repo vocabulary is derived from what *currently
exists*, so a name leaves the vocabulary at the same commit it leaves `repos.yml` — going invisible
exactly when it goes false.

## Expected lift

- The union's mechanical partition by ref-pin, sized (iter-76 found one such false positive by hand).
- A live proof or refutation of the vocabulary hypothesis.
- If confirmed: a **derived** fix (history of the platform artifact, no heuristic), watched RED→GREEN.
- An answer to the fence-vs-restate question **with a measurement**, not a preference.

## Phase plan

A adjudicate the union mechanically · B test the vocabulary hypothesis · C measure the candidate
prose fences and report the negative if they do not clear · D build what survives + positive controls
(§8 rule 5: a no-op control that SURVIVES and an INVERTED mutant) · E repair the adjudicated defects ·
F re-measure reach, honestly.

## Escalation conditions

- A prose fence that cannot reach FP-free precision on the live corpus is **reported as a measured
  negative and NOT shipped** (§4 Trap A: replace the rule, never the threshold).
- More than ~10 adjudicated repair sites → repair the predicate's live set, route the remainder.

## Acceptable close-no-lift outcomes

Refuting the vocabulary hypothesis and reporting the fence measurements is a complete iter.
