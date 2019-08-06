package main

import (
	// "github.com/fatih/color"
	"github.com/tahomatx/hello-golang/hellox"
	"github.com/tahomatx/hello-golang/testvvv"
	// "github.com/tahomatx/hello-golang/redis"
	"fmt"
)

func main() {
	fmt.Printf(Reverse("!oG ,olleH"))
	// color.Red("Hello, world!")
	hellox.Hello()
	testvvv.Hello()

	// redis.Exec()
}

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
