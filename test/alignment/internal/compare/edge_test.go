package compare

import (
	"testing"

	"anthropos.dev/alignment/internal/dna"
)

func opGene(op dna.Operator, norm ...string) dna.Gene {
	return dna.Gene{ID: "C/v", Capability: "C", Operator: op, Weight: 1, Criticality: dna.Standard, Normalize: norm}
}

func TestCompareInvalidJSONValue(t *testing.T) {
	aligned, detail := compareGene(opGene(dna.OpExact), val(`{bad`), val(`{"a":1}`))
	if aligned || detail == "" {
		t.Errorf("invalid source JSON should diverge with a detail; got aligned=%v detail=%q", aligned, detail)
	}
}

func TestShapeArrays(t *testing.T) {
	if a, _ := compareGene(opGene(dna.OpShape), val(`[1,2]`), val(`[9,8]`)); !a {
		t.Error("same-length numeric arrays should match on shape")
	}
	if a, _ := compareGene(opGene(dna.OpShape), val(`[1,2]`), val(`[1]`)); a {
		t.Error("different-length arrays should diverge on shape")
	}
	if a, _ := compareGene(opGene(dna.OpShape), val(`[1]`), val(`["x"]`)); a {
		t.Error("array element type mismatch should diverge on shape")
	}
}

func TestNormalizedNestedPath(t *testing.T) {
	g := opGene(dna.OpNormalized, "meta.id")
	if a, d := compareGene(g, val(`{"v":1,"meta":{"id":"x"}}`), val(`{"v":1,"meta":{"id":"y"}}`)); !a {
		t.Errorf("nested normalize path should ignore meta.id; aligned=%v detail=%q", a, d)
	}
	if a, _ := compareGene(g, val(`{"v":1,"meta":{"id":"x"}}`), val(`{"v":2,"meta":{"id":"y"}}`)); a {
		t.Error("a real difference outside the normalize path should still diverge")
	}
}

// TestLargeIntExactValue pins the precision-safe comparison (the UseNumber fix): large
// integers must compare exactly, and an off-by-one must diverge — no float rounding.
func TestLargeIntExactValue(t *testing.T) {
	big := `9223372036854775807`
	if a, d := compareGene(opGene(dna.OpExact), val(big), val(big)); !a {
		t.Errorf("equal large ints should align; aligned=%v detail=%q", a, d)
	}
	if a, _ := compareGene(opGene(dna.OpExact), val(big), val(`9223372036854775806`)); a {
		t.Error("off-by-one large int should diverge (no float rounding)")
	}
}
