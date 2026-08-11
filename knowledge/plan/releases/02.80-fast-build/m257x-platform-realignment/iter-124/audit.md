# iter-124 — the triage's own accuracy, hand-measured

`triage.py` assigns three of the four fates by rule. **A rule set that is not audited is a promise.**
This is the audit: a seeded random sample of the C1 verdicts, hand-classified independently, then
compared. Drawn and classified before the split was published.

**Invocation** (reproducible — the seed is committed, not chosen after the fact):

```
/usr/bin/python3 iter-124/triage.py /tmp/m257x-iter124-tier2.json --c1-only --audit 30 --seed 124
```

**Denominator: 30 of 344 C1 verdicts.** Sample fate distribution: 30 `cite` (21 by `R3`, 9 by `R4`).
The sample drew no `fix`, `hedge` or `drop` verdict, which is expected at their measured rates
(4 + 7 + 2 of 344) and is a limitation stated rather than hidden: **this audit measures whether `cite`
is over-assigned. It cannot measure whether `hedge`/`drop`/`fix` are under-assigned except through
that same over-assignment.**

| # | site | rule | rule says | hand says | agree |
|---|---|---|---|---|---|
| 1 | `clerkenstein.md:151` | R3 | cite | cite — rext `alignment/cmd/multirun` | ✅ |
| 2 | `external_services.md:125` | R3 | cite | cite — `DIRECTUS_PUBLIC_BASE_ADDR` in compose/`.env_example` | ✅ |
| 3 | `academy-backend.md:50` | R3 | cite | cite — `app` ent schema dir | ✅ |
| 4 | `clerk-integration.md:37` | R4 | cite | cite — app's webhook handler enumerates the event types | ✅ |
| 5 | `ant-academy.md:32` | R3 | cite | cite — ant-academy repo | ✅ |
| 6 | `storage.md:59` | R3 | cite | cite — compose + storage source | ✅ |
| 7 | `clerkenstein.md:218` | R3 | cite | cite — rext `roster.go` | ✅ |
| 8 | `hiring.md:449` | R3 | cite | cite — rext seeder | ✅ |
| 9 | `ai-readiness.md:201` | R3 | cite | cite — `app/internal/data/ent/schema/` | ✅ |
| 10 | `coursebuilder.md:41` | R3 | cite | cite — app source | ✅ |
| 11 | `clerk-integration.md:29` | **R4** | cite | **drop** — *"Clerk's catalog is large; the platform uses a focused subset"* asserts nothing checkable; the bullets under it carry the content | ❌ |
| 12 | `roadrunner.md:119` | R4 | cite | cite — the Judge0 API surface; `judge0` is an org repo | ✅ |
| 13 | `cms.md:126` | R3 | cite | cite — app ent + the cms-in-app commit | ✅ |
| 14 | `storage.md:176` | R3 | cite | cite — storage terraform + `go.mod` | ✅ |
| 15 | `ai-labs.md:59` | R3 | cite | cite — `lab.go` | ✅ |
| 16 | `external_services.md:102` | **R4** | cite | **drop** — *"Full localtunnel setup with Clerk configuration"* is a link label in a see-also table, not an assertion | ❌ |
| 17 | `alignment_testing.md:395` | R3 | cite | cite — the v1.0/M3 milestone records | ✅ |
| 18 | `security_compliance.md:303` | R4 | cite | cite — `flag_use_azure_us` in app source | ✅ |
| 19 | `studio-desk.md:177` | R4 | cite | cite — `VITE_CLERK_PUBLISHABLE_KEY` is required by the app | ✅ |
| 20 | `ai-labs.md:51` | R3 | cite | cite | ✅ |
| 21 | `clerk-integration.md:153` | R3 | cite | cite — `next-web-app` package manifests | ✅ |
| 22 | `clerk-integration.md:144` | **R4** | cite | **hedge** — *"each is a thin adapter over Clerk's shared core"* is a claim about **Clerk's own package architecture**. The SDKs are `node_modules`, not clone content; nothing in any clone set carries it | ❌ |
| 23 | `ant-academy.md:346` | R3 | cite | cite — the `@anthropos.work` domain gate is in ant-academy source | ✅ |
| 24 | `sentinel.md:7` | R3 | cite | cite — sentinel imports no Clerk SDK | ✅ |
| 25 | `cms.md:224` | R3 | cite | cite | ✅ |
| 26 | `clerk-integration.md:132` | R4 | cite | cite — storage/messenger are frozen but clonable | ✅ |
| 27 | `backend.md:320` | R4 | cite | cite — app route registration | ✅ |
| 28 | `next-web-app.md:133` | R3 | cite | cite — `.env_example` | ✅ |
| 29 | `clerk-integration.md:113` | R3 | cite | cite — `colony/authn` claim reader | ✅ |
| 30 | `shared_libraries.md:109` | R3 | cite | cite | ✅ |

## Result — and the error is not spread, it is CONCENTRATED

**27 of 30 agree = 90.0 %.** All three disagreements run one way — the rules assign `cite` where the
hand assigns `drop` or `hedge` — which is the direction `R4` was **declared generous in before the
sample was drawn**. Not one disagreement runs the other way, and none touches `fix`.

**And the error sits entirely in one rule:**

| rule | sampled | correct | accuracy |
|---|---|---|---|
| **R3** — *names a reachable artifact* | 21 | 21 | **100 %** |
| **R4** — *default, presumed citable* | 9 | 6 | **66.7 %** |

**Read that as the finding it is: when a sentence names an artifact, "the evidence exists" holds
without exception in this sample. When it names none, the default is wrong a third of the time.**

## The corrected split, with the correction shown rather than folded in

Applying R4's measured 33.3 % mis-assignment to its 100 C1 members (all of it `cite` → `drop`/`hedge`):

| fate | as printed | corrected estimate | note |
|---|---|---|---|
| **cite** | 331 (96.2 %) | **≈ 298 (86.6 %)** | the R3 sub-population is un-corrected — it audited clean |
| **hedge** + **drop** | 9 (2.6 %) | **≈ 42 (12.2 %)** | the whole correction lands here |
| **fix** | 4 (1.2 %) | **4, and it is a FLOOR** | see below |

**`fix` is a floor and must never be quoted as a rate.** Falsity is not syntactic; the only instrument
that finds it is a reading, and this milestone measured its own reading's test–retest recall at
**~35 %** (iter-119, `FIX-M257x-iter119-instrument-recall-is-35pct`). Recall-correcting 4 gives ≈ 11
of 344 ≈ **3.2 %** — still far under the pre-registered 15 % branch, but that is an *estimate about an
un-fired branch*, not a measurement that the branch is safe. **Both statements are reported; neither
is upgraded.**

## Both pre-registered branches, checked against the numbers

- **`fix ≥ 15 % of C1` → the corpus is unfounded.** Measured **1.2 %**, recall-corrected **≈ 3.2 %**.
  **Does not fire.**
- **`cite < 50 % of C1` → the framing is wrong.** Measured **96.2 %**, corrected **≈ 86.6 %**.
  **Does not fire.**

**So the answer to the question nobody had measured: the corpus is UNDER-CITED, not unfounded.**
