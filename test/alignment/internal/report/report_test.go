package report

import (
	"strings"
	"testing"

	"anthropos.dev/alignment/internal/compare"
	"anthropos.dev/alignment/internal/dna"
)

func TestHumanShowsScoreAndDivergence(t *testing.T) {
	r := &compare.Report{
		Source:     dna.TargetRef{Name: "toy", Version: "v1"},
		Mirror:     dna.TargetRef{Name: "toy-mirror", Version: "v1"},
		TotalGenes: 2, AlignedGenes: 1, Overall: 50, Critical: 100,
		PerCapability: map[string]compare.CapRollup{"Greet": {Aligned: 1, Total: 2}},
		Genes: []compare.GeneResult{
			{GeneID: "Greet/ok", Capability: "Greet", Operator: "exact", Weight: 2, Aligned: true},
			{GeneID: "Greet/bad", Capability: "Greet", Operator: "exact", Weight: 2, Aligned: false, Detail: "value differs"},
		},
	}
	var sb strings.Builder
	Human(&sb, r)
	out := sb.String()
	for _, want := range []string{"overall 50.0%", "critical 100.0%", "Greet/bad", "value differs", "Divergences (1)"} {
		if !strings.Contains(out, want) {
			t.Errorf("human output missing %q:\n%s", want, out)
		}
	}
}

func TestHumanNoDivergence(t *testing.T) {
	r := &compare.Report{
		Source: dna.TargetRef{Name: "a"}, Mirror: dna.TargetRef{Name: "b"},
		TotalGenes: 1, AlignedGenes: 1, Overall: 100, Critical: 100,
		PerCapability: map[string]compare.CapRollup{"C": {Aligned: 1, Total: 1}},
		Genes:         []compare.GeneResult{{GeneID: "C/v", Capability: "C", Aligned: true}},
	}
	var sb strings.Builder
	Human(&sb, r)
	if !strings.Contains(sb.String(), "No divergences") {
		t.Errorf("expected the no-divergence message:\n%s", sb.String())
	}
}
