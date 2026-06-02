package compare

import (
	"testing"

	"anthropos.dev/alignment/internal/dna"
)

// FuzzCompareValue guards the value-comparison core: arbitrary (possibly invalid) JSON on
// either side, under every operator, must never panic — it must diverge gracefully.
func FuzzCompareValue(f *testing.F) {
	f.Add([]byte(`{"a":1}`), []byte(`{"a":2}`))
	f.Add([]byte(`5`), []byte(`"x"`))
	f.Add([]byte(``), []byte(`null`))
	f.Add([]byte(`{bad`), []byte(`[1,2`))
	f.Add([]byte(`[1,2,3]`), []byte(`{"k":[1,{"n":2}]}`))
	f.Fuzz(func(t *testing.T, a, b []byte) {
		for _, op := range []dna.Operator{dna.OpExact, dna.OpShape, dna.OpNormalized, dna.OpErrorClass} {
			g := dna.Gene{ID: "C/v", Operator: op, Normalize: []string{"id"}}
			_, _ = compareGene(g, val(string(a)), val(string(b))) // must not panic
		}
	})
}
