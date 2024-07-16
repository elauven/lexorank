# LexoRank Generator Library

This Golang library provides utilities to generate LexoRank codes, which are used for ordering items in a way that allows efficient insertion of new items between existing ones. The library offers two main functions:

1. `Rank(prev string, next string) string` - Generates a single LexoRank code between two given ranks.
2. `RankN(prev string, next string, n int) []string` - Generates multiple LexoRank codes between two given ranks.

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

### Generating a Single LexoRank

Use the `Rank` function to generate a LexoRank code between two existing ranks:

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
    fmt.Println("Generated LexoRank:", rank)
}
```

### Generating Multiple LexoRanks

Use the `RankN` function to generate multiple LexoRank codes between two existing ranks:

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
    fmt.Println("Generated LexoRanks:", ranks)
}
```

## API Reference

### Rank

```go
func Rank(prev string, next string) string
```

- `prev`: The previous LexoRank code.
- `next`: The next LexoRank code.

Generates a single LexoRank code between `prev` and `next`.

### RankN

```go
func RankN(prev string, next string, n int) []string
```

- `prev`: The previous LexoRank code.
- `next`: The next LexoRank code.
- `n`: The number of LexoRank codes to generate.

Generates `n` LexoRank codes between `prev` and `next`.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## License

This library is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---

Feel free to modify the README as needed and replace placeholders like `danghh-1998` with actual values.