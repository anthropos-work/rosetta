package dna

import "testing"

func oneGene(capID, varID string, weight *int) *DNA {
	return &DNA{
		SchemaVersion: 1,
		Capabilities: []Capability{
			{ID: capID, Criticality: Critical, Variants: []Variant{
				{ID: varID, Operator: OpExact, Weight: weight},
			}},
		},
	}
}

// TestValidateRejectsUnsafeIDs pins the path-traversal / gene-id-format guard: ids that
// could escape the golden dir (or violate the documented grammar) must be rejected.
func TestValidateRejectsUnsafeIDs(t *testing.T) {
	bad := []struct{ cap, vari string }{
		{"../etc", "x"}, {"A", "../escape"},
		{"A/b", "x"}, {"A", "b/c"}, // embedded slash
		{"", "x"}, {"A", ""}, // empty
		{"..", "x"}, {"A", ".."},
		{"-leading", "x"}, {"A", "-leading"}, // leading dash
		{"A b", "x"}, {"A", "a.b"}, // space / dot
	}
	for _, tc := range bad {
		if errs := oneGene(tc.cap, tc.vari, nil).Validate(); len(errs) == 0 {
			t.Errorf("cap=%q variant=%q should be rejected", tc.cap, tc.vari)
		}
	}
	good := []struct{ cap, vari string }{
		{"CreateOrganization", "duplicate-name"},
		{"Add", "two-positives"},
		{"V2Endpoint", "case-1"},
		{"snake_case", "a_b"},
	}
	for _, tc := range good {
		if errs := oneGene(tc.cap, tc.vari, nil).Validate(); len(errs) != 0 {
			t.Errorf("cap=%q variant=%q should be valid, got %v", tc.cap, tc.vari, errs)
		}
	}
}

// TestValidateWeightBounds pins the upper bound that prevents score-sum integer overflow.
func TestValidateWeightBounds(t *testing.T) {
	zero, neg, huge, ok := 0, -1, maxWeight+1, 5
	if errs := oneGene("A", "x", &zero).Validate(); len(errs) == 0 {
		t.Error("weight 0 should be rejected")
	}
	if errs := oneGene("A", "x", &neg).Validate(); len(errs) == 0 {
		t.Error("negative weight should be rejected")
	}
	if errs := oneGene("A", "x", &huge).Validate(); len(errs) == 0 {
		t.Error("weight above maxWeight should be rejected")
	}
	if errs := oneGene("A", "x", &ok).Validate(); len(errs) != 0 {
		t.Errorf("weight 5 should be valid, got %v", errs)
	}
}
