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
	"strings"
)

// enumFunc struct
type enumFunc struct {
	options string
	msg     string
}

// EnumFunc method
func EnumFunc(str string) VFunc {
	if str == "" {
		return nil
	}
	vStr, msg := str, ""
	if spIdx := findSpIdx(str); spIdx != -1 {
		vStr, msg = str[:spIdx], str[spIdx+1:]
	}
	return &enumFunc{vStr, msg}
}

// Accept method
func (f *enumFunc) Accept(typ reflect.Type) bool {
	if isCanEle(typ) {
		return f.Accept(typ.Elem())
	}
	return isNumber(typ, false) || reflectx.IsString(typ)
}

// Pass method
func (f *enumFunc) Pass(value reflect.Value) (bool, string) {
	var passed, msg = true, f.msg
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
		set := make(map[int64]struct{}, 0)
		for _, enum := range strings.Split(f.options, "|") {
			parseInt, err := strconv.ParseInt(enum, 10, 64)
			if err != nil {
				panic(err)
			}
			set[parseInt] = struct{}{}
		}
		_, passed = set[value.Int()]
	case reflectx.IsUint(typ):
		set := make(map[uint64]struct{}, 0)
		for _, enum := range strings.Split(f.options, "|") {
			parseUint, err := strconv.ParseUint(enum, 10, 64)
			if err != nil {
				panic(err)
			}
			set[parseUint] = struct{}{}
		}
		_, passed = set[value.Uint()]
	case reflectx.IsString(typ):
		set := make(map[string]struct{}, 0)
		for _, enum := range strings.Split(f.options, "|") {
			set[enum] = struct{}{}
		}
		_, passed = set[value.String()]
	}
	return passed, msg
}
