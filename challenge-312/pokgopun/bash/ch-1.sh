### https://theweeklychallenge.org/blog/perl-weekly-challenge-312/
<<cmmnt

Task 1: Minimum Time

Submitted by: [52]Mohammad Sajid Anwar
     __________________________________________________________________

   You are given a typewriter with lowercase english letters a to z
   arranged in a circle.

   Task 1

   Typing a character takes 1 sec. You can move pointer one character
   clockwise or anti-clockwise.

   The pointer initially points at a.

   Write a script to return minimum time it takes to print the given
   string.

Example 1

Input: $str = "abc"
Output: 5

The pointer is at 'a' initially.
1 sec - type the letter 'a'
1 sec - move pointer clockwise to 'b'
1 sec - type the letter 'b'
1 sec - move pointer clockwise to 'c'
1 sec - type the letter 'c'

Example 2

Input: $str = "bza"
Output: 7

The pointer is at 'a' initially.
1 sec - move pointer clockwise to 'b'
1 sec - type the letter 'b'
1 sec - move pointer anti-clockwise to 'a'
1 sec - move pointer anti-clockwise to 'z'
1 sec - type the letter 'z'
1 sec - move pointer clockwise to 'a'
1 sec - type the letter 'a'

Example 3

Input: $str = "zjpc"
Output: 34

Task 2: Balls and Boxes
cmmnt
### solution by pokgopun@gmail.com
