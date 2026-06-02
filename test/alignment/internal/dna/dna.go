// Package dna models the Alignment DNA: the enumerated set of (capability × variant)
// "genes" that defines what a faithful mirror of a source engine must reproduce.
//
// The DNA is the score's denominator. See knowledge/plan/.../m0-alignment-framework/spec-notes.md
// for the frozen contract this implements.
package dna

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// Criticality sets a capability's default gene weight and feeds the "critical %" gate.
type Criticality string

const (
	Critical Criticality = "critical"
	Standard Criticality = "standard"
	Optional Criticality = "optional"
)

// Weight returns the default gene weight for a criticality, or 0 if invalid.
func (c Criticality) Weight() int {
	switch c {
	case Critical:
		return 3
	case Standard:
		return 2
	case Optional:
		return 1
	default:
		return 0
	}
}

// Operator is the equivalence test applied when comparing source vs mirror outcomes.
type Operator string

const (
	OpExact      Operator = "exact"       // canonical-JSON-equal value + equal error class
	OpShape      Operator = "shape"       // same JSON structure (keys + value types); values ignored
	OpNormalized Operator = "normalized"  // exact after zeroing the gene's normalize paths
	OpErrorClass Operator = "error_class" // compare only the error class; value ignored
)

// Valid reports whether o is a known operator.
func (o Operator) Valid() bool {
	switch o {
	case OpExact, OpShape, OpNormalized, OpErrorClass:
		return true
	}
	return false
}

// TargetRef identifies a source or mirror engine.
type TargetRef struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Ref     string `json:"ref,omitempty"`
}

// Variant is one input/scenario class for a capability (axis 2).
type Variant struct {
	ID          string          `json:"id"`
	Description string          `json:"description,omitempty"`
	Operator    Operator        `json:"operator"`
	Input       json.RawMessage `json:"input,omitempty"`
	Normalize   []string        `json:"normalize,omitempty"` // dot-paths zeroed before compare (operator=normalized)
	Weight      *int            `json:"weight,omitempty"`    // nil → derive from capability criticality
}

// Capability is one endpoint/function of the source surface (axis 1).
type Capability struct {
	ID          string      `json:"id"`
	Description string      `json:"description,omitempty"`
	Criticality Criticality `json:"criticality"`
	Variants    []Variant   `json:"variants"`
}

// DNA is the full manifest for one source@version mirrored by one mirror.
type DNA struct {
	SchemaVersion int          `json:"schema_version"`
	Source        TargetRef    `json:"source"`
	Mirror        TargetRef    `json:"mirror"`
	Capabilities  []Capability `json:"capabilities"`
}

// Gene is a flattened (capability, variant) pair with its weight resolved.
type Gene struct {
	ID          string
	Capability  string
	Variant     string
	Operator    Operator
	Criticality Criticality
	Input       json.RawMessage
	Normalize   []string
	Weight      int
}

// GeneID is the canonical join key across DNA, goldens, outcomes, test subtests, and reports.
func GeneID(capability, variant string) string { return capability + "/" + variant }

// Genes flattens the DNA into genes (weights resolved), sorted by gene id for determinism.
func (d *DNA) Genes() []Gene {
	var genes []Gene
	for _, c := range d.Capabilities {
		for _, v := range c.Variants {
			w := c.Criticality.Weight()
			if v.Weight != nil {
				w = *v.Weight
			}
			genes = append(genes, Gene{
				ID:          GeneID(c.ID, v.ID),
				Capability:  c.ID,
				Variant:     v.ID,
				Operator:    v.Operator,
				Criticality: c.Criticality,
				Input:       v.Input,
				Normalize:   v.Normalize,
				Weight:      w,
			})
		}
	}
	sort.Slice(genes, func(i, j int) bool { return genes[i].ID < genes[j].ID })
	return genes
}

// GeneIDs returns just the (sorted) gene ids.
func (d *DNA) GeneIDs() []string {
	genes := d.Genes()
	ids := make([]string, len(genes))
	for i, g := range genes {
		ids[i] = g.ID
	}
	return ids
}

// Load reads and JSON-parses a DNA file, rejecting unknown fields.
func Load(path string) (*DNA, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields()
	var d DNA
	if err := dec.Decode(&d); err != nil {
		return nil, fmt.Errorf("parse DNA %s: %w", path, err)
	}
	return &d, nil
}

// Validate returns every structural problem found (empty slice = valid).
func (d *DNA) Validate() []error {
	var errs []error
	if d.SchemaVersion != 1 {
		errs = append(errs, fmt.Errorf("schema_version: want 1, got %d", d.SchemaVersion))
	}
	if len(d.Capabilities) == 0 {
		errs = append(errs, fmt.Errorf("no capabilities"))
	}
	seen := map[string]bool{}
	for _, c := range d.Capabilities {
		if c.ID == "" {
			errs = append(errs, fmt.Errorf("capability with empty id"))
		}
		if c.Criticality.Weight() == 0 {
			errs = append(errs, fmt.Errorf("%s: invalid criticality %q", c.ID, c.Criticality))
		}
		if len(c.Variants) == 0 {
			errs = append(errs, fmt.Errorf("%s: no variants", c.ID))
		}
		for _, v := range c.Variants {
			gid := GeneID(c.ID, v.ID)
			if v.ID == "" {
				errs = append(errs, fmt.Errorf("%s: variant with empty id", c.ID))
			}
			if seen[gid] {
				errs = append(errs, fmt.Errorf("duplicate gene id %q", gid))
			}
			seen[gid] = true
			if !v.Operator.Valid() {
				errs = append(errs, fmt.Errorf("%s: invalid operator %q", gid, v.Operator))
			}
			if v.Operator == OpNormalized && len(v.Normalize) == 0 {
				errs = append(errs, fmt.Errorf("%s: operator=normalized requires non-empty normalize", gid))
			}
			if v.Weight != nil && *v.Weight <= 0 {
				errs = append(errs, fmt.Errorf("%s: weight must be > 0", gid))
			}
		}
	}
	return errs
}
