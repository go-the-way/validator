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
)

// Validator defines validator struct
type Validator struct {
	structPtr interface{}
	fields    []*reflect.StructField
	values    []*reflect.Value
	items     []interface{}
}

// New return new *Validator
func New(structPtr interface{}) *Validator {
	fields, values, tags := reflectx.ParseTag(structPtr, new(Item), "alias", "validate", true)
	return &Validator{structPtr, fields, values, tags}
}

// Validate return validation result
func (v *Validator) Validate() *Result {
	resultItems := make([]*ResultItem, len(v.fields))
	passedCount := 0
	for pos := range v.fields {
		field := v.fields[pos]
		value := v.values[pos]
		item := v.items[pos].(*Item)
		resultItem := validate(item, field, value)
		resultItems[pos] = resultItem
		if resultItem.Passed {
			passedCount++
		}
	}
	return &Result{
		StructPtr: v.structPtr,
		Passed:    len(v.items) == passedCount,
		Items:     resultItems,
	}
}

func validate(item *Item, field *reflect.StructField, value *reflect.Value) *ResultItem {
	passed, msg := item.Validate(*field, *value)
	return &ResultItem{Field: field, Passed: passed, Message: msg}
}
