# iter-20 — decisions

## D104 — G1 read HOST profiles as compose profiles, and the repair is a discriminator, not a waiver

`platform_predicate_guard` was the **last RED fence on the box** after iter-18, on one finding:
*"`docker-desktop-vm` is documented as a profile at 2 site(s) but no service declares it."*

**"profile" is not a compose word.** Since M255 this corpus also has **host** profiles — the
checked-in `stack-core/hostprofiles/*.json` describing a machine's cores, memory and disk, which the
build budget and every headroom clause are written against. G1's noun-phrase detector
(`_PROSE_PROFILE`) had three discriminators — negation, postfix negation, ref-pin — and **no domain
discriminator**, so every *"a `X` profile"* anywhere in the corpus was graded as a compose token.

The finding was therefore *accurate and useless*: it correctly described what
`--profile docker-desktop-vm` would do, about a command nobody would ever type.

**The repair is a third discriminator, deliberately symmetrical with the other two** — same block
window (`_pin_window`), same evidence standard (*the sentence itself says which domain it means*).
`_HOST_PROFILE_DOMAIN` requires the window to **name** the other domain: `host profile`,
`hostprofiles`, a `hostprofiles/x.json` path, or the `profile[...]` subscript form that only a
host-profile dict is written with.

**It deliberately does NOT exempt on the word "host".** `--public-host`, `host-gateway` and
`STACK_PUBLIC_HOST` all appear in genuine compose prose; exempting on proximity to "host" would have
disarmed the arm the fence exists for — *a capability probe that fails OPEN disarms the check it
guards*, which this release has already paid for once.

**RED-proven, three tests** (`test_platform_predicate_guard.py`, 190 passed):

1. the host-profile noun-phrase forms stop being graded;
2. **anti-widening** — three compose claims sitting one clause from the word "host" are **still**
   graded;
3. **mutant** — widening `_HOST_PROFILE_DOMAIN` to always-true launders a live `cms` claim, which
   proves the exemption is what decides these cases rather than an accident of the fixtures.

Shipped as `fast-build-m258-iter-20`, **verified on origin**, declaration re-pinned.

## D105 — one of the two sites needed the prose to name its domain, and that is an improvement

After the discriminator landed, `build-budget.md:551` was exempted and **`:197` was not**. The reason
is the window, not the rule: the phrase *"host profile"* sits at `:194`, and a **blank line at `:195`
is a block boundary** — `_pin_window` stops there, by design, because that boundary is what keeps a
pin from buying silence for its neighbour (`D-M257x-63-1`).

So the sentence at `:197` relied on context two paragraphs up to say which kind of profile it meant.
Naming it inline — *"a `docker-desktop-vm` **host** profile's `cores`"* — makes the prose say what it
means at the point it means it.

⚠️ **This is not the forbidden move.** The rule is *never edit correct prose to make a wrong fence
green*. Here the fence had just been made **right**, and the prose was genuinely ambiguous at that
spot for a human reader too. Widening the window instead would have weakened a discriminator that
exists to be narrow. Recorded explicitly because the two cases look identical from the exit code.

## D106 — the `app` row's anchors had all drifted, and NO fence could see it

`ROUTE-M258-iter18-app-row-anchors-are-at-2035f9a`, discharged. The migration map's `app` row pinned
seven `app/main.go` wiring anchors at `2035f9a`; `origin/main` is `c52dbc51e`. **Every one had
drifted**, by 12–20 lines:

| domain | was | is | construct |
|---|---|---|---|
| customerio-sync | 395 | **396** | `customeriosync.New` |
| storage | 524 | **537** | `internalstorage.NewManager` |
| skiller | 690 | **706** | `skiller.NewSkillerManager` |
| jobsimulation | 721 | **734** | `jobsimwiring.Wire` |
| skillpath | 751 | **764** | `skillpath.NewSessionManager` |
| cms | 1153 | **1167** | `appcms.Wire` |
| messenger | 1471 | **1458** | `msgadapters.Wire` |
| **sentinel** (net-new, 8th) | — | **305** | `sentinel.Open` |

**8/8 verified by reading the target line**, not by trusting the search.

**The finding is that nothing caught this.** They are `app/main.go:NNN` citations, which assertion F
grades **range-only** — there is no block structure in a Go file for it to attribute a line to — and a
1,635-line file swallows a 20-line slip without ever going out of range. The one anchor in this same
table that *was* graded (`messenger`'s `:1485`, cited in a compose-adjacent form) had already slid
onto a **closing brace**, which is how iter-18 noticed the class at all.

*Anchors into a file with no block structure are the class the fences cannot see. Re-derive them at
every ref bump, or do not pin them.*

## D107 — eleven fences, all green, and that is the honest scope of the claim

At close: `platform_alignment_guard` · `platform_predicate_guard` · `anchor_construct_guard` ·
`demo_knob_guard` · `decommissioned_instruction_guard` · `corpus_citation_guard` ·
`corpus_index_guard` · `markdown_structure_guard` · `fence_command_guard` ·
`evidence_visibility_guard` · `dev_flag_guard` — **all `rc=0`**.

Two of the five test modules run this iter still fail, and **both were verified against iter-18's
measured pristine-HEAD baseline as pre-existing** (`test_fence_provenance::test_the_escape_accepts_and_records`,
`test_guard_family_verdict_line_m257x::test_every_member_that_ran_reported_on_its_OWN_summary`) —
members of the `FIX-M258-iter03-guard-scans-its-own-scratch` family, which fires on any box that has
ever run a demo. `test_platform_predicate_guard`, `test_predicate_enumerator` and
`test_service_doc_status_fence` are green.

**What "green" covers, stated rather than implied:** eleven fences whose subject is the corpus and its
agreement with the platform. It does **not** cover the ~46 pre-existing rext-internal census failures,
which are a separate, measured, still-open population.
