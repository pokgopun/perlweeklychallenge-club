### https://theweeklychallenge.org/blog/perl-weekly-challenge-356/
"""

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
"""
### solution by pokgopun@gmail.com

def ksGen(n: int) -> int:
    lst = [1,2,2]
    i = 0
    while i < 3:
        v = lst.pop(0)
        yield v
        i += 1
        if i == n:
            return
    t = i
    while True:
        l = v
        e = 2 - (t % 2)
        for j in range(l):
            lst.append(e)
            yield e
            i += 1
            if i == n:
                return
        t += 1
        #print("lst =", lst, "t =", t, "i =", i)
        v = lst.pop(0)

def ks(n: int) -> int:
    c = 0
    for v in ksGen(n):
        if v == 1:
            c += 1
    return c

import unittest

class TestKs(unittest.TestCase):
    def test(self):
        for inpt, otpt in {
                4: 2,
                5: 3,
                6: 3,
                7: 4,
                8: 4,
                }.items():
            self.assertEqual(ks(inpt),otpt)

unittest.main()
