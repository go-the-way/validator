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

// minLengthFunc struct
type minLengthFunc struct {
	minLength int
	msg       string
}

// MinLengthFunc method
func MinLengthFunc(str string) VFunc {
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
		return &minLengthFunc{int(v), msg}
	}
}

// Accept method
func (f *minLengthFunc) Accept(typ reflect.Type) bool {
	if reflectx.IsPtr(typ) {
		return f.Accept(typ.Elem())
	}
	return reflectx.IsString(typ) || reflectx.IsArray(typ) || reflectx.IsSlice(typ)
}

// Pass method
func (f *minLengthFunc) Pass(value reflect.Value) (bool, string) {
	var passed, msg = true, f.msg
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
			passed = f.minLength <= value.Len()
		}
	case reflectx.IsString(typ):
		passed = f.minLength <= len(value.String())
	}
	return passed, msg
}
