# iter-223 — decisions

## D-M257x-223-1 — sha drift is COUNTED and never a finding

`patch_anchor_guard` reports how many `pre_sha256` baselines no longer match, inside its verdict line,
and **never** changes its exit code on that count.

The obvious fence is the wrong fence here. Since M217 `demopatch.assert_pre_patch` returns `pristine`
when the whole-file sha differs but the anchor is intact exactly once — it WARNs and applies. The strict
sha gate is what silently refused two `app` perf patches, which is the defect M217 exists to have fixed.
A fence reddening on drift would therefore go RED on a set that works, contradict the shipped mechanism,
and get suppressed — and a suppressed fence is worse than no fence.

Pinned by a test that asserts the **non**-behaviour (`DriftIsNotAFinding`), because a deliberate
non-check is invisible unless something says so.

## D-M257x-223-2 — the CAUSE of the 10 drifted baselines is not adjudicated here

Measured: 10 of 23 baselines do not match the pristine file, and the drift set is **identical** at `HEAD`
and at `origin/main` — so they predate the platform's advance.

`demopatch-spec.md` §6 documents at least one as **by design**: the urls.ts chain, where
`next-web-public-website-url`'s `pre_sha256` **is** studio's `post_sha256`, so it reads DRIFTED against a
pristine file on purpose. Whether the other nine are chain members or simply un-repinned is a separate
question with a mechanical answer (`demopatch --repin` refuses unless the pre-image round-trips), and
answering it inside this iter would have been the third line of investigation.

Routed as `ROUTE-M257x-223-classify-the-ten-drifted-baselines`. **The guard reports the count and states
that it does not report the cause** — a fence that implied a verdict it never took would be the failure
this milestone has spent two hundred iters cataloguing.
