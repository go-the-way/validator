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

func TestNewValidFunc(t *testing.T) {
	if v := ValidFunc(""); v != nil {
		t.Fatal("test failed")
	}
}

func TestValidPanic(t *testing.T) {
	for _, tc := range []struct {
		name  string
		str   string
		value reflect.Value
		err   error
	}{
		{"ParseBool", "A", reflect.ValueOf(true), parseNumError("ParseBool", "A")},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				re := recover()
				if re == nil {
					t.Fatalf("%s failed: error expect not nil", t.Name())
				}
				if re == nil || !reflect.DeepEqual(tc.err, re) {
					t.Fatalf("%s failed: error expect [%v], but got [%v]\n", t.Name(), tc.err, re)
				}
			}()
			vF := ValidFunc(tc.str)
			_, _ = vF.Valid(tc.value)
		})
	}
}

func TestValid(t *testing.T) {
	for _, s := range []testCase{
		{"False", ValidFunc("F"), reflect.ValueOf(nil), true, ""},

		{"StructArr", ValidFunc("T"), reflect.ValueOf([2]struct{}{}), true, ""},
		{"StructSlice", ValidFunc("T"), reflect.ValueOf([]struct{}{}), true, ""},

		{"StructPtrArr", ValidFunc("T"), reflect.ValueOf([2]*struct{}{}), false, ""},
		{"StructPtrArr", ValidFunc("T,fail"), reflect.ValueOf([2]*struct{}{}), false, "fail"},
	} {
		t.Run(s.name, s.test)
	}
}
