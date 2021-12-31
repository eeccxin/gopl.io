// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

//dup3
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			//原本
			// countLines(f, counts)
			//题目1.4
			isDup := countLines_1(f, counts)
			if isDup {
				fmt.Printf("文件%s出现重复行\n", arg)
			}

			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//题目1.4  修改 dup2 ，出现重复的行时打印文件名称
func countLines_1(f *os.File, counts map[string]int) (isDup bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		n := counts[input.Text()]
		if n > 1 {
			isDup = true
		}
	}
	return isDup
	// NOTE: ignoring potential errors from input.Err()
}

//!-
