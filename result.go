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

// Result struct
type Result struct {
	StructPtr interface{}
	Passed    bool
	Items     []*ResultItem
	lang      []string
	lm        map[string]int
}

func newResult(structPtr interface{}, items []*ResultItem, passed bool, lang []string) *Result {
	rt := &Result{StructPtr: structPtr, Items: items, Passed: passed, lang: lang}
	// lang: zh-CN,en-US,zh-TW
	var lm = map[string]int{}
	var idx = 0
	for _, lang0 := range lang {
		if lang0 != "" {
			lm[lang0] = idx
			idx++
		}
	}
	rt.lm = lm
	return rt
}

// ResultItem struct
type ResultItem struct {
	Field   *reflect.StructField
	Passed  bool
	Message string
}

// Messages return un-passed messages
func (r *Result) Messages(lang ...string) string {
	langPos := 0
	if len(lang) > 0 {
		if lang0 := lang[0]; len(lang0) > 0 {
			if langPos0, ok := r.lm[lang0]; ok && langPos0 != -1 {
				langPos = langPos0
			}
		}
	}
	messages := make([]string, 0)
	for _, item := range r.Items {
		msg := item.Message
		if !item.Passed && msg != "" {
			msgS := strings.Split(msg, "|")
			if langPos <= len(msgS) {
				msg = msgS[langPos]
			}
			messages = append(messages, msg)
		}
	}
	return strings.Join(messages, ",")
}
