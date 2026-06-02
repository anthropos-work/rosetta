package dna

import (
	"os"
	"path/filepath"
	"testing"
)

// FuzzLoad guards DNA parsing: arbitrary bytes in a DNA file must never panic Load,
// Validate, or Genes — they must error or return a value that's safe to walk.
func FuzzLoad(f *testing.F) {
	f.Add([]byte(`{"schema_version":1,"capabilities":[]}`))
	f.Add([]byte(`{"schema_version":1,"capabilities":[{"id":"C","criticality":"critical","variants":[{"id":"v","operator":"exact"}]}]}`))
	f.Add([]byte(`{`))
	f.Add([]byte(``))
	f.Fuzz(func(t *testing.T, data []byte) {
		p := filepath.Join(t.TempDir(), "d.json")
		if err := os.WriteFile(p, data, 0o644); err != nil {
			t.Skip()
		}
		d, err := Load(p)
		if err == nil && d != nil {
			_ = d.Validate() // must not panic
			_ = d.Genes()
			_ = d.GeneIDs()
		}
	})
}
