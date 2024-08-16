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
	"strings"
)

// Item struct
type Item struct {
	Min          string `alias:"min"`           // Min for min value
	Max          string `alias:"max"`           // Max for max value
	MinLength    string `alias:"minlength"`     // MinLength for min length
	ArrMinLength string `alias:"arr_minlength"` // ArrMinLength for array min length
	MaxLength    string `alias:"maxlength"`     // MaxLength for max length
	ArrMaxLength string `alias:"arr_maxlength"` // ArrMaxLength for array max length
	Length       string `alias:"length"`        // Length for length
	ArrLength    string `alias:"arr_length"`    // ArrLength for array length
	Enum         string `alias:"enum"`          // Enum for enum values
	Regex        string `alias:"regex"`         // Regex for regex pattern
	Msg          string `alias:"msg"`           // Msg for message
	Valid        string `alias:"valid"`         // Valid for valid
	Custom       string `alias:"custom"`        // Custom for custom validator
}

func (i *Item) vfs() []VFunc {
	return []VFunc{
		MinFunc(i.Min),
		MaxFunc(i.Max),
		LengthFunc(i.Length),
		ArrLengthFunc(i.ArrLength),
		MinLengthFunc(i.MinLength),
		ArrMinLengthFunc(i.ArrMinLength),
		MaxLengthFunc(i.MaxLength),
		ArrMaxLengthFunc(i.ArrMaxLength),
		EnumFunc(i.Enum),
		RegexFunc(i.Regex),
		ValidFunc(i.Valid),
		customVFMap[i.Custom],
	}
}

// Validate by fields
func (i *Item) Validate(_ reflect.StructField, value reflect.Value) (bool, string) {
	passed, msg := true, i.Msg
	fs := i.vfs()
	for _, f := range fs {
		if f != nil {
			if passed2, msg2 := f.Valid(value); !passed2 {
				passed = false
				msg2 = strings.TrimSpace(msg2)
				if msg2 != "" {
					msg = msg2
				}
			}
		}
	}
	return passed, msg
}
