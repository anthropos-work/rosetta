package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// chdirRoot changes into the module root for the test so the toy's relative paths and the
// "go run ./examples/toy/cmd/toyrun" runner resolve. Restored on cleanup. These tests call
// the subcommands IN-PROCESS (so run.go/capture.go are coverage-instrumented) against a TEMP
// golden dir — they never touch the committed goldens.
func chdirRoot(t *testing.T) {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(filepath.Clean(filepath.Join(wd, "..", ".."))); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(wd) })
}

func TestRunAndCaptureInProcess(t *testing.T) {
	if testing.Short() {
		t.Skip("skips go-run subprocess in -short")
	}
	chdirRoot(t)
	golden := t.TempDir()
	const (
		dnaFile = "examples/toy/dna.json"
		runner  = "go run ./examples/toy/cmd/toyrun"
	)

	if _, code := captureStdout(t, func() int {
		return captureCmd([]string{"--dna", dnaFile, "--runner", runner, "--golden-dir", golden})
	}); code != 0 {
		t.Fatalf("capture exit=%d", code)
	}
	if n := countFiles(t, golden); n != 6 {
		t.Errorf("expected 6 goldens captured, got %d", n)
	}

	t.Run("golden", func(t *testing.T) {
		out, code := captureStdout(t, func() int {
			return runCmd([]string{"--dna", dnaFile, "--runner", runner, "--golden-dir", golden})
		})
		if code != 0 || !strings.Contains(out, "86.7%") || !strings.Contains(out, "Greet/padded-name") {
			t.Errorf("code=%d out=%s", code, out)
		}
	})
	t.Run("source-live", func(t *testing.T) {
		if _, code := captureStdout(t, func() int {
			return runCmd([]string{"--dna", dnaFile, "--runner", runner, "--source", "live"})
		}); code != 0 {
			t.Errorf("--source live exit=%d", code)
		}
	})
	t.Run("gate-unmet", func(t *testing.T) {
		if _, code := captureStdout(t, func() int {
			return runCmd([]string{"--dna", dnaFile, "--runner", runner, "--golden-dir", golden, "--gate-overall", "95"})
		}); code != 2 {
			t.Errorf("gate unmet should exit 2, got %d", code)
		}
	})
	t.Run("bad-source", func(t *testing.T) {
		if _, code := captureStdout(t, func() int {
			return runCmd([]string{"--dna", dnaFile, "--runner", runner, "--source", "bogus"})
		}); code != 2 {
			t.Errorf("bogus --source should exit 2, got %d", code)
		}
	})
	t.Run("report-file", func(t *testing.T) {
		rep := filepath.Join(t.TempDir(), "rep.json")
		if _, code := captureStdout(t, func() int {
			return runCmd([]string{"--dna", dnaFile, "--runner", runner, "--golden-dir", golden, "--report", rep})
		}); code != 0 {
			t.Fatalf("--report exit=%d", code)
		}
		b, err := os.ReadFile(rep)
		if err != nil || !strings.Contains(string(b), `"overall_score"`) {
			t.Errorf("report file missing or not JSON: err=%v", err)
		}
	})
}

func countFiles(t *testing.T, dir string) int {
	t.Helper()
	n := 0
	if err := filepath.Walk(dir, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			n++
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	return n
}
