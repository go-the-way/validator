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
	"github.com/billcoding/reflectx"
	"reflect"
	"strconv"
)

// validFunc struct
type validFunc struct {
	valid bool
	msg   string
}

// ValidFunc method
func ValidFunc(str string) VFunc {
	if str == "" {
		return nil
	}
	vStr, msg := str, ""
	if spIdx := findSpIdx(str); spIdx != -1 {
		vStr = str[:spIdx]
		msg = str[spIdx+1:]
	}
	if v, err := strconv.ParseBool(vStr); err != nil {
		panic(err)
	} else {
		return &validFunc{v, msg}
	}
}

// Valid method
func (f *validFunc) Valid(value reflect.Value) (bool, string) {
	if !f.valid {
		return true, ""
	}

	if reflectx.IsArray(value.Type()) || reflectx.IsSlice(value.Type()) {
		for i := 0; i < value.Len(); i++ {
			return f.Valid(value.Index(i))
		}
	}

	if reflectx.IsPtr(value.Type()) && value.IsNil() {
		return false, f.msg
	}

	return true, ""
}
