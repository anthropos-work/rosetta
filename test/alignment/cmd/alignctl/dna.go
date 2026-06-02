package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"

	"anthropos.dev/alignment/internal/dna"
)

func dnaCmd(args []string) int {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "alignctl dna: expected list|diff|validate")
		return 2
	}
	switch args[0] {
	case "list":
		return dnaList(args[1:])
	case "diff":
		return dnaDiff(args[1:])
	case "validate":
		return dnaValidate(args[1:])
	default:
		fmt.Fprintf(os.Stderr, "alignctl dna: unknown %q (want list|diff|validate)\n", args[0])
		return 2
	}
}

func dnaList(args []string) int {
	fs := flag.NewFlagSet("dna list", flag.ContinueOnError)
	dnaPath := fs.String("dna", "", "path to DNA JSON")
	asJSON := fs.Bool("json", false, "emit JSON")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	d, err := loadValid(*dnaPath)
	if err != nil {
		return 2
	}
	genes := d.Genes()
	if *asJSON {
		type row struct {
			ID          string `json:"id"`
			Capability  string `json:"capability"`
			Operator    string `json:"operator"`
			Criticality string `json:"criticality"`
			Weight      int    `json:"weight"`
		}
		rows := make([]row, 0, len(genes))
		for _, g := range genes {
			rows = append(rows, row{g.ID, g.Capability, string(g.Operator), string(g.Criticality), g.Weight})
		}
		b, _ := json.MarshalIndent(rows, "", "  ")
		fmt.Println(string(b))
		return 0
	}
	fmt.Printf("%d genes across %d capabilities:\n", len(genes), len(d.Capabilities))
	for _, g := range genes {
		fmt.Printf("  %-36s %-12s w%d  %s\n", g.ID, g.Operator, g.Weight, g.Criticality)
	}
	return 0
}

func dnaValidate(args []string) int {
	fs := flag.NewFlagSet("dna validate", flag.ContinueOnError)
	dnaPath := fs.String("dna", "", "path to DNA JSON")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if *dnaPath == "" {
		fmt.Fprintln(os.Stderr, "alignctl dna validate: --dna required")
		return 2
	}
	d, err := dna.Load(*dnaPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}
	if errs := d.Validate(); len(errs) > 0 {
		for _, e := range errs {
			fmt.Fprintln(os.Stderr, "  -", e)
		}
		return 2
	}
	fmt.Printf("OK: %d genes valid\n", len(d.Genes()))
	return 0
}

// dnaDiff reports added / removed / changed genes between two DNA versions. Exit 1 when
// the DNA moved (so M1b drift detection / CI can branch on it).
func dnaDiff(args []string) int {
	fs := flag.NewFlagSet("dna diff", flag.ContinueOnError)
	oldPath := fs.String("old", "", "old DNA JSON")
	newPath := fs.String("new", "", "new DNA JSON")
	asJSON := fs.Bool("json", false, "emit JSON")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if *oldPath == "" || *newPath == "" {
		fmt.Fprintln(os.Stderr, "alignctl dna diff: --old and --new required")
		return 2
	}
	oldD, err := loadValid(*oldPath)
	if err != nil {
		return 2
	}
	newD, err := loadValid(*newPath)
	if err != nil {
		return 2
	}
	oldG, newG := geneMap(oldD), geneMap(newD)
	var added, removed, changed []string
	for id := range newG {
		if _, ok := oldG[id]; !ok {
			added = append(added, id)
		}
	}
	for id, og := range oldG {
		ng, ok := newG[id]
		if !ok {
			removed = append(removed, id)
		} else if geneChanged(og, ng) {
			changed = append(changed, id)
		}
	}
	sort.Strings(added)
	sort.Strings(removed)
	sort.Strings(changed)

	if *asJSON {
		b, _ := json.MarshalIndent(map[string][]string{
			"added": added, "removed": removed, "changed": changed,
		}, "", "  ")
		fmt.Println(string(b))
	} else {
		fmt.Printf("DNA diff %s -> %s\n", *oldPath, *newPath)
		printList("added", added)
		printList("removed", removed)
		printList("changed", changed)
	}
	if len(added)+len(removed)+len(changed) > 0 {
		return 1
	}
	return 0
}

func printList(label string, ids []string) {
	if len(ids) == 0 {
		fmt.Printf("  %s: none\n", label)
		return
	}
	fmt.Printf("  %s (%d):\n", label, len(ids))
	for _, id := range ids {
		fmt.Printf("    %s\n", id)
	}
}

func geneMap(d *dna.DNA) map[string]dna.Gene {
	m := map[string]dna.Gene{}
	for _, g := range d.Genes() {
		m[g.ID] = g
	}
	return m
}

func geneChanged(a, b dna.Gene) bool {
	if a.Operator != b.Operator || a.Weight != b.Weight || a.Criticality != b.Criticality {
		return true
	}
	// Compare inputs canonically so reformatting / key reordering / integer formatting
	// don't register as spurious drift (M1b must only fire on real contract changes).
	if canonJSON(a.Input) != canonJSON(b.Input) {
		return true
	}
	return joinNorm(a.Normalize) != joinNorm(b.Normalize)
}

// canonJSON returns a canonical, precision-preserving string form of a JSON value
// (sorted keys, normalized whitespace, exact integers), or the raw bytes if unparseable.
func canonJSON(raw json.RawMessage) string {
	if len(raw) == 0 {
		return "null"
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	var v any
	if err := dec.Decode(&v); err != nil {
		return string(raw)
	}
	b, err := json.Marshal(v)
	if err != nil {
		return string(raw)
	}
	return string(b)
}

func joinNorm(s []string) string {
	cp := append([]string{}, s...)
	sort.Strings(cp)
	return fmt.Sprint(cp)
}

func loadValid(path string) (*dna.DNA, error) {
	if path == "" {
		fmt.Fprintln(os.Stderr, "missing --dna")
		return nil, fmt.Errorf("missing --dna")
	}
	d, err := dna.Load(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}
	if errs := d.Validate(); len(errs) > 0 {
		for _, e := range errs {
			fmt.Fprintln(os.Stderr, "  -", e)
		}
		return nil, fmt.Errorf("invalid DNA")
	}
	return d, nil
}
