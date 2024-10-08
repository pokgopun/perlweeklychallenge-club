//# https://theweeklychallenge.org/blog/perl-weekly-challenge-288/
/*#

Task 2: Contiguous Block

Submitted by: [50]Peter Campbell Smith
     __________________________________________________________________

   You are given a rectangular matrix where all the cells contain either x
   or o.

   Write a script to determine the size of the largest contiguous block.

     A contiguous block consists of elements containing the same symbol
     which share an edge (not just a corner) with other elements in the
     block, and where there is a path between any two of these elements
     that crosses only those shared edges.

Example 1

    Input: $matrix = [
                       ['x', 'x', 'x', 'x', 'o'],
                       ['x', 'o', 'o', 'o', 'o'],
                       ['x', 'o', 'o', 'o', 'o'],
                       ['x', 'x', 'x', 'o', 'o'],
                     ]
    Ouput: 11

        There is a block of 9 contiguous cells containing 'x'.
        There is a block of 11 contiguous cells containing 'o'.

Example 2

    Input: $matrix = [
                       ['x', 'x', 'x', 'x', 'x'],
                       ['x', 'o', 'o', 'o', 'o'],
                       ['x', 'x', 'x', 'x', 'o'],
                       ['x', 'o', 'o', 'o', 'o'],
                     ]
    Ouput: 11

        There is a block of 11 contiguous cells containing 'x'.
        There is a block of 9 contiguous cells containing 'o'.

Example 3

    Input: $matrix = [
                       ['x', 'x', 'x', 'o', 'o'],
                       ['o', 'o', 'o', 'x', 'x'],
                       ['o', 'x', 'x', 'o', 'o'],
                       ['o', 'o', 'o', 'x', 'x'],
                     ]
    Ouput: 7

        There is a block of 7 contiguous cells containing 'o'.
        There are two other 2-cell blocks of 'o'.
        There are three 2-cell blocks of 'x' and one 3-cell.
     __________________________________________________________________

   Last date to submit the solution 23:59 (UK Time) Sunday 29th September
   2024.
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

type row []rune

type rows []row

type position struct {
	r, c int
}

type positions []position

func (ps positions) common(p positions) positions {
	m := make(map[position]int)
	var r positions
	for _, pos := range [2]positions{ps, p} {
		for _, v := range pos {
			m[v]++
			if m[v] == 2 {
				r = append(r, v)
			}
		}
	}
	return r
}

type block struct {
	mark      rune
	positions []position
}

type blocks struct {
	largest  int
	pos2mark map[position]rune
	vals     []*block
}

func newBlocks(rws rows) blocks {
	bs := blocks{pos2mark: make(map[position]rune)}
	size := 0
	for r, rw := range rws {
		for c, v := range rw {
			size++
			bs.pos2mark[position{r, c}] = v
		}
	}
	//fmt.Println("size =", size)
	var (
		pos  position
		mark rune
	)
	for len(bs.pos2mark) > 0 {
		for pos, mark = range bs.pos2mark {
			break
		}
		delete(bs.pos2mark, pos)
		b := block{mark, []position{pos}}
		bs.vals = append(bs.vals, &b)
		//bs.largest = max(bs.largest, bs.process(&b, pos, 1))
		// replace recursive bs.process() with bs.cluster() that employ queue and size to safeguard infinite loop
		queue := []position{pos}
		cnt := 0
		for len(queue) > 0 && size > cnt {
			//fmt.Println(string(b.mark), len(bs.pos2mark), len(queue), queue, len(b.positions), b.positions)
			cnt++
			bs.cluster(&b, &queue)
		}
		bs.largest = max(bs.largest, cnt)
	}
	return bs
}

func (bs blocks) cluster(b *block, queue *[]position) {
	pos := (*queue)[0]
	*queue = (*queue)[1:]
	for _, o := range [4][2]int{
		[2]int{0, 1},
		[2]int{1, 0},
		[2]int{-1, 0},
		[2]int{0, -1},
	} {
		p := position{pos.r + o[0], pos.c + o[1]}
		if bs.pos2mark[p] == b.mark {
			delete(bs.pos2mark, p)
			b.positions = append(b.positions, p)
			*queue = append(*queue, p)
		}
	}
}

func (bs blocks) process(b *block, pos position, cnt int) int {
	for _, v := range [4]position{
		position{0, 1},
		position{1, 0},
		position{-1, 0},
		position{0, -1},
	} {
		p := position{pos.r + v.r, pos.c + v.c}
		if bs.pos2mark[p] == b.mark {
			delete(bs.pos2mark, p)
			b.positions = append(b.positions, p)
			cnt += bs.process(b, p, 1)
		}
	}
	return cnt
}

func main() {
	for _, data := range []struct {
		input  rows
		output int
	}{
		{
			rows{
				row{'x', 'x', 'x', 'x', 'o'},
				row{'x', 'o', 'o', 'o', 'o'},
				row{'x', 'o', 'o', 'o', 'o'},
				row{'x', 'x', 'x', 'o', 'o'},
			},
			11,
		},
		{
			rows{
				row{'x', 'x', 'x', 'x', 'x'},
				row{'x', 'o', 'o', 'o', 'o'},
				row{'x', 'x', 'x', 'x', 'o'},
				row{'x', 'o', 'o', 'o', 'o'},
			},
			11,
		},
		{
			rows{
				row{'x', 'x', 'x', 'o', 'o'},
				row{'o', 'o', 'o', 'x', 'x'},
				row{'o', 'x', 'x', 'o', 'o'},
				row{'o', 'o', 'o', 'x', 'x'},
			},
			7,
		},
	} {
		io.WriteString(os.Stdout, cmp.Diff(newBlocks(data.input).largest, data.output)) // blank if ok, otherwise show the difference
	}
}
