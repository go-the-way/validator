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

func TestNewLengthFunc(t *testing.T) {
	if v := LengthFunc(""); v != nil {
		t.Fatal("test failed")
	}
}

func TestNewLengthFuncPanic(t *testing.T) {
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
			_ = LengthFunc(tc.str)
		})
	}
}

func TestLengthNoAccept(t *testing.T) {
	for _, s := range []testCase{
		{"Struct", LengthFunc("0"), reflect.ValueOf(struct{}{}), false, false, ""},
		{"BoolTrue", LengthFunc("0"), reflect.ValueOf(true), false, false, ""},
		{"BoolFalse", LengthFunc("0"), reflect.ValueOf(false), false, false, ""},
		{"Float32", LengthFunc("0"), reflect.ValueOf(float32(12.011111252)), false, false, ""},
		{"Float64", LengthFunc("0"), reflect.ValueOf(12.02511111111111111112), false, false, ""},
		{"Float32-", LengthFunc("0"), reflect.ValueOf(-float32(12.011111252)), false, false, ""},
		{"Float64-", LengthFunc("0"), reflect.ValueOf(-12.02511111111111111112), false, false, ""},
		{"Uintptr", LengthFunc("0"), reflect.ValueOf(uintptr(100)), false, false, ""},
		{"Complex64", LengthFunc("0"), reflect.ValueOf(complex64(100)), false, false, ""},
		{"Complex128", LengthFunc("0"), reflect.ValueOf(complex128(100)), false, false, ""},
		{"Complex", LengthFunc("0"), reflect.ValueOf(complex(100, 100)), false, false, ""},
		{"Chan", LengthFunc("0"), reflect.ValueOf(make(chan bool)), false, false, ""},
		{"Map", LengthFunc("0"), reflect.ValueOf(make(map[struct{}]struct{})), false, false, ""},

		{"StructPtr", LengthFunc("0"), reflect.ValueOf(&struct{}{}), false, false, ""},
		{"BoolTruePtr", LengthFunc("0"), reflect.ValueOf(boolPtr(true)), false, false, ""},
		{"BoolFalsePtr", LengthFunc("0"), reflect.ValueOf(boolPtr(false)), false, false, ""},
		{"Float32Ptr", LengthFunc("0"), reflect.ValueOf(float32Ptr(float32(12.011111252))), false, false, ""},
		{"Float64Ptr", LengthFunc("0"), reflect.ValueOf(float64Ptr(12.02511111111111111112)), false, false, ""},
		{"Float32-Ptr", LengthFunc("0"), reflect.ValueOf(float32Ptr(-float32(12.011111252))), false, false, ""},
		{"Float64-Ptr", LengthFunc("0"), reflect.ValueOf(float64Ptr(-12.02511111111111111112)), false, false, ""},
		{"UintptrPtr", LengthFunc("0"), reflect.ValueOf(uintptrPtr(uintptr(100))), false, false, ""},
		{"Complex64Ptr", LengthFunc("0"), reflect.ValueOf(complex64Ptr(complex64(100))), false, false, ""},
		{"Complex128Ptr", LengthFunc("0"), reflect.ValueOf(complex128Ptr(complex128(100))), false, false, ""},
		{"ChanPtr", LengthFunc("0"), reflect.ValueOf(chanPtr(make(chan interface{}))), false, false, ""},
		{"MapPtr", LengthFunc("0"), reflect.ValueOf(mapPtr(make(map[interface{}]interface{}))), false, false, ""},
	} {
		t.Run(s.name, s.test)
	}
}

func TestLengthUnPassAndEmptyMsg(t *testing.T) {
	for _, s := range []testCase{
		{"String", LengthFunc("0"), reflect.ValueOf("A"), true, false, ""},

		{"RuneArr", LengthFunc("0"), reflect.ValueOf([1]rune{'A'}), true, false, ""},
		{"ByteArr", LengthFunc("0"), reflect.ValueOf([1]byte{'A'}), true, false, ""},
		{"Int8Arr", LengthFunc("0"), reflect.ValueOf([1]int8{'A'}), true, false, ""},
		{"IntArr", LengthFunc("0"), reflect.ValueOf([1]int{'A'}), true, false, ""},
		{"Int16Arr", LengthFunc("0"), reflect.ValueOf([1]int16{'A'}), true, false, ""},
		{"Int32Arr", LengthFunc("0"), reflect.ValueOf([1]int32{int32(byte('A'))}), true, false, ""},
		{"Int64Arr", LengthFunc("0"), reflect.ValueOf([1]int64{'A'}), true, false, ""},
		{"StringArr", LengthFunc("0"), reflect.ValueOf([1]string{"A"}), true, false, ""},

		{"RuneSlice", LengthFunc("0"), reflect.ValueOf([]rune{'A'}), true, false, ""},
		{"ByteSlice", LengthFunc("0"), reflect.ValueOf([]byte{'A'}), true, false, ""},
		{"Int8Slice", LengthFunc("0"), reflect.ValueOf([]int8{'A'}), true, false, ""},
		{"IntSlice", LengthFunc("0"), reflect.ValueOf([]int{'A'}), true, false, ""},
		{"Int16Slice", LengthFunc("0"), reflect.ValueOf([]int16{'A'}), true, false, ""},
		{"Int32Slice", LengthFunc("0"), reflect.ValueOf([]int32{int32(byte('A'))}), true, false, ""},
		{"Int64Slice", LengthFunc("0"), reflect.ValueOf([]int64{'A'}), true, false, ""},
		{"StringSlice", LengthFunc("0"), reflect.ValueOf([]string{"A"}), true, false, ""},

		{"StringPtr", LengthFunc("0"), reflect.ValueOf(stringPtr("A")), true, false, ""},

		{"RunePtrArr", LengthFunc("0"), reflect.ValueOf([1]*rune{runePtr('A')}), true, false, ""},
		{"BytePtrArr", LengthFunc("0"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), true, false, ""},
		{"Int8PtrArr", LengthFunc("0"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), true, false, ""},
		{"IntPtrArr", LengthFunc("0"), reflect.ValueOf([1]*int{intPtr(int('A'))}), true, false, ""},
		{"Int16PtrArr", LengthFunc("0"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), true, false, ""},
		{"Int32PtrArr", LengthFunc("0"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), true, false, ""},
		{"Int64PtrArr", LengthFunc("0"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), true, false, ""},
		{"StringPtrArr", LengthFunc("0"), reflect.ValueOf([1]*string{stringPtr("A")}), true, false, ""},

		{"RunePtrSlice", LengthFunc("0"), reflect.ValueOf([]*rune{runePtr('A')}), true, false, ""},
		{"BytePtrSlice", LengthFunc("0"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), true, false, ""},
		{"Int8PtrSlice", LengthFunc("0"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), true, false, ""},
		{"IntPtrSlice", LengthFunc("0"), reflect.ValueOf([]*int{intPtr(int('A'))}), true, false, ""},
		{"Int16PtrSlice", LengthFunc("0"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), true, false, ""},
		{"Int32PtrSlice", LengthFunc("0"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), true, false, ""},
		{"Int64PtrSlice", LengthFunc("0"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), true, false, ""},
		{"StringPtrSlice", LengthFunc("0"), reflect.ValueOf([]*string{stringPtr("A")}), true, false, ""},
	} {
		t.Run(s.name, s.test)
	}
}

func TestLengthUnPassAndMsg(t *testing.T) {
	for _, s := range []testCase{
		{"String", LengthFunc("0,fail"), reflect.ValueOf("A"), true, false, "fail"},

		{"RuneArr", LengthFunc("0,fail"), reflect.ValueOf([1]rune{'A'}), true, false, "fail"},
		{"ByteArr", LengthFunc("0,fail"), reflect.ValueOf([1]byte{'A'}), true, false, "fail"},
		{"Int8Arr", LengthFunc("0,fail"), reflect.ValueOf([1]int8{'A'}), true, false, "fail"},
		{"IntArr", LengthFunc("0,fail"), reflect.ValueOf([1]int{'A'}), true, false, "fail"},
		{"Int16Arr", LengthFunc("0,fail"), reflect.ValueOf([1]int16{'A'}), true, false, "fail"},
		{"Int32Arr", LengthFunc("0,fail"), reflect.ValueOf([1]int32{int32(byte('A'))}), true, false, "fail"},
		{"Int64Arr", LengthFunc("0,fail"), reflect.ValueOf([1]int64{'A'}), true, false, "fail"},
		{"StringArr", LengthFunc("0,fail"), reflect.ValueOf([1]string{"A"}), true, false, "fail"},

		{"RuneSlice", LengthFunc("0,fail"), reflect.ValueOf([]rune{'A'}), true, false, "fail"},
		{"ByteSlice", LengthFunc("0,fail"), reflect.ValueOf([]byte{'A'}), true, false, "fail"},
		{"Int8Slice", LengthFunc("0,fail"), reflect.ValueOf([]int8{'A'}), true, false, "fail"},
		{"IntSlice", LengthFunc("0,fail"), reflect.ValueOf([]int{'A'}), true, false, "fail"},
		{"Int16Slice", LengthFunc("0,fail"), reflect.ValueOf([]int16{'A'}), true, false, "fail"},
		{"Int32Slice", LengthFunc("0,fail"), reflect.ValueOf([]int32{int32(byte('A'))}), true, false, "fail"},
		{"Int64Slice", LengthFunc("0,fail"), reflect.ValueOf([]int64{'A'}), true, false, "fail"},
		{"StringSlice", LengthFunc("0,fail"), reflect.ValueOf([]string{"A"}), true, false, "fail"},

		{"StringPtr", LengthFunc("0,fail"), reflect.ValueOf(stringPtr("A")), true, false, "fail"},

		{"RunePtrArr", LengthFunc("0,fail"), reflect.ValueOf([1]*rune{runePtr('A')}), true, false, "fail"},
		{"BytePtrArr", LengthFunc("0,fail"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), true, false, "fail"},
		{"Int8PtrArr", LengthFunc("0,fail"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), true, false, "fail"},
		{"IntPtrArr", LengthFunc("0,fail"), reflect.ValueOf([1]*int{intPtr(int('A'))}), true, false, "fail"},
		{"Int16PtrArr", LengthFunc("0,fail"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), true, false, "fail"},
		{"Int32PtrArr", LengthFunc("0,fail"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), true, false, "fail"},
		{"Int64PtrArr", LengthFunc("0,fail"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), true, false, "fail"},
		{"StringPtrArr", LengthFunc("0,fail"), reflect.ValueOf([1]*string{stringPtr("A")}), true, false, "fail"},

		{"RunePtrSlice", LengthFunc("0,fail"), reflect.ValueOf([]*rune{runePtr('A')}), true, false, "fail"},
		{"BytePtrSlice", LengthFunc("0,fail"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), true, false, "fail"},
		{"Int8PtrSlice", LengthFunc("0,fail"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), true, false, "fail"},
		{"IntPtrSlice", LengthFunc("0,fail"), reflect.ValueOf([]*int{intPtr(int('A'))}), true, false, "fail"},
		{"Int16PtrSlice", LengthFunc("0,fail"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), true, false, "fail"},
		{"Int32PtrSlice", LengthFunc("0,fail"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), true, false, "fail"},
		{"Int64PtrSlice", LengthFunc("0,fail"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), true, false, "fail"},
		{"StringPtrSlice", LengthFunc("0,fail"), reflect.ValueOf([]*string{stringPtr("A")}), true, false, "fail"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestLengthUnPassAndChineseMsg(t *testing.T) {
	for _, s := range []testCase{
		{"String", LengthFunc("0,错误"), reflect.ValueOf("A"), true, false, "错误"},

		{"RuneArr", LengthFunc("0,错误"), reflect.ValueOf([1]rune{'A'}), true, false, "错误"},
		{"ByteArr", LengthFunc("0,错误"), reflect.ValueOf([1]byte{'A'}), true, false, "错误"},
		{"Int8Arr", LengthFunc("0,错误"), reflect.ValueOf([1]int8{'A'}), true, false, "错误"},
		{"IntArr", LengthFunc("0,错误"), reflect.ValueOf([1]int{'A'}), true, false, "错误"},
		{"Int16Arr", LengthFunc("0,错误"), reflect.ValueOf([1]int16{'A'}), true, false, "错误"},
		{"Int32Arr", LengthFunc("0,错误"), reflect.ValueOf([1]int32{int32(byte('A'))}), true, false, "错误"},
		{"Int64Arr", LengthFunc("0,错误"), reflect.ValueOf([1]int64{'A'}), true, false, "错误"},
		{"StringArr", LengthFunc("0,错误"), reflect.ValueOf([1]string{"A"}), true, false, "错误"},

		{"RuneSlice", LengthFunc("0,错误"), reflect.ValueOf([]rune{'A'}), true, false, "错误"},
		{"ByteSlice", LengthFunc("0,错误"), reflect.ValueOf([]byte{'A'}), true, false, "错误"},
		{"Int8Slice", LengthFunc("0,错误"), reflect.ValueOf([]int8{'A'}), true, false, "错误"},
		{"IntSlice", LengthFunc("0,错误"), reflect.ValueOf([]int{'A'}), true, false, "错误"},
		{"Int16Slice", LengthFunc("0,错误"), reflect.ValueOf([]int16{'A'}), true, false, "错误"},
		{"Int32Slice", LengthFunc("0,错误"), reflect.ValueOf([]int32{int32(byte('A'))}), true, false, "错误"},
		{"Int64Slice", LengthFunc("0,错误"), reflect.ValueOf([]int64{'A'}), true, false, "错误"},
		{"StringSlice", LengthFunc("0,错误"), reflect.ValueOf([]string{"A"}), true, false, "错误"},

		{"StringPtr", LengthFunc("0,错误"), reflect.ValueOf(stringPtr("A")), true, false, "错误"},

		{"RunePtrArr", LengthFunc("0,错误"), reflect.ValueOf([1]*rune{runePtr('A')}), true, false, "错误"},
		{"BytePtrArr", LengthFunc("0,错误"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), true, false, "错误"},
		{"Int8PtrArr", LengthFunc("0,错误"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), true, false, "错误"},
		{"IntPtrArr", LengthFunc("0,错误"), reflect.ValueOf([1]*int{intPtr(int('A'))}), true, false, "错误"},
		{"Int16PtrArr", LengthFunc("0,错误"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), true, false, "错误"},
		{"Int32PtrArr", LengthFunc("0,错误"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), true, false, "错误"},
		{"Int64PtrArr", LengthFunc("0,错误"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), true, false, "错误"},
		{"StringPtrArr", LengthFunc("0,错误"), reflect.ValueOf([1]*string{stringPtr("A")}), true, false, "错误"},

		{"RunePtrSlice", LengthFunc("0,错误"), reflect.ValueOf([]*rune{runePtr('A')}), true, false, "错误"},
		{"BytePtrSlice", LengthFunc("0,错误"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), true, false, "错误"},
		{"Int8PtrSlice", LengthFunc("0,错误"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), true, false, "错误"},
		{"IntPtrSlice", LengthFunc("0,错误"), reflect.ValueOf([]*int{intPtr(int('A'))}), true, false, "错误"},
		{"Int16PtrSlice", LengthFunc("0,错误"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), true, false, "错误"},
		{"Int32PtrSlice", LengthFunc("0,错误"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), true, false, "错误"},
		{"Int64PtrSlice", LengthFunc("0,错误"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), true, false, "错误"},
		{"StringPtrSlice", LengthFunc("0,错误"), reflect.ValueOf([]*string{stringPtr("A")}), true, false, "错误"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestLengthUnPassAndSpecialEffectMsg(t *testing.T) {
	for _, s := range []testCase{
		{"String", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf("A"), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},

		{"RuneArr", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]rune{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"ByteArr", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]byte{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int8Arr", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]int8{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"IntArr", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]int{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int16Arr", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]int16{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int32Arr", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]int32{int32(byte('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int64Arr", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]int64{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"StringArr", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]string{"A"}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},

		{"RuneSlice", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]rune{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"ByteSlice", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]byte{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int8Slice", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]int8{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"IntSlice", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]int{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int16Slice", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]int16{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int32Slice", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]int32{int32(byte('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int64Slice", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]int64{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"StringSlice", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]string{"A"}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},

		{"StringPtr", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(stringPtr("A")), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},

		{"RunePtrArr", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*rune{runePtr('A')}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"BytePtrArr", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int8PtrArr", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"IntPtrArr", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*int{intPtr(int('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int16PtrArr", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int32PtrArr", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int64PtrArr", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"StringPtrArr", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*string{stringPtr("A")}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},

		{"RunePtrSlice", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*rune{runePtr('A')}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"BytePtrSlice", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int8PtrSlice", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"IntPtrSlice", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*int{intPtr(int('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int16PtrSlice", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int32PtrSlice", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int64PtrSlice", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"StringPtrSlice", LengthFunc("0,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*string{stringPtr("A")}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestLengthUnPassAndArrowMsg(t *testing.T) {
	for _, s := range []testCase{
		{"String", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf("A"), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},

		{"RuneArr", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]rune{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"ByteArr", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]byte{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int8Arr", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]int8{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"IntArr", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]int{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int16Arr", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]int16{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int32Arr", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]int32{int32(byte('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int64Arr", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]int64{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"StringArr", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]string{"A"}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},

		{"RuneSlice", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]rune{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"ByteSlice", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]byte{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int8Slice", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]int8{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"IntSlice", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]int{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int16Slice", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]int16{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int32Slice", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]int32{int32(byte('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int64Slice", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]int64{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"StringSlice", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]string{"A"}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},

		{"StringPtr", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(stringPtr("A")), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},

		{"RunePtrArr", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*rune{runePtr('A')}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"BytePtrArr", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int8PtrArr", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"IntPtrArr", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*int{intPtr(int('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int16PtrArr", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int32PtrArr", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int64PtrArr", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"StringPtrArr", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*string{stringPtr("A")}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},

		{"RunePtrSlice", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*rune{runePtr('A')}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"BytePtrSlice", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int8PtrSlice", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"IntPtrSlice", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*int{intPtr(int('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int16PtrSlice", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int32PtrSlice", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int64PtrSlice", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"StringPtrSlice", LengthFunc("0,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*string{stringPtr("A")}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestLengthUnPassAndForeignMsg(t *testing.T) {
	for _, s := range []testCase{
		{"String", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf("A"), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},

		{"RuneArr", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]rune{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"ByteArr", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]byte{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int8Arr", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]int8{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"IntArr", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]int{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int16Arr", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]int16{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int32Arr", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]int32{int32(byte('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int64Arr", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]int64{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"StringArr", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]string{"A"}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},

		{"RuneSlice", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]rune{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"ByteSlice", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]byte{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int8Slice", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]int8{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"IntSlice", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]int{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int16Slice", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]int16{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int32Slice", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]int32{int32(byte('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int64Slice", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]int64{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"StringSlice", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]string{"A"}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},

		{"StringPtr", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(stringPtr("A")), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},

		{"RunePtrArr", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*rune{runePtr('A')}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"BytePtrArr", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int8PtrArr", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"IntPtrArr", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*int{intPtr(int('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int16PtrArr", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int32PtrArr", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int64PtrArr", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"StringPtrArr", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*string{stringPtr("A")}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},

		{"RunePtrSlice", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*rune{runePtr('A')}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"BytePtrSlice", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int8PtrSlice", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"IntPtrSlice", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*int{intPtr(int('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int16PtrSlice", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int32PtrSlice", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int64PtrSlice", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"StringPtrSlice", LengthFunc("0,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*string{stringPtr("A")}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestLengthUnPassAndSymbolMsg(t *testing.T) {
	for _, s := range []testCase{
		{"String", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf("A"), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},

		{"RuneArr", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]rune{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"ByteArr", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]byte{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int8Arr", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]int8{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"IntArr", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]int{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int16Arr", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]int16{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int32Arr", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]int32{int32(byte('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int64Arr", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]int64{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"StringArr", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]string{"A"}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},

		{"RuneSlice", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]rune{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"ByteSlice", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]byte{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int8Slice", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]int8{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"IntSlice", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]int{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int16Slice", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]int16{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int32Slice", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]int32{int32(byte('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int64Slice", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]int64{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"StringSlice", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]string{"A"}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},

		{"StringPtr", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(stringPtr("A")), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},

		{"RunePtrArr", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*rune{runePtr('A')}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"BytePtrArr", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int8PtrArr", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"IntPtrArr", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*int{intPtr(int('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int16PtrArr", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int32PtrArr", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int64PtrArr", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"StringPtrArr", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*string{stringPtr("A")}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},

		{"RunePtrSlice", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*rune{runePtr('A')}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"BytePtrSlice", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int8PtrSlice", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"IntPtrSlice", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*int{intPtr(int('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int16PtrSlice", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int32PtrSlice", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int64PtrSlice", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"StringPtrSlice", LengthFunc("0,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*string{stringPtr("A")}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestLengthPass(t *testing.T) {
	for _, s := range []testCase{
		{"String", LengthFunc("1"), reflect.ValueOf("1"), true, true, ""},

		{"RuneArr", LengthFunc("1"), reflect.ValueOf([1]rune{1}), true, true, ""},
		{"ByteArr", LengthFunc("1"), reflect.ValueOf([1]byte{1}), true, true, ""},
		{"Int8Arr", LengthFunc("1"), reflect.ValueOf([1]int8{1}), true, true, ""},
		{"IntArr", LengthFunc("1"), reflect.ValueOf([1]int{1}), true, true, ""},
		{"Int16Arr", LengthFunc("1"), reflect.ValueOf([1]int16{1}), true, true, ""},
		{"Int32Arr", LengthFunc("1"), reflect.ValueOf([1]int32{1}), true, true, ""},
		{"Int64Arr", LengthFunc("1"), reflect.ValueOf([1]int64{1}), true, true, ""},
		{"StringArr", LengthFunc("1"), reflect.ValueOf([1]string{"1"}), true, true, ""},

		{"RuneSlice", LengthFunc("1"), reflect.ValueOf([]rune{rune(1)}), true, true, ""},
		{"ByteSlice", LengthFunc("1"), reflect.ValueOf([]byte{byte(1)}), true, true, ""},
		{"Int8Slice", LengthFunc("1"), reflect.ValueOf([]int8{int8(1)}), true, true, ""},
		{"IntSlice", LengthFunc("1"), reflect.ValueOf([]int{1}), true, true, ""},
		{"Int16Slice", LengthFunc("1"), reflect.ValueOf([]int16{int16(1)}), true, true, ""},
		{"Int32Slice", LengthFunc("1"), reflect.ValueOf([]int32{int32(1)}), true, true, ""},
		{"Int64Slice", LengthFunc("1"), reflect.ValueOf([]int64{int64(1)}), true, true, ""},
		{"StringSlice", LengthFunc("1"), reflect.ValueOf([]string{"1"}), true, true, ""},

		{"StringPtr", LengthFunc("1"), reflect.ValueOf(stringPtr("1")), true, true, ""},

		{"RunePtrArr", LengthFunc("1"), reflect.ValueOf([1]*rune{runePtr(rune(1))}), true, true, ""},
		{"BytePtrArr", LengthFunc("1"), reflect.ValueOf([1]*byte{bytePtr(byte(1))}), true, true, ""},
		{"Int8PtrArr", LengthFunc("1"), reflect.ValueOf([1]*int8{int8Ptr(int8(1))}), true, true, ""},
		{"IntPtrArr", LengthFunc("1"), reflect.ValueOf([1]*int{intPtr(1)}), true, true, ""},
		{"Int16PtrArr", LengthFunc("1"), reflect.ValueOf([1]*int16{int16Ptr(int16(1))}), true, true, ""},
		{"Int32PtrArr", LengthFunc("1"), reflect.ValueOf([1]*int32{int32Ptr(int32(1))}), true, true, ""},
		{"Int64PtrArr", LengthFunc("1"), reflect.ValueOf([1]*int64{int64Ptr(int64(1))}), true, true, ""},
		{"StringPtrArr", LengthFunc("1"), reflect.ValueOf([1]*string{stringPtr("1")}), true, true, ""},

		{"RunePtrSlice", LengthFunc("1"), reflect.ValueOf([]*rune{runePtr(rune(1))}), true, true, ""},
		{"BytePtrSlice", LengthFunc("1"), reflect.ValueOf([]*byte{bytePtr(byte(1))}), true, true, ""},
		{"Int8PtrSlice", LengthFunc("1"), reflect.ValueOf([]*int8{int8Ptr(int8(1))}), true, true, ""},
		{"IntPtrSlice", LengthFunc("1"), reflect.ValueOf([]*int{intPtr(1)}), true, true, ""},
		{"Int16PtrSlice", LengthFunc("1"), reflect.ValueOf([]*int16{int16Ptr(int16(1))}), true, true, ""},
		{"Int32PtrSlice", LengthFunc("1"), reflect.ValueOf([]*int32{int32Ptr(1)}), true, true, ""},
		{"Int64PtrSlice", LengthFunc("1"), reflect.ValueOf([]*int64{int64Ptr(int64(1))}), true, true, ""},
		{"StringPtrSlice", LengthFunc("1"), reflect.ValueOf([]*string{stringPtr("1")}), true, true, ""},
	} {
		t.Run(s.name, s.test)
	}
}