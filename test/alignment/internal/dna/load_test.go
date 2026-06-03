package dna

import (
	"os"
	"path/filepath"
	"testing"
)

func writeTmp(t *testing.T, body string) string {
	t.Helper()
	p := filepath.Join(t.TempDir(), "d.json")
	if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}

func TestLoadValid(t *testing.T) {
	p := writeTmp(t, `{"schema_version":1,"source":{"name":"s","version":"v"},"mirror":{"name":"m","version":"v"},"capabilities":[{"id":"C","criticality":"critical","variants":[{"id":"v","operator":"exact"}]}]}`)
	d, err := Load(p)
	if err != nil {
		t.Fatal(err)
	}
	if ids := d.GeneIDs(); len(ids) != 1 || ids[0] != "C/v" {
		t.Errorf("GeneIDs wrong: %v", ids)
	}
}

func TestLoadErrors(t *testing.T) {
	t.Run("missing-file", func(t *testing.T) {
		if _, err := Load(filepath.Join(t.TempDir(), "nope.json")); err == nil {
			t.Error("expected error for missing file")
		}
	})
	t.Run("bad-json", func(t *testing.T) {
		if _, err := Load(writeTmp(t, `{not json`)); err == nil {
			t.Error("expected a parse error")
		}
	})
	t.Run("unknown-field-rejected", func(t *testing.T) {
		if _, err := Load(writeTmp(t, `{"schema_version":1,"bogus":true,"capabilities":[]}`)); err == nil {
			t.Error("DisallowUnknownFields should reject 'bogus'")
		}
	})
}

func TestOperatorValid(t *testing.T) {
	for _, op := range []Operator{OpExact, OpShape, OpNormalized, OpErrorClass} {
		if !op.Valid() {
			t.Errorf("%s should be valid", op)
		}
	}
	if Operator("bogus").Valid() {
		t.Error("bogus operator should be invalid")
	}
}

func TestCriticalityWeight(t *testing.T) {
	if Critical.Weight() != 3 || Standard.Weight() != 2 || Optional.Weight() != 1 {
		t.Error("criticality→weight mapping wrong")
	}
	if Criticality("bogus").Weight() != 0 {
		t.Error("invalid criticality should weigh 0")
	}
}
