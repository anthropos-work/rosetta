// Package toy is a self-contained reference example for the alignment framework: a source
// engine, a mirror with exactly one intentional divergence (Greet/padded-name), the DNA, a
// runner, and build-tagged alignment tests. It exists so the framework can be proven
// end-to-end — including that it *catches* misalignment — without any external dependency.
//
// Run it two ways (they share the same compare core and agree):
//
//	go test -tags alignment ./examples/toy/...                          # developer ergonomic
//	go run ./cmd/alignctl run --dna examples/toy/dna.json \             # engine-agnostic
//	    --runner "go run ./examples/toy/cmd/toyrun" --golden-dir examples/toy/golden
package toy
