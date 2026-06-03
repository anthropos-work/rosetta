// Command alignctl is the reference harness for the alignment test framework: it runs a
// mirror engine against a source (recorded golden, or live), scores how faithfully the
// mirror reproduces the source across an Alignment DNA, and reports divergences.
//
// alignctl is engine-agnostic: it never imports the engine under test. Instead it execs a
// pluggable "runner" command that speaks the outcomes protocol
// (see anthropos.dev/alignment/internal/outcome and the milestone spec-notes).
package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Fprint(os.Stderr, `alignctl - measure how faithfully a mirror engine reproduces a source engine.

usage:
  alignctl run      --dna P --runner CMD [--golden-dir D] [--source golden|live]
                    [--report out.json] [--gate-overall F] [--gate-critical F]
  alignctl capture  --dna P --runner CMD --golden-dir D
  alignctl dna list     --dna P [--json]
  alignctl dna diff     --old P --new P [--json]
  alignctl dna validate --dna P

A runner is any executable invoked as:  CMD --target {source|mirror} --dna P
and printing a JSON map of "<Capability>/<variant>" -> {"value":<json>,"error_class":<str|null>}.
`)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	var code int
	switch os.Args[1] {
	case "run":
		code = runCmd(os.Args[2:])
	case "capture":
		code = captureCmd(os.Args[2:])
	case "dna":
		code = dnaCmd(os.Args[2:])
	case "-h", "--help", "help":
		usage()
	default:
		fmt.Fprintf(os.Stderr, "alignctl: unknown subcommand %q\n\n", os.Args[1])
		usage()
		code = 2
	}
	os.Exit(code)
}
