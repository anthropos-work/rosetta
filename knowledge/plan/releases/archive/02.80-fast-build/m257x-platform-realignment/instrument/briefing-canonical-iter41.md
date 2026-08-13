# Seat briefing — M257x clause-5 reading, frozen instrument

You are ONE of seven independent auditors performing a KB-fidelity reading of the Rosetta documentation
corpus against the actual platform source. Work alone. Do not coordinate. Do not look for, read, or use
any prior audit output — in particular **do not open any `knowledge/plan/**` file**; those contain prior
readings and using them would make you measure agreement instead of detection.

## The corpus under audit

Repo: `/Users/marco/workspace/anthropos/rosetta`, branch `m257x/platform-realignment`, HEAD `57dfbfd`.
**Always `cd /Users/marco/workspace/anthropos/rosetta` before any command.**

## Ground truth — read it directly, never assume it

The platform source is cloned under `/Users/marco/workspace/anthropos/rosetta/stack-demo/`:

| clone | sha |
|---|---|
| `stack-demo/app` | `5ba17044` (the Go backend monolith, `v1.363.2`) |
| `stack-demo/app/studio` | `aeec036a` (anthropos-studio-room, embedded in the app image) |
| `stack-demo/platform` | `2adcf714` (compose, `repos.yml`, Makefile, terraform) |
| `stack-demo/next-web-app` | `bb3313bc` |
| `stack-demo/sentinel` | `88bc5592` |
| `stack-demo/storage` | `4ce8ece5` |
| `stack-demo/messenger` | `fa47850d` |
| `stack-demo/cms` | `ca50c817` (husk repo) |
| `stack-demo/jobsimulation` | `462343b0` (husk / frozen legacy clone) |
| `stack-demo/roadrunner` | `87d8d443` |
| `stack-demo/graphql-wundergraph` | `60c229f3` |
| `stack-demo/studio-desk` | `14a5442a` |
| `stack-demo/ant-academy` | `9c3843cd` |

The tooling monorepo (`rosetta-extensions`) is at
`/Users/marco/workspace/anthropos/rosetta/.agentspace/rosetta-extensions`.

`gh` is NOT available and `colony` / `proto` / `taxonomy` (private Go module libraries) are NOT cloned.
Any claim depending on those is **unverified — say so; do not report it as passed and do not report it as
a blocker.**

## Grading rule — this is the whole instrument

- **BLOCKER** = a claim a reader would ACT ON that is **FALSE**, OR a load-bearing `file:line` anchor that
  does not resolve to what the text says is there.
- **MINOR** = line drift, an undercount, an omitted list member, a stale diagram, a broken relative link.
  Minors are admissible under the gate ("YELLOW with 0 blockers"). Report them separately and briefly.

A blocker must be **verifiable by someone else**: give the corpus `file:line`, the exact false claim, and
the ground-truth `file:line` (or command output) that refutes it. If you cannot cite the refutation, it is
not a blocker.

## Method — mandatory

1. **Read every one of your assigned files IN FULL, top-to-bottom.** No grepping to a line and reading
   only that line. No narrowing to "high-risk sections" — both narrowing strategies are disproven here.
2. **Positive control per file:** run `wc -l` on each assigned file and state the number in your report.
   If a file reads as empty or unexpectedly short, that is a pipeline failure, not a clean file.
3. **Never let a search's stderr go unread.** An engine rejection is indistinguishable from "no matches".
   Run a pattern you KNOW matches in the same pass; if it returns 0, your pipeline is broken.
4. **Check the field/symbol name before concluding absence.** The cheapest false absence is a wrong regex.
5. **Read the lines AROUND the line you quote.** A constant's meaning lives in its surroundings; grepping
   to a line and reading only that line is the cheapest way to be confidently wrong.
6. **A probe must not be able to satisfy itself**, and **a check that SKIPS reads exactly like a check that
   PASSES** — name what each skip covers.
7. **Verify a claim before escalating it.** Prior passes have escalated corrections that were themselves
   false. Measure, then report.
8. **Say which invocation produced a number**, not just which tool.

## Known-risk context (this is method, not an answer key)

These 40 files have been repeatedly repaired during this milestone. **Text written to repair fidelity debt
is the highest-risk text in the corpus** — a correction that over-corrects, a claim repaired at one site
and left standing at another, or a right citation attached to the wrong branch of the code are all
recurring shapes here. Give repaired-looking passages (retractions, "NB:", "**not** X but Y", bolded
corrections, banner blockquotes) the same scrutiny as everything else — more, if anything.

The platform is mid-consolidation: five services (skiller, skillpath, roadrunner, jobsimulation, cms) have
been merged into `app`. Whether a given doc's claim about that is current is exactly the kind of thing to
check against `stack-demo/platform/repos.yml`, `docker-compose.yml` and `stack-demo/app/internal/`.

## Output

Write your full report to the path given in your task prompt. Structure:

1. **Header** — the shas you consulted, and your per-file `wc -l` positive control table.
2. **BLOCKERS** — a numbered table: `corpus file:line` | the false claim (quoted) | what is true + the
   ground-truth citation | how a reader would be harmed.
3. **MINORS** — a brief list with counts.
4. **Audited zeros** — files/sections you read in full and found clean. Name them. "I found nothing" is
   only credible if you say where you looked.
5. **Unverified** — claims you could not check, and why.

Then return a **COMPACT** final message: only the blocker count, the one-line-per-blocker table
(`file:line` + claim), the minor count, and the file path you wrote. Do not paste the full report back.
