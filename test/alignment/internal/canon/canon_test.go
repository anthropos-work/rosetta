package canon

import (
	"encoding/json"
	"testing"
)

func TestStringCanonicalizes(t *testing.T) {
	// key order + whitespace normalize to the same string
	if String(json.RawMessage(`{"b":2,"a":1}`)) != String(json.RawMessage(`{ "a":1, "b":2 }`)) {
		t.Error("key order / whitespace should canonicalize the same")
	}
	// large integer preserved exactly (no float rounding)
	big := `9223372036854775807`
	if got := String(json.RawMessage(big)); got != big {
		t.Errorf("lost integer precision: %s", got)
	}
	if String(nil) != "null" {
		t.Error("empty → null")
	}
	if String(json.RawMessage(`{bad`)) != `{bad` {
		t.Error("unparseable input should pass through unchanged")
	}
}

func TestBytesAndDecode(t *testing.T) {
	b, err := Bytes(json.RawMessage(`{"b":2,"a":1}`))
	if err != nil || string(b) != `{"a":1,"b":2}` {
		t.Errorf("Bytes = %s err=%v, want {\"a\":1,\"b\":2}", b, err)
	}
	if _, err := Bytes(json.RawMessage(`{bad`)); err == nil {
		t.Error("Bytes should error on invalid JSON")
	}
	v, err := Decode(json.RawMessage(`123`))
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := v.(json.Number); !ok {
		t.Errorf("Decode should preserve numbers as json.Number, got %T", v)
	}
}
