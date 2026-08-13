# iter-05 — decisions

## D1 — the host reference is repaired in place; targets are untouched

**Decision.** Re-point the `exit_gate` from `odysseus` (retired, `D-v28-15`) to `macmini`, and change
nothing else about what the gate demands.

**The test applied.** *Does the edit change what "done" means?* A hostname that names a machine which no
longer exists makes the gate un-gradeable — repairing it restores gradeability and nothing more. A target
(≤ 360 s), a rep count (3), a clause (HEADROOM / ISOLATION / G1–G7 / 0 platform edits) or the stretch
(≤ 300 s) *would* change it, and none were touched. `re_scope_trigger`'s 420 s and its whole derivation are
kept literal, with the standing requirement that it be re-derived against `macmini.json`'s `gated_baseline`
once measured — which is exactly what the odysseus-era text already demanded of odysseus.

**Explicitly not a relaxation.** iter-04's arithmetic prices this host at ~420–455 s pre-lever with L1 worth
~136–152 s here. The re-cut therefore points the gate at a host where the *unchanged* 360 s cut is more
reachable than the premise that paused the milestone assumed — the opposite direction from a relaxation.

## D2 — one units definition added to HEADROOM clause 1, declared rather than buried

**Decision.** The gate now states that clause 1's `cores` means *the logical-core count of the machine the
`load1` sample is taken on* — on a `docker-desktop-vm` host, the HOST's count, not the VM allocation.

**Why this is a correction and not a loosening.** `buildbench.py` samples `os.getloadavg()` **on the macOS
host** (12 logical cores) and grades it against `profile["cores"]`, which the profile's own `budget_source`
declares to be *"the Docker Desktop VM allocation"* (8). Two machines' quantities in one comparison. The
same file already draws this exact distinction correctly in **two** other places — `engine_facts()`'s
docstring spells out *"on the M257x sanctioned dev host `os.cpu_count()` is 12 and `docker info` NCPU is 8"*,
and `profile_describes_host()` grades cores *"against the quantity the profile's own `kind` declares it to
be"*. The instrument knows the difference twice and forgets it once.

**The direction it moves the number is stated, not hidden.** Correcting the units raises this host's clause-1
limit from **6** to **10**. That is a consequence of grading the right machine, not a decision to accept more
contention; the clause still fails the gate when tripped. It is recorded here so a reviewer who wants to
object has a single place to object at.

## D3 — retract the prediction, keep the decision

**Decision.** `DOC-M257-hostclass-retraction` retracts the *"a Mac pays no unpack leg"* claim at all three
sites (`state.md`, `roadmap.md` `D-v28-15`, this milestone's `overview.md`) while leaving `D-v28-15`'s actual
ruling — billion is the official demo host, dev/test is local to the new Mac, odysseus and the old laptop are
retired — **completely intact**.

**Why the split matters.** `D-v28-15` is a **user decision**, and a sub-agent does not amend one. But the
paragraph being corrected is not the decision; it is a *predicted cost* the decision was annotated with, and
it was derived by generalising a **per-machine Docker Desktop setting** from a **retired M1 Pro laptop** onto
a machine nobody had probed. Retracting a prediction is documentation work. Retracting the ruling would be
planning. The edits say which is which, in place, so the distinction survives the next reader.

**Evidence standard applied.** Probe, not config string. `docker info` reports `Storage Driver: overlayfs`
here — the exact reading that produced the wrong claim — so the retraction rests on the controlled two-size
unpack probe (0.8 s @ 256 MB → 3.0 s @ 1024 MB) plus the real 4.12 GB image (56.6 s export + 19.3 s unpack):
the leg exists **and** scales with bytes, which no naming coincidence produces.
