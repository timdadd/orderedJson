package orderedjson_test

import (
	"bytes"
	"fmt"
	oj "orderedJson"
	"reflect"
	"testing"
)

// Text examples are (c) 2015 Exponent Labs LLC

var decoderTests = []struct {
	in  string
	err error
	out []*oj.OrderedJson
}{
	{in: `{}`, out: []*oj.OrderedJson{}},
	{in: `{"a":"a", "b":"b", "c":"c"}`,
		out: []*oj.OrderedJson{
			{"a", "a"},
			{"b", "b"},
			{"c", "c"},
		},
	},
	{in: `{"a":{"b":{"c":14}}}`,
		out: []*oj.OrderedJson{
			{"a", []*oj.OrderedJson{
				{"b", []*oj.OrderedJson{
					{"c", float64(14)}}}}},
		},
	},
	{in: `{"a":{"d":{"b":{"c":3}}, "b":{"c":14}}}`,
		out: []*oj.OrderedJson{
			{"a", []*oj.OrderedJson{
				{"d", []*oj.OrderedJson{
					{"b", []*oj.OrderedJson{
						{"c", float64(3)}}}}},
				{"b", []*oj.OrderedJson{{"c", float64(14)}}}}},
		},
	},
	{in: `{"a":{"b":{"c":14},"b2":3}}`,
		out: []*oj.OrderedJson{
			{"a", []*oj.OrderedJson{
				{"b", []*oj.OrderedJson{
					{"c", float64(14)}}},
				{"b2", float64(3)}}}},
	},
	{in: `[]`, out: []*oj.OrderedJson{}},
}

func TestDecoder(t *testing.T) {
	var testDesc string
	var v []*oj.OrderedJson
	var err error

	for ti, tst := range decoderTests {
		d := oj.NewDecoder(bytes.NewBuffer([]byte(tst.in)))
		testDesc = fmt.Sprintf("#%d '%s'", ti, tst.in)
		v, err = d.Decode()

		if !reflect.DeepEqual(err, tst.err) {
			t.Errorf("#%v unexpected error: '%v' expecting '%v' : %v", ti, err, tst.err, testDesc)
		}

		if !reflect.DeepEqual(v, tst.out) {
			t.Errorf("#%v decode: expected %#v, was %#v : %v", ti, tst.out, v, testDesc)
		}
	}
}

//
//{in: `[0,1,2]`, path: []interface{}{2}, match: true, out: float64(2), err: nil},
//{in: `[0,1 , 2]`, path: []interface{}{2}, match: true, out: float64(2), err: nil},
//{in: `[0,{"b":1},2]`, path: []interface{}{1}, match: true, out: map[string]interface{}{"b": float64(1)}, err: nil},
//{in: `[1,{"b":1},2]`, path: []interface{}{2}, match: true, out: float64(2), err: nil},
//{in: `[1,{"b":[1]},2]`, path: []interface{}{2}, match: true, out: float64(2), err: nil},
//{in: `[1,[{"b":[1]},3],2]`, path: []interface{}{2}, match: true, out: float64(2), err: nil},
//
//{in: `[1,[{"b":[1]},3],4]`, path: []interface{}{1, 1}, match: true, out: float64(3), err: nil},
//{in: `[1,[{"b":[1]},3],2]`, path: []interface{}{1, 0, "b", 0}, match: true, out: float64(1), err: nil},
//{in: `[1,[{"b":[1]},3],5]`, path: []interface{}{2}, match: true, out: float64(5), err: nil},
//{in: `{"b":[{"a":0},{"a":1}]}`, path: []interface{}{"b", 0, "a"}, match: true, out: float64(0), err: nil},
//{in: `{"b":[{"a":0},{"a":1}]}`, path: []interface{}{"b", 1, "a"}, match: true, out: float64(1), err: nil},
//{in: `{"a":"b","b":"z","z":"s"}`, path: []interface{}{"b"}, match: true, out: "z", err: nil},
//{in: `{"a":"b","b":"z","l":0,"z":"s"}`, path: []interface{}{"z"}, match: true, out: "s", err: nil},
