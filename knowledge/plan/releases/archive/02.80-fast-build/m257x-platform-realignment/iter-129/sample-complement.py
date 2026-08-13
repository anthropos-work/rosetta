#!/usr/bin/env python3
"""iter-129 — the two blocks iter-128 named UNTESTED rather than clean.

iter-128 published the complement's split as `cite` ≈ 90.0 %, corrected from a printed 99.4 %
by an R4-only hand audit. It named two limits of that number in its own words:

  * **R3 — 486 members, 59 % of the complement — was NOT re-audited.** iter-124's 100 % (21/21,
    measured on C1) was carried forward. The largest untested block inside the corrected figure.
  * **`fix = 0` is a FLOOR of unknown height.** The triage cannot decide falsity, and the
    complement was never read for it.

This file pays the first debt and slices the population for the second. It **imports**
`iter-124/triage-predicate.py`, `iter-124/triage.py` and `iter-128/triage-complement.py`'s
partition rather than re-deriving any of them, so iter-124's and iter-128's published figures
stay byte-reproducible and no drift is possible between the three triages.

  /usr/bin/python3 sample-complement.py <tier2.json> --audit 30 --seed 129 --rule R3
  /usr/bin/python3 sample-complement.py <tier2.json> --slices 4 --out-prefix comp-slice
"""
import argparse
import collections
import importlib.util
import json
import random
import sys
from pathlib import Path

_HERE = Path(__file__).resolve().parent
_M = _HERE.parent


def _load(name, path):
    spec = importlib.util.spec_from_file_location(name, str(path))
    mod = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(mod)
    return mod


def main(argv=None):
    ap = argparse.ArgumentParser(description=__doc__.splitlines()[0])
    ap.add_argument("tier2_json")
    ap.add_argument("--audit", type=int, default=0)
    ap.add_argument("--seed", type=int, default=129)
    ap.add_argument("--rule", default=None, help="restrict the audit pool to a rule prefix, e.g. R3")
    ap.add_argument("--slices", type=int, default=0, help="write N balanced falsity-read slices")
    ap.add_argument("--out-prefix", default="comp-slice")
    a = ap.parse_args(argv)

    tp = _load("tp", _M / "iter-124/triage-predicate.py")
    tr = _load("tr", _M / "iter-124/triage.py")

    items = json.load(open(a.tier2_json))
    c1 = tp.select(items)
    c1_keys = {id(i) for i in c1}
    comp = [i for i in items if id(i) not in c1_keys]
    # Same fail-closed partition assert iter-128 shipped. If this ever trips, every number below
    # is meaningless and the run must stop rather than publish one.
    assert len(c1) + len(comp) == len(items), (
        f"partition broken: |C1|={len(c1)} + |comp|={len(comp)} != {len(items)}")

    verdicts = [(i, *tr.fate(i)) for i in comp]
    rules = collections.Counter(r for _, _, r in verdicts)
    print(f"tier-2 population        : {len(items)}")
    print(f"C1 (sealed, iter-124)    : {len(c1)}")
    print(f"COMPLEMENT               : {len(comp)}")
    for r, c in rules.most_common():
        print(f"   {c:5d}  {r}")

    if a.audit:
        pool = verdicts
        if a.rule:
            pool = [v for v in verdicts if v[2].startswith(a.rule)]
        rng = random.Random(a.seed)
        picked = rng.sample(pool, min(a.audit, len(pool)))
        print(f"\n# audit sample: {len(picked)} of {len(pool)} "
              f"{a.rule or 'complement'} members, seed {a.seed}\n")
        for n, (i, f, r) in enumerate(picked, 1):
            print(f"{n:3d}. {f:5s} [{r}] {i['path']}:{i['start']}\n     {i['sentence'][:320]}\n")
        return 0

    if a.slices:
        # Slice by FILE so a reader holds one document's context at a time; balance by count.
        per = collections.defaultdict(list)
        for i, f, r in verdicts:
            per[i["path"]].append((i, f, r))
        buckets = [[] for _ in range(a.slices)]
        for path in sorted(per, key=lambda p: -len(per[p])):
            buckets.sort(key=len)
            buckets[0].extend(per[path])
        for n, b in enumerate(buckets):
            name = f"{a.out_prefix}-{chr(ord('1')+n)}.txt"
            (_HERE / name).write_text(
                "\n".join(f"{i['path']}:{i['start']}-{i['end']}  [{r}]  {i['sentence']}"
                          for i, f, r in sorted(b, key=lambda t: (t[0]['path'], t[0]['start'])))
                + "\n")
            print(f"  wrote {name}: {len(b)} sentence(s), "
                  f"{len({i['path'] for i, _, _ in b})} file(s)")
        assert sum(len(b) for b in buckets) == len(comp), "slice partition broken"
    return 0


if __name__ == "__main__":
    sys.exit(main())
