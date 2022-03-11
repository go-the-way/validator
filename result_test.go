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

import "testing"

func TestResult(t *testing.T) {
	r := Result{nil, false, []*ResultItem{
		{nil, false, "fail"},
		{nil, false, "fail2"},
		{nil, false, "fail3"},
	}}
	if p := r.Passed; p != false {
		t.Fatalf("test failed: expect passed [false], but got [%v]\n", p)
	}
	if msg := r.Messages(); msg != "fail,fail2,fail3" {
		t.Fatalf("test failed: expect message [fail,fail2,fail3], but got [%s]\n", msg)
	}
}
