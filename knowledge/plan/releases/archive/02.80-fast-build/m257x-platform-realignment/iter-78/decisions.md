# iter-78 — decisions

## `D-M257x-78-1` — 8, 9 and 10 were never three opinions about one number. Two are real sets and one is a count of nothing.

`CHECK-M257x-iter76-compose-service-count` was the one routed item in this milestone left
**explicitly unsettled**: *"8 vs 9 vs 10; my one-line grep disagreed with the tested parser and the
disagreement is recorded rather than resolved by assertion."* Both readings were right, and neither
was complete.

| number | what it counts |
|---|---|
| **8** | services declared by `docker-compose.yml` **alone** |
| **10** | the **effective** topology, once `include: common.yml` adds the always-on floor |
| **9** | nothing, at any ref |

Derived across every ref the corpus cites, `docker-compose.yml`'s own count runs
**12 → 12 → 11 → 8 → 8** (`a2a3ee6` … `236771f` … `2adcf71` … `d11a403` … `0dab54d`), effective
**14 → 14 → 13 → 10 → 10**. There is no ref at which nine services existed. Three documents said it.

**The consequence for the fence is the interesting half:** because 8 and 10 are *both* real and the
corpus legitimately contains one document of each — `architecture_overview.md` says 8,
`dependency_map.md` says 10 — **G10 asserts the pair, never a value.** A fence that picked a side
would have gone RED on a correct document, and the "wrong" one would have been repaired into a
different truth. The finding says *"N is neither"* and names both, because that is the only thing
that is decidable.

Also settled while adjudicating: `platform-alignment.md:194` states **14 compose services** and is
**correct** — it is a dated measurement (*"At 2026-07-31 these read: …"*), and at `b56d731`
(2026-07-31) the effective topology was 12 + 2 = 14. A number that looks stale beside a number that
is current is not a defect when it carries its date.

---

## `D-M257x-78-2` — the broad construct was measured at 44% and replaced, not thresholded.

The obvious rule — *"a number followed by `services`, in a compose context"* — reaches **14 live
sites** and would fire on **9**, of which **4** are true. **44% precision.**

The reason is not sloppiness in the corpus; it is that this corpus counts several genuinely
different sets with the same two words:

| site | what it counts | verdict |
|---|---|---|
| `service_taxonomy.md:57`, `:427` | the **3-service floor** | correct |
| `external_services.md:171` | the **7 application services** (10 − floor) | correct |
| `dependency_map.md:31` | *"the last two **subgraph** services"* | correct |
| `platform-alignment.md:1661` | narrative — *"the eleven services that were perfectly fine"* | correct |

No threshold fixes that. §4 Trap A says replace the **rule**: scoped to a **declaration verb**
(`declares`/`defines`) with a **compose subject in the same block**, the construct reaches **4 sites
at 100% precision** and misses nothing of the class — every *"compose declares N services"* claim in
the tree.

**Two sub-decisions, both measured:**

- **Number WORDS are in the pattern.** All three false sites write *"declares **nine** services"*. A
  digits-only pattern reads **none** of the live defects. A bounded, explicit `zero…twenty` map is
  not English parsing; it is the corpus's own spelling of an integer.
- **The window is the BLOCK, not the line.** Line-scoped, the construct reached **2 of 4** (100%
  precision, 50% recall) because two of the live sites wrap their sentence across lines. Block-scoped
  it reaches all four with no precision lost. **This is the third time in this milestone that a
  "policy hole" has turned out to be a window bug** — `_pin_window` at iter-63, `_NEGATED` at
  iter-68, and now this. It is worth checking the window *first*, before concluding a class is
  unreachable.

---

## `D-M257x-78-3` — iter-77's cross-repo-pin route is confirmed live, and it is not hypothetical.

`external_services.md:296` claims, in words, *"platform **`0dab54d`**'s compose declares nine
services"* — a statement about the guard's own ref. Its block's **leftmost** pin is `b948604`, an
**app** sha, so `_pin_exempts` exempted the whole sentence and the claim about *now* went unchecked.

G10 uses the rule G9 established one iteration earlier — **resolve the ref in the platform clone; a
sha that does not resolve there cannot date a platform file** — and the claim is graded and RED.

This is the second live instance of the mechanism (`roadrunner.md:14` was the first), which promotes
`CHECK-M257x-iter77-cross-repo-pin` from *"unmeasured whether any of the 145 dates a platform
claim"* to **at least two do, and both were false**. The route stays open for the remaining 143;
what is now settled is that the class is real.

---

## Routed forward

- **`CHECK-M257x-iter77-cross-repo-pin`** — upgraded, not closed. 145 pin-exempted blocks name a sha
  that does not resolve in the platform clone; **two are now confirmed to date a platform claim and
  both were false**. The general fix (resolve-in-repo for every assertion whose subject is a
  platform file) is larger than either iter and is the natural next target.
- **`CHECK-M257x-iter78-running-vs-declared`** — `staging-bringup.md:370` and
  `staging_from_dump.md:323` counted **running containers under `--profile all`**, not declared
  services: a different predicate, repaired here by derivation (`|select(all)|` = **8**, and `all` no
  longer includes `messenger` or `storage`) but **not fenced**. G3 fences the *default* profile's
  count; no assertion fences a *named* profile's count stated in prose separated from its command by
  a code fence.
- Unchanged: `CHECK-M257x-iter77-narration-vs-documentation` · `CHECK-M257x-iter77-zsh-modifier` ·
  `CHECK-M257x-iter77-developer-dir` · `CHECK-M257x-iter76-seat-ref-discipline` ·
  `CHECK-M257x-iter70-studio-room-lines` · `RF-M257x-iter71-run-returns-a-tuple` ·
  `FIX-M257x-iter53-union-set` (**PENDING USER DECISION**) · `FIX-M257x-iter56-assignment-flake`
  (**NOT DECIDED**) · `CHECK-M257x-iter38-ai-act-classification` (owner outside this milestone) ·
  RF-2/3/7–13.

**`CHECK-M257x-iter76-compose-service-count` is CLOSED** — settled by derivation across refs, the
three false sites repaired, and the predicate fenced by G10 with its denominators printed on every
run.
