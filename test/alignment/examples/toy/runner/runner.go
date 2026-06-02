// Package runner invokes the toy engine's capabilities for each DNA gene and collects
// normalized outcomes. It is shared by cmd/toyrun (the alignctl runner binary) and the
// build-tagged alignment test, so both produce identical outcomes.
package runner

import (
	"encoding/json"
	"fmt"

	"anthropos.dev/alignment/examples/toy/surface"
	"anthropos.dev/alignment/internal/dna"
	"anthropos.dev/alignment/internal/outcome"
)

// Run invokes eng for every gene in d and returns the outcome set.
func Run(eng surface.Surface, d *dna.DNA) (outcome.Set, error) {
	set := outcome.Set{}
	for _, g := range d.Genes() {
		o, err := Invoke(eng, g)
		if err != nil {
			return nil, fmt.Errorf("gene %s: %w", g.ID, err)
		}
		set[g.ID] = o
	}
	return set, nil
}

// Invoke runs a single gene's input through eng and normalizes the result into an Outcome.
func Invoke(eng surface.Surface, g dna.Gene) (outcome.Outcome, error) {
	switch g.Capability {
	case "Add":
		var in struct {
			A int64 `json:"a"`
			B int64 `json:"b"`
		}
		if err := json.Unmarshal(g.Input, &in); err != nil {
			return outcome.Outcome{}, err
		}
		return result(eng.Add(in.A, in.B))
	case "Greet":
		var in struct {
			Name string `json:"name"`
		}
		if err := json.Unmarshal(g.Input, &in); err != nil {
			return outcome.Outcome{}, err
		}
		return result(eng.Greet(in.Name))
	default:
		return outcome.Outcome{}, fmt.Errorf("unknown capability %q", g.Capability)
	}
}

// result turns a (value, error) pair into a normalized Outcome: on error, the error's
// class string; otherwise the JSON-marshaled value.
func result[T any](value T, err error) (outcome.Outcome, error) {
	if err != nil {
		ec := err.Error()
		return outcome.Outcome{ErrorClass: &ec}, nil
	}
	b, e := json.Marshal(value)
	if e != nil {
		return outcome.Outcome{}, e
	}
	return outcome.Outcome{Value: b}, nil
}
