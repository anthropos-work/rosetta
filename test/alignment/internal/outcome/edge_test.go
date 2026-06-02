package outcome

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestParseSetError(t *testing.T) {
	if _, err := ParseSet([]byte(`{bad`)); err == nil {
		t.Error("expected a parse error on malformed outcomes")
	}
}

func TestGoldenPathNoSlash(t *testing.T) {
	dir := t.TempDir()
	// a gene id without a slash uses the flat fallback path
	if err := WriteGolden(dir, "flatgene", Outcome{Value: json.RawMessage("1")}); err != nil {
		t.Fatal(err)
	}
	o, err := LoadGolden(dir, "flatgene")
	if err != nil {
		t.Fatal(err)
	}
	if string(o.Value) != "1" {
		t.Errorf("got %s", o.Value)
	}
}

func TestLoadGoldenBadJSON(t *testing.T) {
	dir := t.TempDir()
	p := goldenPath(dir, "A/v")
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(p, []byte(`{bad`), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadGolden(dir, "A/v"); err == nil {
		t.Error("expected a parse error on a malformed golden")
	}
}
