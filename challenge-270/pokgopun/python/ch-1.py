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

class Matrix:
    def __init__(self, rows: list):
        self.data = rows
        self.rowCount = len(rows)
        self.rowData = [None for e in range(self.rowCount)]
        self.rowProcessCount = 0
        self.colCount = len(rows[0])
        self.colData = [None for e in range(self.colCount)]
        self.colProcessCount = 0
        self.specialPoint = []
        self.specialPointCount = 0
        self.process()
    def process(self):
        for r in range(self.rowCount):
            for c in range(self.colCount):
                if self.data[r][c] != 1:
                    continue
                if self.rowHasSingleNoneZero(r) == False:
                    continue
                if self.colHasSingleNoneZero(c) == False:
                    continue
                self.specialPoint.append((r,c))
        self.specialPointCount = len(self.specialPoint)
    def rowHasSingleNoneZero(self, r):
        if self.rowData[r] == None:
            self.rowData[r] = self.data[r].count(0) == self.colCount - 1
            self.rowProcessCount += 1
        return self.rowData[r]
    def colHasSingleNoneZero(self, c):
        if self.colData[c] == None:
            self.colData[c] = sum(1 for r in range(self.rowCount) if self.data[r][c]!=0) == 1
            self.colProcessCount += 1
        return self.colData[c]
    def rcProcessCount(self):
        return self.rowProcessCount + self.colProcessCount

import unittest

class TestCntSP(unittest.TestCase):
    def test(self):
        for inpt, otpt in {
                (
                    (1, 0, 0),
                    (0, 0, 1),
                    (1, 0, 0),
                    ): 1,
                (
                    (1, 0, 0),
                    (0, 1, 0),
                    (0, 0, 1),
                    ): 3,
                (
                    (1, 1, 1),
                    (1, 1, 1),
                    (1, 1, 1),
                    ): 0,
                (
                    (0, 0, 0),
                    (0, 0, 0),
                    (0, 0, 0),
                    ): 0,
                (
                    (0, 0, 0),
                    (0, 1, 0),
                    (0, 0, 0),
                    ): 1,
                }.items():
            print(inpt)
            mtx = Matrix(inpt)
            print(Matrix(inpt).specialPoint)
            print(Matrix(inpt).rcProcessCount())
            self.assertEqual(mtx.specialPointCount,otpt)

unittest.main()
