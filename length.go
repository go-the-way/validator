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

// lengthFunc struct
type lengthFunc struct {
	length int
	msg    string
}

// LengthFunc method
func LengthFunc(str string) VFunc {
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
		return &lengthFunc{int(v), msg}
	}
}

// Accept method
func (f *lengthFunc) Accept(typ reflect.Type) bool {
	if reflectx.IsArray(typ) || reflectx.IsSlice(typ) {
		return true
	}
	if reflectx.IsPtr(typ) {
		return f.Accept(typ.Elem())
	}
	return reflectx.IsString(typ)
}

// Pass method
func (f *lengthFunc) Pass(value reflect.Value) (bool, string) {
	passed, msg := false, f.msg
	typ := value.Type()
	switch {
	case reflectx.IsPtr(typ):
		return f.Pass(value.Elem())
	case reflectx.IsArray(typ) || reflectx.IsSlice(typ):
		if reflectx.IsString(value.Type().Elem()) {
			for i := 0; i < value.Len(); i++ {
				passed, msg = f.Pass(value.Index(i))
				if !passed {
					break
				}
			}
		} else {
			passed = f.length == value.Len()
		}
	case reflectx.IsString(typ):
		passed = f.length == len(value.String())
	}
	return passed, msg
}
