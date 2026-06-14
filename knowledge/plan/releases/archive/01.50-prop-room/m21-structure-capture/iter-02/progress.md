**Type:** tik (first tik under TOK-01) — staged-pipeline build toward the binary serve-anonymously gate.

# M21 iter-02 — progress

## Work done (live, against Docker)
1. **Live baseline harness stood up** — throwaway `pgvector/pgvector:pg16` on port 55432 (stack offset N=5 / `dev-5`)
   + a real `directus/directus:11.6.1` bootstrapped on the `directus` schema over a shared docker network. Torn down
   at iter close (no persistent stack mutated).
2. **Stage 2 confirmed LIVE — and found broken, then fixed.** `node cli.js bootstrap` creates the **27** `directus_*`
   system tables (matches the static baseline exactly). But `createDefaultAdmin` **crashed**: `admin@dev-5.local` dies
   `FAILED_VALIDATION`. Isolated the rule on fresh-each-time bootstraps: Directus 11.6.1 **rejects the `.local` TLD**
   (`admin@dev-5.local` x, `admin@dev.local` x, `admin@example.com` ok, `admin@dev-5.example.com` ok -- hyphens/digits
   are fine, `.local` is not). The `provision.go` `DefaultEnvContract` minted `admin@<stack>.local` -> would crash at
   real M22 bring-up. **Fixed** -> `admin@<stack>.example.com` (RFC-2606 reserved), code + comment + tests + a `.local`
   guard. Re-verified live: bootstrap with the fixed email creates the admin (users=1). -> M21-D1.
3. **Baseline refined: the real pipeline replay exits 5, not 4.** The static baseline assumed replay against an *empty*
   directus schema (exit 4 `ErrEmptySchema`). But the real pipeline bootstraps FIRST, so replay runs against a
   **bootstrapped** schema -> `SchemaVersion` returns a non-nil digest -> cache miss at that digest -> **exit 5**.
   Empirically: bootstrapped digest `b4cb55bcee08c76f2c37980da460a683` != prod cache key `6cd35278edbc8a7962053a9d7ebfc480`.
   Confirmed both: empty schema -> exit 4; bootstrapped-but-gap schema -> exit 5. -> M21-D3.
4. **Digest trap crystallized.** `pg.SchemaVersionSQL()` digests *every column of every table* in the `directus`
   schema (27 system + content). The prod key `6cd35278...` encodes the **whole prod directus schema**; a per-stack
   bootstrap can only converge it if the *entire* schema (system tables at the same Directus version + ALL prod
   content collections + identical types) matches -- deeper than "apply structure before replay." -> M21-D5.
5. **Structure round-trip mechanism VALIDATED.** Directus's own `node cli.js schema snapshot` emits a clean YAML
   (`version/directus/vendor/collections/fields/relations`); `node cli.js schema apply <yaml>` of a 1-collection
   snapshot created the user table **and** its `directus_collections` + `directus_fields` registry rows in one step --
   exactly stage 3's need. -> M21-D2.
6. **Structure-source finding (TOK-01 lean partially falsified).** The 1-collection proof used *invented* types. The
   real 9-collection artifact needs prod's **actual column types** (the cached real rows COPY into typed columns at
   stage 4). Pure option (c) (self-contained reference) validates the format/mechanism but **cannot invent correct
   types** -- the real artifact needs a prod-structure read (a/b, or a committed platform schema, or an MCP structural
   read). Routed to iter-03 for resolution. -> M21-D4.

## Test discipline
`go build ./... && go vet ./... && go test ./...` across all 12 `stack-snapshot` packages -- **green** (the directus
package's updated email tests pass; nothing else regressed).

## Close -- 2026-06-11

**Outcome:** Live baseline established; stage 2 secured (the `.local` email bug fixed -- a real stage-2 live blocker);
baseline refined (real pipeline exits 5, not 4); digest trap crystallized; structure round-trip mechanism
(`schema apply` YAML) validated. furthest-passing-stage stays **2** (now LIVE-confirmed + secured) -- the full
9-collection structure artifact is routed to iter-03 pending the structure-source decision (not gamed by authoring a
fake-typed snapshot just to bump the ordinal).
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n -- (2) triggered-tok: n (this is a tik) -- (3) re-scope: n (1 under-target tik, not 2 consecutive; substantial deliverables landed) -- (4) user-blocker: n (the structure-source decision is a routed-forward Fate-3 strategy item under TOK-01; it shapes iter-03, it does not change code landing in THIS iter) -- (5) cap-reached: n (1 tik this session) -- (6) protocol-stop: n -- Outcome: continue
**Decisions:** milestone-root M21-D1 (email fix), M21-D2 (mechanism=schema-apply), M21-D3 (exit-5 baseline), M21-D4 (structure-source finding), M21-D5 (digest trap); iter-local in `decisions.md`.
**Side-deliverables:** none separable -- the email fix is in-scope stage-2 hardening (it makes stage 2 genuinely pass live), committed in rosetta-extensions `98e51b4`.
**Routes carried forward (-> iter-03, Fate-3 under TOK-01):**
  - Resolve the **structure-source** question -- first check the platform repos (cms) for a committed Directus schema
    snapshot / collection definitions (a self-contained, prod-faithful source needing no prod access); else weigh
    a/b/MCP-structural-read. Handler: `STRUCT-M21-iter03-source`.
  - Produce + apply the **real 9-collection structure artifact** (with prod-faithful types) -> advance stage 2 -> 3 (-> 4
    if replay of the cached rows then succeeds). Handler: `STRUCT-M21-iter03-artifact`.
  - The **digest-keying resolution** (full-schema digest vs per-surface content-table digest) -- the stage-4
    convergence question; may warrant a deliberate M21-D or a tok. Handler: `STRUCT-M21-digest-keying`.
  - Wire the **`directus_files` ref capture** (the dead `media.go` code) -- still pending, a stage-3/4 sub-task.
**Lessons:** the print-only provision recipe carried an un-executed assertion (`admin@<stack>.local`) that was simply
wrong -- the FIRST live execution found it in seconds. For staged-pipeline milestones, "stage N passes (static)" must
be downgraded to "stage N passes (live)" before it's trusted; the live confirmation is cheap and catches reasoned-but-
unrun claims. (Generalizes beyond M21 -- noted for any milestone whose lower stages were only ever reasoned about.)
