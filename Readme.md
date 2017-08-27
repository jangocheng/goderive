# goderive

[![Build Status](https://travis-ci.org/awalterschulze/goderive.svg?branch=master)](https://travis-ci.org/awalterschulze/goderive)
[![Go Report Card](https://goreportcard.com/badge/github.com/awalterschulze/goderive)](https://goreportcard.com/report/github.com/awalterschulze/goderive)
![cover.run go](https://cover.run/go/github.com/awalterschulze/goderive/test/normal.svg)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square.svg)](https://godoc.org/github.com/awalterschulze/goderive)

`goderive` derives mundane golang functions that you do not want to maintain and keeps them up to date.

It does thing by parsing your go code for functions which are not implemented and then generates these functions for you by deriving their implementations from the parameter types. 

## Examples

In the following code the `deriveEqual` function will be spotted as a function that was not implemented (or was previously derived) and has a prefix `deriveEqual`.

```go
package main

type MyStruct struct {
	Int64     int64
	StringPtr *string
}

func (this *MyStruct) Equal(that *MyStruct) bool {
	return deriveEqual(this, that)
}
```

goderive will then generate the following code in a `derived.gen.go` file in the same package:

```go
func deriveEqual(this, that *MyStruct) bool {
	return (this == nil && that == nil) ||
		this != nil && that != nil &&
			this.Int64 == that.Int64 &&
			((this.StringPtr == nil && that.StringPtr == nil) || 
        (this.StringPtr != nil && that.StringPtr != nil && *(this.StringPtr) == *(that.StringPtr)))
}
```

Recursive Examples:

  - [Equal](https://github.com/awalterschulze/goderive/tree/master/example/plugin/equal)
  - [Compare](https://github.com/awalterschulze/goderive/tree/master/example/plugin/compare)
  - [DeepCopy](https://github.com/awalterschulze/goderive/tree/master/example/plugin/deepcopy)
  - [GoString](https://github.com/awalterschulze/goderive/tree/master/example/plugin/gostring)

Set Examples:

  - [Keys](https://github.com/awalterschulze/goderive/tree/master/example/plugin/keys)
  - [Sort](https://github.com/awalterschulze/goderive/tree/master/example/plugin/sort)
  - [Unique](https://github.com/awalterschulze/goderive/tree/master/example/plugin/unique)
  - [Set](https://github.com/awalterschulze/goderive/tree/master/example/plugin/set)
  - [Min](https://github.com/awalterschulze/goderive/tree/master/example/plugin/min)
  - [Max](https://github.com/awalterschulze/goderive/tree/master/example/plugin/max)
  - [Contains](https://github.com/awalterschulze/goderive/tree/master/example/plugin/contains)
  - [Intersect](https://github.com/awalterschulze/goderive/tree/master/example/plugin/intersect)
  - [Union](https://github.com/awalterschulze/goderive/tree/master/example/plugin/union)

Concurrency Examples:

  - [Pipeline](https://github.com/awalterschulze/goderive/tree/master/example/plugin/pipeline)
  - [Do](https://github.com/awalterschulze/goderive/tree/master/example/plugin/do)

## Functions

Recursive Functions:

  - [Equal](http://godoc.org/github.com/awalterschulze/goderive/plugin/equal) `deriveEqual(T, T) bool`
  - [Compare](http://godoc.org/github.com/awalterschulze/goderive/plugin/compare) `deriveCompare(T, T) int`
  - [DeepCopy](http://godoc.org/github.com/awalterschulze/goderive/plugin/deepcopy) 
    - `deriveDeepCopy(dst *T, src *T)`
    - `deriveDeepCopy(dst []T, src []T)`
    - `deriveDeepCopy(dst map[A]B, src map[A]B)`
  - [GoString](http://godoc.org/github.com/awalterschulze/goderive/plugin/gostring) `deriveGoString(T) string` 

Set Functions:

  - [Keys](http://godoc.org/github.com/awalterschulze/goderive/plugin/keys) `deriveKeys(map[K]V) []K`
  - [Sort](http://godoc.org/github.com/awalterschulze/goderive/plugin/sort) `deriveSort([]T) []T`
  - [Unique](http://godoc.org/github.com/awalterschulze/goderive/plugin/unique) `deriveUnique([]T) []T`
  - [Set](http://godoc.org/github.com/awalterschulze/goderive/plugin/set) `deriveSet([]T) map[T]struct{}`
  - [Min](http://godoc.org/github.com/awalterschulze/goderive/plugin/min) 
    - `deriveMin(list []T, default T) (min T)`
    - `deriveMin(T, T) T`
  - [Max](http://godoc.org/github.com/awalterschulze/goderive/plugin/max) 
    - `deriveMax(list []T, default T) (max T)`
    - `deriveMax(T, T) T`
  - [Contains](http://godoc.org/github.com/awalterschulze/goderive/plugin/contains) `deriveContains([]T, T) bool`
  - [Intersect](http://godoc.org/github.com/awalterschulze/goderive/plugin/intersect) 
    - `deriveIntersect(a, b []T) []T`
    - `deriveIntersect(a, b map[T]struct{}) map[T]struct{}`
  - [Union](http://godoc.org/github.com/awalterschulze/goderive/plugin/union) 
    - `deriveUnion(a, b []T) []T`
    - `deriveUnion(a, b map[T]struct{}) map[T]struct{}`

Functional Functions:

  - [Fmap](http://godoc.org/github.com/awalterschulze/goderive/plugin/fmap) 
    - `deriveFmap(func(A) B, []A) []B`
    - `deriveFmap(func(rune) B, string) []B` 
    - `deriveFmap(func(A) B, func() (A, error)) (B, error)`
    - `deriveFmap(func(A) (B, error), func() (A, error)) (func() (B, error), error)`
    - `deriveFmap(func(A), func() (A, error)) error`
    - `deriveFmap(func(A) (B, c, d, ...), func() (A, error)) (func() (B, c, d, ...), error)`
  - [Join](http://godoc.org/github.com/awalterschulze/goderive/plugin/join) 
    - `deriveJoin([][]T) []T`
    - `deriveJoin([]string) string`
    - `deriveJoin(func() (T, error), error) func() (T, error)`
    - `deriveJoin(func() (T, ..., error), error) func() (T, ..., error)`
  - [Filter](http://godoc.org/github.com/awalterschulze/goderive/plugin/filter) `deriveFilter(pred func(T) bool, []T) []T`
  - [All](http://godoc.org/github.com/awalterschulze/goderive/plugin/all) `deriveAll(pred func(T) bool, []T) bool`
  - [Any](http://godoc.org/github.com/awalterschulze/goderive/plugin/any) `deriveAny(pred func(T) bool, []T) bool`
  - [TakeWhile](http://godoc.org/github.com/awalterschulze/goderive/plugin/takewhile) `deriveTakeWhile(pred func(T) bool, []T) []T`
  - [Flip](http://godoc.org/github.com/awalterschulze/goderive/plugin/flip) `deriveFlip(f func(A, B, ...) T) func(B, A, ...) T`
  - [Curry](http://godoc.org/github.com/awalterschulze/goderive/plugin/curry) `deriveCurry(f func(A, B, ...) T) func(A) func(B, ...) T`
  - [Uncurry](http://godoc.org/github.com/awalterschulze/goderive/plugin/uncurry) `deriveUncurry(f func(A) func(B, ...) T) func(A, B, ...) T`
  - [Tuple](http://godoc.org/github.com/awalterschulze/goderive/plugin/tuple) `deriveTuple(A, B, ...) func() (A, B, ...)`
  - [Compose](http://godoc.org/github.com/awalterschulze/goderive/plugin/compose) 
    - `deriveCompose(func() (A, error), func(A) (B, error)) (B, error)`
    - `deriveCompose(func(A) (B, error), func(B) (C, error)) func(A) (C, error)`
    - `deriveCompose(func(A...) (B..., error), func(B...) (C..., error)) func(A...) (C..., error)`

Concurrency Functions:
  - [Fmap](http://godoc.org/github.com/awalterschulze/goderive/plugin/fmap)
    - `deriveFmap(func(A) B, <-chan A) <-chan B`
  - [Join](http://godoc.org/github.com/awalterschulze/goderive/plugin/join)
    - `deriveJoin(<-chan <-chan T) <-chan T`
    - `deriveJoin(chan <-chan T) <-chan T`
    - `deriveJoin([]<-chan T) <-chan T`
    - `deriveJoin([]chan T) <-chan T`
    - `deriveJoin(chan T, chan T, ...) <-chan T`
  - [Pipeline](http://godoc.org/github.com/awalterschulze/goderive/plugin/pipeline)
    - `derivePipeline(func(A) <-chan B, func(B) <-chan C) func(A) <-chan C`
  - [Do](http://godoc.org/github.com/awalterschulze/goderive/plugin/do)
    - `deriveDo(func() (A, error), func (B, error)) (A, B, error)`

When goderive walks over your code it is looking for a function that:
  - was not implemented (or was previously derived) and
  - has a predefined prefix.

Functions which have been previously derived will be regenerated to keep them up to date with the latest modifications to your types.  This keeps these functions, which are truly mundane to write, maintainable.

For example when someone in your team adds a new field to a struct and forgets to update the CopyTo method.  This problem is solved by goderive, by generating generated functions given the new types.

Function prefixes are by default `deriveCamelCaseFunctionName`, for example `deriveEqual`.
These are customizable using command line flags.

Let `goderive` edit your function names in your source code, by enabling `autoname` and `dedup` using the command line flags.
These flags respectively makes sure than your functions have unique names and that you don't generate multiple functions that do the same thing.

## How to run

goderive can be run from the command line:

`goderive ./...`

, using the same path semantics as the go tool.

[You can also run goderive using go generate](https://github.com/awalterschulze/goderive/blob/master/example/gogenerate/example.go) 

[And you can customize specific function prefixes](https://github.com/awalterschulze/goderive/blob/master/example/pluginprefix/Makefile)

[Or you can customize all function prefixes](https://github.com/awalterschulze/goderive/blob/master/example/prefix/Makefile)

You can let goderive rename your functions using the `-autoname` and `-dedup` flags.
If these flags are not used, goderive will not touch your code and rather return an error.

## Customization

The derive package allows you to create your own code generator plugins, see all the current plugins for examples.

You can also create your own vanity binary.
Including your own generators and/or customization of function prefixes, etc.
This should be easy to figure out by looking at [main.go](https://github.com/awalterschulze/goderive/blob/master/main.go)

## Inspired By

  - Haskell's deriving
  - Robert Griesemer's talk [Prototype your design!](https://www.youtube.com/watch?v=vLxX3yZmw5Q)
  - [go/types](https://golang.org/pkg/go/types/) standard library

## Users

These projects use goderive:

  - [katydid](https://github.com/katydid/katydid/blob/master/relapse/ast/derived.gen.go)
