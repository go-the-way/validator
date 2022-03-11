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

func TestMinNoAccept(t *testing.T) {
	for _, s := range []testCase{
		{"Struct", MinFunc("100"), reflect.ValueOf(struct{}{}), false, false, ""},
		{"BoolTrue", MinFunc("100"), reflect.ValueOf(true), false, false, ""},
		{"BoolFalse", MinFunc("100"), reflect.ValueOf(false), false, false, ""},
		{"Uintptr", MinFunc("100"), reflect.ValueOf(uintptr(100)), false, false, ""},
		{"Complex64", MinFunc("100"), reflect.ValueOf(complex64(100)), false, false, ""},
		{"Complex128", MinFunc("100"), reflect.ValueOf(complex128(100)), false, false, ""},
		{"Complex", MinFunc("100"), reflect.ValueOf(complex(10100, 100)), false, false, ""},
		{"Chan", MinFunc("100"), reflect.ValueOf(make(chan bool)), false, false, ""},
		{"Map", MinFunc("100"), reflect.ValueOf(make(map[struct{}]struct{})), false, false, ""},

		{"StructArr", MinFunc("100"), reflect.ValueOf([1]struct{}{}), false, false, ""},
		{"BoolTrueArr", MinFunc("100"), reflect.ValueOf([1]bool{true}), false, false, ""},
		{"BoolFalseArr", MinFunc("100"), reflect.ValueOf([1]bool{false}), false, false, ""},
		{"UintptrArr", MinFunc("100"), reflect.ValueOf([1]uintptr{100}), false, false, ""},
		{"Complex64Arr", MinFunc("100"), reflect.ValueOf([1]complex64{100}), false, false, ""},
		{"Complex128Arr", MinFunc("100"), reflect.ValueOf([1]complex128{100}), false, false, ""},
		{"ChanArr", MinFunc("100"), reflect.ValueOf([1]chan bool{make(chan bool)}), false, false, ""},
		{"MapArr", MinFunc("100"), reflect.ValueOf([1]map[struct{}]struct{}{make(map[struct{}]struct{})}), false, false, ""},

		{"StructSlice", MinFunc("100"), reflect.ValueOf([]struct{}{}), false, false, ""},
		{"StructPtrSlice", MinFunc("100"), reflect.ValueOf([]*struct{}{}), false, false, ""},
		{"BoolTrueSlice", MinFunc("100"), reflect.ValueOf([]bool{true}), false, false, ""},
		{"BoolFalseSlice", MinFunc("100"), reflect.ValueOf([]bool{false}), false, false, ""},
		{"UintptrSlice", MinFunc("100"), reflect.ValueOf([]uintptr{100}), false, false, ""},
		{"Complex64Slice", MinFunc("100"), reflect.ValueOf([]complex64{100}), false, false, ""},
		{"Complex128Slice", MinFunc("100"), reflect.ValueOf([]complex128{100}), false, false, ""},
		{"ChanSlice", MinFunc("100"), reflect.ValueOf([]chan bool{make(chan bool)}), false, false, ""},
		{"MapSlice", MinFunc("100"), reflect.ValueOf([]map[struct{}]struct{}{make(map[struct{}]struct{})}), false, false, ""},

		{"StructPtr", MinFunc("100"), reflect.ValueOf(&struct{}{}), false, false, ""},
		{"BoolTruePtr", MinFunc("100"), reflect.ValueOf(boolPtr(true)), false, false, ""},
		{"BoolFalsePtr", MinFunc("100"), reflect.ValueOf(boolPtr(false)), false, false, ""},
		{"UintptrPtr", MinFunc("100"), reflect.ValueOf(uintptrPtr(uintptr(100))), false, false, ""},
		{"Complex64Ptr", MinFunc("100"), reflect.ValueOf(complex64Ptr(complex64(100))), false, false, ""},
		{"Complex128Ptr", MinFunc("100"), reflect.ValueOf(complex128Ptr(complex128(100))), false, false, ""},
		{"ChanPtr", MinFunc("100"), reflect.ValueOf(chanPtr(make(chan interface{}))), false, false, ""},
		{"MapPtr", MinFunc("100"), reflect.ValueOf(mapPtr(make(map[interface{}]interface{}))), false, false, ""},

		{"StructPtrArr", MinFunc("100"), reflect.ValueOf([1]*struct{}{{}}), false, false, ""},
		{"BoolTruePtrArr", MinFunc("100"), reflect.ValueOf([1]*bool{boolPtr(true)}), false, false, ""},
		{"BoolFalsePtrArr", MinFunc("100"), reflect.ValueOf([1]*bool{boolPtr(false)}), false, false, ""},
		{"UintptrPtrArr", MinFunc("100"), reflect.ValueOf([1]*uintptr{uintptrPtr(uintptr(100))}), false, false, ""},
		{"Complex64PtrArr", MinFunc("100"), reflect.ValueOf([1]*complex64{complex64Ptr(complex64(100))}), false, false, ""},
		{"Complex128PtrArr", MinFunc("100"), reflect.ValueOf([1]*complex128{complex128Ptr(complex128(100))}), false, false, ""},
		{"ChanPtrArr", MinFunc("100"), reflect.ValueOf([1]*chan interface{}{chanPtr(make(chan interface{}))}), false, false, ""},
		{"MapPtrArr", MinFunc("100"), reflect.ValueOf([1]*map[interface{}]interface{}{mapPtr(make(map[interface{}]interface{}))}), false, false, ""},

		{"StructPtrSlice", MinFunc("100"), reflect.ValueOf([]*struct{}{{}}), false, false, ""},
		{"BoolTruePtrSlice", MinFunc("100"), reflect.ValueOf([]*bool{boolPtr(true)}), false, false, ""},
		{"BoolFalsePtrSlice", MinFunc("100"), reflect.ValueOf([]*bool{boolPtr(false)}), false, false, ""},
		{"UintptrPtrSlice", MinFunc("100"), reflect.ValueOf([]*uintptr{uintptrPtr(uintptr(100))}), false, false, ""},
		{"Complex64PtrSlice", MinFunc("100"), reflect.ValueOf([]*complex64{complex64Ptr(complex64(100))}), false, false, ""},
		{"Complex128PtrSlice", MinFunc("100"), reflect.ValueOf([]*complex128{complex128Ptr(complex128(100))}), false, false, ""},
		{"ChanPtrSlice", MinFunc("100"), reflect.ValueOf([]*chan interface{}{chanPtr(make(chan interface{}))}), false, false, ""},
		{"MapPtrSlice", MinFunc("100"), reflect.ValueOf([]*map[interface{}]interface{}{mapPtr(make(map[interface{}]interface{}))}), false, false, ""},
	} {
		t.Run(s.name, s.test)
	}
}

func TestMinUnPassAndEmptyMsg(t *testing.T) {
	for _, s := range []testCase{
		{"Rune", MinFunc("100"), reflect.ValueOf('A'), true, false, ""},
		{"Byte", MinFunc("100"), reflect.ValueOf(byte('A')), true, false, ""},
		{"Int8", MinFunc("100"), reflect.ValueOf(int8('A')), true, false, ""},
		{"Int", MinFunc("100"), reflect.ValueOf(int('A')), true, false, ""},
		{"Int16", MinFunc("100"), reflect.ValueOf(int16('A')), true, false, ""},
		{"Int32", MinFunc("100"), reflect.ValueOf(int32(byte('A'))), true, false, ""},
		{"Int64", MinFunc("100"), reflect.ValueOf(int64('A')), true, false, ""},
		{"Float32", MinFunc("100"), reflect.ValueOf(float32('A')), true, false, ""},
		{"Float64", MinFunc("100"), reflect.ValueOf(float64('A')), true, false, ""},

		{"RuneArr", MinFunc("100"), reflect.ValueOf([1]rune{'A'}), true, false, ""},
		{"ByteArr", MinFunc("100"), reflect.ValueOf([1]byte{'A'}), true, false, ""},
		{"Int8Arr", MinFunc("100"), reflect.ValueOf([1]int8{'A'}), true, false, ""},
		{"IntArr", MinFunc("100"), reflect.ValueOf([1]int{'A'}), true, false, ""},
		{"Int16Arr", MinFunc("100"), reflect.ValueOf([1]int16{'A'}), true, false, ""},
		{"Int32Arr", MinFunc("100"), reflect.ValueOf([1]int32{int32(byte('A'))}), true, false, ""},
		{"Int64Arr", MinFunc("100"), reflect.ValueOf([1]int64{'A'}), true, false, ""},
		{"Float32Arr", MinFunc("100"), reflect.ValueOf([1]float32{float32('A')}), true, false, ""},
		{"Float64Arr", MinFunc("100"), reflect.ValueOf([1]float64{float64('A')}), true, false, ""},

		{"RuneSlice", MinFunc("100"), reflect.ValueOf([]rune{'A'}), true, false, ""},
		{"ByteSlice", MinFunc("100"), reflect.ValueOf([]byte{'A'}), true, false, ""},
		{"Int8Slice", MinFunc("100"), reflect.ValueOf([]int8{'A'}), true, false, ""},
		{"IntSlice", MinFunc("100"), reflect.ValueOf([]int{'A'}), true, false, ""},
		{"Int16Slice", MinFunc("100"), reflect.ValueOf([]int16{'A'}), true, false, ""},
		{"Int32Slice", MinFunc("100"), reflect.ValueOf([]int32{int32(byte('A'))}), true, false, ""},
		{"Int64Slice", MinFunc("100"), reflect.ValueOf([]int64{'A'}), true, false, ""},
		{"Float32Slice", MinFunc("100"), reflect.ValueOf([]float32{float32('A')}), true, false, ""},
		{"Float64Slice", MinFunc("100"), reflect.ValueOf([]float64{float64('A')}), true, false, ""},

		{"RunePtr", MinFunc("100"), reflect.ValueOf(runePtr('A')), true, false, ""},
		{"BytePtr", MinFunc("100"), reflect.ValueOf(bytePtr(byte('A'))), true, false, ""},
		{"Int8Ptr", MinFunc("100"), reflect.ValueOf(int8Ptr(int8('A'))), true, false, ""},
		{"IntPtr", MinFunc("100"), reflect.ValueOf(intPtr(int('A'))), true, false, ""},
		{"Int16Ptr", MinFunc("100"), reflect.ValueOf(int16Ptr(int16('A'))), true, false, ""},
		{"Int32Ptr", MinFunc("100"), reflect.ValueOf(int32Ptr(int32(byte('A')))), true, false, ""},
		{"Int64Ptr", MinFunc("100"), reflect.ValueOf(int64Ptr(int64('A'))), true, false, ""},
		{"Float32Ptr", MinFunc("100"), reflect.ValueOf(float32Ptr(float32('A'))), true, false, ""},
		{"Float64Ptr", MinFunc("100"), reflect.ValueOf(float64Ptr(float64('A'))), true, false, ""},

		{"RunePtrArr", MinFunc("100"), reflect.ValueOf([1]*rune{runePtr('A')}), true, false, ""},
		{"BytePtrArr", MinFunc("100"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), true, false, ""},
		{"Int8PtrArr", MinFunc("100"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), true, false, ""},
		{"IntPtrArr", MinFunc("100"), reflect.ValueOf([1]*int{intPtr(int('A'))}), true, false, ""},
		{"Int16PtrArr", MinFunc("100"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), true, false, ""},
		{"Int32PtrArr", MinFunc("100"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), true, false, ""},
		{"Int64PtrArr", MinFunc("100"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), true, false, ""},
		{"Float32PtrArr", MinFunc("100"), reflect.ValueOf([1]*float32{float32Ptr(float32('A'))}), true, false, ""},
		{"Float64PtrArr", MinFunc("100"), reflect.ValueOf([1]*float64{float64Ptr(float64('A'))}), true, false, ""},

		{"RunePtrSlice", MinFunc("100"), reflect.ValueOf([]*rune{runePtr('A')}), true, false, ""},
		{"BytePtrSlice", MinFunc("100"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), true, false, ""},
		{"Int8PtrSlice", MinFunc("100"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), true, false, ""},
		{"IntPtrSlice", MinFunc("100"), reflect.ValueOf([]*int{intPtr(int('A'))}), true, false, ""},
		{"Int16PtrSlice", MinFunc("100"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), true, false, ""},
		{"Int32PtrSlice", MinFunc("100"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), true, false, ""},
		{"Int64PtrSlice", MinFunc("100"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), true, false, ""},
		{"Float32PtrSlice", MinFunc("100"), reflect.ValueOf([]*float32{float32Ptr(float32('A'))}), true, false, ""},
		{"Float64PtrSlice", MinFunc("100"), reflect.ValueOf([]*float64{float64Ptr(float64('A'))}), true, false, ""},
	} {
		t.Run(s.name, s.test)
	}
}

func TestMinUnPassAndMsg(t *testing.T) {
	for _, s := range []testCase{
		{"Rune", MinFunc("100,fail"), reflect.ValueOf('A'), true, false, "fail"},
		{"Byte", MinFunc("100,fail"), reflect.ValueOf(byte('A')), true, false, "fail"},
		{"Int8", MinFunc("100,fail"), reflect.ValueOf(int8('A')), true, false, "fail"},
		{"Int", MinFunc("100,fail"), reflect.ValueOf(int('A')), true, false, "fail"},
		{"Int16", MinFunc("100,fail"), reflect.ValueOf(int16('A')), true, false, "fail"},
		{"Int32", MinFunc("100,fail"), reflect.ValueOf(int32(byte('A'))), true, false, "fail"},
		{"Int64", MinFunc("100,fail"), reflect.ValueOf(int64('A')), true, false, "fail"},
		{"Float32", MinFunc("100,fail"), reflect.ValueOf(float32('A')), true, false, "fail"},
		{"Float64", MinFunc("100,fail"), reflect.ValueOf(float64('A')), true, false, "fail"},

		{"RuneArr", MinFunc("100,fail"), reflect.ValueOf([1]rune{'A'}), true, false, "fail"},
		{"ByteArr", MinFunc("100,fail"), reflect.ValueOf([1]byte{'A'}), true, false, "fail"},
		{"Int8Arr", MinFunc("100,fail"), reflect.ValueOf([1]int8{'A'}), true, false, "fail"},
		{"IntArr", MinFunc("100,fail"), reflect.ValueOf([1]int{'A'}), true, false, "fail"},
		{"Int16Arr", MinFunc("100,fail"), reflect.ValueOf([1]int16{'A'}), true, false, "fail"},
		{"Int32Arr", MinFunc("100,fail"), reflect.ValueOf([1]int32{int32(byte('A'))}), true, false, "fail"},
		{"Int64Arr", MinFunc("100,fail"), reflect.ValueOf([1]int64{'A'}), true, false, "fail"},
		{"Float32Arr", MinFunc("100,fail"), reflect.ValueOf([1]float32{float32('A')}), true, false, "fail"},
		{"Float64Arr", MinFunc("100,fail"), reflect.ValueOf([1]float64{float64('A')}), true, false, "fail"},

		{"RuneSlice", MinFunc("100,fail"), reflect.ValueOf([]rune{'A'}), true, false, "fail"},
		{"ByteSlice", MinFunc("100,fail"), reflect.ValueOf([]byte{'A'}), true, false, "fail"},
		{"Int8Slice", MinFunc("100,fail"), reflect.ValueOf([]int8{'A'}), true, false, "fail"},
		{"IntSlice", MinFunc("100,fail"), reflect.ValueOf([]int{'A'}), true, false, "fail"},
		{"Int16Slice", MinFunc("100,fail"), reflect.ValueOf([]int16{'A'}), true, false, "fail"},
		{"Int32Slice", MinFunc("100,fail"), reflect.ValueOf([]int32{int32(byte('A'))}), true, false, "fail"},
		{"Int64Slice", MinFunc("100,fail"), reflect.ValueOf([]int64{'A'}), true, false, "fail"},
		{"Float32Slice", MinFunc("100,fail"), reflect.ValueOf([]float32{float32('A')}), true, false, "fail"},
		{"Float64Slice", MinFunc("100,fail"), reflect.ValueOf([]float64{float64('A')}), true, false, "fail"},

		{"RunePtr", MinFunc("100,fail"), reflect.ValueOf(runePtr('A')), true, false, "fail"},
		{"BytePtr", MinFunc("100,fail"), reflect.ValueOf(bytePtr(byte('A'))), true, false, "fail"},
		{"Int8Ptr", MinFunc("100,fail"), reflect.ValueOf(int8Ptr(int8('A'))), true, false, "fail"},
		{"IntPtr", MinFunc("100,fail"), reflect.ValueOf(intPtr(int('A'))), true, false, "fail"},
		{"Int16Ptr", MinFunc("100,fail"), reflect.ValueOf(int16Ptr(int16('A'))), true, false, "fail"},
		{"Int32Ptr", MinFunc("100,fail"), reflect.ValueOf(int32Ptr(int32(byte('A')))), true, false, "fail"},
		{"Int64Ptr", MinFunc("100,fail"), reflect.ValueOf(int64Ptr(int64('A'))), true, false, "fail"},
		{"Float32Ptr", MinFunc("100,fail"), reflect.ValueOf(float32Ptr(float32('A'))), true, false, "fail"},
		{"Float64Ptr", MinFunc("100,fail"), reflect.ValueOf(float64Ptr(float64('A'))), true, false, "fail"},

		{"RunePtrArr", MinFunc("100,fail"), reflect.ValueOf([1]*rune{runePtr('A')}), true, false, "fail"},
		{"BytePtrArr", MinFunc("100,fail"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), true, false, "fail"},
		{"Int8PtrArr", MinFunc("100,fail"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), true, false, "fail"},
		{"IntPtrArr", MinFunc("100,fail"), reflect.ValueOf([1]*int{intPtr(int('A'))}), true, false, "fail"},
		{"Int16PtrArr", MinFunc("100,fail"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), true, false, "fail"},
		{"Int32PtrArr", MinFunc("100,fail"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), true, false, "fail"},
		{"Int64PtrArr", MinFunc("100,fail"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), true, false, "fail"},
		{"Float32PtrArr", MinFunc("100,fail"), reflect.ValueOf([1]*float32{float32Ptr(float32('A'))}), true, false, "fail"},
		{"Float64PtrArr", MinFunc("100,fail"), reflect.ValueOf([1]*float64{float64Ptr(float64('A'))}), true, false, "fail"},

		{"RunePtrSlice", MinFunc("100,fail"), reflect.ValueOf([]*rune{runePtr('A')}), true, false, "fail"},
		{"BytePtrSlice", MinFunc("100,fail"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), true, false, "fail"},
		{"Int8PtrSlice", MinFunc("100,fail"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), true, false, "fail"},
		{"IntPtrSlice", MinFunc("100,fail"), reflect.ValueOf([]*int{intPtr(int('A'))}), true, false, "fail"},
		{"Int16PtrSlice", MinFunc("100,fail"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), true, false, "fail"},
		{"Int32PtrSlice", MinFunc("100,fail"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), true, false, "fail"},
		{"Int64PtrSlice", MinFunc("100,fail"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), true, false, "fail"},
		{"Float32PtrSlice", MinFunc("100,fail"), reflect.ValueOf([]*float32{float32Ptr(float32('A'))}), true, false, "fail"},
		{"Float64PtrSlice", MinFunc("100,fail"), reflect.ValueOf([]*float64{float64Ptr(float64('A'))}), true, false, "fail"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestMinUnPassAndChineseMsg(t *testing.T) {
	for _, s := range []testCase{
		{"Rune", MinFunc("100,错误"), reflect.ValueOf('A'), true, false, "错误"},
		{"Byte", MinFunc("100,错误"), reflect.ValueOf(byte('A')), true, false, "错误"},
		{"Int8", MinFunc("100,错误"), reflect.ValueOf(int8('A')), true, false, "错误"},
		{"Int", MinFunc("100,错误"), reflect.ValueOf(int('A')), true, false, "错误"},
		{"Int16", MinFunc("100,错误"), reflect.ValueOf(int16('A')), true, false, "错误"},
		{"Int32", MinFunc("100,错误"), reflect.ValueOf(int32(byte('A'))), true, false, "错误"},
		{"Int64", MinFunc("100,错误"), reflect.ValueOf(int64('A')), true, false, "错误"},
		{"Float32", MinFunc("100,错误"), reflect.ValueOf(float32('A')), true, false, "错误"},
		{"Float64", MinFunc("100,错误"), reflect.ValueOf(float64('A')), true, false, "错误"},

		{"RuneArr", MinFunc("100,错误"), reflect.ValueOf([1]rune{'A'}), true, false, "错误"},
		{"ByteArr", MinFunc("100,错误"), reflect.ValueOf([1]byte{'A'}), true, false, "错误"},
		{"Int8Arr", MinFunc("100,错误"), reflect.ValueOf([1]int8{'A'}), true, false, "错误"},
		{"IntArr", MinFunc("100,错误"), reflect.ValueOf([1]int{'A'}), true, false, "错误"},
		{"Int16Arr", MinFunc("100,错误"), reflect.ValueOf([1]int16{'A'}), true, false, "错误"},
		{"Int32Arr", MinFunc("100,错误"), reflect.ValueOf([1]int32{int32(byte('A'))}), true, false, "错误"},
		{"Int64Arr", MinFunc("100,错误"), reflect.ValueOf([1]int64{'A'}), true, false, "错误"},
		{"Float32Arr", MinFunc("100,错误"), reflect.ValueOf([1]float32{float32('A')}), true, false, "错误"},
		{"Float64Arr", MinFunc("100,错误"), reflect.ValueOf([1]float64{float64('A')}), true, false, "错误"},

		{"RuneSlice", MinFunc("100,错误"), reflect.ValueOf([]rune{'A'}), true, false, "错误"},
		{"ByteSlice", MinFunc("100,错误"), reflect.ValueOf([]byte{'A'}), true, false, "错误"},
		{"Int8Slice", MinFunc("100,错误"), reflect.ValueOf([]int8{'A'}), true, false, "错误"},
		{"IntSlice", MinFunc("100,错误"), reflect.ValueOf([]int{'A'}), true, false, "错误"},
		{"Int16Slice", MinFunc("100,错误"), reflect.ValueOf([]int16{'A'}), true, false, "错误"},
		{"Int32Slice", MinFunc("100,错误"), reflect.ValueOf([]int32{int32(byte('A'))}), true, false, "错误"},
		{"Int64Slice", MinFunc("100,错误"), reflect.ValueOf([]int64{'A'}), true, false, "错误"},
		{"Float32Slice", MinFunc("100,错误"), reflect.ValueOf([]float32{float32('A')}), true, false, "错误"},
		{"Float64Slice", MinFunc("100,错误"), reflect.ValueOf([]float64{float64('A')}), true, false, "错误"},

		{"RunePtr", MinFunc("100,错误"), reflect.ValueOf(runePtr('A')), true, false, "错误"},
		{"BytePtr", MinFunc("100,错误"), reflect.ValueOf(bytePtr(byte('A'))), true, false, "错误"},
		{"Int8Ptr", MinFunc("100,错误"), reflect.ValueOf(int8Ptr(int8('A'))), true, false, "错误"},
		{"IntPtr", MinFunc("100,错误"), reflect.ValueOf(intPtr(int('A'))), true, false, "错误"},
		{"Int16Ptr", MinFunc("100,错误"), reflect.ValueOf(int16Ptr(int16('A'))), true, false, "错误"},
		{"Int32Ptr", MinFunc("100,错误"), reflect.ValueOf(int32Ptr(int32(byte('A')))), true, false, "错误"},
		{"Int64Ptr", MinFunc("100,错误"), reflect.ValueOf(int64Ptr(int64('A'))), true, false, "错误"},
		{"Float32Ptr", MinFunc("100,错误"), reflect.ValueOf(float32Ptr(float32('A'))), true, false, "错误"},
		{"Float64Ptr", MinFunc("100,错误"), reflect.ValueOf(float64Ptr(float64('A'))), true, false, "错误"},

		{"RunePtrArr", MinFunc("100,错误"), reflect.ValueOf([1]*rune{runePtr('A')}), true, false, "错误"},
		{"BytePtrArr", MinFunc("100,错误"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), true, false, "错误"},
		{"Int8PtrArr", MinFunc("100,错误"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), true, false, "错误"},
		{"IntPtrArr", MinFunc("100,错误"), reflect.ValueOf([1]*int{intPtr(int('A'))}), true, false, "错误"},
		{"Int16PtrArr", MinFunc("100,错误"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), true, false, "错误"},
		{"Int32PtrArr", MinFunc("100,错误"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), true, false, "错误"},
		{"Int64PtrArr", MinFunc("100,错误"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), true, false, "错误"},
		{"Float32PtrArr", MinFunc("100,错误"), reflect.ValueOf([1]*float32{float32Ptr(float32('A'))}), true, false, "错误"},
		{"Float64PtrArr", MinFunc("100,错误"), reflect.ValueOf([1]*float64{float64Ptr(float64('A'))}), true, false, "错误"},

		{"RunePtrSlice", MinFunc("100,错误"), reflect.ValueOf([]*rune{runePtr('A')}), true, false, "错误"},
		{"BytePtrSlice", MinFunc("100,错误"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), true, false, "错误"},
		{"Int8PtrSlice", MinFunc("100,错误"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), true, false, "错误"},
		{"IntPtrSlice", MinFunc("100,错误"), reflect.ValueOf([]*int{intPtr(int('A'))}), true, false, "错误"},
		{"Int16PtrSlice", MinFunc("100,错误"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), true, false, "错误"},
		{"Int32PtrSlice", MinFunc("100,错误"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), true, false, "错误"},
		{"Int64PtrSlice", MinFunc("100,错误"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), true, false, "错误"},
		{"Float32PtrSlice", MinFunc("100,错误"), reflect.ValueOf([]*float32{float32Ptr(float32('A'))}), true, false, "错误"},
		{"Float64PtrSlice", MinFunc("100,错误"), reflect.ValueOf([]*float64{float64Ptr(float64('A'))}), true, false, "错误"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestMinUnPassAndSpecialEffectMsg(t *testing.T) {
	for _, s := range []testCase{
		{"Rune", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf('A'), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Byte", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(byte('A')), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int8", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(int8('A')), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(int('A')), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int16", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(int16('A')), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int32", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(int32(byte('A'))), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int64", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(int64('A')), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float32", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(float32('A')), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float64", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(float64('A')), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},

		{"RuneArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]rune{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"ByteArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]byte{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int8Arr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]int8{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"IntArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]int{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int16Arr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]int16{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int32Arr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]int32{int32(byte('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int64Arr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]int64{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float32Arr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]float32{float32('A')}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float64Arr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]float64{float64('A')}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},

		{"RuneSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]rune{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"ByteSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]byte{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int8Slice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]int8{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"IntSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]int{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int16Slice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]int16{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int32Slice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]int32{int32(byte('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int64Slice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]int64{'A'}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float32Slice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]float32{float32('A')}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float64Slice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]float64{float64('A')}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},

		{"RunePtr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(runePtr('A')), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"BytePtr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(bytePtr(byte('A'))), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int8Ptr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(int8Ptr(int8('A'))), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"IntPtr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(intPtr(int('A'))), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int16Ptr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(int16Ptr(int16('A'))), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int32Ptr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(int32Ptr(int32(byte('A')))), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int64Ptr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(int64Ptr(int64('A'))), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float32Ptr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(float32Ptr(float32('A'))), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float64Ptr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(float64Ptr(float64('A'))), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},

		{"RunePtrArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*rune{runePtr('A')}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"BytePtrArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int8PtrArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"IntPtrArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*int{intPtr(int('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int16PtrArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int32PtrArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int64PtrArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float32PtrArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*float32{float32Ptr(float32('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float64PtrArr", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*float64{float64Ptr(float64('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},

		{"RunePtrSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*rune{runePtr('A')}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"BytePtrSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int8PtrSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"IntPtrSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*int{intPtr(int('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int16PtrSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int32PtrSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Int64PtrSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float32PtrSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*float32{float32Ptr(float32('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
		{"Float64PtrSlice", MinFunc("100,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*float64{float64Ptr(float64('A'))}), true, false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestMinUnPassAndArrowMsg(t *testing.T) {
	for _, s := range []testCase{
		{"Rune", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf('A'), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Byte", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(byte('A')), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int8", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(int8('A')), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(int('A')), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int16", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(int16('A')), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int32", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(int32(byte('A'))), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int64", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(int64('A')), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float32", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(float32('A')), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float64", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(float64('A')), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},

		{"RuneArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]rune{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"ByteArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]byte{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int8Arr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]int8{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"IntArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]int{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int16Arr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]int16{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int32Arr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]int32{int32(byte('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int64Arr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]int64{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float32Arr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]float32{float32('A')}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float64Arr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]float64{float64('A')}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},

		{"RuneSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]rune{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"ByteSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]byte{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int8Slice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]int8{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"IntSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]int{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int16Slice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]int16{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int32Slice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]int32{int32(byte('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int64Slice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]int64{'A'}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float32Slice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]float32{float32('A')}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float64Slice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]float64{float64('A')}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},

		{"RunePtr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(runePtr('A')), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"BytePtr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(bytePtr(byte('A'))), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int8Ptr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(int8Ptr(int8('A'))), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"IntPtr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(intPtr(int('A'))), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int16Ptr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(int16Ptr(int16('A'))), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int32Ptr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(int32Ptr(int32(byte('A')))), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int64Ptr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(int64Ptr(int64('A'))), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float32Ptr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(float32Ptr(float32('A'))), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float64Ptr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(float64Ptr(float64('A'))), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},

		{"RunePtrArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*rune{runePtr('A')}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"BytePtrArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int8PtrArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"IntPtrArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*int{intPtr(int('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int16PtrArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int32PtrArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int64PtrArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float32PtrArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*float32{float32Ptr(float32('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float64PtrArr", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*float64{float64Ptr(float64('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},

		{"RunePtrSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*rune{runePtr('A')}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"BytePtrSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int8PtrSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"IntPtrSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*int{intPtr(int('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int16PtrSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int32PtrSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Int64PtrSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float32PtrSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*float32{float32Ptr(float32('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
		{"Float64PtrSlice", MinFunc("100,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*float64{float64Ptr(float64('A'))}), true, false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestMinUnPassAndForeignMsg(t *testing.T) {
	for _, s := range []testCase{
		{"Rune", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf('A'), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Byte", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(byte('A')), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int8", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(int8('A')), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(int('A')), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int16", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(int16('A')), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int32", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(int32(byte('A'))), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int64", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(int64('A')), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float32", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(float32('A')), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float64", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(float64('A')), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},

		{"RuneArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]rune{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"ByteArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]byte{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int8Arr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]int8{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"IntArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]int{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int16Arr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]int16{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int32Arr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]int32{int32(byte('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int64Arr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]int64{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float32Arr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]float32{float32('A')}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float64Arr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]float64{float64('A')}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},

		{"RuneSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]rune{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"ByteSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]byte{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int8Slice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]int8{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"IntSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]int{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int16Slice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]int16{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int32Slice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]int32{int32(byte('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int64Slice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]int64{'A'}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float32Slice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]float32{float32('A')}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float64Slice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]float64{float64('A')}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},

		{"RunePtr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(runePtr('A')), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"BytePtr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(bytePtr(byte('A'))), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int8Ptr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(int8Ptr(int8('A'))), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"IntPtr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(intPtr(int('A'))), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int16Ptr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(int16Ptr(int16('A'))), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int32Ptr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(int32Ptr(int32(byte('A')))), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int64Ptr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(int64Ptr(int64('A'))), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float32Ptr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(float32Ptr(float32('A'))), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float64Ptr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(float64Ptr(float64('A'))), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},

		{"RunePtrArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*rune{runePtr('A')}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"BytePtrArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int8PtrArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"IntPtrArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*int{intPtr(int('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int16PtrArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int32PtrArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int64PtrArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float32PtrArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*float32{float32Ptr(float32('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float64PtrArr", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*float64{float64Ptr(float64('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},

		{"RunePtrSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*rune{runePtr('A')}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"BytePtrSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int8PtrSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"IntPtrSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*int{intPtr(int('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int16PtrSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int32PtrSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Int64PtrSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float32PtrSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*float32{float32Ptr(float32('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
		{"Float64PtrSlice", MinFunc("100,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*float64{float64Ptr(float64('A'))}), true, false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestMinUnPassAndSymbolMsg(t *testing.T) {
	for _, s := range []testCase{
		{"Rune", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf('A'), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Byte", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(byte('A')), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int8", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(int8('A')), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(int('A')), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int16", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(int16('A')), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int32", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(int32(byte('A'))), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int64", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(int64('A')), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float32", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(float32('A')), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float64", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(float64('A')), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},

		{"RuneArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]rune{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"ByteArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]byte{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int8Arr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]int8{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"IntArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]int{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int16Arr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]int16{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int32Arr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]int32{int32(byte('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int64Arr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]int64{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float32Arr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]float32{float32('A')}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float64Arr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]float64{float64('A')}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},

		{"RuneSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]rune{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"ByteSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]byte{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int8Slice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]int8{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"IntSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]int{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int16Slice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]int16{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int32Slice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]int32{int32(byte('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int64Slice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]int64{'A'}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float32Slice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]float32{float32('A')}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float64Slice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]float64{float64('A')}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},

		{"RunePtr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(runePtr('A')), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"BytePtr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(bytePtr(byte('A'))), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int8Ptr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(int8Ptr(int8('A'))), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"IntPtr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(intPtr(int('A'))), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int16Ptr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(int16Ptr(int16('A'))), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int32Ptr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(int32Ptr(int32(byte('A')))), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int64Ptr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(int64Ptr(int64('A'))), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float32Ptr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(float32Ptr(float32('A'))), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float64Ptr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(float64Ptr(float64('A'))), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},

		{"RunePtrArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*rune{runePtr('A')}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"BytePtrArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*byte{bytePtr(byte('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int8PtrArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*int8{int8Ptr(int8('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"IntPtrArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*int{intPtr(int('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int16PtrArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*int16{int16Ptr(int16('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int32PtrArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*int32{int32Ptr(int32(byte('A')))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int64PtrArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*int64{int64Ptr(int64('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float32PtrArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*float32{float32Ptr(float32('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float64PtrArr", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*float64{float64Ptr(float64('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},

		{"RunePtrSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*rune{runePtr('A')}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"BytePtrSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*byte{bytePtr(byte('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int8PtrSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*int8{int8Ptr(int8('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"IntPtrSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*int{intPtr(int('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int16PtrSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*int16{int16Ptr(int16('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int32PtrSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*int32{int32Ptr(int32(byte('A')))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Int64PtrSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*int64{int64Ptr(int64('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float32PtrSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*float32{float32Ptr(float32('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
		{"Float64PtrSlice", MinFunc("100,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*float64{float64Ptr(float64('A'))}), true, false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestMinPass(t *testing.T) {
	for _, s := range []testCase{
		{"Rune", MinFunc("1"), reflect.ValueOf(rune(1)), true, true, ""},
		{"Byte", MinFunc("1"), reflect.ValueOf(byte(1)), true, true, ""},
		{"Int8", MinFunc("1"), reflect.ValueOf(int8(1)), true, true, ""},
		{"Int", MinFunc("1"), reflect.ValueOf(1), true, true, ""},
		{"Int16", MinFunc("1"), reflect.ValueOf(int16(1)), true, true, ""},
		{"Int32", MinFunc("1"), reflect.ValueOf(int32(1)), true, true, ""},
		{"Int64", MinFunc("1"), reflect.ValueOf(int64(1)), true, true, ""},
		{"Float32", MinFunc("1"), reflect.ValueOf(float32(1)), true, true, ""},
		{"Float64", MinFunc("1"), reflect.ValueOf(float64(1)), true, true, ""},

		{"RuneArr", MinFunc("1"), reflect.ValueOf([1]rune{1}), true, true, ""},
		{"ByteArr", MinFunc("1"), reflect.ValueOf([1]byte{1}), true, true, ""},
		{"Int8Arr", MinFunc("1"), reflect.ValueOf([1]int8{1}), true, true, ""},
		{"IntArr", MinFunc("1"), reflect.ValueOf([1]int{1}), true, true, ""},
		{"Int16Arr", MinFunc("1"), reflect.ValueOf([1]int16{1}), true, true, ""},
		{"Int32Arr", MinFunc("1"), reflect.ValueOf([1]int32{1}), true, true, ""},
		{"Int64Arr", MinFunc("1"), reflect.ValueOf([1]int64{1}), true, true, ""},
		{"Float32Arr", MinFunc("1"), reflect.ValueOf([1]float32{float32(1)}), true, true, ""},
		{"Float64Arr", MinFunc("1"), reflect.ValueOf([1]float64{float64(1)}), true, true, ""},

		{"RuneSlice", MinFunc("1"), reflect.ValueOf([]rune{rune(1)}), true, true, ""},
		{"ByteSlice", MinFunc("1"), reflect.ValueOf([]byte{byte(1)}), true, true, ""},
		{"Int8Slice", MinFunc("1"), reflect.ValueOf([]int8{int8(1)}), true, true, ""},
		{"IntSlice", MinFunc("1"), reflect.ValueOf([]int{1}), true, true, ""},
		{"Int16Slice", MinFunc("1"), reflect.ValueOf([]int16{int16(1)}), true, true, ""},
		{"Int32Slice", MinFunc("1"), reflect.ValueOf([]int32{int32(1)}), true, true, ""},
		{"Int64Slice", MinFunc("1"), reflect.ValueOf([]int64{int64(1)}), true, true, ""},
		{"Float32Slice", MinFunc("1"), reflect.ValueOf([]float32{float32(1)}), true, true, ""},
		{"Float64Slice", MinFunc("1"), reflect.ValueOf([]float64{float64(1)}), true, true, ""},

		{"RunePtr", MinFunc("1"), reflect.ValueOf(runePtr(rune(1))), true, true, ""},
		{"BytePtr", MinFunc("1"), reflect.ValueOf(bytePtr(byte(1))), true, true, ""},
		{"Int8Ptr", MinFunc("1"), reflect.ValueOf(int8Ptr(int8(1))), true, true, ""},
		{"IntPtr", MinFunc("1"), reflect.ValueOf(intPtr(1)), true, true, ""},
		{"Int16Ptr", MinFunc("1"), reflect.ValueOf(int16Ptr(int16(1))), true, true, ""},
		{"Int32Ptr", MinFunc("1"), reflect.ValueOf(int32Ptr(1)), true, true, ""},
		{"Int64Ptr", MinFunc("1"), reflect.ValueOf(int64Ptr(int64(1))), true, true, ""},
		{"Float32Ptr", MinFunc("1"), reflect.ValueOf(float32Ptr(float32(1))), true, true, ""},
		{"Float64Ptr", MinFunc("1"), reflect.ValueOf(float64Ptr(float64(1))), true, true, ""},

		{"RunePtrArr", MinFunc("1"), reflect.ValueOf([1]*rune{runePtr(rune(1))}), true, true, ""},
		{"BytePtrArr", MinFunc("1"), reflect.ValueOf([1]*byte{bytePtr(byte(1))}), true, true, ""},
		{"Int8PtrArr", MinFunc("1"), reflect.ValueOf([1]*int8{int8Ptr(int8(1))}), true, true, ""},
		{"IntPtrArr", MinFunc("1"), reflect.ValueOf([1]*int{intPtr(1)}), true, true, ""},
		{"Int16PtrArr", MinFunc("1"), reflect.ValueOf([1]*int16{int16Ptr(int16(1))}), true, true, ""},
		{"Int32PtrArr", MinFunc("1"), reflect.ValueOf([1]*int32{int32Ptr(int32(1))}), true, true, ""},
		{"Int64PtrArr", MinFunc("1"), reflect.ValueOf([1]*int64{int64Ptr(int64(1))}), true, true, ""},
		{"Float32PtrArr", MinFunc("1"), reflect.ValueOf([1]*float32{float32Ptr(float32(1))}), true, true, ""},
		{"Float64PtrArr", MinFunc("1"), reflect.ValueOf([1]*float64{float64Ptr(float64(1))}), true, true, ""},

		{"RunePtrSlice", MinFunc("1"), reflect.ValueOf([]*rune{runePtr(rune(1))}), true, true, ""},
		{"BytePtrSlice", MinFunc("1"), reflect.ValueOf([]*byte{bytePtr(byte(1))}), true, true, ""},
		{"Int8PtrSlice", MinFunc("1"), reflect.ValueOf([]*int8{int8Ptr(int8(1))}), true, true, ""},
		{"IntPtrSlice", MinFunc("1"), reflect.ValueOf([]*int{intPtr(1)}), true, true, ""},
		{"Int16PtrSlice", MinFunc("1"), reflect.ValueOf([]*int16{int16Ptr(int16(1))}), true, true, ""},
		{"Int32PtrSlice", MinFunc("1"), reflect.ValueOf([]*int32{int32Ptr(1)}), true, true, ""},
		{"Int64PtrSlice", MinFunc("1"), reflect.ValueOf([]*int64{int64Ptr(int64(1))}), true, true, ""},
		{"Float32PtrSlice", MinFunc("1"), reflect.ValueOf([]*float32{float32Ptr(float32(1))}), true, true, ""},
		{"Float64PtrSlice", MinFunc("1"), reflect.ValueOf([]*float64{float64Ptr(float64(1))}), true, true, ""},
	} {
		t.Run(s.name, s.test)
	}
}
