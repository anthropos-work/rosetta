package runner

import (
	"encoding/json"
	"testing"

	"anthropos.dev/alignment/examples/toy/mirror"
	"anthropos.dev/alignment/examples/toy/source"
	"anthropos.dev/alignment/internal/dna"
)

func gene(capability, variant, input string) dna.Gene {
	return dna.Gene{
		ID: capability + "/" + variant, Capability: capability, Variant: variant,
		Operator: dna.OpExact, Input: json.RawMessage(input),
	}
}

func TestInvokeAdd(t *testing.T) {
	o, err := Invoke(source.Engine{}, gene("Add", "ok", `{"a":2,"b":3}`))
	if err != nil {
		t.Fatal(err)
	}
	if string(o.Value) != "5" || o.ErrorClass != nil {
		t.Errorf("got value=%s err=%v", o.Value, o.ErrorClass)
	}
}

func TestInvokeOverflowErrorClass(t *testing.T) {
	o, err := Invoke(source.Engine{}, gene("Add", "of", `{"a":9223372036854775807,"b":1}`))
	if err != nil {
		t.Fatal(err)
	}
	if o.ErrorClass == nil || *o.ErrorClass != "overflow" {
		t.Errorf("expected overflow error class, got value=%s err=%v", o.Value, o.ErrorClass)
	}
}

func TestInvokeGreetDiverges(t *testing.T) {
	g := gene("Greet", "padded", `{"name":"  a  b  "}`)
	so, _ := Invoke(source.Engine{}, g)
	mo, _ := Invoke(mirror.Engine{}, g)
	if string(so.Value) == string(mo.Value) {
		t.Errorf("source and mirror should diverge on a padded name; both=%s", so.Value)
	}
	if string(so.Value) != `"Hello, a b!"` {
		t.Errorf("source should collapse whitespace, got %s", so.Value)
	}
}

func TestInvokeUnknownCapability(t *testing.T) {
	if _, err := Invoke(source.Engine{}, gene("Nope", "v", `{}`)); err == nil {
		t.Error("expected an error for an unknown capability")
	}
}

func TestRun(t *testing.T) {
	d := &dna.DNA{SchemaVersion: 1, Capabilities: []dna.Capability{
		{ID: "Add", Criticality: dna.Critical, Variants: []dna.Variant{
			{ID: "ok", Operator: dna.OpExact, Input: json.RawMessage(`{"a":1,"b":1}`)},
		}},
		{ID: "Greet", Criticality: dna.Standard, Variants: []dna.Variant{
			{ID: "b", Operator: dna.OpExact, Input: json.RawMessage(`{"name":"x"}`)},
		}},
	}}
	set, err := Run(source.Engine{}, d)
	if err != nil {
		t.Fatal(err)
	}
	if len(set) != 2 || string(set["Add/ok"].Value) != "2" {
		t.Errorf("Run set wrong: %+v", set)
	}
}
