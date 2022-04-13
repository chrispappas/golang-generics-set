# Golang Generics: Set

A golang 1.18+ implementation of Set using Go generics

## Installation

```bash
$ go get -u github.com/chrispappas/golang-generics-set
```

## Quick Start

```go
package main

import (
	"fmt"

	"github.com/chrispappas/golang-generics-set/set"
)

func main() {
	testSet := set.FromSlice([]int{1, 2, 3, 1, 2, 3})
	fmt.Println(testSet.Values())
}
```

## License

golang-generics-set is released under the [MIT License](https://opensource.org/licenses/MIT).
