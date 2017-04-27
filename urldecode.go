package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
)

func main() {
	help := flag.Bool("h", false, "Display this help message")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [str...]\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(os.Stderr, "  decode url %%XX from arguments or stdin to readable output\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, `
Examples:
  urldecode.exe java.lang%%3Atype%%3DThreading%%2FThreadCount java.lang%%3Atype%%3DThreading%%2FThreadCount
  urldecode.exe < file_filled_urls.txt
`)
	}
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	if flag.NArg() == 0 {
		line := bufio.NewScanner(os.Stdin)
		line.Split(bufio.ScanLines)
		for line.Scan() {
			input := line.Text()
			printPath(input)
		}
		if err := line.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "error: stdin: %s", err)
		}
	} else {
		inputs := flag.Args()
		for _, input := range inputs {
			printPath(input)
		}
	}

}

func printPath(input string) {
	path, err := url.PathUnescape(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s: %s\n", input, err)
	} else {
		fmt.Fprintf(os.Stdout, "%s\n", path)
	}
}
