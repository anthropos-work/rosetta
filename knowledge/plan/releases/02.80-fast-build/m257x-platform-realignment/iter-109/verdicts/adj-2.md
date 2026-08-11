# Adjudicator 2 — seats B and C, readings #27 and #28

**Trees I read.** All grading against `stack-demo/*` at the brief's table shas — platform `0c91421d`,
app `ad9f3c49`, next-web-app `8297c684` (and `f97ba659` as a second witness), cms `ca50c817`,
messenger `fa47850d`, jobsimulation `462343b0`, storage `4ce8ece5`. **No `rosetta-extensions` claim
appears in any of the nine bookings**, so the pinned-vs-authoring split never bites this seat-group and
no `wrong-tree` label is used. No fetch, no git state change, read-only.

**Method.** Every anchor below was opened by me, at the ref the claim names (or, where none is named, at
the checkout **and** at `origin/main`), and read with surrounding context. Where a booking rests on a set
cardinality I re-derived the set first and state it before any arithmetic.

---

## Verdicts

### seat B (r27) B1 | `corpus/services/ai-readiness.md:49-52` | UPHELD | IN-SCOPE | PREDICATE: The `⚠⚠ M51 iter-08/09` block that carries the re-anchored call site is at ai-readiness.md:496.

   evidence: I read `corpus/services/ai-readiness.md` at corpus HEAD. `:52` asserts *"The target is
   therefore named by construct above and pinned at **`:496`** (M257x iter-102)"*. The named target is
   the `⚠⚠ M51 iter-08/09` block, identified at `:46-48` by the quoted parenthetical *"(now
   `aireadiness/readiness.go`, formerly `workforce/ai_readiness.go:512`)"*. Measured: `:496` is
   `> relying on either number; do not re-derive them from prose.` — the **closing line of the
   `> **✅ CORRECTED M219 …**` blockquote**, which opens at `:476`. The M51 block **opens at `:498`**;
   the quoted parenthetical is at **`:500`**. So the third-generation repair lands inside the very
   blockquote the same paragraph (`:50-51`) names as the *wrong* target for the previous generation.
   Not `historical-anchor`: `:496` is asserted as the live, current pin, in a passage published as a
   worked example of a repaired anchor. The self-heal clause (*"when the two disagree, the name wins"*)
   lowers reader cost but does not make the pin true.

### seat B (r28) B1 | `corpus/services/ai-readiness.md:52` | UPHELD | IN-SCOPE | PREDICATE: The `⚠⚠ M51 iter-08/09` block that carries the re-anchored call site is at ai-readiness.md:496.

   evidence: same re-derivation as above — `:496` closes the `✅ CORRECTED M219` blockquote (`:476-496`);
   the M51 block opens at `:498`; the quoted parenthetical is at `:500`. Same anchor, same predicate as
   r27-B1; collapses onto **P1**.

### seat B (r27) B2 | `corpus/services/ai-readiness.md:84`, `:334-335`, `:482-484` | UPHELD | IN-SCOPE | PREDICATE: AIReadinessClient.tsx's orgEnabled gate, effectiveCycleId, isFetched gate and tab-filter read sit at `:133-134`/`:153-154`/`:166-170`/`:599`.

   evidence: I opened
   `next-web-app apps/web/src/app/(authenticated)/(verified)/ai-readiness/AIReadinessClient.tsx` at
   **both** `8297c684` (the checkout) and `f97ba659` (origin/main) — byte-identical at every offset, so
   the finding is ref-independent and no citation-resolution ladder rescues it.
   - `:133-134` = the tail of the `// Feature gate:` comment (*"`featureOn` also keeps the dashboard
     queries from firing while / off (the backend would reject them with ErrAIReadinessDisabled)."*).
     `const { orgEnabled } = useAiReadinessEnabled(true);` is at **`:135`**.
   - `:153-154` = comment (*"// The closed-snapshot path now ALSO joins current membership tags so the /
     // byTeam aggregate populates …"*). `const effectiveCycleId =` is at **`:155-156`**.
   - `:166-170` = `:166-167` comment + `:168-170` `const dataQ = useAIReadiness({` / `cycleId:` /
     `includePeople:`. The cited gate `enabled: featureOn && cyclesQ.isFetched,` is at **`:171`** — the
     range stops one line short of the construct it is offered as evidence for.
   - `:599` = `tab.key === 'how' ||`. The `SHOW_SECONDARY_TABS` read is at **`:601`**.
   The one anchor above the drift point, `:78` `const SHOW_SECONDARY_TABS: boolean = false;`, is exact —
   which is what lets the block read as verified. Not ref-discipline: these four citations name no ref,
   and the document's only `8297c684` re-derivation (`:307-309`) **scopes itself in terms** to *"the
   three bullets below"*, which I confirmed are correct and which do not include any of these four.

### seat B (r28) B2 | `corpus/services/ai-readiness.md:84`, `:335`, `:482`, `:484` | UPHELD | IN-SCOPE | PREDICATE: AIReadinessClient.tsx's orgEnabled gate, effectiveCycleId, isFetched gate and tab-filter read sit at `:133-134`/`:153-154`/`:166-170`/`:599`.

   evidence: identical re-derivation to r27-B2, on the same four constructs at the same two refs. Same
   four anchors, one drift, one cause; collapses onto **P2**. Both seats booked it as one block with an
   explicit anchor count, and I keep that shape so a repair that fixes one anchor does not read as
   discharging the predicate.

### seat B (r27) B3 | `corpus/services/cms.md:287` | UPHELD | IN-SCOPE | PREDICATE: The Go service invokes `studio/gen.py` and `studio/postgen.py` through `bash -c`.

   evidence: `cms.md:287` says the Go service *"invokes `python3 studio/gen.py ...` / `studio/postgen.py`
   … **via `bash -c`**"*, present tense, with **no HISTORICAL marker** — unlike both of its neighbours in
   the same section (`:233` and `:291`, which are marked). The live code is `app` `ad9f3c49`
   `internal/cms/studio/studioManager.go`, and it is the opposite:
   - `:1096-1098` — *"runCommand executes name+args in argv (exec) form — **NEVER through a shell**…
     nothing is string-interpolated into a command line … (M809b H-1/M-1)"*; `:1101`
     `pycmd := exec.CommandContext(ctx, name, args...)`.
   - `:100-103` — *"It MUST NOT be interpolated into a shell … **No `bash -c`**."*; `:119`
     `s.runCommand(ctx, pyBin, append([]string{"studio/gen.py"}, tokens...))`.
   - `:122-131` — the dev-mode venv bootstrap runs `python3 -m venv` and `pip3 install -r
     studio/requirements.txt` as fixed argv, *"previously chained into the same `bash -c` string"*.
   I ran the absence check three ways per rule 3: `git grep -n '"bash"' ad9f3c49 -- '*.go'` over the whole
   `app` tree returns **0**, and `git grep -n 'exec\.Command' ad9f3c49 -- internal/cms/` returns the single
   argv-form call at `:1101`. The claim is still true of the **frozen** `cms` repo
   (`ca50c817:internal/studio/studioManager.go:967` = `exec.Command("bash", "-c", command)`) — which is
   what makes it a defect rather than mere staleness: it is correct about the dead code, wrong about the
   shipped code, and the property it inverts is a deliberate security fix. Context that settles the
   scope: the preceding paragraph (`:265-268`) re-points Python work to *"in `app/studio/`, not
   `cms/studio/`"*, the code block at `:271` opens `cd app/studio`, `:110` states *"the live code is
   `app/internal/cms/`"*, and the note's own closing sentence is live advice (*"use a venv to match the
   service's behavior"*).

### seat B (r28) B3 | `corpus/services/ai-readiness.md:270` | UPHELD | IN-SCOPE | PREDICATE: The email-override validator's placeholder package is `messenger/pkg/aireadinessemail`, not `app/internal/messenger/aireadinessemail`.

   evidence: `:270` says overrides are *"validated against `messenger/pkg/aireadinessemail` placeholders"*.
   The validator is `app/internal/aireadiness/emailoverride/emailoverride.go`, and at `:29` it imports
   `aireadinessemail "github.com/anthropos-work/app/internal/messenger/aireadinessemail"` — **no `pkg/`
   segment**; `git ls-tree ad9f3c49 internal/messenger/aireadinessemail/` lists 9 files there. The
   citation is inside a bullet list whose every sibling path is `app/internal/`-relative
   (`aireadiness/defaults.go`, `aireadiness/provision.go`, `aireadiness/notifications/`,
   `aireadiness/emailoverride/`, `aireadiness/compare.go`, `aireadiness/recommendation_engine.go`), under
   a heading that scopes the whole list to the **live** app package refactor (`:252`). Under the list's
   own convention the cited path resolves to nothing. `messenger/pkg/aireadinessemail` does exist — but
   only in the standalone `messenger` repo at `fa47850d` (`pkg/aireadinessemail/{format_days,override,
   override_test,renderer}.go`), a repo `838d907` removed from `repos.yml`, so `make init` does not clone
   it. The same doc two lines above (`:268`) has just told the reader messenger is *"in-process inside
   `backend` since the v9.0 fold … not a separate service"*, which is what makes this internally
   inconsistent rather than merely terse. Noted against upholding, and recorded: the substance
   (validation against the aireadinessemail placeholder set) is true, the two copies differ only by
   gofmt, and the platform's own stale comment at `emailoverride.go:33` uses the same `pkg/` shorthand —
   so this is the weakest of the six. It is `grep`-unique in the corpus (1 occurrence tree-wide), so
   there is no second anchor.

### seat C (r27) B1 | `corpus/services/jobsimulation.md:160` | UPHELD | IN-SCOPE | PREDICATE: The Directus voice-engine nil-default lives at `cms/directus/collections/jobsimulation.go:1594-1597`.

   evidence: `git -C stack-demo/cms ls-tree --name-only ca50c817` returns
   `.agentspace .github .gitignore CHANGELOG.md CLAUDE.md Dockerfile Dockerfile.dev Makefile atlas.hcl
   cmd cog.toml ent go.mod go.sum internal knowledge main.go singularity-manifest.md terraform tools.go`
   — there is no `directus/` at the cms repo root, and `git ls-files | grep -c '^directus/'` = **0**. The
   two real files are `cms/internal/directus/collections/jobsimulation.go` (frozen) and
   `app/internal/cms/directus/collections/jobsimulation.go` (live, `ad9f3c49`); in the latter I read
   `:1594` `func voiceEngineFromDirectus(...)`, `:1595-1596`
   `if directusVoiceEngine == nil { return simulation.SimulationVoiceEngineGptrealtime }`. So the fact
   and the line range are right and the path is unresolvable. Corpus-wide, `grep -rn
   'directus/collections/jobsimulation.go' corpus/` returns **9** sites; the other **8** all spell it
   correctly (`ai_architecture.md:56,211,240,253`, `external_services.md:566,624`,
   `snapshot-spec.md:565,665`) — this one line is the sole outlier, and it truncates toward the
   decommissioned repo, which `external_services.md:318` explicitly warns against
   (*"`app/internal/cms/directus/` (NOT the frozen cms repo's `internal/directus/`)"*).

### seat C (r28) B1 | `corpus/services/jobsimulation.md:160` | UPHELD | IN-SCOPE | PREDICATE: The Directus voice-engine nil-default lives at `cms/directus/collections/jobsimulation.go:1594-1597`.

   evidence: same re-derivation, same anchor. Collapses onto **P5**.

### seat C (r27) B2 | `corpus/services/customerio-sync.md:163-164` (restated `:85-86`, `:64`) | UPHELD | IN-SCOPE | PREDICATE: customerio-sync was the only platform compose service ever built from a git-URL context.

   evidence: I re-derived the **set** first, independently, over the whole history of
   `stack-demo/platform` — for every commit touching `docker-compose.yml`, every `context:` line matching
   a git URL, de-duplicated. **Cardinality: 18 distinct repo URLs** (app, chronos, cms, customerio-sync,
   graphql, graphql-wundergraph, graphqltmp, intelligence, jobsimulation, messenger, realtime, roadrunner,
   sentinel, simulator, skiller, skillpath, storage, studio-desk) — i.e. **17 services other than
   customerio-sync** were built exactly that way. Building from
   `git@github.com:anthropos-work/<repo>.git#main` was the platform **default** until `a2a3ee6`
   (2026-02-27, *"add Makefile, repos.yml, and switch to local Dockerfile.dev builds"*), and it was not
   unique even after that: at `a2a3ee6` **two** services keep a git URL —
   `:393 customerio-sync.git#main` and `:458 realtime.git#main` — with `realtime` holding it until
   `c17cc9a` (2026-04-15). Only from that date forward was customerio-sync the sole one; at `838d907^` it
   is indeed the only git-URL context among seven. Not ref-discipline: rule 2's class is a *past-tense
   claim refuted by newer evidence*; this is a past-tense claim refuted by **older** evidence, which the
   rule does not cover. What decides it against the contemporaneous reading is `:64`, which frames the
   claim historically and points the reader at the record: *"the build pattern was unique in the platform
   and a reader will meet it in older runbooks and **in `git log`**"* — in `git log` it is manifestly not
   unique. Recorded against: the sentence's operative advice (*"do not reach for it as a precedent"*) and
   `:86`'s qualified twin (*"every **remaining** compose build takes a local sibling directory"*) are both
   correct, so the practical payload survives; it is the unqualified superlative, asserted three times,
   that does not.

---

## PREDICATE ROLL-UP

```
P1 | The ⚠⚠ M51 iter-08/09 block carrying the re-anchored call site is at ai-readiness.md:496 (it opens at :498; the quoted parenthetical is at :500, and :496 closes the ✅ CORRECTED M219 blockquote the same paragraph names as the WRONG target) | anchors: seat B r27 B1 @ corpus/services/ai-readiness.md:49-52, seat B r28 B1 @ corpus/services/ai-readiness.md:52
P2 | AIReadinessClient.tsx's orgEnabled gate, effectiveCycleId, isFetched gate and tab-filter read sit at :133-134/:153-154/:166-170/:599 (they are at :135/:155-156/:171/:601 at both 8297c684 and f97ba659) | anchors: seat B r27 B2 @ corpus/services/ai-readiness.md:84, :334-335, :482-484; seat B r28 B2 @ corpus/services/ai-readiness.md:84, :335, :482, :484
P3 | The Go service invokes studio/gen.py and studio/postgen.py through `bash -c` (the shipped code runs argv form and says "NEVER through a shell") | anchors: seat B r27 B3 @ corpus/services/cms.md:287
P4 | The email-override validator's placeholder package is messenger/pkg/aireadinessemail, not app/internal/messenger/aireadinessemail | anchors: seat B r28 B3 @ corpus/services/ai-readiness.md:270
P5 | The Directus voice-engine nil-default lives at cms/directus/collections/jobsimulation.go:1594-1597 (no such path in any clone) | anchors: seat C r27 B1 @ corpus/services/jobsimulation.md:160, seat C r28 B1 @ corpus/services/jobsimulation.md:160
P6 | customerio-sync was the only platform compose service ever built from a git-URL context (18 distinct repo URLs did) | anchors: seat C r27 B2 @ corpus/services/customerio-sync.md:163-164 (restated :85-86, :64)
```

**Cross-seat collapse:** P1 and P2 each collapse the two seat-B readings (r27+r28) onto one predicate;
P5 collapses the two seat-C readings. P3, P4 and P6 are single-reading predicates. No two anchors were
collapsed on resemblance alone — every collapse is the same construct at the same corpus line.

**Distribution:** 3 predicates in `ai-readiness.md` (P1, P2, P4), 1 in `cms.md` (P3), 1 in
`jobsimulation.md` (P5), 1 in `customerio-sync.md` (P6). All six are in `corpus/services/**`; none in
`corpus/architecture/**`. **Zero blockers were booked against `platform-migration-status.md`,
`alignment_testing.md`, `askengine.md`, `intelligence.md`, `academy-backend.md`, `chronos.md`,
`services/README.md` or `TEMPLATE.md`** by either seat.

**Shape of the six:** five are *citation-resolution* defects (a line number or a path that does not name
the construct it claims) and one (P3) is a *substantive* inversion of a shipped security property. Four
of the six sit in passages that are themselves corrections or worked examples of anti-rot discipline —
P1 is the third generation of one anchor, P2 is the block a same-file re-derivation swept past.

**Summary line:**
`BOOKED=9 UPHELD=9 REJECTED=0 IN-SCOPE-UPHELD-BLOCKERS=9 DISTINCT-IN-SCOPE-PREDICATES=6`
