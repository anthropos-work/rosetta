// Package source is the canonical toy engine — the "source of truth" the mirror must match.
package source

import (
	"math"
	"strings"

	"anthropos.dev/alignment/examples/toy/surface"
)

// Engine is the source implementation of surface.Surface.
type Engine struct{}

var _ surface.Surface = Engine{}

// Add returns a+b, erroring on int64 overflow.
func (Engine) Add(a, b int64) (int64, error) {
	if (b > 0 && a > math.MaxInt64-b) || (b < 0 && a < math.MinInt64-b) {
		return 0, surface.ErrOverflow
	}
	return a + b, nil
}

// Greet normalizes whitespace (trim + collapse internal runs) before formatting — the
// input-cleaning behavior the toy mirror intentionally omits. Stands in for the kind of
// quiet normalization a real SDK does that a naive mirror forgets.
func (Engine) Greet(name string) (string, error) {
	clean := strings.Join(strings.Fields(name), " ")
	return "Hello, " + clean + "!", nil
}
