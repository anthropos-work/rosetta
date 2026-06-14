# M27 — Decisions

_Implementation decisions with rationale, numbered `M27-D1`, `M27-D2`, … . Filled during build._

## M27-D1 — Milestone renumber: M26→M27 (secret-provisioning release shifted to M27→M30)

**Date:** 2026-06-14. **Status:** RESOLVED (user decision).

When the build sub-agent went to author this milestone's code in the `rosetta-extensions` authoring copy
(`.agentspace/rosetta-extensions/`), it found the flat milestone number **M26 already consumed** by an orphaned,
untracked ext effort:
- branch `m26/self-contained-demo` @ `25ab855`, **tagged `prop-room-m26`**, committed 2026-06-14 13:21 (after the
  v1.5 close 07:37, before the v1.6 design) — *"M26: make demo stacks self-contained (own GitHub clone set, like
  stack-dev)"*, +521/−141 across 12 files in `demo-stack/` + `stack-injection/` (notably `ensure-clones.sh +106`);
- local-only (never pushed), unmerged to ext `main`, and absent from the rosetta roadmap/state.

The v1.6 "stage door" secret-provisioning release had been designed (minutes later, from a state.md that read
v1.5 = M21→M25) as **M26→M29**, colliding on M26.

**Decision (user, via the work-milestone blocker escalation):**
1. **Keep `self-contained-demo` as the real M26** — its `prop-room-m26` tag + branch stay; it awaits its own
   `/developer-kit:design-roadmap` pass for a version + scope (a separate task, not part of v1.6).
2. **Renumber the secret-provisioning release to M27→M30** — M27 secret-coverage-dna (this milestone) → M28
   provisioning-engine → M29 docs+skill → M30 field-bake. Roadmap, state, context, vision, and the scaffold dirs
   were shifted accordingly.
3. **The stray uncommitted ext WIP** (`clerkenstein/knowledge/architecture.md`, +32 lines, browser-login handshake
   docs) was preserved by committing it on a dedicated ext branch (`wip/clerkenstein-browser-login`), leaving the
   authoring tree clean.
4. **This milestone (M27) is authored on a fresh ext branch off `main`** (`stage-door-m27`), tagged `stage-door-m27`
   on completion — never on top of the stale `m26/self-contained-demo` branch.

Note the future interaction: self-contained-demo touched `ensure-clones.sh`, which M28's provisioning engine plans
to extend (the single `cp` that M28 folds into `stacksecrets provision`). Whoever lands self-contained-demo's
roadmap home should coordinate with M28.

## M27-D2 — The keep-listed gate is DNA-scoped, not example-mirror-scoped

**Date:** 2026-06-14. **Status:** RESOLVED (surfaced + fixed during build, Fate 1).

Running `stacksecrets diff --stack-root stack-dev` with a naive "every key declared in any
`.env.example` must be a gene" gate produced **111 unlisted-required** findings — because the
`.env.example` files mix the curated *required secrets* the DNA tracks (Clerk pair, GH_PAT, OpenAI,
Directus token, …) with a large body of **config/wiring/optional** lines (Sentry DSNs, PostHog keys,
feature flags, sign-in URLs, AI tuning params like `AI_REQUEST_TIMEOUT_MS`, ports). Treating all of
them as gate-fatal makes the gate permanently red and would pressure listing pure config as secret
genes — the opposite of a curated coverage DNA.

**Decision:** the keep-listed gate (`diff` exit 1) is scoped to the DNA's **own required-secret
universe**, not a 1:1 mirror of every example line. Concretely, the gate-fatal `unlisted-required`
class fires only when a declared key is a **known-required secret** — defined as: a key the DNA
ALREADY lists as a non-waived gene for **some** repo — that appears (declared) for **another** repo
where the DNA does NOT list it. That is exactly the real drift the roadmap's anti-vacuous-green guard
targets: "you added an already-tracked required secret (e.g. `CLERK_SECRET_KEY`) to a new repo's env
but forgot to add its gene there." A brand-new key that has never been a DNA gene anywhere is reported
as **informational** (`unlisted-candidate`), not gate-fatal — it is a candidate for a human to triage
into the curated DNA (or to ignore as config), surfaced by `diff` but never silently auto-promoted
(the DNA stays hand-curated). The `undeclared-gene` asymmetry stays informational as before.

This keeps the gate honest AND usable: it catches the dangerous omission (a tracked secret missing a
gene → vacuously-green coverage) without policing the config noise that legitimately lives in
`.env.example` but isn't a coverage-tracked secret.

## M27-D3 — Implementation choices: stdlib-only, `check`/`measure` folded in, introspect read-only

**Date:** 2026-06-14. **Status:** RESOLVED (build choices).

1. **Stdlib-only module.** Unlike `stack-seeding`/`stack-snapshot` (which need pgx for the live DB), the
   secret-DNA touches no database — the source is `.env`-shaped files on disk. The module is therefore
   dependency-free: JSON for the DNA, `archive/zip` + `bufio` for ingestion. No `go.sum`. This keeps the
   values-blind audit surface trivially small (no third-party code can see a value).

2. **`check`/`measure` folded into M27.** The roadmap puts the coverage *gate wiring* + `provision` in M28, but
   the `check`/`measure` SCORER (run a source against the DNA → Overall/Critical/per-repo rollup) is the natural
   pairing with the DNA and the source reader this milestone already builds — withholding it would be a partial
   landing (the DNA would have no way to be exercised). So `check` (alias `measure`) ships in M27, values-blind,
   with the exit-1-if-critical-<100% contract. M28 still owns `provision` (writing target `.env` files) + the
   non-fatal pre-flight WIRING into `/dev-up`+`/demo-up` + the demo-aware (Clerkenstein-minted) scoring. Fate-1:
   the scorer belongs with its DNA; the gate *plumbing* belongs with the engine.

3. **`introspect --write` is report-only on the gene set.** The DNA is a hand-curated, committable artifact;
   `introspect` surfaces what the hybrid source declares vs what the DNA lists (and `diff` gates on it), but it
   never auto-adds/edits genes. A new declared key is a deliberate human review (curate or ignore-as-config),
   not an automatic mutation — this is what keeps the 55-gene map trustworthy.

4. **Profile scoping (the overview open question): settled to `graphql`.** The DNA carries an explicit
   `profile` field (`"graphql"`); the denominator is scoped to the default profile, with messenger/customerio
   profile keys modeled as `waived-profile-gated`. A per-gene profile tag was considered and deferred — not
   needed for v1; the waived-class device covers the profile-gated keys cleanly. (M28 may revisit if it wires
   non-default-profile bring-ups.)

5. **Encrypted-zip (age/gpg) deferred (the overview open question).** v1 ships plain zip + dir. Encrypted-zip
   support is genuinely out of M27 scope (a new crypto dependency + key-management surface) and no consumer
   needs it yet — recorded here as a known v1 boundary, not a silent gap. If a future need arises it attaches to
   a later milestone; for now the plain dir/zip covers the stack-dev field-bake (M30).

## Adversarial review (close-milestone Phase 2c)

Scenarios considered against the shipped `stack-secrets/` section — the *scenario*, not the fix (the harden
pass had already turned each into a regression test; close confirmed the coverage, no new code fix needed):

- **A leaky value escapes through a "shape" probe or an error message.** The one function that touches a value
  (`ClassifyShape`) is structurally constrained to return a `ValueShape` token; `parseEnv`/`readKeyNames` discard
  `line[eq+1:]` as a local. The harden pass pins this with a 200 KB value + `=`-in-value + quote-only adversarial
  extraction test and an end-to-end `check`/`diff` regression that asserts zero leakage to stdout/stderr against
  the real 55-gene DNA. Re-verified at close by grep (no secret-shaped token in `secret-dna.json`; only one
  value-reading call site, immediately discarded). **Handled.**
- **A stale-backup mirror (`zEnvs/`) or a stray top-level `.env` is silently ingested as "the source".** Ingestion
  is DNA-driven, never glob-driven — the reader opens exactly `<root>/<repo>/<target_file>`, so a non-DNA path is
  structurally invisible. The zip path's suffix-match is `/`-bounded (`matchTarget`), pinned by the
  "evil-app/myapp/notstudio-desk must NOT resolve" decoy test. **Handled by construction.**
- **A malformed/oversized line crashes or silently false-cleans the scan.** `bufio.Scanner` carries a 1 MiB cap;
  an oversized line surfaces `sc.Err()` and is propagated as a fail-loud parse error (FromDir/FromZip and
  ReadDeclaredKeys), pinned by oversized-line tests — never a silent skip. A `ReadDeclaredKeys` read-error in
  `diff`/`introspect` returns `exitFail`, not a false-clean exit 0. **Handled.**
- **The gate goes vacuously green** (DNA under-lists a tracked secret → `measure` reports 100% on a short stack).
  Two defences: `Validate`'s anti-vacuous-100 guard (≥1 required+critical gene) rejects an empty-critical DNA at
  load; the two-tier `diff` keep-listed gate (M27-D2) exits 1 when an already-tracked secret is declared for a
  repo with no gene there. Both pinned. **Handled.**
