# M15 — Spec notes

Technical notes accumulate here during build.

## safety.md — read-side (never reads private data)
_Tenant firewall AssertPublicOnly; public predicates (org_id IS NULL; private=false AND tenant_id IS NULL AND status=published); public-only data-DNA gene; capture read-only._

## safety.md — write-side (never touches prod)
_3-layer isolation guard; never-write shared Directus / prod-S3; capture-source policy; n=0-dev guards; audit-proven zero pollution._

## Dual-repo KB refresh
_rosetta corpus + rosetta-extensions/knowledge/ to the v1.3 converged stack model._
