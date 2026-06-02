package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// captureStdout runs fn with os.Stdout redirected to a pipe and returns what it printed
// plus fn's return code. Used to assert in-process subcommand output (instrumented, unlike
// the out-of-process CLI integration test).
func captureStdout(t *testing.T, fn func() int) (string, int) {
	t.Helper()
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w
	code := fn()
	_ = w.Close()
	os.Stdout = old
	b, _ := io.ReadAll(r)
	return string(b), code
}

const sampleDNA = `{"schema_version":1,"source":{"name":"s","version":"v"},"mirror":{"name":"m","version":"v"},"capabilities":[{"id":"Add","criticality":"critical","variants":[{"id":"x","operator":"exact","input":{"a":1}},{"id":"y","operator":"error_class"}]}]}`

func writeDNA(t *testing.T, body string) string {
	t.Helper()
	p := filepath.Join(t.TempDir(), "d.json")
	if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}

func TestDnaValidateAndList(t *testing.T) {
	p := writeDNA(t, sampleDNA)

	out, code := captureStdout(t, func() int { return dnaValidate([]string{"--dna", p}) })
	if code != 0 || !strings.Contains(out, "2 genes valid") {
		t.Errorf("validate: code=%d out=%s", code, out)
	}
	out, code = captureStdout(t, func() int { return dnaList([]string{"--dna", p}) })
	if code != 0 || !strings.Contains(out, "2 genes across 1 capabilities") {
		t.Errorf("list: code=%d out=%s", code, out)
	}
	out, code = captureStdout(t, func() int { return dnaList([]string{"--dna", p, "--json"}) })
	if code != 0 || !strings.Contains(out, `"operator": "exact"`) {
		t.Errorf("list --json: code=%d out=%s", code, out)
	}
}

func TestDnaDiffInProcess(t *testing.T) {
	a := writeDNA(t, sampleDNA)
	if _, code := captureStdout(t, func() int { return dnaDiff([]string{"--old", a, "--new", a}) }); code != 0 {
		t.Errorf("self diff should exit 0, got %d", code)
	}
	// drop a variant → removed gene → exit 1
	b := writeDNA(t, `{"schema_version":1,"source":{"name":"s","version":"v"},"mirror":{"name":"m","version":"v"},"capabilities":[{"id":"Add","criticality":"critical","variants":[{"id":"x","operator":"exact","input":{"a":1}}]}]}`)
	out, code := captureStdout(t, func() int { return dnaDiff([]string{"--old", a, "--new", b, "--json"}) })
	if code != 1 || !strings.Contains(out, "Add/y") {
		t.Errorf("diff should report removed Add/y and exit 1; code=%d out=%s", code, out)
	}
}

func TestSubcommandErrorExits(t *testing.T) {
	cases := []struct {
		name string
		fn   func() int
	}{
		{"dna-no-args", func() int { return dnaCmd(nil) }},
		{"dna-bogus", func() int { return dnaCmd([]string{"bogus"}) }},
		{"validate-missing-dna", func() int { return dnaValidate(nil) }},
		{"diff-missing-flags", func() int { return dnaDiff([]string{"--old", "x.json"}) }},
		{"run-missing-flags", func() int { return runCmd(nil) }},
		{"capture-missing-flags", func() int { return captureCmd(nil) }},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if code := tc.fn(); code != 2 {
				t.Errorf("expected exit 2, got %d", code)
			}
		})
	}
}
