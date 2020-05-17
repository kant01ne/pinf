package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()

	fn := flag.Arg(0)

	lines := make(map[string]bool)

	if fn != "" {
		// read the whole file into a map if it exists
		r, err := os.Open(fn)
		if err == nil {
			sc := bufio.NewScanner(r)

			for sc.Scan() {
				lines[sc.Text()] = true
			}
			r.Close()
		}
	}

	// read the lines, Print them If Not Found (pinf)
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		line := sc.Text()
		if lines[line] {
			continue
		}

		// add the line to the map so we don't get any duplicates from stdin
		lines[line] = true
		fmt.Println(line)
	}
}