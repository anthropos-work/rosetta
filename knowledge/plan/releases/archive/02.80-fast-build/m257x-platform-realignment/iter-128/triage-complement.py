#!/usr/bin/env python3
"""iter-128 — triage the COMPLEMENT of the sealed consequence class.

iter-124 triaged C1 (the 340 tier-2 sentences whose text touches security, data-handling,
identity, tenancy, residency, backup or access). It named the remaining 820 **untriaged, not
implied clean** — which was the right call and is the debt this file pays.

WHY A SEPARATE FILE INSTEAD OF A FLAG ON `triage.py`
----------------------------------------------------
`iter-124/triage.py` and `iter-124/triage-predicate.py` are SEALED: their outputs are published
figures. This file **imports both and adds no logic of its own** — the same `C1` regex draws the
partition and the same `fate()` assigns the verdicts, so:

    |C1|  +  |complement|  ==  |tier-2 population|          (asserted at runtime, fail-closed)

Any drift between the two triages would therefore be a drift in the shared instrument, not in two
copies of it. The committed iter-124 artifacts are not modified by a single byte.

THE AUDIT IS THE POINT, NOT THE SPLIT
-------------------------------------
The split is cheap and will read ~99 % `cite`. iter-124's own audit showed why that number cannot
be published as-is: rule **R4** (the generous default, applied when a sentence names no artifact)
was wrong **1 time in 3** in the C1 sample. `--audit N` draws a seeded sample from the COMPLEMENT
so R4's accuracy is re-measured **on this population** rather than imported from C1's.
"""
import argparse
import collections
import importlib.util
import json
import random
import sys
from pathlib import Path

_ITER124 = Path(__file__).resolve().parent.parent / "iter-124"


def _load(name, path):
    spec = importlib.util.spec_from_file_location(name, str(path))
    mod = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(mod)
    return mod


def main(argv=None):
    ap = argparse.ArgumentParser(description=__doc__.splitlines()[0])
    ap.add_argument("tier2_json")
    ap.add_argument("--audit", type=int, default=0, help="draw N from the COMPLEMENT for hand audit")
    ap.add_argument("--seed", type=int, default=128)
    ap.add_argument("--r4-only", action="store_true", help="draw the audit sample from R4 members only")
    ap.add_argument("--json", action="store_true")
    a = ap.parse_args(argv)

    tp = _load("tp", _ITER124 / "triage-predicate.py")
    tr = _load("tr", _ITER124 / "triage.py")

    items = json.load(open(a.tier2_json))
    c1 = tp.select(items)
    c1_keys = {id(i) for i in c1}
    comp = [i for i in items if id(i) not in c1_keys]

    # Fail closed: the partition must be exact, or every number below is meaningless.
    assert len(c1) + len(comp) == len(items), (
        f"partition broken: |C1|={len(c1)} + |comp|={len(comp)} != {len(items)}")

    verdicts = [(i, *tr.fate(i)) for i in comp]

    if a.audit:
        pool = [v for v in verdicts if v[2].startswith("R4")] if a.r4_only else verdicts
        rng = random.Random(a.seed)
        picked = rng.sample(pool, min(a.audit, len(pool)))
        print(f"# audit sample: {len(picked)} of {len(pool)} "
              f"{'R4-only' if a.r4_only else 'complement'} members, seed {a.seed}\n")
        for n, (i, f, r) in enumerate(picked, 1):
            print(f"{n:3d}. {f:5s} [{r}] {i['path']}:{i['start']}\n     {i['sentence'][:300]}\n")
        return 0

    counts = collections.Counter(f for _, f, _ in verdicts)
    rules = collections.Counter(r for _, _, r in verdicts)
    n = len(verdicts)
    if a.json:
        print(json.dumps({"population": len(items), "c1": len(c1), "complement": n,
                          "split": dict(counts), "by_rule": dict(rules)}, indent=1))
        return 0

    print(f"tier-2 population        : {len(items)}")
    print(f"C1 (sealed, iter-124)    : {len(c1)}")
    print(f"COMPLEMENT (this triage) : {n}   <- the 'untriaged, not implied clean' set")
    print(f"\ntriage over the complement — denominator {n}")
    for f in ("cite", "hedge", "fix", "drop"):
        print(f"  {f:6s} {counts[f]:5d}  = {100*counts[f]/n:5.1f} %")
    print("  by rule:")
    for r, c in rules.most_common():
        print(f"    {c:5d}  {r}")
    return 0


if __name__ == "__main__":
    sys.exit(main())
