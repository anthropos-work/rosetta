// Command toyrun is the toy engine's alignment runner: it speaks the outcomes protocol
// alignctl expects. Invoked as `toyrun --target {source|mirror} --dna PATH`, it prints a
// JSON map of gene id -> {value, error_class} to stdout. A real mirror (Clerkenstein in M1)
// ships its own equivalent runner.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"anthropos.dev/alignment/examples/toy/mirror"
	"anthropos.dev/alignment/examples/toy/runner"
	"anthropos.dev/alignment/examples/toy/source"
	"anthropos.dev/alignment/examples/toy/surface"
	"anthropos.dev/alignment/internal/dna"
)

func main() {
	target := flag.String("target", "", "source|mirror")
	dnaPath := flag.String("dna", "", "path to DNA JSON")
	flag.Parse()

	var eng surface.Surface
	switch *target {
	case "source":
		eng = source.Engine{}
	case "mirror":
		eng = mirror.Engine{}
	default:
		fmt.Fprintln(os.Stderr, "toyrun: --target must be source|mirror")
		os.Exit(2)
	}
	d, err := dna.Load(*dnaPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "toyrun:", err)
		os.Exit(2)
	}
	set, err := runner.Run(eng, d)
	if err != nil {
		fmt.Fprintln(os.Stderr, "toyrun:", err)
		os.Exit(2)
	}
	b, err := json.MarshalIndent(set, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "toyrun:", err)
		os.Exit(2)
	}
	fmt.Println(string(b))
}
