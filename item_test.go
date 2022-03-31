// Copyright 2022 validator Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validator

import (
	"reflect"
	"testing"
)

func TestItemVF(t *testing.T) {
	{
		i := Item{}
		vFs := i.vfs()
		expectLen := 12
		if len(vFs) != expectLen {
			t.Fatalf("test failed: vFs length expect [%d], but got [%d]\n", expectLen, len(vFs))
		}
		for _, vF := range vFs {
			if vF != nil {
				t.Fatal("test failed: every vF expect nil, but got not nil")
			}
		}
	}

	{
		i := Item{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "F", ""}
		vFs := i.vfs()
		expectLen := 12
		if len(vFs) != expectLen {
			t.Fatalf("test failed: vFs length expect [%d], but got [%d]\n", expectLen, len(vFs))
		}
		for i, vF := range vFs {
			if vF == nil && i != len(vFs)-1 {
				t.Fatal("test failed: every vF expect not nil, but got nil")
			}
		}
	}
}

func TestItemValidate(t *testing.T) {
	for _, tc := range []struct {
		name   string
		i      *Item
		value  reflect.Value
		passed bool
		msg    string
	}{
		// min
		{"minUnPass", &Item{Min: "2"}, reflect.ValueOf(1), false, ""},
		{"minUnPassAndMsg", &Item{Min: "2,fail"}, reflect.ValueOf(1), false, "fail"},
		{"minUnPassAndGlobalMsg", &Item{Min: "2", Msg: "fail"}, reflect.ValueOf(1), false, "fail"},
		{"minPass", &Item{Min: "1"}, reflect.ValueOf(1), true, ""},

		// max
		{"maxUnPass", &Item{Max: "0"}, reflect.ValueOf(1), false, ""},
		{"maxUnPassAndMsg", &Item{Max: "0,fail"}, reflect.ValueOf(1), false, "fail"},
		{"maxUnPassAndGlobalMsg", &Item{Max: "0", Msg: "fail"}, reflect.ValueOf(1), false, "fail"},
		{"maxPass", &Item{Max: "1"}, reflect.ValueOf(1), true, ""},

		// length
		{"lengthUnPass", &Item{Length: "0"}, reflect.ValueOf("1"), false, ""},
		{"lengthUnPassAndMsg", &Item{Length: "0,fail"}, reflect.ValueOf("1"), false, "fail"},
		{"lengthUnPassAndGlobalMsg", &Item{Length: "0", Msg: "fail"}, reflect.ValueOf("1"), false, "fail"},
		{"lengthPass", &Item{Length: "1"}, reflect.ValueOf("1"), true, ""},

		// arrLength
		{"arrLengthUnPass", &Item{ArrLength: "0"}, reflect.ValueOf(make([]interface{}, 1)), false, ""},
		{"arrLengthUnPassAndMsg", &Item{ArrLength: "0,fail"}, reflect.ValueOf(make([]interface{}, 1)), false, "fail"},
		{"arrLengthUnPassAndGlobalMsg", &Item{ArrLength: "0", Msg: "fail"}, reflect.ValueOf(make([]interface{}, 1)), false, "fail"},
		{"arrLengthPass", &Item{ArrLength: "1"}, reflect.ValueOf(make([]interface{}, 1)), true, ""},

		// minLength
		{"minLengthUnPass", &Item{MinLength: "2"}, reflect.ValueOf("1"), false, ""},
		{"minLengthUnPassAndMsg", &Item{MinLength: "2,fail"}, reflect.ValueOf("1"), false, "fail"},
		{"minLengthUnPassAndGlobalMsg", &Item{MinLength: "2", Msg: "fail"}, reflect.ValueOf("1"), false, "fail"},
		{"minLengthPass", &Item{MinLength: "1"}, reflect.ValueOf("1"), true, ""},

		// arrMinLength
		{"arrMinLengthUnPass", &Item{ArrMinLength: "1"}, reflect.ValueOf(make([]interface{}, 0)), false, ""},
		{"arrMinLengthUnPassAndMsg", &Item{ArrMinLength: "1,fail"}, reflect.ValueOf(make([]interface{}, 0)), false, "fail"},
		{"arrMinLengthUnPassAndGlobalMsg", &Item{ArrMinLength: "1", Msg: "fail"}, reflect.ValueOf(make([]interface{}, 0)), false, "fail"},
		{"arrMinLengthPass", &Item{ArrMinLength: "1"}, reflect.ValueOf(make([]interface{}, 1)), true, ""},

		// maxLength
		{"maxLengthUnPass", &Item{MaxLength: "0"}, reflect.ValueOf("1"), false, ""},
		{"maxLengthUnPassAndMsg", &Item{MaxLength: "0,fail"}, reflect.ValueOf("1"), false, "fail"},
		{"maxLengthUnPassAndGlobalMsg", &Item{MaxLength: "0", Msg: "fail"}, reflect.ValueOf("1"), false, "fail"},
		{"maxLengthPass", &Item{MaxLength: "1"}, reflect.ValueOf("1"), true, ""},

		// arrMaxLength
		{"arrMaxLengthUnPass", &Item{ArrMaxLength: "0"}, reflect.ValueOf(make([]interface{}, 1)), false, ""},
		{"arrMaxLengthUnPassAndMsg", &Item{ArrMaxLength: "0,fail"}, reflect.ValueOf(make([]interface{}, 1)), false, "fail"},
		{"arrMaxLengthUnPassAndGlobalMsg", &Item{ArrMaxLength: "0", Msg: "fail"}, reflect.ValueOf(make([]interface{}, 1)), false, "fail"},
		{"arrMaxLengthPass", &Item{ArrMaxLength: "1"}, reflect.ValueOf(make([]interface{}, 1)), true, ""},

		// enum
		{"enumUnPass", &Item{Enum: "0"}, reflect.ValueOf("1"), false, ""},
		{"enumUnPassAndMsg", &Item{Enum: "0,fail"}, reflect.ValueOf("1"), false, "fail"},
		{"enumUnPassAndGlobalMsg", &Item{Enum: "0", Msg: "fail"}, reflect.ValueOf("1"), false, "fail"},
		{"enumPass", &Item{Enum: "1"}, reflect.ValueOf("1"), true, ""},

		// regex
		{"regexUnPass", &Item{Regex: "0"}, reflect.ValueOf("1"), false, ""},
		{"regexUnPassAndMsg", &Item{Regex: "0,fail"}, reflect.ValueOf("1"), false, "fail"},
		{"regexUnPassAndGlobalMsg", &Item{Regex: "0", Msg: "fail"}, reflect.ValueOf("1"), false, "fail"},
		{"regexPass", &Item{Regex: "1"}, reflect.ValueOf("1"), true, ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			passed, msg := tc.i.Validate(reflect.StructField{Type: tc.value.Type()}, tc.value)
			if passed != tc.passed {
				t.Fatalf("%s test failed: expect passed [%v], but got [%v]\n", tc.name, tc.passed, passed)
			}
			if msg != tc.msg {
				t.Fatalf("%s test failed: expect msg [%v], but got [%v]\n", tc.name, tc.msg, msg)
			}
		})
	}
}
