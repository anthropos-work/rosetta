// Package mirror is the toy mirror engine. It reproduces the source faithfully EXCEPT for
// one intentional divergence (Greet's whitespace normalization) so the alignment framework
// has something real to catch — proving it detects misalignment, not just reports green.
package mirror

import (
	"math"

	"anthropos.dev/alignment/examples/toy/surface"
)

// Engine is the mirror implementation of surface.Surface.
type Engine struct{}

var _ surface.Surface = Engine{}

// Add matches the source exactly (including overflow behavior).
func (Engine) Add(a, b int64) (int64, error) {
	if (b > 0 && a > math.MaxInt64-b) || (b < 0 && a < math.MinInt64-b) {
		return 0, surface.ErrOverflow
	}
	return a + b, nil
}

// Greet does NOT normalize whitespace (the intentional divergence). For names without
// surrounding/internal extra whitespace this is identical to the source; for a padded name
// it diverges — exactly the gene the framework must flag.
func (Engine) Greet(name string) (string, error) {
	return "Hello, " + name + "!", nil
}
