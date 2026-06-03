package outcome

import (
	"encoding/json"
	"testing"
)

func TestGoldenRoundTrip(t *testing.T) {
	dir := t.TempDir()
	ec := "overflow"
	cases := map[string]Outcome{
		"Add/two-positives": {Value: json.RawMessage("5")},
		"Add/overflow":      {ErrorClass: &ec},
		"Greet/basic":       {Value: json.RawMessage(`"Hello, World!"`)},
	}
	for id, o := range cases {
		if err := WriteGolden(dir, id, o); err != nil {
			t.Fatalf("write %s: %v", id, err)
		}
	}
	for id, want := range cases {
		got, err := LoadGolden(dir, id)
		if err != nil {
			t.Fatalf("load %s: %v", id, err)
		}
		// A nil value serializes as JSON null per the outcomes protocol, so it round-trips
		// as the bytes "null"; compare canonically rather than byte-for-byte.
		if canon(got.Value) != canon(want.Value) {
			t.Errorf("%s value: got %s want %s", id, canon(got.Value), canon(want.Value))
		}
		if (got.ErrorClass == nil) != (want.ErrorClass == nil) {
			t.Errorf("%s error_class nil-ness mismatch", id)
		}
	}
}

func canon(r json.RawMessage) string {
	if len(r) == 0 {
		return "null"
	}
	return string(r)
}

func TestParseSet(t *testing.T) {
	s, err := ParseSet([]byte(`{"A/v":{"value":5,"error_class":null}}`))
	if err != nil {
		t.Fatal(err)
	}
	if string(s["A/v"].Value) != "5" {
		t.Errorf("got value %q", s["A/v"].Value)
	}
	if s["A/v"].ErrorClass != nil {
		t.Error("error_class should be nil")
	}
}

func TestLoadGoldenSetSkipsMissing(t *testing.T) {
	dir := t.TempDir()
	if err := WriteGolden(dir, "A/v", Outcome{Value: json.RawMessage("1")}); err != nil {
		t.Fatal(err)
	}
	set, err := LoadGoldenSet(dir, []string{"A/v", "A/missing"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := set["A/v"]; !ok {
		t.Error("A/v should be present")
	}
	if _, ok := set["A/missing"]; ok {
		t.Error("A/missing should be skipped, not present")
	}
}
