# M242 — Spec notes

Topic → doc → code triples + cockpit-render findings accumulate here during build.

## (1) Row layout — regroup by requirement tuple
- Regroup by `(sim_type, modality)` → `target | passed login options | not-passed login options` on one row (render-only; fields exist).

## (2) Tab selector — into the white header
- Move into the white header, right, vertically centered (restructure `cockpit.py` header to flex).
- Preserve the byte-identical-when-no-content-manifest invariant.

## (3) Hero icon bg by user-type
- manager = orange / employee = indigo (reuse the badge palette); derive a candidate color = `is_hiring && vantage != manager`.

_(will accumulate topic → doc → code triples during build)_
