//go:build alignment

// Alignment tests are a third test class, beside unit and integration. They are marked by
// the `alignment` build tag (plain `go test` skips them) and reported as one subtest per
// gene (parseable/countable via `go test -tags alignment -json`). This file is the toy's
// suite; a real mirror (M1) ships its own equivalent.
package toy

import (
	"testing"

	"anthropos.dev/alignment/examples/toy/mirror"
	"anthropos.dev/alignment/examples/toy/runner"
	"anthropos.dev/alignment/examples/toy/source"
	"anthropos.dev/alignment/internal/compare"
	"anthropos.dev/alignment/internal/dna"
)

// Gate for the TOY. Consumer mirrors set their own; M1 uses 100% critical / >=95% overall.
const (
	gateOverall  = 80.0
	gateCritical = 100.0
)

func TestAlignment(t *testing.T) {
	d, err := dna.Load("dna.json")
	if err != nil {
		t.Fatalf("load DNA: %v", err)
	}
	if errs := d.Validate(); len(errs) > 0 {
		t.Fatalf("invalid DNA: %v", errs)
	}

	srcSet, err := runner.Run(source.Engine{}, d)
	if err != nil {
		t.Fatalf("source run: %v", err)
	}
	mirSet, err := runner.Run(mirror.Engine{}, d)
	if err != nil {
		t.Fatalf("mirror run: %v", err)
	}
	rep := compare.Evaluate(d, srcSet, mirSet)

	// One subtest per gene — the unit of alignment. Names are gene ids so test2json maps
	// 1:1 back to the DNA.
	for _, g := range rep.Genes {
		t.Run(g.GeneID, func(t *testing.T) {
			switch {
			case g.Aligned:
				// pass
			case g.Critical:
				t.Errorf("CRITICAL divergence: %s", g.Detail)
			default:
				// A non-critical divergence within the overall gate is tolerated but logged,
				// so it never goes silent. (The toy's Greet/padded-name lands here.)
				t.Logf("tolerated non-critical divergence: %s", g.Detail)
			}
		})
	}

	t.Logf("alignment: overall %.1f%%  critical %.1f%%  (%d/%d genes)",
		rep.Overall, rep.Critical, rep.AlignedGenes, rep.TotalGenes)
	if rep.Critical < gateCritical {
		t.Errorf("critical score %.1f%% < gate %.1f%%", rep.Critical, gateCritical)
	}
	if rep.Overall < gateOverall {
		t.Errorf("overall score %.1f%% < gate %.1f%%", rep.Overall, gateOverall)
	}
}
