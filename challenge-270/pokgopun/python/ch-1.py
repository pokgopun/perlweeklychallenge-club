### https://theweeklychallenge.org/blog/perl-weekly-challenge-270/
"""

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
"""
### solution by pokgopun@gmail.com

def hasAllZero(iter):
    for v in iter:
        if v!=0:
            return False
    return True

def cntSP(mn: list):
    m = len(mn)
    n = len(mn[0])
    c = 0
    for i in range(m):
        for j in range(n):
            if mn[i][j] != 1:
                continue
            if hasAllZero((mn[i][e] for e in range(j))) == False:
                continue
            if hasAllZero((mn[i][e] for e in range(j+1,n))) == False:
                continue
            if hasAllZero((mn[e][j] for e in range(i))) == False:
                continue
            if hasAllZero((mn[e][j] for e in range(i+1,m))) == False:
                continue
            c += 1
    return c

import unittest

class TestCntSP(unittest.TestCase):
    def test(self):
        for otpt, inpt in {
                1: [ [1, 0, 0],
                   [0, 0, 1],
                   [1, 0, 0],
                 ],
                3: [ [1, 0, 0],
                   [0, 1, 0],
                   [0, 0, 1],
                 ],
        }.items():
            self.assertEqual(cntSP(inpt),otpt)

unittest.main()
