# iter-154 — decisions

## `D-M257x-154-1` — the class had been HALF-repaired twice, and each half-repair was documented

Both bring-ups derived the platform's service set correctly and then hand-appended literals. The history
is the finding:

- **iter-55** replaced the demo side's hand tuple with `platform_topology.py` and **did not carry the
  repair to its dev twin**. `dev-stack`'s comment says so in its own words 30 iters later.
- **harden pass 3** replaced the dev side's hand tuple with the same derivation — and left **both** sides'
  conditional hand-appends (`next-web-app studio-desk` / `directus`) standing.

So the *base* set was derived twice and the *conditional* set was never derived at all. Each repair wrote
down that its sibling had the same defect, and neither wrote a fence over the pair. `§5` rule 69, in the
plainest instance the milestone has produced: **an observation about a twin is not a fence over it.** Both
are re-pointed here in one commit, with one fence covering the enumerated set.

## `D-M257x-154-2` — the conditional was right; the SET was wrong, and only a generator could say so

`[ "$NO_UI" = 1 ] || verify_svcs="$verify_svcs next-web-app studio-desk"` encodes a true fact — the UI
tier tracks `NO_UI` — with a false enumeration. Measured against `gen_injected_override.py` at platform
`0c91421`, `--no-ui` drops **three** services: the two named plus **`hiring-app`**, the surface
`pt-hiring-recruiter-compare` plays through.

**No amount of reading that line finds this.** The line is internally consistent, its comment is correct,
and its gate is correct. The defect is only visible by asking the generator what it actually emits — which
is `TOK-08`'s whole thesis, arriving at a site nobody had censused because the site *looked* derived.

## `D-M257x-154-3` — the second consecutive iter to find a fence pinned to a SPELLING

iter-153 re-pointed harden pass 35's disclosure fence because it asserted `generate.sh`'s **source**
carried three service-name literals. iter-154 hit the identical shape one section over:
`dev-stack/tests/test_dev_stack.py::test_verify_scopes_directus_only_when_local_content_on` asserted the
literal source line `[ "$local_content" = 1 ] && verify_svcs="$verify_svcs directus"`.

Both fences protected a **real** property. Both encoded it as the **current spelling** of the code that
happened to implement it. Twice in two iters, the correct repair was the same: **re-point to the property,
never delete** — here, the property (`a prod-read dev stack must not be probed for a directus it does not
run`) moves to a behavioural arm that executes the real tail against a real override, and what stays in
the body-text contract is the structural half such a test can honestly assert (*it calls the union*, and
*no verify-scope line names a service*).

**The generalisation, and it is worth more than either fix:** a fence written by quoting the line it
guards will fail on the day that line is IMPROVED, and its failure is indistinguishable from the day that
line is BROKEN. That is not a reason to write fewer fences; it is a reason to write them against
behaviour. Both of this iter's own fences execute their subject.

## `D-M257x-154-4` — non-fatal on every branch, and it is asserted rather than intended

The verify net must never abort a bring-up (M18/M19). The union is a subprocess call, so it has three new
failure modes the tuple did not have: the script is missing, the script fails, or there is no override.
All three leave `verify_svcs` exactly as the platform derivation left it — i.e. pre-iter-154 behaviour
minus the tuple — and all three are **tested**, including a deliberately unresolvable `$HERE`.

One consequence is recorded here rather than left implicit: `scope-union.sh` is invoked **by path**, so
its executable bit is load-bearing on both bring-ups. A lost mode bit would silently leave every stack's
scope un-unioned, with no error — the quiet failure this whole thread is about. Asserted
(`test_the_subjects_exist`).

## `D-M257x-154-5` — the demo suite's 9 failures reproduce iter-147's baseline BY NAME

`demo-stack` closed at **9 failed · 1,055 passed**, and the nine are identical by name to iter-147's
recorded baseline: `test_ssr_origin_chain` ×3 + `test_demopatch` ×2 + `test_ant_academy` ×1 (six
sha-baseline drift, the routed `FIX-M257x-iter145-sha-baseline-drift`) and `test_migrate_race_live` ×3
(host environment — needs a live postgres, `-iter145-migrate-race-needs-a-host-postgres`). **Graded before
quoted**, per `D-M257x-144-2`; zero regressions attributable to this iter.
