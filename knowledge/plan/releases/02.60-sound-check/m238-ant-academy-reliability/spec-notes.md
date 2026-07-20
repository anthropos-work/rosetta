# M238 — Spec notes

Topic → doc → code triples + chapter-body / language-path findings accumulate here during build.

## Chapter-body demo path (#3 Start→404)
- Bodies are backend-authoritative, no FS fallback; the catalog demopatch covers only the catalog.
- Option: a chapter-body FS-fallback demopatch (analogous to `academy-fs-published-fallback`) vs wiring the academy backend for the demo.

## Language error (#2)
- Re-triaged in M237; likely the same backend-null path as #3.

## Academy presence/coverage sweep
- Extend the sweep to assert chapter-body render (not just catalog cards).

_(will accumulate topic → doc → code triples during build)_
