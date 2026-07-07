# koron-go/prefixw - prefix writer

[![GoDoc](https://godoc.org/github.com/koron-go/prefixw?status.svg)](https://godoc.org/github.com/koron-go/prefixw)
[![Actions/Go](https://github.com/koron-go/prefixw/workflows/Go/badge.svg)](https://github.com/koron-go/prefixw/actions?query=workflow%3AGo)

A `io.Writer` appends prefix for each lines.

```go
import "github.com/koron-go/prefixw"
import "os"

func main() {
    w := prefixw.New(os.Stdout, "[PREFIX] ")
    w.Write([]byte("Hello\nWorld\n"))

    // Output:
    // [PREFIX] Hello
    // [PREFIX] World
}
```
