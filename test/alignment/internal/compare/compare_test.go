package compare

import (
	"encoding/json"
	"testing"

	"anthropos.dev/alignment/internal/dna"
	"anthropos.dev/alignment/internal/outcome"
)

func val(s string) outcome.Outcome  { return outcome.Outcome{Value: json.RawMessage(s)} }
func errc(s string) outcome.Outcome { return outcome.Outcome{ErrorClass: &s} }

func TestOperators(t *testing.T) {
	tests := []struct {
		name    string
		op      dna.Operator
		norm    []string
		s, m    outcome.Outcome
		aligned bool
	}{
		{"exact-equal", dna.OpExact, nil, val(`5`), val(`5`), true},
		{"exact-diff", dna.OpExact, nil, val(`5`), val(`6`), false},
		{"exact-obj-key-order", dna.OpExact, nil, val(`{"a":1,"b":2}`), val(`{"b":2,"a":1}`), true},
		{"shape-same", dna.OpShape, nil, val(`{"a":1}`), val(`{"a":99}`), true},
		{"shape-diff-type", dna.OpShape, nil, val(`{"a":1}`), val(`{"a":"x"}`), false},
		{"shape-diff-keys", dna.OpShape, nil, val(`{"a":1}`), val(`{"b":1}`), false},
		{"normalized-ignores-id", dna.OpNormalized, []string{"id"}, val(`{"id":"x","v":1}`), val(`{"id":"y","v":1}`), true},
		{"normalized-real-diff", dna.OpNormalized, []string{"id"}, val(`{"id":"x","v":1}`), val(`{"id":"y","v":2}`), false},
		{"errorclass-match", dna.OpErrorClass, nil, errc("overflow"), errc("overflow"), true},
		{"errorclass-mismatch", dna.OpErrorClass, nil, errc("overflow"), errc("nan"), false},
		{"value-vs-error", dna.OpExact, nil, val(`5`), errc("boom"), false},
		{"both-error-same-under-exact", dna.OpExact, nil, errc("boom"), errc("boom"), true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := dna.Gene{ID: "C/" + tc.name, Capability: "C", Operator: tc.op, Normalize: tc.norm, Weight: 1, Criticality: dna.Standard}
			got, detail := compareGene(g, tc.s, tc.m)
			if got != tc.aligned {
				t.Errorf("aligned=%v want %v (detail=%q)", got, tc.aligned, detail)
			}
			if !got && detail == "" {
				t.Error("a divergence must carry a non-empty detail")
			}
		})
	}
}

func TestEvaluateScoreAndDetection(t *testing.T) {
	w2 := 2
	d := &dna.DNA{
		SchemaVersion: 1,
		Source:        dna.TargetRef{Name: "s"}, Mirror: dna.TargetRef{Name: "m"},
		Capabilities: []dna.Capability{
			{ID: "Add", Criticality: dna.Critical, Variants: []dna.Variant{
				{ID: "ok", Operator: dna.OpExact},
			}},
			{ID: "Greet", Criticality: dna.Standard, Variants: []dna.Variant{
				{ID: "ok", Operator: dna.OpExact},
				{ID: "bad", Operator: dna.OpExact, Weight: &w2},
			}},
		},
	}
	src := outcome.Set{"Add/ok": val(`1`), "Greet/ok": val(`"hi"`), "Greet/bad": val(`"clean"`)}
	mir := outcome.Set{"Add/ok": val(`1`), "Greet/ok": val(`"hi"`), "Greet/bad": val(`"dirty"`)}
	rep := Evaluate(d, src, mir)

	// weights: Add/ok=3, Greet/ok=2, Greet/bad=2(override). total=7, aligned weight=5 → 71.4
	if rep.Overall != 71.4 {
		t.Errorf("overall = %.1f, want 71.4", rep.Overall)
	}
	if rep.Critical != 100.0 {
		t.Errorf("critical = %.1f, want 100.0 (Add aligned)", rep.Critical)
	}
	if rep.AlignedGenes != 2 || rep.TotalGenes != 3 {
		t.Errorf("aligned/total = %d/%d, want 2/3", rep.AlignedGenes, rep.TotalGenes)
	}

	// detection: the divergent gene must be flagged with a detail
	var found bool
	for _, g := range rep.Genes {
		if g.GeneID == "Greet/bad" {
			found = true
			if g.Aligned {
				t.Error("Greet/bad must be flagged as diverged")
			}
			if g.Detail == "" {
				t.Error("diverged gene must carry a detail")
			}
		}
	}
	if !found {
		t.Error("Greet/bad missing from report.Genes")
	}
}

func TestEvaluateMissingOutcomes(t *testing.T) {
	d := &dna.DNA{
		SchemaVersion: 1,
		Capabilities: []dna.Capability{
			{ID: "C", Criticality: dna.Standard, Variants: []dna.Variant{{ID: "v", Operator: dna.OpExact}}},
		},
	}
	rep := Evaluate(d, outcome.Set{}, outcome.Set{}) // both missing
	if rep.AlignedGenes != 0 || rep.Overall != 0 {
		t.Errorf("missing outcomes should score 0, got %d aligned / %.1f", rep.AlignedGenes, rep.Overall)
	}
	if rep.Genes[0].Detail == "" {
		t.Error("missing-outcome gene should explain why")
	}
}

func TestGateMet(t *testing.T) {
	r := &Report{Overall: 95, Critical: 100}
	if !r.GateMet(95, 100) {
		t.Error("should meet the gate exactly")
	}
	if r.GateMet(96, 100) {
		t.Error("should fail the overall gate")
	}
	if r.GateMet(95, 100.1) {
		t.Error("should fail the critical gate")
	}
	if !r.GateMet(0, 0) {
		t.Error("zero gate should always pass")
	}
}
