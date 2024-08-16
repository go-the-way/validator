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
	"testing"
)

func TestResult(t *testing.T) {
	r := Result{nil, false, []*ResultItem{
		{nil, false, "fail"},
		{nil, false, "fail2"},
		{nil, false, "fail3"},
	}, nil, nil}
	if p := r.Passed; p != false {
		t.Fatalf("test failed: expect passed [false], but got [%v]\n", p)
	}
	if msg := r.Messages(); msg != "fail,fail2,fail3" {
		t.Fatalf("test failed: expect message [fail,fail2,fail3], but got [%s]\n", msg)
	}
}

func TestNewResult(t *testing.T) {
	rt := newResult(struct{}{}, nil, false, []string{"zh-CN", "en-US"})
	var (
		zhCNPos = -1
		enUSPos = -1
	)
	if zhCNPos = rt.lm["zh-CN"]; zhCNPos != 0 {
		t.Fatalf("test failed: expect [0], but got [%d]\n", zhCNPos)
	}
	if enUSPos = rt.lm["en-US"]; enUSPos != 1 {
		t.Fatalf("test failed: expect [0], but got [%d]\n", enUSPos)
	}
}

func TestResult_Messages(t *testing.T) {
	zhCN := "zh-CN"
	enUS := "en-US"
	langS := []string{zhCN, enUS}
	msgZhCN := "错误"
	msgEnUS := "Error"
	rt := newResult(struct{}{}, []*ResultItem{{Passed: false, Message: msgZhCN + "|" + msgEnUS}}, false, langS)
	var (
		zhCNPos = -1
		enUSPos = -1
	)
	if zhCNPos = rt.lm[zhCN]; zhCNPos != 0 {
		t.Fatalf("test failed: expect [0], but got [%d]\n", zhCNPos)
	}
	if enUSPos = rt.lm[enUS]; enUSPos != 1 {
		t.Fatalf("test failed: expect [0], but got [%d]\n", enUSPos)
	}
	if msg := rt.Messages(zhCN); msg != msgZhCN {
		t.Fatalf("test failed: expect [%s], but got [%s]\n", msgZhCN, msg)
	}
	if msg := rt.Messages(enUS); msg != msgEnUS {
		t.Fatalf("test failed: expect [%s], but got [%s]\n", msgEnUS, msg)
	}
}
