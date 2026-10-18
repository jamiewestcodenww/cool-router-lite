// cool-router-lite — Platform engineer · Ruby on Rails
// 一个够用就好的小工具
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func collect() ([]string, error) {
	var out []string
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		if line := sc.Text(); line != "" {
			out = append(out, line)
		}
	}
	return out, sc.Err()
}

func main() {
	lines, err := collect()
	if err != nil {
		fmt.Fprintln(os.Stderr, "read error:", err)
		os.Exit(1)
	}
	counts := map[string]int{}
	for _, l := range lines {
		counts[l]++
	}
	keys := make([]string, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%4d  %s\n", counts[k], k)
	}
	fmt.Printf("---\n%d unique / %d lines\n", len(counts), len(lines))
}
