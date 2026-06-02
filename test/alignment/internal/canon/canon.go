// Package canon produces a canonical, precision-preserving JSON form: object keys sorted,
// whitespace normalized, integers exact (UseNumber — so large IDs/timestamps don't lose
// precision via float64). It is shared by the comparator (value equality) and the DNA
// differ (drift detection) so both agree on what "the same JSON" means.
package canon

import (
	"bytes"
	"encoding/json"
)

// Decode parses raw into a generic value with integer precision preserved. Empty/nil → null.
func Decode(raw json.RawMessage) (any, error) {
	dec := json.NewDecoder(bytes.NewReader(nz(raw)))
	dec.UseNumber()
	var v any
	if err := dec.Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}

// Bytes returns the canonical JSON encoding of raw (sorted keys, exact integers).
func Bytes(raw json.RawMessage) ([]byte, error) {
	v, err := Decode(raw)
	if err != nil {
		return nil, err
	}
	return json.Marshal(v)
}

// String returns Bytes as a string, or the raw bytes unchanged if they don't parse (so a
// cheap equality check treats unparseable input as-is rather than erroring). Empty → "null".
func String(raw json.RawMessage) string {
	if len(raw) == 0 {
		return "null"
	}
	b, err := Bytes(raw)
	if err != nil {
		return string(raw)
	}
	return string(b)
}

func nz(r json.RawMessage) []byte {
	if len(r) == 0 {
		return []byte("null")
	}
	return r
}
