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
	"errors"
	"strconv"
	"strings"
)

func findSpIdx(str string) int {
	idx := strings.Index(str, "[,]")
	if idx == -1 {
		idx = strings.Index(str, ",")
	}
	return idx
}

func bytePtr(i byte) *byte                    { return &i }
func runePtr(i rune) *rune                    { return &i }
func int8Ptr(i int8) *int8                    { return &i }
func int16Ptr(i int16) *int16                 { return &i }
func int32Ptr(i int32) *int32                 { return &i }
func int64Ptr(i int64) *int64                 { return &i }
func intPtr(i int) *int                       { return &i }
func stringPtr(i string) *string              { return &i }
func float32Ptr(i float32) *float32           { return &i }
func float64Ptr(i float64) *float64           { return &i }
func arrayPtr(i []interface{}) *[]interface{} { return &i }

func parseNumError(fN, v string) error {
	return &strconv.NumError{Func: fN, Num: v, Err: errors.New("invalid syntax")}
}
