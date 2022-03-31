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

func TestNewArrMaxLength(t *testing.T) {
	if v := ArrMaxLengthFunc(""); v != nil {
		t.Fatal("test failed")
	}
}

func TestNewArrMaxLengthFuncPanic(t *testing.T) {
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
			_ = ArrMaxLengthFunc(tc.str)
		})
	}
}

func TestArrMaxLengthUnPassAndEmptyMsg(t *testing.T) {
	for _, s := range []testCase{
		{"RuneArr", ArrMaxLengthFunc("0"), reflect.ValueOf([1]rune{'A'}), false, ""},
		{"ByteArr", ArrMaxLengthFunc("0"), reflect.ValueOf([1]byte{'A'}), false, ""},
		{"Int8Arr", ArrMaxLengthFunc("0"), reflect.ValueOf([1]int8{'A'}), false, ""},
		{"IntArr", ArrMaxLengthFunc("0"), reflect.ValueOf([1]int{'A'}), false, ""},
		{"Int16Arr", ArrMaxLengthFunc("0"), reflect.ValueOf([1]int16{'A'}), false, ""},
		{"Int32Arr", ArrMaxLengthFunc("0"), reflect.ValueOf([1]int32{int32(byte('A'))}), false, ""},
		{"Int64Arr", ArrMaxLengthFunc("0"), reflect.ValueOf([1]int64{'A'}), false, ""},
		{"StringArr", ArrMaxLengthFunc("0"), reflect.ValueOf([1]string{"A"}), false, ""},

		{"RuneSlice", ArrMaxLengthFunc("0"), reflect.ValueOf([]rune{'A'}), false, ""},
		{"ByteSlice", ArrMaxLengthFunc("0"), reflect.ValueOf([]byte{'A'}), false, ""},
		{"Int8Slice", ArrMaxLengthFunc("0"), reflect.ValueOf([]int8{'A'}), false, ""},
		{"IntSlice", ArrMaxLengthFunc("0"), reflect.ValueOf([]int{'A'}), false, ""},
		{"Int16Slice", ArrMaxLengthFunc("0"), reflect.ValueOf([]int16{'A'}), false, ""},
		{"Int32Slice", ArrMaxLengthFunc("0"), reflect.ValueOf([]int32{int32(byte('A'))}), false, ""},
		{"Int64Slice", ArrMaxLengthFunc("0"), reflect.ValueOf([]int64{'A'}), false, ""},
		{"StringSlice", ArrMaxLengthFunc("0"), reflect.ValueOf([]string{"A"}), false, ""},

		{"RunePtrArr", ArrMaxLengthFunc("0"), reflect.ValueOf([1]*rune{runePtr('A')}), false, ""},
		{"BytePtrArr", ArrMaxLengthFunc("0"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), false, ""},
		{"Int8PtrArr", ArrMaxLengthFunc("0"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), false, ""},
		{"IntPtrArr", ArrMaxLengthFunc("0"), reflect.ValueOf([1]*int{intPtr(int('A'))}), false, ""},
		{"Int16PtrArr", ArrMaxLengthFunc("0"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), false, ""},
		{"Int32PtrArr", ArrMaxLengthFunc("0"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), false, ""},
		{"Int64PtrArr", ArrMaxLengthFunc("0"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), false, ""},
		{"StringPtrArr", ArrMaxLengthFunc("0"), reflect.ValueOf([1]*string{stringPtr("A")}), false, ""},

		{"RunePtrSlice", ArrMaxLengthFunc("0"), reflect.ValueOf([]*rune{runePtr('A')}), false, ""},
		{"BytePtrSlice", ArrMaxLengthFunc("0"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), false, ""},
		{"Int8PtrSlice", ArrMaxLengthFunc("0"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), false, ""},
		{"IntPtrSlice", ArrMaxLengthFunc("0"), reflect.ValueOf([]*int{intPtr(int('A'))}), false, ""},
		{"Int16PtrSlice", ArrMaxLengthFunc("0"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), false, ""},
		{"Int32PtrSlice", ArrMaxLengthFunc("0"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), false, ""},
		{"Int64PtrSlice", ArrMaxLengthFunc("0"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), false, ""},
		{"StringPtrSlice", ArrMaxLengthFunc("0"), reflect.ValueOf([]*string{stringPtr("A")}), false, ""},
	} {
		t.Run(s.name, s.test)
	}
}

func TestArrMaxLengthUnPassAndMsg(t *testing.T) {
	for _, s := range []testCase{
		{"RuneArr", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([1]rune{'A'}), false, "fail"},
		{"ByteArr", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([1]byte{'A'}), false, "fail"},
		{"Int8Arr", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([1]int8{'A'}), false, "fail"},
		{"IntArr", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([1]int{'A'}), false, "fail"},
		{"Int16Arr", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([1]int16{'A'}), false, "fail"},
		{"Int32Arr", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([1]int32{int32(byte('A'))}), false, "fail"},
		{"Int64Arr", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([1]int64{'A'}), false, "fail"},
		{"StringArr", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([1]string{"A"}), false, "fail"},

		{"RuneSlice", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([]rune{'A'}), false, "fail"},
		{"ByteSlice", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([]byte{'A'}), false, "fail"},
		{"Int8Slice", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([]int8{'A'}), false, "fail"},
		{"IntSlice", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([]int{'A'}), false, "fail"},
		{"Int16Slice", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([]int16{'A'}), false, "fail"},
		{"Int32Slice", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([]int32{int32(byte('A'))}), false, "fail"},
		{"Int64Slice", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([]int64{'A'}), false, "fail"},
		{"StringSlice", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([]string{"A"}), false, "fail"},

		{"RunePtrArr", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([1]*rune{runePtr('A')}), false, "fail"},
		{"BytePtrArr", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), false, "fail"},
		{"Int8PtrArr", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), false, "fail"},
		{"IntPtrArr", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([1]*int{intPtr(int('A'))}), false, "fail"},
		{"Int16PtrArr", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), false, "fail"},
		{"Int32PtrArr", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), false, "fail"},
		{"Int64PtrArr", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), false, "fail"},
		{"StringPtrArr", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([1]*string{stringPtr("A")}), false, "fail"},

		{"RunePtrSlice", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([]*rune{runePtr('A')}), false, "fail"},
		{"BytePtrSlice", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), false, "fail"},
		{"Int8PtrSlice", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), false, "fail"},
		{"IntPtrSlice", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([]*int{intPtr(int('A'))}), false, "fail"},
		{"Int16PtrSlice", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), false, "fail"},
		{"Int32PtrSlice", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), false, "fail"},
		{"Int64PtrSlice", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), false, "fail"},
		{"StringPtrSlice", ArrMaxLengthFunc("0,fail"), reflect.ValueOf([]*string{stringPtr("A")}), false, "fail"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestArrMaxLengthUnPassAndChineseMsg(t *testing.T) {
	for _, s := range []testCase{
		{"RuneArr", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([1]rune{'A'}), false, "错误"},
		{"ByteArr", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([1]byte{'A'}), false, "错误"},
		{"Int8Arr", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([1]int8{'A'}), false, "错误"},
		{"IntArr", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([1]int{'A'}), false, "错误"},
		{"Int16Arr", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([1]int16{'A'}), false, "错误"},
		{"Int32Arr", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([1]int32{int32(byte('A'))}), false, "错误"},
		{"Int64Arr", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([1]int64{'A'}), false, "错误"},
		{"StringArr", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([1]string{"A"}), false, "错误"},

		{"RuneSlice", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([]rune{'A'}), false, "错误"},
		{"ByteSlice", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([]byte{'A'}), false, "错误"},
		{"Int8Slice", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([]int8{'A'}), false, "错误"},
		{"IntSlice", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([]int{'A'}), false, "错误"},
		{"Int16Slice", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([]int16{'A'}), false, "错误"},
		{"Int32Slice", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([]int32{int32(byte('A'))}), false, "错误"},
		{"Int64Slice", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([]int64{'A'}), false, "错误"},
		{"StringSlice", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([]string{"A"}), false, "错误"},

		{"RunePtrArr", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([1]*rune{runePtr('A')}), false, "错误"},
		{"BytePtrArr", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), false, "错误"},
		{"Int8PtrArr", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), false, "错误"},
		{"IntPtrArr", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([1]*int{intPtr(int('A'))}), false, "错误"},
		{"Int16PtrArr", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), false, "错误"},
		{"Int32PtrArr", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), false, "错误"},
		{"Int64PtrArr", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), false, "错误"},
		{"StringPtrArr", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([1]*string{stringPtr("A")}), false, "错误"},

		{"RunePtrSlice", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([]*rune{runePtr('A')}), false, "错误"},
		{"BytePtrSlice", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), false, "错误"},
		{"Int8PtrSlice", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), false, "错误"},
		{"IntPtrSlice", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([]*int{intPtr(int('A'))}), false, "错误"},
		{"Int16PtrSlice", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), false, "错误"},
		{"Int32PtrSlice", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), false, "错误"},
		{"Int64PtrSlice", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), false, "错误"},
		{"StringPtrSlice", ArrMaxLengthFunc("0,错误"), reflect.ValueOf([]*string{stringPtr("A")}), false, "错误"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestArrMaxLengthUnPassAndSpecialEffectMsg(t *testing.T) {
	for _, s := range []testCase{
		{"RuneArr", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]rune{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"ByteArr", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]byte{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int8Arr", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]int8{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"IntArr", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]int{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int16Arr", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]int16{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int32Arr", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]int32{int32(byte('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int64Arr", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]int64{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"StringArr", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]string{"A"}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},

		{"RuneSlice", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]rune{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"ByteSlice", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]byte{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int8Slice", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]int8{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"IntSlice", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]int{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int16Slice", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]int16{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int32Slice", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]int32{int32(byte('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int64Slice", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]int64{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"StringSlice", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]string{"A"}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},

		{"RunePtrArr", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*rune{runePtr('A')}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"BytePtrArr", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int8PtrArr", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"IntPtrArr", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*int{intPtr(int('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int16PtrArr", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int32PtrArr", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int64PtrArr", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"StringPtrArr", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*string{stringPtr("A")}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},

		{"RunePtrSlice", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*rune{runePtr('A')}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"BytePtrSlice", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int8PtrSlice", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"IntPtrSlice", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*int{intPtr(int('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int16PtrSlice", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int32PtrSlice", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int64PtrSlice", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"StringPtrSlice", ArrMaxLengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*string{stringPtr("A")}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestArrMaxLengthUnPassAndArrowMsg(t *testing.T) {
	for _, s := range []testCase{
		{"RuneArr", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]rune{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"ByteArr", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]byte{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int8Arr", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]int8{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"IntArr", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]int{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int16Arr", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]int16{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int32Arr", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]int32{int32(byte('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int64Arr", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]int64{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"StringArr", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]string{"A"}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},

		{"RuneSlice", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]rune{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"ByteSlice", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]byte{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int8Slice", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]int8{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"IntSlice", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]int{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int16Slice", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]int16{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int32Slice", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]int32{int32(byte('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int64Slice", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]int64{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"StringSlice", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]string{"A"}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},

		{"RunePtrArr", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*rune{runePtr('A')}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"BytePtrArr", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int8PtrArr", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"IntPtrArr", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*int{intPtr(int('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int16PtrArr", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int32PtrArr", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int64PtrArr", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"StringPtrArr", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*string{stringPtr("A")}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},

		{"RunePtrSlice", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*rune{runePtr('A')}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"BytePtrSlice", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int8PtrSlice", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"IntPtrSlice", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*int{intPtr(int('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int16PtrSlice", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int32PtrSlice", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int64PtrSlice", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"StringPtrSlice", ArrMaxLengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*string{stringPtr("A")}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestArrMaxLengthUnPassAndForeignMsg(t *testing.T) {
	for _, s := range []testCase{
		{"RuneArr", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]rune{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"ByteArr", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]byte{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int8Arr", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]int8{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"IntArr", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]int{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int16Arr", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]int16{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int32Arr", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]int32{int32(byte('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int64Arr", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]int64{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"StringArr", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]string{"A"}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},

		{"RuneSlice", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]rune{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"ByteSlice", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]byte{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int8Slice", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]int8{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"IntSlice", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]int{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int16Slice", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]int16{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int32Slice", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]int32{int32(byte('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int64Slice", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]int64{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"StringSlice", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]string{"A"}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},

		{"RunePtrArr", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*rune{runePtr('A')}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"BytePtrArr", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int8PtrArr", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"IntPtrArr", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*int{intPtr(int('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int16PtrArr", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int32PtrArr", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int64PtrArr", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"StringPtrArr", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*string{stringPtr("A")}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},

		{"RunePtrSlice", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*rune{runePtr('A')}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"BytePtrSlice", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int8PtrSlice", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"IntPtrSlice", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*int{intPtr(int('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int16PtrSlice", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int32PtrSlice", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int64PtrSlice", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"StringPtrSlice", ArrMaxLengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*string{stringPtr("A")}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestArrMaxLengthUnPassAndSymbolMsg(t *testing.T) {
	for _, s := range []testCase{
		{"RuneArr", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]rune{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"ByteArr", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]byte{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int8Arr", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]int8{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"IntArr", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]int{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int16Arr", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]int16{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int32Arr", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]int32{int32(byte('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int64Arr", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]int64{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"StringArr", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]string{"A"}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},

		{"RuneSlice", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]rune{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"ByteSlice", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]byte{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int8Slice", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]int8{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"IntSlice", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]int{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int16Slice", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]int16{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int32Slice", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]int32{int32(byte('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int64Slice", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]int64{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"StringSlice", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]string{"A"}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},

		{"RunePtrArr", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*rune{runePtr('A')}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"BytePtrArr", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int8PtrArr", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"IntPtrArr", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*int{intPtr(int('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int16PtrArr", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int32PtrArr", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int64PtrArr", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"StringPtrArr", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*string{stringPtr("A")}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},

		{"RunePtrSlice", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*rune{runePtr('A')}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"BytePtrSlice", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int8PtrSlice", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"IntPtrSlice", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*int{intPtr(int('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int16PtrSlice", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int32PtrSlice", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int64PtrSlice", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"StringPtrSlice", ArrMaxLengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*string{stringPtr("A")}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestArrMaxLengthPass(t *testing.T) {
	for _, s := range []testCase{
		{"String", ArrMaxLengthFunc("1"), reflect.ValueOf("1"), true, ""},

		{"RuneArr", ArrMaxLengthFunc("1"), reflect.ValueOf([1]rune{1}), true, ""},
		{"ByteArr", ArrMaxLengthFunc("1"), reflect.ValueOf([1]byte{1}), true, ""},
		{"Int8Arr", ArrMaxLengthFunc("1"), reflect.ValueOf([1]int8{1}), true, ""},
		{"IntArr", ArrMaxLengthFunc("1"), reflect.ValueOf([1]int{1}), true, ""},
		{"Int16Arr", ArrMaxLengthFunc("1"), reflect.ValueOf([1]int16{1}), true, ""},
		{"Int32Arr", ArrMaxLengthFunc("1"), reflect.ValueOf([1]int32{1}), true, ""},
		{"Int64Arr", ArrMaxLengthFunc("1"), reflect.ValueOf([1]int64{1}), true, ""},
		{"StringArr", ArrMaxLengthFunc("1"), reflect.ValueOf([1]string{"1"}), true, ""},

		{"RuneSlice", ArrMaxLengthFunc("1"), reflect.ValueOf([]rune{rune(1)}), true, ""},
		{"ByteSlice", ArrMaxLengthFunc("1"), reflect.ValueOf([]byte{byte(1)}), true, ""},
		{"Int8Slice", ArrMaxLengthFunc("1"), reflect.ValueOf([]int8{int8(1)}), true, ""},
		{"IntSlice", ArrMaxLengthFunc("1"), reflect.ValueOf([]int{1}), true, ""},
		{"Int16Slice", ArrMaxLengthFunc("1"), reflect.ValueOf([]int16{int16(1)}), true, ""},
		{"Int32Slice", ArrMaxLengthFunc("1"), reflect.ValueOf([]int32{int32(1)}), true, ""},
		{"Int64Slice", ArrMaxLengthFunc("1"), reflect.ValueOf([]int64{int64(1)}), true, ""},
		{"StringSlice", ArrMaxLengthFunc("1"), reflect.ValueOf([]string{"1"}), true, ""},

		{"StringPtr", ArrMaxLengthFunc("1"), reflect.ValueOf(stringPtr("1")), true, ""},

		{"RunePtrArr", ArrMaxLengthFunc("1"), reflect.ValueOf([1]*rune{runePtr(rune(1))}), true, ""},
		{"BytePtrArr", ArrMaxLengthFunc("1"), reflect.ValueOf([1]*byte{bytePtr(byte(1))}), true, ""},
		{"Int8PtrArr", ArrMaxLengthFunc("1"), reflect.ValueOf([1]*int8{int8Ptr(int8(1))}), true, ""},
		{"IntPtrArr", ArrMaxLengthFunc("1"), reflect.ValueOf([1]*int{intPtr(1)}), true, ""},
		{"Int16PtrArr", ArrMaxLengthFunc("1"), reflect.ValueOf([1]*int16{int16Ptr(int16(1))}), true, ""},
		{"Int32PtrArr", ArrMaxLengthFunc("1"), reflect.ValueOf([1]*int32{int32Ptr(int32(1))}), true, ""},
		{"Int64PtrArr", ArrMaxLengthFunc("1"), reflect.ValueOf([1]*int64{int64Ptr(int64(1))}), true, ""},
		{"StringPtrArr", ArrMaxLengthFunc("1"), reflect.ValueOf([1]*string{stringPtr("1")}), true, ""},

		{"RunePtrSlice", ArrMaxLengthFunc("1"), reflect.ValueOf([]*rune{runePtr(rune(1))}), true, ""},
		{"BytePtrSlice", ArrMaxLengthFunc("1"), reflect.ValueOf([]*byte{bytePtr(byte(1))}), true, ""},
		{"Int8PtrSlice", ArrMaxLengthFunc("1"), reflect.ValueOf([]*int8{int8Ptr(int8(1))}), true, ""},
		{"IntPtrSlice", ArrMaxLengthFunc("1"), reflect.ValueOf([]*int{intPtr(1)}), true, ""},
		{"Int16PtrSlice", ArrMaxLengthFunc("1"), reflect.ValueOf([]*int16{int16Ptr(int16(1))}), true, ""},
		{"Int32PtrSlice", ArrMaxLengthFunc("1"), reflect.ValueOf([]*int32{int32Ptr(1)}), true, ""},
		{"Int64PtrSlice", ArrMaxLengthFunc("1"), reflect.ValueOf([]*int64{int64Ptr(int64(1))}), true, ""},
		{"StringPtrSlice", ArrMaxLengthFunc("1"), reflect.ValueOf([]*string{stringPtr("1")}), true, ""},
	} {
		t.Run(s.name, s.test)
	}
}
