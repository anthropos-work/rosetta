// Package outcome defines the normalized result of running one gene against a target,
// plus the golden (record/replay) file IO. The Outcome type is the engine⇄alignctl contract:
// a runner prints a map of gene id → Outcome as JSON to stdout.
package outcome

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Outcome is the normalized result of one capability+variant invocation.
// Exactly one of (Value present, ErrorClass present) is the meaningful field,
// but both are always serialized so the shape is stable.
type Outcome struct {
	Value      json.RawMessage `json:"value"`
	ErrorClass *string         `json:"error_class"`
}

// Set maps gene id → Outcome, as emitted by a runner and as loaded from goldens.
type Set map[string]Outcome

// ParseSet decodes a runner's stdout (a JSON object of gene id → Outcome).
func ParseSet(b []byte) (Set, error) {
	var s Set
	if err := json.Unmarshal(b, &s); err != nil {
		return nil, fmt.Errorf("parse outcomes: %w", err)
	}
	return s, nil
}

// goldenPath maps a gene id ("<Capability>/<variant>") to its on-disk golden file.
func goldenPath(dir, geneID string) string {
	parts := strings.SplitN(geneID, "/", 2)
	if len(parts) == 2 {
		return filepath.Join(dir, parts[0], parts[1]+".json")
	}
	return filepath.Join(dir, geneID+".json")
}

// WriteGolden records one gene's source outcome under dir.
func WriteGolden(dir, geneID string, o Outcome) error {
	p := goldenPath(dir, geneID)
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(p, append(b, '\n'), 0o644)
}

// LoadGolden reads one gene's recorded source outcome.
func LoadGolden(dir, geneID string) (Outcome, error) {
	b, err := os.ReadFile(goldenPath(dir, geneID))
	if err != nil {
		return Outcome{}, err
	}
	var o Outcome
	if err := json.Unmarshal(b, &o); err != nil {
		return Outcome{}, fmt.Errorf("parse golden for %s: %w", geneID, err)
	}
	return o, nil
}

// LoadGoldenSet loads goldens for every given gene id. A missing golden is left out
// of the set (the comparator reports it as "no source outcome"), not a hard error.
func LoadGoldenSet(dir string, geneIDs []string) (Set, error) {
	s := make(Set, len(geneIDs))
	for _, id := range geneIDs {
		o, err := LoadGolden(dir, id)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return nil, err
		}
		s[id] = o
	}
	return s, nil
}
