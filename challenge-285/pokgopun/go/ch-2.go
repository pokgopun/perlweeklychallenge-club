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

// struct to be processed by a recursive method to yield combinations of coin settings for given amount of money
type Coiner struct {
	ch       chan [5]uint  // a channel to received a result from the recursive method
	cns, cnt [5]uint       // coin values and their counts, five of them for half-dollar, quarter, dime, nickel, penny
	cni      uint          // current coin index to allow the recursive method to progress through the coins
	val      uint          // value of money given to exchange for the coins in unit of penny
	vbs      bool          // option to enable the recursive method to output coin combinations addition to the configuration counts
	n        *uint         // counts of all possible coin combinations for a given amount of money, to be updated by recursive method
	done     chan struct{} // a channel to release the recursive method from writing result for possible coin combinations
	//done1    chan struct{} // temporary channel to allow range-over-func to wait for completion of goroutine before exit
}

// intialize with combination counts with zero and coin values for half-dollar, quarter, dime, nickel and penny
func NewCoiner() Coiner {
	var n uint
	return Coiner{cns: [5]uint{50, 25, 10, 5, 1}, n: &n}
}

// caculate counts of all possible coin combinations (cn.n) for given amount of money (val), it uses "process" method
func (cn Coiner) Count(val uint) uint {
	cn.val = val
	*cn.n = 0
	cn.process()
	return *cn.n
}

// generate all possible coin combinations for given amount of money, it uses the "progress" method with the struct field "vbs" toggled
func (cn Coiner) List(val uint) iter.Seq[[5]uint] { // use iter.Seq output the result from channel, take care of goroutine if early exit
	cn.val = val
	*cn.n = 0
	cn.ch = make(chan [5]uint)
	cn.done = make(chan struct{})
	//cn.done1 = make(chan struct{})
	cn.vbs = true
	return func(yield func([5]uint) bool) {
		go func() {
			cn.process()
			close(cn.ch)
			//fmt.Println("go routine completed")
			//exec.Command("touch", strconv.Itoa(int(*cn.n))).Run()
			//close(cn.done1) // enable this to allow yield to wait for go routine to finish before exit

		}()
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

// recursive method to generate all possible coin combinations, take the struct to use its fields for processing and store results
func (cn Coiner) process() {
	if cn.cni == 4 { // if recurive through coin-index-4 (i.e. the fith coin "penny") which is the smallest unit, output the result
		*cn.n += 1  // update count of coin combiations
		if cn.vbs { // struct field to toggle outputing coin combination
			cn.cnt[4] = cn.val // number of penny coin is just the remaining amount after exchange for all other coins
			select {           // allow the recursive function to be released if iter.Seq exit early
			case cn.ch <- cn.cnt:
			case <-cn.done:
			}
		}
	} else {
		for n := range cn.val/cn.cns[cn.cni] + 1 { // loop all possible number of current coin and recursive for next coin
			c := cn                   // make a copy of coiner for recursive the next coin
			c.cnt[c.cni] = n          // add counts for the current coin
			c.val -= c.cns[c.cni] * n // deduct amount of money in exchanging for the number of current coin
			c.cni += 1                // increase coin index to allow the next recursive to process the next coin in line
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
