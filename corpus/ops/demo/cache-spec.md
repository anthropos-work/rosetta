# The prompt-hash batch cache (v1.10 "method acting" M45)

> The reproducibility + cost-control half of the [AI generation engine](ai-generation-spec.md). Generation
> is expensive (a real LLM call per member) and non-deterministic; the cache makes a batch **reproducible**
> (an unchanged descriptor reseeds **byte-identical**) and **free on rerun** (a cache hit is **$0**, no LLM
> call). It is keyed by the **MOTHER prompt** so any change to the prompt (and thus the intended content)
> invalidates exactly the affected members.

---

## 1. Layout

```
.agentspace/.batchcache/
  batch-${hash}/
    member-0.json
    member-1.json
    ...
    member-${N-1}.json
    .lock                # the concurrency fence (held during a generation run)
    manifest.json        # batch descriptor digest + capture-version + member count
```

- `.agentspace/.batchcache/` is **per-box** and **git-ignored** (it holds generated content + cost
  metadata, never secrets). A shared/committed store is an Open question (flagged, not done in M45).
- `${hash}` is the **batch hash** (see §2). All members of one batch live under one `batch-${hash}/` dir.
- `member-${i}.json` is one generated **envelope** (`{name, email_local, bio, education, prior_roles[],
  current_role, skills_claimed[], skills_verified[], self_eval_bias, location}`) — AI-owned content only,
  no node-ids, no DB rows, **no secret material**.

---

## 2. The cache key — the MOTHER prompt + the capture version

The cache key is **NOT** the high-level descriptor; it is the **MOTHER prompt** — the fully-expanded,
deterministic per-member prompt string `EffectiveBatches()` renders (the Go-template output). Keying on
the mother prompt means:

- two descriptors that expand to the **same** prompts share cache entries (good);
- any change that alters a member's prompt (a different role mix, a reworded template, a new
  reserved-hero-name list) changes that member's hash → a **fresh** generation for exactly the affected
  members, the rest still hit.

**Capture-version extension.** The key is **extended with the taxonomy capture version** (the
snapshot manifest version from `stack-snapshot`). The generated **skill/role names** are only meaningful
against the public taxonomy they'll be resolved into; when the taxonomy is **re-replayed** at a new
capture version, the cache is **invalidated** (coarse granularity — the whole batch re-generates) so a
stale generation can't silently resolve against a different catalog. This is the documented
cache-invalidation granularity (the overview's Open question, resolved **coarse**).

```
batchHash = sha256( joinSorted(memberMotherPrompts) || "\x00" || taxonomyCaptureVersion )
memberKey = sha256( memberMotherPrompt || "\x00" || taxonomyCaptureVersion )   # per-member file selection
```

(The exact composition — batch-level vs per-member hashing — is a code detail; the **invariant** is:
prompt-content + capture-version fully determine the key, and nothing else does. Identical inputs →
identical key → identical bytes.)

---

## 3. Write discipline (atomic + fenced)

- **Atomic writes:** each `member-${i}.json` is written to `member-${i}.json.tmp` then **renamed** into
  place (`rename(2)` is atomic on a local FS). A crashed/aborted run never leaves a half-written member
  file that a later run would read as complete.
- **The `.lock` fence:** a generation run acquires `batch-${hash}/.lock` before writing. A second
  concurrent `cmd/gen-batch` against the same batch sees the lock and refuses (or waits), so two runs
  can't interleave writes into the same batch dir. The lock is released on clean exit.
- **Cache hit = $0:** before a member is generated, `cmd/gen-batch` checks for a present, valid
  `member-${i}.json`. A hit short-circuits the LLM call entirely — **no token spend, no cost**. A batch
  whose members are all cached costs **$0** and makes **zero** API calls.

---

## 4. Reproducibility (the gate's `$0` re-seed)

The exit gate's reproducibility clause — *"an unchanged batch descriptor re-seeds byte-identical from
cache at $0"* — is proven by:

1. Run `cmd/gen-batch` on a descriptor → populates `batch-${hash}/`, reports cost > $0.
2. Run the **same** descriptor again → **every** member is a cache hit → **0 API calls**, **$0** cost.
3. The `GeneratedBatchSeeder` consumes the cache **deterministically** (the resolvers + helpers are pure
   functions of the cached content + the replayed taxonomy) → the resulting DB rows are **byte-identical**
   across the two runs.

Because the seeder is a deterministic transform of `cache → rows`, byte-identical cache ⇒ byte-identical
seed. The cache is the reproducibility anchor.

**Org-scale caveat — the `$0` reseed stays DISTINCT, not necessarily byte-identical to the cache (v1.10 M46
iter-07).** A cache captured at org scale can carry the model's raw attractor duplicates (gpt-4o-mini re-picks
a handful of names hundreds of times across a ~600-member org; the same `email_local` recurs). The
`GeneratedBatchSeeder` applies a **deterministic, `$0` seed-time distinctness backstop** over **two axes** —
**name** (`seeders.DisambiguateGeneratedName` — keep the first name, swap a global-index-keyed surname) and
**email** (`UNIQUE(email)` is enforced by `public.users`; a colliding local part gets the global index
appended). Both are pure functions of `(cached content, global index)`, so the reseed is still **fully
reproducible** (same cache → same distinct identities → a FRESH `/demo-up` reproduces the same org) — it just
isn't *byte-identical to the cache* for the members whose cached name/email collided. The distinctness is what
makes the `$0` cache-hit reseed believable rather than a wall of duplicate names. Proven on the real
614-member cache: **614/614 distinct names + 614/614 distinct emails** seeded at `$0`. See
[`ai-generation-spec.md` §4g](ai-generation-spec.md).

---

## 5. Invalidation summary

| Change | Effect |
|---|---|
| Reword the prompt template / change a role-mix / change reserved-hero-names | affected members' `memberKey` changes → those members regenerate; unaffected members still hit |
| Re-replay the taxonomy at a new capture version | the whole batch's key changes → full regeneration (coarse, deliberate — names must match the live catalog) |
| Nothing changes | full cache hit → $0, byte-identical reseed |
| A `member-${i}.json.tmp` left by a crash | ignored (only the renamed final file is read); regenerated |

---

## See also
- [`ai-generation-spec.md`](ai-generation-spec.md) — the engine + the gen-acceptance protocol this cache
  serves.
- [`snapshot-spec.md`](../snapshot-spec.md) — the taxonomy capture + its **capture version** the cache key
  is extended with.
- [`idempotency.md`](../idempotency.md) — the broader bring-up re-run safety contract the `$0` reseed
  slots into.
