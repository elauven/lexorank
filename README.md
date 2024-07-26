# Lexorank Generator Library

[<img src="https://github.com/danghh-1998/lexorank/actions/workflows/build.yml/badge.svg">](https://github.com/danghh-1998/lexorank/actions)
[<img src="https://pkg.go.dev/badge/danghh-1998/lexorank">](https://pkg.go.dev/github.com/danghh-1998/lexorank)
[![Go Report Card](https://goreportcard.com/badge/github.com/danghh-1998/lexorank)](https://goreportcard.com/report/github.com/danghh-1998/lexorank)

This Golang library provides utilities to generate Lexorank codes, which are used for ordering items in a way that allows efficient insertion of new items between existing ones. The library offers two main functions:

1. `Rank(prev string, next string) string` - Generates a single Lexorank code between two given ranks.
2. `RankN(prev string, next string, n int) []string` - Generates multiple Lexorank codes between two given ranks.

## Installation

To install the library, use the following command:

```sh
go get github.com/danghh-1998/lexorank
```

## Usage

Import the library in your Go code:

```go
import "github.com/danghh-1998/lexorank"
```

### Generating a Single Lexorank

Use the `Rank` function to generate a Lexorank code between two existing ranks:

```go
package main

import (
    "fmt"
    "github.com/danghh-1998/lexorank"
)

func main() {
    prev := "a"
    next := "b"
    rank := lexorank.Rank(prev, next)
    fmt.Println("Generated Lexorank:", rank)
}
```

### Generating Multiple Lexoranks

Use the `RankN` function to generate multiple Lexorank codes between two existing ranks:

```go
package main

import (
    "fmt"
    "github.com/danghh-1998/lexorank"
)

func main() {
    prev := "a"
    next := "b"
    n := 5
    ranks := lexorank.RankN(prev, next, n)
    fmt.Println("Generated Lexoranks:", ranks)
}
```

## API Reference

### Rank

```go
func Rank(prev string, next string) string
```

- `prev`: The previous Lexorank code.
- `next`: The next Lexorank code.

Generates a single Lexorank code between `prev` and `next`.

### RankN

```go
func RankN(prev string, next string, n int) []string
```

- `prev`: The previous Lexorank code.
- `next`: The next Lexorank code.
- `n`: The number of Lexorank codes to generate.

Generates `n` Lexorank codes between `prev` and `next`.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub. See the [CONTRIBUTING](CONTRIBUTING.md) file for more details.

## License

This library is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
