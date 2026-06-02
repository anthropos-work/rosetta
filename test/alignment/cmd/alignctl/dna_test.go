package main

import (
	"encoding/json"
	"testing"

	"anthropos.dev/alignment/internal/dna"
)

func mkGene(op dna.Operator, w int, crit dna.Criticality, input, norm string) dna.Gene {
	var n []string
	if norm != "" {
		n = []string{norm}
	}
	return dna.Gene{
		ID: "C/v", Capability: "C", Variant: "v", Operator: op, Weight: w,
		Criticality: crit, Input: json.RawMessage(input), Normalize: n,
	}
}

func TestGeneChanged(t *testing.T) {
	base := mkGene(dna.OpExact, 3, dna.Critical, `{"a":1}`, "")
	tests := []struct {
		name string
		b    dna.Gene
		want bool
	}{
		{"identical", mkGene(dna.OpExact, 3, dna.Critical, `{"a":1}`, ""), false},
		{"reformatted-input", mkGene(dna.OpExact, 3, dna.Critical, `{ "a":  1 }`, ""), false}, // canonical → no drift
		{"operator", mkGene(dna.OpShape, 3, dna.Critical, `{"a":1}`, ""), true},
		{"weight", mkGene(dna.OpExact, 2, dna.Critical, `{"a":1}`, ""), true},
		{"criticality", mkGene(dna.OpExact, 3, dna.Standard, `{"a":1}`, ""), true},
		{"input-value", mkGene(dna.OpExact, 3, dna.Critical, `{"a":2}`, ""), true},
		{"normalize", mkGene(dna.OpExact, 3, dna.Critical, `{"a":1}`, "id"), true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := geneChanged(base, tc.b); got != tc.want {
				t.Errorf("geneChanged=%v want %v", got, tc.want)
			}
		})
	}
}

func TestJoinNormOrderIndependent(t *testing.T) {
	if joinNorm([]string{"a", "b"}) != joinNorm([]string{"b", "a"}) {
		t.Error("joinNorm should be order-independent")
	}
	if joinNorm([]string{"a"}) == joinNorm([]string{"a", "b"}) {
		t.Error("different normalize sets should differ")
	}
}

func TestGeneMap(t *testing.T) {
	d := &dna.DNA{SchemaVersion: 1, Capabilities: []dna.Capability{
		{ID: "A", Criticality: dna.Critical, Variants: []dna.Variant{
			{ID: "x", Operator: dna.OpExact}, {ID: "y", Operator: dna.OpExact},
		}},
	}}
	m := geneMap(d)
	if len(m) != 2 || m["A/x"].Operator != dna.OpExact {
		t.Errorf("geneMap wrong: %+v", m)
	}
}
