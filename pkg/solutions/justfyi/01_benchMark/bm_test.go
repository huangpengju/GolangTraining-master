// 仅供参考，练习testing
package main

import (
	"testing"
)

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// fmt.Sprintf("hello")
	}
}

// run this at command:
// go test -bench='.*'
