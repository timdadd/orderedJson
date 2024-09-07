package orderedjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type jsonContext int

const (
	none jsonContext = iota
	key
	value
)

type OrderedJson struct {
	K string
	V interface{}
}

// Decoder extends Go encoding/json.Decoder.
type Decoder struct {
	json.Decoder
}

// NewDecoder creates a new instance of the extended JSON Decoder.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{Decoder: *json.NewDecoder(r)}
}

// Unmarshal is vaguely equivalent to encoding/Unmarshal
func Unmarshal(data []byte) ([]*OrderedJson, error) {
	if !json.Valid(data) {
		return nil, fmt.Errorf("json is not well formatted")
	}
	d := NewDecoder(bytes.NewReader(data))
	return d.Decode()
}

// Decode is vaguely equivalent to encoding/json.Decode() it takes an initiated decoder of io.reader
func (d *Decoder) Decode() ([]*OrderedJson, error) {
	var iv interface{}
	var err error
	if iv, _, err = d.token(none); err != nil {
		return nil, err
	}
	switch v := iv.(type) {
	case []*OrderedJson:
		return v, nil
	case []interface{}:
		return []*OrderedJson{}, nil

	}
	return nil, fmt.Errorf("json: cannot unmarshal object into ordered json")
}

// token handles the next token in the json message
// Returns the orderedJson, the last token, and any error
func (d *Decoder) token(ctx jsonContext) (interface{}, json.Token, error) {
	var t json.Token
	var err error
	var v []*OrderedJson
	for {
		if t, err = d.Decoder.Token(); err != nil {
			if err == io.EOF {
				return v, t, nil
			} else {
				return v, t, err
			}
		}

		// Analyse the token
		switch tt := t.(type) {
		case json.Delim:
			switch t {
			case json.Delim('{'):
				ctx = key // We now expect a key
			case json.Delim('}'), json.Delim(']'): // This is either closing an object or an array
				return v, t, nil
			case json.Delim('['): // Starting an array
				array := make([]interface{}, 0)
				var lastToken json.Token
				// Loop until we get array close token
				for lastToken != json.Delim(']') {
					var i interface{}
					if i, lastToken, err = d.token(ctx); err != nil {
						return i, t, err
					}
					if lastToken == json.Delim(']') {
						break
					}
					array = append(array, i)
				}
				return array, t, nil
			}
		case float64, json.Number, bool:
			switch ctx {
			case value:
				return tt, t, nil
			default:
				return nil, nil, fmt.Errorf("value %f received when context is %v", tt, ctx)
			}
		case string:
			switch ctx {
			case key:
				ctx = value // Now we expect a value
				oj := &OrderedJson{K: tt}
				if oj.V, _, err = d.token(ctx); err != nil {
					return nil, nil, fmt.Errorf("could not determine value for key %s", tt)
				}
				v = append(v, oj)
				ctx = key
			case value:
				return tt, t, nil
			default:
				return nil, nil, fmt.Errorf("value %s received when context is %v", tt, ctx)
			}
		}
	}
}
