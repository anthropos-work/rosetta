#!/usr/bin/env python3
"""iter-124 consequence predicate C1 — SEALED before the first triage verdict.

Reproducible denominator for the tier-2 triage. Consumes the census's own
`--tier2-export` JSON; prints |C1| and its per-file spread. Ordering by
CONSEQUENCE, not by class, is the user's rule for run 80.

  cd .agentspace/rosetta-extensions/stack-core
  /usr/bin/python3 claim_census_guard.py --repo-root <rosetta> --census \
      --tier2-export /tmp/tier2.json
  /usr/bin/python3 <this> /tmp/tier2.json
"""
import collections
import json
import re
import sys

C1 = re.compile(
    r"\b(auth\w*|token|secret|password|credential\w*|encrypt\w*|TLS|HTTPS|SSL|PII|GDPR"
    r"|residen\w+|tenan\w+|isolat\w+|backup\w*|retention|RBAC|ABAC|Casbin|JWT|session"
    r"|permission\w*|unauthenticated|public|firewall|VPN|CORS|Clerk|privacy|compliance"
    r"|SOC ?2|ISO ?27001|DPA|sub-?processor|delete\w*|erasure|audit)\b",
    re.I,
)


def select(items):
    return [i for i in items if C1.search(i["sentence"])]


def main(argv):
    items = json.load(open(argv[1]))
    sel = select(items)
    print(f"tier-2 population : {len(items)}")
    print(f"C1 (consequence)  : {len(sel)} = {100*len(sel)/len(items):.1f} % of the population")
    per = collections.Counter(i["path"] for i in sel)
    print(f"spread            : {len(per)} files")
    for p, n in per.most_common():
        print(f"  {n:4d}  {p}")
    return 0


if __name__ == "__main__":
    sys.exit(main(sys.argv))
