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
	"testing"
)

func TestNewMaxLength(t *testing.T) {
	if v := MaxLengthFunc(""); v != nil {
		t.Fatal("test failed")
	}
}

func TestNewMaxLengthFuncPanic(t *testing.T) {
	for _, tc := range []struct {
		name string
		str  string
		err  error
	}{
		{"ParseInt", "A", parseNumError("ParseInt", "A")},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				re := recover()
				if re == nil {
					t.Fatalf("%s failed: error expect not nil", t.Name())
				}
				if re == nil || !reflect.DeepEqual(tc.err, re) {
					t.Fatalf("%s failed: error expect [%v], but got [%v]\n", t.Name(), tc.err, re)
				}
			}()
			_ = MaxLengthFunc(tc.str)
		})
	}
}

func TestMaxLengthUnPassAndEmptyMsg(t *testing.T) {
	for _, s := range []testCase{
		{"String", MaxLengthFunc("0"), reflect.ValueOf("A"), false, ""},

		{"StringArr", MaxLengthFunc("0"), reflect.ValueOf([1]string{"A"}), false, ""},

		{"StringSlice", MaxLengthFunc("0"), reflect.ValueOf([]string{"A"}), false, ""},

		{"StringPtr", MaxLengthFunc("0"), reflect.ValueOf(stringPtr("A")), false, ""},

		{"StringPtrArr", MaxLengthFunc("0"), reflect.ValueOf([1]*string{stringPtr("A")}), false, ""},

		{"StringPtrSlice", MaxLengthFunc("0"), reflect.ValueOf([]*string{stringPtr("A")}), false, ""},
	} {
		t.Run(s.name, s.test)
	}
}

func TestMaxLengthUnPassAndMsg(t *testing.T) {
	for _, s := range []testCase{
		{"String", MaxLengthFunc("0,fail"), reflect.ValueOf("A"), false, "fail"},

		{"StringArr", MaxLengthFunc("0,fail"), reflect.ValueOf([1]string{"A"}), false, "fail"},

		{"StringSlice", MaxLengthFunc("0,fail"), reflect.ValueOf([]string{"A"}), false, "fail"},

		{"StringPtr", MaxLengthFunc("0,fail"), reflect.ValueOf(stringPtr("A")), false, "fail"},

		{"StringPtrArr", MaxLengthFunc("0,fail"), reflect.ValueOf([1]*string{stringPtr("A")}), false, "fail"},

		{"StringPtrSlice", MaxLengthFunc("0,fail"), reflect.ValueOf([]*string{stringPtr("A")}), false, "fail"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestMaxLengthUnPassAndChineseMsg(t *testing.T) {
	for _, s := range []testCase{
		{"String", MaxLengthFunc("0,??????"), reflect.ValueOf("A"), false, "??????"},

		{"StringArr", MaxLengthFunc("0,??????"), reflect.ValueOf([1]string{"A"}), false, "??????"},

		{"StringSlice", MaxLengthFunc("0,??????"), reflect.ValueOf([]string{"A"}), false, "??????"},

		{"StringPtr", MaxLengthFunc("0,??????"), reflect.ValueOf(stringPtr("A")), false, "??????"},

		{"StringPtrArr", MaxLengthFunc("0,??????"), reflect.ValueOf([1]*string{stringPtr("A")}), false, "??????"},

		{"StringPtrSlice", MaxLengthFunc("0,??????"), reflect.ValueOf([]*string{stringPtr("A")}), false, "??????"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestMaxLengthUnPassAndSpecialEffectMsg(t *testing.T) {
	for _, s := range []testCase{
		{"String", MaxLengthFunc("0,??????????????????????????????????????????????????????????????????????????????????????????????? ?????????????????????????????????????????????????????????????????????????????????????? ????????????????????????????????????????????? ??? ??? ??? ??? ??? ??? ??? ??? ??????????????? ?????????????????????????????????????????????"), reflect.ValueOf("A"), false, "??????????????????????????????????????????????????????????????????????????????????????????????? ?????????????????????????????????????????????????????????????????????????????????????? ????????????????????????????????????????????? ??? ??? ??? ??? ??? ??? ??? ??? ??????????????? ?????????????????????????????????????????????"},

		{"StringArr", MaxLengthFunc("0,??????????????????????????????????????????????????????????????????????????????????????????????? ?????????????????????????????????????????????????????????????????????????????????????? ????????????????????????????????????????????? ??? ??? ??? ??? ??? ??? ??? ??? ??????????????? ?????????????????????????????????????????????"), reflect.ValueOf([1]string{"A"}), false, "??????????????????????????????????????????????????????????????????????????????????????????????? ?????????????????????????????????????????????????????????????????????????????????????? ????????????????????????????????????????????? ??? ??? ??? ??? ??? ??? ??? ??? ??????????????? ?????????????????????????????????????????????"},

		{"StringSlice", MaxLengthFunc("0,??????????????????????????????????????????????????????????????????????????????????????????????? ?????????????????????????????????????????????????????????????????????????????????????? ????????????????????????????????????????????? ??? ??? ??? ??? ??? ??? ??? ??? ??????????????? ?????????????????????????????????????????????"), reflect.ValueOf([]string{"A"}), false, "??????????????????????????????????????????????????????????????????????????????????????????????? ?????????????????????????????????????????????????????????????????????????????????????? ????????????????????????????????????????????? ??? ??? ??? ??? ??? ??? ??? ??? ??????????????? ?????????????????????????????????????????????"},

		{"StringPtr", MaxLengthFunc("0,??????????????????????????????????????????????????????????????????????????????????????????????? ?????????????????????????????????????????????????????????????????????????????????????? ????????????????????????????????????????????? ??? ??? ??? ??? ??? ??? ??? ??? ??????????????? ?????????????????????????????????????????????"), reflect.ValueOf(stringPtr("A")), false, "??????????????????????????????????????????????????????????????????????????????????????????????? ?????????????????????????????????????????????????????????????????????????????????????? ????????????????????????????????????????????? ??? ??? ??? ??? ??? ??? ??? ??? ??????????????? ?????????????????????????????????????????????"},

		{"StringPtrArr", MaxLengthFunc("0,??????????????????????????????????????????????????????????????????????????????????????????????? ?????????????????????????????????????????????????????????????????????????????????????? ????????????????????????????????????????????? ??? ??? ??? ??? ??? ??? ??? ??? ??????????????? ?????????????????????????????????????????????"), reflect.ValueOf([1]*string{stringPtr("A")}), false, "??????????????????????????????????????????????????????????????????????????????????????????????? ?????????????????????????????????????????????????????????????????????????????????????? ????????????????????????????????????????????? ??? ??? ??? ??? ??? ??? ??? ??? ??????????????? ?????????????????????????????????????????????"},

		{"StringPtrSlice", MaxLengthFunc("0,??????????????????????????????????????????????????????????????????????????????????????????????? ?????????????????????????????????????????????????????????????????????????????????????? ????????????????????????????????????????????? ??? ??? ??? ??? ??? ??? ??? ??? ??????????????? ?????????????????????????????????????????????"), reflect.ValueOf([]*string{stringPtr("A")}), false, "??????????????????????????????????????????????????????????????????????????????????????????????? ?????????????????????????????????????????????????????????????????????????????????????? ????????????????????????????????????????????? ??? ??? ??? ??? ??? ??? ??? ??? ??????????????? ?????????????????????????????????????????????"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestMaxLengthUnPassAndArrowMsg(t *testing.T) {
	for _, s := range []testCase{
		{"String", MaxLengthFunc("0,??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"), reflect.ValueOf("A"), false, "??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"},

		{"StringArr", MaxLengthFunc("0,??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"), reflect.ValueOf([1]string{"A"}), false, "??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"},

		{"StringSlice", MaxLengthFunc("0,??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"), reflect.ValueOf([]string{"A"}), false, "??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"},

		{"StringPtr", MaxLengthFunc("0,??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"), reflect.ValueOf(stringPtr("A")), false, "??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"},

		{"StringPtrArr", MaxLengthFunc("0,??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"), reflect.ValueOf([1]*string{stringPtr("A")}), false, "??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"},

		{"StringPtrSlice", MaxLengthFunc("0,??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"), reflect.ValueOf([]*string{stringPtr("A")}), false, "??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ?????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestMaxLengthUnPassAndForeignMsg(t *testing.T) {
	for _, s := range []testCase{
		{"String", MaxLengthFunc("0,????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?????????????????????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ??"), reflect.ValueOf("A"), false, "????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?????????????????????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ??"},

		{"StringArr", MaxLengthFunc("0,????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?????????????????????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ??"), reflect.ValueOf([1]string{"A"}), false, "????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?????????????????????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ??"},

		{"StringSlice", MaxLengthFunc("0,????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?????????????????????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ??"), reflect.ValueOf([]string{"A"}), false, "????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?????????????????????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ??"},

		{"StringPtr", MaxLengthFunc("0,????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?????????????????????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ??"), reflect.ValueOf(stringPtr("A")), false, "????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?????????????????????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ??"},

		{"StringPtrArr", MaxLengthFunc("0,????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?????????????????????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ??"), reflect.ValueOf([1]*string{stringPtr("A")}), false, "????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?????????????????????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ??"},

		{"StringPtrSlice", MaxLengthFunc("0,????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?????????????????????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ??"), reflect.ValueOf([]*string{stringPtr("A")}), false, "????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?????????????????????????????????????????????????????????????????????? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ?? ??"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestMaxLengthUnPassAndSymbolMsg(t *testing.T) {
	for _, s := range []testCase{
		{"String", MaxLengthFunc("0,?????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????`~ ????????????????????????????? ??? ??? - ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????? ?????? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"), reflect.ValueOf("A"), false, "?????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????`~ ????????????????????????????? ??? ??? - ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????? ?????? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"},

		{"StringArr", MaxLengthFunc("0,?????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????`~ ????????????????????????????? ??? ??? - ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????? ?????? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"), reflect.ValueOf([1]string{"A"}), false, "?????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????`~ ????????????????????????????? ??? ??? - ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????? ?????? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"},

		{"StringSlice", MaxLengthFunc("0,?????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????`~ ????????????????????????????? ??? ??? - ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????? ?????? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"), reflect.ValueOf([]string{"A"}), false, "?????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????`~ ????????????????????????????? ??? ??? - ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????? ?????? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"},

		{"StringPtr", MaxLengthFunc("0,?????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????`~ ????????????????????????????? ??? ??? - ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????? ?????? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"), reflect.ValueOf(stringPtr("A")), false, "?????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????`~ ????????????????????????????? ??? ??? - ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????? ?????? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"},

		{"StringPtrArr", MaxLengthFunc("0,?????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????`~ ????????????????????????????? ??? ??? - ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????? ?????? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"), reflect.ValueOf([1]*string{stringPtr("A")}), false, "?????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????`~ ????????????????????????????? ??? ??? - ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????? ?????? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"},

		{"StringPtrSlice", MaxLengthFunc("0,?????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????`~ ????????????????????????????? ??? ??? - ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????? ?????? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"), reflect.ValueOf([]*string{stringPtr("A")}), false, "?????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????`~ ????????????????????????????? ??? ??? - ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????? ?????? ??? ??? ??? ??? ????????? ??? ??? ??? ??? ??? ??? ??? ??? ??? ???"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestMaxLengthPass(t *testing.T) {
	for _, s := range []testCase{
		{"String", MaxLengthFunc("1"), reflect.ValueOf("1"), true, ""},

		{"StringArr", MaxLengthFunc("1"), reflect.ValueOf([1]string{"1"}), true, ""},

		{"StringSlice", MaxLengthFunc("1"), reflect.ValueOf([]string{"1"}), true, ""},

		{"StringPtr", MaxLengthFunc("1"), reflect.ValueOf(stringPtr("1")), true, ""},

		{"StringPtrArr", MaxLengthFunc("1"), reflect.ValueOf([1]*string{stringPtr("1")}), true, ""},

		{"StringPtrSlice", MaxLengthFunc("1"), reflect.ValueOf([]*string{stringPtr("1")}), true, ""},
	} {
		t.Run(s.name, s.test)
	}
}
