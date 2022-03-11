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
)

// Item struct
type Item struct {
	Min       string `alias:"min"`       // Min for min value
	Max       string `alias:"max"`       // Max for max value
	MinLength string `alias:"minlength"` // MinLength for min length
	MaxLength string `alias:"maxlength"` // MaxLength for max length
	Length    string `alias:"length"`    // Length for length
	Enum      string `alias:"enum"`      // Enum for enum values
	Regex     string `alias:"regex"`     // Regex for regex pattern
	Msg       string `alias:"msg"`       // Msg for message
}

func (i *Item) vfs() []VFunc {
	return []VFunc{
		MinFunc(i.Min),
		MaxFunc(i.Max),
		LengthFunc(i.Length),
		MinLengthFunc(i.MinLength),
		MaxLengthFunc(i.MaxLength),
		EnumFunc(i.Enum),
		RegexFunc(i.Regex),
	}
}

// Validate by fields
func (i *Item) Validate(field reflect.StructField, value reflect.Value) (bool, string) {
	passed, msg := true, i.Msg
	fs := i.vfs()
	for _, f := range fs {
		if f != nil {
			if f.Accept(field.Type) {
				if passed2, msg2 := f.Pass(value); !passed2 {
					passed = false
					if msg2 != "" {
						msg = msg2
					}
				}
			}
		}
	}

	return passed, msg
}
