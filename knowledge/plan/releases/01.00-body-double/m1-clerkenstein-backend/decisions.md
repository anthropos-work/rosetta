# M1 — decisions

## TOK-01: mirror-by-score, easy-side-first — 2026-06-03

**Tok type:** bootstrap (iter-01)
**Initial strategy:** Build Clerkenstein incrementally in the `clerkenstein` repo (under gitignored
`anthropos-demo/`), measured every step by `/align-run` against a Clerk Alignment DNA. Sequence the
work **easy-offline-side first** so the score starts moving without infrastructure friction:
1. **Author the Clerk DNA** (`/align-dna`) from the platform-consumed surface (spec-notes § A/B),
   splitting it into the **authn** capability set (offline-capturable) and the **orgclient** capability
   set (live-Clerk-dependent).
2. **authn-provider twin first** — local JWT mint/verify with the platform's claim shape + one
   universal credential. Its goldens are offline-capturable (mint a session JWT with a test key against
   a local JWKS; the "source" is the Clerk SDK's verify/decode run locally). Drive the **critical**
   authn genes to 100% — this alone proves backend login works with zero live Clerk.
3. **orgclient twin second** — the networked CRUD methods. Author the mirror against the
   `clerk-sdk-go/v2` response types; capture goldens per the decision below.
4. **Injection** — `go.mod replace` + skip-worktree (resolve the colony-granularity sub-question in
   the first authn tik).
**Rationale:** the M0 framework makes "how faithful is the mirror?" a number; the fastest honest path
to the gate is to land the offline-capturable critical capabilities (authn) first — real score
movement with no blocker — then tackle the live-SaaS orgclient. This front-loads value (backend auth
works) and isolates the one genuine unknown (orgclient golden source).
**Strategy class:** new-direction
**Distance-to-gate context:** gate = 100% critical / ≥95% overall on the Clerk DNA; starting value 0%
(no mirror, no DNA yet). authn ≈ the critical core; orgclient ≈ the bulk of the gene count.
**Next-tik direction:** iter-02 = stand up `anthropos-demo/clerkenstein` + author the Clerk Alignment
DNA (authn + orgclient capabilities × variants) via `/align-dna`. **Blocked on the OPEN DECISION below**
before the first measurement can run.

### OPEN DECISION (D1) — orgclient golden source + workspace setup (blocks iter-02's first measurement)
The exit gate is an alignment score, which requires **goldens for the source side**. The authn side is
offline-capturable. The **orgclient side's source is the LIVE Clerk API** (`api.clerk.com/v1`), so its
goldens need one of:
- **(a) Live capture** — a Clerk dev-app secret key + network: `alignctl capture --source live` against
  real Clerk. Highest fidelity; needs credentials + network this environment may not have.
- **(b) Hand-authored goldens** — derive expected response shapes from the `clerk-sdk-go/v2` response
  structs (offline, deterministic, lower fidelity until reconciled against live once available).
- **(c) Hybrid** — hand-author now to make progress; reconcile against live capture when credentials
  land (the `dna diff` + re-capture path M1b will formalize).
Plus: confirm that **creating the `anthropos-demo/clerkenstein` workspace + cloning the pinned Clerk
SDK** is in-scope/available here (network for `git clone github.com/clerk/clerk-sdk-go`).
**Recommendation:** (c) hybrid — start offline (authn live-capturable locally; orgclient hand-authored
from SDK types) so tiks make real score progress immediately, and reconcile orgclient against live
Clerk when credentials are available. Needs user confirmation before iter-02 runs.
