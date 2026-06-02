package outcome

import "testing"

// FuzzParseSet guards the runner→alignctl boundary: arbitrary bytes from a runner's stdout
// must never panic the parser.
func FuzzParseSet(f *testing.F) {
	f.Add([]byte(`{"A/v":{"value":1,"error_class":null}}`))
	f.Add([]byte(`{}`))
	f.Add([]byte(`not json`))
	f.Add([]byte(``))
	f.Fuzz(func(t *testing.T, data []byte) {
		_, _ = ParseSet(data) // must not panic
	})
}
