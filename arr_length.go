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

// arrLengthFunc struct
type arrLengthFunc struct {
	length int
	msg    string
}

// ArrLengthFunc method
func ArrLengthFunc(str string) VFunc {
	if str == "" {
		return nil
	}
	vStr, msg := str, ""
	if spIdx := findSpIdx(str); spIdx != -1 {
		vStr = str[:spIdx]
		msg = str[spIdx+1:]
	}
	if v, err := strconv.ParseInt(vStr, 10, 64); err != nil {
		panic(err)
	} else {
		return &arrLengthFunc{int(v), msg}
	}
}

// Valid method
func (f *arrLengthFunc) Valid(value reflect.Value) (bool, string) {
	passed, msg := true, f.msg
	typ := value.Type()
	switch {
	case reflectx.IsPtr(typ):
		return f.Valid(value.Elem())
	case reflectx.IsArray(typ) || reflectx.IsSlice(typ):
		passed = f.length == value.Len()
	}
	return passed, msg
}
