# validator

A lightweight model validator written in Go.

[![CircleCI](https://circleci.com/gh/go-the-way/validator/tree/main.svg?style=shield)](https://circleci.com/gh/go-the-way/validator/tree/main)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/go-the-way/validator)
[![codecov](https://codecov.io/gh/go-the-way/validator/branch/main/graph/badge.svg?token=8MAR3J959H)](https://codecov.io/gh/go-the-way/validator)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-the-way/validator)](https://goreportcard.com/report/github.com/go-the-way/validator)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-the-way/validator?status.svg)](https://pkg.go.dev/github.com/go-the-way/validator?tab=doc)
[![Release](https://img.shields.io/github/release/go-the-way/validator.svg?style=flat-square)](https://github.com/go-the-way/validator/releases)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#validation)

## Features
- Supports multiple language.

## quickstart

```go
package main

import (
	"fmt"
	v "github.com/go-the-way/validator"
)

func main() {
	result := v.New(&struct {
		int `validate:"min(10,fail)"`
	}{}).Validate()
	fmt.Println(result.Passed)
	fmt.Println(result.Messages())
}
```

## Custom validation implementation

```go
package main

import (
	"fmt"
	v "github.com/go-the-way/validator"
	"reflect"
)

func main() {
	v.Custom("mycustom", func(value reflect.Value) (bool, string) { return false, "mycustom validation." })
	vv := v.New(&struct {
		Name string `validate:"custom(mycustom)"`
	}{}).Validate()
	fmt.Println(vv.Passed)
	fmt.Println(vv.Messages())
}
```

## Validators

| Name         | Support                                                                         | Example                             | Description                                                                                                                      |
|--------------|---------------------------------------------------------------------------------|-------------------------------------|----------------------------------------------------------------------------------------------------------------------------------|
| Min          | `([])(*)uint{8,64}`, `([])(*)int{8,64}`, `([])(*)float{32,64}`                  | validate:"min(N,invalid)"           | `Every value` must be `>= N`                                                                                                     |
| Max          | `([])(*)uint{8,64}`, `([])(*)int{8,64}`, `([])(*)float{32,64}`                  | validate:"max(N,invalid)"           | `Every value` must be `<= N`                                                                                                     |
| Length       | `(*)string`, `(*)Array[(*)string]`, `(*)Slice[(*)string]`                       | validate:"length(N,invalid)"        | `(*)string`: `Value's Len` must be `== N`<br/>`(*)Array[(*)string]` or `(*)Slice[(*)string]`: `Every Value's Len` must be `== N` |
| ArrLength    | `(*)Array[(*)Any]`, `(*)Slice[(*)Any]`                                          | validate:"arr_length(N,invalid)"    | `(*)Array[(*)Any]` or `(*)Slice[(*)Any]`: `Array` or `Slice's Len` must be `== N`                                                |
| MinLength    | `(*)string`, `(*)Array[(*)string]`, `(*)Slice[(*)string]`                       | validate:"minlength(N,invalid)"     | `(*)string`: `Value's Len` must be `>= N`<br/>`(*)Array[(*)string]` or `(*)Slice[(*)string]`: `Every Value's Len` must be `>= N` |
| ArrMinLength | `(*)Array[(*)Any]`, `(*)Slice[(*)Any]`                                          | validate:"arr_minlength(N,invalid)" | `(*)Array[(*)Any]` or `(*)Slice[(*)Any]`: `Array` or `Slice's Len` must be `>= N`                                                |
| MaxLength    | `(*)string`, `(*)Array[(*)string]`, `(*)Slice[(*)string]`                       | validate:"maxlength(N,invalid)"     | `(*)string`: `Value's Len` must be `<= N`<br/>`(*)Array[(*)string]` or `(*)Slice[(*)string]`: `Every Value's Len` must be `<= N` |
| ArrMaxLength | `(*)Array[(*)Any]`, `(*)Slice[(*)Any]`                                          | validate:"arr_maxlength(N,invalid)" | `(*)Array[(*)Any]` or `(*)Slice[(*)Any]`: `Array` or `Slice's Len` must be `<= N`                                                |
| Enum         | `([])(*)uint{8,64}`, `([])(*)int{8,64}`, `([])(*)float{32,64}`, `([])(*)string` | validate:"enum(O,invalid)"          | `Every value` must be one of `O`                                                                                                 |
| Regex        | `([])(*)string`                                                                 | validate:"regex(RE,invalid)"        | `Every value` must be match `RE`                                                                                                 |
| Valid        | `*struct{}`                                                                     | validate:"valid(T,invalid)"         | `Value` must be not `nil`                                                                                                        |
| Custom       | `any`                                                                           | validate:"custom(CUSTOM)"           | `CUSTOM` validation                                                                                                       |
