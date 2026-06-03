// Package report renders a compare.Report as human-readable text. The JSON form is
// just the compare.Report marshaled directly, so there's no separate JSON renderer here.
package report

import (
	"fmt"
	"io"
	"sort"

	"anthropos.dev/alignment/internal/compare"
)

// Human writes a readable alignment summary: scores, per-capability rollup, and the
// divergence list (the part a developer acts on).
func Human(w io.Writer, r *compare.Report) {
	fmt.Fprintf(w, "Alignment: mirror %s@%s  vs  source %s@%s\n",
		r.Mirror.Name, r.Mirror.Version, r.Source.Name, r.Source.Version)
	fmt.Fprintf(w, "Score: overall %.1f%%   critical %.1f%%   (%d/%d genes aligned)\n",
		r.Overall, r.Critical, r.AlignedGenes, r.TotalGenes)

	fmt.Fprintln(w, "\nPer capability:")
	caps := make([]string, 0, len(r.PerCapability))
	for c := range r.PerCapability {
		caps = append(caps, c)
	}
	sort.Strings(caps)
	for _, c := range caps {
		roll := r.PerCapability[c]
		mark := "ok"
		if roll.Aligned < roll.Total {
			mark = "DIVERGED"
		}
		fmt.Fprintf(w, "  %-28s %d/%d  %s\n", c, roll.Aligned, roll.Total, mark)
	}

	var diverged []compare.GeneResult
	for _, g := range r.Genes {
		if !g.Aligned {
			diverged = append(diverged, g)
		}
	}
	if len(diverged) == 0 {
		fmt.Fprintln(w, "\nNo divergences -- the mirror is indistinguishable across the DNA.")
		return
	}
	fmt.Fprintf(w, "\nDivergences (%d):\n", len(diverged))
	for _, g := range diverged {
		crit := ""
		if g.Critical {
			crit = "  [CRITICAL]"
		}
		fmt.Fprintf(w, "  FAIL %s  (%s, w%d)%s\n       %s\n", g.GeneID, g.Operator, g.Weight, crit, g.Detail)
	}
}
