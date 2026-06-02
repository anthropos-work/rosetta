package main

import (
	"flag"
	"fmt"
	"os"

	"anthropos.dev/alignment/internal/dna"
	"anthropos.dev/alignment/internal/outcome"
)

// captureCmd records the source target's outcomes as goldens (run once, commit, replay
// forever) so a live-SaaS source can be measured reproducibly offline.
func captureCmd(args []string) int {
	fs := flag.NewFlagSet("capture", flag.ContinueOnError)
	dnaPath := fs.String("dna", "", "path to DNA JSON")
	runner := fs.String("runner", "", "runner command (space-separated)")
	goldenDir := fs.String("golden-dir", "", "golden output dir")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if *dnaPath == "" || *runner == "" || *goldenDir == "" {
		fmt.Fprintln(os.Stderr, "alignctl capture: --dna, --runner and --golden-dir are required")
		return 2
	}
	d, err := dna.Load(*dnaPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "alignctl capture:", err)
		return 2
	}
	if errs := d.Validate(); len(errs) > 0 {
		fmt.Fprintln(os.Stderr, "alignctl capture: invalid DNA:")
		for _, e := range errs {
			fmt.Fprintln(os.Stderr, "  -", e)
		}
		return 2
	}
	src, err := runRunner(*runner, "source", *dnaPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "alignctl capture:", err)
		return 2
	}
	written, missing := 0, 0
	for _, id := range d.GeneIDs() {
		o, ok := src[id]
		if !ok {
			fmt.Fprintf(os.Stderr, "  warning: source omitted gene %q\n", id)
			missing++
			continue
		}
		if err := outcome.WriteGolden(*goldenDir, id, o); err != nil {
			fmt.Fprintln(os.Stderr, "alignctl capture:", err)
			return 2
		}
		written++
	}
	fmt.Printf("captured %d goldens to %s (%d missing)\n", written, *goldenDir, missing)
	if missing > 0 {
		return 1
	}
	return 0
}
