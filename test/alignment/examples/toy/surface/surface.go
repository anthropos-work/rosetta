// Package surface defines the contract a toy "engine" exposes. Both the source and the
// mirror implement it. It stands in for a real SDK surface (e.g. Clerk's) so the alignment
// framework can be proven end-to-end without any external dependency.
package surface

// Surface is the toy engine's two-capability public surface.
type Surface interface {
	// Add returns a+b, or an error when the result overflows int64.
	Add(a, b int64) (int64, error)
	// Greet returns a greeting for name. A real engine might Unicode-normalize the name.
	Greet(name string) (string, error)
}

// ErrOverflow is the stable error returned by Add on overflow. Its Error() string is the
// "error class" the alignment framework compares.
type overflowError struct{}

func (overflowError) Error() string { return "overflow" }

// ErrOverflow is returned by Add when a+b overflows int64.
var ErrOverflow error = overflowError{}
