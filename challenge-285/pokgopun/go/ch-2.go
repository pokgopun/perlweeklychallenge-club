//# https://theweeklychallenge.org/blog/perl-weekly-challenge-285/
/*#

Task 2: Making Change

Submitted by: [50]David Ferrone
     __________________________________________________________________

   Compute the number of ways to make change for given amount in cents. By
   using the coins e.g. Penny, Nickel, Dime, Quarter and Half-dollar, in
   how many distinct ways can the total value equal to the given amount?
   Order of coin selection does not matter.
A penny (P) is equal to 1 cent.
A nickel (N) is equal to 5 cents.
A dime (D) is equal to 10 cents.
A quarter (Q) is equal to 25 cents.
A half-dollar (HD) is equal to 50 cents.

Example 1

Input: $amount = 9
Ouput: 2

1: 9P
2: N + 4P

Example 2

Input: $amount = 15
Ouput: 6

1: D + 5P
2: D + N
3: 3N
4: 2N + 5P
5: N + 10P
6: 15P

Example 3

Input: $amount = 100
Ouput: 292
     __________________________________________________________________

   Last date to submit the solution 23:59 (UK Time) Sunday 8th September
   2024.
     __________________________________________________________________

SO WHAT DO YOU THINK ?
#*/
//# solution by pokgopun@gmail.com

package main

import (
	"fmt"
	"io"
	"iter"
	"os"

	"github.com/google/go-cmp/cmp"
)

type Coiner struct {
	ch       chan [5]uint
	cns, cnt [5]uint
	cni, val uint
	vbs      bool
	n        *uint
	done     chan struct{}
	//done1    chan struct{}
}

func NewCoiner() Coiner {
	var n uint
	return Coiner{cns: [5]uint{50, 25, 10, 5, 1}, n: &n}
}

func (cn Coiner) Count(val uint) uint {
	cn.val = val
	*cn.n = 0
	cn.process()
	return *cn.n
}

func (cn Coiner) List(val uint) iter.Seq[[5]uint] {
	cn.val = val
	*cn.n = 0
	cn.ch = make(chan [5]uint)
	cn.done = make(chan struct{})
	//cn.done1 = make(chan struct{})
	cn.vbs = true
	go func() {
		cn.process()
		close(cn.ch)
		//fmt.Println("go routine completed")
		//exec.Command("touch", strconv.Itoa(int(*cn.n))).Run()
		//close(cn.done1) // enable this to allow yield to wait for go routine to finish before exit

	}()
	return func(yield func([5]uint) bool) {
		for v := range cn.ch {
			if !yield(v) {
				close(cn.done) // clean-up the recursive function (i.e. Coiner.process())
				//<-cn.done1     // enable this to allow yield to wait for go routine to finish before exit
				//time.Sleep(1 * time.Millisecond)
				//fmt.Println("count after partial list:", *cn.n)
				return
			}
		}
		//fmt.Println("count after complete list:", *cn.n)
	}
}

func (cn Coiner) process() {
	if cn.cni == 4 {
		*cn.n += 1
		if cn.vbs {
			cn.cnt[4] = cn.val
			select {
			case cn.ch <- cn.cnt:
			case <-cn.done:
			}
		}
	} else {
		for n := range cn.val/cn.cns[cn.cni] + 1 {
			c := cn // make a copy of coiner
			c.cnt[c.cni] = n
			c.val -= c.cns[c.cni] * n
			c.cni += 1
			c.process()
		}
	}
}

func main() {
	cn := NewCoiner()

	for _, data := range []struct {
		input, output uint
	}{
		{9, 2},
		{15, 6},
		{100, 292},
	} {
		io.WriteString(os.Stdout, cmp.Diff(cn.Count(data.input), data.output)) // blank if ok, otherwise show the difference
	}
	for v := range cn.List(15) {
		fmt.Println(v)
	}
}
