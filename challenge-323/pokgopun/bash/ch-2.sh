### https://theweeklychallenge.org/blog/perl-weekly-challenge-323/
<<cmmnt

Task 2: Tax Amount

Submitted by: [48]Mohammad Sajid Anwar
     __________________________________________________________________

   You are given an income amount and tax brackets.

   Write a script to calculate the total tax amount.

Example 1

Input: $income = 10, @tax = ([3, 50], [7, 10], [12,25])
Output: 1.65

1st tax bracket upto  3, tax is 50%.
2nd tax bracket upto  7, tax is 10%.
3rd tax bracket upto 12, tax is 25%.

Total Tax => (3 * 50/100) + (4 * 10/100) + (3 * 25/100)
          => 1.50 + 0.40 + 0.75
          => 2.65

Example 2

Input: $income = 2, @tax = ([1, 0], [4, 25], [5,50])
Output: 0.25

Total Tax => (1 * 0/100) + (1 * 25/100)
          => 0 + 0.25
          => 0.25

Example 3

Input: $income = 0, @tax = ([2, 50])
Output: 0
     __________________________________________________________________

   Last date to submit the solution 23:59 (UK Time) Sunday 1st June 2025.
     __________________________________________________________________

SO WHAT DO YOU THINK ?
cmmnt
### solution by pokgopun@gmail.com
