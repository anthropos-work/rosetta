package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"anthropos.dev/alignment/internal/compare"
	"anthropos.dev/alignment/internal/dna"
	"anthropos.dev/alignment/internal/outcome"
	"anthropos.dev/alignment/internal/report"
)

func runCmd(args []string) int {
	fs := flag.NewFlagSet("run", flag.ContinueOnError)
	dnaPath := fs.String("dna", "", "path to DNA JSON")
	runner := fs.String("runner", "", "runner command (space-separated)")
	goldenDir := fs.String("golden-dir", "", "golden dir (required when --source=golden)")
	source := fs.String("source", "golden", "source outcomes: golden|live")
	reportPath := fs.String("report", "", "also write the JSON report to this path")
	gateOverall := fs.Float64("gate-overall", 0, "minimum overall score to exit 0")
	gateCritical := fs.Float64("gate-critical", 0, "minimum critical score to exit 0")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if *dnaPath == "" || *runner == "" {
		fmt.Fprintln(os.Stderr, "alignctl run: --dna and --runner are required")
		return 2
	}
	d, err := dna.Load(*dnaPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "alignctl run:", err)
		return 2
	}
	if errs := d.Validate(); len(errs) > 0 {
		fmt.Fprintln(os.Stderr, "alignctl run: invalid DNA:")
		for _, e := range errs {
			fmt.Fprintln(os.Stderr, "  -", e)
		}
		return 2
	}

	mirror, err := runRunner(*runner, "mirror", *dnaPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "alignctl run:", err)
		return 2
	}
	var src outcome.Set
	switch *source {
	case "live":
		if src, err = runRunner(*runner, "source", *dnaPath); err != nil {
			fmt.Fprintln(os.Stderr, "alignctl run:", err)
			return 2
		}
	case "golden":
		if *goldenDir == "" {
			fmt.Fprintln(os.Stderr, "alignctl run: --golden-dir is required when --source=golden")
			return 2
		}
		if src, err = outcome.LoadGoldenSet(*goldenDir, d.GeneIDs()); err != nil {
			fmt.Fprintln(os.Stderr, "alignctl run:", err)
			return 2
		}
	default:
		fmt.Fprintf(os.Stderr, "alignctl run: --source must be golden|live, got %q\n", *source)
		return 2
	}

	rep := compare.Evaluate(d, src, mirror)
	report.Human(os.Stdout, rep)
	if *reportPath != "" {
		if err := writeJSON(*reportPath, rep); err != nil {
			fmt.Fprintln(os.Stderr, "alignctl run: write report:", err)
			return 2
		}
	}
	if !rep.GateMet(*gateOverall, *gateCritical) {
		fmt.Fprintf(os.Stderr, "\ngate not met (overall %.1f%% need %.1f%% | critical %.1f%% need %.1f%%)\n",
			rep.Overall, *gateOverall, rep.Critical, *gateCritical)
		return 2
	}
	return 0
}

// runRunner execs the pluggable runner for one target and parses its outcomes JSON.
// The runner is invoked as: CMD [extra args...] --target {source|mirror} --dna PATH
func runRunner(runner, target, dnaPath string) (outcome.Set, error) {
	parts := strings.Fields(runner)
	if len(parts) == 0 {
		return nil, fmt.Errorf("empty --runner")
	}
	cmdArgs := append([]string{}, parts[1:]...)
	cmdArgs = append(cmdArgs, "--target", target, "--dna", dnaPath)
	cmd := exec.Command(parts[0], cmdArgs...)
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("runner %q --target %s: %w", runner, target, err)
	}
	set, err := outcome.ParseSet(out)
	if err != nil {
		return nil, fmt.Errorf("runner %q --target %s: %w", runner, target, err)
	}
	return set, nil
}

func writeJSON(path string, v any) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(b, '\n'), 0o644)
}
