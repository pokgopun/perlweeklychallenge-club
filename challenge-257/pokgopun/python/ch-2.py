### https://theweeklychallenge.org/blog/perl-weekly-challenge-257/
"""

Task 2: Reduced Row Echelon

Submitted by: [45]Ali Moradi
     __________________________________________________________________

   Given a matrix M, check whether the matrix is in reduced row echelon
   form.

   A matrix must have the following properties to be in reduced row
   echelon form:
1. If a row does not consist entirely of zeros, then the first
   nonzero number in the row is a 1. We call this the leading 1.
2. If there are any rows that consist entirely of zeros, then
   they are grouped together at the bottom of the matrix.
3. In any two successive rows that do not consist entirely of zeros,
   the leading 1 in the lower row occurs farther to the right than
   the leading 1 in the higher row.
4. Each column that contains a leading 1 has zeros everywhere else
   in that column.

   For example:
[
   [1,0,0,1],
   [0,1,0,2],
   [0,0,1,3]
]

   The above matrix is in reduced row echelon form since the first nonzero
   number in each row is a 1, leading 1s in each successive row are
   farther to the right, and above and below each leading 1 there are only
   zeros.

   For more information check out this wikipedia [46]article.

Example 1

    Input: $M = [
                  [1, 1, 0],
                  [0, 1, 0],
                  [0, 0, 0]
                ]
    Output: 0

Example 2

    Input: $M = [
                  [0, 1,-2, 0, 1],
                  [0, 0, 0, 1, 3],
                  [0, 0, 0, 0, 0],
                  [0, 0, 0, 0, 0]
                ]
    Output: 1

Example 3

    Input: $M = [
                  [1, 0, 0, 4],
                  [0, 1, 0, 7],
                  [0, 0, 1,-1]
                ]
    Output: 1

Example 4

    Input: $M = [
                  [0, 1,-2, 0, 1],
                  [0, 0, 0, 0, 0],
                  [0, 0, 0, 1, 3],
                  [0, 0, 0, 0, 0]
                ]
    Output: 0

Example 5

    Input: $M = [
                  [0, 1, 0],
                  [1, 0, 0],
                  [0, 0, 0]
                ]
    Output: 0

Example 6

    Input: $M = [
                  [4, 0, 0, 0],
                  [0, 1, 0, 7],
                  [0, 0, 1,-1]
                ]
    Output: 0
     __________________________________________________________________

   Last date to submit the solution 23:59 (UK Time) Sunday 25th February
   2024.
     __________________________________________________________________

SO WHAT DO YOU THINK ?
"""
### solution by pokgopun@gmail.com

def rowContainsAllZeroes(row: tuple):                        ### check if a row contains all zeros
    return sum(
            x!=0 for x in row
            ) == 0

def removeBottomRowWithAllZeroes(matrix: tuple):             ### remove any bottom all-zeroes row
    l = len(matrix)
    lz = l
    while l > 0:
            l -= 1
            if rowContainsAllZeroes(matrix[l]):
                matrix = matrix[:l]
                lz -= 1
            if l != lz:                                      ### stop the removal process when all-zeroes row is no longer continously found from the bottom
                break
    return matrix

def l1pos(row: tuple):                                       ### return position of leading 1, if not found return -1
    for i in range(len(row)):
        if row[i]==1:
            return i
    return -1

def isValidL1col(matrix: tuple, col: int):                   ### check values of the specified column of a matrix if it contains only one value that is not zero
    rc = len(matrix)
    return sum(
            matrix[r][col]==0 for r in range(rc)
            ) !=  rc-1

def isRRE(matrix: tuple):                                    ### check if matrix is Reduced Row Echelon
    l1p = -1                                                 ### var to store l1 position of the previous row, start with -1 so it is always less than l1 of the first row
    matrix = removeBottomRowWithAllZeroes(matrix)            ### removed bottom all-zeroes rows from the matrix
    for row in matrix:
        l1 = l1pos(row)                                      ### find poistion of leading 1
        if l1 < 0 or l1 <= l1p or isValidL1col(matrix,l1):   ### False if l1 not found or found but not valid
            return False
        l1p = l1                                             ### store l1 of the previous row
    return True

import unittest

class TestIsRRE(unittest.TestCase):
    def test(self):
        for inpt,otpt in {
                (
                  (1, 1, 0),
                  (0, 1, 0),
                  (0, 0, 0)
                ): 0,
                (
                  (0, 1,-2, 0, 1),
                  (0, 0, 0, 1, 3),
                  (0, 0, 0, 0, 0),
                  (0, 0, 0, 0, 0)
                ): 1,
                (
                  (1, 0, 0, 4),
                  (0, 1, 0, 7),
                  (0, 0, 1,-1)
                ): 1,
                (
                  (0, 1,-2, 0, 1),
                  (0, 0, 0, 0, 0),
                  (0, 0, 0, 1, 3),
                  (0, 0, 0, 0, 0)
                ): 0,
                (
                  (0, 1, 0),
                  (1, 0, 0),
                  (0, 0, 0)
                ): 0,
                (
                  (4, 0, 0, 0),
                  (0, 1, 0, 7),
                  (0, 0, 1,-1)
                ): 0,
                }.items():
            self.assertEqual(isRRE(inpt),otpt)

unittest.main()
