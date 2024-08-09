# go-cartesian-product

[![Build Status](https://travis-ci.org/veqryn/go-cartesian-product.svg?branch=master)](https://travis-ci.org/schwarmco/go-cartesian-product)
[![GoDoc](https://godoc.org/github.com/veqryn/go-cartesian-product?status.svg)](https://godoc.org/github.com/schwarmco/go-cartesian-product)

a package for building [cartesian products](https://en.wikipedia.org/wiki/Cartesian_product) in golang

keep in mind, that because [how golang handles maps](https://blog.golang.org/go-maps-in-action#TOC_7.) your results will not be "in order"

## Installation

In order to start, `go get` this repository:

```
go get github.com/veqryn/go-cartesian-product
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/veqryn/go-cartesian-product"
)

func main() {

    a := []any{1,2,3}
    b := []any{"a","b","c"}

    c := cartesian.Iter(a, b)

    // receive products through channel
    for product := range c {
        fmt.Println(product)
    }

    // Unordered Output:
    // [1 c]
    // [2 c]
    // [3 c]
    // [1 a]
    // [1 b]
    // [2 a]
    // [2 b]
    // [3 a]
    // [3 b]
}
```

```go
package main

import (
    "fmt"
    "github.com/veqryn/go-cartesian-product"
)

func main() {

	a := map[string][]any{
		"integers": {1, 2, 3},
		"letters":  {"a", "b", "c"},
	}

	products := cartesian.IterMap(a)

	// receive products through channel
	for product := range products {
		fmt.Println(product)
	}

	// Unordered Output:
	// map[integers:1 letters:a]
	// map[integers:2 letters:a]
	// map[integers:3 letters:a]
	// map[integers:1 letters:b]
	// map[integers:2 letters:b]
	// map[integers:3 letters:b]
	// map[integers:1 letters:c]
	// map[integers:2 letters:c]
	// map[integers:3 letters:c]
}
```
