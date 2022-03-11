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

// minFunc struct
type minFunc struct {
	min float64
	msg string
}

// MinFunc method
func MinFunc(str string) VFunc {
	if str == "" {
		return nil
	}
	floatStr := str
	msg := ""
	if spIdx := findSpIdx(str); spIdx != -1 {
		floatStr = str[:spIdx]
		msg = str[spIdx+1:]
	}
	if v, err := strconv.ParseFloat(floatStr, 64); err != nil {
		panic(err)
	} else {
		return &minFunc{v, msg}
	}
}

// Accept method
func (f *minFunc) Accept(typ reflect.Type) bool {
	if isCanEle(typ) {
		return f.Accept(typ.Elem())
	}
	return isNumber(typ, true)
}

// Pass method
func (f *minFunc) Pass(value reflect.Value) (bool, string) {
	passed, msg := true, f.msg
	typ := value.Type()
	switch {
	case reflectx.IsPtr(typ):
		return f.Pass(value.Elem())
	case reflectx.IsArray(typ) || reflectx.IsSlice(typ):
		for i := 0; i < value.Len(); i++ {
			passed, msg = f.Pass(value.Index(i))
			if !passed {
				break
			}
		}
	case reflectx.IsInt(typ):
		passed = f.min <= float64(value.Int())
	case reflectx.IsUint(typ):
		passed = f.min <= float64(value.Uint())
	case reflectx.IsFloat(typ):
		passed = f.min <= value.Float()
	}
	return passed, msg
}
