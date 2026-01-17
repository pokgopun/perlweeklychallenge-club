//# https://theweeklychallenge.org/blog/perl-weekly-challenge-356/
/*#

Task 1: Kolakoski Sequence

Submitted by: [51]Mohammad Sajid Anwar
     __________________________________________________________________

   You are given an integer, $int > 3.

   Write a script to generate the Kolakoski Sequence of given length $int
   and return the count of 1 in the generated sequence. Please follow the
   [52]wikipedia page for more informations.

Example 1

Input: $int = 4
Output: 2

(1)(22)(11)(2) => 1221

Example 2

Input: $int = 5
Output: 3

(1)(22)(11)(2)(1) => 12211

Example 3

Input: $int = 6
Output: 3

(1)(22)(11)(2)(1)(22) => 122112

Example 4

Input: $int = 7
Output: 4

(1)(22)(11)(2)(1)(22)(1) => 1221121

Example 5

Input: $int = 8
Output: 4

(1)(22)(11)(2)(1)(22)(1)(22) => 12211212

Task 2: Who Wins
#*/
//# solution by pokgopun@gmail.com

package main

import (
	"io"
	"iter"
	"os"

	"github.com/google/go-cmp/cmp"
)

func kSeq(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		s := []int{1, 2, 2}
		var i, v, t, l, e int
		for i < 3 {
			v = s[0]
			if !yield(v) {
				return
			}
			i++
			if i == n {
				return
			}
			s = s[1:]
		}
		t = i
		for i < n {
			l = v
			e = 2 - (t % 2)
			for l > 0 {
				if !yield(e) {
					return
				}
				l--
				i++
				if i == n {
					return
				}
				s = append(s, e)
			}
			t++
			//fmt.Println("s =", s, "t =", t, "i =", i)
			v = s[0]
			s = s[1:]
		}
	}
}

func ks(n int) int {
	//fmt.Println("n =", n)
	c := 0
	for v := range kSeq(n) {
		if v == 1 {
			c++
		}
	}
	return c
}

/*
	func ks(n int) int {
		s := []int{1, 2, 2}
		t := len(s)
		c := 1
		m := n - t
		for m > 0 {
			l := s[t-1]
			e := 2 - (t % 2)
			for l > 0 {
				s = append(s, e)
				l--
				if e == 1 {
					c++
				}
				m--
				if m == 0 {
					return c
				}
			}
			t++
		}
		return c
	}
*/
func main() {
	for _, data := range []struct {
		input, output int
	}{
		{4, 2},
		{5, 3},
		{6, 3},
		{7, 4},
		{8, 4},
	} {
		io.WriteString(os.Stdout, cmp.Diff(ks(data.input), data.output)) // blank if ok, otherwise show the difference
	}
}
