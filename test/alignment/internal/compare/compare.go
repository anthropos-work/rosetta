// Package compare is the heart of the framework: it applies a gene's equivalence
// operator to a (source, mirror) outcome pair, and aggregates per-gene verdicts into
// a weighted 0–100% alignment score. This is where divergence is detected.
package compare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"anthropos.dev/alignment/internal/dna"
	"anthropos.dev/alignment/internal/outcome"
)

// GeneResult is the alignment verdict for one gene.
type GeneResult struct {
	GeneID     string `json:"gene_id"`
	Capability string `json:"capability"`
	Operator   string `json:"operator"`
	Weight     int    `json:"weight"`
	Critical   bool   `json:"critical"`
	Aligned    bool   `json:"aligned"`
	Detail     string `json:"detail,omitempty"` // why it diverged (empty when aligned)
}

// CapRollup is the per-capability aligned/total tally.
type CapRollup struct {
	Aligned int `json:"aligned"`
	Total   int `json:"total"`
}

// Report is the full alignment outcome: scores, rollups, and per-gene verdicts.
type Report struct {
	Source        dna.TargetRef        `json:"source"`
	Mirror        dna.TargetRef        `json:"mirror"`
	TotalGenes    int                  `json:"total_genes"`
	AlignedGenes  int                  `json:"aligned_genes"`
	Overall       float64              `json:"overall_score"`
	Critical      float64              `json:"critical_score"`
	PerCapability map[string]CapRollup `json:"per_capability"`
	Genes         []GeneResult         `json:"genes"`
}

// GateMet reports whether the report clears the given thresholds. Zero thresholds
// (the default) are always met, so an un-gated run exits 0.
func (r *Report) GateMet(minOverall, minCritical float64) bool {
	return r.Overall >= minOverall && r.Critical >= minCritical
}

// Evaluate compares source vs mirror outcomes for every gene of the DNA and scores them.
func Evaluate(d *dna.DNA, source, mirror outcome.Set) *Report {
	rep := &Report{Source: d.Source, Mirror: d.Mirror, PerCapability: map[string]CapRollup{}}
	var sumW, alignedW, critTotal, critAligned int
	for _, g := range d.Genes() {
		s, sOK := source[g.ID]
		m, mOK := mirror[g.ID]
		aligned, detail := false, ""
		switch {
		case !sOK:
			detail = "no source outcome (golden missing)"
		case !mOK:
			detail = "no mirror outcome (runner omitted this gene)"
		default:
			aligned, detail = compareGene(g, s, m)
		}

		rep.Genes = append(rep.Genes, GeneResult{
			GeneID: g.ID, Capability: g.Capability, Operator: string(g.Operator),
			Weight: g.Weight, Critical: g.Criticality == dna.Critical, Aligned: aligned, Detail: detail,
		})
		rep.TotalGenes++
		sumW += g.Weight
		roll := rep.PerCapability[g.Capability]
		roll.Total++
		if aligned {
			rep.AlignedGenes++
			alignedW += g.Weight
			roll.Aligned++
		}
		rep.PerCapability[g.Capability] = roll
		if g.Criticality == dna.Critical {
			critTotal++
			if aligned {
				critAligned++
			}
		}
	}
	rep.Overall = pct(alignedW, sumW)
	rep.Critical = pct(critAligned, critTotal)
	return rep
}

// compareGene returns whether one gene's mirror outcome matches the source outcome
// under the gene's operator, plus a human detail string when it diverges.
func compareGene(g dna.Gene, s, m outcome.Outcome) (bool, string) {
	// error_class operator: only the error class matters.
	if g.Operator == dna.OpErrorClass {
		if eqErr(s.ErrorClass, m.ErrorClass) {
			return true, ""
		}
		return false, fmt.Sprintf("error_class: source=%s mirror=%s", errStr(s.ErrorClass), errStr(m.ErrorClass))
	}
	// All value operators require the error class to match first.
	if !eqErr(s.ErrorClass, m.ErrorClass) {
		return false, fmt.Sprintf("error_class differs: source=%s mirror=%s", errStr(s.ErrorClass), errStr(m.ErrorClass))
	}
	// Both errored the same way → aligned (no value to compare).
	if s.ErrorClass != nil {
		return true, ""
	}
	switch g.Operator {
	case dna.OpExact:
		return eqValue(s.Value, m.Value)
	case dna.OpShape:
		return eqShape(s.Value, m.Value)
	case dna.OpNormalized:
		sv, err1 := zeroPaths(s.Value, g.Normalize)
		mv, err2 := zeroPaths(m.Value, g.Normalize)
		if err1 != nil || err2 != nil {
			return false, "normalize: value not a JSON object"
		}
		return eqValue(sv, mv)
	default:
		return false, "unknown operator"
	}
}

func eqErr(a, b *string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	return a == nil || *a == *b
}

func errStr(p *string) string {
	if p == nil {
		return "<none>"
	}
	return *p
}

// eqValue compares two JSON values canonically (object key order is normalized by json.Marshal).
func eqValue(a, b json.RawMessage) (bool, string) {
	ca, err := canonical(a)
	if err != nil {
		return false, "source value: invalid JSON"
	}
	cb, err := canonical(b)
	if err != nil {
		return false, "mirror value: invalid JSON"
	}
	if bytes.Equal(ca, cb) {
		return true, ""
	}
	return false, fmt.Sprintf("value differs: source=%s mirror=%s", truncate(ca), truncate(cb))
}

// decodeAny parses JSON into a generic value, preserving integer precision (UseNumber)
// so large IDs / timestamps compare exactly rather than lossily via float64.
func decodeAny(raw json.RawMessage) (any, error) {
	dec := json.NewDecoder(bytes.NewReader(nz(raw)))
	dec.UseNumber()
	var v any
	if err := dec.Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}

func canonical(raw json.RawMessage) ([]byte, error) {
	v, err := decodeAny(raw)
	if err != nil {
		return nil, err
	}
	return json.Marshal(v) // map keys are emitted sorted → canonical form
}

// eqShape compares JSON structure: matching keys and value *types*, values ignored.
func eqShape(a, b json.RawMessage) (bool, string) {
	va, err := decodeAny(a)
	if err != nil {
		return false, "source value: invalid JSON"
	}
	vb, err := decodeAny(b)
	if err != nil {
		return false, "mirror value: invalid JSON"
	}
	if shapeEq(va, vb) {
		return true, ""
	}
	return false, "shape differs (keys/types)"
}

func shapeEq(a, b any) bool {
	switch av := a.(type) {
	case map[string]any:
		bv, ok := b.(map[string]any)
		if !ok || len(av) != len(bv) {
			return false
		}
		for k, v := range av {
			bvv, ok := bv[k]
			if !ok || !shapeEq(v, bvv) {
				return false
			}
		}
		return true
	case []any:
		bv, ok := b.([]any)
		if !ok || len(av) != len(bv) {
			return false
		}
		for i := range av {
			if !shapeEq(av[i], bv[i]) {
				return false
			}
		}
		return true
	case json.Number:
		_, ok := b.(json.Number)
		return ok
	case string:
		_, ok := b.(string)
		return ok
	case bool:
		_, ok := b.(bool)
		return ok
	case nil:
		return b == nil
	default:
		return false
	}
}

// zeroPaths deletes each dot-path from a JSON object before comparison (operator=normalized).
func zeroPaths(raw json.RawMessage, paths []string) (json.RawMessage, error) {
	if len(paths) == 0 {
		return raw, nil
	}
	v, err := decodeAny(raw)
	if err != nil {
		return nil, err
	}
	for _, p := range paths {
		deletePath(v, strings.Split(p, "."))
	}
	return json.Marshal(v)
}

func deletePath(v any, path []string) {
	m, ok := v.(map[string]any)
	if !ok || len(path) == 0 {
		return
	}
	if len(path) == 1 {
		delete(m, path[0])
		return
	}
	if child, ok := m[path[0]]; ok {
		deletePath(child, path[1:])
	}
}

func nz(r json.RawMessage) []byte {
	if len(r) == 0 {
		return []byte("null")
	}
	return r
}

func pct(n, d int) float64 {
	if d == 0 {
		return 100.0
	}
	return round1(float64(n) / float64(d) * 100)
}

func round1(f float64) float64 { return float64(int(f*10+0.5)) / 10 }

func truncate(b []byte) string {
	const max = 120
	if s := string(b); len(s) > max {
		return s[:max] + "..."
	} else {
		return s
	}
}
