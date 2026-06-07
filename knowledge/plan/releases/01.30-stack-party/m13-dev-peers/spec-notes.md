# M13 — Spec notes

Technical notes accumulate here during build.

## Local per-stack Directus on dev
_Reuse M10 `stack-snapshot/directus/provision.go`; repoint dev CMS at the per-stack Directus._

## Auto-snapshot on dev build
_`stacksnap replay` taxonomy + directus, cache-first; `--no-snapshot` escape._

## dev-min seed preset
_~1 org + ~10 users + minimal activity; n=0 reset guard preserved._
