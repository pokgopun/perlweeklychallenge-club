### https://theweeklychallenge.org/blog/perl-weekly-challenge-339/
=begin

Task 1: Max Diff

Submitted by: [42]Mohammad Sajid Anwar
     __________________________________________________________________

   You are given an array of integers having four or more elements.

   Write a script to find two pairs of numbers from this list (four
   numbers total) so that the difference between their products is as
   large as possible.

   In the end return the max difference.

     With Two pairs (a, b) and (c, d), the product difference is (a * b)
     - (c * d).

Example 1

Input: @ints = (5, 9, 3, 4, 6)
Output: 42

Pair 1: (9, 6)
Pair 2: (3, 4)
Product Diff: (9 * 6) - (3 * 4) => 54 - 12 => 42

Example 2

Input: @ints = (1, -2, 3, -4)
Output: 8

Pair 1: (3, 1)
Pair 2: (-2, -4)

Example 3

Input: @ints = (-3, -1, -2, -4)
Output: 10

Pair 1: (-1, -2)
Pair 2: (-3, -4)

Example 4

Input: @ints = (10, 2, 0, 5, 1)
Output: 50

Pair 1: (10, 5)
Pair 2: (0, 1)

Example 5

Input: @ints = (7, 8, 9, 10, 10)
Output: 44

Pair 1: (10, 10)
Pair 2: (7, 8)

Task 2: Peak Point
=end
### solution by pokgopun@gmail.com
