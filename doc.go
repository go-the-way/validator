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

// Package validator
//
// A lightweight model validator written in Go.
//
// ::usage::
//
//
// package main
//
// import (
//	"fmt"
//
//	 v "github.com/go-the-way/validator"
// )
//
// func main() {
//	result := v.New(&struct {
//		int `validate:"min(10,fail)"`
//	}{}).Validate()
//	fmt.Println(result.Passed)
//	fmt.Println(result.Messages())
// }
package validator
