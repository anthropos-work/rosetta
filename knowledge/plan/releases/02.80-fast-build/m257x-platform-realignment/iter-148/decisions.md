# iter-148 — decisions

## `D-M257x-148-1` — a REPORT discloses where a BRING-UP refuses, and the split is the same rule applied twice

iter-147 made `rosetta-demo`'s bring-up `die` on an underivable profile while `cmd_down` keeps its
non-fatal fallback (`D-M257x-147-2`). The same question arrives here with a third answer.

`generate.sh` **derives** the probe scope, and when it cannot it **discloses and continues** — it does
not refuse. A bring-up that cannot name its scope must not announce a stack; a **report generator** that
cannot name its scope must still produce the report, because refusing leaves the operator with nothing
at all where disclosing leaves them with a caveated reading. The caveat is printed **into the markdown**,
not only to stderr: the report is what gets read and stderr is what gets lost.

Three dispositions, one rule: **the disposition follows what the artifact is FOR.**

## `D-M257x-148-2` — the disclosure is fenced DERIVED-vs-DECLARED, because a warning that names the wrong services is worse than none

The unscoped-run warning names `cms`, `jobsimulation`, `storage`, `roadrunner`. That is a **hand-written
set** — the exact shape iter-145 found rotting as six count literals for four months.

So the fence asserts, both directions, that **`{registry rows} − {services the platform's compose
declares} − {rows this tooling itself injects}` equals the set the disclosure names**, and fails naming
what went away and what is newly orphaned. When the next fold lands, the warning goes RED instead of
quietly reassuring a reader about the wrong four services.

The injected set (`next-web-app`, `studio-desk`, `directus`) is excluded **by name and with the reason**:
those rows are appended by the injection generator, not by the platform, so their absence from the
platform file is expected and is not a merge.

## `D-M257x-148-3` — the RED-proof caught a defect in THIS fence, and that is recorded rather than quietly fixed

The first draft asserted *"the derivation appears before the `verify.sh` invocation"* using
`src.find("platform_topology.py")` and `src.find("live/verify.sh")`. **Both bound to comments** —
`generate.sh`'s usage header names `verify.sh` 88 lines above the command, and the comment block
explaining this very fix names `platform_topology.py`. The check would have kept passing after the code
it guards was deleted.

It was caught by the mutation control, not by review: the de-scoped mutant retained the comment and the
assertion held. **`§5` rule 67 / rule 68(d)'s axis — the same token in a comment and in a command carries
opposite obligations — reproduced inside the fence written to apply it, twice in one iter.** Recorded
because the lesson is not "be careful": it is that **a form-matching fence must bind to executable
content by construction**, and the only thing that reliably detects the miss is a mutation control that
actually runs.

## `D-M257x-148-4` — the `postgres-schemas` failure is a SUBSTRATE artifact and is NOT booked as a defect

It fails in **both** arms with *"cannot derive the expected schema set"*. Its second candidate path is
`stack-verify/lib/../../../platform/repos.yml`, which resolves correctly in a **per-stack consumption
copy** (`stack-demo/rosetta-extensions/…` → `stack-demo/platform/repos.yml`, verified present) and
**not** in the `.agentspace` authoring copy this iter measured from.

Excluded per `D-M257x-122-4` — *before believing a defect, read the substrate line*. The probe is
behaving correctly: it refuses to assert a hand-maintained list it cannot derive, which is the
milestone's own §2 discipline.
