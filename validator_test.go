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

type testCase struct {
	name string
	VFunc
	reflect.Value
	accept bool
	passed bool
	msg    string
}

func (tc *testCase) test(t *testing.T) {
	if tc.VFunc == nil {
		t.Fatalf("%s failed: tc.VFunc is nil\n", t.Name())
	}
	if accept := tc.VFunc.Accept(tc.Value.Type()); accept != tc.accept {
		t.Fatalf("%s failed: accept expect [%v], but got [%v]\n", t.Name(), tc.accept, accept)
	}
	if tc.accept {
		passed, msg := tc.VFunc.Pass(tc.Value)
		if passed != tc.passed {
			t.Fatalf("%s failed: passed expect [%v], but got [%v]\n", t.Name(), tc.passed, passed)
		}
		if msg != tc.msg {
			t.Fatalf("%s failed: msg expect [%v], but got [%v]\n", t.Name(), tc.msg, msg)
		}
	}
}

func TestNew(t *testing.T) {
	for _, tc := range []struct {
		name      string
		structPtr interface{}
		expectLen int
		passed    bool
		msg       string
	}{
		{"NoField", &struct{}{}, 0, true, ""},

		{"IntUnPass", &struct {
			int `validate:"min(1)"`
		}{}, 1, false, ""},
		{"IntUnPassAndMsg", &struct {
			int `validate:"min(1,fail)"`
		}{}, 1, false, "fail"},
		{"IntUnPassAndGlobalMsg", &struct {
			int `validate:"min(1) msg(fail)"`
		}{}, 1, false, "fail"},
		{"IntPass", &struct {
			int `validate:"min(0)"`
		}{}, 1, true, ""},

		{"StringUnPass", &struct {
			string `validate:"minlength(1)"`
		}{}, 1, false, ""},
		{"StringUnPassAndMsg", &struct {
			string `validate:"minlength(1,fail)"`
		}{}, 1, false, "fail"},
		{"StringUnPassAndGlobalMsg", &struct {
			string `validate:"minlength(1) msg(fail)"`
		}{}, 1, false, "fail"},
		{"StringPass", &struct {
			string `validate:"minlength(0)"`
		}{}, 1, true, ""},

		{"FloatUnPass", &struct {
			float64 `validate:"min(1)"`
		}{}, 1, false, ""},
		{"FloatUnPassAndMsg", &struct {
			float64 `validate:"min(1,fail)"`
		}{}, 1, false, "fail"},
		{"FloatUnPassAndGlobalMsg", &struct {
			float64 `validate:"min(1) msg(fail)"`
		}{}, 1, false, "fail"},
		{"FloatPass", &struct {
			float64 `validate:"min(0)"`
		}{}, 1, true, ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			v := New(tc.structPtr)
			if v == nil {
				t.Fatalf("%s test failed: validator is nil\n", tc.name)
			}
			if tc.expectLen != len(v.items) {
				t.Fatalf("%s test failed: expect len [%d], but got [%d]\n", tc.name, tc.expectLen, len(v.items))
			}
			r := v.Validate()
			if tc.passed != r.Passed {
				t.Fatalf("%s test failed: expect passed [%v], but got [%v]\n", tc.name, tc.passed, r.Passed)
			}
			if tc.msg != r.Messages() {
				t.Fatalf("%s test failed: expect msg [%s], but got [%s]\n", tc.name, tc.msg, r.Messages())
			}
		})
	}
}
