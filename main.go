package main

import (
	"fmt"
	"time"

	"github.com/s-urbaniak/worker/mapper"
)

func writer(values []string) {
	i := 0
	for {
		time.Sleep(50 * time.Millisecond)

		mapper.AddWord(values[i])
		i++
		if i == len(values) { // start over
			i = 0
		}
	}
}

func reader() {
	i := 0
	words := []string{"foo", "faa", "fuu", "bar", "baz", "booz", "bla"}
	for {
		fmt.Printf("reader reads word %q cnt %d\n", words[i], mapper.WordCnt(words[i]))
		time.Sleep(200 * time.Millisecond)

		i++
		if i == len(words) {
			i = 0
		}
	}
}

func main() {
	// write worker 1
	go writer([]string{"foo", "faa", "fuu"})
	// write worker 2
	go writer([]string{"bar", "baz", "booz", "bla"})

	// the reader worker, working on the main goroutine
	reader()
}
