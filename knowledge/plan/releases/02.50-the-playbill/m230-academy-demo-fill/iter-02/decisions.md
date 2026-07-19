# iter-02 — decisions (intra-iter)

## D1 (side-observation): `test_demopatch.py` 2 failures are pre-existing next-web clone drift — NOT my regression

`test_pubweb_manifest_chains_on_studio_against_live_clone` + `test_real_manifest_hashes_reproduce_against_live_clone`
fail because the LIVE `stack-demo/next-web-app` `urls.ts` no longer matches the `next-web-public-website-url` /
`next-web-studio-url` manifests' pinned `pre_sha256` (`live != pinned pre` for both). My iter-02 work touched
ONLY ant-academy (a different repo + different manifests) — confirmed by the authoring-copy git status
(`ant-academy.sh` + `academy-fs-published-fallback/` + `apply-academy-fs-published.sh` + the academy test).
Routed forward (a next-web manifest re-anchor; a **cold `/demo-up`** prerequisite, not an academy-patch issue).

## Note
The milestone-level strategy + the USER-BLOCKER surface (the formal cold-`/demo-up` gate) are recorded at the
milestone level in `../decisions.md`.
