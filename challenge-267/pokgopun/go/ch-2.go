//# https://theweeklychallenge.org/blog/perl-weekly-challenge-267/
/*#

Task 2: Line Counts

Submitted by: [43]Mohammad Sajid Anwar
     __________________________________________________________________

   You are given a string, $str, and a 26-items array @widths containing
   the width of each character from a to z.

   Write a script to find out the number of lines and the width of the
   last line needed to display the given string, assuming you can only fit
   100 width units on a line.

Example 1

Input: $str = "abcdefghijklmnopqrstuvwxyz"
       @widths = (10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10
,10,10,10,10,10)
Output: (3, 60)

Line 1: abcdefghij (100 pixels)
Line 2: klmnopqrst (100 pixels)
Line 3: uvwxyz (60 pixels)

Example 2

Input: $str = "bbbcccdddaaa"
       @widths = (4,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,
10,10,10,10,10)
Output: (2, 4)

Line 1: bbbcccdddaa (98 pixels)
Line 2: a (4 pixels)
     __________________________________________________________________

   Last date to submit the solution 23:59 (UK Time) Sunday 5th May 2024.
     __________________________________________________________________

SO WHAT DO YOU THINK ?
#*/
//# solution by pokgopun@gmail.com

package main

import (
	"io"
	"os"

	"github.com/google/go-cmp/cmp"
)

const lim int = 100

type Answer struct {
	Lc, L int
}

type text string

type widths [26]int

func (ws widths) get(r rune) int {
	if r >= 97 && r <= 122 {
		r -= 32
	}
	if r >= 65 && r <= 90 {
		return ws[r-65]
	}
	return 1
}

func (txt text) lc(ws widths) Answer {
	var l, lc, w int
	lc = 1
	var r rune
	for _, r = range txt {
		w = ws.get(r)
		l += w
		if l > lim {
			lc++
			l = w
		}
	}
	return Answer{lc, l}
}

func main() {
	for _, data := range []struct {
		str text
		ws  widths
		ans Answer
	}{
		{"abcdefghijklmnopqrstuvwxyz", widths{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, Answer{3, 60}},
		{"bbbcccdddaaa", widths{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, Answer{2, 4}},
		{"กชค", widths{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, Answer{1, 3}},
		{"กชa", widths{96, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, Answer{1, 98}},
		{"กชคa", widths{96, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, Answer{1, 99}},
		{"กชคงa", widths{96, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, Answer{1, 100}},
	} {
		io.WriteString(os.Stdout, cmp.Diff(data.str.lc(data.ws), data.ans)) // blank if ok,, otherwise show the difference
	}
}
