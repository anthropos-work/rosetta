import re, sys, ast
from pathlib import Path
sys.path.insert(0, "/Users/marco/workspace/anthropos/rosetta/.agentspace/rosetta-extensions/stack-core")
import derivation_registry as dr
ROOT = Path("/Users/marco/workspace/anthropos/rosetta/.agentspace/rosetta-extensions")
LIVE = dr._MEASURED_RE
E = r"(?:\*+|_+)"          # markdown emphasis run ONLY — backticks deliberately excluded
VARIANTS = {
 "close-only":  r"(?<![\w.\-§#])(\d[\d,]*)"+E+r"?[\s\-]+(?:of\s+\d[\d,]*"+E+r"?\s+)?("+dr._MEASURED_NOUNS+r")\b",
 "open-only":   r"(?<![\w.\-§#])(\d[\d,]*)[\s\-]+(?:of\s+\d[\d,]*\s+)?"+E+r"?("+dr._MEASURED_NOUNS+r")\b",
 "both":        r"(?<![\w.\-§#])(\d[\d,]*)"+E+r"?[\s\-]+"+E+r"?(?:of\s+\d[\d,]*"+E+r"?\s+"+E+r"?)?("+dr._MEASURED_NOUNS+r")\b",
}
def newhits(rx):
    R=re.compile(rx, re.IGNORECASE); rows=[]
    for path in sorted(ROOT.rglob("*.py")):
        if any(p in dr._CENSUS_SKIP for p in path.parts): continue
        src=path.read_text(encoding="utf-8",errors="replace")
        try: tree=ast.parse(src)
        except SyntaxError: tree=None
        rel=path.relative_to(ROOT).as_posix(); kind="test" if path.name.startswith("test_") else "src"
        for text,resolver in dr._measurement_units(src,tree):
            live=[(m.start(),m.end()) for m in LIVE.finditer(text)]
            for m in R.finditer(text):
                if any(m.start()>=s and m.start()<e for s,e in live): continue
                ctx=text[max(0,m.start()-45):m.end()+22].replace("\n"," ")
                rows.append((kind,rel,dr._unit_line(resolver,m.start()),dr._classify_measurement(text,m),m.group(0),ctx))
    return rows
for name,rx in VARIANTS.items():
    rows=newhits(rx)
    print(f"--- {name}: NEW {len(rows)}  src={sum(1 for r in rows if r[0]=='src')} test={sum(1 for r in rows if r[0]=='test')} "
          f"standing={sum(1 for r in rows if r[3]=='standing')}")
rows=newhits(VARIANTS["both"])
print("\n=== 'both' hits ===")
for r in rows: print(f"{r[0]:4s} {r[1]}:{r[2]} [{r[3]}] {r[4]!r}\n       …{r[5]}…")
