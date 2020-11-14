package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fmap := make(map[string]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, fmap, "os.Stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fmap, arg)
			f.Close()
		}
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\t%s\n", n, line, fmap[line])
	}
}

func countLines(f *os.File, counts map[string]int, fmap map[string]string, fname string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		fmap[input.Text()] = fname
	}
}
