### https://theweeklychallenge.org/blog/perl-weekly-challenge-235/
"""

Task 1: Remove One

Submitted by: [44]Mohammad S Anwar
     __________________________________________________________________

   You are given an array of integers.

   Write a script to find out if removing ONLY one integer makes it
   strictly increasing order.

Example 1

Input: @ints = (0, 2, 9, 4, 6)
Output: true

Removing ONLY 9 in the given array makes it strictly increasing order.

Example 2

Input: @ints = (5, 1, 3, 2)
Output: false

Example 3

Input: @ints = (2, 2, 3)
Output: true

Task 2: Duplicate Zeros
"""
### simply create a new list by removing an index and then check if it is sorted or not
def isRmoSorted(tup,reverse=False):
    for i in range(len(tup)):
        lst = list(tup)
        del lst[i]
        #print(lst)
        if isSorted(lst,reverse):
            return True
    return False

### function to check if a list is sorted
def isSorted(lst,reverse=False):
    pm, zm = (-1)**reverse, -1*reverse
    for i in range(zm + pm, len(lst) * pm + zm, pm):
        if lst[i-pm] > lst[i]:
            return False
    return True

for inpt,otpt in {
        (0, 2, 9, 4, 6): True,
        (5, 1, 3, 2): False,
        (2, 2, 3): True,
        }.items():
    print(isRmoSorted(inpt)==otpt)
