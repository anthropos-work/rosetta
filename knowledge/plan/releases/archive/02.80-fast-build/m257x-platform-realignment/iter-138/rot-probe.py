"""iter-138 Priority-2 probe: are SAME-FILE bare :NN pins mechanically rot-detectable?

Method, stated before the numbers:
  for each corpus/**.md line containing a bare `:NN` pin (backticked, no path head):
    1. find the commit that last touched the CITING line  (git blame)
    2. read the file at that commit; take the target line's text  -> X
    3. if X is non-blank and X still exists in the file at HEAD at a DIFFERENT line -> ROTTED
  Positive control: a pin whose target text is unchanged at the same line -> STABLE (must be >0)
  Negative control: X absent from HEAD entirely -> UNRESOLVABLE (reported, not counted as rot)
"""
import re, subprocess, sys, collections, os
ROOT="/Users/marco/workspace/anthropos/rosetta"
os.chdir(ROOT)
PIN=re.compile(r'`:(\d{1,4})`')
def sh(*a):
    return subprocess.run(a,capture_output=True,text=True).stdout

files=sh("git","ls-files","corpus").split()
files=[f for f in files if f.endswith(".md")]
stats=collections.Counter(); rotted=[]; stable=[]
for f in files:
    try: cur=open(f,encoding="utf-8").read().split("\n")
    except Exception: continue
    for i,line in enumerate(cur,1):
        pins={int(m) for m in PIN.findall(line)}
        if not pins: continue
        bl=sh("git","blame","-L",f"{i},{i}","--porcelain","--",f)
        if not bl: continue
        sha=bl.split()[0]
        if sha.startswith("0000"): continue
        old=sh("git","show",f"{sha}:{f}")
        if not old: continue
        olds=old.split("\n")
        for p in pins:
            stats["pins"]+=1
            if p>len(olds): stats["out-of-range-then"]+=1; continue
            X=olds[p-1].strip()
            if not X: stats["blank-then"]+=1; continue
            hits=[j for j,l in enumerate(cur,1) if l.strip()==X]
            if not hits: stats["target-text-gone"]+=1; continue
            if p in hits: stats["stable"]+=1; stable.append((f,i,p)); continue
            stats["ROTTED"]+=1; rotted.append((f,i,p,hits[0],X[:70]))
print("=== iter-138 rot probe ===")
for k,v in stats.most_common(): print(f"  {k:24s} {v}")
print(f"\nROTTED pins ({len(rotted)}), first 25:")
for f,i,p,now,x in rotted[:25]:
    print(f"  {f}:{i}  cites :{p} -> content now at :{now}   | {x}")
print(f"\ncontrol STABLE={stats['stable']} (must be >0 or the probe proves nothing)")
