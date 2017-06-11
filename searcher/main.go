package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Fprintf(os.Stderr, "usage: gf <file> <pattern>")
		os.Exit(1)
	}
	pattern := flag.Arg(0)
	filename := flag.Arg(1)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cloud not open file %s: %v", filename, err)
		os.Exit(2)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, pattern) {
			fmt.Println(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error while reading the file: %v", err)
		os.Exit(3)
	}
}
