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
	"regexp"
)

// regexFunc struct
type regexFunc struct {
	patten string
	msg    string
}

// RegexFunc method
func RegexFunc(str string) VFunc {
	if str == "" {
		return nil
	}
	vStr, msg := str, ""
	if spIdx := findSpIdx(str); spIdx != -1 {
		vStr, msg = str[:spIdx], str[spIdx+1:]
	}
	return &regexFunc{vStr, msg}
}

// Valid method
func (f *regexFunc) Valid(value reflect.Value) (bool, string) {
	passed, msg := true, f.msg
	re := regexp.MustCompile(f.patten)
	typ := value.Type()
	switch {
	case reflectx.IsPtr(typ):
		return f.Valid(value.Elem())
	case reflectx.IsArray(typ) || reflectx.IsSlice(typ):
		for i := 0; i < value.Len(); i++ {
			passed, msg = f.Valid(value.Index(i))
			if !passed {
				break
			}
		}
	case reflectx.IsString(typ):
		passed = re.MatchString(value.String())
	}
	return passed, msg
}
