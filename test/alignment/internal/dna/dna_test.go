package dna

import "testing"

func TestGenesWeightDerivation(t *testing.T) {
	override := 5
	d := &DNA{
		SchemaVersion: 1,
		Capabilities: []Capability{
			{ID: "Add", Criticality: Critical, Variants: []Variant{
				{ID: "a", Operator: OpExact},
				{ID: "b", Operator: OpExact, Weight: &override},
			}},
			{ID: "Greet", Criticality: Standard, Variants: []Variant{
				{ID: "c", Operator: OpExact},
			}},
		},
	}
	genes := d.Genes()
	if len(genes) != 3 {
		t.Fatalf("want 3 genes, got %d", len(genes))
	}
	// sorted by id: Add/a, Add/b, Greet/c
	if genes[0].ID != "Add/a" || genes[0].Weight != 3 {
		t.Errorf("Add/a: want weight 3 (critical), got id=%s w=%d", genes[0].ID, genes[0].Weight)
	}
	if genes[1].ID != "Add/b" || genes[1].Weight != 5 {
		t.Errorf("Add/b: want override weight 5, got id=%s w=%d", genes[1].ID, genes[1].Weight)
	}
	if genes[2].ID != "Greet/c" || genes[2].Weight != 2 {
		t.Errorf("Greet/c: want weight 2 (standard), got id=%s w=%d", genes[2].ID, genes[2].Weight)
	}
}

func TestValidateCatchesProblems(t *testing.T) {
	bad := &DNA{
		SchemaVersion: 2, // wrong version
		Capabilities: []Capability{
			{ID: "X", Criticality: "bogus", Variants: []Variant{
				{ID: "v", Operator: "nope"}, // invalid operator
			}},
			{ID: "Y", Criticality: Standard, Variants: []Variant{
				{ID: "n", Operator: OpNormalized}, // normalized without normalize paths
			}},
		},
	}
	errs := bad.Validate()
	if len(errs) < 4 {
		t.Fatalf("want >=4 validation errors, got %d: %v", len(errs), errs)
	}
}

func TestValidateGoodAndDuplicate(t *testing.T) {
	good := &DNA{
		SchemaVersion: 1,
		Capabilities: []Capability{
			{ID: "X", Criticality: Critical, Variants: []Variant{{ID: "v", Operator: OpExact}}},
		},
	}
	if errs := good.Validate(); len(errs) != 0 {
		t.Fatalf("want no errors, got %v", errs)
	}

	dup := &DNA{
		SchemaVersion: 1,
		Capabilities: []Capability{
			{ID: "X", Criticality: Critical, Variants: []Variant{
				{ID: "v", Operator: OpExact},
				{ID: "v", Operator: OpExact}, // duplicate gene id X/v
			}},
		},
	}
	if errs := dup.Validate(); len(errs) == 0 {
		t.Fatal("expected a duplicate-gene-id error")
	}
}
