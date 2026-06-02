package outcome

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
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
	p, err := goldenPath(dir, "A/v")
	if err != nil {
		t.Fatal(err)
	}
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

// TestGoldenPathContainment pins the defense-in-depth guard: a gene id whose path-traversal
// would resolve OUTSIDE the golden dir must be rejected for read and write. (Ids that resolve
// to a weird-but-contained filename, e.g. ".." or "/abs", are caught upstream by dna.Validate;
// the golden IO only has to refuse actual escapes.)
func TestGoldenPathContainment(t *testing.T) {
	dir := t.TempDir()
	for _, bad := range []string{"../escape", "../../etc/passwd", "a/../../x"} {
		if _, err := goldenPath(dir, bad); err == nil {
			t.Errorf("goldenPath(%q) should be rejected as escaping the dir", bad)
		}
		if err := WriteGolden(dir, bad, Outcome{Value: json.RawMessage("1")}); err == nil {
			t.Errorf("WriteGolden(%q) should refuse to escape the dir", bad)
		}
		if _, err := LoadGolden(dir, bad); err == nil {
			t.Errorf("LoadGolden(%q) should refuse to escape the dir", bad)
		}
	}
	// a normal gene id stays contained
	if p, err := goldenPath(dir, "Cap/variant"); err != nil || !strings.HasPrefix(p, dir) {
		t.Errorf("normal gene id should resolve within dir: p=%s err=%v", p, err)
	}
}
