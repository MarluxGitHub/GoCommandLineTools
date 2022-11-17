package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// WcCmd is the wc command.

type Flags struct {
	lines bool
	bytes bool
}

func main() {

	flags := Flags{
		lines: *flag.Bool("l", false, "count lines"),
		bytes: *flag.Bool("b", false, "count bytes"),
	}

	flag.Parse()

	fmt.Println(count(os.Stdin, flags))
}

func count(r io.Reader, flags Flags) int {
	scanner := bufio.NewScanner(r)

	if !flags.lines {
		scanner.Split(bufio.ScanWords)
	}

	if flags.bytes {
		scanner.Split(bufio.ScanBytes)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}