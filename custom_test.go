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

func TestCustom(t *testing.T) {
	for _, tc := range []struct {
		name       string
		customName string
		vF         func(value reflect.Value) (bool, string)
		model      interface{}
		passed     bool
		msg        string
	}{
		{"pass", "pass", func(value reflect.Value) (bool, string) { return false, "" }, &struct {
			Name string `validate:"custom(pass)"`
		}{}, false, ""},

		{"unpass", "unpass", func(value reflect.Value) (bool, string) { return true, "" }, &struct {
			Name string `validate:"custom(unpass)"`
		}{}, true, ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			Custom(tc.customName, tc.vF)
			vv := New(tc.model).Validate()
			if tc.passed != vv.Passed {
				t.Fatalf("%s failed: passed expect [%v], but got [%v]\n", t.Name(), tc.passed, vv.Passed)
			}
			if tc.msg != vv.Messages() {
				t.Fatalf("%s failed: msg expect [%v], but got [%v]\n", t.Name(), tc.msg, vv.Messages())
			}
		})
	}
}
