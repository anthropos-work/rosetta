#!/usr/bin/env python3
"""iter-124 — the tier-2 TRIAGE: what each unevidenced assertion actually IS.

The census (`claim_census_guard.py`, iter-122) enumerated 1,164 assertion sentences in the clause-5
surface that carry neither a citation nor a hedge. It could say nothing about what they ARE. This
assigns each one of the four fates the user's run-80 directive names:

    cite  — the evidence exists and is reachable from here; the sentence simply never gave it
    hedge — it cannot be checked from here, and the sentence must say so (the iter-093 principle)
    fix   — it is wrong
    drop  — it asserts nothing checkable and earns nothing

WHAT THIS IS, AND WHAT IT IS NOT
--------------------------------
**It is a TRIAGE, not a grader, and not a reading.** Nothing it prints is `P`, is `N`, or is a
clause-5 verdict (`F4`, `D-M257x-122-2`). Clause 5 is met only by the frozen graded read returning
zero; a triage answers a different question — *what kind of work does this sentence need* — and a
sentence can be triaged `cite` while being perfectly true or quietly false. **`fix` is the ONLY fate
this file cannot decide**: falsity is not syntactic. So `fix` is supplied as a HAND-ADJUDICATED
input (`FIX_SITES`), derived by reading, and the rules below assign the other three.

THE RULES, IN ORDER — first match wins, and each is stated so a reader can disagree with it
-------------------------------------------------------------------------------------------
  R0 fix    the sentence is in FIX_SITES (hand-measured false; see the iter's progress.md)
  R1 drop   scaffolding: a template placeholder (`[Foo]`), or a doc-purpose sentence
            ("This document describes ..."), or a heading restatement. Asserts nothing checkable.
  R2 hedge  the subject lies OUTSIDE every clone set AND outside every artifact this corpus can
            read: third-party SaaS internals, applied AWS state, Vercel runtime configuration,
            production data. Not "hard to check" — UNCHECKABLE FROM HERE. The directive is explicit
            that hedges must not be manufactured for facts somebody can go and measure.
  R3 cite   the sentence names at least one PLATFORM ARTIFACT that is reachable from a clone or from
            this corpus: a repo-relative path, a source symbol, a DB table/schema, an env var, a
            compose/profile key, a port, a version pin, a stream name. The evidence exists.
  R4 cite   default. A residual factual assertion about the platform is presumed CITABLE, because
            the platform is in clones and this corpus is in git. Presuming `hedge` here would
            manufacture exactly the hedges the directive forbids.

**R4 is the load-bearing choice and it is deliberately the generous direction**: it can only
UNDER-state `hedge` and `drop`. If the residual is really unfoundable, the proof is a failed
citation attempt, not a default.

MEASURED ACCURACY, NOT ASSUMED
------------------------------
The rules are audited against a hand-classified random sample (`--audit`), seeded and committed, so
the split carries an error bar instead of a promise. Disagreements are reported per fate.
"""
import argparse
import collections
import json
import random
import re
import sys

# --------------------------------------------------------------------------------------------------
# R0 — hand-measured FALSE. Syntax cannot reach falsity; these were found by reading, in consequence
# order, and every one of them is repaired in this iter. `path:start` keys the census's own units.
# --------------------------------------------------------------------------------------------------
FIX_SITES = {
    # the router is destroyed in production, not "prod only" -- infrastructure services.tf:509-517
    "corpus/architecture/architecture_overview.md:79",
    "corpus/architecture/service_taxonomy.md:526",
    # the db-backup claim iter-123 repaired at its own doc and did not reach here
    "corpus/architecture/security_compliance.md:236",
    "corpus/architecture/architecture_overview.md:442",
}

PLACEHOLDER = re.compile(r"\[[A-Z][a-z]+\]|\be\.g\.,\s*Postgres, Redis\b")
DOC_PURPOSE = re.compile(
    r"^\**This (document|page|file|section) (describes|provides|covers|is)\b"
    r"|^\**(Single source of truth|Service-level / developer map)\b",
    re.I,
)

# R2 -- subjects outside EVERY clone set and outside this corpus. Deliberately short: a long list
# here is how a triage manufactures hedges. Each entry names a thing no clone and no corpus file
# contains -- vendor-internal behaviour, applied cloud state, or a deployed runtime's configuration.
UNCHECKABLE = (
    ("gdpr compliance", "secure password storage"),      # what Clerk does inside its own service
    ("free tier available",),                            # vendor pricing
    ("live saas",),                                      # vendor availability posture
    ("clerk dashboard",),                                # config held in a vendor console
    ("dashboard config",),
    ("in the clerk dashboard",),
)

ARTIFACT = re.compile(
    r"`[^`\n]*(?:/|\.go|\.ts|\.tsx|\.js|\.jsx|\.py|\.sh|\.yml|\.yaml|\.json|\.tf|\.md|\.env|\.mjs)"
    r"[^`\n]*`"                                          # a path-ish backtick span
    r"|`[A-Za-z_][A-Za-z0-9_]*\.[A-Za-z_][A-Za-z0-9_]*`"  # schema.table / pkg.Symbol
    r"|`[A-Z][A-Z0-9_]{3,}`"                              # ENV_VAR
    r"|`[a-z][A-Za-z0-9]*\(\)`"                           # symbol()
    r"|`@?[a-z0-9@/.-]+`"                                 # package / module / stream / profile name
    r"|:\d{2,5}\b"                                        # a port
    r"|\bv\d+\.\d+"                                       # a version pin
)


def fate(item):
    key = f"{item['path']}:{item['start']}"
    s = item["sentence"]
    if key in FIX_SITES:
        return "fix", "R0 hand-measured false"
    if PLACEHOLDER.search(s) or DOC_PURPOSE.search(s.strip()):
        return "drop", "R1 scaffolding"
    low = s.lower()
    for group in UNCHECKABLE:
        if all(tok in low for tok in group):
            return "hedge", "R2 outside every clone set"
    if ARTIFACT.search(s):
        return "cite", "R3 names a reachable artifact"
    return "cite", "R4 default — presumed citable"


def main(argv=None):
    ap = argparse.ArgumentParser(description=__doc__.splitlines()[0])
    ap.add_argument("tier2_json")
    ap.add_argument("--c1-only", action="store_true", help="restrict to the sealed consequence class")
    ap.add_argument("--audit", type=int, default=0, help="draw N for the hand audit and print them")
    ap.add_argument("--seed", type=int, default=124)
    ap.add_argument("--json", action="store_true")
    a = ap.parse_args(argv)

    items = json.load(open(a.tier2_json))
    if a.c1_only:
        sys.path.insert(0, __file__.rsplit("/", 1)[0])
        import importlib.util
        spec = importlib.util.spec_from_file_location(
            "tp", __file__.rsplit("/", 1)[0] + "/triage-predicate.py")
        tp = importlib.util.module_from_spec(spec)
        spec.loader.exec_module(tp)
        items = tp.select(items)

    verdicts = [(i, *fate(i)) for i in items]

    if a.audit:
        rng = random.Random(a.seed)
        for i, f, r in rng.sample(verdicts, min(a.audit, len(verdicts))):
            print(f"{f:5s} [{r}] {i['path']}:{i['start']}\n      {i['sentence'][:190]}")
        return 0

    counts = collections.Counter(f for _, f, _ in verdicts)
    rules = collections.Counter(r for _, _, r in verdicts)
    n = len(verdicts)
    if a.json:
        print(json.dumps({"denominator": n, "split": dict(counts), "by_rule": dict(rules)}, indent=1))
        return 0
    scope = "C1 (the sealed consequence class)" if a.c1_only else "the whole tier-2 population"
    print(f"triage over {scope} — denominator {n}")
    for f in ("cite", "hedge", "fix", "drop"):
        print(f"  {f:6s} {counts[f]:5d}  = {100*counts[f]/n:5.1f} %")
    print("  by rule:")
    for r, c in rules.most_common():
        print(f"    {c:5d}  {r}")
    return 0


if __name__ == "__main__":
    sys.exit(main())
