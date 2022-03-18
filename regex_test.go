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

func TestNewRegexFunc(t *testing.T) {
	if v := RegexFunc(""); v != nil {
		t.Fatal("test failed")
	}
}

func TestRegexUnPassAndEmptyMsg(t *testing.T) {
	for _, s := range []testCase{
		{"String", RegexFunc("1"), reflect.ValueOf("A"), false, ""},

		{"StringArr", RegexFunc("1"), reflect.ValueOf([1]string{"A"}), false, ""},

		{"StringSlice", RegexFunc("1"), reflect.ValueOf([]string{"A"}), false, ""},

		{"StringPtr", RegexFunc("1"), reflect.ValueOf(stringPtr("A")), false, ""},

		{"StringPtrArr", RegexFunc("1"), reflect.ValueOf([1]*string{stringPtr("A")}), false, ""},

		{"StringPtrSlice", RegexFunc("1"), reflect.ValueOf([]*string{stringPtr("A")}), false, ""},
	} {
		t.Run(s.name, s.test)
	}
}

func TestRegexUnPassAndMsg(t *testing.T) {
	for _, s := range []testCase{
		{"String", RegexFunc("1,fail"), reflect.ValueOf("A"), false, "fail"},

		{"StringArr", RegexFunc("1,fail"), reflect.ValueOf([1]string{"A"}), false, "fail"},

		{"StringSlice", RegexFunc("1,fail"), reflect.ValueOf([]string{"A"}), false, "fail"},

		{"StringPtr", RegexFunc("1,fail"), reflect.ValueOf(stringPtr("A")), false, "fail"},

		{"StringPtrArr", RegexFunc("1,fail"), reflect.ValueOf([1]*string{stringPtr("A")}), false, "fail"},

		{"StringPtrSlice", RegexFunc("1,fail"), reflect.ValueOf([]*string{stringPtr("A")}), false, "fail"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestRegexUnPassAndChineseMsg(t *testing.T) {
	for _, s := range []testCase{
		{"String", RegexFunc("1,错误"), reflect.ValueOf("A"), false, "错误"},

		{"StringArr", RegexFunc("1,错误"), reflect.ValueOf([1]string{"A"}), false, "错误"},

		{"StringSlice", RegexFunc("1,错误"), reflect.ValueOf([]string{"A"}), false, "错误"},

		{"StringPtr", RegexFunc("1,错误"), reflect.ValueOf(stringPtr("A")), false, "错误"},

		{"StringPtrArr", RegexFunc("1,错误"), reflect.ValueOf([1]*string{stringPtr("A")}), false, "错误"},

		{"StringPtrSlice", RegexFunc("1,错误"), reflect.ValueOf([]*string{stringPtr("A")}), false, "错误"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestRegexUnPassAndSpecialEffectMsg(t *testing.T) {
	for _, s := range []testCase{
		{"String", RegexFunc("1,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf("A"), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},

		{"StringArr", RegexFunc("1,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]string{"A"}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},

		{"StringSlice", RegexFunc("1,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]string{"A"}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},

		{"StringPtr", RegexFunc("1,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf(stringPtr("A")), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},

		{"StringPtrArr", RegexFunc("1,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([1]*string{stringPtr("A")}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},

		{"StringPtrSlice", RegexFunc("1,︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"), reflect.ValueOf([]*string{stringPtr("A")}), false, "︻︼︽︾〒↑↓☉⊙●〇◎¤★☆■▓「」『』◆◇▲△▼▽◣◥◢◣◤ ◥№↑↓→←↘↙Ψ※㊣∑⌒∩【】〖〗＠ξζω□∮〓※》∏卐√ ╳々♀♂∞①ㄨ≡╬╭╮╰╯╱╲ ▂ ▂ ▃ ▄ ▅ ▆ ▇ █ ▂▃▅▆█ ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestRegexUnPassAndArrowMsg(t *testing.T) {
	for _, s := range []testCase{
		{"String", RegexFunc("1,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf("A"), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},

		{"StringArr", RegexFunc("1,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]string{"A"}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},

		{"StringSlice", RegexFunc("1,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]string{"A"}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},

		{"StringPtr", RegexFunc("1,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf(stringPtr("A")), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},

		{"StringPtrArr", RegexFunc("1,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([1]*string{stringPtr("A")}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},

		{"StringPtrSlice", RegexFunc("1,↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"), reflect.ValueOf([]*string{stringPtr("A")}), false, "↑ ↓ ← → ↖ ↗ ↙ ↘ ↔ ↕ ➻ ➼ ➽ ➸ ➳ ➺ ➻ ➴ ➵ ➶ ➷ ➹▶ ➩ ➪ ➫ ➬ ➭ ➮➯ ➱ ➲ ➾ ➔ ➘ ➙ ➚ ➛ ➜➝ ➞ ➟ ➠ ➡ ➢ ➣ ➤ ➥ ➦ ➧ ➨ ↚ ↛ ↜ ↝ ↞ ↟ ↠ ↠ ↡↢ ↣ ↤ ↤ ↥ ↦ ↧ ↨ ⇄ ⇅ ⇆ ⇇ ⇈ ⇉ ⇊ ⇋ ⇌ ⇍ ⇎ ⇏ ⇐ ⇑ ⇒ ⇓⇔ ⇖ ⇗ ⇘ ⇙ ⇜ ↩ ↪ ↫ ↬ ↭ ↮ ↯ ↰ ↱ ↲ ↳ ↴ ↵ ↶ ↷ ↸ ↹☇☈ ↼ ↽ ↾ ↿ ⇀ ⇁ ⇂ ⇃ ⇞ ⇟ ⇠ ⇡ ⇢ ⇣ ⇤ ⇥ ⇦ ⇧ ⇨ ⇩ ⇪ ↺ ↻"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestRegexUnPassAndForeignMsg(t *testing.T) {
	for _, s := range []testCase{
		{"String", RegexFunc("1,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf("A"), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},

		{"StringArr", RegexFunc("1,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]string{"A"}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},

		{"StringSlice", RegexFunc("1,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]string{"A"}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},

		{"StringPtr", RegexFunc("1,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf(stringPtr("A")), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},

		{"StringPtrArr", RegexFunc("1,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([1]*string{stringPtr("A")}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},

		{"StringPtrSlice", RegexFunc("1,ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"), reflect.ValueOf([]*string{stringPtr("A")}), false, "ΑΒΓΔΕΖΗΘΙΚ∧ΜΝΞΟ∏Ρ∑ΤΥΦΧΨΩα β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ωАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯа б в г д е ё ж з и й к л м н о п р с т у ф х ц ч ш щ ъ ы ь э ю я"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestRegexUnPassAndSymbolMsg(t *testing.T) {
	for _, s := range []testCase{
		{"String", RegexFunc("1,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf("A"), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},

		{"StringArr", RegexFunc("1,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]string{"A"}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},

		{"StringSlice", RegexFunc("1,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]string{"A"}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},

		{"StringPtr", RegexFunc("1,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf(stringPtr("A")), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},

		{"StringPtrArr", RegexFunc("1,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([1]*string{stringPtr("A")}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},

		{"StringPtrSlice", RegexFunc("1,ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"), reflect.ValueOf([]*string{stringPtr("A")}), false, "ˉˇ¨‘’々～‖∶”’‘｜〃〔〕《》「」『』．〖〗【【】（）〔〕｛｝．。，、；：？！ˉˇ¨`~ 々～‖∶＂＇｀｜·… — ～ - 〃‘’“”〝〞〔〕〈〉《》「」『』〖〗【】（）［］｛｝︻︼﹄﹃＋－×÷﹢﹣±／＝ ∥∠ ≌ ∽ ≦ ≧ ≒﹤﹥ ≈ ≡ ≠ ＝ ≤ ≥ ＜ ＞ ≮ ≯"},
	} {
		t.Run(s.name, s.test)
	}
}

func TestRegexPass(t *testing.T) {
	for _, s := range []testCase{
		{"String", RegexFunc("1"), reflect.ValueOf("1"), true, ""},

		{"StringArr", RegexFunc("1"), reflect.ValueOf([1]string{"1"}), true, ""},

		{"StringSlice", RegexFunc("1"), reflect.ValueOf([]string{"1"}), true, ""},

		{"StringPtr", RegexFunc("1"), reflect.ValueOf(stringPtr("1")), true, ""},

		{"StringPtrArr", RegexFunc("1"), reflect.ValueOf([1]*string{stringPtr("1")}), true, ""},

		{"StringPtrSlice", RegexFunc("1"), reflect.ValueOf([]*string{stringPtr("1")}), true, ""},
	} {
		t.Run(s.name, s.test)
	}
}
