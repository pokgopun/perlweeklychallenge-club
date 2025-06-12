//# https://theweeklychallenge.org/blog/perl-weekly-challenge-325/
/*#

Task 1: Consecutive One

Submitted by: [43]Mohammad Sajid Anwar
     __________________________________________________________________

   You are given a binary array containing only 0 or/and 1.

   Write a script to find out the maximum consecutive 1 in the given
   array.

Example 1

Input: @binary = (0, 1, 1, 0, 1, 1, 1)
Output: 3

Example 2

Input: @binary = (0, 0, 0, 0)
Output: 0

Example 3

Input: @binary = (1, 0, 1, 0, 1, 1)
Output: 2

Task 2: Final Price
#*/
//# solution by pokgopun@gmail.com

package main

import (
	"io"
	"iter"
	"os"
	"slices"

	"github.com/google/go-cmp/cmp"
)

func co(ints iter.Seq[int]) int {
	var c, mx, v int
	for v = range ints {
		if v == 0 {
			if c == 0 {
				continue
			}
			if c > mx {
				//fmt.Printf("mx = %d\n", c)
				mx = c
			}
			c = 0
		} else {
			c++
		}
	}
	if v > 0 {
		//fmt.Println("final max update")
		if c > mx {
			//fmt.Printf("mx = %d\n", c)
			mx = c
		}
	}
	return mx
}

func main() {
	for _, data := range []struct {
		input  iter.Seq[int]
		output int
	}{
		{slices.Values([]int{0, 1, 1, 0, 1, 1, 1}), 3},
		{slices.Values([]int{0, 0, 0, 0}), 0},
		{slices.Values([]int{1, 0, 1, 0, 1, 1}), 2},
		{slices.Values([]int{1, 0, 1, 1, 0, 1, 1}), 2},
		{slices.Values([]int{1, 0, 1, 1, 0, 0, 1, 1}), 2},
		{slices.Values([]int{0, 1, 1, 1, 0, 1, 1, 1}), 3},
		{slices.Values([]int{1, 1, 1, 1, 0, 1, 1, 1, 0}), 4},
		{slices.Values([]int{0, 1, 1, 0, 0, 1, 1, 1, 0}), 3},
	} {
		//fmt.Println(data)
		io.WriteString(os.Stdout, cmp.Diff(co(data.input), data.output)) // blank if ok, otherwise show the difference
	}
}
