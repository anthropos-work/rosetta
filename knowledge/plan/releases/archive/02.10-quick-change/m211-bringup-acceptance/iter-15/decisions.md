# iter-15 — Decisions

**D1 — The demopatch drift is a FILE-level sha re-anchor, not a hunk change.** The skiller-in-app merge
bumped next-web-app v2.89.0 → v2.106.1; sibling exports in `packages/core-js/src/constants/urls.ts` drifted
so the file-level `pre_sha256` G2 gate no longer matched. The STUDIO_URL + PUBLIC_WEBSITE_URL anchor hunks
are byte-identical to v1.10 (grep-confirmed). So the fix is a pure re-anchor (recompute the 4 chained hashes),
the same maintenance M49 #8 did (v2.89.0 re-anchor). Computed via the demopatch's OWN `manifest_loader` +
`body.replace(anchor, replacement, 1)` + `sha256_text` — the exact apply code path — so the apply's
post-condition assert (`sha256_text(patched) == post_sha256`) is guaranteed to pass. Verified live: the
demopatches applied cleanly during the rebuild (no G2 refuse), the clone reverted to pristine (trap).

**D2 — The BUILD reads the CONSUMPTION clone, so the re-pin must land in both places.** The next-web rebuild
runs `stack-demo/rosetta-extensions/demo-stack/up-injected.sh` (the consumption clone at tag
`quick-change-m211`), NOT the authoring copy `.agentspace/rosetta-extensions`. An authoring-only edit is
invisible to the build. So: commit the re-pin to the authoring copy (`84e15e9`), move the local tag
`quick-change-m211` to it, then local-fetch the moved tag into the consumption clone (`git fetch <authoring>
+refs/tags/…` + `checkout -f`). Both clones now carry the new hashes. This is the frontend-repo analog of
iter-10's build-scratch re-sync for the Go services.

**D3 — Rebuild + recreate, never touch postgres.** `docker image rm demo-1-next-web` to force a rebuild (the
reuse guard keys on the offset endpoint, which the URL patches don't change), `build_frontend_next_web` from
the consumption clone (applies the patches, bakes the demo-local URLs into the JS bundle via the .env.local
overlay, trap-reverts the clone), then `docker compose … up -d --no-deps --force-recreate next-web-app` — the
`--no-deps` is mandatory (the M46 warning: a `--force-recreate` without it recreates postgresql and wipes the
seeded org).
