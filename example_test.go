package prefixw_test

import (
	"os"

	"github.com/koron-go/prefixw"
)

func Example() {
	w := prefixw.New(os.Stdout, "PREFIX ")
	w.Write([]byte("Hello\nWorld\n"))

	// Output:
	// PREFIX Hello
	// PREFIX World
}

func ExampleWriter_Write() {
	w1 := prefixw.New(os.Stdout, "PREFIX1 ")
	w1.Write([]byte("Hello\nSufficient"))
	w1.Write([]byte("World\n"))

	w2 := prefixw.New(os.Stdout, "PREFIX2 ")
	w2.Write([]byte("Hello\nInsufficient"))
	w2.Write([]byte("World"))

	// Output:
	// PREFIX1 Hello
	// PREFIX1 SufficientWorld
	// PREFIX2 Hello
}

func ExampleWriter_Close() {
	w1 := prefixw.New(os.Stdout, "PREFIX1 ")
	w1.Write([]byte("Hello\nSufficient\nWorld"))
	w1.Close()

	w2 := prefixw.New(os.Stdout, "PREFIX2 ")
	w2.Write([]byte("Hello\nInsufficient\nWorld"))

	// Output:
	// PREFIX1 Hello
	// PREFIX1 Sufficient
	// PREFIX1 World
	// PREFIX2 Hello
	// PREFIX2 Insufficient
}
