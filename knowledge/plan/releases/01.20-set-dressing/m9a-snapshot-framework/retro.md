# M9a — Retro

**Summary:** Built the **`stack-snapshot` extension** — the dedicated, reusable framework that captures a
*public* reference surface once from a safe (non-primary-blocking) prod source, manifest-caches it in
`.agentspace` (never git), and replays it per-stack — with a **read-side tenant-data firewall**
(`AssertPublicOnly`, the analog of seeding's write-side `AssertClean`), a **production-safe capture-source
policy** (cache-hit → dump-ingest → safe primary read → AWS-gated upgrades), and a **snapshot-fidelity
data-DNA dimension** (5 two-sided operators that count the formerly-`waived` surfaces toward coverage). 9 Go
packages + the `stacksnap` CLI, proven end-to-end on the hermetic `reference-toy` surface; the `/db-query`
skill ported as the prod-read foundation. Shipped clean: **556 Go test funcs (+147)**, flake **0**, `-race`
clean, the close found 1 finding (a trailing tag) and zero code defects.

## Incidents this cycle
- **P3 (close finding, fixed):** the `stack-snapshot-m9a` tag sat at `2e0696d`, but the 2 harden commits +
  the pass-2 commit landed *after* it — so a per-stack clone pinning the tag would have gotten the
  pre-harden (test-only) code. Re-pointed the tag to HEAD `1cc4dd2` + force-pushed at close. No production
  impact (the delta was tests only); the lesson is to tag *after* the final harden pass, not at build-end.
- No P0/P1/P2. No regressions. The harden pass (2 passes) surfaced zero production-code defects — every gap
  was untested behavior, not incorrect behavior.

## What went well
- **The interface-seam discipline paid for itself.** Putting the DB behind small `Capturer` / `Replayer` /
  `FidelityProbe` / `SnapshotStore` interfaces let the *entire* orchestration (plan → firewall → capture →
  serialize → replay → reindex) be tested hermetically against fakes, with the live-DB code isolated to a
  thin `pg.Conn`. That is why coverage is 90–100% on every logic package and the residual uncovered code is
  *irreducibly* live-DB-bound — a clean, defensible stop condition.
- **Defense-in-depth on the load-bearing safety.** The firewall runs twice (plan-time `AssertPlan` +
  post-capture `AssertCaptured`), and `capture.Run` **stashes every payload in memory until all tables
  pass** — so a single leaked tenant row aborts with nothing written. Adversarial review confirmed the
  atomicity holds; `FuzzSanitize` (6.3M execs) confirmed the path-traversal guards.
- **Fate-2 partitioning kept the milestone tight.** M9a is the *framework* milestone; the three real
  surfaces (taxonomy / content / recipes) are cleanly owned by M9b/M10/M11, and the cloud store is a
  pre-signed v1.3 swap behind the stable `SnapshotStore` interface. The deferral audit was GREEN with zero
  repeats — the M7a→M7c precedent applied well.
- **The MVCC correction made the capture-source policy honest.** Recognizing that a read-only `COPY` never
  blocks writers turned "safe primary read" from a scary last resort into a sanctioned fallback, and the
  bounded read-only session (`SET TRANSACTION READ ONLY` + timeouts) caps the impact.

## What didn't / constraints
- **dump-ingest selects but still needs a DSN in M9a.** The offline pg_dump *reader* is deferred to the M9b
  path (the dump-format work belongs with the first real surface); for M9a, `dump-ingest` resolves but the
  CLI requires `--dsn` to stream from a restored dump. Documented in the CLI + spec — not a gap, a scoped
  boundary.
- **The live capture path is proven only on the toy surface.** The real ~2.1 GB taxonomy read (and the
  pgvector index-rebuild cost) is M9b's job; M9a proves the *mechanism* hermetically + on `reference-toy`,
  not against prod scale. That is the intended split.
- **Two-repo close choreography.** The code lives in the gitignored extensions clone (its own git, tagged +
  pushed); the rosetta merge carries only docs + records. The tag-trails-HEAD finding is a symptom of that
  split — worth a checklist note for M9b: tag last.

## Carried forward → M9b
- Prove the framework on the **real public taxonomy** — `skiller.{categories,specializations,skills,
  job_roles}` filtered `organization_id IS NULL`, embeddings/translations via the public parent, bulk-`COPY`
  replay, **rebuild the ~689 MB pgvector index on replay** (carry vectors verbatim).
- Wire the **offline dump-ingest reader** (the M9a-scoped boundary above) as part of the first real capture.
- **Tag after the final harden pass**, not at build-end (the one close finding).

## Metrics delta (from metrics.json)
- **Go test funcs:** 409 → **556** (+147; stack-snapshot 128 new, stack-seeding 145 → 164).
- **Flake:** 0 (5× shuffled sequential, both modules). **Race:** clean. **gofmt + vet:** clean.
- **Coverage (post-harden):** firewall/reference/replay/source 100%; manifest 98.2%; store 90%; capture
  89.7%; cmd/stacksnap 60.1%; pg 52.5% (rest live-DB-bound).
- **Review:** 1 finding (tag re-point), 0 code defects, 5 adversarial scenarios recorded, deferral audit GREEN.
