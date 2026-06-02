package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// buildAlignctl builds the CLI once and returns (binPath, moduleRoot). The module root is
// two levels up from this package dir (cmd/alignctl), and runs use it as the working dir so
// the toy's relative paths (examples/toy/...) resolve.
func buildAlignctl(t *testing.T) (bin, root string) {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	root = filepath.Clean(filepath.Join(wd, "..", ".."))
	bin = filepath.Join(t.TempDir(), "alignctl")
	build := exec.Command("go", "build", "-o", bin, "./cmd/alignctl")
	build.Dir = root
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build alignctl: %v\n%s", err, out)
	}
	return bin, root
}

func runCLI(t *testing.T, bin, root string, args ...string) (string, int) {
	t.Helper()
	cmd := exec.Command(bin, args...)
	cmd.Dir = root
	out, err := cmd.CombinedOutput()
	if ee, ok := err.(*exec.ExitError); ok {
		return string(out), ee.ExitCode()
	}
	if err != nil {
		t.Fatalf("run %v: %v", args, err)
	}
	return string(out), 0
}

// TestCLIEndToEnd exercises the real binary against the toy: exit codes, gate behavior,
// and error handling — the wiring that unit tests of the pure helpers don't cover.
func TestCLIEndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping CLI integration test in -short mode")
	}
	bin, root := buildAlignctl(t)
	const (
		dnaFile   = "examples/toy/dna.json"
		runnerCmd = "go run ./examples/toy/cmd/toyrun"
		goldenDir = "examples/toy/golden"
	)

	t.Run("validate-ok", func(t *testing.T) {
		out, code := runCLI(t, bin, root, "dna", "validate", "--dna", dnaFile)
		if code != 0 || !strings.Contains(out, "6 genes valid") {
			t.Errorf("code=%d out=%s", code, out)
		}
	})
	t.Run("list", func(t *testing.T) {
		out, code := runCLI(t, bin, root, "dna", "list", "--dna", dnaFile)
		if code != 0 || !strings.Contains(out, "6 genes across 2 capabilities") {
			t.Errorf("code=%d out=%s", code, out)
		}
	})
	t.Run("run-no-gate", func(t *testing.T) {
		out, code := runCLI(t, bin, root, "run", "--dna", dnaFile, "--runner", runnerCmd, "--golden-dir", goldenDir)
		if code != 0 || !strings.Contains(out, "86.7%") || !strings.Contains(out, "Greet/padded-name") {
			t.Errorf("code=%d out=%s", code, out)
		}
	})
	t.Run("run-gate-unmet", func(t *testing.T) {
		_, code := runCLI(t, bin, root, "run", "--dna", dnaFile, "--runner", runnerCmd, "--golden-dir", goldenDir, "--gate-overall", "95")
		if code != 2 {
			t.Errorf("gate unmet should exit 2, got %d", code)
		}
	})
	t.Run("run-gate-met", func(t *testing.T) {
		_, code := runCLI(t, bin, root, "run", "--dna", dnaFile, "--runner", runnerCmd, "--golden-dir", goldenDir, "--gate-overall", "80", "--gate-critical", "100")
		if code != 0 {
			t.Errorf("gate met should exit 0, got %d", code)
		}
	})
	t.Run("diff-no-change", func(t *testing.T) {
		_, code := runCLI(t, bin, root, "dna", "diff", "--old", dnaFile, "--new", dnaFile)
		if code != 0 {
			t.Errorf("self-diff should exit 0, got %d", code)
		}
	})
	t.Run("missing-dna", func(t *testing.T) {
		_, code := runCLI(t, bin, root, "dna", "validate", "--dna", "does-not-exist.json")
		if code != 2 {
			t.Errorf("missing dna should exit 2, got %d", code)
		}
	})
	t.Run("unknown-subcommand", func(t *testing.T) {
		_, code := runCLI(t, bin, root, "frobnicate")
		if code != 2 {
			t.Errorf("unknown subcommand should exit 2, got %d", code)
		}
	})
}
