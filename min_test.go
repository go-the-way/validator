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

func TestNewMinFunc(t *testing.T) {
	if v := MinFunc(""); v != nil {
		t.Fatal("test failed")
	}
}

func TestNewMinFuncPanic(t *testing.T) {
	for _, tc := range []struct {
		name string
		str  string
		err  error
	}{
		{"ParseFloat", "A", parseNumError("ParseFloat", "A")},
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
			_ = MinFunc(tc.str)
		})
	}
}

func TestMinUnPassAndEmptyMsg(t *testing.T) {
	for _, s := range []testCase{
		{"Rune", MinFunc("100"), reflect.ValueOf('A'), false, ""},
		{"Byte", MinFunc("100"), reflect.ValueOf(byte('A')), false, ""},
		{"Int8", MinFunc("100"), reflect.ValueOf(int8('A')), false, ""},
		{"Int", MinFunc("100"), reflect.ValueOf(int('A')), false, ""},
		{"Int16", MinFunc("100"), reflect.ValueOf(int16('A')), false, ""},
		{"Int32", MinFunc("100"), reflect.ValueOf(int32(byte('A'))), false, ""},
		{"Int64", MinFunc("100"), reflect.ValueOf(int64('A')), false, ""},
		{"Float32", MinFunc("100"), reflect.ValueOf(float32('A')), false, ""},
		{"Float64", MinFunc("100"), reflect.ValueOf(float64('A')), false, ""},

		{"RuneArr", MinFunc("100"), reflect.ValueOf([1]rune{'A'}), false, ""},
		{"ByteArr", MinFunc("100"), reflect.ValueOf([1]byte{'A'}), false, ""},
		{"Int8Arr", MinFunc("100"), reflect.ValueOf([1]int8{'A'}), false, ""},
		{"IntArr", MinFunc("100"), reflect.ValueOf([1]int{'A'}), false, ""},
		{"Int16Arr", MinFunc("100"), reflect.ValueOf([1]int16{'A'}), false, ""},
		{"Int32Arr", MinFunc("100"), reflect.ValueOf([1]int32{int32(byte('A'))}), false, ""},
		{"Int64Arr", MinFunc("100"), reflect.ValueOf([1]int64{'A'}), false, ""},
		{"Float32Arr", MinFunc("100"), reflect.ValueOf([1]float32{float32('A')}), false, ""},
		{"Float64Arr", MinFunc("100"), reflect.ValueOf([1]float64{float64('A')}), false, ""},

		{"RuneSlice", MinFunc("100"), reflect.ValueOf([]rune{'A'}), false, ""},
		{"ByteSlice", MinFunc("100"), reflect.ValueOf([]byte{'A'}), false, ""},
		{"Int8Slice", MinFunc("100"), reflect.ValueOf([]int8{'A'}), false, ""},
		{"IntSlice", MinFunc("100"), reflect.ValueOf([]int{'A'}), false, ""},
		{"Int16Slice", MinFunc("100"), reflect.ValueOf([]int16{'A'}), false, ""},
		{"Int32Slice", MinFunc("100"), reflect.ValueOf([]int32{int32(byte('A'))}), false, ""},
		{"Int64Slice", MinFunc("100"), reflect.ValueOf([]int64{'A'}), false, ""},
		{"Float32Slice", MinFunc("100"), reflect.ValueOf([]float32{float32('A')}), false, ""},
		{"Float64Slice", MinFunc("100"), reflect.ValueOf([]float64{float64('A')}), false, ""},

		{"RunePtr", MinFunc("100"), reflect.ValueOf(runePtr('A')), false, ""},
		{"BytePtr", MinFunc("100"), reflect.ValueOf(bytePtr(byte('A'))), false, ""},
		{"Int8Ptr", MinFunc("100"), reflect.ValueOf(int8Ptr(int8('A'))), false, ""},
		{"IntPtr", MinFunc("100"), reflect.ValueOf(intPtr(int('A'))), false, ""},
		{"Int16Ptr", MinFunc("100"), reflect.ValueOf(int16Ptr(int16('A'))), false, ""},
		{"Int32Ptr", MinFunc("100"), reflect.ValueOf(int32Ptr(int32(byte('A')))), false, ""},
		{"Int64Ptr", MinFunc("100"), reflect.ValueOf(int64Ptr(int64('A'))), false, ""},
		{"Float32Ptr", MinFunc("100"), reflect.ValueOf(float32Ptr(float32('A'))), false, ""},
		{"Float64Ptr", MinFunc("100"), reflect.ValueOf(float64Ptr(float64('A'))), false, ""},

		{"RunePtrArr", MinFunc("100"), reflect.ValueOf([1]*rune{runePtr('A')}), false, ""},
		{"BytePtrArr", MinFunc("100"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), false, ""},
		{"Int8PtrArr", MinFunc("100"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), false, ""},
		{"IntPtrArr", MinFunc("100"), reflect.ValueOf([1]*int{intPtr(int('A'))}), false, ""},
		{"Int16PtrArr", MinFunc("100"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), false, ""},
		{"Int32PtrArr", MinFunc("100"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), false, ""},
		{"Int64PtrArr", MinFunc("100"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), false, ""},
		{"Float32PtrArr", MinFunc("100"), reflect.ValueOf([1]*float32{float32Ptr(float32('A'))}), false, ""},
		{"Float64PtrArr", MinFunc("100"), reflect.ValueOf([1]*float64{float64Ptr(float64('A'))}), false, ""},

		{"RunePtrSlice", MinFunc("100"), reflect.ValueOf([]*rune{runePtr('A')}), false, ""},
		{"BytePtrSlice", MinFunc("100"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), false, ""},
		{"Int8PtrSlice", MinFunc("100"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), false, ""},
		{"IntPtrSlice", MinFunc("100"), reflect.ValueOf([]*int{intPtr(int('A'))}), false, ""},
		{"Int16PtrSlice", MinFunc("100"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), false, ""},
		{"Int32PtrSlice", MinFunc("100"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), false, ""},
		{"Int64PtrSlice", MinFunc("100"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), false, ""},
		{"Float32PtrSlice", MinFunc("100"), reflect.ValueOf([]*float32{float32Ptr(float32('A'))}), false, ""},
		{"Float64PtrSlice", MinFunc("100"), reflect.ValueOf([]*float64{float64Ptr(float64('A'))}), false, ""},
	} {
		t.Run(s.name, s.test)
	}
}

func TestMinUnPassAndMsg(t *testing.T) {
	for _, s := range []testCase{
		{"Rune", MinFunc("100,fail"), reflect.ValueOf('A'), false, "fail"},
		{"Byte", MinFunc("100,fail"), reflect.ValueOf(byte('A')), false, "fail"},
		{"Int8", MinFunc("100,fail"), reflect.ValueOf(int8('A')), false, "fail"},
		{"Int", MinFunc("100,fail"), reflect.ValueOf(int('A')), false, "fail"},
		{"Int16", MinFunc("100,fail"), reflect.ValueOf(int16('A')), false, "fail"},
		{"Int32", MinFunc("100,fail"), reflect.ValueOf(int32(byte('A'))), false, "fail"},
		{"Int64", MinFunc("100,fail"), reflect.ValueOf(int64('A')), false, "fail"},
		{"Float32", MinFunc("100,fail"), reflect.ValueOf(float32('A')), false, "fail"},
		{"Float64", MinFunc("100,fail"), reflect.ValueOf(float64('A')), false, "fail"},

		{"RuneArr", MinFunc("100,fail"), reflect.ValueOf([1]rune{'A'}), false, "fail"},
		{"ByteArr", MinFunc("100,fail"), reflect.ValueOf([1]byte{'A'}), false, "fail"},
		{"Int8Arr", MinFunc("100,fail"), reflect.ValueOf([1]int8{'A'}), false, "fail"},
		{"IntArr", MinFunc("100,fail"), reflect.ValueOf([1]int{'A'}), false, "fail"},
		{"Int16Arr", MinFunc("100,fail"), reflect.ValueOf([1]int16{'A'}), false, "fail"},
		{"Int32Arr", MinFunc("100,fail"), reflect.ValueOf([1]int32{int32(byte('A'))}), false, "fail"},
		{"Int64Arr", MinFunc("100,fail"), reflect.ValueOf([1]int64{'A'}), false, "fail"},
		{"Float32Arr", MinFunc("100,fail"), reflect.ValueOf([1]float32{float32('A')}), false, "fail"},
		{"Float64Arr", MinFunc("100,fail"), reflect.ValueOf([1]float64{float64('A')}), false, "fail"},

		{"RuneSlice", MinFunc("100,fail"), reflect.ValueOf([]rune{'A'}), false, "fail"},
		{"ByteSlice", MinFunc("100,fail"), reflect.ValueOf([]byte{'A'}), false, "fail"},
		{"Int8Slice", MinFunc("100,fail"), reflect.ValueOf([]int8{'A'}), false, "fail"},
		{"IntSlice", MinFunc("100,fail"), reflect.ValueOf([]int{'A'}), false, "fail"},
		{"Int16Slice", MinFunc("100,fail"), reflect.ValueOf([]int16{'A'}), false, "fail"},
		{"Int32Slice", MinFunc("100,fail"), reflect.ValueOf([]int32{int32(byte('A'))}), false, "fail"},
		{"Int64Slice", MinFunc("100,fail"), reflect.ValueOf([]int64{'A'}), false, "fail"},
		{"Float32Slice", MinFunc("100,fail"), reflect.ValueOf([]float32{float32('A')}), false, "fail"},
		{"Float64Slice", MinFunc("100,fail"), reflect.ValueOf([]float64{float64('A')}), false, "fail"},

		{"RunePtr", MinFunc("100,fail"), reflect.ValueOf(runePtr('A')), false, "fail"},
		{"BytePtr", MinFunc("100,fail"), reflect.ValueOf(bytePtr(byte('A'))), false, "fail"},
		{"Int8Ptr", MinFunc("100,fail"), reflect.ValueOf(int8Ptr(int8('A'))), false, "fail"},
		{"IntPtr", MinFunc("100,fail"), reflect.ValueOf(intPtr(int('A'))), false, "fail"},
		{"Int16Ptr", MinFunc("100,fail"), reflect.ValueOf(int16Ptr(int16('A'))), false, "fail"},
		{"Int32Ptr", MinFunc("100,fail"), reflect.ValueOf(int32Ptr(int32(byte('A')))), false, "fail"},
		{"Int64Ptr", MinFunc("100,fail"), reflect.ValueOf(int64Ptr(int64('A'))), false, "fail"},
		{"Float32Ptr", MinFunc("100,fail"), reflect.ValueOf(float32Ptr(float32('A'))), false, "fail"},
		{"Float64Ptr", MinFunc("100,fail"), reflect.ValueOf(float64Ptr(float64('A'))), false, "fail"},

		{"RunePtrArr", MinFunc("100,fail"), reflect.ValueOf([1]*rune{runePtr('A')}), false, "fail"},
		{"BytePtrArr", MinFunc("100,fail"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), false, "fail"},
		{"Int8PtrArr", MinFunc("100,fail"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), false, "fail"},
		{"IntPtrArr", MinFunc("100,fail"), reflect.ValueOf([1]*int{intPtr(int('A'))}), false, "fail"},
		{"Int16PtrArr", MinFunc("100,fail"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), false, "fail"},
		{"Int32PtrArr", MinFunc("100,fail"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), false, "fail"},
		{"Int64PtrArr", MinFunc("100,fail"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), false, "fail"},
		{"Float32PtrArr", MinFunc("100,fail"), reflect.ValueOf([1]*float32{float32Ptr(float32('A'))}), false, "fail"},
		{"Float64PtrArr", MinFunc("100,fail"), reflect.ValueOf([1]*float64{float64Ptr(float64('A'))}), false, "fail"},

		{"RunePtrSlice", MinFunc("100,fail"), reflect.ValueOf([]*rune{runePtr('A')}), false, "fail"},
		{"BytePtrSlice", MinFunc("100,fail"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), false, "fail"},
		{"Int8PtrSlice", MinFunc("100,fail"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), false, "fail"},
		{"IntPtrSlice", MinFunc("100,fail"), reflect.ValueOf([]*int{intPtr(int('A'))}), false, "fail"},
		{"Int16PtrSlice", MinFunc("100,fail"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), false, "fail"},
		{"Int32PtrSlice", MinFunc("100,fail"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), false, "fail"},
		{"Int64PtrSlice", MinFunc("100,fail"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), false, "fail"},
		{"Float32PtrSlice", MinFunc("100,fail"), reflect.ValueOf([]*float32{float32Ptr(float32('A'))}), false, "fail"},
		{"Float64PtrSlice", MinFunc("100,fail"), reflect.ValueOf([]*float64{float64Ptr(float64('A'))}), false, "fail"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestMinUnPassAndChineseMsg(t *testing.T) {
	for _, s := range []testCase{
		{"Rune", MinFunc("100,错误"), reflect.ValueOf('A'), false, "错误"},
		{"Byte", MinFunc("100,错误"), reflect.ValueOf(byte('A')), false, "错误"},
		{"Int8", MinFunc("100,错误"), reflect.ValueOf(int8('A')), false, "错误"},
		{"Int", MinFunc("100,错误"), reflect.ValueOf(int('A')), false, "错误"},
		{"Int16", MinFunc("100,错误"), reflect.ValueOf(int16('A')), false, "错误"},
		{"Int32", MinFunc("100,错误"), reflect.ValueOf(int32(byte('A'))), false, "错误"},
		{"Int64", MinFunc("100,错误"), reflect.ValueOf(int64('A')), false, "错误"},
		{"Float32", MinFunc("100,错误"), reflect.ValueOf(float32('A')), false, "错误"},
		{"Float64", MinFunc("100,错误"), reflect.ValueOf(float64('A')), false, "错误"},

		{"RuneArr", MinFunc("100,错误"), reflect.ValueOf([1]rune{'A'}), false, "错误"},
		{"ByteArr", MinFunc("100,错误"), reflect.ValueOf([1]byte{'A'}), false, "错误"},
		{"Int8Arr", MinFunc("100,错误"), reflect.ValueOf([1]int8{'A'}), false, "错误"},
		{"IntArr", MinFunc("100,错误"), reflect.ValueOf([1]int{'A'}), false, "错误"},
		{"Int16Arr", MinFunc("100,错误"), reflect.ValueOf([1]int16{'A'}), false, "错误"},
		{"Int32Arr", MinFunc("100,错误"), reflect.ValueOf([1]int32{int32(byte('A'))}), false, "错误"},
		{"Int64Arr", MinFunc("100,错误"), reflect.ValueOf([1]int64{'A'}), false, "错误"},
		{"Float32Arr", MinFunc("100,错误"), reflect.ValueOf([1]float32{float32('A')}), false, "错误"},
		{"Float64Arr", MinFunc("100,错误"), reflect.ValueOf([1]float64{float64('A')}), false, "错误"},

		{"RuneSlice", MinFunc("100,错误"), reflect.ValueOf([]rune{'A'}), false, "错误"},
		{"ByteSlice", MinFunc("100,错误"), reflect.ValueOf([]byte{'A'}), false, "错误"},
		{"Int8Slice", MinFunc("100,错误"), reflect.ValueOf([]int8{'A'}), false, "错误"},
		{"IntSlice", MinFunc("100,错误"), reflect.ValueOf([]int{'A'}), false, "错误"},
		{"Int16Slice", MinFunc("100,错误"), reflect.ValueOf([]int16{'A'}), false, "错误"},
		{"Int32Slice", MinFunc("100,错误"), reflect.ValueOf([]int32{int32(byte('A'))}), false, "错误"},
		{"Int64Slice", MinFunc("100,错误"), reflect.ValueOf([]int64{'A'}), false, "错误"},
		{"Float32Slice", MinFunc("100,错误"), reflect.ValueOf([]float32{float32('A')}), false, "错误"},
		{"Float64Slice", MinFunc("100,错误"), reflect.ValueOf([]float64{float64('A')}), false, "错误"},

		{"RunePtr", MinFunc("100,错误"), reflect.ValueOf(runePtr('A')), false, "错误"},
		{"BytePtr", MinFunc("100,错误"), reflect.ValueOf(bytePtr(byte('A'))), false, "错误"},
		{"Int8Ptr", MinFunc("100,错误"), reflect.ValueOf(int8Ptr(int8('A'))), false, "错误"},
		{"IntPtr", MinFunc("100,错误"), reflect.ValueOf(intPtr(int('A'))), false, "错误"},
		{"Int16Ptr", MinFunc("100,错误"), reflect.ValueOf(int16Ptr(int16('A'))), false, "错误"},
		{"Int32Ptr", MinFunc("100,错误"), reflect.ValueOf(int32Ptr(int32(byte('A')))), false, "错误"},
		{"Int64Ptr", MinFunc("100,错误"), reflect.ValueOf(int64Ptr(int64('A'))), false, "错误"},
		{"Float32Ptr", MinFunc("100,错误"), reflect.ValueOf(float32Ptr(float32('A'))), false, "错误"},
		{"Float64Ptr", MinFunc("100,错误"), reflect.ValueOf(float64Ptr(float64('A'))), false, "错误"},

		{"RunePtrArr", MinFunc("100,错误"), reflect.ValueOf([1]*rune{runePtr('A')}), false, "错误"},
		{"BytePtrArr", MinFunc("100,错误"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), false, "错误"},
		{"Int8PtrArr", MinFunc("100,错误"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), false, "错误"},
		{"IntPtrArr", MinFunc("100,错误"), reflect.ValueOf([1]*int{intPtr(int('A'))}), false, "错误"},
		{"Int16PtrArr", MinFunc("100,错误"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), false, "错误"},
		{"Int32PtrArr", MinFunc("100,错误"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), false, "错误"},
		{"Int64PtrArr", MinFunc("100,错误"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), false, "错误"},
		{"Float32PtrArr", MinFunc("100,错误"), reflect.ValueOf([1]*float32{float32Ptr(float32('A'))}), false, "错误"},
		{"Float64PtrArr", MinFunc("100,错误"), reflect.ValueOf([1]*float64{float64Ptr(float64('A'))}), false, "错误"},

		{"RunePtrSlice", MinFunc("100,错误"), reflect.ValueOf([]*rune{runePtr('A')}), false, "错误"},
		{"BytePtrSlice", MinFunc("100,错误"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), false, "错误"},
		{"Int8PtrSlice", MinFunc("100,错误"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), false, "错误"},
		{"IntPtrSlice", MinFunc("100,错误"), reflect.ValueOf([]*int{intPtr(int('A'))}), false, "错误"},
		{"Int16PtrSlice", MinFunc("100,错误"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), false, "错误"},
		{"Int32PtrSlice", MinFunc("100,错误"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), false, "错误"},
		{"Int64PtrSlice", MinFunc("100,错误"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), false, "错误"},
		{"Float32PtrSlice", MinFunc("100,错误"), reflect.ValueOf([]*float32{float32Ptr(float32('A'))}), false, "错误"},
		{"Float64PtrSlice", MinFunc("100,错误"), reflect.ValueOf([]*float64{float64Ptr(float64('A'))}), false, "错误"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestMinUnPassAndSpecialEffectMsg(t *testing.T) {
	for _, s := range []testCase{
		{"Rune", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf('A'), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Byte", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(byte('A')), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int8", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(int8('A')), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(int('A')), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int16", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(int16('A')), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int32", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(int32(byte('A'))), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int64", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(int64('A')), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float32", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(float32('A')), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float64", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(float64('A')), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},

		{"RuneArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]rune{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"ByteArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]byte{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int8Arr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]int8{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"IntArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]int{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int16Arr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]int16{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int32Arr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]int32{int32(byte('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int64Arr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]int64{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float32Arr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]float32{float32('A')}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float64Arr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]float64{float64('A')}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},

		{"RuneSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]rune{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"ByteSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]byte{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int8Slice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]int8{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"IntSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]int{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int16Slice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]int16{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int32Slice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]int32{int32(byte('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int64Slice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]int64{'A'}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float32Slice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]float32{float32('A')}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float64Slice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]float64{float64('A')}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},

		{"RunePtr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(runePtr('A')), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"BytePtr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(bytePtr(byte('A'))), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int8Ptr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(int8Ptr(int8('A'))), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"IntPtr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(intPtr(int('A'))), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int16Ptr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(int16Ptr(int16('A'))), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int32Ptr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(int32Ptr(int32(byte('A')))), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int64Ptr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(int64Ptr(int64('A'))), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float32Ptr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(float32Ptr(float32('A'))), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float64Ptr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(float64Ptr(float64('A'))), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},

		{"RunePtrArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*rune{runePtr('A')}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"BytePtrArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int8PtrArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"IntPtrArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*int{intPtr(int('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int16PtrArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int32PtrArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int64PtrArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float32PtrArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*float32{float32Ptr(float32('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float64PtrArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*float64{float64Ptr(float64('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},

		{"RunePtrSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*rune{runePtr('A')}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"BytePtrSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int8PtrSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"IntPtrSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*int{intPtr(int('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int16PtrSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int32PtrSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int64PtrSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float32PtrSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*float32{float32Ptr(float32('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float64PtrSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*float64{float64Ptr(float64('A'))}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestMinUnPassAndArrowMsg(t *testing.T) {
	for _, s := range []testCase{
		{"Rune", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf('A'), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Byte", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(byte('A')), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int8", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(int8('A')), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(int('A')), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int16", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(int16('A')), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int32", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(int32(byte('A'))), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int64", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(int64('A')), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float32", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(float32('A')), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float64", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(float64('A')), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},

		{"RuneArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]rune{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"ByteArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]byte{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int8Arr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]int8{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"IntArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]int{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int16Arr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]int16{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int32Arr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]int32{int32(byte('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int64Arr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]int64{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float32Arr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]float32{float32('A')}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float64Arr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]float64{float64('A')}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},

		{"RuneSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]rune{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"ByteSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]byte{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int8Slice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]int8{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"IntSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]int{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int16Slice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]int16{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int32Slice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]int32{int32(byte('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int64Slice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]int64{'A'}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float32Slice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]float32{float32('A')}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float64Slice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]float64{float64('A')}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},

		{"RunePtr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(runePtr('A')), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"BytePtr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(bytePtr(byte('A'))), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int8Ptr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(int8Ptr(int8('A'))), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"IntPtr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(intPtr(int('A'))), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int16Ptr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(int16Ptr(int16('A'))), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int32Ptr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(int32Ptr(int32(byte('A')))), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int64Ptr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(int64Ptr(int64('A'))), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float32Ptr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(float32Ptr(float32('A'))), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float64Ptr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(float64Ptr(float64('A'))), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},

		{"RunePtrArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*rune{runePtr('A')}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"BytePtrArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int8PtrArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"IntPtrArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*int{intPtr(int('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int16PtrArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int32PtrArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int64PtrArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float32PtrArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*float32{float32Ptr(float32('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float64PtrArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*float64{float64Ptr(float64('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},

		{"RunePtrSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*rune{runePtr('A')}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"BytePtrSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int8PtrSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"IntPtrSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*int{intPtr(int('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int16PtrSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int32PtrSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int64PtrSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float32PtrSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*float32{float32Ptr(float32('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float64PtrSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*float64{float64Ptr(float64('A'))}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestMinUnPassAndForeignMsg(t *testing.T) {
	for _, s := range []testCase{
		{"Rune", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf('A'), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Byte", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(byte('A')), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int8", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(int8('A')), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(int('A')), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int16", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(int16('A')), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int32", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(int32(byte('A'))), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int64", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(int64('A')), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float32", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(float32('A')), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float64", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(float64('A')), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},

		{"RuneArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]rune{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"ByteArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]byte{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int8Arr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]int8{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"IntArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]int{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int16Arr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]int16{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int32Arr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]int32{int32(byte('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int64Arr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]int64{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float32Arr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]float32{float32('A')}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float64Arr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]float64{float64('A')}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},

		{"RuneSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]rune{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"ByteSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]byte{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int8Slice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]int8{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"IntSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]int{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int16Slice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]int16{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int32Slice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]int32{int32(byte('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int64Slice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]int64{'A'}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float32Slice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]float32{float32('A')}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float64Slice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]float64{float64('A')}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},

		{"RunePtr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(runePtr('A')), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"BytePtr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(bytePtr(byte('A'))), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int8Ptr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(int8Ptr(int8('A'))), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"IntPtr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(intPtr(int('A'))), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int16Ptr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(int16Ptr(int16('A'))), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int32Ptr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(int32Ptr(int32(byte('A')))), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int64Ptr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(int64Ptr(int64('A'))), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float32Ptr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(float32Ptr(float32('A'))), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float64Ptr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(float64Ptr(float64('A'))), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},

		{"RunePtrArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*rune{runePtr('A')}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"BytePtrArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int8PtrArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"IntPtrArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*int{intPtr(int('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int16PtrArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int32PtrArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int64PtrArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float32PtrArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*float32{float32Ptr(float32('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float64PtrArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*float64{float64Ptr(float64('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},

		{"RunePtrSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*rune{runePtr('A')}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"BytePtrSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int8PtrSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"IntPtrSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*int{intPtr(int('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int16PtrSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int32PtrSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int64PtrSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float32PtrSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*float32{float32Ptr(float32('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float64PtrSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*float64{float64Ptr(float64('A'))}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestMinUnPassAndSymbolMsg(t *testing.T) {
	for _, s := range []testCase{
		{"Rune", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf('A'), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Byte", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(byte('A')), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int8", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(int8('A')), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(int('A')), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int16", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(int16('A')), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int32", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(int32(byte('A'))), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int64", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(int64('A')), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float32", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(float32('A')), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float64", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(float64('A')), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},

		{"RuneArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]rune{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"ByteArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]byte{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int8Arr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]int8{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"IntArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]int{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int16Arr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]int16{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int32Arr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]int32{int32(byte('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int64Arr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]int64{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float32Arr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]float32{float32('A')}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float64Arr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]float64{float64('A')}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},

		{"RuneSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]rune{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"ByteSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]byte{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int8Slice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]int8{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"IntSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]int{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int16Slice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]int16{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int32Slice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]int32{int32(byte('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int64Slice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]int64{'A'}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float32Slice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]float32{float32('A')}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float64Slice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]float64{float64('A')}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},

		{"RunePtr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(runePtr('A')), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"BytePtr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(bytePtr(byte('A'))), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int8Ptr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(int8Ptr(int8('A'))), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"IntPtr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(intPtr(int('A'))), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int16Ptr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(int16Ptr(int16('A'))), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int32Ptr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(int32Ptr(int32(byte('A')))), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int64Ptr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(int64Ptr(int64('A'))), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float32Ptr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(float32Ptr(float32('A'))), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float64Ptr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(float64Ptr(float64('A'))), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},

		{"RunePtrArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*rune{runePtr('A')}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"BytePtrArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int8PtrArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"IntPtrArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*int{intPtr(int('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int16PtrArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int32PtrArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int64PtrArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float32PtrArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*float32{float32Ptr(float32('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float64PtrArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*float64{float64Ptr(float64('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},

		{"RunePtrSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*rune{runePtr('A')}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"BytePtrSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int8PtrSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"IntPtrSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*int{intPtr(int('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int16PtrSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int32PtrSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int64PtrSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float32PtrSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*float32{float32Ptr(float32('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float64PtrSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*float64{float64Ptr(float64('A'))}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestMinPass(t *testing.T) {
	for _, s := range []testCase{
		{"Rune", MinFunc("1"), reflect.ValueOf(rune(1)), true, ""},
		{"Byte", MinFunc("1"), reflect.ValueOf(byte(1)), true, ""},
		{"Int8", MinFunc("1"), reflect.ValueOf(int8(1)), true, ""},
		{"Int", MinFunc("1"), reflect.ValueOf(1), true, ""},
		{"Int16", MinFunc("1"), reflect.ValueOf(int16(1)), true, ""},
		{"Int32", MinFunc("1"), reflect.ValueOf(int32(1)), true, ""},
		{"Int64", MinFunc("1"), reflect.ValueOf(int64(1)), true, ""},
		{"Float32", MinFunc("1"), reflect.ValueOf(float32(1)), true, ""},
		{"Float64", MinFunc("1"), reflect.ValueOf(float64(1)), true, ""},

		{"RuneArr", MinFunc("1"), reflect.ValueOf([1]rune{1}), true, ""},
		{"ByteArr", MinFunc("1"), reflect.ValueOf([1]byte{1}), true, ""},
		{"Int8Arr", MinFunc("1"), reflect.ValueOf([1]int8{1}), true, ""},
		{"IntArr", MinFunc("1"), reflect.ValueOf([1]int{1}), true, ""},
		{"Int16Arr", MinFunc("1"), reflect.ValueOf([1]int16{1}), true, ""},
		{"Int32Arr", MinFunc("1"), reflect.ValueOf([1]int32{1}), true, ""},
		{"Int64Arr", MinFunc("1"), reflect.ValueOf([1]int64{1}), true, ""},
		{"Float32Arr", MinFunc("1"), reflect.ValueOf([1]float32{float32(1)}), true, ""},
		{"Float64Arr", MinFunc("1"), reflect.ValueOf([1]float64{float64(1)}), true, ""},

		{"RuneSlice", MinFunc("1"), reflect.ValueOf([]rune{rune(1)}), true, ""},
		{"ByteSlice", MinFunc("1"), reflect.ValueOf([]byte{byte(1)}), true, ""},
		{"Int8Slice", MinFunc("1"), reflect.ValueOf([]int8{int8(1)}), true, ""},
		{"IntSlice", MinFunc("1"), reflect.ValueOf([]int{1}), true, ""},
		{"Int16Slice", MinFunc("1"), reflect.ValueOf([]int16{int16(1)}), true, ""},
		{"Int32Slice", MinFunc("1"), reflect.ValueOf([]int32{int32(1)}), true, ""},
		{"Int64Slice", MinFunc("1"), reflect.ValueOf([]int64{int64(1)}), true, ""},
		{"Float32Slice", MinFunc("1"), reflect.ValueOf([]float32{float32(1)}), true, ""},
		{"Float64Slice", MinFunc("1"), reflect.ValueOf([]float64{float64(1)}), true, ""},

		{"RunePtr", MinFunc("1"), reflect.ValueOf(runePtr(rune(1))), true, ""},
		{"BytePtr", MinFunc("1"), reflect.ValueOf(bytePtr(byte(1))), true, ""},
		{"Int8Ptr", MinFunc("1"), reflect.ValueOf(int8Ptr(int8(1))), true, ""},
		{"IntPtr", MinFunc("1"), reflect.ValueOf(intPtr(1)), true, ""},
		{"Int16Ptr", MinFunc("1"), reflect.ValueOf(int16Ptr(int16(1))), true, ""},
		{"Int32Ptr", MinFunc("1"), reflect.ValueOf(int32Ptr(1)), true, ""},
		{"Int64Ptr", MinFunc("1"), reflect.ValueOf(int64Ptr(int64(1))), true, ""},
		{"Float32Ptr", MinFunc("1"), reflect.ValueOf(float32Ptr(float32(1))), true, ""},
		{"Float64Ptr", MinFunc("1"), reflect.ValueOf(float64Ptr(float64(1))), true, ""},

		{"RunePtrArr", MinFunc("1"), reflect.ValueOf([1]*rune{runePtr(rune(1))}), true, ""},
		{"BytePtrArr", MinFunc("1"), reflect.ValueOf([1]*byte{bytePtr(byte(1))}), true, ""},
		{"Int8PtrArr", MinFunc("1"), reflect.ValueOf([1]*int8{int8Ptr(int8(1))}), true, ""},
		{"IntPtrArr", MinFunc("1"), reflect.ValueOf([1]*int{intPtr(1)}), true, ""},
		{"Int16PtrArr", MinFunc("1"), reflect.ValueOf([1]*int16{int16Ptr(int16(1))}), true, ""},
		{"Int32PtrArr", MinFunc("1"), reflect.ValueOf([1]*int32{int32Ptr(int32(1))}), true, ""},
		{"Int64PtrArr", MinFunc("1"), reflect.ValueOf([1]*int64{int64Ptr(int64(1))}), true, ""},
		{"Float32PtrArr", MinFunc("1"), reflect.ValueOf([1]*float32{float32Ptr(float32(1))}), true, ""},
		{"Float64PtrArr", MinFunc("1"), reflect.ValueOf([1]*float64{float64Ptr(float64(1))}), true, ""},

		{"RunePtrSlice", MinFunc("1"), reflect.ValueOf([]*rune{runePtr(rune(1))}), true, ""},
		{"BytePtrSlice", MinFunc("1"), reflect.ValueOf([]*byte{bytePtr(byte(1))}), true, ""},
		{"Int8PtrSlice", MinFunc("1"), reflect.ValueOf([]*int8{int8Ptr(int8(1))}), true, ""},
		{"IntPtrSlice", MinFunc("1"), reflect.ValueOf([]*int{intPtr(1)}), true, ""},
		{"Int16PtrSlice", MinFunc("1"), reflect.ValueOf([]*int16{int16Ptr(int16(1))}), true, ""},
		{"Int32PtrSlice", MinFunc("1"), reflect.ValueOf([]*int32{int32Ptr(1)}), true, ""},
		{"Int64PtrSlice", MinFunc("1"), reflect.ValueOf([]*int64{int64Ptr(int64(1))}), true, ""},
		{"Float32PtrSlice", MinFunc("1"), reflect.ValueOf([]*float32{float32Ptr(float32(1))}), true, ""},
		{"Float64PtrSlice", MinFunc("1"), reflect.ValueOf([]*float64{float64Ptr(float64(1))}), true, ""},
	} {
		t.Run(s.name, s.test)
	}
}
