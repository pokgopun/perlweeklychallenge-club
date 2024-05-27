//# https://theweeklychallenge.org/blog/perl-weekly-challenge-270/
/*#

Task 1: Special Positions

Submitted by: [43]Mohammad Sajid Anwar
     __________________________________________________________________

   You are given a m x n binary matrix.

   Write a script to return the number of special positions in the given
   binary matrix.

     A position (i, j) is called special if $matrix[i][j] == 1 and all
     other elements in the row i and column j are 0.

Example 1

Input: $matrix = [ [1, 0, 0],
                   [0, 0, 1],
                   [1, 0, 0],
                 ]
Output: 1

There is only special position (1, 2) as $matrix[1][2] == 1
and all other elements in row 1 and column 2 are 0.

Example 2

Input: $matrix = [ [1, 0, 0],
                   [0, 1, 0],
                   [0, 0, 1],
                 ]
Output: 3

Special positions are (0,0), (1, 1) and (2,2).

Task 2: Distribute Elements
#*/
//# solution by pokgopun@gmail.com

package main

import (
	"io"
	"os"

	"github.com/google/go-cmp/cmp"
)

type row []int

type rows []row

type matrix struct {
	rows               rows
	rowCount, colCount int
	rowColData         map[int]bool
	rowColProcessCount int
	specialPoint       [][2]int
	specialPointCount  int
}

func newMatrix(rws rows) matrix {
	mtx := matrix{
		rows:       rws,
		rowCount:   len(rws),
		colCount:   len(rws[0]),
		rowColData: make(map[int]bool),
	}
	mtx.process()
	return mtx
}

func (mtx *matrix) process() {
	for r := range mtx.rowCount {
		for c := range mtx.colCount {
			if mtx.rows[r][c] == 1 && mtx.hasZeroCross(r, c) {
				mtx.specialPoint = append(mtx.specialPoint, [2]int{r, c})
			}
		}
	}
	mtx.specialPointCount = len(mtx.specialPoint)
}

func (mtx *matrix) hasZeroCross(r, c int) bool {
	var (
		l, k, v int
		ok      bool
	)
	for _, k = range [2]int{-r - 1, c + 1} {
		_, ok = mtx.rowColData[k]
		if !ok {
			mtx.rowColProcessCount++
			if k < 0 {
				l = mtx.colCount
				for v = range l {
					if mtx.rows[r][v] == 0 {
						l--
					}
				}
			} else {
				l = mtx.rowCount
				for v = range l {
					if mtx.rows[v][c] == 0 {
						l--
					}
				}
			}
			mtx.rowColData[k] = l == 1
		}
		if !mtx.rowColData[k] {
			return false
		}
	}
	return true
}
func main() {
	for _, data := range []struct {
		input  rows
		output int
	}{
		{
			rows{
				row{1, 0, 0},
				row{0, 0, 1},
				row{1, 0, 0},
			}, 1,
		},
		{
			rows{
				row{1, 0, 0},
				row{0, 1, 0},
				row{0, 0, 1},
			}, 3,
		},
		{
			rows{
				row{1, 1, 1},
				row{1, 1, 1},
				row{1, 1, 1},
			}, 0,
		},
		{
			rows{
				row{0, 0, 0},
				row{0, 0, 0},
				row{0, 0, 0},
			}, 0,
		},
	} {
		mtx := newMatrix(data.input)
		//fmt.Println(data.output, mtx.rows, mtx.specialPointCount, mtx.specialPoint, mtx.rowColProcessCount, mtx.rowColData)
		io.WriteString(os.Stdout, cmp.Diff(mtx.specialPointCount, data.output)) // blank if ok, otherwise show the difference
	}
}
