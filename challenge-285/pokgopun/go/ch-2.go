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
	"os"
	"strconv"

	"github.com/google/go-cmp/cmp"
)

type money int

const (
	penny   money = 1
	nickel  money = 5
	dime    money = 10
	quarter money = 25
	half    money = 50
)

var coins = [5]money{half, quarter, dime, nickel, penny}

type pocket struct {
	a money
	c [5]int
	l int
	n int
}

func (pk pocket) coining() int {
	if pk.l == 4 {
		///* uncomment to see configurations of coins
		pk.c[4] = int(pk.a)
		fmt.Println(pk.c)
		//*/
		return 1
	}
	for n := range int(pk.a/coins[pk.l]) + 1 {
		p := pk
		p.c[p.l] = n
		p.a -= coins[p.l] * money(n)
		p.l += 1
		p.n = 0 // handle with care
		pk.n += p.coining()
	}
	return pk.n
}

func coiner(n money) int {
	return pocket{a: n}.coining()
}

func main() {
	for _, data := range []struct {
		input  money
		output int
	}{
		{9, 2},
		{15, 6},
		{100, 292},
	} {
		io.WriteString(os.Stdout, strconv.Itoa(int(data.input))+", "+strconv.Itoa(int(data.output))+"\n")
		io.WriteString(os.Stdout, cmp.Diff(coiner(data.input), data.output)) // blank if ok, otherwise show the difference
	}
}
